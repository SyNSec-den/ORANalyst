// Copyright 2016 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:5
package uuid

//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:5
)

import (
	"database/sql/driver"
	"fmt"
)

// Scan implements sql.Scanner so UUIDs can be read from databases transparently.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:12
// Currently, database types that map to string and []byte are supported. Please
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:12
// consult database-specific driver documentation for matching types.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:15
func (uuid *UUID) Scan(src interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:15
	_go_fuzz_dep_.CoverTab[179394]++
										switch src := src.(type) {
	case nil:
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:17
		_go_fuzz_dep_.CoverTab[179396]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:18
		// _ = "end of CoverTab[179396]"

	case string:
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:20
		_go_fuzz_dep_.CoverTab[179397]++

											if src == "" {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:22
			_go_fuzz_dep_.CoverTab[179404]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:23
			// _ = "end of CoverTab[179404]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:24
			_go_fuzz_dep_.CoverTab[179405]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:24
			// _ = "end of CoverTab[179405]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:24
		}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:24
		// _ = "end of CoverTab[179397]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:24
		_go_fuzz_dep_.CoverTab[179398]++

//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:27
		u, err := Parse(src)
		if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:28
			_go_fuzz_dep_.CoverTab[179406]++
												return fmt.Errorf("Scan: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:29
			// _ = "end of CoverTab[179406]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:30
			_go_fuzz_dep_.CoverTab[179407]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:30
			// _ = "end of CoverTab[179407]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:30
		}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:30
		// _ = "end of CoverTab[179398]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:30
		_go_fuzz_dep_.CoverTab[179399]++

											*uuid = u
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:32
		// _ = "end of CoverTab[179399]"

	case []byte:
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:34
		_go_fuzz_dep_.CoverTab[179400]++

											if len(src) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:36
			_go_fuzz_dep_.CoverTab[179408]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:37
			// _ = "end of CoverTab[179408]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:38
			_go_fuzz_dep_.CoverTab[179409]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:38
			// _ = "end of CoverTab[179409]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:38
		}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:38
		// _ = "end of CoverTab[179400]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:38
		_go_fuzz_dep_.CoverTab[179401]++

//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:42
		if len(src) != 16 {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:42
			_go_fuzz_dep_.CoverTab[179410]++
												return uuid.Scan(string(src))
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:43
			// _ = "end of CoverTab[179410]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:44
			_go_fuzz_dep_.CoverTab[179411]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:44
			// _ = "end of CoverTab[179411]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:44
		}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:44
		// _ = "end of CoverTab[179401]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:44
		_go_fuzz_dep_.CoverTab[179402]++
											copy((*uuid)[:], src)
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:45
		// _ = "end of CoverTab[179402]"

	default:
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:47
		_go_fuzz_dep_.CoverTab[179403]++
											return fmt.Errorf("Scan: unable to scan type %T into UUID", src)
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:48
		// _ = "end of CoverTab[179403]"
	}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:49
	// _ = "end of CoverTab[179394]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:49
	_go_fuzz_dep_.CoverTab[179395]++

										return nil
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:51
	// _ = "end of CoverTab[179395]"
}

// Value implements sql.Valuer so that UUIDs can be written to databases
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:54
// transparently. Currently, UUIDs map to strings. Please consult
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:54
// database-specific driver documentation for matching types.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:57
func (uuid UUID) Value() (driver.Value, error) {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:57
	_go_fuzz_dep_.CoverTab[179412]++
										return uuid.String(), nil
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:58
	// _ = "end of CoverTab[179412]"
}

//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:59
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/sql.go:59
var _ = _go_fuzz_dep_.CoverTab
