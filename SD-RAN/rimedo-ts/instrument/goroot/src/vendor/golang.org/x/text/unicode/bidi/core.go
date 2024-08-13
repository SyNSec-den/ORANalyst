// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:5
package bidi

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:5
import (
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:5
)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:5
import (
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:5
)

import (
	"fmt"
	"log"
)

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:53
// level is the embedding level of a character. Even embedding levels indicate
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:53
// left-to-right order and odd levels indicate right-to-left order. The special
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:53
// level of -1 is reserved for undefined order.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:56
type level int8

const implicitLevel level = -1

// in returns if x is equal to any of the values in set.
func (c Class) in(set ...Class) bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:61
	_go_fuzz_dep_.CoverTab[32261]++
										for _, s := range set {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:62
		_go_fuzz_dep_.CoverTab[32263]++
											if c == s {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:63
			_go_fuzz_dep_.CoverTab[32264]++
												return true
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:64
			// _ = "end of CoverTab[32264]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:65
			_go_fuzz_dep_.CoverTab[32265]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:65
			// _ = "end of CoverTab[32265]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:65
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:65
		// _ = "end of CoverTab[32263]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:66
	// _ = "end of CoverTab[32261]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:66
	_go_fuzz_dep_.CoverTab[32262]++
										return false
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:67
	// _ = "end of CoverTab[32262]"
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
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:96
// corresponding to the preprocessed text input. The types correspond to the
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:96
// Unicode BiDi classes for each rune. pairTypes indicates the bracket type for
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:96
// each rune. pairValues provides a unique bracket class identifier for each
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:96
// rune (suggested is the rune of the open bracket for opening and matching
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:96
// close brackets, after normalization). The embedding levels are optional, but
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:96
// may be supplied to encode embedding levels of styled text.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:103
func newParagraph(types []Class, pairTypes []bracketType, pairValues []rune, levels level) (*paragraph, error) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:103
	_go_fuzz_dep_.CoverTab[32266]++
										var err error
										if err = validateTypes(types); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:105
		_go_fuzz_dep_.CoverTab[32271]++
											return nil, err
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:106
		// _ = "end of CoverTab[32271]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:107
		_go_fuzz_dep_.CoverTab[32272]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:107
		// _ = "end of CoverTab[32272]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:107
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:107
	// _ = "end of CoverTab[32266]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:107
	_go_fuzz_dep_.CoverTab[32267]++
										if err = validatePbTypes(pairTypes); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:108
		_go_fuzz_dep_.CoverTab[32273]++
											return nil, err
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:109
		// _ = "end of CoverTab[32273]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:110
		_go_fuzz_dep_.CoverTab[32274]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:110
		// _ = "end of CoverTab[32274]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:110
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:110
	// _ = "end of CoverTab[32267]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:110
	_go_fuzz_dep_.CoverTab[32268]++
										if err = validatePbValues(pairValues, pairTypes); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:111
		_go_fuzz_dep_.CoverTab[32275]++
											return nil, err
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:112
		// _ = "end of CoverTab[32275]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:113
		_go_fuzz_dep_.CoverTab[32276]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:113
		// _ = "end of CoverTab[32276]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:113
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:113
	// _ = "end of CoverTab[32268]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:113
	_go_fuzz_dep_.CoverTab[32269]++
										if err = validateParagraphEmbeddingLevel(levels); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:114
		_go_fuzz_dep_.CoverTab[32277]++
											return nil, err
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:115
		// _ = "end of CoverTab[32277]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:116
		_go_fuzz_dep_.CoverTab[32278]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:116
		// _ = "end of CoverTab[32278]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:116
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:116
	// _ = "end of CoverTab[32269]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:116
	_go_fuzz_dep_.CoverTab[32270]++

										p := &paragraph{
		initialTypes:	append([]Class(nil), types...),
		embeddingLevel:	levels,

		pairTypes:	pairTypes,
		pairValues:	pairValues,

		resultTypes:	append([]Class(nil), types...),
	}
										p.run()
										return p, nil
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:128
	// _ = "end of CoverTab[32270]"
}

func (p *paragraph) Len() int {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:131
	_go_fuzz_dep_.CoverTab[32279]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:131
	return len(p.initialTypes)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:131
	// _ = "end of CoverTab[32279]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:131
}

// The algorithm. Does not include line-based processing (Rules L1, L2).
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:133
// These are applied later in the line-based phase of the algorithm.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:135
func (p *paragraph) run() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:135
	_go_fuzz_dep_.CoverTab[32280]++
										p.determineMatchingIsolates()

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:142
	if p.embeddingLevel == implicitLevel {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:142
		_go_fuzz_dep_.CoverTab[32283]++
											p.embeddingLevel = p.determineParagraphEmbeddingLevel(0, p.Len())
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:143
		// _ = "end of CoverTab[32283]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:144
		_go_fuzz_dep_.CoverTab[32284]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:144
		// _ = "end of CoverTab[32284]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:144
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:144
	// _ = "end of CoverTab[32280]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:144
	_go_fuzz_dep_.CoverTab[32281]++

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:147
	p.resultLevels = make([]level, p.Len())
										setLevels(p.resultLevels, p.embeddingLevel)

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:152
	p.determineExplicitEmbeddingLevels()

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:162
	for _, seq := range p.determineIsolatingRunSequences() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:162
		_go_fuzz_dep_.CoverTab[32285]++

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:165
		seq.resolveWeakTypes()

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:169
		resolvePairedBrackets(seq)

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:173
		seq.resolveNeutralTypes()

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:177
		seq.resolveImplicitLevels()

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:180
		seq.applyLevelsAndTypes()
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:180
		// _ = "end of CoverTab[32285]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:181
	// _ = "end of CoverTab[32281]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:181
	_go_fuzz_dep_.CoverTab[32282]++

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:186
	p.assignLevelsToCharactersRemovedByX9()
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:186
	// _ = "end of CoverTab[32282]"
}

// determineMatchingIsolates determines the matching PDI for each isolate
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:189
// initiator and vice versa.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:189
//
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:189
// Definition BD9.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:189
//
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:189
// At the end of this function:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:189
//
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:189
//   - The member variable matchingPDI is set to point to the index of the
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:189
//     matching PDI character for each isolate initiator character. If there is
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:189
//     no matching PDI, it is set to the length of the input text. For other
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:189
//     characters, it is set to -1.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:189
//   - The member variable matchingIsolateInitiator is set to point to the
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:189
//     index of the matching isolate initiator character for each PDI character.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:189
//     If there is no matching isolate initiator, or the character is not a PDI,
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:189
//     it is set to -1.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:204
func (p *paragraph) determineMatchingIsolates() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:204
	_go_fuzz_dep_.CoverTab[32286]++
										p.matchingPDI = make([]int, p.Len())
										p.matchingIsolateInitiator = make([]int, p.Len())

										for i := range p.matchingIsolateInitiator {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:208
		_go_fuzz_dep_.CoverTab[32288]++
											p.matchingIsolateInitiator[i] = -1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:209
		// _ = "end of CoverTab[32288]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:210
	// _ = "end of CoverTab[32286]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:210
	_go_fuzz_dep_.CoverTab[32287]++

										for i := range p.matchingPDI {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:212
		_go_fuzz_dep_.CoverTab[32289]++
											p.matchingPDI[i] = -1

											if t := p.resultTypes[i]; t.in(LRI, RLI, FSI) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:215
			_go_fuzz_dep_.CoverTab[32290]++
												depthCounter := 1
												for j := i + 1; j < p.Len(); j++ {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:217
				_go_fuzz_dep_.CoverTab[32292]++
													if u := p.resultTypes[j]; u.in(LRI, RLI, FSI) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:218
					_go_fuzz_dep_.CoverTab[32293]++
														depthCounter++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:219
					// _ = "end of CoverTab[32293]"
				} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:220
					_go_fuzz_dep_.CoverTab[32294]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:220
					if u == PDI {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:220
						_go_fuzz_dep_.CoverTab[32295]++
															if depthCounter--; depthCounter == 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:221
							_go_fuzz_dep_.CoverTab[32296]++
																p.matchingPDI[i] = j
																p.matchingIsolateInitiator[j] = i
																break
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:224
							// _ = "end of CoverTab[32296]"
						} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:225
							_go_fuzz_dep_.CoverTab[32297]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:225
							// _ = "end of CoverTab[32297]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:225
						}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:225
						// _ = "end of CoverTab[32295]"
					} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:226
						_go_fuzz_dep_.CoverTab[32298]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:226
						// _ = "end of CoverTab[32298]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:226
					}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:226
					// _ = "end of CoverTab[32294]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:226
				}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:226
				// _ = "end of CoverTab[32292]"
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:227
			// _ = "end of CoverTab[32290]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:227
			_go_fuzz_dep_.CoverTab[32291]++
												if p.matchingPDI[i] == -1 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:228
				_go_fuzz_dep_.CoverTab[32299]++
													p.matchingPDI[i] = p.Len()
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:229
				// _ = "end of CoverTab[32299]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:230
				_go_fuzz_dep_.CoverTab[32300]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:230
				// _ = "end of CoverTab[32300]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:230
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:230
			// _ = "end of CoverTab[32291]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:231
			_go_fuzz_dep_.CoverTab[32301]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:231
			// _ = "end of CoverTab[32301]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:231
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:231
		// _ = "end of CoverTab[32289]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:232
	// _ = "end of CoverTab[32287]"
}

// determineParagraphEmbeddingLevel reports the resolved paragraph direction of
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:235
// the substring limited by the given range [start, end).
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:235
//
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:235
// Determines the paragraph level based on rules P2, P3. This is also used
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:235
// in rule X5c to find if an FSI should resolve to LRI or RLI.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:240
func (p *paragraph) determineParagraphEmbeddingLevel(start, end int) level {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:240
	_go_fuzz_dep_.CoverTab[32302]++
										var strongType Class = unknownClass

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:244
	for i := start; i < end; i++ {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:244
		_go_fuzz_dep_.CoverTab[32304]++
											if t := p.resultTypes[i]; t.in(L, AL, R) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:245
			_go_fuzz_dep_.CoverTab[32305]++
												strongType = t
												break
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:247
			// _ = "end of CoverTab[32305]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:248
			_go_fuzz_dep_.CoverTab[32306]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:248
			if t.in(FSI, LRI, RLI) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:248
				_go_fuzz_dep_.CoverTab[32307]++
													i = p.matchingPDI[i]
													if i > end {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:250
					_go_fuzz_dep_.CoverTab[32308]++
														log.Panic("assert (i <= end)")
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:251
					// _ = "end of CoverTab[32308]"
				} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:252
					_go_fuzz_dep_.CoverTab[32309]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:252
					// _ = "end of CoverTab[32309]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:252
				}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:252
				// _ = "end of CoverTab[32307]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:253
				_go_fuzz_dep_.CoverTab[32310]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:253
				// _ = "end of CoverTab[32310]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:253
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:253
			// _ = "end of CoverTab[32306]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:253
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:253
		// _ = "end of CoverTab[32304]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:254
	// _ = "end of CoverTab[32302]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:254
	_go_fuzz_dep_.CoverTab[32303]++

										switch strongType {
	case unknownClass:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:257
		_go_fuzz_dep_.CoverTab[32311]++

											return 0
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:259
		// _ = "end of CoverTab[32311]"
	case L:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:260
		_go_fuzz_dep_.CoverTab[32312]++
											return 0
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:261
		// _ = "end of CoverTab[32312]"
	default:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:262
		_go_fuzz_dep_.CoverTab[32313]++
											return 1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:263
		// _ = "end of CoverTab[32313]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:264
	// _ = "end of CoverTab[32303]"
}

const maxDepth = 125

// This stack will store the embedding levels and override and isolated
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:269
// statuses
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:271
type directionalStatusStack struct {
	stackCounter		int
	embeddingLevelStack	[maxDepth + 1]level
	overrideStatusStack	[maxDepth + 1]Class
	isolateStatusStack	[maxDepth + 1]bool
}

func (s *directionalStatusStack) empty() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:278
	_go_fuzz_dep_.CoverTab[32314]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:278
	s.stackCounter = 0
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:278
	// _ = "end of CoverTab[32314]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:278
}
func (s *directionalStatusStack) pop() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:279
	_go_fuzz_dep_.CoverTab[32315]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:279
	s.stackCounter--
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:279
	// _ = "end of CoverTab[32315]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:279
}
func (s *directionalStatusStack) depth() int {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:280
	_go_fuzz_dep_.CoverTab[32316]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:280
	return s.stackCounter
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:280
	// _ = "end of CoverTab[32316]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:280
}

func (s *directionalStatusStack) push(level level, overrideStatus Class, isolateStatus bool) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:282
	_go_fuzz_dep_.CoverTab[32317]++
										s.embeddingLevelStack[s.stackCounter] = level
										s.overrideStatusStack[s.stackCounter] = overrideStatus
										s.isolateStatusStack[s.stackCounter] = isolateStatus
										s.stackCounter++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:286
	// _ = "end of CoverTab[32317]"
}

func (s *directionalStatusStack) lastEmbeddingLevel() level {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:289
	_go_fuzz_dep_.CoverTab[32318]++
										return s.embeddingLevelStack[s.stackCounter-1]
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:290
	// _ = "end of CoverTab[32318]"
}

func (s *directionalStatusStack) lastDirectionalOverrideStatus() Class {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:293
	_go_fuzz_dep_.CoverTab[32319]++
										return s.overrideStatusStack[s.stackCounter-1]
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:294
	// _ = "end of CoverTab[32319]"
}

func (s *directionalStatusStack) lastDirectionalIsolateStatus() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:297
	_go_fuzz_dep_.CoverTab[32320]++
										return s.isolateStatusStack[s.stackCounter-1]
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:298
	// _ = "end of CoverTab[32320]"
}

// Determine explicit levels using rules X1 - X8
func (p *paragraph) determineExplicitEmbeddingLevels() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:302
	_go_fuzz_dep_.CoverTab[32321]++
										var stack directionalStatusStack
										var overflowIsolateCount, overflowEmbeddingCount, validIsolateCount int

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:307
	stack.push(p.embeddingLevel, ON, false)

	for i, t := range p.resultTypes {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:309
		_go_fuzz_dep_.CoverTab[32322]++

											switch t {
		case RLE, LRE, RLO, LRO, RLI, LRI, FSI:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:312
			_go_fuzz_dep_.CoverTab[32323]++
												isIsolate := t.in(RLI, LRI, FSI)
												isRTL := t.in(RLE, RLO, RLI)

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:317
			if t == FSI {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:317
				_go_fuzz_dep_.CoverTab[32332]++
													isRTL = (p.determineParagraphEmbeddingLevel(i+1, p.matchingPDI[i]) == 1)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:318
				// _ = "end of CoverTab[32332]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:319
				_go_fuzz_dep_.CoverTab[32333]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:319
				// _ = "end of CoverTab[32333]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:319
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:319
			// _ = "end of CoverTab[32323]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:319
			_go_fuzz_dep_.CoverTab[32324]++
												if isIsolate {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:320
				_go_fuzz_dep_.CoverTab[32334]++
													p.resultLevels[i] = stack.lastEmbeddingLevel()
													if stack.lastDirectionalOverrideStatus() != ON {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:322
					_go_fuzz_dep_.CoverTab[32335]++
														p.resultTypes[i] = stack.lastDirectionalOverrideStatus()
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:323
					// _ = "end of CoverTab[32335]"
				} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:324
					_go_fuzz_dep_.CoverTab[32336]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:324
					// _ = "end of CoverTab[32336]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:324
				}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:324
				// _ = "end of CoverTab[32334]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:325
				_go_fuzz_dep_.CoverTab[32337]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:325
				// _ = "end of CoverTab[32337]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:325
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:325
			// _ = "end of CoverTab[32324]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:325
			_go_fuzz_dep_.CoverTab[32325]++

												var newLevel level
												if isRTL {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:328
				_go_fuzz_dep_.CoverTab[32338]++

													newLevel = (stack.lastEmbeddingLevel() + 1) | 1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:330
				// _ = "end of CoverTab[32338]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:331
				_go_fuzz_dep_.CoverTab[32339]++

													newLevel = (stack.lastEmbeddingLevel() + 2) &^ 1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:333
				// _ = "end of CoverTab[32339]"
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:334
			// _ = "end of CoverTab[32325]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:334
			_go_fuzz_dep_.CoverTab[32326]++

												if newLevel <= maxDepth && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:336
				_go_fuzz_dep_.CoverTab[32340]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:336
				return overflowIsolateCount == 0
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:336
				// _ = "end of CoverTab[32340]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:336
			}() && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:336
				_go_fuzz_dep_.CoverTab[32341]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:336
				return overflowEmbeddingCount == 0
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:336
				// _ = "end of CoverTab[32341]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:336
			}() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:336
				_go_fuzz_dep_.CoverTab[32342]++
													if isIsolate {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:337
					_go_fuzz_dep_.CoverTab[32345]++
														validIsolateCount++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:338
					// _ = "end of CoverTab[32345]"
				} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:339
					_go_fuzz_dep_.CoverTab[32346]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:339
					// _ = "end of CoverTab[32346]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:339
				}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:339
				// _ = "end of CoverTab[32342]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:339
				_go_fuzz_dep_.CoverTab[32343]++

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:344
				switch t {
				case LRO:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:345
					_go_fuzz_dep_.CoverTab[32347]++
														stack.push(newLevel, L, isIsolate)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:346
					// _ = "end of CoverTab[32347]"
				case RLO:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:347
					_go_fuzz_dep_.CoverTab[32348]++
														stack.push(newLevel, R, isIsolate)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:348
					// _ = "end of CoverTab[32348]"
				default:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:349
					_go_fuzz_dep_.CoverTab[32349]++
														stack.push(newLevel, ON, isIsolate)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:350
					// _ = "end of CoverTab[32349]"
				}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:351
				// _ = "end of CoverTab[32343]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:351
				_go_fuzz_dep_.CoverTab[32344]++

													if !isIsolate {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:353
					_go_fuzz_dep_.CoverTab[32350]++
														p.resultLevels[i] = newLevel
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:354
					// _ = "end of CoverTab[32350]"
				} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:355
					_go_fuzz_dep_.CoverTab[32351]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:355
					// _ = "end of CoverTab[32351]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:355
				}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:355
				// _ = "end of CoverTab[32344]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:356
				_go_fuzz_dep_.CoverTab[32352]++

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:359
				if isIsolate {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:359
					_go_fuzz_dep_.CoverTab[32353]++
														overflowIsolateCount++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:360
					// _ = "end of CoverTab[32353]"
				} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:361
					_go_fuzz_dep_.CoverTab[32354]++
														if overflowIsolateCount == 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:362
						_go_fuzz_dep_.CoverTab[32355]++
															overflowEmbeddingCount++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:363
						// _ = "end of CoverTab[32355]"
					} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:364
						_go_fuzz_dep_.CoverTab[32356]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:364
						// _ = "end of CoverTab[32356]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:364
					}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:364
					// _ = "end of CoverTab[32354]"
				}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:365
				// _ = "end of CoverTab[32352]"
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:366
			// _ = "end of CoverTab[32326]"

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:369
		case PDI:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:369
			_go_fuzz_dep_.CoverTab[32327]++
												if overflowIsolateCount > 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:370
				_go_fuzz_dep_.CoverTab[32357]++
													overflowIsolateCount--
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:371
				// _ = "end of CoverTab[32357]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:372
				_go_fuzz_dep_.CoverTab[32358]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:372
				if validIsolateCount == 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:372
					_go_fuzz_dep_.CoverTab[32359]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:372
					// _ = "end of CoverTab[32359]"

				} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:374
					_go_fuzz_dep_.CoverTab[32360]++
														overflowEmbeddingCount = 0
														for !stack.lastDirectionalIsolateStatus() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:376
						_go_fuzz_dep_.CoverTab[32362]++
															stack.pop()
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:377
						// _ = "end of CoverTab[32362]"
					}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:378
					// _ = "end of CoverTab[32360]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:378
					_go_fuzz_dep_.CoverTab[32361]++
														stack.pop()
														validIsolateCount--
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:380
					// _ = "end of CoverTab[32361]"
				}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:381
				// _ = "end of CoverTab[32358]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:381
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:381
			// _ = "end of CoverTab[32327]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:381
			_go_fuzz_dep_.CoverTab[32328]++
												p.resultLevels[i] = stack.lastEmbeddingLevel()
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:382
			// _ = "end of CoverTab[32328]"

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:385
		case PDF:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:385
			_go_fuzz_dep_.CoverTab[32329]++

												p.resultLevels[i] = stack.lastEmbeddingLevel()

												if overflowIsolateCount > 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:389
				_go_fuzz_dep_.CoverTab[32363]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:389
				// _ = "end of CoverTab[32363]"

			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:391
				_go_fuzz_dep_.CoverTab[32364]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:391
				if overflowEmbeddingCount > 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:391
					_go_fuzz_dep_.CoverTab[32365]++
														overflowEmbeddingCount--
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:392
					// _ = "end of CoverTab[32365]"
				} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:393
					_go_fuzz_dep_.CoverTab[32366]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:393
					if !stack.lastDirectionalIsolateStatus() && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:393
						_go_fuzz_dep_.CoverTab[32367]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:393
						return stack.depth() >= 2
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:393
						// _ = "end of CoverTab[32367]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:393
					}() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:393
						_go_fuzz_dep_.CoverTab[32368]++
															stack.pop()
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:394
						// _ = "end of CoverTab[32368]"
					} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:395
						_go_fuzz_dep_.CoverTab[32369]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:395
						// _ = "end of CoverTab[32369]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:395
					}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:395
					// _ = "end of CoverTab[32366]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:395
				}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:395
				// _ = "end of CoverTab[32364]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:395
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:395
			// _ = "end of CoverTab[32329]"

		case B:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:397
			_go_fuzz_dep_.CoverTab[32330]++

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:402
			stack.empty()
												overflowIsolateCount = 0
												overflowEmbeddingCount = 0
												validIsolateCount = 0
												p.resultLevels[i] = p.embeddingLevel
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:406
			// _ = "end of CoverTab[32330]"

		default:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:408
			_go_fuzz_dep_.CoverTab[32331]++
												p.resultLevels[i] = stack.lastEmbeddingLevel()
												if stack.lastDirectionalOverrideStatus() != ON {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:410
				_go_fuzz_dep_.CoverTab[32370]++
													p.resultTypes[i] = stack.lastDirectionalOverrideStatus()
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:411
				// _ = "end of CoverTab[32370]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:412
				_go_fuzz_dep_.CoverTab[32371]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:412
				// _ = "end of CoverTab[32371]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:412
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:412
			// _ = "end of CoverTab[32331]"
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:413
		// _ = "end of CoverTab[32322]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:414
	// _ = "end of CoverTab[32321]"
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
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:428
	_go_fuzz_dep_.CoverTab[32372]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:428
	return len(i.indexes)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:428
	// _ = "end of CoverTab[32372]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:428
}

func maxLevel(a, b level) level {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:430
	_go_fuzz_dep_.CoverTab[32373]++
										if a > b {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:431
		_go_fuzz_dep_.CoverTab[32375]++
											return a
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:432
		// _ = "end of CoverTab[32375]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:433
		_go_fuzz_dep_.CoverTab[32376]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:433
		// _ = "end of CoverTab[32376]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:433
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:433
	// _ = "end of CoverTab[32373]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:433
	_go_fuzz_dep_.CoverTab[32374]++
										return b
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:434
	// _ = "end of CoverTab[32374]"
}

// Rule X10, second bullet: Determine the start-of-sequence (sos) and end-of-sequence (eos) types,
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:437
// either L or R, for each isolating run sequence.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:439
func (p *paragraph) isolatingRunSequence(indexes []int) *isolatingRunSequence {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:439
	_go_fuzz_dep_.CoverTab[32377]++
										length := len(indexes)
										types := make([]Class, length)
										for i, x := range indexes {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:442
		_go_fuzz_dep_.CoverTab[32382]++
											types[i] = p.resultTypes[x]
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:443
		// _ = "end of CoverTab[32382]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:444
	// _ = "end of CoverTab[32377]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:444
	_go_fuzz_dep_.CoverTab[32378]++

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:447
	prevChar := indexes[0] - 1
	for prevChar >= 0 && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:448
		_go_fuzz_dep_.CoverTab[32383]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:448
		return isRemovedByX9(p.initialTypes[prevChar])
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:448
		// _ = "end of CoverTab[32383]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:448
	}() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:448
		_go_fuzz_dep_.CoverTab[32384]++
											prevChar--
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:449
		// _ = "end of CoverTab[32384]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:450
	// _ = "end of CoverTab[32378]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:450
	_go_fuzz_dep_.CoverTab[32379]++
										prevLevel := p.embeddingLevel
										if prevChar >= 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:452
		_go_fuzz_dep_.CoverTab[32385]++
											prevLevel = p.resultLevels[prevChar]
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:453
		// _ = "end of CoverTab[32385]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:454
		_go_fuzz_dep_.CoverTab[32386]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:454
		// _ = "end of CoverTab[32386]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:454
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:454
	// _ = "end of CoverTab[32379]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:454
	_go_fuzz_dep_.CoverTab[32380]++

										var succLevel level
										lastType := types[length-1]
										if lastType.in(LRI, RLI, FSI) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:458
		_go_fuzz_dep_.CoverTab[32387]++
											succLevel = p.embeddingLevel
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:459
		// _ = "end of CoverTab[32387]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:460
		_go_fuzz_dep_.CoverTab[32388]++

											limit := indexes[length-1] + 1
											for ; limit < p.Len() && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:463
			_go_fuzz_dep_.CoverTab[32390]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:463
			return isRemovedByX9(p.initialTypes[limit])
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:463
			// _ = "end of CoverTab[32390]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:463
		}(); limit++ {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:463
			_go_fuzz_dep_.CoverTab[32391]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:463
			// _ = "end of CoverTab[32391]"

		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:465
		// _ = "end of CoverTab[32388]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:465
		_go_fuzz_dep_.CoverTab[32389]++
											succLevel = p.embeddingLevel
											if limit < p.Len() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:467
			_go_fuzz_dep_.CoverTab[32392]++
												succLevel = p.resultLevels[limit]
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:468
			// _ = "end of CoverTab[32392]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:469
			_go_fuzz_dep_.CoverTab[32393]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:469
			// _ = "end of CoverTab[32393]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:469
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:469
		// _ = "end of CoverTab[32389]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:470
	// _ = "end of CoverTab[32380]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:470
	_go_fuzz_dep_.CoverTab[32381]++
										level := p.resultLevels[indexes[0]]
										return &isolatingRunSequence{
		p:		p,
		indexes:	indexes,
		types:		types,
		level:		level,
		sos:		typeForLevel(maxLevel(prevLevel, level)),
		eos:		typeForLevel(maxLevel(succLevel, level)),
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:479
	// _ = "end of CoverTab[32381]"
}

// Resolving weak types Rules W1-W7.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:482
//
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:482
// Note that some weak types (EN, AN) remain after this processing is
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:482
// complete.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:486
func (s *isolatingRunSequence) resolveWeakTypes() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:486
	_go_fuzz_dep_.CoverTab[32394]++

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:489
	s.assertOnly(L, R, AL, EN, ES, ET, AN, CS, B, S, WS, ON, NSM, LRI, RLI, FSI, PDI)

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:493
	precedingCharacterType := s.sos
	for i, t := range s.types {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:494
		_go_fuzz_dep_.CoverTab[32401]++
											if t == NSM {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:495
			_go_fuzz_dep_.CoverTab[32402]++
												s.types[i] = precedingCharacterType
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:496
			// _ = "end of CoverTab[32402]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:497
			_go_fuzz_dep_.CoverTab[32403]++

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:501
			precedingCharacterType = t
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:501
			// _ = "end of CoverTab[32403]"
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:502
		// _ = "end of CoverTab[32401]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:503
	// _ = "end of CoverTab[32394]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:503
	_go_fuzz_dep_.CoverTab[32395]++

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:507
	for i, t := range s.types {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:507
		_go_fuzz_dep_.CoverTab[32404]++
											if t == EN {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:508
			_go_fuzz_dep_.CoverTab[32405]++
												for j := i - 1; j >= 0; j-- {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:509
				_go_fuzz_dep_.CoverTab[32406]++
													if t := s.types[j]; t.in(L, R, AL) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:510
					_go_fuzz_dep_.CoverTab[32407]++
														if t == AL {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:511
						_go_fuzz_dep_.CoverTab[32409]++
															s.types[i] = AN
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:512
						// _ = "end of CoverTab[32409]"
					} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:513
						_go_fuzz_dep_.CoverTab[32410]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:513
						// _ = "end of CoverTab[32410]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:513
					}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:513
					// _ = "end of CoverTab[32407]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:513
					_go_fuzz_dep_.CoverTab[32408]++
														break
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:514
					// _ = "end of CoverTab[32408]"
				} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:515
					_go_fuzz_dep_.CoverTab[32411]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:515
					// _ = "end of CoverTab[32411]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:515
				}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:515
				// _ = "end of CoverTab[32406]"
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:516
			// _ = "end of CoverTab[32405]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:517
			_go_fuzz_dep_.CoverTab[32412]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:517
			// _ = "end of CoverTab[32412]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:517
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:517
		// _ = "end of CoverTab[32404]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:518
	// _ = "end of CoverTab[32395]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:518
	_go_fuzz_dep_.CoverTab[32396]++

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:521
	for i, t := range s.types {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:521
		_go_fuzz_dep_.CoverTab[32413]++
											if t == AL {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:522
			_go_fuzz_dep_.CoverTab[32414]++
												s.types[i] = R
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:523
			// _ = "end of CoverTab[32414]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:524
			_go_fuzz_dep_.CoverTab[32415]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:524
			// _ = "end of CoverTab[32415]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:524
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:524
		// _ = "end of CoverTab[32413]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:525
	// _ = "end of CoverTab[32396]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:525
	_go_fuzz_dep_.CoverTab[32397]++

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:540
	for i := 1; i < s.Len()-1; i++ {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:540
		_go_fuzz_dep_.CoverTab[32416]++
											t := s.types[i]
											if t == ES || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:542
			_go_fuzz_dep_.CoverTab[32417]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:542
			return t == CS
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:542
			// _ = "end of CoverTab[32417]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:542
		}() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:542
			_go_fuzz_dep_.CoverTab[32418]++
												prevSepType := s.types[i-1]
												succSepType := s.types[i+1]
												if prevSepType == EN && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:545
				_go_fuzz_dep_.CoverTab[32419]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:545
				return succSepType == EN
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:545
				// _ = "end of CoverTab[32419]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:545
			}() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:545
				_go_fuzz_dep_.CoverTab[32420]++
													s.types[i] = EN
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:546
				// _ = "end of CoverTab[32420]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:547
				_go_fuzz_dep_.CoverTab[32421]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:547
				if s.types[i] == CS && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:547
					_go_fuzz_dep_.CoverTab[32422]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:547
					return prevSepType == AN
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:547
					// _ = "end of CoverTab[32422]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:547
				}() && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:547
					_go_fuzz_dep_.CoverTab[32423]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:547
					return succSepType == AN
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:547
					// _ = "end of CoverTab[32423]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:547
				}() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:547
					_go_fuzz_dep_.CoverTab[32424]++
														s.types[i] = AN
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:548
					// _ = "end of CoverTab[32424]"
				} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:549
					_go_fuzz_dep_.CoverTab[32425]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:549
					// _ = "end of CoverTab[32425]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:549
				}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:549
				// _ = "end of CoverTab[32421]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:549
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:549
			// _ = "end of CoverTab[32418]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:550
			_go_fuzz_dep_.CoverTab[32426]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:550
			// _ = "end of CoverTab[32426]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:550
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:550
		// _ = "end of CoverTab[32416]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:551
	// _ = "end of CoverTab[32397]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:551
	_go_fuzz_dep_.CoverTab[32398]++

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:554
	for i, t := range s.types {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:554
		_go_fuzz_dep_.CoverTab[32427]++
											if t == ET {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:555
			_go_fuzz_dep_.CoverTab[32428]++

												runStart := i
												runEnd := s.findRunLimit(runStart, ET)

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:561
			t := s.sos
			if runStart > 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:562
				_go_fuzz_dep_.CoverTab[32432]++
													t = s.types[runStart-1]
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:563
				// _ = "end of CoverTab[32432]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:564
				_go_fuzz_dep_.CoverTab[32433]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:564
				// _ = "end of CoverTab[32433]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:564
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:564
			// _ = "end of CoverTab[32428]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:564
			_go_fuzz_dep_.CoverTab[32429]++
												if t != EN {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:565
				_go_fuzz_dep_.CoverTab[32434]++
													t = s.eos
													if runEnd < len(s.types) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:567
					_go_fuzz_dep_.CoverTab[32435]++
														t = s.types[runEnd]
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:568
					// _ = "end of CoverTab[32435]"
				} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:569
					_go_fuzz_dep_.CoverTab[32436]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:569
					// _ = "end of CoverTab[32436]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:569
				}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:569
				// _ = "end of CoverTab[32434]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:570
				_go_fuzz_dep_.CoverTab[32437]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:570
				// _ = "end of CoverTab[32437]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:570
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:570
			// _ = "end of CoverTab[32429]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:570
			_go_fuzz_dep_.CoverTab[32430]++
												if t == EN {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:571
				_go_fuzz_dep_.CoverTab[32438]++
													setTypes(s.types[runStart:runEnd], EN)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:572
				// _ = "end of CoverTab[32438]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:573
				_go_fuzz_dep_.CoverTab[32439]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:573
				// _ = "end of CoverTab[32439]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:573
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:573
			// _ = "end of CoverTab[32430]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:573
			_go_fuzz_dep_.CoverTab[32431]++

												i = runEnd
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:575
			// _ = "end of CoverTab[32431]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:576
			_go_fuzz_dep_.CoverTab[32440]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:576
			// _ = "end of CoverTab[32440]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:576
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:576
		// _ = "end of CoverTab[32427]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:577
	// _ = "end of CoverTab[32398]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:577
	_go_fuzz_dep_.CoverTab[32399]++

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:580
	for i, t := range s.types {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:580
		_go_fuzz_dep_.CoverTab[32441]++
											if t.in(ES, ET, CS) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:581
			_go_fuzz_dep_.CoverTab[32442]++
												s.types[i] = ON
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:582
			// _ = "end of CoverTab[32442]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:583
			_go_fuzz_dep_.CoverTab[32443]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:583
			// _ = "end of CoverTab[32443]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:583
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:583
		// _ = "end of CoverTab[32441]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:584
	// _ = "end of CoverTab[32399]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:584
	_go_fuzz_dep_.CoverTab[32400]++

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:587
	for i, t := range s.types {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:587
		_go_fuzz_dep_.CoverTab[32444]++
											if t == EN {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:588
			_go_fuzz_dep_.CoverTab[32445]++

												prevStrongType := s.sos
												for j := i - 1; j >= 0; j-- {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:591
				_go_fuzz_dep_.CoverTab[32447]++
													t = s.types[j]
													if t == L || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:593
					_go_fuzz_dep_.CoverTab[32448]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:593
					return t == R
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:593
					// _ = "end of CoverTab[32448]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:593
				}() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:593
					_go_fuzz_dep_.CoverTab[32449]++
														prevStrongType = t
														break
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:595
					// _ = "end of CoverTab[32449]"
				} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:596
					_go_fuzz_dep_.CoverTab[32450]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:596
					// _ = "end of CoverTab[32450]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:596
				}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:596
				// _ = "end of CoverTab[32447]"
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:597
			// _ = "end of CoverTab[32445]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:597
			_go_fuzz_dep_.CoverTab[32446]++
												if prevStrongType == L {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:598
				_go_fuzz_dep_.CoverTab[32451]++
													s.types[i] = L
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:599
				// _ = "end of CoverTab[32451]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:600
				_go_fuzz_dep_.CoverTab[32452]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:600
				// _ = "end of CoverTab[32452]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:600
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:600
			// _ = "end of CoverTab[32446]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:601
			_go_fuzz_dep_.CoverTab[32453]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:601
			// _ = "end of CoverTab[32453]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:601
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:601
		// _ = "end of CoverTab[32444]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:602
	// _ = "end of CoverTab[32400]"
}

// 6) resolving neutral types Rules N1-N2.
func (s *isolatingRunSequence) resolveNeutralTypes() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:606
	_go_fuzz_dep_.CoverTab[32454]++

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:609
	s.assertOnly(L, R, EN, AN, B, S, WS, ON, RLI, LRI, FSI, PDI)

	for i, t := range s.types {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:611
		_go_fuzz_dep_.CoverTab[32455]++
											switch t {
		case WS, ON, B, S, RLI, LRI, FSI, PDI:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:613
			_go_fuzz_dep_.CoverTab[32456]++

												runStart := i
												runEnd := s.findRunLimit(runStart, B, S, WS, ON, RLI, LRI, FSI, PDI)

												// determine effective types at ends of run
												var leadType, trailType Class

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:623
			if runStart == 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:623
				_go_fuzz_dep_.CoverTab[32461]++
													leadType = s.sos
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:624
				// _ = "end of CoverTab[32461]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:625
				_go_fuzz_dep_.CoverTab[32462]++
													leadType = s.types[runStart-1]
													if leadType.in(AN, EN) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:627
					_go_fuzz_dep_.CoverTab[32463]++
														leadType = R
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:628
					// _ = "end of CoverTab[32463]"
				} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:629
					_go_fuzz_dep_.CoverTab[32464]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:629
					// _ = "end of CoverTab[32464]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:629
				}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:629
				// _ = "end of CoverTab[32462]"
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:630
			// _ = "end of CoverTab[32456]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:630
			_go_fuzz_dep_.CoverTab[32457]++
												if runEnd == len(s.types) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:631
				_go_fuzz_dep_.CoverTab[32465]++
													trailType = s.eos
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:632
				// _ = "end of CoverTab[32465]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:633
				_go_fuzz_dep_.CoverTab[32466]++
													trailType = s.types[runEnd]
													if trailType.in(AN, EN) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:635
					_go_fuzz_dep_.CoverTab[32467]++
														trailType = R
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:636
					// _ = "end of CoverTab[32467]"
				} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:637
					_go_fuzz_dep_.CoverTab[32468]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:637
					// _ = "end of CoverTab[32468]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:637
				}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:637
				// _ = "end of CoverTab[32466]"
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:638
			// _ = "end of CoverTab[32457]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:638
			_go_fuzz_dep_.CoverTab[32458]++

												var resolvedType Class
												if leadType == trailType {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:641
				_go_fuzz_dep_.CoverTab[32469]++

													resolvedType = leadType
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:643
				// _ = "end of CoverTab[32469]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:644
				_go_fuzz_dep_.CoverTab[32470]++

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:648
				resolvedType = typeForLevel(s.level)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:648
				// _ = "end of CoverTab[32470]"
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:649
			// _ = "end of CoverTab[32458]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:649
			_go_fuzz_dep_.CoverTab[32459]++

												setTypes(s.types[runStart:runEnd], resolvedType)

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:654
			i = runEnd
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:654
			// _ = "end of CoverTab[32459]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:654
		default:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:654
			_go_fuzz_dep_.CoverTab[32460]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:654
			// _ = "end of CoverTab[32460]"
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:655
		// _ = "end of CoverTab[32455]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:656
	// _ = "end of CoverTab[32454]"
}

func setLevels(levels []level, newLevel level) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:659
	_go_fuzz_dep_.CoverTab[32471]++
										for i := range levels {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:660
		_go_fuzz_dep_.CoverTab[32472]++
											levels[i] = newLevel
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:661
		// _ = "end of CoverTab[32472]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:662
	// _ = "end of CoverTab[32471]"
}

func setTypes(types []Class, newType Class) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:665
	_go_fuzz_dep_.CoverTab[32473]++
										for i := range types {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:666
		_go_fuzz_dep_.CoverTab[32474]++
											types[i] = newType
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:667
		// _ = "end of CoverTab[32474]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:668
	// _ = "end of CoverTab[32473]"
}

// 7) resolving implicit embedding levels Rules I1, I2.
func (s *isolatingRunSequence) resolveImplicitLevels() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:672
	_go_fuzz_dep_.CoverTab[32475]++

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:675
	s.assertOnly(L, R, EN, AN)

	s.resolvedLevels = make([]level, len(s.types))
	setLevels(s.resolvedLevels, s.level)

	if (s.level & 1) == 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:680
		_go_fuzz_dep_.CoverTab[32476]++
											for i, t := range s.types {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:681
			_go_fuzz_dep_.CoverTab[32477]++

												if t == L {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:683
				_go_fuzz_dep_.CoverTab[32478]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:683
				// _ = "end of CoverTab[32478]"

			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:685
				_go_fuzz_dep_.CoverTab[32479]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:685
				if t == R {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:685
					_go_fuzz_dep_.CoverTab[32480]++
														s.resolvedLevels[i] += 1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:686
					// _ = "end of CoverTab[32480]"
				} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:687
					_go_fuzz_dep_.CoverTab[32481]++
														s.resolvedLevels[i] += 2
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:688
					// _ = "end of CoverTab[32481]"
				}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:689
				// _ = "end of CoverTab[32479]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:689
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:689
			// _ = "end of CoverTab[32477]"
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:690
		// _ = "end of CoverTab[32476]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:691
		_go_fuzz_dep_.CoverTab[32482]++
											for i, t := range s.types {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:692
			_go_fuzz_dep_.CoverTab[32483]++

												if t == R {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:694
				_go_fuzz_dep_.CoverTab[32484]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:694
				// _ = "end of CoverTab[32484]"

			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:696
				_go_fuzz_dep_.CoverTab[32485]++
													s.resolvedLevels[i] += 1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:697
				// _ = "end of CoverTab[32485]"
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:698
			// _ = "end of CoverTab[32483]"
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:699
		// _ = "end of CoverTab[32482]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:700
	// _ = "end of CoverTab[32475]"
}

// Applies the levels and types resolved in rules W1-I2 to the
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:703
// resultLevels array.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:705
func (s *isolatingRunSequence) applyLevelsAndTypes() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:705
	_go_fuzz_dep_.CoverTab[32486]++
										for i, x := range s.indexes {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:706
		_go_fuzz_dep_.CoverTab[32487]++
											s.p.resultTypes[x] = s.types[i]
											s.p.resultLevels[x] = s.resolvedLevels[i]
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:708
		// _ = "end of CoverTab[32487]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:709
	// _ = "end of CoverTab[32486]"
}

// Return the limit of the run consisting only of the types in validSet
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:712
// starting at index. This checks the value at index, and will return
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:712
// index if that value is not in validSet.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:715
func (s *isolatingRunSequence) findRunLimit(index int, validSet ...Class) int {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:715
	_go_fuzz_dep_.CoverTab[32488]++
loop:
	for ; index < len(s.types); index++ {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:717
		_go_fuzz_dep_.CoverTab[32490]++
											t := s.types[index]
											for _, valid := range validSet {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:719
			_go_fuzz_dep_.CoverTab[32492]++
												if t == valid {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:720
				_go_fuzz_dep_.CoverTab[32493]++
													continue loop
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:721
				// _ = "end of CoverTab[32493]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:722
				_go_fuzz_dep_.CoverTab[32494]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:722
				// _ = "end of CoverTab[32494]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:722
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:722
			// _ = "end of CoverTab[32492]"
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:723
		// _ = "end of CoverTab[32490]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:723
		_go_fuzz_dep_.CoverTab[32491]++
											return index
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:724
		// _ = "end of CoverTab[32491]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:725
	// _ = "end of CoverTab[32488]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:725
	_go_fuzz_dep_.CoverTab[32489]++
										return len(s.types)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:726
	// _ = "end of CoverTab[32489]"
}

// Algorithm validation. Assert that all values in types are in the
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:729
// provided set.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:731
func (s *isolatingRunSequence) assertOnly(codes ...Class) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:731
	_go_fuzz_dep_.CoverTab[32495]++
loop:
	for i, t := range s.types {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:733
		_go_fuzz_dep_.CoverTab[32496]++
											for _, c := range codes {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:734
			_go_fuzz_dep_.CoverTab[32498]++
												if t == c {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:735
				_go_fuzz_dep_.CoverTab[32499]++
													continue loop
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:736
				// _ = "end of CoverTab[32499]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:737
				_go_fuzz_dep_.CoverTab[32500]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:737
				// _ = "end of CoverTab[32500]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:737
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:737
			// _ = "end of CoverTab[32498]"
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:738
		// _ = "end of CoverTab[32496]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:738
		_go_fuzz_dep_.CoverTab[32497]++
											log.Panicf("invalid bidi code %v present in assertOnly at position %d", t, s.indexes[i])
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:739
		// _ = "end of CoverTab[32497]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:740
	// _ = "end of CoverTab[32495]"
}

// determineLevelRuns returns an array of level runs. Each level run is
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:743
// described as an array of indexes into the input string.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:743
//
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:743
// Determines the level runs. Rule X9 will be applied in determining the
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:743
// runs, in the way that makes sure the characters that are supposed to be
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:743
// removed are not included in the runs.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:749
func (p *paragraph) determineLevelRuns() [][]int {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:749
	_go_fuzz_dep_.CoverTab[32501]++
										run := []int{}
										allRuns := [][]int{}
										currentLevel := implicitLevel

										for i := range p.initialTypes {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:754
		_go_fuzz_dep_.CoverTab[32504]++
											if !isRemovedByX9(p.initialTypes[i]) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:755
			_go_fuzz_dep_.CoverTab[32505]++
												if p.resultLevels[i] != currentLevel {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:756
				_go_fuzz_dep_.CoverTab[32507]++

													if currentLevel >= 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:758
					_go_fuzz_dep_.CoverTab[32509]++
														allRuns = append(allRuns, run)
														run = nil
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:760
					// _ = "end of CoverTab[32509]"
				} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:761
					_go_fuzz_dep_.CoverTab[32510]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:761
					// _ = "end of CoverTab[32510]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:761
				}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:761
				// _ = "end of CoverTab[32507]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:761
				_go_fuzz_dep_.CoverTab[32508]++

													currentLevel = p.resultLevels[i]
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:763
				// _ = "end of CoverTab[32508]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:764
				_go_fuzz_dep_.CoverTab[32511]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:764
				// _ = "end of CoverTab[32511]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:764
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:764
			// _ = "end of CoverTab[32505]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:764
			_go_fuzz_dep_.CoverTab[32506]++
												run = append(run, i)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:765
			// _ = "end of CoverTab[32506]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:766
			_go_fuzz_dep_.CoverTab[32512]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:766
			// _ = "end of CoverTab[32512]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:766
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:766
		// _ = "end of CoverTab[32504]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:767
	// _ = "end of CoverTab[32501]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:767
	_go_fuzz_dep_.CoverTab[32502]++

										if len(run) > 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:769
		_go_fuzz_dep_.CoverTab[32513]++
											allRuns = append(allRuns, run)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:770
		// _ = "end of CoverTab[32513]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:771
		_go_fuzz_dep_.CoverTab[32514]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:771
		// _ = "end of CoverTab[32514]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:771
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:771
	// _ = "end of CoverTab[32502]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:771
	_go_fuzz_dep_.CoverTab[32503]++
										return allRuns
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:772
	// _ = "end of CoverTab[32503]"
}

// Definition BD13. Determine isolating run sequences.
func (p *paragraph) determineIsolatingRunSequences() []*isolatingRunSequence {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:776
	_go_fuzz_dep_.CoverTab[32515]++
										levelRuns := p.determineLevelRuns()

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:780
	runForCharacter := make([]int, p.Len())
	for i, run := range levelRuns {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:781
		_go_fuzz_dep_.CoverTab[32518]++
											for _, index := range run {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:782
			_go_fuzz_dep_.CoverTab[32519]++
												runForCharacter[index] = i
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:783
			// _ = "end of CoverTab[32519]"
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:784
		// _ = "end of CoverTab[32518]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:785
	// _ = "end of CoverTab[32515]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:785
	_go_fuzz_dep_.CoverTab[32516]++

										sequences := []*isolatingRunSequence{}

										var currentRunSequence []int

										for _, run := range levelRuns {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:791
		_go_fuzz_dep_.CoverTab[32520]++
											first := run[0]
											if p.initialTypes[first] != PDI || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:793
			_go_fuzz_dep_.CoverTab[32521]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:793
			return p.matchingIsolateInitiator[first] == -1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:793
			// _ = "end of CoverTab[32521]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:793
		}() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:793
			_go_fuzz_dep_.CoverTab[32522]++
												currentRunSequence = nil

												for {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:796
				_go_fuzz_dep_.CoverTab[32524]++

													currentRunSequence = append(currentRunSequence, run...)

													last := currentRunSequence[len(currentRunSequence)-1]
													lastT := p.initialTypes[last]
													if lastT.in(LRI, RLI, FSI) && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:802
					_go_fuzz_dep_.CoverTab[32525]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:802
					return p.matchingPDI[last] != p.Len()
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:802
					// _ = "end of CoverTab[32525]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:802
				}() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:802
					_go_fuzz_dep_.CoverTab[32526]++
														run = levelRuns[runForCharacter[p.matchingPDI[last]]]
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:803
					// _ = "end of CoverTab[32526]"
				} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:804
					_go_fuzz_dep_.CoverTab[32527]++
														break
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:805
					// _ = "end of CoverTab[32527]"
				}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:806
				// _ = "end of CoverTab[32524]"
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:807
			// _ = "end of CoverTab[32522]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:807
			_go_fuzz_dep_.CoverTab[32523]++
												sequences = append(sequences, p.isolatingRunSequence(currentRunSequence))
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:808
			// _ = "end of CoverTab[32523]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:809
			_go_fuzz_dep_.CoverTab[32528]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:809
			// _ = "end of CoverTab[32528]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:809
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:809
		// _ = "end of CoverTab[32520]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:810
	// _ = "end of CoverTab[32516]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:810
	_go_fuzz_dep_.CoverTab[32517]++
										return sequences
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:811
	// _ = "end of CoverTab[32517]"
}

// Assign level information to characters removed by rule X9. This is for
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:814
// ease of relating the level information to the original input data. Note
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:814
// that the levels assigned to these codes are arbitrary, they're chosen so
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:814
// as to avoid breaking level runs.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:818
func (p *paragraph) assignLevelsToCharactersRemovedByX9() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:818
	_go_fuzz_dep_.CoverTab[32529]++
										for i, t := range p.initialTypes {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:819
		_go_fuzz_dep_.CoverTab[32532]++
											if t.in(LRE, RLE, LRO, RLO, PDF, BN) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:820
			_go_fuzz_dep_.CoverTab[32533]++
												p.resultTypes[i] = t
												p.resultLevels[i] = -1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:822
			// _ = "end of CoverTab[32533]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:823
			_go_fuzz_dep_.CoverTab[32534]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:823
			// _ = "end of CoverTab[32534]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:823
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:823
		// _ = "end of CoverTab[32532]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:824
	// _ = "end of CoverTab[32529]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:824
	_go_fuzz_dep_.CoverTab[32530]++

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:829
	if p.resultLevels[0] == -1 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:829
		_go_fuzz_dep_.CoverTab[32535]++
											p.resultLevels[0] = p.embeddingLevel
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:830
		// _ = "end of CoverTab[32535]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:831
		_go_fuzz_dep_.CoverTab[32536]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:831
		// _ = "end of CoverTab[32536]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:831
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:831
	// _ = "end of CoverTab[32530]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:831
	_go_fuzz_dep_.CoverTab[32531]++
										for i := 1; i < len(p.initialTypes); i++ {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:832
		_go_fuzz_dep_.CoverTab[32537]++
											if p.resultLevels[i] == -1 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:833
			_go_fuzz_dep_.CoverTab[32538]++
												p.resultLevels[i] = p.resultLevels[i-1]
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:834
			// _ = "end of CoverTab[32538]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:835
			_go_fuzz_dep_.CoverTab[32539]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:835
			// _ = "end of CoverTab[32539]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:835
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:835
		// _ = "end of CoverTab[32537]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:836
	// _ = "end of CoverTab[32531]"

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:839
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:845
// getLevels computes levels array breaking lines at offsets in linebreaks.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:845
// Rule L1.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:845
//
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:845
// The linebreaks array must include at least one value. The values must be
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:845
// in strictly increasing order (no duplicates) between 1 and the length of
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:845
// the text, inclusive. The last value must be the length of the text.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:851
func (p *paragraph) getLevels(linebreaks []int) []level {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:851
	_go_fuzz_dep_.CoverTab[32540]++

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:863
	validateLineBreaks(linebreaks, p.Len())

										result := append([]level(nil), p.resultLevels...)

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:870
	for i, t := range p.initialTypes {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:870
		_go_fuzz_dep_.CoverTab[32543]++
											if t.in(B, S) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:871
			_go_fuzz_dep_.CoverTab[32544]++

												result[i] = p.embeddingLevel

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:876
			for j := i - 1; j >= 0; j-- {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:876
				_go_fuzz_dep_.CoverTab[32545]++
													if isWhitespace(p.initialTypes[j]) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:877
					_go_fuzz_dep_.CoverTab[32546]++
														result[j] = p.embeddingLevel
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:878
					// _ = "end of CoverTab[32546]"
				} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:879
					_go_fuzz_dep_.CoverTab[32547]++
														break
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:880
					// _ = "end of CoverTab[32547]"
				}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:881
				// _ = "end of CoverTab[32545]"
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:882
			// _ = "end of CoverTab[32544]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:883
			_go_fuzz_dep_.CoverTab[32548]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:883
			// _ = "end of CoverTab[32548]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:883
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:883
		// _ = "end of CoverTab[32543]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:884
	// _ = "end of CoverTab[32540]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:884
	_go_fuzz_dep_.CoverTab[32541]++

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:887
	start := 0
	for _, limit := range linebreaks {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:888
		_go_fuzz_dep_.CoverTab[32549]++
											for j := limit - 1; j >= start; j-- {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:889
			_go_fuzz_dep_.CoverTab[32551]++
												if isWhitespace(p.initialTypes[j]) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:890
				_go_fuzz_dep_.CoverTab[32552]++
													result[j] = p.embeddingLevel
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:891
				// _ = "end of CoverTab[32552]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:892
				_go_fuzz_dep_.CoverTab[32553]++
													break
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:893
				// _ = "end of CoverTab[32553]"
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:894
			// _ = "end of CoverTab[32551]"
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:895
		// _ = "end of CoverTab[32549]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:895
		_go_fuzz_dep_.CoverTab[32550]++
											start = limit
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:896
		// _ = "end of CoverTab[32550]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:897
	// _ = "end of CoverTab[32541]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:897
	_go_fuzz_dep_.CoverTab[32542]++

										return result
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:899
	// _ = "end of CoverTab[32542]"
}

// getReordering returns the reordering of lines from a visual index to a
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:902
// logical index for line breaks at the given offsets.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:902
//
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:902
// Lines are concatenated from left to right. So for example, the fifth
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:902
// character from the left on the third line is
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:902
//
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:902
//	getReordering(linebreaks)[linebreaks[1] + 4]
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:902
//
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:902
// (linebreaks[1] is the position after the last character of the second
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:902
// line, which is also the index of the first character on the third line,
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:902
// and adding four gets the fifth character from the left).
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:902
//
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:902
// The linebreaks array must include at least one value. The values must be
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:902
// in strictly increasing order (no duplicates) between 1 and the length of
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:902
// the text, inclusive. The last value must be the length of the text.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:917
func (p *paragraph) getReordering(linebreaks []int) []int {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:917
	_go_fuzz_dep_.CoverTab[32554]++
										validateLineBreaks(linebreaks, p.Len())

										return computeMultilineReordering(p.getLevels(linebreaks), linebreaks)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:920
	// _ = "end of CoverTab[32554]"
}

// Return multiline reordering array for a given level array. Reordering
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:923
// does not occur across a line break.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:925
func computeMultilineReordering(levels []level, linebreaks []int) []int {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:925
	_go_fuzz_dep_.CoverTab[32555]++
										result := make([]int, len(levels))

										start := 0
										for _, limit := range linebreaks {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:929
		_go_fuzz_dep_.CoverTab[32557]++
											tempLevels := make([]level, limit-start)
											copy(tempLevels, levels[start:])

											for j, order := range computeReordering(tempLevels) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:933
			_go_fuzz_dep_.CoverTab[32559]++
												result[start+j] = order + start
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:934
			// _ = "end of CoverTab[32559]"
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:935
		// _ = "end of CoverTab[32557]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:935
		_go_fuzz_dep_.CoverTab[32558]++
											start = limit
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:936
		// _ = "end of CoverTab[32558]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:937
	// _ = "end of CoverTab[32555]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:937
	_go_fuzz_dep_.CoverTab[32556]++
										return result
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:938
	// _ = "end of CoverTab[32556]"
}

// Return reordering array for a given level array. This reorders a single
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:941
// line. The reordering is a visual to logical map. For example, the
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:941
// leftmost char is string.charAt(order[0]). Rule L2.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:944
func computeReordering(levels []level) []int {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:944
	_go_fuzz_dep_.CoverTab[32560]++
										result := make([]int, len(levels))

										for i := range result {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:947
		_go_fuzz_dep_.CoverTab[32564]++
											result[i] = i
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:948
		// _ = "end of CoverTab[32564]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:949
	// _ = "end of CoverTab[32560]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:949
	_go_fuzz_dep_.CoverTab[32561]++

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:954
	highestLevel := level(0)
	lowestOddLevel := level(maxDepth + 2)
	for _, level := range levels {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:956
		_go_fuzz_dep_.CoverTab[32565]++
											if level > highestLevel {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:957
			_go_fuzz_dep_.CoverTab[32567]++
												highestLevel = level
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:958
			// _ = "end of CoverTab[32567]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:959
			_go_fuzz_dep_.CoverTab[32568]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:959
			// _ = "end of CoverTab[32568]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:959
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:959
		// _ = "end of CoverTab[32565]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:959
		_go_fuzz_dep_.CoverTab[32566]++
											if level&1 != 0 && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:960
			_go_fuzz_dep_.CoverTab[32569]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:960
			return level < lowestOddLevel
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:960
			// _ = "end of CoverTab[32569]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:960
		}() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:960
			_go_fuzz_dep_.CoverTab[32570]++
												lowestOddLevel = level
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:961
			// _ = "end of CoverTab[32570]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:962
			_go_fuzz_dep_.CoverTab[32571]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:962
			// _ = "end of CoverTab[32571]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:962
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:962
		// _ = "end of CoverTab[32566]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:963
	// _ = "end of CoverTab[32561]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:963
	_go_fuzz_dep_.CoverTab[32562]++

										for level := highestLevel; level >= lowestOddLevel; level-- {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:965
		_go_fuzz_dep_.CoverTab[32572]++
											for i := 0; i < len(levels); i++ {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:966
			_go_fuzz_dep_.CoverTab[32573]++
												if levels[i] >= level {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:967
				_go_fuzz_dep_.CoverTab[32574]++

													start := i
													limit := i + 1
													for limit < len(levels) && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:971
					_go_fuzz_dep_.CoverTab[32577]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:971
					return levels[limit] >= level
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:971
					// _ = "end of CoverTab[32577]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:971
				}() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:971
					_go_fuzz_dep_.CoverTab[32578]++
														limit++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:972
					// _ = "end of CoverTab[32578]"
				}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:973
				// _ = "end of CoverTab[32574]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:973
				_go_fuzz_dep_.CoverTab[32575]++

													for j, k := start, limit-1; j < k; j, k = j+1, k-1 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:975
					_go_fuzz_dep_.CoverTab[32579]++
														result[j], result[k] = result[k], result[j]
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:976
					// _ = "end of CoverTab[32579]"
				}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:977
				// _ = "end of CoverTab[32575]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:977
				_go_fuzz_dep_.CoverTab[32576]++

													i = limit
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:979
				// _ = "end of CoverTab[32576]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:980
				_go_fuzz_dep_.CoverTab[32580]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:980
				// _ = "end of CoverTab[32580]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:980
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:980
			// _ = "end of CoverTab[32573]"
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:981
		// _ = "end of CoverTab[32572]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:982
	// _ = "end of CoverTab[32562]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:982
	_go_fuzz_dep_.CoverTab[32563]++

										return result
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:984
	// _ = "end of CoverTab[32563]"
}

// isWhitespace reports whether the type is considered a whitespace type for the
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:987
// line break rules.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:989
func isWhitespace(c Class) bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:989
	_go_fuzz_dep_.CoverTab[32581]++
										switch c {
	case LRE, RLE, LRO, RLO, PDF, LRI, RLI, FSI, PDI, BN, WS:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:991
		_go_fuzz_dep_.CoverTab[32583]++
											return true
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:992
		// _ = "end of CoverTab[32583]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:992
	default:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:992
		_go_fuzz_dep_.CoverTab[32584]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:992
		// _ = "end of CoverTab[32584]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:993
	// _ = "end of CoverTab[32581]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:993
	_go_fuzz_dep_.CoverTab[32582]++
										return false
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:994
	// _ = "end of CoverTab[32582]"
}

// isRemovedByX9 reports whether the type is one of the types removed in X9.
func isRemovedByX9(c Class) bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:998
	_go_fuzz_dep_.CoverTab[32585]++
										switch c {
	case LRE, RLE, LRO, RLO, PDF, BN:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1000
		_go_fuzz_dep_.CoverTab[32587]++
											return true
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1001
		// _ = "end of CoverTab[32587]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1001
	default:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1001
		_go_fuzz_dep_.CoverTab[32588]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1001
		// _ = "end of CoverTab[32588]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1002
	// _ = "end of CoverTab[32585]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1002
	_go_fuzz_dep_.CoverTab[32586]++
										return false
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1003
	// _ = "end of CoverTab[32586]"
}

// typeForLevel reports the strong type (L or R) corresponding to the level.
func typeForLevel(level level) Class {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1007
	_go_fuzz_dep_.CoverTab[32589]++
										if (level & 0x1) == 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1008
		_go_fuzz_dep_.CoverTab[32591]++
											return L
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1009
		// _ = "end of CoverTab[32591]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1010
		_go_fuzz_dep_.CoverTab[32592]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1010
		// _ = "end of CoverTab[32592]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1010
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1010
	// _ = "end of CoverTab[32589]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1010
	_go_fuzz_dep_.CoverTab[32590]++
										return R
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1011
	// _ = "end of CoverTab[32590]"
}

func validateTypes(types []Class) error {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1014
	_go_fuzz_dep_.CoverTab[32593]++
										if len(types) == 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1015
		_go_fuzz_dep_.CoverTab[32596]++
											return fmt.Errorf("types is null")
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1016
		// _ = "end of CoverTab[32596]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1017
		_go_fuzz_dep_.CoverTab[32597]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1017
		// _ = "end of CoverTab[32597]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1017
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1017
	// _ = "end of CoverTab[32593]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1017
	_go_fuzz_dep_.CoverTab[32594]++
										for i, t := range types[:len(types)-1] {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1018
		_go_fuzz_dep_.CoverTab[32598]++
											if t == B {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1019
			_go_fuzz_dep_.CoverTab[32599]++
												return fmt.Errorf("B type before end of paragraph at index: %d", i)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1020
			// _ = "end of CoverTab[32599]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1021
			_go_fuzz_dep_.CoverTab[32600]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1021
			// _ = "end of CoverTab[32600]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1021
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1021
		// _ = "end of CoverTab[32598]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1022
	// _ = "end of CoverTab[32594]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1022
	_go_fuzz_dep_.CoverTab[32595]++
										return nil
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1023
	// _ = "end of CoverTab[32595]"
}

func validateParagraphEmbeddingLevel(embeddingLevel level) error {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1026
	_go_fuzz_dep_.CoverTab[32601]++
										if embeddingLevel != implicitLevel && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1027
		_go_fuzz_dep_.CoverTab[32603]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1027
		return embeddingLevel != 0
											// _ = "end of CoverTab[32603]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1028
	}() && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1028
		_go_fuzz_dep_.CoverTab[32604]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1028
		return embeddingLevel != 1
											// _ = "end of CoverTab[32604]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1029
	}() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1029
		_go_fuzz_dep_.CoverTab[32605]++
											return fmt.Errorf("illegal paragraph embedding level: %d", embeddingLevel)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1030
		// _ = "end of CoverTab[32605]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1031
		_go_fuzz_dep_.CoverTab[32606]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1031
		// _ = "end of CoverTab[32606]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1031
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1031
	// _ = "end of CoverTab[32601]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1031
	_go_fuzz_dep_.CoverTab[32602]++
										return nil
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1032
	// _ = "end of CoverTab[32602]"
}

func validateLineBreaks(linebreaks []int, textLength int) error {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1035
	_go_fuzz_dep_.CoverTab[32607]++
										prev := 0
										for i, next := range linebreaks {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1037
		_go_fuzz_dep_.CoverTab[32610]++
											if next <= prev {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1038
			_go_fuzz_dep_.CoverTab[32612]++
												return fmt.Errorf("bad linebreak: %d at index: %d", next, i)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1039
			// _ = "end of CoverTab[32612]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1040
			_go_fuzz_dep_.CoverTab[32613]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1040
			// _ = "end of CoverTab[32613]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1040
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1040
		// _ = "end of CoverTab[32610]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1040
		_go_fuzz_dep_.CoverTab[32611]++
											prev = next
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1041
		// _ = "end of CoverTab[32611]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1042
	// _ = "end of CoverTab[32607]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1042
	_go_fuzz_dep_.CoverTab[32608]++
										if prev != textLength {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1043
		_go_fuzz_dep_.CoverTab[32614]++
											return fmt.Errorf("last linebreak was %d, want %d", prev, textLength)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1044
		// _ = "end of CoverTab[32614]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1045
		_go_fuzz_dep_.CoverTab[32615]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1045
		// _ = "end of CoverTab[32615]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1045
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1045
	// _ = "end of CoverTab[32608]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1045
	_go_fuzz_dep_.CoverTab[32609]++
										return nil
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1046
	// _ = "end of CoverTab[32609]"
}

func validatePbTypes(pairTypes []bracketType) error {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1049
	_go_fuzz_dep_.CoverTab[32616]++
										if len(pairTypes) == 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1050
		_go_fuzz_dep_.CoverTab[32619]++
											return fmt.Errorf("pairTypes is null")
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1051
		// _ = "end of CoverTab[32619]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1052
		_go_fuzz_dep_.CoverTab[32620]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1052
		// _ = "end of CoverTab[32620]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1052
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1052
	// _ = "end of CoverTab[32616]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1052
	_go_fuzz_dep_.CoverTab[32617]++
										for i, pt := range pairTypes {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1053
		_go_fuzz_dep_.CoverTab[32621]++
											switch pt {
		case bpNone, bpOpen, bpClose:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1055
			_go_fuzz_dep_.CoverTab[32622]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1055
			// _ = "end of CoverTab[32622]"
		default:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1056
			_go_fuzz_dep_.CoverTab[32623]++
												return fmt.Errorf("illegal pairType value at %d: %v", i, pairTypes[i])
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1057
			// _ = "end of CoverTab[32623]"
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1058
		// _ = "end of CoverTab[32621]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1059
	// _ = "end of CoverTab[32617]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1059
	_go_fuzz_dep_.CoverTab[32618]++
										return nil
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1060
	// _ = "end of CoverTab[32618]"
}

func validatePbValues(pairValues []rune, pairTypes []bracketType) error {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1063
	_go_fuzz_dep_.CoverTab[32624]++
										if pairValues == nil {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1064
		_go_fuzz_dep_.CoverTab[32627]++
											return fmt.Errorf("pairValues is null")
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1065
		// _ = "end of CoverTab[32627]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1066
		_go_fuzz_dep_.CoverTab[32628]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1066
		// _ = "end of CoverTab[32628]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1066
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1066
	// _ = "end of CoverTab[32624]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1066
	_go_fuzz_dep_.CoverTab[32625]++
										if len(pairTypes) != len(pairValues) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1067
		_go_fuzz_dep_.CoverTab[32629]++
											return fmt.Errorf("pairTypes is different length from pairValues")
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1068
		// _ = "end of CoverTab[32629]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1069
		_go_fuzz_dep_.CoverTab[32630]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1069
		// _ = "end of CoverTab[32630]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1069
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1069
	// _ = "end of CoverTab[32625]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1069
	_go_fuzz_dep_.CoverTab[32626]++
										return nil
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1070
	// _ = "end of CoverTab[32626]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1071
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/core.go:1071
var _ = _go_fuzz_dep_.CoverTab
