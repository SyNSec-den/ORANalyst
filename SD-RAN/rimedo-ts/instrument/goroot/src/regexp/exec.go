// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/regexp/exec.go:5
package regexp

//line /usr/local/go/src/regexp/exec.go:5
import (
//line /usr/local/go/src/regexp/exec.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/regexp/exec.go:5
)
//line /usr/local/go/src/regexp/exec.go:5
import (
//line /usr/local/go/src/regexp/exec.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/regexp/exec.go:5
)

import (
	"io"
	"regexp/syntax"
	"sync"
)

// A queue is a 'sparse array' holding pending threads of execution.
//line /usr/local/go/src/regexp/exec.go:13
// See https://research.swtch.com/2008/03/using-uninitialized-memory-for-fun-and.html
//line /usr/local/go/src/regexp/exec.go:15
type queue struct {
	sparse	[]uint32
	dense	[]entry
}

// An entry is an entry on a queue.
//line /usr/local/go/src/regexp/exec.go:20
// It holds both the instruction pc and the actual thread.
//line /usr/local/go/src/regexp/exec.go:20
// Some queue entries are just place holders so that the machine
//line /usr/local/go/src/regexp/exec.go:20
// knows it has considered that pc. Such entries have t == nil.
//line /usr/local/go/src/regexp/exec.go:24
type entry struct {
	pc	uint32
	t	*thread
}

// A thread is the state of a single path through the machine:
//line /usr/local/go/src/regexp/exec.go:29
// an instruction and a corresponding capture array.
//line /usr/local/go/src/regexp/exec.go:29
// See https://swtch.com/~rsc/regexp/regexp2.html
//line /usr/local/go/src/regexp/exec.go:32
type thread struct {
	inst	*syntax.Inst
	cap	[]int
}

// A machine holds all the state during an NFA simulation for p.
type machine struct {
	re		*Regexp		// corresponding Regexp
	p		*syntax.Prog	// compiled program
	q0, q1		queue		// two queues for runq, nextq
	pool		[]*thread	// pool of available threads
	matched		bool		// whether a match was found
	matchcap	[]int		// capture information for the match

	inputs	inputs
}

type inputs struct {
	// cached inputs, to avoid allocation
	bytes	inputBytes
	string	inputString
	reader	inputReader
}

func (i *inputs) newBytes(b []byte) input {
//line /usr/local/go/src/regexp/exec.go:56
	_go_fuzz_dep_.CoverTab[64541]++
						i.bytes.str = b
						return &i.bytes
//line /usr/local/go/src/regexp/exec.go:58
	// _ = "end of CoverTab[64541]"
}

func (i *inputs) newString(s string) input {
//line /usr/local/go/src/regexp/exec.go:61
	_go_fuzz_dep_.CoverTab[64542]++
						i.string.str = s
						return &i.string
//line /usr/local/go/src/regexp/exec.go:63
	// _ = "end of CoverTab[64542]"
}

func (i *inputs) newReader(r io.RuneReader) input {
//line /usr/local/go/src/regexp/exec.go:66
	_go_fuzz_dep_.CoverTab[64543]++
						i.reader.r = r
						i.reader.atEOT = false
						i.reader.pos = 0
						return &i.reader
//line /usr/local/go/src/regexp/exec.go:70
	// _ = "end of CoverTab[64543]"
}

func (i *inputs) clear() {
//line /usr/local/go/src/regexp/exec.go:73
	_go_fuzz_dep_.CoverTab[64544]++

//line /usr/local/go/src/regexp/exec.go:76
	if i.bytes.str != nil {
//line /usr/local/go/src/regexp/exec.go:76
		_go_fuzz_dep_.CoverTab[64545]++
							i.bytes.str = nil
//line /usr/local/go/src/regexp/exec.go:77
		// _ = "end of CoverTab[64545]"
	} else {
//line /usr/local/go/src/regexp/exec.go:78
		_go_fuzz_dep_.CoverTab[64546]++
//line /usr/local/go/src/regexp/exec.go:78
		if i.reader.r != nil {
//line /usr/local/go/src/regexp/exec.go:78
			_go_fuzz_dep_.CoverTab[64547]++
								i.reader.r = nil
//line /usr/local/go/src/regexp/exec.go:79
			// _ = "end of CoverTab[64547]"
		} else {
//line /usr/local/go/src/regexp/exec.go:80
			_go_fuzz_dep_.CoverTab[64548]++
								i.string.str = ""
//line /usr/local/go/src/regexp/exec.go:81
			// _ = "end of CoverTab[64548]"
		}
//line /usr/local/go/src/regexp/exec.go:82
		// _ = "end of CoverTab[64546]"
//line /usr/local/go/src/regexp/exec.go:82
	}
//line /usr/local/go/src/regexp/exec.go:82
	// _ = "end of CoverTab[64544]"
}

func (i *inputs) init(r io.RuneReader, b []byte, s string) (input, int) {
	if r != nil {
		return i.newReader(r), 0
	}
	if b != nil {
		return i.newBytes(b), len(b)
	}
	return i.newString(s), len(s)
}

func (m *machine) init(ncap int) {
	for _, t := range m.pool {
		t.cap = t.cap[:ncap]
	}
	m.matchcap = m.matchcap[:ncap]
}

// alloc allocates a new thread with the given instruction.
//line /usr/local/go/src/regexp/exec.go:102
// It uses the free pool if possible.
//line /usr/local/go/src/regexp/exec.go:104
func (m *machine) alloc(i *syntax.Inst) *thread {
//line /usr/local/go/src/regexp/exec.go:104
	_go_fuzz_dep_.CoverTab[64549]++
						var t *thread
						if n := len(m.pool); n > 0 {
//line /usr/local/go/src/regexp/exec.go:106
		_go_fuzz_dep_.CoverTab[64551]++
							t = m.pool[n-1]
							m.pool = m.pool[:n-1]
//line /usr/local/go/src/regexp/exec.go:108
		// _ = "end of CoverTab[64551]"
	} else {
//line /usr/local/go/src/regexp/exec.go:109
		_go_fuzz_dep_.CoverTab[64552]++
							t = new(thread)
							t.cap = make([]int, len(m.matchcap), cap(m.matchcap))
//line /usr/local/go/src/regexp/exec.go:111
		// _ = "end of CoverTab[64552]"
	}
//line /usr/local/go/src/regexp/exec.go:112
	// _ = "end of CoverTab[64549]"
//line /usr/local/go/src/regexp/exec.go:112
	_go_fuzz_dep_.CoverTab[64550]++
						t.inst = i
						return t
//line /usr/local/go/src/regexp/exec.go:114
	// _ = "end of CoverTab[64550]"
}

// A lazyFlag is a lazily-evaluated syntax.EmptyOp,
//line /usr/local/go/src/regexp/exec.go:117
// for checking zero-width flags like ^ $ \A \z \B \b.
//line /usr/local/go/src/regexp/exec.go:117
// It records the pair of relevant runes and does not
//line /usr/local/go/src/regexp/exec.go:117
// determine the implied flags until absolutely necessary
//line /usr/local/go/src/regexp/exec.go:117
// (most of the time, that means never).
//line /usr/local/go/src/regexp/exec.go:122
type lazyFlag uint64

func newLazyFlag(r1, r2 rune) lazyFlag {
//line /usr/local/go/src/regexp/exec.go:124
	_go_fuzz_dep_.CoverTab[64553]++
						return lazyFlag(uint64(r1)<<32 | uint64(uint32(r2)))
//line /usr/local/go/src/regexp/exec.go:125
	// _ = "end of CoverTab[64553]"
}

func (f lazyFlag) match(op syntax.EmptyOp) bool {
//line /usr/local/go/src/regexp/exec.go:128
	_go_fuzz_dep_.CoverTab[64554]++
						if op == 0 {
//line /usr/local/go/src/regexp/exec.go:129
		_go_fuzz_dep_.CoverTab[64563]++
							return true
//line /usr/local/go/src/regexp/exec.go:130
		// _ = "end of CoverTab[64563]"
	} else {
//line /usr/local/go/src/regexp/exec.go:131
		_go_fuzz_dep_.CoverTab[64564]++
//line /usr/local/go/src/regexp/exec.go:131
		// _ = "end of CoverTab[64564]"
//line /usr/local/go/src/regexp/exec.go:131
	}
//line /usr/local/go/src/regexp/exec.go:131
	// _ = "end of CoverTab[64554]"
//line /usr/local/go/src/regexp/exec.go:131
	_go_fuzz_dep_.CoverTab[64555]++
						r1 := rune(f >> 32)
						if op&syntax.EmptyBeginLine != 0 {
//line /usr/local/go/src/regexp/exec.go:133
		_go_fuzz_dep_.CoverTab[64565]++
							if r1 != '\n' && func() bool {
//line /usr/local/go/src/regexp/exec.go:134
			_go_fuzz_dep_.CoverTab[64567]++
//line /usr/local/go/src/regexp/exec.go:134
			return r1 >= 0
//line /usr/local/go/src/regexp/exec.go:134
			// _ = "end of CoverTab[64567]"
//line /usr/local/go/src/regexp/exec.go:134
		}() {
//line /usr/local/go/src/regexp/exec.go:134
			_go_fuzz_dep_.CoverTab[64568]++
								return false
//line /usr/local/go/src/regexp/exec.go:135
			// _ = "end of CoverTab[64568]"
		} else {
//line /usr/local/go/src/regexp/exec.go:136
			_go_fuzz_dep_.CoverTab[64569]++
//line /usr/local/go/src/regexp/exec.go:136
			// _ = "end of CoverTab[64569]"
//line /usr/local/go/src/regexp/exec.go:136
		}
//line /usr/local/go/src/regexp/exec.go:136
		// _ = "end of CoverTab[64565]"
//line /usr/local/go/src/regexp/exec.go:136
		_go_fuzz_dep_.CoverTab[64566]++
							op &^= syntax.EmptyBeginLine
//line /usr/local/go/src/regexp/exec.go:137
		// _ = "end of CoverTab[64566]"
	} else {
//line /usr/local/go/src/regexp/exec.go:138
		_go_fuzz_dep_.CoverTab[64570]++
//line /usr/local/go/src/regexp/exec.go:138
		// _ = "end of CoverTab[64570]"
//line /usr/local/go/src/regexp/exec.go:138
	}
//line /usr/local/go/src/regexp/exec.go:138
	// _ = "end of CoverTab[64555]"
//line /usr/local/go/src/regexp/exec.go:138
	_go_fuzz_dep_.CoverTab[64556]++
						if op&syntax.EmptyBeginText != 0 {
//line /usr/local/go/src/regexp/exec.go:139
		_go_fuzz_dep_.CoverTab[64571]++
							if r1 >= 0 {
//line /usr/local/go/src/regexp/exec.go:140
			_go_fuzz_dep_.CoverTab[64573]++
								return false
//line /usr/local/go/src/regexp/exec.go:141
			// _ = "end of CoverTab[64573]"
		} else {
//line /usr/local/go/src/regexp/exec.go:142
			_go_fuzz_dep_.CoverTab[64574]++
//line /usr/local/go/src/regexp/exec.go:142
			// _ = "end of CoverTab[64574]"
//line /usr/local/go/src/regexp/exec.go:142
		}
//line /usr/local/go/src/regexp/exec.go:142
		// _ = "end of CoverTab[64571]"
//line /usr/local/go/src/regexp/exec.go:142
		_go_fuzz_dep_.CoverTab[64572]++
							op &^= syntax.EmptyBeginText
//line /usr/local/go/src/regexp/exec.go:143
		// _ = "end of CoverTab[64572]"
	} else {
//line /usr/local/go/src/regexp/exec.go:144
		_go_fuzz_dep_.CoverTab[64575]++
//line /usr/local/go/src/regexp/exec.go:144
		// _ = "end of CoverTab[64575]"
//line /usr/local/go/src/regexp/exec.go:144
	}
//line /usr/local/go/src/regexp/exec.go:144
	// _ = "end of CoverTab[64556]"
//line /usr/local/go/src/regexp/exec.go:144
	_go_fuzz_dep_.CoverTab[64557]++
						if op == 0 {
//line /usr/local/go/src/regexp/exec.go:145
		_go_fuzz_dep_.CoverTab[64576]++
							return true
//line /usr/local/go/src/regexp/exec.go:146
		// _ = "end of CoverTab[64576]"
	} else {
//line /usr/local/go/src/regexp/exec.go:147
		_go_fuzz_dep_.CoverTab[64577]++
//line /usr/local/go/src/regexp/exec.go:147
		// _ = "end of CoverTab[64577]"
//line /usr/local/go/src/regexp/exec.go:147
	}
//line /usr/local/go/src/regexp/exec.go:147
	// _ = "end of CoverTab[64557]"
//line /usr/local/go/src/regexp/exec.go:147
	_go_fuzz_dep_.CoverTab[64558]++
						r2 := rune(f)
						if op&syntax.EmptyEndLine != 0 {
//line /usr/local/go/src/regexp/exec.go:149
		_go_fuzz_dep_.CoverTab[64578]++
							if r2 != '\n' && func() bool {
//line /usr/local/go/src/regexp/exec.go:150
			_go_fuzz_dep_.CoverTab[64580]++
//line /usr/local/go/src/regexp/exec.go:150
			return r2 >= 0
//line /usr/local/go/src/regexp/exec.go:150
			// _ = "end of CoverTab[64580]"
//line /usr/local/go/src/regexp/exec.go:150
		}() {
//line /usr/local/go/src/regexp/exec.go:150
			_go_fuzz_dep_.CoverTab[64581]++
								return false
//line /usr/local/go/src/regexp/exec.go:151
			// _ = "end of CoverTab[64581]"
		} else {
//line /usr/local/go/src/regexp/exec.go:152
			_go_fuzz_dep_.CoverTab[64582]++
//line /usr/local/go/src/regexp/exec.go:152
			// _ = "end of CoverTab[64582]"
//line /usr/local/go/src/regexp/exec.go:152
		}
//line /usr/local/go/src/regexp/exec.go:152
		// _ = "end of CoverTab[64578]"
//line /usr/local/go/src/regexp/exec.go:152
		_go_fuzz_dep_.CoverTab[64579]++
							op &^= syntax.EmptyEndLine
//line /usr/local/go/src/regexp/exec.go:153
		// _ = "end of CoverTab[64579]"
	} else {
//line /usr/local/go/src/regexp/exec.go:154
		_go_fuzz_dep_.CoverTab[64583]++
//line /usr/local/go/src/regexp/exec.go:154
		// _ = "end of CoverTab[64583]"
//line /usr/local/go/src/regexp/exec.go:154
	}
//line /usr/local/go/src/regexp/exec.go:154
	// _ = "end of CoverTab[64558]"
//line /usr/local/go/src/regexp/exec.go:154
	_go_fuzz_dep_.CoverTab[64559]++
						if op&syntax.EmptyEndText != 0 {
//line /usr/local/go/src/regexp/exec.go:155
		_go_fuzz_dep_.CoverTab[64584]++
							if r2 >= 0 {
//line /usr/local/go/src/regexp/exec.go:156
			_go_fuzz_dep_.CoverTab[64586]++
								return false
//line /usr/local/go/src/regexp/exec.go:157
			// _ = "end of CoverTab[64586]"
		} else {
//line /usr/local/go/src/regexp/exec.go:158
			_go_fuzz_dep_.CoverTab[64587]++
//line /usr/local/go/src/regexp/exec.go:158
			// _ = "end of CoverTab[64587]"
//line /usr/local/go/src/regexp/exec.go:158
		}
//line /usr/local/go/src/regexp/exec.go:158
		// _ = "end of CoverTab[64584]"
//line /usr/local/go/src/regexp/exec.go:158
		_go_fuzz_dep_.CoverTab[64585]++
							op &^= syntax.EmptyEndText
//line /usr/local/go/src/regexp/exec.go:159
		// _ = "end of CoverTab[64585]"
	} else {
//line /usr/local/go/src/regexp/exec.go:160
		_go_fuzz_dep_.CoverTab[64588]++
//line /usr/local/go/src/regexp/exec.go:160
		// _ = "end of CoverTab[64588]"
//line /usr/local/go/src/regexp/exec.go:160
	}
//line /usr/local/go/src/regexp/exec.go:160
	// _ = "end of CoverTab[64559]"
//line /usr/local/go/src/regexp/exec.go:160
	_go_fuzz_dep_.CoverTab[64560]++
						if op == 0 {
//line /usr/local/go/src/regexp/exec.go:161
		_go_fuzz_dep_.CoverTab[64589]++
							return true
//line /usr/local/go/src/regexp/exec.go:162
		// _ = "end of CoverTab[64589]"
	} else {
//line /usr/local/go/src/regexp/exec.go:163
		_go_fuzz_dep_.CoverTab[64590]++
//line /usr/local/go/src/regexp/exec.go:163
		// _ = "end of CoverTab[64590]"
//line /usr/local/go/src/regexp/exec.go:163
	}
//line /usr/local/go/src/regexp/exec.go:163
	// _ = "end of CoverTab[64560]"
//line /usr/local/go/src/regexp/exec.go:163
	_go_fuzz_dep_.CoverTab[64561]++
						if syntax.IsWordChar(r1) != syntax.IsWordChar(r2) {
//line /usr/local/go/src/regexp/exec.go:164
		_go_fuzz_dep_.CoverTab[64591]++
							op &^= syntax.EmptyWordBoundary
//line /usr/local/go/src/regexp/exec.go:165
		// _ = "end of CoverTab[64591]"
	} else {
//line /usr/local/go/src/regexp/exec.go:166
		_go_fuzz_dep_.CoverTab[64592]++
							op &^= syntax.EmptyNoWordBoundary
//line /usr/local/go/src/regexp/exec.go:167
		// _ = "end of CoverTab[64592]"
	}
//line /usr/local/go/src/regexp/exec.go:168
	// _ = "end of CoverTab[64561]"
//line /usr/local/go/src/regexp/exec.go:168
	_go_fuzz_dep_.CoverTab[64562]++
						return op == 0
//line /usr/local/go/src/regexp/exec.go:169
	// _ = "end of CoverTab[64562]"
}

// match runs the machine over the input starting at pos.
//line /usr/local/go/src/regexp/exec.go:172
// It reports whether a match was found.
//line /usr/local/go/src/regexp/exec.go:172
// If so, m.matchcap holds the submatch information.
//line /usr/local/go/src/regexp/exec.go:175
func (m *machine) match(i input, pos int) bool {
//line /usr/local/go/src/regexp/exec.go:175
	_go_fuzz_dep_.CoverTab[64593]++
						startCond := m.re.cond
						if startCond == ^syntax.EmptyOp(0) {
//line /usr/local/go/src/regexp/exec.go:177
		_go_fuzz_dep_.CoverTab[64599]++
							return false
//line /usr/local/go/src/regexp/exec.go:178
		// _ = "end of CoverTab[64599]"
	} else {
//line /usr/local/go/src/regexp/exec.go:179
		_go_fuzz_dep_.CoverTab[64600]++
//line /usr/local/go/src/regexp/exec.go:179
		// _ = "end of CoverTab[64600]"
//line /usr/local/go/src/regexp/exec.go:179
	}
//line /usr/local/go/src/regexp/exec.go:179
	// _ = "end of CoverTab[64593]"
//line /usr/local/go/src/regexp/exec.go:179
	_go_fuzz_dep_.CoverTab[64594]++
						m.matched = false
						for i := range m.matchcap {
//line /usr/local/go/src/regexp/exec.go:181
		_go_fuzz_dep_.CoverTab[64601]++
							m.matchcap[i] = -1
//line /usr/local/go/src/regexp/exec.go:182
		// _ = "end of CoverTab[64601]"
	}
//line /usr/local/go/src/regexp/exec.go:183
	// _ = "end of CoverTab[64594]"
//line /usr/local/go/src/regexp/exec.go:183
	_go_fuzz_dep_.CoverTab[64595]++
						runq, nextq := &m.q0, &m.q1
						r, r1 := endOfText, endOfText
						width, width1 := 0, 0
						r, width = i.step(pos)
						if r != endOfText {
//line /usr/local/go/src/regexp/exec.go:188
		_go_fuzz_dep_.CoverTab[64602]++
							r1, width1 = i.step(pos + width)
//line /usr/local/go/src/regexp/exec.go:189
		// _ = "end of CoverTab[64602]"
	} else {
//line /usr/local/go/src/regexp/exec.go:190
		_go_fuzz_dep_.CoverTab[64603]++
//line /usr/local/go/src/regexp/exec.go:190
		// _ = "end of CoverTab[64603]"
//line /usr/local/go/src/regexp/exec.go:190
	}
//line /usr/local/go/src/regexp/exec.go:190
	// _ = "end of CoverTab[64595]"
//line /usr/local/go/src/regexp/exec.go:190
	_go_fuzz_dep_.CoverTab[64596]++
						var flag lazyFlag
						if pos == 0 {
//line /usr/local/go/src/regexp/exec.go:192
		_go_fuzz_dep_.CoverTab[64604]++
							flag = newLazyFlag(-1, r)
//line /usr/local/go/src/regexp/exec.go:193
		// _ = "end of CoverTab[64604]"
	} else {
//line /usr/local/go/src/regexp/exec.go:194
		_go_fuzz_dep_.CoverTab[64605]++
							flag = i.context(pos)
//line /usr/local/go/src/regexp/exec.go:195
		// _ = "end of CoverTab[64605]"
	}
//line /usr/local/go/src/regexp/exec.go:196
	// _ = "end of CoverTab[64596]"
//line /usr/local/go/src/regexp/exec.go:196
	_go_fuzz_dep_.CoverTab[64597]++
						for {
//line /usr/local/go/src/regexp/exec.go:197
		_go_fuzz_dep_.CoverTab[64606]++
							if len(runq.dense) == 0 {
//line /usr/local/go/src/regexp/exec.go:198
			_go_fuzz_dep_.CoverTab[64612]++
								if startCond&syntax.EmptyBeginText != 0 && func() bool {
//line /usr/local/go/src/regexp/exec.go:199
				_go_fuzz_dep_.CoverTab[64615]++
//line /usr/local/go/src/regexp/exec.go:199
				return pos != 0
//line /usr/local/go/src/regexp/exec.go:199
				// _ = "end of CoverTab[64615]"
//line /usr/local/go/src/regexp/exec.go:199
			}() {
//line /usr/local/go/src/regexp/exec.go:199
				_go_fuzz_dep_.CoverTab[64616]++

									break
//line /usr/local/go/src/regexp/exec.go:201
				// _ = "end of CoverTab[64616]"
			} else {
//line /usr/local/go/src/regexp/exec.go:202
				_go_fuzz_dep_.CoverTab[64617]++
//line /usr/local/go/src/regexp/exec.go:202
				// _ = "end of CoverTab[64617]"
//line /usr/local/go/src/regexp/exec.go:202
			}
//line /usr/local/go/src/regexp/exec.go:202
			// _ = "end of CoverTab[64612]"
//line /usr/local/go/src/regexp/exec.go:202
			_go_fuzz_dep_.CoverTab[64613]++
								if m.matched {
//line /usr/local/go/src/regexp/exec.go:203
				_go_fuzz_dep_.CoverTab[64618]++

									break
//line /usr/local/go/src/regexp/exec.go:205
				// _ = "end of CoverTab[64618]"
			} else {
//line /usr/local/go/src/regexp/exec.go:206
				_go_fuzz_dep_.CoverTab[64619]++
//line /usr/local/go/src/regexp/exec.go:206
				// _ = "end of CoverTab[64619]"
//line /usr/local/go/src/regexp/exec.go:206
			}
//line /usr/local/go/src/regexp/exec.go:206
			// _ = "end of CoverTab[64613]"
//line /usr/local/go/src/regexp/exec.go:206
			_go_fuzz_dep_.CoverTab[64614]++
								if len(m.re.prefix) > 0 && func() bool {
//line /usr/local/go/src/regexp/exec.go:207
				_go_fuzz_dep_.CoverTab[64620]++
//line /usr/local/go/src/regexp/exec.go:207
				return r1 != m.re.prefixRune
//line /usr/local/go/src/regexp/exec.go:207
				// _ = "end of CoverTab[64620]"
//line /usr/local/go/src/regexp/exec.go:207
			}() && func() bool {
//line /usr/local/go/src/regexp/exec.go:207
				_go_fuzz_dep_.CoverTab[64621]++
//line /usr/local/go/src/regexp/exec.go:207
				return i.canCheckPrefix()
//line /usr/local/go/src/regexp/exec.go:207
				// _ = "end of CoverTab[64621]"
//line /usr/local/go/src/regexp/exec.go:207
			}() {
//line /usr/local/go/src/regexp/exec.go:207
				_go_fuzz_dep_.CoverTab[64622]++

									advance := i.index(m.re, pos)
									if advance < 0 {
//line /usr/local/go/src/regexp/exec.go:210
					_go_fuzz_dep_.CoverTab[64624]++
										break
//line /usr/local/go/src/regexp/exec.go:211
					// _ = "end of CoverTab[64624]"
				} else {
//line /usr/local/go/src/regexp/exec.go:212
					_go_fuzz_dep_.CoverTab[64625]++
//line /usr/local/go/src/regexp/exec.go:212
					// _ = "end of CoverTab[64625]"
//line /usr/local/go/src/regexp/exec.go:212
				}
//line /usr/local/go/src/regexp/exec.go:212
				// _ = "end of CoverTab[64622]"
//line /usr/local/go/src/regexp/exec.go:212
				_go_fuzz_dep_.CoverTab[64623]++
									pos += advance
									r, width = i.step(pos)
									r1, width1 = i.step(pos + width)
//line /usr/local/go/src/regexp/exec.go:215
				// _ = "end of CoverTab[64623]"
			} else {
//line /usr/local/go/src/regexp/exec.go:216
				_go_fuzz_dep_.CoverTab[64626]++
//line /usr/local/go/src/regexp/exec.go:216
				// _ = "end of CoverTab[64626]"
//line /usr/local/go/src/regexp/exec.go:216
			}
//line /usr/local/go/src/regexp/exec.go:216
			// _ = "end of CoverTab[64614]"
		} else {
//line /usr/local/go/src/regexp/exec.go:217
			_go_fuzz_dep_.CoverTab[64627]++
//line /usr/local/go/src/regexp/exec.go:217
			// _ = "end of CoverTab[64627]"
//line /usr/local/go/src/regexp/exec.go:217
		}
//line /usr/local/go/src/regexp/exec.go:217
		// _ = "end of CoverTab[64606]"
//line /usr/local/go/src/regexp/exec.go:217
		_go_fuzz_dep_.CoverTab[64607]++
							if !m.matched {
//line /usr/local/go/src/regexp/exec.go:218
			_go_fuzz_dep_.CoverTab[64628]++
								if len(m.matchcap) > 0 {
//line /usr/local/go/src/regexp/exec.go:219
				_go_fuzz_dep_.CoverTab[64630]++
									m.matchcap[0] = pos
//line /usr/local/go/src/regexp/exec.go:220
				// _ = "end of CoverTab[64630]"
			} else {
//line /usr/local/go/src/regexp/exec.go:221
				_go_fuzz_dep_.CoverTab[64631]++
//line /usr/local/go/src/regexp/exec.go:221
				// _ = "end of CoverTab[64631]"
//line /usr/local/go/src/regexp/exec.go:221
			}
//line /usr/local/go/src/regexp/exec.go:221
			// _ = "end of CoverTab[64628]"
//line /usr/local/go/src/regexp/exec.go:221
			_go_fuzz_dep_.CoverTab[64629]++
								m.add(runq, uint32(m.p.Start), pos, m.matchcap, &flag, nil)
//line /usr/local/go/src/regexp/exec.go:222
			// _ = "end of CoverTab[64629]"
		} else {
//line /usr/local/go/src/regexp/exec.go:223
			_go_fuzz_dep_.CoverTab[64632]++
//line /usr/local/go/src/regexp/exec.go:223
			// _ = "end of CoverTab[64632]"
//line /usr/local/go/src/regexp/exec.go:223
		}
//line /usr/local/go/src/regexp/exec.go:223
		// _ = "end of CoverTab[64607]"
//line /usr/local/go/src/regexp/exec.go:223
		_go_fuzz_dep_.CoverTab[64608]++
							flag = newLazyFlag(r, r1)
							m.step(runq, nextq, pos, pos+width, r, &flag)
							if width == 0 {
//line /usr/local/go/src/regexp/exec.go:226
			_go_fuzz_dep_.CoverTab[64633]++
								break
//line /usr/local/go/src/regexp/exec.go:227
			// _ = "end of CoverTab[64633]"
		} else {
//line /usr/local/go/src/regexp/exec.go:228
			_go_fuzz_dep_.CoverTab[64634]++
//line /usr/local/go/src/regexp/exec.go:228
			// _ = "end of CoverTab[64634]"
//line /usr/local/go/src/regexp/exec.go:228
		}
//line /usr/local/go/src/regexp/exec.go:228
		// _ = "end of CoverTab[64608]"
//line /usr/local/go/src/regexp/exec.go:228
		_go_fuzz_dep_.CoverTab[64609]++
							if len(m.matchcap) == 0 && func() bool {
//line /usr/local/go/src/regexp/exec.go:229
			_go_fuzz_dep_.CoverTab[64635]++
//line /usr/local/go/src/regexp/exec.go:229
			return m.matched
//line /usr/local/go/src/regexp/exec.go:229
			// _ = "end of CoverTab[64635]"
//line /usr/local/go/src/regexp/exec.go:229
		}() {
//line /usr/local/go/src/regexp/exec.go:229
			_go_fuzz_dep_.CoverTab[64636]++

//line /usr/local/go/src/regexp/exec.go:232
			break
//line /usr/local/go/src/regexp/exec.go:232
			// _ = "end of CoverTab[64636]"
		} else {
//line /usr/local/go/src/regexp/exec.go:233
			_go_fuzz_dep_.CoverTab[64637]++
//line /usr/local/go/src/regexp/exec.go:233
			// _ = "end of CoverTab[64637]"
//line /usr/local/go/src/regexp/exec.go:233
		}
//line /usr/local/go/src/regexp/exec.go:233
		// _ = "end of CoverTab[64609]"
//line /usr/local/go/src/regexp/exec.go:233
		_go_fuzz_dep_.CoverTab[64610]++
							pos += width
							r, width = r1, width1
							if r != endOfText {
//line /usr/local/go/src/regexp/exec.go:236
			_go_fuzz_dep_.CoverTab[64638]++
								r1, width1 = i.step(pos + width)
//line /usr/local/go/src/regexp/exec.go:237
			// _ = "end of CoverTab[64638]"
		} else {
//line /usr/local/go/src/regexp/exec.go:238
			_go_fuzz_dep_.CoverTab[64639]++
//line /usr/local/go/src/regexp/exec.go:238
			// _ = "end of CoverTab[64639]"
//line /usr/local/go/src/regexp/exec.go:238
		}
//line /usr/local/go/src/regexp/exec.go:238
		// _ = "end of CoverTab[64610]"
//line /usr/local/go/src/regexp/exec.go:238
		_go_fuzz_dep_.CoverTab[64611]++
							runq, nextq = nextq, runq
//line /usr/local/go/src/regexp/exec.go:239
		// _ = "end of CoverTab[64611]"
	}
//line /usr/local/go/src/regexp/exec.go:240
	// _ = "end of CoverTab[64597]"
//line /usr/local/go/src/regexp/exec.go:240
	_go_fuzz_dep_.CoverTab[64598]++
						m.clear(nextq)
						return m.matched
//line /usr/local/go/src/regexp/exec.go:242
	// _ = "end of CoverTab[64598]"
}

// clear frees all threads on the thread queue.
func (m *machine) clear(q *queue) {
//line /usr/local/go/src/regexp/exec.go:246
	_go_fuzz_dep_.CoverTab[64640]++
						for _, d := range q.dense {
//line /usr/local/go/src/regexp/exec.go:247
		_go_fuzz_dep_.CoverTab[64642]++
							if d.t != nil {
//line /usr/local/go/src/regexp/exec.go:248
			_go_fuzz_dep_.CoverTab[64643]++
								m.pool = append(m.pool, d.t)
//line /usr/local/go/src/regexp/exec.go:249
			// _ = "end of CoverTab[64643]"
		} else {
//line /usr/local/go/src/regexp/exec.go:250
			_go_fuzz_dep_.CoverTab[64644]++
//line /usr/local/go/src/regexp/exec.go:250
			// _ = "end of CoverTab[64644]"
//line /usr/local/go/src/regexp/exec.go:250
		}
//line /usr/local/go/src/regexp/exec.go:250
		// _ = "end of CoverTab[64642]"
	}
//line /usr/local/go/src/regexp/exec.go:251
	// _ = "end of CoverTab[64640]"
//line /usr/local/go/src/regexp/exec.go:251
	_go_fuzz_dep_.CoverTab[64641]++
						q.dense = q.dense[:0]
//line /usr/local/go/src/regexp/exec.go:252
	// _ = "end of CoverTab[64641]"
}

// step executes one step of the machine, running each of the threads
//line /usr/local/go/src/regexp/exec.go:255
// on runq and appending new threads to nextq.
//line /usr/local/go/src/regexp/exec.go:255
// The step processes the rune c (which may be endOfText),
//line /usr/local/go/src/regexp/exec.go:255
// which starts at position pos and ends at nextPos.
//line /usr/local/go/src/regexp/exec.go:255
// nextCond gives the setting for the empty-width flags after c.
//line /usr/local/go/src/regexp/exec.go:260
func (m *machine) step(runq, nextq *queue, pos, nextPos int, c rune, nextCond *lazyFlag) {
//line /usr/local/go/src/regexp/exec.go:260
	_go_fuzz_dep_.CoverTab[64645]++
						longest := m.re.longest
						for j := 0; j < len(runq.dense); j++ {
//line /usr/local/go/src/regexp/exec.go:262
		_go_fuzz_dep_.CoverTab[64647]++
							d := &runq.dense[j]
							t := d.t
							if t == nil {
//line /usr/local/go/src/regexp/exec.go:265
			_go_fuzz_dep_.CoverTab[64652]++
								continue
//line /usr/local/go/src/regexp/exec.go:266
			// _ = "end of CoverTab[64652]"
		} else {
//line /usr/local/go/src/regexp/exec.go:267
			_go_fuzz_dep_.CoverTab[64653]++
//line /usr/local/go/src/regexp/exec.go:267
			// _ = "end of CoverTab[64653]"
//line /usr/local/go/src/regexp/exec.go:267
		}
//line /usr/local/go/src/regexp/exec.go:267
		// _ = "end of CoverTab[64647]"
//line /usr/local/go/src/regexp/exec.go:267
		_go_fuzz_dep_.CoverTab[64648]++
							if longest && func() bool {
//line /usr/local/go/src/regexp/exec.go:268
			_go_fuzz_dep_.CoverTab[64654]++
//line /usr/local/go/src/regexp/exec.go:268
			return m.matched
//line /usr/local/go/src/regexp/exec.go:268
			// _ = "end of CoverTab[64654]"
//line /usr/local/go/src/regexp/exec.go:268
		}() && func() bool {
//line /usr/local/go/src/regexp/exec.go:268
			_go_fuzz_dep_.CoverTab[64655]++
//line /usr/local/go/src/regexp/exec.go:268
			return len(t.cap) > 0
//line /usr/local/go/src/regexp/exec.go:268
			// _ = "end of CoverTab[64655]"
//line /usr/local/go/src/regexp/exec.go:268
		}() && func() bool {
//line /usr/local/go/src/regexp/exec.go:268
			_go_fuzz_dep_.CoverTab[64656]++
//line /usr/local/go/src/regexp/exec.go:268
			return m.matchcap[0] < t.cap[0]
//line /usr/local/go/src/regexp/exec.go:268
			// _ = "end of CoverTab[64656]"
//line /usr/local/go/src/regexp/exec.go:268
		}() {
//line /usr/local/go/src/regexp/exec.go:268
			_go_fuzz_dep_.CoverTab[64657]++
								m.pool = append(m.pool, t)
								continue
//line /usr/local/go/src/regexp/exec.go:270
			// _ = "end of CoverTab[64657]"
		} else {
//line /usr/local/go/src/regexp/exec.go:271
			_go_fuzz_dep_.CoverTab[64658]++
//line /usr/local/go/src/regexp/exec.go:271
			// _ = "end of CoverTab[64658]"
//line /usr/local/go/src/regexp/exec.go:271
		}
//line /usr/local/go/src/regexp/exec.go:271
		// _ = "end of CoverTab[64648]"
//line /usr/local/go/src/regexp/exec.go:271
		_go_fuzz_dep_.CoverTab[64649]++
							i := t.inst
							add := false
							switch i.Op {
		default:
//line /usr/local/go/src/regexp/exec.go:275
			_go_fuzz_dep_.CoverTab[64659]++
								panic("bad inst")
//line /usr/local/go/src/regexp/exec.go:276
			// _ = "end of CoverTab[64659]"

		case syntax.InstMatch:
//line /usr/local/go/src/regexp/exec.go:278
			_go_fuzz_dep_.CoverTab[64660]++
								if len(t.cap) > 0 && func() bool {
//line /usr/local/go/src/regexp/exec.go:279
				_go_fuzz_dep_.CoverTab[64667]++
//line /usr/local/go/src/regexp/exec.go:279
				return (!longest || func() bool {
//line /usr/local/go/src/regexp/exec.go:279
					_go_fuzz_dep_.CoverTab[64668]++
//line /usr/local/go/src/regexp/exec.go:279
					return !m.matched
//line /usr/local/go/src/regexp/exec.go:279
					// _ = "end of CoverTab[64668]"
//line /usr/local/go/src/regexp/exec.go:279
				}() || func() bool {
//line /usr/local/go/src/regexp/exec.go:279
					_go_fuzz_dep_.CoverTab[64669]++
//line /usr/local/go/src/regexp/exec.go:279
					return m.matchcap[1] < pos
//line /usr/local/go/src/regexp/exec.go:279
					// _ = "end of CoverTab[64669]"
//line /usr/local/go/src/regexp/exec.go:279
				}())
//line /usr/local/go/src/regexp/exec.go:279
				// _ = "end of CoverTab[64667]"
//line /usr/local/go/src/regexp/exec.go:279
			}() {
//line /usr/local/go/src/regexp/exec.go:279
				_go_fuzz_dep_.CoverTab[64670]++
									t.cap[1] = pos
									copy(m.matchcap, t.cap)
//line /usr/local/go/src/regexp/exec.go:281
				// _ = "end of CoverTab[64670]"
			} else {
//line /usr/local/go/src/regexp/exec.go:282
				_go_fuzz_dep_.CoverTab[64671]++
//line /usr/local/go/src/regexp/exec.go:282
				// _ = "end of CoverTab[64671]"
//line /usr/local/go/src/regexp/exec.go:282
			}
//line /usr/local/go/src/regexp/exec.go:282
			// _ = "end of CoverTab[64660]"
//line /usr/local/go/src/regexp/exec.go:282
			_go_fuzz_dep_.CoverTab[64661]++
								if !longest {
//line /usr/local/go/src/regexp/exec.go:283
				_go_fuzz_dep_.CoverTab[64672]++

									for _, d := range runq.dense[j+1:] {
//line /usr/local/go/src/regexp/exec.go:285
					_go_fuzz_dep_.CoverTab[64674]++
										if d.t != nil {
//line /usr/local/go/src/regexp/exec.go:286
						_go_fuzz_dep_.CoverTab[64675]++
											m.pool = append(m.pool, d.t)
//line /usr/local/go/src/regexp/exec.go:287
						// _ = "end of CoverTab[64675]"
					} else {
//line /usr/local/go/src/regexp/exec.go:288
						_go_fuzz_dep_.CoverTab[64676]++
//line /usr/local/go/src/regexp/exec.go:288
						// _ = "end of CoverTab[64676]"
//line /usr/local/go/src/regexp/exec.go:288
					}
//line /usr/local/go/src/regexp/exec.go:288
					// _ = "end of CoverTab[64674]"
				}
//line /usr/local/go/src/regexp/exec.go:289
				// _ = "end of CoverTab[64672]"
//line /usr/local/go/src/regexp/exec.go:289
				_go_fuzz_dep_.CoverTab[64673]++
									runq.dense = runq.dense[:0]
//line /usr/local/go/src/regexp/exec.go:290
				// _ = "end of CoverTab[64673]"
			} else {
//line /usr/local/go/src/regexp/exec.go:291
				_go_fuzz_dep_.CoverTab[64677]++
//line /usr/local/go/src/regexp/exec.go:291
				// _ = "end of CoverTab[64677]"
//line /usr/local/go/src/regexp/exec.go:291
			}
//line /usr/local/go/src/regexp/exec.go:291
			// _ = "end of CoverTab[64661]"
//line /usr/local/go/src/regexp/exec.go:291
			_go_fuzz_dep_.CoverTab[64662]++
								m.matched = true
//line /usr/local/go/src/regexp/exec.go:292
			// _ = "end of CoverTab[64662]"

		case syntax.InstRune:
//line /usr/local/go/src/regexp/exec.go:294
			_go_fuzz_dep_.CoverTab[64663]++
								add = i.MatchRune(c)
//line /usr/local/go/src/regexp/exec.go:295
			// _ = "end of CoverTab[64663]"
		case syntax.InstRune1:
//line /usr/local/go/src/regexp/exec.go:296
			_go_fuzz_dep_.CoverTab[64664]++
								add = c == i.Rune[0]
//line /usr/local/go/src/regexp/exec.go:297
			// _ = "end of CoverTab[64664]"
		case syntax.InstRuneAny:
//line /usr/local/go/src/regexp/exec.go:298
			_go_fuzz_dep_.CoverTab[64665]++
								add = true
//line /usr/local/go/src/regexp/exec.go:299
			// _ = "end of CoverTab[64665]"
		case syntax.InstRuneAnyNotNL:
//line /usr/local/go/src/regexp/exec.go:300
			_go_fuzz_dep_.CoverTab[64666]++
								add = c != '\n'
//line /usr/local/go/src/regexp/exec.go:301
			// _ = "end of CoverTab[64666]"
		}
//line /usr/local/go/src/regexp/exec.go:302
		// _ = "end of CoverTab[64649]"
//line /usr/local/go/src/regexp/exec.go:302
		_go_fuzz_dep_.CoverTab[64650]++
							if add {
//line /usr/local/go/src/regexp/exec.go:303
			_go_fuzz_dep_.CoverTab[64678]++
								t = m.add(nextq, i.Out, nextPos, t.cap, nextCond, t)
//line /usr/local/go/src/regexp/exec.go:304
			// _ = "end of CoverTab[64678]"
		} else {
//line /usr/local/go/src/regexp/exec.go:305
			_go_fuzz_dep_.CoverTab[64679]++
//line /usr/local/go/src/regexp/exec.go:305
			// _ = "end of CoverTab[64679]"
//line /usr/local/go/src/regexp/exec.go:305
		}
//line /usr/local/go/src/regexp/exec.go:305
		// _ = "end of CoverTab[64650]"
//line /usr/local/go/src/regexp/exec.go:305
		_go_fuzz_dep_.CoverTab[64651]++
							if t != nil {
//line /usr/local/go/src/regexp/exec.go:306
			_go_fuzz_dep_.CoverTab[64680]++
								m.pool = append(m.pool, t)
//line /usr/local/go/src/regexp/exec.go:307
			// _ = "end of CoverTab[64680]"
		} else {
//line /usr/local/go/src/regexp/exec.go:308
			_go_fuzz_dep_.CoverTab[64681]++
//line /usr/local/go/src/regexp/exec.go:308
			// _ = "end of CoverTab[64681]"
//line /usr/local/go/src/regexp/exec.go:308
		}
//line /usr/local/go/src/regexp/exec.go:308
		// _ = "end of CoverTab[64651]"
	}
//line /usr/local/go/src/regexp/exec.go:309
	// _ = "end of CoverTab[64645]"
//line /usr/local/go/src/regexp/exec.go:309
	_go_fuzz_dep_.CoverTab[64646]++
						runq.dense = runq.dense[:0]
//line /usr/local/go/src/regexp/exec.go:310
	// _ = "end of CoverTab[64646]"
}

// add adds an entry to q for pc, unless the q already has such an entry.
//line /usr/local/go/src/regexp/exec.go:313
// It also recursively adds an entry for all instructions reachable from pc by following
//line /usr/local/go/src/regexp/exec.go:313
// empty-width conditions satisfied by cond.  pos gives the current position
//line /usr/local/go/src/regexp/exec.go:313
// in the input.
//line /usr/local/go/src/regexp/exec.go:317
func (m *machine) add(q *queue, pc uint32, pos int, cap []int, cond *lazyFlag, t *thread) *thread {
//line /usr/local/go/src/regexp/exec.go:317
	_go_fuzz_dep_.CoverTab[64682]++
Again:
	if pc == 0 {
//line /usr/local/go/src/regexp/exec.go:319
		_go_fuzz_dep_.CoverTab[64686]++
							return t
//line /usr/local/go/src/regexp/exec.go:320
		// _ = "end of CoverTab[64686]"
	} else {
//line /usr/local/go/src/regexp/exec.go:321
		_go_fuzz_dep_.CoverTab[64687]++
//line /usr/local/go/src/regexp/exec.go:321
		// _ = "end of CoverTab[64687]"
//line /usr/local/go/src/regexp/exec.go:321
	}
//line /usr/local/go/src/regexp/exec.go:321
	// _ = "end of CoverTab[64682]"
//line /usr/local/go/src/regexp/exec.go:321
	_go_fuzz_dep_.CoverTab[64683]++
						if j := q.sparse[pc]; j < uint32(len(q.dense)) && func() bool {
//line /usr/local/go/src/regexp/exec.go:322
		_go_fuzz_dep_.CoverTab[64688]++
//line /usr/local/go/src/regexp/exec.go:322
		return q.dense[j].pc == pc
//line /usr/local/go/src/regexp/exec.go:322
		// _ = "end of CoverTab[64688]"
//line /usr/local/go/src/regexp/exec.go:322
	}() {
//line /usr/local/go/src/regexp/exec.go:322
		_go_fuzz_dep_.CoverTab[64689]++
							return t
//line /usr/local/go/src/regexp/exec.go:323
		// _ = "end of CoverTab[64689]"
	} else {
//line /usr/local/go/src/regexp/exec.go:324
		_go_fuzz_dep_.CoverTab[64690]++
//line /usr/local/go/src/regexp/exec.go:324
		// _ = "end of CoverTab[64690]"
//line /usr/local/go/src/regexp/exec.go:324
	}
//line /usr/local/go/src/regexp/exec.go:324
	// _ = "end of CoverTab[64683]"
//line /usr/local/go/src/regexp/exec.go:324
	_go_fuzz_dep_.CoverTab[64684]++

						j := len(q.dense)
						q.dense = q.dense[:j+1]
						d := &q.dense[j]
						d.t = nil
						d.pc = pc
						q.sparse[pc] = uint32(j)

						i := &m.p.Inst[pc]
						switch i.Op {
	default:
//line /usr/local/go/src/regexp/exec.go:335
		_go_fuzz_dep_.CoverTab[64691]++
							panic("unhandled")
//line /usr/local/go/src/regexp/exec.go:336
		// _ = "end of CoverTab[64691]"
	case syntax.InstFail:
//line /usr/local/go/src/regexp/exec.go:337
		_go_fuzz_dep_.CoverTab[64692]++
//line /usr/local/go/src/regexp/exec.go:337
		// _ = "end of CoverTab[64692]"

	case syntax.InstAlt, syntax.InstAltMatch:
//line /usr/local/go/src/regexp/exec.go:339
		_go_fuzz_dep_.CoverTab[64693]++
							t = m.add(q, i.Out, pos, cap, cond, t)
							pc = i.Arg
							goto Again
//line /usr/local/go/src/regexp/exec.go:342
		// _ = "end of CoverTab[64693]"
	case syntax.InstEmptyWidth:
//line /usr/local/go/src/regexp/exec.go:343
		_go_fuzz_dep_.CoverTab[64694]++
							if cond.match(syntax.EmptyOp(i.Arg)) {
//line /usr/local/go/src/regexp/exec.go:344
			_go_fuzz_dep_.CoverTab[64700]++
								pc = i.Out
								goto Again
//line /usr/local/go/src/regexp/exec.go:346
			// _ = "end of CoverTab[64700]"
		} else {
//line /usr/local/go/src/regexp/exec.go:347
			_go_fuzz_dep_.CoverTab[64701]++
//line /usr/local/go/src/regexp/exec.go:347
			// _ = "end of CoverTab[64701]"
//line /usr/local/go/src/regexp/exec.go:347
		}
//line /usr/local/go/src/regexp/exec.go:347
		// _ = "end of CoverTab[64694]"
	case syntax.InstNop:
//line /usr/local/go/src/regexp/exec.go:348
		_go_fuzz_dep_.CoverTab[64695]++
							pc = i.Out
							goto Again
//line /usr/local/go/src/regexp/exec.go:350
		// _ = "end of CoverTab[64695]"
	case syntax.InstCapture:
//line /usr/local/go/src/regexp/exec.go:351
		_go_fuzz_dep_.CoverTab[64696]++
							if int(i.Arg) < len(cap) {
//line /usr/local/go/src/regexp/exec.go:352
			_go_fuzz_dep_.CoverTab[64702]++
								opos := cap[i.Arg]
								cap[i.Arg] = pos
								m.add(q, i.Out, pos, cap, cond, nil)
								cap[i.Arg] = opos
//line /usr/local/go/src/regexp/exec.go:356
			// _ = "end of CoverTab[64702]"
		} else {
//line /usr/local/go/src/regexp/exec.go:357
			_go_fuzz_dep_.CoverTab[64703]++
								pc = i.Out
								goto Again
//line /usr/local/go/src/regexp/exec.go:359
			// _ = "end of CoverTab[64703]"
		}
//line /usr/local/go/src/regexp/exec.go:360
		// _ = "end of CoverTab[64696]"
	case syntax.InstMatch, syntax.InstRune, syntax.InstRune1, syntax.InstRuneAny, syntax.InstRuneAnyNotNL:
//line /usr/local/go/src/regexp/exec.go:361
		_go_fuzz_dep_.CoverTab[64697]++
							if t == nil {
//line /usr/local/go/src/regexp/exec.go:362
			_go_fuzz_dep_.CoverTab[64704]++
								t = m.alloc(i)
//line /usr/local/go/src/regexp/exec.go:363
			// _ = "end of CoverTab[64704]"
		} else {
//line /usr/local/go/src/regexp/exec.go:364
			_go_fuzz_dep_.CoverTab[64705]++
								t.inst = i
//line /usr/local/go/src/regexp/exec.go:365
			// _ = "end of CoverTab[64705]"
		}
//line /usr/local/go/src/regexp/exec.go:366
		// _ = "end of CoverTab[64697]"
//line /usr/local/go/src/regexp/exec.go:366
		_go_fuzz_dep_.CoverTab[64698]++
							if len(cap) > 0 && func() bool {
//line /usr/local/go/src/regexp/exec.go:367
			_go_fuzz_dep_.CoverTab[64706]++
//line /usr/local/go/src/regexp/exec.go:367
			return &t.cap[0] != &cap[0]
//line /usr/local/go/src/regexp/exec.go:367
			// _ = "end of CoverTab[64706]"
//line /usr/local/go/src/regexp/exec.go:367
		}() {
//line /usr/local/go/src/regexp/exec.go:367
			_go_fuzz_dep_.CoverTab[64707]++
								copy(t.cap, cap)
//line /usr/local/go/src/regexp/exec.go:368
			// _ = "end of CoverTab[64707]"
		} else {
//line /usr/local/go/src/regexp/exec.go:369
			_go_fuzz_dep_.CoverTab[64708]++
//line /usr/local/go/src/regexp/exec.go:369
			// _ = "end of CoverTab[64708]"
//line /usr/local/go/src/regexp/exec.go:369
		}
//line /usr/local/go/src/regexp/exec.go:369
		// _ = "end of CoverTab[64698]"
//line /usr/local/go/src/regexp/exec.go:369
		_go_fuzz_dep_.CoverTab[64699]++
							d.t = t
							t = nil
//line /usr/local/go/src/regexp/exec.go:371
		// _ = "end of CoverTab[64699]"
	}
//line /usr/local/go/src/regexp/exec.go:372
	// _ = "end of CoverTab[64684]"
//line /usr/local/go/src/regexp/exec.go:372
	_go_fuzz_dep_.CoverTab[64685]++
						return t
//line /usr/local/go/src/regexp/exec.go:373
	// _ = "end of CoverTab[64685]"
}

type onePassMachine struct {
	inputs		inputs
	matchcap	[]int
}

var onePassPool sync.Pool

func newOnePassMachine() *onePassMachine {
//line /usr/local/go/src/regexp/exec.go:383
	_go_fuzz_dep_.CoverTab[64709]++
						m, ok := onePassPool.Get().(*onePassMachine)
						if !ok {
//line /usr/local/go/src/regexp/exec.go:385
		_go_fuzz_dep_.CoverTab[64711]++
							m = new(onePassMachine)
//line /usr/local/go/src/regexp/exec.go:386
		// _ = "end of CoverTab[64711]"
	} else {
//line /usr/local/go/src/regexp/exec.go:387
		_go_fuzz_dep_.CoverTab[64712]++
//line /usr/local/go/src/regexp/exec.go:387
		// _ = "end of CoverTab[64712]"
//line /usr/local/go/src/regexp/exec.go:387
	}
//line /usr/local/go/src/regexp/exec.go:387
	// _ = "end of CoverTab[64709]"
//line /usr/local/go/src/regexp/exec.go:387
	_go_fuzz_dep_.CoverTab[64710]++
						return m
//line /usr/local/go/src/regexp/exec.go:388
	// _ = "end of CoverTab[64710]"
}

func freeOnePassMachine(m *onePassMachine) {
//line /usr/local/go/src/regexp/exec.go:391
	_go_fuzz_dep_.CoverTab[64713]++
						m.inputs.clear()
						onePassPool.Put(m)
//line /usr/local/go/src/regexp/exec.go:393
	// _ = "end of CoverTab[64713]"
}

// doOnePass implements r.doExecute using the one-pass execution engine.
func (re *Regexp) doOnePass(ir io.RuneReader, ib []byte, is string, pos, ncap int, dstCap []int) []int {
//line /usr/local/go/src/regexp/exec.go:397
	_go_fuzz_dep_.CoverTab[64714]++
						startCond := re.cond
						if startCond == ^syntax.EmptyOp(0) {
//line /usr/local/go/src/regexp/exec.go:399
		_go_fuzz_dep_.CoverTab[64723]++
							return nil
//line /usr/local/go/src/regexp/exec.go:400
		// _ = "end of CoverTab[64723]"
	} else {
//line /usr/local/go/src/regexp/exec.go:401
		_go_fuzz_dep_.CoverTab[64724]++
//line /usr/local/go/src/regexp/exec.go:401
		// _ = "end of CoverTab[64724]"
//line /usr/local/go/src/regexp/exec.go:401
	}
//line /usr/local/go/src/regexp/exec.go:401
	// _ = "end of CoverTab[64714]"
//line /usr/local/go/src/regexp/exec.go:401
	_go_fuzz_dep_.CoverTab[64715]++

						m := newOnePassMachine()
						if cap(m.matchcap) < ncap {
//line /usr/local/go/src/regexp/exec.go:404
		_go_fuzz_dep_.CoverTab[64725]++
							m.matchcap = make([]int, ncap)
//line /usr/local/go/src/regexp/exec.go:405
		// _ = "end of CoverTab[64725]"
	} else {
//line /usr/local/go/src/regexp/exec.go:406
		_go_fuzz_dep_.CoverTab[64726]++
							m.matchcap = m.matchcap[:ncap]
//line /usr/local/go/src/regexp/exec.go:407
		// _ = "end of CoverTab[64726]"
	}
//line /usr/local/go/src/regexp/exec.go:408
	// _ = "end of CoverTab[64715]"
//line /usr/local/go/src/regexp/exec.go:408
	_go_fuzz_dep_.CoverTab[64716]++

						matched := false
						for i := range m.matchcap {
//line /usr/local/go/src/regexp/exec.go:411
		_go_fuzz_dep_.CoverTab[64727]++
							m.matchcap[i] = -1
//line /usr/local/go/src/regexp/exec.go:412
		// _ = "end of CoverTab[64727]"
	}
//line /usr/local/go/src/regexp/exec.go:413
	// _ = "end of CoverTab[64716]"
//line /usr/local/go/src/regexp/exec.go:413
	_go_fuzz_dep_.CoverTab[64717]++

						i, _ := m.inputs.init(ir, ib, is)

						r, r1 := endOfText, endOfText
						width, width1 := 0, 0
						r, width = i.step(pos)
						if r != endOfText {
//line /usr/local/go/src/regexp/exec.go:420
		_go_fuzz_dep_.CoverTab[64728]++
							r1, width1 = i.step(pos + width)
//line /usr/local/go/src/regexp/exec.go:421
		// _ = "end of CoverTab[64728]"
	} else {
//line /usr/local/go/src/regexp/exec.go:422
		_go_fuzz_dep_.CoverTab[64729]++
//line /usr/local/go/src/regexp/exec.go:422
		// _ = "end of CoverTab[64729]"
//line /usr/local/go/src/regexp/exec.go:422
	}
//line /usr/local/go/src/regexp/exec.go:422
	// _ = "end of CoverTab[64717]"
//line /usr/local/go/src/regexp/exec.go:422
	_go_fuzz_dep_.CoverTab[64718]++
						var flag lazyFlag
						if pos == 0 {
//line /usr/local/go/src/regexp/exec.go:424
		_go_fuzz_dep_.CoverTab[64730]++
							flag = newLazyFlag(-1, r)
//line /usr/local/go/src/regexp/exec.go:425
		// _ = "end of CoverTab[64730]"
	} else {
//line /usr/local/go/src/regexp/exec.go:426
		_go_fuzz_dep_.CoverTab[64731]++
							flag = i.context(pos)
//line /usr/local/go/src/regexp/exec.go:427
		// _ = "end of CoverTab[64731]"
	}
//line /usr/local/go/src/regexp/exec.go:428
	// _ = "end of CoverTab[64718]"
//line /usr/local/go/src/regexp/exec.go:428
	_go_fuzz_dep_.CoverTab[64719]++
						pc := re.onepass.Start
						inst := &re.onepass.Inst[pc]

						if pos == 0 && func() bool {
//line /usr/local/go/src/regexp/exec.go:432
		_go_fuzz_dep_.CoverTab[64732]++
//line /usr/local/go/src/regexp/exec.go:432
		return flag.match(syntax.EmptyOp(inst.Arg))
//line /usr/local/go/src/regexp/exec.go:432
		// _ = "end of CoverTab[64732]"
//line /usr/local/go/src/regexp/exec.go:432
	}() && func() bool {
//line /usr/local/go/src/regexp/exec.go:432
		_go_fuzz_dep_.CoverTab[64733]++
//line /usr/local/go/src/regexp/exec.go:432
		return len(re.prefix) > 0
							// _ = "end of CoverTab[64733]"
//line /usr/local/go/src/regexp/exec.go:433
	}() && func() bool {
//line /usr/local/go/src/regexp/exec.go:433
		_go_fuzz_dep_.CoverTab[64734]++
//line /usr/local/go/src/regexp/exec.go:433
		return i.canCheckPrefix()
//line /usr/local/go/src/regexp/exec.go:433
		// _ = "end of CoverTab[64734]"
//line /usr/local/go/src/regexp/exec.go:433
	}() {
//line /usr/local/go/src/regexp/exec.go:433
		_go_fuzz_dep_.CoverTab[64735]++

							if !i.hasPrefix(re) {
//line /usr/local/go/src/regexp/exec.go:435
			_go_fuzz_dep_.CoverTab[64737]++
								goto Return
//line /usr/local/go/src/regexp/exec.go:436
			// _ = "end of CoverTab[64737]"
		} else {
//line /usr/local/go/src/regexp/exec.go:437
			_go_fuzz_dep_.CoverTab[64738]++
//line /usr/local/go/src/regexp/exec.go:437
			// _ = "end of CoverTab[64738]"
//line /usr/local/go/src/regexp/exec.go:437
		}
//line /usr/local/go/src/regexp/exec.go:437
		// _ = "end of CoverTab[64735]"
//line /usr/local/go/src/regexp/exec.go:437
		_go_fuzz_dep_.CoverTab[64736]++
							pos += len(re.prefix)
							r, width = i.step(pos)
							r1, width1 = i.step(pos + width)
							flag = i.context(pos)
							pc = int(re.prefixEnd)
//line /usr/local/go/src/regexp/exec.go:442
		// _ = "end of CoverTab[64736]"
	} else {
//line /usr/local/go/src/regexp/exec.go:443
		_go_fuzz_dep_.CoverTab[64739]++
//line /usr/local/go/src/regexp/exec.go:443
		// _ = "end of CoverTab[64739]"
//line /usr/local/go/src/regexp/exec.go:443
	}
//line /usr/local/go/src/regexp/exec.go:443
	// _ = "end of CoverTab[64719]"
//line /usr/local/go/src/regexp/exec.go:443
	_go_fuzz_dep_.CoverTab[64720]++
						for {
//line /usr/local/go/src/regexp/exec.go:444
		_go_fuzz_dep_.CoverTab[64740]++
							inst = &re.onepass.Inst[pc]
							pc = int(inst.Out)
							switch inst.Op {
		default:
//line /usr/local/go/src/regexp/exec.go:448
			_go_fuzz_dep_.CoverTab[64743]++
								panic("bad inst")
//line /usr/local/go/src/regexp/exec.go:449
			// _ = "end of CoverTab[64743]"
		case syntax.InstMatch:
//line /usr/local/go/src/regexp/exec.go:450
			_go_fuzz_dep_.CoverTab[64744]++
								matched = true
								if len(m.matchcap) > 0 {
//line /usr/local/go/src/regexp/exec.go:452
				_go_fuzz_dep_.CoverTab[64757]++
									m.matchcap[0] = 0
									m.matchcap[1] = pos
//line /usr/local/go/src/regexp/exec.go:454
				// _ = "end of CoverTab[64757]"
			} else {
//line /usr/local/go/src/regexp/exec.go:455
				_go_fuzz_dep_.CoverTab[64758]++
//line /usr/local/go/src/regexp/exec.go:455
				// _ = "end of CoverTab[64758]"
//line /usr/local/go/src/regexp/exec.go:455
			}
//line /usr/local/go/src/regexp/exec.go:455
			// _ = "end of CoverTab[64744]"
//line /usr/local/go/src/regexp/exec.go:455
			_go_fuzz_dep_.CoverTab[64745]++
								goto Return
//line /usr/local/go/src/regexp/exec.go:456
			// _ = "end of CoverTab[64745]"
		case syntax.InstRune:
//line /usr/local/go/src/regexp/exec.go:457
			_go_fuzz_dep_.CoverTab[64746]++
								if !inst.MatchRune(r) {
//line /usr/local/go/src/regexp/exec.go:458
				_go_fuzz_dep_.CoverTab[64759]++
									goto Return
//line /usr/local/go/src/regexp/exec.go:459
				// _ = "end of CoverTab[64759]"
			} else {
//line /usr/local/go/src/regexp/exec.go:460
				_go_fuzz_dep_.CoverTab[64760]++
//line /usr/local/go/src/regexp/exec.go:460
				// _ = "end of CoverTab[64760]"
//line /usr/local/go/src/regexp/exec.go:460
			}
//line /usr/local/go/src/regexp/exec.go:460
			// _ = "end of CoverTab[64746]"
		case syntax.InstRune1:
//line /usr/local/go/src/regexp/exec.go:461
			_go_fuzz_dep_.CoverTab[64747]++
								if r != inst.Rune[0] {
//line /usr/local/go/src/regexp/exec.go:462
				_go_fuzz_dep_.CoverTab[64761]++
									goto Return
//line /usr/local/go/src/regexp/exec.go:463
				// _ = "end of CoverTab[64761]"
			} else {
//line /usr/local/go/src/regexp/exec.go:464
				_go_fuzz_dep_.CoverTab[64762]++
//line /usr/local/go/src/regexp/exec.go:464
				// _ = "end of CoverTab[64762]"
//line /usr/local/go/src/regexp/exec.go:464
			}
//line /usr/local/go/src/regexp/exec.go:464
			// _ = "end of CoverTab[64747]"
		case syntax.InstRuneAny:
//line /usr/local/go/src/regexp/exec.go:465
			_go_fuzz_dep_.CoverTab[64748]++
//line /usr/local/go/src/regexp/exec.go:465
			// _ = "end of CoverTab[64748]"

		case syntax.InstRuneAnyNotNL:
//line /usr/local/go/src/regexp/exec.go:467
			_go_fuzz_dep_.CoverTab[64749]++
								if r == '\n' {
//line /usr/local/go/src/regexp/exec.go:468
				_go_fuzz_dep_.CoverTab[64763]++
									goto Return
//line /usr/local/go/src/regexp/exec.go:469
				// _ = "end of CoverTab[64763]"
			} else {
//line /usr/local/go/src/regexp/exec.go:470
				_go_fuzz_dep_.CoverTab[64764]++
//line /usr/local/go/src/regexp/exec.go:470
				// _ = "end of CoverTab[64764]"
//line /usr/local/go/src/regexp/exec.go:470
			}
//line /usr/local/go/src/regexp/exec.go:470
			// _ = "end of CoverTab[64749]"

		case syntax.InstAlt, syntax.InstAltMatch:
//line /usr/local/go/src/regexp/exec.go:472
			_go_fuzz_dep_.CoverTab[64750]++
								pc = int(onePassNext(inst, r))
								continue
//line /usr/local/go/src/regexp/exec.go:474
			// _ = "end of CoverTab[64750]"
		case syntax.InstFail:
//line /usr/local/go/src/regexp/exec.go:475
			_go_fuzz_dep_.CoverTab[64751]++
								goto Return
//line /usr/local/go/src/regexp/exec.go:476
			// _ = "end of CoverTab[64751]"
		case syntax.InstNop:
//line /usr/local/go/src/regexp/exec.go:477
			_go_fuzz_dep_.CoverTab[64752]++
								continue
//line /usr/local/go/src/regexp/exec.go:478
			// _ = "end of CoverTab[64752]"
		case syntax.InstEmptyWidth:
//line /usr/local/go/src/regexp/exec.go:479
			_go_fuzz_dep_.CoverTab[64753]++
								if !flag.match(syntax.EmptyOp(inst.Arg)) {
//line /usr/local/go/src/regexp/exec.go:480
				_go_fuzz_dep_.CoverTab[64765]++
									goto Return
//line /usr/local/go/src/regexp/exec.go:481
				// _ = "end of CoverTab[64765]"
			} else {
//line /usr/local/go/src/regexp/exec.go:482
				_go_fuzz_dep_.CoverTab[64766]++
//line /usr/local/go/src/regexp/exec.go:482
				// _ = "end of CoverTab[64766]"
//line /usr/local/go/src/regexp/exec.go:482
			}
//line /usr/local/go/src/regexp/exec.go:482
			// _ = "end of CoverTab[64753]"
//line /usr/local/go/src/regexp/exec.go:482
			_go_fuzz_dep_.CoverTab[64754]++
								continue
//line /usr/local/go/src/regexp/exec.go:483
			// _ = "end of CoverTab[64754]"
		case syntax.InstCapture:
//line /usr/local/go/src/regexp/exec.go:484
			_go_fuzz_dep_.CoverTab[64755]++
								if int(inst.Arg) < len(m.matchcap) {
//line /usr/local/go/src/regexp/exec.go:485
				_go_fuzz_dep_.CoverTab[64767]++
									m.matchcap[inst.Arg] = pos
//line /usr/local/go/src/regexp/exec.go:486
				// _ = "end of CoverTab[64767]"
			} else {
//line /usr/local/go/src/regexp/exec.go:487
				_go_fuzz_dep_.CoverTab[64768]++
//line /usr/local/go/src/regexp/exec.go:487
				// _ = "end of CoverTab[64768]"
//line /usr/local/go/src/regexp/exec.go:487
			}
//line /usr/local/go/src/regexp/exec.go:487
			// _ = "end of CoverTab[64755]"
//line /usr/local/go/src/regexp/exec.go:487
			_go_fuzz_dep_.CoverTab[64756]++
								continue
//line /usr/local/go/src/regexp/exec.go:488
			// _ = "end of CoverTab[64756]"
		}
//line /usr/local/go/src/regexp/exec.go:489
		// _ = "end of CoverTab[64740]"
//line /usr/local/go/src/regexp/exec.go:489
		_go_fuzz_dep_.CoverTab[64741]++
							if width == 0 {
//line /usr/local/go/src/regexp/exec.go:490
			_go_fuzz_dep_.CoverTab[64769]++
								break
//line /usr/local/go/src/regexp/exec.go:491
			// _ = "end of CoverTab[64769]"
		} else {
//line /usr/local/go/src/regexp/exec.go:492
			_go_fuzz_dep_.CoverTab[64770]++
//line /usr/local/go/src/regexp/exec.go:492
			// _ = "end of CoverTab[64770]"
//line /usr/local/go/src/regexp/exec.go:492
		}
//line /usr/local/go/src/regexp/exec.go:492
		// _ = "end of CoverTab[64741]"
//line /usr/local/go/src/regexp/exec.go:492
		_go_fuzz_dep_.CoverTab[64742]++
							flag = newLazyFlag(r, r1)
							pos += width
							r, width = r1, width1
							if r != endOfText {
//line /usr/local/go/src/regexp/exec.go:496
			_go_fuzz_dep_.CoverTab[64771]++
								r1, width1 = i.step(pos + width)
//line /usr/local/go/src/regexp/exec.go:497
			// _ = "end of CoverTab[64771]"
		} else {
//line /usr/local/go/src/regexp/exec.go:498
			_go_fuzz_dep_.CoverTab[64772]++
//line /usr/local/go/src/regexp/exec.go:498
			// _ = "end of CoverTab[64772]"
//line /usr/local/go/src/regexp/exec.go:498
		}
//line /usr/local/go/src/regexp/exec.go:498
		// _ = "end of CoverTab[64742]"
	}
//line /usr/local/go/src/regexp/exec.go:499
	// _ = "end of CoverTab[64720]"
//line /usr/local/go/src/regexp/exec.go:499
	_go_fuzz_dep_.CoverTab[64721]++

Return:
	if !matched {
//line /usr/local/go/src/regexp/exec.go:502
		_go_fuzz_dep_.CoverTab[64773]++
							freeOnePassMachine(m)
							return nil
//line /usr/local/go/src/regexp/exec.go:504
		// _ = "end of CoverTab[64773]"
	} else {
//line /usr/local/go/src/regexp/exec.go:505
		_go_fuzz_dep_.CoverTab[64774]++
//line /usr/local/go/src/regexp/exec.go:505
		// _ = "end of CoverTab[64774]"
//line /usr/local/go/src/regexp/exec.go:505
	}
//line /usr/local/go/src/regexp/exec.go:505
	// _ = "end of CoverTab[64721]"
//line /usr/local/go/src/regexp/exec.go:505
	_go_fuzz_dep_.CoverTab[64722]++

						dstCap = append(dstCap, m.matchcap...)
						freeOnePassMachine(m)
						return dstCap
//line /usr/local/go/src/regexp/exec.go:509
	// _ = "end of CoverTab[64722]"
}

// doMatch reports whether either r, b or s match the regexp.
func (re *Regexp) doMatch(r io.RuneReader, b []byte, s string) bool {
//line /usr/local/go/src/regexp/exec.go:513
	_go_fuzz_dep_.CoverTab[64775]++
						return re.doExecute(r, b, s, 0, 0, nil) != nil
//line /usr/local/go/src/regexp/exec.go:514
	// _ = "end of CoverTab[64775]"
}

// doExecute finds the leftmost match in the input, appends the position
//line /usr/local/go/src/regexp/exec.go:517
// of its subexpressions to dstCap and returns dstCap.
//line /usr/local/go/src/regexp/exec.go:517
//
//line /usr/local/go/src/regexp/exec.go:517
// nil is returned if no matches are found and non-nil if matches are found.
//line /usr/local/go/src/regexp/exec.go:521
func (re *Regexp) doExecute(r io.RuneReader, b []byte, s string, pos int, ncap int, dstCap []int) []int {
//line /usr/local/go/src/regexp/exec.go:521
	_go_fuzz_dep_.CoverTab[64776]++
						if dstCap == nil {
//line /usr/local/go/src/regexp/exec.go:522
		_go_fuzz_dep_.CoverTab[64782]++

							dstCap = arrayNoInts[:0:0]
//line /usr/local/go/src/regexp/exec.go:524
		// _ = "end of CoverTab[64782]"
	} else {
//line /usr/local/go/src/regexp/exec.go:525
		_go_fuzz_dep_.CoverTab[64783]++
//line /usr/local/go/src/regexp/exec.go:525
		// _ = "end of CoverTab[64783]"
//line /usr/local/go/src/regexp/exec.go:525
	}
//line /usr/local/go/src/regexp/exec.go:525
	// _ = "end of CoverTab[64776]"
//line /usr/local/go/src/regexp/exec.go:525
	_go_fuzz_dep_.CoverTab[64777]++

						if r == nil && func() bool {
//line /usr/local/go/src/regexp/exec.go:527
		_go_fuzz_dep_.CoverTab[64784]++
//line /usr/local/go/src/regexp/exec.go:527
		return len(b)+len(s) < re.minInputLen
//line /usr/local/go/src/regexp/exec.go:527
		// _ = "end of CoverTab[64784]"
//line /usr/local/go/src/regexp/exec.go:527
	}() {
//line /usr/local/go/src/regexp/exec.go:527
		_go_fuzz_dep_.CoverTab[64785]++
							return nil
//line /usr/local/go/src/regexp/exec.go:528
		// _ = "end of CoverTab[64785]"
	} else {
//line /usr/local/go/src/regexp/exec.go:529
		_go_fuzz_dep_.CoverTab[64786]++
//line /usr/local/go/src/regexp/exec.go:529
		// _ = "end of CoverTab[64786]"
//line /usr/local/go/src/regexp/exec.go:529
	}
//line /usr/local/go/src/regexp/exec.go:529
	// _ = "end of CoverTab[64777]"
//line /usr/local/go/src/regexp/exec.go:529
	_go_fuzz_dep_.CoverTab[64778]++

						if re.onepass != nil {
//line /usr/local/go/src/regexp/exec.go:531
		_go_fuzz_dep_.CoverTab[64787]++
							return re.doOnePass(r, b, s, pos, ncap, dstCap)
//line /usr/local/go/src/regexp/exec.go:532
		// _ = "end of CoverTab[64787]"
	} else {
//line /usr/local/go/src/regexp/exec.go:533
		_go_fuzz_dep_.CoverTab[64788]++
//line /usr/local/go/src/regexp/exec.go:533
		// _ = "end of CoverTab[64788]"
//line /usr/local/go/src/regexp/exec.go:533
	}
//line /usr/local/go/src/regexp/exec.go:533
	// _ = "end of CoverTab[64778]"
//line /usr/local/go/src/regexp/exec.go:533
	_go_fuzz_dep_.CoverTab[64779]++
						if r == nil && func() bool {
//line /usr/local/go/src/regexp/exec.go:534
		_go_fuzz_dep_.CoverTab[64789]++
//line /usr/local/go/src/regexp/exec.go:534
		return len(b)+len(s) < re.maxBitStateLen
//line /usr/local/go/src/regexp/exec.go:534
		// _ = "end of CoverTab[64789]"
//line /usr/local/go/src/regexp/exec.go:534
	}() {
//line /usr/local/go/src/regexp/exec.go:534
		_go_fuzz_dep_.CoverTab[64790]++
							return re.backtrack(b, s, pos, ncap, dstCap)
//line /usr/local/go/src/regexp/exec.go:535
		// _ = "end of CoverTab[64790]"
	} else {
//line /usr/local/go/src/regexp/exec.go:536
		_go_fuzz_dep_.CoverTab[64791]++
//line /usr/local/go/src/regexp/exec.go:536
		// _ = "end of CoverTab[64791]"
//line /usr/local/go/src/regexp/exec.go:536
	}
//line /usr/local/go/src/regexp/exec.go:536
	// _ = "end of CoverTab[64779]"
//line /usr/local/go/src/regexp/exec.go:536
	_go_fuzz_dep_.CoverTab[64780]++

						m := re.get()
						i, _ := m.inputs.init(r, b, s)

						m.init(ncap)
						if !m.match(i, pos) {
//line /usr/local/go/src/regexp/exec.go:542
		_go_fuzz_dep_.CoverTab[64792]++
							re.put(m)
							return nil
//line /usr/local/go/src/regexp/exec.go:544
		// _ = "end of CoverTab[64792]"
	} else {
//line /usr/local/go/src/regexp/exec.go:545
		_go_fuzz_dep_.CoverTab[64793]++
//line /usr/local/go/src/regexp/exec.go:545
		// _ = "end of CoverTab[64793]"
//line /usr/local/go/src/regexp/exec.go:545
	}
//line /usr/local/go/src/regexp/exec.go:545
	// _ = "end of CoverTab[64780]"
//line /usr/local/go/src/regexp/exec.go:545
	_go_fuzz_dep_.CoverTab[64781]++

						dstCap = append(dstCap, m.matchcap...)
						re.put(m)
						return dstCap
//line /usr/local/go/src/regexp/exec.go:549
	// _ = "end of CoverTab[64781]"
}

// arrayNoInts is returned by doExecute match if nil dstCap is passed
//line /usr/local/go/src/regexp/exec.go:552
// to it with ncap=0.
//line /usr/local/go/src/regexp/exec.go:554
var arrayNoInts [0]int
//line /usr/local/go/src/regexp/exec.go:554
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/regexp/exec.go:554
var _ = _go_fuzz_dep_.CoverTab
