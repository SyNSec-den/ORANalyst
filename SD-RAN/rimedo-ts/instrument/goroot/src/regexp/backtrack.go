// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// backtrack is a regular expression search with submatch
// tracking for small regular expressions and texts. It allocates
// a bit vector with (length of input) * (length of prog) bits,
// to make sure it never explores the same (character position, instruction)
// state multiple times. This limits the search to run in time linear in
// the length of the test.
//
// backtrack is a fast replacement for the NFA code on small
// regexps when onepass cannot be used.

//line /usr/local/go/src/regexp/backtrack.go:15
package regexp

//line /usr/local/go/src/regexp/backtrack.go:15
import (
//line /usr/local/go/src/regexp/backtrack.go:15
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/regexp/backtrack.go:15
)
//line /usr/local/go/src/regexp/backtrack.go:15
import (
//line /usr/local/go/src/regexp/backtrack.go:15
	_atomic_ "sync/atomic"
//line /usr/local/go/src/regexp/backtrack.go:15
)

import (
	"regexp/syntax"
	"sync"
)

// A job is an entry on the backtracker's job stack. It holds
//line /usr/local/go/src/regexp/backtrack.go:22
// the instruction pc and the position in the input.
//line /usr/local/go/src/regexp/backtrack.go:24
type job struct {
	pc	uint32
	arg	bool
	pos	int
}

const (
	visitedBits		= 32
	maxBacktrackProg	= 500		// len(prog.Inst) <= max
	maxBacktrackVector	= 256 * 1024	// bit vector size <= max (bits)
)

// bitState holds state for the backtracker.
type bitState struct {
	end		int
	cap		[]int
	matchcap	[]int
	jobs		[]job
	visited		[]uint32

	inputs	inputs
}

var bitStatePool sync.Pool

func newBitState() *bitState {
//line /usr/local/go/src/regexp/backtrack.go:49
	_go_fuzz_dep_.CoverTab[64409]++
							b, ok := bitStatePool.Get().(*bitState)
							if !ok {
//line /usr/local/go/src/regexp/backtrack.go:51
		_go_fuzz_dep_.CoverTab[64411]++
								b = new(bitState)
//line /usr/local/go/src/regexp/backtrack.go:52
		// _ = "end of CoverTab[64411]"
	} else {
//line /usr/local/go/src/regexp/backtrack.go:53
		_go_fuzz_dep_.CoverTab[64412]++
//line /usr/local/go/src/regexp/backtrack.go:53
		// _ = "end of CoverTab[64412]"
//line /usr/local/go/src/regexp/backtrack.go:53
	}
//line /usr/local/go/src/regexp/backtrack.go:53
	// _ = "end of CoverTab[64409]"
//line /usr/local/go/src/regexp/backtrack.go:53
	_go_fuzz_dep_.CoverTab[64410]++
							return b
//line /usr/local/go/src/regexp/backtrack.go:54
	// _ = "end of CoverTab[64410]"
}

func freeBitState(b *bitState) {
//line /usr/local/go/src/regexp/backtrack.go:57
	_go_fuzz_dep_.CoverTab[64413]++
							b.inputs.clear()
							bitStatePool.Put(b)
//line /usr/local/go/src/regexp/backtrack.go:59
	// _ = "end of CoverTab[64413]"
}

// maxBitStateLen returns the maximum length of a string to search with
//line /usr/local/go/src/regexp/backtrack.go:62
// the backtracker using prog.
//line /usr/local/go/src/regexp/backtrack.go:64
func maxBitStateLen(prog *syntax.Prog) int {
//line /usr/local/go/src/regexp/backtrack.go:64
	_go_fuzz_dep_.CoverTab[64414]++
							if !shouldBacktrack(prog) {
//line /usr/local/go/src/regexp/backtrack.go:65
		_go_fuzz_dep_.CoverTab[64416]++
								return 0
//line /usr/local/go/src/regexp/backtrack.go:66
		// _ = "end of CoverTab[64416]"
	} else {
//line /usr/local/go/src/regexp/backtrack.go:67
		_go_fuzz_dep_.CoverTab[64417]++
//line /usr/local/go/src/regexp/backtrack.go:67
		// _ = "end of CoverTab[64417]"
//line /usr/local/go/src/regexp/backtrack.go:67
	}
//line /usr/local/go/src/regexp/backtrack.go:67
	// _ = "end of CoverTab[64414]"
//line /usr/local/go/src/regexp/backtrack.go:67
	_go_fuzz_dep_.CoverTab[64415]++
							return maxBacktrackVector / len(prog.Inst)
//line /usr/local/go/src/regexp/backtrack.go:68
	// _ = "end of CoverTab[64415]"
}

// shouldBacktrack reports whether the program is too
//line /usr/local/go/src/regexp/backtrack.go:71
// long for the backtracker to run.
//line /usr/local/go/src/regexp/backtrack.go:73
func shouldBacktrack(prog *syntax.Prog) bool {
//line /usr/local/go/src/regexp/backtrack.go:73
	_go_fuzz_dep_.CoverTab[64418]++
							return len(prog.Inst) <= maxBacktrackProg
//line /usr/local/go/src/regexp/backtrack.go:74
	// _ = "end of CoverTab[64418]"
}

// reset resets the state of the backtracker.
//line /usr/local/go/src/regexp/backtrack.go:77
// end is the end position in the input.
//line /usr/local/go/src/regexp/backtrack.go:77
// ncap is the number of captures.
//line /usr/local/go/src/regexp/backtrack.go:80
func (b *bitState) reset(prog *syntax.Prog, end int, ncap int) {
//line /usr/local/go/src/regexp/backtrack.go:80
	_go_fuzz_dep_.CoverTab[64419]++
							b.end = end

							if cap(b.jobs) == 0 {
//line /usr/local/go/src/regexp/backtrack.go:83
		_go_fuzz_dep_.CoverTab[64425]++
								b.jobs = make([]job, 0, 256)
//line /usr/local/go/src/regexp/backtrack.go:84
		// _ = "end of CoverTab[64425]"
	} else {
//line /usr/local/go/src/regexp/backtrack.go:85
		_go_fuzz_dep_.CoverTab[64426]++
								b.jobs = b.jobs[:0]
//line /usr/local/go/src/regexp/backtrack.go:86
		// _ = "end of CoverTab[64426]"
	}
//line /usr/local/go/src/regexp/backtrack.go:87
	// _ = "end of CoverTab[64419]"
//line /usr/local/go/src/regexp/backtrack.go:87
	_go_fuzz_dep_.CoverTab[64420]++

							visitedSize := (len(prog.Inst)*(end+1) + visitedBits - 1) / visitedBits
							if cap(b.visited) < visitedSize {
//line /usr/local/go/src/regexp/backtrack.go:90
		_go_fuzz_dep_.CoverTab[64427]++
								b.visited = make([]uint32, visitedSize, maxBacktrackVector/visitedBits)
//line /usr/local/go/src/regexp/backtrack.go:91
		// _ = "end of CoverTab[64427]"
	} else {
//line /usr/local/go/src/regexp/backtrack.go:92
		_go_fuzz_dep_.CoverTab[64428]++
								b.visited = b.visited[:visitedSize]
								for i := range b.visited {
//line /usr/local/go/src/regexp/backtrack.go:94
			_go_fuzz_dep_.CoverTab[64429]++
									b.visited[i] = 0
//line /usr/local/go/src/regexp/backtrack.go:95
			// _ = "end of CoverTab[64429]"
		}
//line /usr/local/go/src/regexp/backtrack.go:96
		// _ = "end of CoverTab[64428]"
	}
//line /usr/local/go/src/regexp/backtrack.go:97
	// _ = "end of CoverTab[64420]"
//line /usr/local/go/src/regexp/backtrack.go:97
	_go_fuzz_dep_.CoverTab[64421]++

							if cap(b.cap) < ncap {
//line /usr/local/go/src/regexp/backtrack.go:99
		_go_fuzz_dep_.CoverTab[64430]++
								b.cap = make([]int, ncap)
//line /usr/local/go/src/regexp/backtrack.go:100
		// _ = "end of CoverTab[64430]"
	} else {
//line /usr/local/go/src/regexp/backtrack.go:101
		_go_fuzz_dep_.CoverTab[64431]++
								b.cap = b.cap[:ncap]
//line /usr/local/go/src/regexp/backtrack.go:102
		// _ = "end of CoverTab[64431]"
	}
//line /usr/local/go/src/regexp/backtrack.go:103
	// _ = "end of CoverTab[64421]"
//line /usr/local/go/src/regexp/backtrack.go:103
	_go_fuzz_dep_.CoverTab[64422]++
							for i := range b.cap {
//line /usr/local/go/src/regexp/backtrack.go:104
		_go_fuzz_dep_.CoverTab[64432]++
								b.cap[i] = -1
//line /usr/local/go/src/regexp/backtrack.go:105
		// _ = "end of CoverTab[64432]"
	}
//line /usr/local/go/src/regexp/backtrack.go:106
	// _ = "end of CoverTab[64422]"
//line /usr/local/go/src/regexp/backtrack.go:106
	_go_fuzz_dep_.CoverTab[64423]++

							if cap(b.matchcap) < ncap {
//line /usr/local/go/src/regexp/backtrack.go:108
		_go_fuzz_dep_.CoverTab[64433]++
								b.matchcap = make([]int, ncap)
//line /usr/local/go/src/regexp/backtrack.go:109
		// _ = "end of CoverTab[64433]"
	} else {
//line /usr/local/go/src/regexp/backtrack.go:110
		_go_fuzz_dep_.CoverTab[64434]++
								b.matchcap = b.matchcap[:ncap]
//line /usr/local/go/src/regexp/backtrack.go:111
		// _ = "end of CoverTab[64434]"
	}
//line /usr/local/go/src/regexp/backtrack.go:112
	// _ = "end of CoverTab[64423]"
//line /usr/local/go/src/regexp/backtrack.go:112
	_go_fuzz_dep_.CoverTab[64424]++
							for i := range b.matchcap {
//line /usr/local/go/src/regexp/backtrack.go:113
		_go_fuzz_dep_.CoverTab[64435]++
								b.matchcap[i] = -1
//line /usr/local/go/src/regexp/backtrack.go:114
		// _ = "end of CoverTab[64435]"
	}
//line /usr/local/go/src/regexp/backtrack.go:115
	// _ = "end of CoverTab[64424]"
}

// shouldVisit reports whether the combination of (pc, pos) has not
//line /usr/local/go/src/regexp/backtrack.go:118
// been visited yet.
//line /usr/local/go/src/regexp/backtrack.go:120
func (b *bitState) shouldVisit(pc uint32, pos int) bool {
//line /usr/local/go/src/regexp/backtrack.go:120
	_go_fuzz_dep_.CoverTab[64436]++
							n := uint(int(pc)*(b.end+1) + pos)
							if b.visited[n/visitedBits]&(1<<(n&(visitedBits-1))) != 0 {
//line /usr/local/go/src/regexp/backtrack.go:122
		_go_fuzz_dep_.CoverTab[64438]++
								return false
//line /usr/local/go/src/regexp/backtrack.go:123
		// _ = "end of CoverTab[64438]"
	} else {
//line /usr/local/go/src/regexp/backtrack.go:124
		_go_fuzz_dep_.CoverTab[64439]++
//line /usr/local/go/src/regexp/backtrack.go:124
		// _ = "end of CoverTab[64439]"
//line /usr/local/go/src/regexp/backtrack.go:124
	}
//line /usr/local/go/src/regexp/backtrack.go:124
	// _ = "end of CoverTab[64436]"
//line /usr/local/go/src/regexp/backtrack.go:124
	_go_fuzz_dep_.CoverTab[64437]++
							b.visited[n/visitedBits] |= 1 << (n & (visitedBits - 1))
							return true
//line /usr/local/go/src/regexp/backtrack.go:126
	// _ = "end of CoverTab[64437]"
}

// push pushes (pc, pos, arg) onto the job stack if it should be
//line /usr/local/go/src/regexp/backtrack.go:129
// visited.
//line /usr/local/go/src/regexp/backtrack.go:131
func (b *bitState) push(re *Regexp, pc uint32, pos int, arg bool) {
//line /usr/local/go/src/regexp/backtrack.go:131
	_go_fuzz_dep_.CoverTab[64440]++

//line /usr/local/go/src/regexp/backtrack.go:134
	if re.prog.Inst[pc].Op != syntax.InstFail && func() bool {
//line /usr/local/go/src/regexp/backtrack.go:134
		_go_fuzz_dep_.CoverTab[64441]++
//line /usr/local/go/src/regexp/backtrack.go:134
		return (arg || func() bool {
//line /usr/local/go/src/regexp/backtrack.go:134
			_go_fuzz_dep_.CoverTab[64442]++
//line /usr/local/go/src/regexp/backtrack.go:134
			return b.shouldVisit(pc, pos)
//line /usr/local/go/src/regexp/backtrack.go:134
			// _ = "end of CoverTab[64442]"
//line /usr/local/go/src/regexp/backtrack.go:134
		}())
//line /usr/local/go/src/regexp/backtrack.go:134
		// _ = "end of CoverTab[64441]"
//line /usr/local/go/src/regexp/backtrack.go:134
	}() {
//line /usr/local/go/src/regexp/backtrack.go:134
		_go_fuzz_dep_.CoverTab[64443]++
								b.jobs = append(b.jobs, job{pc: pc, arg: arg, pos: pos})
//line /usr/local/go/src/regexp/backtrack.go:135
		// _ = "end of CoverTab[64443]"
	} else {
//line /usr/local/go/src/regexp/backtrack.go:136
		_go_fuzz_dep_.CoverTab[64444]++
//line /usr/local/go/src/regexp/backtrack.go:136
		// _ = "end of CoverTab[64444]"
//line /usr/local/go/src/regexp/backtrack.go:136
	}
//line /usr/local/go/src/regexp/backtrack.go:136
	// _ = "end of CoverTab[64440]"
}

// tryBacktrack runs a backtracking search starting at pos.
func (re *Regexp) tryBacktrack(b *bitState, i input, pc uint32, pos int) bool {
//line /usr/local/go/src/regexp/backtrack.go:140
	_go_fuzz_dep_.CoverTab[64445]++
							longest := re.longest

							b.push(re, pc, pos, false)
							for len(b.jobs) > 0 {
//line /usr/local/go/src/regexp/backtrack.go:144
		_go_fuzz_dep_.CoverTab[64447]++
								l := len(b.jobs) - 1

								pc := b.jobs[l].pc
								pos := b.jobs[l].pos
								arg := b.jobs[l].arg
								b.jobs = b.jobs[:l]

//line /usr/local/go/src/regexp/backtrack.go:159
		goto Skip
//line /usr/local/go/src/regexp/backtrack.go:159
		// _ = "end of CoverTab[64447]"
//line /usr/local/go/src/regexp/backtrack.go:159
		_go_fuzz_dep_.CoverTab[64448]++
	CheckAndLoop:
		if !b.shouldVisit(pc, pos) {
//line /usr/local/go/src/regexp/backtrack.go:161
			_go_fuzz_dep_.CoverTab[64450]++
									continue
//line /usr/local/go/src/regexp/backtrack.go:162
			// _ = "end of CoverTab[64450]"
		} else {
//line /usr/local/go/src/regexp/backtrack.go:163
			_go_fuzz_dep_.CoverTab[64451]++
//line /usr/local/go/src/regexp/backtrack.go:163
			// _ = "end of CoverTab[64451]"
//line /usr/local/go/src/regexp/backtrack.go:163
		}
//line /usr/local/go/src/regexp/backtrack.go:163
		// _ = "end of CoverTab[64448]"
//line /usr/local/go/src/regexp/backtrack.go:163
		_go_fuzz_dep_.CoverTab[64449]++
	Skip:

		inst := &re.prog.Inst[pc]

		switch inst.Op {
		default:
//line /usr/local/go/src/regexp/backtrack.go:169
			_go_fuzz_dep_.CoverTab[64452]++
									panic("bad inst")
//line /usr/local/go/src/regexp/backtrack.go:170
			// _ = "end of CoverTab[64452]"
		case syntax.InstFail:
//line /usr/local/go/src/regexp/backtrack.go:171
			_go_fuzz_dep_.CoverTab[64453]++
									panic("unexpected InstFail")
//line /usr/local/go/src/regexp/backtrack.go:172
			// _ = "end of CoverTab[64453]"
		case syntax.InstAlt:
//line /usr/local/go/src/regexp/backtrack.go:173
			_go_fuzz_dep_.CoverTab[64454]++

//line /usr/local/go/src/regexp/backtrack.go:182
			if arg {
//line /usr/local/go/src/regexp/backtrack.go:182
				_go_fuzz_dep_.CoverTab[64475]++

										arg = false
										pc = inst.Arg
										goto CheckAndLoop
//line /usr/local/go/src/regexp/backtrack.go:186
				// _ = "end of CoverTab[64475]"
			} else {
//line /usr/local/go/src/regexp/backtrack.go:187
				_go_fuzz_dep_.CoverTab[64476]++
										b.push(re, pc, pos, true)
										pc = inst.Out
										goto CheckAndLoop
//line /usr/local/go/src/regexp/backtrack.go:190
				// _ = "end of CoverTab[64476]"
			}
//line /usr/local/go/src/regexp/backtrack.go:191
			// _ = "end of CoverTab[64454]"

		case syntax.InstAltMatch:
//line /usr/local/go/src/regexp/backtrack.go:193
			_go_fuzz_dep_.CoverTab[64455]++

									switch re.prog.Inst[inst.Out].Op {
			case syntax.InstRune, syntax.InstRune1, syntax.InstRuneAny, syntax.InstRuneAnyNotNL:
//line /usr/local/go/src/regexp/backtrack.go:196
				_go_fuzz_dep_.CoverTab[64477]++

										b.push(re, inst.Arg, pos, false)
										pc = inst.Arg
										pos = b.end
										goto CheckAndLoop
//line /usr/local/go/src/regexp/backtrack.go:201
				// _ = "end of CoverTab[64477]"
//line /usr/local/go/src/regexp/backtrack.go:201
			default:
//line /usr/local/go/src/regexp/backtrack.go:201
				_go_fuzz_dep_.CoverTab[64478]++
//line /usr/local/go/src/regexp/backtrack.go:201
				// _ = "end of CoverTab[64478]"
			}
//line /usr/local/go/src/regexp/backtrack.go:202
			// _ = "end of CoverTab[64455]"
//line /usr/local/go/src/regexp/backtrack.go:202
			_go_fuzz_dep_.CoverTab[64456]++

									b.push(re, inst.Out, b.end, false)
									pc = inst.Out
									goto CheckAndLoop
//line /usr/local/go/src/regexp/backtrack.go:206
			// _ = "end of CoverTab[64456]"

		case syntax.InstRune:
//line /usr/local/go/src/regexp/backtrack.go:208
			_go_fuzz_dep_.CoverTab[64457]++
									r, width := i.step(pos)
									if !inst.MatchRune(r) {
//line /usr/local/go/src/regexp/backtrack.go:210
				_go_fuzz_dep_.CoverTab[64479]++
										continue
//line /usr/local/go/src/regexp/backtrack.go:211
				// _ = "end of CoverTab[64479]"
			} else {
//line /usr/local/go/src/regexp/backtrack.go:212
				_go_fuzz_dep_.CoverTab[64480]++
//line /usr/local/go/src/regexp/backtrack.go:212
				// _ = "end of CoverTab[64480]"
//line /usr/local/go/src/regexp/backtrack.go:212
			}
//line /usr/local/go/src/regexp/backtrack.go:212
			// _ = "end of CoverTab[64457]"
//line /usr/local/go/src/regexp/backtrack.go:212
			_go_fuzz_dep_.CoverTab[64458]++
									pos += width
									pc = inst.Out
									goto CheckAndLoop
//line /usr/local/go/src/regexp/backtrack.go:215
			// _ = "end of CoverTab[64458]"

		case syntax.InstRune1:
//line /usr/local/go/src/regexp/backtrack.go:217
			_go_fuzz_dep_.CoverTab[64459]++
									r, width := i.step(pos)
									if r != inst.Rune[0] {
//line /usr/local/go/src/regexp/backtrack.go:219
				_go_fuzz_dep_.CoverTab[64481]++
										continue
//line /usr/local/go/src/regexp/backtrack.go:220
				// _ = "end of CoverTab[64481]"
			} else {
//line /usr/local/go/src/regexp/backtrack.go:221
				_go_fuzz_dep_.CoverTab[64482]++
//line /usr/local/go/src/regexp/backtrack.go:221
				// _ = "end of CoverTab[64482]"
//line /usr/local/go/src/regexp/backtrack.go:221
			}
//line /usr/local/go/src/regexp/backtrack.go:221
			// _ = "end of CoverTab[64459]"
//line /usr/local/go/src/regexp/backtrack.go:221
			_go_fuzz_dep_.CoverTab[64460]++
									pos += width
									pc = inst.Out
									goto CheckAndLoop
//line /usr/local/go/src/regexp/backtrack.go:224
			// _ = "end of CoverTab[64460]"

		case syntax.InstRuneAnyNotNL:
//line /usr/local/go/src/regexp/backtrack.go:226
			_go_fuzz_dep_.CoverTab[64461]++
									r, width := i.step(pos)
									if r == '\n' || func() bool {
//line /usr/local/go/src/regexp/backtrack.go:228
				_go_fuzz_dep_.CoverTab[64483]++
//line /usr/local/go/src/regexp/backtrack.go:228
				return r == endOfText
//line /usr/local/go/src/regexp/backtrack.go:228
				// _ = "end of CoverTab[64483]"
//line /usr/local/go/src/regexp/backtrack.go:228
			}() {
//line /usr/local/go/src/regexp/backtrack.go:228
				_go_fuzz_dep_.CoverTab[64484]++
										continue
//line /usr/local/go/src/regexp/backtrack.go:229
				// _ = "end of CoverTab[64484]"
			} else {
//line /usr/local/go/src/regexp/backtrack.go:230
				_go_fuzz_dep_.CoverTab[64485]++
//line /usr/local/go/src/regexp/backtrack.go:230
				// _ = "end of CoverTab[64485]"
//line /usr/local/go/src/regexp/backtrack.go:230
			}
//line /usr/local/go/src/regexp/backtrack.go:230
			// _ = "end of CoverTab[64461]"
//line /usr/local/go/src/regexp/backtrack.go:230
			_go_fuzz_dep_.CoverTab[64462]++
									pos += width
									pc = inst.Out
									goto CheckAndLoop
//line /usr/local/go/src/regexp/backtrack.go:233
			// _ = "end of CoverTab[64462]"

		case syntax.InstRuneAny:
//line /usr/local/go/src/regexp/backtrack.go:235
			_go_fuzz_dep_.CoverTab[64463]++
									r, width := i.step(pos)
									if r == endOfText {
//line /usr/local/go/src/regexp/backtrack.go:237
				_go_fuzz_dep_.CoverTab[64486]++
										continue
//line /usr/local/go/src/regexp/backtrack.go:238
				// _ = "end of CoverTab[64486]"
			} else {
//line /usr/local/go/src/regexp/backtrack.go:239
				_go_fuzz_dep_.CoverTab[64487]++
//line /usr/local/go/src/regexp/backtrack.go:239
				// _ = "end of CoverTab[64487]"
//line /usr/local/go/src/regexp/backtrack.go:239
			}
//line /usr/local/go/src/regexp/backtrack.go:239
			// _ = "end of CoverTab[64463]"
//line /usr/local/go/src/regexp/backtrack.go:239
			_go_fuzz_dep_.CoverTab[64464]++
									pos += width
									pc = inst.Out
									goto CheckAndLoop
//line /usr/local/go/src/regexp/backtrack.go:242
			// _ = "end of CoverTab[64464]"

		case syntax.InstCapture:
//line /usr/local/go/src/regexp/backtrack.go:244
			_go_fuzz_dep_.CoverTab[64465]++
									if arg {
//line /usr/local/go/src/regexp/backtrack.go:245
				_go_fuzz_dep_.CoverTab[64488]++

										b.cap[inst.Arg] = pos
										continue
//line /usr/local/go/src/regexp/backtrack.go:248
				// _ = "end of CoverTab[64488]"
			} else {
//line /usr/local/go/src/regexp/backtrack.go:249
				_go_fuzz_dep_.CoverTab[64489]++
										if inst.Arg < uint32(len(b.cap)) {
//line /usr/local/go/src/regexp/backtrack.go:250
					_go_fuzz_dep_.CoverTab[64491]++

											b.push(re, pc, b.cap[inst.Arg], true)
											b.cap[inst.Arg] = pos
//line /usr/local/go/src/regexp/backtrack.go:253
					// _ = "end of CoverTab[64491]"
				} else {
//line /usr/local/go/src/regexp/backtrack.go:254
					_go_fuzz_dep_.CoverTab[64492]++
//line /usr/local/go/src/regexp/backtrack.go:254
					// _ = "end of CoverTab[64492]"
//line /usr/local/go/src/regexp/backtrack.go:254
				}
//line /usr/local/go/src/regexp/backtrack.go:254
				// _ = "end of CoverTab[64489]"
//line /usr/local/go/src/regexp/backtrack.go:254
				_go_fuzz_dep_.CoverTab[64490]++
										pc = inst.Out
										goto CheckAndLoop
//line /usr/local/go/src/regexp/backtrack.go:256
				// _ = "end of CoverTab[64490]"
			}
//line /usr/local/go/src/regexp/backtrack.go:257
			// _ = "end of CoverTab[64465]"

		case syntax.InstEmptyWidth:
//line /usr/local/go/src/regexp/backtrack.go:259
			_go_fuzz_dep_.CoverTab[64466]++
									flag := i.context(pos)
									if !flag.match(syntax.EmptyOp(inst.Arg)) {
//line /usr/local/go/src/regexp/backtrack.go:261
				_go_fuzz_dep_.CoverTab[64493]++
										continue
//line /usr/local/go/src/regexp/backtrack.go:262
				// _ = "end of CoverTab[64493]"
			} else {
//line /usr/local/go/src/regexp/backtrack.go:263
				_go_fuzz_dep_.CoverTab[64494]++
//line /usr/local/go/src/regexp/backtrack.go:263
				// _ = "end of CoverTab[64494]"
//line /usr/local/go/src/regexp/backtrack.go:263
			}
//line /usr/local/go/src/regexp/backtrack.go:263
			// _ = "end of CoverTab[64466]"
//line /usr/local/go/src/regexp/backtrack.go:263
			_go_fuzz_dep_.CoverTab[64467]++
									pc = inst.Out
									goto CheckAndLoop
//line /usr/local/go/src/regexp/backtrack.go:265
			// _ = "end of CoverTab[64467]"

		case syntax.InstNop:
//line /usr/local/go/src/regexp/backtrack.go:267
			_go_fuzz_dep_.CoverTab[64468]++
									pc = inst.Out
									goto CheckAndLoop
//line /usr/local/go/src/regexp/backtrack.go:269
			// _ = "end of CoverTab[64468]"

		case syntax.InstMatch:
//line /usr/local/go/src/regexp/backtrack.go:271
			_go_fuzz_dep_.CoverTab[64469]++

//line /usr/local/go/src/regexp/backtrack.go:274
			if len(b.cap) == 0 {
//line /usr/local/go/src/regexp/backtrack.go:274
				_go_fuzz_dep_.CoverTab[64495]++
										return true
//line /usr/local/go/src/regexp/backtrack.go:275
				// _ = "end of CoverTab[64495]"
			} else {
//line /usr/local/go/src/regexp/backtrack.go:276
				_go_fuzz_dep_.CoverTab[64496]++
//line /usr/local/go/src/regexp/backtrack.go:276
				// _ = "end of CoverTab[64496]"
//line /usr/local/go/src/regexp/backtrack.go:276
			}
//line /usr/local/go/src/regexp/backtrack.go:276
			// _ = "end of CoverTab[64469]"
//line /usr/local/go/src/regexp/backtrack.go:276
			_go_fuzz_dep_.CoverTab[64470]++

//line /usr/local/go/src/regexp/backtrack.go:281
			if len(b.cap) > 1 {
//line /usr/local/go/src/regexp/backtrack.go:281
				_go_fuzz_dep_.CoverTab[64497]++
										b.cap[1] = pos
//line /usr/local/go/src/regexp/backtrack.go:282
				// _ = "end of CoverTab[64497]"
			} else {
//line /usr/local/go/src/regexp/backtrack.go:283
				_go_fuzz_dep_.CoverTab[64498]++
//line /usr/local/go/src/regexp/backtrack.go:283
				// _ = "end of CoverTab[64498]"
//line /usr/local/go/src/regexp/backtrack.go:283
			}
//line /usr/local/go/src/regexp/backtrack.go:283
			// _ = "end of CoverTab[64470]"
//line /usr/local/go/src/regexp/backtrack.go:283
			_go_fuzz_dep_.CoverTab[64471]++
									if old := b.matchcap[1]; old == -1 || func() bool {
//line /usr/local/go/src/regexp/backtrack.go:284
				_go_fuzz_dep_.CoverTab[64499]++
//line /usr/local/go/src/regexp/backtrack.go:284
				return (longest && func() bool {
//line /usr/local/go/src/regexp/backtrack.go:284
					_go_fuzz_dep_.CoverTab[64500]++
//line /usr/local/go/src/regexp/backtrack.go:284
					return pos > 0
//line /usr/local/go/src/regexp/backtrack.go:284
					// _ = "end of CoverTab[64500]"
//line /usr/local/go/src/regexp/backtrack.go:284
				}() && func() bool {
//line /usr/local/go/src/regexp/backtrack.go:284
					_go_fuzz_dep_.CoverTab[64501]++
//line /usr/local/go/src/regexp/backtrack.go:284
					return pos > old
//line /usr/local/go/src/regexp/backtrack.go:284
					// _ = "end of CoverTab[64501]"
//line /usr/local/go/src/regexp/backtrack.go:284
				}())
//line /usr/local/go/src/regexp/backtrack.go:284
				// _ = "end of CoverTab[64499]"
//line /usr/local/go/src/regexp/backtrack.go:284
			}() {
//line /usr/local/go/src/regexp/backtrack.go:284
				_go_fuzz_dep_.CoverTab[64502]++
										copy(b.matchcap, b.cap)
//line /usr/local/go/src/regexp/backtrack.go:285
				// _ = "end of CoverTab[64502]"
			} else {
//line /usr/local/go/src/regexp/backtrack.go:286
				_go_fuzz_dep_.CoverTab[64503]++
//line /usr/local/go/src/regexp/backtrack.go:286
				// _ = "end of CoverTab[64503]"
//line /usr/local/go/src/regexp/backtrack.go:286
			}
//line /usr/local/go/src/regexp/backtrack.go:286
			// _ = "end of CoverTab[64471]"
//line /usr/local/go/src/regexp/backtrack.go:286
			_go_fuzz_dep_.CoverTab[64472]++

//line /usr/local/go/src/regexp/backtrack.go:289
			if !longest {
//line /usr/local/go/src/regexp/backtrack.go:289
				_go_fuzz_dep_.CoverTab[64504]++
										return true
//line /usr/local/go/src/regexp/backtrack.go:290
				// _ = "end of CoverTab[64504]"
			} else {
//line /usr/local/go/src/regexp/backtrack.go:291
				_go_fuzz_dep_.CoverTab[64505]++
//line /usr/local/go/src/regexp/backtrack.go:291
				// _ = "end of CoverTab[64505]"
//line /usr/local/go/src/regexp/backtrack.go:291
			}
//line /usr/local/go/src/regexp/backtrack.go:291
			// _ = "end of CoverTab[64472]"
//line /usr/local/go/src/regexp/backtrack.go:291
			_go_fuzz_dep_.CoverTab[64473]++

//line /usr/local/go/src/regexp/backtrack.go:294
			if pos == b.end {
//line /usr/local/go/src/regexp/backtrack.go:294
				_go_fuzz_dep_.CoverTab[64506]++
										return true
//line /usr/local/go/src/regexp/backtrack.go:295
				// _ = "end of CoverTab[64506]"
			} else {
//line /usr/local/go/src/regexp/backtrack.go:296
				_go_fuzz_dep_.CoverTab[64507]++
//line /usr/local/go/src/regexp/backtrack.go:296
				// _ = "end of CoverTab[64507]"
//line /usr/local/go/src/regexp/backtrack.go:296
			}
//line /usr/local/go/src/regexp/backtrack.go:296
			// _ = "end of CoverTab[64473]"
//line /usr/local/go/src/regexp/backtrack.go:296
			_go_fuzz_dep_.CoverTab[64474]++

//line /usr/local/go/src/regexp/backtrack.go:299
			continue
//line /usr/local/go/src/regexp/backtrack.go:299
			// _ = "end of CoverTab[64474]"
		}
//line /usr/local/go/src/regexp/backtrack.go:300
		// _ = "end of CoverTab[64449]"
	}
//line /usr/local/go/src/regexp/backtrack.go:301
	// _ = "end of CoverTab[64445]"
//line /usr/local/go/src/regexp/backtrack.go:301
	_go_fuzz_dep_.CoverTab[64446]++

							return longest && func() bool {
//line /usr/local/go/src/regexp/backtrack.go:303
		_go_fuzz_dep_.CoverTab[64508]++
//line /usr/local/go/src/regexp/backtrack.go:303
		return len(b.matchcap) > 1
//line /usr/local/go/src/regexp/backtrack.go:303
		// _ = "end of CoverTab[64508]"
//line /usr/local/go/src/regexp/backtrack.go:303
	}() && func() bool {
//line /usr/local/go/src/regexp/backtrack.go:303
		_go_fuzz_dep_.CoverTab[64509]++
//line /usr/local/go/src/regexp/backtrack.go:303
		return b.matchcap[1] >= 0
//line /usr/local/go/src/regexp/backtrack.go:303
		// _ = "end of CoverTab[64509]"
//line /usr/local/go/src/regexp/backtrack.go:303
	}()
//line /usr/local/go/src/regexp/backtrack.go:303
	// _ = "end of CoverTab[64446]"
}

// backtrack runs a backtracking search of prog on the input starting at pos.
func (re *Regexp) backtrack(ib []byte, is string, pos int, ncap int, dstCap []int) []int {
//line /usr/local/go/src/regexp/backtrack.go:307
	_go_fuzz_dep_.CoverTab[64510]++
							startCond := re.cond
							if startCond == ^syntax.EmptyOp(0) {
//line /usr/local/go/src/regexp/backtrack.go:309
		_go_fuzz_dep_.CoverTab[64514]++
								return nil
//line /usr/local/go/src/regexp/backtrack.go:310
		// _ = "end of CoverTab[64514]"
	} else {
//line /usr/local/go/src/regexp/backtrack.go:311
		_go_fuzz_dep_.CoverTab[64515]++
//line /usr/local/go/src/regexp/backtrack.go:311
		// _ = "end of CoverTab[64515]"
//line /usr/local/go/src/regexp/backtrack.go:311
	}
//line /usr/local/go/src/regexp/backtrack.go:311
	// _ = "end of CoverTab[64510]"
//line /usr/local/go/src/regexp/backtrack.go:311
	_go_fuzz_dep_.CoverTab[64511]++
							if startCond&syntax.EmptyBeginText != 0 && func() bool {
//line /usr/local/go/src/regexp/backtrack.go:312
		_go_fuzz_dep_.CoverTab[64516]++
//line /usr/local/go/src/regexp/backtrack.go:312
		return pos != 0
//line /usr/local/go/src/regexp/backtrack.go:312
		// _ = "end of CoverTab[64516]"
//line /usr/local/go/src/regexp/backtrack.go:312
	}() {
//line /usr/local/go/src/regexp/backtrack.go:312
		_go_fuzz_dep_.CoverTab[64517]++

								return nil
//line /usr/local/go/src/regexp/backtrack.go:314
		// _ = "end of CoverTab[64517]"
	} else {
//line /usr/local/go/src/regexp/backtrack.go:315
		_go_fuzz_dep_.CoverTab[64518]++
//line /usr/local/go/src/regexp/backtrack.go:315
		// _ = "end of CoverTab[64518]"
//line /usr/local/go/src/regexp/backtrack.go:315
	}
//line /usr/local/go/src/regexp/backtrack.go:315
	// _ = "end of CoverTab[64511]"
//line /usr/local/go/src/regexp/backtrack.go:315
	_go_fuzz_dep_.CoverTab[64512]++

							b := newBitState()
							i, end := b.inputs.init(nil, ib, is)
							b.reset(re.prog, end, ncap)

//line /usr/local/go/src/regexp/backtrack.go:322
	if startCond&syntax.EmptyBeginText != 0 {
//line /usr/local/go/src/regexp/backtrack.go:322
		_go_fuzz_dep_.CoverTab[64519]++
								if len(b.cap) > 0 {
//line /usr/local/go/src/regexp/backtrack.go:323
			_go_fuzz_dep_.CoverTab[64521]++
									b.cap[0] = pos
//line /usr/local/go/src/regexp/backtrack.go:324
			// _ = "end of CoverTab[64521]"
		} else {
//line /usr/local/go/src/regexp/backtrack.go:325
			_go_fuzz_dep_.CoverTab[64522]++
//line /usr/local/go/src/regexp/backtrack.go:325
			// _ = "end of CoverTab[64522]"
//line /usr/local/go/src/regexp/backtrack.go:325
		}
//line /usr/local/go/src/regexp/backtrack.go:325
		// _ = "end of CoverTab[64519]"
//line /usr/local/go/src/regexp/backtrack.go:325
		_go_fuzz_dep_.CoverTab[64520]++
								if !re.tryBacktrack(b, i, uint32(re.prog.Start), pos) {
//line /usr/local/go/src/regexp/backtrack.go:326
			_go_fuzz_dep_.CoverTab[64523]++
									freeBitState(b)
									return nil
//line /usr/local/go/src/regexp/backtrack.go:328
			// _ = "end of CoverTab[64523]"
		} else {
//line /usr/local/go/src/regexp/backtrack.go:329
			_go_fuzz_dep_.CoverTab[64524]++
//line /usr/local/go/src/regexp/backtrack.go:329
			// _ = "end of CoverTab[64524]"
//line /usr/local/go/src/regexp/backtrack.go:329
		}
//line /usr/local/go/src/regexp/backtrack.go:329
		// _ = "end of CoverTab[64520]"
	} else {
//line /usr/local/go/src/regexp/backtrack.go:330
		_go_fuzz_dep_.CoverTab[64525]++

//line /usr/local/go/src/regexp/backtrack.go:338
		width := -1
		for ; pos <= end && func() bool {
//line /usr/local/go/src/regexp/backtrack.go:339
			_go_fuzz_dep_.CoverTab[64527]++
//line /usr/local/go/src/regexp/backtrack.go:339
			return width != 0
//line /usr/local/go/src/regexp/backtrack.go:339
			// _ = "end of CoverTab[64527]"
//line /usr/local/go/src/regexp/backtrack.go:339
		}(); pos += width {
//line /usr/local/go/src/regexp/backtrack.go:339
			_go_fuzz_dep_.CoverTab[64528]++
									if len(re.prefix) > 0 {
//line /usr/local/go/src/regexp/backtrack.go:340
				_go_fuzz_dep_.CoverTab[64532]++

										advance := i.index(re, pos)
										if advance < 0 {
//line /usr/local/go/src/regexp/backtrack.go:343
					_go_fuzz_dep_.CoverTab[64534]++
											freeBitState(b)
											return nil
//line /usr/local/go/src/regexp/backtrack.go:345
					// _ = "end of CoverTab[64534]"
				} else {
//line /usr/local/go/src/regexp/backtrack.go:346
					_go_fuzz_dep_.CoverTab[64535]++
//line /usr/local/go/src/regexp/backtrack.go:346
					// _ = "end of CoverTab[64535]"
//line /usr/local/go/src/regexp/backtrack.go:346
				}
//line /usr/local/go/src/regexp/backtrack.go:346
				// _ = "end of CoverTab[64532]"
//line /usr/local/go/src/regexp/backtrack.go:346
				_go_fuzz_dep_.CoverTab[64533]++
										pos += advance
//line /usr/local/go/src/regexp/backtrack.go:347
				// _ = "end of CoverTab[64533]"
			} else {
//line /usr/local/go/src/regexp/backtrack.go:348
				_go_fuzz_dep_.CoverTab[64536]++
//line /usr/local/go/src/regexp/backtrack.go:348
				// _ = "end of CoverTab[64536]"
//line /usr/local/go/src/regexp/backtrack.go:348
			}
//line /usr/local/go/src/regexp/backtrack.go:348
			// _ = "end of CoverTab[64528]"
//line /usr/local/go/src/regexp/backtrack.go:348
			_go_fuzz_dep_.CoverTab[64529]++

									if len(b.cap) > 0 {
//line /usr/local/go/src/regexp/backtrack.go:350
				_go_fuzz_dep_.CoverTab[64537]++
										b.cap[0] = pos
//line /usr/local/go/src/regexp/backtrack.go:351
				// _ = "end of CoverTab[64537]"
			} else {
//line /usr/local/go/src/regexp/backtrack.go:352
				_go_fuzz_dep_.CoverTab[64538]++
//line /usr/local/go/src/regexp/backtrack.go:352
				// _ = "end of CoverTab[64538]"
//line /usr/local/go/src/regexp/backtrack.go:352
			}
//line /usr/local/go/src/regexp/backtrack.go:352
			// _ = "end of CoverTab[64529]"
//line /usr/local/go/src/regexp/backtrack.go:352
			_go_fuzz_dep_.CoverTab[64530]++
									if re.tryBacktrack(b, i, uint32(re.prog.Start), pos) {
//line /usr/local/go/src/regexp/backtrack.go:353
				_go_fuzz_dep_.CoverTab[64539]++

										goto Match
//line /usr/local/go/src/regexp/backtrack.go:355
				// _ = "end of CoverTab[64539]"
			} else {
//line /usr/local/go/src/regexp/backtrack.go:356
				_go_fuzz_dep_.CoverTab[64540]++
//line /usr/local/go/src/regexp/backtrack.go:356
				// _ = "end of CoverTab[64540]"
//line /usr/local/go/src/regexp/backtrack.go:356
			}
//line /usr/local/go/src/regexp/backtrack.go:356
			// _ = "end of CoverTab[64530]"
//line /usr/local/go/src/regexp/backtrack.go:356
			_go_fuzz_dep_.CoverTab[64531]++
									_, width = i.step(pos)
//line /usr/local/go/src/regexp/backtrack.go:357
			// _ = "end of CoverTab[64531]"
		}
//line /usr/local/go/src/regexp/backtrack.go:358
		// _ = "end of CoverTab[64525]"
//line /usr/local/go/src/regexp/backtrack.go:358
		_go_fuzz_dep_.CoverTab[64526]++
								freeBitState(b)
								return nil
//line /usr/local/go/src/regexp/backtrack.go:360
		// _ = "end of CoverTab[64526]"
	}
//line /usr/local/go/src/regexp/backtrack.go:361
	// _ = "end of CoverTab[64512]"
//line /usr/local/go/src/regexp/backtrack.go:361
	_go_fuzz_dep_.CoverTab[64513]++

Match:
							dstCap = append(dstCap, b.matchcap...)
							freeBitState(b)
							return dstCap
//line /usr/local/go/src/regexp/backtrack.go:366
	// _ = "end of CoverTab[64513]"
}

//line /usr/local/go/src/regexp/backtrack.go:367
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/regexp/backtrack.go:367
var _ = _go_fuzz_dep_.CoverTab
