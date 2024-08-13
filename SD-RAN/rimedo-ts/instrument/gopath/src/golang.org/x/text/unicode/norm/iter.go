// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:5
package norm

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:5
)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:5
)

import (
	"fmt"
	"unicode/utf8"
)

// MaxSegmentSize is the maximum size of a byte buffer needed to consider any
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:12
// sequence of starter and non-starter runes for the purpose of normalization.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:14
const MaxSegmentSize = maxByteBufferSize

// An Iter iterates over a string or byte slice, while normalizing it
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:16
// to a given Form.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:18
type Iter struct {
	rb	reorderBuffer
	buf	[maxByteBufferSize]byte
	info	Properties	// first character saved from previous iteration
	next	iterFunc	// implementation of next depends on form
	asciiF	iterFunc

	p		int	// current position in input source
	multiSeg	[]byte	// remainder of multi-segment decomposition
}

type iterFunc func(*Iter) []byte

// Init initializes i to iterate over src after normalizing it to Form f.
func (i *Iter) Init(f Form, src []byte) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:32
	_go_fuzz_dep_.CoverTab[70483]++
											i.p = 0
											if len(src) == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:34
		_go_fuzz_dep_.CoverTab[70485]++
												i.setDone()
												i.rb.nsrc = 0
												return
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:37
		// _ = "end of CoverTab[70485]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:38
		_go_fuzz_dep_.CoverTab[70486]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:38
		// _ = "end of CoverTab[70486]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:38
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:38
	// _ = "end of CoverTab[70483]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:38
	_go_fuzz_dep_.CoverTab[70484]++
											i.multiSeg = nil
											i.rb.init(f, src)
											i.next = i.rb.f.nextMain
											i.asciiF = nextASCIIBytes
											i.info = i.rb.f.info(i.rb.src, i.p)
											i.rb.ss.first(i.info)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:44
	// _ = "end of CoverTab[70484]"
}

// InitString initializes i to iterate over src after normalizing it to Form f.
func (i *Iter) InitString(f Form, src string) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:48
	_go_fuzz_dep_.CoverTab[70487]++
											i.p = 0
											if len(src) == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:50
		_go_fuzz_dep_.CoverTab[70489]++
												i.setDone()
												i.rb.nsrc = 0
												return
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:53
		// _ = "end of CoverTab[70489]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:54
		_go_fuzz_dep_.CoverTab[70490]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:54
		// _ = "end of CoverTab[70490]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:54
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:54
	// _ = "end of CoverTab[70487]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:54
	_go_fuzz_dep_.CoverTab[70488]++
											i.multiSeg = nil
											i.rb.initString(f, src)
											i.next = i.rb.f.nextMain
											i.asciiF = nextASCIIString
											i.info = i.rb.f.info(i.rb.src, i.p)
											i.rb.ss.first(i.info)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:60
	// _ = "end of CoverTab[70488]"
}

// Seek sets the segment to be returned by the next call to Next to start
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:63
// at position p.  It is the responsibility of the caller to set p to the
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:63
// start of a segment.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:66
func (i *Iter) Seek(offset int64, whence int) (int64, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:66
	_go_fuzz_dep_.CoverTab[70491]++
											var abs int64
											switch whence {
	case 0:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:69
		_go_fuzz_dep_.CoverTab[70495]++
												abs = offset
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:70
		// _ = "end of CoverTab[70495]"
	case 1:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:71
		_go_fuzz_dep_.CoverTab[70496]++
												abs = int64(i.p) + offset
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:72
		// _ = "end of CoverTab[70496]"
	case 2:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:73
		_go_fuzz_dep_.CoverTab[70497]++
												abs = int64(i.rb.nsrc) + offset
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:74
		// _ = "end of CoverTab[70497]"
	default:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:75
		_go_fuzz_dep_.CoverTab[70498]++
												return 0, fmt.Errorf("norm: invalid whence")
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:76
		// _ = "end of CoverTab[70498]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:77
	// _ = "end of CoverTab[70491]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:77
	_go_fuzz_dep_.CoverTab[70492]++
											if abs < 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:78
		_go_fuzz_dep_.CoverTab[70499]++
												return 0, fmt.Errorf("norm: negative position")
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:79
		// _ = "end of CoverTab[70499]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:80
		_go_fuzz_dep_.CoverTab[70500]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:80
		// _ = "end of CoverTab[70500]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:80
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:80
	// _ = "end of CoverTab[70492]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:80
	_go_fuzz_dep_.CoverTab[70493]++
											if int(abs) >= i.rb.nsrc {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:81
		_go_fuzz_dep_.CoverTab[70501]++
												i.setDone()
												return int64(i.p), nil
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:83
		// _ = "end of CoverTab[70501]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:84
		_go_fuzz_dep_.CoverTab[70502]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:84
		// _ = "end of CoverTab[70502]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:84
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:84
	// _ = "end of CoverTab[70493]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:84
	_go_fuzz_dep_.CoverTab[70494]++
											i.p = int(abs)
											i.multiSeg = nil
											i.next = i.rb.f.nextMain
											i.info = i.rb.f.info(i.rb.src, i.p)
											i.rb.ss.first(i.info)
											return abs, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:90
	// _ = "end of CoverTab[70494]"
}

// returnSlice returns a slice of the underlying input type as a byte slice.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:93
// If the underlying is of type []byte, it will simply return a slice.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:93
// If the underlying is of type string, it will copy the slice to the buffer
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:93
// and return that.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:97
func (i *Iter) returnSlice(a, b int) []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:97
	_go_fuzz_dep_.CoverTab[70503]++
											if i.rb.src.bytes == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:98
		_go_fuzz_dep_.CoverTab[70505]++
												return i.buf[:copy(i.buf[:], i.rb.src.str[a:b])]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:99
		// _ = "end of CoverTab[70505]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:100
		_go_fuzz_dep_.CoverTab[70506]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:100
		// _ = "end of CoverTab[70506]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:100
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:100
	// _ = "end of CoverTab[70503]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:100
	_go_fuzz_dep_.CoverTab[70504]++
											return i.rb.src.bytes[a:b]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:101
	// _ = "end of CoverTab[70504]"
}

// Pos returns the byte position at which the next call to Next will commence processing.
func (i *Iter) Pos() int {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:105
	_go_fuzz_dep_.CoverTab[70507]++
											return i.p
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:106
	// _ = "end of CoverTab[70507]"
}

func (i *Iter) setDone() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:109
	_go_fuzz_dep_.CoverTab[70508]++
											i.next = nextDone
											i.p = i.rb.nsrc
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:111
	// _ = "end of CoverTab[70508]"
}

// Done returns true if there is no more input to process.
func (i *Iter) Done() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:115
	_go_fuzz_dep_.CoverTab[70509]++
											return i.p >= i.rb.nsrc
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:116
	// _ = "end of CoverTab[70509]"
}

// Next returns f(i.input[i.Pos():n]), where n is a boundary of i.input.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:119
// For any input a and b for which f(a) == f(b), subsequent calls
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:119
// to Next will return the same segments.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:119
// Modifying runes are grouped together with the preceding starter, if such a starter exists.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:119
// Although not guaranteed, n will typically be the smallest possible n.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:124
func (i *Iter) Next() []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:124
	_go_fuzz_dep_.CoverTab[70510]++
											return i.next(i)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:125
	// _ = "end of CoverTab[70510]"
}

func nextASCIIBytes(i *Iter) []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:128
	_go_fuzz_dep_.CoverTab[70511]++
											p := i.p + 1
											if p >= i.rb.nsrc {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:130
		_go_fuzz_dep_.CoverTab[70514]++
												p0 := i.p
												i.setDone()
												return i.rb.src.bytes[p0:p]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:133
		// _ = "end of CoverTab[70514]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:134
		_go_fuzz_dep_.CoverTab[70515]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:134
		// _ = "end of CoverTab[70515]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:134
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:134
	// _ = "end of CoverTab[70511]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:134
	_go_fuzz_dep_.CoverTab[70512]++
											if i.rb.src.bytes[p] < utf8.RuneSelf {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:135
		_go_fuzz_dep_.CoverTab[70516]++
												p0 := i.p
												i.p = p
												return i.rb.src.bytes[p0:p]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:138
		// _ = "end of CoverTab[70516]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:139
		_go_fuzz_dep_.CoverTab[70517]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:139
		// _ = "end of CoverTab[70517]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:139
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:139
	// _ = "end of CoverTab[70512]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:139
	_go_fuzz_dep_.CoverTab[70513]++
											i.info = i.rb.f.info(i.rb.src, i.p)
											i.next = i.rb.f.nextMain
											return i.next(i)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:142
	// _ = "end of CoverTab[70513]"
}

func nextASCIIString(i *Iter) []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:145
	_go_fuzz_dep_.CoverTab[70518]++
											p := i.p + 1
											if p >= i.rb.nsrc {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:147
		_go_fuzz_dep_.CoverTab[70521]++
												i.buf[0] = i.rb.src.str[i.p]
												i.setDone()
												return i.buf[:1]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:150
		// _ = "end of CoverTab[70521]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:151
		_go_fuzz_dep_.CoverTab[70522]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:151
		// _ = "end of CoverTab[70522]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:151
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:151
	// _ = "end of CoverTab[70518]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:151
	_go_fuzz_dep_.CoverTab[70519]++
											if i.rb.src.str[p] < utf8.RuneSelf {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:152
		_go_fuzz_dep_.CoverTab[70523]++
												i.buf[0] = i.rb.src.str[i.p]
												i.p = p
												return i.buf[:1]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:155
		// _ = "end of CoverTab[70523]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:156
		_go_fuzz_dep_.CoverTab[70524]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:156
		// _ = "end of CoverTab[70524]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:156
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:156
	// _ = "end of CoverTab[70519]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:156
	_go_fuzz_dep_.CoverTab[70520]++
											i.info = i.rb.f.info(i.rb.src, i.p)
											i.next = i.rb.f.nextMain
											return i.next(i)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:159
	// _ = "end of CoverTab[70520]"
}

func nextHangul(i *Iter) []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:162
	_go_fuzz_dep_.CoverTab[70525]++
											p := i.p
											next := p + hangulUTF8Size
											if next >= i.rb.nsrc {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:165
		_go_fuzz_dep_.CoverTab[70527]++
												i.setDone()
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:166
		// _ = "end of CoverTab[70527]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:167
		_go_fuzz_dep_.CoverTab[70528]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:167
		if i.rb.src.hangul(next) == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:167
			_go_fuzz_dep_.CoverTab[70529]++
													i.rb.ss.next(i.info)
													i.info = i.rb.f.info(i.rb.src, i.p)
													i.next = i.rb.f.nextMain
													return i.next(i)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:171
			// _ = "end of CoverTab[70529]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:172
			_go_fuzz_dep_.CoverTab[70530]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:172
			// _ = "end of CoverTab[70530]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:172
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:172
		// _ = "end of CoverTab[70528]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:172
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:172
	// _ = "end of CoverTab[70525]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:172
	_go_fuzz_dep_.CoverTab[70526]++
											i.p = next
											return i.buf[:decomposeHangul(i.buf[:], i.rb.src.hangul(p))]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:174
	// _ = "end of CoverTab[70526]"
}

func nextDone(i *Iter) []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:177
	_go_fuzz_dep_.CoverTab[70531]++
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:178
	// _ = "end of CoverTab[70531]"
}

// nextMulti is used for iterating over multi-segment decompositions
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:181
// for decomposing normal forms.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:183
func nextMulti(i *Iter) []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:183
	_go_fuzz_dep_.CoverTab[70532]++
											j := 0
											d := i.multiSeg

											for j = 1; j < len(d) && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:187
		_go_fuzz_dep_.CoverTab[70535]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:187
		return !utf8.RuneStart(d[j])
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:187
		// _ = "end of CoverTab[70535]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:187
	}(); j++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:187
		_go_fuzz_dep_.CoverTab[70536]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:187
		// _ = "end of CoverTab[70536]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:188
	// _ = "end of CoverTab[70532]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:188
	_go_fuzz_dep_.CoverTab[70533]++
											for j < len(d) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:189
		_go_fuzz_dep_.CoverTab[70537]++
												info := i.rb.f.info(input{bytes: d}, j)
												if info.BoundaryBefore() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:191
			_go_fuzz_dep_.CoverTab[70539]++
													i.multiSeg = d[j:]
													return d[:j]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:193
			// _ = "end of CoverTab[70539]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:194
			_go_fuzz_dep_.CoverTab[70540]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:194
			// _ = "end of CoverTab[70540]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:194
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:194
		// _ = "end of CoverTab[70537]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:194
		_go_fuzz_dep_.CoverTab[70538]++
												j += int(info.size)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:195
		// _ = "end of CoverTab[70538]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:196
	// _ = "end of CoverTab[70533]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:196
	_go_fuzz_dep_.CoverTab[70534]++

											i.next = i.rb.f.nextMain
											return i.next(i)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:199
	// _ = "end of CoverTab[70534]"
}

// nextMultiNorm is used for iterating over multi-segment decompositions
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:202
// for composing normal forms.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:204
func nextMultiNorm(i *Iter) []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:204
	_go_fuzz_dep_.CoverTab[70541]++
											j := 0
											d := i.multiSeg
											for j < len(d) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:207
		_go_fuzz_dep_.CoverTab[70543]++
												info := i.rb.f.info(input{bytes: d}, j)
												if info.BoundaryBefore() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:209
			_go_fuzz_dep_.CoverTab[70545]++
													i.rb.compose()
													seg := i.buf[:i.rb.flushCopy(i.buf[:])]
													i.rb.insertUnsafe(input{bytes: d}, j, info)
													i.multiSeg = d[j+int(info.size):]
													return seg
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:214
			// _ = "end of CoverTab[70545]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:215
			_go_fuzz_dep_.CoverTab[70546]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:215
			// _ = "end of CoverTab[70546]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:215
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:215
		// _ = "end of CoverTab[70543]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:215
		_go_fuzz_dep_.CoverTab[70544]++
												i.rb.insertUnsafe(input{bytes: d}, j, info)
												j += int(info.size)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:217
		// _ = "end of CoverTab[70544]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:218
	// _ = "end of CoverTab[70541]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:218
	_go_fuzz_dep_.CoverTab[70542]++
											i.multiSeg = nil
											i.next = nextComposed
											return doNormComposed(i)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:221
	// _ = "end of CoverTab[70542]"
}

// nextDecomposed is the implementation of Next for forms NFD and NFKD.
func nextDecomposed(i *Iter) (next []byte) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:225
	_go_fuzz_dep_.CoverTab[70547]++
											outp := 0
											inCopyStart, outCopyStart := i.p, 0
											for {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:228
		_go_fuzz_dep_.CoverTab[70550]++
												if sz := int(i.info.size); sz <= 1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:229
			_go_fuzz_dep_.CoverTab[70554]++
													i.rb.ss = 0
													p := i.p
													i.p++
													if i.p >= i.rb.nsrc {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:233
				_go_fuzz_dep_.CoverTab[70556]++
														i.setDone()
														return i.returnSlice(p, i.p)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:235
				// _ = "end of CoverTab[70556]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:236
				_go_fuzz_dep_.CoverTab[70557]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:236
				if i.rb.src._byte(i.p) < utf8.RuneSelf {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:236
					_go_fuzz_dep_.CoverTab[70558]++
															i.next = i.asciiF
															return i.returnSlice(p, i.p)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:238
					// _ = "end of CoverTab[70558]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:239
					_go_fuzz_dep_.CoverTab[70559]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:239
					// _ = "end of CoverTab[70559]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:239
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:239
				// _ = "end of CoverTab[70557]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:239
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:239
			// _ = "end of CoverTab[70554]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:239
			_go_fuzz_dep_.CoverTab[70555]++
													outp++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:240
			// _ = "end of CoverTab[70555]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:241
			_go_fuzz_dep_.CoverTab[70560]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:241
			if d := i.info.Decomposition(); d != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:241
				_go_fuzz_dep_.CoverTab[70561]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:246
				p := outp + len(d)
				if outp > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:247
					_go_fuzz_dep_.CoverTab[70566]++
															i.rb.src.copySlice(i.buf[outCopyStart:], inCopyStart, i.p)

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:251
					if p > len(i.buf) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:251
						_go_fuzz_dep_.CoverTab[70567]++
																return i.buf[:outp]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:252
						// _ = "end of CoverTab[70567]"
					} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:253
						_go_fuzz_dep_.CoverTab[70568]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:253
						// _ = "end of CoverTab[70568]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:253
					}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:253
					// _ = "end of CoverTab[70566]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:254
					_go_fuzz_dep_.CoverTab[70569]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:254
					if i.info.multiSegment() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:254
						_go_fuzz_dep_.CoverTab[70570]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:257
						if i.multiSeg == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:257
							_go_fuzz_dep_.CoverTab[70572]++
																	i.multiSeg = d
																	i.next = nextMulti
																	return nextMulti(i)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:260
							// _ = "end of CoverTab[70572]"
						} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:261
							_go_fuzz_dep_.CoverTab[70573]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:261
							// _ = "end of CoverTab[70573]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:261
						}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:261
						// _ = "end of CoverTab[70570]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:261
						_go_fuzz_dep_.CoverTab[70571]++

																d = i.multiSeg
																i.multiSeg = nil
																p = len(d)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:265
						// _ = "end of CoverTab[70571]"
					} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:266
						_go_fuzz_dep_.CoverTab[70574]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:266
						// _ = "end of CoverTab[70574]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:266
					}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:266
					// _ = "end of CoverTab[70569]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:266
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:266
				// _ = "end of CoverTab[70561]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:266
				_go_fuzz_dep_.CoverTab[70562]++
														prevCC := i.info.tccc
														if i.p += sz; i.p >= i.rb.nsrc {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:268
					_go_fuzz_dep_.CoverTab[70575]++
															i.setDone()
															i.info = Properties{}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:270
					// _ = "end of CoverTab[70575]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:271
					_go_fuzz_dep_.CoverTab[70576]++
															i.info = i.rb.f.info(i.rb.src, i.p)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:272
					// _ = "end of CoverTab[70576]"
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:273
				// _ = "end of CoverTab[70562]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:273
				_go_fuzz_dep_.CoverTab[70563]++
														switch i.rb.ss.next(i.info) {
				case ssOverflow:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:275
					_go_fuzz_dep_.CoverTab[70577]++
															i.next = nextCGJDecompose
															fallthrough
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:277
					// _ = "end of CoverTab[70577]"
				case ssStarter:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:278
					_go_fuzz_dep_.CoverTab[70578]++
															if outp > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:279
						_go_fuzz_dep_.CoverTab[70581]++
																copy(i.buf[outp:], d)
																return i.buf[:p]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:281
						// _ = "end of CoverTab[70581]"
					} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:282
						_go_fuzz_dep_.CoverTab[70582]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:282
						// _ = "end of CoverTab[70582]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:282
					}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:282
					// _ = "end of CoverTab[70578]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:282
					_go_fuzz_dep_.CoverTab[70579]++
															return d
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:283
					// _ = "end of CoverTab[70579]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:283
				default:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:283
					_go_fuzz_dep_.CoverTab[70580]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:283
					// _ = "end of CoverTab[70580]"
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:284
				// _ = "end of CoverTab[70563]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:284
				_go_fuzz_dep_.CoverTab[70564]++
														copy(i.buf[outp:], d)
														outp = p
														inCopyStart, outCopyStart = i.p, outp
														if i.info.ccc < prevCC {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:288
					_go_fuzz_dep_.CoverTab[70583]++
															goto doNorm
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:289
					// _ = "end of CoverTab[70583]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:290
					_go_fuzz_dep_.CoverTab[70584]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:290
					// _ = "end of CoverTab[70584]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:290
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:290
				// _ = "end of CoverTab[70564]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:290
				_go_fuzz_dep_.CoverTab[70565]++
														continue
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:291
				// _ = "end of CoverTab[70565]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:292
				_go_fuzz_dep_.CoverTab[70585]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:292
				if r := i.rb.src.hangul(i.p); r != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:292
					_go_fuzz_dep_.CoverTab[70586]++
															outp = decomposeHangul(i.buf[:], r)
															i.p += hangulUTF8Size
															inCopyStart, outCopyStart = i.p, outp
															if i.p >= i.rb.nsrc {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:296
						_go_fuzz_dep_.CoverTab[70587]++
																i.setDone()
																break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:298
						// _ = "end of CoverTab[70587]"
					} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:299
						_go_fuzz_dep_.CoverTab[70588]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:299
						if i.rb.src.hangul(i.p) != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:299
							_go_fuzz_dep_.CoverTab[70589]++
																	i.next = nextHangul
																	return i.buf[:outp]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:301
							// _ = "end of CoverTab[70589]"
						} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:302
							_go_fuzz_dep_.CoverTab[70590]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:302
							// _ = "end of CoverTab[70590]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:302
						}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:302
						// _ = "end of CoverTab[70588]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:302
					}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:302
					// _ = "end of CoverTab[70586]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:303
					_go_fuzz_dep_.CoverTab[70591]++
															p := outp + sz
															if p > len(i.buf) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:305
						_go_fuzz_dep_.CoverTab[70593]++
																break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:306
						// _ = "end of CoverTab[70593]"
					} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:307
						_go_fuzz_dep_.CoverTab[70594]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:307
						// _ = "end of CoverTab[70594]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:307
					}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:307
					// _ = "end of CoverTab[70591]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:307
					_go_fuzz_dep_.CoverTab[70592]++
															outp = p
															i.p += sz
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:309
					// _ = "end of CoverTab[70592]"
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:310
				// _ = "end of CoverTab[70585]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:310
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:310
			// _ = "end of CoverTab[70560]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:310
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:310
		// _ = "end of CoverTab[70550]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:310
		_go_fuzz_dep_.CoverTab[70551]++
												if i.p >= i.rb.nsrc {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:311
			_go_fuzz_dep_.CoverTab[70595]++
													i.setDone()
													break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:313
			// _ = "end of CoverTab[70595]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:314
			_go_fuzz_dep_.CoverTab[70596]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:314
			// _ = "end of CoverTab[70596]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:314
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:314
		// _ = "end of CoverTab[70551]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:314
		_go_fuzz_dep_.CoverTab[70552]++
												prevCC := i.info.tccc
												i.info = i.rb.f.info(i.rb.src, i.p)
												if v := i.rb.ss.next(i.info); v == ssStarter {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:317
			_go_fuzz_dep_.CoverTab[70597]++
													break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:318
			// _ = "end of CoverTab[70597]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:319
			_go_fuzz_dep_.CoverTab[70598]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:319
			if v == ssOverflow {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:319
				_go_fuzz_dep_.CoverTab[70599]++
														i.next = nextCGJDecompose
														break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:321
				// _ = "end of CoverTab[70599]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:322
				_go_fuzz_dep_.CoverTab[70600]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:322
				// _ = "end of CoverTab[70600]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:322
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:322
			// _ = "end of CoverTab[70598]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:322
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:322
		// _ = "end of CoverTab[70552]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:322
		_go_fuzz_dep_.CoverTab[70553]++
												if i.info.ccc < prevCC {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:323
			_go_fuzz_dep_.CoverTab[70601]++
													goto doNorm
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:324
			// _ = "end of CoverTab[70601]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:325
			_go_fuzz_dep_.CoverTab[70602]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:325
			// _ = "end of CoverTab[70602]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:325
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:325
		// _ = "end of CoverTab[70553]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:326
	// _ = "end of CoverTab[70547]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:326
	_go_fuzz_dep_.CoverTab[70548]++
											if outCopyStart == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:327
		_go_fuzz_dep_.CoverTab[70603]++
												return i.returnSlice(inCopyStart, i.p)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:328
		// _ = "end of CoverTab[70603]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:329
		_go_fuzz_dep_.CoverTab[70604]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:329
		if inCopyStart < i.p {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:329
			_go_fuzz_dep_.CoverTab[70605]++
													i.rb.src.copySlice(i.buf[outCopyStart:], inCopyStart, i.p)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:330
			// _ = "end of CoverTab[70605]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:331
			_go_fuzz_dep_.CoverTab[70606]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:331
			// _ = "end of CoverTab[70606]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:331
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:331
		// _ = "end of CoverTab[70604]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:331
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:331
	// _ = "end of CoverTab[70548]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:331
	_go_fuzz_dep_.CoverTab[70549]++
											return i.buf[:outp]
doNorm:

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:336
	i.rb.src.copySlice(i.buf[outCopyStart:], inCopyStart, i.p)
											i.rb.insertDecomposed(i.buf[0:outp])
											return doNormDecomposed(i)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:338
	// _ = "end of CoverTab[70549]"
}

func doNormDecomposed(i *Iter) []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:341
	_go_fuzz_dep_.CoverTab[70607]++
											for {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:342
		_go_fuzz_dep_.CoverTab[70609]++
												i.rb.insertUnsafe(i.rb.src, i.p, i.info)
												if i.p += int(i.info.size); i.p >= i.rb.nsrc {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:344
			_go_fuzz_dep_.CoverTab[70612]++
													i.setDone()
													break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:346
			// _ = "end of CoverTab[70612]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:347
			_go_fuzz_dep_.CoverTab[70613]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:347
			// _ = "end of CoverTab[70613]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:347
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:347
		// _ = "end of CoverTab[70609]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:347
		_go_fuzz_dep_.CoverTab[70610]++
												i.info = i.rb.f.info(i.rb.src, i.p)
												if i.info.ccc == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:349
			_go_fuzz_dep_.CoverTab[70614]++
													break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:350
			// _ = "end of CoverTab[70614]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:351
			_go_fuzz_dep_.CoverTab[70615]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:351
			// _ = "end of CoverTab[70615]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:351
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:351
		// _ = "end of CoverTab[70610]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:351
		_go_fuzz_dep_.CoverTab[70611]++
												if s := i.rb.ss.next(i.info); s == ssOverflow {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:352
			_go_fuzz_dep_.CoverTab[70616]++
													i.next = nextCGJDecompose
													break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:354
			// _ = "end of CoverTab[70616]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:355
			_go_fuzz_dep_.CoverTab[70617]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:355
			// _ = "end of CoverTab[70617]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:355
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:355
		// _ = "end of CoverTab[70611]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:356
	// _ = "end of CoverTab[70607]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:356
	_go_fuzz_dep_.CoverTab[70608]++

											return i.buf[:i.rb.flushCopy(i.buf[:])]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:358
	// _ = "end of CoverTab[70608]"
}

func nextCGJDecompose(i *Iter) []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:361
	_go_fuzz_dep_.CoverTab[70618]++
											i.rb.ss = 0
											i.rb.insertCGJ()
											i.next = nextDecomposed
											i.rb.ss.first(i.info)
											buf := doNormDecomposed(i)
											return buf
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:367
	// _ = "end of CoverTab[70618]"
}

// nextComposed is the implementation of Next for forms NFC and NFKC.
func nextComposed(i *Iter) []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:371
	_go_fuzz_dep_.CoverTab[70619]++
											outp, startp := 0, i.p
											var prevCC uint8
											for {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:374
		_go_fuzz_dep_.CoverTab[70622]++
												if !i.info.isYesC() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:375
			_go_fuzz_dep_.CoverTab[70628]++
													goto doNorm
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:376
			// _ = "end of CoverTab[70628]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:377
			_go_fuzz_dep_.CoverTab[70629]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:377
			// _ = "end of CoverTab[70629]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:377
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:377
		// _ = "end of CoverTab[70622]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:377
		_go_fuzz_dep_.CoverTab[70623]++
												prevCC = i.info.tccc
												sz := int(i.info.size)
												if sz == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:380
			_go_fuzz_dep_.CoverTab[70630]++
													sz = 1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:381
			// _ = "end of CoverTab[70630]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:382
			_go_fuzz_dep_.CoverTab[70631]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:382
			// _ = "end of CoverTab[70631]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:382
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:382
		// _ = "end of CoverTab[70623]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:382
		_go_fuzz_dep_.CoverTab[70624]++
												p := outp + sz
												if p > len(i.buf) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:384
			_go_fuzz_dep_.CoverTab[70632]++
													break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:385
			// _ = "end of CoverTab[70632]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:386
			_go_fuzz_dep_.CoverTab[70633]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:386
			// _ = "end of CoverTab[70633]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:386
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:386
		// _ = "end of CoverTab[70624]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:386
		_go_fuzz_dep_.CoverTab[70625]++
												outp = p
												i.p += sz
												if i.p >= i.rb.nsrc {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:389
			_go_fuzz_dep_.CoverTab[70634]++
													i.setDone()
													break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:391
			// _ = "end of CoverTab[70634]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:392
			_go_fuzz_dep_.CoverTab[70635]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:392
			if i.rb.src._byte(i.p) < utf8.RuneSelf {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:392
				_go_fuzz_dep_.CoverTab[70636]++
														i.rb.ss = 0
														i.next = i.asciiF
														break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:395
				// _ = "end of CoverTab[70636]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:396
				_go_fuzz_dep_.CoverTab[70637]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:396
				// _ = "end of CoverTab[70637]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:396
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:396
			// _ = "end of CoverTab[70635]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:396
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:396
		// _ = "end of CoverTab[70625]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:396
		_go_fuzz_dep_.CoverTab[70626]++
												i.info = i.rb.f.info(i.rb.src, i.p)
												if v := i.rb.ss.next(i.info); v == ssStarter {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:398
			_go_fuzz_dep_.CoverTab[70638]++
													break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:399
			// _ = "end of CoverTab[70638]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:400
			_go_fuzz_dep_.CoverTab[70639]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:400
			if v == ssOverflow {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:400
				_go_fuzz_dep_.CoverTab[70640]++
														i.next = nextCGJCompose
														break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:402
				// _ = "end of CoverTab[70640]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:403
				_go_fuzz_dep_.CoverTab[70641]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:403
				// _ = "end of CoverTab[70641]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:403
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:403
			// _ = "end of CoverTab[70639]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:403
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:403
		// _ = "end of CoverTab[70626]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:403
		_go_fuzz_dep_.CoverTab[70627]++
												if i.info.ccc < prevCC {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:404
			_go_fuzz_dep_.CoverTab[70642]++
													goto doNorm
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:405
			// _ = "end of CoverTab[70642]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:406
			_go_fuzz_dep_.CoverTab[70643]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:406
			// _ = "end of CoverTab[70643]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:406
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:406
		// _ = "end of CoverTab[70627]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:407
	// _ = "end of CoverTab[70619]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:407
	_go_fuzz_dep_.CoverTab[70620]++
											return i.returnSlice(startp, i.p)
doNorm:

	i.p = startp
	i.info = i.rb.f.info(i.rb.src, i.p)
	i.rb.ss.first(i.info)
	if i.info.multiSegment() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:414
		_go_fuzz_dep_.CoverTab[70644]++
												d := i.info.Decomposition()
												info := i.rb.f.info(input{bytes: d}, 0)
												i.rb.insertUnsafe(input{bytes: d}, 0, info)
												i.multiSeg = d[int(info.size):]
												i.next = nextMultiNorm
												return nextMultiNorm(i)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:420
		// _ = "end of CoverTab[70644]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:421
		_go_fuzz_dep_.CoverTab[70645]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:421
		// _ = "end of CoverTab[70645]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:421
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:421
	// _ = "end of CoverTab[70620]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:421
	_go_fuzz_dep_.CoverTab[70621]++
											i.rb.ss.first(i.info)
											i.rb.insertUnsafe(i.rb.src, i.p, i.info)
											return doNormComposed(i)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:424
	// _ = "end of CoverTab[70621]"
}

func doNormComposed(i *Iter) []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:427
	_go_fuzz_dep_.CoverTab[70646]++

											for {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:429
		_go_fuzz_dep_.CoverTab[70648]++
												if i.p += int(i.info.size); i.p >= i.rb.nsrc {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:430
			_go_fuzz_dep_.CoverTab[70651]++
													i.setDone()
													break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:432
			// _ = "end of CoverTab[70651]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:433
			_go_fuzz_dep_.CoverTab[70652]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:433
			// _ = "end of CoverTab[70652]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:433
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:433
		// _ = "end of CoverTab[70648]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:433
		_go_fuzz_dep_.CoverTab[70649]++
												i.info = i.rb.f.info(i.rb.src, i.p)
												if s := i.rb.ss.next(i.info); s == ssStarter {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:435
			_go_fuzz_dep_.CoverTab[70653]++
													break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:436
			// _ = "end of CoverTab[70653]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:437
			_go_fuzz_dep_.CoverTab[70654]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:437
			if s == ssOverflow {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:437
				_go_fuzz_dep_.CoverTab[70655]++
														i.next = nextCGJCompose
														break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:439
				// _ = "end of CoverTab[70655]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:440
				_go_fuzz_dep_.CoverTab[70656]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:440
				// _ = "end of CoverTab[70656]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:440
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:440
			// _ = "end of CoverTab[70654]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:440
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:440
		// _ = "end of CoverTab[70649]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:440
		_go_fuzz_dep_.CoverTab[70650]++
												i.rb.insertUnsafe(i.rb.src, i.p, i.info)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:441
		// _ = "end of CoverTab[70650]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:442
	// _ = "end of CoverTab[70646]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:442
	_go_fuzz_dep_.CoverTab[70647]++
											i.rb.compose()
											seg := i.buf[:i.rb.flushCopy(i.buf[:])]
											return seg
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:445
	// _ = "end of CoverTab[70647]"
}

func nextCGJCompose(i *Iter) []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:448
	_go_fuzz_dep_.CoverTab[70657]++
											i.rb.ss = 0
											i.rb.insertCGJ()
											i.next = nextComposed

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:455
	i.rb.ss.first(i.info)
											i.rb.insertUnsafe(i.rb.src, i.p, i.info)
											return doNormComposed(i)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:457
	// _ = "end of CoverTab[70657]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:458
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/iter.go:458
var _ = _go_fuzz_dep_.CoverTab
