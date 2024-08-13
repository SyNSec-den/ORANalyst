// Copyright (c) 2015-2016 Dave Collins <dave@davec.name>
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

// NOTE: Due to the following build constraints, this file will only be compiled
// when the code is not running on Google App Engine, compiled by GopherJS, and
// "-tags safe" is not added to the go build command line.  The "disableunsafe"
// tag is deprecated and thus should not be used.
// Go versions prior to 1.4 are disabled because they use a different layout
// for interfaces which make the implementation of unsafeReflectValue more complex.
// +build !js,!appengine,!safe,!disableunsafe,go1.4

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:23
package spew

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:23
import (
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:23
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:23
)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:23
import (
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:23
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:23
)

import (
	"reflect"
	"unsafe"
)

const (
	// UnsafeDisabled is a build-time constant which specifies whether or
	// not access to the unsafe package is available.
	UnsafeDisabled	= false

	// ptrSize is the size of a pointer on the current arch.
	ptrSize	= unsafe.Sizeof((*byte)(nil))
)

type flag uintptr

var (
	// flagRO indicates whether the value field of a reflect.Value
	// is read-only.
	flagRO	flag

	// flagAddr indicates whether the address of the reflect.Value's
	// value may be taken.
	flagAddr	flag
)

// flagKindMask holds the bits that make up the kind
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:51
// part of the flags field. In all the supported versions,
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:51
// it is in the lower 5 bits.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:54
const flagKindMask = flag(0x1f)

// Different versions of Go have used different
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:56
// bit layouts for the flags type. This table
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:56
// records the known combinations.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:59
var okFlags = []struct {
	ro, addr flag
}{{

	ro:	1 << 5,
	addr:	1 << 7,
}, {

	ro:	1<<5 | 1<<6,
	addr:	1 << 8,
}}

var flagValOffset = func() uintptr {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:71
	_go_fuzz_dep_.CoverTab[81530]++
											field, ok := reflect.TypeOf(reflect.Value{}).FieldByName("flag")
											if !ok {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:73
		_go_fuzz_dep_.CoverTab[81532]++
												panic("reflect.Value has no flag field")
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:74
		// _ = "end of CoverTab[81532]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:75
		_go_fuzz_dep_.CoverTab[81533]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:75
		// _ = "end of CoverTab[81533]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:75
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:75
	// _ = "end of CoverTab[81530]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:75
	_go_fuzz_dep_.CoverTab[81531]++
											return field.Offset
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:76
	// _ = "end of CoverTab[81531]"
}()

// flagField returns a pointer to the flag field of a reflect.Value.
func flagField(v *reflect.Value) *flag {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:80
	_go_fuzz_dep_.CoverTab[81534]++
											return (*flag)(unsafe.Pointer(uintptr(unsafe.Pointer(v)) + flagValOffset))
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:81
	// _ = "end of CoverTab[81534]"
}

// unsafeReflectValue converts the passed reflect.Value into a one that bypasses
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:84
// the typical safety restrictions preventing access to unaddressable and
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:84
// unexported data.  It works by digging the raw pointer to the underlying
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:84
// value out of the protected value and generating a new unprotected (unsafe)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:84
// reflect.Value to it.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:84
//
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:84
// This allows us to check for implementations of the Stringer and error
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:84
// interfaces to be used for pretty printing ordinarily unaddressable and
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:84
// inaccessible values such as unexported struct fields.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:93
func unsafeReflectValue(v reflect.Value) reflect.Value {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:93
	_go_fuzz_dep_.CoverTab[81535]++
											if !v.IsValid() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:94
		_go_fuzz_dep_.CoverTab[81537]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:94
		return (v.CanInterface() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:94
			_go_fuzz_dep_.CoverTab[81538]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:94
			return v.CanAddr()
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:94
			// _ = "end of CoverTab[81538]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:94
		}())
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:94
		// _ = "end of CoverTab[81537]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:94
	}() {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:94
		_go_fuzz_dep_.CoverTab[81539]++
												return v
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:95
		// _ = "end of CoverTab[81539]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:96
		_go_fuzz_dep_.CoverTab[81540]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:96
		// _ = "end of CoverTab[81540]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:96
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:96
	// _ = "end of CoverTab[81535]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:96
	_go_fuzz_dep_.CoverTab[81536]++
											flagFieldPtr := flagField(&v)
											*flagFieldPtr &^= flagRO
											*flagFieldPtr |= flagAddr
											return v
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:100
	// _ = "end of CoverTab[81536]"
}

// Sanity checks against future reflect package changes
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:103
// to the type or semantics of the Value.flag field.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:105
func init() {
	field, ok := reflect.TypeOf(reflect.Value{}).FieldByName("flag")
	if !ok {
		panic("reflect.Value has no flag field")
	}
	if field.Type.Kind() != reflect.TypeOf(flag(0)).Kind() {
		panic("reflect.Value flag field has changed kind")
	}
	type t0 int
	var t struct {
		A	t0
		// t0 will have flagEmbedRO set.
		t0
		// a will have flagStickyRO set
		a	t0
	}
											vA := reflect.ValueOf(t).FieldByName("A")
											va := reflect.ValueOf(t).FieldByName("a")
											vt0 := reflect.ValueOf(t).FieldByName("t0")

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:127
	flagPublic := *flagField(&vA)
											flagWithRO := *flagField(&va) | *flagField(&vt0)
											flagRO = flagPublic ^ flagWithRO

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:133
	vPtrA := reflect.ValueOf(&t).Elem().FieldByName("A")
											flagNoPtr := *flagField(&vA)
											flagPtr := *flagField(&vPtrA)
											flagAddr = flagNoPtr ^ flagPtr

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:139
	for _, f := range okFlags {
		if flagRO == f.ro && flagAddr == f.addr {
			return
		}
	}
	panic("reflect.Value read-only flag has changed semantics")
}

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:145
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/bypass.go:145
var _ = _go_fuzz_dep_.CoverTab
