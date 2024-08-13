// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:5
package bidi

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:5
import (
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:5
)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:5
import (
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:5
)

import (
	"container/list"
	"fmt"
	"sort"
)

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:35
// Bidi_Paired_Bracket_Type
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:35
// BD14. An opening paired bracket is a character whose
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:35
// Bidi_Paired_Bracket_Type property value is Open.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:35
//
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:35
// BD15. A closing paired bracket is a character whose
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:35
// Bidi_Paired_Bracket_Type property value is Close.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:41
type bracketType byte

const (
	bpNone	bracketType	= iota
	bpOpen
	bpClose
)

// bracketPair holds a pair of index values for opening and closing bracket
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:49
// location of a bracket pair.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:51
type bracketPair struct {
	opener	int
	closer	int
}

func (b *bracketPair) String() string {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:56
	_go_fuzz_dep_.CoverTab[32194]++
										return fmt.Sprintf("(%v, %v)", b.opener, b.closer)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:57
	// _ = "end of CoverTab[32194]"
}

// bracketPairs is a slice of bracketPairs with a sort.Interface implementation.
type bracketPairs []bracketPair

func (b bracketPairs) Len() int {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:63
	_go_fuzz_dep_.CoverTab[32195]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:63
	return len(b)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:63
	// _ = "end of CoverTab[32195]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:63
}
func (b bracketPairs) Swap(i, j int) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:64
	_go_fuzz_dep_.CoverTab[32196]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:64
	b[i], b[j] = b[j], b[i]
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:64
	// _ = "end of CoverTab[32196]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:64
}
func (b bracketPairs) Less(i, j int) bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:65
	_go_fuzz_dep_.CoverTab[32197]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:65
	return b[i].opener < b[j].opener
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:65
	// _ = "end of CoverTab[32197]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:65
}

// resolvePairedBrackets runs the paired bracket part of the UBA algorithm.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:67
//
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:67
// For each rune, it takes the indexes into the original string, the class the
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:67
// bracket type (in pairTypes) and the bracket identifier (pairValues). It also
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:67
// takes the direction type for the start-of-sentence and the embedding level.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:67
//
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:67
// The identifiers for bracket types are the rune of the canonicalized opening
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:67
// bracket for brackets (open or close) or 0 for runes that are not brackets.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:75
func resolvePairedBrackets(s *isolatingRunSequence) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:75
	_go_fuzz_dep_.CoverTab[32198]++
										p := bracketPairer{
		sos:			s.sos,
		openers:		list.New(),
		codesIsolatedRun:	s.types,
		indexes:		s.indexes,
	}
	dirEmbed := L
	if s.level&1 != 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:83
		_go_fuzz_dep_.CoverTab[32200]++
											dirEmbed = R
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:84
		// _ = "end of CoverTab[32200]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:85
		_go_fuzz_dep_.CoverTab[32201]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:85
		// _ = "end of CoverTab[32201]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:85
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:85
	// _ = "end of CoverTab[32198]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:85
	_go_fuzz_dep_.CoverTab[32199]++
										p.locateBrackets(s.p.pairTypes, s.p.pairValues)
										p.resolveBrackets(dirEmbed, s.p.initialTypes)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:87
	// _ = "end of CoverTab[32199]"
}

type bracketPairer struct {
										sos	Class	// direction corresponding to start of sequence

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:112
	openers	*list.List	// list of positions for opening brackets

	// bracket pair positions sorted by location of opening bracket
	pairPositions	bracketPairs

	codesIsolatedRun	[]Class	// directional bidi codes for an isolated run
	indexes			[]int	// array of index values into the original string

}

// matchOpener reports whether characters at given positions form a matching
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:122
// bracket pair.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:124
func (p *bracketPairer) matchOpener(pairValues []rune, opener, closer int) bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:124
	_go_fuzz_dep_.CoverTab[32202]++
										return pairValues[p.indexes[opener]] == pairValues[p.indexes[closer]]
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:125
	// _ = "end of CoverTab[32202]"
}

const maxPairingDepth = 63

// locateBrackets locates matching bracket pairs according to BD16.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:130
//
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:130
// This implementation uses a linked list instead of a stack, because, while
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:130
// elements are added at the front (like a push) they are not generally removed
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:130
// in atomic 'pop' operations, reducing the benefit of the stack archetype.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:135
func (p *bracketPairer) locateBrackets(pairTypes []bracketType, pairValues []rune) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:135
	_go_fuzz_dep_.CoverTab[32203]++

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:138
	for i, index := range p.indexes {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:138
		_go_fuzz_dep_.CoverTab[32204]++

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:141
		if pairTypes[index] == bpNone || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:141
			_go_fuzz_dep_.CoverTab[32206]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:141
			return p.codesIsolatedRun[i] != ON
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:141
			// _ = "end of CoverTab[32206]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:141
		}() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:141
			_go_fuzz_dep_.CoverTab[32207]++

												continue
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:143
			// _ = "end of CoverTab[32207]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:144
			_go_fuzz_dep_.CoverTab[32208]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:144
			// _ = "end of CoverTab[32208]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:144
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:144
		// _ = "end of CoverTab[32204]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:144
		_go_fuzz_dep_.CoverTab[32205]++
											switch pairTypes[index] {
		case bpOpen:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:146
			_go_fuzz_dep_.CoverTab[32209]++

												if p.openers.Len() == maxPairingDepth {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:148
				_go_fuzz_dep_.CoverTab[32214]++
													p.openers.Init()
													return
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:150
				// _ = "end of CoverTab[32214]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:151
				_go_fuzz_dep_.CoverTab[32215]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:151
				// _ = "end of CoverTab[32215]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:151
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:151
			// _ = "end of CoverTab[32209]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:151
			_go_fuzz_dep_.CoverTab[32210]++

												p.openers.PushFront(i)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:153
			// _ = "end of CoverTab[32210]"

		case bpClose:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:155
			_go_fuzz_dep_.CoverTab[32211]++

												count := 0
												for elem := p.openers.Front(); elem != nil; elem = elem.Next() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:158
				_go_fuzz_dep_.CoverTab[32216]++
													count++
													opener := elem.Value.(int)
													if p.matchOpener(pairValues, opener, i) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:161
					_go_fuzz_dep_.CoverTab[32217]++

														p.pairPositions = append(p.pairPositions, bracketPair{opener, i})

														for ; count > 0; count-- {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:165
						_go_fuzz_dep_.CoverTab[32219]++
															p.openers.Remove(p.openers.Front())
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:166
						// _ = "end of CoverTab[32219]"
					}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:167
					// _ = "end of CoverTab[32217]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:167
					_go_fuzz_dep_.CoverTab[32218]++
														break
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:168
					// _ = "end of CoverTab[32218]"
				} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:169
					_go_fuzz_dep_.CoverTab[32220]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:169
					// _ = "end of CoverTab[32220]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:169
				}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:169
				// _ = "end of CoverTab[32216]"
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:170
			// _ = "end of CoverTab[32211]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:170
			_go_fuzz_dep_.CoverTab[32212]++
												sort.Sort(p.pairPositions)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:171
			// _ = "end of CoverTab[32212]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:171
		default:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:171
			_go_fuzz_dep_.CoverTab[32213]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:171
			// _ = "end of CoverTab[32213]"

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:174
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:174
		// _ = "end of CoverTab[32205]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:175
	// _ = "end of CoverTab[32203]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:229
// getStrongTypeN0 maps character's directional code to strong type as required
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:229
// by rule N0.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:229
//
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:229
// TODO: have separate type for "strong" directionality.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:233
func (p *bracketPairer) getStrongTypeN0(index int) Class {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:233
	_go_fuzz_dep_.CoverTab[32221]++
										switch p.codesIsolatedRun[index] {

	case EN, AN, AL, R:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:236
		_go_fuzz_dep_.CoverTab[32222]++
											return R
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:237
		// _ = "end of CoverTab[32222]"
	case L:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:238
		_go_fuzz_dep_.CoverTab[32223]++
											return L
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:239
		// _ = "end of CoverTab[32223]"
	default:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:240
		_go_fuzz_dep_.CoverTab[32224]++
											return ON
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:241
		// _ = "end of CoverTab[32224]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:242
	// _ = "end of CoverTab[32221]"
}

// classifyPairContent reports the strong types contained inside a Bracket Pair,
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:245
// assuming the given embedding direction.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:245
//
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:245
// It returns ON if no strong type is found. If a single strong type is found,
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:245
// it returns this type. Otherwise it returns the embedding direction.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:245
//
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:245
// TODO: use separate type for "strong" directionality.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:252
func (p *bracketPairer) classifyPairContent(loc bracketPair, dirEmbed Class) Class {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:252
	_go_fuzz_dep_.CoverTab[32225]++
										dirOpposite := ON
										for i := loc.opener + 1; i < loc.closer; i++ {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:254
		_go_fuzz_dep_.CoverTab[32227]++
											dir := p.getStrongTypeN0(i)
											if dir == ON {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:256
			_go_fuzz_dep_.CoverTab[32230]++
												continue
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:257
			// _ = "end of CoverTab[32230]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:258
			_go_fuzz_dep_.CoverTab[32231]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:258
			// _ = "end of CoverTab[32231]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:258
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:258
		// _ = "end of CoverTab[32227]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:258
		_go_fuzz_dep_.CoverTab[32228]++
											if dir == dirEmbed {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:259
			_go_fuzz_dep_.CoverTab[32232]++
												return dir
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:260
			// _ = "end of CoverTab[32232]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:261
			_go_fuzz_dep_.CoverTab[32233]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:261
			// _ = "end of CoverTab[32233]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:261
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:261
		// _ = "end of CoverTab[32228]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:261
		_go_fuzz_dep_.CoverTab[32229]++
											dirOpposite = dir
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:262
		// _ = "end of CoverTab[32229]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:263
	// _ = "end of CoverTab[32225]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:263
	_go_fuzz_dep_.CoverTab[32226]++

										return dirOpposite
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:265
	// _ = "end of CoverTab[32226]"
}

// classBeforePair determines which strong types are present before a Bracket
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:268
// Pair. Return R or L if strong type found, otherwise ON.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:270
func (p *bracketPairer) classBeforePair(loc bracketPair) Class {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:270
	_go_fuzz_dep_.CoverTab[32234]++
										for i := loc.opener - 1; i >= 0; i-- {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:271
		_go_fuzz_dep_.CoverTab[32236]++
											if dir := p.getStrongTypeN0(i); dir != ON {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:272
			_go_fuzz_dep_.CoverTab[32237]++
												return dir
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:273
			// _ = "end of CoverTab[32237]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:274
			_go_fuzz_dep_.CoverTab[32238]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:274
			// _ = "end of CoverTab[32238]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:274
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:274
		// _ = "end of CoverTab[32236]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:275
	// _ = "end of CoverTab[32234]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:275
	_go_fuzz_dep_.CoverTab[32235]++

										return p.sos
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:277
	// _ = "end of CoverTab[32235]"
}

// assignBracketType implements rule N0 for a single bracket pair.
func (p *bracketPairer) assignBracketType(loc bracketPair, dirEmbed Class, initialTypes []Class) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:281
	_go_fuzz_dep_.CoverTab[32239]++

										dirPair := p.classifyPairContent(loc, dirEmbed)

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:289
	if dirPair == ON {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:289
		_go_fuzz_dep_.CoverTab[32242]++
											return
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:290
		// _ = "end of CoverTab[32242]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:291
		_go_fuzz_dep_.CoverTab[32243]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:291
		// _ = "end of CoverTab[32243]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:291
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:291
	// _ = "end of CoverTab[32239]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:291
	_go_fuzz_dep_.CoverTab[32240]++

										if dirPair != dirEmbed {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:293
		_go_fuzz_dep_.CoverTab[32244]++

											dirPair = p.classBeforePair(loc)
											if dirPair == dirEmbed || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:296
			_go_fuzz_dep_.CoverTab[32245]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:296
			return dirPair == ON
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:296
			// _ = "end of CoverTab[32245]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:296
		}() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:296
			_go_fuzz_dep_.CoverTab[32246]++

												dirPair = dirEmbed
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:298
			// _ = "end of CoverTab[32246]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:299
			_go_fuzz_dep_.CoverTab[32247]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:299
			// _ = "end of CoverTab[32247]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:299
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:299
		// _ = "end of CoverTab[32244]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:300
		_go_fuzz_dep_.CoverTab[32248]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:300
		// _ = "end of CoverTab[32248]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:300
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:300
	// _ = "end of CoverTab[32240]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:300
	_go_fuzz_dep_.CoverTab[32241]++

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:306
	p.setBracketsToType(loc, dirPair, initialTypes)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:306
	// _ = "end of CoverTab[32241]"
}

func (p *bracketPairer) setBracketsToType(loc bracketPair, dirPair Class, initialTypes []Class) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:309
	_go_fuzz_dep_.CoverTab[32249]++
										p.codesIsolatedRun[loc.opener] = dirPair
										p.codesIsolatedRun[loc.closer] = dirPair

										for i := loc.opener + 1; i < loc.closer; i++ {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:313
		_go_fuzz_dep_.CoverTab[32251]++
											index := p.indexes[i]
											if initialTypes[index] != NSM {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:315
			_go_fuzz_dep_.CoverTab[32253]++
												break
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:316
			// _ = "end of CoverTab[32253]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:317
			_go_fuzz_dep_.CoverTab[32254]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:317
			// _ = "end of CoverTab[32254]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:317
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:317
		// _ = "end of CoverTab[32251]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:317
		_go_fuzz_dep_.CoverTab[32252]++
											p.codesIsolatedRun[i] = dirPair
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:318
		// _ = "end of CoverTab[32252]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:319
	// _ = "end of CoverTab[32249]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:319
	_go_fuzz_dep_.CoverTab[32250]++

										for i := loc.closer + 1; i < len(p.indexes); i++ {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:321
		_go_fuzz_dep_.CoverTab[32255]++
											index := p.indexes[i]
											if initialTypes[index] != NSM {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:323
			_go_fuzz_dep_.CoverTab[32257]++
												break
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:324
			// _ = "end of CoverTab[32257]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:325
			_go_fuzz_dep_.CoverTab[32258]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:325
			// _ = "end of CoverTab[32258]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:325
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:325
		// _ = "end of CoverTab[32255]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:325
		_go_fuzz_dep_.CoverTab[32256]++
											p.codesIsolatedRun[i] = dirPair
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:326
		// _ = "end of CoverTab[32256]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:327
	// _ = "end of CoverTab[32250]"
}

// resolveBrackets implements rule N0 for a list of pairs.
func (p *bracketPairer) resolveBrackets(dirEmbed Class, initialTypes []Class) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:331
	_go_fuzz_dep_.CoverTab[32259]++
										for _, loc := range p.pairPositions {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:332
		_go_fuzz_dep_.CoverTab[32260]++
											p.assignBracketType(loc, dirEmbed, initialTypes)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:333
		// _ = "end of CoverTab[32260]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:334
	// _ = "end of CoverTab[32259]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:335
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/bracket.go:335
var _ = _go_fuzz_dep_.CoverTab
