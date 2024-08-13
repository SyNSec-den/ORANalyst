// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/errors.go:5
// Package errors implements functions to manipulate errors.
package errors

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/errors.go:6
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/errors.go:6
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/errors.go:6
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/errors.go:6
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/errors.go:6
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/errors.go:6
)

import (
	"errors"
	"fmt"

	"google.golang.org/protobuf/internal/detrand"
)

// Error is a sentinel matching all errors produced by this package.
var Error = errors.New("protobuf error")

// New formats a string according to the format specifier and arguments and
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/errors.go:18
// returns an error that has a "proto" prefix.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/errors.go:20
func New(f string, x ...interface{}) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/errors.go:20
	_go_fuzz_dep_.CoverTab[48309]++
													return &prefixError{s: format(f, x...)}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/errors.go:21
	// _ = "end of CoverTab[48309]"
}

type prefixError struct{ s string }

var prefix = func() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/errors.go:26
	_go_fuzz_dep_.CoverTab[48310]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/errors.go:29
	if detrand.Bool() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/errors.go:29
		_go_fuzz_dep_.CoverTab[48311]++
														return "proto: "
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/errors.go:30
		// _ = "end of CoverTab[48311]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/errors.go:31
		_go_fuzz_dep_.CoverTab[48312]++
														return "proto: "
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/errors.go:32
		// _ = "end of CoverTab[48312]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/errors.go:33
	// _ = "end of CoverTab[48310]"
}()

func (e *prefixError) Error() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/errors.go:36
	_go_fuzz_dep_.CoverTab[48313]++
													return prefix + e.s
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/errors.go:37
	// _ = "end of CoverTab[48313]"
}

func (e *prefixError) Unwrap() error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/errors.go:40
	_go_fuzz_dep_.CoverTab[48314]++
													return Error
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/errors.go:41
	// _ = "end of CoverTab[48314]"
}

// Wrap returns an error that has a "proto" prefix, the formatted string described
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/errors.go:44
// by the format specifier and arguments, and a suffix of err. The error wraps err.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/errors.go:46
func Wrap(err error, f string, x ...interface{}) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/errors.go:46
	_go_fuzz_dep_.CoverTab[48315]++
													return &wrapError{
		s:	format(f, x...),
		err:	err,
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/errors.go:50
	// _ = "end of CoverTab[48315]"
}

type wrapError struct {
	s	string
	err	error
}

func (e *wrapError) Error() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/errors.go:58
	_go_fuzz_dep_.CoverTab[48316]++
													return format("%v%v: %v", prefix, e.s, e.err)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/errors.go:59
	// _ = "end of CoverTab[48316]"
}

func (e *wrapError) Unwrap() error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/errors.go:62
	_go_fuzz_dep_.CoverTab[48317]++
													return e.err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/errors.go:63
	// _ = "end of CoverTab[48317]"
}

func (e *wrapError) Is(target error) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/errors.go:66
	_go_fuzz_dep_.CoverTab[48318]++
													return target == Error
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/errors.go:67
	// _ = "end of CoverTab[48318]"
}

func format(f string, x ...interface{}) string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/errors.go:70
	_go_fuzz_dep_.CoverTab[48319]++

													for i := 0; i < len(x); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/errors.go:72
		_go_fuzz_dep_.CoverTab[48321]++
														switch e := x[i].(type) {
		case *prefixError:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/errors.go:74
			_go_fuzz_dep_.CoverTab[48322]++
															x[i] = e.s
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/errors.go:75
			// _ = "end of CoverTab[48322]"
		case *wrapError:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/errors.go:76
			_go_fuzz_dep_.CoverTab[48323]++
															x[i] = format("%v: %v", e.s, e.err)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/errors.go:77
			// _ = "end of CoverTab[48323]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/errors.go:78
		// _ = "end of CoverTab[48321]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/errors.go:79
	// _ = "end of CoverTab[48319]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/errors.go:79
	_go_fuzz_dep_.CoverTab[48320]++
													return fmt.Sprintf(f, x...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/errors.go:80
	// _ = "end of CoverTab[48320]"
}

func InvalidUTF8(name string) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/errors.go:83
	_go_fuzz_dep_.CoverTab[48324]++
													return New("field %v contains invalid UTF-8", name)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/errors.go:84
	// _ = "end of CoverTab[48324]"
}

func RequiredNotSet(name string) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/errors.go:87
	_go_fuzz_dep_.CoverTab[48325]++
													return New("required field %v not set", name)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/errors.go:88
	// _ = "end of CoverTab[48325]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/errors.go:89
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/errors.go:89
var _ = _go_fuzz_dep_.CoverTab
