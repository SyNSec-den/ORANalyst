// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build gofuzz

//line /usr/local/go/src/html/fuzz.go:7
package html

//line /usr/local/go/src/html/fuzz.go:7
import (
//line /usr/local/go/src/html/fuzz.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/html/fuzz.go:7
)
//line /usr/local/go/src/html/fuzz.go:7
import (
//line /usr/local/go/src/html/fuzz.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/html/fuzz.go:7
)

import (
	"fmt"
)

func Fuzz(data []byte) int {
//line /usr/local/go/src/html/fuzz.go:13
	_go_fuzz_dep_.CoverTab[28910]++
						v := string(data)

						e := EscapeString(v)
						u := UnescapeString(e)
						if v != u {
//line /usr/local/go/src/html/fuzz.go:18
		_go_fuzz_dep_.CoverTab[28912]++
							fmt.Printf("v = %q\n", v)
							fmt.Printf("e = %q\n", e)
							fmt.Printf("u = %q\n", u)
							panic("not equal")
//line /usr/local/go/src/html/fuzz.go:22
		// _ = "end of CoverTab[28912]"
	} else {
//line /usr/local/go/src/html/fuzz.go:23
		_go_fuzz_dep_.CoverTab[28913]++
//line /usr/local/go/src/html/fuzz.go:23
		// _ = "end of CoverTab[28913]"
//line /usr/local/go/src/html/fuzz.go:23
	}
//line /usr/local/go/src/html/fuzz.go:23
	// _ = "end of CoverTab[28910]"
//line /usr/local/go/src/html/fuzz.go:23
	_go_fuzz_dep_.CoverTab[28911]++

//line /usr/local/go/src/html/fuzz.go:28
	EscapeString(UnescapeString(v))

						return 0
//line /usr/local/go/src/html/fuzz.go:30
	// _ = "end of CoverTab[28911]"
}

//line /usr/local/go/src/html/fuzz.go:31
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/html/fuzz.go:31
var _ = _go_fuzz_dep_.CoverTab
