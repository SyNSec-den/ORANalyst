// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:5
package runes

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:5
)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:5
)

import (
	"unicode/utf8"

	"golang.org/x/text/transform"
)

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:28
// If returns a transformer that applies tIn to consecutive runes for which
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:28
// s.Contains(r) and tNotIn to consecutive runes for which !s.Contains(r). Reset
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:28
// is called on tIn and tNotIn at the start of each run. A Nop transformer will
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:28
// substitute a nil value passed to tIn or tNotIn. Invalid UTF-8 is translated
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:28
// to RuneError to determine which transformer to apply, but is passed as is to
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:28
// the respective transformer.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:34
func If(s Set, tIn, tNotIn transform.Transformer) Transformer {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:34
	_go_fuzz_dep_.CoverTab[116918]++
										if tIn == nil && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:35
		_go_fuzz_dep_.CoverTab[116924]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:35
		return tNotIn == nil
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:35
		// _ = "end of CoverTab[116924]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:35
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:35
		_go_fuzz_dep_.CoverTab[116925]++
											return Transformer{transform.Nop}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:36
		// _ = "end of CoverTab[116925]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:37
		_go_fuzz_dep_.CoverTab[116926]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:37
		// _ = "end of CoverTab[116926]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:37
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:37
	// _ = "end of CoverTab[116918]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:37
	_go_fuzz_dep_.CoverTab[116919]++
										if tIn == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:38
		_go_fuzz_dep_.CoverTab[116927]++
											tIn = transform.Nop
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:39
		// _ = "end of CoverTab[116927]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:40
		_go_fuzz_dep_.CoverTab[116928]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:40
		// _ = "end of CoverTab[116928]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:40
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:40
	// _ = "end of CoverTab[116919]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:40
	_go_fuzz_dep_.CoverTab[116920]++
										if tNotIn == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:41
		_go_fuzz_dep_.CoverTab[116929]++
											tNotIn = transform.Nop
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:42
		// _ = "end of CoverTab[116929]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:43
		_go_fuzz_dep_.CoverTab[116930]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:43
		// _ = "end of CoverTab[116930]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:43
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:43
	// _ = "end of CoverTab[116920]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:43
	_go_fuzz_dep_.CoverTab[116921]++
										sIn, ok := tIn.(transform.SpanningTransformer)
										if !ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:45
		_go_fuzz_dep_.CoverTab[116931]++
											sIn = dummySpan{tIn}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:46
		// _ = "end of CoverTab[116931]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:47
		_go_fuzz_dep_.CoverTab[116932]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:47
		// _ = "end of CoverTab[116932]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:47
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:47
	// _ = "end of CoverTab[116921]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:47
	_go_fuzz_dep_.CoverTab[116922]++
										sNotIn, ok := tNotIn.(transform.SpanningTransformer)
										if !ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:49
		_go_fuzz_dep_.CoverTab[116933]++
											sNotIn = dummySpan{tNotIn}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:50
		// _ = "end of CoverTab[116933]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:51
		_go_fuzz_dep_.CoverTab[116934]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:51
		// _ = "end of CoverTab[116934]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:51
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:51
	// _ = "end of CoverTab[116922]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:51
	_go_fuzz_dep_.CoverTab[116923]++

										a := &cond{
		tIn:	sIn,
		tNotIn:	sNotIn,
		f:	s.Contains,
	}
										a.Reset()
										return Transformer{a}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:59
	// _ = "end of CoverTab[116923]"
}

type dummySpan struct{ transform.Transformer }

func (d dummySpan) Span(src []byte, atEOF bool) (n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:64
	_go_fuzz_dep_.CoverTab[116935]++
										return 0, transform.ErrEndOfSpan
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:65
	// _ = "end of CoverTab[116935]"
}

type cond struct {
	tIn, tNotIn	transform.SpanningTransformer
	f		func(rune) bool
	check		func(rune) bool			// current check to perform
	t		transform.SpanningTransformer	// current transformer to use
}

// Reset implements transform.Transformer.
func (t *cond) Reset() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:76
	_go_fuzz_dep_.CoverTab[116936]++
										t.check = t.is
										t.t = t.tIn
										t.t.Reset()
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:79
	// _ = "end of CoverTab[116936]"
}

func (t *cond) is(r rune) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:82
	_go_fuzz_dep_.CoverTab[116937]++
										if t.f(r) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:83
		_go_fuzz_dep_.CoverTab[116939]++
											return true
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:84
		// _ = "end of CoverTab[116939]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:85
		_go_fuzz_dep_.CoverTab[116940]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:85
		// _ = "end of CoverTab[116940]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:85
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:85
	// _ = "end of CoverTab[116937]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:85
	_go_fuzz_dep_.CoverTab[116938]++
										t.check = t.isNot
										t.t = t.tNotIn
										t.tNotIn.Reset()
										return false
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:89
	// _ = "end of CoverTab[116938]"
}

func (t *cond) isNot(r rune) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:92
	_go_fuzz_dep_.CoverTab[116941]++
										if !t.f(r) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:93
		_go_fuzz_dep_.CoverTab[116943]++
											return true
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:94
		// _ = "end of CoverTab[116943]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:95
		_go_fuzz_dep_.CoverTab[116944]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:95
		// _ = "end of CoverTab[116944]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:95
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:95
	// _ = "end of CoverTab[116941]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:95
	_go_fuzz_dep_.CoverTab[116942]++
										t.check = t.is
										t.t = t.tIn
										t.tIn.Reset()
										return false
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:99
	// _ = "end of CoverTab[116942]"
}

// This implementation of Span doesn't help all too much, but it needs to be
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:102
// there to satisfy this package's Transformer interface.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:102
// TODO: there are certainly room for improvements, though. For example, if
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:102
// t.t == transform.Nop (which will a common occurrence) it will save a bundle
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:102
// to special-case that loop.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:107
func (t *cond) Span(src []byte, atEOF bool) (n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:107
	_go_fuzz_dep_.CoverTab[116945]++
										p := 0
										for n < len(src) && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:109
		_go_fuzz_dep_.CoverTab[116947]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:109
		return err == nil
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:109
		// _ = "end of CoverTab[116947]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:109
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:109
		_go_fuzz_dep_.CoverTab[116948]++
		// Don't process too much at a time as the Spanner that will be
		// called on this block may terminate early.
		const maxChunk = 4096
		max := len(src)
		if v := n + maxChunk; v < max {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:114
			_go_fuzz_dep_.CoverTab[116952]++
												max = v
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:115
			// _ = "end of CoverTab[116952]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:116
			_go_fuzz_dep_.CoverTab[116953]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:116
			// _ = "end of CoverTab[116953]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:116
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:116
		// _ = "end of CoverTab[116948]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:116
		_go_fuzz_dep_.CoverTab[116949]++
											atEnd := false
											size := 0
											current := t.t
											for ; p < max; p += size {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:120
			_go_fuzz_dep_.CoverTab[116954]++
												r := rune(src[p])
												if r < utf8.RuneSelf {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:122
				_go_fuzz_dep_.CoverTab[116956]++
													size = 1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:123
				// _ = "end of CoverTab[116956]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:124
				_go_fuzz_dep_.CoverTab[116957]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:124
				if r, size = utf8.DecodeRune(src[p:]); size == 1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:124
					_go_fuzz_dep_.CoverTab[116958]++
														if !atEOF && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:125
						_go_fuzz_dep_.CoverTab[116959]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:125
						return !utf8.FullRune(src[p:])
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:125
						// _ = "end of CoverTab[116959]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:125
					}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:125
						_go_fuzz_dep_.CoverTab[116960]++
															err = transform.ErrShortSrc
															break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:127
						// _ = "end of CoverTab[116960]"
					} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:128
						_go_fuzz_dep_.CoverTab[116961]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:128
						// _ = "end of CoverTab[116961]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:128
					}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:128
					// _ = "end of CoverTab[116958]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:129
					_go_fuzz_dep_.CoverTab[116962]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:129
					// _ = "end of CoverTab[116962]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:129
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:129
				// _ = "end of CoverTab[116957]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:129
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:129
			// _ = "end of CoverTab[116954]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:129
			_go_fuzz_dep_.CoverTab[116955]++
												if !t.check(r) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:130
				_go_fuzz_dep_.CoverTab[116963]++

													atEnd = true
													break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:133
				// _ = "end of CoverTab[116963]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:134
				_go_fuzz_dep_.CoverTab[116964]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:134
				// _ = "end of CoverTab[116964]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:134
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:134
			// _ = "end of CoverTab[116955]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:135
		// _ = "end of CoverTab[116949]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:135
		_go_fuzz_dep_.CoverTab[116950]++
											n2, err2 := current.Span(src[n:p], atEnd || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:136
			_go_fuzz_dep_.CoverTab[116965]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:136
			return (atEOF && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:136
				_go_fuzz_dep_.CoverTab[116966]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:136
				return p == len(src)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:136
				// _ = "end of CoverTab[116966]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:136
			}())
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:136
			// _ = "end of CoverTab[116965]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:136
		}())
											n += n2
											if err2 != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:138
			_go_fuzz_dep_.CoverTab[116967]++
												return n, err2
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:139
			// _ = "end of CoverTab[116967]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:140
			_go_fuzz_dep_.CoverTab[116968]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:140
			// _ = "end of CoverTab[116968]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:140
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:140
		// _ = "end of CoverTab[116950]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:140
		_go_fuzz_dep_.CoverTab[116951]++

											p = n + size
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:142
		// _ = "end of CoverTab[116951]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:143
	// _ = "end of CoverTab[116945]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:143
	_go_fuzz_dep_.CoverTab[116946]++
										return n, err
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:144
	// _ = "end of CoverTab[116946]"
}

func (t *cond) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:147
	_go_fuzz_dep_.CoverTab[116969]++
										p := 0
										for nSrc < len(src) && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:149
		_go_fuzz_dep_.CoverTab[116971]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:149
		return err == nil
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:149
		// _ = "end of CoverTab[116971]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:149
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:149
		_go_fuzz_dep_.CoverTab[116972]++
		// Don't process too much at a time, as the work might be wasted if the
		// destination buffer isn't large enough to hold the result or a
		// transform returns an error early.
		const maxChunk = 4096
		max := len(src)
		if n := nSrc + maxChunk; n < len(src) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:155
			_go_fuzz_dep_.CoverTab[116976]++
												max = n
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:156
			// _ = "end of CoverTab[116976]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:157
			_go_fuzz_dep_.CoverTab[116977]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:157
			// _ = "end of CoverTab[116977]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:157
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:157
		// _ = "end of CoverTab[116972]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:157
		_go_fuzz_dep_.CoverTab[116973]++
											atEnd := false
											size := 0
											current := t.t
											for ; p < max; p += size {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:161
			_go_fuzz_dep_.CoverTab[116978]++
												r := rune(src[p])
												if r < utf8.RuneSelf {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:163
				_go_fuzz_dep_.CoverTab[116980]++
													size = 1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:164
				// _ = "end of CoverTab[116980]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:165
				_go_fuzz_dep_.CoverTab[116981]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:165
				if r, size = utf8.DecodeRune(src[p:]); size == 1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:165
					_go_fuzz_dep_.CoverTab[116982]++
														if !atEOF && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:166
						_go_fuzz_dep_.CoverTab[116983]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:166
						return !utf8.FullRune(src[p:])
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:166
						// _ = "end of CoverTab[116983]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:166
					}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:166
						_go_fuzz_dep_.CoverTab[116984]++
															err = transform.ErrShortSrc
															break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:168
						// _ = "end of CoverTab[116984]"
					} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:169
						_go_fuzz_dep_.CoverTab[116985]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:169
						// _ = "end of CoverTab[116985]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:169
					}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:169
					// _ = "end of CoverTab[116982]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:170
					_go_fuzz_dep_.CoverTab[116986]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:170
					// _ = "end of CoverTab[116986]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:170
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:170
				// _ = "end of CoverTab[116981]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:170
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:170
			// _ = "end of CoverTab[116978]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:170
			_go_fuzz_dep_.CoverTab[116979]++
												if !t.check(r) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:171
				_go_fuzz_dep_.CoverTab[116987]++

													atEnd = true
													break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:174
				// _ = "end of CoverTab[116987]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:175
				_go_fuzz_dep_.CoverTab[116988]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:175
				// _ = "end of CoverTab[116988]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:175
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:175
			// _ = "end of CoverTab[116979]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:176
		// _ = "end of CoverTab[116973]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:176
		_go_fuzz_dep_.CoverTab[116974]++
											nDst2, nSrc2, err2 := current.Transform(dst[nDst:], src[nSrc:p], atEnd || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:177
			_go_fuzz_dep_.CoverTab[116989]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:177
			return (atEOF && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:177
				_go_fuzz_dep_.CoverTab[116990]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:177
				return p == len(src)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:177
				// _ = "end of CoverTab[116990]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:177
			}())
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:177
			// _ = "end of CoverTab[116989]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:177
		}())
											nDst += nDst2
											nSrc += nSrc2
											if err2 != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:180
			_go_fuzz_dep_.CoverTab[116991]++
												return nDst, nSrc, err2
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:181
			// _ = "end of CoverTab[116991]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:182
			_go_fuzz_dep_.CoverTab[116992]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:182
			// _ = "end of CoverTab[116992]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:182
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:182
		// _ = "end of CoverTab[116974]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:182
		_go_fuzz_dep_.CoverTab[116975]++

											p = nSrc + size
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:184
		// _ = "end of CoverTab[116975]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:185
	// _ = "end of CoverTab[116969]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:185
	_go_fuzz_dep_.CoverTab[116970]++
										return nDst, nSrc, err
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:186
	// _ = "end of CoverTab[116970]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:187
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/cond.go:187
var _ = _go_fuzz_dep_.CoverTab
