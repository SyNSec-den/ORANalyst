// Copyright 2018 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:5
package uuid

//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:5
)

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"strings"
	"sync"
)

// A UUID is a 128 bit (16 byte) Universal Unique IDentifier as defined in RFC
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:18
// 4122.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:20
type UUID [16]byte

// A Version represents a UUID's version.
type Version byte

// A Variant represents a UUID's variant.
type Variant byte

// Constants returned by Variant.
const (
	Invalid		= Variant(iota)	// Invalid UUID
	RFC4122				// The variant specified in RFC4122
	Reserved			// Reserved, NCS backward compatibility.
	Microsoft			// Reserved, Microsoft Corporation backward compatibility.
	Future				// Reserved for future definition.
)

const randPoolSize = 16 * 16

var (
	rander		= rand.Reader	// random function
	poolEnabled	= false
	poolMu		sync.Mutex
	poolPos		= randPoolSize		// protected with poolMu
	pool		[randPoolSize]byte	// protected with poolMu
)

type invalidLengthError struct{ len int }

func (err invalidLengthError) Error() string {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:49
	_go_fuzz_dep_.CoverTab[179441]++
										return fmt.Sprintf("invalid UUID length: %d", err.len)
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:50
	// _ = "end of CoverTab[179441]"
}

// IsInvalidLengthError is matcher function for custom error invalidLengthError
func IsInvalidLengthError(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:54
	_go_fuzz_dep_.CoverTab[179442]++
										_, ok := err.(invalidLengthError)
										return ok
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:56
	// _ = "end of CoverTab[179442]"
}

// Parse decodes s into a UUID or returns an error.  Both the standard UUID
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:59
// forms of xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx and
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:59
// urn:uuid:xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx are decoded as well as the
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:59
// Microsoft encoding {xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx} and the raw hex
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:59
// encoding: xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:64
func Parse(s string) (UUID, error) {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:64
	_go_fuzz_dep_.CoverTab[179443]++
										var uuid UUID
										switch len(s) {

	case 36:
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:68
		_go_fuzz_dep_.CoverTab[179447]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:68
		// _ = "end of CoverTab[179447]"

//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:71
	case 36 + 9:
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:71
		_go_fuzz_dep_.CoverTab[179448]++
											if strings.ToLower(s[:9]) != "urn:uuid:" {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:72
			_go_fuzz_dep_.CoverTab[179454]++
												return uuid, fmt.Errorf("invalid urn prefix: %q", s[:9])
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:73
			// _ = "end of CoverTab[179454]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:74
			_go_fuzz_dep_.CoverTab[179455]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:74
			// _ = "end of CoverTab[179455]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:74
		}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:74
		// _ = "end of CoverTab[179448]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:74
		_go_fuzz_dep_.CoverTab[179449]++
											s = s[9:]
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:75
		// _ = "end of CoverTab[179449]"

//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:78
	case 36 + 2:
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:78
		_go_fuzz_dep_.CoverTab[179450]++
											s = s[1:]
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:79
		// _ = "end of CoverTab[179450]"

//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:82
	case 32:
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:82
		_go_fuzz_dep_.CoverTab[179451]++
											var ok bool
											for i := range uuid {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:84
			_go_fuzz_dep_.CoverTab[179456]++
												uuid[i], ok = xtob(s[i*2], s[i*2+1])
												if !ok {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:86
				_go_fuzz_dep_.CoverTab[179457]++
													return uuid, errors.New("invalid UUID format")
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:87
				// _ = "end of CoverTab[179457]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:88
				_go_fuzz_dep_.CoverTab[179458]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:88
				// _ = "end of CoverTab[179458]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:88
			}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:88
			// _ = "end of CoverTab[179456]"
		}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:89
		// _ = "end of CoverTab[179451]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:89
		_go_fuzz_dep_.CoverTab[179452]++
											return uuid, nil
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:90
		// _ = "end of CoverTab[179452]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:91
		_go_fuzz_dep_.CoverTab[179453]++
											return uuid, invalidLengthError{len(s)}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:92
		// _ = "end of CoverTab[179453]"
	}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:93
	// _ = "end of CoverTab[179443]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:93
	_go_fuzz_dep_.CoverTab[179444]++

//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:96
	if s[8] != '-' || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:96
		_go_fuzz_dep_.CoverTab[179459]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:96
		return s[13] != '-'
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:96
		// _ = "end of CoverTab[179459]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:96
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:96
		_go_fuzz_dep_.CoverTab[179460]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:96
		return s[18] != '-'
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:96
		// _ = "end of CoverTab[179460]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:96
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:96
		_go_fuzz_dep_.CoverTab[179461]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:96
		return s[23] != '-'
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:96
		// _ = "end of CoverTab[179461]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:96
	}() {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:96
		_go_fuzz_dep_.CoverTab[179462]++
											return uuid, errors.New("invalid UUID format")
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:97
		// _ = "end of CoverTab[179462]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:98
		_go_fuzz_dep_.CoverTab[179463]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:98
		// _ = "end of CoverTab[179463]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:98
	}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:98
	// _ = "end of CoverTab[179444]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:98
	_go_fuzz_dep_.CoverTab[179445]++
										for i, x := range [16]int{
		0, 2, 4, 6,
		9, 11,
		14, 16,
		19, 21,
		24, 26, 28, 30, 32, 34} {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:104
		_go_fuzz_dep_.CoverTab[179464]++
											v, ok := xtob(s[x], s[x+1])
											if !ok {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:106
			_go_fuzz_dep_.CoverTab[179466]++
												return uuid, errors.New("invalid UUID format")
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:107
			// _ = "end of CoverTab[179466]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:108
			_go_fuzz_dep_.CoverTab[179467]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:108
			// _ = "end of CoverTab[179467]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:108
		}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:108
		// _ = "end of CoverTab[179464]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:108
		_go_fuzz_dep_.CoverTab[179465]++
											uuid[i] = v
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:109
		// _ = "end of CoverTab[179465]"
	}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:110
	// _ = "end of CoverTab[179445]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:110
	_go_fuzz_dep_.CoverTab[179446]++
										return uuid, nil
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:111
	// _ = "end of CoverTab[179446]"
}

// ParseBytes is like Parse, except it parses a byte slice instead of a string.
func ParseBytes(b []byte) (UUID, error) {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:115
	_go_fuzz_dep_.CoverTab[179468]++
										var uuid UUID
										switch len(b) {
	case 36:
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:118
		_go_fuzz_dep_.CoverTab[179472]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:118
		// _ = "end of CoverTab[179472]"
	case 36 + 9:
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:119
		_go_fuzz_dep_.CoverTab[179473]++
											if !bytes.Equal(bytes.ToLower(b[:9]), []byte("urn:uuid:")) {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:120
			_go_fuzz_dep_.CoverTab[179479]++
												return uuid, fmt.Errorf("invalid urn prefix: %q", b[:9])
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:121
			// _ = "end of CoverTab[179479]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:122
			_go_fuzz_dep_.CoverTab[179480]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:122
			// _ = "end of CoverTab[179480]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:122
		}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:122
		// _ = "end of CoverTab[179473]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:122
		_go_fuzz_dep_.CoverTab[179474]++
											b = b[9:]
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:123
		// _ = "end of CoverTab[179474]"
	case 36 + 2:
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:124
		_go_fuzz_dep_.CoverTab[179475]++
											b = b[1:]
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:125
		// _ = "end of CoverTab[179475]"
	case 32:
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:126
		_go_fuzz_dep_.CoverTab[179476]++
											var ok bool
											for i := 0; i < 32; i += 2 {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:128
			_go_fuzz_dep_.CoverTab[179481]++
												uuid[i/2], ok = xtob(b[i], b[i+1])
												if !ok {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:130
				_go_fuzz_dep_.CoverTab[179482]++
													return uuid, errors.New("invalid UUID format")
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:131
				// _ = "end of CoverTab[179482]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:132
				_go_fuzz_dep_.CoverTab[179483]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:132
				// _ = "end of CoverTab[179483]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:132
			}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:132
			// _ = "end of CoverTab[179481]"
		}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:133
		// _ = "end of CoverTab[179476]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:133
		_go_fuzz_dep_.CoverTab[179477]++
											return uuid, nil
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:134
		// _ = "end of CoverTab[179477]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:135
		_go_fuzz_dep_.CoverTab[179478]++
											return uuid, invalidLengthError{len(b)}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:136
		// _ = "end of CoverTab[179478]"
	}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:137
	// _ = "end of CoverTab[179468]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:137
	_go_fuzz_dep_.CoverTab[179469]++

//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:140
	if b[8] != '-' || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:140
		_go_fuzz_dep_.CoverTab[179484]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:140
		return b[13] != '-'
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:140
		// _ = "end of CoverTab[179484]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:140
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:140
		_go_fuzz_dep_.CoverTab[179485]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:140
		return b[18] != '-'
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:140
		// _ = "end of CoverTab[179485]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:140
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:140
		_go_fuzz_dep_.CoverTab[179486]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:140
		return b[23] != '-'
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:140
		// _ = "end of CoverTab[179486]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:140
	}() {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:140
		_go_fuzz_dep_.CoverTab[179487]++
											return uuid, errors.New("invalid UUID format")
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:141
		// _ = "end of CoverTab[179487]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:142
		_go_fuzz_dep_.CoverTab[179488]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:142
		// _ = "end of CoverTab[179488]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:142
	}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:142
	// _ = "end of CoverTab[179469]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:142
	_go_fuzz_dep_.CoverTab[179470]++
										for i, x := range [16]int{
		0, 2, 4, 6,
		9, 11,
		14, 16,
		19, 21,
		24, 26, 28, 30, 32, 34} {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:148
		_go_fuzz_dep_.CoverTab[179489]++
											v, ok := xtob(b[x], b[x+1])
											if !ok {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:150
			_go_fuzz_dep_.CoverTab[179491]++
												return uuid, errors.New("invalid UUID format")
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:151
			// _ = "end of CoverTab[179491]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:152
			_go_fuzz_dep_.CoverTab[179492]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:152
			// _ = "end of CoverTab[179492]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:152
		}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:152
		// _ = "end of CoverTab[179489]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:152
		_go_fuzz_dep_.CoverTab[179490]++
											uuid[i] = v
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:153
		// _ = "end of CoverTab[179490]"
	}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:154
	// _ = "end of CoverTab[179470]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:154
	_go_fuzz_dep_.CoverTab[179471]++
										return uuid, nil
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:155
	// _ = "end of CoverTab[179471]"
}

// MustParse is like Parse but panics if the string cannot be parsed.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:158
// It simplifies safe initialization of global variables holding compiled UUIDs.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:160
func MustParse(s string) UUID {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:160
	_go_fuzz_dep_.CoverTab[179493]++
										uuid, err := Parse(s)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:162
		_go_fuzz_dep_.CoverTab[179495]++
											panic(`uuid: Parse(` + s + `): ` + err.Error())
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:163
		// _ = "end of CoverTab[179495]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:164
		_go_fuzz_dep_.CoverTab[179496]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:164
		// _ = "end of CoverTab[179496]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:164
	}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:164
	// _ = "end of CoverTab[179493]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:164
	_go_fuzz_dep_.CoverTab[179494]++
										return uuid
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:165
	// _ = "end of CoverTab[179494]"
}

// FromBytes creates a new UUID from a byte slice. Returns an error if the slice
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:168
// does not have a length of 16. The bytes are copied from the slice.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:170
func FromBytes(b []byte) (uuid UUID, err error) {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:170
	_go_fuzz_dep_.CoverTab[179497]++
										err = uuid.UnmarshalBinary(b)
										return uuid, err
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:172
	// _ = "end of CoverTab[179497]"
}

// Must returns uuid if err is nil and panics otherwise.
func Must(uuid UUID, err error) UUID {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:176
	_go_fuzz_dep_.CoverTab[179498]++
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:177
		_go_fuzz_dep_.CoverTab[179500]++
											panic(err)
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:178
		// _ = "end of CoverTab[179500]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:179
		_go_fuzz_dep_.CoverTab[179501]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:179
		// _ = "end of CoverTab[179501]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:179
	}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:179
	// _ = "end of CoverTab[179498]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:179
	_go_fuzz_dep_.CoverTab[179499]++
										return uuid
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:180
	// _ = "end of CoverTab[179499]"
}

// String returns the string form of uuid, xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:183
// , or "" if uuid is invalid.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:185
func (uuid UUID) String() string {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:185
	_go_fuzz_dep_.CoverTab[179502]++
										var buf [36]byte
										encodeHex(buf[:], uuid)
										return string(buf[:])
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:188
	// _ = "end of CoverTab[179502]"
}

// URN returns the RFC 2141 URN form of uuid,
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:191
// urn:uuid:xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx,  or "" if uuid is invalid.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:193
func (uuid UUID) URN() string {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:193
	_go_fuzz_dep_.CoverTab[179503]++
										var buf [36 + 9]byte
										copy(buf[:], "urn:uuid:")
										encodeHex(buf[9:], uuid)
										return string(buf[:])
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:197
	// _ = "end of CoverTab[179503]"
}

func encodeHex(dst []byte, uuid UUID) {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:200
	_go_fuzz_dep_.CoverTab[179504]++
										hex.Encode(dst, uuid[:4])
										dst[8] = '-'
										hex.Encode(dst[9:13], uuid[4:6])
										dst[13] = '-'
										hex.Encode(dst[14:18], uuid[6:8])
										dst[18] = '-'
										hex.Encode(dst[19:23], uuid[8:10])
										dst[23] = '-'
										hex.Encode(dst[24:], uuid[10:])
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:209
	// _ = "end of CoverTab[179504]"
}

// Variant returns the variant encoded in uuid.
func (uuid UUID) Variant() Variant {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:213
	_go_fuzz_dep_.CoverTab[179505]++
										switch {
	case (uuid[8] & 0xc0) == 0x80:
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:215
		_go_fuzz_dep_.CoverTab[179506]++
											return RFC4122
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:216
		// _ = "end of CoverTab[179506]"
	case (uuid[8] & 0xe0) == 0xc0:
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:217
		_go_fuzz_dep_.CoverTab[179507]++
											return Microsoft
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:218
		// _ = "end of CoverTab[179507]"
	case (uuid[8] & 0xe0) == 0xe0:
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:219
		_go_fuzz_dep_.CoverTab[179508]++
											return Future
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:220
		// _ = "end of CoverTab[179508]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:221
		_go_fuzz_dep_.CoverTab[179509]++
											return Reserved
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:222
		// _ = "end of CoverTab[179509]"
	}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:223
	// _ = "end of CoverTab[179505]"
}

// Version returns the version of uuid.
func (uuid UUID) Version() Version {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:227
	_go_fuzz_dep_.CoverTab[179510]++
										return Version(uuid[6] >> 4)
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:228
	// _ = "end of CoverTab[179510]"
}

func (v Version) String() string {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:231
	_go_fuzz_dep_.CoverTab[179511]++
										if v > 15 {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:232
		_go_fuzz_dep_.CoverTab[179513]++
											return fmt.Sprintf("BAD_VERSION_%d", v)
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:233
		// _ = "end of CoverTab[179513]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:234
		_go_fuzz_dep_.CoverTab[179514]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:234
		// _ = "end of CoverTab[179514]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:234
	}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:234
	// _ = "end of CoverTab[179511]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:234
	_go_fuzz_dep_.CoverTab[179512]++
										return fmt.Sprintf("VERSION_%d", v)
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:235
	// _ = "end of CoverTab[179512]"
}

func (v Variant) String() string {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:238
	_go_fuzz_dep_.CoverTab[179515]++
										switch v {
	case RFC4122:
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:240
		_go_fuzz_dep_.CoverTab[179517]++
											return "RFC4122"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:241
		// _ = "end of CoverTab[179517]"
	case Reserved:
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:242
		_go_fuzz_dep_.CoverTab[179518]++
											return "Reserved"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:243
		// _ = "end of CoverTab[179518]"
	case Microsoft:
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:244
		_go_fuzz_dep_.CoverTab[179519]++
											return "Microsoft"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:245
		// _ = "end of CoverTab[179519]"
	case Future:
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:246
		_go_fuzz_dep_.CoverTab[179520]++
											return "Future"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:247
		// _ = "end of CoverTab[179520]"
	case Invalid:
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:248
		_go_fuzz_dep_.CoverTab[179521]++
											return "Invalid"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:249
		// _ = "end of CoverTab[179521]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:249
	default:
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:249
		_go_fuzz_dep_.CoverTab[179522]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:249
		// _ = "end of CoverTab[179522]"
	}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:250
	// _ = "end of CoverTab[179515]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:250
	_go_fuzz_dep_.CoverTab[179516]++
										return fmt.Sprintf("BadVariant%d", int(v))
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:251
	// _ = "end of CoverTab[179516]"
}

// SetRand sets the random number generator to r, which implements io.Reader.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:254
// If r.Read returns an error when the package requests random data then
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:254
// a panic will be issued.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:254
//
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:254
// Calling SetRand with nil sets the random number generator to the default
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:254
// generator.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:260
func SetRand(r io.Reader) {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:260
	_go_fuzz_dep_.CoverTab[179523]++
										if r == nil {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:261
		_go_fuzz_dep_.CoverTab[179525]++
											rander = rand.Reader
											return
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:263
		// _ = "end of CoverTab[179525]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:264
		_go_fuzz_dep_.CoverTab[179526]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:264
		// _ = "end of CoverTab[179526]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:264
	}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:264
	// _ = "end of CoverTab[179523]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:264
	_go_fuzz_dep_.CoverTab[179524]++
										rander = r
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:265
	// _ = "end of CoverTab[179524]"
}

// EnableRandPool enables internal randomness pool used for Random
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:268
// (Version 4) UUID generation. The pool contains random bytes read from
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:268
// the random number generator on demand in batches. Enabling the pool
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:268
// may improve the UUID generation throughput significantly.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:268
//
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:268
// Since the pool is stored on the Go heap, this feature may be a bad fit
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:268
// for security sensitive applications.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:268
//
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:268
// Both EnableRandPool and DisableRandPool are not thread-safe and should
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:268
// only be called when there is no possibility that New or any other
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:268
// UUID Version 4 generation function will be called concurrently.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:279
func EnableRandPool() {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:279
	_go_fuzz_dep_.CoverTab[179527]++
										poolEnabled = true
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:280
	// _ = "end of CoverTab[179527]"
}

// DisableRandPool disables the randomness pool if it was previously
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:283
// enabled with EnableRandPool.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:283
//
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:283
// Both EnableRandPool and DisableRandPool are not thread-safe and should
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:283
// only be called when there is no possibility that New or any other
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:283
// UUID Version 4 generation function will be called concurrently.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:289
func DisableRandPool() {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:289
	_go_fuzz_dep_.CoverTab[179528]++
										poolEnabled = false
										defer poolMu.Unlock()
										poolMu.Lock()
										poolPos = randPoolSize
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:293
	// _ = "end of CoverTab[179528]"
}

//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:294
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/uuid.go:294
var _ = _go_fuzz_dep_.CoverTab
