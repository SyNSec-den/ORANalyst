// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/encoding/gob/error.go:5
package gob

//line /usr/local/go/src/encoding/gob/error.go:5
import (
//line /usr/local/go/src/encoding/gob/error.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/encoding/gob/error.go:5
)
//line /usr/local/go/src/encoding/gob/error.go:5
import (
//line /usr/local/go/src/encoding/gob/error.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/encoding/gob/error.go:5
)

import "fmt"

//line /usr/local/go/src/encoding/gob/error.go:16
// A gobError is used to distinguish errors (panics) generated in this package.
type gobError struct {
	err error
}

// errorf is like error_ but takes Printf-style arguments to construct an error.
//line /usr/local/go/src/encoding/gob/error.go:21
// It always prefixes the message with "gob: ".
//line /usr/local/go/src/encoding/gob/error.go:23
func errorf(format string, args ...any) {
//line /usr/local/go/src/encoding/gob/error.go:23
	_go_fuzz_dep_.CoverTab[85178]++
							error_(fmt.Errorf("gob: "+format, args...))
//line /usr/local/go/src/encoding/gob/error.go:24
	// _ = "end of CoverTab[85178]"
}

// error_ wraps the argument error and uses it as the argument to panic.
func error_(err error) {
//line /usr/local/go/src/encoding/gob/error.go:28
	_go_fuzz_dep_.CoverTab[85179]++
							panic(gobError{err})
//line /usr/local/go/src/encoding/gob/error.go:29
	// _ = "end of CoverTab[85179]"
}

// catchError is meant to be used as a deferred function to turn a panic(gobError) into a
//line /usr/local/go/src/encoding/gob/error.go:32
// plain error. It overwrites the error return of the function that deferred its call.
//line /usr/local/go/src/encoding/gob/error.go:34
func catchError(err *error) {
//line /usr/local/go/src/encoding/gob/error.go:34
	_go_fuzz_dep_.CoverTab[85180]++
							if e := recover(); e != nil {
//line /usr/local/go/src/encoding/gob/error.go:35
		_go_fuzz_dep_.CoverTab[85181]++
								ge, ok := e.(gobError)
								if !ok {
//line /usr/local/go/src/encoding/gob/error.go:37
			_go_fuzz_dep_.CoverTab[85183]++
									panic(e)
//line /usr/local/go/src/encoding/gob/error.go:38
			// _ = "end of CoverTab[85183]"
		} else {
//line /usr/local/go/src/encoding/gob/error.go:39
			_go_fuzz_dep_.CoverTab[85184]++
//line /usr/local/go/src/encoding/gob/error.go:39
			// _ = "end of CoverTab[85184]"
//line /usr/local/go/src/encoding/gob/error.go:39
		}
//line /usr/local/go/src/encoding/gob/error.go:39
		// _ = "end of CoverTab[85181]"
//line /usr/local/go/src/encoding/gob/error.go:39
		_go_fuzz_dep_.CoverTab[85182]++
								*err = ge.err
//line /usr/local/go/src/encoding/gob/error.go:40
		// _ = "end of CoverTab[85182]"
	} else {
//line /usr/local/go/src/encoding/gob/error.go:41
		_go_fuzz_dep_.CoverTab[85185]++
//line /usr/local/go/src/encoding/gob/error.go:41
		// _ = "end of CoverTab[85185]"
//line /usr/local/go/src/encoding/gob/error.go:41
	}
//line /usr/local/go/src/encoding/gob/error.go:41
	// _ = "end of CoverTab[85180]"
}

//line /usr/local/go/src/encoding/gob/error.go:42
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/encoding/gob/error.go:42
var _ = _go_fuzz_dep_.CoverTab
