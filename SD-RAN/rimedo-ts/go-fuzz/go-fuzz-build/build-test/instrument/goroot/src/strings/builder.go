// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/strings/builder.go:5
package strings

//line /usr/local/go/src/strings/builder.go:5
import (
//line /usr/local/go/src/strings/builder.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/strings/builder.go:5
)
//line /usr/local/go/src/strings/builder.go:5
import (
//line /usr/local/go/src/strings/builder.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/strings/builder.go:5
)

import (
	"unicode/utf8"
	"unsafe"
)

//line /usr/local/go/src/strings/builder.go:15
type Builder struct {
	addr	*Builder
	buf	[]byte
}

//line /usr/local/go/src/strings/builder.go:26
//go:nosplit
//go:nocheckptr
func noescape(p unsafe.Pointer) unsafe.Pointer {
//line /usr/local/go/src/strings/builder.go:28
	_go_fuzz_dep_.CoverTab[864]++
						x := uintptr(p)
						return unsafe.Pointer(x ^ 0)
//line /usr/local/go/src/strings/builder.go:30
	// _ = "end of CoverTab[864]"
}

func (b *Builder) copyCheck() {
//line /usr/local/go/src/strings/builder.go:33
	_go_fuzz_dep_.CoverTab[865]++
						if b.addr == nil {
//line /usr/local/go/src/strings/builder.go:34
		_go_fuzz_dep_.CoverTab[866]++

//line /usr/local/go/src/strings/builder.go:40
		b.addr = (*Builder)(noescape(unsafe.Pointer(b)))
//line /usr/local/go/src/strings/builder.go:40
		// _ = "end of CoverTab[866]"
	} else {
//line /usr/local/go/src/strings/builder.go:41
		_go_fuzz_dep_.CoverTab[867]++
//line /usr/local/go/src/strings/builder.go:41
		if b.addr != b {
//line /usr/local/go/src/strings/builder.go:41
			_go_fuzz_dep_.CoverTab[868]++
								panic("strings: illegal use of non-zero Builder copied by value")
//line /usr/local/go/src/strings/builder.go:42
			// _ = "end of CoverTab[868]"
		} else {
//line /usr/local/go/src/strings/builder.go:43
			_go_fuzz_dep_.CoverTab[869]++
//line /usr/local/go/src/strings/builder.go:43
			// _ = "end of CoverTab[869]"
//line /usr/local/go/src/strings/builder.go:43
		}
//line /usr/local/go/src/strings/builder.go:43
		// _ = "end of CoverTab[867]"
//line /usr/local/go/src/strings/builder.go:43
	}
//line /usr/local/go/src/strings/builder.go:43
	// _ = "end of CoverTab[865]"
}

//line /usr/local/go/src/strings/builder.go:47
func (b *Builder) String() string {
//line /usr/local/go/src/strings/builder.go:47
	_go_fuzz_dep_.CoverTab[870]++
						return unsafe.String(unsafe.SliceData(b.buf), len(b.buf))
//line /usr/local/go/src/strings/builder.go:48
	// _ = "end of CoverTab[870]"
}

//line /usr/local/go/src/strings/builder.go:52
func (b *Builder) Len() int {
//line /usr/local/go/src/strings/builder.go:52
	_go_fuzz_dep_.CoverTab[871]++
//line /usr/local/go/src/strings/builder.go:52
	return len(b.buf)
//line /usr/local/go/src/strings/builder.go:52
	// _ = "end of CoverTab[871]"
//line /usr/local/go/src/strings/builder.go:52
}

//line /usr/local/go/src/strings/builder.go:57
func (b *Builder) Cap() int {
//line /usr/local/go/src/strings/builder.go:57
	_go_fuzz_dep_.CoverTab[872]++
//line /usr/local/go/src/strings/builder.go:57
	return cap(b.buf)
//line /usr/local/go/src/strings/builder.go:57
	// _ = "end of CoverTab[872]"
//line /usr/local/go/src/strings/builder.go:57
}

//line /usr/local/go/src/strings/builder.go:60
func (b *Builder) Reset() {
//line /usr/local/go/src/strings/builder.go:60
	_go_fuzz_dep_.CoverTab[873]++
						b.addr = nil
						b.buf = nil
//line /usr/local/go/src/strings/builder.go:62
	// _ = "end of CoverTab[873]"
}

//line /usr/local/go/src/strings/builder.go:67
func (b *Builder) grow(n int) {
//line /usr/local/go/src/strings/builder.go:67
	_go_fuzz_dep_.CoverTab[874]++
						buf := make([]byte, len(b.buf), 2*cap(b.buf)+n)
						copy(buf, b.buf)
						b.buf = buf
//line /usr/local/go/src/strings/builder.go:70
	// _ = "end of CoverTab[874]"
}

//line /usr/local/go/src/strings/builder.go:76
func (b *Builder) Grow(n int) {
//line /usr/local/go/src/strings/builder.go:76
	_go_fuzz_dep_.CoverTab[875]++
						b.copyCheck()
						if n < 0 {
//line /usr/local/go/src/strings/builder.go:78
		_go_fuzz_dep_.CoverTab[877]++
							panic("strings.Builder.Grow: negative count")
//line /usr/local/go/src/strings/builder.go:79
		// _ = "end of CoverTab[877]"
	} else {
//line /usr/local/go/src/strings/builder.go:80
		_go_fuzz_dep_.CoverTab[878]++
//line /usr/local/go/src/strings/builder.go:80
		// _ = "end of CoverTab[878]"
//line /usr/local/go/src/strings/builder.go:80
	}
//line /usr/local/go/src/strings/builder.go:80
	// _ = "end of CoverTab[875]"
//line /usr/local/go/src/strings/builder.go:80
	_go_fuzz_dep_.CoverTab[876]++
						if cap(b.buf)-len(b.buf) < n {
//line /usr/local/go/src/strings/builder.go:81
		_go_fuzz_dep_.CoverTab[879]++
							b.grow(n)
//line /usr/local/go/src/strings/builder.go:82
		// _ = "end of CoverTab[879]"
	} else {
//line /usr/local/go/src/strings/builder.go:83
		_go_fuzz_dep_.CoverTab[880]++
//line /usr/local/go/src/strings/builder.go:83
		// _ = "end of CoverTab[880]"
//line /usr/local/go/src/strings/builder.go:83
	}
//line /usr/local/go/src/strings/builder.go:83
	// _ = "end of CoverTab[876]"
}

//line /usr/local/go/src/strings/builder.go:88
func (b *Builder) Write(p []byte) (int, error) {
//line /usr/local/go/src/strings/builder.go:88
	_go_fuzz_dep_.CoverTab[881]++
						b.copyCheck()
						b.buf = append(b.buf, p...)
						return len(p), nil
//line /usr/local/go/src/strings/builder.go:91
	// _ = "end of CoverTab[881]"
}

//line /usr/local/go/src/strings/builder.go:96
func (b *Builder) WriteByte(c byte) error {
//line /usr/local/go/src/strings/builder.go:96
	_go_fuzz_dep_.CoverTab[882]++
						b.copyCheck()
						b.buf = append(b.buf, c)
						return nil
//line /usr/local/go/src/strings/builder.go:99
	// _ = "end of CoverTab[882]"
}

//line /usr/local/go/src/strings/builder.go:104
func (b *Builder) WriteRune(r rune) (int, error) {
//line /usr/local/go/src/strings/builder.go:104
	_go_fuzz_dep_.CoverTab[883]++
							b.copyCheck()
							n := len(b.buf)
							b.buf = utf8.AppendRune(b.buf, r)
							return len(b.buf) - n, nil
//line /usr/local/go/src/strings/builder.go:108
	// _ = "end of CoverTab[883]"
}

//line /usr/local/go/src/strings/builder.go:113
func (b *Builder) WriteString(s string) (int, error) {
//line /usr/local/go/src/strings/builder.go:113
	_go_fuzz_dep_.CoverTab[884]++
							b.copyCheck()
							b.buf = append(b.buf, s...)
							return len(s), nil
//line /usr/local/go/src/strings/builder.go:116
	// _ = "end of CoverTab[884]"
}

//line /usr/local/go/src/strings/builder.go:117
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/strings/builder.go:117
var _ = _go_fuzz_dep_.CoverTab
