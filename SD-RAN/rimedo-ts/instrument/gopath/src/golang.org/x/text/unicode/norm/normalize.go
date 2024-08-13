// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Note: the file data_test.go that is generated should not be checked in.
//go:generate go run maketables.go triegen.go
//go:generate go test -tags test

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:6
//go:generate go run maketables.go triegen.go
//go:generate go test -tags test

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:10
package norm

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:10
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:10
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:10
)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:10
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:10
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:10
)

import (
	"unicode/utf8"

	"golang.org/x/text/transform"
)

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:35
type Form int

const (
	NFC	Form	= iota
	NFD
	NFKC
	NFKD
)

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:45
func (f Form) Bytes(b []byte) []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:45
	_go_fuzz_dep_.CoverTab[70658]++
												src := inputBytes(b)
												ft := formTable[f]
												n, ok := ft.quickSpan(src, 0, len(b), true)
												if ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:49
		_go_fuzz_dep_.CoverTab[70660]++
													return b
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:50
		// _ = "end of CoverTab[70660]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:51
		_go_fuzz_dep_.CoverTab[70661]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:51
		// _ = "end of CoverTab[70661]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:51
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:51
	// _ = "end of CoverTab[70658]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:51
	_go_fuzz_dep_.CoverTab[70659]++
												out := make([]byte, n, len(b))
												copy(out, b[0:n])
												rb := reorderBuffer{f: *ft, src: src, nsrc: len(b), out: out, flushF: appendFlush}
												return doAppendInner(&rb, n)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:55
	// _ = "end of CoverTab[70659]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:59
func (f Form) String(s string) string {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:59
	_go_fuzz_dep_.CoverTab[70662]++
												src := inputString(s)
												ft := formTable[f]
												n, ok := ft.quickSpan(src, 0, len(s), true)
												if ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:63
		_go_fuzz_dep_.CoverTab[70664]++
													return s
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:64
		// _ = "end of CoverTab[70664]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:65
		_go_fuzz_dep_.CoverTab[70665]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:65
		// _ = "end of CoverTab[70665]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:65
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:65
	// _ = "end of CoverTab[70662]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:65
	_go_fuzz_dep_.CoverTab[70663]++
												out := make([]byte, n, len(s))
												copy(out, s[0:n])
												rb := reorderBuffer{f: *ft, src: src, nsrc: len(s), out: out, flushF: appendFlush}
												return string(doAppendInner(&rb, n))
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:69
	// _ = "end of CoverTab[70663]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:73
func (f Form) IsNormal(b []byte) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:73
	_go_fuzz_dep_.CoverTab[70666]++
												src := inputBytes(b)
												ft := formTable[f]
												bp, ok := ft.quickSpan(src, 0, len(b), true)
												if ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:77
		_go_fuzz_dep_.CoverTab[70669]++
													return true
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:78
		// _ = "end of CoverTab[70669]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:79
		_go_fuzz_dep_.CoverTab[70670]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:79
		// _ = "end of CoverTab[70670]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:79
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:79
	// _ = "end of CoverTab[70666]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:79
	_go_fuzz_dep_.CoverTab[70667]++
												rb := reorderBuffer{f: *ft, src: src, nsrc: len(b)}
												rb.setFlusher(nil, cmpNormalBytes)
												for bp < len(b) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:82
		_go_fuzz_dep_.CoverTab[70671]++
													rb.out = b[bp:]
													if bp = decomposeSegment(&rb, bp, true); bp < 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:84
			_go_fuzz_dep_.CoverTab[70673]++
														return false
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:85
			// _ = "end of CoverTab[70673]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:86
			_go_fuzz_dep_.CoverTab[70674]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:86
			// _ = "end of CoverTab[70674]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:86
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:86
		// _ = "end of CoverTab[70671]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:86
		_go_fuzz_dep_.CoverTab[70672]++
													bp, _ = rb.f.quickSpan(rb.src, bp, len(b), true)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:87
		// _ = "end of CoverTab[70672]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:88
	// _ = "end of CoverTab[70667]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:88
	_go_fuzz_dep_.CoverTab[70668]++
												return true
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:89
	// _ = "end of CoverTab[70668]"
}

func cmpNormalBytes(rb *reorderBuffer) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:92
	_go_fuzz_dep_.CoverTab[70675]++
												b := rb.out
												for i := 0; i < rb.nrune; i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:94
		_go_fuzz_dep_.CoverTab[70677]++
													info := rb.rune[i]
													if int(info.size) > len(b) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:96
			_go_fuzz_dep_.CoverTab[70679]++
														return false
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:97
			// _ = "end of CoverTab[70679]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:98
			_go_fuzz_dep_.CoverTab[70680]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:98
			// _ = "end of CoverTab[70680]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:98
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:98
		// _ = "end of CoverTab[70677]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:98
		_go_fuzz_dep_.CoverTab[70678]++
													p := info.pos
													pe := p + info.size
													for ; p < pe; p++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:101
			_go_fuzz_dep_.CoverTab[70681]++
														if b[0] != rb.byte[p] {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:102
				_go_fuzz_dep_.CoverTab[70683]++
															return false
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:103
				// _ = "end of CoverTab[70683]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:104
				_go_fuzz_dep_.CoverTab[70684]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:104
				// _ = "end of CoverTab[70684]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:104
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:104
			// _ = "end of CoverTab[70681]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:104
			_go_fuzz_dep_.CoverTab[70682]++
														b = b[1:]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:105
			// _ = "end of CoverTab[70682]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:106
		// _ = "end of CoverTab[70678]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:107
	// _ = "end of CoverTab[70675]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:107
	_go_fuzz_dep_.CoverTab[70676]++
												return true
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:108
	// _ = "end of CoverTab[70676]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:112
func (f Form) IsNormalString(s string) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:112
	_go_fuzz_dep_.CoverTab[70685]++
												src := inputString(s)
												ft := formTable[f]
												bp, ok := ft.quickSpan(src, 0, len(s), true)
												if ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:116
		_go_fuzz_dep_.CoverTab[70689]++
													return true
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:117
		// _ = "end of CoverTab[70689]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:118
		_go_fuzz_dep_.CoverTab[70690]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:118
		// _ = "end of CoverTab[70690]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:118
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:118
	// _ = "end of CoverTab[70685]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:118
	_go_fuzz_dep_.CoverTab[70686]++
												rb := reorderBuffer{f: *ft, src: src, nsrc: len(s)}
												rb.setFlusher(nil, func(rb *reorderBuffer) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:120
		_go_fuzz_dep_.CoverTab[70691]++
													for i := 0; i < rb.nrune; i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:121
			_go_fuzz_dep_.CoverTab[70693]++
														info := rb.rune[i]
														if bp+int(info.size) > len(s) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:123
				_go_fuzz_dep_.CoverTab[70695]++
															return false
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:124
				// _ = "end of CoverTab[70695]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:125
				_go_fuzz_dep_.CoverTab[70696]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:125
				// _ = "end of CoverTab[70696]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:125
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:125
			// _ = "end of CoverTab[70693]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:125
			_go_fuzz_dep_.CoverTab[70694]++
														p := info.pos
														pe := p + info.size
														for ; p < pe; p++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:128
				_go_fuzz_dep_.CoverTab[70697]++
															if s[bp] != rb.byte[p] {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:129
					_go_fuzz_dep_.CoverTab[70699]++
																return false
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:130
					// _ = "end of CoverTab[70699]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:131
					_go_fuzz_dep_.CoverTab[70700]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:131
					// _ = "end of CoverTab[70700]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:131
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:131
				// _ = "end of CoverTab[70697]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:131
				_go_fuzz_dep_.CoverTab[70698]++
															bp++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:132
				// _ = "end of CoverTab[70698]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:133
			// _ = "end of CoverTab[70694]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:134
		// _ = "end of CoverTab[70691]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:134
		_go_fuzz_dep_.CoverTab[70692]++
													return true
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:135
		// _ = "end of CoverTab[70692]"
	})
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:136
	// _ = "end of CoverTab[70686]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:136
	_go_fuzz_dep_.CoverTab[70687]++
												for bp < len(s) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:137
		_go_fuzz_dep_.CoverTab[70701]++
													if bp = decomposeSegment(&rb, bp, true); bp < 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:138
			_go_fuzz_dep_.CoverTab[70703]++
														return false
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:139
			// _ = "end of CoverTab[70703]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:140
			_go_fuzz_dep_.CoverTab[70704]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:140
			// _ = "end of CoverTab[70704]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:140
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:140
		// _ = "end of CoverTab[70701]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:140
		_go_fuzz_dep_.CoverTab[70702]++
													bp, _ = rb.f.quickSpan(rb.src, bp, len(s), true)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:141
		// _ = "end of CoverTab[70702]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:142
	// _ = "end of CoverTab[70687]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:142
	_go_fuzz_dep_.CoverTab[70688]++
												return true
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:143
	// _ = "end of CoverTab[70688]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:149
func patchTail(rb *reorderBuffer) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:149
	_go_fuzz_dep_.CoverTab[70705]++
												info, p := lastRuneStart(&rb.f, rb.out)
												if p == -1 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:151
		_go_fuzz_dep_.CoverTab[70709]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:151
		return info.size == 0
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:151
		// _ = "end of CoverTab[70709]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:151
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:151
		_go_fuzz_dep_.CoverTab[70710]++
													return true
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:152
		// _ = "end of CoverTab[70710]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:153
		_go_fuzz_dep_.CoverTab[70711]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:153
		// _ = "end of CoverTab[70711]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:153
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:153
	// _ = "end of CoverTab[70705]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:153
	_go_fuzz_dep_.CoverTab[70706]++
												end := p + int(info.size)
												extra := len(rb.out) - end
												if extra > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:156
		_go_fuzz_dep_.CoverTab[70712]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:159
		x := make([]byte, 0)
													x = append(x, rb.out[len(rb.out)-extra:]...)
													rb.out = rb.out[:end]
													decomposeToLastBoundary(rb)
													rb.doFlush()
													rb.out = append(rb.out, x...)
													return false
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:165
		// _ = "end of CoverTab[70712]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:166
		_go_fuzz_dep_.CoverTab[70713]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:166
		// _ = "end of CoverTab[70713]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:166
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:166
	// _ = "end of CoverTab[70706]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:166
	_go_fuzz_dep_.CoverTab[70707]++
												buf := rb.out[p:]
												rb.out = rb.out[:p]
												decomposeToLastBoundary(rb)
												if s := rb.ss.next(info); s == ssStarter {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:170
		_go_fuzz_dep_.CoverTab[70714]++
													rb.doFlush()
													rb.ss.first(info)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:172
		// _ = "end of CoverTab[70714]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:173
		_go_fuzz_dep_.CoverTab[70715]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:173
		if s == ssOverflow {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:173
			_go_fuzz_dep_.CoverTab[70716]++
														rb.doFlush()
														rb.insertCGJ()
														rb.ss = 0
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:176
			// _ = "end of CoverTab[70716]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:177
			_go_fuzz_dep_.CoverTab[70717]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:177
			// _ = "end of CoverTab[70717]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:177
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:177
		// _ = "end of CoverTab[70715]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:177
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:177
	// _ = "end of CoverTab[70707]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:177
	_go_fuzz_dep_.CoverTab[70708]++
												rb.insertUnsafe(inputBytes(buf), 0, info)
												return true
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:179
	// _ = "end of CoverTab[70708]"
}

func appendQuick(rb *reorderBuffer, i int) int {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:182
	_go_fuzz_dep_.CoverTab[70718]++
												if rb.nsrc == i {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:183
		_go_fuzz_dep_.CoverTab[70720]++
													return i
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:184
		// _ = "end of CoverTab[70720]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:185
		_go_fuzz_dep_.CoverTab[70721]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:185
		// _ = "end of CoverTab[70721]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:185
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:185
	// _ = "end of CoverTab[70718]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:185
	_go_fuzz_dep_.CoverTab[70719]++
												end, _ := rb.f.quickSpan(rb.src, i, rb.nsrc, true)
												rb.out = rb.src.appendSlice(rb.out, i, end)
												return end
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:188
	// _ = "end of CoverTab[70719]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:193
func (f Form) Append(out []byte, src ...byte) []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:193
	_go_fuzz_dep_.CoverTab[70722]++
												return f.doAppend(out, inputBytes(src), len(src))
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:194
	// _ = "end of CoverTab[70722]"
}

func (f Form) doAppend(out []byte, src input, n int) []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:197
	_go_fuzz_dep_.CoverTab[70723]++
												if n == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:198
		_go_fuzz_dep_.CoverTab[70726]++
													return out
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:199
		// _ = "end of CoverTab[70726]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:200
		_go_fuzz_dep_.CoverTab[70727]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:200
		// _ = "end of CoverTab[70727]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:200
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:200
	// _ = "end of CoverTab[70723]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:200
	_go_fuzz_dep_.CoverTab[70724]++
												ft := formTable[f]

												if len(out) == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:203
		_go_fuzz_dep_.CoverTab[70728]++
													p, _ := ft.quickSpan(src, 0, n, true)
													out = src.appendSlice(out, 0, p)
													if p == n {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:206
			_go_fuzz_dep_.CoverTab[70730]++
														return out
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:207
			// _ = "end of CoverTab[70730]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:208
			_go_fuzz_dep_.CoverTab[70731]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:208
			// _ = "end of CoverTab[70731]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:208
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:208
		// _ = "end of CoverTab[70728]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:208
		_go_fuzz_dep_.CoverTab[70729]++
													rb := reorderBuffer{f: *ft, src: src, nsrc: n, out: out, flushF: appendFlush}
													return doAppendInner(&rb, p)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:210
		// _ = "end of CoverTab[70729]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:211
		_go_fuzz_dep_.CoverTab[70732]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:211
		// _ = "end of CoverTab[70732]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:211
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:211
	// _ = "end of CoverTab[70724]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:211
	_go_fuzz_dep_.CoverTab[70725]++
												rb := reorderBuffer{f: *ft, src: src, nsrc: n}
												return doAppend(&rb, out, 0)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:213
	// _ = "end of CoverTab[70725]"
}

func doAppend(rb *reorderBuffer, out []byte, p int) []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:216
	_go_fuzz_dep_.CoverTab[70733]++
												rb.setFlusher(out, appendFlush)
												src, n := rb.src, rb.nsrc
												doMerge := len(out) > 0
												if q := src.skipContinuationBytes(p); q > p {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:220
		_go_fuzz_dep_.CoverTab[70736]++

													rb.out = src.appendSlice(rb.out, p, q)
													p = q
													doMerge = patchTail(rb)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:224
		// _ = "end of CoverTab[70736]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:225
		_go_fuzz_dep_.CoverTab[70737]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:225
		// _ = "end of CoverTab[70737]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:225
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:225
	// _ = "end of CoverTab[70733]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:225
	_go_fuzz_dep_.CoverTab[70734]++
												fd := &rb.f
												if doMerge {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:227
		_go_fuzz_dep_.CoverTab[70738]++
													var info Properties
													if p < n {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:229
			_go_fuzz_dep_.CoverTab[70741]++
														info = fd.info(src, p)
														if !info.BoundaryBefore() || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:231
				_go_fuzz_dep_.CoverTab[70742]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:231
				return info.nLeadingNonStarters() > 0
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:231
				// _ = "end of CoverTab[70742]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:231
			}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:231
				_go_fuzz_dep_.CoverTab[70743]++
															if p == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:232
					_go_fuzz_dep_.CoverTab[70745]++
																decomposeToLastBoundary(rb)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:233
					// _ = "end of CoverTab[70745]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:234
					_go_fuzz_dep_.CoverTab[70746]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:234
					// _ = "end of CoverTab[70746]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:234
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:234
				// _ = "end of CoverTab[70743]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:234
				_go_fuzz_dep_.CoverTab[70744]++
															p = decomposeSegment(rb, p, true)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:235
				// _ = "end of CoverTab[70744]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:236
				_go_fuzz_dep_.CoverTab[70747]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:236
				// _ = "end of CoverTab[70747]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:236
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:236
			// _ = "end of CoverTab[70741]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:237
			_go_fuzz_dep_.CoverTab[70748]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:237
			// _ = "end of CoverTab[70748]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:237
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:237
		// _ = "end of CoverTab[70738]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:237
		_go_fuzz_dep_.CoverTab[70739]++
													if info.size == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:238
			_go_fuzz_dep_.CoverTab[70749]++
														rb.doFlush()

														return src.appendSlice(rb.out, p, n)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:241
			// _ = "end of CoverTab[70749]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:242
			_go_fuzz_dep_.CoverTab[70750]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:242
			// _ = "end of CoverTab[70750]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:242
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:242
		// _ = "end of CoverTab[70739]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:242
		_go_fuzz_dep_.CoverTab[70740]++
													if rb.nrune > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:243
			_go_fuzz_dep_.CoverTab[70751]++
														return doAppendInner(rb, p)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:244
			// _ = "end of CoverTab[70751]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:245
			_go_fuzz_dep_.CoverTab[70752]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:245
			// _ = "end of CoverTab[70752]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:245
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:245
		// _ = "end of CoverTab[70740]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:246
		_go_fuzz_dep_.CoverTab[70753]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:246
		// _ = "end of CoverTab[70753]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:246
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:246
	// _ = "end of CoverTab[70734]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:246
	_go_fuzz_dep_.CoverTab[70735]++
												p = appendQuick(rb, p)
												return doAppendInner(rb, p)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:248
	// _ = "end of CoverTab[70735]"
}

func doAppendInner(rb *reorderBuffer, p int) []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:251
	_go_fuzz_dep_.CoverTab[70754]++
												for n := rb.nsrc; p < n; {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:252
		_go_fuzz_dep_.CoverTab[70756]++
													p = decomposeSegment(rb, p, true)
													p = appendQuick(rb, p)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:254
		// _ = "end of CoverTab[70756]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:255
	// _ = "end of CoverTab[70754]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:255
	_go_fuzz_dep_.CoverTab[70755]++
												return rb.out
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:256
	// _ = "end of CoverTab[70755]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:261
func (f Form) AppendString(out []byte, src string) []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:261
	_go_fuzz_dep_.CoverTab[70757]++
												return f.doAppend(out, inputString(src), len(src))
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:262
	// _ = "end of CoverTab[70757]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:267
func (f Form) QuickSpan(b []byte) int {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:267
	_go_fuzz_dep_.CoverTab[70758]++
												n, _ := formTable[f].quickSpan(inputBytes(b), 0, len(b), true)
												return n
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:269
	// _ = "end of CoverTab[70758]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:274
func (f Form) Span(b []byte, atEOF bool) (n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:274
	_go_fuzz_dep_.CoverTab[70759]++
												n, ok := formTable[f].quickSpan(inputBytes(b), 0, len(b), atEOF)
												if n < len(b) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:276
		_go_fuzz_dep_.CoverTab[70761]++
													if !ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:277
			_go_fuzz_dep_.CoverTab[70762]++
														err = transform.ErrEndOfSpan
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:278
			// _ = "end of CoverTab[70762]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:279
			_go_fuzz_dep_.CoverTab[70763]++
														err = transform.ErrShortSrc
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:280
			// _ = "end of CoverTab[70763]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:281
		// _ = "end of CoverTab[70761]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:282
		_go_fuzz_dep_.CoverTab[70764]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:282
		// _ = "end of CoverTab[70764]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:282
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:282
	// _ = "end of CoverTab[70759]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:282
	_go_fuzz_dep_.CoverTab[70760]++
												return n, err
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:283
	// _ = "end of CoverTab[70760]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:288
func (f Form) SpanString(s string, atEOF bool) (n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:288
	_go_fuzz_dep_.CoverTab[70765]++
												n, ok := formTable[f].quickSpan(inputString(s), 0, len(s), atEOF)
												if n < len(s) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:290
		_go_fuzz_dep_.CoverTab[70767]++
													if !ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:291
			_go_fuzz_dep_.CoverTab[70768]++
														err = transform.ErrEndOfSpan
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:292
			// _ = "end of CoverTab[70768]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:293
			_go_fuzz_dep_.CoverTab[70769]++
														err = transform.ErrShortSrc
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:294
			// _ = "end of CoverTab[70769]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:295
		// _ = "end of CoverTab[70767]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:296
		_go_fuzz_dep_.CoverTab[70770]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:296
		// _ = "end of CoverTab[70770]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:296
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:296
	// _ = "end of CoverTab[70765]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:296
	_go_fuzz_dep_.CoverTab[70766]++
												return n, err
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:297
	// _ = "end of CoverTab[70766]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:304
func (f *formInfo) quickSpan(src input, i, end int, atEOF bool) (n int, ok bool) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:304
	_go_fuzz_dep_.CoverTab[70771]++
												var lastCC uint8
												ss := streamSafe(0)
												lastSegStart := i
												for n = end; i < n; {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:308
		_go_fuzz_dep_.CoverTab[70774]++
													if j := src.skipASCII(i, n); i != j {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:309
			_go_fuzz_dep_.CoverTab[70779]++
														i = j
														lastSegStart = i - 1
														lastCC = 0
														ss = 0
														continue
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:314
			// _ = "end of CoverTab[70779]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:315
			_go_fuzz_dep_.CoverTab[70780]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:315
			// _ = "end of CoverTab[70780]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:315
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:315
		// _ = "end of CoverTab[70774]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:315
		_go_fuzz_dep_.CoverTab[70775]++
													info := f.info(src, i)
													if info.size == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:317
			_go_fuzz_dep_.CoverTab[70781]++
														if atEOF {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:318
				_go_fuzz_dep_.CoverTab[70783]++

															return n, true
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:320
				// _ = "end of CoverTab[70783]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:321
				_go_fuzz_dep_.CoverTab[70784]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:321
				// _ = "end of CoverTab[70784]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:321
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:321
			// _ = "end of CoverTab[70781]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:321
			_go_fuzz_dep_.CoverTab[70782]++
														return lastSegStart, true
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:322
			// _ = "end of CoverTab[70782]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:323
			_go_fuzz_dep_.CoverTab[70785]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:323
			// _ = "end of CoverTab[70785]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:323
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:323
		// _ = "end of CoverTab[70775]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:323
		_go_fuzz_dep_.CoverTab[70776]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:326
		switch ss.next(info) {
		case ssStarter:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:327
			_go_fuzz_dep_.CoverTab[70786]++
														lastSegStart = i
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:328
			// _ = "end of CoverTab[70786]"
		case ssOverflow:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:329
			_go_fuzz_dep_.CoverTab[70787]++
														return lastSegStart, false
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:330
			// _ = "end of CoverTab[70787]"
		case ssSuccess:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:331
			_go_fuzz_dep_.CoverTab[70788]++
														if lastCC > info.ccc {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:332
				_go_fuzz_dep_.CoverTab[70790]++
															return lastSegStart, false
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:333
				// _ = "end of CoverTab[70790]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:334
				_go_fuzz_dep_.CoverTab[70791]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:334
				// _ = "end of CoverTab[70791]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:334
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:334
			// _ = "end of CoverTab[70788]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:334
		default:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:334
			_go_fuzz_dep_.CoverTab[70789]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:334
			// _ = "end of CoverTab[70789]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:335
		// _ = "end of CoverTab[70776]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:335
		_go_fuzz_dep_.CoverTab[70777]++
													if f.composing {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:336
			_go_fuzz_dep_.CoverTab[70792]++
														if !info.isYesC() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:337
				_go_fuzz_dep_.CoverTab[70793]++
															break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:338
				// _ = "end of CoverTab[70793]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:339
				_go_fuzz_dep_.CoverTab[70794]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:339
				// _ = "end of CoverTab[70794]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:339
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:339
			// _ = "end of CoverTab[70792]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:340
			_go_fuzz_dep_.CoverTab[70795]++
														if !info.isYesD() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:341
				_go_fuzz_dep_.CoverTab[70796]++
															break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:342
				// _ = "end of CoverTab[70796]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:343
				_go_fuzz_dep_.CoverTab[70797]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:343
				// _ = "end of CoverTab[70797]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:343
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:343
			// _ = "end of CoverTab[70795]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:344
		// _ = "end of CoverTab[70777]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:344
		_go_fuzz_dep_.CoverTab[70778]++
													lastCC = info.ccc
													i += int(info.size)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:346
		// _ = "end of CoverTab[70778]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:347
	// _ = "end of CoverTab[70771]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:347
	_go_fuzz_dep_.CoverTab[70772]++
												if i == n {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:348
		_go_fuzz_dep_.CoverTab[70798]++
													if !atEOF {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:349
			_go_fuzz_dep_.CoverTab[70800]++
														n = lastSegStart
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:350
			// _ = "end of CoverTab[70800]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:351
			_go_fuzz_dep_.CoverTab[70801]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:351
			// _ = "end of CoverTab[70801]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:351
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:351
		// _ = "end of CoverTab[70798]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:351
		_go_fuzz_dep_.CoverTab[70799]++
													return n, true
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:352
		// _ = "end of CoverTab[70799]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:353
		_go_fuzz_dep_.CoverTab[70802]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:353
		// _ = "end of CoverTab[70802]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:353
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:353
	// _ = "end of CoverTab[70772]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:353
	_go_fuzz_dep_.CoverTab[70773]++
												return lastSegStart, false
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:354
	// _ = "end of CoverTab[70773]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:359
func (f Form) QuickSpanString(s string) int {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:359
	_go_fuzz_dep_.CoverTab[70803]++
												n, _ := formTable[f].quickSpan(inputString(s), 0, len(s), true)
												return n
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:361
	// _ = "end of CoverTab[70803]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:366
func (f Form) FirstBoundary(b []byte) int {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:366
	_go_fuzz_dep_.CoverTab[70804]++
												return f.firstBoundary(inputBytes(b), len(b))
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:367
	// _ = "end of CoverTab[70804]"
}

func (f Form) firstBoundary(src input, nsrc int) int {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:370
	_go_fuzz_dep_.CoverTab[70805]++
												i := src.skipContinuationBytes(0)
												if i >= nsrc {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:372
		_go_fuzz_dep_.CoverTab[70807]++
													return -1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:373
		// _ = "end of CoverTab[70807]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:374
		_go_fuzz_dep_.CoverTab[70808]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:374
		// _ = "end of CoverTab[70808]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:374
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:374
	// _ = "end of CoverTab[70805]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:374
	_go_fuzz_dep_.CoverTab[70806]++
												fd := formTable[f]
												ss := streamSafe(0)

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:380
	for {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:380
		_go_fuzz_dep_.CoverTab[70809]++
													info := fd.info(src, i)
													if info.size == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:382
			_go_fuzz_dep_.CoverTab[70812]++
														return -1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:383
			// _ = "end of CoverTab[70812]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:384
			_go_fuzz_dep_.CoverTab[70813]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:384
			// _ = "end of CoverTab[70813]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:384
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:384
		// _ = "end of CoverTab[70809]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:384
		_go_fuzz_dep_.CoverTab[70810]++
													if s := ss.next(info); s != ssSuccess {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:385
			_go_fuzz_dep_.CoverTab[70814]++
														return i
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:386
			// _ = "end of CoverTab[70814]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:387
			_go_fuzz_dep_.CoverTab[70815]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:387
			// _ = "end of CoverTab[70815]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:387
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:387
		// _ = "end of CoverTab[70810]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:387
		_go_fuzz_dep_.CoverTab[70811]++
													i += int(info.size)
													if i >= nsrc {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:389
			_go_fuzz_dep_.CoverTab[70816]++
														if !info.BoundaryAfter() && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:390
				_go_fuzz_dep_.CoverTab[70818]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:390
				return !ss.isMax()
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:390
				// _ = "end of CoverTab[70818]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:390
			}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:390
				_go_fuzz_dep_.CoverTab[70819]++
															return -1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:391
				// _ = "end of CoverTab[70819]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:392
				_go_fuzz_dep_.CoverTab[70820]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:392
				// _ = "end of CoverTab[70820]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:392
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:392
			// _ = "end of CoverTab[70816]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:392
			_go_fuzz_dep_.CoverTab[70817]++
														return nsrc
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:393
			// _ = "end of CoverTab[70817]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:394
			_go_fuzz_dep_.CoverTab[70821]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:394
			// _ = "end of CoverTab[70821]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:394
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:394
		// _ = "end of CoverTab[70811]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:395
	// _ = "end of CoverTab[70806]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:400
func (f Form) FirstBoundaryInString(s string) int {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:400
	_go_fuzz_dep_.CoverTab[70822]++
												return f.firstBoundary(inputString(s), len(s))
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:401
	// _ = "end of CoverTab[70822]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:407
func (f Form) NextBoundary(b []byte, atEOF bool) int {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:407
	_go_fuzz_dep_.CoverTab[70823]++
												return f.nextBoundary(inputBytes(b), len(b), atEOF)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:408
	// _ = "end of CoverTab[70823]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:414
func (f Form) NextBoundaryInString(s string, atEOF bool) int {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:414
	_go_fuzz_dep_.CoverTab[70824]++
												return f.nextBoundary(inputString(s), len(s), atEOF)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:415
	// _ = "end of CoverTab[70824]"
}

func (f Form) nextBoundary(src input, nsrc int, atEOF bool) int {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:418
	_go_fuzz_dep_.CoverTab[70825]++
												if nsrc == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:419
		_go_fuzz_dep_.CoverTab[70830]++
													if atEOF {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:420
			_go_fuzz_dep_.CoverTab[70832]++
														return 0
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:421
			// _ = "end of CoverTab[70832]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:422
			_go_fuzz_dep_.CoverTab[70833]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:422
			// _ = "end of CoverTab[70833]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:422
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:422
		// _ = "end of CoverTab[70830]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:422
		_go_fuzz_dep_.CoverTab[70831]++
													return -1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:423
		// _ = "end of CoverTab[70831]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:424
		_go_fuzz_dep_.CoverTab[70834]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:424
		// _ = "end of CoverTab[70834]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:424
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:424
	// _ = "end of CoverTab[70825]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:424
	_go_fuzz_dep_.CoverTab[70826]++
												fd := formTable[f]
												info := fd.info(src, 0)
												if info.size == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:427
		_go_fuzz_dep_.CoverTab[70835]++
													if atEOF {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:428
			_go_fuzz_dep_.CoverTab[70837]++
														return 1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:429
			// _ = "end of CoverTab[70837]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:430
			_go_fuzz_dep_.CoverTab[70838]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:430
			// _ = "end of CoverTab[70838]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:430
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:430
		// _ = "end of CoverTab[70835]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:430
		_go_fuzz_dep_.CoverTab[70836]++
													return -1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:431
		// _ = "end of CoverTab[70836]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:432
		_go_fuzz_dep_.CoverTab[70839]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:432
		// _ = "end of CoverTab[70839]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:432
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:432
	// _ = "end of CoverTab[70826]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:432
	_go_fuzz_dep_.CoverTab[70827]++
												ss := streamSafe(0)
												ss.first(info)

												for i := int(info.size); i < nsrc; i += int(info.size) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:436
		_go_fuzz_dep_.CoverTab[70840]++
													info = fd.info(src, i)
													if info.size == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:438
			_go_fuzz_dep_.CoverTab[70842]++
														if atEOF {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:439
				_go_fuzz_dep_.CoverTab[70844]++
															return i
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:440
				// _ = "end of CoverTab[70844]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:441
				_go_fuzz_dep_.CoverTab[70845]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:441
				// _ = "end of CoverTab[70845]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:441
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:441
			// _ = "end of CoverTab[70842]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:441
			_go_fuzz_dep_.CoverTab[70843]++
														return -1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:442
			// _ = "end of CoverTab[70843]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:443
			_go_fuzz_dep_.CoverTab[70846]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:443
			// _ = "end of CoverTab[70846]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:443
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:443
		// _ = "end of CoverTab[70840]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:443
		_go_fuzz_dep_.CoverTab[70841]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:446
		if s := ss.next(info); s != ssSuccess {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:446
			_go_fuzz_dep_.CoverTab[70847]++
														return i
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:447
			// _ = "end of CoverTab[70847]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:448
			_go_fuzz_dep_.CoverTab[70848]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:448
			// _ = "end of CoverTab[70848]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:448
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:448
		// _ = "end of CoverTab[70841]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:449
	// _ = "end of CoverTab[70827]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:449
	_go_fuzz_dep_.CoverTab[70828]++
												if !atEOF && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:450
		_go_fuzz_dep_.CoverTab[70849]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:450
		return !info.BoundaryAfter()
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:450
		// _ = "end of CoverTab[70849]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:450
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:450
		_go_fuzz_dep_.CoverTab[70850]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:450
		return !ss.isMax()
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:450
		// _ = "end of CoverTab[70850]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:450
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:450
		_go_fuzz_dep_.CoverTab[70851]++
													return -1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:451
		// _ = "end of CoverTab[70851]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:452
		_go_fuzz_dep_.CoverTab[70852]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:452
		// _ = "end of CoverTab[70852]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:452
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:452
	// _ = "end of CoverTab[70828]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:452
	_go_fuzz_dep_.CoverTab[70829]++
												return nsrc
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:453
	// _ = "end of CoverTab[70829]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:458
func (f Form) LastBoundary(b []byte) int {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:458
	_go_fuzz_dep_.CoverTab[70853]++
												return lastBoundary(formTable[f], b)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:459
	// _ = "end of CoverTab[70853]"
}

func lastBoundary(fd *formInfo, b []byte) int {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:462
	_go_fuzz_dep_.CoverTab[70854]++
												i := len(b)
												info, p := lastRuneStart(fd, b)
												if p == -1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:465
		_go_fuzz_dep_.CoverTab[70860]++
													return -1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:466
		// _ = "end of CoverTab[70860]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:467
		_go_fuzz_dep_.CoverTab[70861]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:467
		// _ = "end of CoverTab[70861]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:467
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:467
	// _ = "end of CoverTab[70854]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:467
	_go_fuzz_dep_.CoverTab[70855]++
												if info.size == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:468
		_go_fuzz_dep_.CoverTab[70862]++
													if p == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:469
			_go_fuzz_dep_.CoverTab[70864]++
														return -1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:470
			// _ = "end of CoverTab[70864]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:471
			_go_fuzz_dep_.CoverTab[70865]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:471
			// _ = "end of CoverTab[70865]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:471
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:471
		// _ = "end of CoverTab[70862]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:471
		_go_fuzz_dep_.CoverTab[70863]++
													i = p
													info, p = lastRuneStart(fd, b[:i])
													if p == -1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:474
			_go_fuzz_dep_.CoverTab[70866]++
														return i
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:475
			// _ = "end of CoverTab[70866]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:476
			_go_fuzz_dep_.CoverTab[70867]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:476
			// _ = "end of CoverTab[70867]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:476
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:476
		// _ = "end of CoverTab[70863]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:477
		_go_fuzz_dep_.CoverTab[70868]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:477
		// _ = "end of CoverTab[70868]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:477
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:477
	// _ = "end of CoverTab[70855]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:477
	_go_fuzz_dep_.CoverTab[70856]++
												if p+int(info.size) != i {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:478
		_go_fuzz_dep_.CoverTab[70869]++
													return i
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:479
		// _ = "end of CoverTab[70869]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:480
		_go_fuzz_dep_.CoverTab[70870]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:480
		// _ = "end of CoverTab[70870]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:480
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:480
	// _ = "end of CoverTab[70856]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:480
	_go_fuzz_dep_.CoverTab[70857]++
												if info.BoundaryAfter() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:481
		_go_fuzz_dep_.CoverTab[70871]++
													return i
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:482
		// _ = "end of CoverTab[70871]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:483
		_go_fuzz_dep_.CoverTab[70872]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:483
		// _ = "end of CoverTab[70872]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:483
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:483
	// _ = "end of CoverTab[70857]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:483
	_go_fuzz_dep_.CoverTab[70858]++
												ss := streamSafe(0)
												v := ss.backwards(info)
												for i = p; i >= 0 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:486
		_go_fuzz_dep_.CoverTab[70873]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:486
		return v != ssStarter
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:486
		// _ = "end of CoverTab[70873]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:486
	}(); i = p {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:486
		_go_fuzz_dep_.CoverTab[70874]++
													info, p = lastRuneStart(fd, b[:i])
													if v = ss.backwards(info); v == ssOverflow {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:488
			_go_fuzz_dep_.CoverTab[70876]++
														break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:489
			// _ = "end of CoverTab[70876]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:490
			_go_fuzz_dep_.CoverTab[70877]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:490
			// _ = "end of CoverTab[70877]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:490
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:490
		// _ = "end of CoverTab[70874]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:490
		_go_fuzz_dep_.CoverTab[70875]++
													if p+int(info.size) != i {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:491
			_go_fuzz_dep_.CoverTab[70878]++
														if p == -1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:492
				_go_fuzz_dep_.CoverTab[70880]++
															return -1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:493
				// _ = "end of CoverTab[70880]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:494
				_go_fuzz_dep_.CoverTab[70881]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:494
				// _ = "end of CoverTab[70881]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:494
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:494
			// _ = "end of CoverTab[70878]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:494
			_go_fuzz_dep_.CoverTab[70879]++
														return i
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:495
			// _ = "end of CoverTab[70879]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:496
			_go_fuzz_dep_.CoverTab[70882]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:496
			// _ = "end of CoverTab[70882]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:496
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:496
		// _ = "end of CoverTab[70875]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:497
	// _ = "end of CoverTab[70858]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:497
	_go_fuzz_dep_.CoverTab[70859]++
												return i
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:498
	// _ = "end of CoverTab[70859]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:504
func decomposeSegment(rb *reorderBuffer, sp int, atEOF bool) int {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:504
	_go_fuzz_dep_.CoverTab[70883]++

												info := rb.f.info(rb.src, sp)
												if info.size == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:507
		_go_fuzz_dep_.CoverTab[70889]++
													return 0
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:508
		// _ = "end of CoverTab[70889]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:509
		_go_fuzz_dep_.CoverTab[70890]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:509
		// _ = "end of CoverTab[70890]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:509
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:509
	// _ = "end of CoverTab[70883]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:509
	_go_fuzz_dep_.CoverTab[70884]++
												if s := rb.ss.next(info); s == ssStarter {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:510
		_go_fuzz_dep_.CoverTab[70891]++

													if rb.nrune > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:512
			_go_fuzz_dep_.CoverTab[70892]++
														goto end
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:513
			// _ = "end of CoverTab[70892]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:514
			_go_fuzz_dep_.CoverTab[70893]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:514
			// _ = "end of CoverTab[70893]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:514
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:514
		// _ = "end of CoverTab[70891]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:515
		_go_fuzz_dep_.CoverTab[70894]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:515
		if s == ssOverflow {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:515
			_go_fuzz_dep_.CoverTab[70895]++
														rb.insertCGJ()
														goto end
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:517
			// _ = "end of CoverTab[70895]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:518
			_go_fuzz_dep_.CoverTab[70896]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:518
			// _ = "end of CoverTab[70896]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:518
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:518
		// _ = "end of CoverTab[70894]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:518
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:518
	// _ = "end of CoverTab[70884]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:518
	_go_fuzz_dep_.CoverTab[70885]++
												if err := rb.insertFlush(rb.src, sp, info); err != iSuccess {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:519
		_go_fuzz_dep_.CoverTab[70897]++
													return int(err)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:520
		// _ = "end of CoverTab[70897]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:521
		_go_fuzz_dep_.CoverTab[70898]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:521
		// _ = "end of CoverTab[70898]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:521
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:521
	// _ = "end of CoverTab[70885]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:521
	_go_fuzz_dep_.CoverTab[70886]++
												for {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:522
		_go_fuzz_dep_.CoverTab[70899]++
													sp += int(info.size)
													if sp >= rb.nsrc {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:524
			_go_fuzz_dep_.CoverTab[70903]++
														if !atEOF && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:525
				_go_fuzz_dep_.CoverTab[70905]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:525
				return !info.BoundaryAfter()
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:525
				// _ = "end of CoverTab[70905]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:525
			}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:525
				_go_fuzz_dep_.CoverTab[70906]++
															return int(iShortSrc)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:526
				// _ = "end of CoverTab[70906]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:527
				_go_fuzz_dep_.CoverTab[70907]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:527
				// _ = "end of CoverTab[70907]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:527
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:527
			// _ = "end of CoverTab[70903]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:527
			_go_fuzz_dep_.CoverTab[70904]++
														break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:528
			// _ = "end of CoverTab[70904]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:529
			_go_fuzz_dep_.CoverTab[70908]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:529
			// _ = "end of CoverTab[70908]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:529
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:529
		// _ = "end of CoverTab[70899]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:529
		_go_fuzz_dep_.CoverTab[70900]++
													info = rb.f.info(rb.src, sp)
													if info.size == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:531
			_go_fuzz_dep_.CoverTab[70909]++
														if !atEOF {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:532
				_go_fuzz_dep_.CoverTab[70911]++
															return int(iShortSrc)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:533
				// _ = "end of CoverTab[70911]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:534
				_go_fuzz_dep_.CoverTab[70912]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:534
				// _ = "end of CoverTab[70912]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:534
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:534
			// _ = "end of CoverTab[70909]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:534
			_go_fuzz_dep_.CoverTab[70910]++
														break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:535
			// _ = "end of CoverTab[70910]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:536
			_go_fuzz_dep_.CoverTab[70913]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:536
			// _ = "end of CoverTab[70913]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:536
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:536
		// _ = "end of CoverTab[70900]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:536
		_go_fuzz_dep_.CoverTab[70901]++
													if s := rb.ss.next(info); s == ssStarter {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:537
			_go_fuzz_dep_.CoverTab[70914]++
														break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:538
			// _ = "end of CoverTab[70914]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:539
			_go_fuzz_dep_.CoverTab[70915]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:539
			if s == ssOverflow {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:539
				_go_fuzz_dep_.CoverTab[70916]++
															rb.insertCGJ()
															break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:541
				// _ = "end of CoverTab[70916]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:542
				_go_fuzz_dep_.CoverTab[70917]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:542
				// _ = "end of CoverTab[70917]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:542
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:542
			// _ = "end of CoverTab[70915]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:542
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:542
		// _ = "end of CoverTab[70901]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:542
		_go_fuzz_dep_.CoverTab[70902]++
													if err := rb.insertFlush(rb.src, sp, info); err != iSuccess {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:543
			_go_fuzz_dep_.CoverTab[70918]++
														return int(err)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:544
			// _ = "end of CoverTab[70918]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:545
			_go_fuzz_dep_.CoverTab[70919]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:545
			// _ = "end of CoverTab[70919]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:545
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:545
		// _ = "end of CoverTab[70902]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:546
	// _ = "end of CoverTab[70886]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:546
	_go_fuzz_dep_.CoverTab[70887]++
end:
	if !rb.doFlush() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:548
		_go_fuzz_dep_.CoverTab[70920]++
													return int(iShortDst)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:549
		// _ = "end of CoverTab[70920]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:550
		_go_fuzz_dep_.CoverTab[70921]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:550
		// _ = "end of CoverTab[70921]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:550
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:550
	// _ = "end of CoverTab[70887]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:550
	_go_fuzz_dep_.CoverTab[70888]++
												return sp
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:551
	// _ = "end of CoverTab[70888]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:556
func lastRuneStart(fd *formInfo, buf []byte) (Properties, int) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:556
	_go_fuzz_dep_.CoverTab[70922]++
												p := len(buf) - 1
												for ; p >= 0 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:558
		_go_fuzz_dep_.CoverTab[70925]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:558
		return !utf8.RuneStart(buf[p])
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:558
		// _ = "end of CoverTab[70925]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:558
	}(); p-- {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:558
		_go_fuzz_dep_.CoverTab[70926]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:558
		// _ = "end of CoverTab[70926]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:559
	// _ = "end of CoverTab[70922]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:559
	_go_fuzz_dep_.CoverTab[70923]++
												if p < 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:560
		_go_fuzz_dep_.CoverTab[70927]++
													return Properties{}, -1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:561
		// _ = "end of CoverTab[70927]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:562
		_go_fuzz_dep_.CoverTab[70928]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:562
		// _ = "end of CoverTab[70928]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:562
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:562
	// _ = "end of CoverTab[70923]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:562
	_go_fuzz_dep_.CoverTab[70924]++
												return fd.info(inputBytes(buf), p), p
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:563
	// _ = "end of CoverTab[70924]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:568
func decomposeToLastBoundary(rb *reorderBuffer) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:568
	_go_fuzz_dep_.CoverTab[70929]++
												fd := &rb.f
												info, i := lastRuneStart(fd, rb.out)
												if int(info.size) != len(rb.out)-i {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:571
		_go_fuzz_dep_.CoverTab[70933]++

													return
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:573
		// _ = "end of CoverTab[70933]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:574
		_go_fuzz_dep_.CoverTab[70934]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:574
		// _ = "end of CoverTab[70934]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:574
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:574
	// _ = "end of CoverTab[70929]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:574
	_go_fuzz_dep_.CoverTab[70930]++
												if info.BoundaryAfter() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:575
		_go_fuzz_dep_.CoverTab[70935]++
													return
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:576
		// _ = "end of CoverTab[70935]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:577
		_go_fuzz_dep_.CoverTab[70936]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:577
		// _ = "end of CoverTab[70936]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:577
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:577
	// _ = "end of CoverTab[70930]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:577
	_go_fuzz_dep_.CoverTab[70931]++
												var add [maxNonStarters + 1]Properties
												padd := 0
												ss := streamSafe(0)
												p := len(rb.out)
												for {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:582
		_go_fuzz_dep_.CoverTab[70937]++
													add[padd] = info
													v := ss.backwards(info)
													if v == ssOverflow {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:585
			_go_fuzz_dep_.CoverTab[70940]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:588
			break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:588
			// _ = "end of CoverTab[70940]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:589
			_go_fuzz_dep_.CoverTab[70941]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:589
			// _ = "end of CoverTab[70941]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:589
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:589
		// _ = "end of CoverTab[70937]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:589
		_go_fuzz_dep_.CoverTab[70938]++
													padd++
													p -= int(info.size)
													if v == ssStarter || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:592
			_go_fuzz_dep_.CoverTab[70942]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:592
			return p < 0
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:592
			// _ = "end of CoverTab[70942]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:592
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:592
			_go_fuzz_dep_.CoverTab[70943]++
														break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:593
			// _ = "end of CoverTab[70943]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:594
			_go_fuzz_dep_.CoverTab[70944]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:594
			// _ = "end of CoverTab[70944]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:594
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:594
		// _ = "end of CoverTab[70938]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:594
		_go_fuzz_dep_.CoverTab[70939]++
													info, i = lastRuneStart(fd, rb.out[:p])
													if int(info.size) != p-i {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:596
			_go_fuzz_dep_.CoverTab[70945]++
														break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:597
			// _ = "end of CoverTab[70945]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:598
			_go_fuzz_dep_.CoverTab[70946]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:598
			// _ = "end of CoverTab[70946]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:598
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:598
		// _ = "end of CoverTab[70939]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:599
	// _ = "end of CoverTab[70931]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:599
	_go_fuzz_dep_.CoverTab[70932]++
												rb.ss = ss

												var buf [maxBufferSize * utf8.UTFMax]byte
												cp := buf[:copy(buf[:], rb.out[p:])]
												rb.out = rb.out[:p]
												for padd--; padd >= 0; padd-- {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:605
		_go_fuzz_dep_.CoverTab[70947]++
													info = add[padd]
													rb.insertUnsafe(inputBytes(cp), 0, info)
													cp = cp[info.size:]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:608
		// _ = "end of CoverTab[70947]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:609
	// _ = "end of CoverTab[70932]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:610
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/normalize.go:610
var _ = _go_fuzz_dep_.CoverTab
