// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/regexp/regexp.go:5
// Package regexp implements regular expression search.
//line /usr/local/go/src/regexp/regexp.go:5
//
//line /usr/local/go/src/regexp/regexp.go:5
// The syntax of the regular expressions accepted is the same
//line /usr/local/go/src/regexp/regexp.go:5
// general syntax used by Perl, Python, and other languages.
//line /usr/local/go/src/regexp/regexp.go:5
// More precisely, it is the syntax accepted by RE2 and described at
//line /usr/local/go/src/regexp/regexp.go:5
// https://golang.org/s/re2syntax, except for \C.
//line /usr/local/go/src/regexp/regexp.go:5
// For an overview of the syntax, run
//line /usr/local/go/src/regexp/regexp.go:5
//
//line /usr/local/go/src/regexp/regexp.go:5
//	go doc regexp/syntax
//line /usr/local/go/src/regexp/regexp.go:5
//
//line /usr/local/go/src/regexp/regexp.go:5
// The regexp implementation provided by this package is
//line /usr/local/go/src/regexp/regexp.go:5
// guaranteed to run in time linear in the size of the input.
//line /usr/local/go/src/regexp/regexp.go:5
// (This is a property not guaranteed by most open source
//line /usr/local/go/src/regexp/regexp.go:5
// implementations of regular expressions.) For more information
//line /usr/local/go/src/regexp/regexp.go:5
// about this property, see
//line /usr/local/go/src/regexp/regexp.go:5
//
//line /usr/local/go/src/regexp/regexp.go:5
//	https://swtch.com/~rsc/regexp/regexp1.html
//line /usr/local/go/src/regexp/regexp.go:5
//
//line /usr/local/go/src/regexp/regexp.go:5
// or any book about automata theory.
//line /usr/local/go/src/regexp/regexp.go:5
//
//line /usr/local/go/src/regexp/regexp.go:5
// All characters are UTF-8-encoded code points.
//line /usr/local/go/src/regexp/regexp.go:5
// Following utf8.DecodeRune, each byte of an invalid UTF-8 sequence
//line /usr/local/go/src/regexp/regexp.go:5
// is treated as if it encoded utf8.RuneError (U+FFFD).
//line /usr/local/go/src/regexp/regexp.go:5
//
//line /usr/local/go/src/regexp/regexp.go:5
// There are 16 methods of Regexp that match a regular expression and identify
//line /usr/local/go/src/regexp/regexp.go:5
// the matched text. Their names are matched by this regular expression:
//line /usr/local/go/src/regexp/regexp.go:5
//
//line /usr/local/go/src/regexp/regexp.go:5
//	Find(All)?(String)?(Submatch)?(Index)?
//line /usr/local/go/src/regexp/regexp.go:5
//
//line /usr/local/go/src/regexp/regexp.go:5
// If 'All' is present, the routine matches successive non-overlapping
//line /usr/local/go/src/regexp/regexp.go:5
// matches of the entire expression. Empty matches abutting a preceding
//line /usr/local/go/src/regexp/regexp.go:5
// match are ignored. The return value is a slice containing the successive
//line /usr/local/go/src/regexp/regexp.go:5
// return values of the corresponding non-'All' routine. These routines take
//line /usr/local/go/src/regexp/regexp.go:5
// an extra integer argument, n. If n >= 0, the function returns at most n
//line /usr/local/go/src/regexp/regexp.go:5
// matches/submatches; otherwise, it returns all of them.
//line /usr/local/go/src/regexp/regexp.go:5
//
//line /usr/local/go/src/regexp/regexp.go:5
// If 'String' is present, the argument is a string; otherwise it is a slice
//line /usr/local/go/src/regexp/regexp.go:5
// of bytes; return values are adjusted as appropriate.
//line /usr/local/go/src/regexp/regexp.go:5
//
//line /usr/local/go/src/regexp/regexp.go:5
// If 'Submatch' is present, the return value is a slice identifying the
//line /usr/local/go/src/regexp/regexp.go:5
// successive submatches of the expression. Submatches are matches of
//line /usr/local/go/src/regexp/regexp.go:5
// parenthesized subexpressions (also known as capturing groups) within the
//line /usr/local/go/src/regexp/regexp.go:5
// regular expression, numbered from left to right in order of opening
//line /usr/local/go/src/regexp/regexp.go:5
// parenthesis. Submatch 0 is the match of the entire expression, submatch 1 is
//line /usr/local/go/src/regexp/regexp.go:5
// the match of the first parenthesized subexpression, and so on.
//line /usr/local/go/src/regexp/regexp.go:5
//
//line /usr/local/go/src/regexp/regexp.go:5
// If 'Index' is present, matches and submatches are identified by byte index
//line /usr/local/go/src/regexp/regexp.go:5
// pairs within the input string: result[2*n:2*n+2] identifies the indexes of
//line /usr/local/go/src/regexp/regexp.go:5
// the nth submatch. The pair for n==0 identifies the match of the entire
//line /usr/local/go/src/regexp/regexp.go:5
// expression. If 'Index' is not present, the match is identified by the text
//line /usr/local/go/src/regexp/regexp.go:5
// of the match/submatch. If an index is negative or text is nil, it means that
//line /usr/local/go/src/regexp/regexp.go:5
// subexpression did not match any string in the input. For 'String' versions
//line /usr/local/go/src/regexp/regexp.go:5
// an empty string means either no match or an empty match.
//line /usr/local/go/src/regexp/regexp.go:5
//
//line /usr/local/go/src/regexp/regexp.go:5
// There is also a subset of the methods that can be applied to text read
//line /usr/local/go/src/regexp/regexp.go:5
// from a RuneReader:
//line /usr/local/go/src/regexp/regexp.go:5
//
//line /usr/local/go/src/regexp/regexp.go:5
//	MatchReader, FindReaderIndex, FindReaderSubmatchIndex
//line /usr/local/go/src/regexp/regexp.go:5
//
//line /usr/local/go/src/regexp/regexp.go:5
// This set may grow. Note that regular expression matches may need to
//line /usr/local/go/src/regexp/regexp.go:5
// examine text beyond the text returned by a match, so the methods that
//line /usr/local/go/src/regexp/regexp.go:5
// match text from a RuneReader may read arbitrarily far into the input
//line /usr/local/go/src/regexp/regexp.go:5
// before returning.
//line /usr/local/go/src/regexp/regexp.go:5
//
//line /usr/local/go/src/regexp/regexp.go:5
// (There are a few other methods that do not match this pattern.)
//line /usr/local/go/src/regexp/regexp.go:70
package regexp

//line /usr/local/go/src/regexp/regexp.go:70
import (
//line /usr/local/go/src/regexp/regexp.go:70
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/regexp/regexp.go:70
)
//line /usr/local/go/src/regexp/regexp.go:70
import (
//line /usr/local/go/src/regexp/regexp.go:70
	_atomic_ "sync/atomic"
//line /usr/local/go/src/regexp/regexp.go:70
)

import (
	"bytes"
	"io"
	"regexp/syntax"
	"strconv"
	"strings"
	"sync"
	"unicode"
	"unicode/utf8"
)

// Regexp is the representation of a compiled regular expression.
//line /usr/local/go/src/regexp/regexp.go:83
// A Regexp is safe for concurrent use by multiple goroutines,
//line /usr/local/go/src/regexp/regexp.go:83
// except for configuration methods, such as Longest.
//line /usr/local/go/src/regexp/regexp.go:86
type Regexp struct {
	expr		string		// as passed to Compile
	prog		*syntax.Prog	// compiled program
	onepass		*onePassProg	// onepass program or nil
	numSubexp	int
	maxBitStateLen	int
	subexpNames	[]string
	prefix		string		// required prefix in unanchored matches
	prefixBytes	[]byte		// prefix, as a []byte
	prefixRune	rune		// first rune in prefix
	prefixEnd	uint32		// pc for last rune in prefix
	mpool		int		// pool for machines
	matchcap	int		// size of recorded match lengths
	prefixComplete	bool		// prefix is the entire regexp
	cond		syntax.EmptyOp	// empty-width conditions required at start of match
	minInputLen	int		// minimum length of the input in bytes

	// This field can be modified by the Longest method,
	// but it is otherwise read-only.
	longest	bool	// whether regexp prefers leftmost-longest match
}

// String returns the source text used to compile the regular expression.
func (re *Regexp) String() string {
//line /usr/local/go/src/regexp/regexp.go:109
	_go_fuzz_dep_.CoverTab[65004]++
						return re.expr
//line /usr/local/go/src/regexp/regexp.go:110
	// _ = "end of CoverTab[65004]"
}

// Copy returns a new Regexp object copied from re.
//line /usr/local/go/src/regexp/regexp.go:113
// Calling Longest on one copy does not affect another.
//line /usr/local/go/src/regexp/regexp.go:113
//
//line /usr/local/go/src/regexp/regexp.go:113
// Deprecated: In earlier releases, when using a Regexp in multiple goroutines,
//line /usr/local/go/src/regexp/regexp.go:113
// giving each goroutine its own copy helped to avoid lock contention.
//line /usr/local/go/src/regexp/regexp.go:113
// As of Go 1.12, using Copy is no longer necessary to avoid lock contention.
//line /usr/local/go/src/regexp/regexp.go:113
// Copy may still be appropriate if the reason for its use is to make
//line /usr/local/go/src/regexp/regexp.go:113
// two copies with different Longest settings.
//line /usr/local/go/src/regexp/regexp.go:121
func (re *Regexp) Copy() *Regexp {
//line /usr/local/go/src/regexp/regexp.go:121
	_go_fuzz_dep_.CoverTab[65005]++
						re2 := *re
						return &re2
//line /usr/local/go/src/regexp/regexp.go:123
	// _ = "end of CoverTab[65005]"
}

// Compile parses a regular expression and returns, if successful,
//line /usr/local/go/src/regexp/regexp.go:126
// a Regexp object that can be used to match against text.
//line /usr/local/go/src/regexp/regexp.go:126
//
//line /usr/local/go/src/regexp/regexp.go:126
// When matching against text, the regexp returns a match that
//line /usr/local/go/src/regexp/regexp.go:126
// begins as early as possible in the input (leftmost), and among those
//line /usr/local/go/src/regexp/regexp.go:126
// it chooses the one that a backtracking search would have found first.
//line /usr/local/go/src/regexp/regexp.go:126
// This so-called leftmost-first matching is the same semantics
//line /usr/local/go/src/regexp/regexp.go:126
// that Perl, Python, and other implementations use, although this
//line /usr/local/go/src/regexp/regexp.go:126
// package implements it without the expense of backtracking.
//line /usr/local/go/src/regexp/regexp.go:126
// For POSIX leftmost-longest matching, see CompilePOSIX.
//line /usr/local/go/src/regexp/regexp.go:136
func Compile(expr string) (*Regexp, error) {
//line /usr/local/go/src/regexp/regexp.go:136
	_go_fuzz_dep_.CoverTab[65006]++
						return compile(expr, syntax.Perl, false)
//line /usr/local/go/src/regexp/regexp.go:137
	// _ = "end of CoverTab[65006]"
}

// CompilePOSIX is like Compile but restricts the regular expression
//line /usr/local/go/src/regexp/regexp.go:140
// to POSIX ERE (egrep) syntax and changes the match semantics to
//line /usr/local/go/src/regexp/regexp.go:140
// leftmost-longest.
//line /usr/local/go/src/regexp/regexp.go:140
//
//line /usr/local/go/src/regexp/regexp.go:140
// That is, when matching against text, the regexp returns a match that
//line /usr/local/go/src/regexp/regexp.go:140
// begins as early as possible in the input (leftmost), and among those
//line /usr/local/go/src/regexp/regexp.go:140
// it chooses a match that is as long as possible.
//line /usr/local/go/src/regexp/regexp.go:140
// This so-called leftmost-longest matching is the same semantics
//line /usr/local/go/src/regexp/regexp.go:140
// that early regular expression implementations used and that POSIX
//line /usr/local/go/src/regexp/regexp.go:140
// specifies.
//line /usr/local/go/src/regexp/regexp.go:140
//
//line /usr/local/go/src/regexp/regexp.go:140
// However, there can be multiple leftmost-longest matches, with different
//line /usr/local/go/src/regexp/regexp.go:140
// submatch choices, and here this package diverges from POSIX.
//line /usr/local/go/src/regexp/regexp.go:140
// Among the possible leftmost-longest matches, this package chooses
//line /usr/local/go/src/regexp/regexp.go:140
// the one that a backtracking search would have found first, while POSIX
//line /usr/local/go/src/regexp/regexp.go:140
// specifies that the match be chosen to maximize the length of the first
//line /usr/local/go/src/regexp/regexp.go:140
// subexpression, then the second, and so on from left to right.
//line /usr/local/go/src/regexp/regexp.go:140
// The POSIX rule is computationally prohibitive and not even well-defined.
//line /usr/local/go/src/regexp/regexp.go:140
// See https://swtch.com/~rsc/regexp/regexp2.html#posix for details.
//line /usr/local/go/src/regexp/regexp.go:159
func CompilePOSIX(expr string) (*Regexp, error) {
//line /usr/local/go/src/regexp/regexp.go:159
	_go_fuzz_dep_.CoverTab[65007]++
						return compile(expr, syntax.POSIX, true)
//line /usr/local/go/src/regexp/regexp.go:160
	// _ = "end of CoverTab[65007]"
}

// Longest makes future searches prefer the leftmost-longest match.
//line /usr/local/go/src/regexp/regexp.go:163
// That is, when matching against text, the regexp returns a match that
//line /usr/local/go/src/regexp/regexp.go:163
// begins as early as possible in the input (leftmost), and among those
//line /usr/local/go/src/regexp/regexp.go:163
// it chooses a match that is as long as possible.
//line /usr/local/go/src/regexp/regexp.go:163
// This method modifies the Regexp and may not be called concurrently
//line /usr/local/go/src/regexp/regexp.go:163
// with any other methods.
//line /usr/local/go/src/regexp/regexp.go:169
func (re *Regexp) Longest() {
//line /usr/local/go/src/regexp/regexp.go:169
	_go_fuzz_dep_.CoverTab[65008]++
						re.longest = true
//line /usr/local/go/src/regexp/regexp.go:170
	// _ = "end of CoverTab[65008]"
}

func compile(expr string, mode syntax.Flags, longest bool) (*Regexp, error) {
//line /usr/local/go/src/regexp/regexp.go:173
	_go_fuzz_dep_.CoverTab[65009]++
						re, err := syntax.Parse(expr, mode)
						if err != nil {
//line /usr/local/go/src/regexp/regexp.go:175
		_go_fuzz_dep_.CoverTab[65016]++
							return nil, err
//line /usr/local/go/src/regexp/regexp.go:176
		// _ = "end of CoverTab[65016]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:177
		_go_fuzz_dep_.CoverTab[65017]++
//line /usr/local/go/src/regexp/regexp.go:177
		// _ = "end of CoverTab[65017]"
//line /usr/local/go/src/regexp/regexp.go:177
	}
//line /usr/local/go/src/regexp/regexp.go:177
	// _ = "end of CoverTab[65009]"
//line /usr/local/go/src/regexp/regexp.go:177
	_go_fuzz_dep_.CoverTab[65010]++
						maxCap := re.MaxCap()
						capNames := re.CapNames()

						re = re.Simplify()
						prog, err := syntax.Compile(re)
						if err != nil {
//line /usr/local/go/src/regexp/regexp.go:183
		_go_fuzz_dep_.CoverTab[65018]++
							return nil, err
//line /usr/local/go/src/regexp/regexp.go:184
		// _ = "end of CoverTab[65018]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:185
		_go_fuzz_dep_.CoverTab[65019]++
//line /usr/local/go/src/regexp/regexp.go:185
		// _ = "end of CoverTab[65019]"
//line /usr/local/go/src/regexp/regexp.go:185
	}
//line /usr/local/go/src/regexp/regexp.go:185
	// _ = "end of CoverTab[65010]"
//line /usr/local/go/src/regexp/regexp.go:185
	_go_fuzz_dep_.CoverTab[65011]++
						matchcap := prog.NumCap
						if matchcap < 2 {
//line /usr/local/go/src/regexp/regexp.go:187
		_go_fuzz_dep_.CoverTab[65020]++
							matchcap = 2
//line /usr/local/go/src/regexp/regexp.go:188
		// _ = "end of CoverTab[65020]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:189
		_go_fuzz_dep_.CoverTab[65021]++
//line /usr/local/go/src/regexp/regexp.go:189
		// _ = "end of CoverTab[65021]"
//line /usr/local/go/src/regexp/regexp.go:189
	}
//line /usr/local/go/src/regexp/regexp.go:189
	// _ = "end of CoverTab[65011]"
//line /usr/local/go/src/regexp/regexp.go:189
	_go_fuzz_dep_.CoverTab[65012]++
						regexp := &Regexp{
		expr:		expr,
		prog:		prog,
		onepass:	compileOnePass(prog),
		numSubexp:	maxCap,
		subexpNames:	capNames,
		cond:		prog.StartCond(),
		longest:	longest,
		matchcap:	matchcap,
		minInputLen:	minInputLen(re),
	}
	if regexp.onepass == nil {
//line /usr/local/go/src/regexp/regexp.go:201
		_go_fuzz_dep_.CoverTab[65022]++
							regexp.prefix, regexp.prefixComplete = prog.Prefix()
							regexp.maxBitStateLen = maxBitStateLen(prog)
//line /usr/local/go/src/regexp/regexp.go:203
		// _ = "end of CoverTab[65022]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:204
		_go_fuzz_dep_.CoverTab[65023]++
							regexp.prefix, regexp.prefixComplete, regexp.prefixEnd = onePassPrefix(prog)
//line /usr/local/go/src/regexp/regexp.go:205
		// _ = "end of CoverTab[65023]"
	}
//line /usr/local/go/src/regexp/regexp.go:206
	// _ = "end of CoverTab[65012]"
//line /usr/local/go/src/regexp/regexp.go:206
	_go_fuzz_dep_.CoverTab[65013]++
						if regexp.prefix != "" {
//line /usr/local/go/src/regexp/regexp.go:207
		_go_fuzz_dep_.CoverTab[65024]++

//line /usr/local/go/src/regexp/regexp.go:210
		regexp.prefixBytes = []byte(regexp.prefix)
							regexp.prefixRune, _ = utf8.DecodeRuneInString(regexp.prefix)
//line /usr/local/go/src/regexp/regexp.go:211
		// _ = "end of CoverTab[65024]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:212
		_go_fuzz_dep_.CoverTab[65025]++
//line /usr/local/go/src/regexp/regexp.go:212
		// _ = "end of CoverTab[65025]"
//line /usr/local/go/src/regexp/regexp.go:212
	}
//line /usr/local/go/src/regexp/regexp.go:212
	// _ = "end of CoverTab[65013]"
//line /usr/local/go/src/regexp/regexp.go:212
	_go_fuzz_dep_.CoverTab[65014]++

						n := len(prog.Inst)
						i := 0
						for matchSize[i] != 0 && func() bool {
//line /usr/local/go/src/regexp/regexp.go:216
		_go_fuzz_dep_.CoverTab[65026]++
//line /usr/local/go/src/regexp/regexp.go:216
		return matchSize[i] < n
//line /usr/local/go/src/regexp/regexp.go:216
		// _ = "end of CoverTab[65026]"
//line /usr/local/go/src/regexp/regexp.go:216
	}() {
//line /usr/local/go/src/regexp/regexp.go:216
		_go_fuzz_dep_.CoverTab[65027]++
							i++
//line /usr/local/go/src/regexp/regexp.go:217
		// _ = "end of CoverTab[65027]"
	}
//line /usr/local/go/src/regexp/regexp.go:218
	// _ = "end of CoverTab[65014]"
//line /usr/local/go/src/regexp/regexp.go:218
	_go_fuzz_dep_.CoverTab[65015]++
						regexp.mpool = i

						return regexp, nil
//line /usr/local/go/src/regexp/regexp.go:221
	// _ = "end of CoverTab[65015]"
}

// Pools of *machine for use during (*Regexp).doExecute,
//line /usr/local/go/src/regexp/regexp.go:224
// split up by the size of the execution queues.
//line /usr/local/go/src/regexp/regexp.go:224
// matchPool[i] machines have queue size matchSize[i].
//line /usr/local/go/src/regexp/regexp.go:224
// On a 64-bit system each queue entry is 16 bytes,
//line /usr/local/go/src/regexp/regexp.go:224
// so matchPool[0] has 16*2*128 = 4kB queues, etc.
//line /usr/local/go/src/regexp/regexp.go:224
// The final matchPool is a catch-all for very large queues.
//line /usr/local/go/src/regexp/regexp.go:230
var (
	matchSize	= [...]int{128, 512, 2048, 16384, 0}
	matchPool	[len(matchSize)]sync.Pool
)

// get returns a machine to use for matching re.
//line /usr/local/go/src/regexp/regexp.go:235
// It uses the re's machine cache if possible, to avoid
//line /usr/local/go/src/regexp/regexp.go:235
// unnecessary allocation.
//line /usr/local/go/src/regexp/regexp.go:238
func (re *Regexp) get() *machine {
//line /usr/local/go/src/regexp/regexp.go:238
	_go_fuzz_dep_.CoverTab[65028]++
						m, ok := matchPool[re.mpool].Get().(*machine)
						if !ok {
//line /usr/local/go/src/regexp/regexp.go:240
		_go_fuzz_dep_.CoverTab[65033]++
							m = new(machine)
//line /usr/local/go/src/regexp/regexp.go:241
		// _ = "end of CoverTab[65033]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:242
		_go_fuzz_dep_.CoverTab[65034]++
//line /usr/local/go/src/regexp/regexp.go:242
		// _ = "end of CoverTab[65034]"
//line /usr/local/go/src/regexp/regexp.go:242
	}
//line /usr/local/go/src/regexp/regexp.go:242
	// _ = "end of CoverTab[65028]"
//line /usr/local/go/src/regexp/regexp.go:242
	_go_fuzz_dep_.CoverTab[65029]++
						m.re = re
						m.p = re.prog
						if cap(m.matchcap) < re.matchcap {
//line /usr/local/go/src/regexp/regexp.go:245
		_go_fuzz_dep_.CoverTab[65035]++
							m.matchcap = make([]int, re.matchcap)
							for _, t := range m.pool {
//line /usr/local/go/src/regexp/regexp.go:247
			_go_fuzz_dep_.CoverTab[65036]++
								t.cap = make([]int, re.matchcap)
//line /usr/local/go/src/regexp/regexp.go:248
			// _ = "end of CoverTab[65036]"
		}
//line /usr/local/go/src/regexp/regexp.go:249
		// _ = "end of CoverTab[65035]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:250
		_go_fuzz_dep_.CoverTab[65037]++
//line /usr/local/go/src/regexp/regexp.go:250
		// _ = "end of CoverTab[65037]"
//line /usr/local/go/src/regexp/regexp.go:250
	}
//line /usr/local/go/src/regexp/regexp.go:250
	// _ = "end of CoverTab[65029]"
//line /usr/local/go/src/regexp/regexp.go:250
	_go_fuzz_dep_.CoverTab[65030]++

//line /usr/local/go/src/regexp/regexp.go:254
	n := matchSize[re.mpool]
	if n == 0 {
//line /usr/local/go/src/regexp/regexp.go:255
		_go_fuzz_dep_.CoverTab[65038]++
							n = len(re.prog.Inst)
//line /usr/local/go/src/regexp/regexp.go:256
		// _ = "end of CoverTab[65038]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:257
		_go_fuzz_dep_.CoverTab[65039]++
//line /usr/local/go/src/regexp/regexp.go:257
		// _ = "end of CoverTab[65039]"
//line /usr/local/go/src/regexp/regexp.go:257
	}
//line /usr/local/go/src/regexp/regexp.go:257
	// _ = "end of CoverTab[65030]"
//line /usr/local/go/src/regexp/regexp.go:257
	_go_fuzz_dep_.CoverTab[65031]++
						if len(m.q0.sparse) < n {
//line /usr/local/go/src/regexp/regexp.go:258
		_go_fuzz_dep_.CoverTab[65040]++
							m.q0 = queue{make([]uint32, n), make([]entry, 0, n)}
							m.q1 = queue{make([]uint32, n), make([]entry, 0, n)}
//line /usr/local/go/src/regexp/regexp.go:260
		// _ = "end of CoverTab[65040]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:261
		_go_fuzz_dep_.CoverTab[65041]++
//line /usr/local/go/src/regexp/regexp.go:261
		// _ = "end of CoverTab[65041]"
//line /usr/local/go/src/regexp/regexp.go:261
	}
//line /usr/local/go/src/regexp/regexp.go:261
	// _ = "end of CoverTab[65031]"
//line /usr/local/go/src/regexp/regexp.go:261
	_go_fuzz_dep_.CoverTab[65032]++
						return m
//line /usr/local/go/src/regexp/regexp.go:262
	// _ = "end of CoverTab[65032]"
}

// put returns a machine to the correct machine pool.
func (re *Regexp) put(m *machine) {
//line /usr/local/go/src/regexp/regexp.go:266
	_go_fuzz_dep_.CoverTab[65042]++
						m.re = nil
						m.p = nil
						m.inputs.clear()
						matchPool[re.mpool].Put(m)
//line /usr/local/go/src/regexp/regexp.go:270
	// _ = "end of CoverTab[65042]"
}

// minInputLen walks the regexp to find the minimum length of any matchable input.
func minInputLen(re *syntax.Regexp) int {
//line /usr/local/go/src/regexp/regexp.go:274
	_go_fuzz_dep_.CoverTab[65043]++
						switch re.Op {
	default:
//line /usr/local/go/src/regexp/regexp.go:276
		_go_fuzz_dep_.CoverTab[65044]++
							return 0
//line /usr/local/go/src/regexp/regexp.go:277
		// _ = "end of CoverTab[65044]"
	case syntax.OpAnyChar, syntax.OpAnyCharNotNL, syntax.OpCharClass:
//line /usr/local/go/src/regexp/regexp.go:278
		_go_fuzz_dep_.CoverTab[65045]++
							return 1
//line /usr/local/go/src/regexp/regexp.go:279
		// _ = "end of CoverTab[65045]"
	case syntax.OpLiteral:
//line /usr/local/go/src/regexp/regexp.go:280
		_go_fuzz_dep_.CoverTab[65046]++
							l := 0
							for _, r := range re.Rune {
//line /usr/local/go/src/regexp/regexp.go:282
			_go_fuzz_dep_.CoverTab[65054]++
								if r == utf8.RuneError {
//line /usr/local/go/src/regexp/regexp.go:283
				_go_fuzz_dep_.CoverTab[65055]++
									l++
//line /usr/local/go/src/regexp/regexp.go:284
				// _ = "end of CoverTab[65055]"
			} else {
//line /usr/local/go/src/regexp/regexp.go:285
				_go_fuzz_dep_.CoverTab[65056]++
									l += utf8.RuneLen(r)
//line /usr/local/go/src/regexp/regexp.go:286
				// _ = "end of CoverTab[65056]"
			}
//line /usr/local/go/src/regexp/regexp.go:287
			// _ = "end of CoverTab[65054]"
		}
//line /usr/local/go/src/regexp/regexp.go:288
		// _ = "end of CoverTab[65046]"
//line /usr/local/go/src/regexp/regexp.go:288
		_go_fuzz_dep_.CoverTab[65047]++
							return l
//line /usr/local/go/src/regexp/regexp.go:289
		// _ = "end of CoverTab[65047]"
	case syntax.OpCapture, syntax.OpPlus:
//line /usr/local/go/src/regexp/regexp.go:290
		_go_fuzz_dep_.CoverTab[65048]++
							return minInputLen(re.Sub[0])
//line /usr/local/go/src/regexp/regexp.go:291
		// _ = "end of CoverTab[65048]"
	case syntax.OpRepeat:
//line /usr/local/go/src/regexp/regexp.go:292
		_go_fuzz_dep_.CoverTab[65049]++
							return re.Min * minInputLen(re.Sub[0])
//line /usr/local/go/src/regexp/regexp.go:293
		// _ = "end of CoverTab[65049]"
	case syntax.OpConcat:
//line /usr/local/go/src/regexp/regexp.go:294
		_go_fuzz_dep_.CoverTab[65050]++
							l := 0
							for _, sub := range re.Sub {
//line /usr/local/go/src/regexp/regexp.go:296
			_go_fuzz_dep_.CoverTab[65057]++
								l += minInputLen(sub)
//line /usr/local/go/src/regexp/regexp.go:297
			// _ = "end of CoverTab[65057]"
		}
//line /usr/local/go/src/regexp/regexp.go:298
		// _ = "end of CoverTab[65050]"
//line /usr/local/go/src/regexp/regexp.go:298
		_go_fuzz_dep_.CoverTab[65051]++
							return l
//line /usr/local/go/src/regexp/regexp.go:299
		// _ = "end of CoverTab[65051]"
	case syntax.OpAlternate:
//line /usr/local/go/src/regexp/regexp.go:300
		_go_fuzz_dep_.CoverTab[65052]++
							l := minInputLen(re.Sub[0])
							var lnext int
							for _, sub := range re.Sub[1:] {
//line /usr/local/go/src/regexp/regexp.go:303
			_go_fuzz_dep_.CoverTab[65058]++
								lnext = minInputLen(sub)
								if lnext < l {
//line /usr/local/go/src/regexp/regexp.go:305
				_go_fuzz_dep_.CoverTab[65059]++
									l = lnext
//line /usr/local/go/src/regexp/regexp.go:306
				// _ = "end of CoverTab[65059]"
			} else {
//line /usr/local/go/src/regexp/regexp.go:307
				_go_fuzz_dep_.CoverTab[65060]++
//line /usr/local/go/src/regexp/regexp.go:307
				// _ = "end of CoverTab[65060]"
//line /usr/local/go/src/regexp/regexp.go:307
			}
//line /usr/local/go/src/regexp/regexp.go:307
			// _ = "end of CoverTab[65058]"
		}
//line /usr/local/go/src/regexp/regexp.go:308
		// _ = "end of CoverTab[65052]"
//line /usr/local/go/src/regexp/regexp.go:308
		_go_fuzz_dep_.CoverTab[65053]++
							return l
//line /usr/local/go/src/regexp/regexp.go:309
		// _ = "end of CoverTab[65053]"
	}
//line /usr/local/go/src/regexp/regexp.go:310
	// _ = "end of CoverTab[65043]"
}

// MustCompile is like Compile but panics if the expression cannot be parsed.
//line /usr/local/go/src/regexp/regexp.go:313
// It simplifies safe initialization of global variables holding compiled regular
//line /usr/local/go/src/regexp/regexp.go:313
// expressions.
//line /usr/local/go/src/regexp/regexp.go:316
func MustCompile(str string) *Regexp {
//line /usr/local/go/src/regexp/regexp.go:316
	_go_fuzz_dep_.CoverTab[65061]++
						regexp, err := Compile(str)
						if err != nil {
//line /usr/local/go/src/regexp/regexp.go:318
		_go_fuzz_dep_.CoverTab[65063]++
							panic(`regexp: Compile(` + quote(str) + `): ` + err.Error())
//line /usr/local/go/src/regexp/regexp.go:319
		// _ = "end of CoverTab[65063]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:320
		_go_fuzz_dep_.CoverTab[65064]++
//line /usr/local/go/src/regexp/regexp.go:320
		// _ = "end of CoverTab[65064]"
//line /usr/local/go/src/regexp/regexp.go:320
	}
//line /usr/local/go/src/regexp/regexp.go:320
	// _ = "end of CoverTab[65061]"
//line /usr/local/go/src/regexp/regexp.go:320
	_go_fuzz_dep_.CoverTab[65062]++
						return regexp
//line /usr/local/go/src/regexp/regexp.go:321
	// _ = "end of CoverTab[65062]"
}

// MustCompilePOSIX is like CompilePOSIX but panics if the expression cannot be parsed.
//line /usr/local/go/src/regexp/regexp.go:324
// It simplifies safe initialization of global variables holding compiled regular
//line /usr/local/go/src/regexp/regexp.go:324
// expressions.
//line /usr/local/go/src/regexp/regexp.go:327
func MustCompilePOSIX(str string) *Regexp {
//line /usr/local/go/src/regexp/regexp.go:327
	_go_fuzz_dep_.CoverTab[65065]++
						regexp, err := CompilePOSIX(str)
						if err != nil {
//line /usr/local/go/src/regexp/regexp.go:329
		_go_fuzz_dep_.CoverTab[65067]++
							panic(`regexp: CompilePOSIX(` + quote(str) + `): ` + err.Error())
//line /usr/local/go/src/regexp/regexp.go:330
		// _ = "end of CoverTab[65067]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:331
		_go_fuzz_dep_.CoverTab[65068]++
//line /usr/local/go/src/regexp/regexp.go:331
		// _ = "end of CoverTab[65068]"
//line /usr/local/go/src/regexp/regexp.go:331
	}
//line /usr/local/go/src/regexp/regexp.go:331
	// _ = "end of CoverTab[65065]"
//line /usr/local/go/src/regexp/regexp.go:331
	_go_fuzz_dep_.CoverTab[65066]++
						return regexp
//line /usr/local/go/src/regexp/regexp.go:332
	// _ = "end of CoverTab[65066]"
}

func quote(s string) string {
//line /usr/local/go/src/regexp/regexp.go:335
	_go_fuzz_dep_.CoverTab[65069]++
						if strconv.CanBackquote(s) {
//line /usr/local/go/src/regexp/regexp.go:336
		_go_fuzz_dep_.CoverTab[65071]++
							return "`" + s + "`"
//line /usr/local/go/src/regexp/regexp.go:337
		// _ = "end of CoverTab[65071]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:338
		_go_fuzz_dep_.CoverTab[65072]++
//line /usr/local/go/src/regexp/regexp.go:338
		// _ = "end of CoverTab[65072]"
//line /usr/local/go/src/regexp/regexp.go:338
	}
//line /usr/local/go/src/regexp/regexp.go:338
	// _ = "end of CoverTab[65069]"
//line /usr/local/go/src/regexp/regexp.go:338
	_go_fuzz_dep_.CoverTab[65070]++
						return strconv.Quote(s)
//line /usr/local/go/src/regexp/regexp.go:339
	// _ = "end of CoverTab[65070]"
}

// NumSubexp returns the number of parenthesized subexpressions in this Regexp.
func (re *Regexp) NumSubexp() int {
//line /usr/local/go/src/regexp/regexp.go:343
	_go_fuzz_dep_.CoverTab[65073]++
						return re.numSubexp
//line /usr/local/go/src/regexp/regexp.go:344
	// _ = "end of CoverTab[65073]"
}

// SubexpNames returns the names of the parenthesized subexpressions
//line /usr/local/go/src/regexp/regexp.go:347
// in this Regexp. The name for the first sub-expression is names[1],
//line /usr/local/go/src/regexp/regexp.go:347
// so that if m is a match slice, the name for m[i] is SubexpNames()[i].
//line /usr/local/go/src/regexp/regexp.go:347
// Since the Regexp as a whole cannot be named, names[0] is always
//line /usr/local/go/src/regexp/regexp.go:347
// the empty string. The slice should not be modified.
//line /usr/local/go/src/regexp/regexp.go:352
func (re *Regexp) SubexpNames() []string {
//line /usr/local/go/src/regexp/regexp.go:352
	_go_fuzz_dep_.CoverTab[65074]++
						return re.subexpNames
//line /usr/local/go/src/regexp/regexp.go:353
	// _ = "end of CoverTab[65074]"
}

// SubexpIndex returns the index of the first subexpression with the given name,
//line /usr/local/go/src/regexp/regexp.go:356
// or -1 if there is no subexpression with that name.
//line /usr/local/go/src/regexp/regexp.go:356
//
//line /usr/local/go/src/regexp/regexp.go:356
// Note that multiple subexpressions can be written using the same name, as in
//line /usr/local/go/src/regexp/regexp.go:356
// (?P<bob>a+)(?P<bob>b+), which declares two subexpressions named "bob".
//line /usr/local/go/src/regexp/regexp.go:356
// In this case, SubexpIndex returns the index of the leftmost such subexpression
//line /usr/local/go/src/regexp/regexp.go:356
// in the regular expression.
//line /usr/local/go/src/regexp/regexp.go:363
func (re *Regexp) SubexpIndex(name string) int {
//line /usr/local/go/src/regexp/regexp.go:363
	_go_fuzz_dep_.CoverTab[65075]++
						if name != "" {
//line /usr/local/go/src/regexp/regexp.go:364
		_go_fuzz_dep_.CoverTab[65077]++
							for i, s := range re.subexpNames {
//line /usr/local/go/src/regexp/regexp.go:365
			_go_fuzz_dep_.CoverTab[65078]++
								if name == s {
//line /usr/local/go/src/regexp/regexp.go:366
				_go_fuzz_dep_.CoverTab[65079]++
									return i
//line /usr/local/go/src/regexp/regexp.go:367
				// _ = "end of CoverTab[65079]"
			} else {
//line /usr/local/go/src/regexp/regexp.go:368
				_go_fuzz_dep_.CoverTab[65080]++
//line /usr/local/go/src/regexp/regexp.go:368
				// _ = "end of CoverTab[65080]"
//line /usr/local/go/src/regexp/regexp.go:368
			}
//line /usr/local/go/src/regexp/regexp.go:368
			// _ = "end of CoverTab[65078]"
		}
//line /usr/local/go/src/regexp/regexp.go:369
		// _ = "end of CoverTab[65077]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:370
		_go_fuzz_dep_.CoverTab[65081]++
//line /usr/local/go/src/regexp/regexp.go:370
		// _ = "end of CoverTab[65081]"
//line /usr/local/go/src/regexp/regexp.go:370
	}
//line /usr/local/go/src/regexp/regexp.go:370
	// _ = "end of CoverTab[65075]"
//line /usr/local/go/src/regexp/regexp.go:370
	_go_fuzz_dep_.CoverTab[65076]++
						return -1
//line /usr/local/go/src/regexp/regexp.go:371
	// _ = "end of CoverTab[65076]"
}

const endOfText rune = -1

// input abstracts different representations of the input text. It provides
//line /usr/local/go/src/regexp/regexp.go:376
// one-character lookahead.
//line /usr/local/go/src/regexp/regexp.go:378
type input interface {
	step(pos int) (r rune, width int)	// advance one rune
	canCheckPrefix() bool			// can we look ahead without losing info?
	hasPrefix(re *Regexp) bool
	index(re *Regexp, pos int) int
	context(pos int) lazyFlag
}

// inputString scans a string.
type inputString struct {
	str string
}

func (i *inputString) step(pos int) (rune, int) {
//line /usr/local/go/src/regexp/regexp.go:391
	_go_fuzz_dep_.CoverTab[65082]++
						if pos < len(i.str) {
//line /usr/local/go/src/regexp/regexp.go:392
		_go_fuzz_dep_.CoverTab[65084]++
							c := i.str[pos]
							if c < utf8.RuneSelf {
//line /usr/local/go/src/regexp/regexp.go:394
			_go_fuzz_dep_.CoverTab[65086]++
								return rune(c), 1
//line /usr/local/go/src/regexp/regexp.go:395
			// _ = "end of CoverTab[65086]"
		} else {
//line /usr/local/go/src/regexp/regexp.go:396
			_go_fuzz_dep_.CoverTab[65087]++
//line /usr/local/go/src/regexp/regexp.go:396
			// _ = "end of CoverTab[65087]"
//line /usr/local/go/src/regexp/regexp.go:396
		}
//line /usr/local/go/src/regexp/regexp.go:396
		// _ = "end of CoverTab[65084]"
//line /usr/local/go/src/regexp/regexp.go:396
		_go_fuzz_dep_.CoverTab[65085]++
							return utf8.DecodeRuneInString(i.str[pos:])
//line /usr/local/go/src/regexp/regexp.go:397
		// _ = "end of CoverTab[65085]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:398
		_go_fuzz_dep_.CoverTab[65088]++
//line /usr/local/go/src/regexp/regexp.go:398
		// _ = "end of CoverTab[65088]"
//line /usr/local/go/src/regexp/regexp.go:398
	}
//line /usr/local/go/src/regexp/regexp.go:398
	// _ = "end of CoverTab[65082]"
//line /usr/local/go/src/regexp/regexp.go:398
	_go_fuzz_dep_.CoverTab[65083]++
						return endOfText, 0
//line /usr/local/go/src/regexp/regexp.go:399
	// _ = "end of CoverTab[65083]"
}

func (i *inputString) canCheckPrefix() bool {
//line /usr/local/go/src/regexp/regexp.go:402
	_go_fuzz_dep_.CoverTab[65089]++
						return true
//line /usr/local/go/src/regexp/regexp.go:403
	// _ = "end of CoverTab[65089]"
}

func (i *inputString) hasPrefix(re *Regexp) bool {
//line /usr/local/go/src/regexp/regexp.go:406
	_go_fuzz_dep_.CoverTab[65090]++
						return strings.HasPrefix(i.str, re.prefix)
//line /usr/local/go/src/regexp/regexp.go:407
	// _ = "end of CoverTab[65090]"
}

func (i *inputString) index(re *Regexp, pos int) int {
//line /usr/local/go/src/regexp/regexp.go:410
	_go_fuzz_dep_.CoverTab[65091]++
						return strings.Index(i.str[pos:], re.prefix)
//line /usr/local/go/src/regexp/regexp.go:411
	// _ = "end of CoverTab[65091]"
}

func (i *inputString) context(pos int) lazyFlag {
//line /usr/local/go/src/regexp/regexp.go:414
	_go_fuzz_dep_.CoverTab[65092]++
						r1, r2 := endOfText, endOfText

						if uint(pos-1) < uint(len(i.str)) {
//line /usr/local/go/src/regexp/regexp.go:417
		_go_fuzz_dep_.CoverTab[65095]++
							r1 = rune(i.str[pos-1])
							if r1 >= utf8.RuneSelf {
//line /usr/local/go/src/regexp/regexp.go:419
			_go_fuzz_dep_.CoverTab[65096]++
								r1, _ = utf8.DecodeLastRuneInString(i.str[:pos])
//line /usr/local/go/src/regexp/regexp.go:420
			// _ = "end of CoverTab[65096]"
		} else {
//line /usr/local/go/src/regexp/regexp.go:421
			_go_fuzz_dep_.CoverTab[65097]++
//line /usr/local/go/src/regexp/regexp.go:421
			// _ = "end of CoverTab[65097]"
//line /usr/local/go/src/regexp/regexp.go:421
		}
//line /usr/local/go/src/regexp/regexp.go:421
		// _ = "end of CoverTab[65095]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:422
		_go_fuzz_dep_.CoverTab[65098]++
//line /usr/local/go/src/regexp/regexp.go:422
		// _ = "end of CoverTab[65098]"
//line /usr/local/go/src/regexp/regexp.go:422
	}
//line /usr/local/go/src/regexp/regexp.go:422
	// _ = "end of CoverTab[65092]"
//line /usr/local/go/src/regexp/regexp.go:422
	_go_fuzz_dep_.CoverTab[65093]++

						if uint(pos) < uint(len(i.str)) {
//line /usr/local/go/src/regexp/regexp.go:424
		_go_fuzz_dep_.CoverTab[65099]++
							r2 = rune(i.str[pos])
							if r2 >= utf8.RuneSelf {
//line /usr/local/go/src/regexp/regexp.go:426
			_go_fuzz_dep_.CoverTab[65100]++
								r2, _ = utf8.DecodeRuneInString(i.str[pos:])
//line /usr/local/go/src/regexp/regexp.go:427
			// _ = "end of CoverTab[65100]"
		} else {
//line /usr/local/go/src/regexp/regexp.go:428
			_go_fuzz_dep_.CoverTab[65101]++
//line /usr/local/go/src/regexp/regexp.go:428
			// _ = "end of CoverTab[65101]"
//line /usr/local/go/src/regexp/regexp.go:428
		}
//line /usr/local/go/src/regexp/regexp.go:428
		// _ = "end of CoverTab[65099]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:429
		_go_fuzz_dep_.CoverTab[65102]++
//line /usr/local/go/src/regexp/regexp.go:429
		// _ = "end of CoverTab[65102]"
//line /usr/local/go/src/regexp/regexp.go:429
	}
//line /usr/local/go/src/regexp/regexp.go:429
	// _ = "end of CoverTab[65093]"
//line /usr/local/go/src/regexp/regexp.go:429
	_go_fuzz_dep_.CoverTab[65094]++
						return newLazyFlag(r1, r2)
//line /usr/local/go/src/regexp/regexp.go:430
	// _ = "end of CoverTab[65094]"
}

// inputBytes scans a byte slice.
type inputBytes struct {
	str []byte
}

func (i *inputBytes) step(pos int) (rune, int) {
//line /usr/local/go/src/regexp/regexp.go:438
	_go_fuzz_dep_.CoverTab[65103]++
						if pos < len(i.str) {
//line /usr/local/go/src/regexp/regexp.go:439
		_go_fuzz_dep_.CoverTab[65105]++
							c := i.str[pos]
							if c < utf8.RuneSelf {
//line /usr/local/go/src/regexp/regexp.go:441
			_go_fuzz_dep_.CoverTab[65107]++
								return rune(c), 1
//line /usr/local/go/src/regexp/regexp.go:442
			// _ = "end of CoverTab[65107]"
		} else {
//line /usr/local/go/src/regexp/regexp.go:443
			_go_fuzz_dep_.CoverTab[65108]++
//line /usr/local/go/src/regexp/regexp.go:443
			// _ = "end of CoverTab[65108]"
//line /usr/local/go/src/regexp/regexp.go:443
		}
//line /usr/local/go/src/regexp/regexp.go:443
		// _ = "end of CoverTab[65105]"
//line /usr/local/go/src/regexp/regexp.go:443
		_go_fuzz_dep_.CoverTab[65106]++
							return utf8.DecodeRune(i.str[pos:])
//line /usr/local/go/src/regexp/regexp.go:444
		// _ = "end of CoverTab[65106]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:445
		_go_fuzz_dep_.CoverTab[65109]++
//line /usr/local/go/src/regexp/regexp.go:445
		// _ = "end of CoverTab[65109]"
//line /usr/local/go/src/regexp/regexp.go:445
	}
//line /usr/local/go/src/regexp/regexp.go:445
	// _ = "end of CoverTab[65103]"
//line /usr/local/go/src/regexp/regexp.go:445
	_go_fuzz_dep_.CoverTab[65104]++
						return endOfText, 0
//line /usr/local/go/src/regexp/regexp.go:446
	// _ = "end of CoverTab[65104]"
}

func (i *inputBytes) canCheckPrefix() bool {
//line /usr/local/go/src/regexp/regexp.go:449
	_go_fuzz_dep_.CoverTab[65110]++
						return true
//line /usr/local/go/src/regexp/regexp.go:450
	// _ = "end of CoverTab[65110]"
}

func (i *inputBytes) hasPrefix(re *Regexp) bool {
//line /usr/local/go/src/regexp/regexp.go:453
	_go_fuzz_dep_.CoverTab[65111]++
						return bytes.HasPrefix(i.str, re.prefixBytes)
//line /usr/local/go/src/regexp/regexp.go:454
	// _ = "end of CoverTab[65111]"
}

func (i *inputBytes) index(re *Regexp, pos int) int {
//line /usr/local/go/src/regexp/regexp.go:457
	_go_fuzz_dep_.CoverTab[65112]++
						return bytes.Index(i.str[pos:], re.prefixBytes)
//line /usr/local/go/src/regexp/regexp.go:458
	// _ = "end of CoverTab[65112]"
}

func (i *inputBytes) context(pos int) lazyFlag {
//line /usr/local/go/src/regexp/regexp.go:461
	_go_fuzz_dep_.CoverTab[65113]++
						r1, r2 := endOfText, endOfText

						if uint(pos-1) < uint(len(i.str)) {
//line /usr/local/go/src/regexp/regexp.go:464
		_go_fuzz_dep_.CoverTab[65116]++
							r1 = rune(i.str[pos-1])
							if r1 >= utf8.RuneSelf {
//line /usr/local/go/src/regexp/regexp.go:466
			_go_fuzz_dep_.CoverTab[65117]++
								r1, _ = utf8.DecodeLastRune(i.str[:pos])
//line /usr/local/go/src/regexp/regexp.go:467
			// _ = "end of CoverTab[65117]"
		} else {
//line /usr/local/go/src/regexp/regexp.go:468
			_go_fuzz_dep_.CoverTab[65118]++
//line /usr/local/go/src/regexp/regexp.go:468
			// _ = "end of CoverTab[65118]"
//line /usr/local/go/src/regexp/regexp.go:468
		}
//line /usr/local/go/src/regexp/regexp.go:468
		// _ = "end of CoverTab[65116]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:469
		_go_fuzz_dep_.CoverTab[65119]++
//line /usr/local/go/src/regexp/regexp.go:469
		// _ = "end of CoverTab[65119]"
//line /usr/local/go/src/regexp/regexp.go:469
	}
//line /usr/local/go/src/regexp/regexp.go:469
	// _ = "end of CoverTab[65113]"
//line /usr/local/go/src/regexp/regexp.go:469
	_go_fuzz_dep_.CoverTab[65114]++

						if uint(pos) < uint(len(i.str)) {
//line /usr/local/go/src/regexp/regexp.go:471
		_go_fuzz_dep_.CoverTab[65120]++
							r2 = rune(i.str[pos])
							if r2 >= utf8.RuneSelf {
//line /usr/local/go/src/regexp/regexp.go:473
			_go_fuzz_dep_.CoverTab[65121]++
								r2, _ = utf8.DecodeRune(i.str[pos:])
//line /usr/local/go/src/regexp/regexp.go:474
			// _ = "end of CoverTab[65121]"
		} else {
//line /usr/local/go/src/regexp/regexp.go:475
			_go_fuzz_dep_.CoverTab[65122]++
//line /usr/local/go/src/regexp/regexp.go:475
			// _ = "end of CoverTab[65122]"
//line /usr/local/go/src/regexp/regexp.go:475
		}
//line /usr/local/go/src/regexp/regexp.go:475
		// _ = "end of CoverTab[65120]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:476
		_go_fuzz_dep_.CoverTab[65123]++
//line /usr/local/go/src/regexp/regexp.go:476
		// _ = "end of CoverTab[65123]"
//line /usr/local/go/src/regexp/regexp.go:476
	}
//line /usr/local/go/src/regexp/regexp.go:476
	// _ = "end of CoverTab[65114]"
//line /usr/local/go/src/regexp/regexp.go:476
	_go_fuzz_dep_.CoverTab[65115]++
						return newLazyFlag(r1, r2)
//line /usr/local/go/src/regexp/regexp.go:477
	// _ = "end of CoverTab[65115]"
}

// inputReader scans a RuneReader.
type inputReader struct {
	r	io.RuneReader
	atEOT	bool
	pos	int
}

func (i *inputReader) step(pos int) (rune, int) {
//line /usr/local/go/src/regexp/regexp.go:487
	_go_fuzz_dep_.CoverTab[65124]++
						if !i.atEOT && func() bool {
//line /usr/local/go/src/regexp/regexp.go:488
		_go_fuzz_dep_.CoverTab[65127]++
//line /usr/local/go/src/regexp/regexp.go:488
		return pos != i.pos
//line /usr/local/go/src/regexp/regexp.go:488
		// _ = "end of CoverTab[65127]"
//line /usr/local/go/src/regexp/regexp.go:488
	}() {
//line /usr/local/go/src/regexp/regexp.go:488
		_go_fuzz_dep_.CoverTab[65128]++
							return endOfText, 0
//line /usr/local/go/src/regexp/regexp.go:489
		// _ = "end of CoverTab[65128]"

	} else {
//line /usr/local/go/src/regexp/regexp.go:491
		_go_fuzz_dep_.CoverTab[65129]++
//line /usr/local/go/src/regexp/regexp.go:491
		// _ = "end of CoverTab[65129]"
//line /usr/local/go/src/regexp/regexp.go:491
	}
//line /usr/local/go/src/regexp/regexp.go:491
	// _ = "end of CoverTab[65124]"
//line /usr/local/go/src/regexp/regexp.go:491
	_go_fuzz_dep_.CoverTab[65125]++
						r, w, err := i.r.ReadRune()
						if err != nil {
//line /usr/local/go/src/regexp/regexp.go:493
		_go_fuzz_dep_.CoverTab[65130]++
							i.atEOT = true
							return endOfText, 0
//line /usr/local/go/src/regexp/regexp.go:495
		// _ = "end of CoverTab[65130]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:496
		_go_fuzz_dep_.CoverTab[65131]++
//line /usr/local/go/src/regexp/regexp.go:496
		// _ = "end of CoverTab[65131]"
//line /usr/local/go/src/regexp/regexp.go:496
	}
//line /usr/local/go/src/regexp/regexp.go:496
	// _ = "end of CoverTab[65125]"
//line /usr/local/go/src/regexp/regexp.go:496
	_go_fuzz_dep_.CoverTab[65126]++
						i.pos += w
						return r, w
//line /usr/local/go/src/regexp/regexp.go:498
	// _ = "end of CoverTab[65126]"
}

func (i *inputReader) canCheckPrefix() bool {
//line /usr/local/go/src/regexp/regexp.go:501
	_go_fuzz_dep_.CoverTab[65132]++
						return false
//line /usr/local/go/src/regexp/regexp.go:502
	// _ = "end of CoverTab[65132]"
}

func (i *inputReader) hasPrefix(re *Regexp) bool {
//line /usr/local/go/src/regexp/regexp.go:505
	_go_fuzz_dep_.CoverTab[65133]++
						return false
//line /usr/local/go/src/regexp/regexp.go:506
	// _ = "end of CoverTab[65133]"
}

func (i *inputReader) index(re *Regexp, pos int) int {
//line /usr/local/go/src/regexp/regexp.go:509
	_go_fuzz_dep_.CoverTab[65134]++
						return -1
//line /usr/local/go/src/regexp/regexp.go:510
	// _ = "end of CoverTab[65134]"
}

func (i *inputReader) context(pos int) lazyFlag {
//line /usr/local/go/src/regexp/regexp.go:513
	_go_fuzz_dep_.CoverTab[65135]++
						return 0
//line /usr/local/go/src/regexp/regexp.go:514
	// _ = "end of CoverTab[65135]"
}

// LiteralPrefix returns a literal string that must begin any match
//line /usr/local/go/src/regexp/regexp.go:517
// of the regular expression re. It returns the boolean true if the
//line /usr/local/go/src/regexp/regexp.go:517
// literal string comprises the entire regular expression.
//line /usr/local/go/src/regexp/regexp.go:520
func (re *Regexp) LiteralPrefix() (prefix string, complete bool) {
//line /usr/local/go/src/regexp/regexp.go:520
	_go_fuzz_dep_.CoverTab[65136]++
						return re.prefix, re.prefixComplete
//line /usr/local/go/src/regexp/regexp.go:521
	// _ = "end of CoverTab[65136]"
}

// MatchReader reports whether the text returned by the RuneReader
//line /usr/local/go/src/regexp/regexp.go:524
// contains any match of the regular expression re.
//line /usr/local/go/src/regexp/regexp.go:526
func (re *Regexp) MatchReader(r io.RuneReader) bool {
//line /usr/local/go/src/regexp/regexp.go:526
	_go_fuzz_dep_.CoverTab[65137]++
						return re.doMatch(r, nil, "")
//line /usr/local/go/src/regexp/regexp.go:527
	// _ = "end of CoverTab[65137]"
}

// MatchString reports whether the string s
//line /usr/local/go/src/regexp/regexp.go:530
// contains any match of the regular expression re.
//line /usr/local/go/src/regexp/regexp.go:532
func (re *Regexp) MatchString(s string) bool {
//line /usr/local/go/src/regexp/regexp.go:532
	_go_fuzz_dep_.CoverTab[65138]++
						return re.doMatch(nil, nil, s)
//line /usr/local/go/src/regexp/regexp.go:533
	// _ = "end of CoverTab[65138]"
}

// Match reports whether the byte slice b
//line /usr/local/go/src/regexp/regexp.go:536
// contains any match of the regular expression re.
//line /usr/local/go/src/regexp/regexp.go:538
func (re *Regexp) Match(b []byte) bool {
//line /usr/local/go/src/regexp/regexp.go:538
	_go_fuzz_dep_.CoverTab[65139]++
						return re.doMatch(nil, b, "")
//line /usr/local/go/src/regexp/regexp.go:539
	// _ = "end of CoverTab[65139]"
}

// MatchReader reports whether the text returned by the RuneReader
//line /usr/local/go/src/regexp/regexp.go:542
// contains any match of the regular expression pattern.
//line /usr/local/go/src/regexp/regexp.go:542
// More complicated queries need to use Compile and the full Regexp interface.
//line /usr/local/go/src/regexp/regexp.go:545
func MatchReader(pattern string, r io.RuneReader) (matched bool, err error) {
//line /usr/local/go/src/regexp/regexp.go:545
	_go_fuzz_dep_.CoverTab[65140]++
						re, err := Compile(pattern)
						if err != nil {
//line /usr/local/go/src/regexp/regexp.go:547
		_go_fuzz_dep_.CoverTab[65142]++
							return false, err
//line /usr/local/go/src/regexp/regexp.go:548
		// _ = "end of CoverTab[65142]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:549
		_go_fuzz_dep_.CoverTab[65143]++
//line /usr/local/go/src/regexp/regexp.go:549
		// _ = "end of CoverTab[65143]"
//line /usr/local/go/src/regexp/regexp.go:549
	}
//line /usr/local/go/src/regexp/regexp.go:549
	// _ = "end of CoverTab[65140]"
//line /usr/local/go/src/regexp/regexp.go:549
	_go_fuzz_dep_.CoverTab[65141]++
						return re.MatchReader(r), nil
//line /usr/local/go/src/regexp/regexp.go:550
	// _ = "end of CoverTab[65141]"
}

// MatchString reports whether the string s
//line /usr/local/go/src/regexp/regexp.go:553
// contains any match of the regular expression pattern.
//line /usr/local/go/src/regexp/regexp.go:553
// More complicated queries need to use Compile and the full Regexp interface.
//line /usr/local/go/src/regexp/regexp.go:556
func MatchString(pattern string, s string) (matched bool, err error) {
//line /usr/local/go/src/regexp/regexp.go:556
	_go_fuzz_dep_.CoverTab[65144]++
						re, err := Compile(pattern)
						if err != nil {
//line /usr/local/go/src/regexp/regexp.go:558
		_go_fuzz_dep_.CoverTab[65146]++
							return false, err
//line /usr/local/go/src/regexp/regexp.go:559
		// _ = "end of CoverTab[65146]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:560
		_go_fuzz_dep_.CoverTab[65147]++
//line /usr/local/go/src/regexp/regexp.go:560
		// _ = "end of CoverTab[65147]"
//line /usr/local/go/src/regexp/regexp.go:560
	}
//line /usr/local/go/src/regexp/regexp.go:560
	// _ = "end of CoverTab[65144]"
//line /usr/local/go/src/regexp/regexp.go:560
	_go_fuzz_dep_.CoverTab[65145]++
						return re.MatchString(s), nil
//line /usr/local/go/src/regexp/regexp.go:561
	// _ = "end of CoverTab[65145]"
}

// Match reports whether the byte slice b
//line /usr/local/go/src/regexp/regexp.go:564
// contains any match of the regular expression pattern.
//line /usr/local/go/src/regexp/regexp.go:564
// More complicated queries need to use Compile and the full Regexp interface.
//line /usr/local/go/src/regexp/regexp.go:567
func Match(pattern string, b []byte) (matched bool, err error) {
//line /usr/local/go/src/regexp/regexp.go:567
	_go_fuzz_dep_.CoverTab[65148]++
						re, err := Compile(pattern)
						if err != nil {
//line /usr/local/go/src/regexp/regexp.go:569
		_go_fuzz_dep_.CoverTab[65150]++
							return false, err
//line /usr/local/go/src/regexp/regexp.go:570
		// _ = "end of CoverTab[65150]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:571
		_go_fuzz_dep_.CoverTab[65151]++
//line /usr/local/go/src/regexp/regexp.go:571
		// _ = "end of CoverTab[65151]"
//line /usr/local/go/src/regexp/regexp.go:571
	}
//line /usr/local/go/src/regexp/regexp.go:571
	// _ = "end of CoverTab[65148]"
//line /usr/local/go/src/regexp/regexp.go:571
	_go_fuzz_dep_.CoverTab[65149]++
						return re.Match(b), nil
//line /usr/local/go/src/regexp/regexp.go:572
	// _ = "end of CoverTab[65149]"
}

// ReplaceAllString returns a copy of src, replacing matches of the Regexp
//line /usr/local/go/src/regexp/regexp.go:575
// with the replacement string repl. Inside repl, $ signs are interpreted as
//line /usr/local/go/src/regexp/regexp.go:575
// in Expand, so for instance $1 represents the text of the first submatch.
//line /usr/local/go/src/regexp/regexp.go:578
func (re *Regexp) ReplaceAllString(src, repl string) string {
//line /usr/local/go/src/regexp/regexp.go:578
	_go_fuzz_dep_.CoverTab[65152]++
						n := 2
						if strings.Contains(repl, "$") {
//line /usr/local/go/src/regexp/regexp.go:580
		_go_fuzz_dep_.CoverTab[65155]++
							n = 2 * (re.numSubexp + 1)
//line /usr/local/go/src/regexp/regexp.go:581
		// _ = "end of CoverTab[65155]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:582
		_go_fuzz_dep_.CoverTab[65156]++
//line /usr/local/go/src/regexp/regexp.go:582
		// _ = "end of CoverTab[65156]"
//line /usr/local/go/src/regexp/regexp.go:582
	}
//line /usr/local/go/src/regexp/regexp.go:582
	// _ = "end of CoverTab[65152]"
//line /usr/local/go/src/regexp/regexp.go:582
	_go_fuzz_dep_.CoverTab[65153]++
						b := re.replaceAll(nil, src, n, func(dst []byte, match []int) []byte {
//line /usr/local/go/src/regexp/regexp.go:583
		_go_fuzz_dep_.CoverTab[65157]++
							return re.expand(dst, repl, nil, src, match)
//line /usr/local/go/src/regexp/regexp.go:584
		// _ = "end of CoverTab[65157]"
	})
//line /usr/local/go/src/regexp/regexp.go:585
	// _ = "end of CoverTab[65153]"
//line /usr/local/go/src/regexp/regexp.go:585
	_go_fuzz_dep_.CoverTab[65154]++
						return string(b)
//line /usr/local/go/src/regexp/regexp.go:586
	// _ = "end of CoverTab[65154]"
}

// ReplaceAllLiteralString returns a copy of src, replacing matches of the Regexp
//line /usr/local/go/src/regexp/regexp.go:589
// with the replacement string repl. The replacement repl is substituted directly,
//line /usr/local/go/src/regexp/regexp.go:589
// without using Expand.
//line /usr/local/go/src/regexp/regexp.go:592
func (re *Regexp) ReplaceAllLiteralString(src, repl string) string {
//line /usr/local/go/src/regexp/regexp.go:592
	_go_fuzz_dep_.CoverTab[65158]++
						return string(re.replaceAll(nil, src, 2, func(dst []byte, match []int) []byte {
//line /usr/local/go/src/regexp/regexp.go:593
		_go_fuzz_dep_.CoverTab[65159]++
							return append(dst, repl...)
//line /usr/local/go/src/regexp/regexp.go:594
		// _ = "end of CoverTab[65159]"
	}))
//line /usr/local/go/src/regexp/regexp.go:595
	// _ = "end of CoverTab[65158]"
}

// ReplaceAllStringFunc returns a copy of src in which all matches of the
//line /usr/local/go/src/regexp/regexp.go:598
// Regexp have been replaced by the return value of function repl applied
//line /usr/local/go/src/regexp/regexp.go:598
// to the matched substring. The replacement returned by repl is substituted
//line /usr/local/go/src/regexp/regexp.go:598
// directly, without using Expand.
//line /usr/local/go/src/regexp/regexp.go:602
func (re *Regexp) ReplaceAllStringFunc(src string, repl func(string) string) string {
//line /usr/local/go/src/regexp/regexp.go:602
	_go_fuzz_dep_.CoverTab[65160]++
						b := re.replaceAll(nil, src, 2, func(dst []byte, match []int) []byte {
//line /usr/local/go/src/regexp/regexp.go:603
		_go_fuzz_dep_.CoverTab[65162]++
							return append(dst, repl(src[match[0]:match[1]])...)
//line /usr/local/go/src/regexp/regexp.go:604
		// _ = "end of CoverTab[65162]"
	})
//line /usr/local/go/src/regexp/regexp.go:605
	// _ = "end of CoverTab[65160]"
//line /usr/local/go/src/regexp/regexp.go:605
	_go_fuzz_dep_.CoverTab[65161]++
						return string(b)
//line /usr/local/go/src/regexp/regexp.go:606
	// _ = "end of CoverTab[65161]"
}

func (re *Regexp) replaceAll(bsrc []byte, src string, nmatch int, repl func(dst []byte, m []int) []byte) []byte {
//line /usr/local/go/src/regexp/regexp.go:609
	_go_fuzz_dep_.CoverTab[65163]++
						lastMatchEnd := 0
						searchPos := 0
						var buf []byte
						var endPos int
						if bsrc != nil {
//line /usr/local/go/src/regexp/regexp.go:614
		_go_fuzz_dep_.CoverTab[65168]++
							endPos = len(bsrc)
//line /usr/local/go/src/regexp/regexp.go:615
		// _ = "end of CoverTab[65168]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:616
		_go_fuzz_dep_.CoverTab[65169]++
							endPos = len(src)
//line /usr/local/go/src/regexp/regexp.go:617
		// _ = "end of CoverTab[65169]"
	}
//line /usr/local/go/src/regexp/regexp.go:618
	// _ = "end of CoverTab[65163]"
//line /usr/local/go/src/regexp/regexp.go:618
	_go_fuzz_dep_.CoverTab[65164]++
						if nmatch > re.prog.NumCap {
//line /usr/local/go/src/regexp/regexp.go:619
		_go_fuzz_dep_.CoverTab[65170]++
							nmatch = re.prog.NumCap
//line /usr/local/go/src/regexp/regexp.go:620
		// _ = "end of CoverTab[65170]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:621
		_go_fuzz_dep_.CoverTab[65171]++
//line /usr/local/go/src/regexp/regexp.go:621
		// _ = "end of CoverTab[65171]"
//line /usr/local/go/src/regexp/regexp.go:621
	}
//line /usr/local/go/src/regexp/regexp.go:621
	// _ = "end of CoverTab[65164]"
//line /usr/local/go/src/regexp/regexp.go:621
	_go_fuzz_dep_.CoverTab[65165]++

						var dstCap [2]int
						for searchPos <= endPos {
//line /usr/local/go/src/regexp/regexp.go:624
		_go_fuzz_dep_.CoverTab[65172]++
							a := re.doExecute(nil, bsrc, src, searchPos, nmatch, dstCap[:0])
							if len(a) == 0 {
//line /usr/local/go/src/regexp/regexp.go:626
			_go_fuzz_dep_.CoverTab[65177]++
								break
//line /usr/local/go/src/regexp/regexp.go:627
			// _ = "end of CoverTab[65177]"
		} else {
//line /usr/local/go/src/regexp/regexp.go:628
			_go_fuzz_dep_.CoverTab[65178]++
//line /usr/local/go/src/regexp/regexp.go:628
			// _ = "end of CoverTab[65178]"
//line /usr/local/go/src/regexp/regexp.go:628
		}
//line /usr/local/go/src/regexp/regexp.go:628
		// _ = "end of CoverTab[65172]"
//line /usr/local/go/src/regexp/regexp.go:628
		_go_fuzz_dep_.CoverTab[65173]++

//line /usr/local/go/src/regexp/regexp.go:631
		if bsrc != nil {
//line /usr/local/go/src/regexp/regexp.go:631
			_go_fuzz_dep_.CoverTab[65179]++
								buf = append(buf, bsrc[lastMatchEnd:a[0]]...)
//line /usr/local/go/src/regexp/regexp.go:632
			// _ = "end of CoverTab[65179]"
		} else {
//line /usr/local/go/src/regexp/regexp.go:633
			_go_fuzz_dep_.CoverTab[65180]++
								buf = append(buf, src[lastMatchEnd:a[0]]...)
//line /usr/local/go/src/regexp/regexp.go:634
			// _ = "end of CoverTab[65180]"
		}
//line /usr/local/go/src/regexp/regexp.go:635
		// _ = "end of CoverTab[65173]"
//line /usr/local/go/src/regexp/regexp.go:635
		_go_fuzz_dep_.CoverTab[65174]++

//line /usr/local/go/src/regexp/regexp.go:641
		if a[1] > lastMatchEnd || func() bool {
//line /usr/local/go/src/regexp/regexp.go:641
			_go_fuzz_dep_.CoverTab[65181]++
//line /usr/local/go/src/regexp/regexp.go:641
			return a[0] == 0
//line /usr/local/go/src/regexp/regexp.go:641
			// _ = "end of CoverTab[65181]"
//line /usr/local/go/src/regexp/regexp.go:641
		}() {
//line /usr/local/go/src/regexp/regexp.go:641
			_go_fuzz_dep_.CoverTab[65182]++
								buf = repl(buf, a)
//line /usr/local/go/src/regexp/regexp.go:642
			// _ = "end of CoverTab[65182]"
		} else {
//line /usr/local/go/src/regexp/regexp.go:643
			_go_fuzz_dep_.CoverTab[65183]++
//line /usr/local/go/src/regexp/regexp.go:643
			// _ = "end of CoverTab[65183]"
//line /usr/local/go/src/regexp/regexp.go:643
		}
//line /usr/local/go/src/regexp/regexp.go:643
		// _ = "end of CoverTab[65174]"
//line /usr/local/go/src/regexp/regexp.go:643
		_go_fuzz_dep_.CoverTab[65175]++
							lastMatchEnd = a[1]

		// Advance past this match; always advance at least one character.
		var width int
		if bsrc != nil {
//line /usr/local/go/src/regexp/regexp.go:648
			_go_fuzz_dep_.CoverTab[65184]++
								_, width = utf8.DecodeRune(bsrc[searchPos:])
//line /usr/local/go/src/regexp/regexp.go:649
			// _ = "end of CoverTab[65184]"
		} else {
//line /usr/local/go/src/regexp/regexp.go:650
			_go_fuzz_dep_.CoverTab[65185]++
								_, width = utf8.DecodeRuneInString(src[searchPos:])
//line /usr/local/go/src/regexp/regexp.go:651
			// _ = "end of CoverTab[65185]"
		}
//line /usr/local/go/src/regexp/regexp.go:652
		// _ = "end of CoverTab[65175]"
//line /usr/local/go/src/regexp/regexp.go:652
		_go_fuzz_dep_.CoverTab[65176]++
							if searchPos+width > a[1] {
//line /usr/local/go/src/regexp/regexp.go:653
			_go_fuzz_dep_.CoverTab[65186]++
								searchPos += width
//line /usr/local/go/src/regexp/regexp.go:654
			// _ = "end of CoverTab[65186]"
		} else {
//line /usr/local/go/src/regexp/regexp.go:655
			_go_fuzz_dep_.CoverTab[65187]++
//line /usr/local/go/src/regexp/regexp.go:655
			if searchPos+1 > a[1] {
//line /usr/local/go/src/regexp/regexp.go:655
				_go_fuzz_dep_.CoverTab[65188]++

//line /usr/local/go/src/regexp/regexp.go:658
				searchPos++
//line /usr/local/go/src/regexp/regexp.go:658
				// _ = "end of CoverTab[65188]"
			} else {
//line /usr/local/go/src/regexp/regexp.go:659
				_go_fuzz_dep_.CoverTab[65189]++
									searchPos = a[1]
//line /usr/local/go/src/regexp/regexp.go:660
				// _ = "end of CoverTab[65189]"
			}
//line /usr/local/go/src/regexp/regexp.go:661
			// _ = "end of CoverTab[65187]"
//line /usr/local/go/src/regexp/regexp.go:661
		}
//line /usr/local/go/src/regexp/regexp.go:661
		// _ = "end of CoverTab[65176]"
	}
//line /usr/local/go/src/regexp/regexp.go:662
	// _ = "end of CoverTab[65165]"
//line /usr/local/go/src/regexp/regexp.go:662
	_go_fuzz_dep_.CoverTab[65166]++

//line /usr/local/go/src/regexp/regexp.go:665
	if bsrc != nil {
//line /usr/local/go/src/regexp/regexp.go:665
		_go_fuzz_dep_.CoverTab[65190]++
							buf = append(buf, bsrc[lastMatchEnd:]...)
//line /usr/local/go/src/regexp/regexp.go:666
		// _ = "end of CoverTab[65190]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:667
		_go_fuzz_dep_.CoverTab[65191]++
							buf = append(buf, src[lastMatchEnd:]...)
//line /usr/local/go/src/regexp/regexp.go:668
		// _ = "end of CoverTab[65191]"
	}
//line /usr/local/go/src/regexp/regexp.go:669
	// _ = "end of CoverTab[65166]"
//line /usr/local/go/src/regexp/regexp.go:669
	_go_fuzz_dep_.CoverTab[65167]++

						return buf
//line /usr/local/go/src/regexp/regexp.go:671
	// _ = "end of CoverTab[65167]"
}

// ReplaceAll returns a copy of src, replacing matches of the Regexp
//line /usr/local/go/src/regexp/regexp.go:674
// with the replacement text repl. Inside repl, $ signs are interpreted as
//line /usr/local/go/src/regexp/regexp.go:674
// in Expand, so for instance $1 represents the text of the first submatch.
//line /usr/local/go/src/regexp/regexp.go:677
func (re *Regexp) ReplaceAll(src, repl []byte) []byte {
//line /usr/local/go/src/regexp/regexp.go:677
	_go_fuzz_dep_.CoverTab[65192]++
						n := 2
						if bytes.IndexByte(repl, '$') >= 0 {
//line /usr/local/go/src/regexp/regexp.go:679
		_go_fuzz_dep_.CoverTab[65195]++
							n = 2 * (re.numSubexp + 1)
//line /usr/local/go/src/regexp/regexp.go:680
		// _ = "end of CoverTab[65195]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:681
		_go_fuzz_dep_.CoverTab[65196]++
//line /usr/local/go/src/regexp/regexp.go:681
		// _ = "end of CoverTab[65196]"
//line /usr/local/go/src/regexp/regexp.go:681
	}
//line /usr/local/go/src/regexp/regexp.go:681
	// _ = "end of CoverTab[65192]"
//line /usr/local/go/src/regexp/regexp.go:681
	_go_fuzz_dep_.CoverTab[65193]++
						srepl := ""
						b := re.replaceAll(src, "", n, func(dst []byte, match []int) []byte {
//line /usr/local/go/src/regexp/regexp.go:683
		_go_fuzz_dep_.CoverTab[65197]++
							if len(srepl) != len(repl) {
//line /usr/local/go/src/regexp/regexp.go:684
			_go_fuzz_dep_.CoverTab[65199]++
								srepl = string(repl)
//line /usr/local/go/src/regexp/regexp.go:685
			// _ = "end of CoverTab[65199]"
		} else {
//line /usr/local/go/src/regexp/regexp.go:686
			_go_fuzz_dep_.CoverTab[65200]++
//line /usr/local/go/src/regexp/regexp.go:686
			// _ = "end of CoverTab[65200]"
//line /usr/local/go/src/regexp/regexp.go:686
		}
//line /usr/local/go/src/regexp/regexp.go:686
		// _ = "end of CoverTab[65197]"
//line /usr/local/go/src/regexp/regexp.go:686
		_go_fuzz_dep_.CoverTab[65198]++
							return re.expand(dst, srepl, src, "", match)
//line /usr/local/go/src/regexp/regexp.go:687
		// _ = "end of CoverTab[65198]"
	})
//line /usr/local/go/src/regexp/regexp.go:688
	// _ = "end of CoverTab[65193]"
//line /usr/local/go/src/regexp/regexp.go:688
	_go_fuzz_dep_.CoverTab[65194]++
						return b
//line /usr/local/go/src/regexp/regexp.go:689
	// _ = "end of CoverTab[65194]"
}

// ReplaceAllLiteral returns a copy of src, replacing matches of the Regexp
//line /usr/local/go/src/regexp/regexp.go:692
// with the replacement bytes repl. The replacement repl is substituted directly,
//line /usr/local/go/src/regexp/regexp.go:692
// without using Expand.
//line /usr/local/go/src/regexp/regexp.go:695
func (re *Regexp) ReplaceAllLiteral(src, repl []byte) []byte {
//line /usr/local/go/src/regexp/regexp.go:695
	_go_fuzz_dep_.CoverTab[65201]++
						return re.replaceAll(src, "", 2, func(dst []byte, match []int) []byte {
//line /usr/local/go/src/regexp/regexp.go:696
		_go_fuzz_dep_.CoverTab[65202]++
							return append(dst, repl...)
//line /usr/local/go/src/regexp/regexp.go:697
		// _ = "end of CoverTab[65202]"
	})
//line /usr/local/go/src/regexp/regexp.go:698
	// _ = "end of CoverTab[65201]"
}

// ReplaceAllFunc returns a copy of src in which all matches of the
//line /usr/local/go/src/regexp/regexp.go:701
// Regexp have been replaced by the return value of function repl applied
//line /usr/local/go/src/regexp/regexp.go:701
// to the matched byte slice. The replacement returned by repl is substituted
//line /usr/local/go/src/regexp/regexp.go:701
// directly, without using Expand.
//line /usr/local/go/src/regexp/regexp.go:705
func (re *Regexp) ReplaceAllFunc(src []byte, repl func([]byte) []byte) []byte {
//line /usr/local/go/src/regexp/regexp.go:705
	_go_fuzz_dep_.CoverTab[65203]++
						return re.replaceAll(src, "", 2, func(dst []byte, match []int) []byte {
//line /usr/local/go/src/regexp/regexp.go:706
		_go_fuzz_dep_.CoverTab[65204]++
							return append(dst, repl(src[match[0]:match[1]])...)
//line /usr/local/go/src/regexp/regexp.go:707
		// _ = "end of CoverTab[65204]"
	})
//line /usr/local/go/src/regexp/regexp.go:708
	// _ = "end of CoverTab[65203]"
}

// Bitmap used by func special to check whether a character needs to be escaped.
var specialBytes [16]byte

// special reports whether byte b needs to be escaped by QuoteMeta.
func special(b byte) bool {
//line /usr/local/go/src/regexp/regexp.go:715
	_go_fuzz_dep_.CoverTab[65205]++
						return b < utf8.RuneSelf && func() bool {
//line /usr/local/go/src/regexp/regexp.go:716
		_go_fuzz_dep_.CoverTab[65206]++
//line /usr/local/go/src/regexp/regexp.go:716
		return specialBytes[b%16]&(1<<(b/16)) != 0
//line /usr/local/go/src/regexp/regexp.go:716
		// _ = "end of CoverTab[65206]"
//line /usr/local/go/src/regexp/regexp.go:716
	}()
//line /usr/local/go/src/regexp/regexp.go:716
	// _ = "end of CoverTab[65205]"
}

func init() {
	for _, b := range []byte(`\.+*?()|[]{}^$`) {
		specialBytes[b%16] |= 1 << (b / 16)
	}
}

// QuoteMeta returns a string that escapes all regular expression metacharacters
//line /usr/local/go/src/regexp/regexp.go:725
// inside the argument text; the returned string is a regular expression matching
//line /usr/local/go/src/regexp/regexp.go:725
// the literal text.
//line /usr/local/go/src/regexp/regexp.go:728
func QuoteMeta(s string) string {
//line /usr/local/go/src/regexp/regexp.go:728
	_go_fuzz_dep_.CoverTab[65207]++
	// A byte loop is correct because all metacharacters are ASCII.
	var i int
	for i = 0; i < len(s); i++ {
//line /usr/local/go/src/regexp/regexp.go:731
		_go_fuzz_dep_.CoverTab[65211]++
							if special(s[i]) {
//line /usr/local/go/src/regexp/regexp.go:732
			_go_fuzz_dep_.CoverTab[65212]++
								break
//line /usr/local/go/src/regexp/regexp.go:733
			// _ = "end of CoverTab[65212]"
		} else {
//line /usr/local/go/src/regexp/regexp.go:734
			_go_fuzz_dep_.CoverTab[65213]++
//line /usr/local/go/src/regexp/regexp.go:734
			// _ = "end of CoverTab[65213]"
//line /usr/local/go/src/regexp/regexp.go:734
		}
//line /usr/local/go/src/regexp/regexp.go:734
		// _ = "end of CoverTab[65211]"
	}
//line /usr/local/go/src/regexp/regexp.go:735
	// _ = "end of CoverTab[65207]"
//line /usr/local/go/src/regexp/regexp.go:735
	_go_fuzz_dep_.CoverTab[65208]++

						if i >= len(s) {
//line /usr/local/go/src/regexp/regexp.go:737
		_go_fuzz_dep_.CoverTab[65214]++
							return s
//line /usr/local/go/src/regexp/regexp.go:738
		// _ = "end of CoverTab[65214]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:739
		_go_fuzz_dep_.CoverTab[65215]++
//line /usr/local/go/src/regexp/regexp.go:739
		// _ = "end of CoverTab[65215]"
//line /usr/local/go/src/regexp/regexp.go:739
	}
//line /usr/local/go/src/regexp/regexp.go:739
	// _ = "end of CoverTab[65208]"
//line /usr/local/go/src/regexp/regexp.go:739
	_go_fuzz_dep_.CoverTab[65209]++

						b := make([]byte, 2*len(s)-i)
						copy(b, s[:i])
						j := i
						for ; i < len(s); i++ {
//line /usr/local/go/src/regexp/regexp.go:744
		_go_fuzz_dep_.CoverTab[65216]++
							if special(s[i]) {
//line /usr/local/go/src/regexp/regexp.go:745
			_go_fuzz_dep_.CoverTab[65218]++
								b[j] = '\\'
								j++
//line /usr/local/go/src/regexp/regexp.go:747
			// _ = "end of CoverTab[65218]"
		} else {
//line /usr/local/go/src/regexp/regexp.go:748
			_go_fuzz_dep_.CoverTab[65219]++
//line /usr/local/go/src/regexp/regexp.go:748
			// _ = "end of CoverTab[65219]"
//line /usr/local/go/src/regexp/regexp.go:748
		}
//line /usr/local/go/src/regexp/regexp.go:748
		// _ = "end of CoverTab[65216]"
//line /usr/local/go/src/regexp/regexp.go:748
		_go_fuzz_dep_.CoverTab[65217]++
							b[j] = s[i]
							j++
//line /usr/local/go/src/regexp/regexp.go:750
		// _ = "end of CoverTab[65217]"
	}
//line /usr/local/go/src/regexp/regexp.go:751
	// _ = "end of CoverTab[65209]"
//line /usr/local/go/src/regexp/regexp.go:751
	_go_fuzz_dep_.CoverTab[65210]++
						return string(b[:j])
//line /usr/local/go/src/regexp/regexp.go:752
	// _ = "end of CoverTab[65210]"
}

// The number of capture values in the program may correspond
//line /usr/local/go/src/regexp/regexp.go:755
// to fewer capturing expressions than are in the regexp.
//line /usr/local/go/src/regexp/regexp.go:755
// For example, "(a){0}" turns into an empty program, so the
//line /usr/local/go/src/regexp/regexp.go:755
// maximum capture in the program is 0 but we need to return
//line /usr/local/go/src/regexp/regexp.go:755
// an expression for \1.  Pad appends -1s to the slice a as needed.
//line /usr/local/go/src/regexp/regexp.go:760
func (re *Regexp) pad(a []int) []int {
//line /usr/local/go/src/regexp/regexp.go:760
	_go_fuzz_dep_.CoverTab[65220]++
						if a == nil {
//line /usr/local/go/src/regexp/regexp.go:761
		_go_fuzz_dep_.CoverTab[65223]++

							return nil
//line /usr/local/go/src/regexp/regexp.go:763
		// _ = "end of CoverTab[65223]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:764
		_go_fuzz_dep_.CoverTab[65224]++
//line /usr/local/go/src/regexp/regexp.go:764
		// _ = "end of CoverTab[65224]"
//line /usr/local/go/src/regexp/regexp.go:764
	}
//line /usr/local/go/src/regexp/regexp.go:764
	// _ = "end of CoverTab[65220]"
//line /usr/local/go/src/regexp/regexp.go:764
	_go_fuzz_dep_.CoverTab[65221]++
						n := (1 + re.numSubexp) * 2
						for len(a) < n {
//line /usr/local/go/src/regexp/regexp.go:766
		_go_fuzz_dep_.CoverTab[65225]++
							a = append(a, -1)
//line /usr/local/go/src/regexp/regexp.go:767
		// _ = "end of CoverTab[65225]"
	}
//line /usr/local/go/src/regexp/regexp.go:768
	// _ = "end of CoverTab[65221]"
//line /usr/local/go/src/regexp/regexp.go:768
	_go_fuzz_dep_.CoverTab[65222]++
						return a
//line /usr/local/go/src/regexp/regexp.go:769
	// _ = "end of CoverTab[65222]"
}

// allMatches calls deliver at most n times
//line /usr/local/go/src/regexp/regexp.go:772
// with the location of successive matches in the input text.
//line /usr/local/go/src/regexp/regexp.go:772
// The input text is b if non-nil, otherwise s.
//line /usr/local/go/src/regexp/regexp.go:775
func (re *Regexp) allMatches(s string, b []byte, n int, deliver func([]int)) {
//line /usr/local/go/src/regexp/regexp.go:775
	_go_fuzz_dep_.CoverTab[65226]++
						var end int
						if b == nil {
//line /usr/local/go/src/regexp/regexp.go:777
		_go_fuzz_dep_.CoverTab[65228]++
							end = len(s)
//line /usr/local/go/src/regexp/regexp.go:778
		// _ = "end of CoverTab[65228]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:779
		_go_fuzz_dep_.CoverTab[65229]++
							end = len(b)
//line /usr/local/go/src/regexp/regexp.go:780
		// _ = "end of CoverTab[65229]"
	}
//line /usr/local/go/src/regexp/regexp.go:781
	// _ = "end of CoverTab[65226]"
//line /usr/local/go/src/regexp/regexp.go:781
	_go_fuzz_dep_.CoverTab[65227]++

						for pos, i, prevMatchEnd := 0, 0, -1; i < n && func() bool {
//line /usr/local/go/src/regexp/regexp.go:783
		_go_fuzz_dep_.CoverTab[65230]++
//line /usr/local/go/src/regexp/regexp.go:783
		return pos <= end
//line /usr/local/go/src/regexp/regexp.go:783
		// _ = "end of CoverTab[65230]"
//line /usr/local/go/src/regexp/regexp.go:783
	}(); {
//line /usr/local/go/src/regexp/regexp.go:783
		_go_fuzz_dep_.CoverTab[65231]++
							matches := re.doExecute(nil, b, s, pos, re.prog.NumCap, nil)
							if len(matches) == 0 {
//line /usr/local/go/src/regexp/regexp.go:785
			_go_fuzz_dep_.CoverTab[65234]++
								break
//line /usr/local/go/src/regexp/regexp.go:786
			// _ = "end of CoverTab[65234]"
		} else {
//line /usr/local/go/src/regexp/regexp.go:787
			_go_fuzz_dep_.CoverTab[65235]++
//line /usr/local/go/src/regexp/regexp.go:787
			// _ = "end of CoverTab[65235]"
//line /usr/local/go/src/regexp/regexp.go:787
		}
//line /usr/local/go/src/regexp/regexp.go:787
		// _ = "end of CoverTab[65231]"
//line /usr/local/go/src/regexp/regexp.go:787
		_go_fuzz_dep_.CoverTab[65232]++

							accept := true
							if matches[1] == pos {
//line /usr/local/go/src/regexp/regexp.go:790
			_go_fuzz_dep_.CoverTab[65236]++

								if matches[0] == prevMatchEnd {
//line /usr/local/go/src/regexp/regexp.go:792
				_go_fuzz_dep_.CoverTab[65239]++

//line /usr/local/go/src/regexp/regexp.go:795
				accept = false
//line /usr/local/go/src/regexp/regexp.go:795
				// _ = "end of CoverTab[65239]"
			} else {
//line /usr/local/go/src/regexp/regexp.go:796
				_go_fuzz_dep_.CoverTab[65240]++
//line /usr/local/go/src/regexp/regexp.go:796
				// _ = "end of CoverTab[65240]"
//line /usr/local/go/src/regexp/regexp.go:796
			}
//line /usr/local/go/src/regexp/regexp.go:796
			// _ = "end of CoverTab[65236]"
//line /usr/local/go/src/regexp/regexp.go:796
			_go_fuzz_dep_.CoverTab[65237]++
								var width int
								if b == nil {
//line /usr/local/go/src/regexp/regexp.go:798
				_go_fuzz_dep_.CoverTab[65241]++
									is := inputString{str: s}
									_, width = is.step(pos)
//line /usr/local/go/src/regexp/regexp.go:800
				// _ = "end of CoverTab[65241]"
			} else {
//line /usr/local/go/src/regexp/regexp.go:801
				_go_fuzz_dep_.CoverTab[65242]++
									ib := inputBytes{str: b}
									_, width = ib.step(pos)
//line /usr/local/go/src/regexp/regexp.go:803
				// _ = "end of CoverTab[65242]"
			}
//line /usr/local/go/src/regexp/regexp.go:804
			// _ = "end of CoverTab[65237]"
//line /usr/local/go/src/regexp/regexp.go:804
			_go_fuzz_dep_.CoverTab[65238]++
								if width > 0 {
//line /usr/local/go/src/regexp/regexp.go:805
				_go_fuzz_dep_.CoverTab[65243]++
									pos += width
//line /usr/local/go/src/regexp/regexp.go:806
				// _ = "end of CoverTab[65243]"
			} else {
//line /usr/local/go/src/regexp/regexp.go:807
				_go_fuzz_dep_.CoverTab[65244]++
									pos = end + 1
//line /usr/local/go/src/regexp/regexp.go:808
				// _ = "end of CoverTab[65244]"
			}
//line /usr/local/go/src/regexp/regexp.go:809
			// _ = "end of CoverTab[65238]"
		} else {
//line /usr/local/go/src/regexp/regexp.go:810
			_go_fuzz_dep_.CoverTab[65245]++
								pos = matches[1]
//line /usr/local/go/src/regexp/regexp.go:811
			// _ = "end of CoverTab[65245]"
		}
//line /usr/local/go/src/regexp/regexp.go:812
		// _ = "end of CoverTab[65232]"
//line /usr/local/go/src/regexp/regexp.go:812
		_go_fuzz_dep_.CoverTab[65233]++
							prevMatchEnd = matches[1]

							if accept {
//line /usr/local/go/src/regexp/regexp.go:815
			_go_fuzz_dep_.CoverTab[65246]++
								deliver(re.pad(matches))
								i++
//line /usr/local/go/src/regexp/regexp.go:817
			// _ = "end of CoverTab[65246]"
		} else {
//line /usr/local/go/src/regexp/regexp.go:818
			_go_fuzz_dep_.CoverTab[65247]++
//line /usr/local/go/src/regexp/regexp.go:818
			// _ = "end of CoverTab[65247]"
//line /usr/local/go/src/regexp/regexp.go:818
		}
//line /usr/local/go/src/regexp/regexp.go:818
		// _ = "end of CoverTab[65233]"
	}
//line /usr/local/go/src/regexp/regexp.go:819
	// _ = "end of CoverTab[65227]"
}

// Find returns a slice holding the text of the leftmost match in b of the regular expression.
//line /usr/local/go/src/regexp/regexp.go:822
// A return value of nil indicates no match.
//line /usr/local/go/src/regexp/regexp.go:824
func (re *Regexp) Find(b []byte) []byte {
//line /usr/local/go/src/regexp/regexp.go:824
	_go_fuzz_dep_.CoverTab[65248]++
						var dstCap [2]int
						a := re.doExecute(nil, b, "", 0, 2, dstCap[:0])
						if a == nil {
//line /usr/local/go/src/regexp/regexp.go:827
		_go_fuzz_dep_.CoverTab[65250]++
							return nil
//line /usr/local/go/src/regexp/regexp.go:828
		// _ = "end of CoverTab[65250]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:829
		_go_fuzz_dep_.CoverTab[65251]++
//line /usr/local/go/src/regexp/regexp.go:829
		// _ = "end of CoverTab[65251]"
//line /usr/local/go/src/regexp/regexp.go:829
	}
//line /usr/local/go/src/regexp/regexp.go:829
	// _ = "end of CoverTab[65248]"
//line /usr/local/go/src/regexp/regexp.go:829
	_go_fuzz_dep_.CoverTab[65249]++
						return b[a[0]:a[1]:a[1]]
//line /usr/local/go/src/regexp/regexp.go:830
	// _ = "end of CoverTab[65249]"
}

// FindIndex returns a two-element slice of integers defining the location of
//line /usr/local/go/src/regexp/regexp.go:833
// the leftmost match in b of the regular expression. The match itself is at
//line /usr/local/go/src/regexp/regexp.go:833
// b[loc[0]:loc[1]].
//line /usr/local/go/src/regexp/regexp.go:833
// A return value of nil indicates no match.
//line /usr/local/go/src/regexp/regexp.go:837
func (re *Regexp) FindIndex(b []byte) (loc []int) {
//line /usr/local/go/src/regexp/regexp.go:837
	_go_fuzz_dep_.CoverTab[65252]++
						a := re.doExecute(nil, b, "", 0, 2, nil)
						if a == nil {
//line /usr/local/go/src/regexp/regexp.go:839
		_go_fuzz_dep_.CoverTab[65254]++
							return nil
//line /usr/local/go/src/regexp/regexp.go:840
		// _ = "end of CoverTab[65254]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:841
		_go_fuzz_dep_.CoverTab[65255]++
//line /usr/local/go/src/regexp/regexp.go:841
		// _ = "end of CoverTab[65255]"
//line /usr/local/go/src/regexp/regexp.go:841
	}
//line /usr/local/go/src/regexp/regexp.go:841
	// _ = "end of CoverTab[65252]"
//line /usr/local/go/src/regexp/regexp.go:841
	_go_fuzz_dep_.CoverTab[65253]++
						return a[0:2]
//line /usr/local/go/src/regexp/regexp.go:842
	// _ = "end of CoverTab[65253]"
}

// FindString returns a string holding the text of the leftmost match in s of the regular
//line /usr/local/go/src/regexp/regexp.go:845
// expression. If there is no match, the return value is an empty string,
//line /usr/local/go/src/regexp/regexp.go:845
// but it will also be empty if the regular expression successfully matches
//line /usr/local/go/src/regexp/regexp.go:845
// an empty string. Use FindStringIndex or FindStringSubmatch if it is
//line /usr/local/go/src/regexp/regexp.go:845
// necessary to distinguish these cases.
//line /usr/local/go/src/regexp/regexp.go:850
func (re *Regexp) FindString(s string) string {
//line /usr/local/go/src/regexp/regexp.go:850
	_go_fuzz_dep_.CoverTab[65256]++
						var dstCap [2]int
						a := re.doExecute(nil, nil, s, 0, 2, dstCap[:0])
						if a == nil {
//line /usr/local/go/src/regexp/regexp.go:853
		_go_fuzz_dep_.CoverTab[65258]++
							return ""
//line /usr/local/go/src/regexp/regexp.go:854
		// _ = "end of CoverTab[65258]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:855
		_go_fuzz_dep_.CoverTab[65259]++
//line /usr/local/go/src/regexp/regexp.go:855
		// _ = "end of CoverTab[65259]"
//line /usr/local/go/src/regexp/regexp.go:855
	}
//line /usr/local/go/src/regexp/regexp.go:855
	// _ = "end of CoverTab[65256]"
//line /usr/local/go/src/regexp/regexp.go:855
	_go_fuzz_dep_.CoverTab[65257]++
						return s[a[0]:a[1]]
//line /usr/local/go/src/regexp/regexp.go:856
	// _ = "end of CoverTab[65257]"
}

// FindStringIndex returns a two-element slice of integers defining the
//line /usr/local/go/src/regexp/regexp.go:859
// location of the leftmost match in s of the regular expression. The match
//line /usr/local/go/src/regexp/regexp.go:859
// itself is at s[loc[0]:loc[1]].
//line /usr/local/go/src/regexp/regexp.go:859
// A return value of nil indicates no match.
//line /usr/local/go/src/regexp/regexp.go:863
func (re *Regexp) FindStringIndex(s string) (loc []int) {
//line /usr/local/go/src/regexp/regexp.go:863
	_go_fuzz_dep_.CoverTab[65260]++
						a := re.doExecute(nil, nil, s, 0, 2, nil)
						if a == nil {
//line /usr/local/go/src/regexp/regexp.go:865
		_go_fuzz_dep_.CoverTab[65262]++
							return nil
//line /usr/local/go/src/regexp/regexp.go:866
		// _ = "end of CoverTab[65262]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:867
		_go_fuzz_dep_.CoverTab[65263]++
//line /usr/local/go/src/regexp/regexp.go:867
		// _ = "end of CoverTab[65263]"
//line /usr/local/go/src/regexp/regexp.go:867
	}
//line /usr/local/go/src/regexp/regexp.go:867
	// _ = "end of CoverTab[65260]"
//line /usr/local/go/src/regexp/regexp.go:867
	_go_fuzz_dep_.CoverTab[65261]++
						return a[0:2]
//line /usr/local/go/src/regexp/regexp.go:868
	// _ = "end of CoverTab[65261]"
}

// FindReaderIndex returns a two-element slice of integers defining the
//line /usr/local/go/src/regexp/regexp.go:871
// location of the leftmost match of the regular expression in text read from
//line /usr/local/go/src/regexp/regexp.go:871
// the RuneReader. The match text was found in the input stream at
//line /usr/local/go/src/regexp/regexp.go:871
// byte offset loc[0] through loc[1]-1.
//line /usr/local/go/src/regexp/regexp.go:871
// A return value of nil indicates no match.
//line /usr/local/go/src/regexp/regexp.go:876
func (re *Regexp) FindReaderIndex(r io.RuneReader) (loc []int) {
//line /usr/local/go/src/regexp/regexp.go:876
	_go_fuzz_dep_.CoverTab[65264]++
						a := re.doExecute(r, nil, "", 0, 2, nil)
						if a == nil {
//line /usr/local/go/src/regexp/regexp.go:878
		_go_fuzz_dep_.CoverTab[65266]++
							return nil
//line /usr/local/go/src/regexp/regexp.go:879
		// _ = "end of CoverTab[65266]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:880
		_go_fuzz_dep_.CoverTab[65267]++
//line /usr/local/go/src/regexp/regexp.go:880
		// _ = "end of CoverTab[65267]"
//line /usr/local/go/src/regexp/regexp.go:880
	}
//line /usr/local/go/src/regexp/regexp.go:880
	// _ = "end of CoverTab[65264]"
//line /usr/local/go/src/regexp/regexp.go:880
	_go_fuzz_dep_.CoverTab[65265]++
						return a[0:2]
//line /usr/local/go/src/regexp/regexp.go:881
	// _ = "end of CoverTab[65265]"
}

// FindSubmatch returns a slice of slices holding the text of the leftmost
//line /usr/local/go/src/regexp/regexp.go:884
// match of the regular expression in b and the matches, if any, of its
//line /usr/local/go/src/regexp/regexp.go:884
// subexpressions, as defined by the 'Submatch' descriptions in the package
//line /usr/local/go/src/regexp/regexp.go:884
// comment.
//line /usr/local/go/src/regexp/regexp.go:884
// A return value of nil indicates no match.
//line /usr/local/go/src/regexp/regexp.go:889
func (re *Regexp) FindSubmatch(b []byte) [][]byte {
//line /usr/local/go/src/regexp/regexp.go:889
	_go_fuzz_dep_.CoverTab[65268]++
						var dstCap [4]int
						a := re.doExecute(nil, b, "", 0, re.prog.NumCap, dstCap[:0])
						if a == nil {
//line /usr/local/go/src/regexp/regexp.go:892
		_go_fuzz_dep_.CoverTab[65271]++
							return nil
//line /usr/local/go/src/regexp/regexp.go:893
		// _ = "end of CoverTab[65271]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:894
		_go_fuzz_dep_.CoverTab[65272]++
//line /usr/local/go/src/regexp/regexp.go:894
		// _ = "end of CoverTab[65272]"
//line /usr/local/go/src/regexp/regexp.go:894
	}
//line /usr/local/go/src/regexp/regexp.go:894
	// _ = "end of CoverTab[65268]"
//line /usr/local/go/src/regexp/regexp.go:894
	_go_fuzz_dep_.CoverTab[65269]++
						ret := make([][]byte, 1+re.numSubexp)
						for i := range ret {
//line /usr/local/go/src/regexp/regexp.go:896
		_go_fuzz_dep_.CoverTab[65273]++
							if 2*i < len(a) && func() bool {
//line /usr/local/go/src/regexp/regexp.go:897
			_go_fuzz_dep_.CoverTab[65274]++
//line /usr/local/go/src/regexp/regexp.go:897
			return a[2*i] >= 0
//line /usr/local/go/src/regexp/regexp.go:897
			// _ = "end of CoverTab[65274]"
//line /usr/local/go/src/regexp/regexp.go:897
		}() {
//line /usr/local/go/src/regexp/regexp.go:897
			_go_fuzz_dep_.CoverTab[65275]++
								ret[i] = b[a[2*i]:a[2*i+1]:a[2*i+1]]
//line /usr/local/go/src/regexp/regexp.go:898
			// _ = "end of CoverTab[65275]"
		} else {
//line /usr/local/go/src/regexp/regexp.go:899
			_go_fuzz_dep_.CoverTab[65276]++
//line /usr/local/go/src/regexp/regexp.go:899
			// _ = "end of CoverTab[65276]"
//line /usr/local/go/src/regexp/regexp.go:899
		}
//line /usr/local/go/src/regexp/regexp.go:899
		// _ = "end of CoverTab[65273]"
	}
//line /usr/local/go/src/regexp/regexp.go:900
	// _ = "end of CoverTab[65269]"
//line /usr/local/go/src/regexp/regexp.go:900
	_go_fuzz_dep_.CoverTab[65270]++
						return ret
//line /usr/local/go/src/regexp/regexp.go:901
	// _ = "end of CoverTab[65270]"
}

// Expand appends template to dst and returns the result; during the
//line /usr/local/go/src/regexp/regexp.go:904
// append, Expand replaces variables in the template with corresponding
//line /usr/local/go/src/regexp/regexp.go:904
// matches drawn from src. The match slice should have been returned by
//line /usr/local/go/src/regexp/regexp.go:904
// FindSubmatchIndex.
//line /usr/local/go/src/regexp/regexp.go:904
//
//line /usr/local/go/src/regexp/regexp.go:904
// In the template, a variable is denoted by a substring of the form
//line /usr/local/go/src/regexp/regexp.go:904
// $name or ${name}, where name is a non-empty sequence of letters,
//line /usr/local/go/src/regexp/regexp.go:904
// digits, and underscores. A purely numeric name like $1 refers to
//line /usr/local/go/src/regexp/regexp.go:904
// the submatch with the corresponding index; other names refer to
//line /usr/local/go/src/regexp/regexp.go:904
// capturing parentheses named with the (?P<name>...) syntax. A
//line /usr/local/go/src/regexp/regexp.go:904
// reference to an out of range or unmatched index or a name that is not
//line /usr/local/go/src/regexp/regexp.go:904
// present in the regular expression is replaced with an empty slice.
//line /usr/local/go/src/regexp/regexp.go:904
//
//line /usr/local/go/src/regexp/regexp.go:904
// In the $name form, name is taken to be as long as possible: $1x is
//line /usr/local/go/src/regexp/regexp.go:904
// equivalent to ${1x}, not ${1}x, and, $10 is equivalent to ${10}, not ${1}0.
//line /usr/local/go/src/regexp/regexp.go:904
//
//line /usr/local/go/src/regexp/regexp.go:904
// To insert a literal $ in the output, use $$ in the template.
//line /usr/local/go/src/regexp/regexp.go:921
func (re *Regexp) Expand(dst []byte, template []byte, src []byte, match []int) []byte {
//line /usr/local/go/src/regexp/regexp.go:921
	_go_fuzz_dep_.CoverTab[65277]++
						return re.expand(dst, string(template), src, "", match)
//line /usr/local/go/src/regexp/regexp.go:922
	// _ = "end of CoverTab[65277]"
}

// ExpandString is like Expand but the template and source are strings.
//line /usr/local/go/src/regexp/regexp.go:925
// It appends to and returns a byte slice in order to give the calling
//line /usr/local/go/src/regexp/regexp.go:925
// code control over allocation.
//line /usr/local/go/src/regexp/regexp.go:928
func (re *Regexp) ExpandString(dst []byte, template string, src string, match []int) []byte {
//line /usr/local/go/src/regexp/regexp.go:928
	_go_fuzz_dep_.CoverTab[65278]++
						return re.expand(dst, template, nil, src, match)
//line /usr/local/go/src/regexp/regexp.go:929
	// _ = "end of CoverTab[65278]"
}

func (re *Regexp) expand(dst []byte, template string, bsrc []byte, src string, match []int) []byte {
//line /usr/local/go/src/regexp/regexp.go:932
	_go_fuzz_dep_.CoverTab[65279]++
						for len(template) > 0 {
//line /usr/local/go/src/regexp/regexp.go:933
		_go_fuzz_dep_.CoverTab[65281]++
							before, after, ok := strings.Cut(template, "$")
							if !ok {
//line /usr/local/go/src/regexp/regexp.go:935
			_go_fuzz_dep_.CoverTab[65285]++
								break
//line /usr/local/go/src/regexp/regexp.go:936
			// _ = "end of CoverTab[65285]"
		} else {
//line /usr/local/go/src/regexp/regexp.go:937
			_go_fuzz_dep_.CoverTab[65286]++
//line /usr/local/go/src/regexp/regexp.go:937
			// _ = "end of CoverTab[65286]"
//line /usr/local/go/src/regexp/regexp.go:937
		}
//line /usr/local/go/src/regexp/regexp.go:937
		// _ = "end of CoverTab[65281]"
//line /usr/local/go/src/regexp/regexp.go:937
		_go_fuzz_dep_.CoverTab[65282]++
							dst = append(dst, before...)
							template = after
							if template != "" && func() bool {
//line /usr/local/go/src/regexp/regexp.go:940
			_go_fuzz_dep_.CoverTab[65287]++
//line /usr/local/go/src/regexp/regexp.go:940
			return template[0] == '$'
//line /usr/local/go/src/regexp/regexp.go:940
			// _ = "end of CoverTab[65287]"
//line /usr/local/go/src/regexp/regexp.go:940
		}() {
//line /usr/local/go/src/regexp/regexp.go:940
			_go_fuzz_dep_.CoverTab[65288]++

								dst = append(dst, '$')
								template = template[1:]
								continue
//line /usr/local/go/src/regexp/regexp.go:944
			// _ = "end of CoverTab[65288]"
		} else {
//line /usr/local/go/src/regexp/regexp.go:945
			_go_fuzz_dep_.CoverTab[65289]++
//line /usr/local/go/src/regexp/regexp.go:945
			// _ = "end of CoverTab[65289]"
//line /usr/local/go/src/regexp/regexp.go:945
		}
//line /usr/local/go/src/regexp/regexp.go:945
		// _ = "end of CoverTab[65282]"
//line /usr/local/go/src/regexp/regexp.go:945
		_go_fuzz_dep_.CoverTab[65283]++
							name, num, rest, ok := extract(template)
							if !ok {
//line /usr/local/go/src/regexp/regexp.go:947
			_go_fuzz_dep_.CoverTab[65290]++

								dst = append(dst, '$')
								continue
//line /usr/local/go/src/regexp/regexp.go:950
			// _ = "end of CoverTab[65290]"
		} else {
//line /usr/local/go/src/regexp/regexp.go:951
			_go_fuzz_dep_.CoverTab[65291]++
//line /usr/local/go/src/regexp/regexp.go:951
			// _ = "end of CoverTab[65291]"
//line /usr/local/go/src/regexp/regexp.go:951
		}
//line /usr/local/go/src/regexp/regexp.go:951
		// _ = "end of CoverTab[65283]"
//line /usr/local/go/src/regexp/regexp.go:951
		_go_fuzz_dep_.CoverTab[65284]++
							template = rest
							if num >= 0 {
//line /usr/local/go/src/regexp/regexp.go:953
			_go_fuzz_dep_.CoverTab[65292]++
								if 2*num+1 < len(match) && func() bool {
//line /usr/local/go/src/regexp/regexp.go:954
				_go_fuzz_dep_.CoverTab[65293]++
//line /usr/local/go/src/regexp/regexp.go:954
				return match[2*num] >= 0
//line /usr/local/go/src/regexp/regexp.go:954
				// _ = "end of CoverTab[65293]"
//line /usr/local/go/src/regexp/regexp.go:954
			}() {
//line /usr/local/go/src/regexp/regexp.go:954
				_go_fuzz_dep_.CoverTab[65294]++
									if bsrc != nil {
//line /usr/local/go/src/regexp/regexp.go:955
					_go_fuzz_dep_.CoverTab[65295]++
										dst = append(dst, bsrc[match[2*num]:match[2*num+1]]...)
//line /usr/local/go/src/regexp/regexp.go:956
					// _ = "end of CoverTab[65295]"
				} else {
//line /usr/local/go/src/regexp/regexp.go:957
					_go_fuzz_dep_.CoverTab[65296]++
										dst = append(dst, src[match[2*num]:match[2*num+1]]...)
//line /usr/local/go/src/regexp/regexp.go:958
					// _ = "end of CoverTab[65296]"
				}
//line /usr/local/go/src/regexp/regexp.go:959
				// _ = "end of CoverTab[65294]"
			} else {
//line /usr/local/go/src/regexp/regexp.go:960
				_go_fuzz_dep_.CoverTab[65297]++
//line /usr/local/go/src/regexp/regexp.go:960
				// _ = "end of CoverTab[65297]"
//line /usr/local/go/src/regexp/regexp.go:960
			}
//line /usr/local/go/src/regexp/regexp.go:960
			// _ = "end of CoverTab[65292]"
		} else {
//line /usr/local/go/src/regexp/regexp.go:961
			_go_fuzz_dep_.CoverTab[65298]++
								for i, namei := range re.subexpNames {
//line /usr/local/go/src/regexp/regexp.go:962
				_go_fuzz_dep_.CoverTab[65299]++
									if name == namei && func() bool {
//line /usr/local/go/src/regexp/regexp.go:963
					_go_fuzz_dep_.CoverTab[65300]++
//line /usr/local/go/src/regexp/regexp.go:963
					return 2*i+1 < len(match)
//line /usr/local/go/src/regexp/regexp.go:963
					// _ = "end of CoverTab[65300]"
//line /usr/local/go/src/regexp/regexp.go:963
				}() && func() bool {
//line /usr/local/go/src/regexp/regexp.go:963
					_go_fuzz_dep_.CoverTab[65301]++
//line /usr/local/go/src/regexp/regexp.go:963
					return match[2*i] >= 0
//line /usr/local/go/src/regexp/regexp.go:963
					// _ = "end of CoverTab[65301]"
//line /usr/local/go/src/regexp/regexp.go:963
				}() {
//line /usr/local/go/src/regexp/regexp.go:963
					_go_fuzz_dep_.CoverTab[65302]++
										if bsrc != nil {
//line /usr/local/go/src/regexp/regexp.go:964
						_go_fuzz_dep_.CoverTab[65304]++
											dst = append(dst, bsrc[match[2*i]:match[2*i+1]]...)
//line /usr/local/go/src/regexp/regexp.go:965
						// _ = "end of CoverTab[65304]"
					} else {
//line /usr/local/go/src/regexp/regexp.go:966
						_go_fuzz_dep_.CoverTab[65305]++
											dst = append(dst, src[match[2*i]:match[2*i+1]]...)
//line /usr/local/go/src/regexp/regexp.go:967
						// _ = "end of CoverTab[65305]"
					}
//line /usr/local/go/src/regexp/regexp.go:968
					// _ = "end of CoverTab[65302]"
//line /usr/local/go/src/regexp/regexp.go:968
					_go_fuzz_dep_.CoverTab[65303]++
										break
//line /usr/local/go/src/regexp/regexp.go:969
					// _ = "end of CoverTab[65303]"
				} else {
//line /usr/local/go/src/regexp/regexp.go:970
					_go_fuzz_dep_.CoverTab[65306]++
//line /usr/local/go/src/regexp/regexp.go:970
					// _ = "end of CoverTab[65306]"
//line /usr/local/go/src/regexp/regexp.go:970
				}
//line /usr/local/go/src/regexp/regexp.go:970
				// _ = "end of CoverTab[65299]"
			}
//line /usr/local/go/src/regexp/regexp.go:971
			// _ = "end of CoverTab[65298]"
		}
//line /usr/local/go/src/regexp/regexp.go:972
		// _ = "end of CoverTab[65284]"
	}
//line /usr/local/go/src/regexp/regexp.go:973
	// _ = "end of CoverTab[65279]"
//line /usr/local/go/src/regexp/regexp.go:973
	_go_fuzz_dep_.CoverTab[65280]++
						dst = append(dst, template...)
						return dst
//line /usr/local/go/src/regexp/regexp.go:975
	// _ = "end of CoverTab[65280]"
}

// extract returns the name from a leading "name" or "{name}" in str.
//line /usr/local/go/src/regexp/regexp.go:978
// (The $ has already been removed by the caller.)
//line /usr/local/go/src/regexp/regexp.go:978
// If it is a number, extract returns num set to that number; otherwise num = -1.
//line /usr/local/go/src/regexp/regexp.go:981
func extract(str string) (name string, num int, rest string, ok bool) {
//line /usr/local/go/src/regexp/regexp.go:981
	_go_fuzz_dep_.CoverTab[65307]++
						if str == "" {
//line /usr/local/go/src/regexp/regexp.go:982
		_go_fuzz_dep_.CoverTab[65315]++
							return
//line /usr/local/go/src/regexp/regexp.go:983
		// _ = "end of CoverTab[65315]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:984
		_go_fuzz_dep_.CoverTab[65316]++
//line /usr/local/go/src/regexp/regexp.go:984
		// _ = "end of CoverTab[65316]"
//line /usr/local/go/src/regexp/regexp.go:984
	}
//line /usr/local/go/src/regexp/regexp.go:984
	// _ = "end of CoverTab[65307]"
//line /usr/local/go/src/regexp/regexp.go:984
	_go_fuzz_dep_.CoverTab[65308]++
						brace := false
						if str[0] == '{' {
//line /usr/local/go/src/regexp/regexp.go:986
		_go_fuzz_dep_.CoverTab[65317]++
							brace = true
							str = str[1:]
//line /usr/local/go/src/regexp/regexp.go:988
		// _ = "end of CoverTab[65317]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:989
		_go_fuzz_dep_.CoverTab[65318]++
//line /usr/local/go/src/regexp/regexp.go:989
		// _ = "end of CoverTab[65318]"
//line /usr/local/go/src/regexp/regexp.go:989
	}
//line /usr/local/go/src/regexp/regexp.go:989
	// _ = "end of CoverTab[65308]"
//line /usr/local/go/src/regexp/regexp.go:989
	_go_fuzz_dep_.CoverTab[65309]++
						i := 0
						for i < len(str) {
//line /usr/local/go/src/regexp/regexp.go:991
		_go_fuzz_dep_.CoverTab[65319]++
							rune, size := utf8.DecodeRuneInString(str[i:])
							if !unicode.IsLetter(rune) && func() bool {
//line /usr/local/go/src/regexp/regexp.go:993
			_go_fuzz_dep_.CoverTab[65321]++
//line /usr/local/go/src/regexp/regexp.go:993
			return !unicode.IsDigit(rune)
//line /usr/local/go/src/regexp/regexp.go:993
			// _ = "end of CoverTab[65321]"
//line /usr/local/go/src/regexp/regexp.go:993
		}() && func() bool {
//line /usr/local/go/src/regexp/regexp.go:993
			_go_fuzz_dep_.CoverTab[65322]++
//line /usr/local/go/src/regexp/regexp.go:993
			return rune != '_'
//line /usr/local/go/src/regexp/regexp.go:993
			// _ = "end of CoverTab[65322]"
//line /usr/local/go/src/regexp/regexp.go:993
		}() {
//line /usr/local/go/src/regexp/regexp.go:993
			_go_fuzz_dep_.CoverTab[65323]++
								break
//line /usr/local/go/src/regexp/regexp.go:994
			// _ = "end of CoverTab[65323]"
		} else {
//line /usr/local/go/src/regexp/regexp.go:995
			_go_fuzz_dep_.CoverTab[65324]++
//line /usr/local/go/src/regexp/regexp.go:995
			// _ = "end of CoverTab[65324]"
//line /usr/local/go/src/regexp/regexp.go:995
		}
//line /usr/local/go/src/regexp/regexp.go:995
		// _ = "end of CoverTab[65319]"
//line /usr/local/go/src/regexp/regexp.go:995
		_go_fuzz_dep_.CoverTab[65320]++
							i += size
//line /usr/local/go/src/regexp/regexp.go:996
		// _ = "end of CoverTab[65320]"
	}
//line /usr/local/go/src/regexp/regexp.go:997
	// _ = "end of CoverTab[65309]"
//line /usr/local/go/src/regexp/regexp.go:997
	_go_fuzz_dep_.CoverTab[65310]++
						if i == 0 {
//line /usr/local/go/src/regexp/regexp.go:998
		_go_fuzz_dep_.CoverTab[65325]++

							return
//line /usr/local/go/src/regexp/regexp.go:1000
		// _ = "end of CoverTab[65325]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:1001
		_go_fuzz_dep_.CoverTab[65326]++
//line /usr/local/go/src/regexp/regexp.go:1001
		// _ = "end of CoverTab[65326]"
//line /usr/local/go/src/regexp/regexp.go:1001
	}
//line /usr/local/go/src/regexp/regexp.go:1001
	// _ = "end of CoverTab[65310]"
//line /usr/local/go/src/regexp/regexp.go:1001
	_go_fuzz_dep_.CoverTab[65311]++
						name = str[:i]
						if brace {
//line /usr/local/go/src/regexp/regexp.go:1003
		_go_fuzz_dep_.CoverTab[65327]++
							if i >= len(str) || func() bool {
//line /usr/local/go/src/regexp/regexp.go:1004
			_go_fuzz_dep_.CoverTab[65329]++
//line /usr/local/go/src/regexp/regexp.go:1004
			return str[i] != '}'
//line /usr/local/go/src/regexp/regexp.go:1004
			// _ = "end of CoverTab[65329]"
//line /usr/local/go/src/regexp/regexp.go:1004
		}() {
//line /usr/local/go/src/regexp/regexp.go:1004
			_go_fuzz_dep_.CoverTab[65330]++

								return
//line /usr/local/go/src/regexp/regexp.go:1006
			// _ = "end of CoverTab[65330]"
		} else {
//line /usr/local/go/src/regexp/regexp.go:1007
			_go_fuzz_dep_.CoverTab[65331]++
//line /usr/local/go/src/regexp/regexp.go:1007
			// _ = "end of CoverTab[65331]"
//line /usr/local/go/src/regexp/regexp.go:1007
		}
//line /usr/local/go/src/regexp/regexp.go:1007
		// _ = "end of CoverTab[65327]"
//line /usr/local/go/src/regexp/regexp.go:1007
		_go_fuzz_dep_.CoverTab[65328]++
							i++
//line /usr/local/go/src/regexp/regexp.go:1008
		// _ = "end of CoverTab[65328]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:1009
		_go_fuzz_dep_.CoverTab[65332]++
//line /usr/local/go/src/regexp/regexp.go:1009
		// _ = "end of CoverTab[65332]"
//line /usr/local/go/src/regexp/regexp.go:1009
	}
//line /usr/local/go/src/regexp/regexp.go:1009
	// _ = "end of CoverTab[65311]"
//line /usr/local/go/src/regexp/regexp.go:1009
	_go_fuzz_dep_.CoverTab[65312]++

//line /usr/local/go/src/regexp/regexp.go:1012
	num = 0
	for i := 0; i < len(name); i++ {
//line /usr/local/go/src/regexp/regexp.go:1013
		_go_fuzz_dep_.CoverTab[65333]++
							if name[i] < '0' || func() bool {
//line /usr/local/go/src/regexp/regexp.go:1014
			_go_fuzz_dep_.CoverTab[65335]++
//line /usr/local/go/src/regexp/regexp.go:1014
			return '9' < name[i]
//line /usr/local/go/src/regexp/regexp.go:1014
			// _ = "end of CoverTab[65335]"
//line /usr/local/go/src/regexp/regexp.go:1014
		}() || func() bool {
//line /usr/local/go/src/regexp/regexp.go:1014
			_go_fuzz_dep_.CoverTab[65336]++
//line /usr/local/go/src/regexp/regexp.go:1014
			return num >= 1e8
//line /usr/local/go/src/regexp/regexp.go:1014
			// _ = "end of CoverTab[65336]"
//line /usr/local/go/src/regexp/regexp.go:1014
		}() {
//line /usr/local/go/src/regexp/regexp.go:1014
			_go_fuzz_dep_.CoverTab[65337]++
								num = -1
								break
//line /usr/local/go/src/regexp/regexp.go:1016
			// _ = "end of CoverTab[65337]"
		} else {
//line /usr/local/go/src/regexp/regexp.go:1017
			_go_fuzz_dep_.CoverTab[65338]++
//line /usr/local/go/src/regexp/regexp.go:1017
			// _ = "end of CoverTab[65338]"
//line /usr/local/go/src/regexp/regexp.go:1017
		}
//line /usr/local/go/src/regexp/regexp.go:1017
		// _ = "end of CoverTab[65333]"
//line /usr/local/go/src/regexp/regexp.go:1017
		_go_fuzz_dep_.CoverTab[65334]++
							num = num*10 + int(name[i]) - '0'
//line /usr/local/go/src/regexp/regexp.go:1018
		// _ = "end of CoverTab[65334]"
	}
//line /usr/local/go/src/regexp/regexp.go:1019
	// _ = "end of CoverTab[65312]"
//line /usr/local/go/src/regexp/regexp.go:1019
	_go_fuzz_dep_.CoverTab[65313]++

						if name[0] == '0' && func() bool {
//line /usr/local/go/src/regexp/regexp.go:1021
		_go_fuzz_dep_.CoverTab[65339]++
//line /usr/local/go/src/regexp/regexp.go:1021
		return len(name) > 1
//line /usr/local/go/src/regexp/regexp.go:1021
		// _ = "end of CoverTab[65339]"
//line /usr/local/go/src/regexp/regexp.go:1021
	}() {
//line /usr/local/go/src/regexp/regexp.go:1021
		_go_fuzz_dep_.CoverTab[65340]++
							num = -1
//line /usr/local/go/src/regexp/regexp.go:1022
		// _ = "end of CoverTab[65340]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:1023
		_go_fuzz_dep_.CoverTab[65341]++
//line /usr/local/go/src/regexp/regexp.go:1023
		// _ = "end of CoverTab[65341]"
//line /usr/local/go/src/regexp/regexp.go:1023
	}
//line /usr/local/go/src/regexp/regexp.go:1023
	// _ = "end of CoverTab[65313]"
//line /usr/local/go/src/regexp/regexp.go:1023
	_go_fuzz_dep_.CoverTab[65314]++

						rest = str[i:]
						ok = true
						return
//line /usr/local/go/src/regexp/regexp.go:1027
	// _ = "end of CoverTab[65314]"
}

// FindSubmatchIndex returns a slice holding the index pairs identifying the
//line /usr/local/go/src/regexp/regexp.go:1030
// leftmost match of the regular expression in b and the matches, if any, of
//line /usr/local/go/src/regexp/regexp.go:1030
// its subexpressions, as defined by the 'Submatch' and 'Index' descriptions
//line /usr/local/go/src/regexp/regexp.go:1030
// in the package comment.
//line /usr/local/go/src/regexp/regexp.go:1030
// A return value of nil indicates no match.
//line /usr/local/go/src/regexp/regexp.go:1035
func (re *Regexp) FindSubmatchIndex(b []byte) []int {
//line /usr/local/go/src/regexp/regexp.go:1035
	_go_fuzz_dep_.CoverTab[65342]++
						return re.pad(re.doExecute(nil, b, "", 0, re.prog.NumCap, nil))
//line /usr/local/go/src/regexp/regexp.go:1036
	// _ = "end of CoverTab[65342]"
}

// FindStringSubmatch returns a slice of strings holding the text of the
//line /usr/local/go/src/regexp/regexp.go:1039
// leftmost match of the regular expression in s and the matches, if any, of
//line /usr/local/go/src/regexp/regexp.go:1039
// its subexpressions, as defined by the 'Submatch' description in the
//line /usr/local/go/src/regexp/regexp.go:1039
// package comment.
//line /usr/local/go/src/regexp/regexp.go:1039
// A return value of nil indicates no match.
//line /usr/local/go/src/regexp/regexp.go:1044
func (re *Regexp) FindStringSubmatch(s string) []string {
//line /usr/local/go/src/regexp/regexp.go:1044
	_go_fuzz_dep_.CoverTab[65343]++
						var dstCap [4]int
						a := re.doExecute(nil, nil, s, 0, re.prog.NumCap, dstCap[:0])
						if a == nil {
//line /usr/local/go/src/regexp/regexp.go:1047
		_go_fuzz_dep_.CoverTab[65346]++
							return nil
//line /usr/local/go/src/regexp/regexp.go:1048
		// _ = "end of CoverTab[65346]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:1049
		_go_fuzz_dep_.CoverTab[65347]++
//line /usr/local/go/src/regexp/regexp.go:1049
		// _ = "end of CoverTab[65347]"
//line /usr/local/go/src/regexp/regexp.go:1049
	}
//line /usr/local/go/src/regexp/regexp.go:1049
	// _ = "end of CoverTab[65343]"
//line /usr/local/go/src/regexp/regexp.go:1049
	_go_fuzz_dep_.CoverTab[65344]++
						ret := make([]string, 1+re.numSubexp)
						for i := range ret {
//line /usr/local/go/src/regexp/regexp.go:1051
		_go_fuzz_dep_.CoverTab[65348]++
							if 2*i < len(a) && func() bool {
//line /usr/local/go/src/regexp/regexp.go:1052
			_go_fuzz_dep_.CoverTab[65349]++
//line /usr/local/go/src/regexp/regexp.go:1052
			return a[2*i] >= 0
//line /usr/local/go/src/regexp/regexp.go:1052
			// _ = "end of CoverTab[65349]"
//line /usr/local/go/src/regexp/regexp.go:1052
		}() {
//line /usr/local/go/src/regexp/regexp.go:1052
			_go_fuzz_dep_.CoverTab[65350]++
								ret[i] = s[a[2*i]:a[2*i+1]]
//line /usr/local/go/src/regexp/regexp.go:1053
			// _ = "end of CoverTab[65350]"
		} else {
//line /usr/local/go/src/regexp/regexp.go:1054
			_go_fuzz_dep_.CoverTab[65351]++
//line /usr/local/go/src/regexp/regexp.go:1054
			// _ = "end of CoverTab[65351]"
//line /usr/local/go/src/regexp/regexp.go:1054
		}
//line /usr/local/go/src/regexp/regexp.go:1054
		// _ = "end of CoverTab[65348]"
	}
//line /usr/local/go/src/regexp/regexp.go:1055
	// _ = "end of CoverTab[65344]"
//line /usr/local/go/src/regexp/regexp.go:1055
	_go_fuzz_dep_.CoverTab[65345]++
						return ret
//line /usr/local/go/src/regexp/regexp.go:1056
	// _ = "end of CoverTab[65345]"
}

// FindStringSubmatchIndex returns a slice holding the index pairs
//line /usr/local/go/src/regexp/regexp.go:1059
// identifying the leftmost match of the regular expression in s and the
//line /usr/local/go/src/regexp/regexp.go:1059
// matches, if any, of its subexpressions, as defined by the 'Submatch' and
//line /usr/local/go/src/regexp/regexp.go:1059
// 'Index' descriptions in the package comment.
//line /usr/local/go/src/regexp/regexp.go:1059
// A return value of nil indicates no match.
//line /usr/local/go/src/regexp/regexp.go:1064
func (re *Regexp) FindStringSubmatchIndex(s string) []int {
//line /usr/local/go/src/regexp/regexp.go:1064
	_go_fuzz_dep_.CoverTab[65352]++
						return re.pad(re.doExecute(nil, nil, s, 0, re.prog.NumCap, nil))
//line /usr/local/go/src/regexp/regexp.go:1065
	// _ = "end of CoverTab[65352]"
}

// FindReaderSubmatchIndex returns a slice holding the index pairs
//line /usr/local/go/src/regexp/regexp.go:1068
// identifying the leftmost match of the regular expression of text read by
//line /usr/local/go/src/regexp/regexp.go:1068
// the RuneReader, and the matches, if any, of its subexpressions, as defined
//line /usr/local/go/src/regexp/regexp.go:1068
// by the 'Submatch' and 'Index' descriptions in the package comment. A
//line /usr/local/go/src/regexp/regexp.go:1068
// return value of nil indicates no match.
//line /usr/local/go/src/regexp/regexp.go:1073
func (re *Regexp) FindReaderSubmatchIndex(r io.RuneReader) []int {
//line /usr/local/go/src/regexp/regexp.go:1073
	_go_fuzz_dep_.CoverTab[65353]++
						return re.pad(re.doExecute(r, nil, "", 0, re.prog.NumCap, nil))
//line /usr/local/go/src/regexp/regexp.go:1074
	// _ = "end of CoverTab[65353]"
}

const startSize = 10	// The size at which to start a slice in the 'All' routines.

// FindAll is the 'All' version of Find; it returns a slice of all successive
//line /usr/local/go/src/regexp/regexp.go:1079
// matches of the expression, as defined by the 'All' description in the
//line /usr/local/go/src/regexp/regexp.go:1079
// package comment.
//line /usr/local/go/src/regexp/regexp.go:1079
// A return value of nil indicates no match.
//line /usr/local/go/src/regexp/regexp.go:1083
func (re *Regexp) FindAll(b []byte, n int) [][]byte {
//line /usr/local/go/src/regexp/regexp.go:1083
	_go_fuzz_dep_.CoverTab[65354]++
						if n < 0 {
//line /usr/local/go/src/regexp/regexp.go:1084
		_go_fuzz_dep_.CoverTab[65357]++
							n = len(b) + 1
//line /usr/local/go/src/regexp/regexp.go:1085
		// _ = "end of CoverTab[65357]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:1086
		_go_fuzz_dep_.CoverTab[65358]++
//line /usr/local/go/src/regexp/regexp.go:1086
		// _ = "end of CoverTab[65358]"
//line /usr/local/go/src/regexp/regexp.go:1086
	}
//line /usr/local/go/src/regexp/regexp.go:1086
	// _ = "end of CoverTab[65354]"
//line /usr/local/go/src/regexp/regexp.go:1086
	_go_fuzz_dep_.CoverTab[65355]++
						var result [][]byte
						re.allMatches("", b, n, func(match []int) {
//line /usr/local/go/src/regexp/regexp.go:1088
		_go_fuzz_dep_.CoverTab[65359]++
							if result == nil {
//line /usr/local/go/src/regexp/regexp.go:1089
			_go_fuzz_dep_.CoverTab[65361]++
								result = make([][]byte, 0, startSize)
//line /usr/local/go/src/regexp/regexp.go:1090
			// _ = "end of CoverTab[65361]"
		} else {
//line /usr/local/go/src/regexp/regexp.go:1091
			_go_fuzz_dep_.CoverTab[65362]++
//line /usr/local/go/src/regexp/regexp.go:1091
			// _ = "end of CoverTab[65362]"
//line /usr/local/go/src/regexp/regexp.go:1091
		}
//line /usr/local/go/src/regexp/regexp.go:1091
		// _ = "end of CoverTab[65359]"
//line /usr/local/go/src/regexp/regexp.go:1091
		_go_fuzz_dep_.CoverTab[65360]++
							result = append(result, b[match[0]:match[1]:match[1]])
//line /usr/local/go/src/regexp/regexp.go:1092
		// _ = "end of CoverTab[65360]"
	})
//line /usr/local/go/src/regexp/regexp.go:1093
	// _ = "end of CoverTab[65355]"
//line /usr/local/go/src/regexp/regexp.go:1093
	_go_fuzz_dep_.CoverTab[65356]++
						return result
//line /usr/local/go/src/regexp/regexp.go:1094
	// _ = "end of CoverTab[65356]"
}

// FindAllIndex is the 'All' version of FindIndex; it returns a slice of all
//line /usr/local/go/src/regexp/regexp.go:1097
// successive matches of the expression, as defined by the 'All' description
//line /usr/local/go/src/regexp/regexp.go:1097
// in the package comment.
//line /usr/local/go/src/regexp/regexp.go:1097
// A return value of nil indicates no match.
//line /usr/local/go/src/regexp/regexp.go:1101
func (re *Regexp) FindAllIndex(b []byte, n int) [][]int {
//line /usr/local/go/src/regexp/regexp.go:1101
	_go_fuzz_dep_.CoverTab[65363]++
						if n < 0 {
//line /usr/local/go/src/regexp/regexp.go:1102
		_go_fuzz_dep_.CoverTab[65366]++
							n = len(b) + 1
//line /usr/local/go/src/regexp/regexp.go:1103
		// _ = "end of CoverTab[65366]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:1104
		_go_fuzz_dep_.CoverTab[65367]++
//line /usr/local/go/src/regexp/regexp.go:1104
		// _ = "end of CoverTab[65367]"
//line /usr/local/go/src/regexp/regexp.go:1104
	}
//line /usr/local/go/src/regexp/regexp.go:1104
	// _ = "end of CoverTab[65363]"
//line /usr/local/go/src/regexp/regexp.go:1104
	_go_fuzz_dep_.CoverTab[65364]++
						var result [][]int
						re.allMatches("", b, n, func(match []int) {
//line /usr/local/go/src/regexp/regexp.go:1106
		_go_fuzz_dep_.CoverTab[65368]++
							if result == nil {
//line /usr/local/go/src/regexp/regexp.go:1107
			_go_fuzz_dep_.CoverTab[65370]++
								result = make([][]int, 0, startSize)
//line /usr/local/go/src/regexp/regexp.go:1108
			// _ = "end of CoverTab[65370]"
		} else {
//line /usr/local/go/src/regexp/regexp.go:1109
			_go_fuzz_dep_.CoverTab[65371]++
//line /usr/local/go/src/regexp/regexp.go:1109
			// _ = "end of CoverTab[65371]"
//line /usr/local/go/src/regexp/regexp.go:1109
		}
//line /usr/local/go/src/regexp/regexp.go:1109
		// _ = "end of CoverTab[65368]"
//line /usr/local/go/src/regexp/regexp.go:1109
		_go_fuzz_dep_.CoverTab[65369]++
							result = append(result, match[0:2])
//line /usr/local/go/src/regexp/regexp.go:1110
		// _ = "end of CoverTab[65369]"
	})
//line /usr/local/go/src/regexp/regexp.go:1111
	// _ = "end of CoverTab[65364]"
//line /usr/local/go/src/regexp/regexp.go:1111
	_go_fuzz_dep_.CoverTab[65365]++
						return result
//line /usr/local/go/src/regexp/regexp.go:1112
	// _ = "end of CoverTab[65365]"
}

// FindAllString is the 'All' version of FindString; it returns a slice of all
//line /usr/local/go/src/regexp/regexp.go:1115
// successive matches of the expression, as defined by the 'All' description
//line /usr/local/go/src/regexp/regexp.go:1115
// in the package comment.
//line /usr/local/go/src/regexp/regexp.go:1115
// A return value of nil indicates no match.
//line /usr/local/go/src/regexp/regexp.go:1119
func (re *Regexp) FindAllString(s string, n int) []string {
//line /usr/local/go/src/regexp/regexp.go:1119
	_go_fuzz_dep_.CoverTab[65372]++
						if n < 0 {
//line /usr/local/go/src/regexp/regexp.go:1120
		_go_fuzz_dep_.CoverTab[65375]++
							n = len(s) + 1
//line /usr/local/go/src/regexp/regexp.go:1121
		// _ = "end of CoverTab[65375]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:1122
		_go_fuzz_dep_.CoverTab[65376]++
//line /usr/local/go/src/regexp/regexp.go:1122
		// _ = "end of CoverTab[65376]"
//line /usr/local/go/src/regexp/regexp.go:1122
	}
//line /usr/local/go/src/regexp/regexp.go:1122
	// _ = "end of CoverTab[65372]"
//line /usr/local/go/src/regexp/regexp.go:1122
	_go_fuzz_dep_.CoverTab[65373]++
						var result []string
						re.allMatches(s, nil, n, func(match []int) {
//line /usr/local/go/src/regexp/regexp.go:1124
		_go_fuzz_dep_.CoverTab[65377]++
							if result == nil {
//line /usr/local/go/src/regexp/regexp.go:1125
			_go_fuzz_dep_.CoverTab[65379]++
								result = make([]string, 0, startSize)
//line /usr/local/go/src/regexp/regexp.go:1126
			// _ = "end of CoverTab[65379]"
		} else {
//line /usr/local/go/src/regexp/regexp.go:1127
			_go_fuzz_dep_.CoverTab[65380]++
//line /usr/local/go/src/regexp/regexp.go:1127
			// _ = "end of CoverTab[65380]"
//line /usr/local/go/src/regexp/regexp.go:1127
		}
//line /usr/local/go/src/regexp/regexp.go:1127
		// _ = "end of CoverTab[65377]"
//line /usr/local/go/src/regexp/regexp.go:1127
		_go_fuzz_dep_.CoverTab[65378]++
							result = append(result, s[match[0]:match[1]])
//line /usr/local/go/src/regexp/regexp.go:1128
		// _ = "end of CoverTab[65378]"
	})
//line /usr/local/go/src/regexp/regexp.go:1129
	// _ = "end of CoverTab[65373]"
//line /usr/local/go/src/regexp/regexp.go:1129
	_go_fuzz_dep_.CoverTab[65374]++
						return result
//line /usr/local/go/src/regexp/regexp.go:1130
	// _ = "end of CoverTab[65374]"
}

// FindAllStringIndex is the 'All' version of FindStringIndex; it returns a
//line /usr/local/go/src/regexp/regexp.go:1133
// slice of all successive matches of the expression, as defined by the 'All'
//line /usr/local/go/src/regexp/regexp.go:1133
// description in the package comment.
//line /usr/local/go/src/regexp/regexp.go:1133
// A return value of nil indicates no match.
//line /usr/local/go/src/regexp/regexp.go:1137
func (re *Regexp) FindAllStringIndex(s string, n int) [][]int {
//line /usr/local/go/src/regexp/regexp.go:1137
	_go_fuzz_dep_.CoverTab[65381]++
						if n < 0 {
//line /usr/local/go/src/regexp/regexp.go:1138
		_go_fuzz_dep_.CoverTab[65384]++
							n = len(s) + 1
//line /usr/local/go/src/regexp/regexp.go:1139
		// _ = "end of CoverTab[65384]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:1140
		_go_fuzz_dep_.CoverTab[65385]++
//line /usr/local/go/src/regexp/regexp.go:1140
		// _ = "end of CoverTab[65385]"
//line /usr/local/go/src/regexp/regexp.go:1140
	}
//line /usr/local/go/src/regexp/regexp.go:1140
	// _ = "end of CoverTab[65381]"
//line /usr/local/go/src/regexp/regexp.go:1140
	_go_fuzz_dep_.CoverTab[65382]++
						var result [][]int
						re.allMatches(s, nil, n, func(match []int) {
//line /usr/local/go/src/regexp/regexp.go:1142
		_go_fuzz_dep_.CoverTab[65386]++
							if result == nil {
//line /usr/local/go/src/regexp/regexp.go:1143
			_go_fuzz_dep_.CoverTab[65388]++
								result = make([][]int, 0, startSize)
//line /usr/local/go/src/regexp/regexp.go:1144
			// _ = "end of CoverTab[65388]"
		} else {
//line /usr/local/go/src/regexp/regexp.go:1145
			_go_fuzz_dep_.CoverTab[65389]++
//line /usr/local/go/src/regexp/regexp.go:1145
			// _ = "end of CoverTab[65389]"
//line /usr/local/go/src/regexp/regexp.go:1145
		}
//line /usr/local/go/src/regexp/regexp.go:1145
		// _ = "end of CoverTab[65386]"
//line /usr/local/go/src/regexp/regexp.go:1145
		_go_fuzz_dep_.CoverTab[65387]++
							result = append(result, match[0:2])
//line /usr/local/go/src/regexp/regexp.go:1146
		// _ = "end of CoverTab[65387]"
	})
//line /usr/local/go/src/regexp/regexp.go:1147
	// _ = "end of CoverTab[65382]"
//line /usr/local/go/src/regexp/regexp.go:1147
	_go_fuzz_dep_.CoverTab[65383]++
						return result
//line /usr/local/go/src/regexp/regexp.go:1148
	// _ = "end of CoverTab[65383]"
}

// FindAllSubmatch is the 'All' version of FindSubmatch; it returns a slice
//line /usr/local/go/src/regexp/regexp.go:1151
// of all successive matches of the expression, as defined by the 'All'
//line /usr/local/go/src/regexp/regexp.go:1151
// description in the package comment.
//line /usr/local/go/src/regexp/regexp.go:1151
// A return value of nil indicates no match.
//line /usr/local/go/src/regexp/regexp.go:1155
func (re *Regexp) FindAllSubmatch(b []byte, n int) [][][]byte {
//line /usr/local/go/src/regexp/regexp.go:1155
	_go_fuzz_dep_.CoverTab[65390]++
						if n < 0 {
//line /usr/local/go/src/regexp/regexp.go:1156
		_go_fuzz_dep_.CoverTab[65393]++
							n = len(b) + 1
//line /usr/local/go/src/regexp/regexp.go:1157
		// _ = "end of CoverTab[65393]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:1158
		_go_fuzz_dep_.CoverTab[65394]++
//line /usr/local/go/src/regexp/regexp.go:1158
		// _ = "end of CoverTab[65394]"
//line /usr/local/go/src/regexp/regexp.go:1158
	}
//line /usr/local/go/src/regexp/regexp.go:1158
	// _ = "end of CoverTab[65390]"
//line /usr/local/go/src/regexp/regexp.go:1158
	_go_fuzz_dep_.CoverTab[65391]++
						var result [][][]byte
						re.allMatches("", b, n, func(match []int) {
//line /usr/local/go/src/regexp/regexp.go:1160
		_go_fuzz_dep_.CoverTab[65395]++
							if result == nil {
//line /usr/local/go/src/regexp/regexp.go:1161
			_go_fuzz_dep_.CoverTab[65398]++
								result = make([][][]byte, 0, startSize)
//line /usr/local/go/src/regexp/regexp.go:1162
			// _ = "end of CoverTab[65398]"
		} else {
//line /usr/local/go/src/regexp/regexp.go:1163
			_go_fuzz_dep_.CoverTab[65399]++
//line /usr/local/go/src/regexp/regexp.go:1163
			// _ = "end of CoverTab[65399]"
//line /usr/local/go/src/regexp/regexp.go:1163
		}
//line /usr/local/go/src/regexp/regexp.go:1163
		// _ = "end of CoverTab[65395]"
//line /usr/local/go/src/regexp/regexp.go:1163
		_go_fuzz_dep_.CoverTab[65396]++
							slice := make([][]byte, len(match)/2)
							for j := range slice {
//line /usr/local/go/src/regexp/regexp.go:1165
			_go_fuzz_dep_.CoverTab[65400]++
								if match[2*j] >= 0 {
//line /usr/local/go/src/regexp/regexp.go:1166
				_go_fuzz_dep_.CoverTab[65401]++
									slice[j] = b[match[2*j]:match[2*j+1]:match[2*j+1]]
//line /usr/local/go/src/regexp/regexp.go:1167
				// _ = "end of CoverTab[65401]"
			} else {
//line /usr/local/go/src/regexp/regexp.go:1168
				_go_fuzz_dep_.CoverTab[65402]++
//line /usr/local/go/src/regexp/regexp.go:1168
				// _ = "end of CoverTab[65402]"
//line /usr/local/go/src/regexp/regexp.go:1168
			}
//line /usr/local/go/src/regexp/regexp.go:1168
			// _ = "end of CoverTab[65400]"
		}
//line /usr/local/go/src/regexp/regexp.go:1169
		// _ = "end of CoverTab[65396]"
//line /usr/local/go/src/regexp/regexp.go:1169
		_go_fuzz_dep_.CoverTab[65397]++
							result = append(result, slice)
//line /usr/local/go/src/regexp/regexp.go:1170
		// _ = "end of CoverTab[65397]"
	})
//line /usr/local/go/src/regexp/regexp.go:1171
	// _ = "end of CoverTab[65391]"
//line /usr/local/go/src/regexp/regexp.go:1171
	_go_fuzz_dep_.CoverTab[65392]++
						return result
//line /usr/local/go/src/regexp/regexp.go:1172
	// _ = "end of CoverTab[65392]"
}

// FindAllSubmatchIndex is the 'All' version of FindSubmatchIndex; it returns
//line /usr/local/go/src/regexp/regexp.go:1175
// a slice of all successive matches of the expression, as defined by the
//line /usr/local/go/src/regexp/regexp.go:1175
// 'All' description in the package comment.
//line /usr/local/go/src/regexp/regexp.go:1175
// A return value of nil indicates no match.
//line /usr/local/go/src/regexp/regexp.go:1179
func (re *Regexp) FindAllSubmatchIndex(b []byte, n int) [][]int {
//line /usr/local/go/src/regexp/regexp.go:1179
	_go_fuzz_dep_.CoverTab[65403]++
						if n < 0 {
//line /usr/local/go/src/regexp/regexp.go:1180
		_go_fuzz_dep_.CoverTab[65406]++
							n = len(b) + 1
//line /usr/local/go/src/regexp/regexp.go:1181
		// _ = "end of CoverTab[65406]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:1182
		_go_fuzz_dep_.CoverTab[65407]++
//line /usr/local/go/src/regexp/regexp.go:1182
		// _ = "end of CoverTab[65407]"
//line /usr/local/go/src/regexp/regexp.go:1182
	}
//line /usr/local/go/src/regexp/regexp.go:1182
	// _ = "end of CoverTab[65403]"
//line /usr/local/go/src/regexp/regexp.go:1182
	_go_fuzz_dep_.CoverTab[65404]++
						var result [][]int
						re.allMatches("", b, n, func(match []int) {
//line /usr/local/go/src/regexp/regexp.go:1184
		_go_fuzz_dep_.CoverTab[65408]++
							if result == nil {
//line /usr/local/go/src/regexp/regexp.go:1185
			_go_fuzz_dep_.CoverTab[65410]++
								result = make([][]int, 0, startSize)
//line /usr/local/go/src/regexp/regexp.go:1186
			// _ = "end of CoverTab[65410]"
		} else {
//line /usr/local/go/src/regexp/regexp.go:1187
			_go_fuzz_dep_.CoverTab[65411]++
//line /usr/local/go/src/regexp/regexp.go:1187
			// _ = "end of CoverTab[65411]"
//line /usr/local/go/src/regexp/regexp.go:1187
		}
//line /usr/local/go/src/regexp/regexp.go:1187
		// _ = "end of CoverTab[65408]"
//line /usr/local/go/src/regexp/regexp.go:1187
		_go_fuzz_dep_.CoverTab[65409]++
							result = append(result, match)
//line /usr/local/go/src/regexp/regexp.go:1188
		// _ = "end of CoverTab[65409]"
	})
//line /usr/local/go/src/regexp/regexp.go:1189
	// _ = "end of CoverTab[65404]"
//line /usr/local/go/src/regexp/regexp.go:1189
	_go_fuzz_dep_.CoverTab[65405]++
						return result
//line /usr/local/go/src/regexp/regexp.go:1190
	// _ = "end of CoverTab[65405]"
}

// FindAllStringSubmatch is the 'All' version of FindStringSubmatch; it
//line /usr/local/go/src/regexp/regexp.go:1193
// returns a slice of all successive matches of the expression, as defined by
//line /usr/local/go/src/regexp/regexp.go:1193
// the 'All' description in the package comment.
//line /usr/local/go/src/regexp/regexp.go:1193
// A return value of nil indicates no match.
//line /usr/local/go/src/regexp/regexp.go:1197
func (re *Regexp) FindAllStringSubmatch(s string, n int) [][]string {
//line /usr/local/go/src/regexp/regexp.go:1197
	_go_fuzz_dep_.CoverTab[65412]++
						if n < 0 {
//line /usr/local/go/src/regexp/regexp.go:1198
		_go_fuzz_dep_.CoverTab[65415]++
							n = len(s) + 1
//line /usr/local/go/src/regexp/regexp.go:1199
		// _ = "end of CoverTab[65415]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:1200
		_go_fuzz_dep_.CoverTab[65416]++
//line /usr/local/go/src/regexp/regexp.go:1200
		// _ = "end of CoverTab[65416]"
//line /usr/local/go/src/regexp/regexp.go:1200
	}
//line /usr/local/go/src/regexp/regexp.go:1200
	// _ = "end of CoverTab[65412]"
//line /usr/local/go/src/regexp/regexp.go:1200
	_go_fuzz_dep_.CoverTab[65413]++
						var result [][]string
						re.allMatches(s, nil, n, func(match []int) {
//line /usr/local/go/src/regexp/regexp.go:1202
		_go_fuzz_dep_.CoverTab[65417]++
							if result == nil {
//line /usr/local/go/src/regexp/regexp.go:1203
			_go_fuzz_dep_.CoverTab[65420]++
								result = make([][]string, 0, startSize)
//line /usr/local/go/src/regexp/regexp.go:1204
			// _ = "end of CoverTab[65420]"
		} else {
//line /usr/local/go/src/regexp/regexp.go:1205
			_go_fuzz_dep_.CoverTab[65421]++
//line /usr/local/go/src/regexp/regexp.go:1205
			// _ = "end of CoverTab[65421]"
//line /usr/local/go/src/regexp/regexp.go:1205
		}
//line /usr/local/go/src/regexp/regexp.go:1205
		// _ = "end of CoverTab[65417]"
//line /usr/local/go/src/regexp/regexp.go:1205
		_go_fuzz_dep_.CoverTab[65418]++
							slice := make([]string, len(match)/2)
							for j := range slice {
//line /usr/local/go/src/regexp/regexp.go:1207
			_go_fuzz_dep_.CoverTab[65422]++
								if match[2*j] >= 0 {
//line /usr/local/go/src/regexp/regexp.go:1208
				_go_fuzz_dep_.CoverTab[65423]++
									slice[j] = s[match[2*j]:match[2*j+1]]
//line /usr/local/go/src/regexp/regexp.go:1209
				// _ = "end of CoverTab[65423]"
			} else {
//line /usr/local/go/src/regexp/regexp.go:1210
				_go_fuzz_dep_.CoverTab[65424]++
//line /usr/local/go/src/regexp/regexp.go:1210
				// _ = "end of CoverTab[65424]"
//line /usr/local/go/src/regexp/regexp.go:1210
			}
//line /usr/local/go/src/regexp/regexp.go:1210
			// _ = "end of CoverTab[65422]"
		}
//line /usr/local/go/src/regexp/regexp.go:1211
		// _ = "end of CoverTab[65418]"
//line /usr/local/go/src/regexp/regexp.go:1211
		_go_fuzz_dep_.CoverTab[65419]++
							result = append(result, slice)
//line /usr/local/go/src/regexp/regexp.go:1212
		// _ = "end of CoverTab[65419]"
	})
//line /usr/local/go/src/regexp/regexp.go:1213
	// _ = "end of CoverTab[65413]"
//line /usr/local/go/src/regexp/regexp.go:1213
	_go_fuzz_dep_.CoverTab[65414]++
						return result
//line /usr/local/go/src/regexp/regexp.go:1214
	// _ = "end of CoverTab[65414]"
}

// FindAllStringSubmatchIndex is the 'All' version of
//line /usr/local/go/src/regexp/regexp.go:1217
// FindStringSubmatchIndex; it returns a slice of all successive matches of
//line /usr/local/go/src/regexp/regexp.go:1217
// the expression, as defined by the 'All' description in the package
//line /usr/local/go/src/regexp/regexp.go:1217
// comment.
//line /usr/local/go/src/regexp/regexp.go:1217
// A return value of nil indicates no match.
//line /usr/local/go/src/regexp/regexp.go:1222
func (re *Regexp) FindAllStringSubmatchIndex(s string, n int) [][]int {
//line /usr/local/go/src/regexp/regexp.go:1222
	_go_fuzz_dep_.CoverTab[65425]++
						if n < 0 {
//line /usr/local/go/src/regexp/regexp.go:1223
		_go_fuzz_dep_.CoverTab[65428]++
							n = len(s) + 1
//line /usr/local/go/src/regexp/regexp.go:1224
		// _ = "end of CoverTab[65428]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:1225
		_go_fuzz_dep_.CoverTab[65429]++
//line /usr/local/go/src/regexp/regexp.go:1225
		// _ = "end of CoverTab[65429]"
//line /usr/local/go/src/regexp/regexp.go:1225
	}
//line /usr/local/go/src/regexp/regexp.go:1225
	// _ = "end of CoverTab[65425]"
//line /usr/local/go/src/regexp/regexp.go:1225
	_go_fuzz_dep_.CoverTab[65426]++
						var result [][]int
						re.allMatches(s, nil, n, func(match []int) {
//line /usr/local/go/src/regexp/regexp.go:1227
		_go_fuzz_dep_.CoverTab[65430]++
							if result == nil {
//line /usr/local/go/src/regexp/regexp.go:1228
			_go_fuzz_dep_.CoverTab[65432]++
								result = make([][]int, 0, startSize)
//line /usr/local/go/src/regexp/regexp.go:1229
			// _ = "end of CoverTab[65432]"
		} else {
//line /usr/local/go/src/regexp/regexp.go:1230
			_go_fuzz_dep_.CoverTab[65433]++
//line /usr/local/go/src/regexp/regexp.go:1230
			// _ = "end of CoverTab[65433]"
//line /usr/local/go/src/regexp/regexp.go:1230
		}
//line /usr/local/go/src/regexp/regexp.go:1230
		// _ = "end of CoverTab[65430]"
//line /usr/local/go/src/regexp/regexp.go:1230
		_go_fuzz_dep_.CoverTab[65431]++
							result = append(result, match)
//line /usr/local/go/src/regexp/regexp.go:1231
		// _ = "end of CoverTab[65431]"
	})
//line /usr/local/go/src/regexp/regexp.go:1232
	// _ = "end of CoverTab[65426]"
//line /usr/local/go/src/regexp/regexp.go:1232
	_go_fuzz_dep_.CoverTab[65427]++
						return result
//line /usr/local/go/src/regexp/regexp.go:1233
	// _ = "end of CoverTab[65427]"
}

// Split slices s into substrings separated by the expression and returns a slice of
//line /usr/local/go/src/regexp/regexp.go:1236
// the substrings between those expression matches.
//line /usr/local/go/src/regexp/regexp.go:1236
//
//line /usr/local/go/src/regexp/regexp.go:1236
// The slice returned by this method consists of all the substrings of s
//line /usr/local/go/src/regexp/regexp.go:1236
// not contained in the slice returned by FindAllString. When called on an expression
//line /usr/local/go/src/regexp/regexp.go:1236
// that contains no metacharacters, it is equivalent to strings.SplitN.
//line /usr/local/go/src/regexp/regexp.go:1236
//
//line /usr/local/go/src/regexp/regexp.go:1236
// Example:
//line /usr/local/go/src/regexp/regexp.go:1236
//
//line /usr/local/go/src/regexp/regexp.go:1236
//	s := regexp.MustCompile("a*").Split("abaabaccadaaae", 5)
//line /usr/local/go/src/regexp/regexp.go:1236
//	// s: ["", "b", "b", "c", "cadaaae"]
//line /usr/local/go/src/regexp/regexp.go:1236
//
//line /usr/local/go/src/regexp/regexp.go:1236
// The count determines the number of substrings to return:
//line /usr/local/go/src/regexp/regexp.go:1236
//
//line /usr/local/go/src/regexp/regexp.go:1236
//	n > 0: at most n substrings; the last substring will be the unsplit remainder.
//line /usr/local/go/src/regexp/regexp.go:1236
//	n == 0: the result is nil (zero substrings)
//line /usr/local/go/src/regexp/regexp.go:1236
//	n < 0: all substrings
//line /usr/local/go/src/regexp/regexp.go:1253
func (re *Regexp) Split(s string, n int) []string {
//line /usr/local/go/src/regexp/regexp.go:1253
	_go_fuzz_dep_.CoverTab[65434]++

						if n == 0 {
//line /usr/local/go/src/regexp/regexp.go:1255
		_go_fuzz_dep_.CoverTab[65439]++
							return nil
//line /usr/local/go/src/regexp/regexp.go:1256
		// _ = "end of CoverTab[65439]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:1257
		_go_fuzz_dep_.CoverTab[65440]++
//line /usr/local/go/src/regexp/regexp.go:1257
		// _ = "end of CoverTab[65440]"
//line /usr/local/go/src/regexp/regexp.go:1257
	}
//line /usr/local/go/src/regexp/regexp.go:1257
	// _ = "end of CoverTab[65434]"
//line /usr/local/go/src/regexp/regexp.go:1257
	_go_fuzz_dep_.CoverTab[65435]++

						if len(re.expr) > 0 && func() bool {
//line /usr/local/go/src/regexp/regexp.go:1259
		_go_fuzz_dep_.CoverTab[65441]++
//line /usr/local/go/src/regexp/regexp.go:1259
		return len(s) == 0
//line /usr/local/go/src/regexp/regexp.go:1259
		// _ = "end of CoverTab[65441]"
//line /usr/local/go/src/regexp/regexp.go:1259
	}() {
//line /usr/local/go/src/regexp/regexp.go:1259
		_go_fuzz_dep_.CoverTab[65442]++
							return []string{""}
//line /usr/local/go/src/regexp/regexp.go:1260
		// _ = "end of CoverTab[65442]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:1261
		_go_fuzz_dep_.CoverTab[65443]++
//line /usr/local/go/src/regexp/regexp.go:1261
		// _ = "end of CoverTab[65443]"
//line /usr/local/go/src/regexp/regexp.go:1261
	}
//line /usr/local/go/src/regexp/regexp.go:1261
	// _ = "end of CoverTab[65435]"
//line /usr/local/go/src/regexp/regexp.go:1261
	_go_fuzz_dep_.CoverTab[65436]++

						matches := re.FindAllStringIndex(s, n)
						strings := make([]string, 0, len(matches))

						beg := 0
						end := 0
						for _, match := range matches {
//line /usr/local/go/src/regexp/regexp.go:1268
		_go_fuzz_dep_.CoverTab[65444]++
							if n > 0 && func() bool {
//line /usr/local/go/src/regexp/regexp.go:1269
			_go_fuzz_dep_.CoverTab[65447]++
//line /usr/local/go/src/regexp/regexp.go:1269
			return len(strings) >= n-1
//line /usr/local/go/src/regexp/regexp.go:1269
			// _ = "end of CoverTab[65447]"
//line /usr/local/go/src/regexp/regexp.go:1269
		}() {
//line /usr/local/go/src/regexp/regexp.go:1269
			_go_fuzz_dep_.CoverTab[65448]++
								break
//line /usr/local/go/src/regexp/regexp.go:1270
			// _ = "end of CoverTab[65448]"
		} else {
//line /usr/local/go/src/regexp/regexp.go:1271
			_go_fuzz_dep_.CoverTab[65449]++
//line /usr/local/go/src/regexp/regexp.go:1271
			// _ = "end of CoverTab[65449]"
//line /usr/local/go/src/regexp/regexp.go:1271
		}
//line /usr/local/go/src/regexp/regexp.go:1271
		// _ = "end of CoverTab[65444]"
//line /usr/local/go/src/regexp/regexp.go:1271
		_go_fuzz_dep_.CoverTab[65445]++

							end = match[0]
							if match[1] != 0 {
//line /usr/local/go/src/regexp/regexp.go:1274
			_go_fuzz_dep_.CoverTab[65450]++
								strings = append(strings, s[beg:end])
//line /usr/local/go/src/regexp/regexp.go:1275
			// _ = "end of CoverTab[65450]"
		} else {
//line /usr/local/go/src/regexp/regexp.go:1276
			_go_fuzz_dep_.CoverTab[65451]++
//line /usr/local/go/src/regexp/regexp.go:1276
			// _ = "end of CoverTab[65451]"
//line /usr/local/go/src/regexp/regexp.go:1276
		}
//line /usr/local/go/src/regexp/regexp.go:1276
		// _ = "end of CoverTab[65445]"
//line /usr/local/go/src/regexp/regexp.go:1276
		_go_fuzz_dep_.CoverTab[65446]++
							beg = match[1]
//line /usr/local/go/src/regexp/regexp.go:1277
		// _ = "end of CoverTab[65446]"
	}
//line /usr/local/go/src/regexp/regexp.go:1278
	// _ = "end of CoverTab[65436]"
//line /usr/local/go/src/regexp/regexp.go:1278
	_go_fuzz_dep_.CoverTab[65437]++

						if end != len(s) {
//line /usr/local/go/src/regexp/regexp.go:1280
		_go_fuzz_dep_.CoverTab[65452]++
							strings = append(strings, s[beg:])
//line /usr/local/go/src/regexp/regexp.go:1281
		// _ = "end of CoverTab[65452]"
	} else {
//line /usr/local/go/src/regexp/regexp.go:1282
		_go_fuzz_dep_.CoverTab[65453]++
//line /usr/local/go/src/regexp/regexp.go:1282
		// _ = "end of CoverTab[65453]"
//line /usr/local/go/src/regexp/regexp.go:1282
	}
//line /usr/local/go/src/regexp/regexp.go:1282
	// _ = "end of CoverTab[65437]"
//line /usr/local/go/src/regexp/regexp.go:1282
	_go_fuzz_dep_.CoverTab[65438]++

						return strings
//line /usr/local/go/src/regexp/regexp.go:1284
	// _ = "end of CoverTab[65438]"
}

//line /usr/local/go/src/regexp/regexp.go:1285
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/regexp/regexp.go:1285
var _ = _go_fuzz_dep_.CoverTab
