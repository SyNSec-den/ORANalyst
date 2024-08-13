// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /snap/go/10455/src/strings/builder.go:5
package strings

//line /snap/go/10455/src/strings/builder.go:5
import (
//line /snap/go/10455/src/strings/builder.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/strings/builder.go:5
)
//line /snap/go/10455/src/strings/builder.go:5
import (
//line /snap/go/10455/src/strings/builder.go:5
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/strings/builder.go:5
)

import (
	"internal/bytealg"
	"unicode/utf8"
	"unsafe"
)

//line /snap/go/10455/src/strings/builder.go:16
type Builder struct {
	addr	*Builder
	buf	[]byte
}

//line /snap/go/10455/src/strings/builder.go:27
//go:nosplit
//go:nocheckptr
func noescape(p unsafe.Pointer) unsafe.Pointer {
//line /snap/go/10455/src/strings/builder.go:29
	_go_fuzz_dep_.CoverTab[877]++
							x := uintptr(p)
							return unsafe.Pointer(x ^ 0)
//line /snap/go/10455/src/strings/builder.go:31
	// _ = "end of CoverTab[877]"
}

func (b *Builder) copyCheck() {
//line /snap/go/10455/src/strings/builder.go:34
	_go_fuzz_dep_.CoverTab[878]++
							if b.addr == nil {
//line /snap/go/10455/src/strings/builder.go:35
		_go_fuzz_dep_.CoverTab[524929]++
//line /snap/go/10455/src/strings/builder.go:35
		_go_fuzz_dep_.CoverTab[879]++

//line /snap/go/10455/src/strings/builder.go:41
		b.addr = (*Builder)(noescape(unsafe.Pointer(b)))
//line /snap/go/10455/src/strings/builder.go:41
		// _ = "end of CoverTab[879]"
	} else {
//line /snap/go/10455/src/strings/builder.go:42
		_go_fuzz_dep_.CoverTab[524930]++
//line /snap/go/10455/src/strings/builder.go:42
		_go_fuzz_dep_.CoverTab[880]++
//line /snap/go/10455/src/strings/builder.go:42
		if b.addr != b {
//line /snap/go/10455/src/strings/builder.go:42
			_go_fuzz_dep_.CoverTab[524931]++
//line /snap/go/10455/src/strings/builder.go:42
			_go_fuzz_dep_.CoverTab[881]++
									panic("strings: illegal use of non-zero Builder copied by value")
//line /snap/go/10455/src/strings/builder.go:43
			// _ = "end of CoverTab[881]"
		} else {
//line /snap/go/10455/src/strings/builder.go:44
			_go_fuzz_dep_.CoverTab[524932]++
//line /snap/go/10455/src/strings/builder.go:44
			_go_fuzz_dep_.CoverTab[882]++
//line /snap/go/10455/src/strings/builder.go:44
			// _ = "end of CoverTab[882]"
//line /snap/go/10455/src/strings/builder.go:44
		}
//line /snap/go/10455/src/strings/builder.go:44
		// _ = "end of CoverTab[880]"
//line /snap/go/10455/src/strings/builder.go:44
	}
//line /snap/go/10455/src/strings/builder.go:44
	// _ = "end of CoverTab[878]"
}

//line /snap/go/10455/src/strings/builder.go:48
func (b *Builder) String() string {
//line /snap/go/10455/src/strings/builder.go:48
	_go_fuzz_dep_.CoverTab[883]++
							return unsafe.String(unsafe.SliceData(b.buf), len(b.buf))
//line /snap/go/10455/src/strings/builder.go:49
	// _ = "end of CoverTab[883]"
}

//line /snap/go/10455/src/strings/builder.go:53
func (b *Builder) Len() int {
//line /snap/go/10455/src/strings/builder.go:53
	_go_fuzz_dep_.CoverTab[884]++
//line /snap/go/10455/src/strings/builder.go:53
	return len(b.buf)
//line /snap/go/10455/src/strings/builder.go:53
	// _ = "end of CoverTab[884]"
//line /snap/go/10455/src/strings/builder.go:53
}

//line /snap/go/10455/src/strings/builder.go:58
func (b *Builder) Cap() int {
//line /snap/go/10455/src/strings/builder.go:58
	_go_fuzz_dep_.CoverTab[885]++
//line /snap/go/10455/src/strings/builder.go:58
	return cap(b.buf)
//line /snap/go/10455/src/strings/builder.go:58
	// _ = "end of CoverTab[885]"
//line /snap/go/10455/src/strings/builder.go:58
}

//line /snap/go/10455/src/strings/builder.go:61
func (b *Builder) Reset() {
//line /snap/go/10455/src/strings/builder.go:61
	_go_fuzz_dep_.CoverTab[886]++
							b.addr = nil
							b.buf = nil
//line /snap/go/10455/src/strings/builder.go:63
	// _ = "end of CoverTab[886]"
}

//line /snap/go/10455/src/strings/builder.go:68
func (b *Builder) grow(n int) {
//line /snap/go/10455/src/strings/builder.go:68
	_go_fuzz_dep_.CoverTab[887]++
							buf := bytealg.MakeNoZero(2*cap(b.buf) + n)[:len(b.buf)]
							copy(buf, b.buf)
							b.buf = buf
//line /snap/go/10455/src/strings/builder.go:71
	// _ = "end of CoverTab[887]"
}

//line /snap/go/10455/src/strings/builder.go:77
func (b *Builder) Grow(n int) {
//line /snap/go/10455/src/strings/builder.go:77
	_go_fuzz_dep_.CoverTab[888]++
							b.copyCheck()
							if n < 0 {
//line /snap/go/10455/src/strings/builder.go:79
		_go_fuzz_dep_.CoverTab[524933]++
//line /snap/go/10455/src/strings/builder.go:79
		_go_fuzz_dep_.CoverTab[890]++
								panic("strings.Builder.Grow: negative count")
//line /snap/go/10455/src/strings/builder.go:80
		// _ = "end of CoverTab[890]"
	} else {
//line /snap/go/10455/src/strings/builder.go:81
		_go_fuzz_dep_.CoverTab[524934]++
//line /snap/go/10455/src/strings/builder.go:81
		_go_fuzz_dep_.CoverTab[891]++
//line /snap/go/10455/src/strings/builder.go:81
		// _ = "end of CoverTab[891]"
//line /snap/go/10455/src/strings/builder.go:81
	}
//line /snap/go/10455/src/strings/builder.go:81
	// _ = "end of CoverTab[888]"
//line /snap/go/10455/src/strings/builder.go:81
	_go_fuzz_dep_.CoverTab[889]++
							if cap(b.buf)-len(b.buf) < n {
//line /snap/go/10455/src/strings/builder.go:82
		_go_fuzz_dep_.CoverTab[524935]++
//line /snap/go/10455/src/strings/builder.go:82
		_go_fuzz_dep_.CoverTab[892]++
								b.grow(n)
//line /snap/go/10455/src/strings/builder.go:83
		// _ = "end of CoverTab[892]"
	} else {
//line /snap/go/10455/src/strings/builder.go:84
		_go_fuzz_dep_.CoverTab[524936]++
//line /snap/go/10455/src/strings/builder.go:84
		_go_fuzz_dep_.CoverTab[893]++
//line /snap/go/10455/src/strings/builder.go:84
		// _ = "end of CoverTab[893]"
//line /snap/go/10455/src/strings/builder.go:84
	}
//line /snap/go/10455/src/strings/builder.go:84
	// _ = "end of CoverTab[889]"
}

//line /snap/go/10455/src/strings/builder.go:89
func (b *Builder) Write(p []byte) (int, error) {
//line /snap/go/10455/src/strings/builder.go:89
	_go_fuzz_dep_.CoverTab[894]++
							b.copyCheck()
							b.buf = append(b.buf, p...)
							return len(p), nil
//line /snap/go/10455/src/strings/builder.go:92
	// _ = "end of CoverTab[894]"
}

//line /snap/go/10455/src/strings/builder.go:97
func (b *Builder) WriteByte(c byte) error {
//line /snap/go/10455/src/strings/builder.go:97
	_go_fuzz_dep_.CoverTab[895]++
							b.copyCheck()
							b.buf = append(b.buf, c)
							return nil
//line /snap/go/10455/src/strings/builder.go:100
	// _ = "end of CoverTab[895]"
}

//line /snap/go/10455/src/strings/builder.go:105
func (b *Builder) WriteRune(r rune) (int, error) {
//line /snap/go/10455/src/strings/builder.go:105
	_go_fuzz_dep_.CoverTab[896]++
							b.copyCheck()
							n := len(b.buf)
							b.buf = utf8.AppendRune(b.buf, r)
							return len(b.buf) - n, nil
//line /snap/go/10455/src/strings/builder.go:109
	// _ = "end of CoverTab[896]"
}

//line /snap/go/10455/src/strings/builder.go:114
func (b *Builder) WriteString(s string) (int, error) {
//line /snap/go/10455/src/strings/builder.go:114
	_go_fuzz_dep_.CoverTab[897]++
							b.copyCheck()
							b.buf = append(b.buf, s...)
							return len(s), nil
//line /snap/go/10455/src/strings/builder.go:117
	// _ = "end of CoverTab[897]"
}

//line /snap/go/10455/src/strings/builder.go:118
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/strings/builder.go:118
var _ = _go_fuzz_dep_.CoverTab
