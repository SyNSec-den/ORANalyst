// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build gofuzz

//line /usr/local/go/src/encoding/json/fuzz.go:7
package json

//line /usr/local/go/src/encoding/json/fuzz.go:7
import (
//line /usr/local/go/src/encoding/json/fuzz.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/encoding/json/fuzz.go:7
)
//line /usr/local/go/src/encoding/json/fuzz.go:7
import (
//line /usr/local/go/src/encoding/json/fuzz.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/encoding/json/fuzz.go:7
)

import (
	"fmt"
)

func Fuzz(data []byte) (score int) {
//line /usr/local/go/src/encoding/json/fuzz.go:13
	_go_fuzz_dep_.CoverTab[28035]++
							for _, ctor := range []func() any{
		func() any { _go_fuzz_dep_.CoverTab[28037]++; return new(any); // _ = "end of CoverTab[28037]" },
		func() any {
//line /usr/local/go/src/encoding/json/fuzz.go:16
			_go_fuzz_dep_.CoverTab[28038]++
//line /usr/local/go/src/encoding/json/fuzz.go:16
			return new(map[string]any)
//line /usr/local/go/src/encoding/json/fuzz.go:16
			// _ = "end of CoverTab[28038]"
//line /usr/local/go/src/encoding/json/fuzz.go:16
		},
		func() any { _go_fuzz_dep_.CoverTab[28039]++; return new([]any); // _ = "end of CoverTab[28039]" },
	} {
//line /usr/local/go/src/encoding/json/fuzz.go:18
		_go_fuzz_dep_.CoverTab[28040]++
								v := ctor()
								err := Unmarshal(data, v)
								if err != nil {
//line /usr/local/go/src/encoding/json/fuzz.go:21
			_go_fuzz_dep_.CoverTab[28043]++
									continue
//line /usr/local/go/src/encoding/json/fuzz.go:22
			// _ = "end of CoverTab[28043]"
		} else {
//line /usr/local/go/src/encoding/json/fuzz.go:23
			_go_fuzz_dep_.CoverTab[28044]++
//line /usr/local/go/src/encoding/json/fuzz.go:23
			// _ = "end of CoverTab[28044]"
//line /usr/local/go/src/encoding/json/fuzz.go:23
		}
//line /usr/local/go/src/encoding/json/fuzz.go:23
		// _ = "end of CoverTab[28040]"
//line /usr/local/go/src/encoding/json/fuzz.go:23
		_go_fuzz_dep_.CoverTab[28041]++
								score = 1

								m, err := Marshal(v)
								if err != nil {
//line /usr/local/go/src/encoding/json/fuzz.go:27
			_go_fuzz_dep_.CoverTab[28045]++
									fmt.Printf("v=%#v\n", v)
									panic(err)
//line /usr/local/go/src/encoding/json/fuzz.go:29
			// _ = "end of CoverTab[28045]"
		} else {
//line /usr/local/go/src/encoding/json/fuzz.go:30
			_go_fuzz_dep_.CoverTab[28046]++
//line /usr/local/go/src/encoding/json/fuzz.go:30
			// _ = "end of CoverTab[28046]"
//line /usr/local/go/src/encoding/json/fuzz.go:30
		}
//line /usr/local/go/src/encoding/json/fuzz.go:30
		// _ = "end of CoverTab[28041]"
//line /usr/local/go/src/encoding/json/fuzz.go:30
		_go_fuzz_dep_.CoverTab[28042]++

								u := ctor()
								err = Unmarshal(m, u)
								if err != nil {
//line /usr/local/go/src/encoding/json/fuzz.go:34
			_go_fuzz_dep_.CoverTab[28047]++
									fmt.Printf("v=%#v\n", v)
									fmt.Printf("m=%s\n", m)
									panic(err)
//line /usr/local/go/src/encoding/json/fuzz.go:37
			// _ = "end of CoverTab[28047]"
		} else {
//line /usr/local/go/src/encoding/json/fuzz.go:38
			_go_fuzz_dep_.CoverTab[28048]++
//line /usr/local/go/src/encoding/json/fuzz.go:38
			// _ = "end of CoverTab[28048]"
//line /usr/local/go/src/encoding/json/fuzz.go:38
		}
//line /usr/local/go/src/encoding/json/fuzz.go:38
		// _ = "end of CoverTab[28042]"
	}
//line /usr/local/go/src/encoding/json/fuzz.go:39
	// _ = "end of CoverTab[28035]"
//line /usr/local/go/src/encoding/json/fuzz.go:39
	_go_fuzz_dep_.CoverTab[28036]++

							return
//line /usr/local/go/src/encoding/json/fuzz.go:41
	// _ = "end of CoverTab[28036]"
}

//line /usr/local/go/src/encoding/json/fuzz.go:42
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/encoding/json/fuzz.go:42
var _ = _go_fuzz_dep_.CoverTab
