// Copyright 2016 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/marshal.go:5
package uuid

//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/marshal.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/marshal.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/marshal.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/marshal.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/marshal.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/marshal.go:5
)

import "fmt"

// MarshalText implements encoding.TextMarshaler.
func (uuid UUID) MarshalText() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/marshal.go:10
	_go_fuzz_dep_.CoverTab[179318]++
										var js [36]byte
										encodeHex(js[:], uuid)
										return js[:], nil
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/marshal.go:13
	// _ = "end of CoverTab[179318]"
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (uuid *UUID) UnmarshalText(data []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/marshal.go:17
	_go_fuzz_dep_.CoverTab[179319]++
										id, err := ParseBytes(data)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/marshal.go:19
		_go_fuzz_dep_.CoverTab[179321]++
											return err
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/marshal.go:20
		// _ = "end of CoverTab[179321]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/marshal.go:21
		_go_fuzz_dep_.CoverTab[179322]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/marshal.go:21
		// _ = "end of CoverTab[179322]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/marshal.go:21
	}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/marshal.go:21
	// _ = "end of CoverTab[179319]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/marshal.go:21
	_go_fuzz_dep_.CoverTab[179320]++
										*uuid = id
										return nil
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/marshal.go:23
	// _ = "end of CoverTab[179320]"
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (uuid UUID) MarshalBinary() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/marshal.go:27
	_go_fuzz_dep_.CoverTab[179323]++
										return uuid[:], nil
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/marshal.go:28
	// _ = "end of CoverTab[179323]"
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (uuid *UUID) UnmarshalBinary(data []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/marshal.go:32
	_go_fuzz_dep_.CoverTab[179324]++
										if len(data) != 16 {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/marshal.go:33
		_go_fuzz_dep_.CoverTab[179326]++
											return fmt.Errorf("invalid UUID (got %d bytes)", len(data))
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/marshal.go:34
		// _ = "end of CoverTab[179326]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/marshal.go:35
		_go_fuzz_dep_.CoverTab[179327]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/marshal.go:35
		// _ = "end of CoverTab[179327]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/marshal.go:35
	}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/marshal.go:35
	// _ = "end of CoverTab[179324]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/marshal.go:35
	_go_fuzz_dep_.CoverTab[179325]++
										copy(uuid[:], data)
										return nil
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/marshal.go:37
	// _ = "end of CoverTab[179325]"
}

//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/marshal.go:38
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/marshal.go:38
var _ = _go_fuzz_dep_.CoverTab
