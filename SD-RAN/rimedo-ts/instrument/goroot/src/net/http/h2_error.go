// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !nethttpomithttp2
// +build !nethttpomithttp2

//line /usr/local/go/src/net/http/h2_error.go:8
package http

//line /usr/local/go/src/net/http/h2_error.go:8
import (
//line /usr/local/go/src/net/http/h2_error.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/http/h2_error.go:8
)
//line /usr/local/go/src/net/http/h2_error.go:8
import (
//line /usr/local/go/src/net/http/h2_error.go:8
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/http/h2_error.go:8
)

import (
	"reflect"
)

func (e http2StreamError) As(target any) bool {
//line /usr/local/go/src/net/http/h2_error.go:14
	_go_fuzz_dep_.CoverTab[41344]++
							dst := reflect.ValueOf(target).Elem()
							dstType := dst.Type()
							if dstType.Kind() != reflect.Struct {
//line /usr/local/go/src/net/http/h2_error.go:17
		_go_fuzz_dep_.CoverTab[41349]++
								return false
//line /usr/local/go/src/net/http/h2_error.go:18
		// _ = "end of CoverTab[41349]"
	} else {
//line /usr/local/go/src/net/http/h2_error.go:19
		_go_fuzz_dep_.CoverTab[41350]++
//line /usr/local/go/src/net/http/h2_error.go:19
		// _ = "end of CoverTab[41350]"
//line /usr/local/go/src/net/http/h2_error.go:19
	}
//line /usr/local/go/src/net/http/h2_error.go:19
	// _ = "end of CoverTab[41344]"
//line /usr/local/go/src/net/http/h2_error.go:19
	_go_fuzz_dep_.CoverTab[41345]++
							src := reflect.ValueOf(e)
							srcType := src.Type()
							numField := srcType.NumField()
							if dstType.NumField() != numField {
//line /usr/local/go/src/net/http/h2_error.go:23
		_go_fuzz_dep_.CoverTab[41351]++
								return false
//line /usr/local/go/src/net/http/h2_error.go:24
		// _ = "end of CoverTab[41351]"
	} else {
//line /usr/local/go/src/net/http/h2_error.go:25
		_go_fuzz_dep_.CoverTab[41352]++
//line /usr/local/go/src/net/http/h2_error.go:25
		// _ = "end of CoverTab[41352]"
//line /usr/local/go/src/net/http/h2_error.go:25
	}
//line /usr/local/go/src/net/http/h2_error.go:25
	// _ = "end of CoverTab[41345]"
//line /usr/local/go/src/net/http/h2_error.go:25
	_go_fuzz_dep_.CoverTab[41346]++
							for i := 0; i < numField; i++ {
//line /usr/local/go/src/net/http/h2_error.go:26
		_go_fuzz_dep_.CoverTab[41353]++
								sf := srcType.Field(i)
								df := dstType.Field(i)
								if sf.Name != df.Name || func() bool {
//line /usr/local/go/src/net/http/h2_error.go:29
			_go_fuzz_dep_.CoverTab[41354]++
//line /usr/local/go/src/net/http/h2_error.go:29
			return !sf.Type.ConvertibleTo(df.Type)
//line /usr/local/go/src/net/http/h2_error.go:29
			// _ = "end of CoverTab[41354]"
//line /usr/local/go/src/net/http/h2_error.go:29
		}() {
//line /usr/local/go/src/net/http/h2_error.go:29
			_go_fuzz_dep_.CoverTab[41355]++
									return false
//line /usr/local/go/src/net/http/h2_error.go:30
			// _ = "end of CoverTab[41355]"
		} else {
//line /usr/local/go/src/net/http/h2_error.go:31
			_go_fuzz_dep_.CoverTab[41356]++
//line /usr/local/go/src/net/http/h2_error.go:31
			// _ = "end of CoverTab[41356]"
//line /usr/local/go/src/net/http/h2_error.go:31
		}
//line /usr/local/go/src/net/http/h2_error.go:31
		// _ = "end of CoverTab[41353]"
	}
//line /usr/local/go/src/net/http/h2_error.go:32
	// _ = "end of CoverTab[41346]"
//line /usr/local/go/src/net/http/h2_error.go:32
	_go_fuzz_dep_.CoverTab[41347]++
							for i := 0; i < numField; i++ {
//line /usr/local/go/src/net/http/h2_error.go:33
		_go_fuzz_dep_.CoverTab[41357]++
								df := dst.Field(i)
								df.Set(src.Field(i).Convert(df.Type()))
//line /usr/local/go/src/net/http/h2_error.go:35
		// _ = "end of CoverTab[41357]"
	}
//line /usr/local/go/src/net/http/h2_error.go:36
	// _ = "end of CoverTab[41347]"
//line /usr/local/go/src/net/http/h2_error.go:36
	_go_fuzz_dep_.CoverTab[41348]++
							return true
//line /usr/local/go/src/net/http/h2_error.go:37
	// _ = "end of CoverTab[41348]"
}

//line /usr/local/go/src/net/http/h2_error.go:38
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/http/h2_error.go:38
var _ = _go_fuzz_dep_.CoverTab
