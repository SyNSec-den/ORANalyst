// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:5
// Package runes provide transforms for UTF-8 encoded text.
package runes

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:6
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:6
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:6
)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:6
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:6
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:6
)

import (
	"unicode"
	"unicode/utf8"

	"golang.org/x/text/transform"
)

// A Set is a collection of runes.
type Set interface {
	// Contains returns true if r is contained in the set.
	Contains(r rune) bool
}

type setFunc func(rune) bool

func (s setFunc) Contains(r rune) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:23
	_go_fuzz_dep_.CoverTab[116993]++
										return s(r)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:24
	// _ = "end of CoverTab[116993]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:30
// In creates a Set with a Contains method that returns true for all runes in
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:30
// the given RangeTable.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:32
func In(rt *unicode.RangeTable) Set {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:32
	_go_fuzz_dep_.CoverTab[116994]++
										return setFunc(func(r rune) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:33
		_go_fuzz_dep_.CoverTab[116995]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:33
		return unicode.Is(rt, r)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:33
		// _ = "end of CoverTab[116995]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:33
	})
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:33
	// _ = "end of CoverTab[116994]"
}

// NotIn creates a Set with a Contains method that returns true for all runes not
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:36
// in the given RangeTable.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:38
func NotIn(rt *unicode.RangeTable) Set {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:38
	_go_fuzz_dep_.CoverTab[116996]++
										return setFunc(func(r rune) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:39
		_go_fuzz_dep_.CoverTab[116997]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:39
		return !unicode.Is(rt, r)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:39
		// _ = "end of CoverTab[116997]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:39
	})
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:39
	// _ = "end of CoverTab[116996]"
}

// Predicate creates a Set with a Contains method that returns f(r).
func Predicate(f func(rune) bool) Set {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:43
	_go_fuzz_dep_.CoverTab[116998]++
										return setFunc(f)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:44
	// _ = "end of CoverTab[116998]"
}

// Transformer implements the transform.Transformer interface.
type Transformer struct {
	t transform.SpanningTransformer
}

func (t Transformer) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:52
	_go_fuzz_dep_.CoverTab[116999]++
										return t.t.Transform(dst, src, atEOF)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:53
	// _ = "end of CoverTab[116999]"
}

func (t Transformer) Span(b []byte, atEOF bool) (n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:56
	_go_fuzz_dep_.CoverTab[117000]++
										return t.t.Span(b, atEOF)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:57
	// _ = "end of CoverTab[117000]"
}

func (t Transformer) Reset() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:60
	_go_fuzz_dep_.CoverTab[117001]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:60
	t.t.Reset()
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:60
	// _ = "end of CoverTab[117001]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:60
}

// Bytes returns a new byte slice with the result of converting b using t.  It
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:62
// calls Reset on t. It returns nil if any error was found. This can only happen
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:62
// if an error-producing Transformer is passed to If.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:65
func (t Transformer) Bytes(b []byte) []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:65
	_go_fuzz_dep_.CoverTab[117002]++
										b, _, err := transform.Bytes(t, b)
										if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:67
		_go_fuzz_dep_.CoverTab[117004]++
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:68
		// _ = "end of CoverTab[117004]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:69
		_go_fuzz_dep_.CoverTab[117005]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:69
		// _ = "end of CoverTab[117005]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:69
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:69
	// _ = "end of CoverTab[117002]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:69
	_go_fuzz_dep_.CoverTab[117003]++
										return b
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:70
	// _ = "end of CoverTab[117003]"
}

// String returns a string with the result of converting s using t. It calls
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:73
// Reset on t. It returns the empty string if any error was found. This can only
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:73
// happen if an error-producing Transformer is passed to If.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:76
func (t Transformer) String(s string) string {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:76
	_go_fuzz_dep_.CoverTab[117006]++
										s, _, err := transform.String(t, s)
										if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:78
		_go_fuzz_dep_.CoverTab[117008]++
											return ""
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:79
		// _ = "end of CoverTab[117008]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:80
		_go_fuzz_dep_.CoverTab[117009]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:80
		// _ = "end of CoverTab[117009]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:80
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:80
	// _ = "end of CoverTab[117006]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:80
	_go_fuzz_dep_.CoverTab[117007]++
										return s
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:81
	// _ = "end of CoverTab[117007]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:89
const runeErrorString = string(utf8.RuneError)

// Remove returns a Transformer that removes runes r for which s.Contains(r).
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:91
// Illegal input bytes are replaced by RuneError before being passed to f.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:93
func Remove(s Set) Transformer {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:93
	_go_fuzz_dep_.CoverTab[117010]++
										if f, ok := s.(setFunc); ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:94
		_go_fuzz_dep_.CoverTab[117012]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:98
		return Transformer{remove(f)}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:98
		// _ = "end of CoverTab[117012]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:99
		_go_fuzz_dep_.CoverTab[117013]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:99
		// _ = "end of CoverTab[117013]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:99
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:99
	// _ = "end of CoverTab[117010]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:99
	_go_fuzz_dep_.CoverTab[117011]++
										return Transformer{remove(s.Contains)}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:100
	// _ = "end of CoverTab[117011]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:105
type remove func(r rune) bool

func (remove) Reset()	{ _go_fuzz_dep_.CoverTab[117014]++; // _ = "end of CoverTab[117014]" }

// Span implements transform.Spanner.
func (t remove) Span(src []byte, atEOF bool) (n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:110
	_go_fuzz_dep_.CoverTab[117015]++
										for r, size := rune(0), 0; n < len(src); {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:111
		_go_fuzz_dep_.CoverTab[117017]++
											if r = rune(src[n]); r < utf8.RuneSelf {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:112
			_go_fuzz_dep_.CoverTab[117020]++
												size = 1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:113
			// _ = "end of CoverTab[117020]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:114
			_go_fuzz_dep_.CoverTab[117021]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:114
			if r, size = utf8.DecodeRune(src[n:]); size == 1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:114
				_go_fuzz_dep_.CoverTab[117022]++

													if !atEOF && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:116
					_go_fuzz_dep_.CoverTab[117024]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:116
					return !utf8.FullRune(src[n:])
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:116
					// _ = "end of CoverTab[117024]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:116
				}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:116
					_go_fuzz_dep_.CoverTab[117025]++
														err = transform.ErrShortSrc
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:117
					// _ = "end of CoverTab[117025]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:118
					_go_fuzz_dep_.CoverTab[117026]++
														err = transform.ErrEndOfSpan
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:119
					// _ = "end of CoverTab[117026]"
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:120
				// _ = "end of CoverTab[117022]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:120
				_go_fuzz_dep_.CoverTab[117023]++
													break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:121
				// _ = "end of CoverTab[117023]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:122
				_go_fuzz_dep_.CoverTab[117027]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:122
				// _ = "end of CoverTab[117027]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:122
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:122
			// _ = "end of CoverTab[117021]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:122
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:122
		// _ = "end of CoverTab[117017]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:122
		_go_fuzz_dep_.CoverTab[117018]++
											if t(r) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:123
			_go_fuzz_dep_.CoverTab[117028]++
												err = transform.ErrEndOfSpan
												break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:125
			// _ = "end of CoverTab[117028]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:126
			_go_fuzz_dep_.CoverTab[117029]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:126
			// _ = "end of CoverTab[117029]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:126
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:126
		// _ = "end of CoverTab[117018]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:126
		_go_fuzz_dep_.CoverTab[117019]++
											n += size
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:127
		// _ = "end of CoverTab[117019]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:128
	// _ = "end of CoverTab[117015]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:128
	_go_fuzz_dep_.CoverTab[117016]++
										return
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:129
	// _ = "end of CoverTab[117016]"
}

// Transform implements transform.Transformer.
func (t remove) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:133
	_go_fuzz_dep_.CoverTab[117030]++
										for r, size := rune(0), 0; nSrc < len(src); {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:134
		_go_fuzz_dep_.CoverTab[117032]++
											if r = rune(src[nSrc]); r < utf8.RuneSelf {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:135
			_go_fuzz_dep_.CoverTab[117036]++
												size = 1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:136
			// _ = "end of CoverTab[117036]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:137
			_go_fuzz_dep_.CoverTab[117037]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:137
			if r, size = utf8.DecodeRune(src[nSrc:]); size == 1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:137
				_go_fuzz_dep_.CoverTab[117038]++

													if !atEOF && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:139
					_go_fuzz_dep_.CoverTab[117041]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:139
					return !utf8.FullRune(src[nSrc:])
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:139
					// _ = "end of CoverTab[117041]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:139
				}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:139
					_go_fuzz_dep_.CoverTab[117042]++
														err = transform.ErrShortSrc
														break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:141
					// _ = "end of CoverTab[117042]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:142
					_go_fuzz_dep_.CoverTab[117043]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:142
					// _ = "end of CoverTab[117043]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:142
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:142
				// _ = "end of CoverTab[117038]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:142
				_go_fuzz_dep_.CoverTab[117039]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:147
				if !t(utf8.RuneError) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:147
					_go_fuzz_dep_.CoverTab[117044]++
														if nDst+3 > len(dst) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:148
						_go_fuzz_dep_.CoverTab[117046]++
															err = transform.ErrShortDst
															break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:150
						// _ = "end of CoverTab[117046]"
					} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:151
						_go_fuzz_dep_.CoverTab[117047]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:151
						// _ = "end of CoverTab[117047]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:151
					}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:151
					// _ = "end of CoverTab[117044]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:151
					_go_fuzz_dep_.CoverTab[117045]++
														dst[nDst+0] = runeErrorString[0]
														dst[nDst+1] = runeErrorString[1]
														dst[nDst+2] = runeErrorString[2]
														nDst += 3
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:155
					// _ = "end of CoverTab[117045]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:156
					_go_fuzz_dep_.CoverTab[117048]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:156
					// _ = "end of CoverTab[117048]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:156
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:156
				// _ = "end of CoverTab[117039]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:156
				_go_fuzz_dep_.CoverTab[117040]++
													nSrc++
													continue
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:158
				// _ = "end of CoverTab[117040]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:159
				_go_fuzz_dep_.CoverTab[117049]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:159
				// _ = "end of CoverTab[117049]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:159
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:159
			// _ = "end of CoverTab[117037]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:159
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:159
		// _ = "end of CoverTab[117032]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:159
		_go_fuzz_dep_.CoverTab[117033]++
											if t(r) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:160
			_go_fuzz_dep_.CoverTab[117050]++
												nSrc += size
												continue
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:162
			// _ = "end of CoverTab[117050]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:163
			_go_fuzz_dep_.CoverTab[117051]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:163
			// _ = "end of CoverTab[117051]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:163
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:163
		// _ = "end of CoverTab[117033]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:163
		_go_fuzz_dep_.CoverTab[117034]++
											if nDst+size > len(dst) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:164
			_go_fuzz_dep_.CoverTab[117052]++
												err = transform.ErrShortDst
												break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:166
			// _ = "end of CoverTab[117052]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:167
			_go_fuzz_dep_.CoverTab[117053]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:167
			// _ = "end of CoverTab[117053]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:167
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:167
		// _ = "end of CoverTab[117034]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:167
		_go_fuzz_dep_.CoverTab[117035]++
											for i := 0; i < size; i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:168
			_go_fuzz_dep_.CoverTab[117054]++
												dst[nDst] = src[nSrc]
												nDst++
												nSrc++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:171
			// _ = "end of CoverTab[117054]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:172
		// _ = "end of CoverTab[117035]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:173
	// _ = "end of CoverTab[117030]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:173
	_go_fuzz_dep_.CoverTab[117031]++
										return
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:174
	// _ = "end of CoverTab[117031]"
}

// Map returns a Transformer that maps the runes in the input using the given
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:177
// mapping. Illegal bytes in the input are converted to utf8.RuneError before
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:177
// being passed to the mapping func.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:180
func Map(mapping func(rune) rune) Transformer {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:180
	_go_fuzz_dep_.CoverTab[117055]++
										return Transformer{mapper(mapping)}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:181
	// _ = "end of CoverTab[117055]"
}

type mapper func(rune) rune

func (mapper) Reset()	{ _go_fuzz_dep_.CoverTab[117056]++; // _ = "end of CoverTab[117056]" }

// Span implements transform.Spanner.
func (t mapper) Span(src []byte, atEOF bool) (n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:189
	_go_fuzz_dep_.CoverTab[117057]++
										for r, size := rune(0), 0; n < len(src); n += size {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:190
		_go_fuzz_dep_.CoverTab[117059]++
											if r = rune(src[n]); r < utf8.RuneSelf {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:191
			_go_fuzz_dep_.CoverTab[117061]++
												size = 1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:192
			// _ = "end of CoverTab[117061]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:193
			_go_fuzz_dep_.CoverTab[117062]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:193
			if r, size = utf8.DecodeRune(src[n:]); size == 1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:193
				_go_fuzz_dep_.CoverTab[117063]++

													if !atEOF && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:195
					_go_fuzz_dep_.CoverTab[117065]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:195
					return !utf8.FullRune(src[n:])
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:195
					// _ = "end of CoverTab[117065]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:195
				}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:195
					_go_fuzz_dep_.CoverTab[117066]++
														err = transform.ErrShortSrc
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:196
					// _ = "end of CoverTab[117066]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:197
					_go_fuzz_dep_.CoverTab[117067]++
														err = transform.ErrEndOfSpan
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:198
					// _ = "end of CoverTab[117067]"
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:199
				// _ = "end of CoverTab[117063]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:199
				_go_fuzz_dep_.CoverTab[117064]++
													break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:200
				// _ = "end of CoverTab[117064]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:201
				_go_fuzz_dep_.CoverTab[117068]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:201
				// _ = "end of CoverTab[117068]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:201
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:201
			// _ = "end of CoverTab[117062]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:201
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:201
		// _ = "end of CoverTab[117059]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:201
		_go_fuzz_dep_.CoverTab[117060]++
											if t(r) != r {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:202
			_go_fuzz_dep_.CoverTab[117069]++
												err = transform.ErrEndOfSpan
												break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:204
			// _ = "end of CoverTab[117069]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:205
			_go_fuzz_dep_.CoverTab[117070]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:205
			// _ = "end of CoverTab[117070]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:205
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:205
		// _ = "end of CoverTab[117060]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:206
	// _ = "end of CoverTab[117057]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:206
	_go_fuzz_dep_.CoverTab[117058]++
										return n, err
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:207
	// _ = "end of CoverTab[117058]"
}

// Transform implements transform.Transformer.
func (t mapper) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:211
	_go_fuzz_dep_.CoverTab[117071]++
										var replacement rune
										var b [utf8.UTFMax]byte

										for r, size := rune(0), 0; nSrc < len(src); {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:215
		_go_fuzz_dep_.CoverTab[117073]++
											if r = rune(src[nSrc]); r < utf8.RuneSelf {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:216
			_go_fuzz_dep_.CoverTab[117077]++
												if replacement = t(r); replacement < utf8.RuneSelf {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:217
				_go_fuzz_dep_.CoverTab[117079]++
													if nDst == len(dst) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:218
					_go_fuzz_dep_.CoverTab[117081]++
														err = transform.ErrShortDst
														break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:220
					// _ = "end of CoverTab[117081]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:221
					_go_fuzz_dep_.CoverTab[117082]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:221
					// _ = "end of CoverTab[117082]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:221
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:221
				// _ = "end of CoverTab[117079]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:221
				_go_fuzz_dep_.CoverTab[117080]++
													dst[nDst] = byte(replacement)
													nDst++
													nSrc++
													continue
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:225
				// _ = "end of CoverTab[117080]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:226
				_go_fuzz_dep_.CoverTab[117083]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:226
				// _ = "end of CoverTab[117083]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:226
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:226
			// _ = "end of CoverTab[117077]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:226
			_go_fuzz_dep_.CoverTab[117078]++
												size = 1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:227
			// _ = "end of CoverTab[117078]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:228
			_go_fuzz_dep_.CoverTab[117084]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:228
			if r, size = utf8.DecodeRune(src[nSrc:]); size == 1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:228
				_go_fuzz_dep_.CoverTab[117085]++

													if !atEOF && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:230
					_go_fuzz_dep_.CoverTab[117087]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:230
					return !utf8.FullRune(src[nSrc:])
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:230
					// _ = "end of CoverTab[117087]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:230
				}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:230
					_go_fuzz_dep_.CoverTab[117088]++
														err = transform.ErrShortSrc
														break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:232
					// _ = "end of CoverTab[117088]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:233
					_go_fuzz_dep_.CoverTab[117089]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:233
					// _ = "end of CoverTab[117089]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:233
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:233
				// _ = "end of CoverTab[117085]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:233
				_go_fuzz_dep_.CoverTab[117086]++

													if replacement = t(utf8.RuneError); replacement == utf8.RuneError {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:235
					_go_fuzz_dep_.CoverTab[117090]++
														if nDst+3 > len(dst) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:236
						_go_fuzz_dep_.CoverTab[117092]++
															err = transform.ErrShortDst
															break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:238
						// _ = "end of CoverTab[117092]"
					} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:239
						_go_fuzz_dep_.CoverTab[117093]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:239
						// _ = "end of CoverTab[117093]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:239
					}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:239
					// _ = "end of CoverTab[117090]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:239
					_go_fuzz_dep_.CoverTab[117091]++
														dst[nDst+0] = runeErrorString[0]
														dst[nDst+1] = runeErrorString[1]
														dst[nDst+2] = runeErrorString[2]
														nDst += 3
														nSrc++
														continue
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:245
					// _ = "end of CoverTab[117091]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:246
					_go_fuzz_dep_.CoverTab[117094]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:246
					// _ = "end of CoverTab[117094]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:246
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:246
				// _ = "end of CoverTab[117086]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:247
				_go_fuzz_dep_.CoverTab[117095]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:247
				if replacement = t(r); replacement == r {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:247
					_go_fuzz_dep_.CoverTab[117096]++
														if nDst+size > len(dst) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:248
						_go_fuzz_dep_.CoverTab[117099]++
															err = transform.ErrShortDst
															break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:250
						// _ = "end of CoverTab[117099]"
					} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:251
						_go_fuzz_dep_.CoverTab[117100]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:251
						// _ = "end of CoverTab[117100]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:251
					}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:251
					// _ = "end of CoverTab[117096]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:251
					_go_fuzz_dep_.CoverTab[117097]++
														for i := 0; i < size; i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:252
						_go_fuzz_dep_.CoverTab[117101]++
															dst[nDst] = src[nSrc]
															nDst++
															nSrc++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:255
						// _ = "end of CoverTab[117101]"
					}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:256
					// _ = "end of CoverTab[117097]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:256
					_go_fuzz_dep_.CoverTab[117098]++
														continue
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:257
					// _ = "end of CoverTab[117098]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:258
					_go_fuzz_dep_.CoverTab[117102]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:258
					// _ = "end of CoverTab[117102]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:258
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:258
				// _ = "end of CoverTab[117095]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:258
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:258
			// _ = "end of CoverTab[117084]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:258
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:258
		// _ = "end of CoverTab[117073]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:258
		_go_fuzz_dep_.CoverTab[117074]++

											n := utf8.EncodeRune(b[:], replacement)

											if nDst+n > len(dst) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:262
			_go_fuzz_dep_.CoverTab[117103]++
												err = transform.ErrShortDst
												break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:264
			// _ = "end of CoverTab[117103]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:265
			_go_fuzz_dep_.CoverTab[117104]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:265
			// _ = "end of CoverTab[117104]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:265
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:265
		// _ = "end of CoverTab[117074]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:265
		_go_fuzz_dep_.CoverTab[117075]++
											for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:266
			_go_fuzz_dep_.CoverTab[117105]++
												dst[nDst] = b[i]
												nDst++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:268
			// _ = "end of CoverTab[117105]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:269
		// _ = "end of CoverTab[117075]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:269
		_go_fuzz_dep_.CoverTab[117076]++
											nSrc += size
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:270
		// _ = "end of CoverTab[117076]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:271
	// _ = "end of CoverTab[117071]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:271
	_go_fuzz_dep_.CoverTab[117072]++
										return
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:272
	// _ = "end of CoverTab[117072]"
}

// ReplaceIllFormed returns a transformer that replaces all input bytes that are
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:275
// not part of a well-formed UTF-8 code sequence with utf8.RuneError.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:277
func ReplaceIllFormed() Transformer {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:277
	_go_fuzz_dep_.CoverTab[117106]++
										return Transformer{&replaceIllFormed{}}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:278
	// _ = "end of CoverTab[117106]"
}

type replaceIllFormed struct{ transform.NopResetter }

func (t replaceIllFormed) Span(src []byte, atEOF bool) (n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:283
	_go_fuzz_dep_.CoverTab[117107]++
										for n < len(src) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:284
		_go_fuzz_dep_.CoverTab[117109]++

											if src[n] < utf8.RuneSelf {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:286
			_go_fuzz_dep_.CoverTab[117113]++
												n++
												continue
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:288
			// _ = "end of CoverTab[117113]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:289
			_go_fuzz_dep_.CoverTab[117114]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:289
			// _ = "end of CoverTab[117114]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:289
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:289
		// _ = "end of CoverTab[117109]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:289
		_go_fuzz_dep_.CoverTab[117110]++

											r, size := utf8.DecodeRune(src[n:])

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:294
		if r != utf8.RuneError || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:294
			_go_fuzz_dep_.CoverTab[117115]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:294
			return size != 1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:294
			// _ = "end of CoverTab[117115]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:294
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:294
			_go_fuzz_dep_.CoverTab[117116]++
												n += size
												continue
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:296
			// _ = "end of CoverTab[117116]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:297
			_go_fuzz_dep_.CoverTab[117117]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:297
			// _ = "end of CoverTab[117117]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:297
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:297
		// _ = "end of CoverTab[117110]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:297
		_go_fuzz_dep_.CoverTab[117111]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:300
		if !atEOF && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:300
			_go_fuzz_dep_.CoverTab[117118]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:300
			return !utf8.FullRune(src[n:])
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:300
			// _ = "end of CoverTab[117118]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:300
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:300
			_go_fuzz_dep_.CoverTab[117119]++
												err = transform.ErrShortSrc
												break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:302
			// _ = "end of CoverTab[117119]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:303
			_go_fuzz_dep_.CoverTab[117120]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:303
			// _ = "end of CoverTab[117120]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:303
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:303
		// _ = "end of CoverTab[117111]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:303
		_go_fuzz_dep_.CoverTab[117112]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:306
		err = transform.ErrEndOfSpan
											break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:307
		// _ = "end of CoverTab[117112]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:308
	// _ = "end of CoverTab[117107]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:308
	_go_fuzz_dep_.CoverTab[117108]++
										return n, err
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:309
	// _ = "end of CoverTab[117108]"
}

func (t replaceIllFormed) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:312
	_go_fuzz_dep_.CoverTab[117121]++
										for nSrc < len(src) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:313
		_go_fuzz_dep_.CoverTab[117123]++

											if r := src[nSrc]; r < utf8.RuneSelf {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:315
			_go_fuzz_dep_.CoverTab[117128]++
												if nDst == len(dst) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:316
				_go_fuzz_dep_.CoverTab[117130]++
													err = transform.ErrShortDst
													break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:318
				// _ = "end of CoverTab[117130]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:319
				_go_fuzz_dep_.CoverTab[117131]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:319
				// _ = "end of CoverTab[117131]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:319
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:319
			// _ = "end of CoverTab[117128]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:319
			_go_fuzz_dep_.CoverTab[117129]++
												dst[nDst] = r
												nDst++
												nSrc++
												continue
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:323
			// _ = "end of CoverTab[117129]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:324
			_go_fuzz_dep_.CoverTab[117132]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:324
			// _ = "end of CoverTab[117132]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:324
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:324
		// _ = "end of CoverTab[117123]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:324
		_go_fuzz_dep_.CoverTab[117124]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:327
		if _, size := utf8.DecodeRune(src[nSrc:]); size != 1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:327
			_go_fuzz_dep_.CoverTab[117133]++
												if size != copy(dst[nDst:], src[nSrc:nSrc+size]) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:328
				_go_fuzz_dep_.CoverTab[117135]++
													err = transform.ErrShortDst
													break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:330
				// _ = "end of CoverTab[117135]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:331
				_go_fuzz_dep_.CoverTab[117136]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:331
				// _ = "end of CoverTab[117136]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:331
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:331
			// _ = "end of CoverTab[117133]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:331
			_go_fuzz_dep_.CoverTab[117134]++
												nDst += size
												nSrc += size
												continue
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:334
			// _ = "end of CoverTab[117134]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:335
			_go_fuzz_dep_.CoverTab[117137]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:335
			// _ = "end of CoverTab[117137]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:335
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:335
		// _ = "end of CoverTab[117124]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:335
		_go_fuzz_dep_.CoverTab[117125]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:338
		if !atEOF && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:338
			_go_fuzz_dep_.CoverTab[117138]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:338
			return !utf8.FullRune(src[nSrc:])
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:338
			// _ = "end of CoverTab[117138]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:338
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:338
			_go_fuzz_dep_.CoverTab[117139]++
												err = transform.ErrShortSrc
												break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:340
			// _ = "end of CoverTab[117139]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:341
			_go_fuzz_dep_.CoverTab[117140]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:341
			// _ = "end of CoverTab[117140]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:341
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:341
		// _ = "end of CoverTab[117125]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:341
		_go_fuzz_dep_.CoverTab[117126]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:344
		if nDst+3 > len(dst) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:344
			_go_fuzz_dep_.CoverTab[117141]++
												err = transform.ErrShortDst
												break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:346
			// _ = "end of CoverTab[117141]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:347
			_go_fuzz_dep_.CoverTab[117142]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:347
			// _ = "end of CoverTab[117142]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:347
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:347
		// _ = "end of CoverTab[117126]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:347
		_go_fuzz_dep_.CoverTab[117127]++
											dst[nDst+0] = runeErrorString[0]
											dst[nDst+1] = runeErrorString[1]
											dst[nDst+2] = runeErrorString[2]
											nDst += 3
											nSrc++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:352
		// _ = "end of CoverTab[117127]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:353
	// _ = "end of CoverTab[117121]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:353
	_go_fuzz_dep_.CoverTab[117122]++
										return nDst, nSrc, err
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:354
	// _ = "end of CoverTab[117122]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:355
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/runes/runes.go:355
var _ = _go_fuzz_dep_.CoverTab
