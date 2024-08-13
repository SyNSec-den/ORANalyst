// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/regexp/syntax/regexp.go:5
package syntax

//line /usr/local/go/src/regexp/syntax/regexp.go:5
import (
//line /usr/local/go/src/regexp/syntax/regexp.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/regexp/syntax/regexp.go:5
)
//line /usr/local/go/src/regexp/syntax/regexp.go:5
import (
//line /usr/local/go/src/regexp/syntax/regexp.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/regexp/syntax/regexp.go:5
)

//line /usr/local/go/src/regexp/syntax/regexp.go:10
import (
	"strconv"
	"strings"
	"unicode"
)

//line /usr/local/go/src/regexp/syntax/regexp.go:17
type Regexp struct {
	Op		Op
	Flags		Flags
	Sub		[]*Regexp
	Sub0		[1]*Regexp
	Rune		[]rune
	Rune0		[2]rune
	Min, Max	int
	Cap		int
	Name		string
}

//go:generate stringer -type Op -trimprefix Op

//line /usr/local/go/src/regexp/syntax/regexp.go:32
type Op uint8

//line /usr/local/go/src/regexp/syntax/regexp.go:38
const (
	OpNoMatch		Op	= 1 + iota
	OpEmptyMatch
	OpLiteral
	OpCharClass
	OpAnyCharNotNL
	OpAnyChar
	OpBeginLine
	OpEndLine
	OpBeginText
	OpEndText
	OpWordBoundary
	OpNoWordBoundary
	OpCapture
	OpStar
	OpPlus
	OpQuest
	OpRepeat
	OpConcat
	OpAlternate
)

const opPseudo Op = 128

//line /usr/local/go/src/regexp/syntax/regexp.go:63
func (x *Regexp) Equal(y *Regexp) bool {
//line /usr/local/go/src/regexp/syntax/regexp.go:63
	_go_fuzz_dep_.CoverTab[64192]++
							if x == nil || func() bool {
//line /usr/local/go/src/regexp/syntax/regexp.go:64
		_go_fuzz_dep_.CoverTab[64196]++
//line /usr/local/go/src/regexp/syntax/regexp.go:64
		return y == nil
//line /usr/local/go/src/regexp/syntax/regexp.go:64
		// _ = "end of CoverTab[64196]"
//line /usr/local/go/src/regexp/syntax/regexp.go:64
	}() {
//line /usr/local/go/src/regexp/syntax/regexp.go:64
		_go_fuzz_dep_.CoverTab[64197]++
								return x == y
//line /usr/local/go/src/regexp/syntax/regexp.go:65
		// _ = "end of CoverTab[64197]"
	} else {
//line /usr/local/go/src/regexp/syntax/regexp.go:66
		_go_fuzz_dep_.CoverTab[64198]++
//line /usr/local/go/src/regexp/syntax/regexp.go:66
		// _ = "end of CoverTab[64198]"
//line /usr/local/go/src/regexp/syntax/regexp.go:66
	}
//line /usr/local/go/src/regexp/syntax/regexp.go:66
	// _ = "end of CoverTab[64192]"
//line /usr/local/go/src/regexp/syntax/regexp.go:66
	_go_fuzz_dep_.CoverTab[64193]++
							if x.Op != y.Op {
//line /usr/local/go/src/regexp/syntax/regexp.go:67
		_go_fuzz_dep_.CoverTab[64199]++
								return false
//line /usr/local/go/src/regexp/syntax/regexp.go:68
		// _ = "end of CoverTab[64199]"
	} else {
//line /usr/local/go/src/regexp/syntax/regexp.go:69
		_go_fuzz_dep_.CoverTab[64200]++
//line /usr/local/go/src/regexp/syntax/regexp.go:69
		// _ = "end of CoverTab[64200]"
//line /usr/local/go/src/regexp/syntax/regexp.go:69
	}
//line /usr/local/go/src/regexp/syntax/regexp.go:69
	// _ = "end of CoverTab[64193]"
//line /usr/local/go/src/regexp/syntax/regexp.go:69
	_go_fuzz_dep_.CoverTab[64194]++
							switch x.Op {
	case OpEndText:
//line /usr/local/go/src/regexp/syntax/regexp.go:71
		_go_fuzz_dep_.CoverTab[64201]++

								if x.Flags&WasDollar != y.Flags&WasDollar {
//line /usr/local/go/src/regexp/syntax/regexp.go:73
			_go_fuzz_dep_.CoverTab[64210]++
									return false
//line /usr/local/go/src/regexp/syntax/regexp.go:74
			// _ = "end of CoverTab[64210]"
		} else {
//line /usr/local/go/src/regexp/syntax/regexp.go:75
			_go_fuzz_dep_.CoverTab[64211]++
//line /usr/local/go/src/regexp/syntax/regexp.go:75
			// _ = "end of CoverTab[64211]"
//line /usr/local/go/src/regexp/syntax/regexp.go:75
		}
//line /usr/local/go/src/regexp/syntax/regexp.go:75
		// _ = "end of CoverTab[64201]"

	case OpLiteral, OpCharClass:
//line /usr/local/go/src/regexp/syntax/regexp.go:77
		_go_fuzz_dep_.CoverTab[64202]++
								if len(x.Rune) != len(y.Rune) {
//line /usr/local/go/src/regexp/syntax/regexp.go:78
			_go_fuzz_dep_.CoverTab[64212]++
									return false
//line /usr/local/go/src/regexp/syntax/regexp.go:79
			// _ = "end of CoverTab[64212]"
		} else {
//line /usr/local/go/src/regexp/syntax/regexp.go:80
			_go_fuzz_dep_.CoverTab[64213]++
//line /usr/local/go/src/regexp/syntax/regexp.go:80
			// _ = "end of CoverTab[64213]"
//line /usr/local/go/src/regexp/syntax/regexp.go:80
		}
//line /usr/local/go/src/regexp/syntax/regexp.go:80
		// _ = "end of CoverTab[64202]"
//line /usr/local/go/src/regexp/syntax/regexp.go:80
		_go_fuzz_dep_.CoverTab[64203]++
								for i, r := range x.Rune {
//line /usr/local/go/src/regexp/syntax/regexp.go:81
			_go_fuzz_dep_.CoverTab[64214]++
									if r != y.Rune[i] {
//line /usr/local/go/src/regexp/syntax/regexp.go:82
				_go_fuzz_dep_.CoverTab[64215]++
										return false
//line /usr/local/go/src/regexp/syntax/regexp.go:83
				// _ = "end of CoverTab[64215]"
			} else {
//line /usr/local/go/src/regexp/syntax/regexp.go:84
				_go_fuzz_dep_.CoverTab[64216]++
//line /usr/local/go/src/regexp/syntax/regexp.go:84
				// _ = "end of CoverTab[64216]"
//line /usr/local/go/src/regexp/syntax/regexp.go:84
			}
//line /usr/local/go/src/regexp/syntax/regexp.go:84
			// _ = "end of CoverTab[64214]"
		}
//line /usr/local/go/src/regexp/syntax/regexp.go:85
		// _ = "end of CoverTab[64203]"

	case OpAlternate, OpConcat:
//line /usr/local/go/src/regexp/syntax/regexp.go:87
		_go_fuzz_dep_.CoverTab[64204]++
								if len(x.Sub) != len(y.Sub) {
//line /usr/local/go/src/regexp/syntax/regexp.go:88
			_go_fuzz_dep_.CoverTab[64217]++
									return false
//line /usr/local/go/src/regexp/syntax/regexp.go:89
			// _ = "end of CoverTab[64217]"
		} else {
//line /usr/local/go/src/regexp/syntax/regexp.go:90
			_go_fuzz_dep_.CoverTab[64218]++
//line /usr/local/go/src/regexp/syntax/regexp.go:90
			// _ = "end of CoverTab[64218]"
//line /usr/local/go/src/regexp/syntax/regexp.go:90
		}
//line /usr/local/go/src/regexp/syntax/regexp.go:90
		// _ = "end of CoverTab[64204]"
//line /usr/local/go/src/regexp/syntax/regexp.go:90
		_go_fuzz_dep_.CoverTab[64205]++
								for i, sub := range x.Sub {
//line /usr/local/go/src/regexp/syntax/regexp.go:91
			_go_fuzz_dep_.CoverTab[64219]++
									if !sub.Equal(y.Sub[i]) {
//line /usr/local/go/src/regexp/syntax/regexp.go:92
				_go_fuzz_dep_.CoverTab[64220]++
										return false
//line /usr/local/go/src/regexp/syntax/regexp.go:93
				// _ = "end of CoverTab[64220]"
			} else {
//line /usr/local/go/src/regexp/syntax/regexp.go:94
				_go_fuzz_dep_.CoverTab[64221]++
//line /usr/local/go/src/regexp/syntax/regexp.go:94
				// _ = "end of CoverTab[64221]"
//line /usr/local/go/src/regexp/syntax/regexp.go:94
			}
//line /usr/local/go/src/regexp/syntax/regexp.go:94
			// _ = "end of CoverTab[64219]"
		}
//line /usr/local/go/src/regexp/syntax/regexp.go:95
		// _ = "end of CoverTab[64205]"

	case OpStar, OpPlus, OpQuest:
//line /usr/local/go/src/regexp/syntax/regexp.go:97
		_go_fuzz_dep_.CoverTab[64206]++
								if x.Flags&NonGreedy != y.Flags&NonGreedy || func() bool {
//line /usr/local/go/src/regexp/syntax/regexp.go:98
			_go_fuzz_dep_.CoverTab[64222]++
//line /usr/local/go/src/regexp/syntax/regexp.go:98
			return !x.Sub[0].Equal(y.Sub[0])
//line /usr/local/go/src/regexp/syntax/regexp.go:98
			// _ = "end of CoverTab[64222]"
//line /usr/local/go/src/regexp/syntax/regexp.go:98
		}() {
//line /usr/local/go/src/regexp/syntax/regexp.go:98
			_go_fuzz_dep_.CoverTab[64223]++
									return false
//line /usr/local/go/src/regexp/syntax/regexp.go:99
			// _ = "end of CoverTab[64223]"
		} else {
//line /usr/local/go/src/regexp/syntax/regexp.go:100
			_go_fuzz_dep_.CoverTab[64224]++
//line /usr/local/go/src/regexp/syntax/regexp.go:100
			// _ = "end of CoverTab[64224]"
//line /usr/local/go/src/regexp/syntax/regexp.go:100
		}
//line /usr/local/go/src/regexp/syntax/regexp.go:100
		// _ = "end of CoverTab[64206]"

	case OpRepeat:
//line /usr/local/go/src/regexp/syntax/regexp.go:102
		_go_fuzz_dep_.CoverTab[64207]++
								if x.Flags&NonGreedy != y.Flags&NonGreedy || func() bool {
//line /usr/local/go/src/regexp/syntax/regexp.go:103
			_go_fuzz_dep_.CoverTab[64225]++
//line /usr/local/go/src/regexp/syntax/regexp.go:103
			return x.Min != y.Min
//line /usr/local/go/src/regexp/syntax/regexp.go:103
			// _ = "end of CoverTab[64225]"
//line /usr/local/go/src/regexp/syntax/regexp.go:103
		}() || func() bool {
//line /usr/local/go/src/regexp/syntax/regexp.go:103
			_go_fuzz_dep_.CoverTab[64226]++
//line /usr/local/go/src/regexp/syntax/regexp.go:103
			return x.Max != y.Max
//line /usr/local/go/src/regexp/syntax/regexp.go:103
			// _ = "end of CoverTab[64226]"
//line /usr/local/go/src/regexp/syntax/regexp.go:103
		}() || func() bool {
//line /usr/local/go/src/regexp/syntax/regexp.go:103
			_go_fuzz_dep_.CoverTab[64227]++
//line /usr/local/go/src/regexp/syntax/regexp.go:103
			return !x.Sub[0].Equal(y.Sub[0])
//line /usr/local/go/src/regexp/syntax/regexp.go:103
			// _ = "end of CoverTab[64227]"
//line /usr/local/go/src/regexp/syntax/regexp.go:103
		}() {
//line /usr/local/go/src/regexp/syntax/regexp.go:103
			_go_fuzz_dep_.CoverTab[64228]++
									return false
//line /usr/local/go/src/regexp/syntax/regexp.go:104
			// _ = "end of CoverTab[64228]"
		} else {
//line /usr/local/go/src/regexp/syntax/regexp.go:105
			_go_fuzz_dep_.CoverTab[64229]++
//line /usr/local/go/src/regexp/syntax/regexp.go:105
			// _ = "end of CoverTab[64229]"
//line /usr/local/go/src/regexp/syntax/regexp.go:105
		}
//line /usr/local/go/src/regexp/syntax/regexp.go:105
		// _ = "end of CoverTab[64207]"

	case OpCapture:
//line /usr/local/go/src/regexp/syntax/regexp.go:107
		_go_fuzz_dep_.CoverTab[64208]++
								if x.Cap != y.Cap || func() bool {
//line /usr/local/go/src/regexp/syntax/regexp.go:108
			_go_fuzz_dep_.CoverTab[64230]++
//line /usr/local/go/src/regexp/syntax/regexp.go:108
			return x.Name != y.Name
//line /usr/local/go/src/regexp/syntax/regexp.go:108
			// _ = "end of CoverTab[64230]"
//line /usr/local/go/src/regexp/syntax/regexp.go:108
		}() || func() bool {
//line /usr/local/go/src/regexp/syntax/regexp.go:108
			_go_fuzz_dep_.CoverTab[64231]++
//line /usr/local/go/src/regexp/syntax/regexp.go:108
			return !x.Sub[0].Equal(y.Sub[0])
//line /usr/local/go/src/regexp/syntax/regexp.go:108
			// _ = "end of CoverTab[64231]"
//line /usr/local/go/src/regexp/syntax/regexp.go:108
		}() {
//line /usr/local/go/src/regexp/syntax/regexp.go:108
			_go_fuzz_dep_.CoverTab[64232]++
									return false
//line /usr/local/go/src/regexp/syntax/regexp.go:109
			// _ = "end of CoverTab[64232]"
		} else {
//line /usr/local/go/src/regexp/syntax/regexp.go:110
			_go_fuzz_dep_.CoverTab[64233]++
//line /usr/local/go/src/regexp/syntax/regexp.go:110
			// _ = "end of CoverTab[64233]"
//line /usr/local/go/src/regexp/syntax/regexp.go:110
		}
//line /usr/local/go/src/regexp/syntax/regexp.go:110
		// _ = "end of CoverTab[64208]"
//line /usr/local/go/src/regexp/syntax/regexp.go:110
	default:
//line /usr/local/go/src/regexp/syntax/regexp.go:110
		_go_fuzz_dep_.CoverTab[64209]++
//line /usr/local/go/src/regexp/syntax/regexp.go:110
		// _ = "end of CoverTab[64209]"
	}
//line /usr/local/go/src/regexp/syntax/regexp.go:111
	// _ = "end of CoverTab[64194]"
//line /usr/local/go/src/regexp/syntax/regexp.go:111
	_go_fuzz_dep_.CoverTab[64195]++
							return true
//line /usr/local/go/src/regexp/syntax/regexp.go:112
	// _ = "end of CoverTab[64195]"
}

//line /usr/local/go/src/regexp/syntax/regexp.go:116
func writeRegexp(b *strings.Builder, re *Regexp) {
//line /usr/local/go/src/regexp/syntax/regexp.go:116
	_go_fuzz_dep_.CoverTab[64234]++
							switch re.Op {
	default:
//line /usr/local/go/src/regexp/syntax/regexp.go:118
		_go_fuzz_dep_.CoverTab[64235]++
								b.WriteString("<invalid op" + strconv.Itoa(int(re.Op)) + ">")
//line /usr/local/go/src/regexp/syntax/regexp.go:119
		// _ = "end of CoverTab[64235]"
	case OpNoMatch:
//line /usr/local/go/src/regexp/syntax/regexp.go:120
		_go_fuzz_dep_.CoverTab[64236]++
								b.WriteString(`[^\x00-\x{10FFFF}]`)
//line /usr/local/go/src/regexp/syntax/regexp.go:121
		// _ = "end of CoverTab[64236]"
	case OpEmptyMatch:
//line /usr/local/go/src/regexp/syntax/regexp.go:122
		_go_fuzz_dep_.CoverTab[64237]++
								b.WriteString(`(?:)`)
//line /usr/local/go/src/regexp/syntax/regexp.go:123
		// _ = "end of CoverTab[64237]"
	case OpLiteral:
//line /usr/local/go/src/regexp/syntax/regexp.go:124
		_go_fuzz_dep_.CoverTab[64238]++
								if re.Flags&FoldCase != 0 {
//line /usr/local/go/src/regexp/syntax/regexp.go:125
			_go_fuzz_dep_.CoverTab[64260]++
									b.WriteString(`(?i:`)
//line /usr/local/go/src/regexp/syntax/regexp.go:126
			// _ = "end of CoverTab[64260]"
		} else {
//line /usr/local/go/src/regexp/syntax/regexp.go:127
			_go_fuzz_dep_.CoverTab[64261]++
//line /usr/local/go/src/regexp/syntax/regexp.go:127
			// _ = "end of CoverTab[64261]"
//line /usr/local/go/src/regexp/syntax/regexp.go:127
		}
//line /usr/local/go/src/regexp/syntax/regexp.go:127
		// _ = "end of CoverTab[64238]"
//line /usr/local/go/src/regexp/syntax/regexp.go:127
		_go_fuzz_dep_.CoverTab[64239]++
								for _, r := range re.Rune {
//line /usr/local/go/src/regexp/syntax/regexp.go:128
			_go_fuzz_dep_.CoverTab[64262]++
									escape(b, r, false)
//line /usr/local/go/src/regexp/syntax/regexp.go:129
			// _ = "end of CoverTab[64262]"
		}
//line /usr/local/go/src/regexp/syntax/regexp.go:130
		// _ = "end of CoverTab[64239]"
//line /usr/local/go/src/regexp/syntax/regexp.go:130
		_go_fuzz_dep_.CoverTab[64240]++
								if re.Flags&FoldCase != 0 {
//line /usr/local/go/src/regexp/syntax/regexp.go:131
			_go_fuzz_dep_.CoverTab[64263]++
									b.WriteString(`)`)
//line /usr/local/go/src/regexp/syntax/regexp.go:132
			// _ = "end of CoverTab[64263]"
		} else {
//line /usr/local/go/src/regexp/syntax/regexp.go:133
			_go_fuzz_dep_.CoverTab[64264]++
//line /usr/local/go/src/regexp/syntax/regexp.go:133
			// _ = "end of CoverTab[64264]"
//line /usr/local/go/src/regexp/syntax/regexp.go:133
		}
//line /usr/local/go/src/regexp/syntax/regexp.go:133
		// _ = "end of CoverTab[64240]"
	case OpCharClass:
//line /usr/local/go/src/regexp/syntax/regexp.go:134
		_go_fuzz_dep_.CoverTab[64241]++
								if len(re.Rune)%2 != 0 {
//line /usr/local/go/src/regexp/syntax/regexp.go:135
			_go_fuzz_dep_.CoverTab[64265]++
									b.WriteString(`[invalid char class]`)
									break
//line /usr/local/go/src/regexp/syntax/regexp.go:137
			// _ = "end of CoverTab[64265]"
		} else {
//line /usr/local/go/src/regexp/syntax/regexp.go:138
			_go_fuzz_dep_.CoverTab[64266]++
//line /usr/local/go/src/regexp/syntax/regexp.go:138
			// _ = "end of CoverTab[64266]"
//line /usr/local/go/src/regexp/syntax/regexp.go:138
		}
//line /usr/local/go/src/regexp/syntax/regexp.go:138
		// _ = "end of CoverTab[64241]"
//line /usr/local/go/src/regexp/syntax/regexp.go:138
		_go_fuzz_dep_.CoverTab[64242]++
								b.WriteRune('[')
								if len(re.Rune) == 0 {
//line /usr/local/go/src/regexp/syntax/regexp.go:140
			_go_fuzz_dep_.CoverTab[64267]++
									b.WriteString(`^\x00-\x{10FFFF}`)
//line /usr/local/go/src/regexp/syntax/regexp.go:141
			// _ = "end of CoverTab[64267]"
		} else {
//line /usr/local/go/src/regexp/syntax/regexp.go:142
			_go_fuzz_dep_.CoverTab[64268]++
//line /usr/local/go/src/regexp/syntax/regexp.go:142
			if re.Rune[0] == 0 && func() bool {
//line /usr/local/go/src/regexp/syntax/regexp.go:142
				_go_fuzz_dep_.CoverTab[64269]++
//line /usr/local/go/src/regexp/syntax/regexp.go:142
				return re.Rune[len(re.Rune)-1] == unicode.MaxRune
//line /usr/local/go/src/regexp/syntax/regexp.go:142
				// _ = "end of CoverTab[64269]"
//line /usr/local/go/src/regexp/syntax/regexp.go:142
			}() && func() bool {
//line /usr/local/go/src/regexp/syntax/regexp.go:142
				_go_fuzz_dep_.CoverTab[64270]++
//line /usr/local/go/src/regexp/syntax/regexp.go:142
				return len(re.Rune) > 2
//line /usr/local/go/src/regexp/syntax/regexp.go:142
				// _ = "end of CoverTab[64270]"
//line /usr/local/go/src/regexp/syntax/regexp.go:142
			}() {
//line /usr/local/go/src/regexp/syntax/regexp.go:142
				_go_fuzz_dep_.CoverTab[64271]++

//line /usr/local/go/src/regexp/syntax/regexp.go:145
				b.WriteRune('^')
				for i := 1; i < len(re.Rune)-1; i += 2 {
//line /usr/local/go/src/regexp/syntax/regexp.go:146
					_go_fuzz_dep_.CoverTab[64272]++
											lo, hi := re.Rune[i]+1, re.Rune[i+1]-1
											escape(b, lo, lo == '-')
											if lo != hi {
//line /usr/local/go/src/regexp/syntax/regexp.go:149
						_go_fuzz_dep_.CoverTab[64273]++
												b.WriteRune('-')
												escape(b, hi, hi == '-')
//line /usr/local/go/src/regexp/syntax/regexp.go:151
						// _ = "end of CoverTab[64273]"
					} else {
//line /usr/local/go/src/regexp/syntax/regexp.go:152
						_go_fuzz_dep_.CoverTab[64274]++
//line /usr/local/go/src/regexp/syntax/regexp.go:152
						// _ = "end of CoverTab[64274]"
//line /usr/local/go/src/regexp/syntax/regexp.go:152
					}
//line /usr/local/go/src/regexp/syntax/regexp.go:152
					// _ = "end of CoverTab[64272]"
				}
//line /usr/local/go/src/regexp/syntax/regexp.go:153
				// _ = "end of CoverTab[64271]"
			} else {
//line /usr/local/go/src/regexp/syntax/regexp.go:154
				_go_fuzz_dep_.CoverTab[64275]++
										for i := 0; i < len(re.Rune); i += 2 {
//line /usr/local/go/src/regexp/syntax/regexp.go:155
					_go_fuzz_dep_.CoverTab[64276]++
											lo, hi := re.Rune[i], re.Rune[i+1]
											escape(b, lo, lo == '-')
											if lo != hi {
//line /usr/local/go/src/regexp/syntax/regexp.go:158
						_go_fuzz_dep_.CoverTab[64277]++
												b.WriteRune('-')
												escape(b, hi, hi == '-')
//line /usr/local/go/src/regexp/syntax/regexp.go:160
						// _ = "end of CoverTab[64277]"
					} else {
//line /usr/local/go/src/regexp/syntax/regexp.go:161
						_go_fuzz_dep_.CoverTab[64278]++
//line /usr/local/go/src/regexp/syntax/regexp.go:161
						// _ = "end of CoverTab[64278]"
//line /usr/local/go/src/regexp/syntax/regexp.go:161
					}
//line /usr/local/go/src/regexp/syntax/regexp.go:161
					// _ = "end of CoverTab[64276]"
				}
//line /usr/local/go/src/regexp/syntax/regexp.go:162
				// _ = "end of CoverTab[64275]"
			}
//line /usr/local/go/src/regexp/syntax/regexp.go:163
			// _ = "end of CoverTab[64268]"
//line /usr/local/go/src/regexp/syntax/regexp.go:163
		}
//line /usr/local/go/src/regexp/syntax/regexp.go:163
		// _ = "end of CoverTab[64242]"
//line /usr/local/go/src/regexp/syntax/regexp.go:163
		_go_fuzz_dep_.CoverTab[64243]++
								b.WriteRune(']')
//line /usr/local/go/src/regexp/syntax/regexp.go:164
		// _ = "end of CoverTab[64243]"
	case OpAnyCharNotNL:
//line /usr/local/go/src/regexp/syntax/regexp.go:165
		_go_fuzz_dep_.CoverTab[64244]++
								b.WriteString(`(?-s:.)`)
//line /usr/local/go/src/regexp/syntax/regexp.go:166
		// _ = "end of CoverTab[64244]"
	case OpAnyChar:
//line /usr/local/go/src/regexp/syntax/regexp.go:167
		_go_fuzz_dep_.CoverTab[64245]++
								b.WriteString(`(?s:.)`)
//line /usr/local/go/src/regexp/syntax/regexp.go:168
		// _ = "end of CoverTab[64245]"
	case OpBeginLine:
//line /usr/local/go/src/regexp/syntax/regexp.go:169
		_go_fuzz_dep_.CoverTab[64246]++
								b.WriteString(`(?m:^)`)
//line /usr/local/go/src/regexp/syntax/regexp.go:170
		// _ = "end of CoverTab[64246]"
	case OpEndLine:
//line /usr/local/go/src/regexp/syntax/regexp.go:171
		_go_fuzz_dep_.CoverTab[64247]++
								b.WriteString(`(?m:$)`)
//line /usr/local/go/src/regexp/syntax/regexp.go:172
		// _ = "end of CoverTab[64247]"
	case OpBeginText:
//line /usr/local/go/src/regexp/syntax/regexp.go:173
		_go_fuzz_dep_.CoverTab[64248]++
								b.WriteString(`\A`)
//line /usr/local/go/src/regexp/syntax/regexp.go:174
		// _ = "end of CoverTab[64248]"
	case OpEndText:
//line /usr/local/go/src/regexp/syntax/regexp.go:175
		_go_fuzz_dep_.CoverTab[64249]++
								if re.Flags&WasDollar != 0 {
//line /usr/local/go/src/regexp/syntax/regexp.go:176
			_go_fuzz_dep_.CoverTab[64279]++
									b.WriteString(`(?-m:$)`)
//line /usr/local/go/src/regexp/syntax/regexp.go:177
			// _ = "end of CoverTab[64279]"
		} else {
//line /usr/local/go/src/regexp/syntax/regexp.go:178
			_go_fuzz_dep_.CoverTab[64280]++
									b.WriteString(`\z`)
//line /usr/local/go/src/regexp/syntax/regexp.go:179
			// _ = "end of CoverTab[64280]"
		}
//line /usr/local/go/src/regexp/syntax/regexp.go:180
		// _ = "end of CoverTab[64249]"
	case OpWordBoundary:
//line /usr/local/go/src/regexp/syntax/regexp.go:181
		_go_fuzz_dep_.CoverTab[64250]++
								b.WriteString(`\b`)
//line /usr/local/go/src/regexp/syntax/regexp.go:182
		// _ = "end of CoverTab[64250]"
	case OpNoWordBoundary:
//line /usr/local/go/src/regexp/syntax/regexp.go:183
		_go_fuzz_dep_.CoverTab[64251]++
								b.WriteString(`\B`)
//line /usr/local/go/src/regexp/syntax/regexp.go:184
		// _ = "end of CoverTab[64251]"
	case OpCapture:
//line /usr/local/go/src/regexp/syntax/regexp.go:185
		_go_fuzz_dep_.CoverTab[64252]++
								if re.Name != "" {
//line /usr/local/go/src/regexp/syntax/regexp.go:186
			_go_fuzz_dep_.CoverTab[64281]++
									b.WriteString(`(?P<`)
									b.WriteString(re.Name)
									b.WriteRune('>')
//line /usr/local/go/src/regexp/syntax/regexp.go:189
			// _ = "end of CoverTab[64281]"
		} else {
//line /usr/local/go/src/regexp/syntax/regexp.go:190
			_go_fuzz_dep_.CoverTab[64282]++
									b.WriteRune('(')
//line /usr/local/go/src/regexp/syntax/regexp.go:191
			// _ = "end of CoverTab[64282]"
		}
//line /usr/local/go/src/regexp/syntax/regexp.go:192
		// _ = "end of CoverTab[64252]"
//line /usr/local/go/src/regexp/syntax/regexp.go:192
		_go_fuzz_dep_.CoverTab[64253]++
								if re.Sub[0].Op != OpEmptyMatch {
//line /usr/local/go/src/regexp/syntax/regexp.go:193
			_go_fuzz_dep_.CoverTab[64283]++
									writeRegexp(b, re.Sub[0])
//line /usr/local/go/src/regexp/syntax/regexp.go:194
			// _ = "end of CoverTab[64283]"
		} else {
//line /usr/local/go/src/regexp/syntax/regexp.go:195
			_go_fuzz_dep_.CoverTab[64284]++
//line /usr/local/go/src/regexp/syntax/regexp.go:195
			// _ = "end of CoverTab[64284]"
//line /usr/local/go/src/regexp/syntax/regexp.go:195
		}
//line /usr/local/go/src/regexp/syntax/regexp.go:195
		// _ = "end of CoverTab[64253]"
//line /usr/local/go/src/regexp/syntax/regexp.go:195
		_go_fuzz_dep_.CoverTab[64254]++
								b.WriteRune(')')
//line /usr/local/go/src/regexp/syntax/regexp.go:196
		// _ = "end of CoverTab[64254]"
	case OpStar, OpPlus, OpQuest, OpRepeat:
//line /usr/local/go/src/regexp/syntax/regexp.go:197
		_go_fuzz_dep_.CoverTab[64255]++
								if sub := re.Sub[0]; sub.Op > OpCapture || func() bool {
//line /usr/local/go/src/regexp/syntax/regexp.go:198
			_go_fuzz_dep_.CoverTab[64285]++
//line /usr/local/go/src/regexp/syntax/regexp.go:198
			return sub.Op == OpLiteral && func() bool {
//line /usr/local/go/src/regexp/syntax/regexp.go:198
				_go_fuzz_dep_.CoverTab[64286]++
//line /usr/local/go/src/regexp/syntax/regexp.go:198
				return len(sub.Rune) > 1
//line /usr/local/go/src/regexp/syntax/regexp.go:198
				// _ = "end of CoverTab[64286]"
//line /usr/local/go/src/regexp/syntax/regexp.go:198
			}()
//line /usr/local/go/src/regexp/syntax/regexp.go:198
			// _ = "end of CoverTab[64285]"
//line /usr/local/go/src/regexp/syntax/regexp.go:198
		}() {
//line /usr/local/go/src/regexp/syntax/regexp.go:198
			_go_fuzz_dep_.CoverTab[64287]++
									b.WriteString(`(?:`)
									writeRegexp(b, sub)
									b.WriteString(`)`)
//line /usr/local/go/src/regexp/syntax/regexp.go:201
			// _ = "end of CoverTab[64287]"
		} else {
//line /usr/local/go/src/regexp/syntax/regexp.go:202
			_go_fuzz_dep_.CoverTab[64288]++
									writeRegexp(b, sub)
//line /usr/local/go/src/regexp/syntax/regexp.go:203
			// _ = "end of CoverTab[64288]"
		}
//line /usr/local/go/src/regexp/syntax/regexp.go:204
		// _ = "end of CoverTab[64255]"
//line /usr/local/go/src/regexp/syntax/regexp.go:204
		_go_fuzz_dep_.CoverTab[64256]++
								switch re.Op {
		case OpStar:
//line /usr/local/go/src/regexp/syntax/regexp.go:206
			_go_fuzz_dep_.CoverTab[64289]++
									b.WriteRune('*')
//line /usr/local/go/src/regexp/syntax/regexp.go:207
			// _ = "end of CoverTab[64289]"
		case OpPlus:
//line /usr/local/go/src/regexp/syntax/regexp.go:208
			_go_fuzz_dep_.CoverTab[64290]++
									b.WriteRune('+')
//line /usr/local/go/src/regexp/syntax/regexp.go:209
			// _ = "end of CoverTab[64290]"
		case OpQuest:
//line /usr/local/go/src/regexp/syntax/regexp.go:210
			_go_fuzz_dep_.CoverTab[64291]++
									b.WriteRune('?')
//line /usr/local/go/src/regexp/syntax/regexp.go:211
			// _ = "end of CoverTab[64291]"
		case OpRepeat:
//line /usr/local/go/src/regexp/syntax/regexp.go:212
			_go_fuzz_dep_.CoverTab[64292]++
									b.WriteRune('{')
									b.WriteString(strconv.Itoa(re.Min))
									if re.Max != re.Min {
//line /usr/local/go/src/regexp/syntax/regexp.go:215
				_go_fuzz_dep_.CoverTab[64295]++
										b.WriteRune(',')
										if re.Max >= 0 {
//line /usr/local/go/src/regexp/syntax/regexp.go:217
					_go_fuzz_dep_.CoverTab[64296]++
											b.WriteString(strconv.Itoa(re.Max))
//line /usr/local/go/src/regexp/syntax/regexp.go:218
					// _ = "end of CoverTab[64296]"
				} else {
//line /usr/local/go/src/regexp/syntax/regexp.go:219
					_go_fuzz_dep_.CoverTab[64297]++
//line /usr/local/go/src/regexp/syntax/regexp.go:219
					// _ = "end of CoverTab[64297]"
//line /usr/local/go/src/regexp/syntax/regexp.go:219
				}
//line /usr/local/go/src/regexp/syntax/regexp.go:219
				// _ = "end of CoverTab[64295]"
			} else {
//line /usr/local/go/src/regexp/syntax/regexp.go:220
				_go_fuzz_dep_.CoverTab[64298]++
//line /usr/local/go/src/regexp/syntax/regexp.go:220
				// _ = "end of CoverTab[64298]"
//line /usr/local/go/src/regexp/syntax/regexp.go:220
			}
//line /usr/local/go/src/regexp/syntax/regexp.go:220
			// _ = "end of CoverTab[64292]"
//line /usr/local/go/src/regexp/syntax/regexp.go:220
			_go_fuzz_dep_.CoverTab[64293]++
									b.WriteRune('}')
//line /usr/local/go/src/regexp/syntax/regexp.go:221
			// _ = "end of CoverTab[64293]"
//line /usr/local/go/src/regexp/syntax/regexp.go:221
		default:
//line /usr/local/go/src/regexp/syntax/regexp.go:221
			_go_fuzz_dep_.CoverTab[64294]++
//line /usr/local/go/src/regexp/syntax/regexp.go:221
			// _ = "end of CoverTab[64294]"
		}
//line /usr/local/go/src/regexp/syntax/regexp.go:222
		// _ = "end of CoverTab[64256]"
//line /usr/local/go/src/regexp/syntax/regexp.go:222
		_go_fuzz_dep_.CoverTab[64257]++
								if re.Flags&NonGreedy != 0 {
//line /usr/local/go/src/regexp/syntax/regexp.go:223
			_go_fuzz_dep_.CoverTab[64299]++
									b.WriteRune('?')
//line /usr/local/go/src/regexp/syntax/regexp.go:224
			// _ = "end of CoverTab[64299]"
		} else {
//line /usr/local/go/src/regexp/syntax/regexp.go:225
			_go_fuzz_dep_.CoverTab[64300]++
//line /usr/local/go/src/regexp/syntax/regexp.go:225
			// _ = "end of CoverTab[64300]"
//line /usr/local/go/src/regexp/syntax/regexp.go:225
		}
//line /usr/local/go/src/regexp/syntax/regexp.go:225
		// _ = "end of CoverTab[64257]"
	case OpConcat:
//line /usr/local/go/src/regexp/syntax/regexp.go:226
		_go_fuzz_dep_.CoverTab[64258]++
								for _, sub := range re.Sub {
//line /usr/local/go/src/regexp/syntax/regexp.go:227
			_go_fuzz_dep_.CoverTab[64301]++
									if sub.Op == OpAlternate {
//line /usr/local/go/src/regexp/syntax/regexp.go:228
				_go_fuzz_dep_.CoverTab[64302]++
										b.WriteString(`(?:`)
										writeRegexp(b, sub)
										b.WriteString(`)`)
//line /usr/local/go/src/regexp/syntax/regexp.go:231
				// _ = "end of CoverTab[64302]"
			} else {
//line /usr/local/go/src/regexp/syntax/regexp.go:232
				_go_fuzz_dep_.CoverTab[64303]++
										writeRegexp(b, sub)
//line /usr/local/go/src/regexp/syntax/regexp.go:233
				// _ = "end of CoverTab[64303]"
			}
//line /usr/local/go/src/regexp/syntax/regexp.go:234
			// _ = "end of CoverTab[64301]"
		}
//line /usr/local/go/src/regexp/syntax/regexp.go:235
		// _ = "end of CoverTab[64258]"
	case OpAlternate:
//line /usr/local/go/src/regexp/syntax/regexp.go:236
		_go_fuzz_dep_.CoverTab[64259]++
								for i, sub := range re.Sub {
//line /usr/local/go/src/regexp/syntax/regexp.go:237
			_go_fuzz_dep_.CoverTab[64304]++
									if i > 0 {
//line /usr/local/go/src/regexp/syntax/regexp.go:238
				_go_fuzz_dep_.CoverTab[64306]++
										b.WriteRune('|')
//line /usr/local/go/src/regexp/syntax/regexp.go:239
				// _ = "end of CoverTab[64306]"
			} else {
//line /usr/local/go/src/regexp/syntax/regexp.go:240
				_go_fuzz_dep_.CoverTab[64307]++
//line /usr/local/go/src/regexp/syntax/regexp.go:240
				// _ = "end of CoverTab[64307]"
//line /usr/local/go/src/regexp/syntax/regexp.go:240
			}
//line /usr/local/go/src/regexp/syntax/regexp.go:240
			// _ = "end of CoverTab[64304]"
//line /usr/local/go/src/regexp/syntax/regexp.go:240
			_go_fuzz_dep_.CoverTab[64305]++
									writeRegexp(b, sub)
//line /usr/local/go/src/regexp/syntax/regexp.go:241
			// _ = "end of CoverTab[64305]"
		}
//line /usr/local/go/src/regexp/syntax/regexp.go:242
		// _ = "end of CoverTab[64259]"
	}
//line /usr/local/go/src/regexp/syntax/regexp.go:243
	// _ = "end of CoverTab[64234]"
}

func (re *Regexp) String() string {
//line /usr/local/go/src/regexp/syntax/regexp.go:246
	_go_fuzz_dep_.CoverTab[64308]++
							var b strings.Builder
							writeRegexp(&b, re)
							return b.String()
//line /usr/local/go/src/regexp/syntax/regexp.go:249
	// _ = "end of CoverTab[64308]"
}

const meta = `\.+*?()|[]{}^$`

func escape(b *strings.Builder, r rune, force bool) {
//line /usr/local/go/src/regexp/syntax/regexp.go:254
	_go_fuzz_dep_.CoverTab[64309]++
							if unicode.IsPrint(r) {
//line /usr/local/go/src/regexp/syntax/regexp.go:255
		_go_fuzz_dep_.CoverTab[64311]++
								if strings.ContainsRune(meta, r) || func() bool {
//line /usr/local/go/src/regexp/syntax/regexp.go:256
			_go_fuzz_dep_.CoverTab[64313]++
//line /usr/local/go/src/regexp/syntax/regexp.go:256
			return force
//line /usr/local/go/src/regexp/syntax/regexp.go:256
			// _ = "end of CoverTab[64313]"
//line /usr/local/go/src/regexp/syntax/regexp.go:256
		}() {
//line /usr/local/go/src/regexp/syntax/regexp.go:256
			_go_fuzz_dep_.CoverTab[64314]++
									b.WriteRune('\\')
//line /usr/local/go/src/regexp/syntax/regexp.go:257
			// _ = "end of CoverTab[64314]"
		} else {
//line /usr/local/go/src/regexp/syntax/regexp.go:258
			_go_fuzz_dep_.CoverTab[64315]++
//line /usr/local/go/src/regexp/syntax/regexp.go:258
			// _ = "end of CoverTab[64315]"
//line /usr/local/go/src/regexp/syntax/regexp.go:258
		}
//line /usr/local/go/src/regexp/syntax/regexp.go:258
		// _ = "end of CoverTab[64311]"
//line /usr/local/go/src/regexp/syntax/regexp.go:258
		_go_fuzz_dep_.CoverTab[64312]++
								b.WriteRune(r)
								return
//line /usr/local/go/src/regexp/syntax/regexp.go:260
		// _ = "end of CoverTab[64312]"
	} else {
//line /usr/local/go/src/regexp/syntax/regexp.go:261
		_go_fuzz_dep_.CoverTab[64316]++
//line /usr/local/go/src/regexp/syntax/regexp.go:261
		// _ = "end of CoverTab[64316]"
//line /usr/local/go/src/regexp/syntax/regexp.go:261
	}
//line /usr/local/go/src/regexp/syntax/regexp.go:261
	// _ = "end of CoverTab[64309]"
//line /usr/local/go/src/regexp/syntax/regexp.go:261
	_go_fuzz_dep_.CoverTab[64310]++

							switch r {
	case '\a':
//line /usr/local/go/src/regexp/syntax/regexp.go:264
		_go_fuzz_dep_.CoverTab[64317]++
								b.WriteString(`\a`)
//line /usr/local/go/src/regexp/syntax/regexp.go:265
		// _ = "end of CoverTab[64317]"
	case '\f':
//line /usr/local/go/src/regexp/syntax/regexp.go:266
		_go_fuzz_dep_.CoverTab[64318]++
								b.WriteString(`\f`)
//line /usr/local/go/src/regexp/syntax/regexp.go:267
		// _ = "end of CoverTab[64318]"
	case '\n':
//line /usr/local/go/src/regexp/syntax/regexp.go:268
		_go_fuzz_dep_.CoverTab[64319]++
								b.WriteString(`\n`)
//line /usr/local/go/src/regexp/syntax/regexp.go:269
		// _ = "end of CoverTab[64319]"
	case '\r':
//line /usr/local/go/src/regexp/syntax/regexp.go:270
		_go_fuzz_dep_.CoverTab[64320]++
								b.WriteString(`\r`)
//line /usr/local/go/src/regexp/syntax/regexp.go:271
		// _ = "end of CoverTab[64320]"
	case '\t':
//line /usr/local/go/src/regexp/syntax/regexp.go:272
		_go_fuzz_dep_.CoverTab[64321]++
								b.WriteString(`\t`)
//line /usr/local/go/src/regexp/syntax/regexp.go:273
		// _ = "end of CoverTab[64321]"
	case '\v':
//line /usr/local/go/src/regexp/syntax/regexp.go:274
		_go_fuzz_dep_.CoverTab[64322]++
								b.WriteString(`\v`)
//line /usr/local/go/src/regexp/syntax/regexp.go:275
		// _ = "end of CoverTab[64322]"
	default:
//line /usr/local/go/src/regexp/syntax/regexp.go:276
		_go_fuzz_dep_.CoverTab[64323]++
								if r < 0x100 {
//line /usr/local/go/src/regexp/syntax/regexp.go:277
			_go_fuzz_dep_.CoverTab[64325]++
									b.WriteString(`\x`)
									s := strconv.FormatInt(int64(r), 16)
									if len(s) == 1 {
//line /usr/local/go/src/regexp/syntax/regexp.go:280
				_go_fuzz_dep_.CoverTab[64327]++
										b.WriteRune('0')
//line /usr/local/go/src/regexp/syntax/regexp.go:281
				// _ = "end of CoverTab[64327]"
			} else {
//line /usr/local/go/src/regexp/syntax/regexp.go:282
				_go_fuzz_dep_.CoverTab[64328]++
//line /usr/local/go/src/regexp/syntax/regexp.go:282
				// _ = "end of CoverTab[64328]"
//line /usr/local/go/src/regexp/syntax/regexp.go:282
			}
//line /usr/local/go/src/regexp/syntax/regexp.go:282
			// _ = "end of CoverTab[64325]"
//line /usr/local/go/src/regexp/syntax/regexp.go:282
			_go_fuzz_dep_.CoverTab[64326]++
									b.WriteString(s)
									break
//line /usr/local/go/src/regexp/syntax/regexp.go:284
			// _ = "end of CoverTab[64326]"
		} else {
//line /usr/local/go/src/regexp/syntax/regexp.go:285
			_go_fuzz_dep_.CoverTab[64329]++
//line /usr/local/go/src/regexp/syntax/regexp.go:285
			// _ = "end of CoverTab[64329]"
//line /usr/local/go/src/regexp/syntax/regexp.go:285
		}
//line /usr/local/go/src/regexp/syntax/regexp.go:285
		// _ = "end of CoverTab[64323]"
//line /usr/local/go/src/regexp/syntax/regexp.go:285
		_go_fuzz_dep_.CoverTab[64324]++
								b.WriteString(`\x{`)
								b.WriteString(strconv.FormatInt(int64(r), 16))
								b.WriteString(`}`)
//line /usr/local/go/src/regexp/syntax/regexp.go:288
		// _ = "end of CoverTab[64324]"
	}
//line /usr/local/go/src/regexp/syntax/regexp.go:289
	// _ = "end of CoverTab[64310]"
}

//line /usr/local/go/src/regexp/syntax/regexp.go:293
func (re *Regexp) MaxCap() int {
//line /usr/local/go/src/regexp/syntax/regexp.go:293
	_go_fuzz_dep_.CoverTab[64330]++
							m := 0
							if re.Op == OpCapture {
//line /usr/local/go/src/regexp/syntax/regexp.go:295
		_go_fuzz_dep_.CoverTab[64333]++
								m = re.Cap
//line /usr/local/go/src/regexp/syntax/regexp.go:296
		// _ = "end of CoverTab[64333]"
	} else {
//line /usr/local/go/src/regexp/syntax/regexp.go:297
		_go_fuzz_dep_.CoverTab[64334]++
//line /usr/local/go/src/regexp/syntax/regexp.go:297
		// _ = "end of CoverTab[64334]"
//line /usr/local/go/src/regexp/syntax/regexp.go:297
	}
//line /usr/local/go/src/regexp/syntax/regexp.go:297
	// _ = "end of CoverTab[64330]"
//line /usr/local/go/src/regexp/syntax/regexp.go:297
	_go_fuzz_dep_.CoverTab[64331]++
							for _, sub := range re.Sub {
//line /usr/local/go/src/regexp/syntax/regexp.go:298
		_go_fuzz_dep_.CoverTab[64335]++
								if n := sub.MaxCap(); m < n {
//line /usr/local/go/src/regexp/syntax/regexp.go:299
			_go_fuzz_dep_.CoverTab[64336]++
									m = n
//line /usr/local/go/src/regexp/syntax/regexp.go:300
			// _ = "end of CoverTab[64336]"
		} else {
//line /usr/local/go/src/regexp/syntax/regexp.go:301
			_go_fuzz_dep_.CoverTab[64337]++
//line /usr/local/go/src/regexp/syntax/regexp.go:301
			// _ = "end of CoverTab[64337]"
//line /usr/local/go/src/regexp/syntax/regexp.go:301
		}
//line /usr/local/go/src/regexp/syntax/regexp.go:301
		// _ = "end of CoverTab[64335]"
	}
//line /usr/local/go/src/regexp/syntax/regexp.go:302
	// _ = "end of CoverTab[64331]"
//line /usr/local/go/src/regexp/syntax/regexp.go:302
	_go_fuzz_dep_.CoverTab[64332]++
							return m
//line /usr/local/go/src/regexp/syntax/regexp.go:303
	// _ = "end of CoverTab[64332]"
}

//line /usr/local/go/src/regexp/syntax/regexp.go:307
func (re *Regexp) CapNames() []string {
//line /usr/local/go/src/regexp/syntax/regexp.go:307
	_go_fuzz_dep_.CoverTab[64338]++
							names := make([]string, re.MaxCap()+1)
							re.capNames(names)
							return names
//line /usr/local/go/src/regexp/syntax/regexp.go:310
	// _ = "end of CoverTab[64338]"
}

func (re *Regexp) capNames(names []string) {
//line /usr/local/go/src/regexp/syntax/regexp.go:313
	_go_fuzz_dep_.CoverTab[64339]++
							if re.Op == OpCapture {
//line /usr/local/go/src/regexp/syntax/regexp.go:314
		_go_fuzz_dep_.CoverTab[64341]++
								names[re.Cap] = re.Name
//line /usr/local/go/src/regexp/syntax/regexp.go:315
		// _ = "end of CoverTab[64341]"
	} else {
//line /usr/local/go/src/regexp/syntax/regexp.go:316
		_go_fuzz_dep_.CoverTab[64342]++
//line /usr/local/go/src/regexp/syntax/regexp.go:316
		// _ = "end of CoverTab[64342]"
//line /usr/local/go/src/regexp/syntax/regexp.go:316
	}
//line /usr/local/go/src/regexp/syntax/regexp.go:316
	// _ = "end of CoverTab[64339]"
//line /usr/local/go/src/regexp/syntax/regexp.go:316
	_go_fuzz_dep_.CoverTab[64340]++
							for _, sub := range re.Sub {
//line /usr/local/go/src/regexp/syntax/regexp.go:317
		_go_fuzz_dep_.CoverTab[64343]++
								sub.capNames(names)
//line /usr/local/go/src/regexp/syntax/regexp.go:318
		// _ = "end of CoverTab[64343]"
	}
//line /usr/local/go/src/regexp/syntax/regexp.go:319
	// _ = "end of CoverTab[64340]"
}

//line /usr/local/go/src/regexp/syntax/regexp.go:320
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/regexp/syntax/regexp.go:320
var _ = _go_fuzz_dep_.CoverTab
