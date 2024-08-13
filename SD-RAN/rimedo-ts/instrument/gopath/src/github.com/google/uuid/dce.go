// Copyright 2016 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:5
package uuid

//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:5
)

import (
	"encoding/binary"
	"fmt"
	"os"
)

// A Domain represents a Version 2 domain
type Domain byte

// Domain constants for DCE Security (Version 2) UUIDs.
const (
	Person	= Domain(0)
	Group	= Domain(1)
	Org	= Domain(2)
)

// NewDCESecurity returns a DCE Security (Version 2) UUID.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:23
//
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:23
// The domain should be one of Person, Group or Org.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:23
// On a POSIX system the id should be the users UID for the Person
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:23
// domain and the users GID for the Group.  The meaning of id for
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:23
// the domain Org or on non-POSIX systems is site defined.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:23
//
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:23
// For a given domain/id pair the same token may be returned for up to
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:23
// 7 minutes and 10 seconds.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:32
func NewDCESecurity(domain Domain, id uint32) (UUID, error) {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:32
	_go_fuzz_dep_.CoverTab[179301]++
										uuid, err := NewUUID()
										if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:34
		_go_fuzz_dep_.CoverTab[179303]++
											uuid[6] = (uuid[6] & 0x0f) | 0x20
											uuid[9] = byte(domain)
											binary.BigEndian.PutUint32(uuid[0:], id)
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:37
		// _ = "end of CoverTab[179303]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:38
		_go_fuzz_dep_.CoverTab[179304]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:38
		// _ = "end of CoverTab[179304]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:38
	}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:38
	// _ = "end of CoverTab[179301]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:38
	_go_fuzz_dep_.CoverTab[179302]++
										return uuid, err
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:39
	// _ = "end of CoverTab[179302]"
}

// NewDCEPerson returns a DCE Security (Version 2) UUID in the person
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:42
// domain with the id returned by os.Getuid.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:42
//
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:42
//	NewDCESecurity(Person, uint32(os.Getuid()))
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:46
func NewDCEPerson() (UUID, error) {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:46
	_go_fuzz_dep_.CoverTab[179305]++
										return NewDCESecurity(Person, uint32(os.Getuid()))
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:47
	// _ = "end of CoverTab[179305]"
}

// NewDCEGroup returns a DCE Security (Version 2) UUID in the group
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:50
// domain with the id returned by os.Getgid.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:50
//
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:50
//	NewDCESecurity(Group, uint32(os.Getgid()))
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:54
func NewDCEGroup() (UUID, error) {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:54
	_go_fuzz_dep_.CoverTab[179306]++
										return NewDCESecurity(Group, uint32(os.Getgid()))
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:55
	// _ = "end of CoverTab[179306]"
}

// Domain returns the domain for a Version 2 UUID.  Domains are only defined
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:58
// for Version 2 UUIDs.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:60
func (uuid UUID) Domain() Domain {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:60
	_go_fuzz_dep_.CoverTab[179307]++
										return Domain(uuid[9])
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:61
	// _ = "end of CoverTab[179307]"
}

// ID returns the id for a Version 2 UUID. IDs are only defined for Version 2
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:64
// UUIDs.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:66
func (uuid UUID) ID() uint32 {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:66
	_go_fuzz_dep_.CoverTab[179308]++
										return binary.BigEndian.Uint32(uuid[0:4])
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:67
	// _ = "end of CoverTab[179308]"
}

func (d Domain) String() string {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:70
	_go_fuzz_dep_.CoverTab[179309]++
										switch d {
	case Person:
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:72
		_go_fuzz_dep_.CoverTab[179311]++
											return "Person"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:73
		// _ = "end of CoverTab[179311]"
	case Group:
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:74
		_go_fuzz_dep_.CoverTab[179312]++
											return "Group"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:75
		// _ = "end of CoverTab[179312]"
	case Org:
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:76
		_go_fuzz_dep_.CoverTab[179313]++
											return "Org"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:77
		// _ = "end of CoverTab[179313]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:77
	default:
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:77
		_go_fuzz_dep_.CoverTab[179314]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:77
		// _ = "end of CoverTab[179314]"
	}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:78
	// _ = "end of CoverTab[179309]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:78
	_go_fuzz_dep_.CoverTab[179310]++
										return fmt.Sprintf("Domain%d", int(d))
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:79
	// _ = "end of CoverTab[179310]"
}

//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:80
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/dce.go:80
var _ = _go_fuzz_dep_.CoverTab
