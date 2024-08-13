// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/os/user/user.go:5
/*
Package user allows user account lookups by name or id.

For most Unix systems, this package has two internal implementations of
resolving user and group ids to names, and listing supplementary group IDs.
One is written in pure Go and parses /etc/passwd and /etc/group. The other
is cgo-based and relies on the standard C library (libc) routines such as
getpwuid_r, getgrnam_r, and getgrouplist.

When cgo is available, and the required routines are implemented in libc
for a particular platform, cgo-based (libc-backed) code is used.
This can be overridden by using osusergo build tag, which enforces
the pure Go implementation.
*/
package user

//line /usr/local/go/src/os/user/user.go:19
import (
//line /usr/local/go/src/os/user/user.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/os/user/user.go:19
)
//line /usr/local/go/src/os/user/user.go:19
import (
//line /usr/local/go/src/os/user/user.go:19
	_atomic_ "sync/atomic"
//line /usr/local/go/src/os/user/user.go:19
)

import (
	"strconv"
)

// These may be set to false in init() for a particular platform and/or
//line /usr/local/go/src/os/user/user.go:25
// build flags to let the tests know to skip tests of some features.
//line /usr/local/go/src/os/user/user.go:27
var (
	userImplemented		= true
	groupImplemented	= true
	groupListImplemented	= true
)

// User represents a user account.
type User struct {
	// Uid is the user ID.
	// On POSIX systems, this is a decimal number representing the uid.
	// On Windows, this is a security identifier (SID) in a string format.
	// On Plan 9, this is the contents of /dev/user.
	Uid	string
	// Gid is the primary group ID.
	// On POSIX systems, this is a decimal number representing the gid.
	// On Windows, this is a SID in a string format.
	// On Plan 9, this is the contents of /dev/user.
	Gid	string
	// Username is the login name.
	Username	string
	// Name is the user's real or display name.
	// It might be blank.
	// On POSIX systems, this is the first (or only) entry in the GECOS field
	// list.
	// On Windows, this is the user's display name.
	// On Plan 9, this is the contents of /dev/user.
	Name	string
	// HomeDir is the path to the user's home directory (if they have one).
	HomeDir	string
}

// Group represents a grouping of users.
//line /usr/local/go/src/os/user/user.go:58
//
//line /usr/local/go/src/os/user/user.go:58
// On POSIX systems Gid contains a decimal number representing the group ID.
//line /usr/local/go/src/os/user/user.go:61
type Group struct {
	Gid	string	// group ID
	Name	string	// group name
}

// UnknownUserIdError is returned by LookupId when a user cannot be found.
type UnknownUserIdError int

func (e UnknownUserIdError) Error() string {
//line /usr/local/go/src/os/user/user.go:69
	_go_fuzz_dep_.CoverTab[83252]++
						return "user: unknown userid " + strconv.Itoa(int(e))
//line /usr/local/go/src/os/user/user.go:70
	// _ = "end of CoverTab[83252]"
}

// UnknownUserError is returned by Lookup when
//line /usr/local/go/src/os/user/user.go:73
// a user cannot be found.
//line /usr/local/go/src/os/user/user.go:75
type UnknownUserError string

func (e UnknownUserError) Error() string {
//line /usr/local/go/src/os/user/user.go:77
	_go_fuzz_dep_.CoverTab[83253]++
						return "user: unknown user " + string(e)
//line /usr/local/go/src/os/user/user.go:78
	// _ = "end of CoverTab[83253]"
}

// UnknownGroupIdError is returned by LookupGroupId when
//line /usr/local/go/src/os/user/user.go:81
// a group cannot be found.
//line /usr/local/go/src/os/user/user.go:83
type UnknownGroupIdError string

func (e UnknownGroupIdError) Error() string {
//line /usr/local/go/src/os/user/user.go:85
	_go_fuzz_dep_.CoverTab[83254]++
						return "group: unknown groupid " + string(e)
//line /usr/local/go/src/os/user/user.go:86
	// _ = "end of CoverTab[83254]"
}

// UnknownGroupError is returned by LookupGroup when
//line /usr/local/go/src/os/user/user.go:89
// a group cannot be found.
//line /usr/local/go/src/os/user/user.go:91
type UnknownGroupError string

func (e UnknownGroupError) Error() string {
//line /usr/local/go/src/os/user/user.go:93
	_go_fuzz_dep_.CoverTab[83255]++
						return "group: unknown group " + string(e)
//line /usr/local/go/src/os/user/user.go:94
	// _ = "end of CoverTab[83255]"
}

//line /usr/local/go/src/os/user/user.go:95
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/os/user/user.go:95
var _ = _go_fuzz_dep_.CoverTab
