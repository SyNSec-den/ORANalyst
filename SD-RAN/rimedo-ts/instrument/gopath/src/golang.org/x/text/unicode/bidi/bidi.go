// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run gen.go gen_trieval.go gen_ranges.go

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:5
//go:generate go run gen.go gen_trieval.go gen_ranges.go

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:13
package bidi

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:13
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:13
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:13
)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:13
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:13
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:13
)

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:19
import (
	"bytes"
)

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:29
type Direction int

const (
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:35
	LeftToRight	Direction	= iota

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:40
	RightToLeft

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:44
	Mixed

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:48
	Neutral
)

type options struct {
	defaultDirection Direction
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:56
type Option func(*options)

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:69
func DefaultDirection(d Direction) Option {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:69
	_go_fuzz_dep_.CoverTab[69408]++
											return func(opts *options) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:70
		_go_fuzz_dep_.CoverTab[69409]++
												opts.defaultDirection = d
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:71
		// _ = "end of CoverTab[69409]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:72
	// _ = "end of CoverTab[69408]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:76
type Paragraph struct {
	p		[]byte
	o		Ordering
	opts		[]Option
	types		[]Class
	pairTypes	[]bracketType
	pairValues	[]rune
	runes		[]rune
	options		options
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:96
func (p *Paragraph) prepareInput() (n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:96
	_go_fuzz_dep_.CoverTab[69410]++
											p.runes = bytes.Runes(p.p)
											bytecount := 0

											p.pairTypes = nil
											p.pairValues = nil
											p.types = nil

											for _, r := range p.runes {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:104
		_go_fuzz_dep_.CoverTab[69412]++
												props, i := LookupRune(r)
												bytecount += i
												cls := props.Class()
												if cls == B {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:108
			_go_fuzz_dep_.CoverTab[69414]++
													return bytecount, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:109
			// _ = "end of CoverTab[69414]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:110
			_go_fuzz_dep_.CoverTab[69415]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:110
			// _ = "end of CoverTab[69415]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:110
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:110
		// _ = "end of CoverTab[69412]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:110
		_go_fuzz_dep_.CoverTab[69413]++
												p.types = append(p.types, cls)
												if props.IsOpeningBracket() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:112
			_go_fuzz_dep_.CoverTab[69416]++
													p.pairTypes = append(p.pairTypes, bpOpen)
													p.pairValues = append(p.pairValues, r)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:114
			// _ = "end of CoverTab[69416]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:115
			_go_fuzz_dep_.CoverTab[69417]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:115
			if props.IsBracket() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:115
				_go_fuzz_dep_.CoverTab[69418]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:118
				p.pairTypes = append(p.pairTypes, bpClose)
														p.pairValues = append(p.pairValues, r)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:119
				// _ = "end of CoverTab[69418]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:120
				_go_fuzz_dep_.CoverTab[69419]++
														p.pairTypes = append(p.pairTypes, bpNone)
														p.pairValues = append(p.pairValues, 0)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:122
				// _ = "end of CoverTab[69419]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:123
			// _ = "end of CoverTab[69417]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:123
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:123
		// _ = "end of CoverTab[69413]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:124
	// _ = "end of CoverTab[69410]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:124
	_go_fuzz_dep_.CoverTab[69411]++
											return bytecount, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:125
	// _ = "end of CoverTab[69411]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:133
func (p *Paragraph) SetBytes(b []byte, opts ...Option) (n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:133
	_go_fuzz_dep_.CoverTab[69420]++
											p.p = b
											p.opts = opts
											return p.prepareInput()
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:136
	// _ = "end of CoverTab[69420]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:144
func (p *Paragraph) SetString(s string, opts ...Option) (n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:144
	_go_fuzz_dep_.CoverTab[69421]++
											p.p = []byte(s)
											p.opts = opts
											return p.prepareInput()
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:147
	// _ = "end of CoverTab[69421]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:153
func (p *Paragraph) IsLeftToRight() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:153
	_go_fuzz_dep_.CoverTab[69422]++
											return p.Direction() == LeftToRight
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:154
	// _ = "end of CoverTab[69422]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:160
func (p *Paragraph) Direction() Direction {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:160
	_go_fuzz_dep_.CoverTab[69423]++
											return p.o.Direction()
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:161
	// _ = "end of CoverTab[69423]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:169
func (p *Paragraph) RunAt(pos int) Run {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:169
	_go_fuzz_dep_.CoverTab[69424]++
											c := 0
											runNumber := 0
											for i, r := range p.o.runes {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:172
		_go_fuzz_dep_.CoverTab[69426]++
												c += len(r)
												if pos < c {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:174
			_go_fuzz_dep_.CoverTab[69427]++
													runNumber = i
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:175
			// _ = "end of CoverTab[69427]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:176
			_go_fuzz_dep_.CoverTab[69428]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:176
			// _ = "end of CoverTab[69428]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:176
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:176
		// _ = "end of CoverTab[69426]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:177
	// _ = "end of CoverTab[69424]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:177
	_go_fuzz_dep_.CoverTab[69425]++
											return p.o.Run(runNumber)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:178
	// _ = "end of CoverTab[69425]"
}

func calculateOrdering(levels []level, runes []rune) Ordering {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:181
	_go_fuzz_dep_.CoverTab[69429]++
											var curDir Direction

											prevDir := Neutral
											prevI := 0

											o := Ordering{}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:190
	for i, lvl := range levels {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:190
		_go_fuzz_dep_.CoverTab[69431]++
												if lvl%2 == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:191
			_go_fuzz_dep_.CoverTab[69433]++
													curDir = LeftToRight
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:192
			// _ = "end of CoverTab[69433]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:193
			_go_fuzz_dep_.CoverTab[69434]++
													curDir = RightToLeft
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:194
			// _ = "end of CoverTab[69434]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:195
		// _ = "end of CoverTab[69431]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:195
		_go_fuzz_dep_.CoverTab[69432]++
												if curDir != prevDir {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:196
			_go_fuzz_dep_.CoverTab[69435]++
													if i > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:197
				_go_fuzz_dep_.CoverTab[69437]++
														o.runes = append(o.runes, runes[prevI:i])
														o.directions = append(o.directions, prevDir)
														o.startpos = append(o.startpos, prevI)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:200
				// _ = "end of CoverTab[69437]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:201
				_go_fuzz_dep_.CoverTab[69438]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:201
				// _ = "end of CoverTab[69438]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:201
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:201
			// _ = "end of CoverTab[69435]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:201
			_go_fuzz_dep_.CoverTab[69436]++
													prevI = i
													prevDir = curDir
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:203
			// _ = "end of CoverTab[69436]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:204
			_go_fuzz_dep_.CoverTab[69439]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:204
			// _ = "end of CoverTab[69439]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:204
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:204
		// _ = "end of CoverTab[69432]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:205
	// _ = "end of CoverTab[69429]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:205
	_go_fuzz_dep_.CoverTab[69430]++
											o.runes = append(o.runes, runes[prevI:])
											o.directions = append(o.directions, prevDir)
											o.startpos = append(o.startpos, prevI)
											return o
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:209
	// _ = "end of CoverTab[69430]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:213
func (p *Paragraph) Order() (Ordering, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:213
	_go_fuzz_dep_.CoverTab[69440]++
											if len(p.types) == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:214
		_go_fuzz_dep_.CoverTab[69445]++
												return Ordering{}, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:215
		// _ = "end of CoverTab[69445]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:216
		_go_fuzz_dep_.CoverTab[69446]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:216
		// _ = "end of CoverTab[69446]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:216
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:216
	// _ = "end of CoverTab[69440]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:216
	_go_fuzz_dep_.CoverTab[69441]++

											for _, fn := range p.opts {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:218
		_go_fuzz_dep_.CoverTab[69447]++
												fn(&p.options)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:219
		// _ = "end of CoverTab[69447]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:220
	// _ = "end of CoverTab[69441]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:220
	_go_fuzz_dep_.CoverTab[69442]++
											lvl := level(-1)
											if p.options.defaultDirection == RightToLeft {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:222
		_go_fuzz_dep_.CoverTab[69448]++
												lvl = 1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:223
		// _ = "end of CoverTab[69448]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:224
		_go_fuzz_dep_.CoverTab[69449]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:224
		// _ = "end of CoverTab[69449]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:224
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:224
	// _ = "end of CoverTab[69442]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:224
	_go_fuzz_dep_.CoverTab[69443]++
											para, err := newParagraph(p.types, p.pairTypes, p.pairValues, lvl)
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:226
		_go_fuzz_dep_.CoverTab[69450]++
												return Ordering{}, err
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:227
		// _ = "end of CoverTab[69450]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:228
		_go_fuzz_dep_.CoverTab[69451]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:228
		// _ = "end of CoverTab[69451]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:228
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:228
	// _ = "end of CoverTab[69443]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:228
	_go_fuzz_dep_.CoverTab[69444]++

											levels := para.getLevels([]int{len(p.types)})

											p.o = calculateOrdering(levels, p.runes)
											return p.o, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:233
	// _ = "end of CoverTab[69444]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:238
func (p *Paragraph) Line(start, end int) (Ordering, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:238
	_go_fuzz_dep_.CoverTab[69452]++
											lineTypes := p.types[start:end]
											para, err := newParagraph(lineTypes, p.pairTypes[start:end], p.pairValues[start:end], -1)
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:241
		_go_fuzz_dep_.CoverTab[69454]++
												return Ordering{}, err
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:242
		// _ = "end of CoverTab[69454]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:243
		_go_fuzz_dep_.CoverTab[69455]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:243
		// _ = "end of CoverTab[69455]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:243
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:243
	// _ = "end of CoverTab[69452]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:243
	_go_fuzz_dep_.CoverTab[69453]++
											levels := para.getLevels([]int{len(lineTypes)})
											o := calculateOrdering(levels, p.runes[start:end])
											return o, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:246
	// _ = "end of CoverTab[69453]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:252
type Ordering struct {
	runes		[][]rune
	directions	[]Direction
	startpos	[]int
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:261
func (o *Ordering) Direction() Direction {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:261
	_go_fuzz_dep_.CoverTab[69456]++
											return o.directions[0]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:262
	// _ = "end of CoverTab[69456]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:266
func (o *Ordering) NumRuns() int {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:266
	_go_fuzz_dep_.CoverTab[69457]++
											return len(o.runes)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:267
	// _ = "end of CoverTab[69457]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:271
func (o *Ordering) Run(i int) Run {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:271
	_go_fuzz_dep_.CoverTab[69458]++
											r := Run{
		runes:		o.runes[i],
		direction:	o.directions[i],
		startpos:	o.startpos[i],
	}
											return r
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:277
	// _ = "end of CoverTab[69458]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:288
type Run struct {
	runes		[]rune
	direction	Direction
	startpos	int
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:295
func (r *Run) String() string {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:295
	_go_fuzz_dep_.CoverTab[69459]++
											return string(r.runes)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:296
	// _ = "end of CoverTab[69459]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:300
func (r *Run) Bytes() []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:300
	_go_fuzz_dep_.CoverTab[69460]++
											return []byte(r.String())
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:301
	// _ = "end of CoverTab[69460]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:310
func (r *Run) Direction() Direction {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:310
	_go_fuzz_dep_.CoverTab[69461]++
											return r.direction
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:311
	// _ = "end of CoverTab[69461]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:316
func (r *Run) Pos() (start, end int) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:316
	_go_fuzz_dep_.CoverTab[69462]++
											return r.startpos, r.startpos + len(r.runes) - 1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:317
	// _ = "end of CoverTab[69462]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:323
func AppendReverse(out, in []byte) []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:323
	_go_fuzz_dep_.CoverTab[69463]++
											ret := make([]byte, len(in)+len(out))
											copy(ret, out)
											inRunes := bytes.Runes(in)

											for i, r := range inRunes {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:328
		_go_fuzz_dep_.CoverTab[69466]++
												prop, _ := LookupRune(r)
												if prop.IsBracket() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:330
			_go_fuzz_dep_.CoverTab[69467]++
													inRunes[i] = prop.reverseBracket(r)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:331
			// _ = "end of CoverTab[69467]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:332
			_go_fuzz_dep_.CoverTab[69468]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:332
			// _ = "end of CoverTab[69468]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:332
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:332
		// _ = "end of CoverTab[69466]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:333
	// _ = "end of CoverTab[69463]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:333
	_go_fuzz_dep_.CoverTab[69464]++

											for i, j := 0, len(inRunes)-1; i < j; i, j = i+1, j-1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:335
		_go_fuzz_dep_.CoverTab[69469]++
												inRunes[i], inRunes[j] = inRunes[j], inRunes[i]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:336
		// _ = "end of CoverTab[69469]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:337
	// _ = "end of CoverTab[69464]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:337
	_go_fuzz_dep_.CoverTab[69465]++
											copy(ret[len(out):], string(inRunes))

											return ret
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:340
	// _ = "end of CoverTab[69465]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:346
func ReverseString(s string) string {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:346
	_go_fuzz_dep_.CoverTab[69470]++
											input := []rune(s)
											li := len(input)
											ret := make([]rune, li)
											for i, r := range input {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:350
		_go_fuzz_dep_.CoverTab[69472]++
												prop, _ := LookupRune(r)
												if prop.IsBracket() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:352
			_go_fuzz_dep_.CoverTab[69473]++
													ret[li-i-1] = prop.reverseBracket(r)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:353
			// _ = "end of CoverTab[69473]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:354
			_go_fuzz_dep_.CoverTab[69474]++
													ret[li-i-1] = r
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:355
			// _ = "end of CoverTab[69474]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:356
		// _ = "end of CoverTab[69472]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:357
	// _ = "end of CoverTab[69470]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:357
	_go_fuzz_dep_.CoverTab[69471]++
											return string(ret)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:358
	// _ = "end of CoverTab[69471]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:359
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bidi.go:359
var _ = _go_fuzz_dep_.CoverTab
