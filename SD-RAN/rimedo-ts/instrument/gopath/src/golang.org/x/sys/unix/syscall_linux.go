// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Linux system calls.
// This file is compiled as ordinary Go code,
// but it is also input to mksyscall,
// which parses the //sys lines and generates system call stubs.
// Note that sometimes we use a lowercase //sys name and
// wrap it in our own nicer implementation.

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:12
package unix

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:12
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:12
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:12
)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:12
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:12
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:12
)

import (
	"encoding/binary"
	"strconv"
	"syscall"
	"time"
	"unsafe"
)

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:26
func Access(path string, mode uint32) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:26
	_go_fuzz_dep_.CoverTab[45944]++
											return Faccessat(AT_FDCWD, path, mode, 0)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:27
	// _ = "end of CoverTab[45944]"
}

func Chmod(path string, mode uint32) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:30
	_go_fuzz_dep_.CoverTab[45945]++
											return Fchmodat(AT_FDCWD, path, mode, 0)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:31
	// _ = "end of CoverTab[45945]"
}

func Chown(path string, uid int, gid int) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:34
	_go_fuzz_dep_.CoverTab[45946]++
											return Fchownat(AT_FDCWD, path, uid, gid, 0)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:35
	// _ = "end of CoverTab[45946]"
}

func Creat(path string, mode uint32) (fd int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:38
	_go_fuzz_dep_.CoverTab[45947]++
											return Open(path, O_CREAT|O_WRONLY|O_TRUNC, mode)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:39
	// _ = "end of CoverTab[45947]"
}

func EpollCreate(size int) (fd int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:42
	_go_fuzz_dep_.CoverTab[45948]++
											if size <= 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:43
		_go_fuzz_dep_.CoverTab[45950]++
												return -1, EINVAL
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:44
		// _ = "end of CoverTab[45950]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:45
		_go_fuzz_dep_.CoverTab[45951]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:45
		// _ = "end of CoverTab[45951]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:45
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:45
	// _ = "end of CoverTab[45948]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:45
	_go_fuzz_dep_.CoverTab[45949]++
											return EpollCreate1(0)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:46
	// _ = "end of CoverTab[45949]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:52
func FanotifyMark(fd int, flags uint, mask uint64, dirFd int, pathname string) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:52
	_go_fuzz_dep_.CoverTab[45952]++
											if pathname == "" {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:53
		_go_fuzz_dep_.CoverTab[45955]++
												return fanotifyMark(fd, flags, mask, dirFd, nil)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:54
		// _ = "end of CoverTab[45955]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:55
		_go_fuzz_dep_.CoverTab[45956]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:55
		// _ = "end of CoverTab[45956]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:55
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:55
	// _ = "end of CoverTab[45952]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:55
	_go_fuzz_dep_.CoverTab[45953]++
											p, err := BytePtrFromString(pathname)
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:57
		_go_fuzz_dep_.CoverTab[45957]++
												return err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:58
		// _ = "end of CoverTab[45957]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:59
		_go_fuzz_dep_.CoverTab[45958]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:59
		// _ = "end of CoverTab[45958]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:59
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:59
	// _ = "end of CoverTab[45953]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:59
	_go_fuzz_dep_.CoverTab[45954]++
											return fanotifyMark(fd, flags, mask, dirFd, p)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:60
	// _ = "end of CoverTab[45954]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:65
func Fchmodat(dirfd int, path string, mode uint32, flags int) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:65
	_go_fuzz_dep_.CoverTab[45959]++

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:69
	if flags&^AT_SYMLINK_NOFOLLOW != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:69
		_go_fuzz_dep_.CoverTab[45961]++
												return EINVAL
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:70
		// _ = "end of CoverTab[45961]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:71
		_go_fuzz_dep_.CoverTab[45962]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:71
		if flags&AT_SYMLINK_NOFOLLOW != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:71
			_go_fuzz_dep_.CoverTab[45963]++
													return EOPNOTSUPP
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:72
			// _ = "end of CoverTab[45963]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:73
			_go_fuzz_dep_.CoverTab[45964]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:73
			// _ = "end of CoverTab[45964]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:73
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:73
		// _ = "end of CoverTab[45962]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:73
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:73
	// _ = "end of CoverTab[45959]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:73
	_go_fuzz_dep_.CoverTab[45960]++
											return fchmodat(dirfd, path, mode)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:74
	// _ = "end of CoverTab[45960]"
}

func InotifyInit() (fd int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:77
	_go_fuzz_dep_.CoverTab[45965]++
											return InotifyInit1(0)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:78
	// _ = "end of CoverTab[45965]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:96
func Link(oldpath string, newpath string) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:96
	_go_fuzz_dep_.CoverTab[45966]++
											return Linkat(AT_FDCWD, oldpath, AT_FDCWD, newpath, 0)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:97
	// _ = "end of CoverTab[45966]"
}

func Mkdir(path string, mode uint32) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:100
	_go_fuzz_dep_.CoverTab[45967]++
											return Mkdirat(AT_FDCWD, path, mode)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:101
	// _ = "end of CoverTab[45967]"
}

func Mknod(path string, mode uint32, dev int) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:104
	_go_fuzz_dep_.CoverTab[45968]++
											return Mknodat(AT_FDCWD, path, mode, dev)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:105
	// _ = "end of CoverTab[45968]"
}

func Open(path string, mode int, perm uint32) (fd int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:108
	_go_fuzz_dep_.CoverTab[45969]++
											return openat(AT_FDCWD, path, mode|O_LARGEFILE, perm)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:109
	// _ = "end of CoverTab[45969]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:114
func Openat(dirfd int, path string, flags int, mode uint32) (fd int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:114
	_go_fuzz_dep_.CoverTab[45970]++
											return openat(dirfd, path, flags|O_LARGEFILE, mode)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:115
	// _ = "end of CoverTab[45970]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:120
func Openat2(dirfd int, path string, how *OpenHow) (fd int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:120
	_go_fuzz_dep_.CoverTab[45971]++
											return openat2(dirfd, path, how, SizeofOpenHow)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:121
	// _ = "end of CoverTab[45971]"
}

func Pipe(p []int) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:124
	_go_fuzz_dep_.CoverTab[45972]++
											return Pipe2(p, 0)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:125
	// _ = "end of CoverTab[45972]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:130
func Pipe2(p []int, flags int) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:130
	_go_fuzz_dep_.CoverTab[45973]++
											if len(p) != 2 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:131
		_go_fuzz_dep_.CoverTab[45976]++
												return EINVAL
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:132
		// _ = "end of CoverTab[45976]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:133
		_go_fuzz_dep_.CoverTab[45977]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:133
		// _ = "end of CoverTab[45977]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:133
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:133
	// _ = "end of CoverTab[45973]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:133
	_go_fuzz_dep_.CoverTab[45974]++
											var pp [2]_C_int
											err := pipe2(&pp, flags)
											if err == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:136
		_go_fuzz_dep_.CoverTab[45978]++
												p[0] = int(pp[0])
												p[1] = int(pp[1])
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:138
		// _ = "end of CoverTab[45978]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:139
		_go_fuzz_dep_.CoverTab[45979]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:139
		// _ = "end of CoverTab[45979]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:139
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:139
	// _ = "end of CoverTab[45974]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:139
	_go_fuzz_dep_.CoverTab[45975]++
											return err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:140
	// _ = "end of CoverTab[45975]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:145
func Ppoll(fds []PollFd, timeout *Timespec, sigmask *Sigset_t) (n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:145
	_go_fuzz_dep_.CoverTab[45980]++
											if len(fds) == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:146
		_go_fuzz_dep_.CoverTab[45982]++
												return ppoll(nil, 0, timeout, sigmask)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:147
		// _ = "end of CoverTab[45982]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:148
		_go_fuzz_dep_.CoverTab[45983]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:148
		// _ = "end of CoverTab[45983]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:148
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:148
	// _ = "end of CoverTab[45980]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:148
	_go_fuzz_dep_.CoverTab[45981]++
											return ppoll(&fds[0], len(fds), timeout, sigmask)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:149
	// _ = "end of CoverTab[45981]"
}

func Poll(fds []PollFd, timeout int) (n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:152
	_go_fuzz_dep_.CoverTab[45984]++
											var ts *Timespec
											if timeout >= 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:154
		_go_fuzz_dep_.CoverTab[45986]++
												ts = new(Timespec)
												*ts = NsecToTimespec(int64(timeout) * 1e6)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:156
		// _ = "end of CoverTab[45986]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:157
		_go_fuzz_dep_.CoverTab[45987]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:157
		// _ = "end of CoverTab[45987]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:157
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:157
	// _ = "end of CoverTab[45984]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:157
	_go_fuzz_dep_.CoverTab[45985]++
											return Ppoll(fds, ts, nil)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:158
	// _ = "end of CoverTab[45985]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:163
func Readlink(path string, buf []byte) (n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:163
	_go_fuzz_dep_.CoverTab[45988]++
											return Readlinkat(AT_FDCWD, path, buf)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:164
	// _ = "end of CoverTab[45988]"
}

func Rename(oldpath string, newpath string) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:167
	_go_fuzz_dep_.CoverTab[45989]++
											return Renameat(AT_FDCWD, oldpath, AT_FDCWD, newpath)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:168
	// _ = "end of CoverTab[45989]"
}

func Rmdir(path string) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:171
	_go_fuzz_dep_.CoverTab[45990]++
											return Unlinkat(AT_FDCWD, path, AT_REMOVEDIR)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:172
	// _ = "end of CoverTab[45990]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:177
func Symlink(oldpath string, newpath string) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:177
	_go_fuzz_dep_.CoverTab[45991]++
											return Symlinkat(oldpath, AT_FDCWD, newpath)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:178
	// _ = "end of CoverTab[45991]"
}

func Unlink(path string) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:181
	_go_fuzz_dep_.CoverTab[45992]++
											return Unlinkat(AT_FDCWD, path, 0)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:182
	// _ = "end of CoverTab[45992]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:187
func Utimes(path string, tv []Timeval) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:187
	_go_fuzz_dep_.CoverTab[45993]++
											if tv == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:188
		_go_fuzz_dep_.CoverTab[45997]++
												err := utimensat(AT_FDCWD, path, nil, 0)
												if err != ENOSYS {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:190
			_go_fuzz_dep_.CoverTab[45999]++
													return err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:191
			// _ = "end of CoverTab[45999]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:192
			_go_fuzz_dep_.CoverTab[46000]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:192
			// _ = "end of CoverTab[46000]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:192
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:192
		// _ = "end of CoverTab[45997]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:192
		_go_fuzz_dep_.CoverTab[45998]++
												return utimes(path, nil)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:193
		// _ = "end of CoverTab[45998]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:194
		_go_fuzz_dep_.CoverTab[46001]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:194
		// _ = "end of CoverTab[46001]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:194
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:194
	// _ = "end of CoverTab[45993]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:194
	_go_fuzz_dep_.CoverTab[45994]++
											if len(tv) != 2 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:195
		_go_fuzz_dep_.CoverTab[46002]++
												return EINVAL
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:196
		// _ = "end of CoverTab[46002]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:197
		_go_fuzz_dep_.CoverTab[46003]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:197
		// _ = "end of CoverTab[46003]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:197
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:197
	// _ = "end of CoverTab[45994]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:197
	_go_fuzz_dep_.CoverTab[45995]++
											var ts [2]Timespec
											ts[0] = NsecToTimespec(TimevalToNsec(tv[0]))
											ts[1] = NsecToTimespec(TimevalToNsec(tv[1]))
											err := utimensat(AT_FDCWD, path, (*[2]Timespec)(unsafe.Pointer(&ts[0])), 0)
											if err != ENOSYS {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:202
		_go_fuzz_dep_.CoverTab[46004]++
												return err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:203
		// _ = "end of CoverTab[46004]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:204
		_go_fuzz_dep_.CoverTab[46005]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:204
		// _ = "end of CoverTab[46005]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:204
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:204
	// _ = "end of CoverTab[45995]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:204
	_go_fuzz_dep_.CoverTab[45996]++
											return utimes(path, (*[2]Timeval)(unsafe.Pointer(&tv[0])))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:205
	// _ = "end of CoverTab[45996]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:210
func UtimesNano(path string, ts []Timespec) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:210
	_go_fuzz_dep_.CoverTab[46006]++
											return UtimesNanoAt(AT_FDCWD, path, ts, 0)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:211
	// _ = "end of CoverTab[46006]"
}

func UtimesNanoAt(dirfd int, path string, ts []Timespec, flags int) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:214
	_go_fuzz_dep_.CoverTab[46007]++
											if ts == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:215
		_go_fuzz_dep_.CoverTab[46010]++
												return utimensat(dirfd, path, nil, flags)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:216
		// _ = "end of CoverTab[46010]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:217
		_go_fuzz_dep_.CoverTab[46011]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:217
		// _ = "end of CoverTab[46011]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:217
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:217
	// _ = "end of CoverTab[46007]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:217
	_go_fuzz_dep_.CoverTab[46008]++
											if len(ts) != 2 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:218
		_go_fuzz_dep_.CoverTab[46012]++
												return EINVAL
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:219
		// _ = "end of CoverTab[46012]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:220
		_go_fuzz_dep_.CoverTab[46013]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:220
		// _ = "end of CoverTab[46013]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:220
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:220
	// _ = "end of CoverTab[46008]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:220
	_go_fuzz_dep_.CoverTab[46009]++
											return utimensat(dirfd, path, (*[2]Timespec)(unsafe.Pointer(&ts[0])), flags)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:221
	// _ = "end of CoverTab[46009]"
}

func Futimesat(dirfd int, path string, tv []Timeval) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:224
	_go_fuzz_dep_.CoverTab[46014]++
											if tv == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:225
		_go_fuzz_dep_.CoverTab[46017]++
												return futimesat(dirfd, path, nil)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:226
		// _ = "end of CoverTab[46017]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:227
		_go_fuzz_dep_.CoverTab[46018]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:227
		// _ = "end of CoverTab[46018]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:227
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:227
	// _ = "end of CoverTab[46014]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:227
	_go_fuzz_dep_.CoverTab[46015]++
											if len(tv) != 2 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:228
		_go_fuzz_dep_.CoverTab[46019]++
												return EINVAL
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:229
		// _ = "end of CoverTab[46019]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:230
		_go_fuzz_dep_.CoverTab[46020]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:230
		// _ = "end of CoverTab[46020]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:230
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:230
	// _ = "end of CoverTab[46015]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:230
	_go_fuzz_dep_.CoverTab[46016]++
											return futimesat(dirfd, path, (*[2]Timeval)(unsafe.Pointer(&tv[0])))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:231
	// _ = "end of CoverTab[46016]"
}

func Futimes(fd int, tv []Timeval) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:234
	_go_fuzz_dep_.CoverTab[46021]++

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:237
	return Utimes("/proc/self/fd/"+strconv.Itoa(fd), tv)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:237
	// _ = "end of CoverTab[46021]"
}

const ImplementsGetwd = true

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:244
func Getwd() (wd string, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:244
	_go_fuzz_dep_.CoverTab[46022]++
											var buf [PathMax]byte
											n, err := Getcwd(buf[0:])
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:247
		_go_fuzz_dep_.CoverTab[46026]++
												return "", err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:248
		// _ = "end of CoverTab[46026]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:249
		_go_fuzz_dep_.CoverTab[46027]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:249
		// _ = "end of CoverTab[46027]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:249
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:249
	// _ = "end of CoverTab[46022]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:249
	_go_fuzz_dep_.CoverTab[46023]++

											if n < 1 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:251
		_go_fuzz_dep_.CoverTab[46028]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:251
		return n > len(buf)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:251
		// _ = "end of CoverTab[46028]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:251
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:251
		_go_fuzz_dep_.CoverTab[46029]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:251
		return buf[n-1] != 0
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:251
		// _ = "end of CoverTab[46029]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:251
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:251
		_go_fuzz_dep_.CoverTab[46030]++
												return "", EINVAL
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:252
		// _ = "end of CoverTab[46030]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:253
		_go_fuzz_dep_.CoverTab[46031]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:253
		// _ = "end of CoverTab[46031]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:253
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:253
	// _ = "end of CoverTab[46023]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:253
	_go_fuzz_dep_.CoverTab[46024]++

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:257
	if buf[0] != '/' {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:257
		_go_fuzz_dep_.CoverTab[46032]++
												return "", ENOENT
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:258
		// _ = "end of CoverTab[46032]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:259
		_go_fuzz_dep_.CoverTab[46033]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:259
		// _ = "end of CoverTab[46033]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:259
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:259
	// _ = "end of CoverTab[46024]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:259
	_go_fuzz_dep_.CoverTab[46025]++

											return string(buf[0 : n-1]), nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:261
	// _ = "end of CoverTab[46025]"
}

func Getgroups() (gids []int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:264
	_go_fuzz_dep_.CoverTab[46034]++
											n, err := getgroups(0, nil)
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:266
		_go_fuzz_dep_.CoverTab[46040]++
												return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:267
		// _ = "end of CoverTab[46040]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:268
		_go_fuzz_dep_.CoverTab[46041]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:268
		// _ = "end of CoverTab[46041]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:268
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:268
	// _ = "end of CoverTab[46034]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:268
	_go_fuzz_dep_.CoverTab[46035]++
											if n == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:269
		_go_fuzz_dep_.CoverTab[46042]++
												return nil, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:270
		// _ = "end of CoverTab[46042]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:271
		_go_fuzz_dep_.CoverTab[46043]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:271
		// _ = "end of CoverTab[46043]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:271
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:271
	// _ = "end of CoverTab[46035]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:271
	_go_fuzz_dep_.CoverTab[46036]++

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:274
	if n < 0 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:274
		_go_fuzz_dep_.CoverTab[46044]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:274
		return n > 1<<20
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:274
		// _ = "end of CoverTab[46044]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:274
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:274
		_go_fuzz_dep_.CoverTab[46045]++
												return nil, EINVAL
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:275
		// _ = "end of CoverTab[46045]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:276
		_go_fuzz_dep_.CoverTab[46046]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:276
		// _ = "end of CoverTab[46046]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:276
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:276
	// _ = "end of CoverTab[46036]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:276
	_go_fuzz_dep_.CoverTab[46037]++

											a := make([]_Gid_t, n)
											n, err = getgroups(n, &a[0])
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:280
		_go_fuzz_dep_.CoverTab[46047]++
												return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:281
		// _ = "end of CoverTab[46047]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:282
		_go_fuzz_dep_.CoverTab[46048]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:282
		// _ = "end of CoverTab[46048]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:282
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:282
	// _ = "end of CoverTab[46037]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:282
	_go_fuzz_dep_.CoverTab[46038]++
											gids = make([]int, n)
											for i, v := range a[0:n] {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:284
		_go_fuzz_dep_.CoverTab[46049]++
												gids[i] = int(v)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:285
		// _ = "end of CoverTab[46049]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:286
	// _ = "end of CoverTab[46038]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:286
	_go_fuzz_dep_.CoverTab[46039]++
											return
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:287
	// _ = "end of CoverTab[46039]"
}

func Setgroups(gids []int) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:290
	_go_fuzz_dep_.CoverTab[46050]++
											if len(gids) == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:291
		_go_fuzz_dep_.CoverTab[46053]++
												return setgroups(0, nil)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:292
		// _ = "end of CoverTab[46053]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:293
		_go_fuzz_dep_.CoverTab[46054]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:293
		// _ = "end of CoverTab[46054]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:293
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:293
	// _ = "end of CoverTab[46050]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:293
	_go_fuzz_dep_.CoverTab[46051]++

											a := make([]_Gid_t, len(gids))
											for i, v := range gids {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:296
		_go_fuzz_dep_.CoverTab[46055]++
												a[i] = _Gid_t(v)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:297
		// _ = "end of CoverTab[46055]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:298
	// _ = "end of CoverTab[46051]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:298
	_go_fuzz_dep_.CoverTab[46052]++
											return setgroups(len(a), &a[0])
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:299
	// _ = "end of CoverTab[46052]"
}

type WaitStatus uint32

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:313
const (
	mask	= 0x7F
	core	= 0x80
	exited	= 0x00
	stopped	= 0x7F
	shift	= 8
)

func (w WaitStatus) Exited() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:321
	_go_fuzz_dep_.CoverTab[46056]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:321
	return w&mask == exited
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:321
	// _ = "end of CoverTab[46056]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:321
}

func (w WaitStatus) Signaled() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:323
	_go_fuzz_dep_.CoverTab[46057]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:323
	return w&mask != stopped && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:323
		_go_fuzz_dep_.CoverTab[46058]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:323
		return w&mask != exited
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:323
		// _ = "end of CoverTab[46058]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:323
	}()
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:323
	// _ = "end of CoverTab[46057]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:323
}

func (w WaitStatus) Stopped() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:325
	_go_fuzz_dep_.CoverTab[46059]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:325
	return w&0xFF == stopped
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:325
	// _ = "end of CoverTab[46059]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:325
}

func (w WaitStatus) Continued() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:327
	_go_fuzz_dep_.CoverTab[46060]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:327
	return w == 0xFFFF
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:327
	// _ = "end of CoverTab[46060]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:327
}

func (w WaitStatus) CoreDump() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:329
	_go_fuzz_dep_.CoverTab[46061]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:329
	return w.Signaled() && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:329
		_go_fuzz_dep_.CoverTab[46062]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:329
		return w&core != 0
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:329
		// _ = "end of CoverTab[46062]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:329
	}()
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:329
	// _ = "end of CoverTab[46061]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:329
}

func (w WaitStatus) ExitStatus() int {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:331
	_go_fuzz_dep_.CoverTab[46063]++
											if !w.Exited() {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:332
		_go_fuzz_dep_.CoverTab[46065]++
												return -1
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:333
		// _ = "end of CoverTab[46065]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:334
		_go_fuzz_dep_.CoverTab[46066]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:334
		// _ = "end of CoverTab[46066]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:334
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:334
	// _ = "end of CoverTab[46063]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:334
	_go_fuzz_dep_.CoverTab[46064]++
											return int(w>>shift) & 0xFF
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:335
	// _ = "end of CoverTab[46064]"
}

func (w WaitStatus) Signal() syscall.Signal {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:338
	_go_fuzz_dep_.CoverTab[46067]++
											if !w.Signaled() {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:339
		_go_fuzz_dep_.CoverTab[46069]++
												return -1
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:340
		// _ = "end of CoverTab[46069]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:341
		_go_fuzz_dep_.CoverTab[46070]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:341
		// _ = "end of CoverTab[46070]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:341
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:341
	// _ = "end of CoverTab[46067]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:341
	_go_fuzz_dep_.CoverTab[46068]++
											return syscall.Signal(w & mask)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:342
	// _ = "end of CoverTab[46068]"
}

func (w WaitStatus) StopSignal() syscall.Signal {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:345
	_go_fuzz_dep_.CoverTab[46071]++
											if !w.Stopped() {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:346
		_go_fuzz_dep_.CoverTab[46073]++
												return -1
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:347
		// _ = "end of CoverTab[46073]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:348
		_go_fuzz_dep_.CoverTab[46074]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:348
		// _ = "end of CoverTab[46074]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:348
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:348
	// _ = "end of CoverTab[46071]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:348
	_go_fuzz_dep_.CoverTab[46072]++
											return syscall.Signal(w>>shift) & 0xFF
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:349
	// _ = "end of CoverTab[46072]"
}

func (w WaitStatus) TrapCause() int {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:352
	_go_fuzz_dep_.CoverTab[46075]++
											if w.StopSignal() != SIGTRAP {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:353
		_go_fuzz_dep_.CoverTab[46077]++
												return -1
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:354
		// _ = "end of CoverTab[46077]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:355
		_go_fuzz_dep_.CoverTab[46078]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:355
		// _ = "end of CoverTab[46078]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:355
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:355
	// _ = "end of CoverTab[46075]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:355
	_go_fuzz_dep_.CoverTab[46076]++
											return int(w>>shift) >> 8
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:356
	// _ = "end of CoverTab[46076]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:361
func Wait4(pid int, wstatus *WaitStatus, options int, rusage *Rusage) (wpid int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:361
	_go_fuzz_dep_.CoverTab[46079]++
											var status _C_int
											wpid, err = wait4(pid, &status, options, rusage)
											if wstatus != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:364
		_go_fuzz_dep_.CoverTab[46081]++
												*wstatus = WaitStatus(status)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:365
		// _ = "end of CoverTab[46081]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:366
		_go_fuzz_dep_.CoverTab[46082]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:366
		// _ = "end of CoverTab[46082]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:366
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:366
	// _ = "end of CoverTab[46079]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:366
	_go_fuzz_dep_.CoverTab[46080]++
											return
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:367
	// _ = "end of CoverTab[46080]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:372
func Mkfifo(path string, mode uint32) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:372
	_go_fuzz_dep_.CoverTab[46083]++
											return Mknod(path, mode|S_IFIFO, 0)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:373
	// _ = "end of CoverTab[46083]"
}

func Mkfifoat(dirfd int, path string, mode uint32) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:376
	_go_fuzz_dep_.CoverTab[46084]++
											return Mknodat(dirfd, path, mode|S_IFIFO, 0)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:377
	// _ = "end of CoverTab[46084]"
}

func (sa *SockaddrInet4) sockaddr() (unsafe.Pointer, _Socklen, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:380
	_go_fuzz_dep_.CoverTab[46085]++
											if sa.Port < 0 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:381
		_go_fuzz_dep_.CoverTab[46087]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:381
		return sa.Port > 0xFFFF
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:381
		// _ = "end of CoverTab[46087]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:381
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:381
		_go_fuzz_dep_.CoverTab[46088]++
												return nil, 0, EINVAL
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:382
		// _ = "end of CoverTab[46088]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:383
		_go_fuzz_dep_.CoverTab[46089]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:383
		// _ = "end of CoverTab[46089]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:383
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:383
	// _ = "end of CoverTab[46085]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:383
	_go_fuzz_dep_.CoverTab[46086]++
											sa.raw.Family = AF_INET
											p := (*[2]byte)(unsafe.Pointer(&sa.raw.Port))
											p[0] = byte(sa.Port >> 8)
											p[1] = byte(sa.Port)
											sa.raw.Addr = sa.Addr
											return unsafe.Pointer(&sa.raw), SizeofSockaddrInet4, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:389
	// _ = "end of CoverTab[46086]"
}

func (sa *SockaddrInet6) sockaddr() (unsafe.Pointer, _Socklen, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:392
	_go_fuzz_dep_.CoverTab[46090]++
											if sa.Port < 0 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:393
		_go_fuzz_dep_.CoverTab[46092]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:393
		return sa.Port > 0xFFFF
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:393
		// _ = "end of CoverTab[46092]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:393
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:393
		_go_fuzz_dep_.CoverTab[46093]++
												return nil, 0, EINVAL
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:394
		// _ = "end of CoverTab[46093]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:395
		_go_fuzz_dep_.CoverTab[46094]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:395
		// _ = "end of CoverTab[46094]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:395
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:395
	// _ = "end of CoverTab[46090]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:395
	_go_fuzz_dep_.CoverTab[46091]++
											sa.raw.Family = AF_INET6
											p := (*[2]byte)(unsafe.Pointer(&sa.raw.Port))
											p[0] = byte(sa.Port >> 8)
											p[1] = byte(sa.Port)
											sa.raw.Scope_id = sa.ZoneId
											sa.raw.Addr = sa.Addr
											return unsafe.Pointer(&sa.raw), SizeofSockaddrInet6, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:402
	// _ = "end of CoverTab[46091]"
}

func (sa *SockaddrUnix) sockaddr() (unsafe.Pointer, _Socklen, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:405
	_go_fuzz_dep_.CoverTab[46095]++
											name := sa.Name
											n := len(name)
											if n >= len(sa.raw.Path) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:408
		_go_fuzz_dep_.CoverTab[46100]++
												return nil, 0, EINVAL
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:409
		// _ = "end of CoverTab[46100]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:410
		_go_fuzz_dep_.CoverTab[46101]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:410
		// _ = "end of CoverTab[46101]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:410
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:410
	// _ = "end of CoverTab[46095]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:410
	_go_fuzz_dep_.CoverTab[46096]++
											sa.raw.Family = AF_UNIX
											for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:412
		_go_fuzz_dep_.CoverTab[46102]++
												sa.raw.Path[i] = int8(name[i])
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:413
		// _ = "end of CoverTab[46102]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:414
	// _ = "end of CoverTab[46096]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:414
	_go_fuzz_dep_.CoverTab[46097]++

											sl := _Socklen(2)
											if n > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:417
		_go_fuzz_dep_.CoverTab[46103]++
												sl += _Socklen(n) + 1
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:418
		// _ = "end of CoverTab[46103]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:419
		_go_fuzz_dep_.CoverTab[46104]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:419
		// _ = "end of CoverTab[46104]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:419
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:419
	// _ = "end of CoverTab[46097]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:419
	_go_fuzz_dep_.CoverTab[46098]++
											if sa.raw.Path[0] == '@' {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:420
		_go_fuzz_dep_.CoverTab[46105]++
												sa.raw.Path[0] = 0

												sl--
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:423
		// _ = "end of CoverTab[46105]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:424
		_go_fuzz_dep_.CoverTab[46106]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:424
		// _ = "end of CoverTab[46106]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:424
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:424
	// _ = "end of CoverTab[46098]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:424
	_go_fuzz_dep_.CoverTab[46099]++

											return unsafe.Pointer(&sa.raw), sl, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:426
	// _ = "end of CoverTab[46099]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:430
type SockaddrLinklayer struct {
	Protocol	uint16
	Ifindex		int
	Hatype		uint16
	Pkttype		uint8
	Halen		uint8
	Addr		[8]byte
	raw		RawSockaddrLinklayer
}

func (sa *SockaddrLinklayer) sockaddr() (unsafe.Pointer, _Socklen, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:440
	_go_fuzz_dep_.CoverTab[46107]++
											if sa.Ifindex < 0 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:441
		_go_fuzz_dep_.CoverTab[46109]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:441
		return sa.Ifindex > 0x7fffffff
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:441
		// _ = "end of CoverTab[46109]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:441
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:441
		_go_fuzz_dep_.CoverTab[46110]++
												return nil, 0, EINVAL
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:442
		// _ = "end of CoverTab[46110]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:443
		_go_fuzz_dep_.CoverTab[46111]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:443
		// _ = "end of CoverTab[46111]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:443
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:443
	// _ = "end of CoverTab[46107]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:443
	_go_fuzz_dep_.CoverTab[46108]++
											sa.raw.Family = AF_PACKET
											sa.raw.Protocol = sa.Protocol
											sa.raw.Ifindex = int32(sa.Ifindex)
											sa.raw.Hatype = sa.Hatype
											sa.raw.Pkttype = sa.Pkttype
											sa.raw.Halen = sa.Halen
											sa.raw.Addr = sa.Addr
											return unsafe.Pointer(&sa.raw), SizeofSockaddrLinklayer, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:451
	// _ = "end of CoverTab[46108]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:455
type SockaddrNetlink struct {
	Family	uint16
	Pad	uint16
	Pid	uint32
	Groups	uint32
	raw	RawSockaddrNetlink
}

func (sa *SockaddrNetlink) sockaddr() (unsafe.Pointer, _Socklen, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:463
	_go_fuzz_dep_.CoverTab[46112]++
											sa.raw.Family = AF_NETLINK
											sa.raw.Pad = sa.Pad
											sa.raw.Pid = sa.Pid
											sa.raw.Groups = sa.Groups
											return unsafe.Pointer(&sa.raw), SizeofSockaddrNetlink, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:468
	// _ = "end of CoverTab[46112]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:473
type SockaddrHCI struct {
	Dev	uint16
	Channel	uint16
	raw	RawSockaddrHCI
}

func (sa *SockaddrHCI) sockaddr() (unsafe.Pointer, _Socklen, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:479
	_go_fuzz_dep_.CoverTab[46113]++
											sa.raw.Family = AF_BLUETOOTH
											sa.raw.Dev = sa.Dev
											sa.raw.Channel = sa.Channel
											return unsafe.Pointer(&sa.raw), SizeofSockaddrHCI, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:483
	// _ = "end of CoverTab[46113]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:488
type SockaddrL2 struct {
	PSM		uint16
	CID		uint16
	Addr		[6]uint8
	AddrType	uint8
	raw		RawSockaddrL2
}

func (sa *SockaddrL2) sockaddr() (unsafe.Pointer, _Socklen, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:496
	_go_fuzz_dep_.CoverTab[46114]++
											sa.raw.Family = AF_BLUETOOTH
											psm := (*[2]byte)(unsafe.Pointer(&sa.raw.Psm))
											psm[0] = byte(sa.PSM)
											psm[1] = byte(sa.PSM >> 8)
											for i := 0; i < len(sa.Addr); i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:501
		_go_fuzz_dep_.CoverTab[46116]++
												sa.raw.Bdaddr[i] = sa.Addr[len(sa.Addr)-1-i]
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:502
		// _ = "end of CoverTab[46116]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:503
	// _ = "end of CoverTab[46114]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:503
	_go_fuzz_dep_.CoverTab[46115]++
											cid := (*[2]byte)(unsafe.Pointer(&sa.raw.Cid))
											cid[0] = byte(sa.CID)
											cid[1] = byte(sa.CID >> 8)
											sa.raw.Bdaddr_type = sa.AddrType
											return unsafe.Pointer(&sa.raw), SizeofSockaddrL2, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:508
	// _ = "end of CoverTab[46115]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:534
type SockaddrRFCOMM struct {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:536
	Addr	[6]uint8

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:540
	Channel	uint8

	raw	RawSockaddrRFCOMM
}

func (sa *SockaddrRFCOMM) sockaddr() (unsafe.Pointer, _Socklen, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:545
	_go_fuzz_dep_.CoverTab[46117]++
											sa.raw.Family = AF_BLUETOOTH
											sa.raw.Channel = sa.Channel
											sa.raw.Bdaddr = sa.Addr
											return unsafe.Pointer(&sa.raw), SizeofSockaddrRFCOMM, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:549
	// _ = "end of CoverTab[46117]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:569
type SockaddrCAN struct {
	Ifindex	int
	RxID	uint32
	TxID	uint32
	raw	RawSockaddrCAN
}

func (sa *SockaddrCAN) sockaddr() (unsafe.Pointer, _Socklen, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:576
	_go_fuzz_dep_.CoverTab[46118]++
											if sa.Ifindex < 0 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:577
		_go_fuzz_dep_.CoverTab[46122]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:577
		return sa.Ifindex > 0x7fffffff
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:577
		// _ = "end of CoverTab[46122]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:577
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:577
		_go_fuzz_dep_.CoverTab[46123]++
												return nil, 0, EINVAL
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:578
		// _ = "end of CoverTab[46123]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:579
		_go_fuzz_dep_.CoverTab[46124]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:579
		// _ = "end of CoverTab[46124]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:579
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:579
	// _ = "end of CoverTab[46118]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:579
	_go_fuzz_dep_.CoverTab[46119]++
											sa.raw.Family = AF_CAN
											sa.raw.Ifindex = int32(sa.Ifindex)
											rx := (*[4]byte)(unsafe.Pointer(&sa.RxID))
											for i := 0; i < 4; i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:583
		_go_fuzz_dep_.CoverTab[46125]++
												sa.raw.Addr[i] = rx[i]
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:584
		// _ = "end of CoverTab[46125]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:585
	// _ = "end of CoverTab[46119]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:585
	_go_fuzz_dep_.CoverTab[46120]++
											tx := (*[4]byte)(unsafe.Pointer(&sa.TxID))
											for i := 0; i < 4; i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:587
		_go_fuzz_dep_.CoverTab[46126]++
												sa.raw.Addr[i+4] = tx[i]
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:588
		// _ = "end of CoverTab[46126]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:589
	// _ = "end of CoverTab[46120]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:589
	_go_fuzz_dep_.CoverTab[46121]++
											return unsafe.Pointer(&sa.raw), SizeofSockaddrCAN, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:590
	// _ = "end of CoverTab[46121]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:597
type SockaddrCANJ1939 struct {
	Ifindex	int
	Name	uint64
	PGN	uint32
	Addr	uint8
	raw	RawSockaddrCAN
}

func (sa *SockaddrCANJ1939) sockaddr() (unsafe.Pointer, _Socklen, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:605
	_go_fuzz_dep_.CoverTab[46127]++
											if sa.Ifindex < 0 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:606
		_go_fuzz_dep_.CoverTab[46131]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:606
		return sa.Ifindex > 0x7fffffff
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:606
		// _ = "end of CoverTab[46131]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:606
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:606
		_go_fuzz_dep_.CoverTab[46132]++
												return nil, 0, EINVAL
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:607
		// _ = "end of CoverTab[46132]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:608
		_go_fuzz_dep_.CoverTab[46133]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:608
		// _ = "end of CoverTab[46133]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:608
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:608
	// _ = "end of CoverTab[46127]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:608
	_go_fuzz_dep_.CoverTab[46128]++
											sa.raw.Family = AF_CAN
											sa.raw.Ifindex = int32(sa.Ifindex)
											n := (*[8]byte)(unsafe.Pointer(&sa.Name))
											for i := 0; i < 8; i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:612
		_go_fuzz_dep_.CoverTab[46134]++
												sa.raw.Addr[i] = n[i]
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:613
		// _ = "end of CoverTab[46134]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:614
	// _ = "end of CoverTab[46128]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:614
	_go_fuzz_dep_.CoverTab[46129]++
											p := (*[4]byte)(unsafe.Pointer(&sa.PGN))
											for i := 0; i < 4; i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:616
		_go_fuzz_dep_.CoverTab[46135]++
												sa.raw.Addr[i+8] = p[i]
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:617
		// _ = "end of CoverTab[46135]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:618
	// _ = "end of CoverTab[46129]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:618
	_go_fuzz_dep_.CoverTab[46130]++
											sa.raw.Addr[12] = sa.Addr
											return unsafe.Pointer(&sa.raw), SizeofSockaddrCAN, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:620
	// _ = "end of CoverTab[46130]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:686
type SockaddrALG struct {
	Type	string
	Name	string
	Feature	uint32
	Mask	uint32
	raw	RawSockaddrALG
}

func (sa *SockaddrALG) sockaddr() (unsafe.Pointer, _Socklen, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:694
	_go_fuzz_dep_.CoverTab[46136]++

											if len(sa.Type) > 13 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:696
		_go_fuzz_dep_.CoverTab[46141]++
												return nil, 0, EINVAL
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:697
		// _ = "end of CoverTab[46141]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:698
		_go_fuzz_dep_.CoverTab[46142]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:698
		// _ = "end of CoverTab[46142]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:698
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:698
	// _ = "end of CoverTab[46136]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:698
	_go_fuzz_dep_.CoverTab[46137]++
											if len(sa.Name) > 63 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:699
		_go_fuzz_dep_.CoverTab[46143]++
												return nil, 0, EINVAL
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:700
		// _ = "end of CoverTab[46143]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:701
		_go_fuzz_dep_.CoverTab[46144]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:701
		// _ = "end of CoverTab[46144]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:701
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:701
	// _ = "end of CoverTab[46137]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:701
	_go_fuzz_dep_.CoverTab[46138]++

											sa.raw.Family = AF_ALG
											sa.raw.Feat = sa.Feature
											sa.raw.Mask = sa.Mask

											typ, err := ByteSliceFromString(sa.Type)
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:708
		_go_fuzz_dep_.CoverTab[46145]++
												return nil, 0, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:709
		// _ = "end of CoverTab[46145]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:710
		_go_fuzz_dep_.CoverTab[46146]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:710
		// _ = "end of CoverTab[46146]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:710
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:710
	// _ = "end of CoverTab[46138]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:710
	_go_fuzz_dep_.CoverTab[46139]++
											name, err := ByteSliceFromString(sa.Name)
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:712
		_go_fuzz_dep_.CoverTab[46147]++
												return nil, 0, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:713
		// _ = "end of CoverTab[46147]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:714
		_go_fuzz_dep_.CoverTab[46148]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:714
		// _ = "end of CoverTab[46148]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:714
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:714
	// _ = "end of CoverTab[46139]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:714
	_go_fuzz_dep_.CoverTab[46140]++

											copy(sa.raw.Type[:], typ)
											copy(sa.raw.Name[:], name)

											return unsafe.Pointer(&sa.raw), SizeofSockaddrALG, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:719
	// _ = "end of CoverTab[46140]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:726
type SockaddrVM struct {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:732
	CID	uint32
	Port	uint32
	Flags	uint8
	raw	RawSockaddrVM
}

func (sa *SockaddrVM) sockaddr() (unsafe.Pointer, _Socklen, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:738
	_go_fuzz_dep_.CoverTab[46149]++
											sa.raw.Family = AF_VSOCK
											sa.raw.Port = sa.Port
											sa.raw.Cid = sa.CID
											sa.raw.Flags = sa.Flags

											return unsafe.Pointer(&sa.raw), SizeofSockaddrVM, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:744
	// _ = "end of CoverTab[46149]"
}

type SockaddrXDP struct {
	Flags		uint16
	Ifindex		uint32
	QueueID		uint32
	SharedUmemFD	uint32
	raw		RawSockaddrXDP
}

func (sa *SockaddrXDP) sockaddr() (unsafe.Pointer, _Socklen, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:755
	_go_fuzz_dep_.CoverTab[46150]++
											sa.raw.Family = AF_XDP
											sa.raw.Flags = sa.Flags
											sa.raw.Ifindex = sa.Ifindex
											sa.raw.Queue_id = sa.QueueID
											sa.raw.Shared_umem_fd = sa.SharedUmemFD

											return unsafe.Pointer(&sa.raw), SizeofSockaddrXDP, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:762
	// _ = "end of CoverTab[46150]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:772
const px_proto_oe = 0

type SockaddrPPPoE struct {
	SID	uint16
	Remote	[]byte
	Dev	string
	raw	RawSockaddrPPPoX
}

func (sa *SockaddrPPPoE) sockaddr() (unsafe.Pointer, _Socklen, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:781
	_go_fuzz_dep_.CoverTab[46151]++
											if len(sa.Remote) != 6 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:782
		_go_fuzz_dep_.CoverTab[46155]++
												return nil, 0, EINVAL
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:783
		// _ = "end of CoverTab[46155]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:784
		_go_fuzz_dep_.CoverTab[46156]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:784
		// _ = "end of CoverTab[46156]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:784
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:784
	// _ = "end of CoverTab[46151]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:784
	_go_fuzz_dep_.CoverTab[46152]++
											if len(sa.Dev) > IFNAMSIZ-1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:785
		_go_fuzz_dep_.CoverTab[46157]++
												return nil, 0, EINVAL
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:786
		// _ = "end of CoverTab[46157]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:787
		_go_fuzz_dep_.CoverTab[46158]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:787
		// _ = "end of CoverTab[46158]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:787
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:787
	// _ = "end of CoverTab[46152]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:787
	_go_fuzz_dep_.CoverTab[46153]++

											*(*uint16)(unsafe.Pointer(&sa.raw[0])) = AF_PPPOX

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:798
	binary.BigEndian.PutUint32(sa.raw[2:6], px_proto_oe)

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:801
	binary.BigEndian.PutUint16(sa.raw[6:8], sa.SID)
	copy(sa.raw[8:14], sa.Remote)
	for i := 14; i < 14+IFNAMSIZ; i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:803
		_go_fuzz_dep_.CoverTab[46159]++
												sa.raw[i] = 0
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:804
		// _ = "end of CoverTab[46159]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:805
	// _ = "end of CoverTab[46153]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:805
	_go_fuzz_dep_.CoverTab[46154]++
											copy(sa.raw[14:], sa.Dev)
											return unsafe.Pointer(&sa.raw), SizeofSockaddrPPPoX, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:807
	// _ = "end of CoverTab[46154]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:812
type SockaddrTIPC struct {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:815
	Scope	int

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:824
	Addr	TIPCAddr

	raw	RawSockaddrTIPC
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:832
type TIPCAddr interface {
	tipcAddrtype() uint8
	tipcAddr() [12]byte
}

func (sa *TIPCSocketAddr) tipcAddr() [12]byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:837
	_go_fuzz_dep_.CoverTab[46160]++
											var out [12]byte
											copy(out[:], (*(*[unsafe.Sizeof(TIPCSocketAddr{})]byte)(unsafe.Pointer(sa)))[:])
											return out
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:840
	// _ = "end of CoverTab[46160]"
}

func (sa *TIPCSocketAddr) tipcAddrtype() uint8 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:843
	_go_fuzz_dep_.CoverTab[46161]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:843
	return TIPC_SOCKET_ADDR
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:843
	// _ = "end of CoverTab[46161]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:843
}

func (sa *TIPCServiceRange) tipcAddr() [12]byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:845
	_go_fuzz_dep_.CoverTab[46162]++
											var out [12]byte
											copy(out[:], (*(*[unsafe.Sizeof(TIPCServiceRange{})]byte)(unsafe.Pointer(sa)))[:])
											return out
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:848
	// _ = "end of CoverTab[46162]"
}

func (sa *TIPCServiceRange) tipcAddrtype() uint8 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:851
	_go_fuzz_dep_.CoverTab[46163]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:851
	return TIPC_SERVICE_RANGE
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:851
	// _ = "end of CoverTab[46163]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:851
}

func (sa *TIPCServiceName) tipcAddr() [12]byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:853
	_go_fuzz_dep_.CoverTab[46164]++
											var out [12]byte
											copy(out[:], (*(*[unsafe.Sizeof(TIPCServiceName{})]byte)(unsafe.Pointer(sa)))[:])
											return out
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:856
	// _ = "end of CoverTab[46164]"
}

func (sa *TIPCServiceName) tipcAddrtype() uint8 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:859
	_go_fuzz_dep_.CoverTab[46165]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:859
	return TIPC_SERVICE_ADDR
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:859
	// _ = "end of CoverTab[46165]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:859
}

func (sa *SockaddrTIPC) sockaddr() (unsafe.Pointer, _Socklen, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:861
	_go_fuzz_dep_.CoverTab[46166]++
											if sa.Addr == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:862
		_go_fuzz_dep_.CoverTab[46168]++
												return nil, 0, EINVAL
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:863
		// _ = "end of CoverTab[46168]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:864
		_go_fuzz_dep_.CoverTab[46169]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:864
		// _ = "end of CoverTab[46169]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:864
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:864
	// _ = "end of CoverTab[46166]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:864
	_go_fuzz_dep_.CoverTab[46167]++
											sa.raw.Family = AF_TIPC
											sa.raw.Scope = int8(sa.Scope)
											sa.raw.Addrtype = sa.Addr.tipcAddrtype()
											sa.raw.Addr = sa.Addr.tipcAddr()
											return unsafe.Pointer(&sa.raw), SizeofSockaddrTIPC, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:869
	// _ = "end of CoverTab[46167]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:873
type SockaddrL2TPIP struct {
	Addr	[4]byte
	ConnId	uint32
	raw	RawSockaddrL2TPIP
}

func (sa *SockaddrL2TPIP) sockaddr() (unsafe.Pointer, _Socklen, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:879
	_go_fuzz_dep_.CoverTab[46170]++
											sa.raw.Family = AF_INET
											sa.raw.Conn_id = sa.ConnId
											sa.raw.Addr = sa.Addr
											return unsafe.Pointer(&sa.raw), SizeofSockaddrL2TPIP, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:883
	// _ = "end of CoverTab[46170]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:887
type SockaddrL2TPIP6 struct {
	Addr	[16]byte
	ZoneId	uint32
	ConnId	uint32
	raw	RawSockaddrL2TPIP6
}

func (sa *SockaddrL2TPIP6) sockaddr() (unsafe.Pointer, _Socklen, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:894
	_go_fuzz_dep_.CoverTab[46171]++
											sa.raw.Family = AF_INET6
											sa.raw.Conn_id = sa.ConnId
											sa.raw.Scope_id = sa.ZoneId
											sa.raw.Addr = sa.Addr
											return unsafe.Pointer(&sa.raw), SizeofSockaddrL2TPIP6, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:899
	// _ = "end of CoverTab[46171]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:903
type SockaddrIUCV struct {
	UserID	string
	Name	string
	raw	RawSockaddrIUCV
}

func (sa *SockaddrIUCV) sockaddr() (unsafe.Pointer, _Socklen, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:909
	_go_fuzz_dep_.CoverTab[46172]++
											sa.raw.Family = AF_IUCV

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:914
	for i := 0; i < 8; i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:914
		_go_fuzz_dep_.CoverTab[46177]++
												sa.raw.Nodeid[i] = ' '
												sa.raw.User_id[i] = ' '
												sa.raw.Name[i] = ' '
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:917
		// _ = "end of CoverTab[46177]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:918
	// _ = "end of CoverTab[46172]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:918
	_go_fuzz_dep_.CoverTab[46173]++
											if len(sa.UserID) > 8 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:919
		_go_fuzz_dep_.CoverTab[46178]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:919
		return len(sa.Name) > 8
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:919
		// _ = "end of CoverTab[46178]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:919
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:919
		_go_fuzz_dep_.CoverTab[46179]++
												return nil, 0, EINVAL
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:920
		// _ = "end of CoverTab[46179]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:921
		_go_fuzz_dep_.CoverTab[46180]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:921
		// _ = "end of CoverTab[46180]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:921
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:921
	// _ = "end of CoverTab[46173]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:921
	_go_fuzz_dep_.CoverTab[46174]++
											for i, b := range []byte(sa.UserID[:]) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:922
		_go_fuzz_dep_.CoverTab[46181]++
												sa.raw.User_id[i] = int8(b)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:923
		// _ = "end of CoverTab[46181]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:924
	// _ = "end of CoverTab[46174]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:924
	_go_fuzz_dep_.CoverTab[46175]++
											for i, b := range []byte(sa.Name[:]) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:925
		_go_fuzz_dep_.CoverTab[46182]++
												sa.raw.Name[i] = int8(b)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:926
		// _ = "end of CoverTab[46182]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:927
	// _ = "end of CoverTab[46175]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:927
	_go_fuzz_dep_.CoverTab[46176]++
											return unsafe.Pointer(&sa.raw), SizeofSockaddrIUCV, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:928
	// _ = "end of CoverTab[46176]"
}

type SockaddrNFC struct {
	DeviceIdx	uint32
	TargetIdx	uint32
	NFCProtocol	uint32
	raw		RawSockaddrNFC
}

func (sa *SockaddrNFC) sockaddr() (unsafe.Pointer, _Socklen, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:938
	_go_fuzz_dep_.CoverTab[46183]++
											sa.raw.Sa_family = AF_NFC
											sa.raw.Dev_idx = sa.DeviceIdx
											sa.raw.Target_idx = sa.TargetIdx
											sa.raw.Nfc_protocol = sa.NFCProtocol
											return unsafe.Pointer(&sa.raw), SizeofSockaddrNFC, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:943
	// _ = "end of CoverTab[46183]"
}

type SockaddrNFCLLCP struct {
	DeviceIdx	uint32
	TargetIdx	uint32
	NFCProtocol	uint32
	DestinationSAP	uint8
	SourceSAP	uint8
	ServiceName	string
	raw		RawSockaddrNFCLLCP
}

func (sa *SockaddrNFCLLCP) sockaddr() (unsafe.Pointer, _Socklen, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:956
	_go_fuzz_dep_.CoverTab[46184]++
											sa.raw.Sa_family = AF_NFC
											sa.raw.Dev_idx = sa.DeviceIdx
											sa.raw.Target_idx = sa.TargetIdx
											sa.raw.Nfc_protocol = sa.NFCProtocol
											sa.raw.Dsap = sa.DestinationSAP
											sa.raw.Ssap = sa.SourceSAP
											if len(sa.ServiceName) > len(sa.raw.Service_name) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:963
		_go_fuzz_dep_.CoverTab[46186]++
												return nil, 0, EINVAL
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:964
		// _ = "end of CoverTab[46186]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:965
		_go_fuzz_dep_.CoverTab[46187]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:965
		// _ = "end of CoverTab[46187]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:965
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:965
	// _ = "end of CoverTab[46184]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:965
	_go_fuzz_dep_.CoverTab[46185]++
											copy(sa.raw.Service_name[:], sa.ServiceName)
											sa.raw.SetServiceNameLen(len(sa.ServiceName))
											return unsafe.Pointer(&sa.raw), SizeofSockaddrNFCLLCP, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:968
	// _ = "end of CoverTab[46185]"
}

var socketProtocol = func(fd int) (int, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:971
	_go_fuzz_dep_.CoverTab[46188]++
											return GetsockoptInt(fd, SOL_SOCKET, SO_PROTOCOL)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:972
	// _ = "end of CoverTab[46188]"
}

func anyToSockaddr(fd int, rsa *RawSockaddrAny) (Sockaddr, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:975
	_go_fuzz_dep_.CoverTab[46189]++
											switch rsa.Addr.Family {
	case AF_NETLINK:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:977
		_go_fuzz_dep_.CoverTab[46191]++
												pp := (*RawSockaddrNetlink)(unsafe.Pointer(rsa))
												sa := new(SockaddrNetlink)
												sa.Family = pp.Family
												sa.Pad = pp.Pad
												sa.Pid = pp.Pid
												sa.Groups = pp.Groups
												return sa, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:984
		// _ = "end of CoverTab[46191]"

	case AF_PACKET:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:986
		_go_fuzz_dep_.CoverTab[46192]++
												pp := (*RawSockaddrLinklayer)(unsafe.Pointer(rsa))
												sa := new(SockaddrLinklayer)
												sa.Protocol = pp.Protocol
												sa.Ifindex = int(pp.Ifindex)
												sa.Hatype = pp.Hatype
												sa.Pkttype = pp.Pkttype
												sa.Halen = pp.Halen
												sa.Addr = pp.Addr
												return sa, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:995
		// _ = "end of CoverTab[46192]"

	case AF_UNIX:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:997
		_go_fuzz_dep_.CoverTab[46193]++
												pp := (*RawSockaddrUnix)(unsafe.Pointer(rsa))
												sa := new(SockaddrUnix)
												if pp.Path[0] == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1000
			_go_fuzz_dep_.CoverTab[46216]++

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1006
			pp.Path[0] = '@'
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1006
			// _ = "end of CoverTab[46216]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1007
			_go_fuzz_dep_.CoverTab[46217]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1007
			// _ = "end of CoverTab[46217]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1007
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1007
		// _ = "end of CoverTab[46193]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1007
		_go_fuzz_dep_.CoverTab[46194]++

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1014
		n := 0
		for n < len(pp.Path) && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1015
			_go_fuzz_dep_.CoverTab[46218]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1015
			return pp.Path[n] != 0
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1015
			// _ = "end of CoverTab[46218]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1015
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1015
			_go_fuzz_dep_.CoverTab[46219]++
													n++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1016
			// _ = "end of CoverTab[46219]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1017
		// _ = "end of CoverTab[46194]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1017
		_go_fuzz_dep_.CoverTab[46195]++
												sa.Name = string(unsafe.Slice((*byte)(unsafe.Pointer(&pp.Path[0])), n))
												return sa, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1019
		// _ = "end of CoverTab[46195]"

	case AF_INET:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1021
		_go_fuzz_dep_.CoverTab[46196]++
												proto, err := socketProtocol(fd)
												if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1023
			_go_fuzz_dep_.CoverTab[46220]++
													return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1024
			// _ = "end of CoverTab[46220]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1025
			_go_fuzz_dep_.CoverTab[46221]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1025
			// _ = "end of CoverTab[46221]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1025
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1025
		// _ = "end of CoverTab[46196]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1025
		_go_fuzz_dep_.CoverTab[46197]++

												switch proto {
		case IPPROTO_L2TP:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1028
			_go_fuzz_dep_.CoverTab[46222]++
													pp := (*RawSockaddrL2TPIP)(unsafe.Pointer(rsa))
													sa := new(SockaddrL2TPIP)
													sa.ConnId = pp.Conn_id
													sa.Addr = pp.Addr
													return sa, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1033
			// _ = "end of CoverTab[46222]"
		default:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1034
			_go_fuzz_dep_.CoverTab[46223]++
													pp := (*RawSockaddrInet4)(unsafe.Pointer(rsa))
													sa := new(SockaddrInet4)
													p := (*[2]byte)(unsafe.Pointer(&pp.Port))
													sa.Port = int(p[0])<<8 + int(p[1])
													sa.Addr = pp.Addr
													return sa, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1040
			// _ = "end of CoverTab[46223]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1041
		// _ = "end of CoverTab[46197]"

	case AF_INET6:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1043
		_go_fuzz_dep_.CoverTab[46198]++
												proto, err := socketProtocol(fd)
												if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1045
			_go_fuzz_dep_.CoverTab[46224]++
													return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1046
			// _ = "end of CoverTab[46224]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1047
			_go_fuzz_dep_.CoverTab[46225]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1047
			// _ = "end of CoverTab[46225]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1047
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1047
		// _ = "end of CoverTab[46198]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1047
		_go_fuzz_dep_.CoverTab[46199]++

												switch proto {
		case IPPROTO_L2TP:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1050
			_go_fuzz_dep_.CoverTab[46226]++
													pp := (*RawSockaddrL2TPIP6)(unsafe.Pointer(rsa))
													sa := new(SockaddrL2TPIP6)
													sa.ConnId = pp.Conn_id
													sa.ZoneId = pp.Scope_id
													sa.Addr = pp.Addr
													return sa, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1056
			// _ = "end of CoverTab[46226]"
		default:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1057
			_go_fuzz_dep_.CoverTab[46227]++
													pp := (*RawSockaddrInet6)(unsafe.Pointer(rsa))
													sa := new(SockaddrInet6)
													p := (*[2]byte)(unsafe.Pointer(&pp.Port))
													sa.Port = int(p[0])<<8 + int(p[1])
													sa.ZoneId = pp.Scope_id
													sa.Addr = pp.Addr
													return sa, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1064
			// _ = "end of CoverTab[46227]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1065
		// _ = "end of CoverTab[46199]"

	case AF_VSOCK:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1067
		_go_fuzz_dep_.CoverTab[46200]++
												pp := (*RawSockaddrVM)(unsafe.Pointer(rsa))
												sa := &SockaddrVM{
			CID:	pp.Cid,
			Port:	pp.Port,
			Flags:	pp.Flags,
		}
												return sa, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1074
		// _ = "end of CoverTab[46200]"
	case AF_BLUETOOTH:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1075
		_go_fuzz_dep_.CoverTab[46201]++
												proto, err := socketProtocol(fd)
												if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1077
			_go_fuzz_dep_.CoverTab[46228]++
													return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1078
			// _ = "end of CoverTab[46228]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1079
			_go_fuzz_dep_.CoverTab[46229]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1079
			// _ = "end of CoverTab[46229]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1079
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1079
		// _ = "end of CoverTab[46201]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1079
		_go_fuzz_dep_.CoverTab[46202]++

												switch proto {
		case BTPROTO_L2CAP:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1082
			_go_fuzz_dep_.CoverTab[46230]++
													pp := (*RawSockaddrL2)(unsafe.Pointer(rsa))
													sa := &SockaddrL2{
				PSM:		pp.Psm,
				CID:		pp.Cid,
				Addr:		pp.Bdaddr,
				AddrType:	pp.Bdaddr_type,
			}
													return sa, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1090
			// _ = "end of CoverTab[46230]"
		case BTPROTO_RFCOMM:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1091
			_go_fuzz_dep_.CoverTab[46231]++
													pp := (*RawSockaddrRFCOMM)(unsafe.Pointer(rsa))
													sa := &SockaddrRFCOMM{
				Channel:	pp.Channel,
				Addr:		pp.Bdaddr,
			}
													return sa, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1097
			// _ = "end of CoverTab[46231]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1097
		default:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1097
			_go_fuzz_dep_.CoverTab[46232]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1097
			// _ = "end of CoverTab[46232]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1098
		// _ = "end of CoverTab[46202]"
	case AF_XDP:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1099
		_go_fuzz_dep_.CoverTab[46203]++
												pp := (*RawSockaddrXDP)(unsafe.Pointer(rsa))
												sa := &SockaddrXDP{
			Flags:		pp.Flags,
			Ifindex:	pp.Ifindex,
			QueueID:	pp.Queue_id,
			SharedUmemFD:	pp.Shared_umem_fd,
		}
												return sa, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1107
		// _ = "end of CoverTab[46203]"
	case AF_PPPOX:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1108
		_go_fuzz_dep_.CoverTab[46204]++
												pp := (*RawSockaddrPPPoX)(unsafe.Pointer(rsa))
												if binary.BigEndian.Uint32(pp[2:6]) != px_proto_oe {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1110
			_go_fuzz_dep_.CoverTab[46233]++
													return nil, EINVAL
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1111
			// _ = "end of CoverTab[46233]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1112
			_go_fuzz_dep_.CoverTab[46234]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1112
			// _ = "end of CoverTab[46234]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1112
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1112
		// _ = "end of CoverTab[46204]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1112
		_go_fuzz_dep_.CoverTab[46205]++
												sa := &SockaddrPPPoE{
			SID:	binary.BigEndian.Uint16(pp[6:8]),
			Remote:	pp[8:14],
		}
		for i := 14; i < 14+IFNAMSIZ; i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1117
			_go_fuzz_dep_.CoverTab[46235]++
													if pp[i] == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1118
				_go_fuzz_dep_.CoverTab[46236]++
														sa.Dev = string(pp[14:i])
														break
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1120
				// _ = "end of CoverTab[46236]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1121
				_go_fuzz_dep_.CoverTab[46237]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1121
				// _ = "end of CoverTab[46237]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1121
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1121
			// _ = "end of CoverTab[46235]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1122
		// _ = "end of CoverTab[46205]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1122
		_go_fuzz_dep_.CoverTab[46206]++
												return sa, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1123
		// _ = "end of CoverTab[46206]"
	case AF_TIPC:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1124
		_go_fuzz_dep_.CoverTab[46207]++
												pp := (*RawSockaddrTIPC)(unsafe.Pointer(rsa))

												sa := &SockaddrTIPC{
			Scope: int(pp.Scope),
		}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1133
		switch pp.Addrtype {
		case TIPC_SERVICE_RANGE:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1134
			_go_fuzz_dep_.CoverTab[46238]++
													sa.Addr = (*TIPCServiceRange)(unsafe.Pointer(&pp.Addr))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1135
			// _ = "end of CoverTab[46238]"
		case TIPC_SERVICE_ADDR:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1136
			_go_fuzz_dep_.CoverTab[46239]++
													sa.Addr = (*TIPCServiceName)(unsafe.Pointer(&pp.Addr))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1137
			// _ = "end of CoverTab[46239]"
		case TIPC_SOCKET_ADDR:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1138
			_go_fuzz_dep_.CoverTab[46240]++
													sa.Addr = (*TIPCSocketAddr)(unsafe.Pointer(&pp.Addr))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1139
			// _ = "end of CoverTab[46240]"
		default:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1140
			_go_fuzz_dep_.CoverTab[46241]++
													return nil, EINVAL
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1141
			// _ = "end of CoverTab[46241]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1142
		// _ = "end of CoverTab[46207]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1142
		_go_fuzz_dep_.CoverTab[46208]++

												return sa, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1144
		// _ = "end of CoverTab[46208]"
	case AF_IUCV:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1145
		_go_fuzz_dep_.CoverTab[46209]++
												pp := (*RawSockaddrIUCV)(unsafe.Pointer(rsa))

												var user [8]byte
												var name [8]byte

												for i := 0; i < 8; i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1151
			_go_fuzz_dep_.CoverTab[46242]++
													user[i] = byte(pp.User_id[i])
													name[i] = byte(pp.Name[i])
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1153
			// _ = "end of CoverTab[46242]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1154
		// _ = "end of CoverTab[46209]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1154
		_go_fuzz_dep_.CoverTab[46210]++

												sa := &SockaddrIUCV{
			UserID:	string(user[:]),
			Name:	string(name[:]),
		}
												return sa, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1160
		// _ = "end of CoverTab[46210]"

	case AF_CAN:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1162
		_go_fuzz_dep_.CoverTab[46211]++
												proto, err := socketProtocol(fd)
												if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1164
			_go_fuzz_dep_.CoverTab[46243]++
													return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1165
			// _ = "end of CoverTab[46243]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1166
			_go_fuzz_dep_.CoverTab[46244]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1166
			// _ = "end of CoverTab[46244]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1166
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1166
		// _ = "end of CoverTab[46211]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1166
		_go_fuzz_dep_.CoverTab[46212]++

												pp := (*RawSockaddrCAN)(unsafe.Pointer(rsa))

												switch proto {
		case CAN_J1939:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1171
			_go_fuzz_dep_.CoverTab[46245]++
													sa := &SockaddrCANJ1939{
				Ifindex: int(pp.Ifindex),
			}
			name := (*[8]byte)(unsafe.Pointer(&sa.Name))
			for i := 0; i < 8; i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1176
				_go_fuzz_dep_.CoverTab[46251]++
														name[i] = pp.Addr[i]
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1177
				// _ = "end of CoverTab[46251]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1178
			// _ = "end of CoverTab[46245]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1178
			_go_fuzz_dep_.CoverTab[46246]++
													pgn := (*[4]byte)(unsafe.Pointer(&sa.PGN))
													for i := 0; i < 4; i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1180
				_go_fuzz_dep_.CoverTab[46252]++
														pgn[i] = pp.Addr[i+8]
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1181
				// _ = "end of CoverTab[46252]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1182
			// _ = "end of CoverTab[46246]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1182
			_go_fuzz_dep_.CoverTab[46247]++
													addr := (*[1]byte)(unsafe.Pointer(&sa.Addr))
													addr[0] = pp.Addr[12]
													return sa, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1185
			// _ = "end of CoverTab[46247]"
		default:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1186
			_go_fuzz_dep_.CoverTab[46248]++
													sa := &SockaddrCAN{
				Ifindex: int(pp.Ifindex),
			}
			rx := (*[4]byte)(unsafe.Pointer(&sa.RxID))
			for i := 0; i < 4; i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1191
				_go_fuzz_dep_.CoverTab[46253]++
														rx[i] = pp.Addr[i]
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1192
				// _ = "end of CoverTab[46253]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1193
			// _ = "end of CoverTab[46248]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1193
			_go_fuzz_dep_.CoverTab[46249]++
													tx := (*[4]byte)(unsafe.Pointer(&sa.TxID))
													for i := 0; i < 4; i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1195
				_go_fuzz_dep_.CoverTab[46254]++
														tx[i] = pp.Addr[i+4]
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1196
				// _ = "end of CoverTab[46254]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1197
			// _ = "end of CoverTab[46249]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1197
			_go_fuzz_dep_.CoverTab[46250]++
													return sa, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1198
			// _ = "end of CoverTab[46250]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1199
		// _ = "end of CoverTab[46212]"
	case AF_NFC:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1200
		_go_fuzz_dep_.CoverTab[46213]++
												proto, err := socketProtocol(fd)
												if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1202
			_go_fuzz_dep_.CoverTab[46255]++
													return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1203
			// _ = "end of CoverTab[46255]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1204
			_go_fuzz_dep_.CoverTab[46256]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1204
			// _ = "end of CoverTab[46256]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1204
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1204
		// _ = "end of CoverTab[46213]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1204
		_go_fuzz_dep_.CoverTab[46214]++
												switch proto {
		case NFC_SOCKPROTO_RAW:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1206
			_go_fuzz_dep_.CoverTab[46257]++
													pp := (*RawSockaddrNFC)(unsafe.Pointer(rsa))
													sa := &SockaddrNFC{
				DeviceIdx:	pp.Dev_idx,
				TargetIdx:	pp.Target_idx,
				NFCProtocol:	pp.Nfc_protocol,
			}
													return sa, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1213
			// _ = "end of CoverTab[46257]"
		case NFC_SOCKPROTO_LLCP:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1214
			_go_fuzz_dep_.CoverTab[46258]++
													pp := (*RawSockaddrNFCLLCP)(unsafe.Pointer(rsa))
													if uint64(pp.Service_name_len) > uint64(len(pp.Service_name)) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1216
				_go_fuzz_dep_.CoverTab[46261]++
														return nil, EINVAL
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1217
				// _ = "end of CoverTab[46261]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1218
				_go_fuzz_dep_.CoverTab[46262]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1218
				// _ = "end of CoverTab[46262]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1218
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1218
			// _ = "end of CoverTab[46258]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1218
			_go_fuzz_dep_.CoverTab[46259]++
													sa := &SockaddrNFCLLCP{
				DeviceIdx:	pp.Dev_idx,
				TargetIdx:	pp.Target_idx,
				NFCProtocol:	pp.Nfc_protocol,
				DestinationSAP:	pp.Dsap,
				SourceSAP:	pp.Ssap,
				ServiceName:	string(pp.Service_name[:pp.Service_name_len]),
			}
													return sa, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1227
			// _ = "end of CoverTab[46259]"
		default:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1228
			_go_fuzz_dep_.CoverTab[46260]++
													return nil, EINVAL
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1229
			// _ = "end of CoverTab[46260]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1230
		// _ = "end of CoverTab[46214]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1230
	default:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1230
		_go_fuzz_dep_.CoverTab[46215]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1230
		// _ = "end of CoverTab[46215]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1231
	// _ = "end of CoverTab[46189]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1231
	_go_fuzz_dep_.CoverTab[46190]++
											return nil, EAFNOSUPPORT
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1232
	// _ = "end of CoverTab[46190]"
}

func Accept(fd int) (nfd int, sa Sockaddr, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1235
	_go_fuzz_dep_.CoverTab[46263]++
											var rsa RawSockaddrAny
											var len _Socklen = SizeofSockaddrAny
											nfd, err = accept4(fd, &rsa, &len, 0)
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1239
		_go_fuzz_dep_.CoverTab[46266]++
												return
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1240
		// _ = "end of CoverTab[46266]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1241
		_go_fuzz_dep_.CoverTab[46267]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1241
		// _ = "end of CoverTab[46267]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1241
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1241
	// _ = "end of CoverTab[46263]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1241
	_go_fuzz_dep_.CoverTab[46264]++
											sa, err = anyToSockaddr(fd, &rsa)
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1243
		_go_fuzz_dep_.CoverTab[46268]++
												Close(nfd)
												nfd = 0
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1245
		// _ = "end of CoverTab[46268]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1246
		_go_fuzz_dep_.CoverTab[46269]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1246
		// _ = "end of CoverTab[46269]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1246
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1246
	// _ = "end of CoverTab[46264]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1246
	_go_fuzz_dep_.CoverTab[46265]++
											return
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1247
	// _ = "end of CoverTab[46265]"
}

func Accept4(fd int, flags int) (nfd int, sa Sockaddr, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1250
	_go_fuzz_dep_.CoverTab[46270]++
											var rsa RawSockaddrAny
											var len _Socklen = SizeofSockaddrAny
											nfd, err = accept4(fd, &rsa, &len, flags)
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1254
		_go_fuzz_dep_.CoverTab[46274]++
												return
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1255
		// _ = "end of CoverTab[46274]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1256
		_go_fuzz_dep_.CoverTab[46275]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1256
		// _ = "end of CoverTab[46275]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1256
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1256
	// _ = "end of CoverTab[46270]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1256
	_go_fuzz_dep_.CoverTab[46271]++
											if len > SizeofSockaddrAny {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1257
		_go_fuzz_dep_.CoverTab[46276]++
												panic("RawSockaddrAny too small")
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1258
		// _ = "end of CoverTab[46276]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1259
		_go_fuzz_dep_.CoverTab[46277]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1259
		// _ = "end of CoverTab[46277]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1259
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1259
	// _ = "end of CoverTab[46271]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1259
	_go_fuzz_dep_.CoverTab[46272]++
											sa, err = anyToSockaddr(fd, &rsa)
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1261
		_go_fuzz_dep_.CoverTab[46278]++
												Close(nfd)
												nfd = 0
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1263
		// _ = "end of CoverTab[46278]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1264
		_go_fuzz_dep_.CoverTab[46279]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1264
		// _ = "end of CoverTab[46279]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1264
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1264
	// _ = "end of CoverTab[46272]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1264
	_go_fuzz_dep_.CoverTab[46273]++
											return
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1265
	// _ = "end of CoverTab[46273]"
}

func Getsockname(fd int) (sa Sockaddr, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1268
	_go_fuzz_dep_.CoverTab[46280]++
											var rsa RawSockaddrAny
											var len _Socklen = SizeofSockaddrAny
											if err = getsockname(fd, &rsa, &len); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1271
		_go_fuzz_dep_.CoverTab[46282]++
												return
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1272
		// _ = "end of CoverTab[46282]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1273
		_go_fuzz_dep_.CoverTab[46283]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1273
		// _ = "end of CoverTab[46283]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1273
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1273
	// _ = "end of CoverTab[46280]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1273
	_go_fuzz_dep_.CoverTab[46281]++
											return anyToSockaddr(fd, &rsa)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1274
	// _ = "end of CoverTab[46281]"
}

func GetsockoptIPMreqn(fd, level, opt int) (*IPMreqn, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1277
	_go_fuzz_dep_.CoverTab[46284]++
											var value IPMreqn
											vallen := _Socklen(SizeofIPMreqn)
											err := getsockopt(fd, level, opt, unsafe.Pointer(&value), &vallen)
											return &value, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1281
	// _ = "end of CoverTab[46284]"
}

func GetsockoptUcred(fd, level, opt int) (*Ucred, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1284
	_go_fuzz_dep_.CoverTab[46285]++
											var value Ucred
											vallen := _Socklen(SizeofUcred)
											err := getsockopt(fd, level, opt, unsafe.Pointer(&value), &vallen)
											return &value, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1288
	// _ = "end of CoverTab[46285]"
}

func GetsockoptTCPInfo(fd, level, opt int) (*TCPInfo, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1291
	_go_fuzz_dep_.CoverTab[46286]++
											var value TCPInfo
											vallen := _Socklen(SizeofTCPInfo)
											err := getsockopt(fd, level, opt, unsafe.Pointer(&value), &vallen)
											return &value, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1295
	// _ = "end of CoverTab[46286]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1300
func GetsockoptString(fd, level, opt int) (string, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1300
	_go_fuzz_dep_.CoverTab[46287]++
											buf := make([]byte, 256)
											vallen := _Socklen(len(buf))
											err := getsockopt(fd, level, opt, unsafe.Pointer(&buf[0]), &vallen)
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1304
		_go_fuzz_dep_.CoverTab[46289]++
												if err == ERANGE {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1305
			_go_fuzz_dep_.CoverTab[46291]++
													buf = make([]byte, vallen)
													err = getsockopt(fd, level, opt, unsafe.Pointer(&buf[0]), &vallen)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1307
			// _ = "end of CoverTab[46291]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1308
			_go_fuzz_dep_.CoverTab[46292]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1308
			// _ = "end of CoverTab[46292]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1308
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1308
		// _ = "end of CoverTab[46289]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1308
		_go_fuzz_dep_.CoverTab[46290]++
												if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1309
			_go_fuzz_dep_.CoverTab[46293]++
													return "", err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1310
			// _ = "end of CoverTab[46293]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1311
			_go_fuzz_dep_.CoverTab[46294]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1311
			// _ = "end of CoverTab[46294]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1311
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1311
		// _ = "end of CoverTab[46290]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1312
		_go_fuzz_dep_.CoverTab[46295]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1312
		// _ = "end of CoverTab[46295]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1312
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1312
	// _ = "end of CoverTab[46287]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1312
	_go_fuzz_dep_.CoverTab[46288]++
											return string(buf[:vallen-1]), nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1313
	// _ = "end of CoverTab[46288]"
}

func GetsockoptTpacketStats(fd, level, opt int) (*TpacketStats, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1316
	_go_fuzz_dep_.CoverTab[46296]++
											var value TpacketStats
											vallen := _Socklen(SizeofTpacketStats)
											err := getsockopt(fd, level, opt, unsafe.Pointer(&value), &vallen)
											return &value, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1320
	// _ = "end of CoverTab[46296]"
}

func GetsockoptTpacketStatsV3(fd, level, opt int) (*TpacketStatsV3, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1323
	_go_fuzz_dep_.CoverTab[46297]++
											var value TpacketStatsV3
											vallen := _Socklen(SizeofTpacketStatsV3)
											err := getsockopt(fd, level, opt, unsafe.Pointer(&value), &vallen)
											return &value, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1327
	// _ = "end of CoverTab[46297]"
}

func SetsockoptIPMreqn(fd, level, opt int, mreq *IPMreqn) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1330
	_go_fuzz_dep_.CoverTab[46298]++
											return setsockopt(fd, level, opt, unsafe.Pointer(mreq), unsafe.Sizeof(*mreq))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1331
	// _ = "end of CoverTab[46298]"
}

func SetsockoptPacketMreq(fd, level, opt int, mreq *PacketMreq) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1334
	_go_fuzz_dep_.CoverTab[46299]++
											return setsockopt(fd, level, opt, unsafe.Pointer(mreq), unsafe.Sizeof(*mreq))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1335
	// _ = "end of CoverTab[46299]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1340
func SetsockoptSockFprog(fd, level, opt int, fprog *SockFprog) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1340
	_go_fuzz_dep_.CoverTab[46300]++
											return setsockopt(fd, level, opt, unsafe.Pointer(fprog), unsafe.Sizeof(*fprog))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1341
	// _ = "end of CoverTab[46300]"
}

func SetsockoptCanRawFilter(fd, level, opt int, filter []CanFilter) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1344
	_go_fuzz_dep_.CoverTab[46301]++
											var p unsafe.Pointer
											if len(filter) > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1346
		_go_fuzz_dep_.CoverTab[46303]++
												p = unsafe.Pointer(&filter[0])
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1347
		// _ = "end of CoverTab[46303]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1348
		_go_fuzz_dep_.CoverTab[46304]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1348
		// _ = "end of CoverTab[46304]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1348
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1348
	// _ = "end of CoverTab[46301]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1348
	_go_fuzz_dep_.CoverTab[46302]++
											return setsockopt(fd, level, opt, p, uintptr(len(filter)*SizeofCanFilter))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1349
	// _ = "end of CoverTab[46302]"
}

func SetsockoptTpacketReq(fd, level, opt int, tp *TpacketReq) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1352
	_go_fuzz_dep_.CoverTab[46305]++
											return setsockopt(fd, level, opt, unsafe.Pointer(tp), unsafe.Sizeof(*tp))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1353
	// _ = "end of CoverTab[46305]"
}

func SetsockoptTpacketReq3(fd, level, opt int, tp *TpacketReq3) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1356
	_go_fuzz_dep_.CoverTab[46306]++
											return setsockopt(fd, level, opt, unsafe.Pointer(tp), unsafe.Sizeof(*tp))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1357
	// _ = "end of CoverTab[46306]"
}

func SetsockoptTCPRepairOpt(fd, level, opt int, o []TCPRepairOpt) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1360
	_go_fuzz_dep_.CoverTab[46307]++
											if len(o) == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1361
		_go_fuzz_dep_.CoverTab[46309]++
												return EINVAL
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1362
		// _ = "end of CoverTab[46309]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1363
		_go_fuzz_dep_.CoverTab[46310]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1363
		// _ = "end of CoverTab[46310]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1363
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1363
	// _ = "end of CoverTab[46307]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1363
	_go_fuzz_dep_.CoverTab[46308]++
											return setsockopt(fd, level, opt, unsafe.Pointer(&o[0]), uintptr(SizeofTCPRepairOpt*len(o)))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1364
	// _ = "end of CoverTab[46308]"
}

func SetsockoptTCPMD5Sig(fd, level, opt int, s *TCPMD5Sig) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1367
	_go_fuzz_dep_.CoverTab[46311]++
											return setsockopt(fd, level, opt, unsafe.Pointer(s), unsafe.Sizeof(*s))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1368
	// _ = "end of CoverTab[46311]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1387
func KeyctlString(cmd int, id int) (string, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1387
	_go_fuzz_dep_.CoverTab[46312]++

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1392
	var buffer []byte
	for {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1393
		_go_fuzz_dep_.CoverTab[46313]++

												length, err := KeyctlBuffer(cmd, id, buffer, 0)
												if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1396
			_go_fuzz_dep_.CoverTab[46316]++
													return "", err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1397
			// _ = "end of CoverTab[46316]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1398
			_go_fuzz_dep_.CoverTab[46317]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1398
			// _ = "end of CoverTab[46317]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1398
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1398
		// _ = "end of CoverTab[46313]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1398
		_go_fuzz_dep_.CoverTab[46314]++

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1401
		if length <= len(buffer) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1401
			_go_fuzz_dep_.CoverTab[46318]++

													return string(buffer[:length-1]), nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1403
			// _ = "end of CoverTab[46318]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1404
			_go_fuzz_dep_.CoverTab[46319]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1404
			// _ = "end of CoverTab[46319]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1404
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1404
		// _ = "end of CoverTab[46314]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1404
		_go_fuzz_dep_.CoverTab[46315]++

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1407
		buffer = make([]byte, length)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1407
		// _ = "end of CoverTab[46315]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1408
	// _ = "end of CoverTab[46312]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1416
func KeyctlGetKeyringID(id int, create bool) (ringid int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1416
	_go_fuzz_dep_.CoverTab[46320]++
											createInt := 0
											if create {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1418
		_go_fuzz_dep_.CoverTab[46322]++
												createInt = 1
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1419
		// _ = "end of CoverTab[46322]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1420
		_go_fuzz_dep_.CoverTab[46323]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1420
		// _ = "end of CoverTab[46323]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1420
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1420
	// _ = "end of CoverTab[46320]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1420
	_go_fuzz_dep_.CoverTab[46321]++
											return KeyctlInt(KEYCTL_GET_KEYRING_ID, id, createInt, 0, 0)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1421
	// _ = "end of CoverTab[46321]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1429
func KeyctlSetperm(id int, perm uint32) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1429
	_go_fuzz_dep_.CoverTab[46324]++
											_, err := KeyctlInt(KEYCTL_SETPERM, id, int(perm), 0, 0)
											return err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1431
	// _ = "end of CoverTab[46324]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1439
func KeyctlJoinSessionKeyring(name string) (ringid int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1439
	_go_fuzz_dep_.CoverTab[46325]++
											return keyctlJoin(KEYCTL_JOIN_SESSION_KEYRING, name)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1440
	// _ = "end of CoverTab[46325]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1448
func KeyctlSearch(ringid int, keyType, description string, destRingid int) (id int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1448
	_go_fuzz_dep_.CoverTab[46326]++
											return keyctlSearch(KEYCTL_SEARCH, ringid, keyType, description, destRingid)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1449
	// _ = "end of CoverTab[46326]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1459
func KeyctlInstantiateIOV(id int, payload []Iovec, ringid int) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1459
	_go_fuzz_dep_.CoverTab[46327]++
											return keyctlIOV(KEYCTL_INSTANTIATE_IOV, id, payload, ringid)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1460
	// _ = "end of CoverTab[46327]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1474
func KeyctlDHCompute(params *KeyctlDHParams, buffer []byte) (size int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1474
	_go_fuzz_dep_.CoverTab[46328]++
											return keyctlDH(KEYCTL_DH_COMPUTE, params, buffer)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1475
	// _ = "end of CoverTab[46328]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1496
func KeyctlRestrictKeyring(ringid int, keyType string, restriction string) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1496
	_go_fuzz_dep_.CoverTab[46329]++
											if keyType == "" {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1497
		_go_fuzz_dep_.CoverTab[46331]++
												return keyctlRestrictKeyring(KEYCTL_RESTRICT_KEYRING, ringid)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1498
		// _ = "end of CoverTab[46331]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1499
		_go_fuzz_dep_.CoverTab[46332]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1499
		// _ = "end of CoverTab[46332]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1499
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1499
	// _ = "end of CoverTab[46329]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1499
	_go_fuzz_dep_.CoverTab[46330]++
											return keyctlRestrictKeyringByType(KEYCTL_RESTRICT_KEYRING, ringid, keyType, restriction)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1500
	// _ = "end of CoverTab[46330]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1506
func recvmsgRaw(fd int, iov []Iovec, oob []byte, flags int, rsa *RawSockaddrAny) (n, oobn int, recvflags int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1506
	_go_fuzz_dep_.CoverTab[46333]++
											var msg Msghdr
											msg.Name = (*byte)(unsafe.Pointer(rsa))
											msg.Namelen = uint32(SizeofSockaddrAny)
											var dummy byte
											if len(oob) > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1511
		_go_fuzz_dep_.CoverTab[46337]++
												if emptyIovecs(iov) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1512
			_go_fuzz_dep_.CoverTab[46339]++
													var sockType int
													sockType, err = GetsockoptInt(fd, SOL_SOCKET, SO_TYPE)
													if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1515
				_go_fuzz_dep_.CoverTab[46341]++
														return
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1516
				// _ = "end of CoverTab[46341]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1517
				_go_fuzz_dep_.CoverTab[46342]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1517
				// _ = "end of CoverTab[46342]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1517
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1517
			// _ = "end of CoverTab[46339]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1517
			_go_fuzz_dep_.CoverTab[46340]++

													if sockType != SOCK_DGRAM {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1519
				_go_fuzz_dep_.CoverTab[46343]++
														var iova [1]Iovec
														iova[0].Base = &dummy
														iova[0].SetLen(1)
														iov = iova[:]
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1523
				// _ = "end of CoverTab[46343]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1524
				_go_fuzz_dep_.CoverTab[46344]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1524
				// _ = "end of CoverTab[46344]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1524
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1524
			// _ = "end of CoverTab[46340]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1525
			_go_fuzz_dep_.CoverTab[46345]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1525
			// _ = "end of CoverTab[46345]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1525
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1525
		// _ = "end of CoverTab[46337]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1525
		_go_fuzz_dep_.CoverTab[46338]++
												msg.Control = &oob[0]
												msg.SetControllen(len(oob))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1527
		// _ = "end of CoverTab[46338]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1528
		_go_fuzz_dep_.CoverTab[46346]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1528
		// _ = "end of CoverTab[46346]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1528
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1528
	// _ = "end of CoverTab[46333]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1528
	_go_fuzz_dep_.CoverTab[46334]++
											if len(iov) > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1529
		_go_fuzz_dep_.CoverTab[46347]++
												msg.Iov = &iov[0]
												msg.SetIovlen(len(iov))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1531
		// _ = "end of CoverTab[46347]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1532
		_go_fuzz_dep_.CoverTab[46348]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1532
		// _ = "end of CoverTab[46348]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1532
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1532
	// _ = "end of CoverTab[46334]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1532
	_go_fuzz_dep_.CoverTab[46335]++
											if n, err = recvmsg(fd, &msg, flags); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1533
		_go_fuzz_dep_.CoverTab[46349]++
												return
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1534
		// _ = "end of CoverTab[46349]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1535
		_go_fuzz_dep_.CoverTab[46350]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1535
		// _ = "end of CoverTab[46350]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1535
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1535
	// _ = "end of CoverTab[46335]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1535
	_go_fuzz_dep_.CoverTab[46336]++
											oobn = int(msg.Controllen)
											recvflags = int(msg.Flags)
											return
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1538
	// _ = "end of CoverTab[46336]"
}

func sendmsgN(fd int, iov []Iovec, oob []byte, ptr unsafe.Pointer, salen _Socklen, flags int) (n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1541
	_go_fuzz_dep_.CoverTab[46351]++
											var msg Msghdr
											msg.Name = (*byte)(ptr)
											msg.Namelen = uint32(salen)
											var dummy byte
											var empty bool
											if len(oob) > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1547
		_go_fuzz_dep_.CoverTab[46356]++
												empty = emptyIovecs(iov)
												if empty {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1549
			_go_fuzz_dep_.CoverTab[46358]++
													var sockType int
													sockType, err = GetsockoptInt(fd, SOL_SOCKET, SO_TYPE)
													if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1552
				_go_fuzz_dep_.CoverTab[46360]++
														return 0, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1553
				// _ = "end of CoverTab[46360]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1554
				_go_fuzz_dep_.CoverTab[46361]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1554
				// _ = "end of CoverTab[46361]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1554
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1554
			// _ = "end of CoverTab[46358]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1554
			_go_fuzz_dep_.CoverTab[46359]++

													if sockType != SOCK_DGRAM {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1556
				_go_fuzz_dep_.CoverTab[46362]++
														var iova [1]Iovec
														iova[0].Base = &dummy
														iova[0].SetLen(1)
														iov = iova[:]
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1560
				// _ = "end of CoverTab[46362]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1561
				_go_fuzz_dep_.CoverTab[46363]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1561
				// _ = "end of CoverTab[46363]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1561
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1561
			// _ = "end of CoverTab[46359]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1562
			_go_fuzz_dep_.CoverTab[46364]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1562
			// _ = "end of CoverTab[46364]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1562
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1562
		// _ = "end of CoverTab[46356]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1562
		_go_fuzz_dep_.CoverTab[46357]++
												msg.Control = &oob[0]
												msg.SetControllen(len(oob))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1564
		// _ = "end of CoverTab[46357]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1565
		_go_fuzz_dep_.CoverTab[46365]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1565
		// _ = "end of CoverTab[46365]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1565
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1565
	// _ = "end of CoverTab[46351]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1565
	_go_fuzz_dep_.CoverTab[46352]++
											if len(iov) > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1566
		_go_fuzz_dep_.CoverTab[46366]++
												msg.Iov = &iov[0]
												msg.SetIovlen(len(iov))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1568
		// _ = "end of CoverTab[46366]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1569
		_go_fuzz_dep_.CoverTab[46367]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1569
		// _ = "end of CoverTab[46367]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1569
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1569
	// _ = "end of CoverTab[46352]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1569
	_go_fuzz_dep_.CoverTab[46353]++
											if n, err = sendmsg(fd, &msg, flags); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1570
		_go_fuzz_dep_.CoverTab[46368]++
												return 0, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1571
		// _ = "end of CoverTab[46368]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1572
		_go_fuzz_dep_.CoverTab[46369]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1572
		// _ = "end of CoverTab[46369]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1572
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1572
	// _ = "end of CoverTab[46353]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1572
	_go_fuzz_dep_.CoverTab[46354]++
											if len(oob) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1573
		_go_fuzz_dep_.CoverTab[46370]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1573
		return empty
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1573
		// _ = "end of CoverTab[46370]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1573
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1573
		_go_fuzz_dep_.CoverTab[46371]++
												n = 0
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1574
		// _ = "end of CoverTab[46371]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1575
		_go_fuzz_dep_.CoverTab[46372]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1575
		// _ = "end of CoverTab[46372]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1575
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1575
	// _ = "end of CoverTab[46354]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1575
	_go_fuzz_dep_.CoverTab[46355]++
											return n, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1576
	// _ = "end of CoverTab[46355]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1580
func BindToDevice(fd int, device string) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1580
	_go_fuzz_dep_.CoverTab[46373]++
											return SetsockoptString(fd, SOL_SOCKET, SO_BINDTODEVICE, device)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1581
	// _ = "end of CoverTab[46373]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1587
func ptracePeek(req int, pid int, addr uintptr, out []byte) (count int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1587
	_go_fuzz_dep_.CoverTab[46374]++

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1594
	var buf [SizeofPtr]byte

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1601
	n := 0
	if addr%SizeofPtr != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1602
		_go_fuzz_dep_.CoverTab[46377]++
												err = ptracePtr(req, pid, addr-addr%SizeofPtr, unsafe.Pointer(&buf[0]))
												if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1604
			_go_fuzz_dep_.CoverTab[46379]++
													return 0, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1605
			// _ = "end of CoverTab[46379]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1606
			_go_fuzz_dep_.CoverTab[46380]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1606
			// _ = "end of CoverTab[46380]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1606
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1606
		// _ = "end of CoverTab[46377]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1606
		_go_fuzz_dep_.CoverTab[46378]++
												n += copy(out, buf[addr%SizeofPtr:])
												out = out[n:]
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1608
		// _ = "end of CoverTab[46378]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1609
		_go_fuzz_dep_.CoverTab[46381]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1609
		// _ = "end of CoverTab[46381]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1609
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1609
	// _ = "end of CoverTab[46374]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1609
	_go_fuzz_dep_.CoverTab[46375]++

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1612
	for len(out) > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1612
		_go_fuzz_dep_.CoverTab[46382]++

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1615
		err = ptracePtr(req, pid, addr+uintptr(n), unsafe.Pointer(&buf[0]))
		if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1616
			_go_fuzz_dep_.CoverTab[46384]++
													return n, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1617
			// _ = "end of CoverTab[46384]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1618
			_go_fuzz_dep_.CoverTab[46385]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1618
			// _ = "end of CoverTab[46385]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1618
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1618
		// _ = "end of CoverTab[46382]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1618
		_go_fuzz_dep_.CoverTab[46383]++
												copied := copy(out, buf[0:])
												n += copied
												out = out[copied:]
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1621
		// _ = "end of CoverTab[46383]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1622
	// _ = "end of CoverTab[46375]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1622
	_go_fuzz_dep_.CoverTab[46376]++

											return n, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1624
	// _ = "end of CoverTab[46376]"
}

func PtracePeekText(pid int, addr uintptr, out []byte) (count int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1627
	_go_fuzz_dep_.CoverTab[46386]++
											return ptracePeek(PTRACE_PEEKTEXT, pid, addr, out)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1628
	// _ = "end of CoverTab[46386]"
}

func PtracePeekData(pid int, addr uintptr, out []byte) (count int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1631
	_go_fuzz_dep_.CoverTab[46387]++
											return ptracePeek(PTRACE_PEEKDATA, pid, addr, out)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1632
	// _ = "end of CoverTab[46387]"
}

func PtracePeekUser(pid int, addr uintptr, out []byte) (count int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1635
	_go_fuzz_dep_.CoverTab[46388]++
											return ptracePeek(PTRACE_PEEKUSR, pid, addr, out)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1636
	// _ = "end of CoverTab[46388]"
}

func ptracePoke(pokeReq int, peekReq int, pid int, addr uintptr, data []byte) (count int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1639
	_go_fuzz_dep_.CoverTab[46389]++

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1644
	n := 0
	if addr%SizeofPtr != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1645
		_go_fuzz_dep_.CoverTab[46393]++
												var buf [SizeofPtr]byte
												err = ptracePtr(peekReq, pid, addr-addr%SizeofPtr, unsafe.Pointer(&buf[0]))
												if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1648
			_go_fuzz_dep_.CoverTab[46396]++
													return 0, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1649
			// _ = "end of CoverTab[46396]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1650
			_go_fuzz_dep_.CoverTab[46397]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1650
			// _ = "end of CoverTab[46397]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1650
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1650
		// _ = "end of CoverTab[46393]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1650
		_go_fuzz_dep_.CoverTab[46394]++
												n += copy(buf[addr%SizeofPtr:], data)
												word := *((*uintptr)(unsafe.Pointer(&buf[0])))
												err = ptrace(pokeReq, pid, addr-addr%SizeofPtr, word)
												if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1654
			_go_fuzz_dep_.CoverTab[46398]++
													return 0, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1655
			// _ = "end of CoverTab[46398]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1656
			_go_fuzz_dep_.CoverTab[46399]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1656
			// _ = "end of CoverTab[46399]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1656
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1656
		// _ = "end of CoverTab[46394]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1656
		_go_fuzz_dep_.CoverTab[46395]++
												data = data[n:]
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1657
		// _ = "end of CoverTab[46395]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1658
		_go_fuzz_dep_.CoverTab[46400]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1658
		// _ = "end of CoverTab[46400]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1658
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1658
	// _ = "end of CoverTab[46389]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1658
	_go_fuzz_dep_.CoverTab[46390]++

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1661
	for len(data) > SizeofPtr {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1661
		_go_fuzz_dep_.CoverTab[46401]++
												word := *((*uintptr)(unsafe.Pointer(&data[0])))
												err = ptrace(pokeReq, pid, addr+uintptr(n), word)
												if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1664
			_go_fuzz_dep_.CoverTab[46403]++
													return n, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1665
			// _ = "end of CoverTab[46403]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1666
			_go_fuzz_dep_.CoverTab[46404]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1666
			// _ = "end of CoverTab[46404]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1666
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1666
		// _ = "end of CoverTab[46401]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1666
		_go_fuzz_dep_.CoverTab[46402]++
												n += SizeofPtr
												data = data[SizeofPtr:]
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1668
		// _ = "end of CoverTab[46402]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1669
	// _ = "end of CoverTab[46390]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1669
	_go_fuzz_dep_.CoverTab[46391]++

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1672
	if len(data) > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1672
		_go_fuzz_dep_.CoverTab[46405]++
												var buf [SizeofPtr]byte
												err = ptracePtr(peekReq, pid, addr+uintptr(n), unsafe.Pointer(&buf[0]))
												if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1675
			_go_fuzz_dep_.CoverTab[46408]++
													return n, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1676
			// _ = "end of CoverTab[46408]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1677
			_go_fuzz_dep_.CoverTab[46409]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1677
			// _ = "end of CoverTab[46409]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1677
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1677
		// _ = "end of CoverTab[46405]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1677
		_go_fuzz_dep_.CoverTab[46406]++
												copy(buf[0:], data)
												word := *((*uintptr)(unsafe.Pointer(&buf[0])))
												err = ptrace(pokeReq, pid, addr+uintptr(n), word)
												if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1681
			_go_fuzz_dep_.CoverTab[46410]++
													return n, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1682
			// _ = "end of CoverTab[46410]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1683
			_go_fuzz_dep_.CoverTab[46411]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1683
			// _ = "end of CoverTab[46411]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1683
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1683
		// _ = "end of CoverTab[46406]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1683
		_go_fuzz_dep_.CoverTab[46407]++
												n += len(data)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1684
		// _ = "end of CoverTab[46407]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1685
		_go_fuzz_dep_.CoverTab[46412]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1685
		// _ = "end of CoverTab[46412]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1685
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1685
	// _ = "end of CoverTab[46391]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1685
	_go_fuzz_dep_.CoverTab[46392]++

											return n, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1687
	// _ = "end of CoverTab[46392]"
}

func PtracePokeText(pid int, addr uintptr, data []byte) (count int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1690
	_go_fuzz_dep_.CoverTab[46413]++
											return ptracePoke(PTRACE_POKETEXT, PTRACE_PEEKTEXT, pid, addr, data)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1691
	// _ = "end of CoverTab[46413]"
}

func PtracePokeData(pid int, addr uintptr, data []byte) (count int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1694
	_go_fuzz_dep_.CoverTab[46414]++
											return ptracePoke(PTRACE_POKEDATA, PTRACE_PEEKDATA, pid, addr, data)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1695
	// _ = "end of CoverTab[46414]"
}

func PtracePokeUser(pid int, addr uintptr, data []byte) (count int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1698
	_go_fuzz_dep_.CoverTab[46415]++
											return ptracePoke(PTRACE_POKEUSR, PTRACE_PEEKUSR, pid, addr, data)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1699
	// _ = "end of CoverTab[46415]"
}

func PtraceGetRegs(pid int, regsout *PtraceRegs) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1702
	_go_fuzz_dep_.CoverTab[46416]++
											return ptracePtr(PTRACE_GETREGS, pid, 0, unsafe.Pointer(regsout))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1703
	// _ = "end of CoverTab[46416]"
}

func PtraceSetRegs(pid int, regs *PtraceRegs) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1706
	_go_fuzz_dep_.CoverTab[46417]++
											return ptracePtr(PTRACE_SETREGS, pid, 0, unsafe.Pointer(regs))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1707
	// _ = "end of CoverTab[46417]"
}

func PtraceSetOptions(pid int, options int) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1710
	_go_fuzz_dep_.CoverTab[46418]++
											return ptrace(PTRACE_SETOPTIONS, pid, 0, uintptr(options))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1711
	// _ = "end of CoverTab[46418]"
}

func PtraceGetEventMsg(pid int) (msg uint, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1714
	_go_fuzz_dep_.CoverTab[46419]++
											var data _C_long
											err = ptracePtr(PTRACE_GETEVENTMSG, pid, 0, unsafe.Pointer(&data))
											msg = uint(data)
											return
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1718
	// _ = "end of CoverTab[46419]"
}

func PtraceCont(pid int, signal int) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1721
	_go_fuzz_dep_.CoverTab[46420]++
											return ptrace(PTRACE_CONT, pid, 0, uintptr(signal))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1722
	// _ = "end of CoverTab[46420]"
}

func PtraceSyscall(pid int, signal int) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1725
	_go_fuzz_dep_.CoverTab[46421]++
											return ptrace(PTRACE_SYSCALL, pid, 0, uintptr(signal))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1726
	// _ = "end of CoverTab[46421]"
}

func PtraceSingleStep(pid int) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1729
	_go_fuzz_dep_.CoverTab[46422]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1729
	return ptrace(PTRACE_SINGLESTEP, pid, 0, 0)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1729
	// _ = "end of CoverTab[46422]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1729
}

func PtraceInterrupt(pid int) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1731
	_go_fuzz_dep_.CoverTab[46423]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1731
	return ptrace(PTRACE_INTERRUPT, pid, 0, 0)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1731
	// _ = "end of CoverTab[46423]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1731
}

func PtraceAttach(pid int) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1733
	_go_fuzz_dep_.CoverTab[46424]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1733
	return ptrace(PTRACE_ATTACH, pid, 0, 0)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1733
	// _ = "end of CoverTab[46424]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1733
}

func PtraceSeize(pid int) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1735
	_go_fuzz_dep_.CoverTab[46425]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1735
	return ptrace(PTRACE_SEIZE, pid, 0, 0)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1735
	// _ = "end of CoverTab[46425]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1735
}

func PtraceDetach(pid int) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1737
	_go_fuzz_dep_.CoverTab[46426]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1737
	return ptrace(PTRACE_DETACH, pid, 0, 0)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1737
	// _ = "end of CoverTab[46426]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1737
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1741
func Reboot(cmd int) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1741
	_go_fuzz_dep_.CoverTab[46427]++
											return reboot(LINUX_REBOOT_MAGIC1, LINUX_REBOOT_MAGIC2, cmd, "")
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1742
	// _ = "end of CoverTab[46427]"
}

func direntIno(buf []byte) (uint64, bool) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1745
	_go_fuzz_dep_.CoverTab[46428]++
											return readInt(buf, unsafe.Offsetof(Dirent{}.Ino), unsafe.Sizeof(Dirent{}.Ino))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1746
	// _ = "end of CoverTab[46428]"
}

func direntReclen(buf []byte) (uint64, bool) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1749
	_go_fuzz_dep_.CoverTab[46429]++
											return readInt(buf, unsafe.Offsetof(Dirent{}.Reclen), unsafe.Sizeof(Dirent{}.Reclen))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1750
	// _ = "end of CoverTab[46429]"
}

func direntNamlen(buf []byte) (uint64, bool) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1753
	_go_fuzz_dep_.CoverTab[46430]++
											reclen, ok := direntReclen(buf)
											if !ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1755
		_go_fuzz_dep_.CoverTab[46432]++
												return 0, false
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1756
		// _ = "end of CoverTab[46432]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1757
		_go_fuzz_dep_.CoverTab[46433]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1757
		// _ = "end of CoverTab[46433]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1757
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1757
	// _ = "end of CoverTab[46430]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1757
	_go_fuzz_dep_.CoverTab[46431]++
											return reclen - uint64(unsafe.Offsetof(Dirent{}.Name)), true
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1758
	// _ = "end of CoverTab[46431]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1763
func Mount(source string, target string, fstype string, flags uintptr, data string) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1763
	_go_fuzz_dep_.CoverTab[46434]++

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1766
	if data == "" {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1766
		_go_fuzz_dep_.CoverTab[46437]++
												return mount(source, target, fstype, flags, nil)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1767
		// _ = "end of CoverTab[46437]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1768
		_go_fuzz_dep_.CoverTab[46438]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1768
		// _ = "end of CoverTab[46438]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1768
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1768
	// _ = "end of CoverTab[46434]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1768
	_go_fuzz_dep_.CoverTab[46435]++
											datap, err := BytePtrFromString(data)
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1770
		_go_fuzz_dep_.CoverTab[46439]++
												return err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1771
		// _ = "end of CoverTab[46439]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1772
		_go_fuzz_dep_.CoverTab[46440]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1772
		// _ = "end of CoverTab[46440]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1772
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1772
	// _ = "end of CoverTab[46435]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1772
	_go_fuzz_dep_.CoverTab[46436]++
											return mount(source, target, fstype, flags, datap)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1773
	// _ = "end of CoverTab[46436]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1782
func MountSetattr(dirfd int, pathname string, flags uint, attr *MountAttr) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1782
	_go_fuzz_dep_.CoverTab[46441]++
											return mountSetattr(dirfd, pathname, flags, attr, unsafe.Sizeof(*attr))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1783
	// _ = "end of CoverTab[46441]"
}

func Sendfile(outfd int, infd int, offset *int64, count int) (written int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1786
	_go_fuzz_dep_.CoverTab[46442]++
											if raceenabled {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1787
		_go_fuzz_dep_.CoverTab[46444]++
												raceReleaseMerge(unsafe.Pointer(&ioSync))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1788
		// _ = "end of CoverTab[46444]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1789
		_go_fuzz_dep_.CoverTab[46445]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1789
		// _ = "end of CoverTab[46445]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1789
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1789
	// _ = "end of CoverTab[46442]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1789
	_go_fuzz_dep_.CoverTab[46443]++
											return sendfile(outfd, infd, offset, count)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1790
	// _ = "end of CoverTab[46443]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1817
func Dup2(oldfd, newfd int) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1817
	_go_fuzz_dep_.CoverTab[46446]++
											return Dup3(oldfd, newfd, 0)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1818
	// _ = "end of CoverTab[46446]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1844
func Getpgrp() (pid int) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1844
	_go_fuzz_dep_.CoverTab[46447]++
											pid, _ = Getpgid(0)
											return
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1846
	// _ = "end of CoverTab[46447]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1889
//go:linkname syscall_prlimit syscall.prlimit
func syscall_prlimit(pid, resource int, newlimit, old *syscall.Rlimit) error

func Prlimit(pid, resource int, newlimit, old *Rlimit) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1892
	_go_fuzz_dep_.CoverTab[46448]++

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1895
	return syscall_prlimit(pid, resource, (*syscall.Rlimit)(newlimit), (*syscall.Rlimit)(old))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1895
	// _ = "end of CoverTab[46448]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1901
func PrctlRetInt(option int, arg2 uintptr, arg3 uintptr, arg4 uintptr, arg5 uintptr) (int, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1901
	_go_fuzz_dep_.CoverTab[46449]++
											ret, _, err := Syscall6(SYS_PRCTL, uintptr(option), uintptr(arg2), uintptr(arg3), uintptr(arg4), uintptr(arg5), 0)
											if err != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1903
		_go_fuzz_dep_.CoverTab[46451]++
												return 0, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1904
		// _ = "end of CoverTab[46451]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1905
		_go_fuzz_dep_.CoverTab[46452]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1905
		// _ = "end of CoverTab[46452]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1905
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1905
	// _ = "end of CoverTab[46449]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1905
	_go_fuzz_dep_.CoverTab[46450]++
											return int(ret), nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1906
	// _ = "end of CoverTab[46450]"
}

func Setuid(uid int) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1909
	_go_fuzz_dep_.CoverTab[46453]++
											return syscall.Setuid(uid)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1910
	// _ = "end of CoverTab[46453]"
}

func Setgid(gid int) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1913
	_go_fuzz_dep_.CoverTab[46454]++
											return syscall.Setgid(gid)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1914
	// _ = "end of CoverTab[46454]"
}

func Setreuid(ruid, euid int) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1917
	_go_fuzz_dep_.CoverTab[46455]++
											return syscall.Setreuid(ruid, euid)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1918
	// _ = "end of CoverTab[46455]"
}

func Setregid(rgid, egid int) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1921
	_go_fuzz_dep_.CoverTab[46456]++
											return syscall.Setregid(rgid, egid)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1922
	// _ = "end of CoverTab[46456]"
}

func Setresuid(ruid, euid, suid int) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1925
	_go_fuzz_dep_.CoverTab[46457]++
											return syscall.Setresuid(ruid, euid, suid)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1926
	// _ = "end of CoverTab[46457]"
}

func Setresgid(rgid, egid, sgid int) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1929
	_go_fuzz_dep_.CoverTab[46458]++
											return syscall.Setresgid(rgid, egid, sgid)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1930
	// _ = "end of CoverTab[46458]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1936
func SetfsgidRetGid(gid int) (int, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1936
	_go_fuzz_dep_.CoverTab[46459]++
											return setfsgid(gid)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1937
	// _ = "end of CoverTab[46459]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1943
func SetfsuidRetUid(uid int) (int, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1943
	_go_fuzz_dep_.CoverTab[46460]++
											return setfsuid(uid)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1944
	// _ = "end of CoverTab[46460]"
}

func Setfsgid(gid int) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1947
	_go_fuzz_dep_.CoverTab[46461]++
											_, err := setfsgid(gid)
											return err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1949
	// _ = "end of CoverTab[46461]"
}

func Setfsuid(uid int) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1952
	_go_fuzz_dep_.CoverTab[46462]++
											_, err := setfsuid(uid)
											return err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1954
	// _ = "end of CoverTab[46462]"
}

func Signalfd(fd int, sigmask *Sigset_t, flags int) (newfd int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1957
	_go_fuzz_dep_.CoverTab[46463]++
											return signalfd(fd, sigmask, _C__NSIG/8, flags)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1958
	// _ = "end of CoverTab[46463]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1995
const minIovec = 8

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1998
func appendBytes(vecs []Iovec, bs [][]byte) []Iovec {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1998
	_go_fuzz_dep_.CoverTab[46464]++
											for _, b := range bs {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:1999
		_go_fuzz_dep_.CoverTab[46466]++
												var v Iovec
												v.SetLen(len(b))
												if len(b) > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2002
			_go_fuzz_dep_.CoverTab[46468]++
													v.Base = &b[0]
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2003
			// _ = "end of CoverTab[46468]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2004
			_go_fuzz_dep_.CoverTab[46469]++
													v.Base = (*byte)(unsafe.Pointer(&_zero))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2005
			// _ = "end of CoverTab[46469]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2006
		// _ = "end of CoverTab[46466]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2006
		_go_fuzz_dep_.CoverTab[46467]++
												vecs = append(vecs, v)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2007
		// _ = "end of CoverTab[46467]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2008
	// _ = "end of CoverTab[46464]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2008
	_go_fuzz_dep_.CoverTab[46465]++
											return vecs
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2009
	// _ = "end of CoverTab[46465]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2013
func offs2lohi(offs int64) (lo, hi uintptr) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2013
	_go_fuzz_dep_.CoverTab[46470]++
											const longBits = SizeofLong * 8
											return uintptr(offs), uintptr(uint64(offs) >> (longBits - 1) >> 1)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2015
	// _ = "end of CoverTab[46470]"
}

func Readv(fd int, iovs [][]byte) (n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2018
	_go_fuzz_dep_.CoverTab[46471]++
											iovecs := make([]Iovec, 0, minIovec)
											iovecs = appendBytes(iovecs, iovs)
											n, err = readv(fd, iovecs)
											readvRacedetect(iovecs, n, err)
											return n, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2023
	// _ = "end of CoverTab[46471]"
}

func Preadv(fd int, iovs [][]byte, offset int64) (n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2026
	_go_fuzz_dep_.CoverTab[46472]++
											iovecs := make([]Iovec, 0, minIovec)
											iovecs = appendBytes(iovecs, iovs)
											lo, hi := offs2lohi(offset)
											n, err = preadv(fd, iovecs, lo, hi)
											readvRacedetect(iovecs, n, err)
											return n, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2032
	// _ = "end of CoverTab[46472]"
}

func Preadv2(fd int, iovs [][]byte, offset int64, flags int) (n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2035
	_go_fuzz_dep_.CoverTab[46473]++
											iovecs := make([]Iovec, 0, minIovec)
											iovecs = appendBytes(iovecs, iovs)
											lo, hi := offs2lohi(offset)
											n, err = preadv2(fd, iovecs, lo, hi, flags)
											readvRacedetect(iovecs, n, err)
											return n, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2041
	// _ = "end of CoverTab[46473]"
}

func readvRacedetect(iovecs []Iovec, n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2044
	_go_fuzz_dep_.CoverTab[46474]++
											if !raceenabled {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2045
		_go_fuzz_dep_.CoverTab[46477]++
												return
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2046
		// _ = "end of CoverTab[46477]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2047
		_go_fuzz_dep_.CoverTab[46478]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2047
		// _ = "end of CoverTab[46478]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2047
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2047
	// _ = "end of CoverTab[46474]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2047
	_go_fuzz_dep_.CoverTab[46475]++
											for i := 0; n > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2048
		_go_fuzz_dep_.CoverTab[46479]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2048
		return i < len(iovecs)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2048
		// _ = "end of CoverTab[46479]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2048
	}(); i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2048
		_go_fuzz_dep_.CoverTab[46480]++
												m := int(iovecs[i].Len)
												if m > n {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2050
			_go_fuzz_dep_.CoverTab[46482]++
													m = n
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2051
			// _ = "end of CoverTab[46482]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2052
			_go_fuzz_dep_.CoverTab[46483]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2052
			// _ = "end of CoverTab[46483]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2052
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2052
		// _ = "end of CoverTab[46480]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2052
		_go_fuzz_dep_.CoverTab[46481]++
												n -= m
												if m > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2054
			_go_fuzz_dep_.CoverTab[46484]++
													raceWriteRange(unsafe.Pointer(iovecs[i].Base), m)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2055
			// _ = "end of CoverTab[46484]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2056
			_go_fuzz_dep_.CoverTab[46485]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2056
			// _ = "end of CoverTab[46485]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2056
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2056
		// _ = "end of CoverTab[46481]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2057
	// _ = "end of CoverTab[46475]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2057
	_go_fuzz_dep_.CoverTab[46476]++
											if err == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2058
		_go_fuzz_dep_.CoverTab[46486]++
												raceAcquire(unsafe.Pointer(&ioSync))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2059
		// _ = "end of CoverTab[46486]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2060
		_go_fuzz_dep_.CoverTab[46487]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2060
		// _ = "end of CoverTab[46487]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2060
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2060
	// _ = "end of CoverTab[46476]"
}

func Writev(fd int, iovs [][]byte) (n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2063
	_go_fuzz_dep_.CoverTab[46488]++
											iovecs := make([]Iovec, 0, minIovec)
											iovecs = appendBytes(iovecs, iovs)
											if raceenabled {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2066
		_go_fuzz_dep_.CoverTab[46490]++
												raceReleaseMerge(unsafe.Pointer(&ioSync))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2067
		// _ = "end of CoverTab[46490]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2068
		_go_fuzz_dep_.CoverTab[46491]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2068
		// _ = "end of CoverTab[46491]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2068
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2068
	// _ = "end of CoverTab[46488]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2068
	_go_fuzz_dep_.CoverTab[46489]++
											n, err = writev(fd, iovecs)
											writevRacedetect(iovecs, n)
											return n, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2071
	// _ = "end of CoverTab[46489]"
}

func Pwritev(fd int, iovs [][]byte, offset int64) (n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2074
	_go_fuzz_dep_.CoverTab[46492]++
											iovecs := make([]Iovec, 0, minIovec)
											iovecs = appendBytes(iovecs, iovs)
											if raceenabled {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2077
		_go_fuzz_dep_.CoverTab[46494]++
												raceReleaseMerge(unsafe.Pointer(&ioSync))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2078
		// _ = "end of CoverTab[46494]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2079
		_go_fuzz_dep_.CoverTab[46495]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2079
		// _ = "end of CoverTab[46495]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2079
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2079
	// _ = "end of CoverTab[46492]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2079
	_go_fuzz_dep_.CoverTab[46493]++
											lo, hi := offs2lohi(offset)
											n, err = pwritev(fd, iovecs, lo, hi)
											writevRacedetect(iovecs, n)
											return n, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2083
	// _ = "end of CoverTab[46493]"
}

func Pwritev2(fd int, iovs [][]byte, offset int64, flags int) (n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2086
	_go_fuzz_dep_.CoverTab[46496]++
											iovecs := make([]Iovec, 0, minIovec)
											iovecs = appendBytes(iovecs, iovs)
											if raceenabled {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2089
		_go_fuzz_dep_.CoverTab[46498]++
												raceReleaseMerge(unsafe.Pointer(&ioSync))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2090
		// _ = "end of CoverTab[46498]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2091
		_go_fuzz_dep_.CoverTab[46499]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2091
		// _ = "end of CoverTab[46499]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2091
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2091
	// _ = "end of CoverTab[46496]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2091
	_go_fuzz_dep_.CoverTab[46497]++
											lo, hi := offs2lohi(offset)
											n, err = pwritev2(fd, iovecs, lo, hi, flags)
											writevRacedetect(iovecs, n)
											return n, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2095
	// _ = "end of CoverTab[46497]"
}

func writevRacedetect(iovecs []Iovec, n int) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2098
	_go_fuzz_dep_.CoverTab[46500]++
											if !raceenabled {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2099
		_go_fuzz_dep_.CoverTab[46502]++
												return
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2100
		// _ = "end of CoverTab[46502]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2101
		_go_fuzz_dep_.CoverTab[46503]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2101
		// _ = "end of CoverTab[46503]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2101
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2101
	// _ = "end of CoverTab[46500]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2101
	_go_fuzz_dep_.CoverTab[46501]++
											for i := 0; n > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2102
		_go_fuzz_dep_.CoverTab[46504]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2102
		return i < len(iovecs)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2102
		// _ = "end of CoverTab[46504]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2102
	}(); i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2102
		_go_fuzz_dep_.CoverTab[46505]++
												m := int(iovecs[i].Len)
												if m > n {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2104
			_go_fuzz_dep_.CoverTab[46507]++
													m = n
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2105
			// _ = "end of CoverTab[46507]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2106
			_go_fuzz_dep_.CoverTab[46508]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2106
			// _ = "end of CoverTab[46508]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2106
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2106
		// _ = "end of CoverTab[46505]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2106
		_go_fuzz_dep_.CoverTab[46506]++
												n -= m
												if m > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2108
			_go_fuzz_dep_.CoverTab[46509]++
													raceReadRange(unsafe.Pointer(iovecs[i].Base), m)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2109
			// _ = "end of CoverTab[46509]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2110
			_go_fuzz_dep_.CoverTab[46510]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2110
			// _ = "end of CoverTab[46510]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2110
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2110
		// _ = "end of CoverTab[46506]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2111
	// _ = "end of CoverTab[46501]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2117
var mapper = &mmapper{
	active:	make(map[*byte][]byte),
	mmap:	mmap,
	munmap:	munmap,
}

func Mmap(fd int, offset int64, length int, prot int, flags int) (data []byte, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2123
	_go_fuzz_dep_.CoverTab[46511]++
											return mapper.Mmap(fd, offset, length, prot, flags)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2124
	// _ = "end of CoverTab[46511]"
}

func Munmap(b []byte) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2127
	_go_fuzz_dep_.CoverTab[46512]++
											return mapper.Munmap(b)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2128
	// _ = "end of CoverTab[46512]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2141
func Vmsplice(fd int, iovs []Iovec, flags int) (int, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2141
	_go_fuzz_dep_.CoverTab[46513]++
											var p unsafe.Pointer
											if len(iovs) > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2143
		_go_fuzz_dep_.CoverTab[46516]++
												p = unsafe.Pointer(&iovs[0])
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2144
		// _ = "end of CoverTab[46516]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2145
		_go_fuzz_dep_.CoverTab[46517]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2145
		// _ = "end of CoverTab[46517]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2145
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2145
	// _ = "end of CoverTab[46513]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2145
	_go_fuzz_dep_.CoverTab[46514]++

											n, _, errno := Syscall6(SYS_VMSPLICE, uintptr(fd), uintptr(p), uintptr(len(iovs)), uintptr(flags), 0, 0)
											if errno != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2148
		_go_fuzz_dep_.CoverTab[46518]++
												return 0, syscall.Errno(errno)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2149
		// _ = "end of CoverTab[46518]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2150
		_go_fuzz_dep_.CoverTab[46519]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2150
		// _ = "end of CoverTab[46519]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2150
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2150
	// _ = "end of CoverTab[46514]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2150
	_go_fuzz_dep_.CoverTab[46515]++

											return int(n), nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2152
	// _ = "end of CoverTab[46515]"
}

func isGroupMember(gid int) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2155
	_go_fuzz_dep_.CoverTab[46520]++
											groups, err := Getgroups()
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2157
		_go_fuzz_dep_.CoverTab[46523]++
												return false
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2158
		// _ = "end of CoverTab[46523]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2159
		_go_fuzz_dep_.CoverTab[46524]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2159
		// _ = "end of CoverTab[46524]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2159
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2159
	// _ = "end of CoverTab[46520]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2159
	_go_fuzz_dep_.CoverTab[46521]++

											for _, g := range groups {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2161
		_go_fuzz_dep_.CoverTab[46525]++
												if g == gid {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2162
			_go_fuzz_dep_.CoverTab[46526]++
													return true
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2163
			// _ = "end of CoverTab[46526]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2164
			_go_fuzz_dep_.CoverTab[46527]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2164
			// _ = "end of CoverTab[46527]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2164
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2164
		// _ = "end of CoverTab[46525]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2165
	// _ = "end of CoverTab[46521]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2165
	_go_fuzz_dep_.CoverTab[46522]++
											return false
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2166
	// _ = "end of CoverTab[46522]"
}

func isCapDacOverrideSet() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2169
	_go_fuzz_dep_.CoverTab[46528]++
											hdr := CapUserHeader{Version: LINUX_CAPABILITY_VERSION_3}
											data := [2]CapUserData{}
											err := Capget(&hdr, &data[0])

											return err == nil && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2174
		_go_fuzz_dep_.CoverTab[46529]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2174
		return data[0].Effective&(1<<CAP_DAC_OVERRIDE) != 0
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2174
		// _ = "end of CoverTab[46529]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2174
	}()
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2174
	// _ = "end of CoverTab[46528]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2180
func Faccessat(dirfd int, path string, mode uint32, flags int) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2180
	_go_fuzz_dep_.CoverTab[46530]++
											if flags == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2181
		_go_fuzz_dep_.CoverTab[46540]++
												return faccessat(dirfd, path, mode)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2182
		// _ = "end of CoverTab[46540]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2183
		_go_fuzz_dep_.CoverTab[46541]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2183
		// _ = "end of CoverTab[46541]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2183
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2183
	// _ = "end of CoverTab[46530]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2183
	_go_fuzz_dep_.CoverTab[46531]++

											if err := Faccessat2(dirfd, path, mode, flags); err != ENOSYS && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2185
		_go_fuzz_dep_.CoverTab[46542]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2185
		return err != EPERM
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2185
		// _ = "end of CoverTab[46542]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2185
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2185
		_go_fuzz_dep_.CoverTab[46543]++
												return err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2186
		// _ = "end of CoverTab[46543]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2187
		_go_fuzz_dep_.CoverTab[46544]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2187
		// _ = "end of CoverTab[46544]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2187
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2187
	// _ = "end of CoverTab[46531]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2187
	_go_fuzz_dep_.CoverTab[46532]++

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2195
	if flags & ^(AT_SYMLINK_NOFOLLOW|AT_EACCESS) != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2195
		_go_fuzz_dep_.CoverTab[46545]++
												return EINVAL
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2196
		// _ = "end of CoverTab[46545]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2197
		_go_fuzz_dep_.CoverTab[46546]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2197
		// _ = "end of CoverTab[46546]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2197
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2197
	// _ = "end of CoverTab[46532]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2197
	_go_fuzz_dep_.CoverTab[46533]++

											var st Stat_t
											if err := Fstatat(dirfd, path, &st, flags&AT_SYMLINK_NOFOLLOW); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2200
		_go_fuzz_dep_.CoverTab[46547]++
												return err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2201
		// _ = "end of CoverTab[46547]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2202
		_go_fuzz_dep_.CoverTab[46548]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2202
		// _ = "end of CoverTab[46548]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2202
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2202
	// _ = "end of CoverTab[46533]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2202
	_go_fuzz_dep_.CoverTab[46534]++

											mode &= 7
											if mode == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2205
		_go_fuzz_dep_.CoverTab[46549]++
												return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2206
		// _ = "end of CoverTab[46549]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2207
		_go_fuzz_dep_.CoverTab[46550]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2207
		// _ = "end of CoverTab[46550]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2207
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2207
	// _ = "end of CoverTab[46534]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2207
	_go_fuzz_dep_.CoverTab[46535]++

											var uid int
											if flags&AT_EACCESS != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2210
		_go_fuzz_dep_.CoverTab[46551]++
												uid = Geteuid()
												if uid != 0 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2212
			_go_fuzz_dep_.CoverTab[46552]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2212
			return isCapDacOverrideSet()
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2212
			// _ = "end of CoverTab[46552]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2212
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2212
			_go_fuzz_dep_.CoverTab[46553]++

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2216
			uid = 0
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2216
			// _ = "end of CoverTab[46553]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2217
			_go_fuzz_dep_.CoverTab[46554]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2217
			// _ = "end of CoverTab[46554]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2217
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2217
		// _ = "end of CoverTab[46551]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2218
		_go_fuzz_dep_.CoverTab[46555]++
												uid = Getuid()
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2219
		// _ = "end of CoverTab[46555]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2220
	// _ = "end of CoverTab[46535]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2220
	_go_fuzz_dep_.CoverTab[46536]++

											if uid == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2222
		_go_fuzz_dep_.CoverTab[46556]++
												if mode&1 == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2223
			_go_fuzz_dep_.CoverTab[46559]++

													return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2225
			// _ = "end of CoverTab[46559]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2226
			_go_fuzz_dep_.CoverTab[46560]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2226
			// _ = "end of CoverTab[46560]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2226
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2226
		// _ = "end of CoverTab[46556]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2226
		_go_fuzz_dep_.CoverTab[46557]++
												if st.Mode&0111 != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2227
			_go_fuzz_dep_.CoverTab[46561]++

													return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2229
			// _ = "end of CoverTab[46561]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2230
			_go_fuzz_dep_.CoverTab[46562]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2230
			// _ = "end of CoverTab[46562]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2230
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2230
		// _ = "end of CoverTab[46557]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2230
		_go_fuzz_dep_.CoverTab[46558]++
												return EACCES
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2231
		// _ = "end of CoverTab[46558]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2232
		_go_fuzz_dep_.CoverTab[46563]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2232
		// _ = "end of CoverTab[46563]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2232
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2232
	// _ = "end of CoverTab[46536]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2232
	_go_fuzz_dep_.CoverTab[46537]++

											var fmode uint32
											if uint32(uid) == st.Uid {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2235
		_go_fuzz_dep_.CoverTab[46564]++
												fmode = (st.Mode >> 6) & 7
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2236
		// _ = "end of CoverTab[46564]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2237
		_go_fuzz_dep_.CoverTab[46565]++
												var gid int
												if flags&AT_EACCESS != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2239
			_go_fuzz_dep_.CoverTab[46567]++
													gid = Getegid()
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2240
			// _ = "end of CoverTab[46567]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2241
			_go_fuzz_dep_.CoverTab[46568]++
													gid = Getgid()
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2242
			// _ = "end of CoverTab[46568]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2243
		// _ = "end of CoverTab[46565]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2243
		_go_fuzz_dep_.CoverTab[46566]++

												if uint32(gid) == st.Gid || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2245
			_go_fuzz_dep_.CoverTab[46569]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2245
			return isGroupMember(int(st.Gid))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2245
			// _ = "end of CoverTab[46569]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2245
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2245
			_go_fuzz_dep_.CoverTab[46570]++
													fmode = (st.Mode >> 3) & 7
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2246
			// _ = "end of CoverTab[46570]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2247
			_go_fuzz_dep_.CoverTab[46571]++
													fmode = st.Mode & 7
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2248
			// _ = "end of CoverTab[46571]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2249
		// _ = "end of CoverTab[46566]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2250
	// _ = "end of CoverTab[46537]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2250
	_go_fuzz_dep_.CoverTab[46538]++

											if fmode&mode == mode {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2252
		_go_fuzz_dep_.CoverTab[46572]++
												return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2253
		// _ = "end of CoverTab[46572]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2254
		_go_fuzz_dep_.CoverTab[46573]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2254
		// _ = "end of CoverTab[46573]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2254
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2254
	// _ = "end of CoverTab[46538]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2254
	_go_fuzz_dep_.CoverTab[46539]++

											return EACCES
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2256
	// _ = "end of CoverTab[46539]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2267
type fileHandle struct {
	Bytes	uint32
	Type	int32
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2275
type FileHandle struct {
	*fileHandle
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2280
func NewFileHandle(handleType int32, handle []byte) FileHandle {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2280
	_go_fuzz_dep_.CoverTab[46574]++
											const hdrSize = unsafe.Sizeof(fileHandle{})
											buf := make([]byte, hdrSize+uintptr(len(handle)))
											copy(buf[hdrSize:], handle)
											fh := (*fileHandle)(unsafe.Pointer(&buf[0]))
											fh.Type = handleType
											fh.Bytes = uint32(len(handle))
											return FileHandle{fh}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2287
	// _ = "end of CoverTab[46574]"
}

func (fh *FileHandle) Size() int {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2290
	_go_fuzz_dep_.CoverTab[46575]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2290
	return int(fh.fileHandle.Bytes)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2290
	// _ = "end of CoverTab[46575]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2290
}
func (fh *FileHandle) Type() int32 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2291
	_go_fuzz_dep_.CoverTab[46576]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2291
	return fh.fileHandle.Type
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2291
	// _ = "end of CoverTab[46576]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2291
}
func (fh *FileHandle) Bytes() []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2292
	_go_fuzz_dep_.CoverTab[46577]++
											n := fh.Size()
											if n == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2294
		_go_fuzz_dep_.CoverTab[46579]++
												return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2295
		// _ = "end of CoverTab[46579]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2296
		_go_fuzz_dep_.CoverTab[46580]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2296
		// _ = "end of CoverTab[46580]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2296
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2296
	// _ = "end of CoverTab[46577]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2296
	_go_fuzz_dep_.CoverTab[46578]++
											return unsafe.Slice((*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&fh.fileHandle.Type))+4)), n)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2297
	// _ = "end of CoverTab[46578]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2302
func NameToHandleAt(dirfd int, path string, flags int) (handle FileHandle, mountID int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2302
	_go_fuzz_dep_.CoverTab[46581]++
											var mid _C_int

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2306
	size := uint32(32 + unsafe.Sizeof(fileHandle{}))
	didResize := false
	for {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2308
		_go_fuzz_dep_.CoverTab[46582]++
												buf := make([]byte, size)
												fh := (*fileHandle)(unsafe.Pointer(&buf[0]))
												fh.Bytes = size - uint32(unsafe.Sizeof(fileHandle{}))
												err = nameToHandleAt(dirfd, path, fh, &mid, flags)
												if err == EOVERFLOW {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2313
			_go_fuzz_dep_.CoverTab[46585]++
													if didResize {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2314
				_go_fuzz_dep_.CoverTab[46587]++

														return
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2316
				// _ = "end of CoverTab[46587]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2317
				_go_fuzz_dep_.CoverTab[46588]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2317
				// _ = "end of CoverTab[46588]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2317
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2317
			// _ = "end of CoverTab[46585]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2317
			_go_fuzz_dep_.CoverTab[46586]++
													didResize = true
													size = fh.Bytes + uint32(unsafe.Sizeof(fileHandle{}))
													continue
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2320
			// _ = "end of CoverTab[46586]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2321
			_go_fuzz_dep_.CoverTab[46589]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2321
			// _ = "end of CoverTab[46589]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2321
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2321
		// _ = "end of CoverTab[46582]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2321
		_go_fuzz_dep_.CoverTab[46583]++
												if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2322
			_go_fuzz_dep_.CoverTab[46590]++
													return
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2323
			// _ = "end of CoverTab[46590]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2324
			_go_fuzz_dep_.CoverTab[46591]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2324
			// _ = "end of CoverTab[46591]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2324
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2324
		// _ = "end of CoverTab[46583]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2324
		_go_fuzz_dep_.CoverTab[46584]++
												return FileHandle{fh}, int(mid), nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2325
		// _ = "end of CoverTab[46584]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2326
	// _ = "end of CoverTab[46581]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2331
func OpenByHandleAt(mountFD int, handle FileHandle, flags int) (fd int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2331
	_go_fuzz_dep_.CoverTab[46592]++
											return openByHandleAt(mountFD, handle.fileHandle, flags)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2332
	// _ = "end of CoverTab[46592]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2337
func Klogset(typ int, arg int) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2337
	_go_fuzz_dep_.CoverTab[46593]++
											var p unsafe.Pointer
											_, _, errno := Syscall(SYS_SYSLOG, uintptr(typ), uintptr(p), uintptr(arg))
											if errno != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2340
		_go_fuzz_dep_.CoverTab[46595]++
												return errnoErr(errno)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2341
		// _ = "end of CoverTab[46595]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2342
		_go_fuzz_dep_.CoverTab[46596]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2342
		// _ = "end of CoverTab[46596]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2342
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2342
	// _ = "end of CoverTab[46593]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2342
	_go_fuzz_dep_.CoverTab[46594]++
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2343
	// _ = "end of CoverTab[46594]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2350
type RemoteIovec struct {
	Base	uintptr
	Len	int
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2371
func MakeItimerval(interval, value time.Duration) Itimerval {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2371
	_go_fuzz_dep_.CoverTab[46597]++
											return Itimerval{
		Interval:	NsecToTimeval(interval.Nanoseconds()),
		Value:		NsecToTimeval(value.Nanoseconds()),
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2375
	// _ = "end of CoverTab[46597]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2380
type ItimerWhich int

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2383
const (
	ItimerReal	ItimerWhich	= ITIMER_REAL
	ItimerVirtual	ItimerWhich	= ITIMER_VIRTUAL
	ItimerProf	ItimerWhich	= ITIMER_PROF
)

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2391
func Getitimer(which ItimerWhich) (Itimerval, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2391
	_go_fuzz_dep_.CoverTab[46598]++
											var it Itimerval
											if err := getitimer(int(which), &it); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2393
		_go_fuzz_dep_.CoverTab[46600]++
												return Itimerval{}, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2394
		// _ = "end of CoverTab[46600]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2395
		_go_fuzz_dep_.CoverTab[46601]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2395
		// _ = "end of CoverTab[46601]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2395
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2395
	// _ = "end of CoverTab[46598]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2395
	_go_fuzz_dep_.CoverTab[46599]++

											return it, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2397
	// _ = "end of CoverTab[46599]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2404
func Setitimer(which ItimerWhich, it Itimerval) (Itimerval, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2404
	_go_fuzz_dep_.CoverTab[46602]++
											var prev Itimerval
											if err := setitimer(int(which), &it, &prev); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2406
		_go_fuzz_dep_.CoverTab[46604]++
												return Itimerval{}, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2407
		// _ = "end of CoverTab[46604]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2408
		_go_fuzz_dep_.CoverTab[46605]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2408
		// _ = "end of CoverTab[46605]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2408
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2408
	// _ = "end of CoverTab[46602]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2408
	_go_fuzz_dep_.CoverTab[46603]++

											return prev, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2410
	// _ = "end of CoverTab[46603]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2415
func PthreadSigmask(how int, set, oldset *Sigset_t) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2415
	_go_fuzz_dep_.CoverTab[46606]++
											if oldset != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2416
		_go_fuzz_dep_.CoverTab[46608]++

												*oldset = Sigset_t{}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2418
		// _ = "end of CoverTab[46608]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2419
		_go_fuzz_dep_.CoverTab[46609]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2419
		// _ = "end of CoverTab[46609]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2419
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2419
	// _ = "end of CoverTab[46606]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2419
	_go_fuzz_dep_.CoverTab[46607]++
											return rtSigprocmask(how, set, oldset, _C__NSIG/8)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2420
	// _ = "end of CoverTab[46607]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2421
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux.go:2421
var _ = _go_fuzz_dep_.CoverTab
