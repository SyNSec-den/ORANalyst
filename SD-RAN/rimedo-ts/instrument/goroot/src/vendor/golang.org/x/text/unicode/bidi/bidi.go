// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run gen.go gen_trieval.go gen_ranges.go

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:5
//go:generate go run gen.go gen_trieval.go gen_ranges.go

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:13
package bidi

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:13
import (
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:13
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:13
)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:13
import (
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:13
	_atomic_ "sync/atomic"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:13
)

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:19
import (
	"bytes"
)

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:29
type Direction int

const (
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:35
	LeftToRight	Direction	= iota

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:40
	RightToLeft

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:44
	Mixed

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:48
	Neutral
)

type options struct {
	defaultDirection Direction
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:56
type Option func(*options)

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:69
func DefaultDirection(d Direction) Option {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:69
	_go_fuzz_dep_.CoverTab[32127]++
										return func(opts *options) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:70
		_go_fuzz_dep_.CoverTab[32128]++
											opts.defaultDirection = d
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:71
		// _ = "end of CoverTab[32128]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:72
	// _ = "end of CoverTab[32127]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:76
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

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:96
func (p *Paragraph) prepareInput() (n int, err error) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:96
	_go_fuzz_dep_.CoverTab[32129]++
										p.runes = bytes.Runes(p.p)
										bytecount := 0

										p.pairTypes = nil
										p.pairValues = nil
										p.types = nil

										for _, r := range p.runes {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:104
		_go_fuzz_dep_.CoverTab[32131]++
											props, i := LookupRune(r)
											bytecount += i
											cls := props.Class()
											if cls == B {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:108
			_go_fuzz_dep_.CoverTab[32133]++
												return bytecount, nil
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:109
			// _ = "end of CoverTab[32133]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:110
			_go_fuzz_dep_.CoverTab[32134]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:110
			// _ = "end of CoverTab[32134]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:110
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:110
		// _ = "end of CoverTab[32131]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:110
		_go_fuzz_dep_.CoverTab[32132]++
											p.types = append(p.types, cls)
											if props.IsOpeningBracket() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:112
			_go_fuzz_dep_.CoverTab[32135]++
												p.pairTypes = append(p.pairTypes, bpOpen)
												p.pairValues = append(p.pairValues, r)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:114
			// _ = "end of CoverTab[32135]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:115
			_go_fuzz_dep_.CoverTab[32136]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:115
			if props.IsBracket() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:115
				_go_fuzz_dep_.CoverTab[32137]++

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:118
				p.pairTypes = append(p.pairTypes, bpClose)
													p.pairValues = append(p.pairValues, r)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:119
				// _ = "end of CoverTab[32137]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:120
				_go_fuzz_dep_.CoverTab[32138]++
													p.pairTypes = append(p.pairTypes, bpNone)
													p.pairValues = append(p.pairValues, 0)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:122
				// _ = "end of CoverTab[32138]"
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:123
			// _ = "end of CoverTab[32136]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:123
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:123
		// _ = "end of CoverTab[32132]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:124
	// _ = "end of CoverTab[32129]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:124
	_go_fuzz_dep_.CoverTab[32130]++
										return bytecount, nil
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:125
	// _ = "end of CoverTab[32130]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:133
func (p *Paragraph) SetBytes(b []byte, opts ...Option) (n int, err error) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:133
	_go_fuzz_dep_.CoverTab[32139]++
										p.p = b
										p.opts = opts
										return p.prepareInput()
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:136
	// _ = "end of CoverTab[32139]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:144
func (p *Paragraph) SetString(s string, opts ...Option) (n int, err error) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:144
	_go_fuzz_dep_.CoverTab[32140]++
										p.p = []byte(s)
										p.opts = opts
										return p.prepareInput()
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:147
	// _ = "end of CoverTab[32140]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:153
func (p *Paragraph) IsLeftToRight() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:153
	_go_fuzz_dep_.CoverTab[32141]++
										return p.Direction() == LeftToRight
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:154
	// _ = "end of CoverTab[32141]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:160
func (p *Paragraph) Direction() Direction {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:160
	_go_fuzz_dep_.CoverTab[32142]++
										return p.o.Direction()
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:161
	// _ = "end of CoverTab[32142]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:169
func (p *Paragraph) RunAt(pos int) Run {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:169
	_go_fuzz_dep_.CoverTab[32143]++
										c := 0
										runNumber := 0
										for i, r := range p.o.runes {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:172
		_go_fuzz_dep_.CoverTab[32145]++
											c += len(r)
											if pos < c {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:174
			_go_fuzz_dep_.CoverTab[32146]++
												runNumber = i
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:175
			// _ = "end of CoverTab[32146]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:176
			_go_fuzz_dep_.CoverTab[32147]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:176
			// _ = "end of CoverTab[32147]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:176
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:176
		// _ = "end of CoverTab[32145]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:177
	// _ = "end of CoverTab[32143]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:177
	_go_fuzz_dep_.CoverTab[32144]++
										return p.o.Run(runNumber)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:178
	// _ = "end of CoverTab[32144]"
}

func calculateOrdering(levels []level, runes []rune) Ordering {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:181
	_go_fuzz_dep_.CoverTab[32148]++
										var curDir Direction

										prevDir := Neutral
										prevI := 0

										o := Ordering{}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:190
	for i, lvl := range levels {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:190
		_go_fuzz_dep_.CoverTab[32150]++
											if lvl%2 == 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:191
			_go_fuzz_dep_.CoverTab[32152]++
												curDir = LeftToRight
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:192
			// _ = "end of CoverTab[32152]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:193
			_go_fuzz_dep_.CoverTab[32153]++
												curDir = RightToLeft
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:194
			// _ = "end of CoverTab[32153]"
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:195
		// _ = "end of CoverTab[32150]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:195
		_go_fuzz_dep_.CoverTab[32151]++
											if curDir != prevDir {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:196
			_go_fuzz_dep_.CoverTab[32154]++
												if i > 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:197
				_go_fuzz_dep_.CoverTab[32156]++
													o.runes = append(o.runes, runes[prevI:i])
													o.directions = append(o.directions, prevDir)
													o.startpos = append(o.startpos, prevI)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:200
				// _ = "end of CoverTab[32156]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:201
				_go_fuzz_dep_.CoverTab[32157]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:201
				// _ = "end of CoverTab[32157]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:201
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:201
			// _ = "end of CoverTab[32154]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:201
			_go_fuzz_dep_.CoverTab[32155]++
												prevI = i
												prevDir = curDir
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:203
			// _ = "end of CoverTab[32155]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:204
			_go_fuzz_dep_.CoverTab[32158]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:204
			// _ = "end of CoverTab[32158]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:204
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:204
		// _ = "end of CoverTab[32151]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:205
	// _ = "end of CoverTab[32148]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:205
	_go_fuzz_dep_.CoverTab[32149]++
										o.runes = append(o.runes, runes[prevI:])
										o.directions = append(o.directions, prevDir)
										o.startpos = append(o.startpos, prevI)
										return o
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:209
	// _ = "end of CoverTab[32149]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:213
func (p *Paragraph) Order() (Ordering, error) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:213
	_go_fuzz_dep_.CoverTab[32159]++
										if len(p.types) == 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:214
		_go_fuzz_dep_.CoverTab[32164]++
											return Ordering{}, nil
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:215
		// _ = "end of CoverTab[32164]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:216
		_go_fuzz_dep_.CoverTab[32165]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:216
		// _ = "end of CoverTab[32165]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:216
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:216
	// _ = "end of CoverTab[32159]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:216
	_go_fuzz_dep_.CoverTab[32160]++

										for _, fn := range p.opts {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:218
		_go_fuzz_dep_.CoverTab[32166]++
											fn(&p.options)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:219
		// _ = "end of CoverTab[32166]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:220
	// _ = "end of CoverTab[32160]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:220
	_go_fuzz_dep_.CoverTab[32161]++
										lvl := level(-1)
										if p.options.defaultDirection == RightToLeft {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:222
		_go_fuzz_dep_.CoverTab[32167]++
											lvl = 1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:223
		// _ = "end of CoverTab[32167]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:224
		_go_fuzz_dep_.CoverTab[32168]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:224
		// _ = "end of CoverTab[32168]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:224
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:224
	// _ = "end of CoverTab[32161]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:224
	_go_fuzz_dep_.CoverTab[32162]++
										para, err := newParagraph(p.types, p.pairTypes, p.pairValues, lvl)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:226
		_go_fuzz_dep_.CoverTab[32169]++
											return Ordering{}, err
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:227
		// _ = "end of CoverTab[32169]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:228
		_go_fuzz_dep_.CoverTab[32170]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:228
		// _ = "end of CoverTab[32170]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:228
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:228
	// _ = "end of CoverTab[32162]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:228
	_go_fuzz_dep_.CoverTab[32163]++

										levels := para.getLevels([]int{len(p.types)})

										p.o = calculateOrdering(levels, p.runes)
										return p.o, nil
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:233
	// _ = "end of CoverTab[32163]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:238
func (p *Paragraph) Line(start, end int) (Ordering, error) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:238
	_go_fuzz_dep_.CoverTab[32171]++
										lineTypes := p.types[start:end]
										para, err := newParagraph(lineTypes, p.pairTypes[start:end], p.pairValues[start:end], -1)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:241
		_go_fuzz_dep_.CoverTab[32173]++
											return Ordering{}, err
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:242
		// _ = "end of CoverTab[32173]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:243
		_go_fuzz_dep_.CoverTab[32174]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:243
		// _ = "end of CoverTab[32174]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:243
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:243
	// _ = "end of CoverTab[32171]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:243
	_go_fuzz_dep_.CoverTab[32172]++
										levels := para.getLevels([]int{len(lineTypes)})
										o := calculateOrdering(levels, p.runes[start:end])
										return o, nil
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:246
	// _ = "end of CoverTab[32172]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:252
type Ordering struct {
	runes		[][]rune
	directions	[]Direction
	startpos	[]int
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:261
func (o *Ordering) Direction() Direction {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:261
	_go_fuzz_dep_.CoverTab[32175]++
										return o.directions[0]
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:262
	// _ = "end of CoverTab[32175]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:266
func (o *Ordering) NumRuns() int {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:266
	_go_fuzz_dep_.CoverTab[32176]++
										return len(o.runes)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:267
	// _ = "end of CoverTab[32176]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:271
func (o *Ordering) Run(i int) Run {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:271
	_go_fuzz_dep_.CoverTab[32177]++
										r := Run{
		runes:		o.runes[i],
		direction:	o.directions[i],
		startpos:	o.startpos[i],
	}
										return r
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:277
	// _ = "end of CoverTab[32177]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:288
type Run struct {
	runes		[]rune
	direction	Direction
	startpos	int
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:295
func (r *Run) String() string {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:295
	_go_fuzz_dep_.CoverTab[32178]++
										return string(r.runes)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:296
	// _ = "end of CoverTab[32178]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:300
func (r *Run) Bytes() []byte {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:300
	_go_fuzz_dep_.CoverTab[32179]++
										return []byte(r.String())
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:301
	// _ = "end of CoverTab[32179]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:310
func (r *Run) Direction() Direction {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:310
	_go_fuzz_dep_.CoverTab[32180]++
										return r.direction
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:311
	// _ = "end of CoverTab[32180]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:316
func (r *Run) Pos() (start, end int) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:316
	_go_fuzz_dep_.CoverTab[32181]++
										return r.startpos, r.startpos + len(r.runes) - 1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:317
	// _ = "end of CoverTab[32181]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:323
func AppendReverse(out, in []byte) []byte {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:323
	_go_fuzz_dep_.CoverTab[32182]++
										ret := make([]byte, len(in)+len(out))
										copy(ret, out)
										inRunes := bytes.Runes(in)

										for i, r := range inRunes {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:328
		_go_fuzz_dep_.CoverTab[32185]++
											prop, _ := LookupRune(r)
											if prop.IsBracket() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:330
			_go_fuzz_dep_.CoverTab[32186]++
												inRunes[i] = prop.reverseBracket(r)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:331
			// _ = "end of CoverTab[32186]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:332
			_go_fuzz_dep_.CoverTab[32187]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:332
			// _ = "end of CoverTab[32187]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:332
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:332
		// _ = "end of CoverTab[32185]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:333
	// _ = "end of CoverTab[32182]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:333
	_go_fuzz_dep_.CoverTab[32183]++

										for i, j := 0, len(inRunes)-1; i < j; i, j = i+1, j-1 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:335
		_go_fuzz_dep_.CoverTab[32188]++
											inRunes[i], inRunes[j] = inRunes[j], inRunes[i]
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:336
		// _ = "end of CoverTab[32188]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:337
	// _ = "end of CoverTab[32183]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:337
	_go_fuzz_dep_.CoverTab[32184]++
										copy(ret[len(out):], string(inRunes))

										return ret
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:340
	// _ = "end of CoverTab[32184]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:346
func ReverseString(s string) string {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:346
	_go_fuzz_dep_.CoverTab[32189]++
										input := []rune(s)
										li := len(input)
										ret := make([]rune, li)
										for i, r := range input {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:350
		_go_fuzz_dep_.CoverTab[32191]++
											prop, _ := LookupRune(r)
											if prop.IsBracket() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:352
			_go_fuzz_dep_.CoverTab[32192]++
												ret[li-i-1] = prop.reverseBracket(r)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:353
			// _ = "end of CoverTab[32192]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:354
			_go_fuzz_dep_.CoverTab[32193]++
												ret[li-i-1] = r
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:355
			// _ = "end of CoverTab[32193]"
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:356
		// _ = "end of CoverTab[32191]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:357
	// _ = "end of CoverTab[32189]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:357
	_go_fuzz_dep_.CoverTab[32190]++
										return string(ret)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:358
	// _ = "end of CoverTab[32190]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:359
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bidi.go:359
var _ = _go_fuzz_dep_.CoverTab
