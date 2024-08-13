// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/regexp/onepass.go:5
package regexp

//line /usr/local/go/src/regexp/onepass.go:5
import (
//line /usr/local/go/src/regexp/onepass.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/regexp/onepass.go:5
)
//line /usr/local/go/src/regexp/onepass.go:5
import (
//line /usr/local/go/src/regexp/onepass.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/regexp/onepass.go:5
)

import (
	"regexp/syntax"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

//line /usr/local/go/src/regexp/onepass.go:21
// A onePassProg is a compiled one-pass regular expression program.
//line /usr/local/go/src/regexp/onepass.go:21
// It is the same as syntax.Prog except for the use of onePassInst.
//line /usr/local/go/src/regexp/onepass.go:23
type onePassProg struct {
	Inst	[]onePassInst
	Start	int	// index of start instruction
	NumCap	int	// number of InstCapture insts in re
}

// A onePassInst is a single instruction in a one-pass regular expression program.
//line /usr/local/go/src/regexp/onepass.go:29
// It is the same as syntax.Inst except for the new 'Next' field.
//line /usr/local/go/src/regexp/onepass.go:31
type onePassInst struct {
	syntax.Inst
	Next	[]uint32
}

// onePassPrefix returns a literal string that all matches for the
//line /usr/local/go/src/regexp/onepass.go:36
// regexp must start with. Complete is true if the prefix
//line /usr/local/go/src/regexp/onepass.go:36
// is the entire match. Pc is the index of the last rune instruction
//line /usr/local/go/src/regexp/onepass.go:36
// in the string. The onePassPrefix skips over the mandatory
//line /usr/local/go/src/regexp/onepass.go:36
// EmptyBeginText.
//line /usr/local/go/src/regexp/onepass.go:41
func onePassPrefix(p *syntax.Prog) (prefix string, complete bool, pc uint32) {
//line /usr/local/go/src/regexp/onepass.go:41
	_go_fuzz_dep_.CoverTab[64794]++
						i := &p.Inst[p.Start]
						if i.Op != syntax.InstEmptyWidth || func() bool {
//line /usr/local/go/src/regexp/onepass.go:43
		_go_fuzz_dep_.CoverTab[64800]++
//line /usr/local/go/src/regexp/onepass.go:43
		return (syntax.EmptyOp(i.Arg))&syntax.EmptyBeginText == 0
//line /usr/local/go/src/regexp/onepass.go:43
		// _ = "end of CoverTab[64800]"
//line /usr/local/go/src/regexp/onepass.go:43
	}() {
//line /usr/local/go/src/regexp/onepass.go:43
		_go_fuzz_dep_.CoverTab[64801]++
							return "", i.Op == syntax.InstMatch, uint32(p.Start)
//line /usr/local/go/src/regexp/onepass.go:44
		// _ = "end of CoverTab[64801]"
	} else {
//line /usr/local/go/src/regexp/onepass.go:45
		_go_fuzz_dep_.CoverTab[64802]++
//line /usr/local/go/src/regexp/onepass.go:45
		// _ = "end of CoverTab[64802]"
//line /usr/local/go/src/regexp/onepass.go:45
	}
//line /usr/local/go/src/regexp/onepass.go:45
	// _ = "end of CoverTab[64794]"
//line /usr/local/go/src/regexp/onepass.go:45
	_go_fuzz_dep_.CoverTab[64795]++
						pc = i.Out
						i = &p.Inst[pc]
						for i.Op == syntax.InstNop {
//line /usr/local/go/src/regexp/onepass.go:48
		_go_fuzz_dep_.CoverTab[64803]++
							pc = i.Out
							i = &p.Inst[pc]
//line /usr/local/go/src/regexp/onepass.go:50
		// _ = "end of CoverTab[64803]"
	}
//line /usr/local/go/src/regexp/onepass.go:51
	// _ = "end of CoverTab[64795]"
//line /usr/local/go/src/regexp/onepass.go:51
	_go_fuzz_dep_.CoverTab[64796]++

						if iop(i) != syntax.InstRune || func() bool {
//line /usr/local/go/src/regexp/onepass.go:53
		_go_fuzz_dep_.CoverTab[64804]++
//line /usr/local/go/src/regexp/onepass.go:53
		return len(i.Rune) != 1
//line /usr/local/go/src/regexp/onepass.go:53
		// _ = "end of CoverTab[64804]"
//line /usr/local/go/src/regexp/onepass.go:53
	}() {
//line /usr/local/go/src/regexp/onepass.go:53
		_go_fuzz_dep_.CoverTab[64805]++
							return "", i.Op == syntax.InstMatch, uint32(p.Start)
//line /usr/local/go/src/regexp/onepass.go:54
		// _ = "end of CoverTab[64805]"
	} else {
//line /usr/local/go/src/regexp/onepass.go:55
		_go_fuzz_dep_.CoverTab[64806]++
//line /usr/local/go/src/regexp/onepass.go:55
		// _ = "end of CoverTab[64806]"
//line /usr/local/go/src/regexp/onepass.go:55
	}
//line /usr/local/go/src/regexp/onepass.go:55
	// _ = "end of CoverTab[64796]"
//line /usr/local/go/src/regexp/onepass.go:55
	_go_fuzz_dep_.CoverTab[64797]++

	// Have prefix; gather characters.
	var buf strings.Builder
	for iop(i) == syntax.InstRune && func() bool {
//line /usr/local/go/src/regexp/onepass.go:59
		_go_fuzz_dep_.CoverTab[64807]++
//line /usr/local/go/src/regexp/onepass.go:59
		return len(i.Rune) == 1
//line /usr/local/go/src/regexp/onepass.go:59
		// _ = "end of CoverTab[64807]"
//line /usr/local/go/src/regexp/onepass.go:59
	}() && func() bool {
//line /usr/local/go/src/regexp/onepass.go:59
		_go_fuzz_dep_.CoverTab[64808]++
//line /usr/local/go/src/regexp/onepass.go:59
		return syntax.Flags(i.Arg)&syntax.FoldCase == 0
//line /usr/local/go/src/regexp/onepass.go:59
		// _ = "end of CoverTab[64808]"
//line /usr/local/go/src/regexp/onepass.go:59
	}() && func() bool {
//line /usr/local/go/src/regexp/onepass.go:59
		_go_fuzz_dep_.CoverTab[64809]++
//line /usr/local/go/src/regexp/onepass.go:59
		return i.Rune[0] != utf8.RuneError
//line /usr/local/go/src/regexp/onepass.go:59
		// _ = "end of CoverTab[64809]"
//line /usr/local/go/src/regexp/onepass.go:59
	}() {
//line /usr/local/go/src/regexp/onepass.go:59
		_go_fuzz_dep_.CoverTab[64810]++
							buf.WriteRune(i.Rune[0])
							pc, i = i.Out, &p.Inst[i.Out]
//line /usr/local/go/src/regexp/onepass.go:61
		// _ = "end of CoverTab[64810]"
	}
//line /usr/local/go/src/regexp/onepass.go:62
	// _ = "end of CoverTab[64797]"
//line /usr/local/go/src/regexp/onepass.go:62
	_go_fuzz_dep_.CoverTab[64798]++
						if i.Op == syntax.InstEmptyWidth && func() bool {
//line /usr/local/go/src/regexp/onepass.go:63
		_go_fuzz_dep_.CoverTab[64811]++
//line /usr/local/go/src/regexp/onepass.go:63
		return syntax.EmptyOp(i.Arg)&syntax.EmptyEndText != 0
							// _ = "end of CoverTab[64811]"
//line /usr/local/go/src/regexp/onepass.go:64
	}() && func() bool {
//line /usr/local/go/src/regexp/onepass.go:64
		_go_fuzz_dep_.CoverTab[64812]++
//line /usr/local/go/src/regexp/onepass.go:64
		return p.Inst[i.Out].Op == syntax.InstMatch
							// _ = "end of CoverTab[64812]"
//line /usr/local/go/src/regexp/onepass.go:65
	}() {
//line /usr/local/go/src/regexp/onepass.go:65
		_go_fuzz_dep_.CoverTab[64813]++
							complete = true
//line /usr/local/go/src/regexp/onepass.go:66
		// _ = "end of CoverTab[64813]"
	} else {
//line /usr/local/go/src/regexp/onepass.go:67
		_go_fuzz_dep_.CoverTab[64814]++
//line /usr/local/go/src/regexp/onepass.go:67
		// _ = "end of CoverTab[64814]"
//line /usr/local/go/src/regexp/onepass.go:67
	}
//line /usr/local/go/src/regexp/onepass.go:67
	// _ = "end of CoverTab[64798]"
//line /usr/local/go/src/regexp/onepass.go:67
	_go_fuzz_dep_.CoverTab[64799]++
						return buf.String(), complete, pc
//line /usr/local/go/src/regexp/onepass.go:68
	// _ = "end of CoverTab[64799]"
}

// onePassNext selects the next actionable state of the prog, based on the input character.
//line /usr/local/go/src/regexp/onepass.go:71
// It should only be called when i.Op == InstAlt or InstAltMatch, and from the one-pass machine.
//line /usr/local/go/src/regexp/onepass.go:71
// One of the alternates may ultimately lead without input to end of line. If the instruction
//line /usr/local/go/src/regexp/onepass.go:71
// is InstAltMatch the path to the InstMatch is in i.Out, the normal node in i.Next.
//line /usr/local/go/src/regexp/onepass.go:75
func onePassNext(i *onePassInst, r rune) uint32 {
//line /usr/local/go/src/regexp/onepass.go:75
	_go_fuzz_dep_.CoverTab[64815]++
						next := i.MatchRunePos(r)
						if next >= 0 {
//line /usr/local/go/src/regexp/onepass.go:77
		_go_fuzz_dep_.CoverTab[64818]++
							return i.Next[next]
//line /usr/local/go/src/regexp/onepass.go:78
		// _ = "end of CoverTab[64818]"
	} else {
//line /usr/local/go/src/regexp/onepass.go:79
		_go_fuzz_dep_.CoverTab[64819]++
//line /usr/local/go/src/regexp/onepass.go:79
		// _ = "end of CoverTab[64819]"
//line /usr/local/go/src/regexp/onepass.go:79
	}
//line /usr/local/go/src/regexp/onepass.go:79
	// _ = "end of CoverTab[64815]"
//line /usr/local/go/src/regexp/onepass.go:79
	_go_fuzz_dep_.CoverTab[64816]++
						if i.Op == syntax.InstAltMatch {
//line /usr/local/go/src/regexp/onepass.go:80
		_go_fuzz_dep_.CoverTab[64820]++
							return i.Out
//line /usr/local/go/src/regexp/onepass.go:81
		// _ = "end of CoverTab[64820]"
	} else {
//line /usr/local/go/src/regexp/onepass.go:82
		_go_fuzz_dep_.CoverTab[64821]++
//line /usr/local/go/src/regexp/onepass.go:82
		// _ = "end of CoverTab[64821]"
//line /usr/local/go/src/regexp/onepass.go:82
	}
//line /usr/local/go/src/regexp/onepass.go:82
	// _ = "end of CoverTab[64816]"
//line /usr/local/go/src/regexp/onepass.go:82
	_go_fuzz_dep_.CoverTab[64817]++
						return 0
//line /usr/local/go/src/regexp/onepass.go:83
	// _ = "end of CoverTab[64817]"
}

func iop(i *syntax.Inst) syntax.InstOp {
//line /usr/local/go/src/regexp/onepass.go:86
	_go_fuzz_dep_.CoverTab[64822]++
						op := i.Op
						switch op {
	case syntax.InstRune1, syntax.InstRuneAny, syntax.InstRuneAnyNotNL:
//line /usr/local/go/src/regexp/onepass.go:89
		_go_fuzz_dep_.CoverTab[64824]++
							op = syntax.InstRune
//line /usr/local/go/src/regexp/onepass.go:90
		// _ = "end of CoverTab[64824]"
//line /usr/local/go/src/regexp/onepass.go:90
	default:
//line /usr/local/go/src/regexp/onepass.go:90
		_go_fuzz_dep_.CoverTab[64825]++
//line /usr/local/go/src/regexp/onepass.go:90
		// _ = "end of CoverTab[64825]"
	}
//line /usr/local/go/src/regexp/onepass.go:91
	// _ = "end of CoverTab[64822]"
//line /usr/local/go/src/regexp/onepass.go:91
	_go_fuzz_dep_.CoverTab[64823]++
						return op
//line /usr/local/go/src/regexp/onepass.go:92
	// _ = "end of CoverTab[64823]"
}

// Sparse Array implementation is used as a queueOnePass.
type queueOnePass struct {
	sparse		[]uint32
	dense		[]uint32
	size, nextIndex	uint32
}

func (q *queueOnePass) empty() bool {
//line /usr/local/go/src/regexp/onepass.go:102
	_go_fuzz_dep_.CoverTab[64826]++
						return q.nextIndex >= q.size
//line /usr/local/go/src/regexp/onepass.go:103
	// _ = "end of CoverTab[64826]"
}

func (q *queueOnePass) next() (n uint32) {
//line /usr/local/go/src/regexp/onepass.go:106
	_go_fuzz_dep_.CoverTab[64827]++
						n = q.dense[q.nextIndex]
						q.nextIndex++
						return
//line /usr/local/go/src/regexp/onepass.go:109
	// _ = "end of CoverTab[64827]"
}

func (q *queueOnePass) clear() {
//line /usr/local/go/src/regexp/onepass.go:112
	_go_fuzz_dep_.CoverTab[64828]++
						q.size = 0
						q.nextIndex = 0
//line /usr/local/go/src/regexp/onepass.go:114
	// _ = "end of CoverTab[64828]"
}

func (q *queueOnePass) contains(u uint32) bool {
//line /usr/local/go/src/regexp/onepass.go:117
	_go_fuzz_dep_.CoverTab[64829]++
						if u >= uint32(len(q.sparse)) {
//line /usr/local/go/src/regexp/onepass.go:118
		_go_fuzz_dep_.CoverTab[64831]++
							return false
//line /usr/local/go/src/regexp/onepass.go:119
		// _ = "end of CoverTab[64831]"
	} else {
//line /usr/local/go/src/regexp/onepass.go:120
		_go_fuzz_dep_.CoverTab[64832]++
//line /usr/local/go/src/regexp/onepass.go:120
		// _ = "end of CoverTab[64832]"
//line /usr/local/go/src/regexp/onepass.go:120
	}
//line /usr/local/go/src/regexp/onepass.go:120
	// _ = "end of CoverTab[64829]"
//line /usr/local/go/src/regexp/onepass.go:120
	_go_fuzz_dep_.CoverTab[64830]++
						return q.sparse[u] < q.size && func() bool {
//line /usr/local/go/src/regexp/onepass.go:121
		_go_fuzz_dep_.CoverTab[64833]++
//line /usr/local/go/src/regexp/onepass.go:121
		return q.dense[q.sparse[u]] == u
//line /usr/local/go/src/regexp/onepass.go:121
		// _ = "end of CoverTab[64833]"
//line /usr/local/go/src/regexp/onepass.go:121
	}()
//line /usr/local/go/src/regexp/onepass.go:121
	// _ = "end of CoverTab[64830]"
}

func (q *queueOnePass) insert(u uint32) {
//line /usr/local/go/src/regexp/onepass.go:124
	_go_fuzz_dep_.CoverTab[64834]++
						if !q.contains(u) {
//line /usr/local/go/src/regexp/onepass.go:125
		_go_fuzz_dep_.CoverTab[64835]++
							q.insertNew(u)
//line /usr/local/go/src/regexp/onepass.go:126
		// _ = "end of CoverTab[64835]"
	} else {
//line /usr/local/go/src/regexp/onepass.go:127
		_go_fuzz_dep_.CoverTab[64836]++
//line /usr/local/go/src/regexp/onepass.go:127
		// _ = "end of CoverTab[64836]"
//line /usr/local/go/src/regexp/onepass.go:127
	}
//line /usr/local/go/src/regexp/onepass.go:127
	// _ = "end of CoverTab[64834]"
}

func (q *queueOnePass) insertNew(u uint32) {
//line /usr/local/go/src/regexp/onepass.go:130
	_go_fuzz_dep_.CoverTab[64837]++
						if u >= uint32(len(q.sparse)) {
//line /usr/local/go/src/regexp/onepass.go:131
		_go_fuzz_dep_.CoverTab[64839]++
							return
//line /usr/local/go/src/regexp/onepass.go:132
		// _ = "end of CoverTab[64839]"
	} else {
//line /usr/local/go/src/regexp/onepass.go:133
		_go_fuzz_dep_.CoverTab[64840]++
//line /usr/local/go/src/regexp/onepass.go:133
		// _ = "end of CoverTab[64840]"
//line /usr/local/go/src/regexp/onepass.go:133
	}
//line /usr/local/go/src/regexp/onepass.go:133
	// _ = "end of CoverTab[64837]"
//line /usr/local/go/src/regexp/onepass.go:133
	_go_fuzz_dep_.CoverTab[64838]++
						q.sparse[u] = q.size
						q.dense[q.size] = u
						q.size++
//line /usr/local/go/src/regexp/onepass.go:136
	// _ = "end of CoverTab[64838]"
}

func newQueue(size int) (q *queueOnePass) {
//line /usr/local/go/src/regexp/onepass.go:139
	_go_fuzz_dep_.CoverTab[64841]++
						return &queueOnePass{
		sparse:	make([]uint32, size),
		dense:	make([]uint32, size),
	}
//line /usr/local/go/src/regexp/onepass.go:143
	// _ = "end of CoverTab[64841]"
}

// mergeRuneSets merges two non-intersecting runesets, and returns the merged result,
//line /usr/local/go/src/regexp/onepass.go:146
// and a NextIp array. The idea is that if a rune matches the OnePassRunes at index
//line /usr/local/go/src/regexp/onepass.go:146
// i, NextIp[i/2] is the target. If the input sets intersect, an empty runeset and a
//line /usr/local/go/src/regexp/onepass.go:146
// NextIp array with the single element mergeFailed is returned.
//line /usr/local/go/src/regexp/onepass.go:146
// The code assumes that both inputs contain ordered and non-intersecting rune pairs.
//line /usr/local/go/src/regexp/onepass.go:151
const mergeFailed = uint32(0xffffffff)

var (
	noRune	= []rune{}
	noNext	= []uint32{mergeFailed}
)

func mergeRuneSets(leftRunes, rightRunes *[]rune, leftPC, rightPC uint32) ([]rune, []uint32) {
//line /usr/local/go/src/regexp/onepass.go:158
	_go_fuzz_dep_.CoverTab[64842]++
						leftLen := len(*leftRunes)
						rightLen := len(*rightRunes)
						if leftLen&0x1 != 0 || func() bool {
//line /usr/local/go/src/regexp/onepass.go:161
		_go_fuzz_dep_.CoverTab[64847]++
//line /usr/local/go/src/regexp/onepass.go:161
		return rightLen&0x1 != 0
//line /usr/local/go/src/regexp/onepass.go:161
		// _ = "end of CoverTab[64847]"
//line /usr/local/go/src/regexp/onepass.go:161
	}() {
//line /usr/local/go/src/regexp/onepass.go:161
		_go_fuzz_dep_.CoverTab[64848]++
							panic("mergeRuneSets odd length []rune")
//line /usr/local/go/src/regexp/onepass.go:162
		// _ = "end of CoverTab[64848]"
	} else {
//line /usr/local/go/src/regexp/onepass.go:163
		_go_fuzz_dep_.CoverTab[64849]++
//line /usr/local/go/src/regexp/onepass.go:163
		// _ = "end of CoverTab[64849]"
//line /usr/local/go/src/regexp/onepass.go:163
	}
//line /usr/local/go/src/regexp/onepass.go:163
	// _ = "end of CoverTab[64842]"
//line /usr/local/go/src/regexp/onepass.go:163
	_go_fuzz_dep_.CoverTab[64843]++
						var (
		lx, rx int
	)
	merged := make([]rune, 0)
	next := make([]uint32, 0)
	ok := true
	defer func() {
//line /usr/local/go/src/regexp/onepass.go:170
		_go_fuzz_dep_.CoverTab[64850]++
							if !ok {
//line /usr/local/go/src/regexp/onepass.go:171
			_go_fuzz_dep_.CoverTab[64851]++
								merged = nil
								next = nil
//line /usr/local/go/src/regexp/onepass.go:173
			// _ = "end of CoverTab[64851]"
		} else {
//line /usr/local/go/src/regexp/onepass.go:174
			_go_fuzz_dep_.CoverTab[64852]++
//line /usr/local/go/src/regexp/onepass.go:174
			// _ = "end of CoverTab[64852]"
//line /usr/local/go/src/regexp/onepass.go:174
		}
//line /usr/local/go/src/regexp/onepass.go:174
		// _ = "end of CoverTab[64850]"
	}()
//line /usr/local/go/src/regexp/onepass.go:175
	// _ = "end of CoverTab[64843]"
//line /usr/local/go/src/regexp/onepass.go:175
	_go_fuzz_dep_.CoverTab[64844]++

						ix := -1
						extend := func(newLow *int, newArray *[]rune, pc uint32) bool {
//line /usr/local/go/src/regexp/onepass.go:178
		_go_fuzz_dep_.CoverTab[64853]++
							if ix > 0 && func() bool {
//line /usr/local/go/src/regexp/onepass.go:179
			_go_fuzz_dep_.CoverTab[64855]++
//line /usr/local/go/src/regexp/onepass.go:179
			return (*newArray)[*newLow] <= merged[ix]
//line /usr/local/go/src/regexp/onepass.go:179
			// _ = "end of CoverTab[64855]"
//line /usr/local/go/src/regexp/onepass.go:179
		}() {
//line /usr/local/go/src/regexp/onepass.go:179
			_go_fuzz_dep_.CoverTab[64856]++
								return false
//line /usr/local/go/src/regexp/onepass.go:180
			// _ = "end of CoverTab[64856]"
		} else {
//line /usr/local/go/src/regexp/onepass.go:181
			_go_fuzz_dep_.CoverTab[64857]++
//line /usr/local/go/src/regexp/onepass.go:181
			// _ = "end of CoverTab[64857]"
//line /usr/local/go/src/regexp/onepass.go:181
		}
//line /usr/local/go/src/regexp/onepass.go:181
		// _ = "end of CoverTab[64853]"
//line /usr/local/go/src/regexp/onepass.go:181
		_go_fuzz_dep_.CoverTab[64854]++
							merged = append(merged, (*newArray)[*newLow], (*newArray)[*newLow+1])
							*newLow += 2
							ix += 2
							next = append(next, pc)
							return true
//line /usr/local/go/src/regexp/onepass.go:186
		// _ = "end of CoverTab[64854]"
	}
//line /usr/local/go/src/regexp/onepass.go:187
	// _ = "end of CoverTab[64844]"
//line /usr/local/go/src/regexp/onepass.go:187
	_go_fuzz_dep_.CoverTab[64845]++

						for lx < leftLen || func() bool {
//line /usr/local/go/src/regexp/onepass.go:189
		_go_fuzz_dep_.CoverTab[64858]++
//line /usr/local/go/src/regexp/onepass.go:189
		return rx < rightLen
//line /usr/local/go/src/regexp/onepass.go:189
		// _ = "end of CoverTab[64858]"
//line /usr/local/go/src/regexp/onepass.go:189
	}() {
//line /usr/local/go/src/regexp/onepass.go:189
		_go_fuzz_dep_.CoverTab[64859]++
							switch {
		case rx >= rightLen:
//line /usr/local/go/src/regexp/onepass.go:191
			_go_fuzz_dep_.CoverTab[64861]++
								ok = extend(&lx, leftRunes, leftPC)
//line /usr/local/go/src/regexp/onepass.go:192
			// _ = "end of CoverTab[64861]"
		case lx >= leftLen:
//line /usr/local/go/src/regexp/onepass.go:193
			_go_fuzz_dep_.CoverTab[64862]++
								ok = extend(&rx, rightRunes, rightPC)
//line /usr/local/go/src/regexp/onepass.go:194
			// _ = "end of CoverTab[64862]"
		case (*rightRunes)[rx] < (*leftRunes)[lx]:
//line /usr/local/go/src/regexp/onepass.go:195
			_go_fuzz_dep_.CoverTab[64863]++
								ok = extend(&rx, rightRunes, rightPC)
//line /usr/local/go/src/regexp/onepass.go:196
			// _ = "end of CoverTab[64863]"
		default:
//line /usr/local/go/src/regexp/onepass.go:197
			_go_fuzz_dep_.CoverTab[64864]++
								ok = extend(&lx, leftRunes, leftPC)
//line /usr/local/go/src/regexp/onepass.go:198
			// _ = "end of CoverTab[64864]"
		}
//line /usr/local/go/src/regexp/onepass.go:199
		// _ = "end of CoverTab[64859]"
//line /usr/local/go/src/regexp/onepass.go:199
		_go_fuzz_dep_.CoverTab[64860]++
							if !ok {
//line /usr/local/go/src/regexp/onepass.go:200
			_go_fuzz_dep_.CoverTab[64865]++
								return noRune, noNext
//line /usr/local/go/src/regexp/onepass.go:201
			// _ = "end of CoverTab[64865]"
		} else {
//line /usr/local/go/src/regexp/onepass.go:202
			_go_fuzz_dep_.CoverTab[64866]++
//line /usr/local/go/src/regexp/onepass.go:202
			// _ = "end of CoverTab[64866]"
//line /usr/local/go/src/regexp/onepass.go:202
		}
//line /usr/local/go/src/regexp/onepass.go:202
		// _ = "end of CoverTab[64860]"
	}
//line /usr/local/go/src/regexp/onepass.go:203
	// _ = "end of CoverTab[64845]"
//line /usr/local/go/src/regexp/onepass.go:203
	_go_fuzz_dep_.CoverTab[64846]++
						return merged, next
//line /usr/local/go/src/regexp/onepass.go:204
	// _ = "end of CoverTab[64846]"
}

// cleanupOnePass drops working memory, and restores certain shortcut instructions.
func cleanupOnePass(prog *onePassProg, original *syntax.Prog) {
//line /usr/local/go/src/regexp/onepass.go:208
	_go_fuzz_dep_.CoverTab[64867]++
						for ix, instOriginal := range original.Inst {
//line /usr/local/go/src/regexp/onepass.go:209
		_go_fuzz_dep_.CoverTab[64868]++
							switch instOriginal.Op {
		case syntax.InstAlt, syntax.InstAltMatch, syntax.InstRune:
//line /usr/local/go/src/regexp/onepass.go:211
			_go_fuzz_dep_.CoverTab[64869]++
//line /usr/local/go/src/regexp/onepass.go:211
			// _ = "end of CoverTab[64869]"
		case syntax.InstCapture, syntax.InstEmptyWidth, syntax.InstNop, syntax.InstMatch, syntax.InstFail:
//line /usr/local/go/src/regexp/onepass.go:212
			_go_fuzz_dep_.CoverTab[64870]++
								prog.Inst[ix].Next = nil
//line /usr/local/go/src/regexp/onepass.go:213
			// _ = "end of CoverTab[64870]"
		case syntax.InstRune1, syntax.InstRuneAny, syntax.InstRuneAnyNotNL:
//line /usr/local/go/src/regexp/onepass.go:214
			_go_fuzz_dep_.CoverTab[64871]++
								prog.Inst[ix].Next = nil
								prog.Inst[ix] = onePassInst{Inst: instOriginal}
//line /usr/local/go/src/regexp/onepass.go:216
			// _ = "end of CoverTab[64871]"
//line /usr/local/go/src/regexp/onepass.go:216
		default:
//line /usr/local/go/src/regexp/onepass.go:216
			_go_fuzz_dep_.CoverTab[64872]++
//line /usr/local/go/src/regexp/onepass.go:216
			// _ = "end of CoverTab[64872]"
		}
//line /usr/local/go/src/regexp/onepass.go:217
		// _ = "end of CoverTab[64868]"
	}
//line /usr/local/go/src/regexp/onepass.go:218
	// _ = "end of CoverTab[64867]"
}

// onePassCopy creates a copy of the original Prog, as we'll be modifying it.
func onePassCopy(prog *syntax.Prog) *onePassProg {
//line /usr/local/go/src/regexp/onepass.go:222
	_go_fuzz_dep_.CoverTab[64873]++
						p := &onePassProg{
		Start:	prog.Start,
		NumCap:	prog.NumCap,
		Inst:	make([]onePassInst, len(prog.Inst)),
	}
	for i, inst := range prog.Inst {
//line /usr/local/go/src/regexp/onepass.go:228
		_go_fuzz_dep_.CoverTab[64876]++
							p.Inst[i] = onePassInst{Inst: inst}
//line /usr/local/go/src/regexp/onepass.go:229
		// _ = "end of CoverTab[64876]"
	}
//line /usr/local/go/src/regexp/onepass.go:230
	// _ = "end of CoverTab[64873]"
//line /usr/local/go/src/regexp/onepass.go:230
	_go_fuzz_dep_.CoverTab[64874]++

//line /usr/local/go/src/regexp/onepass.go:237
	for pc := range p.Inst {
//line /usr/local/go/src/regexp/onepass.go:237
		_go_fuzz_dep_.CoverTab[64877]++
							switch p.Inst[pc].Op {
		default:
//line /usr/local/go/src/regexp/onepass.go:239
			_go_fuzz_dep_.CoverTab[64878]++
								continue
//line /usr/local/go/src/regexp/onepass.go:240
			// _ = "end of CoverTab[64878]"
		case syntax.InstAlt, syntax.InstAltMatch:
//line /usr/local/go/src/regexp/onepass.go:241
			_go_fuzz_dep_.CoverTab[64879]++

								p_A_Other := &p.Inst[pc].Out
								p_A_Alt := &p.Inst[pc].Arg

								instAlt := p.Inst[*p_A_Alt]
								if !(instAlt.Op == syntax.InstAlt || func() bool {
//line /usr/local/go/src/regexp/onepass.go:247
				_go_fuzz_dep_.CoverTab[64884]++
//line /usr/local/go/src/regexp/onepass.go:247
				return instAlt.Op == syntax.InstAltMatch
//line /usr/local/go/src/regexp/onepass.go:247
				// _ = "end of CoverTab[64884]"
//line /usr/local/go/src/regexp/onepass.go:247
			}()) {
//line /usr/local/go/src/regexp/onepass.go:247
				_go_fuzz_dep_.CoverTab[64885]++
									p_A_Alt, p_A_Other = p_A_Other, p_A_Alt
									instAlt = p.Inst[*p_A_Alt]
									if !(instAlt.Op == syntax.InstAlt || func() bool {
//line /usr/local/go/src/regexp/onepass.go:250
					_go_fuzz_dep_.CoverTab[64886]++
//line /usr/local/go/src/regexp/onepass.go:250
					return instAlt.Op == syntax.InstAltMatch
//line /usr/local/go/src/regexp/onepass.go:250
					// _ = "end of CoverTab[64886]"
//line /usr/local/go/src/regexp/onepass.go:250
				}()) {
//line /usr/local/go/src/regexp/onepass.go:250
					_go_fuzz_dep_.CoverTab[64887]++
										continue
//line /usr/local/go/src/regexp/onepass.go:251
					// _ = "end of CoverTab[64887]"
				} else {
//line /usr/local/go/src/regexp/onepass.go:252
					_go_fuzz_dep_.CoverTab[64888]++
//line /usr/local/go/src/regexp/onepass.go:252
					// _ = "end of CoverTab[64888]"
//line /usr/local/go/src/regexp/onepass.go:252
				}
//line /usr/local/go/src/regexp/onepass.go:252
				// _ = "end of CoverTab[64885]"
			} else {
//line /usr/local/go/src/regexp/onepass.go:253
				_go_fuzz_dep_.CoverTab[64889]++
//line /usr/local/go/src/regexp/onepass.go:253
				// _ = "end of CoverTab[64889]"
//line /usr/local/go/src/regexp/onepass.go:253
			}
//line /usr/local/go/src/regexp/onepass.go:253
			// _ = "end of CoverTab[64879]"
//line /usr/local/go/src/regexp/onepass.go:253
			_go_fuzz_dep_.CoverTab[64880]++
								instOther := p.Inst[*p_A_Other]

								if instOther.Op == syntax.InstAlt || func() bool {
//line /usr/local/go/src/regexp/onepass.go:256
				_go_fuzz_dep_.CoverTab[64890]++
//line /usr/local/go/src/regexp/onepass.go:256
				return instOther.Op == syntax.InstAltMatch
//line /usr/local/go/src/regexp/onepass.go:256
				// _ = "end of CoverTab[64890]"
//line /usr/local/go/src/regexp/onepass.go:256
			}() {
//line /usr/local/go/src/regexp/onepass.go:256
				_go_fuzz_dep_.CoverTab[64891]++

									continue
//line /usr/local/go/src/regexp/onepass.go:258
				// _ = "end of CoverTab[64891]"
			} else {
//line /usr/local/go/src/regexp/onepass.go:259
				_go_fuzz_dep_.CoverTab[64892]++
//line /usr/local/go/src/regexp/onepass.go:259
				// _ = "end of CoverTab[64892]"
//line /usr/local/go/src/regexp/onepass.go:259
			}
//line /usr/local/go/src/regexp/onepass.go:259
			// _ = "end of CoverTab[64880]"
//line /usr/local/go/src/regexp/onepass.go:259
			_go_fuzz_dep_.CoverTab[64881]++

//line /usr/local/go/src/regexp/onepass.go:262
			p_B_Alt := &p.Inst[*p_A_Alt].Out
			p_B_Other := &p.Inst[*p_A_Alt].Arg
			patch := false
			if instAlt.Out == uint32(pc) {
//line /usr/local/go/src/regexp/onepass.go:265
				_go_fuzz_dep_.CoverTab[64893]++
									patch = true
//line /usr/local/go/src/regexp/onepass.go:266
				// _ = "end of CoverTab[64893]"
			} else {
//line /usr/local/go/src/regexp/onepass.go:267
				_go_fuzz_dep_.CoverTab[64894]++
//line /usr/local/go/src/regexp/onepass.go:267
				if instAlt.Arg == uint32(pc) {
//line /usr/local/go/src/regexp/onepass.go:267
					_go_fuzz_dep_.CoverTab[64895]++
										patch = true
										p_B_Alt, p_B_Other = p_B_Other, p_B_Alt
//line /usr/local/go/src/regexp/onepass.go:269
					// _ = "end of CoverTab[64895]"
				} else {
//line /usr/local/go/src/regexp/onepass.go:270
					_go_fuzz_dep_.CoverTab[64896]++
//line /usr/local/go/src/regexp/onepass.go:270
					// _ = "end of CoverTab[64896]"
//line /usr/local/go/src/regexp/onepass.go:270
				}
//line /usr/local/go/src/regexp/onepass.go:270
				// _ = "end of CoverTab[64894]"
//line /usr/local/go/src/regexp/onepass.go:270
			}
//line /usr/local/go/src/regexp/onepass.go:270
			// _ = "end of CoverTab[64881]"
//line /usr/local/go/src/regexp/onepass.go:270
			_go_fuzz_dep_.CoverTab[64882]++
								if patch {
//line /usr/local/go/src/regexp/onepass.go:271
				_go_fuzz_dep_.CoverTab[64897]++
									*p_B_Alt = *p_A_Other
//line /usr/local/go/src/regexp/onepass.go:272
				// _ = "end of CoverTab[64897]"
			} else {
//line /usr/local/go/src/regexp/onepass.go:273
				_go_fuzz_dep_.CoverTab[64898]++
//line /usr/local/go/src/regexp/onepass.go:273
				// _ = "end of CoverTab[64898]"
//line /usr/local/go/src/regexp/onepass.go:273
			}
//line /usr/local/go/src/regexp/onepass.go:273
			// _ = "end of CoverTab[64882]"
//line /usr/local/go/src/regexp/onepass.go:273
			_go_fuzz_dep_.CoverTab[64883]++

//line /usr/local/go/src/regexp/onepass.go:277
			if *p_A_Other == *p_B_Alt {
//line /usr/local/go/src/regexp/onepass.go:277
				_go_fuzz_dep_.CoverTab[64899]++
									*p_A_Alt = *p_B_Other
//line /usr/local/go/src/regexp/onepass.go:278
				// _ = "end of CoverTab[64899]"
			} else {
//line /usr/local/go/src/regexp/onepass.go:279
				_go_fuzz_dep_.CoverTab[64900]++
//line /usr/local/go/src/regexp/onepass.go:279
				// _ = "end of CoverTab[64900]"
//line /usr/local/go/src/regexp/onepass.go:279
			}
//line /usr/local/go/src/regexp/onepass.go:279
			// _ = "end of CoverTab[64883]"
		}
//line /usr/local/go/src/regexp/onepass.go:280
		// _ = "end of CoverTab[64877]"
	}
//line /usr/local/go/src/regexp/onepass.go:281
	// _ = "end of CoverTab[64874]"
//line /usr/local/go/src/regexp/onepass.go:281
	_go_fuzz_dep_.CoverTab[64875]++
						return p
//line /usr/local/go/src/regexp/onepass.go:282
	// _ = "end of CoverTab[64875]"
}

// runeSlice exists to permit sorting the case-folded rune sets.
type runeSlice []rune

func (p runeSlice) Len() int {
//line /usr/local/go/src/regexp/onepass.go:288
	_go_fuzz_dep_.CoverTab[64901]++
//line /usr/local/go/src/regexp/onepass.go:288
	return len(p)
//line /usr/local/go/src/regexp/onepass.go:288
	// _ = "end of CoverTab[64901]"
//line /usr/local/go/src/regexp/onepass.go:288
}
func (p runeSlice) Less(i, j int) bool {
//line /usr/local/go/src/regexp/onepass.go:289
	_go_fuzz_dep_.CoverTab[64902]++
//line /usr/local/go/src/regexp/onepass.go:289
	return p[i] < p[j]
//line /usr/local/go/src/regexp/onepass.go:289
	// _ = "end of CoverTab[64902]"
//line /usr/local/go/src/regexp/onepass.go:289
}
func (p runeSlice) Swap(i, j int) {
//line /usr/local/go/src/regexp/onepass.go:290
	_go_fuzz_dep_.CoverTab[64903]++
//line /usr/local/go/src/regexp/onepass.go:290
	p[i], p[j] = p[j], p[i]
//line /usr/local/go/src/regexp/onepass.go:290
	// _ = "end of CoverTab[64903]"
//line /usr/local/go/src/regexp/onepass.go:290
}

var anyRuneNotNL = []rune{0, '\n' - 1, '\n' + 1, unicode.MaxRune}
var anyRune = []rune{0, unicode.MaxRune}

// makeOnePass creates a onepass Prog, if possible. It is possible if at any alt,
//line /usr/local/go/src/regexp/onepass.go:295
// the match engine can always tell which branch to take. The routine may modify
//line /usr/local/go/src/regexp/onepass.go:295
// p if it is turned into a onepass Prog. If it isn't possible for this to be a
//line /usr/local/go/src/regexp/onepass.go:295
// onepass Prog, the Prog nil is returned. makeOnePass is recursive
//line /usr/local/go/src/regexp/onepass.go:295
// to the size of the Prog.
//line /usr/local/go/src/regexp/onepass.go:300
func makeOnePass(p *onePassProg) *onePassProg {
//line /usr/local/go/src/regexp/onepass.go:300
	_go_fuzz_dep_.CoverTab[64904]++

						if len(p.Inst) >= 1000 {
//line /usr/local/go/src/regexp/onepass.go:302
		_go_fuzz_dep_.CoverTab[64909]++
							return nil
//line /usr/local/go/src/regexp/onepass.go:303
		// _ = "end of CoverTab[64909]"
	} else {
//line /usr/local/go/src/regexp/onepass.go:304
		_go_fuzz_dep_.CoverTab[64910]++
//line /usr/local/go/src/regexp/onepass.go:304
		// _ = "end of CoverTab[64910]"
//line /usr/local/go/src/regexp/onepass.go:304
	}
//line /usr/local/go/src/regexp/onepass.go:304
	// _ = "end of CoverTab[64904]"
//line /usr/local/go/src/regexp/onepass.go:304
	_go_fuzz_dep_.CoverTab[64905]++

						var (
		instQueue	= newQueue(len(p.Inst))
		visitQueue	= newQueue(len(p.Inst))
		check		func(uint32, []bool) bool
		onePassRunes	= make([][]rune, len(p.Inst))
	)

//line /usr/local/go/src/regexp/onepass.go:315
	check = func(pc uint32, m []bool) (ok bool) {
//line /usr/local/go/src/regexp/onepass.go:315
		_go_fuzz_dep_.CoverTab[64911]++
							ok = true
							inst := &p.Inst[pc]
							if visitQueue.contains(pc) {
//line /usr/local/go/src/regexp/onepass.go:318
			_go_fuzz_dep_.CoverTab[64914]++
								return
//line /usr/local/go/src/regexp/onepass.go:319
			// _ = "end of CoverTab[64914]"
		} else {
//line /usr/local/go/src/regexp/onepass.go:320
			_go_fuzz_dep_.CoverTab[64915]++
//line /usr/local/go/src/regexp/onepass.go:320
			// _ = "end of CoverTab[64915]"
//line /usr/local/go/src/regexp/onepass.go:320
		}
//line /usr/local/go/src/regexp/onepass.go:320
		// _ = "end of CoverTab[64911]"
//line /usr/local/go/src/regexp/onepass.go:320
		_go_fuzz_dep_.CoverTab[64912]++
							visitQueue.insert(pc)
							switch inst.Op {
		case syntax.InstAlt, syntax.InstAltMatch:
//line /usr/local/go/src/regexp/onepass.go:323
			_go_fuzz_dep_.CoverTab[64916]++
								ok = check(inst.Out, m) && func() bool {
//line /usr/local/go/src/regexp/onepass.go:324
				_go_fuzz_dep_.CoverTab[64937]++
//line /usr/local/go/src/regexp/onepass.go:324
				return check(inst.Arg, m)
//line /usr/local/go/src/regexp/onepass.go:324
				// _ = "end of CoverTab[64937]"
//line /usr/local/go/src/regexp/onepass.go:324
			}()

								matchOut := m[inst.Out]
								matchArg := m[inst.Arg]
								if matchOut && func() bool {
//line /usr/local/go/src/regexp/onepass.go:328
				_go_fuzz_dep_.CoverTab[64938]++
//line /usr/local/go/src/regexp/onepass.go:328
				return matchArg
//line /usr/local/go/src/regexp/onepass.go:328
				// _ = "end of CoverTab[64938]"
//line /usr/local/go/src/regexp/onepass.go:328
			}() {
//line /usr/local/go/src/regexp/onepass.go:328
				_go_fuzz_dep_.CoverTab[64939]++
									ok = false
									break
//line /usr/local/go/src/regexp/onepass.go:330
				// _ = "end of CoverTab[64939]"
			} else {
//line /usr/local/go/src/regexp/onepass.go:331
				_go_fuzz_dep_.CoverTab[64940]++
//line /usr/local/go/src/regexp/onepass.go:331
				// _ = "end of CoverTab[64940]"
//line /usr/local/go/src/regexp/onepass.go:331
			}
//line /usr/local/go/src/regexp/onepass.go:331
			// _ = "end of CoverTab[64916]"
//line /usr/local/go/src/regexp/onepass.go:331
			_go_fuzz_dep_.CoverTab[64917]++

								if matchArg {
//line /usr/local/go/src/regexp/onepass.go:333
				_go_fuzz_dep_.CoverTab[64941]++
									inst.Out, inst.Arg = inst.Arg, inst.Out
									matchOut, matchArg = matchArg, matchOut
//line /usr/local/go/src/regexp/onepass.go:335
				// _ = "end of CoverTab[64941]"
			} else {
//line /usr/local/go/src/regexp/onepass.go:336
				_go_fuzz_dep_.CoverTab[64942]++
//line /usr/local/go/src/regexp/onepass.go:336
				// _ = "end of CoverTab[64942]"
//line /usr/local/go/src/regexp/onepass.go:336
			}
//line /usr/local/go/src/regexp/onepass.go:336
			// _ = "end of CoverTab[64917]"
//line /usr/local/go/src/regexp/onepass.go:336
			_go_fuzz_dep_.CoverTab[64918]++
								if matchOut {
//line /usr/local/go/src/regexp/onepass.go:337
				_go_fuzz_dep_.CoverTab[64943]++
									m[pc] = true
									inst.Op = syntax.InstAltMatch
//line /usr/local/go/src/regexp/onepass.go:339
				// _ = "end of CoverTab[64943]"
			} else {
//line /usr/local/go/src/regexp/onepass.go:340
				_go_fuzz_dep_.CoverTab[64944]++
//line /usr/local/go/src/regexp/onepass.go:340
				// _ = "end of CoverTab[64944]"
//line /usr/local/go/src/regexp/onepass.go:340
			}
//line /usr/local/go/src/regexp/onepass.go:340
			// _ = "end of CoverTab[64918]"
//line /usr/local/go/src/regexp/onepass.go:340
			_go_fuzz_dep_.CoverTab[64919]++

//line /usr/local/go/src/regexp/onepass.go:343
			onePassRunes[pc], inst.Next = mergeRuneSets(
				&onePassRunes[inst.Out], &onePassRunes[inst.Arg], inst.Out, inst.Arg)
			if len(inst.Next) > 0 && func() bool {
//line /usr/local/go/src/regexp/onepass.go:345
				_go_fuzz_dep_.CoverTab[64945]++
//line /usr/local/go/src/regexp/onepass.go:345
				return inst.Next[0] == mergeFailed
//line /usr/local/go/src/regexp/onepass.go:345
				// _ = "end of CoverTab[64945]"
//line /usr/local/go/src/regexp/onepass.go:345
			}() {
//line /usr/local/go/src/regexp/onepass.go:345
				_go_fuzz_dep_.CoverTab[64946]++
									ok = false
									break
//line /usr/local/go/src/regexp/onepass.go:347
				// _ = "end of CoverTab[64946]"
			} else {
//line /usr/local/go/src/regexp/onepass.go:348
				_go_fuzz_dep_.CoverTab[64947]++
//line /usr/local/go/src/regexp/onepass.go:348
				// _ = "end of CoverTab[64947]"
//line /usr/local/go/src/regexp/onepass.go:348
			}
//line /usr/local/go/src/regexp/onepass.go:348
			// _ = "end of CoverTab[64919]"
		case syntax.InstCapture, syntax.InstNop:
//line /usr/local/go/src/regexp/onepass.go:349
			_go_fuzz_dep_.CoverTab[64920]++
								ok = check(inst.Out, m)
								m[pc] = m[inst.Out]

								onePassRunes[pc] = append([]rune{}, onePassRunes[inst.Out]...)
								inst.Next = make([]uint32, len(onePassRunes[pc])/2+1)
								for i := range inst.Next {
//line /usr/local/go/src/regexp/onepass.go:355
				_go_fuzz_dep_.CoverTab[64948]++
									inst.Next[i] = inst.Out
//line /usr/local/go/src/regexp/onepass.go:356
				// _ = "end of CoverTab[64948]"
			}
//line /usr/local/go/src/regexp/onepass.go:357
			// _ = "end of CoverTab[64920]"
		case syntax.InstEmptyWidth:
//line /usr/local/go/src/regexp/onepass.go:358
			_go_fuzz_dep_.CoverTab[64921]++
								ok = check(inst.Out, m)
								m[pc] = m[inst.Out]
								onePassRunes[pc] = append([]rune{}, onePassRunes[inst.Out]...)
								inst.Next = make([]uint32, len(onePassRunes[pc])/2+1)
								for i := range inst.Next {
//line /usr/local/go/src/regexp/onepass.go:363
				_go_fuzz_dep_.CoverTab[64949]++
									inst.Next[i] = inst.Out
//line /usr/local/go/src/regexp/onepass.go:364
				// _ = "end of CoverTab[64949]"
			}
//line /usr/local/go/src/regexp/onepass.go:365
			// _ = "end of CoverTab[64921]"
		case syntax.InstMatch, syntax.InstFail:
//line /usr/local/go/src/regexp/onepass.go:366
			_go_fuzz_dep_.CoverTab[64922]++
								m[pc] = inst.Op == syntax.InstMatch
//line /usr/local/go/src/regexp/onepass.go:367
			// _ = "end of CoverTab[64922]"
		case syntax.InstRune:
//line /usr/local/go/src/regexp/onepass.go:368
			_go_fuzz_dep_.CoverTab[64923]++
								m[pc] = false
								if len(inst.Next) > 0 {
//line /usr/local/go/src/regexp/onepass.go:370
				_go_fuzz_dep_.CoverTab[64950]++
									break
//line /usr/local/go/src/regexp/onepass.go:371
				// _ = "end of CoverTab[64950]"
			} else {
//line /usr/local/go/src/regexp/onepass.go:372
				_go_fuzz_dep_.CoverTab[64951]++
//line /usr/local/go/src/regexp/onepass.go:372
				// _ = "end of CoverTab[64951]"
//line /usr/local/go/src/regexp/onepass.go:372
			}
//line /usr/local/go/src/regexp/onepass.go:372
			// _ = "end of CoverTab[64923]"
//line /usr/local/go/src/regexp/onepass.go:372
			_go_fuzz_dep_.CoverTab[64924]++
								instQueue.insert(inst.Out)
								if len(inst.Rune) == 0 {
//line /usr/local/go/src/regexp/onepass.go:374
				_go_fuzz_dep_.CoverTab[64952]++
									onePassRunes[pc] = []rune{}
									inst.Next = []uint32{inst.Out}
									break
//line /usr/local/go/src/regexp/onepass.go:377
				// _ = "end of CoverTab[64952]"
			} else {
//line /usr/local/go/src/regexp/onepass.go:378
				_go_fuzz_dep_.CoverTab[64953]++
//line /usr/local/go/src/regexp/onepass.go:378
				// _ = "end of CoverTab[64953]"
//line /usr/local/go/src/regexp/onepass.go:378
			}
//line /usr/local/go/src/regexp/onepass.go:378
			// _ = "end of CoverTab[64924]"
//line /usr/local/go/src/regexp/onepass.go:378
			_go_fuzz_dep_.CoverTab[64925]++
								runes := make([]rune, 0)
								if len(inst.Rune) == 1 && func() bool {
//line /usr/local/go/src/regexp/onepass.go:380
				_go_fuzz_dep_.CoverTab[64954]++
//line /usr/local/go/src/regexp/onepass.go:380
				return syntax.Flags(inst.Arg)&syntax.FoldCase != 0
//line /usr/local/go/src/regexp/onepass.go:380
				// _ = "end of CoverTab[64954]"
//line /usr/local/go/src/regexp/onepass.go:380
			}() {
//line /usr/local/go/src/regexp/onepass.go:380
				_go_fuzz_dep_.CoverTab[64955]++
									r0 := inst.Rune[0]
									runes = append(runes, r0, r0)
									for r1 := unicode.SimpleFold(r0); r1 != r0; r1 = unicode.SimpleFold(r1) {
//line /usr/local/go/src/regexp/onepass.go:383
					_go_fuzz_dep_.CoverTab[64957]++
										runes = append(runes, r1, r1)
//line /usr/local/go/src/regexp/onepass.go:384
					// _ = "end of CoverTab[64957]"
				}
//line /usr/local/go/src/regexp/onepass.go:385
				// _ = "end of CoverTab[64955]"
//line /usr/local/go/src/regexp/onepass.go:385
				_go_fuzz_dep_.CoverTab[64956]++
									sort.Sort(runeSlice(runes))
//line /usr/local/go/src/regexp/onepass.go:386
				// _ = "end of CoverTab[64956]"
			} else {
//line /usr/local/go/src/regexp/onepass.go:387
				_go_fuzz_dep_.CoverTab[64958]++
									runes = append(runes, inst.Rune...)
//line /usr/local/go/src/regexp/onepass.go:388
				// _ = "end of CoverTab[64958]"
			}
//line /usr/local/go/src/regexp/onepass.go:389
			// _ = "end of CoverTab[64925]"
//line /usr/local/go/src/regexp/onepass.go:389
			_go_fuzz_dep_.CoverTab[64926]++
								onePassRunes[pc] = runes
								inst.Next = make([]uint32, len(onePassRunes[pc])/2+1)
								for i := range inst.Next {
//line /usr/local/go/src/regexp/onepass.go:392
				_go_fuzz_dep_.CoverTab[64959]++
									inst.Next[i] = inst.Out
//line /usr/local/go/src/regexp/onepass.go:393
				// _ = "end of CoverTab[64959]"
			}
//line /usr/local/go/src/regexp/onepass.go:394
			// _ = "end of CoverTab[64926]"
//line /usr/local/go/src/regexp/onepass.go:394
			_go_fuzz_dep_.CoverTab[64927]++
								inst.Op = syntax.InstRune
//line /usr/local/go/src/regexp/onepass.go:395
			// _ = "end of CoverTab[64927]"
		case syntax.InstRune1:
//line /usr/local/go/src/regexp/onepass.go:396
			_go_fuzz_dep_.CoverTab[64928]++
								m[pc] = false
								if len(inst.Next) > 0 {
//line /usr/local/go/src/regexp/onepass.go:398
				_go_fuzz_dep_.CoverTab[64960]++
									break
//line /usr/local/go/src/regexp/onepass.go:399
				// _ = "end of CoverTab[64960]"
			} else {
//line /usr/local/go/src/regexp/onepass.go:400
				_go_fuzz_dep_.CoverTab[64961]++
//line /usr/local/go/src/regexp/onepass.go:400
				// _ = "end of CoverTab[64961]"
//line /usr/local/go/src/regexp/onepass.go:400
			}
//line /usr/local/go/src/regexp/onepass.go:400
			// _ = "end of CoverTab[64928]"
//line /usr/local/go/src/regexp/onepass.go:400
			_go_fuzz_dep_.CoverTab[64929]++
								instQueue.insert(inst.Out)
								runes := []rune{}

								if syntax.Flags(inst.Arg)&syntax.FoldCase != 0 {
//line /usr/local/go/src/regexp/onepass.go:404
				_go_fuzz_dep_.CoverTab[64962]++
									r0 := inst.Rune[0]
									runes = append(runes, r0, r0)
									for r1 := unicode.SimpleFold(r0); r1 != r0; r1 = unicode.SimpleFold(r1) {
//line /usr/local/go/src/regexp/onepass.go:407
					_go_fuzz_dep_.CoverTab[64964]++
										runes = append(runes, r1, r1)
//line /usr/local/go/src/regexp/onepass.go:408
					// _ = "end of CoverTab[64964]"
				}
//line /usr/local/go/src/regexp/onepass.go:409
				// _ = "end of CoverTab[64962]"
//line /usr/local/go/src/regexp/onepass.go:409
				_go_fuzz_dep_.CoverTab[64963]++
									sort.Sort(runeSlice(runes))
//line /usr/local/go/src/regexp/onepass.go:410
				// _ = "end of CoverTab[64963]"
			} else {
//line /usr/local/go/src/regexp/onepass.go:411
				_go_fuzz_dep_.CoverTab[64965]++
									runes = append(runes, inst.Rune[0], inst.Rune[0])
//line /usr/local/go/src/regexp/onepass.go:412
				// _ = "end of CoverTab[64965]"
			}
//line /usr/local/go/src/regexp/onepass.go:413
			// _ = "end of CoverTab[64929]"
//line /usr/local/go/src/regexp/onepass.go:413
			_go_fuzz_dep_.CoverTab[64930]++
								onePassRunes[pc] = runes
								inst.Next = make([]uint32, len(onePassRunes[pc])/2+1)
								for i := range inst.Next {
//line /usr/local/go/src/regexp/onepass.go:416
				_go_fuzz_dep_.CoverTab[64966]++
									inst.Next[i] = inst.Out
//line /usr/local/go/src/regexp/onepass.go:417
				// _ = "end of CoverTab[64966]"
			}
//line /usr/local/go/src/regexp/onepass.go:418
			// _ = "end of CoverTab[64930]"
//line /usr/local/go/src/regexp/onepass.go:418
			_go_fuzz_dep_.CoverTab[64931]++
								inst.Op = syntax.InstRune
//line /usr/local/go/src/regexp/onepass.go:419
			// _ = "end of CoverTab[64931]"
		case syntax.InstRuneAny:
//line /usr/local/go/src/regexp/onepass.go:420
			_go_fuzz_dep_.CoverTab[64932]++
								m[pc] = false
								if len(inst.Next) > 0 {
//line /usr/local/go/src/regexp/onepass.go:422
				_go_fuzz_dep_.CoverTab[64967]++
									break
//line /usr/local/go/src/regexp/onepass.go:423
				// _ = "end of CoverTab[64967]"
			} else {
//line /usr/local/go/src/regexp/onepass.go:424
				_go_fuzz_dep_.CoverTab[64968]++
//line /usr/local/go/src/regexp/onepass.go:424
				// _ = "end of CoverTab[64968]"
//line /usr/local/go/src/regexp/onepass.go:424
			}
//line /usr/local/go/src/regexp/onepass.go:424
			// _ = "end of CoverTab[64932]"
//line /usr/local/go/src/regexp/onepass.go:424
			_go_fuzz_dep_.CoverTab[64933]++
								instQueue.insert(inst.Out)
								onePassRunes[pc] = append([]rune{}, anyRune...)
								inst.Next = []uint32{inst.Out}
//line /usr/local/go/src/regexp/onepass.go:427
			// _ = "end of CoverTab[64933]"
		case syntax.InstRuneAnyNotNL:
//line /usr/local/go/src/regexp/onepass.go:428
			_go_fuzz_dep_.CoverTab[64934]++
								m[pc] = false
								if len(inst.Next) > 0 {
//line /usr/local/go/src/regexp/onepass.go:430
				_go_fuzz_dep_.CoverTab[64969]++
									break
//line /usr/local/go/src/regexp/onepass.go:431
				// _ = "end of CoverTab[64969]"
			} else {
//line /usr/local/go/src/regexp/onepass.go:432
				_go_fuzz_dep_.CoverTab[64970]++
//line /usr/local/go/src/regexp/onepass.go:432
				// _ = "end of CoverTab[64970]"
//line /usr/local/go/src/regexp/onepass.go:432
			}
//line /usr/local/go/src/regexp/onepass.go:432
			// _ = "end of CoverTab[64934]"
//line /usr/local/go/src/regexp/onepass.go:432
			_go_fuzz_dep_.CoverTab[64935]++
								instQueue.insert(inst.Out)
								onePassRunes[pc] = append([]rune{}, anyRuneNotNL...)
								inst.Next = make([]uint32, len(onePassRunes[pc])/2+1)
								for i := range inst.Next {
//line /usr/local/go/src/regexp/onepass.go:436
				_go_fuzz_dep_.CoverTab[64971]++
									inst.Next[i] = inst.Out
//line /usr/local/go/src/regexp/onepass.go:437
				// _ = "end of CoverTab[64971]"
			}
//line /usr/local/go/src/regexp/onepass.go:438
			// _ = "end of CoverTab[64935]"
//line /usr/local/go/src/regexp/onepass.go:438
		default:
//line /usr/local/go/src/regexp/onepass.go:438
			_go_fuzz_dep_.CoverTab[64936]++
//line /usr/local/go/src/regexp/onepass.go:438
			// _ = "end of CoverTab[64936]"
		}
//line /usr/local/go/src/regexp/onepass.go:439
		// _ = "end of CoverTab[64912]"
//line /usr/local/go/src/regexp/onepass.go:439
		_go_fuzz_dep_.CoverTab[64913]++
							return
//line /usr/local/go/src/regexp/onepass.go:440
		// _ = "end of CoverTab[64913]"
	}
//line /usr/local/go/src/regexp/onepass.go:441
	// _ = "end of CoverTab[64905]"
//line /usr/local/go/src/regexp/onepass.go:441
	_go_fuzz_dep_.CoverTab[64906]++

						instQueue.clear()
						instQueue.insert(uint32(p.Start))
						m := make([]bool, len(p.Inst))
						for !instQueue.empty() {
//line /usr/local/go/src/regexp/onepass.go:446
		_go_fuzz_dep_.CoverTab[64972]++
							visitQueue.clear()
							pc := instQueue.next()
							if !check(pc, m) {
//line /usr/local/go/src/regexp/onepass.go:449
			_go_fuzz_dep_.CoverTab[64973]++
								p = nil
								break
//line /usr/local/go/src/regexp/onepass.go:451
			// _ = "end of CoverTab[64973]"
		} else {
//line /usr/local/go/src/regexp/onepass.go:452
			_go_fuzz_dep_.CoverTab[64974]++
//line /usr/local/go/src/regexp/onepass.go:452
			// _ = "end of CoverTab[64974]"
//line /usr/local/go/src/regexp/onepass.go:452
		}
//line /usr/local/go/src/regexp/onepass.go:452
		// _ = "end of CoverTab[64972]"
	}
//line /usr/local/go/src/regexp/onepass.go:453
	// _ = "end of CoverTab[64906]"
//line /usr/local/go/src/regexp/onepass.go:453
	_go_fuzz_dep_.CoverTab[64907]++
						if p != nil {
//line /usr/local/go/src/regexp/onepass.go:454
		_go_fuzz_dep_.CoverTab[64975]++
							for i := range p.Inst {
//line /usr/local/go/src/regexp/onepass.go:455
			_go_fuzz_dep_.CoverTab[64976]++
								p.Inst[i].Rune = onePassRunes[i]
//line /usr/local/go/src/regexp/onepass.go:456
			// _ = "end of CoverTab[64976]"
		}
//line /usr/local/go/src/regexp/onepass.go:457
		// _ = "end of CoverTab[64975]"
	} else {
//line /usr/local/go/src/regexp/onepass.go:458
		_go_fuzz_dep_.CoverTab[64977]++
//line /usr/local/go/src/regexp/onepass.go:458
		// _ = "end of CoverTab[64977]"
//line /usr/local/go/src/regexp/onepass.go:458
	}
//line /usr/local/go/src/regexp/onepass.go:458
	// _ = "end of CoverTab[64907]"
//line /usr/local/go/src/regexp/onepass.go:458
	_go_fuzz_dep_.CoverTab[64908]++
						return p
//line /usr/local/go/src/regexp/onepass.go:459
	// _ = "end of CoverTab[64908]"
}

// compileOnePass returns a new *syntax.Prog suitable for onePass execution if the original Prog
//line /usr/local/go/src/regexp/onepass.go:462
// can be recharacterized as a one-pass regexp program, or syntax.nil if the
//line /usr/local/go/src/regexp/onepass.go:462
// Prog cannot be converted. For a one pass prog, the fundamental condition that must
//line /usr/local/go/src/regexp/onepass.go:462
// be true is: at any InstAlt, there must be no ambiguity about what branch to  take.
//line /usr/local/go/src/regexp/onepass.go:466
func compileOnePass(prog *syntax.Prog) (p *onePassProg) {
//line /usr/local/go/src/regexp/onepass.go:466
	_go_fuzz_dep_.CoverTab[64978]++
						if prog.Start == 0 {
//line /usr/local/go/src/regexp/onepass.go:467
		_go_fuzz_dep_.CoverTab[64983]++
							return nil
//line /usr/local/go/src/regexp/onepass.go:468
		// _ = "end of CoverTab[64983]"
	} else {
//line /usr/local/go/src/regexp/onepass.go:469
		_go_fuzz_dep_.CoverTab[64984]++
//line /usr/local/go/src/regexp/onepass.go:469
		// _ = "end of CoverTab[64984]"
//line /usr/local/go/src/regexp/onepass.go:469
	}
//line /usr/local/go/src/regexp/onepass.go:469
	// _ = "end of CoverTab[64978]"
//line /usr/local/go/src/regexp/onepass.go:469
	_go_fuzz_dep_.CoverTab[64979]++

						if prog.Inst[prog.Start].Op != syntax.InstEmptyWidth || func() bool {
//line /usr/local/go/src/regexp/onepass.go:471
		_go_fuzz_dep_.CoverTab[64985]++
//line /usr/local/go/src/regexp/onepass.go:471
		return syntax.EmptyOp(prog.Inst[prog.Start].Arg)&syntax.EmptyBeginText != syntax.EmptyBeginText
							// _ = "end of CoverTab[64985]"
//line /usr/local/go/src/regexp/onepass.go:472
	}() {
//line /usr/local/go/src/regexp/onepass.go:472
		_go_fuzz_dep_.CoverTab[64986]++
							return nil
//line /usr/local/go/src/regexp/onepass.go:473
		// _ = "end of CoverTab[64986]"
	} else {
//line /usr/local/go/src/regexp/onepass.go:474
		_go_fuzz_dep_.CoverTab[64987]++
//line /usr/local/go/src/regexp/onepass.go:474
		// _ = "end of CoverTab[64987]"
//line /usr/local/go/src/regexp/onepass.go:474
	}
//line /usr/local/go/src/regexp/onepass.go:474
	// _ = "end of CoverTab[64979]"
//line /usr/local/go/src/regexp/onepass.go:474
	_go_fuzz_dep_.CoverTab[64980]++

						for _, inst := range prog.Inst {
//line /usr/local/go/src/regexp/onepass.go:476
		_go_fuzz_dep_.CoverTab[64988]++
							opOut := prog.Inst[inst.Out].Op
							switch inst.Op {
		default:
//line /usr/local/go/src/regexp/onepass.go:479
			_go_fuzz_dep_.CoverTab[64989]++
								if opOut == syntax.InstMatch {
//line /usr/local/go/src/regexp/onepass.go:480
				_go_fuzz_dep_.CoverTab[64992]++
									return nil
//line /usr/local/go/src/regexp/onepass.go:481
				// _ = "end of CoverTab[64992]"
			} else {
//line /usr/local/go/src/regexp/onepass.go:482
				_go_fuzz_dep_.CoverTab[64993]++
//line /usr/local/go/src/regexp/onepass.go:482
				// _ = "end of CoverTab[64993]"
//line /usr/local/go/src/regexp/onepass.go:482
			}
//line /usr/local/go/src/regexp/onepass.go:482
			// _ = "end of CoverTab[64989]"
		case syntax.InstAlt, syntax.InstAltMatch:
//line /usr/local/go/src/regexp/onepass.go:483
			_go_fuzz_dep_.CoverTab[64990]++
								if opOut == syntax.InstMatch || func() bool {
//line /usr/local/go/src/regexp/onepass.go:484
				_go_fuzz_dep_.CoverTab[64994]++
//line /usr/local/go/src/regexp/onepass.go:484
				return prog.Inst[inst.Arg].Op == syntax.InstMatch
//line /usr/local/go/src/regexp/onepass.go:484
				// _ = "end of CoverTab[64994]"
//line /usr/local/go/src/regexp/onepass.go:484
			}() {
//line /usr/local/go/src/regexp/onepass.go:484
				_go_fuzz_dep_.CoverTab[64995]++
									return nil
//line /usr/local/go/src/regexp/onepass.go:485
				// _ = "end of CoverTab[64995]"
			} else {
//line /usr/local/go/src/regexp/onepass.go:486
				_go_fuzz_dep_.CoverTab[64996]++
//line /usr/local/go/src/regexp/onepass.go:486
				// _ = "end of CoverTab[64996]"
//line /usr/local/go/src/regexp/onepass.go:486
			}
//line /usr/local/go/src/regexp/onepass.go:486
			// _ = "end of CoverTab[64990]"
		case syntax.InstEmptyWidth:
//line /usr/local/go/src/regexp/onepass.go:487
			_go_fuzz_dep_.CoverTab[64991]++
								if opOut == syntax.InstMatch {
//line /usr/local/go/src/regexp/onepass.go:488
				_go_fuzz_dep_.CoverTab[64997]++
									if syntax.EmptyOp(inst.Arg)&syntax.EmptyEndText == syntax.EmptyEndText {
//line /usr/local/go/src/regexp/onepass.go:489
					_go_fuzz_dep_.CoverTab[64999]++
										continue
//line /usr/local/go/src/regexp/onepass.go:490
					// _ = "end of CoverTab[64999]"
				} else {
//line /usr/local/go/src/regexp/onepass.go:491
					_go_fuzz_dep_.CoverTab[65000]++
//line /usr/local/go/src/regexp/onepass.go:491
					// _ = "end of CoverTab[65000]"
//line /usr/local/go/src/regexp/onepass.go:491
				}
//line /usr/local/go/src/regexp/onepass.go:491
				// _ = "end of CoverTab[64997]"
//line /usr/local/go/src/regexp/onepass.go:491
				_go_fuzz_dep_.CoverTab[64998]++
									return nil
//line /usr/local/go/src/regexp/onepass.go:492
				// _ = "end of CoverTab[64998]"
			} else {
//line /usr/local/go/src/regexp/onepass.go:493
				_go_fuzz_dep_.CoverTab[65001]++
//line /usr/local/go/src/regexp/onepass.go:493
				// _ = "end of CoverTab[65001]"
//line /usr/local/go/src/regexp/onepass.go:493
			}
//line /usr/local/go/src/regexp/onepass.go:493
			// _ = "end of CoverTab[64991]"
		}
//line /usr/local/go/src/regexp/onepass.go:494
		// _ = "end of CoverTab[64988]"
	}
//line /usr/local/go/src/regexp/onepass.go:495
	// _ = "end of CoverTab[64980]"
//line /usr/local/go/src/regexp/onepass.go:495
	_go_fuzz_dep_.CoverTab[64981]++

//line /usr/local/go/src/regexp/onepass.go:498
	p = onePassCopy(prog)

//line /usr/local/go/src/regexp/onepass.go:501
	p = makeOnePass(p)

	if p != nil {
//line /usr/local/go/src/regexp/onepass.go:503
		_go_fuzz_dep_.CoverTab[65002]++
							cleanupOnePass(p, prog)
//line /usr/local/go/src/regexp/onepass.go:504
		// _ = "end of CoverTab[65002]"
	} else {
//line /usr/local/go/src/regexp/onepass.go:505
		_go_fuzz_dep_.CoverTab[65003]++
//line /usr/local/go/src/regexp/onepass.go:505
		// _ = "end of CoverTab[65003]"
//line /usr/local/go/src/regexp/onepass.go:505
	}
//line /usr/local/go/src/regexp/onepass.go:505
	// _ = "end of CoverTab[64981]"
//line /usr/local/go/src/regexp/onepass.go:505
	_go_fuzz_dep_.CoverTab[64982]++
						return p
//line /usr/local/go/src/regexp/onepass.go:506
	// _ = "end of CoverTab[64982]"
}

//line /usr/local/go/src/regexp/onepass.go:507
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/regexp/onepass.go:507
var _ = _go_fuzz_dep_.CoverTab
