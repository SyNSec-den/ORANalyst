// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/regexp/syntax/simplify.go:5
package syntax

//line /usr/local/go/src/regexp/syntax/simplify.go:5
import (
//line /usr/local/go/src/regexp/syntax/simplify.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/regexp/syntax/simplify.go:5
)
//line /usr/local/go/src/regexp/syntax/simplify.go:5
import (
//line /usr/local/go/src/regexp/syntax/simplify.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/regexp/syntax/simplify.go:5
)

// Simplify returns a regexp equivalent to re but without counted repetitions
//line /usr/local/go/src/regexp/syntax/simplify.go:7
// and with various other simplifications, such as rewriting /(?:a+)+/ to /a+/.
//line /usr/local/go/src/regexp/syntax/simplify.go:7
// The resulting regexp will execute correctly but its string representation
//line /usr/local/go/src/regexp/syntax/simplify.go:7
// will not produce the same parse tree, because capturing parentheses
//line /usr/local/go/src/regexp/syntax/simplify.go:7
// may have been duplicated or removed. For example, the simplified form
//line /usr/local/go/src/regexp/syntax/simplify.go:7
// for /(x){1,2}/ is /(x)(x)?/ but both parentheses capture as $1.
//line /usr/local/go/src/regexp/syntax/simplify.go:7
// The returned regexp may share structure with or be the original.
//line /usr/local/go/src/regexp/syntax/simplify.go:14
func (re *Regexp) Simplify() *Regexp {
//line /usr/local/go/src/regexp/syntax/simplify.go:14
	_go_fuzz_dep_.CoverTab[64344]++
							if re == nil {
//line /usr/local/go/src/regexp/syntax/simplify.go:15
		_go_fuzz_dep_.CoverTab[64347]++
								return nil
//line /usr/local/go/src/regexp/syntax/simplify.go:16
		// _ = "end of CoverTab[64347]"
	} else {
//line /usr/local/go/src/regexp/syntax/simplify.go:17
		_go_fuzz_dep_.CoverTab[64348]++
//line /usr/local/go/src/regexp/syntax/simplify.go:17
		// _ = "end of CoverTab[64348]"
//line /usr/local/go/src/regexp/syntax/simplify.go:17
	}
//line /usr/local/go/src/regexp/syntax/simplify.go:17
	// _ = "end of CoverTab[64344]"
//line /usr/local/go/src/regexp/syntax/simplify.go:17
	_go_fuzz_dep_.CoverTab[64345]++
							switch re.Op {
	case OpCapture, OpConcat, OpAlternate:
//line /usr/local/go/src/regexp/syntax/simplify.go:19
		_go_fuzz_dep_.CoverTab[64349]++

								nre := re
								for i, sub := range re.Sub {
//line /usr/local/go/src/regexp/syntax/simplify.go:22
			_go_fuzz_dep_.CoverTab[64360]++
									nsub := sub.Simplify()
									if nre == re && func() bool {
//line /usr/local/go/src/regexp/syntax/simplify.go:24
				_go_fuzz_dep_.CoverTab[64362]++
//line /usr/local/go/src/regexp/syntax/simplify.go:24
				return nsub != sub
//line /usr/local/go/src/regexp/syntax/simplify.go:24
				// _ = "end of CoverTab[64362]"
//line /usr/local/go/src/regexp/syntax/simplify.go:24
			}() {
//line /usr/local/go/src/regexp/syntax/simplify.go:24
				_go_fuzz_dep_.CoverTab[64363]++

										nre = new(Regexp)
										*nre = *re
										nre.Rune = nil
										nre.Sub = append(nre.Sub0[:0], re.Sub[:i]...)
//line /usr/local/go/src/regexp/syntax/simplify.go:29
				// _ = "end of CoverTab[64363]"
			} else {
//line /usr/local/go/src/regexp/syntax/simplify.go:30
				_go_fuzz_dep_.CoverTab[64364]++
//line /usr/local/go/src/regexp/syntax/simplify.go:30
				// _ = "end of CoverTab[64364]"
//line /usr/local/go/src/regexp/syntax/simplify.go:30
			}
//line /usr/local/go/src/regexp/syntax/simplify.go:30
			// _ = "end of CoverTab[64360]"
//line /usr/local/go/src/regexp/syntax/simplify.go:30
			_go_fuzz_dep_.CoverTab[64361]++
									if nre != re {
//line /usr/local/go/src/regexp/syntax/simplify.go:31
				_go_fuzz_dep_.CoverTab[64365]++
										nre.Sub = append(nre.Sub, nsub)
//line /usr/local/go/src/regexp/syntax/simplify.go:32
				// _ = "end of CoverTab[64365]"
			} else {
//line /usr/local/go/src/regexp/syntax/simplify.go:33
				_go_fuzz_dep_.CoverTab[64366]++
//line /usr/local/go/src/regexp/syntax/simplify.go:33
				// _ = "end of CoverTab[64366]"
//line /usr/local/go/src/regexp/syntax/simplify.go:33
			}
//line /usr/local/go/src/regexp/syntax/simplify.go:33
			// _ = "end of CoverTab[64361]"
		}
//line /usr/local/go/src/regexp/syntax/simplify.go:34
		// _ = "end of CoverTab[64349]"
//line /usr/local/go/src/regexp/syntax/simplify.go:34
		_go_fuzz_dep_.CoverTab[64350]++
								return nre
//line /usr/local/go/src/regexp/syntax/simplify.go:35
		// _ = "end of CoverTab[64350]"

	case OpStar, OpPlus, OpQuest:
//line /usr/local/go/src/regexp/syntax/simplify.go:37
		_go_fuzz_dep_.CoverTab[64351]++
								sub := re.Sub[0].Simplify()
								return simplify1(re.Op, re.Flags, sub, re)
//line /usr/local/go/src/regexp/syntax/simplify.go:39
		// _ = "end of CoverTab[64351]"

	case OpRepeat:
//line /usr/local/go/src/regexp/syntax/simplify.go:41
		_go_fuzz_dep_.CoverTab[64352]++

//line /usr/local/go/src/regexp/syntax/simplify.go:44
		if re.Min == 0 && func() bool {
//line /usr/local/go/src/regexp/syntax/simplify.go:44
			_go_fuzz_dep_.CoverTab[64367]++
//line /usr/local/go/src/regexp/syntax/simplify.go:44
			return re.Max == 0
//line /usr/local/go/src/regexp/syntax/simplify.go:44
			// _ = "end of CoverTab[64367]"
//line /usr/local/go/src/regexp/syntax/simplify.go:44
		}() {
//line /usr/local/go/src/regexp/syntax/simplify.go:44
			_go_fuzz_dep_.CoverTab[64368]++
									return &Regexp{Op: OpEmptyMatch}
//line /usr/local/go/src/regexp/syntax/simplify.go:45
			// _ = "end of CoverTab[64368]"
		} else {
//line /usr/local/go/src/regexp/syntax/simplify.go:46
			_go_fuzz_dep_.CoverTab[64369]++
//line /usr/local/go/src/regexp/syntax/simplify.go:46
			// _ = "end of CoverTab[64369]"
//line /usr/local/go/src/regexp/syntax/simplify.go:46
		}
//line /usr/local/go/src/regexp/syntax/simplify.go:46
		// _ = "end of CoverTab[64352]"
//line /usr/local/go/src/regexp/syntax/simplify.go:46
		_go_fuzz_dep_.CoverTab[64353]++

//line /usr/local/go/src/regexp/syntax/simplify.go:49
		sub := re.Sub[0].Simplify()

//line /usr/local/go/src/regexp/syntax/simplify.go:52
		if re.Max == -1 {
//line /usr/local/go/src/regexp/syntax/simplify.go:52
			_go_fuzz_dep_.CoverTab[64370]++

									if re.Min == 0 {
//line /usr/local/go/src/regexp/syntax/simplify.go:54
				_go_fuzz_dep_.CoverTab[64374]++
										return simplify1(OpStar, re.Flags, sub, nil)
//line /usr/local/go/src/regexp/syntax/simplify.go:55
				// _ = "end of CoverTab[64374]"
			} else {
//line /usr/local/go/src/regexp/syntax/simplify.go:56
				_go_fuzz_dep_.CoverTab[64375]++
//line /usr/local/go/src/regexp/syntax/simplify.go:56
				// _ = "end of CoverTab[64375]"
//line /usr/local/go/src/regexp/syntax/simplify.go:56
			}
//line /usr/local/go/src/regexp/syntax/simplify.go:56
			// _ = "end of CoverTab[64370]"
//line /usr/local/go/src/regexp/syntax/simplify.go:56
			_go_fuzz_dep_.CoverTab[64371]++

//line /usr/local/go/src/regexp/syntax/simplify.go:59
			if re.Min == 1 {
//line /usr/local/go/src/regexp/syntax/simplify.go:59
				_go_fuzz_dep_.CoverTab[64376]++
										return simplify1(OpPlus, re.Flags, sub, nil)
//line /usr/local/go/src/regexp/syntax/simplify.go:60
				// _ = "end of CoverTab[64376]"
			} else {
//line /usr/local/go/src/regexp/syntax/simplify.go:61
				_go_fuzz_dep_.CoverTab[64377]++
//line /usr/local/go/src/regexp/syntax/simplify.go:61
				// _ = "end of CoverTab[64377]"
//line /usr/local/go/src/regexp/syntax/simplify.go:61
			}
//line /usr/local/go/src/regexp/syntax/simplify.go:61
			// _ = "end of CoverTab[64371]"
//line /usr/local/go/src/regexp/syntax/simplify.go:61
			_go_fuzz_dep_.CoverTab[64372]++

//line /usr/local/go/src/regexp/syntax/simplify.go:64
			nre := &Regexp{Op: OpConcat}
			nre.Sub = nre.Sub0[:0]
			for i := 0; i < re.Min-1; i++ {
//line /usr/local/go/src/regexp/syntax/simplify.go:66
				_go_fuzz_dep_.CoverTab[64378]++
										nre.Sub = append(nre.Sub, sub)
//line /usr/local/go/src/regexp/syntax/simplify.go:67
				// _ = "end of CoverTab[64378]"
			}
//line /usr/local/go/src/regexp/syntax/simplify.go:68
			// _ = "end of CoverTab[64372]"
//line /usr/local/go/src/regexp/syntax/simplify.go:68
			_go_fuzz_dep_.CoverTab[64373]++
									nre.Sub = append(nre.Sub, simplify1(OpPlus, re.Flags, sub, nil))
									return nre
//line /usr/local/go/src/regexp/syntax/simplify.go:70
			// _ = "end of CoverTab[64373]"
		} else {
//line /usr/local/go/src/regexp/syntax/simplify.go:71
			_go_fuzz_dep_.CoverTab[64379]++
//line /usr/local/go/src/regexp/syntax/simplify.go:71
			// _ = "end of CoverTab[64379]"
//line /usr/local/go/src/regexp/syntax/simplify.go:71
		}
//line /usr/local/go/src/regexp/syntax/simplify.go:71
		// _ = "end of CoverTab[64353]"
//line /usr/local/go/src/regexp/syntax/simplify.go:71
		_go_fuzz_dep_.CoverTab[64354]++

//line /usr/local/go/src/regexp/syntax/simplify.go:76
		if re.Min == 1 && func() bool {
//line /usr/local/go/src/regexp/syntax/simplify.go:76
			_go_fuzz_dep_.CoverTab[64380]++
//line /usr/local/go/src/regexp/syntax/simplify.go:76
			return re.Max == 1
//line /usr/local/go/src/regexp/syntax/simplify.go:76
			// _ = "end of CoverTab[64380]"
//line /usr/local/go/src/regexp/syntax/simplify.go:76
		}() {
//line /usr/local/go/src/regexp/syntax/simplify.go:76
			_go_fuzz_dep_.CoverTab[64381]++
									return sub
//line /usr/local/go/src/regexp/syntax/simplify.go:77
			// _ = "end of CoverTab[64381]"
		} else {
//line /usr/local/go/src/regexp/syntax/simplify.go:78
			_go_fuzz_dep_.CoverTab[64382]++
//line /usr/local/go/src/regexp/syntax/simplify.go:78
			// _ = "end of CoverTab[64382]"
//line /usr/local/go/src/regexp/syntax/simplify.go:78
		}
//line /usr/local/go/src/regexp/syntax/simplify.go:78
		// _ = "end of CoverTab[64354]"
//line /usr/local/go/src/regexp/syntax/simplify.go:78
		_go_fuzz_dep_.CoverTab[64355]++

//line /usr/local/go/src/regexp/syntax/simplify.go:84
		// Build leading prefix: xx.
								var prefix *Regexp
								if re.Min > 0 {
//line /usr/local/go/src/regexp/syntax/simplify.go:86
			_go_fuzz_dep_.CoverTab[64383]++
									prefix = &Regexp{Op: OpConcat}
									prefix.Sub = prefix.Sub0[:0]
									for i := 0; i < re.Min; i++ {
//line /usr/local/go/src/regexp/syntax/simplify.go:89
				_go_fuzz_dep_.CoverTab[64384]++
										prefix.Sub = append(prefix.Sub, sub)
//line /usr/local/go/src/regexp/syntax/simplify.go:90
				// _ = "end of CoverTab[64384]"
			}
//line /usr/local/go/src/regexp/syntax/simplify.go:91
			// _ = "end of CoverTab[64383]"
		} else {
//line /usr/local/go/src/regexp/syntax/simplify.go:92
			_go_fuzz_dep_.CoverTab[64385]++
//line /usr/local/go/src/regexp/syntax/simplify.go:92
			// _ = "end of CoverTab[64385]"
//line /usr/local/go/src/regexp/syntax/simplify.go:92
		}
//line /usr/local/go/src/regexp/syntax/simplify.go:92
		// _ = "end of CoverTab[64355]"
//line /usr/local/go/src/regexp/syntax/simplify.go:92
		_go_fuzz_dep_.CoverTab[64356]++

//line /usr/local/go/src/regexp/syntax/simplify.go:95
		if re.Max > re.Min {
//line /usr/local/go/src/regexp/syntax/simplify.go:95
			_go_fuzz_dep_.CoverTab[64386]++
									suffix := simplify1(OpQuest, re.Flags, sub, nil)
									for i := re.Min + 1; i < re.Max; i++ {
//line /usr/local/go/src/regexp/syntax/simplify.go:97
				_go_fuzz_dep_.CoverTab[64389]++
										nre2 := &Regexp{Op: OpConcat}
										nre2.Sub = append(nre2.Sub0[:0], sub, suffix)
										suffix = simplify1(OpQuest, re.Flags, nre2, nil)
//line /usr/local/go/src/regexp/syntax/simplify.go:100
				// _ = "end of CoverTab[64389]"
			}
//line /usr/local/go/src/regexp/syntax/simplify.go:101
			// _ = "end of CoverTab[64386]"
//line /usr/local/go/src/regexp/syntax/simplify.go:101
			_go_fuzz_dep_.CoverTab[64387]++
									if prefix == nil {
//line /usr/local/go/src/regexp/syntax/simplify.go:102
				_go_fuzz_dep_.CoverTab[64390]++
										return suffix
//line /usr/local/go/src/regexp/syntax/simplify.go:103
				// _ = "end of CoverTab[64390]"
			} else {
//line /usr/local/go/src/regexp/syntax/simplify.go:104
				_go_fuzz_dep_.CoverTab[64391]++
//line /usr/local/go/src/regexp/syntax/simplify.go:104
				// _ = "end of CoverTab[64391]"
//line /usr/local/go/src/regexp/syntax/simplify.go:104
			}
//line /usr/local/go/src/regexp/syntax/simplify.go:104
			// _ = "end of CoverTab[64387]"
//line /usr/local/go/src/regexp/syntax/simplify.go:104
			_go_fuzz_dep_.CoverTab[64388]++
									prefix.Sub = append(prefix.Sub, suffix)
//line /usr/local/go/src/regexp/syntax/simplify.go:105
			// _ = "end of CoverTab[64388]"
		} else {
//line /usr/local/go/src/regexp/syntax/simplify.go:106
			_go_fuzz_dep_.CoverTab[64392]++
//line /usr/local/go/src/regexp/syntax/simplify.go:106
			// _ = "end of CoverTab[64392]"
//line /usr/local/go/src/regexp/syntax/simplify.go:106
		}
//line /usr/local/go/src/regexp/syntax/simplify.go:106
		// _ = "end of CoverTab[64356]"
//line /usr/local/go/src/regexp/syntax/simplify.go:106
		_go_fuzz_dep_.CoverTab[64357]++
								if prefix != nil {
//line /usr/local/go/src/regexp/syntax/simplify.go:107
			_go_fuzz_dep_.CoverTab[64393]++
									return prefix
//line /usr/local/go/src/regexp/syntax/simplify.go:108
			// _ = "end of CoverTab[64393]"
		} else {
//line /usr/local/go/src/regexp/syntax/simplify.go:109
			_go_fuzz_dep_.CoverTab[64394]++
//line /usr/local/go/src/regexp/syntax/simplify.go:109
			// _ = "end of CoverTab[64394]"
//line /usr/local/go/src/regexp/syntax/simplify.go:109
		}
//line /usr/local/go/src/regexp/syntax/simplify.go:109
		// _ = "end of CoverTab[64357]"
//line /usr/local/go/src/regexp/syntax/simplify.go:109
		_go_fuzz_dep_.CoverTab[64358]++

//line /usr/local/go/src/regexp/syntax/simplify.go:113
		return &Regexp{Op: OpNoMatch}
//line /usr/local/go/src/regexp/syntax/simplify.go:113
		// _ = "end of CoverTab[64358]"
//line /usr/local/go/src/regexp/syntax/simplify.go:113
	default:
//line /usr/local/go/src/regexp/syntax/simplify.go:113
		_go_fuzz_dep_.CoverTab[64359]++
//line /usr/local/go/src/regexp/syntax/simplify.go:113
		// _ = "end of CoverTab[64359]"
	}
//line /usr/local/go/src/regexp/syntax/simplify.go:114
	// _ = "end of CoverTab[64345]"
//line /usr/local/go/src/regexp/syntax/simplify.go:114
	_go_fuzz_dep_.CoverTab[64346]++

							return re
//line /usr/local/go/src/regexp/syntax/simplify.go:116
	// _ = "end of CoverTab[64346]"
}

// simplify1 implements Simplify for the unary OpStar,
//line /usr/local/go/src/regexp/syntax/simplify.go:119
// OpPlus, and OpQuest operators. It returns the simple regexp
//line /usr/local/go/src/regexp/syntax/simplify.go:119
// equivalent to
//line /usr/local/go/src/regexp/syntax/simplify.go:119
//
//line /usr/local/go/src/regexp/syntax/simplify.go:119
//	Regexp{Op: op, Flags: flags, Sub: {sub}}
//line /usr/local/go/src/regexp/syntax/simplify.go:119
//
//line /usr/local/go/src/regexp/syntax/simplify.go:119
// under the assumption that sub is already simple, and
//line /usr/local/go/src/regexp/syntax/simplify.go:119
// without first allocating that structure. If the regexp
//line /usr/local/go/src/regexp/syntax/simplify.go:119
// to be returned turns out to be equivalent to re, simplify1
//line /usr/local/go/src/regexp/syntax/simplify.go:119
// returns re instead.
//line /usr/local/go/src/regexp/syntax/simplify.go:119
//
//line /usr/local/go/src/regexp/syntax/simplify.go:119
// simplify1 is factored out of Simplify because the implementation
//line /usr/local/go/src/regexp/syntax/simplify.go:119
// for other operators generates these unary expressions.
//line /usr/local/go/src/regexp/syntax/simplify.go:119
// Letting them call simplify1 makes sure the expressions they
//line /usr/local/go/src/regexp/syntax/simplify.go:119
// generate are simple.
//line /usr/local/go/src/regexp/syntax/simplify.go:134
func simplify1(op Op, flags Flags, sub, re *Regexp) *Regexp {
//line /usr/local/go/src/regexp/syntax/simplify.go:134
	_go_fuzz_dep_.CoverTab[64395]++

//line /usr/local/go/src/regexp/syntax/simplify.go:137
	if sub.Op == OpEmptyMatch {
//line /usr/local/go/src/regexp/syntax/simplify.go:137
		_go_fuzz_dep_.CoverTab[64399]++
								return sub
//line /usr/local/go/src/regexp/syntax/simplify.go:138
		// _ = "end of CoverTab[64399]"
	} else {
//line /usr/local/go/src/regexp/syntax/simplify.go:139
		_go_fuzz_dep_.CoverTab[64400]++
//line /usr/local/go/src/regexp/syntax/simplify.go:139
		// _ = "end of CoverTab[64400]"
//line /usr/local/go/src/regexp/syntax/simplify.go:139
	}
//line /usr/local/go/src/regexp/syntax/simplify.go:139
	// _ = "end of CoverTab[64395]"
//line /usr/local/go/src/regexp/syntax/simplify.go:139
	_go_fuzz_dep_.CoverTab[64396]++

							if op == sub.Op && func() bool {
//line /usr/local/go/src/regexp/syntax/simplify.go:141
		_go_fuzz_dep_.CoverTab[64401]++
//line /usr/local/go/src/regexp/syntax/simplify.go:141
		return flags&NonGreedy == sub.Flags&NonGreedy
//line /usr/local/go/src/regexp/syntax/simplify.go:141
		// _ = "end of CoverTab[64401]"
//line /usr/local/go/src/regexp/syntax/simplify.go:141
	}() {
//line /usr/local/go/src/regexp/syntax/simplify.go:141
		_go_fuzz_dep_.CoverTab[64402]++
								return sub
//line /usr/local/go/src/regexp/syntax/simplify.go:142
		// _ = "end of CoverTab[64402]"
	} else {
//line /usr/local/go/src/regexp/syntax/simplify.go:143
		_go_fuzz_dep_.CoverTab[64403]++
//line /usr/local/go/src/regexp/syntax/simplify.go:143
		// _ = "end of CoverTab[64403]"
//line /usr/local/go/src/regexp/syntax/simplify.go:143
	}
//line /usr/local/go/src/regexp/syntax/simplify.go:143
	// _ = "end of CoverTab[64396]"
//line /usr/local/go/src/regexp/syntax/simplify.go:143
	_go_fuzz_dep_.CoverTab[64397]++
							if re != nil && func() bool {
//line /usr/local/go/src/regexp/syntax/simplify.go:144
		_go_fuzz_dep_.CoverTab[64404]++
//line /usr/local/go/src/regexp/syntax/simplify.go:144
		return re.Op == op
//line /usr/local/go/src/regexp/syntax/simplify.go:144
		// _ = "end of CoverTab[64404]"
//line /usr/local/go/src/regexp/syntax/simplify.go:144
	}() && func() bool {
//line /usr/local/go/src/regexp/syntax/simplify.go:144
		_go_fuzz_dep_.CoverTab[64405]++
//line /usr/local/go/src/regexp/syntax/simplify.go:144
		return re.Flags&NonGreedy == flags&NonGreedy
//line /usr/local/go/src/regexp/syntax/simplify.go:144
		// _ = "end of CoverTab[64405]"
//line /usr/local/go/src/regexp/syntax/simplify.go:144
	}() && func() bool {
//line /usr/local/go/src/regexp/syntax/simplify.go:144
		_go_fuzz_dep_.CoverTab[64406]++
//line /usr/local/go/src/regexp/syntax/simplify.go:144
		return sub == re.Sub[0]
//line /usr/local/go/src/regexp/syntax/simplify.go:144
		// _ = "end of CoverTab[64406]"
//line /usr/local/go/src/regexp/syntax/simplify.go:144
	}() {
//line /usr/local/go/src/regexp/syntax/simplify.go:144
		_go_fuzz_dep_.CoverTab[64407]++
								return re
//line /usr/local/go/src/regexp/syntax/simplify.go:145
		// _ = "end of CoverTab[64407]"
	} else {
//line /usr/local/go/src/regexp/syntax/simplify.go:146
		_go_fuzz_dep_.CoverTab[64408]++
//line /usr/local/go/src/regexp/syntax/simplify.go:146
		// _ = "end of CoverTab[64408]"
//line /usr/local/go/src/regexp/syntax/simplify.go:146
	}
//line /usr/local/go/src/regexp/syntax/simplify.go:146
	// _ = "end of CoverTab[64397]"
//line /usr/local/go/src/regexp/syntax/simplify.go:146
	_go_fuzz_dep_.CoverTab[64398]++

							re = &Regexp{Op: op, Flags: flags}
							re.Sub = append(re.Sub0[:0], sub)
							return re
//line /usr/local/go/src/regexp/syntax/simplify.go:150
	// _ = "end of CoverTab[64398]"
}

//line /usr/local/go/src/regexp/syntax/simplify.go:151
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/regexp/syntax/simplify.go:151
var _ = _go_fuzz_dep_.CoverTab
