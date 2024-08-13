// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:5
package bidi

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:5
)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:5
)

import (
	"fmt"
	"log"
)

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:53
// level is the embedding level of a character. Even embedding levels indicate
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:53
// left-to-right order and odd levels indicate right-to-left order. The special
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:53
// level of -1 is reserved for undefined order.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:56
type level int8

const implicitLevel level = -1

// in returns if x is equal to any of the values in set.
func (c Class) in(set ...Class) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:61
	_go_fuzz_dep_.CoverTab[69542]++
											for _, s := range set {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:62
		_go_fuzz_dep_.CoverTab[69544]++
												if c == s {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:63
			_go_fuzz_dep_.CoverTab[69545]++
													return true
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:64
			// _ = "end of CoverTab[69545]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:65
			_go_fuzz_dep_.CoverTab[69546]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:65
			// _ = "end of CoverTab[69546]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:65
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:65
		// _ = "end of CoverTab[69544]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:66
	// _ = "end of CoverTab[69542]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:66
	_go_fuzz_dep_.CoverTab[69543]++
											return false
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:67
	// _ = "end of CoverTab[69543]"
}

// A paragraph contains the state of a paragraph.
type paragraph struct {
	initialTypes	[]Class

	// Arrays of properties needed for paired bracket evaluation in N0
	pairTypes	[]bracketType	// paired Bracket types for paragraph
	pairValues	[]rune		// rune for opening bracket or pbOpen and pbClose; 0 for pbNone

	embeddingLevel	level	// default: = implicitLevel;

	// at the paragraph levels
	resultTypes	[]Class
	resultLevels	[]level

	// Index of matching PDI for isolate initiator characters. For other
	// characters, the value of matchingPDI will be set to -1. For isolate
	// initiators with no matching PDI, matchingPDI will be set to the length of
	// the input string.
	matchingPDI	[]int

	// Index of matching isolate initiator for PDI characters. For other
	// characters, and for PDIs with no matching isolate initiator, the value of
	// matchingIsolateInitiator will be set to -1.
	matchingIsolateInitiator	[]int
}

// newParagraph initializes a paragraph. The user needs to supply a few arrays
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:96
// corresponding to the preprocessed text input. The types correspond to the
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:96
// Unicode BiDi classes for each rune. pairTypes indicates the bracket type for
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:96
// each rune. pairValues provides a unique bracket class identifier for each
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:96
// rune (suggested is the rune of the open bracket for opening and matching
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:96
// close brackets, after normalization). The embedding levels are optional, but
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:96
// may be supplied to encode embedding levels of styled text.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:103
func newParagraph(types []Class, pairTypes []bracketType, pairValues []rune, levels level) (*paragraph, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:103
	_go_fuzz_dep_.CoverTab[69547]++
											var err error
											if err = validateTypes(types); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:105
		_go_fuzz_dep_.CoverTab[69552]++
												return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:106
		// _ = "end of CoverTab[69552]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:107
		_go_fuzz_dep_.CoverTab[69553]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:107
		// _ = "end of CoverTab[69553]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:107
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:107
	// _ = "end of CoverTab[69547]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:107
	_go_fuzz_dep_.CoverTab[69548]++
											if err = validatePbTypes(pairTypes); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:108
		_go_fuzz_dep_.CoverTab[69554]++
												return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:109
		// _ = "end of CoverTab[69554]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:110
		_go_fuzz_dep_.CoverTab[69555]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:110
		// _ = "end of CoverTab[69555]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:110
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:110
	// _ = "end of CoverTab[69548]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:110
	_go_fuzz_dep_.CoverTab[69549]++
											if err = validatePbValues(pairValues, pairTypes); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:111
		_go_fuzz_dep_.CoverTab[69556]++
												return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:112
		// _ = "end of CoverTab[69556]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:113
		_go_fuzz_dep_.CoverTab[69557]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:113
		// _ = "end of CoverTab[69557]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:113
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:113
	// _ = "end of CoverTab[69549]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:113
	_go_fuzz_dep_.CoverTab[69550]++
											if err = validateParagraphEmbeddingLevel(levels); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:114
		_go_fuzz_dep_.CoverTab[69558]++
												return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:115
		// _ = "end of CoverTab[69558]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:116
		_go_fuzz_dep_.CoverTab[69559]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:116
		// _ = "end of CoverTab[69559]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:116
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:116
	// _ = "end of CoverTab[69550]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:116
	_go_fuzz_dep_.CoverTab[69551]++

											p := &paragraph{
		initialTypes:	append([]Class(nil), types...),
		embeddingLevel:	levels,

		pairTypes:	pairTypes,
		pairValues:	pairValues,

		resultTypes:	append([]Class(nil), types...),
	}
											p.run()
											return p, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:128
	// _ = "end of CoverTab[69551]"
}

func (p *paragraph) Len() int {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:131
	_go_fuzz_dep_.CoverTab[69560]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:131
	return len(p.initialTypes)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:131
	// _ = "end of CoverTab[69560]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:131
}

// The algorithm. Does not include line-based processing (Rules L1, L2).
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:133
// These are applied later in the line-based phase of the algorithm.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:135
func (p *paragraph) run() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:135
	_go_fuzz_dep_.CoverTab[69561]++
											p.determineMatchingIsolates()

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:142
	if p.embeddingLevel == implicitLevel {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:142
		_go_fuzz_dep_.CoverTab[69564]++
												p.embeddingLevel = p.determineParagraphEmbeddingLevel(0, p.Len())
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:143
		// _ = "end of CoverTab[69564]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:144
		_go_fuzz_dep_.CoverTab[69565]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:144
		// _ = "end of CoverTab[69565]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:144
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:144
	// _ = "end of CoverTab[69561]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:144
	_go_fuzz_dep_.CoverTab[69562]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:147
	p.resultLevels = make([]level, p.Len())
											setLevels(p.resultLevels, p.embeddingLevel)

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:152
	p.determineExplicitEmbeddingLevels()

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:162
	for _, seq := range p.determineIsolatingRunSequences() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:162
		_go_fuzz_dep_.CoverTab[69566]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:165
		seq.resolveWeakTypes()

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:169
		resolvePairedBrackets(seq)

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:173
		seq.resolveNeutralTypes()

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:177
		seq.resolveImplicitLevels()

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:180
		seq.applyLevelsAndTypes()
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:180
		// _ = "end of CoverTab[69566]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:181
	// _ = "end of CoverTab[69562]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:181
	_go_fuzz_dep_.CoverTab[69563]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:186
	p.assignLevelsToCharactersRemovedByX9()
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:186
	// _ = "end of CoverTab[69563]"
}

// determineMatchingIsolates determines the matching PDI for each isolate
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:189
// initiator and vice versa.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:189
//
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:189
// Definition BD9.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:189
//
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:189
// At the end of this function:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:189
//
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:189
//   - The member variable matchingPDI is set to point to the index of the
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:189
//     matching PDI character for each isolate initiator character. If there is
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:189
//     no matching PDI, it is set to the length of the input text. For other
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:189
//     characters, it is set to -1.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:189
//   - The member variable matchingIsolateInitiator is set to point to the
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:189
//     index of the matching isolate initiator character for each PDI character.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:189
//     If there is no matching isolate initiator, or the character is not a PDI,
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:189
//     it is set to -1.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:204
func (p *paragraph) determineMatchingIsolates() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:204
	_go_fuzz_dep_.CoverTab[69567]++
											p.matchingPDI = make([]int, p.Len())
											p.matchingIsolateInitiator = make([]int, p.Len())

											for i := range p.matchingIsolateInitiator {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:208
		_go_fuzz_dep_.CoverTab[69569]++
												p.matchingIsolateInitiator[i] = -1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:209
		// _ = "end of CoverTab[69569]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:210
	// _ = "end of CoverTab[69567]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:210
	_go_fuzz_dep_.CoverTab[69568]++

											for i := range p.matchingPDI {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:212
		_go_fuzz_dep_.CoverTab[69570]++
												p.matchingPDI[i] = -1

												if t := p.resultTypes[i]; t.in(LRI, RLI, FSI) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:215
			_go_fuzz_dep_.CoverTab[69571]++
													depthCounter := 1
													for j := i + 1; j < p.Len(); j++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:217
				_go_fuzz_dep_.CoverTab[69573]++
														if u := p.resultTypes[j]; u.in(LRI, RLI, FSI) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:218
					_go_fuzz_dep_.CoverTab[69574]++
															depthCounter++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:219
					// _ = "end of CoverTab[69574]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:220
					_go_fuzz_dep_.CoverTab[69575]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:220
					if u == PDI {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:220
						_go_fuzz_dep_.CoverTab[69576]++
																if depthCounter--; depthCounter == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:221
							_go_fuzz_dep_.CoverTab[69577]++
																	p.matchingPDI[i] = j
																	p.matchingIsolateInitiator[j] = i
																	break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:224
							// _ = "end of CoverTab[69577]"
						} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:225
							_go_fuzz_dep_.CoverTab[69578]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:225
							// _ = "end of CoverTab[69578]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:225
						}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:225
						// _ = "end of CoverTab[69576]"
					} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:226
						_go_fuzz_dep_.CoverTab[69579]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:226
						// _ = "end of CoverTab[69579]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:226
					}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:226
					// _ = "end of CoverTab[69575]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:226
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:226
				// _ = "end of CoverTab[69573]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:227
			// _ = "end of CoverTab[69571]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:227
			_go_fuzz_dep_.CoverTab[69572]++
													if p.matchingPDI[i] == -1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:228
				_go_fuzz_dep_.CoverTab[69580]++
														p.matchingPDI[i] = p.Len()
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:229
				// _ = "end of CoverTab[69580]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:230
				_go_fuzz_dep_.CoverTab[69581]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:230
				// _ = "end of CoverTab[69581]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:230
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:230
			// _ = "end of CoverTab[69572]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:231
			_go_fuzz_dep_.CoverTab[69582]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:231
			// _ = "end of CoverTab[69582]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:231
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:231
		// _ = "end of CoverTab[69570]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:232
	// _ = "end of CoverTab[69568]"
}

// determineParagraphEmbeddingLevel reports the resolved paragraph direction of
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:235
// the substring limited by the given range [start, end).
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:235
//
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:235
// Determines the paragraph level based on rules P2, P3. This is also used
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:235
// in rule X5c to find if an FSI should resolve to LRI or RLI.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:240
func (p *paragraph) determineParagraphEmbeddingLevel(start, end int) level {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:240
	_go_fuzz_dep_.CoverTab[69583]++
											var strongType Class = unknownClass

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:244
	for i := start; i < end; i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:244
		_go_fuzz_dep_.CoverTab[69585]++
												if t := p.resultTypes[i]; t.in(L, AL, R) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:245
			_go_fuzz_dep_.CoverTab[69586]++
													strongType = t
													break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:247
			// _ = "end of CoverTab[69586]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:248
			_go_fuzz_dep_.CoverTab[69587]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:248
			if t.in(FSI, LRI, RLI) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:248
				_go_fuzz_dep_.CoverTab[69588]++
														i = p.matchingPDI[i]
														if i > end {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:250
					_go_fuzz_dep_.CoverTab[69589]++
															log.Panic("assert (i <= end)")
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:251
					// _ = "end of CoverTab[69589]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:252
					_go_fuzz_dep_.CoverTab[69590]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:252
					// _ = "end of CoverTab[69590]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:252
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:252
				// _ = "end of CoverTab[69588]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:253
				_go_fuzz_dep_.CoverTab[69591]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:253
				// _ = "end of CoverTab[69591]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:253
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:253
			// _ = "end of CoverTab[69587]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:253
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:253
		// _ = "end of CoverTab[69585]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:254
	// _ = "end of CoverTab[69583]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:254
	_go_fuzz_dep_.CoverTab[69584]++

											switch strongType {
	case unknownClass:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:257
		_go_fuzz_dep_.CoverTab[69592]++

												return 0
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:259
		// _ = "end of CoverTab[69592]"
	case L:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:260
		_go_fuzz_dep_.CoverTab[69593]++
												return 0
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:261
		// _ = "end of CoverTab[69593]"
	default:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:262
		_go_fuzz_dep_.CoverTab[69594]++
												return 1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:263
		// _ = "end of CoverTab[69594]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:264
	// _ = "end of CoverTab[69584]"
}

const maxDepth = 125

// This stack will store the embedding levels and override and isolated
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:269
// statuses
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:271
type directionalStatusStack struct {
	stackCounter		int
	embeddingLevelStack	[maxDepth + 1]level
	overrideStatusStack	[maxDepth + 1]Class
	isolateStatusStack	[maxDepth + 1]bool
}

func (s *directionalStatusStack) empty() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:278
	_go_fuzz_dep_.CoverTab[69595]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:278
	s.stackCounter = 0
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:278
	// _ = "end of CoverTab[69595]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:278
}
func (s *directionalStatusStack) pop() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:279
	_go_fuzz_dep_.CoverTab[69596]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:279
	s.stackCounter--
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:279
	// _ = "end of CoverTab[69596]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:279
}
func (s *directionalStatusStack) depth() int {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:280
	_go_fuzz_dep_.CoverTab[69597]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:280
	return s.stackCounter
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:280
	// _ = "end of CoverTab[69597]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:280
}

func (s *directionalStatusStack) push(level level, overrideStatus Class, isolateStatus bool) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:282
	_go_fuzz_dep_.CoverTab[69598]++
											s.embeddingLevelStack[s.stackCounter] = level
											s.overrideStatusStack[s.stackCounter] = overrideStatus
											s.isolateStatusStack[s.stackCounter] = isolateStatus
											s.stackCounter++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:286
	// _ = "end of CoverTab[69598]"
}

func (s *directionalStatusStack) lastEmbeddingLevel() level {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:289
	_go_fuzz_dep_.CoverTab[69599]++
											return s.embeddingLevelStack[s.stackCounter-1]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:290
	// _ = "end of CoverTab[69599]"
}

func (s *directionalStatusStack) lastDirectionalOverrideStatus() Class {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:293
	_go_fuzz_dep_.CoverTab[69600]++
											return s.overrideStatusStack[s.stackCounter-1]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:294
	// _ = "end of CoverTab[69600]"
}

func (s *directionalStatusStack) lastDirectionalIsolateStatus() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:297
	_go_fuzz_dep_.CoverTab[69601]++
											return s.isolateStatusStack[s.stackCounter-1]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:298
	// _ = "end of CoverTab[69601]"
}

// Determine explicit levels using rules X1 - X8
func (p *paragraph) determineExplicitEmbeddingLevels() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:302
	_go_fuzz_dep_.CoverTab[69602]++
											var stack directionalStatusStack
											var overflowIsolateCount, overflowEmbeddingCount, validIsolateCount int

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:307
	stack.push(p.embeddingLevel, ON, false)

	for i, t := range p.resultTypes {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:309
		_go_fuzz_dep_.CoverTab[69603]++

												switch t {
		case RLE, LRE, RLO, LRO, RLI, LRI, FSI:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:312
			_go_fuzz_dep_.CoverTab[69604]++
													isIsolate := t.in(RLI, LRI, FSI)
													isRTL := t.in(RLE, RLO, RLI)

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:317
			if t == FSI {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:317
				_go_fuzz_dep_.CoverTab[69613]++
														isRTL = (p.determineParagraphEmbeddingLevel(i+1, p.matchingPDI[i]) == 1)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:318
				// _ = "end of CoverTab[69613]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:319
				_go_fuzz_dep_.CoverTab[69614]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:319
				// _ = "end of CoverTab[69614]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:319
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:319
			// _ = "end of CoverTab[69604]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:319
			_go_fuzz_dep_.CoverTab[69605]++
													if isIsolate {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:320
				_go_fuzz_dep_.CoverTab[69615]++
														p.resultLevels[i] = stack.lastEmbeddingLevel()
														if stack.lastDirectionalOverrideStatus() != ON {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:322
					_go_fuzz_dep_.CoverTab[69616]++
															p.resultTypes[i] = stack.lastDirectionalOverrideStatus()
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:323
					// _ = "end of CoverTab[69616]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:324
					_go_fuzz_dep_.CoverTab[69617]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:324
					// _ = "end of CoverTab[69617]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:324
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:324
				// _ = "end of CoverTab[69615]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:325
				_go_fuzz_dep_.CoverTab[69618]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:325
				// _ = "end of CoverTab[69618]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:325
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:325
			// _ = "end of CoverTab[69605]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:325
			_go_fuzz_dep_.CoverTab[69606]++

													var newLevel level
													if isRTL {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:328
				_go_fuzz_dep_.CoverTab[69619]++

														newLevel = (stack.lastEmbeddingLevel() + 1) | 1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:330
				// _ = "end of CoverTab[69619]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:331
				_go_fuzz_dep_.CoverTab[69620]++

														newLevel = (stack.lastEmbeddingLevel() + 2) &^ 1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:333
				// _ = "end of CoverTab[69620]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:334
			// _ = "end of CoverTab[69606]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:334
			_go_fuzz_dep_.CoverTab[69607]++

													if newLevel <= maxDepth && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:336
				_go_fuzz_dep_.CoverTab[69621]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:336
				return overflowIsolateCount == 0
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:336
				// _ = "end of CoverTab[69621]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:336
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:336
				_go_fuzz_dep_.CoverTab[69622]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:336
				return overflowEmbeddingCount == 0
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:336
				// _ = "end of CoverTab[69622]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:336
			}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:336
				_go_fuzz_dep_.CoverTab[69623]++
														if isIsolate {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:337
					_go_fuzz_dep_.CoverTab[69626]++
															validIsolateCount++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:338
					// _ = "end of CoverTab[69626]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:339
					_go_fuzz_dep_.CoverTab[69627]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:339
					// _ = "end of CoverTab[69627]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:339
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:339
				// _ = "end of CoverTab[69623]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:339
				_go_fuzz_dep_.CoverTab[69624]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:344
				switch t {
				case LRO:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:345
					_go_fuzz_dep_.CoverTab[69628]++
															stack.push(newLevel, L, isIsolate)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:346
					// _ = "end of CoverTab[69628]"
				case RLO:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:347
					_go_fuzz_dep_.CoverTab[69629]++
															stack.push(newLevel, R, isIsolate)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:348
					// _ = "end of CoverTab[69629]"
				default:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:349
					_go_fuzz_dep_.CoverTab[69630]++
															stack.push(newLevel, ON, isIsolate)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:350
					// _ = "end of CoverTab[69630]"
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:351
				// _ = "end of CoverTab[69624]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:351
				_go_fuzz_dep_.CoverTab[69625]++

														if !isIsolate {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:353
					_go_fuzz_dep_.CoverTab[69631]++
															p.resultLevels[i] = newLevel
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:354
					// _ = "end of CoverTab[69631]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:355
					_go_fuzz_dep_.CoverTab[69632]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:355
					// _ = "end of CoverTab[69632]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:355
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:355
				// _ = "end of CoverTab[69625]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:356
				_go_fuzz_dep_.CoverTab[69633]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:359
				if isIsolate {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:359
					_go_fuzz_dep_.CoverTab[69634]++
															overflowIsolateCount++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:360
					// _ = "end of CoverTab[69634]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:361
					_go_fuzz_dep_.CoverTab[69635]++
															if overflowIsolateCount == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:362
						_go_fuzz_dep_.CoverTab[69636]++
																overflowEmbeddingCount++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:363
						// _ = "end of CoverTab[69636]"
					} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:364
						_go_fuzz_dep_.CoverTab[69637]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:364
						// _ = "end of CoverTab[69637]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:364
					}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:364
					// _ = "end of CoverTab[69635]"
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:365
				// _ = "end of CoverTab[69633]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:366
			// _ = "end of CoverTab[69607]"

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:369
		case PDI:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:369
			_go_fuzz_dep_.CoverTab[69608]++
													if overflowIsolateCount > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:370
				_go_fuzz_dep_.CoverTab[69638]++
														overflowIsolateCount--
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:371
				// _ = "end of CoverTab[69638]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:372
				_go_fuzz_dep_.CoverTab[69639]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:372
				if validIsolateCount == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:372
					_go_fuzz_dep_.CoverTab[69640]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:372
					// _ = "end of CoverTab[69640]"

				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:374
					_go_fuzz_dep_.CoverTab[69641]++
															overflowEmbeddingCount = 0
															for !stack.lastDirectionalIsolateStatus() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:376
						_go_fuzz_dep_.CoverTab[69643]++
																stack.pop()
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:377
						// _ = "end of CoverTab[69643]"
					}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:378
					// _ = "end of CoverTab[69641]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:378
					_go_fuzz_dep_.CoverTab[69642]++
															stack.pop()
															validIsolateCount--
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:380
					// _ = "end of CoverTab[69642]"
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:381
				// _ = "end of CoverTab[69639]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:381
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:381
			// _ = "end of CoverTab[69608]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:381
			_go_fuzz_dep_.CoverTab[69609]++
													p.resultLevels[i] = stack.lastEmbeddingLevel()
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:382
			// _ = "end of CoverTab[69609]"

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:385
		case PDF:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:385
			_go_fuzz_dep_.CoverTab[69610]++

													p.resultLevels[i] = stack.lastEmbeddingLevel()

													if overflowIsolateCount > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:389
				_go_fuzz_dep_.CoverTab[69644]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:389
				// _ = "end of CoverTab[69644]"

			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:391
				_go_fuzz_dep_.CoverTab[69645]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:391
				if overflowEmbeddingCount > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:391
					_go_fuzz_dep_.CoverTab[69646]++
															overflowEmbeddingCount--
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:392
					// _ = "end of CoverTab[69646]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:393
					_go_fuzz_dep_.CoverTab[69647]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:393
					if !stack.lastDirectionalIsolateStatus() && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:393
						_go_fuzz_dep_.CoverTab[69648]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:393
						return stack.depth() >= 2
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:393
						// _ = "end of CoverTab[69648]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:393
					}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:393
						_go_fuzz_dep_.CoverTab[69649]++
																stack.pop()
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:394
						// _ = "end of CoverTab[69649]"
					} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:395
						_go_fuzz_dep_.CoverTab[69650]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:395
						// _ = "end of CoverTab[69650]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:395
					}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:395
					// _ = "end of CoverTab[69647]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:395
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:395
				// _ = "end of CoverTab[69645]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:395
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:395
			// _ = "end of CoverTab[69610]"

		case B:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:397
			_go_fuzz_dep_.CoverTab[69611]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:402
			stack.empty()
													overflowIsolateCount = 0
													overflowEmbeddingCount = 0
													validIsolateCount = 0
													p.resultLevels[i] = p.embeddingLevel
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:406
			// _ = "end of CoverTab[69611]"

		default:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:408
			_go_fuzz_dep_.CoverTab[69612]++
													p.resultLevels[i] = stack.lastEmbeddingLevel()
													if stack.lastDirectionalOverrideStatus() != ON {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:410
				_go_fuzz_dep_.CoverTab[69651]++
														p.resultTypes[i] = stack.lastDirectionalOverrideStatus()
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:411
				// _ = "end of CoverTab[69651]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:412
				_go_fuzz_dep_.CoverTab[69652]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:412
				// _ = "end of CoverTab[69652]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:412
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:412
			// _ = "end of CoverTab[69612]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:413
		// _ = "end of CoverTab[69603]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:414
	// _ = "end of CoverTab[69602]"
}

type isolatingRunSequence struct {
	p	*paragraph

	indexes	[]int	// indexes to the original string

	types		[]Class	// type of each character using the index
	resolvedLevels	[]level	// resolved levels after application of rules
	level		level
	sos, eos	Class
}

func (i *isolatingRunSequence) Len() int {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:428
	_go_fuzz_dep_.CoverTab[69653]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:428
	return len(i.indexes)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:428
	// _ = "end of CoverTab[69653]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:428
}

func maxLevel(a, b level) level {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:430
	_go_fuzz_dep_.CoverTab[69654]++
											if a > b {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:431
		_go_fuzz_dep_.CoverTab[69656]++
												return a
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:432
		// _ = "end of CoverTab[69656]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:433
		_go_fuzz_dep_.CoverTab[69657]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:433
		// _ = "end of CoverTab[69657]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:433
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:433
	// _ = "end of CoverTab[69654]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:433
	_go_fuzz_dep_.CoverTab[69655]++
											return b
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:434
	// _ = "end of CoverTab[69655]"
}

// Rule X10, second bullet: Determine the start-of-sequence (sos) and end-of-sequence (eos) types,
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:437
// either L or R, for each isolating run sequence.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:439
func (p *paragraph) isolatingRunSequence(indexes []int) *isolatingRunSequence {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:439
	_go_fuzz_dep_.CoverTab[69658]++
											length := len(indexes)
											types := make([]Class, length)
											for i, x := range indexes {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:442
		_go_fuzz_dep_.CoverTab[69663]++
												types[i] = p.resultTypes[x]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:443
		// _ = "end of CoverTab[69663]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:444
	// _ = "end of CoverTab[69658]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:444
	_go_fuzz_dep_.CoverTab[69659]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:447
	prevChar := indexes[0] - 1
	for prevChar >= 0 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:448
		_go_fuzz_dep_.CoverTab[69664]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:448
		return isRemovedByX9(p.initialTypes[prevChar])
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:448
		// _ = "end of CoverTab[69664]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:448
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:448
		_go_fuzz_dep_.CoverTab[69665]++
												prevChar--
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:449
		// _ = "end of CoverTab[69665]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:450
	// _ = "end of CoverTab[69659]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:450
	_go_fuzz_dep_.CoverTab[69660]++
											prevLevel := p.embeddingLevel
											if prevChar >= 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:452
		_go_fuzz_dep_.CoverTab[69666]++
												prevLevel = p.resultLevels[prevChar]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:453
		// _ = "end of CoverTab[69666]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:454
		_go_fuzz_dep_.CoverTab[69667]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:454
		// _ = "end of CoverTab[69667]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:454
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:454
	// _ = "end of CoverTab[69660]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:454
	_go_fuzz_dep_.CoverTab[69661]++

											var succLevel level
											lastType := types[length-1]
											if lastType.in(LRI, RLI, FSI) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:458
		_go_fuzz_dep_.CoverTab[69668]++
												succLevel = p.embeddingLevel
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:459
		// _ = "end of CoverTab[69668]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:460
		_go_fuzz_dep_.CoverTab[69669]++

												limit := indexes[length-1] + 1
												for ; limit < p.Len() && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:463
			_go_fuzz_dep_.CoverTab[69671]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:463
			return isRemovedByX9(p.initialTypes[limit])
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:463
			// _ = "end of CoverTab[69671]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:463
		}(); limit++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:463
			_go_fuzz_dep_.CoverTab[69672]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:463
			// _ = "end of CoverTab[69672]"

		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:465
		// _ = "end of CoverTab[69669]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:465
		_go_fuzz_dep_.CoverTab[69670]++
												succLevel = p.embeddingLevel
												if limit < p.Len() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:467
			_go_fuzz_dep_.CoverTab[69673]++
													succLevel = p.resultLevels[limit]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:468
			// _ = "end of CoverTab[69673]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:469
			_go_fuzz_dep_.CoverTab[69674]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:469
			// _ = "end of CoverTab[69674]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:469
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:469
		// _ = "end of CoverTab[69670]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:470
	// _ = "end of CoverTab[69661]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:470
	_go_fuzz_dep_.CoverTab[69662]++
											level := p.resultLevels[indexes[0]]
											return &isolatingRunSequence{
		p:		p,
		indexes:	indexes,
		types:		types,
		level:		level,
		sos:		typeForLevel(maxLevel(prevLevel, level)),
		eos:		typeForLevel(maxLevel(succLevel, level)),
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:479
	// _ = "end of CoverTab[69662]"
}

// Resolving weak types Rules W1-W7.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:482
//
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:482
// Note that some weak types (EN, AN) remain after this processing is
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:482
// complete.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:486
func (s *isolatingRunSequence) resolveWeakTypes() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:486
	_go_fuzz_dep_.CoverTab[69675]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:489
	s.assertOnly(L, R, AL, EN, ES, ET, AN, CS, B, S, WS, ON, NSM, LRI, RLI, FSI, PDI)

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:493
	precedingCharacterType := s.sos
	for i, t := range s.types {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:494
		_go_fuzz_dep_.CoverTab[69682]++
												if t == NSM {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:495
			_go_fuzz_dep_.CoverTab[69683]++
													s.types[i] = precedingCharacterType
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:496
			// _ = "end of CoverTab[69683]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:497
			_go_fuzz_dep_.CoverTab[69684]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:501
			precedingCharacterType = t
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:501
			// _ = "end of CoverTab[69684]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:502
		// _ = "end of CoverTab[69682]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:503
	// _ = "end of CoverTab[69675]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:503
	_go_fuzz_dep_.CoverTab[69676]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:507
	for i, t := range s.types {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:507
		_go_fuzz_dep_.CoverTab[69685]++
												if t == EN {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:508
			_go_fuzz_dep_.CoverTab[69686]++
													for j := i - 1; j >= 0; j-- {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:509
				_go_fuzz_dep_.CoverTab[69687]++
														if t := s.types[j]; t.in(L, R, AL) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:510
					_go_fuzz_dep_.CoverTab[69688]++
															if t == AL {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:511
						_go_fuzz_dep_.CoverTab[69690]++
																s.types[i] = AN
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:512
						// _ = "end of CoverTab[69690]"
					} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:513
						_go_fuzz_dep_.CoverTab[69691]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:513
						// _ = "end of CoverTab[69691]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:513
					}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:513
					// _ = "end of CoverTab[69688]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:513
					_go_fuzz_dep_.CoverTab[69689]++
															break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:514
					// _ = "end of CoverTab[69689]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:515
					_go_fuzz_dep_.CoverTab[69692]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:515
					// _ = "end of CoverTab[69692]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:515
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:515
				// _ = "end of CoverTab[69687]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:516
			// _ = "end of CoverTab[69686]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:517
			_go_fuzz_dep_.CoverTab[69693]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:517
			// _ = "end of CoverTab[69693]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:517
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:517
		// _ = "end of CoverTab[69685]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:518
	// _ = "end of CoverTab[69676]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:518
	_go_fuzz_dep_.CoverTab[69677]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:521
	for i, t := range s.types {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:521
		_go_fuzz_dep_.CoverTab[69694]++
												if t == AL {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:522
			_go_fuzz_dep_.CoverTab[69695]++
													s.types[i] = R
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:523
			// _ = "end of CoverTab[69695]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:524
			_go_fuzz_dep_.CoverTab[69696]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:524
			// _ = "end of CoverTab[69696]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:524
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:524
		// _ = "end of CoverTab[69694]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:525
	// _ = "end of CoverTab[69677]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:525
	_go_fuzz_dep_.CoverTab[69678]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:540
	for i := 1; i < s.Len()-1; i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:540
		_go_fuzz_dep_.CoverTab[69697]++
												t := s.types[i]
												if t == ES || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:542
			_go_fuzz_dep_.CoverTab[69698]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:542
			return t == CS
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:542
			// _ = "end of CoverTab[69698]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:542
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:542
			_go_fuzz_dep_.CoverTab[69699]++
													prevSepType := s.types[i-1]
													succSepType := s.types[i+1]
													if prevSepType == EN && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:545
				_go_fuzz_dep_.CoverTab[69700]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:545
				return succSepType == EN
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:545
				// _ = "end of CoverTab[69700]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:545
			}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:545
				_go_fuzz_dep_.CoverTab[69701]++
														s.types[i] = EN
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:546
				// _ = "end of CoverTab[69701]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:547
				_go_fuzz_dep_.CoverTab[69702]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:547
				if s.types[i] == CS && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:547
					_go_fuzz_dep_.CoverTab[69703]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:547
					return prevSepType == AN
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:547
					// _ = "end of CoverTab[69703]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:547
				}() && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:547
					_go_fuzz_dep_.CoverTab[69704]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:547
					return succSepType == AN
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:547
					// _ = "end of CoverTab[69704]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:547
				}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:547
					_go_fuzz_dep_.CoverTab[69705]++
															s.types[i] = AN
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:548
					// _ = "end of CoverTab[69705]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:549
					_go_fuzz_dep_.CoverTab[69706]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:549
					// _ = "end of CoverTab[69706]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:549
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:549
				// _ = "end of CoverTab[69702]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:549
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:549
			// _ = "end of CoverTab[69699]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:550
			_go_fuzz_dep_.CoverTab[69707]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:550
			// _ = "end of CoverTab[69707]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:550
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:550
		// _ = "end of CoverTab[69697]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:551
	// _ = "end of CoverTab[69678]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:551
	_go_fuzz_dep_.CoverTab[69679]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:554
	for i, t := range s.types {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:554
		_go_fuzz_dep_.CoverTab[69708]++
												if t == ET {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:555
			_go_fuzz_dep_.CoverTab[69709]++

													runStart := i
													runEnd := s.findRunLimit(runStart, ET)

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:561
			t := s.sos
			if runStart > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:562
				_go_fuzz_dep_.CoverTab[69713]++
														t = s.types[runStart-1]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:563
				// _ = "end of CoverTab[69713]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:564
				_go_fuzz_dep_.CoverTab[69714]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:564
				// _ = "end of CoverTab[69714]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:564
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:564
			// _ = "end of CoverTab[69709]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:564
			_go_fuzz_dep_.CoverTab[69710]++
													if t != EN {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:565
				_go_fuzz_dep_.CoverTab[69715]++
														t = s.eos
														if runEnd < len(s.types) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:567
					_go_fuzz_dep_.CoverTab[69716]++
															t = s.types[runEnd]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:568
					// _ = "end of CoverTab[69716]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:569
					_go_fuzz_dep_.CoverTab[69717]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:569
					// _ = "end of CoverTab[69717]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:569
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:569
				// _ = "end of CoverTab[69715]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:570
				_go_fuzz_dep_.CoverTab[69718]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:570
				// _ = "end of CoverTab[69718]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:570
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:570
			// _ = "end of CoverTab[69710]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:570
			_go_fuzz_dep_.CoverTab[69711]++
													if t == EN {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:571
				_go_fuzz_dep_.CoverTab[69719]++
														setTypes(s.types[runStart:runEnd], EN)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:572
				// _ = "end of CoverTab[69719]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:573
				_go_fuzz_dep_.CoverTab[69720]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:573
				// _ = "end of CoverTab[69720]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:573
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:573
			// _ = "end of CoverTab[69711]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:573
			_go_fuzz_dep_.CoverTab[69712]++

													i = runEnd
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:575
			// _ = "end of CoverTab[69712]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:576
			_go_fuzz_dep_.CoverTab[69721]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:576
			// _ = "end of CoverTab[69721]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:576
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:576
		// _ = "end of CoverTab[69708]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:577
	// _ = "end of CoverTab[69679]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:577
	_go_fuzz_dep_.CoverTab[69680]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:580
	for i, t := range s.types {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:580
		_go_fuzz_dep_.CoverTab[69722]++
												if t.in(ES, ET, CS) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:581
			_go_fuzz_dep_.CoverTab[69723]++
													s.types[i] = ON
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:582
			// _ = "end of CoverTab[69723]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:583
			_go_fuzz_dep_.CoverTab[69724]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:583
			// _ = "end of CoverTab[69724]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:583
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:583
		// _ = "end of CoverTab[69722]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:584
	// _ = "end of CoverTab[69680]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:584
	_go_fuzz_dep_.CoverTab[69681]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:587
	for i, t := range s.types {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:587
		_go_fuzz_dep_.CoverTab[69725]++
												if t == EN {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:588
			_go_fuzz_dep_.CoverTab[69726]++

													prevStrongType := s.sos
													for j := i - 1; j >= 0; j-- {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:591
				_go_fuzz_dep_.CoverTab[69728]++
														t = s.types[j]
														if t == L || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:593
					_go_fuzz_dep_.CoverTab[69729]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:593
					return t == R
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:593
					// _ = "end of CoverTab[69729]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:593
				}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:593
					_go_fuzz_dep_.CoverTab[69730]++
															prevStrongType = t
															break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:595
					// _ = "end of CoverTab[69730]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:596
					_go_fuzz_dep_.CoverTab[69731]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:596
					// _ = "end of CoverTab[69731]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:596
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:596
				// _ = "end of CoverTab[69728]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:597
			// _ = "end of CoverTab[69726]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:597
			_go_fuzz_dep_.CoverTab[69727]++
													if prevStrongType == L {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:598
				_go_fuzz_dep_.CoverTab[69732]++
														s.types[i] = L
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:599
				// _ = "end of CoverTab[69732]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:600
				_go_fuzz_dep_.CoverTab[69733]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:600
				// _ = "end of CoverTab[69733]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:600
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:600
			// _ = "end of CoverTab[69727]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:601
			_go_fuzz_dep_.CoverTab[69734]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:601
			// _ = "end of CoverTab[69734]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:601
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:601
		// _ = "end of CoverTab[69725]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:602
	// _ = "end of CoverTab[69681]"
}

// 6) resolving neutral types Rules N1-N2.
func (s *isolatingRunSequence) resolveNeutralTypes() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:606
	_go_fuzz_dep_.CoverTab[69735]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:609
	s.assertOnly(L, R, EN, AN, B, S, WS, ON, RLI, LRI, FSI, PDI)

	for i, t := range s.types {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:611
		_go_fuzz_dep_.CoverTab[69736]++
												switch t {
		case WS, ON, B, S, RLI, LRI, FSI, PDI:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:613
			_go_fuzz_dep_.CoverTab[69737]++

													runStart := i
													runEnd := s.findRunLimit(runStart, B, S, WS, ON, RLI, LRI, FSI, PDI)

													// determine effective types at ends of run
													var leadType, trailType Class

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:623
			if runStart == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:623
				_go_fuzz_dep_.CoverTab[69742]++
														leadType = s.sos
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:624
				// _ = "end of CoverTab[69742]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:625
				_go_fuzz_dep_.CoverTab[69743]++
														leadType = s.types[runStart-1]
														if leadType.in(AN, EN) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:627
					_go_fuzz_dep_.CoverTab[69744]++
															leadType = R
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:628
					// _ = "end of CoverTab[69744]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:629
					_go_fuzz_dep_.CoverTab[69745]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:629
					// _ = "end of CoverTab[69745]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:629
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:629
				// _ = "end of CoverTab[69743]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:630
			// _ = "end of CoverTab[69737]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:630
			_go_fuzz_dep_.CoverTab[69738]++
													if runEnd == len(s.types) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:631
				_go_fuzz_dep_.CoverTab[69746]++
														trailType = s.eos
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:632
				// _ = "end of CoverTab[69746]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:633
				_go_fuzz_dep_.CoverTab[69747]++
														trailType = s.types[runEnd]
														if trailType.in(AN, EN) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:635
					_go_fuzz_dep_.CoverTab[69748]++
															trailType = R
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:636
					// _ = "end of CoverTab[69748]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:637
					_go_fuzz_dep_.CoverTab[69749]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:637
					// _ = "end of CoverTab[69749]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:637
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:637
				// _ = "end of CoverTab[69747]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:638
			// _ = "end of CoverTab[69738]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:638
			_go_fuzz_dep_.CoverTab[69739]++

													var resolvedType Class
													if leadType == trailType {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:641
				_go_fuzz_dep_.CoverTab[69750]++

														resolvedType = leadType
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:643
				// _ = "end of CoverTab[69750]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:644
				_go_fuzz_dep_.CoverTab[69751]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:648
				resolvedType = typeForLevel(s.level)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:648
				// _ = "end of CoverTab[69751]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:649
			// _ = "end of CoverTab[69739]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:649
			_go_fuzz_dep_.CoverTab[69740]++

													setTypes(s.types[runStart:runEnd], resolvedType)

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:654
			i = runEnd
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:654
			// _ = "end of CoverTab[69740]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:654
		default:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:654
			_go_fuzz_dep_.CoverTab[69741]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:654
			// _ = "end of CoverTab[69741]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:655
		// _ = "end of CoverTab[69736]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:656
	// _ = "end of CoverTab[69735]"
}

func setLevels(levels []level, newLevel level) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:659
	_go_fuzz_dep_.CoverTab[69752]++
											for i := range levels {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:660
		_go_fuzz_dep_.CoverTab[69753]++
												levels[i] = newLevel
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:661
		// _ = "end of CoverTab[69753]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:662
	// _ = "end of CoverTab[69752]"
}

func setTypes(types []Class, newType Class) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:665
	_go_fuzz_dep_.CoverTab[69754]++
											for i := range types {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:666
		_go_fuzz_dep_.CoverTab[69755]++
												types[i] = newType
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:667
		// _ = "end of CoverTab[69755]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:668
	// _ = "end of CoverTab[69754]"
}

// 7) resolving implicit embedding levels Rules I1, I2.
func (s *isolatingRunSequence) resolveImplicitLevels() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:672
	_go_fuzz_dep_.CoverTab[69756]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:675
	s.assertOnly(L, R, EN, AN)

	s.resolvedLevels = make([]level, len(s.types))
	setLevels(s.resolvedLevels, s.level)

	if (s.level & 1) == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:680
		_go_fuzz_dep_.CoverTab[69757]++
												for i, t := range s.types {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:681
			_go_fuzz_dep_.CoverTab[69758]++

													if t == L {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:683
				_go_fuzz_dep_.CoverTab[69759]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:683
				// _ = "end of CoverTab[69759]"

			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:685
				_go_fuzz_dep_.CoverTab[69760]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:685
				if t == R {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:685
					_go_fuzz_dep_.CoverTab[69761]++
															s.resolvedLevels[i] += 1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:686
					// _ = "end of CoverTab[69761]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:687
					_go_fuzz_dep_.CoverTab[69762]++
															s.resolvedLevels[i] += 2
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:688
					// _ = "end of CoverTab[69762]"
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:689
				// _ = "end of CoverTab[69760]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:689
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:689
			// _ = "end of CoverTab[69758]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:690
		// _ = "end of CoverTab[69757]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:691
		_go_fuzz_dep_.CoverTab[69763]++
												for i, t := range s.types {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:692
			_go_fuzz_dep_.CoverTab[69764]++

													if t == R {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:694
				_go_fuzz_dep_.CoverTab[69765]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:694
				// _ = "end of CoverTab[69765]"

			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:696
				_go_fuzz_dep_.CoverTab[69766]++
														s.resolvedLevels[i] += 1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:697
				// _ = "end of CoverTab[69766]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:698
			// _ = "end of CoverTab[69764]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:699
		// _ = "end of CoverTab[69763]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:700
	// _ = "end of CoverTab[69756]"
}

// Applies the levels and types resolved in rules W1-I2 to the
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:703
// resultLevels array.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:705
func (s *isolatingRunSequence) applyLevelsAndTypes() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:705
	_go_fuzz_dep_.CoverTab[69767]++
											for i, x := range s.indexes {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:706
		_go_fuzz_dep_.CoverTab[69768]++
												s.p.resultTypes[x] = s.types[i]
												s.p.resultLevels[x] = s.resolvedLevels[i]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:708
		// _ = "end of CoverTab[69768]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:709
	// _ = "end of CoverTab[69767]"
}

// Return the limit of the run consisting only of the types in validSet
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:712
// starting at index. This checks the value at index, and will return
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:712
// index if that value is not in validSet.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:715
func (s *isolatingRunSequence) findRunLimit(index int, validSet ...Class) int {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:715
	_go_fuzz_dep_.CoverTab[69769]++
loop:
	for ; index < len(s.types); index++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:717
		_go_fuzz_dep_.CoverTab[69771]++
												t := s.types[index]
												for _, valid := range validSet {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:719
			_go_fuzz_dep_.CoverTab[69773]++
													if t == valid {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:720
				_go_fuzz_dep_.CoverTab[69774]++
														continue loop
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:721
				// _ = "end of CoverTab[69774]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:722
				_go_fuzz_dep_.CoverTab[69775]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:722
				// _ = "end of CoverTab[69775]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:722
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:722
			// _ = "end of CoverTab[69773]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:723
		// _ = "end of CoverTab[69771]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:723
		_go_fuzz_dep_.CoverTab[69772]++
												return index
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:724
		// _ = "end of CoverTab[69772]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:725
	// _ = "end of CoverTab[69769]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:725
	_go_fuzz_dep_.CoverTab[69770]++
											return len(s.types)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:726
	// _ = "end of CoverTab[69770]"
}

// Algorithm validation. Assert that all values in types are in the
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:729
// provided set.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:731
func (s *isolatingRunSequence) assertOnly(codes ...Class) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:731
	_go_fuzz_dep_.CoverTab[69776]++
loop:
	for i, t := range s.types {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:733
		_go_fuzz_dep_.CoverTab[69777]++
												for _, c := range codes {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:734
			_go_fuzz_dep_.CoverTab[69779]++
													if t == c {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:735
				_go_fuzz_dep_.CoverTab[69780]++
														continue loop
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:736
				// _ = "end of CoverTab[69780]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:737
				_go_fuzz_dep_.CoverTab[69781]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:737
				// _ = "end of CoverTab[69781]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:737
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:737
			// _ = "end of CoverTab[69779]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:738
		// _ = "end of CoverTab[69777]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:738
		_go_fuzz_dep_.CoverTab[69778]++
												log.Panicf("invalid bidi code %v present in assertOnly at position %d", t, s.indexes[i])
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:739
		// _ = "end of CoverTab[69778]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:740
	// _ = "end of CoverTab[69776]"
}

// determineLevelRuns returns an array of level runs. Each level run is
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:743
// described as an array of indexes into the input string.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:743
//
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:743
// Determines the level runs. Rule X9 will be applied in determining the
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:743
// runs, in the way that makes sure the characters that are supposed to be
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:743
// removed are not included in the runs.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:749
func (p *paragraph) determineLevelRuns() [][]int {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:749
	_go_fuzz_dep_.CoverTab[69782]++
											run := []int{}
											allRuns := [][]int{}
											currentLevel := implicitLevel

											for i := range p.initialTypes {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:754
		_go_fuzz_dep_.CoverTab[69785]++
												if !isRemovedByX9(p.initialTypes[i]) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:755
			_go_fuzz_dep_.CoverTab[69786]++
													if p.resultLevels[i] != currentLevel {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:756
				_go_fuzz_dep_.CoverTab[69788]++

														if currentLevel >= 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:758
					_go_fuzz_dep_.CoverTab[69790]++
															allRuns = append(allRuns, run)
															run = nil
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:760
					// _ = "end of CoverTab[69790]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:761
					_go_fuzz_dep_.CoverTab[69791]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:761
					// _ = "end of CoverTab[69791]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:761
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:761
				// _ = "end of CoverTab[69788]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:761
				_go_fuzz_dep_.CoverTab[69789]++

														currentLevel = p.resultLevels[i]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:763
				// _ = "end of CoverTab[69789]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:764
				_go_fuzz_dep_.CoverTab[69792]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:764
				// _ = "end of CoverTab[69792]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:764
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:764
			// _ = "end of CoverTab[69786]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:764
			_go_fuzz_dep_.CoverTab[69787]++
													run = append(run, i)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:765
			// _ = "end of CoverTab[69787]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:766
			_go_fuzz_dep_.CoverTab[69793]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:766
			// _ = "end of CoverTab[69793]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:766
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:766
		// _ = "end of CoverTab[69785]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:767
	// _ = "end of CoverTab[69782]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:767
	_go_fuzz_dep_.CoverTab[69783]++

											if len(run) > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:769
		_go_fuzz_dep_.CoverTab[69794]++
												allRuns = append(allRuns, run)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:770
		// _ = "end of CoverTab[69794]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:771
		_go_fuzz_dep_.CoverTab[69795]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:771
		// _ = "end of CoverTab[69795]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:771
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:771
	// _ = "end of CoverTab[69783]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:771
	_go_fuzz_dep_.CoverTab[69784]++
											return allRuns
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:772
	// _ = "end of CoverTab[69784]"
}

// Definition BD13. Determine isolating run sequences.
func (p *paragraph) determineIsolatingRunSequences() []*isolatingRunSequence {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:776
	_go_fuzz_dep_.CoverTab[69796]++
											levelRuns := p.determineLevelRuns()

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:780
	runForCharacter := make([]int, p.Len())
	for i, run := range levelRuns {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:781
		_go_fuzz_dep_.CoverTab[69799]++
												for _, index := range run {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:782
			_go_fuzz_dep_.CoverTab[69800]++
													runForCharacter[index] = i
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:783
			// _ = "end of CoverTab[69800]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:784
		// _ = "end of CoverTab[69799]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:785
	// _ = "end of CoverTab[69796]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:785
	_go_fuzz_dep_.CoverTab[69797]++

											sequences := []*isolatingRunSequence{}

											var currentRunSequence []int

											for _, run := range levelRuns {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:791
		_go_fuzz_dep_.CoverTab[69801]++
												first := run[0]
												if p.initialTypes[first] != PDI || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:793
			_go_fuzz_dep_.CoverTab[69802]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:793
			return p.matchingIsolateInitiator[first] == -1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:793
			// _ = "end of CoverTab[69802]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:793
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:793
			_go_fuzz_dep_.CoverTab[69803]++
													currentRunSequence = nil

													for {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:796
				_go_fuzz_dep_.CoverTab[69805]++

														currentRunSequence = append(currentRunSequence, run...)

														last := currentRunSequence[len(currentRunSequence)-1]
														lastT := p.initialTypes[last]
														if lastT.in(LRI, RLI, FSI) && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:802
					_go_fuzz_dep_.CoverTab[69806]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:802
					return p.matchingPDI[last] != p.Len()
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:802
					// _ = "end of CoverTab[69806]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:802
				}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:802
					_go_fuzz_dep_.CoverTab[69807]++
															run = levelRuns[runForCharacter[p.matchingPDI[last]]]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:803
					// _ = "end of CoverTab[69807]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:804
					_go_fuzz_dep_.CoverTab[69808]++
															break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:805
					// _ = "end of CoverTab[69808]"
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:806
				// _ = "end of CoverTab[69805]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:807
			// _ = "end of CoverTab[69803]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:807
			_go_fuzz_dep_.CoverTab[69804]++
													sequences = append(sequences, p.isolatingRunSequence(currentRunSequence))
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:808
			// _ = "end of CoverTab[69804]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:809
			_go_fuzz_dep_.CoverTab[69809]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:809
			// _ = "end of CoverTab[69809]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:809
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:809
		// _ = "end of CoverTab[69801]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:810
	// _ = "end of CoverTab[69797]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:810
	_go_fuzz_dep_.CoverTab[69798]++
											return sequences
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:811
	// _ = "end of CoverTab[69798]"
}

// Assign level information to characters removed by rule X9. This is for
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:814
// ease of relating the level information to the original input data. Note
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:814
// that the levels assigned to these codes are arbitrary, they're chosen so
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:814
// as to avoid breaking level runs.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:818
func (p *paragraph) assignLevelsToCharactersRemovedByX9() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:818
	_go_fuzz_dep_.CoverTab[69810]++
											for i, t := range p.initialTypes {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:819
		_go_fuzz_dep_.CoverTab[69813]++
												if t.in(LRE, RLE, LRO, RLO, PDF, BN) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:820
			_go_fuzz_dep_.CoverTab[69814]++
													p.resultTypes[i] = t
													p.resultLevels[i] = -1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:822
			// _ = "end of CoverTab[69814]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:823
			_go_fuzz_dep_.CoverTab[69815]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:823
			// _ = "end of CoverTab[69815]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:823
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:823
		// _ = "end of CoverTab[69813]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:824
	// _ = "end of CoverTab[69810]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:824
	_go_fuzz_dep_.CoverTab[69811]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:829
	if p.resultLevels[0] == -1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:829
		_go_fuzz_dep_.CoverTab[69816]++
												p.resultLevels[0] = p.embeddingLevel
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:830
		// _ = "end of CoverTab[69816]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:831
		_go_fuzz_dep_.CoverTab[69817]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:831
		// _ = "end of CoverTab[69817]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:831
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:831
	// _ = "end of CoverTab[69811]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:831
	_go_fuzz_dep_.CoverTab[69812]++
											for i := 1; i < len(p.initialTypes); i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:832
		_go_fuzz_dep_.CoverTab[69818]++
												if p.resultLevels[i] == -1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:833
			_go_fuzz_dep_.CoverTab[69819]++
													p.resultLevels[i] = p.resultLevels[i-1]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:834
			// _ = "end of CoverTab[69819]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:835
			_go_fuzz_dep_.CoverTab[69820]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:835
			// _ = "end of CoverTab[69820]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:835
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:835
		// _ = "end of CoverTab[69818]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:836
	// _ = "end of CoverTab[69812]"

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:839
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:845
// getLevels computes levels array breaking lines at offsets in linebreaks.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:845
// Rule L1.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:845
//
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:845
// The linebreaks array must include at least one value. The values must be
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:845
// in strictly increasing order (no duplicates) between 1 and the length of
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:845
// the text, inclusive. The last value must be the length of the text.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:851
func (p *paragraph) getLevels(linebreaks []int) []level {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:851
	_go_fuzz_dep_.CoverTab[69821]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:863
	validateLineBreaks(linebreaks, p.Len())

											result := append([]level(nil), p.resultLevels...)

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:870
	for i, t := range p.initialTypes {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:870
		_go_fuzz_dep_.CoverTab[69824]++
												if t.in(B, S) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:871
			_go_fuzz_dep_.CoverTab[69825]++

													result[i] = p.embeddingLevel

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:876
			for j := i - 1; j >= 0; j-- {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:876
				_go_fuzz_dep_.CoverTab[69826]++
														if isWhitespace(p.initialTypes[j]) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:877
					_go_fuzz_dep_.CoverTab[69827]++
															result[j] = p.embeddingLevel
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:878
					// _ = "end of CoverTab[69827]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:879
					_go_fuzz_dep_.CoverTab[69828]++
															break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:880
					// _ = "end of CoverTab[69828]"
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:881
				// _ = "end of CoverTab[69826]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:882
			// _ = "end of CoverTab[69825]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:883
			_go_fuzz_dep_.CoverTab[69829]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:883
			// _ = "end of CoverTab[69829]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:883
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:883
		// _ = "end of CoverTab[69824]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:884
	// _ = "end of CoverTab[69821]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:884
	_go_fuzz_dep_.CoverTab[69822]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:887
	start := 0
	for _, limit := range linebreaks {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:888
		_go_fuzz_dep_.CoverTab[69830]++
												for j := limit - 1; j >= start; j-- {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:889
			_go_fuzz_dep_.CoverTab[69832]++
													if isWhitespace(p.initialTypes[j]) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:890
				_go_fuzz_dep_.CoverTab[69833]++
														result[j] = p.embeddingLevel
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:891
				// _ = "end of CoverTab[69833]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:892
				_go_fuzz_dep_.CoverTab[69834]++
														break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:893
				// _ = "end of CoverTab[69834]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:894
			// _ = "end of CoverTab[69832]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:895
		// _ = "end of CoverTab[69830]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:895
		_go_fuzz_dep_.CoverTab[69831]++
												start = limit
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:896
		// _ = "end of CoverTab[69831]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:897
	// _ = "end of CoverTab[69822]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:897
	_go_fuzz_dep_.CoverTab[69823]++

											return result
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:899
	// _ = "end of CoverTab[69823]"
}

// getReordering returns the reordering of lines from a visual index to a
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:902
// logical index for line breaks at the given offsets.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:902
//
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:902
// Lines are concatenated from left to right. So for example, the fifth
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:902
// character from the left on the third line is
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:902
//
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:902
//	getReordering(linebreaks)[linebreaks[1] + 4]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:902
//
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:902
// (linebreaks[1] is the position after the last character of the second
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:902
// line, which is also the index of the first character on the third line,
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:902
// and adding four gets the fifth character from the left).
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:902
//
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:902
// The linebreaks array must include at least one value. The values must be
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:902
// in strictly increasing order (no duplicates) between 1 and the length of
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:902
// the text, inclusive. The last value must be the length of the text.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:917
func (p *paragraph) getReordering(linebreaks []int) []int {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:917
	_go_fuzz_dep_.CoverTab[69835]++
											validateLineBreaks(linebreaks, p.Len())

											return computeMultilineReordering(p.getLevels(linebreaks), linebreaks)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:920
	// _ = "end of CoverTab[69835]"
}

// Return multiline reordering array for a given level array. Reordering
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:923
// does not occur across a line break.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:925
func computeMultilineReordering(levels []level, linebreaks []int) []int {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:925
	_go_fuzz_dep_.CoverTab[69836]++
											result := make([]int, len(levels))

											start := 0
											for _, limit := range linebreaks {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:929
		_go_fuzz_dep_.CoverTab[69838]++
												tempLevels := make([]level, limit-start)
												copy(tempLevels, levels[start:])

												for j, order := range computeReordering(tempLevels) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:933
			_go_fuzz_dep_.CoverTab[69840]++
													result[start+j] = order + start
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:934
			// _ = "end of CoverTab[69840]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:935
		// _ = "end of CoverTab[69838]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:935
		_go_fuzz_dep_.CoverTab[69839]++
												start = limit
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:936
		// _ = "end of CoverTab[69839]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:937
	// _ = "end of CoverTab[69836]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:937
	_go_fuzz_dep_.CoverTab[69837]++
											return result
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:938
	// _ = "end of CoverTab[69837]"
}

// Return reordering array for a given level array. This reorders a single
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:941
// line. The reordering is a visual to logical map. For example, the
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:941
// leftmost char is string.charAt(order[0]). Rule L2.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:944
func computeReordering(levels []level) []int {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:944
	_go_fuzz_dep_.CoverTab[69841]++
											result := make([]int, len(levels))

											for i := range result {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:947
		_go_fuzz_dep_.CoverTab[69845]++
												result[i] = i
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:948
		// _ = "end of CoverTab[69845]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:949
	// _ = "end of CoverTab[69841]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:949
	_go_fuzz_dep_.CoverTab[69842]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:954
	highestLevel := level(0)
	lowestOddLevel := level(maxDepth + 2)
	for _, level := range levels {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:956
		_go_fuzz_dep_.CoverTab[69846]++
												if level > highestLevel {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:957
			_go_fuzz_dep_.CoverTab[69848]++
													highestLevel = level
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:958
			// _ = "end of CoverTab[69848]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:959
			_go_fuzz_dep_.CoverTab[69849]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:959
			// _ = "end of CoverTab[69849]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:959
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:959
		// _ = "end of CoverTab[69846]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:959
		_go_fuzz_dep_.CoverTab[69847]++
												if level&1 != 0 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:960
			_go_fuzz_dep_.CoverTab[69850]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:960
			return level < lowestOddLevel
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:960
			// _ = "end of CoverTab[69850]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:960
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:960
			_go_fuzz_dep_.CoverTab[69851]++
													lowestOddLevel = level
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:961
			// _ = "end of CoverTab[69851]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:962
			_go_fuzz_dep_.CoverTab[69852]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:962
			// _ = "end of CoverTab[69852]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:962
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:962
		// _ = "end of CoverTab[69847]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:963
	// _ = "end of CoverTab[69842]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:963
	_go_fuzz_dep_.CoverTab[69843]++

											for level := highestLevel; level >= lowestOddLevel; level-- {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:965
		_go_fuzz_dep_.CoverTab[69853]++
												for i := 0; i < len(levels); i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:966
			_go_fuzz_dep_.CoverTab[69854]++
													if levels[i] >= level {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:967
				_go_fuzz_dep_.CoverTab[69855]++

														start := i
														limit := i + 1
														for limit < len(levels) && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:971
					_go_fuzz_dep_.CoverTab[69858]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:971
					return levels[limit] >= level
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:971
					// _ = "end of CoverTab[69858]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:971
				}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:971
					_go_fuzz_dep_.CoverTab[69859]++
															limit++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:972
					// _ = "end of CoverTab[69859]"
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:973
				// _ = "end of CoverTab[69855]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:973
				_go_fuzz_dep_.CoverTab[69856]++

														for j, k := start, limit-1; j < k; j, k = j+1, k-1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:975
					_go_fuzz_dep_.CoverTab[69860]++
															result[j], result[k] = result[k], result[j]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:976
					// _ = "end of CoverTab[69860]"
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:977
				// _ = "end of CoverTab[69856]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:977
				_go_fuzz_dep_.CoverTab[69857]++

														i = limit
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:979
				// _ = "end of CoverTab[69857]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:980
				_go_fuzz_dep_.CoverTab[69861]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:980
				// _ = "end of CoverTab[69861]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:980
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:980
			// _ = "end of CoverTab[69854]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:981
		// _ = "end of CoverTab[69853]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:982
	// _ = "end of CoverTab[69843]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:982
	_go_fuzz_dep_.CoverTab[69844]++

											return result
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:984
	// _ = "end of CoverTab[69844]"
}

// isWhitespace reports whether the type is considered a whitespace type for the
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:987
// line break rules.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:989
func isWhitespace(c Class) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:989
	_go_fuzz_dep_.CoverTab[69862]++
											switch c {
	case LRE, RLE, LRO, RLO, PDF, LRI, RLI, FSI, PDI, BN, WS:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:991
		_go_fuzz_dep_.CoverTab[69864]++
												return true
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:992
		// _ = "end of CoverTab[69864]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:992
	default:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:992
		_go_fuzz_dep_.CoverTab[69865]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:992
		// _ = "end of CoverTab[69865]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:993
	// _ = "end of CoverTab[69862]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:993
	_go_fuzz_dep_.CoverTab[69863]++
											return false
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:994
	// _ = "end of CoverTab[69863]"
}

// isRemovedByX9 reports whether the type is one of the types removed in X9.
func isRemovedByX9(c Class) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:998
	_go_fuzz_dep_.CoverTab[69866]++
											switch c {
	case LRE, RLE, LRO, RLO, PDF, BN:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1000
		_go_fuzz_dep_.CoverTab[69868]++
												return true
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1001
		// _ = "end of CoverTab[69868]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1001
	default:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1001
		_go_fuzz_dep_.CoverTab[69869]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1001
		// _ = "end of CoverTab[69869]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1002
	// _ = "end of CoverTab[69866]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1002
	_go_fuzz_dep_.CoverTab[69867]++
											return false
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1003
	// _ = "end of CoverTab[69867]"
}

// typeForLevel reports the strong type (L or R) corresponding to the level.
func typeForLevel(level level) Class {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1007
	_go_fuzz_dep_.CoverTab[69870]++
											if (level & 0x1) == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1008
		_go_fuzz_dep_.CoverTab[69872]++
												return L
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1009
		// _ = "end of CoverTab[69872]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1010
		_go_fuzz_dep_.CoverTab[69873]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1010
		// _ = "end of CoverTab[69873]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1010
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1010
	// _ = "end of CoverTab[69870]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1010
	_go_fuzz_dep_.CoverTab[69871]++
											return R
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1011
	// _ = "end of CoverTab[69871]"
}

func validateTypes(types []Class) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1014
	_go_fuzz_dep_.CoverTab[69874]++
											if len(types) == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1015
		_go_fuzz_dep_.CoverTab[69877]++
												return fmt.Errorf("types is null")
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1016
		// _ = "end of CoverTab[69877]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1017
		_go_fuzz_dep_.CoverTab[69878]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1017
		// _ = "end of CoverTab[69878]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1017
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1017
	// _ = "end of CoverTab[69874]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1017
	_go_fuzz_dep_.CoverTab[69875]++
											for i, t := range types[:len(types)-1] {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1018
		_go_fuzz_dep_.CoverTab[69879]++
												if t == B {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1019
			_go_fuzz_dep_.CoverTab[69880]++
													return fmt.Errorf("B type before end of paragraph at index: %d", i)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1020
			// _ = "end of CoverTab[69880]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1021
			_go_fuzz_dep_.CoverTab[69881]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1021
			// _ = "end of CoverTab[69881]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1021
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1021
		// _ = "end of CoverTab[69879]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1022
	// _ = "end of CoverTab[69875]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1022
	_go_fuzz_dep_.CoverTab[69876]++
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1023
	// _ = "end of CoverTab[69876]"
}

func validateParagraphEmbeddingLevel(embeddingLevel level) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1026
	_go_fuzz_dep_.CoverTab[69882]++
											if embeddingLevel != implicitLevel && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1027
		_go_fuzz_dep_.CoverTab[69884]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1027
		return embeddingLevel != 0
												// _ = "end of CoverTab[69884]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1028
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1028
		_go_fuzz_dep_.CoverTab[69885]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1028
		return embeddingLevel != 1
												// _ = "end of CoverTab[69885]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1029
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1029
		_go_fuzz_dep_.CoverTab[69886]++
												return fmt.Errorf("illegal paragraph embedding level: %d", embeddingLevel)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1030
		// _ = "end of CoverTab[69886]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1031
		_go_fuzz_dep_.CoverTab[69887]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1031
		// _ = "end of CoverTab[69887]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1031
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1031
	// _ = "end of CoverTab[69882]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1031
	_go_fuzz_dep_.CoverTab[69883]++
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1032
	// _ = "end of CoverTab[69883]"
}

func validateLineBreaks(linebreaks []int, textLength int) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1035
	_go_fuzz_dep_.CoverTab[69888]++
											prev := 0
											for i, next := range linebreaks {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1037
		_go_fuzz_dep_.CoverTab[69891]++
												if next <= prev {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1038
			_go_fuzz_dep_.CoverTab[69893]++
													return fmt.Errorf("bad linebreak: %d at index: %d", next, i)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1039
			// _ = "end of CoverTab[69893]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1040
			_go_fuzz_dep_.CoverTab[69894]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1040
			// _ = "end of CoverTab[69894]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1040
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1040
		// _ = "end of CoverTab[69891]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1040
		_go_fuzz_dep_.CoverTab[69892]++
												prev = next
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1041
		// _ = "end of CoverTab[69892]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1042
	// _ = "end of CoverTab[69888]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1042
	_go_fuzz_dep_.CoverTab[69889]++
											if prev != textLength {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1043
		_go_fuzz_dep_.CoverTab[69895]++
												return fmt.Errorf("last linebreak was %d, want %d", prev, textLength)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1044
		// _ = "end of CoverTab[69895]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1045
		_go_fuzz_dep_.CoverTab[69896]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1045
		// _ = "end of CoverTab[69896]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1045
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1045
	// _ = "end of CoverTab[69889]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1045
	_go_fuzz_dep_.CoverTab[69890]++
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1046
	// _ = "end of CoverTab[69890]"
}

func validatePbTypes(pairTypes []bracketType) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1049
	_go_fuzz_dep_.CoverTab[69897]++
											if len(pairTypes) == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1050
		_go_fuzz_dep_.CoverTab[69900]++
												return fmt.Errorf("pairTypes is null")
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1051
		// _ = "end of CoverTab[69900]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1052
		_go_fuzz_dep_.CoverTab[69901]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1052
		// _ = "end of CoverTab[69901]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1052
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1052
	// _ = "end of CoverTab[69897]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1052
	_go_fuzz_dep_.CoverTab[69898]++
											for i, pt := range pairTypes {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1053
		_go_fuzz_dep_.CoverTab[69902]++
												switch pt {
		case bpNone, bpOpen, bpClose:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1055
			_go_fuzz_dep_.CoverTab[69903]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1055
			// _ = "end of CoverTab[69903]"
		default:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1056
			_go_fuzz_dep_.CoverTab[69904]++
													return fmt.Errorf("illegal pairType value at %d: %v", i, pairTypes[i])
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1057
			// _ = "end of CoverTab[69904]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1058
		// _ = "end of CoverTab[69902]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1059
	// _ = "end of CoverTab[69898]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1059
	_go_fuzz_dep_.CoverTab[69899]++
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1060
	// _ = "end of CoverTab[69899]"
}

func validatePbValues(pairValues []rune, pairTypes []bracketType) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1063
	_go_fuzz_dep_.CoverTab[69905]++
											if pairValues == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1064
		_go_fuzz_dep_.CoverTab[69908]++
												return fmt.Errorf("pairValues is null")
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1065
		// _ = "end of CoverTab[69908]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1066
		_go_fuzz_dep_.CoverTab[69909]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1066
		// _ = "end of CoverTab[69909]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1066
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1066
	// _ = "end of CoverTab[69905]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1066
	_go_fuzz_dep_.CoverTab[69906]++
											if len(pairTypes) != len(pairValues) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1067
		_go_fuzz_dep_.CoverTab[69910]++
												return fmt.Errorf("pairTypes is different length from pairValues")
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1068
		// _ = "end of CoverTab[69910]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1069
		_go_fuzz_dep_.CoverTab[69911]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1069
		// _ = "end of CoverTab[69911]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1069
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1069
	// _ = "end of CoverTab[69906]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1069
	_go_fuzz_dep_.CoverTab[69907]++
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1070
	// _ = "end of CoverTab[69907]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1071
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/core.go:1071
var _ = _go_fuzz_dep_.CoverTab
