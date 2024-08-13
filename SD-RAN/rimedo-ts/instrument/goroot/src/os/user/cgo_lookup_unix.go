// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build (cgo || darwin) && !osusergo && unix && !android

//line /usr/local/go/src/os/user/cgo_lookup_unix.go:7
package user

//line /usr/local/go/src/os/user/cgo_lookup_unix.go:7
import (
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:7
)
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:7
import (
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:7
)

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"unsafe"
)

func current() (*User, error) {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:18
	_go_fuzz_dep_.CoverTab[83163]++
							return lookupUnixUid(syscall.Getuid())
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:19
	// _ = "end of CoverTab[83163]"
}

func lookupUser(username string) (*User, error) {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:22
	_go_fuzz_dep_.CoverTab[83164]++
							var pwd _C_struct_passwd
							var found bool
							nameC := make([]byte, len(username)+1)
							copy(nameC, username)

							err := retryWithBuffer(userBuffer, func(buf []byte) syscall.Errno {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:28
		_go_fuzz_dep_.CoverTab[83168]++
								var errno syscall.Errno
								pwd, found, errno = _C_getpwnam_r((*_C_char)(unsafe.Pointer(&nameC[0])),
			(*_C_char)(unsafe.Pointer(&buf[0])), _C_size_t(len(buf)))
								return errno
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:32
		// _ = "end of CoverTab[83168]"
	})
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:33
	// _ = "end of CoverTab[83164]"
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:33
	_go_fuzz_dep_.CoverTab[83165]++
							if err != nil {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:34
		_go_fuzz_dep_.CoverTab[83169]++
								return nil, fmt.Errorf("user: lookup username %s: %v", username, err)
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:35
		// _ = "end of CoverTab[83169]"
	} else {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:36
		_go_fuzz_dep_.CoverTab[83170]++
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:36
		// _ = "end of CoverTab[83170]"
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:36
	}
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:36
	// _ = "end of CoverTab[83165]"
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:36
	_go_fuzz_dep_.CoverTab[83166]++
							if !found {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:37
		_go_fuzz_dep_.CoverTab[83171]++
								return nil, UnknownUserError(username)
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:38
		// _ = "end of CoverTab[83171]"
	} else {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:39
		_go_fuzz_dep_.CoverTab[83172]++
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:39
		// _ = "end of CoverTab[83172]"
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:39
	}
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:39
	// _ = "end of CoverTab[83166]"
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:39
	_go_fuzz_dep_.CoverTab[83167]++
							return buildUser(&pwd), err
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:40
	// _ = "end of CoverTab[83167]"
}

func lookupUserId(uid string) (*User, error) {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:43
	_go_fuzz_dep_.CoverTab[83173]++
							i, e := strconv.Atoi(uid)
							if e != nil {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:45
		_go_fuzz_dep_.CoverTab[83175]++
								return nil, e
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:46
		// _ = "end of CoverTab[83175]"
	} else {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:47
		_go_fuzz_dep_.CoverTab[83176]++
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:47
		// _ = "end of CoverTab[83176]"
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:47
	}
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:47
	// _ = "end of CoverTab[83173]"
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:47
	_go_fuzz_dep_.CoverTab[83174]++
							return lookupUnixUid(i)
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:48
	// _ = "end of CoverTab[83174]"
}

func lookupUnixUid(uid int) (*User, error) {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:51
	_go_fuzz_dep_.CoverTab[83177]++
							var pwd _C_struct_passwd
							var found bool

							err := retryWithBuffer(userBuffer, func(buf []byte) syscall.Errno {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:55
		_go_fuzz_dep_.CoverTab[83181]++
								var errno syscall.Errno
								pwd, found, errno = _C_getpwuid_r(_C_uid_t(uid),
			(*_C_char)(unsafe.Pointer(&buf[0])), _C_size_t(len(buf)))
								return errno
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:59
		// _ = "end of CoverTab[83181]"
	})
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:60
	// _ = "end of CoverTab[83177]"
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:60
	_go_fuzz_dep_.CoverTab[83178]++
							if err != nil {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:61
		_go_fuzz_dep_.CoverTab[83182]++
								return nil, fmt.Errorf("user: lookup userid %d: %v", uid, err)
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:62
		// _ = "end of CoverTab[83182]"
	} else {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:63
		_go_fuzz_dep_.CoverTab[83183]++
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:63
		// _ = "end of CoverTab[83183]"
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:63
	}
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:63
	// _ = "end of CoverTab[83178]"
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:63
	_go_fuzz_dep_.CoverTab[83179]++
							if !found {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:64
		_go_fuzz_dep_.CoverTab[83184]++
								return nil, UnknownUserIdError(uid)
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:65
		// _ = "end of CoverTab[83184]"
	} else {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:66
		_go_fuzz_dep_.CoverTab[83185]++
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:66
		// _ = "end of CoverTab[83185]"
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:66
	}
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:66
	// _ = "end of CoverTab[83179]"
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:66
	_go_fuzz_dep_.CoverTab[83180]++
							return buildUser(&pwd), nil
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:67
	// _ = "end of CoverTab[83180]"
}

func buildUser(pwd *_C_struct_passwd) *User {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:70
	_go_fuzz_dep_.CoverTab[83186]++
							u := &User{
		Uid:		strconv.FormatUint(uint64(_C_pw_uid(pwd)), 10),
		Gid:		strconv.FormatUint(uint64(_C_pw_gid(pwd)), 10),
		Username:	_C_GoString(_C_pw_name(pwd)),
		Name:		_C_GoString(_C_pw_gecos(pwd)),
		HomeDir:	_C_GoString(_C_pw_dir(pwd)),
	}

//line /usr/local/go/src/os/user/cgo_lookup_unix.go:82
	u.Name, _, _ = strings.Cut(u.Name, ",")
							return u
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:83
	// _ = "end of CoverTab[83186]"
}

func lookupGroup(groupname string) (*Group, error) {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:86
	_go_fuzz_dep_.CoverTab[83187]++
							var grp _C_struct_group
							var found bool

							cname := make([]byte, len(groupname)+1)
							copy(cname, groupname)

							err := retryWithBuffer(groupBuffer, func(buf []byte) syscall.Errno {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:93
		_go_fuzz_dep_.CoverTab[83191]++
								var errno syscall.Errno
								grp, found, errno = _C_getgrnam_r((*_C_char)(unsafe.Pointer(&cname[0])),
			(*_C_char)(unsafe.Pointer(&buf[0])), _C_size_t(len(buf)))
								return errno
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:97
		// _ = "end of CoverTab[83191]"
	})
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:98
	// _ = "end of CoverTab[83187]"
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:98
	_go_fuzz_dep_.CoverTab[83188]++
							if err != nil {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:99
			_go_fuzz_dep_.CoverTab[83192]++
									return nil, fmt.Errorf("user: lookup groupname %s: %v", groupname, err)
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:100
		// _ = "end of CoverTab[83192]"
	} else {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:101
		_go_fuzz_dep_.CoverTab[83193]++
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:101
		// _ = "end of CoverTab[83193]"
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:101
	}
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:101
	// _ = "end of CoverTab[83188]"
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:101
	_go_fuzz_dep_.CoverTab[83189]++
								if !found {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:102
		_go_fuzz_dep_.CoverTab[83194]++
									return nil, UnknownGroupError(groupname)
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:103
		// _ = "end of CoverTab[83194]"
	} else {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:104
		_go_fuzz_dep_.CoverTab[83195]++
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:104
		// _ = "end of CoverTab[83195]"
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:104
	}
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:104
	// _ = "end of CoverTab[83189]"
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:104
	_go_fuzz_dep_.CoverTab[83190]++
								return buildGroup(&grp), nil
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:105
	// _ = "end of CoverTab[83190]"
}

func lookupGroupId(gid string) (*Group, error) {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:108
	_go_fuzz_dep_.CoverTab[83196]++
								i, e := strconv.Atoi(gid)
								if e != nil {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:110
		_go_fuzz_dep_.CoverTab[83198]++
									return nil, e
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:111
		// _ = "end of CoverTab[83198]"
	} else {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:112
		_go_fuzz_dep_.CoverTab[83199]++
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:112
		// _ = "end of CoverTab[83199]"
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:112
	}
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:112
	// _ = "end of CoverTab[83196]"
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:112
	_go_fuzz_dep_.CoverTab[83197]++
								return lookupUnixGid(i)
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:113
	// _ = "end of CoverTab[83197]"
}

func lookupUnixGid(gid int) (*Group, error) {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:116
	_go_fuzz_dep_.CoverTab[83200]++
								var grp _C_struct_group
								var found bool

								err := retryWithBuffer(groupBuffer, func(buf []byte) syscall.Errno {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:120
		_go_fuzz_dep_.CoverTab[83204]++
									var errno syscall.Errno
									grp, found, errno = _C_getgrgid_r(_C_gid_t(gid),
			(*_C_char)(unsafe.Pointer(&buf[0])), _C_size_t(len(buf)))
									return syscall.Errno(errno)
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:124
		// _ = "end of CoverTab[83204]"
	})
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:125
	// _ = "end of CoverTab[83200]"
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:125
	_go_fuzz_dep_.CoverTab[83201]++
								if err != nil {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:126
		_go_fuzz_dep_.CoverTab[83205]++
									return nil, fmt.Errorf("user: lookup groupid %d: %v", gid, err)
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:127
		// _ = "end of CoverTab[83205]"
	} else {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:128
		_go_fuzz_dep_.CoverTab[83206]++
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:128
		// _ = "end of CoverTab[83206]"
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:128
	}
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:128
	// _ = "end of CoverTab[83201]"
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:128
	_go_fuzz_dep_.CoverTab[83202]++
								if !found {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:129
		_go_fuzz_dep_.CoverTab[83207]++
									return nil, UnknownGroupIdError(strconv.Itoa(gid))
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:130
		// _ = "end of CoverTab[83207]"
	} else {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:131
		_go_fuzz_dep_.CoverTab[83208]++
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:131
		// _ = "end of CoverTab[83208]"
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:131
	}
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:131
	// _ = "end of CoverTab[83202]"
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:131
	_go_fuzz_dep_.CoverTab[83203]++
								return buildGroup(&grp), nil
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:132
	// _ = "end of CoverTab[83203]"
}

func buildGroup(grp *_C_struct_group) *Group {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:135
	_go_fuzz_dep_.CoverTab[83209]++
								g := &Group{
		Gid:	strconv.Itoa(int(_C_gr_gid(grp))),
		Name:	_C_GoString(_C_gr_name(grp)),
	}
								return g
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:140
	// _ = "end of CoverTab[83209]"
}

type bufferKind _C_int

var (
	userBuffer	= bufferKind(_C__SC_GETPW_R_SIZE_MAX)
	groupBuffer	= bufferKind(_C__SC_GETGR_R_SIZE_MAX)
)

func (k bufferKind) initialSize() _C_size_t {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:150
	_go_fuzz_dep_.CoverTab[83210]++
								sz := _C_sysconf(_C_int(k))
								if sz == -1 {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:152
		_go_fuzz_dep_.CoverTab[83213]++

//line /usr/local/go/src/os/user/cgo_lookup_unix.go:156
		return 1024
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:156
		// _ = "end of CoverTab[83213]"
	} else {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:157
		_go_fuzz_dep_.CoverTab[83214]++
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:157
		// _ = "end of CoverTab[83214]"
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:157
	}
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:157
	// _ = "end of CoverTab[83210]"
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:157
	_go_fuzz_dep_.CoverTab[83211]++
								if !isSizeReasonable(int64(sz)) {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:158
		_go_fuzz_dep_.CoverTab[83215]++

									return maxBufferSize
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:160
		// _ = "end of CoverTab[83215]"
	} else {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:161
		_go_fuzz_dep_.CoverTab[83216]++
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:161
		// _ = "end of CoverTab[83216]"
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:161
	}
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:161
	// _ = "end of CoverTab[83211]"
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:161
	_go_fuzz_dep_.CoverTab[83212]++
								return _C_size_t(sz)
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:162
	// _ = "end of CoverTab[83212]"
}

// retryWithBuffer repeatedly calls f(), increasing the size of the
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:165
// buffer each time, until f succeeds, fails with a non-ERANGE error,
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:165
// or the buffer exceeds a reasonable limit.
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:168
func retryWithBuffer(startSize bufferKind, f func([]byte) syscall.Errno) error {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:168
	_go_fuzz_dep_.CoverTab[83217]++
								buf := make([]byte, startSize)
								for {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:170
		_go_fuzz_dep_.CoverTab[83218]++
									errno := f(buf)
									if errno == 0 {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:172
			_go_fuzz_dep_.CoverTab[83221]++
										return nil
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:173
			// _ = "end of CoverTab[83221]"
		} else {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:174
			_go_fuzz_dep_.CoverTab[83222]++
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:174
			if runtime.GOOS == "aix" && func() bool {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:174
				_go_fuzz_dep_.CoverTab[83223]++
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:174
				return errno+1 == 0
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:174
				// _ = "end of CoverTab[83223]"
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:174
			}() {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:174
				_go_fuzz_dep_.CoverTab[83224]++
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:174
				// _ = "end of CoverTab[83224]"

//line /usr/local/go/src/os/user/cgo_lookup_unix.go:177
			} else {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:177
				_go_fuzz_dep_.CoverTab[83225]++
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:177
				if errno != syscall.ERANGE {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:177
					_go_fuzz_dep_.CoverTab[83226]++
												return errno
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:178
					// _ = "end of CoverTab[83226]"
				} else {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:179
					_go_fuzz_dep_.CoverTab[83227]++
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:179
					// _ = "end of CoverTab[83227]"
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:179
				}
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:179
				// _ = "end of CoverTab[83225]"
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:179
			}
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:179
			// _ = "end of CoverTab[83222]"
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:179
		}
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:179
		// _ = "end of CoverTab[83218]"
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:179
		_go_fuzz_dep_.CoverTab[83219]++
									newSize := len(buf) * 2
									if !isSizeReasonable(int64(newSize)) {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:181
			_go_fuzz_dep_.CoverTab[83228]++
										return fmt.Errorf("internal buffer exceeds %d bytes", maxBufferSize)
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:182
			// _ = "end of CoverTab[83228]"
		} else {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:183
			_go_fuzz_dep_.CoverTab[83229]++
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:183
			// _ = "end of CoverTab[83229]"
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:183
		}
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:183
		// _ = "end of CoverTab[83219]"
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:183
		_go_fuzz_dep_.CoverTab[83220]++
									buf = make([]byte, newSize)
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:184
		// _ = "end of CoverTab[83220]"
	}
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:185
	// _ = "end of CoverTab[83217]"
}

const maxBufferSize = 1 << 20

func isSizeReasonable(sz int64) bool {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:190
	_go_fuzz_dep_.CoverTab[83230]++
								return sz > 0 && func() bool {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:191
		_go_fuzz_dep_.CoverTab[83231]++
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:191
		return sz <= maxBufferSize
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:191
		// _ = "end of CoverTab[83231]"
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:191
	}()
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:191
	// _ = "end of CoverTab[83230]"
}

// Because we can't use cgo in tests:
func structPasswdForNegativeTest() _C_struct_passwd {
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:195
	_go_fuzz_dep_.CoverTab[83232]++
								sp := _C_struct_passwd{}
								*_C_pw_uidp(&sp) = 1<<32 - 2
								*_C_pw_gidp(&sp) = 1<<32 - 3
								return sp
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:199
	// _ = "end of CoverTab[83232]"
}

//line /usr/local/go/src/os/user/cgo_lookup_unix.go:200
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/os/user/cgo_lookup_unix.go:200
var _ = _go_fuzz_dep_.CoverTab
