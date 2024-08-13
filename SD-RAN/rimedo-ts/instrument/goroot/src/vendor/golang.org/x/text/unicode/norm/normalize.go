// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Note: the file data_test.go that is generated should not be checked in.
//go:generate go run maketables.go triegen.go
//go:generate go test -tags test

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:6
//go:generate go run maketables.go triegen.go
//go:generate go test -tags test

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:10
package norm

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:10
import (
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:10
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:10
)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:10
import (
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:10
	_atomic_ "sync/atomic"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:10
)

import (
	"unicode/utf8"

	"golang.org/x/text/transform"
)

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:35
type Form int

const (
	NFC	Form	= iota
	NFD
	NFKC
	NFKD
)

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:45
func (f Form) Bytes(b []byte) []byte {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:45
	_go_fuzz_dep_.CoverTab[33377]++
										src := inputBytes(b)
										ft := formTable[f]
										n, ok := ft.quickSpan(src, 0, len(b), true)
										if ok {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:49
		_go_fuzz_dep_.CoverTab[33379]++
											return b
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:50
		// _ = "end of CoverTab[33379]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:51
		_go_fuzz_dep_.CoverTab[33380]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:51
		// _ = "end of CoverTab[33380]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:51
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:51
	// _ = "end of CoverTab[33377]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:51
	_go_fuzz_dep_.CoverTab[33378]++
										out := make([]byte, n, len(b))
										copy(out, b[0:n])
										rb := reorderBuffer{f: *ft, src: src, nsrc: len(b), out: out, flushF: appendFlush}
										return doAppendInner(&rb, n)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:55
	// _ = "end of CoverTab[33378]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:59
func (f Form) String(s string) string {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:59
	_go_fuzz_dep_.CoverTab[33381]++
										src := inputString(s)
										ft := formTable[f]
										n, ok := ft.quickSpan(src, 0, len(s), true)
										if ok {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:63
		_go_fuzz_dep_.CoverTab[33383]++
											return s
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:64
		// _ = "end of CoverTab[33383]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:65
		_go_fuzz_dep_.CoverTab[33384]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:65
		// _ = "end of CoverTab[33384]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:65
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:65
	// _ = "end of CoverTab[33381]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:65
	_go_fuzz_dep_.CoverTab[33382]++
										out := make([]byte, n, len(s))
										copy(out, s[0:n])
										rb := reorderBuffer{f: *ft, src: src, nsrc: len(s), out: out, flushF: appendFlush}
										return string(doAppendInner(&rb, n))
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:69
	// _ = "end of CoverTab[33382]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:73
func (f Form) IsNormal(b []byte) bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:73
	_go_fuzz_dep_.CoverTab[33385]++
										src := inputBytes(b)
										ft := formTable[f]
										bp, ok := ft.quickSpan(src, 0, len(b), true)
										if ok {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:77
		_go_fuzz_dep_.CoverTab[33388]++
											return true
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:78
		// _ = "end of CoverTab[33388]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:79
		_go_fuzz_dep_.CoverTab[33389]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:79
		// _ = "end of CoverTab[33389]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:79
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:79
	// _ = "end of CoverTab[33385]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:79
	_go_fuzz_dep_.CoverTab[33386]++
										rb := reorderBuffer{f: *ft, src: src, nsrc: len(b)}
										rb.setFlusher(nil, cmpNormalBytes)
										for bp < len(b) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:82
		_go_fuzz_dep_.CoverTab[33390]++
											rb.out = b[bp:]
											if bp = decomposeSegment(&rb, bp, true); bp < 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:84
			_go_fuzz_dep_.CoverTab[33392]++
												return false
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:85
			// _ = "end of CoverTab[33392]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:86
			_go_fuzz_dep_.CoverTab[33393]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:86
			// _ = "end of CoverTab[33393]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:86
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:86
		// _ = "end of CoverTab[33390]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:86
		_go_fuzz_dep_.CoverTab[33391]++
											bp, _ = rb.f.quickSpan(rb.src, bp, len(b), true)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:87
		// _ = "end of CoverTab[33391]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:88
	// _ = "end of CoverTab[33386]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:88
	_go_fuzz_dep_.CoverTab[33387]++
										return true
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:89
	// _ = "end of CoverTab[33387]"
}

func cmpNormalBytes(rb *reorderBuffer) bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:92
	_go_fuzz_dep_.CoverTab[33394]++
										b := rb.out
										for i := 0; i < rb.nrune; i++ {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:94
		_go_fuzz_dep_.CoverTab[33396]++
											info := rb.rune[i]
											if int(info.size) > len(b) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:96
			_go_fuzz_dep_.CoverTab[33398]++
												return false
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:97
			// _ = "end of CoverTab[33398]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:98
			_go_fuzz_dep_.CoverTab[33399]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:98
			// _ = "end of CoverTab[33399]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:98
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:98
		// _ = "end of CoverTab[33396]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:98
		_go_fuzz_dep_.CoverTab[33397]++
											p := info.pos
											pe := p + info.size
											for ; p < pe; p++ {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:101
			_go_fuzz_dep_.CoverTab[33400]++
													if b[0] != rb.byte[p] {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:102
				_go_fuzz_dep_.CoverTab[33402]++
														return false
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:103
				// _ = "end of CoverTab[33402]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:104
				_go_fuzz_dep_.CoverTab[33403]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:104
				// _ = "end of CoverTab[33403]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:104
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:104
			// _ = "end of CoverTab[33400]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:104
			_go_fuzz_dep_.CoverTab[33401]++
													b = b[1:]
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:105
			// _ = "end of CoverTab[33401]"
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:106
		// _ = "end of CoverTab[33397]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:107
	// _ = "end of CoverTab[33394]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:107
	_go_fuzz_dep_.CoverTab[33395]++
											return true
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:108
	// _ = "end of CoverTab[33395]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:112
func (f Form) IsNormalString(s string) bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:112
	_go_fuzz_dep_.CoverTab[33404]++
											src := inputString(s)
											ft := formTable[f]
											bp, ok := ft.quickSpan(src, 0, len(s), true)
											if ok {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:116
		_go_fuzz_dep_.CoverTab[33408]++
												return true
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:117
		// _ = "end of CoverTab[33408]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:118
		_go_fuzz_dep_.CoverTab[33409]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:118
		// _ = "end of CoverTab[33409]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:118
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:118
	// _ = "end of CoverTab[33404]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:118
	_go_fuzz_dep_.CoverTab[33405]++
											rb := reorderBuffer{f: *ft, src: src, nsrc: len(s)}
											rb.setFlusher(nil, func(rb *reorderBuffer) bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:120
		_go_fuzz_dep_.CoverTab[33410]++
												for i := 0; i < rb.nrune; i++ {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:121
			_go_fuzz_dep_.CoverTab[33412]++
													info := rb.rune[i]
													if bp+int(info.size) > len(s) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:123
				_go_fuzz_dep_.CoverTab[33414]++
														return false
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:124
				// _ = "end of CoverTab[33414]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:125
				_go_fuzz_dep_.CoverTab[33415]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:125
				// _ = "end of CoverTab[33415]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:125
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:125
			// _ = "end of CoverTab[33412]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:125
			_go_fuzz_dep_.CoverTab[33413]++
													p := info.pos
													pe := p + info.size
													for ; p < pe; p++ {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:128
				_go_fuzz_dep_.CoverTab[33416]++
														if s[bp] != rb.byte[p] {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:129
					_go_fuzz_dep_.CoverTab[33418]++
															return false
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:130
					// _ = "end of CoverTab[33418]"
				} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:131
					_go_fuzz_dep_.CoverTab[33419]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:131
					// _ = "end of CoverTab[33419]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:131
				}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:131
				// _ = "end of CoverTab[33416]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:131
				_go_fuzz_dep_.CoverTab[33417]++
														bp++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:132
				// _ = "end of CoverTab[33417]"
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:133
			// _ = "end of CoverTab[33413]"
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:134
		// _ = "end of CoverTab[33410]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:134
		_go_fuzz_dep_.CoverTab[33411]++
												return true
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:135
		// _ = "end of CoverTab[33411]"
	})
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:136
	// _ = "end of CoverTab[33405]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:136
	_go_fuzz_dep_.CoverTab[33406]++
											for bp < len(s) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:137
		_go_fuzz_dep_.CoverTab[33420]++
												if bp = decomposeSegment(&rb, bp, true); bp < 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:138
			_go_fuzz_dep_.CoverTab[33422]++
													return false
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:139
			// _ = "end of CoverTab[33422]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:140
			_go_fuzz_dep_.CoverTab[33423]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:140
			// _ = "end of CoverTab[33423]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:140
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:140
		// _ = "end of CoverTab[33420]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:140
		_go_fuzz_dep_.CoverTab[33421]++
												bp, _ = rb.f.quickSpan(rb.src, bp, len(s), true)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:141
		// _ = "end of CoverTab[33421]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:142
	// _ = "end of CoverTab[33406]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:142
	_go_fuzz_dep_.CoverTab[33407]++
											return true
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:143
	// _ = "end of CoverTab[33407]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:149
func patchTail(rb *reorderBuffer) bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:149
	_go_fuzz_dep_.CoverTab[33424]++
											info, p := lastRuneStart(&rb.f, rb.out)
											if p == -1 || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:151
		_go_fuzz_dep_.CoverTab[33428]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:151
		return info.size == 0
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:151
		// _ = "end of CoverTab[33428]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:151
	}() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:151
		_go_fuzz_dep_.CoverTab[33429]++
												return true
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:152
		// _ = "end of CoverTab[33429]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:153
		_go_fuzz_dep_.CoverTab[33430]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:153
		// _ = "end of CoverTab[33430]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:153
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:153
	// _ = "end of CoverTab[33424]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:153
	_go_fuzz_dep_.CoverTab[33425]++
											end := p + int(info.size)
											extra := len(rb.out) - end
											if extra > 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:156
		_go_fuzz_dep_.CoverTab[33431]++

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:159
		x := make([]byte, 0)
												x = append(x, rb.out[len(rb.out)-extra:]...)
												rb.out = rb.out[:end]
												decomposeToLastBoundary(rb)
												rb.doFlush()
												rb.out = append(rb.out, x...)
												return false
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:165
		// _ = "end of CoverTab[33431]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:166
		_go_fuzz_dep_.CoverTab[33432]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:166
		// _ = "end of CoverTab[33432]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:166
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:166
	// _ = "end of CoverTab[33425]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:166
	_go_fuzz_dep_.CoverTab[33426]++
											buf := rb.out[p:]
											rb.out = rb.out[:p]
											decomposeToLastBoundary(rb)
											if s := rb.ss.next(info); s == ssStarter {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:170
		_go_fuzz_dep_.CoverTab[33433]++
												rb.doFlush()
												rb.ss.first(info)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:172
		// _ = "end of CoverTab[33433]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:173
		_go_fuzz_dep_.CoverTab[33434]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:173
		if s == ssOverflow {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:173
			_go_fuzz_dep_.CoverTab[33435]++
													rb.doFlush()
													rb.insertCGJ()
													rb.ss = 0
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:176
			// _ = "end of CoverTab[33435]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:177
			_go_fuzz_dep_.CoverTab[33436]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:177
			// _ = "end of CoverTab[33436]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:177
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:177
		// _ = "end of CoverTab[33434]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:177
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:177
	// _ = "end of CoverTab[33426]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:177
	_go_fuzz_dep_.CoverTab[33427]++
											rb.insertUnsafe(inputBytes(buf), 0, info)
											return true
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:179
	// _ = "end of CoverTab[33427]"
}

func appendQuick(rb *reorderBuffer, i int) int {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:182
	_go_fuzz_dep_.CoverTab[33437]++
											if rb.nsrc == i {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:183
		_go_fuzz_dep_.CoverTab[33439]++
												return i
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:184
		// _ = "end of CoverTab[33439]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:185
		_go_fuzz_dep_.CoverTab[33440]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:185
		// _ = "end of CoverTab[33440]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:185
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:185
	// _ = "end of CoverTab[33437]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:185
	_go_fuzz_dep_.CoverTab[33438]++
											end, _ := rb.f.quickSpan(rb.src, i, rb.nsrc, true)
											rb.out = rb.src.appendSlice(rb.out, i, end)
											return end
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:188
	// _ = "end of CoverTab[33438]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:193
func (f Form) Append(out []byte, src ...byte) []byte {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:193
	_go_fuzz_dep_.CoverTab[33441]++
											return f.doAppend(out, inputBytes(src), len(src))
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:194
	// _ = "end of CoverTab[33441]"
}

func (f Form) doAppend(out []byte, src input, n int) []byte {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:197
	_go_fuzz_dep_.CoverTab[33442]++
											if n == 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:198
		_go_fuzz_dep_.CoverTab[33445]++
												return out
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:199
		// _ = "end of CoverTab[33445]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:200
		_go_fuzz_dep_.CoverTab[33446]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:200
		// _ = "end of CoverTab[33446]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:200
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:200
	// _ = "end of CoverTab[33442]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:200
	_go_fuzz_dep_.CoverTab[33443]++
											ft := formTable[f]

											if len(out) == 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:203
		_go_fuzz_dep_.CoverTab[33447]++
												p, _ := ft.quickSpan(src, 0, n, true)
												out = src.appendSlice(out, 0, p)
												if p == n {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:206
			_go_fuzz_dep_.CoverTab[33449]++
													return out
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:207
			// _ = "end of CoverTab[33449]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:208
			_go_fuzz_dep_.CoverTab[33450]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:208
			// _ = "end of CoverTab[33450]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:208
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:208
		// _ = "end of CoverTab[33447]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:208
		_go_fuzz_dep_.CoverTab[33448]++
												rb := reorderBuffer{f: *ft, src: src, nsrc: n, out: out, flushF: appendFlush}
												return doAppendInner(&rb, p)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:210
		// _ = "end of CoverTab[33448]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:211
		_go_fuzz_dep_.CoverTab[33451]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:211
		// _ = "end of CoverTab[33451]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:211
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:211
	// _ = "end of CoverTab[33443]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:211
	_go_fuzz_dep_.CoverTab[33444]++
											rb := reorderBuffer{f: *ft, src: src, nsrc: n}
											return doAppend(&rb, out, 0)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:213
	// _ = "end of CoverTab[33444]"
}

func doAppend(rb *reorderBuffer, out []byte, p int) []byte {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:216
	_go_fuzz_dep_.CoverTab[33452]++
											rb.setFlusher(out, appendFlush)
											src, n := rb.src, rb.nsrc
											doMerge := len(out) > 0
											if q := src.skipContinuationBytes(p); q > p {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:220
		_go_fuzz_dep_.CoverTab[33455]++

												rb.out = src.appendSlice(rb.out, p, q)
												p = q
												doMerge = patchTail(rb)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:224
		// _ = "end of CoverTab[33455]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:225
		_go_fuzz_dep_.CoverTab[33456]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:225
		// _ = "end of CoverTab[33456]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:225
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:225
	// _ = "end of CoverTab[33452]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:225
	_go_fuzz_dep_.CoverTab[33453]++
											fd := &rb.f
											if doMerge {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:227
		_go_fuzz_dep_.CoverTab[33457]++
												var info Properties
												if p < n {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:229
			_go_fuzz_dep_.CoverTab[33460]++
													info = fd.info(src, p)
													if !info.BoundaryBefore() || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:231
				_go_fuzz_dep_.CoverTab[33461]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:231
				return info.nLeadingNonStarters() > 0
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:231
				// _ = "end of CoverTab[33461]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:231
			}() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:231
				_go_fuzz_dep_.CoverTab[33462]++
														if p == 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:232
					_go_fuzz_dep_.CoverTab[33464]++
															decomposeToLastBoundary(rb)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:233
					// _ = "end of CoverTab[33464]"
				} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:234
					_go_fuzz_dep_.CoverTab[33465]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:234
					// _ = "end of CoverTab[33465]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:234
				}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:234
				// _ = "end of CoverTab[33462]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:234
				_go_fuzz_dep_.CoverTab[33463]++
														p = decomposeSegment(rb, p, true)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:235
				// _ = "end of CoverTab[33463]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:236
				_go_fuzz_dep_.CoverTab[33466]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:236
				// _ = "end of CoverTab[33466]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:236
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:236
			// _ = "end of CoverTab[33460]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:237
			_go_fuzz_dep_.CoverTab[33467]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:237
			// _ = "end of CoverTab[33467]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:237
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:237
		// _ = "end of CoverTab[33457]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:237
		_go_fuzz_dep_.CoverTab[33458]++
												if info.size == 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:238
			_go_fuzz_dep_.CoverTab[33468]++
													rb.doFlush()

													return src.appendSlice(rb.out, p, n)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:241
			// _ = "end of CoverTab[33468]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:242
			_go_fuzz_dep_.CoverTab[33469]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:242
			// _ = "end of CoverTab[33469]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:242
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:242
		// _ = "end of CoverTab[33458]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:242
		_go_fuzz_dep_.CoverTab[33459]++
												if rb.nrune > 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:243
			_go_fuzz_dep_.CoverTab[33470]++
													return doAppendInner(rb, p)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:244
			// _ = "end of CoverTab[33470]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:245
			_go_fuzz_dep_.CoverTab[33471]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:245
			// _ = "end of CoverTab[33471]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:245
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:245
		// _ = "end of CoverTab[33459]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:246
		_go_fuzz_dep_.CoverTab[33472]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:246
		// _ = "end of CoverTab[33472]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:246
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:246
	// _ = "end of CoverTab[33453]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:246
	_go_fuzz_dep_.CoverTab[33454]++
											p = appendQuick(rb, p)
											return doAppendInner(rb, p)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:248
	// _ = "end of CoverTab[33454]"
}

func doAppendInner(rb *reorderBuffer, p int) []byte {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:251
	_go_fuzz_dep_.CoverTab[33473]++
											for n := rb.nsrc; p < n; {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:252
		_go_fuzz_dep_.CoverTab[33475]++
												p = decomposeSegment(rb, p, true)
												p = appendQuick(rb, p)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:254
		// _ = "end of CoverTab[33475]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:255
	// _ = "end of CoverTab[33473]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:255
	_go_fuzz_dep_.CoverTab[33474]++
											return rb.out
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:256
	// _ = "end of CoverTab[33474]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:261
func (f Form) AppendString(out []byte, src string) []byte {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:261
	_go_fuzz_dep_.CoverTab[33476]++
											return f.doAppend(out, inputString(src), len(src))
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:262
	// _ = "end of CoverTab[33476]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:267
func (f Form) QuickSpan(b []byte) int {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:267
	_go_fuzz_dep_.CoverTab[33477]++
											n, _ := formTable[f].quickSpan(inputBytes(b), 0, len(b), true)
											return n
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:269
	// _ = "end of CoverTab[33477]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:274
func (f Form) Span(b []byte, atEOF bool) (n int, err error) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:274
	_go_fuzz_dep_.CoverTab[33478]++
											n, ok := formTable[f].quickSpan(inputBytes(b), 0, len(b), atEOF)
											if n < len(b) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:276
		_go_fuzz_dep_.CoverTab[33480]++
												if !ok {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:277
			_go_fuzz_dep_.CoverTab[33481]++
													err = transform.ErrEndOfSpan
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:278
			// _ = "end of CoverTab[33481]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:279
			_go_fuzz_dep_.CoverTab[33482]++
													err = transform.ErrShortSrc
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:280
			// _ = "end of CoverTab[33482]"
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:281
		// _ = "end of CoverTab[33480]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:282
		_go_fuzz_dep_.CoverTab[33483]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:282
		// _ = "end of CoverTab[33483]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:282
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:282
	// _ = "end of CoverTab[33478]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:282
	_go_fuzz_dep_.CoverTab[33479]++
											return n, err
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:283
	// _ = "end of CoverTab[33479]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:288
func (f Form) SpanString(s string, atEOF bool) (n int, err error) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:288
	_go_fuzz_dep_.CoverTab[33484]++
											n, ok := formTable[f].quickSpan(inputString(s), 0, len(s), atEOF)
											if n < len(s) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:290
		_go_fuzz_dep_.CoverTab[33486]++
												if !ok {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:291
			_go_fuzz_dep_.CoverTab[33487]++
													err = transform.ErrEndOfSpan
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:292
			// _ = "end of CoverTab[33487]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:293
			_go_fuzz_dep_.CoverTab[33488]++
													err = transform.ErrShortSrc
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:294
			// _ = "end of CoverTab[33488]"
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:295
		// _ = "end of CoverTab[33486]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:296
		_go_fuzz_dep_.CoverTab[33489]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:296
		// _ = "end of CoverTab[33489]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:296
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:296
	// _ = "end of CoverTab[33484]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:296
	_go_fuzz_dep_.CoverTab[33485]++
											return n, err
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:297
	// _ = "end of CoverTab[33485]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:304
func (f *formInfo) quickSpan(src input, i, end int, atEOF bool) (n int, ok bool) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:304
	_go_fuzz_dep_.CoverTab[33490]++
											var lastCC uint8
											ss := streamSafe(0)
											lastSegStart := i
											for n = end; i < n; {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:308
		_go_fuzz_dep_.CoverTab[33493]++
												if j := src.skipASCII(i, n); i != j {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:309
			_go_fuzz_dep_.CoverTab[33498]++
													i = j
													lastSegStart = i - 1
													lastCC = 0
													ss = 0
													continue
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:314
			// _ = "end of CoverTab[33498]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:315
			_go_fuzz_dep_.CoverTab[33499]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:315
			// _ = "end of CoverTab[33499]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:315
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:315
		// _ = "end of CoverTab[33493]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:315
		_go_fuzz_dep_.CoverTab[33494]++
												info := f.info(src, i)
												if info.size == 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:317
			_go_fuzz_dep_.CoverTab[33500]++
													if atEOF {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:318
				_go_fuzz_dep_.CoverTab[33502]++

														return n, true
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:320
				// _ = "end of CoverTab[33502]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:321
				_go_fuzz_dep_.CoverTab[33503]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:321
				// _ = "end of CoverTab[33503]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:321
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:321
			// _ = "end of CoverTab[33500]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:321
			_go_fuzz_dep_.CoverTab[33501]++
													return lastSegStart, true
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:322
			// _ = "end of CoverTab[33501]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:323
			_go_fuzz_dep_.CoverTab[33504]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:323
			// _ = "end of CoverTab[33504]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:323
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:323
		// _ = "end of CoverTab[33494]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:323
		_go_fuzz_dep_.CoverTab[33495]++

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:326
		switch ss.next(info) {
		case ssStarter:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:327
			_go_fuzz_dep_.CoverTab[33505]++
													lastSegStart = i
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:328
			// _ = "end of CoverTab[33505]"
		case ssOverflow:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:329
			_go_fuzz_dep_.CoverTab[33506]++
													return lastSegStart, false
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:330
			// _ = "end of CoverTab[33506]"
		case ssSuccess:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:331
			_go_fuzz_dep_.CoverTab[33507]++
													if lastCC > info.ccc {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:332
				_go_fuzz_dep_.CoverTab[33509]++
														return lastSegStart, false
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:333
				// _ = "end of CoverTab[33509]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:334
				_go_fuzz_dep_.CoverTab[33510]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:334
				// _ = "end of CoverTab[33510]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:334
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:334
			// _ = "end of CoverTab[33507]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:334
		default:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:334
			_go_fuzz_dep_.CoverTab[33508]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:334
			// _ = "end of CoverTab[33508]"
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:335
		// _ = "end of CoverTab[33495]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:335
		_go_fuzz_dep_.CoverTab[33496]++
												if f.composing {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:336
			_go_fuzz_dep_.CoverTab[33511]++
													if !info.isYesC() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:337
				_go_fuzz_dep_.CoverTab[33512]++
														break
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:338
				// _ = "end of CoverTab[33512]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:339
				_go_fuzz_dep_.CoverTab[33513]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:339
				// _ = "end of CoverTab[33513]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:339
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:339
			// _ = "end of CoverTab[33511]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:340
			_go_fuzz_dep_.CoverTab[33514]++
													if !info.isYesD() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:341
				_go_fuzz_dep_.CoverTab[33515]++
														break
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:342
				// _ = "end of CoverTab[33515]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:343
				_go_fuzz_dep_.CoverTab[33516]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:343
				// _ = "end of CoverTab[33516]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:343
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:343
			// _ = "end of CoverTab[33514]"
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:344
		// _ = "end of CoverTab[33496]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:344
		_go_fuzz_dep_.CoverTab[33497]++
												lastCC = info.ccc
												i += int(info.size)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:346
		// _ = "end of CoverTab[33497]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:347
	// _ = "end of CoverTab[33490]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:347
	_go_fuzz_dep_.CoverTab[33491]++
											if i == n {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:348
		_go_fuzz_dep_.CoverTab[33517]++
												if !atEOF {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:349
			_go_fuzz_dep_.CoverTab[33519]++
													n = lastSegStart
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:350
			// _ = "end of CoverTab[33519]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:351
			_go_fuzz_dep_.CoverTab[33520]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:351
			// _ = "end of CoverTab[33520]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:351
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:351
		// _ = "end of CoverTab[33517]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:351
		_go_fuzz_dep_.CoverTab[33518]++
												return n, true
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:352
		// _ = "end of CoverTab[33518]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:353
		_go_fuzz_dep_.CoverTab[33521]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:353
		// _ = "end of CoverTab[33521]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:353
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:353
	// _ = "end of CoverTab[33491]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:353
	_go_fuzz_dep_.CoverTab[33492]++
											return lastSegStart, false
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:354
	// _ = "end of CoverTab[33492]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:359
func (f Form) QuickSpanString(s string) int {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:359
	_go_fuzz_dep_.CoverTab[33522]++
											n, _ := formTable[f].quickSpan(inputString(s), 0, len(s), true)
											return n
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:361
	// _ = "end of CoverTab[33522]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:366
func (f Form) FirstBoundary(b []byte) int {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:366
	_go_fuzz_dep_.CoverTab[33523]++
											return f.firstBoundary(inputBytes(b), len(b))
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:367
	// _ = "end of CoverTab[33523]"
}

func (f Form) firstBoundary(src input, nsrc int) int {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:370
	_go_fuzz_dep_.CoverTab[33524]++
											i := src.skipContinuationBytes(0)
											if i >= nsrc {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:372
		_go_fuzz_dep_.CoverTab[33526]++
												return -1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:373
		// _ = "end of CoverTab[33526]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:374
		_go_fuzz_dep_.CoverTab[33527]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:374
		// _ = "end of CoverTab[33527]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:374
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:374
	// _ = "end of CoverTab[33524]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:374
	_go_fuzz_dep_.CoverTab[33525]++
											fd := formTable[f]
											ss := streamSafe(0)

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:380
	for {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:380
		_go_fuzz_dep_.CoverTab[33528]++
												info := fd.info(src, i)
												if info.size == 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:382
			_go_fuzz_dep_.CoverTab[33531]++
													return -1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:383
			// _ = "end of CoverTab[33531]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:384
			_go_fuzz_dep_.CoverTab[33532]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:384
			// _ = "end of CoverTab[33532]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:384
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:384
		// _ = "end of CoverTab[33528]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:384
		_go_fuzz_dep_.CoverTab[33529]++
												if s := ss.next(info); s != ssSuccess {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:385
			_go_fuzz_dep_.CoverTab[33533]++
													return i
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:386
			// _ = "end of CoverTab[33533]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:387
			_go_fuzz_dep_.CoverTab[33534]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:387
			// _ = "end of CoverTab[33534]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:387
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:387
		// _ = "end of CoverTab[33529]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:387
		_go_fuzz_dep_.CoverTab[33530]++
												i += int(info.size)
												if i >= nsrc {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:389
			_go_fuzz_dep_.CoverTab[33535]++
													if !info.BoundaryAfter() && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:390
				_go_fuzz_dep_.CoverTab[33537]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:390
				return !ss.isMax()
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:390
				// _ = "end of CoverTab[33537]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:390
			}() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:390
				_go_fuzz_dep_.CoverTab[33538]++
														return -1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:391
				// _ = "end of CoverTab[33538]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:392
				_go_fuzz_dep_.CoverTab[33539]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:392
				// _ = "end of CoverTab[33539]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:392
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:392
			// _ = "end of CoverTab[33535]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:392
			_go_fuzz_dep_.CoverTab[33536]++
													return nsrc
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:393
			// _ = "end of CoverTab[33536]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:394
			_go_fuzz_dep_.CoverTab[33540]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:394
			// _ = "end of CoverTab[33540]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:394
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:394
		// _ = "end of CoverTab[33530]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:395
	// _ = "end of CoverTab[33525]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:400
func (f Form) FirstBoundaryInString(s string) int {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:400
	_go_fuzz_dep_.CoverTab[33541]++
											return f.firstBoundary(inputString(s), len(s))
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:401
	// _ = "end of CoverTab[33541]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:407
func (f Form) NextBoundary(b []byte, atEOF bool) int {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:407
	_go_fuzz_dep_.CoverTab[33542]++
											return f.nextBoundary(inputBytes(b), len(b), atEOF)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:408
	// _ = "end of CoverTab[33542]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:414
func (f Form) NextBoundaryInString(s string, atEOF bool) int {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:414
	_go_fuzz_dep_.CoverTab[33543]++
											return f.nextBoundary(inputString(s), len(s), atEOF)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:415
	// _ = "end of CoverTab[33543]"
}

func (f Form) nextBoundary(src input, nsrc int, atEOF bool) int {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:418
	_go_fuzz_dep_.CoverTab[33544]++
											if nsrc == 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:419
		_go_fuzz_dep_.CoverTab[33549]++
												if atEOF {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:420
			_go_fuzz_dep_.CoverTab[33551]++
													return 0
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:421
			// _ = "end of CoverTab[33551]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:422
			_go_fuzz_dep_.CoverTab[33552]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:422
			// _ = "end of CoverTab[33552]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:422
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:422
		// _ = "end of CoverTab[33549]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:422
		_go_fuzz_dep_.CoverTab[33550]++
												return -1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:423
		// _ = "end of CoverTab[33550]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:424
		_go_fuzz_dep_.CoverTab[33553]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:424
		// _ = "end of CoverTab[33553]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:424
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:424
	// _ = "end of CoverTab[33544]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:424
	_go_fuzz_dep_.CoverTab[33545]++
											fd := formTable[f]
											info := fd.info(src, 0)
											if info.size == 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:427
		_go_fuzz_dep_.CoverTab[33554]++
												if atEOF {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:428
			_go_fuzz_dep_.CoverTab[33556]++
													return 1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:429
			// _ = "end of CoverTab[33556]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:430
			_go_fuzz_dep_.CoverTab[33557]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:430
			// _ = "end of CoverTab[33557]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:430
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:430
		// _ = "end of CoverTab[33554]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:430
		_go_fuzz_dep_.CoverTab[33555]++
												return -1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:431
		// _ = "end of CoverTab[33555]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:432
		_go_fuzz_dep_.CoverTab[33558]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:432
		// _ = "end of CoverTab[33558]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:432
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:432
	// _ = "end of CoverTab[33545]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:432
	_go_fuzz_dep_.CoverTab[33546]++
											ss := streamSafe(0)
											ss.first(info)

											for i := int(info.size); i < nsrc; i += int(info.size) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:436
		_go_fuzz_dep_.CoverTab[33559]++
												info = fd.info(src, i)
												if info.size == 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:438
			_go_fuzz_dep_.CoverTab[33561]++
													if atEOF {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:439
				_go_fuzz_dep_.CoverTab[33563]++
														return i
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:440
				// _ = "end of CoverTab[33563]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:441
				_go_fuzz_dep_.CoverTab[33564]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:441
				// _ = "end of CoverTab[33564]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:441
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:441
			// _ = "end of CoverTab[33561]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:441
			_go_fuzz_dep_.CoverTab[33562]++
													return -1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:442
			// _ = "end of CoverTab[33562]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:443
			_go_fuzz_dep_.CoverTab[33565]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:443
			// _ = "end of CoverTab[33565]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:443
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:443
		// _ = "end of CoverTab[33559]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:443
		_go_fuzz_dep_.CoverTab[33560]++

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:446
		if s := ss.next(info); s != ssSuccess {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:446
			_go_fuzz_dep_.CoverTab[33566]++
													return i
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:447
			// _ = "end of CoverTab[33566]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:448
			_go_fuzz_dep_.CoverTab[33567]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:448
			// _ = "end of CoverTab[33567]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:448
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:448
		// _ = "end of CoverTab[33560]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:449
	// _ = "end of CoverTab[33546]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:449
	_go_fuzz_dep_.CoverTab[33547]++
											if !atEOF && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:450
		_go_fuzz_dep_.CoverTab[33568]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:450
		return !info.BoundaryAfter()
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:450
		// _ = "end of CoverTab[33568]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:450
	}() && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:450
		_go_fuzz_dep_.CoverTab[33569]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:450
		return !ss.isMax()
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:450
		// _ = "end of CoverTab[33569]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:450
	}() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:450
		_go_fuzz_dep_.CoverTab[33570]++
												return -1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:451
		// _ = "end of CoverTab[33570]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:452
		_go_fuzz_dep_.CoverTab[33571]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:452
		// _ = "end of CoverTab[33571]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:452
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:452
	// _ = "end of CoverTab[33547]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:452
	_go_fuzz_dep_.CoverTab[33548]++
											return nsrc
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:453
	// _ = "end of CoverTab[33548]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:458
func (f Form) LastBoundary(b []byte) int {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:458
	_go_fuzz_dep_.CoverTab[33572]++
											return lastBoundary(formTable[f], b)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:459
	// _ = "end of CoverTab[33572]"
}

func lastBoundary(fd *formInfo, b []byte) int {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:462
	_go_fuzz_dep_.CoverTab[33573]++
											i := len(b)
											info, p := lastRuneStart(fd, b)
											if p == -1 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:465
		_go_fuzz_dep_.CoverTab[33579]++
												return -1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:466
		// _ = "end of CoverTab[33579]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:467
		_go_fuzz_dep_.CoverTab[33580]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:467
		// _ = "end of CoverTab[33580]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:467
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:467
	// _ = "end of CoverTab[33573]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:467
	_go_fuzz_dep_.CoverTab[33574]++
											if info.size == 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:468
		_go_fuzz_dep_.CoverTab[33581]++
												if p == 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:469
			_go_fuzz_dep_.CoverTab[33583]++
													return -1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:470
			// _ = "end of CoverTab[33583]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:471
			_go_fuzz_dep_.CoverTab[33584]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:471
			// _ = "end of CoverTab[33584]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:471
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:471
		// _ = "end of CoverTab[33581]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:471
		_go_fuzz_dep_.CoverTab[33582]++
												i = p
												info, p = lastRuneStart(fd, b[:i])
												if p == -1 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:474
			_go_fuzz_dep_.CoverTab[33585]++
													return i
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:475
			// _ = "end of CoverTab[33585]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:476
			_go_fuzz_dep_.CoverTab[33586]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:476
			// _ = "end of CoverTab[33586]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:476
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:476
		// _ = "end of CoverTab[33582]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:477
		_go_fuzz_dep_.CoverTab[33587]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:477
		// _ = "end of CoverTab[33587]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:477
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:477
	// _ = "end of CoverTab[33574]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:477
	_go_fuzz_dep_.CoverTab[33575]++
											if p+int(info.size) != i {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:478
		_go_fuzz_dep_.CoverTab[33588]++
												return i
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:479
		// _ = "end of CoverTab[33588]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:480
		_go_fuzz_dep_.CoverTab[33589]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:480
		// _ = "end of CoverTab[33589]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:480
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:480
	// _ = "end of CoverTab[33575]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:480
	_go_fuzz_dep_.CoverTab[33576]++
											if info.BoundaryAfter() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:481
		_go_fuzz_dep_.CoverTab[33590]++
												return i
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:482
		// _ = "end of CoverTab[33590]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:483
		_go_fuzz_dep_.CoverTab[33591]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:483
		// _ = "end of CoverTab[33591]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:483
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:483
	// _ = "end of CoverTab[33576]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:483
	_go_fuzz_dep_.CoverTab[33577]++
											ss := streamSafe(0)
											v := ss.backwards(info)
											for i = p; i >= 0 && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:486
		_go_fuzz_dep_.CoverTab[33592]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:486
		return v != ssStarter
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:486
		// _ = "end of CoverTab[33592]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:486
	}(); i = p {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:486
		_go_fuzz_dep_.CoverTab[33593]++
												info, p = lastRuneStart(fd, b[:i])
												if v = ss.backwards(info); v == ssOverflow {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:488
			_go_fuzz_dep_.CoverTab[33595]++
													break
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:489
			// _ = "end of CoverTab[33595]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:490
			_go_fuzz_dep_.CoverTab[33596]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:490
			// _ = "end of CoverTab[33596]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:490
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:490
		// _ = "end of CoverTab[33593]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:490
		_go_fuzz_dep_.CoverTab[33594]++
												if p+int(info.size) != i {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:491
			_go_fuzz_dep_.CoverTab[33597]++
													if p == -1 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:492
				_go_fuzz_dep_.CoverTab[33599]++
														return -1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:493
				// _ = "end of CoverTab[33599]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:494
				_go_fuzz_dep_.CoverTab[33600]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:494
				// _ = "end of CoverTab[33600]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:494
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:494
			// _ = "end of CoverTab[33597]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:494
			_go_fuzz_dep_.CoverTab[33598]++
													return i
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:495
			// _ = "end of CoverTab[33598]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:496
			_go_fuzz_dep_.CoverTab[33601]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:496
			// _ = "end of CoverTab[33601]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:496
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:496
		// _ = "end of CoverTab[33594]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:497
	// _ = "end of CoverTab[33577]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:497
	_go_fuzz_dep_.CoverTab[33578]++
											return i
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:498
	// _ = "end of CoverTab[33578]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:504
func decomposeSegment(rb *reorderBuffer, sp int, atEOF bool) int {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:504
	_go_fuzz_dep_.CoverTab[33602]++

											info := rb.f.info(rb.src, sp)
											if info.size == 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:507
		_go_fuzz_dep_.CoverTab[33608]++
												return 0
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:508
		// _ = "end of CoverTab[33608]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:509
		_go_fuzz_dep_.CoverTab[33609]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:509
		// _ = "end of CoverTab[33609]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:509
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:509
	// _ = "end of CoverTab[33602]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:509
	_go_fuzz_dep_.CoverTab[33603]++
											if s := rb.ss.next(info); s == ssStarter {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:510
		_go_fuzz_dep_.CoverTab[33610]++

												if rb.nrune > 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:512
			_go_fuzz_dep_.CoverTab[33611]++
													goto end
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:513
			// _ = "end of CoverTab[33611]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:514
			_go_fuzz_dep_.CoverTab[33612]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:514
			// _ = "end of CoverTab[33612]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:514
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:514
		// _ = "end of CoverTab[33610]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:515
		_go_fuzz_dep_.CoverTab[33613]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:515
		if s == ssOverflow {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:515
			_go_fuzz_dep_.CoverTab[33614]++
													rb.insertCGJ()
													goto end
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:517
			// _ = "end of CoverTab[33614]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:518
			_go_fuzz_dep_.CoverTab[33615]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:518
			// _ = "end of CoverTab[33615]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:518
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:518
		// _ = "end of CoverTab[33613]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:518
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:518
	// _ = "end of CoverTab[33603]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:518
	_go_fuzz_dep_.CoverTab[33604]++
											if err := rb.insertFlush(rb.src, sp, info); err != iSuccess {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:519
		_go_fuzz_dep_.CoverTab[33616]++
												return int(err)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:520
		// _ = "end of CoverTab[33616]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:521
		_go_fuzz_dep_.CoverTab[33617]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:521
		// _ = "end of CoverTab[33617]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:521
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:521
	// _ = "end of CoverTab[33604]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:521
	_go_fuzz_dep_.CoverTab[33605]++
											for {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:522
		_go_fuzz_dep_.CoverTab[33618]++
												sp += int(info.size)
												if sp >= rb.nsrc {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:524
			_go_fuzz_dep_.CoverTab[33622]++
													if !atEOF && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:525
				_go_fuzz_dep_.CoverTab[33624]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:525
				return !info.BoundaryAfter()
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:525
				// _ = "end of CoverTab[33624]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:525
			}() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:525
				_go_fuzz_dep_.CoverTab[33625]++
														return int(iShortSrc)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:526
				// _ = "end of CoverTab[33625]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:527
				_go_fuzz_dep_.CoverTab[33626]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:527
				// _ = "end of CoverTab[33626]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:527
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:527
			// _ = "end of CoverTab[33622]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:527
			_go_fuzz_dep_.CoverTab[33623]++
													break
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:528
			// _ = "end of CoverTab[33623]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:529
			_go_fuzz_dep_.CoverTab[33627]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:529
			// _ = "end of CoverTab[33627]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:529
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:529
		// _ = "end of CoverTab[33618]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:529
		_go_fuzz_dep_.CoverTab[33619]++
												info = rb.f.info(rb.src, sp)
												if info.size == 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:531
			_go_fuzz_dep_.CoverTab[33628]++
													if !atEOF {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:532
				_go_fuzz_dep_.CoverTab[33630]++
														return int(iShortSrc)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:533
				// _ = "end of CoverTab[33630]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:534
				_go_fuzz_dep_.CoverTab[33631]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:534
				// _ = "end of CoverTab[33631]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:534
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:534
			// _ = "end of CoverTab[33628]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:534
			_go_fuzz_dep_.CoverTab[33629]++
													break
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:535
			// _ = "end of CoverTab[33629]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:536
			_go_fuzz_dep_.CoverTab[33632]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:536
			// _ = "end of CoverTab[33632]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:536
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:536
		// _ = "end of CoverTab[33619]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:536
		_go_fuzz_dep_.CoverTab[33620]++
												if s := rb.ss.next(info); s == ssStarter {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:537
			_go_fuzz_dep_.CoverTab[33633]++
													break
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:538
			// _ = "end of CoverTab[33633]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:539
			_go_fuzz_dep_.CoverTab[33634]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:539
			if s == ssOverflow {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:539
				_go_fuzz_dep_.CoverTab[33635]++
														rb.insertCGJ()
														break
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:541
				// _ = "end of CoverTab[33635]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:542
				_go_fuzz_dep_.CoverTab[33636]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:542
				// _ = "end of CoverTab[33636]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:542
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:542
			// _ = "end of CoverTab[33634]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:542
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:542
		// _ = "end of CoverTab[33620]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:542
		_go_fuzz_dep_.CoverTab[33621]++
												if err := rb.insertFlush(rb.src, sp, info); err != iSuccess {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:543
			_go_fuzz_dep_.CoverTab[33637]++
													return int(err)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:544
			// _ = "end of CoverTab[33637]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:545
			_go_fuzz_dep_.CoverTab[33638]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:545
			// _ = "end of CoverTab[33638]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:545
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:545
		// _ = "end of CoverTab[33621]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:546
	// _ = "end of CoverTab[33605]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:546
	_go_fuzz_dep_.CoverTab[33606]++
end:
	if !rb.doFlush() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:548
		_go_fuzz_dep_.CoverTab[33639]++
												return int(iShortDst)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:549
		// _ = "end of CoverTab[33639]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:550
		_go_fuzz_dep_.CoverTab[33640]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:550
		// _ = "end of CoverTab[33640]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:550
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:550
	// _ = "end of CoverTab[33606]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:550
	_go_fuzz_dep_.CoverTab[33607]++
											return sp
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:551
	// _ = "end of CoverTab[33607]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:556
func lastRuneStart(fd *formInfo, buf []byte) (Properties, int) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:556
	_go_fuzz_dep_.CoverTab[33641]++
											p := len(buf) - 1
											for ; p >= 0 && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:558
		_go_fuzz_dep_.CoverTab[33644]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:558
		return !utf8.RuneStart(buf[p])
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:558
		// _ = "end of CoverTab[33644]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:558
	}(); p-- {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:558
		_go_fuzz_dep_.CoverTab[33645]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:558
		// _ = "end of CoverTab[33645]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:559
	// _ = "end of CoverTab[33641]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:559
	_go_fuzz_dep_.CoverTab[33642]++
											if p < 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:560
		_go_fuzz_dep_.CoverTab[33646]++
												return Properties{}, -1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:561
		// _ = "end of CoverTab[33646]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:562
		_go_fuzz_dep_.CoverTab[33647]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:562
		// _ = "end of CoverTab[33647]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:562
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:562
	// _ = "end of CoverTab[33642]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:562
	_go_fuzz_dep_.CoverTab[33643]++
											return fd.info(inputBytes(buf), p), p
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:563
	// _ = "end of CoverTab[33643]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:568
func decomposeToLastBoundary(rb *reorderBuffer) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:568
	_go_fuzz_dep_.CoverTab[33648]++
											fd := &rb.f
											info, i := lastRuneStart(fd, rb.out)
											if int(info.size) != len(rb.out)-i {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:571
		_go_fuzz_dep_.CoverTab[33652]++

												return
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:573
		// _ = "end of CoverTab[33652]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:574
		_go_fuzz_dep_.CoverTab[33653]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:574
		// _ = "end of CoverTab[33653]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:574
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:574
	// _ = "end of CoverTab[33648]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:574
	_go_fuzz_dep_.CoverTab[33649]++
											if info.BoundaryAfter() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:575
		_go_fuzz_dep_.CoverTab[33654]++
												return
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:576
		// _ = "end of CoverTab[33654]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:577
		_go_fuzz_dep_.CoverTab[33655]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:577
		// _ = "end of CoverTab[33655]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:577
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:577
	// _ = "end of CoverTab[33649]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:577
	_go_fuzz_dep_.CoverTab[33650]++
											var add [maxNonStarters + 1]Properties
											padd := 0
											ss := streamSafe(0)
											p := len(rb.out)
											for {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:582
		_go_fuzz_dep_.CoverTab[33656]++
												add[padd] = info
												v := ss.backwards(info)
												if v == ssOverflow {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:585
			_go_fuzz_dep_.CoverTab[33659]++

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:588
			break
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:588
			// _ = "end of CoverTab[33659]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:589
			_go_fuzz_dep_.CoverTab[33660]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:589
			// _ = "end of CoverTab[33660]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:589
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:589
		// _ = "end of CoverTab[33656]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:589
		_go_fuzz_dep_.CoverTab[33657]++
												padd++
												p -= int(info.size)
												if v == ssStarter || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:592
			_go_fuzz_dep_.CoverTab[33661]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:592
			return p < 0
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:592
			// _ = "end of CoverTab[33661]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:592
		}() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:592
			_go_fuzz_dep_.CoverTab[33662]++
													break
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:593
			// _ = "end of CoverTab[33662]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:594
			_go_fuzz_dep_.CoverTab[33663]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:594
			// _ = "end of CoverTab[33663]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:594
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:594
		// _ = "end of CoverTab[33657]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:594
		_go_fuzz_dep_.CoverTab[33658]++
												info, i = lastRuneStart(fd, rb.out[:p])
												if int(info.size) != p-i {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:596
			_go_fuzz_dep_.CoverTab[33664]++
													break
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:597
			// _ = "end of CoverTab[33664]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:598
			_go_fuzz_dep_.CoverTab[33665]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:598
			// _ = "end of CoverTab[33665]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:598
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:598
		// _ = "end of CoverTab[33658]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:599
	// _ = "end of CoverTab[33650]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:599
	_go_fuzz_dep_.CoverTab[33651]++
											rb.ss = ss

											var buf [maxBufferSize * utf8.UTFMax]byte
											cp := buf[:copy(buf[:], rb.out[p:])]
											rb.out = rb.out[:p]
											for padd--; padd >= 0; padd-- {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:605
		_go_fuzz_dep_.CoverTab[33666]++
												info = add[padd]
												rb.insertUnsafe(inputBytes(cp), 0, info)
												cp = cp[info.size:]
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:608
		// _ = "end of CoverTab[33666]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:609
	// _ = "end of CoverTab[33651]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:610
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/normalize.go:610
var _ = _go_fuzz_dep_.CoverTab
