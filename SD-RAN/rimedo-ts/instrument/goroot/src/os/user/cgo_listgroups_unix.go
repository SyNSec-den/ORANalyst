// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build (cgo || darwin) && !osusergo && (darwin || dragonfly || freebsd || (linux && !android) || netbsd || openbsd || (solaris && !illumos))

//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:7
package user

//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:7
import (
//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:7
)
//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:7
import (
//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:7
)

import (
	"fmt"
	"strconv"
	"unsafe"
)

const maxGroups = 2048

func listGroups(u *User) ([]string, error) {
//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:17
	_go_fuzz_dep_.CoverTab[83145]++
								ug, err := strconv.Atoi(u.Gid)
								if err != nil {
//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:19
		_go_fuzz_dep_.CoverTab[83149]++
									return nil, fmt.Errorf("user: list groups for %s: invalid gid %q", u.Username, u.Gid)
//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:20
		// _ = "end of CoverTab[83149]"
	} else {
//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:21
		_go_fuzz_dep_.CoverTab[83150]++
//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:21
		// _ = "end of CoverTab[83150]"
//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:21
	}
//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:21
	// _ = "end of CoverTab[83145]"
//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:21
	_go_fuzz_dep_.CoverTab[83146]++
								userGID := _C_gid_t(ug)
								nameC := make([]byte, len(u.Username)+1)
								copy(nameC, u.Username)

								n := _C_int(256)
								gidsC := make([]_C_gid_t, n)
								rv := getGroupList((*_C_char)(unsafe.Pointer(&nameC[0])), userGID, &gidsC[0], &n)
								if rv == -1 {
//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:29
		_go_fuzz_dep_.CoverTab[83151]++

//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:32
		if err := groupRetry(u.Username, nameC, userGID, &gidsC, &n); err != nil {
//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:32
			_go_fuzz_dep_.CoverTab[83152]++
										return nil, err
//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:33
			// _ = "end of CoverTab[83152]"
		} else {
//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:34
			_go_fuzz_dep_.CoverTab[83153]++
//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:34
			// _ = "end of CoverTab[83153]"
//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:34
		}
//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:34
		// _ = "end of CoverTab[83151]"
	} else {
//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:35
		_go_fuzz_dep_.CoverTab[83154]++
//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:35
		// _ = "end of CoverTab[83154]"
//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:35
	}
//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:35
	// _ = "end of CoverTab[83146]"
//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:35
	_go_fuzz_dep_.CoverTab[83147]++
								gidsC = gidsC[:n]
								gids := make([]string, 0, n)
								for _, g := range gidsC[:n] {
//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:38
		_go_fuzz_dep_.CoverTab[83155]++
									gids = append(gids, strconv.Itoa(int(g)))
//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:39
		// _ = "end of CoverTab[83155]"
	}
//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:40
	// _ = "end of CoverTab[83147]"
//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:40
	_go_fuzz_dep_.CoverTab[83148]++
								return gids, nil
//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:41
	// _ = "end of CoverTab[83148]"
}

// groupRetry retries getGroupList with much larger size for n. The result is
//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:44
// stored in gids.
//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:46
func groupRetry(username string, name []byte, userGID _C_gid_t, gids *[]_C_gid_t, n *_C_int) error {
//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:46
	_go_fuzz_dep_.CoverTab[83156]++

								if *n > maxGroups {
//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:48
		_go_fuzz_dep_.CoverTab[83159]++
									return fmt.Errorf("user: %q is a member of more than %d groups", username, maxGroups)
//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:49
		// _ = "end of CoverTab[83159]"
	} else {
//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:50
		_go_fuzz_dep_.CoverTab[83160]++
//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:50
		// _ = "end of CoverTab[83160]"
//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:50
	}
//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:50
	// _ = "end of CoverTab[83156]"
//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:50
	_go_fuzz_dep_.CoverTab[83157]++
								*gids = make([]_C_gid_t, *n)
								rv := getGroupList((*_C_char)(unsafe.Pointer(&name[0])), userGID, &(*gids)[0], n)
								if rv == -1 {
//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:53
		_go_fuzz_dep_.CoverTab[83161]++
									return fmt.Errorf("user: list groups for %s failed", username)
//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:54
		// _ = "end of CoverTab[83161]"
	} else {
//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:55
		_go_fuzz_dep_.CoverTab[83162]++
//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:55
		// _ = "end of CoverTab[83162]"
//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:55
	}
//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:55
	// _ = "end of CoverTab[83157]"
//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:55
	_go_fuzz_dep_.CoverTab[83158]++
								return nil
//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:56
	// _ = "end of CoverTab[83158]"
}

//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:57
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/os/user/cgo_listgroups_unix.go:57
var _ = _go_fuzz_dep_.CoverTab
