// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:5
package norm

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:5
import (
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:5
)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:5
import (
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:5
)

import (
	"fmt"
	"unicode/utf8"
)

// MaxSegmentSize is the maximum size of a byte buffer needed to consider any
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:12
// sequence of starter and non-starter runes for the purpose of normalization.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:14
const MaxSegmentSize = maxByteBufferSize

// An Iter iterates over a string or byte slice, while normalizing it
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:16
// to a given Form.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:18
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
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:32
	_go_fuzz_dep_.CoverTab[33202]++
										i.p = 0
										if len(src) == 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:34
		_go_fuzz_dep_.CoverTab[33204]++
											i.setDone()
											i.rb.nsrc = 0
											return
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:37
		// _ = "end of CoverTab[33204]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:38
		_go_fuzz_dep_.CoverTab[33205]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:38
		// _ = "end of CoverTab[33205]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:38
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:38
	// _ = "end of CoverTab[33202]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:38
	_go_fuzz_dep_.CoverTab[33203]++
										i.multiSeg = nil
										i.rb.init(f, src)
										i.next = i.rb.f.nextMain
										i.asciiF = nextASCIIBytes
										i.info = i.rb.f.info(i.rb.src, i.p)
										i.rb.ss.first(i.info)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:44
	// _ = "end of CoverTab[33203]"
}

// InitString initializes i to iterate over src after normalizing it to Form f.
func (i *Iter) InitString(f Form, src string) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:48
	_go_fuzz_dep_.CoverTab[33206]++
										i.p = 0
										if len(src) == 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:50
		_go_fuzz_dep_.CoverTab[33208]++
											i.setDone()
											i.rb.nsrc = 0
											return
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:53
		// _ = "end of CoverTab[33208]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:54
		_go_fuzz_dep_.CoverTab[33209]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:54
		// _ = "end of CoverTab[33209]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:54
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:54
	// _ = "end of CoverTab[33206]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:54
	_go_fuzz_dep_.CoverTab[33207]++
										i.multiSeg = nil
										i.rb.initString(f, src)
										i.next = i.rb.f.nextMain
										i.asciiF = nextASCIIString
										i.info = i.rb.f.info(i.rb.src, i.p)
										i.rb.ss.first(i.info)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:60
	// _ = "end of CoverTab[33207]"
}

// Seek sets the segment to be returned by the next call to Next to start
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:63
// at position p.  It is the responsibility of the caller to set p to the
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:63
// start of a segment.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:66
func (i *Iter) Seek(offset int64, whence int) (int64, error) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:66
	_go_fuzz_dep_.CoverTab[33210]++
										var abs int64
										switch whence {
	case 0:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:69
		_go_fuzz_dep_.CoverTab[33214]++
											abs = offset
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:70
		// _ = "end of CoverTab[33214]"
	case 1:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:71
		_go_fuzz_dep_.CoverTab[33215]++
											abs = int64(i.p) + offset
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:72
		// _ = "end of CoverTab[33215]"
	case 2:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:73
		_go_fuzz_dep_.CoverTab[33216]++
											abs = int64(i.rb.nsrc) + offset
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:74
		// _ = "end of CoverTab[33216]"
	default:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:75
		_go_fuzz_dep_.CoverTab[33217]++
											return 0, fmt.Errorf("norm: invalid whence")
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:76
		// _ = "end of CoverTab[33217]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:77
	// _ = "end of CoverTab[33210]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:77
	_go_fuzz_dep_.CoverTab[33211]++
										if abs < 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:78
		_go_fuzz_dep_.CoverTab[33218]++
											return 0, fmt.Errorf("norm: negative position")
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:79
		// _ = "end of CoverTab[33218]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:80
		_go_fuzz_dep_.CoverTab[33219]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:80
		// _ = "end of CoverTab[33219]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:80
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:80
	// _ = "end of CoverTab[33211]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:80
	_go_fuzz_dep_.CoverTab[33212]++
										if int(abs) >= i.rb.nsrc {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:81
		_go_fuzz_dep_.CoverTab[33220]++
											i.setDone()
											return int64(i.p), nil
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:83
		// _ = "end of CoverTab[33220]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:84
		_go_fuzz_dep_.CoverTab[33221]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:84
		// _ = "end of CoverTab[33221]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:84
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:84
	// _ = "end of CoverTab[33212]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:84
	_go_fuzz_dep_.CoverTab[33213]++
										i.p = int(abs)
										i.multiSeg = nil
										i.next = i.rb.f.nextMain
										i.info = i.rb.f.info(i.rb.src, i.p)
										i.rb.ss.first(i.info)
										return abs, nil
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:90
	// _ = "end of CoverTab[33213]"
}

// returnSlice returns a slice of the underlying input type as a byte slice.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:93
// If the underlying is of type []byte, it will simply return a slice.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:93
// If the underlying is of type string, it will copy the slice to the buffer
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:93
// and return that.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:97
func (i *Iter) returnSlice(a, b int) []byte {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:97
	_go_fuzz_dep_.CoverTab[33222]++
										if i.rb.src.bytes == nil {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:98
		_go_fuzz_dep_.CoverTab[33224]++
											return i.buf[:copy(i.buf[:], i.rb.src.str[a:b])]
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:99
		// _ = "end of CoverTab[33224]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:100
		_go_fuzz_dep_.CoverTab[33225]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:100
		// _ = "end of CoverTab[33225]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:100
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:100
	// _ = "end of CoverTab[33222]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:100
	_go_fuzz_dep_.CoverTab[33223]++
										return i.rb.src.bytes[a:b]
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:101
	// _ = "end of CoverTab[33223]"
}

// Pos returns the byte position at which the next call to Next will commence processing.
func (i *Iter) Pos() int {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:105
	_go_fuzz_dep_.CoverTab[33226]++
										return i.p
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:106
	// _ = "end of CoverTab[33226]"
}

func (i *Iter) setDone() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:109
	_go_fuzz_dep_.CoverTab[33227]++
										i.next = nextDone
										i.p = i.rb.nsrc
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:111
	// _ = "end of CoverTab[33227]"
}

// Done returns true if there is no more input to process.
func (i *Iter) Done() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:115
	_go_fuzz_dep_.CoverTab[33228]++
										return i.p >= i.rb.nsrc
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:116
	// _ = "end of CoverTab[33228]"
}

// Next returns f(i.input[i.Pos():n]), where n is a boundary of i.input.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:119
// For any input a and b for which f(a) == f(b), subsequent calls
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:119
// to Next will return the same segments.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:119
// Modifying runes are grouped together with the preceding starter, if such a starter exists.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:119
// Although not guaranteed, n will typically be the smallest possible n.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:124
func (i *Iter) Next() []byte {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:124
	_go_fuzz_dep_.CoverTab[33229]++
										return i.next(i)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:125
	// _ = "end of CoverTab[33229]"
}

func nextASCIIBytes(i *Iter) []byte {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:128
	_go_fuzz_dep_.CoverTab[33230]++
										p := i.p + 1
										if p >= i.rb.nsrc {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:130
		_go_fuzz_dep_.CoverTab[33233]++
											p0 := i.p
											i.setDone()
											return i.rb.src.bytes[p0:p]
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:133
		// _ = "end of CoverTab[33233]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:134
		_go_fuzz_dep_.CoverTab[33234]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:134
		// _ = "end of CoverTab[33234]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:134
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:134
	// _ = "end of CoverTab[33230]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:134
	_go_fuzz_dep_.CoverTab[33231]++
										if i.rb.src.bytes[p] < utf8.RuneSelf {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:135
		_go_fuzz_dep_.CoverTab[33235]++
											p0 := i.p
											i.p = p
											return i.rb.src.bytes[p0:p]
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:138
		// _ = "end of CoverTab[33235]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:139
		_go_fuzz_dep_.CoverTab[33236]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:139
		// _ = "end of CoverTab[33236]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:139
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:139
	// _ = "end of CoverTab[33231]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:139
	_go_fuzz_dep_.CoverTab[33232]++
										i.info = i.rb.f.info(i.rb.src, i.p)
										i.next = i.rb.f.nextMain
										return i.next(i)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:142
	// _ = "end of CoverTab[33232]"
}

func nextASCIIString(i *Iter) []byte {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:145
	_go_fuzz_dep_.CoverTab[33237]++
										p := i.p + 1
										if p >= i.rb.nsrc {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:147
		_go_fuzz_dep_.CoverTab[33240]++
											i.buf[0] = i.rb.src.str[i.p]
											i.setDone()
											return i.buf[:1]
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:150
		// _ = "end of CoverTab[33240]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:151
		_go_fuzz_dep_.CoverTab[33241]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:151
		// _ = "end of CoverTab[33241]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:151
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:151
	// _ = "end of CoverTab[33237]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:151
	_go_fuzz_dep_.CoverTab[33238]++
										if i.rb.src.str[p] < utf8.RuneSelf {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:152
		_go_fuzz_dep_.CoverTab[33242]++
											i.buf[0] = i.rb.src.str[i.p]
											i.p = p
											return i.buf[:1]
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:155
		// _ = "end of CoverTab[33242]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:156
		_go_fuzz_dep_.CoverTab[33243]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:156
		// _ = "end of CoverTab[33243]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:156
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:156
	// _ = "end of CoverTab[33238]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:156
	_go_fuzz_dep_.CoverTab[33239]++
										i.info = i.rb.f.info(i.rb.src, i.p)
										i.next = i.rb.f.nextMain
										return i.next(i)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:159
	// _ = "end of CoverTab[33239]"
}

func nextHangul(i *Iter) []byte {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:162
	_go_fuzz_dep_.CoverTab[33244]++
										p := i.p
										next := p + hangulUTF8Size
										if next >= i.rb.nsrc {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:165
		_go_fuzz_dep_.CoverTab[33246]++
											i.setDone()
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:166
		// _ = "end of CoverTab[33246]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:167
		_go_fuzz_dep_.CoverTab[33247]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:167
		if i.rb.src.hangul(next) == 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:167
			_go_fuzz_dep_.CoverTab[33248]++
												i.rb.ss.next(i.info)
												i.info = i.rb.f.info(i.rb.src, i.p)
												i.next = i.rb.f.nextMain
												return i.next(i)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:171
			// _ = "end of CoverTab[33248]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:172
			_go_fuzz_dep_.CoverTab[33249]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:172
			// _ = "end of CoverTab[33249]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:172
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:172
		// _ = "end of CoverTab[33247]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:172
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:172
	// _ = "end of CoverTab[33244]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:172
	_go_fuzz_dep_.CoverTab[33245]++
										i.p = next
										return i.buf[:decomposeHangul(i.buf[:], i.rb.src.hangul(p))]
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:174
	// _ = "end of CoverTab[33245]"
}

func nextDone(i *Iter) []byte {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:177
	_go_fuzz_dep_.CoverTab[33250]++
										return nil
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:178
	// _ = "end of CoverTab[33250]"
}

// nextMulti is used for iterating over multi-segment decompositions
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:181
// for decomposing normal forms.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:183
func nextMulti(i *Iter) []byte {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:183
	_go_fuzz_dep_.CoverTab[33251]++
										j := 0
										d := i.multiSeg

										for j = 1; j < len(d) && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:187
		_go_fuzz_dep_.CoverTab[33254]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:187
		return !utf8.RuneStart(d[j])
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:187
		// _ = "end of CoverTab[33254]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:187
	}(); j++ {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:187
		_go_fuzz_dep_.CoverTab[33255]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:187
		// _ = "end of CoverTab[33255]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:188
	// _ = "end of CoverTab[33251]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:188
	_go_fuzz_dep_.CoverTab[33252]++
										for j < len(d) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:189
		_go_fuzz_dep_.CoverTab[33256]++
											info := i.rb.f.info(input{bytes: d}, j)
											if info.BoundaryBefore() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:191
			_go_fuzz_dep_.CoverTab[33258]++
												i.multiSeg = d[j:]
												return d[:j]
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:193
			// _ = "end of CoverTab[33258]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:194
			_go_fuzz_dep_.CoverTab[33259]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:194
			// _ = "end of CoverTab[33259]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:194
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:194
		// _ = "end of CoverTab[33256]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:194
		_go_fuzz_dep_.CoverTab[33257]++
											j += int(info.size)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:195
		// _ = "end of CoverTab[33257]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:196
	// _ = "end of CoverTab[33252]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:196
	_go_fuzz_dep_.CoverTab[33253]++

										i.next = i.rb.f.nextMain
										return i.next(i)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:199
	// _ = "end of CoverTab[33253]"
}

// nextMultiNorm is used for iterating over multi-segment decompositions
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:202
// for composing normal forms.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:204
func nextMultiNorm(i *Iter) []byte {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:204
	_go_fuzz_dep_.CoverTab[33260]++
										j := 0
										d := i.multiSeg
										for j < len(d) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:207
		_go_fuzz_dep_.CoverTab[33262]++
											info := i.rb.f.info(input{bytes: d}, j)
											if info.BoundaryBefore() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:209
			_go_fuzz_dep_.CoverTab[33264]++
												i.rb.compose()
												seg := i.buf[:i.rb.flushCopy(i.buf[:])]
												i.rb.insertUnsafe(input{bytes: d}, j, info)
												i.multiSeg = d[j+int(info.size):]
												return seg
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:214
			// _ = "end of CoverTab[33264]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:215
			_go_fuzz_dep_.CoverTab[33265]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:215
			// _ = "end of CoverTab[33265]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:215
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:215
		// _ = "end of CoverTab[33262]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:215
		_go_fuzz_dep_.CoverTab[33263]++
											i.rb.insertUnsafe(input{bytes: d}, j, info)
											j += int(info.size)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:217
		// _ = "end of CoverTab[33263]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:218
	// _ = "end of CoverTab[33260]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:218
	_go_fuzz_dep_.CoverTab[33261]++
										i.multiSeg = nil
										i.next = nextComposed
										return doNormComposed(i)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:221
	// _ = "end of CoverTab[33261]"
}

// nextDecomposed is the implementation of Next for forms NFD and NFKD.
func nextDecomposed(i *Iter) (next []byte) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:225
	_go_fuzz_dep_.CoverTab[33266]++
										outp := 0
										inCopyStart, outCopyStart := i.p, 0
										for {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:228
		_go_fuzz_dep_.CoverTab[33269]++
											if sz := int(i.info.size); sz <= 1 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:229
			_go_fuzz_dep_.CoverTab[33273]++
												i.rb.ss = 0
												p := i.p
												i.p++
												if i.p >= i.rb.nsrc {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:233
				_go_fuzz_dep_.CoverTab[33275]++
													i.setDone()
													return i.returnSlice(p, i.p)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:235
				// _ = "end of CoverTab[33275]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:236
				_go_fuzz_dep_.CoverTab[33276]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:236
				if i.rb.src._byte(i.p) < utf8.RuneSelf {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:236
					_go_fuzz_dep_.CoverTab[33277]++
														i.next = i.asciiF
														return i.returnSlice(p, i.p)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:238
					// _ = "end of CoverTab[33277]"
				} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:239
					_go_fuzz_dep_.CoverTab[33278]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:239
					// _ = "end of CoverTab[33278]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:239
				}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:239
				// _ = "end of CoverTab[33276]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:239
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:239
			// _ = "end of CoverTab[33273]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:239
			_go_fuzz_dep_.CoverTab[33274]++
												outp++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:240
			// _ = "end of CoverTab[33274]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:241
			_go_fuzz_dep_.CoverTab[33279]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:241
			if d := i.info.Decomposition(); d != nil {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:241
				_go_fuzz_dep_.CoverTab[33280]++

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:246
				p := outp + len(d)
				if outp > 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:247
					_go_fuzz_dep_.CoverTab[33285]++
														i.rb.src.copySlice(i.buf[outCopyStart:], inCopyStart, i.p)

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:251
					if p > len(i.buf) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:251
						_go_fuzz_dep_.CoverTab[33286]++
															return i.buf[:outp]
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:252
						// _ = "end of CoverTab[33286]"
					} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:253
						_go_fuzz_dep_.CoverTab[33287]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:253
						// _ = "end of CoverTab[33287]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:253
					}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:253
					// _ = "end of CoverTab[33285]"
				} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:254
					_go_fuzz_dep_.CoverTab[33288]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:254
					if i.info.multiSegment() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:254
						_go_fuzz_dep_.CoverTab[33289]++

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:257
						if i.multiSeg == nil {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:257
							_go_fuzz_dep_.CoverTab[33291]++
																i.multiSeg = d
																i.next = nextMulti
																return nextMulti(i)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:260
							// _ = "end of CoverTab[33291]"
						} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:261
							_go_fuzz_dep_.CoverTab[33292]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:261
							// _ = "end of CoverTab[33292]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:261
						}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:261
						// _ = "end of CoverTab[33289]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:261
						_go_fuzz_dep_.CoverTab[33290]++

															d = i.multiSeg
															i.multiSeg = nil
															p = len(d)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:265
						// _ = "end of CoverTab[33290]"
					} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:266
						_go_fuzz_dep_.CoverTab[33293]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:266
						// _ = "end of CoverTab[33293]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:266
					}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:266
					// _ = "end of CoverTab[33288]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:266
				}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:266
				// _ = "end of CoverTab[33280]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:266
				_go_fuzz_dep_.CoverTab[33281]++
													prevCC := i.info.tccc
													if i.p += sz; i.p >= i.rb.nsrc {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:268
					_go_fuzz_dep_.CoverTab[33294]++
														i.setDone()
														i.info = Properties{}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:270
					// _ = "end of CoverTab[33294]"
				} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:271
					_go_fuzz_dep_.CoverTab[33295]++
														i.info = i.rb.f.info(i.rb.src, i.p)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:272
					// _ = "end of CoverTab[33295]"
				}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:273
				// _ = "end of CoverTab[33281]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:273
				_go_fuzz_dep_.CoverTab[33282]++
													switch i.rb.ss.next(i.info) {
				case ssOverflow:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:275
					_go_fuzz_dep_.CoverTab[33296]++
														i.next = nextCGJDecompose
														fallthrough
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:277
					// _ = "end of CoverTab[33296]"
				case ssStarter:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:278
					_go_fuzz_dep_.CoverTab[33297]++
														if outp > 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:279
						_go_fuzz_dep_.CoverTab[33300]++
															copy(i.buf[outp:], d)
															return i.buf[:p]
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:281
						// _ = "end of CoverTab[33300]"
					} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:282
						_go_fuzz_dep_.CoverTab[33301]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:282
						// _ = "end of CoverTab[33301]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:282
					}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:282
					// _ = "end of CoverTab[33297]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:282
					_go_fuzz_dep_.CoverTab[33298]++
														return d
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:283
					// _ = "end of CoverTab[33298]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:283
				default:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:283
					_go_fuzz_dep_.CoverTab[33299]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:283
					// _ = "end of CoverTab[33299]"
				}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:284
				// _ = "end of CoverTab[33282]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:284
				_go_fuzz_dep_.CoverTab[33283]++
													copy(i.buf[outp:], d)
													outp = p
													inCopyStart, outCopyStart = i.p, outp
													if i.info.ccc < prevCC {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:288
					_go_fuzz_dep_.CoverTab[33302]++
														goto doNorm
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:289
					// _ = "end of CoverTab[33302]"
				} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:290
					_go_fuzz_dep_.CoverTab[33303]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:290
					// _ = "end of CoverTab[33303]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:290
				}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:290
				// _ = "end of CoverTab[33283]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:290
				_go_fuzz_dep_.CoverTab[33284]++
													continue
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:291
				// _ = "end of CoverTab[33284]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:292
				_go_fuzz_dep_.CoverTab[33304]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:292
				if r := i.rb.src.hangul(i.p); r != 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:292
					_go_fuzz_dep_.CoverTab[33305]++
														outp = decomposeHangul(i.buf[:], r)
														i.p += hangulUTF8Size
														inCopyStart, outCopyStart = i.p, outp
														if i.p >= i.rb.nsrc {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:296
						_go_fuzz_dep_.CoverTab[33306]++
															i.setDone()
															break
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:298
						// _ = "end of CoverTab[33306]"
					} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:299
						_go_fuzz_dep_.CoverTab[33307]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:299
						if i.rb.src.hangul(i.p) != 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:299
							_go_fuzz_dep_.CoverTab[33308]++
																i.next = nextHangul
																return i.buf[:outp]
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:301
							// _ = "end of CoverTab[33308]"
						} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:302
							_go_fuzz_dep_.CoverTab[33309]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:302
							// _ = "end of CoverTab[33309]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:302
						}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:302
						// _ = "end of CoverTab[33307]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:302
					}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:302
					// _ = "end of CoverTab[33305]"
				} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:303
					_go_fuzz_dep_.CoverTab[33310]++
														p := outp + sz
														if p > len(i.buf) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:305
						_go_fuzz_dep_.CoverTab[33312]++
															break
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:306
						// _ = "end of CoverTab[33312]"
					} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:307
						_go_fuzz_dep_.CoverTab[33313]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:307
						// _ = "end of CoverTab[33313]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:307
					}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:307
					// _ = "end of CoverTab[33310]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:307
					_go_fuzz_dep_.CoverTab[33311]++
														outp = p
														i.p += sz
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:309
					// _ = "end of CoverTab[33311]"
				}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:310
				// _ = "end of CoverTab[33304]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:310
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:310
			// _ = "end of CoverTab[33279]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:310
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:310
		// _ = "end of CoverTab[33269]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:310
		_go_fuzz_dep_.CoverTab[33270]++
											if i.p >= i.rb.nsrc {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:311
			_go_fuzz_dep_.CoverTab[33314]++
												i.setDone()
												break
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:313
			// _ = "end of CoverTab[33314]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:314
			_go_fuzz_dep_.CoverTab[33315]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:314
			// _ = "end of CoverTab[33315]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:314
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:314
		// _ = "end of CoverTab[33270]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:314
		_go_fuzz_dep_.CoverTab[33271]++
											prevCC := i.info.tccc
											i.info = i.rb.f.info(i.rb.src, i.p)
											if v := i.rb.ss.next(i.info); v == ssStarter {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:317
			_go_fuzz_dep_.CoverTab[33316]++
												break
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:318
			// _ = "end of CoverTab[33316]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:319
			_go_fuzz_dep_.CoverTab[33317]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:319
			if v == ssOverflow {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:319
				_go_fuzz_dep_.CoverTab[33318]++
													i.next = nextCGJDecompose
													break
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:321
				// _ = "end of CoverTab[33318]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:322
				_go_fuzz_dep_.CoverTab[33319]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:322
				// _ = "end of CoverTab[33319]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:322
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:322
			// _ = "end of CoverTab[33317]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:322
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:322
		// _ = "end of CoverTab[33271]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:322
		_go_fuzz_dep_.CoverTab[33272]++
											if i.info.ccc < prevCC {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:323
			_go_fuzz_dep_.CoverTab[33320]++
												goto doNorm
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:324
			// _ = "end of CoverTab[33320]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:325
			_go_fuzz_dep_.CoverTab[33321]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:325
			// _ = "end of CoverTab[33321]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:325
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:325
		// _ = "end of CoverTab[33272]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:326
	// _ = "end of CoverTab[33266]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:326
	_go_fuzz_dep_.CoverTab[33267]++
										if outCopyStart == 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:327
		_go_fuzz_dep_.CoverTab[33322]++
											return i.returnSlice(inCopyStart, i.p)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:328
		// _ = "end of CoverTab[33322]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:329
		_go_fuzz_dep_.CoverTab[33323]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:329
		if inCopyStart < i.p {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:329
			_go_fuzz_dep_.CoverTab[33324]++
												i.rb.src.copySlice(i.buf[outCopyStart:], inCopyStart, i.p)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:330
			// _ = "end of CoverTab[33324]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:331
			_go_fuzz_dep_.CoverTab[33325]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:331
			// _ = "end of CoverTab[33325]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:331
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:331
		// _ = "end of CoverTab[33323]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:331
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:331
	// _ = "end of CoverTab[33267]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:331
	_go_fuzz_dep_.CoverTab[33268]++
										return i.buf[:outp]
doNorm:

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:336
	i.rb.src.copySlice(i.buf[outCopyStart:], inCopyStart, i.p)
										i.rb.insertDecomposed(i.buf[0:outp])
										return doNormDecomposed(i)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:338
	// _ = "end of CoverTab[33268]"
}

func doNormDecomposed(i *Iter) []byte {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:341
	_go_fuzz_dep_.CoverTab[33326]++
										for {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:342
		_go_fuzz_dep_.CoverTab[33328]++
											i.rb.insertUnsafe(i.rb.src, i.p, i.info)
											if i.p += int(i.info.size); i.p >= i.rb.nsrc {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:344
			_go_fuzz_dep_.CoverTab[33331]++
												i.setDone()
												break
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:346
			// _ = "end of CoverTab[33331]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:347
			_go_fuzz_dep_.CoverTab[33332]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:347
			// _ = "end of CoverTab[33332]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:347
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:347
		// _ = "end of CoverTab[33328]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:347
		_go_fuzz_dep_.CoverTab[33329]++
											i.info = i.rb.f.info(i.rb.src, i.p)
											if i.info.ccc == 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:349
			_go_fuzz_dep_.CoverTab[33333]++
												break
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:350
			// _ = "end of CoverTab[33333]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:351
			_go_fuzz_dep_.CoverTab[33334]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:351
			// _ = "end of CoverTab[33334]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:351
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:351
		// _ = "end of CoverTab[33329]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:351
		_go_fuzz_dep_.CoverTab[33330]++
											if s := i.rb.ss.next(i.info); s == ssOverflow {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:352
			_go_fuzz_dep_.CoverTab[33335]++
												i.next = nextCGJDecompose
												break
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:354
			// _ = "end of CoverTab[33335]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:355
			_go_fuzz_dep_.CoverTab[33336]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:355
			// _ = "end of CoverTab[33336]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:355
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:355
		// _ = "end of CoverTab[33330]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:356
	// _ = "end of CoverTab[33326]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:356
	_go_fuzz_dep_.CoverTab[33327]++

										return i.buf[:i.rb.flushCopy(i.buf[:])]
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:358
	// _ = "end of CoverTab[33327]"
}

func nextCGJDecompose(i *Iter) []byte {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:361
	_go_fuzz_dep_.CoverTab[33337]++
										i.rb.ss = 0
										i.rb.insertCGJ()
										i.next = nextDecomposed
										i.rb.ss.first(i.info)
										buf := doNormDecomposed(i)
										return buf
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:367
	// _ = "end of CoverTab[33337]"
}

// nextComposed is the implementation of Next for forms NFC and NFKC.
func nextComposed(i *Iter) []byte {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:371
	_go_fuzz_dep_.CoverTab[33338]++
										outp, startp := 0, i.p
										var prevCC uint8
										for {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:374
		_go_fuzz_dep_.CoverTab[33341]++
											if !i.info.isYesC() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:375
			_go_fuzz_dep_.CoverTab[33347]++
												goto doNorm
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:376
			// _ = "end of CoverTab[33347]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:377
			_go_fuzz_dep_.CoverTab[33348]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:377
			// _ = "end of CoverTab[33348]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:377
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:377
		// _ = "end of CoverTab[33341]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:377
		_go_fuzz_dep_.CoverTab[33342]++
											prevCC = i.info.tccc
											sz := int(i.info.size)
											if sz == 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:380
			_go_fuzz_dep_.CoverTab[33349]++
												sz = 1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:381
			// _ = "end of CoverTab[33349]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:382
			_go_fuzz_dep_.CoverTab[33350]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:382
			// _ = "end of CoverTab[33350]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:382
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:382
		// _ = "end of CoverTab[33342]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:382
		_go_fuzz_dep_.CoverTab[33343]++
											p := outp + sz
											if p > len(i.buf) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:384
			_go_fuzz_dep_.CoverTab[33351]++
												break
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:385
			// _ = "end of CoverTab[33351]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:386
			_go_fuzz_dep_.CoverTab[33352]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:386
			// _ = "end of CoverTab[33352]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:386
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:386
		// _ = "end of CoverTab[33343]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:386
		_go_fuzz_dep_.CoverTab[33344]++
											outp = p
											i.p += sz
											if i.p >= i.rb.nsrc {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:389
			_go_fuzz_dep_.CoverTab[33353]++
												i.setDone()
												break
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:391
			// _ = "end of CoverTab[33353]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:392
			_go_fuzz_dep_.CoverTab[33354]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:392
			if i.rb.src._byte(i.p) < utf8.RuneSelf {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:392
				_go_fuzz_dep_.CoverTab[33355]++
													i.rb.ss = 0
													i.next = i.asciiF
													break
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:395
				// _ = "end of CoverTab[33355]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:396
				_go_fuzz_dep_.CoverTab[33356]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:396
				// _ = "end of CoverTab[33356]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:396
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:396
			// _ = "end of CoverTab[33354]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:396
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:396
		// _ = "end of CoverTab[33344]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:396
		_go_fuzz_dep_.CoverTab[33345]++
											i.info = i.rb.f.info(i.rb.src, i.p)
											if v := i.rb.ss.next(i.info); v == ssStarter {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:398
			_go_fuzz_dep_.CoverTab[33357]++
												break
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:399
			// _ = "end of CoverTab[33357]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:400
			_go_fuzz_dep_.CoverTab[33358]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:400
			if v == ssOverflow {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:400
				_go_fuzz_dep_.CoverTab[33359]++
													i.next = nextCGJCompose
													break
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:402
				// _ = "end of CoverTab[33359]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:403
				_go_fuzz_dep_.CoverTab[33360]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:403
				// _ = "end of CoverTab[33360]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:403
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:403
			// _ = "end of CoverTab[33358]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:403
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:403
		// _ = "end of CoverTab[33345]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:403
		_go_fuzz_dep_.CoverTab[33346]++
											if i.info.ccc < prevCC {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:404
			_go_fuzz_dep_.CoverTab[33361]++
												goto doNorm
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:405
			// _ = "end of CoverTab[33361]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:406
			_go_fuzz_dep_.CoverTab[33362]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:406
			// _ = "end of CoverTab[33362]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:406
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:406
		// _ = "end of CoverTab[33346]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:407
	// _ = "end of CoverTab[33338]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:407
	_go_fuzz_dep_.CoverTab[33339]++
										return i.returnSlice(startp, i.p)
doNorm:

	i.p = startp
	i.info = i.rb.f.info(i.rb.src, i.p)
	i.rb.ss.first(i.info)
	if i.info.multiSegment() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:414
		_go_fuzz_dep_.CoverTab[33363]++
											d := i.info.Decomposition()
											info := i.rb.f.info(input{bytes: d}, 0)
											i.rb.insertUnsafe(input{bytes: d}, 0, info)
											i.multiSeg = d[int(info.size):]
											i.next = nextMultiNorm
											return nextMultiNorm(i)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:420
		// _ = "end of CoverTab[33363]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:421
		_go_fuzz_dep_.CoverTab[33364]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:421
		// _ = "end of CoverTab[33364]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:421
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:421
	// _ = "end of CoverTab[33339]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:421
	_go_fuzz_dep_.CoverTab[33340]++
										i.rb.ss.first(i.info)
										i.rb.insertUnsafe(i.rb.src, i.p, i.info)
										return doNormComposed(i)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:424
	// _ = "end of CoverTab[33340]"
}

func doNormComposed(i *Iter) []byte {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:427
	_go_fuzz_dep_.CoverTab[33365]++

										for {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:429
		_go_fuzz_dep_.CoverTab[33367]++
											if i.p += int(i.info.size); i.p >= i.rb.nsrc {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:430
			_go_fuzz_dep_.CoverTab[33370]++
												i.setDone()
												break
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:432
			// _ = "end of CoverTab[33370]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:433
			_go_fuzz_dep_.CoverTab[33371]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:433
			// _ = "end of CoverTab[33371]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:433
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:433
		// _ = "end of CoverTab[33367]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:433
		_go_fuzz_dep_.CoverTab[33368]++
											i.info = i.rb.f.info(i.rb.src, i.p)
											if s := i.rb.ss.next(i.info); s == ssStarter {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:435
			_go_fuzz_dep_.CoverTab[33372]++
												break
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:436
			// _ = "end of CoverTab[33372]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:437
			_go_fuzz_dep_.CoverTab[33373]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:437
			if s == ssOverflow {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:437
				_go_fuzz_dep_.CoverTab[33374]++
													i.next = nextCGJCompose
													break
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:439
				// _ = "end of CoverTab[33374]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:440
				_go_fuzz_dep_.CoverTab[33375]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:440
				// _ = "end of CoverTab[33375]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:440
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:440
			// _ = "end of CoverTab[33373]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:440
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:440
		// _ = "end of CoverTab[33368]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:440
		_go_fuzz_dep_.CoverTab[33369]++
											i.rb.insertUnsafe(i.rb.src, i.p, i.info)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:441
		// _ = "end of CoverTab[33369]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:442
	// _ = "end of CoverTab[33365]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:442
	_go_fuzz_dep_.CoverTab[33366]++
										i.rb.compose()
										seg := i.buf[:i.rb.flushCopy(i.buf[:])]
										return seg
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:445
	// _ = "end of CoverTab[33366]"
}

func nextCGJCompose(i *Iter) []byte {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:448
	_go_fuzz_dep_.CoverTab[33376]++
										i.rb.ss = 0
										i.rb.insertCGJ()
										i.next = nextComposed

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:455
	i.rb.ss.first(i.info)
										i.rb.insertUnsafe(i.rb.src, i.p, i.info)
										return doNormComposed(i)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:457
	// _ = "end of CoverTab[33376]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:458
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/iter.go:458
var _ = _go_fuzz_dep_.CoverTab
