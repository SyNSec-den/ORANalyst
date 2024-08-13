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
	_go_fuzz_dep_.CoverTab[3119]++
						x := uintptr(p)
						return unsafe.Pointer(x ^ 0)
//line /usr/local/go/src/strings/builder.go:30
	// _ = "end of CoverTab[3119]"
}

func (b *Builder) copyCheck() {
//line /usr/local/go/src/strings/builder.go:33
	_go_fuzz_dep_.CoverTab[3120]++
						if b.addr == nil {
//line /usr/local/go/src/strings/builder.go:34
		_go_fuzz_dep_.CoverTab[3121]++

//line /usr/local/go/src/strings/builder.go:40
		b.addr = (*Builder)(noescape(unsafe.Pointer(b)))
//line /usr/local/go/src/strings/builder.go:40
		// _ = "end of CoverTab[3121]"
	} else {
//line /usr/local/go/src/strings/builder.go:41
		_go_fuzz_dep_.CoverTab[3122]++
//line /usr/local/go/src/strings/builder.go:41
		if b.addr != b {
//line /usr/local/go/src/strings/builder.go:41
			_go_fuzz_dep_.CoverTab[3123]++
								panic("strings: illegal use of non-zero Builder copied by value")
//line /usr/local/go/src/strings/builder.go:42
			// _ = "end of CoverTab[3123]"
		} else {
//line /usr/local/go/src/strings/builder.go:43
			_go_fuzz_dep_.CoverTab[3124]++
//line /usr/local/go/src/strings/builder.go:43
			// _ = "end of CoverTab[3124]"
//line /usr/local/go/src/strings/builder.go:43
		}
//line /usr/local/go/src/strings/builder.go:43
		// _ = "end of CoverTab[3122]"
//line /usr/local/go/src/strings/builder.go:43
	}
//line /usr/local/go/src/strings/builder.go:43
	// _ = "end of CoverTab[3120]"
}

//line /usr/local/go/src/strings/builder.go:47
func (b *Builder) String() string {
//line /usr/local/go/src/strings/builder.go:47
	_go_fuzz_dep_.CoverTab[3125]++
						return unsafe.String(unsafe.SliceData(b.buf), len(b.buf))
//line /usr/local/go/src/strings/builder.go:48
	// _ = "end of CoverTab[3125]"
}

//line /usr/local/go/src/strings/builder.go:52
func (b *Builder) Len() int {
//line /usr/local/go/src/strings/builder.go:52
	_go_fuzz_dep_.CoverTab[3126]++
//line /usr/local/go/src/strings/builder.go:52
	return len(b.buf)
//line /usr/local/go/src/strings/builder.go:52
	// _ = "end of CoverTab[3126]"
//line /usr/local/go/src/strings/builder.go:52
}

//line /usr/local/go/src/strings/builder.go:57
func (b *Builder) Cap() int {
//line /usr/local/go/src/strings/builder.go:57
	_go_fuzz_dep_.CoverTab[3127]++
//line /usr/local/go/src/strings/builder.go:57
	return cap(b.buf)
//line /usr/local/go/src/strings/builder.go:57
	// _ = "end of CoverTab[3127]"
//line /usr/local/go/src/strings/builder.go:57
}

//line /usr/local/go/src/strings/builder.go:60
func (b *Builder) Reset() {
//line /usr/local/go/src/strings/builder.go:60
	_go_fuzz_dep_.CoverTab[3128]++
						b.addr = nil
						b.buf = nil
//line /usr/local/go/src/strings/builder.go:62
	// _ = "end of CoverTab[3128]"
}

//line /usr/local/go/src/strings/builder.go:67
func (b *Builder) grow(n int) {
//line /usr/local/go/src/strings/builder.go:67
	_go_fuzz_dep_.CoverTab[3129]++
						buf := make([]byte, len(b.buf), 2*cap(b.buf)+n)
						copy(buf, b.buf)
						b.buf = buf
//line /usr/local/go/src/strings/builder.go:70
	// _ = "end of CoverTab[3129]"
}

//line /usr/local/go/src/strings/builder.go:76
func (b *Builder) Grow(n int) {
//line /usr/local/go/src/strings/builder.go:76
	_go_fuzz_dep_.CoverTab[3130]++
						b.copyCheck()
						if n < 0 {
//line /usr/local/go/src/strings/builder.go:78
		_go_fuzz_dep_.CoverTab[3132]++
							panic("strings.Builder.Grow: negative count")
//line /usr/local/go/src/strings/builder.go:79
		// _ = "end of CoverTab[3132]"
	} else {
//line /usr/local/go/src/strings/builder.go:80
		_go_fuzz_dep_.CoverTab[3133]++
//line /usr/local/go/src/strings/builder.go:80
		// _ = "end of CoverTab[3133]"
//line /usr/local/go/src/strings/builder.go:80
	}
//line /usr/local/go/src/strings/builder.go:80
	// _ = "end of CoverTab[3130]"
//line /usr/local/go/src/strings/builder.go:80
	_go_fuzz_dep_.CoverTab[3131]++
						if cap(b.buf)-len(b.buf) < n {
//line /usr/local/go/src/strings/builder.go:81
		_go_fuzz_dep_.CoverTab[3134]++
							b.grow(n)
//line /usr/local/go/src/strings/builder.go:82
		// _ = "end of CoverTab[3134]"
	} else {
//line /usr/local/go/src/strings/builder.go:83
		_go_fuzz_dep_.CoverTab[3135]++
//line /usr/local/go/src/strings/builder.go:83
		// _ = "end of CoverTab[3135]"
//line /usr/local/go/src/strings/builder.go:83
	}
//line /usr/local/go/src/strings/builder.go:83
	// _ = "end of CoverTab[3131]"
}

//line /usr/local/go/src/strings/builder.go:88
func (b *Builder) Write(p []byte) (int, error) {
//line /usr/local/go/src/strings/builder.go:88
	_go_fuzz_dep_.CoverTab[3136]++
						b.copyCheck()
						b.buf = append(b.buf, p...)
						return len(p), nil
//line /usr/local/go/src/strings/builder.go:91
	// _ = "end of CoverTab[3136]"
}

//line /usr/local/go/src/strings/builder.go:96
func (b *Builder) WriteByte(c byte) error {
//line /usr/local/go/src/strings/builder.go:96
	_go_fuzz_dep_.CoverTab[3137]++
						b.copyCheck()
						b.buf = append(b.buf, c)
						return nil
//line /usr/local/go/src/strings/builder.go:99
	// _ = "end of CoverTab[3137]"
}

//line /usr/local/go/src/strings/builder.go:104
func (b *Builder) WriteRune(r rune) (int, error) {
//line /usr/local/go/src/strings/builder.go:104
	_go_fuzz_dep_.CoverTab[3138]++
							b.copyCheck()
							n := len(b.buf)
							b.buf = utf8.AppendRune(b.buf, r)
							return len(b.buf) - n, nil
//line /usr/local/go/src/strings/builder.go:108
	// _ = "end of CoverTab[3138]"
}

//line /usr/local/go/src/strings/builder.go:113
func (b *Builder) WriteString(s string) (int, error) {
//line /usr/local/go/src/strings/builder.go:113
	_go_fuzz_dep_.CoverTab[3139]++
							b.copyCheck()
							b.buf = append(b.buf, s...)
							return len(s), nil
//line /usr/local/go/src/strings/builder.go:116
	// _ = "end of CoverTab[3139]"
}

//line /usr/local/go/src/strings/builder.go:117
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/strings/builder.go:117
var _ = _go_fuzz_dep_.CoverTab
