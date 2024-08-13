// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/os/user/lookup.go:5
package user

//line /usr/local/go/src/os/user/lookup.go:5
import (
//line /usr/local/go/src/os/user/lookup.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/os/user/lookup.go:5
)
//line /usr/local/go/src/os/user/lookup.go:5
import (
//line /usr/local/go/src/os/user/lookup.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/os/user/lookup.go:5
)

import "sync"

const (
	userFile	= "/etc/passwd"
	groupFile	= "/etc/group"
)

var colon = []byte{':'}

// Current returns the current user.
//line /usr/local/go/src/os/user/lookup.go:16
//
//line /usr/local/go/src/os/user/lookup.go:16
// The first call will cache the current user information.
//line /usr/local/go/src/os/user/lookup.go:16
// Subsequent calls will return the cached value and will not reflect
//line /usr/local/go/src/os/user/lookup.go:16
// changes to the current user.
//line /usr/local/go/src/os/user/lookup.go:21
func Current() (*User, error) {
//line /usr/local/go/src/os/user/lookup.go:21
	_go_fuzz_dep_.CoverTab[83233]++
						cache.Do(func() {
//line /usr/local/go/src/os/user/lookup.go:22
		_go_fuzz_dep_.CoverTab[83236]++
//line /usr/local/go/src/os/user/lookup.go:22
		cache.u, cache.err = current()
//line /usr/local/go/src/os/user/lookup.go:22
		// _ = "end of CoverTab[83236]"
//line /usr/local/go/src/os/user/lookup.go:22
	})
//line /usr/local/go/src/os/user/lookup.go:22
	// _ = "end of CoverTab[83233]"
//line /usr/local/go/src/os/user/lookup.go:22
	_go_fuzz_dep_.CoverTab[83234]++
						if cache.err != nil {
//line /usr/local/go/src/os/user/lookup.go:23
		_go_fuzz_dep_.CoverTab[83237]++
							return nil, cache.err
//line /usr/local/go/src/os/user/lookup.go:24
		// _ = "end of CoverTab[83237]"
	} else {
//line /usr/local/go/src/os/user/lookup.go:25
		_go_fuzz_dep_.CoverTab[83238]++
//line /usr/local/go/src/os/user/lookup.go:25
		// _ = "end of CoverTab[83238]"
//line /usr/local/go/src/os/user/lookup.go:25
	}
//line /usr/local/go/src/os/user/lookup.go:25
	// _ = "end of CoverTab[83234]"
//line /usr/local/go/src/os/user/lookup.go:25
	_go_fuzz_dep_.CoverTab[83235]++
						u := *cache.u
						return &u, nil
//line /usr/local/go/src/os/user/lookup.go:27
	// _ = "end of CoverTab[83235]"
}

// cache of the current user
var cache struct {
	sync.Once
	u	*User
	err	error
}

// Lookup looks up a user by username. If the user cannot be found, the
//line /usr/local/go/src/os/user/lookup.go:37
// returned error is of type UnknownUserError.
//line /usr/local/go/src/os/user/lookup.go:39
func Lookup(username string) (*User, error) {
//line /usr/local/go/src/os/user/lookup.go:39
	_go_fuzz_dep_.CoverTab[83239]++
						if u, err := Current(); err == nil && func() bool {
//line /usr/local/go/src/os/user/lookup.go:40
		_go_fuzz_dep_.CoverTab[83241]++
//line /usr/local/go/src/os/user/lookup.go:40
		return u.Username == username
//line /usr/local/go/src/os/user/lookup.go:40
		// _ = "end of CoverTab[83241]"
//line /usr/local/go/src/os/user/lookup.go:40
	}() {
//line /usr/local/go/src/os/user/lookup.go:40
		_go_fuzz_dep_.CoverTab[83242]++
							return u, err
//line /usr/local/go/src/os/user/lookup.go:41
		// _ = "end of CoverTab[83242]"
	} else {
//line /usr/local/go/src/os/user/lookup.go:42
		_go_fuzz_dep_.CoverTab[83243]++
//line /usr/local/go/src/os/user/lookup.go:42
		// _ = "end of CoverTab[83243]"
//line /usr/local/go/src/os/user/lookup.go:42
	}
//line /usr/local/go/src/os/user/lookup.go:42
	// _ = "end of CoverTab[83239]"
//line /usr/local/go/src/os/user/lookup.go:42
	_go_fuzz_dep_.CoverTab[83240]++
						return lookupUser(username)
//line /usr/local/go/src/os/user/lookup.go:43
	// _ = "end of CoverTab[83240]"
}

// LookupId looks up a user by userid. If the user cannot be found, the
//line /usr/local/go/src/os/user/lookup.go:46
// returned error is of type UnknownUserIdError.
//line /usr/local/go/src/os/user/lookup.go:48
func LookupId(uid string) (*User, error) {
//line /usr/local/go/src/os/user/lookup.go:48
	_go_fuzz_dep_.CoverTab[83244]++
						if u, err := Current(); err == nil && func() bool {
//line /usr/local/go/src/os/user/lookup.go:49
		_go_fuzz_dep_.CoverTab[83246]++
//line /usr/local/go/src/os/user/lookup.go:49
		return u.Uid == uid
//line /usr/local/go/src/os/user/lookup.go:49
		// _ = "end of CoverTab[83246]"
//line /usr/local/go/src/os/user/lookup.go:49
	}() {
//line /usr/local/go/src/os/user/lookup.go:49
		_go_fuzz_dep_.CoverTab[83247]++
							return u, err
//line /usr/local/go/src/os/user/lookup.go:50
		// _ = "end of CoverTab[83247]"
	} else {
//line /usr/local/go/src/os/user/lookup.go:51
		_go_fuzz_dep_.CoverTab[83248]++
//line /usr/local/go/src/os/user/lookup.go:51
		// _ = "end of CoverTab[83248]"
//line /usr/local/go/src/os/user/lookup.go:51
	}
//line /usr/local/go/src/os/user/lookup.go:51
	// _ = "end of CoverTab[83244]"
//line /usr/local/go/src/os/user/lookup.go:51
	_go_fuzz_dep_.CoverTab[83245]++
						return lookupUserId(uid)
//line /usr/local/go/src/os/user/lookup.go:52
	// _ = "end of CoverTab[83245]"
}

// LookupGroup looks up a group by name. If the group cannot be found, the
//line /usr/local/go/src/os/user/lookup.go:55
// returned error is of type UnknownGroupError.
//line /usr/local/go/src/os/user/lookup.go:57
func LookupGroup(name string) (*Group, error) {
//line /usr/local/go/src/os/user/lookup.go:57
	_go_fuzz_dep_.CoverTab[83249]++
						return lookupGroup(name)
//line /usr/local/go/src/os/user/lookup.go:58
	// _ = "end of CoverTab[83249]"
}

// LookupGroupId looks up a group by groupid. If the group cannot be found, the
//line /usr/local/go/src/os/user/lookup.go:61
// returned error is of type UnknownGroupIdError.
//line /usr/local/go/src/os/user/lookup.go:63
func LookupGroupId(gid string) (*Group, error) {
//line /usr/local/go/src/os/user/lookup.go:63
	_go_fuzz_dep_.CoverTab[83250]++
						return lookupGroupId(gid)
//line /usr/local/go/src/os/user/lookup.go:64
	// _ = "end of CoverTab[83250]"
}

// GroupIds returns the list of group IDs that the user is a member of.
func (u *User) GroupIds() ([]string, error) {
//line /usr/local/go/src/os/user/lookup.go:68
	_go_fuzz_dep_.CoverTab[83251]++
						return listGroups(u)
//line /usr/local/go/src/os/user/lookup.go:69
	// _ = "end of CoverTab[83251]"
}

//line /usr/local/go/src/os/user/lookup.go:70
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/os/user/lookup.go:70
var _ = _go_fuzz_dep_.CoverTab
