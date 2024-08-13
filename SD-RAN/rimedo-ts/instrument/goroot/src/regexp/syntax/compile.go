// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/regexp/syntax/compile.go:5
package syntax

//line /usr/local/go/src/regexp/syntax/compile.go:5
import (
//line /usr/local/go/src/regexp/syntax/compile.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/regexp/syntax/compile.go:5
)
//line /usr/local/go/src/regexp/syntax/compile.go:5
import (
//line /usr/local/go/src/regexp/syntax/compile.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/regexp/syntax/compile.go:5
)

import "unicode"

// A patchList is a list of instruction pointers that need to be filled in (patched).
//line /usr/local/go/src/regexp/syntax/compile.go:9
// Because the pointers haven't been filled in yet, we can reuse their storage
//line /usr/local/go/src/regexp/syntax/compile.go:9
// to hold the list. It's kind of sleazy, but works well in practice.
//line /usr/local/go/src/regexp/syntax/compile.go:9
// See https://swtch.com/~rsc/regexp/regexp1.html for inspiration.
//line /usr/local/go/src/regexp/syntax/compile.go:9
//
//line /usr/local/go/src/regexp/syntax/compile.go:9
// These aren't really pointers: they're integers, so we can reinterpret them
//line /usr/local/go/src/regexp/syntax/compile.go:9
// this way without using package unsafe. A value l.head denotes
//line /usr/local/go/src/regexp/syntax/compile.go:9
// p.inst[l.head>>1].Out (l.head&1==0) or .Arg (l.head&1==1).
//line /usr/local/go/src/regexp/syntax/compile.go:9
// head == 0 denotes the empty list, okay because we start every program
//line /usr/local/go/src/regexp/syntax/compile.go:9
// with a fail instruction, so we'll never want to point at its output link.
//line /usr/local/go/src/regexp/syntax/compile.go:19
type patchList struct {
	head, tail uint32
}

func makePatchList(n uint32) patchList {
//line /usr/local/go/src/regexp/syntax/compile.go:23
	_go_fuzz_dep_.CoverTab[62958]++
							return patchList{n, n}
//line /usr/local/go/src/regexp/syntax/compile.go:24
	// _ = "end of CoverTab[62958]"
}

func (l patchList) patch(p *Prog, val uint32) {
//line /usr/local/go/src/regexp/syntax/compile.go:27
	_go_fuzz_dep_.CoverTab[62959]++
							head := l.head
							for head != 0 {
//line /usr/local/go/src/regexp/syntax/compile.go:29
		_go_fuzz_dep_.CoverTab[62960]++
								i := &p.Inst[head>>1]
								if head&1 == 0 {
//line /usr/local/go/src/regexp/syntax/compile.go:31
			_go_fuzz_dep_.CoverTab[62961]++
									head = i.Out
									i.Out = val
//line /usr/local/go/src/regexp/syntax/compile.go:33
			// _ = "end of CoverTab[62961]"
		} else {
//line /usr/local/go/src/regexp/syntax/compile.go:34
			_go_fuzz_dep_.CoverTab[62962]++
									head = i.Arg
									i.Arg = val
//line /usr/local/go/src/regexp/syntax/compile.go:36
			// _ = "end of CoverTab[62962]"
		}
//line /usr/local/go/src/regexp/syntax/compile.go:37
		// _ = "end of CoverTab[62960]"
	}
//line /usr/local/go/src/regexp/syntax/compile.go:38
	// _ = "end of CoverTab[62959]"
}

func (l1 patchList) append(p *Prog, l2 patchList) patchList {
//line /usr/local/go/src/regexp/syntax/compile.go:41
	_go_fuzz_dep_.CoverTab[62963]++
							if l1.head == 0 {
//line /usr/local/go/src/regexp/syntax/compile.go:42
		_go_fuzz_dep_.CoverTab[62967]++
								return l2
//line /usr/local/go/src/regexp/syntax/compile.go:43
		// _ = "end of CoverTab[62967]"
	} else {
//line /usr/local/go/src/regexp/syntax/compile.go:44
		_go_fuzz_dep_.CoverTab[62968]++
//line /usr/local/go/src/regexp/syntax/compile.go:44
		// _ = "end of CoverTab[62968]"
//line /usr/local/go/src/regexp/syntax/compile.go:44
	}
//line /usr/local/go/src/regexp/syntax/compile.go:44
	// _ = "end of CoverTab[62963]"
//line /usr/local/go/src/regexp/syntax/compile.go:44
	_go_fuzz_dep_.CoverTab[62964]++
							if l2.head == 0 {
//line /usr/local/go/src/regexp/syntax/compile.go:45
		_go_fuzz_dep_.CoverTab[62969]++
								return l1
//line /usr/local/go/src/regexp/syntax/compile.go:46
		// _ = "end of CoverTab[62969]"
	} else {
//line /usr/local/go/src/regexp/syntax/compile.go:47
		_go_fuzz_dep_.CoverTab[62970]++
//line /usr/local/go/src/regexp/syntax/compile.go:47
		// _ = "end of CoverTab[62970]"
//line /usr/local/go/src/regexp/syntax/compile.go:47
	}
//line /usr/local/go/src/regexp/syntax/compile.go:47
	// _ = "end of CoverTab[62964]"
//line /usr/local/go/src/regexp/syntax/compile.go:47
	_go_fuzz_dep_.CoverTab[62965]++

							i := &p.Inst[l1.tail>>1]
							if l1.tail&1 == 0 {
//line /usr/local/go/src/regexp/syntax/compile.go:50
		_go_fuzz_dep_.CoverTab[62971]++
								i.Out = l2.head
//line /usr/local/go/src/regexp/syntax/compile.go:51
		// _ = "end of CoverTab[62971]"
	} else {
//line /usr/local/go/src/regexp/syntax/compile.go:52
		_go_fuzz_dep_.CoverTab[62972]++
								i.Arg = l2.head
//line /usr/local/go/src/regexp/syntax/compile.go:53
		// _ = "end of CoverTab[62972]"
	}
//line /usr/local/go/src/regexp/syntax/compile.go:54
	// _ = "end of CoverTab[62965]"
//line /usr/local/go/src/regexp/syntax/compile.go:54
	_go_fuzz_dep_.CoverTab[62966]++
							return patchList{l1.head, l2.tail}
//line /usr/local/go/src/regexp/syntax/compile.go:55
	// _ = "end of CoverTab[62966]"
}

// A frag represents a compiled program fragment.
type frag struct {
	i		uint32		// index of first instruction
	out		patchList	// where to record end instruction
	nullable	bool		// whether fragment can match empty string
}

type compiler struct {
	p *Prog
}

// Compile compiles the regexp into a program to be executed.
//line /usr/local/go/src/regexp/syntax/compile.go:69
// The regexp should have been simplified already (returned from re.Simplify).
//line /usr/local/go/src/regexp/syntax/compile.go:71
func Compile(re *Regexp) (*Prog, error) {
//line /usr/local/go/src/regexp/syntax/compile.go:71
	_go_fuzz_dep_.CoverTab[62973]++
							var c compiler
							c.init()
							f := c.compile(re)
							f.out.patch(c.p, c.inst(InstMatch).i)
							c.p.Start = int(f.i)
							return c.p, nil
//line /usr/local/go/src/regexp/syntax/compile.go:77
	// _ = "end of CoverTab[62973]"
}

func (c *compiler) init() {
	c.p = new(Prog)
	c.p.NumCap = 2
	c.inst(InstFail)
}

var anyRuneNotNL = []rune{0, '\n' - 1, '\n' + 1, unicode.MaxRune}
var anyRune = []rune{0, unicode.MaxRune}

func (c *compiler) compile(re *Regexp) frag {
//line /usr/local/go/src/regexp/syntax/compile.go:89
	_go_fuzz_dep_.CoverTab[62974]++
							switch re.Op {
	case OpNoMatch:
//line /usr/local/go/src/regexp/syntax/compile.go:91
		_go_fuzz_dep_.CoverTab[62976]++
								return c.fail()
//line /usr/local/go/src/regexp/syntax/compile.go:92
		// _ = "end of CoverTab[62976]"
	case OpEmptyMatch:
//line /usr/local/go/src/regexp/syntax/compile.go:93
		_go_fuzz_dep_.CoverTab[62977]++
								return c.nop()
//line /usr/local/go/src/regexp/syntax/compile.go:94
		// _ = "end of CoverTab[62977]"
	case OpLiteral:
//line /usr/local/go/src/regexp/syntax/compile.go:95
		_go_fuzz_dep_.CoverTab[62978]++
								if len(re.Rune) == 0 {
//line /usr/local/go/src/regexp/syntax/compile.go:96
			_go_fuzz_dep_.CoverTab[63000]++
									return c.nop()
//line /usr/local/go/src/regexp/syntax/compile.go:97
			// _ = "end of CoverTab[63000]"
		} else {
//line /usr/local/go/src/regexp/syntax/compile.go:98
			_go_fuzz_dep_.CoverTab[63001]++
//line /usr/local/go/src/regexp/syntax/compile.go:98
			// _ = "end of CoverTab[63001]"
//line /usr/local/go/src/regexp/syntax/compile.go:98
		}
//line /usr/local/go/src/regexp/syntax/compile.go:98
		// _ = "end of CoverTab[62978]"
//line /usr/local/go/src/regexp/syntax/compile.go:98
		_go_fuzz_dep_.CoverTab[62979]++
								var f frag
								for j := range re.Rune {
//line /usr/local/go/src/regexp/syntax/compile.go:100
			_go_fuzz_dep_.CoverTab[63002]++
									f1 := c.rune(re.Rune[j:j+1], re.Flags)
									if j == 0 {
//line /usr/local/go/src/regexp/syntax/compile.go:102
				_go_fuzz_dep_.CoverTab[63003]++
										f = f1
//line /usr/local/go/src/regexp/syntax/compile.go:103
				// _ = "end of CoverTab[63003]"
			} else {
//line /usr/local/go/src/regexp/syntax/compile.go:104
				_go_fuzz_dep_.CoverTab[63004]++
										f = c.cat(f, f1)
//line /usr/local/go/src/regexp/syntax/compile.go:105
				// _ = "end of CoverTab[63004]"
			}
//line /usr/local/go/src/regexp/syntax/compile.go:106
			// _ = "end of CoverTab[63002]"
		}
//line /usr/local/go/src/regexp/syntax/compile.go:107
		// _ = "end of CoverTab[62979]"
//line /usr/local/go/src/regexp/syntax/compile.go:107
		_go_fuzz_dep_.CoverTab[62980]++
								return f
//line /usr/local/go/src/regexp/syntax/compile.go:108
		// _ = "end of CoverTab[62980]"
	case OpCharClass:
//line /usr/local/go/src/regexp/syntax/compile.go:109
		_go_fuzz_dep_.CoverTab[62981]++
								return c.rune(re.Rune, re.Flags)
//line /usr/local/go/src/regexp/syntax/compile.go:110
		// _ = "end of CoverTab[62981]"
	case OpAnyCharNotNL:
//line /usr/local/go/src/regexp/syntax/compile.go:111
		_go_fuzz_dep_.CoverTab[62982]++
								return c.rune(anyRuneNotNL, 0)
//line /usr/local/go/src/regexp/syntax/compile.go:112
		// _ = "end of CoverTab[62982]"
	case OpAnyChar:
//line /usr/local/go/src/regexp/syntax/compile.go:113
		_go_fuzz_dep_.CoverTab[62983]++
								return c.rune(anyRune, 0)
//line /usr/local/go/src/regexp/syntax/compile.go:114
		// _ = "end of CoverTab[62983]"
	case OpBeginLine:
//line /usr/local/go/src/regexp/syntax/compile.go:115
		_go_fuzz_dep_.CoverTab[62984]++
								return c.empty(EmptyBeginLine)
//line /usr/local/go/src/regexp/syntax/compile.go:116
		// _ = "end of CoverTab[62984]"
	case OpEndLine:
//line /usr/local/go/src/regexp/syntax/compile.go:117
		_go_fuzz_dep_.CoverTab[62985]++
								return c.empty(EmptyEndLine)
//line /usr/local/go/src/regexp/syntax/compile.go:118
		// _ = "end of CoverTab[62985]"
	case OpBeginText:
//line /usr/local/go/src/regexp/syntax/compile.go:119
		_go_fuzz_dep_.CoverTab[62986]++
								return c.empty(EmptyBeginText)
//line /usr/local/go/src/regexp/syntax/compile.go:120
		// _ = "end of CoverTab[62986]"
	case OpEndText:
//line /usr/local/go/src/regexp/syntax/compile.go:121
		_go_fuzz_dep_.CoverTab[62987]++
								return c.empty(EmptyEndText)
//line /usr/local/go/src/regexp/syntax/compile.go:122
		// _ = "end of CoverTab[62987]"
	case OpWordBoundary:
//line /usr/local/go/src/regexp/syntax/compile.go:123
		_go_fuzz_dep_.CoverTab[62988]++
								return c.empty(EmptyWordBoundary)
//line /usr/local/go/src/regexp/syntax/compile.go:124
		// _ = "end of CoverTab[62988]"
	case OpNoWordBoundary:
//line /usr/local/go/src/regexp/syntax/compile.go:125
		_go_fuzz_dep_.CoverTab[62989]++
								return c.empty(EmptyNoWordBoundary)
//line /usr/local/go/src/regexp/syntax/compile.go:126
		// _ = "end of CoverTab[62989]"
	case OpCapture:
//line /usr/local/go/src/regexp/syntax/compile.go:127
		_go_fuzz_dep_.CoverTab[62990]++
								bra := c.cap(uint32(re.Cap << 1))
								sub := c.compile(re.Sub[0])
								ket := c.cap(uint32(re.Cap<<1 | 1))
								return c.cat(c.cat(bra, sub), ket)
//line /usr/local/go/src/regexp/syntax/compile.go:131
		// _ = "end of CoverTab[62990]"
	case OpStar:
//line /usr/local/go/src/regexp/syntax/compile.go:132
		_go_fuzz_dep_.CoverTab[62991]++
								return c.star(c.compile(re.Sub[0]), re.Flags&NonGreedy != 0)
//line /usr/local/go/src/regexp/syntax/compile.go:133
		// _ = "end of CoverTab[62991]"
	case OpPlus:
//line /usr/local/go/src/regexp/syntax/compile.go:134
		_go_fuzz_dep_.CoverTab[62992]++
								return c.plus(c.compile(re.Sub[0]), re.Flags&NonGreedy != 0)
//line /usr/local/go/src/regexp/syntax/compile.go:135
		// _ = "end of CoverTab[62992]"
	case OpQuest:
//line /usr/local/go/src/regexp/syntax/compile.go:136
		_go_fuzz_dep_.CoverTab[62993]++
								return c.quest(c.compile(re.Sub[0]), re.Flags&NonGreedy != 0)
//line /usr/local/go/src/regexp/syntax/compile.go:137
		// _ = "end of CoverTab[62993]"
	case OpConcat:
//line /usr/local/go/src/regexp/syntax/compile.go:138
		_go_fuzz_dep_.CoverTab[62994]++
								if len(re.Sub) == 0 {
//line /usr/local/go/src/regexp/syntax/compile.go:139
			_go_fuzz_dep_.CoverTab[63005]++
									return c.nop()
//line /usr/local/go/src/regexp/syntax/compile.go:140
			// _ = "end of CoverTab[63005]"
		} else {
//line /usr/local/go/src/regexp/syntax/compile.go:141
			_go_fuzz_dep_.CoverTab[63006]++
//line /usr/local/go/src/regexp/syntax/compile.go:141
			// _ = "end of CoverTab[63006]"
//line /usr/local/go/src/regexp/syntax/compile.go:141
		}
//line /usr/local/go/src/regexp/syntax/compile.go:141
		// _ = "end of CoverTab[62994]"
//line /usr/local/go/src/regexp/syntax/compile.go:141
		_go_fuzz_dep_.CoverTab[62995]++
								var f frag
								for i, sub := range re.Sub {
//line /usr/local/go/src/regexp/syntax/compile.go:143
			_go_fuzz_dep_.CoverTab[63007]++
									if i == 0 {
//line /usr/local/go/src/regexp/syntax/compile.go:144
				_go_fuzz_dep_.CoverTab[63008]++
										f = c.compile(sub)
//line /usr/local/go/src/regexp/syntax/compile.go:145
				// _ = "end of CoverTab[63008]"
			} else {
//line /usr/local/go/src/regexp/syntax/compile.go:146
				_go_fuzz_dep_.CoverTab[63009]++
										f = c.cat(f, c.compile(sub))
//line /usr/local/go/src/regexp/syntax/compile.go:147
				// _ = "end of CoverTab[63009]"
			}
//line /usr/local/go/src/regexp/syntax/compile.go:148
			// _ = "end of CoverTab[63007]"
		}
//line /usr/local/go/src/regexp/syntax/compile.go:149
		// _ = "end of CoverTab[62995]"
//line /usr/local/go/src/regexp/syntax/compile.go:149
		_go_fuzz_dep_.CoverTab[62996]++
								return f
//line /usr/local/go/src/regexp/syntax/compile.go:150
		// _ = "end of CoverTab[62996]"
	case OpAlternate:
//line /usr/local/go/src/regexp/syntax/compile.go:151
		_go_fuzz_dep_.CoverTab[62997]++
								var f frag
								for _, sub := range re.Sub {
//line /usr/local/go/src/regexp/syntax/compile.go:153
			_go_fuzz_dep_.CoverTab[63010]++
									f = c.alt(f, c.compile(sub))
//line /usr/local/go/src/regexp/syntax/compile.go:154
			// _ = "end of CoverTab[63010]"
		}
//line /usr/local/go/src/regexp/syntax/compile.go:155
		// _ = "end of CoverTab[62997]"
//line /usr/local/go/src/regexp/syntax/compile.go:155
		_go_fuzz_dep_.CoverTab[62998]++
								return f
//line /usr/local/go/src/regexp/syntax/compile.go:156
		// _ = "end of CoverTab[62998]"
//line /usr/local/go/src/regexp/syntax/compile.go:156
	default:
//line /usr/local/go/src/regexp/syntax/compile.go:156
		_go_fuzz_dep_.CoverTab[62999]++
//line /usr/local/go/src/regexp/syntax/compile.go:156
		// _ = "end of CoverTab[62999]"
	}
//line /usr/local/go/src/regexp/syntax/compile.go:157
	// _ = "end of CoverTab[62974]"
//line /usr/local/go/src/regexp/syntax/compile.go:157
	_go_fuzz_dep_.CoverTab[62975]++
							panic("regexp: unhandled case in compile")
//line /usr/local/go/src/regexp/syntax/compile.go:158
	// _ = "end of CoverTab[62975]"
}

func (c *compiler) inst(op InstOp) frag {
//line /usr/local/go/src/regexp/syntax/compile.go:161
	_go_fuzz_dep_.CoverTab[63011]++

							f := frag{i: uint32(len(c.p.Inst)), nullable: true}
							c.p.Inst = append(c.p.Inst, Inst{Op: op})
							return f
//line /usr/local/go/src/regexp/syntax/compile.go:165
	// _ = "end of CoverTab[63011]"
}

func (c *compiler) nop() frag {
//line /usr/local/go/src/regexp/syntax/compile.go:168
	_go_fuzz_dep_.CoverTab[63012]++
							f := c.inst(InstNop)
							f.out = makePatchList(f.i << 1)
							return f
//line /usr/local/go/src/regexp/syntax/compile.go:171
	// _ = "end of CoverTab[63012]"
}

func (c *compiler) fail() frag {
//line /usr/local/go/src/regexp/syntax/compile.go:174
	_go_fuzz_dep_.CoverTab[63013]++
							return frag{}
//line /usr/local/go/src/regexp/syntax/compile.go:175
	// _ = "end of CoverTab[63013]"
}

func (c *compiler) cap(arg uint32) frag {
//line /usr/local/go/src/regexp/syntax/compile.go:178
	_go_fuzz_dep_.CoverTab[63014]++
							f := c.inst(InstCapture)
							f.out = makePatchList(f.i << 1)
							c.p.Inst[f.i].Arg = arg

							if c.p.NumCap < int(arg)+1 {
//line /usr/local/go/src/regexp/syntax/compile.go:183
		_go_fuzz_dep_.CoverTab[63016]++
								c.p.NumCap = int(arg) + 1
//line /usr/local/go/src/regexp/syntax/compile.go:184
		// _ = "end of CoverTab[63016]"
	} else {
//line /usr/local/go/src/regexp/syntax/compile.go:185
		_go_fuzz_dep_.CoverTab[63017]++
//line /usr/local/go/src/regexp/syntax/compile.go:185
		// _ = "end of CoverTab[63017]"
//line /usr/local/go/src/regexp/syntax/compile.go:185
	}
//line /usr/local/go/src/regexp/syntax/compile.go:185
	// _ = "end of CoverTab[63014]"
//line /usr/local/go/src/regexp/syntax/compile.go:185
	_go_fuzz_dep_.CoverTab[63015]++
							return f
//line /usr/local/go/src/regexp/syntax/compile.go:186
	// _ = "end of CoverTab[63015]"
}

func (c *compiler) cat(f1, f2 frag) frag {
//line /usr/local/go/src/regexp/syntax/compile.go:189
	_go_fuzz_dep_.CoverTab[63018]++

							if f1.i == 0 || func() bool {
//line /usr/local/go/src/regexp/syntax/compile.go:191
		_go_fuzz_dep_.CoverTab[63020]++
//line /usr/local/go/src/regexp/syntax/compile.go:191
		return f2.i == 0
//line /usr/local/go/src/regexp/syntax/compile.go:191
		// _ = "end of CoverTab[63020]"
//line /usr/local/go/src/regexp/syntax/compile.go:191
	}() {
//line /usr/local/go/src/regexp/syntax/compile.go:191
		_go_fuzz_dep_.CoverTab[63021]++
								return frag{}
//line /usr/local/go/src/regexp/syntax/compile.go:192
		// _ = "end of CoverTab[63021]"
	} else {
//line /usr/local/go/src/regexp/syntax/compile.go:193
		_go_fuzz_dep_.CoverTab[63022]++
//line /usr/local/go/src/regexp/syntax/compile.go:193
		// _ = "end of CoverTab[63022]"
//line /usr/local/go/src/regexp/syntax/compile.go:193
	}
//line /usr/local/go/src/regexp/syntax/compile.go:193
	// _ = "end of CoverTab[63018]"
//line /usr/local/go/src/regexp/syntax/compile.go:193
	_go_fuzz_dep_.CoverTab[63019]++

//line /usr/local/go/src/regexp/syntax/compile.go:197
	f1.out.patch(c.p, f2.i)
	return frag{f1.i, f2.out, f1.nullable && func() bool {
//line /usr/local/go/src/regexp/syntax/compile.go:198
		_go_fuzz_dep_.CoverTab[63023]++
//line /usr/local/go/src/regexp/syntax/compile.go:198
		return f2.nullable
//line /usr/local/go/src/regexp/syntax/compile.go:198
		// _ = "end of CoverTab[63023]"
//line /usr/local/go/src/regexp/syntax/compile.go:198
	}()}
//line /usr/local/go/src/regexp/syntax/compile.go:198
	// _ = "end of CoverTab[63019]"
}

func (c *compiler) alt(f1, f2 frag) frag {
//line /usr/local/go/src/regexp/syntax/compile.go:201
	_go_fuzz_dep_.CoverTab[63024]++

							if f1.i == 0 {
//line /usr/local/go/src/regexp/syntax/compile.go:203
		_go_fuzz_dep_.CoverTab[63027]++
								return f2
//line /usr/local/go/src/regexp/syntax/compile.go:204
		// _ = "end of CoverTab[63027]"
	} else {
//line /usr/local/go/src/regexp/syntax/compile.go:205
		_go_fuzz_dep_.CoverTab[63028]++
//line /usr/local/go/src/regexp/syntax/compile.go:205
		// _ = "end of CoverTab[63028]"
//line /usr/local/go/src/regexp/syntax/compile.go:205
	}
//line /usr/local/go/src/regexp/syntax/compile.go:205
	// _ = "end of CoverTab[63024]"
//line /usr/local/go/src/regexp/syntax/compile.go:205
	_go_fuzz_dep_.CoverTab[63025]++
							if f2.i == 0 {
//line /usr/local/go/src/regexp/syntax/compile.go:206
		_go_fuzz_dep_.CoverTab[63029]++
								return f1
//line /usr/local/go/src/regexp/syntax/compile.go:207
		// _ = "end of CoverTab[63029]"
	} else {
//line /usr/local/go/src/regexp/syntax/compile.go:208
		_go_fuzz_dep_.CoverTab[63030]++
//line /usr/local/go/src/regexp/syntax/compile.go:208
		// _ = "end of CoverTab[63030]"
//line /usr/local/go/src/regexp/syntax/compile.go:208
	}
//line /usr/local/go/src/regexp/syntax/compile.go:208
	// _ = "end of CoverTab[63025]"
//line /usr/local/go/src/regexp/syntax/compile.go:208
	_go_fuzz_dep_.CoverTab[63026]++

							f := c.inst(InstAlt)
							i := &c.p.Inst[f.i]
							i.Out = f1.i
							i.Arg = f2.i
							f.out = f1.out.append(c.p, f2.out)
							f.nullable = f1.nullable || func() bool {
//line /usr/local/go/src/regexp/syntax/compile.go:215
		_go_fuzz_dep_.CoverTab[63031]++
//line /usr/local/go/src/regexp/syntax/compile.go:215
		return f2.nullable
//line /usr/local/go/src/regexp/syntax/compile.go:215
		// _ = "end of CoverTab[63031]"
//line /usr/local/go/src/regexp/syntax/compile.go:215
	}()
							return f
//line /usr/local/go/src/regexp/syntax/compile.go:216
	// _ = "end of CoverTab[63026]"
}

func (c *compiler) quest(f1 frag, nongreedy bool) frag {
//line /usr/local/go/src/regexp/syntax/compile.go:219
	_go_fuzz_dep_.CoverTab[63032]++
							f := c.inst(InstAlt)
							i := &c.p.Inst[f.i]
							if nongreedy {
//line /usr/local/go/src/regexp/syntax/compile.go:222
		_go_fuzz_dep_.CoverTab[63034]++
								i.Arg = f1.i
								f.out = makePatchList(f.i << 1)
//line /usr/local/go/src/regexp/syntax/compile.go:224
		// _ = "end of CoverTab[63034]"
	} else {
//line /usr/local/go/src/regexp/syntax/compile.go:225
		_go_fuzz_dep_.CoverTab[63035]++
								i.Out = f1.i
								f.out = makePatchList(f.i<<1 | 1)
//line /usr/local/go/src/regexp/syntax/compile.go:227
		// _ = "end of CoverTab[63035]"
	}
//line /usr/local/go/src/regexp/syntax/compile.go:228
	// _ = "end of CoverTab[63032]"
//line /usr/local/go/src/regexp/syntax/compile.go:228
	_go_fuzz_dep_.CoverTab[63033]++
							f.out = f.out.append(c.p, f1.out)
							return f
//line /usr/local/go/src/regexp/syntax/compile.go:230
	// _ = "end of CoverTab[63033]"
}

// loop returns the fragment for the main loop of a plus or star.
//line /usr/local/go/src/regexp/syntax/compile.go:233
// For plus, it can be used after changing the entry to f1.i.
//line /usr/local/go/src/regexp/syntax/compile.go:233
// For star, it can be used directly when f1 can't match an empty string.
//line /usr/local/go/src/regexp/syntax/compile.go:233
// (When f1 can match an empty string, f1* must be implemented as (f1+)?
//line /usr/local/go/src/regexp/syntax/compile.go:233
// to get the priority match order correct.)
//line /usr/local/go/src/regexp/syntax/compile.go:238
func (c *compiler) loop(f1 frag, nongreedy bool) frag {
//line /usr/local/go/src/regexp/syntax/compile.go:238
	_go_fuzz_dep_.CoverTab[63036]++
							f := c.inst(InstAlt)
							i := &c.p.Inst[f.i]
							if nongreedy {
//line /usr/local/go/src/regexp/syntax/compile.go:241
		_go_fuzz_dep_.CoverTab[63038]++
								i.Arg = f1.i
								f.out = makePatchList(f.i << 1)
//line /usr/local/go/src/regexp/syntax/compile.go:243
		// _ = "end of CoverTab[63038]"
	} else {
//line /usr/local/go/src/regexp/syntax/compile.go:244
		_go_fuzz_dep_.CoverTab[63039]++
								i.Out = f1.i
								f.out = makePatchList(f.i<<1 | 1)
//line /usr/local/go/src/regexp/syntax/compile.go:246
		// _ = "end of CoverTab[63039]"
	}
//line /usr/local/go/src/regexp/syntax/compile.go:247
	// _ = "end of CoverTab[63036]"
//line /usr/local/go/src/regexp/syntax/compile.go:247
	_go_fuzz_dep_.CoverTab[63037]++
							f1.out.patch(c.p, f.i)
							return f
//line /usr/local/go/src/regexp/syntax/compile.go:249
	// _ = "end of CoverTab[63037]"
}

func (c *compiler) star(f1 frag, nongreedy bool) frag {
//line /usr/local/go/src/regexp/syntax/compile.go:252
	_go_fuzz_dep_.CoverTab[63040]++
							if f1.nullable {
//line /usr/local/go/src/regexp/syntax/compile.go:253
		_go_fuzz_dep_.CoverTab[63042]++

//line /usr/local/go/src/regexp/syntax/compile.go:256
		return c.quest(c.plus(f1, nongreedy), nongreedy)
//line /usr/local/go/src/regexp/syntax/compile.go:256
		// _ = "end of CoverTab[63042]"
	} else {
//line /usr/local/go/src/regexp/syntax/compile.go:257
		_go_fuzz_dep_.CoverTab[63043]++
//line /usr/local/go/src/regexp/syntax/compile.go:257
		// _ = "end of CoverTab[63043]"
//line /usr/local/go/src/regexp/syntax/compile.go:257
	}
//line /usr/local/go/src/regexp/syntax/compile.go:257
	// _ = "end of CoverTab[63040]"
//line /usr/local/go/src/regexp/syntax/compile.go:257
	_go_fuzz_dep_.CoverTab[63041]++
							return c.loop(f1, nongreedy)
//line /usr/local/go/src/regexp/syntax/compile.go:258
	// _ = "end of CoverTab[63041]"
}

func (c *compiler) plus(f1 frag, nongreedy bool) frag {
//line /usr/local/go/src/regexp/syntax/compile.go:261
	_go_fuzz_dep_.CoverTab[63044]++
							return frag{f1.i, c.loop(f1, nongreedy).out, f1.nullable}
//line /usr/local/go/src/regexp/syntax/compile.go:262
	// _ = "end of CoverTab[63044]"
}

func (c *compiler) empty(op EmptyOp) frag {
//line /usr/local/go/src/regexp/syntax/compile.go:265
	_go_fuzz_dep_.CoverTab[63045]++
							f := c.inst(InstEmptyWidth)
							c.p.Inst[f.i].Arg = uint32(op)
							f.out = makePatchList(f.i << 1)
							return f
//line /usr/local/go/src/regexp/syntax/compile.go:269
	// _ = "end of CoverTab[63045]"
}

func (c *compiler) rune(r []rune, flags Flags) frag {
//line /usr/local/go/src/regexp/syntax/compile.go:272
	_go_fuzz_dep_.CoverTab[63046]++
							f := c.inst(InstRune)
							f.nullable = false
							i := &c.p.Inst[f.i]
							i.Rune = r
							flags &= FoldCase
							if len(r) != 1 || func() bool {
//line /usr/local/go/src/regexp/syntax/compile.go:278
		_go_fuzz_dep_.CoverTab[63049]++
//line /usr/local/go/src/regexp/syntax/compile.go:278
		return unicode.SimpleFold(r[0]) == r[0]
//line /usr/local/go/src/regexp/syntax/compile.go:278
		// _ = "end of CoverTab[63049]"
//line /usr/local/go/src/regexp/syntax/compile.go:278
	}() {
//line /usr/local/go/src/regexp/syntax/compile.go:278
		_go_fuzz_dep_.CoverTab[63050]++

								flags &^= FoldCase
//line /usr/local/go/src/regexp/syntax/compile.go:280
		// _ = "end of CoverTab[63050]"
	} else {
//line /usr/local/go/src/regexp/syntax/compile.go:281
		_go_fuzz_dep_.CoverTab[63051]++
//line /usr/local/go/src/regexp/syntax/compile.go:281
		// _ = "end of CoverTab[63051]"
//line /usr/local/go/src/regexp/syntax/compile.go:281
	}
//line /usr/local/go/src/regexp/syntax/compile.go:281
	// _ = "end of CoverTab[63046]"
//line /usr/local/go/src/regexp/syntax/compile.go:281
	_go_fuzz_dep_.CoverTab[63047]++
							i.Arg = uint32(flags)
							f.out = makePatchList(f.i << 1)

//line /usr/local/go/src/regexp/syntax/compile.go:286
	switch {
	case flags&FoldCase == 0 && func() bool {
//line /usr/local/go/src/regexp/syntax/compile.go:287
		_go_fuzz_dep_.CoverTab[63056]++
//line /usr/local/go/src/regexp/syntax/compile.go:287
		return (len(r) == 1 || func() bool {
//line /usr/local/go/src/regexp/syntax/compile.go:287
			_go_fuzz_dep_.CoverTab[63057]++
//line /usr/local/go/src/regexp/syntax/compile.go:287
			return len(r) == 2 && func() bool {
//line /usr/local/go/src/regexp/syntax/compile.go:287
				_go_fuzz_dep_.CoverTab[63058]++
//line /usr/local/go/src/regexp/syntax/compile.go:287
				return r[0] == r[1]
//line /usr/local/go/src/regexp/syntax/compile.go:287
				// _ = "end of CoverTab[63058]"
//line /usr/local/go/src/regexp/syntax/compile.go:287
			}()
//line /usr/local/go/src/regexp/syntax/compile.go:287
			// _ = "end of CoverTab[63057]"
//line /usr/local/go/src/regexp/syntax/compile.go:287
		}())
//line /usr/local/go/src/regexp/syntax/compile.go:287
		// _ = "end of CoverTab[63056]"
//line /usr/local/go/src/regexp/syntax/compile.go:287
	}():
//line /usr/local/go/src/regexp/syntax/compile.go:287
		_go_fuzz_dep_.CoverTab[63052]++
								i.Op = InstRune1
//line /usr/local/go/src/regexp/syntax/compile.go:288
		// _ = "end of CoverTab[63052]"
	case len(r) == 2 && func() bool {
//line /usr/local/go/src/regexp/syntax/compile.go:289
		_go_fuzz_dep_.CoverTab[63059]++
//line /usr/local/go/src/regexp/syntax/compile.go:289
		return r[0] == 0
//line /usr/local/go/src/regexp/syntax/compile.go:289
		// _ = "end of CoverTab[63059]"
//line /usr/local/go/src/regexp/syntax/compile.go:289
	}() && func() bool {
//line /usr/local/go/src/regexp/syntax/compile.go:289
		_go_fuzz_dep_.CoverTab[63060]++
//line /usr/local/go/src/regexp/syntax/compile.go:289
		return r[1] == unicode.MaxRune
//line /usr/local/go/src/regexp/syntax/compile.go:289
		// _ = "end of CoverTab[63060]"
//line /usr/local/go/src/regexp/syntax/compile.go:289
	}():
//line /usr/local/go/src/regexp/syntax/compile.go:289
		_go_fuzz_dep_.CoverTab[63053]++
								i.Op = InstRuneAny
//line /usr/local/go/src/regexp/syntax/compile.go:290
		// _ = "end of CoverTab[63053]"
	case len(r) == 4 && func() bool {
//line /usr/local/go/src/regexp/syntax/compile.go:291
		_go_fuzz_dep_.CoverTab[63061]++
//line /usr/local/go/src/regexp/syntax/compile.go:291
		return r[0] == 0
//line /usr/local/go/src/regexp/syntax/compile.go:291
		// _ = "end of CoverTab[63061]"
//line /usr/local/go/src/regexp/syntax/compile.go:291
	}() && func() bool {
//line /usr/local/go/src/regexp/syntax/compile.go:291
		_go_fuzz_dep_.CoverTab[63062]++
//line /usr/local/go/src/regexp/syntax/compile.go:291
		return r[1] == '\n'-1
//line /usr/local/go/src/regexp/syntax/compile.go:291
		// _ = "end of CoverTab[63062]"
//line /usr/local/go/src/regexp/syntax/compile.go:291
	}() && func() bool {
//line /usr/local/go/src/regexp/syntax/compile.go:291
		_go_fuzz_dep_.CoverTab[63063]++
//line /usr/local/go/src/regexp/syntax/compile.go:291
		return r[2] == '\n'+1
//line /usr/local/go/src/regexp/syntax/compile.go:291
		// _ = "end of CoverTab[63063]"
//line /usr/local/go/src/regexp/syntax/compile.go:291
	}() && func() bool {
//line /usr/local/go/src/regexp/syntax/compile.go:291
		_go_fuzz_dep_.CoverTab[63064]++
//line /usr/local/go/src/regexp/syntax/compile.go:291
		return r[3] == unicode.MaxRune
//line /usr/local/go/src/regexp/syntax/compile.go:291
		// _ = "end of CoverTab[63064]"
//line /usr/local/go/src/regexp/syntax/compile.go:291
	}():
//line /usr/local/go/src/regexp/syntax/compile.go:291
		_go_fuzz_dep_.CoverTab[63054]++
								i.Op = InstRuneAnyNotNL
//line /usr/local/go/src/regexp/syntax/compile.go:292
		// _ = "end of CoverTab[63054]"
//line /usr/local/go/src/regexp/syntax/compile.go:292
	default:
//line /usr/local/go/src/regexp/syntax/compile.go:292
		_go_fuzz_dep_.CoverTab[63055]++
//line /usr/local/go/src/regexp/syntax/compile.go:292
		// _ = "end of CoverTab[63055]"
	}
//line /usr/local/go/src/regexp/syntax/compile.go:293
	// _ = "end of CoverTab[63047]"
//line /usr/local/go/src/regexp/syntax/compile.go:293
	_go_fuzz_dep_.CoverTab[63048]++

							return f
//line /usr/local/go/src/regexp/syntax/compile.go:295
	// _ = "end of CoverTab[63048]"
}

//line /usr/local/go/src/regexp/syntax/compile.go:296
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/regexp/syntax/compile.go:296
var _ = _go_fuzz_dep_.CoverTab
