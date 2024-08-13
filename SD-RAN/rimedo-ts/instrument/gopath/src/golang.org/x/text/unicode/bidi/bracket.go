// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:5
package bidi

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:5
)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:5
)

import (
	"container/list"
	"fmt"
	"sort"
)

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:35
// Bidi_Paired_Bracket_Type
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:35
// BD14. An opening paired bracket is a character whose
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:35
// Bidi_Paired_Bracket_Type property value is Open.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:35
//
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:35
// BD15. A closing paired bracket is a character whose
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:35
// Bidi_Paired_Bracket_Type property value is Close.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:41
type bracketType byte

const (
	bpNone	bracketType	= iota
	bpOpen
	bpClose
)

// bracketPair holds a pair of index values for opening and closing bracket
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:49
// location of a bracket pair.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:51
type bracketPair struct {
	opener	int
	closer	int
}

func (b *bracketPair) String() string {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:56
	_go_fuzz_dep_.CoverTab[69475]++
											return fmt.Sprintf("(%v, %v)", b.opener, b.closer)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:57
	// _ = "end of CoverTab[69475]"
}

// bracketPairs is a slice of bracketPairs with a sort.Interface implementation.
type bracketPairs []bracketPair

func (b bracketPairs) Len() int {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:63
	_go_fuzz_dep_.CoverTab[69476]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:63
	return len(b)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:63
	// _ = "end of CoverTab[69476]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:63
}
func (b bracketPairs) Swap(i, j int) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:64
	_go_fuzz_dep_.CoverTab[69477]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:64
	b[i], b[j] = b[j], b[i]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:64
	// _ = "end of CoverTab[69477]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:64
}
func (b bracketPairs) Less(i, j int) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:65
	_go_fuzz_dep_.CoverTab[69478]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:65
	return b[i].opener < b[j].opener
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:65
	// _ = "end of CoverTab[69478]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:65
}

// resolvePairedBrackets runs the paired bracket part of the UBA algorithm.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:67
//
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:67
// For each rune, it takes the indexes into the original string, the class the
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:67
// bracket type (in pairTypes) and the bracket identifier (pairValues). It also
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:67
// takes the direction type for the start-of-sentence and the embedding level.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:67
//
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:67
// The identifiers for bracket types are the rune of the canonicalized opening
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:67
// bracket for brackets (open or close) or 0 for runes that are not brackets.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:75
func resolvePairedBrackets(s *isolatingRunSequence) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:75
	_go_fuzz_dep_.CoverTab[69479]++
											p := bracketPairer{
		sos:			s.sos,
		openers:		list.New(),
		codesIsolatedRun:	s.types,
		indexes:		s.indexes,
	}
	dirEmbed := L
	if s.level&1 != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:83
		_go_fuzz_dep_.CoverTab[69481]++
												dirEmbed = R
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:84
		// _ = "end of CoverTab[69481]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:85
		_go_fuzz_dep_.CoverTab[69482]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:85
		// _ = "end of CoverTab[69482]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:85
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:85
	// _ = "end of CoverTab[69479]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:85
	_go_fuzz_dep_.CoverTab[69480]++
											p.locateBrackets(s.p.pairTypes, s.p.pairValues)
											p.resolveBrackets(dirEmbed, s.p.initialTypes)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:87
	// _ = "end of CoverTab[69480]"
}

type bracketPairer struct {
											sos	Class	// direction corresponding to start of sequence

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:112
	openers	*list.List	// list of positions for opening brackets

	// bracket pair positions sorted by location of opening bracket
	pairPositions	bracketPairs

	codesIsolatedRun	[]Class	// directional bidi codes for an isolated run
	indexes			[]int	// array of index values into the original string

}

// matchOpener reports whether characters at given positions form a matching
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:122
// bracket pair.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:124
func (p *bracketPairer) matchOpener(pairValues []rune, opener, closer int) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:124
	_go_fuzz_dep_.CoverTab[69483]++
											return pairValues[p.indexes[opener]] == pairValues[p.indexes[closer]]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:125
	// _ = "end of CoverTab[69483]"
}

const maxPairingDepth = 63

// locateBrackets locates matching bracket pairs according to BD16.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:130
//
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:130
// This implementation uses a linked list instead of a stack, because, while
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:130
// elements are added at the front (like a push) they are not generally removed
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:130
// in atomic 'pop' operations, reducing the benefit of the stack archetype.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:135
func (p *bracketPairer) locateBrackets(pairTypes []bracketType, pairValues []rune) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:135
	_go_fuzz_dep_.CoverTab[69484]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:138
	for i, index := range p.indexes {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:138
		_go_fuzz_dep_.CoverTab[69485]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:141
		if pairTypes[index] == bpNone || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:141
			_go_fuzz_dep_.CoverTab[69487]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:141
			return p.codesIsolatedRun[i] != ON
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:141
			// _ = "end of CoverTab[69487]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:141
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:141
			_go_fuzz_dep_.CoverTab[69488]++

													continue
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:143
			// _ = "end of CoverTab[69488]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:144
			_go_fuzz_dep_.CoverTab[69489]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:144
			// _ = "end of CoverTab[69489]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:144
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:144
		// _ = "end of CoverTab[69485]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:144
		_go_fuzz_dep_.CoverTab[69486]++
												switch pairTypes[index] {
		case bpOpen:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:146
			_go_fuzz_dep_.CoverTab[69490]++

													if p.openers.Len() == maxPairingDepth {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:148
				_go_fuzz_dep_.CoverTab[69495]++
														p.openers.Init()
														return
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:150
				// _ = "end of CoverTab[69495]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:151
				_go_fuzz_dep_.CoverTab[69496]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:151
				// _ = "end of CoverTab[69496]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:151
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:151
			// _ = "end of CoverTab[69490]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:151
			_go_fuzz_dep_.CoverTab[69491]++

													p.openers.PushFront(i)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:153
			// _ = "end of CoverTab[69491]"

		case bpClose:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:155
			_go_fuzz_dep_.CoverTab[69492]++

													count := 0
													for elem := p.openers.Front(); elem != nil; elem = elem.Next() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:158
				_go_fuzz_dep_.CoverTab[69497]++
														count++
														opener := elem.Value.(int)
														if p.matchOpener(pairValues, opener, i) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:161
					_go_fuzz_dep_.CoverTab[69498]++

															p.pairPositions = append(p.pairPositions, bracketPair{opener, i})

															for ; count > 0; count-- {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:165
						_go_fuzz_dep_.CoverTab[69500]++
																p.openers.Remove(p.openers.Front())
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:166
						// _ = "end of CoverTab[69500]"
					}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:167
					// _ = "end of CoverTab[69498]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:167
					_go_fuzz_dep_.CoverTab[69499]++
															break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:168
					// _ = "end of CoverTab[69499]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:169
					_go_fuzz_dep_.CoverTab[69501]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:169
					// _ = "end of CoverTab[69501]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:169
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:169
				// _ = "end of CoverTab[69497]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:170
			// _ = "end of CoverTab[69492]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:170
			_go_fuzz_dep_.CoverTab[69493]++
													sort.Sort(p.pairPositions)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:171
			// _ = "end of CoverTab[69493]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:171
		default:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:171
			_go_fuzz_dep_.CoverTab[69494]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:171
			// _ = "end of CoverTab[69494]"

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:174
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:174
		// _ = "end of CoverTab[69486]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:175
	// _ = "end of CoverTab[69484]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:229
// getStrongTypeN0 maps character's directional code to strong type as required
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:229
// by rule N0.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:229
//
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:229
// TODO: have separate type for "strong" directionality.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:233
func (p *bracketPairer) getStrongTypeN0(index int) Class {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:233
	_go_fuzz_dep_.CoverTab[69502]++
											switch p.codesIsolatedRun[index] {

	case EN, AN, AL, R:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:236
		_go_fuzz_dep_.CoverTab[69503]++
												return R
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:237
		// _ = "end of CoverTab[69503]"
	case L:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:238
		_go_fuzz_dep_.CoverTab[69504]++
												return L
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:239
		// _ = "end of CoverTab[69504]"
	default:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:240
		_go_fuzz_dep_.CoverTab[69505]++
												return ON
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:241
		// _ = "end of CoverTab[69505]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:242
	// _ = "end of CoverTab[69502]"
}

// classifyPairContent reports the strong types contained inside a Bracket Pair,
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:245
// assuming the given embedding direction.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:245
//
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:245
// It returns ON if no strong type is found. If a single strong type is found,
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:245
// it returns this type. Otherwise it returns the embedding direction.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:245
//
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:245
// TODO: use separate type for "strong" directionality.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:252
func (p *bracketPairer) classifyPairContent(loc bracketPair, dirEmbed Class) Class {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:252
	_go_fuzz_dep_.CoverTab[69506]++
											dirOpposite := ON
											for i := loc.opener + 1; i < loc.closer; i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:254
		_go_fuzz_dep_.CoverTab[69508]++
												dir := p.getStrongTypeN0(i)
												if dir == ON {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:256
			_go_fuzz_dep_.CoverTab[69511]++
													continue
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:257
			// _ = "end of CoverTab[69511]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:258
			_go_fuzz_dep_.CoverTab[69512]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:258
			// _ = "end of CoverTab[69512]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:258
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:258
		// _ = "end of CoverTab[69508]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:258
		_go_fuzz_dep_.CoverTab[69509]++
												if dir == dirEmbed {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:259
			_go_fuzz_dep_.CoverTab[69513]++
													return dir
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:260
			// _ = "end of CoverTab[69513]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:261
			_go_fuzz_dep_.CoverTab[69514]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:261
			// _ = "end of CoverTab[69514]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:261
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:261
		// _ = "end of CoverTab[69509]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:261
		_go_fuzz_dep_.CoverTab[69510]++
												dirOpposite = dir
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:262
		// _ = "end of CoverTab[69510]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:263
	// _ = "end of CoverTab[69506]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:263
	_go_fuzz_dep_.CoverTab[69507]++

											return dirOpposite
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:265
	// _ = "end of CoverTab[69507]"
}

// classBeforePair determines which strong types are present before a Bracket
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:268
// Pair. Return R or L if strong type found, otherwise ON.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:270
func (p *bracketPairer) classBeforePair(loc bracketPair) Class {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:270
	_go_fuzz_dep_.CoverTab[69515]++
											for i := loc.opener - 1; i >= 0; i-- {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:271
		_go_fuzz_dep_.CoverTab[69517]++
												if dir := p.getStrongTypeN0(i); dir != ON {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:272
			_go_fuzz_dep_.CoverTab[69518]++
													return dir
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:273
			// _ = "end of CoverTab[69518]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:274
			_go_fuzz_dep_.CoverTab[69519]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:274
			// _ = "end of CoverTab[69519]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:274
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:274
		// _ = "end of CoverTab[69517]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:275
	// _ = "end of CoverTab[69515]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:275
	_go_fuzz_dep_.CoverTab[69516]++

											return p.sos
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:277
	// _ = "end of CoverTab[69516]"
}

// assignBracketType implements rule N0 for a single bracket pair.
func (p *bracketPairer) assignBracketType(loc bracketPair, dirEmbed Class, initialTypes []Class) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:281
	_go_fuzz_dep_.CoverTab[69520]++

											dirPair := p.classifyPairContent(loc, dirEmbed)

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:289
	if dirPair == ON {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:289
		_go_fuzz_dep_.CoverTab[69523]++
												return
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:290
		// _ = "end of CoverTab[69523]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:291
		_go_fuzz_dep_.CoverTab[69524]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:291
		// _ = "end of CoverTab[69524]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:291
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:291
	// _ = "end of CoverTab[69520]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:291
	_go_fuzz_dep_.CoverTab[69521]++

											if dirPair != dirEmbed {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:293
		_go_fuzz_dep_.CoverTab[69525]++

												dirPair = p.classBeforePair(loc)
												if dirPair == dirEmbed || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:296
			_go_fuzz_dep_.CoverTab[69526]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:296
			return dirPair == ON
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:296
			// _ = "end of CoverTab[69526]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:296
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:296
			_go_fuzz_dep_.CoverTab[69527]++

													dirPair = dirEmbed
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:298
			// _ = "end of CoverTab[69527]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:299
			_go_fuzz_dep_.CoverTab[69528]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:299
			// _ = "end of CoverTab[69528]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:299
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:299
		// _ = "end of CoverTab[69525]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:300
		_go_fuzz_dep_.CoverTab[69529]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:300
		// _ = "end of CoverTab[69529]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:300
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:300
	// _ = "end of CoverTab[69521]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:300
	_go_fuzz_dep_.CoverTab[69522]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:306
	p.setBracketsToType(loc, dirPair, initialTypes)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:306
	// _ = "end of CoverTab[69522]"
}

func (p *bracketPairer) setBracketsToType(loc bracketPair, dirPair Class, initialTypes []Class) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:309
	_go_fuzz_dep_.CoverTab[69530]++
											p.codesIsolatedRun[loc.opener] = dirPair
											p.codesIsolatedRun[loc.closer] = dirPair

											for i := loc.opener + 1; i < loc.closer; i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:313
		_go_fuzz_dep_.CoverTab[69532]++
												index := p.indexes[i]
												if initialTypes[index] != NSM {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:315
			_go_fuzz_dep_.CoverTab[69534]++
													break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:316
			// _ = "end of CoverTab[69534]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:317
			_go_fuzz_dep_.CoverTab[69535]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:317
			// _ = "end of CoverTab[69535]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:317
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:317
		// _ = "end of CoverTab[69532]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:317
		_go_fuzz_dep_.CoverTab[69533]++
												p.codesIsolatedRun[i] = dirPair
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:318
		// _ = "end of CoverTab[69533]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:319
	// _ = "end of CoverTab[69530]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:319
	_go_fuzz_dep_.CoverTab[69531]++

											for i := loc.closer + 1; i < len(p.indexes); i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:321
		_go_fuzz_dep_.CoverTab[69536]++
												index := p.indexes[i]
												if initialTypes[index] != NSM {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:323
			_go_fuzz_dep_.CoverTab[69538]++
													break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:324
			// _ = "end of CoverTab[69538]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:325
			_go_fuzz_dep_.CoverTab[69539]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:325
			// _ = "end of CoverTab[69539]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:325
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:325
		// _ = "end of CoverTab[69536]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:325
		_go_fuzz_dep_.CoverTab[69537]++
												p.codesIsolatedRun[i] = dirPair
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:326
		// _ = "end of CoverTab[69537]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:327
	// _ = "end of CoverTab[69531]"
}

// resolveBrackets implements rule N0 for a list of pairs.
func (p *bracketPairer) resolveBrackets(dirEmbed Class, initialTypes []Class) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:331
	_go_fuzz_dep_.CoverTab[69540]++
											for _, loc := range p.pairPositions {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:332
		_go_fuzz_dep_.CoverTab[69541]++
												p.assignBracketType(loc, dirEmbed, initialTypes)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:333
		// _ = "end of CoverTab[69541]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:334
	// _ = "end of CoverTab[69540]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:335
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/bracket.go:335
var _ = _go_fuzz_dep_.CoverTab
