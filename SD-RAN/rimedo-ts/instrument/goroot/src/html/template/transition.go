// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/html/template/transition.go:5
package template

//line /usr/local/go/src/html/template/transition.go:5
import (
//line /usr/local/go/src/html/template/transition.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/html/template/transition.go:5
)
//line /usr/local/go/src/html/template/transition.go:5
import (
//line /usr/local/go/src/html/template/transition.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/html/template/transition.go:5
)

import (
	"bytes"
	"strings"
)

// transitionFunc is the array of context transition functions for text nodes.
//line /usr/local/go/src/html/template/transition.go:12
// A transition function takes a context and template text input, and returns
//line /usr/local/go/src/html/template/transition.go:12
// the updated context and the number of bytes consumed from the front of the
//line /usr/local/go/src/html/template/transition.go:12
// input.
//line /usr/local/go/src/html/template/transition.go:16
var transitionFunc = [...]func(context, []byte) (context, int){
	stateText:		tText,
	stateTag:		tTag,
	stateAttrName:		tAttrName,
	stateAfterName:		tAfterName,
	stateBeforeValue:	tBeforeValue,
	stateHTMLCmt:		tHTMLCmt,
	stateRCDATA:		tSpecialTagEnd,
	stateAttr:		tAttr,
	stateURL:		tURL,
	stateSrcset:		tURL,
	stateJS:		tJS,
	stateJSDqStr:		tJSDelimited,
	stateJSSqStr:		tJSDelimited,
	stateJSBqStr:		tJSDelimited,
	stateJSRegexp:		tJSDelimited,
	stateJSBlockCmt:	tBlockCmt,
	stateJSLineCmt:		tLineCmt,
	stateCSS:		tCSS,
	stateCSSDqStr:		tCSSStr,
	stateCSSSqStr:		tCSSStr,
	stateCSSDqURL:		tCSSStr,
	stateCSSSqURL:		tCSSStr,
	stateCSSURL:		tCSSStr,
	stateCSSBlockCmt:	tBlockCmt,
	stateCSSLineCmt:	tLineCmt,
	stateError:		tError,
}

var commentStart = []byte("<!--")
var commentEnd = []byte("-->")

// tText is the context transition function for the text state.
func tText(c context, s []byte) (context, int) {
//line /usr/local/go/src/html/template/transition.go:49
	_go_fuzz_dep_.CoverTab[31581]++
								k := 0
								for {
//line /usr/local/go/src/html/template/transition.go:51
		_go_fuzz_dep_.CoverTab[31582]++
									i := k + bytes.IndexByte(s[k:], '<')
									if i < k || func() bool {
//line /usr/local/go/src/html/template/transition.go:53
			_go_fuzz_dep_.CoverTab[31586]++
//line /usr/local/go/src/html/template/transition.go:53
			return i+1 == len(s)
//line /usr/local/go/src/html/template/transition.go:53
			// _ = "end of CoverTab[31586]"
//line /usr/local/go/src/html/template/transition.go:53
		}() {
//line /usr/local/go/src/html/template/transition.go:53
			_go_fuzz_dep_.CoverTab[31587]++
										return c, len(s)
//line /usr/local/go/src/html/template/transition.go:54
			// _ = "end of CoverTab[31587]"
		} else {
//line /usr/local/go/src/html/template/transition.go:55
			_go_fuzz_dep_.CoverTab[31588]++
//line /usr/local/go/src/html/template/transition.go:55
			if i+4 <= len(s) && func() bool {
//line /usr/local/go/src/html/template/transition.go:55
				_go_fuzz_dep_.CoverTab[31589]++
//line /usr/local/go/src/html/template/transition.go:55
				return bytes.Equal(commentStart, s[i:i+4])
//line /usr/local/go/src/html/template/transition.go:55
				// _ = "end of CoverTab[31589]"
//line /usr/local/go/src/html/template/transition.go:55
			}() {
//line /usr/local/go/src/html/template/transition.go:55
				_go_fuzz_dep_.CoverTab[31590]++
											return context{state: stateHTMLCmt}, i + 4
//line /usr/local/go/src/html/template/transition.go:56
				// _ = "end of CoverTab[31590]"
			} else {
//line /usr/local/go/src/html/template/transition.go:57
				_go_fuzz_dep_.CoverTab[31591]++
//line /usr/local/go/src/html/template/transition.go:57
				// _ = "end of CoverTab[31591]"
//line /usr/local/go/src/html/template/transition.go:57
			}
//line /usr/local/go/src/html/template/transition.go:57
			// _ = "end of CoverTab[31588]"
//line /usr/local/go/src/html/template/transition.go:57
		}
//line /usr/local/go/src/html/template/transition.go:57
		// _ = "end of CoverTab[31582]"
//line /usr/local/go/src/html/template/transition.go:57
		_go_fuzz_dep_.CoverTab[31583]++
									i++
									end := false
									if s[i] == '/' {
//line /usr/local/go/src/html/template/transition.go:60
			_go_fuzz_dep_.CoverTab[31592]++
										if i+1 == len(s) {
//line /usr/local/go/src/html/template/transition.go:61
				_go_fuzz_dep_.CoverTab[31594]++
											return c, len(s)
//line /usr/local/go/src/html/template/transition.go:62
				// _ = "end of CoverTab[31594]"
			} else {
//line /usr/local/go/src/html/template/transition.go:63
				_go_fuzz_dep_.CoverTab[31595]++
//line /usr/local/go/src/html/template/transition.go:63
				// _ = "end of CoverTab[31595]"
//line /usr/local/go/src/html/template/transition.go:63
			}
//line /usr/local/go/src/html/template/transition.go:63
			// _ = "end of CoverTab[31592]"
//line /usr/local/go/src/html/template/transition.go:63
			_go_fuzz_dep_.CoverTab[31593]++
										end, i = true, i+1
//line /usr/local/go/src/html/template/transition.go:64
			// _ = "end of CoverTab[31593]"
		} else {
//line /usr/local/go/src/html/template/transition.go:65
			_go_fuzz_dep_.CoverTab[31596]++
//line /usr/local/go/src/html/template/transition.go:65
			// _ = "end of CoverTab[31596]"
//line /usr/local/go/src/html/template/transition.go:65
		}
//line /usr/local/go/src/html/template/transition.go:65
		// _ = "end of CoverTab[31583]"
//line /usr/local/go/src/html/template/transition.go:65
		_go_fuzz_dep_.CoverTab[31584]++
									j, e := eatTagName(s, i)
									if j != i {
//line /usr/local/go/src/html/template/transition.go:67
			_go_fuzz_dep_.CoverTab[31597]++
										if end {
//line /usr/local/go/src/html/template/transition.go:68
				_go_fuzz_dep_.CoverTab[31599]++
											e = elementNone
//line /usr/local/go/src/html/template/transition.go:69
				// _ = "end of CoverTab[31599]"
			} else {
//line /usr/local/go/src/html/template/transition.go:70
				_go_fuzz_dep_.CoverTab[31600]++
//line /usr/local/go/src/html/template/transition.go:70
				// _ = "end of CoverTab[31600]"
//line /usr/local/go/src/html/template/transition.go:70
			}
//line /usr/local/go/src/html/template/transition.go:70
			// _ = "end of CoverTab[31597]"
//line /usr/local/go/src/html/template/transition.go:70
			_go_fuzz_dep_.CoverTab[31598]++

										return context{state: stateTag, element: e}, j
//line /usr/local/go/src/html/template/transition.go:72
			// _ = "end of CoverTab[31598]"
		} else {
//line /usr/local/go/src/html/template/transition.go:73
			_go_fuzz_dep_.CoverTab[31601]++
//line /usr/local/go/src/html/template/transition.go:73
			// _ = "end of CoverTab[31601]"
//line /usr/local/go/src/html/template/transition.go:73
		}
//line /usr/local/go/src/html/template/transition.go:73
		// _ = "end of CoverTab[31584]"
//line /usr/local/go/src/html/template/transition.go:73
		_go_fuzz_dep_.CoverTab[31585]++
									k = j
//line /usr/local/go/src/html/template/transition.go:74
		// _ = "end of CoverTab[31585]"
	}
//line /usr/local/go/src/html/template/transition.go:75
	// _ = "end of CoverTab[31581]"
}

var elementContentType = [...]state{
	elementNone:		stateText,
	elementScript:		stateJS,
	elementStyle:		stateCSS,
	elementTextarea:	stateRCDATA,
	elementTitle:		stateRCDATA,
}

// tTag is the context transition function for the tag state.
func tTag(c context, s []byte) (context, int) {
//line /usr/local/go/src/html/template/transition.go:87
	_go_fuzz_dep_.CoverTab[31602]++

								i := eatWhiteSpace(s, 0)
								if i == len(s) {
//line /usr/local/go/src/html/template/transition.go:90
		_go_fuzz_dep_.CoverTab[31609]++
									return c, len(s)
//line /usr/local/go/src/html/template/transition.go:91
		// _ = "end of CoverTab[31609]"
	} else {
//line /usr/local/go/src/html/template/transition.go:92
		_go_fuzz_dep_.CoverTab[31610]++
//line /usr/local/go/src/html/template/transition.go:92
		// _ = "end of CoverTab[31610]"
//line /usr/local/go/src/html/template/transition.go:92
	}
//line /usr/local/go/src/html/template/transition.go:92
	// _ = "end of CoverTab[31602]"
//line /usr/local/go/src/html/template/transition.go:92
	_go_fuzz_dep_.CoverTab[31603]++
								if s[i] == '>' {
//line /usr/local/go/src/html/template/transition.go:93
		_go_fuzz_dep_.CoverTab[31611]++
									return context{
			state:		elementContentType[c.element],
			element:	c.element,
		}, i + 1
//line /usr/local/go/src/html/template/transition.go:97
		// _ = "end of CoverTab[31611]"
	} else {
//line /usr/local/go/src/html/template/transition.go:98
		_go_fuzz_dep_.CoverTab[31612]++
//line /usr/local/go/src/html/template/transition.go:98
		// _ = "end of CoverTab[31612]"
//line /usr/local/go/src/html/template/transition.go:98
	}
//line /usr/local/go/src/html/template/transition.go:98
	// _ = "end of CoverTab[31603]"
//line /usr/local/go/src/html/template/transition.go:98
	_go_fuzz_dep_.CoverTab[31604]++
								j, err := eatAttrName(s, i)
								if err != nil {
//line /usr/local/go/src/html/template/transition.go:100
		_go_fuzz_dep_.CoverTab[31613]++
									return context{state: stateError, err: err}, len(s)
//line /usr/local/go/src/html/template/transition.go:101
		// _ = "end of CoverTab[31613]"
	} else {
//line /usr/local/go/src/html/template/transition.go:102
		_go_fuzz_dep_.CoverTab[31614]++
//line /usr/local/go/src/html/template/transition.go:102
		// _ = "end of CoverTab[31614]"
//line /usr/local/go/src/html/template/transition.go:102
	}
//line /usr/local/go/src/html/template/transition.go:102
	// _ = "end of CoverTab[31604]"
//line /usr/local/go/src/html/template/transition.go:102
	_go_fuzz_dep_.CoverTab[31605]++
								state, attr := stateTag, attrNone
								if i == j {
//line /usr/local/go/src/html/template/transition.go:104
		_go_fuzz_dep_.CoverTab[31615]++
									return context{
			state:	stateError,
			err:	errorf(ErrBadHTML, nil, 0, "expected space, attr name, or end of tag, but got %q", s[i:]),
		}, len(s)
//line /usr/local/go/src/html/template/transition.go:108
		// _ = "end of CoverTab[31615]"
	} else {
//line /usr/local/go/src/html/template/transition.go:109
		_go_fuzz_dep_.CoverTab[31616]++
//line /usr/local/go/src/html/template/transition.go:109
		// _ = "end of CoverTab[31616]"
//line /usr/local/go/src/html/template/transition.go:109
	}
//line /usr/local/go/src/html/template/transition.go:109
	// _ = "end of CoverTab[31605]"
//line /usr/local/go/src/html/template/transition.go:109
	_go_fuzz_dep_.CoverTab[31606]++

								attrName := strings.ToLower(string(s[i:j]))
								if c.element == elementScript && func() bool {
//line /usr/local/go/src/html/template/transition.go:112
		_go_fuzz_dep_.CoverTab[31617]++
//line /usr/local/go/src/html/template/transition.go:112
		return attrName == "type"
//line /usr/local/go/src/html/template/transition.go:112
		// _ = "end of CoverTab[31617]"
//line /usr/local/go/src/html/template/transition.go:112
	}() {
//line /usr/local/go/src/html/template/transition.go:112
		_go_fuzz_dep_.CoverTab[31618]++
									attr = attrScriptType
//line /usr/local/go/src/html/template/transition.go:113
		// _ = "end of CoverTab[31618]"
	} else {
//line /usr/local/go/src/html/template/transition.go:114
		_go_fuzz_dep_.CoverTab[31619]++
									switch attrType(attrName) {
		case contentTypeURL:
//line /usr/local/go/src/html/template/transition.go:116
			_go_fuzz_dep_.CoverTab[31620]++
										attr = attrURL
//line /usr/local/go/src/html/template/transition.go:117
			// _ = "end of CoverTab[31620]"
		case contentTypeCSS:
//line /usr/local/go/src/html/template/transition.go:118
			_go_fuzz_dep_.CoverTab[31621]++
										attr = attrStyle
//line /usr/local/go/src/html/template/transition.go:119
			// _ = "end of CoverTab[31621]"
		case contentTypeJS:
//line /usr/local/go/src/html/template/transition.go:120
			_go_fuzz_dep_.CoverTab[31622]++
										attr = attrScript
//line /usr/local/go/src/html/template/transition.go:121
			// _ = "end of CoverTab[31622]"
		case contentTypeSrcset:
//line /usr/local/go/src/html/template/transition.go:122
			_go_fuzz_dep_.CoverTab[31623]++
										attr = attrSrcset
//line /usr/local/go/src/html/template/transition.go:123
			// _ = "end of CoverTab[31623]"
//line /usr/local/go/src/html/template/transition.go:123
		default:
//line /usr/local/go/src/html/template/transition.go:123
			_go_fuzz_dep_.CoverTab[31624]++
//line /usr/local/go/src/html/template/transition.go:123
			// _ = "end of CoverTab[31624]"
		}
//line /usr/local/go/src/html/template/transition.go:124
		// _ = "end of CoverTab[31619]"
	}
//line /usr/local/go/src/html/template/transition.go:125
	// _ = "end of CoverTab[31606]"
//line /usr/local/go/src/html/template/transition.go:125
	_go_fuzz_dep_.CoverTab[31607]++

								if j == len(s) {
//line /usr/local/go/src/html/template/transition.go:127
		_go_fuzz_dep_.CoverTab[31625]++
									state = stateAttrName
//line /usr/local/go/src/html/template/transition.go:128
		// _ = "end of CoverTab[31625]"
	} else {
//line /usr/local/go/src/html/template/transition.go:129
		_go_fuzz_dep_.CoverTab[31626]++
									state = stateAfterName
//line /usr/local/go/src/html/template/transition.go:130
		// _ = "end of CoverTab[31626]"
	}
//line /usr/local/go/src/html/template/transition.go:131
	// _ = "end of CoverTab[31607]"
//line /usr/local/go/src/html/template/transition.go:131
	_go_fuzz_dep_.CoverTab[31608]++
								return context{state: state, element: c.element, attr: attr}, j
//line /usr/local/go/src/html/template/transition.go:132
	// _ = "end of CoverTab[31608]"
}

// tAttrName is the context transition function for stateAttrName.
func tAttrName(c context, s []byte) (context, int) {
//line /usr/local/go/src/html/template/transition.go:136
	_go_fuzz_dep_.CoverTab[31627]++
								i, err := eatAttrName(s, 0)
								if err != nil {
//line /usr/local/go/src/html/template/transition.go:138
		_go_fuzz_dep_.CoverTab[31629]++
									return context{state: stateError, err: err}, len(s)
//line /usr/local/go/src/html/template/transition.go:139
		// _ = "end of CoverTab[31629]"
	} else {
//line /usr/local/go/src/html/template/transition.go:140
		_go_fuzz_dep_.CoverTab[31630]++
//line /usr/local/go/src/html/template/transition.go:140
		if i != len(s) {
//line /usr/local/go/src/html/template/transition.go:140
			_go_fuzz_dep_.CoverTab[31631]++
										c.state = stateAfterName
//line /usr/local/go/src/html/template/transition.go:141
			// _ = "end of CoverTab[31631]"
		} else {
//line /usr/local/go/src/html/template/transition.go:142
			_go_fuzz_dep_.CoverTab[31632]++
//line /usr/local/go/src/html/template/transition.go:142
			// _ = "end of CoverTab[31632]"
//line /usr/local/go/src/html/template/transition.go:142
		}
//line /usr/local/go/src/html/template/transition.go:142
		// _ = "end of CoverTab[31630]"
//line /usr/local/go/src/html/template/transition.go:142
	}
//line /usr/local/go/src/html/template/transition.go:142
	// _ = "end of CoverTab[31627]"
//line /usr/local/go/src/html/template/transition.go:142
	_go_fuzz_dep_.CoverTab[31628]++
								return c, i
//line /usr/local/go/src/html/template/transition.go:143
	// _ = "end of CoverTab[31628]"
}

// tAfterName is the context transition function for stateAfterName.
func tAfterName(c context, s []byte) (context, int) {
//line /usr/local/go/src/html/template/transition.go:147
	_go_fuzz_dep_.CoverTab[31633]++

								i := eatWhiteSpace(s, 0)
								if i == len(s) {
//line /usr/local/go/src/html/template/transition.go:150
		_go_fuzz_dep_.CoverTab[31635]++
									return c, len(s)
//line /usr/local/go/src/html/template/transition.go:151
		// _ = "end of CoverTab[31635]"
	} else {
//line /usr/local/go/src/html/template/transition.go:152
		_go_fuzz_dep_.CoverTab[31636]++
//line /usr/local/go/src/html/template/transition.go:152
		if s[i] != '=' {
//line /usr/local/go/src/html/template/transition.go:152
			_go_fuzz_dep_.CoverTab[31637]++

										c.state = stateTag
										return c, i
//line /usr/local/go/src/html/template/transition.go:155
			// _ = "end of CoverTab[31637]"
		} else {
//line /usr/local/go/src/html/template/transition.go:156
			_go_fuzz_dep_.CoverTab[31638]++
//line /usr/local/go/src/html/template/transition.go:156
			// _ = "end of CoverTab[31638]"
//line /usr/local/go/src/html/template/transition.go:156
		}
//line /usr/local/go/src/html/template/transition.go:156
		// _ = "end of CoverTab[31636]"
//line /usr/local/go/src/html/template/transition.go:156
	}
//line /usr/local/go/src/html/template/transition.go:156
	// _ = "end of CoverTab[31633]"
//line /usr/local/go/src/html/template/transition.go:156
	_go_fuzz_dep_.CoverTab[31634]++
								c.state = stateBeforeValue

								return c, i + 1
//line /usr/local/go/src/html/template/transition.go:159
	// _ = "end of CoverTab[31634]"
}

var attrStartStates = [...]state{
	attrNone:	stateAttr,
	attrScript:	stateJS,
	attrScriptType:	stateAttr,
	attrStyle:	stateCSS,
	attrURL:	stateURL,
	attrSrcset:	stateSrcset,
}

// tBeforeValue is the context transition function for stateBeforeValue.
func tBeforeValue(c context, s []byte) (context, int) {
//line /usr/local/go/src/html/template/transition.go:172
	_go_fuzz_dep_.CoverTab[31639]++
								i := eatWhiteSpace(s, 0)
								if i == len(s) {
//line /usr/local/go/src/html/template/transition.go:174
		_go_fuzz_dep_.CoverTab[31642]++
									return c, len(s)
//line /usr/local/go/src/html/template/transition.go:175
		// _ = "end of CoverTab[31642]"
	} else {
//line /usr/local/go/src/html/template/transition.go:176
		_go_fuzz_dep_.CoverTab[31643]++
//line /usr/local/go/src/html/template/transition.go:176
		// _ = "end of CoverTab[31643]"
//line /usr/local/go/src/html/template/transition.go:176
	}
//line /usr/local/go/src/html/template/transition.go:176
	// _ = "end of CoverTab[31639]"
//line /usr/local/go/src/html/template/transition.go:176
	_go_fuzz_dep_.CoverTab[31640]++

								delim := delimSpaceOrTagEnd
								switch s[i] {
	case '\'':
//line /usr/local/go/src/html/template/transition.go:180
		_go_fuzz_dep_.CoverTab[31644]++
									delim, i = delimSingleQuote, i+1
//line /usr/local/go/src/html/template/transition.go:181
		// _ = "end of CoverTab[31644]"
	case '"':
//line /usr/local/go/src/html/template/transition.go:182
		_go_fuzz_dep_.CoverTab[31645]++
									delim, i = delimDoubleQuote, i+1
//line /usr/local/go/src/html/template/transition.go:183
		// _ = "end of CoverTab[31645]"
//line /usr/local/go/src/html/template/transition.go:183
	default:
//line /usr/local/go/src/html/template/transition.go:183
		_go_fuzz_dep_.CoverTab[31646]++
//line /usr/local/go/src/html/template/transition.go:183
		// _ = "end of CoverTab[31646]"
	}
//line /usr/local/go/src/html/template/transition.go:184
	// _ = "end of CoverTab[31640]"
//line /usr/local/go/src/html/template/transition.go:184
	_go_fuzz_dep_.CoverTab[31641]++
								c.state, c.delim = attrStartStates[c.attr], delim
								return c, i
//line /usr/local/go/src/html/template/transition.go:186
	// _ = "end of CoverTab[31641]"
}

// tHTMLCmt is the context transition function for stateHTMLCmt.
func tHTMLCmt(c context, s []byte) (context, int) {
//line /usr/local/go/src/html/template/transition.go:190
	_go_fuzz_dep_.CoverTab[31647]++
								if i := bytes.Index(s, commentEnd); i != -1 {
//line /usr/local/go/src/html/template/transition.go:191
		_go_fuzz_dep_.CoverTab[31649]++
									return context{}, i + 3
//line /usr/local/go/src/html/template/transition.go:192
		// _ = "end of CoverTab[31649]"
	} else {
//line /usr/local/go/src/html/template/transition.go:193
		_go_fuzz_dep_.CoverTab[31650]++
//line /usr/local/go/src/html/template/transition.go:193
		// _ = "end of CoverTab[31650]"
//line /usr/local/go/src/html/template/transition.go:193
	}
//line /usr/local/go/src/html/template/transition.go:193
	// _ = "end of CoverTab[31647]"
//line /usr/local/go/src/html/template/transition.go:193
	_go_fuzz_dep_.CoverTab[31648]++
								return c, len(s)
//line /usr/local/go/src/html/template/transition.go:194
	// _ = "end of CoverTab[31648]"
}

// specialTagEndMarkers maps element types to the character sequence that
//line /usr/local/go/src/html/template/transition.go:197
// case-insensitively signals the end of the special tag body.
//line /usr/local/go/src/html/template/transition.go:199
var specialTagEndMarkers = [...][]byte{
	elementScript:		[]byte("script"),
	elementStyle:		[]byte("style"),
	elementTextarea:	[]byte("textarea"),
	elementTitle:		[]byte("title"),
}

var (
	specialTagEndPrefix	= []byte("</")
	tagEndSeparators	= []byte("> \t\n\f/")
)

// tSpecialTagEnd is the context transition function for raw text and RCDATA
//line /usr/local/go/src/html/template/transition.go:211
// element states.
//line /usr/local/go/src/html/template/transition.go:213
func tSpecialTagEnd(c context, s []byte) (context, int) {
//line /usr/local/go/src/html/template/transition.go:213
	_go_fuzz_dep_.CoverTab[31651]++
								if c.element != elementNone {
//line /usr/local/go/src/html/template/transition.go:214
		_go_fuzz_dep_.CoverTab[31653]++
									if i := indexTagEnd(s, specialTagEndMarkers[c.element]); i != -1 {
//line /usr/local/go/src/html/template/transition.go:215
			_go_fuzz_dep_.CoverTab[31654]++
										return context{}, i
//line /usr/local/go/src/html/template/transition.go:216
			// _ = "end of CoverTab[31654]"
		} else {
//line /usr/local/go/src/html/template/transition.go:217
			_go_fuzz_dep_.CoverTab[31655]++
//line /usr/local/go/src/html/template/transition.go:217
			// _ = "end of CoverTab[31655]"
//line /usr/local/go/src/html/template/transition.go:217
		}
//line /usr/local/go/src/html/template/transition.go:217
		// _ = "end of CoverTab[31653]"
	} else {
//line /usr/local/go/src/html/template/transition.go:218
		_go_fuzz_dep_.CoverTab[31656]++
//line /usr/local/go/src/html/template/transition.go:218
		// _ = "end of CoverTab[31656]"
//line /usr/local/go/src/html/template/transition.go:218
	}
//line /usr/local/go/src/html/template/transition.go:218
	// _ = "end of CoverTab[31651]"
//line /usr/local/go/src/html/template/transition.go:218
	_go_fuzz_dep_.CoverTab[31652]++
								return c, len(s)
//line /usr/local/go/src/html/template/transition.go:219
	// _ = "end of CoverTab[31652]"
}

// indexTagEnd finds the index of a special tag end in a case insensitive way, or returns -1
func indexTagEnd(s []byte, tag []byte) int {
//line /usr/local/go/src/html/template/transition.go:223
	_go_fuzz_dep_.CoverTab[31657]++
								res := 0
								plen := len(specialTagEndPrefix)
								for len(s) > 0 {
//line /usr/local/go/src/html/template/transition.go:226
		_go_fuzz_dep_.CoverTab[31659]++

									i := bytes.Index(s, specialTagEndPrefix)
									if i == -1 {
//line /usr/local/go/src/html/template/transition.go:229
			_go_fuzz_dep_.CoverTab[31662]++
										return i
//line /usr/local/go/src/html/template/transition.go:230
			// _ = "end of CoverTab[31662]"
		} else {
//line /usr/local/go/src/html/template/transition.go:231
			_go_fuzz_dep_.CoverTab[31663]++
//line /usr/local/go/src/html/template/transition.go:231
			// _ = "end of CoverTab[31663]"
//line /usr/local/go/src/html/template/transition.go:231
		}
//line /usr/local/go/src/html/template/transition.go:231
		// _ = "end of CoverTab[31659]"
//line /usr/local/go/src/html/template/transition.go:231
		_go_fuzz_dep_.CoverTab[31660]++
									s = s[i+plen:]

									if len(tag) <= len(s) && func() bool {
//line /usr/local/go/src/html/template/transition.go:234
			_go_fuzz_dep_.CoverTab[31664]++
//line /usr/local/go/src/html/template/transition.go:234
			return bytes.EqualFold(tag, s[:len(tag)])
//line /usr/local/go/src/html/template/transition.go:234
			// _ = "end of CoverTab[31664]"
//line /usr/local/go/src/html/template/transition.go:234
		}() {
//line /usr/local/go/src/html/template/transition.go:234
			_go_fuzz_dep_.CoverTab[31665]++
										s = s[len(tag):]

										if len(s) > 0 && func() bool {
//line /usr/local/go/src/html/template/transition.go:237
				_go_fuzz_dep_.CoverTab[31667]++
//line /usr/local/go/src/html/template/transition.go:237
				return bytes.IndexByte(tagEndSeparators, s[0]) != -1
//line /usr/local/go/src/html/template/transition.go:237
				// _ = "end of CoverTab[31667]"
//line /usr/local/go/src/html/template/transition.go:237
			}() {
//line /usr/local/go/src/html/template/transition.go:237
				_go_fuzz_dep_.CoverTab[31668]++
											return res + i
//line /usr/local/go/src/html/template/transition.go:238
				// _ = "end of CoverTab[31668]"
			} else {
//line /usr/local/go/src/html/template/transition.go:239
				_go_fuzz_dep_.CoverTab[31669]++
//line /usr/local/go/src/html/template/transition.go:239
				// _ = "end of CoverTab[31669]"
//line /usr/local/go/src/html/template/transition.go:239
			}
//line /usr/local/go/src/html/template/transition.go:239
			// _ = "end of CoverTab[31665]"
//line /usr/local/go/src/html/template/transition.go:239
			_go_fuzz_dep_.CoverTab[31666]++
										res += len(tag)
//line /usr/local/go/src/html/template/transition.go:240
			// _ = "end of CoverTab[31666]"
		} else {
//line /usr/local/go/src/html/template/transition.go:241
			_go_fuzz_dep_.CoverTab[31670]++
//line /usr/local/go/src/html/template/transition.go:241
			// _ = "end of CoverTab[31670]"
//line /usr/local/go/src/html/template/transition.go:241
		}
//line /usr/local/go/src/html/template/transition.go:241
		// _ = "end of CoverTab[31660]"
//line /usr/local/go/src/html/template/transition.go:241
		_go_fuzz_dep_.CoverTab[31661]++
									res += i + plen
//line /usr/local/go/src/html/template/transition.go:242
		// _ = "end of CoverTab[31661]"
	}
//line /usr/local/go/src/html/template/transition.go:243
	// _ = "end of CoverTab[31657]"
//line /usr/local/go/src/html/template/transition.go:243
	_go_fuzz_dep_.CoverTab[31658]++
								return -1
//line /usr/local/go/src/html/template/transition.go:244
	// _ = "end of CoverTab[31658]"
}

// tAttr is the context transition function for the attribute state.
func tAttr(c context, s []byte) (context, int) {
//line /usr/local/go/src/html/template/transition.go:248
	_go_fuzz_dep_.CoverTab[31671]++
								return c, len(s)
//line /usr/local/go/src/html/template/transition.go:249
	// _ = "end of CoverTab[31671]"
}

// tURL is the context transition function for the URL state.
func tURL(c context, s []byte) (context, int) {
//line /usr/local/go/src/html/template/transition.go:253
	_go_fuzz_dep_.CoverTab[31672]++
								if bytes.ContainsAny(s, "#?") {
//line /usr/local/go/src/html/template/transition.go:254
		_go_fuzz_dep_.CoverTab[31674]++
									c.urlPart = urlPartQueryOrFrag
//line /usr/local/go/src/html/template/transition.go:255
		// _ = "end of CoverTab[31674]"
	} else {
//line /usr/local/go/src/html/template/transition.go:256
		_go_fuzz_dep_.CoverTab[31675]++
//line /usr/local/go/src/html/template/transition.go:256
		if len(s) != eatWhiteSpace(s, 0) && func() bool {
//line /usr/local/go/src/html/template/transition.go:256
			_go_fuzz_dep_.CoverTab[31676]++
//line /usr/local/go/src/html/template/transition.go:256
			return c.urlPart == urlPartNone
//line /usr/local/go/src/html/template/transition.go:256
			// _ = "end of CoverTab[31676]"
//line /usr/local/go/src/html/template/transition.go:256
		}() {
//line /usr/local/go/src/html/template/transition.go:256
			_go_fuzz_dep_.CoverTab[31677]++

//line /usr/local/go/src/html/template/transition.go:259
			c.urlPart = urlPartPreQuery
//line /usr/local/go/src/html/template/transition.go:259
			// _ = "end of CoverTab[31677]"
		} else {
//line /usr/local/go/src/html/template/transition.go:260
			_go_fuzz_dep_.CoverTab[31678]++
//line /usr/local/go/src/html/template/transition.go:260
			// _ = "end of CoverTab[31678]"
//line /usr/local/go/src/html/template/transition.go:260
		}
//line /usr/local/go/src/html/template/transition.go:260
		// _ = "end of CoverTab[31675]"
//line /usr/local/go/src/html/template/transition.go:260
	}
//line /usr/local/go/src/html/template/transition.go:260
	// _ = "end of CoverTab[31672]"
//line /usr/local/go/src/html/template/transition.go:260
	_go_fuzz_dep_.CoverTab[31673]++
								return c, len(s)
//line /usr/local/go/src/html/template/transition.go:261
	// _ = "end of CoverTab[31673]"
}

// tJS is the context transition function for the JS state.
func tJS(c context, s []byte) (context, int) {
//line /usr/local/go/src/html/template/transition.go:265
	_go_fuzz_dep_.CoverTab[31679]++
								i := bytes.IndexAny(s, "\"`'/")
								if i == -1 {
//line /usr/local/go/src/html/template/transition.go:267
		_go_fuzz_dep_.CoverTab[31682]++

									c.jsCtx = nextJSCtx(s, c.jsCtx)
									return c, len(s)
//line /usr/local/go/src/html/template/transition.go:270
		// _ = "end of CoverTab[31682]"
	} else {
//line /usr/local/go/src/html/template/transition.go:271
		_go_fuzz_dep_.CoverTab[31683]++
//line /usr/local/go/src/html/template/transition.go:271
		// _ = "end of CoverTab[31683]"
//line /usr/local/go/src/html/template/transition.go:271
	}
//line /usr/local/go/src/html/template/transition.go:271
	// _ = "end of CoverTab[31679]"
//line /usr/local/go/src/html/template/transition.go:271
	_go_fuzz_dep_.CoverTab[31680]++
								c.jsCtx = nextJSCtx(s[:i], c.jsCtx)
								switch s[i] {
	case '"':
//line /usr/local/go/src/html/template/transition.go:274
		_go_fuzz_dep_.CoverTab[31684]++
									c.state, c.jsCtx = stateJSDqStr, jsCtxRegexp
//line /usr/local/go/src/html/template/transition.go:275
		// _ = "end of CoverTab[31684]"
	case '\'':
//line /usr/local/go/src/html/template/transition.go:276
		_go_fuzz_dep_.CoverTab[31685]++
									c.state, c.jsCtx = stateJSSqStr, jsCtxRegexp
//line /usr/local/go/src/html/template/transition.go:277
		// _ = "end of CoverTab[31685]"
	case '`':
//line /usr/local/go/src/html/template/transition.go:278
		_go_fuzz_dep_.CoverTab[31686]++
									c.state, c.jsCtx = stateJSBqStr, jsCtxRegexp
//line /usr/local/go/src/html/template/transition.go:279
		// _ = "end of CoverTab[31686]"
	case '/':
//line /usr/local/go/src/html/template/transition.go:280
		_go_fuzz_dep_.CoverTab[31687]++
									switch {
		case i+1 < len(s) && func() bool {
//line /usr/local/go/src/html/template/transition.go:282
			_go_fuzz_dep_.CoverTab[31694]++
//line /usr/local/go/src/html/template/transition.go:282
			return s[i+1] == '/'
//line /usr/local/go/src/html/template/transition.go:282
			// _ = "end of CoverTab[31694]"
//line /usr/local/go/src/html/template/transition.go:282
		}():
//line /usr/local/go/src/html/template/transition.go:282
			_go_fuzz_dep_.CoverTab[31689]++
										c.state, i = stateJSLineCmt, i+1
//line /usr/local/go/src/html/template/transition.go:283
			// _ = "end of CoverTab[31689]"
		case i+1 < len(s) && func() bool {
//line /usr/local/go/src/html/template/transition.go:284
			_go_fuzz_dep_.CoverTab[31695]++
//line /usr/local/go/src/html/template/transition.go:284
			return s[i+1] == '*'
//line /usr/local/go/src/html/template/transition.go:284
			// _ = "end of CoverTab[31695]"
//line /usr/local/go/src/html/template/transition.go:284
		}():
//line /usr/local/go/src/html/template/transition.go:284
			_go_fuzz_dep_.CoverTab[31690]++
										c.state, i = stateJSBlockCmt, i+1
//line /usr/local/go/src/html/template/transition.go:285
			// _ = "end of CoverTab[31690]"
		case c.jsCtx == jsCtxRegexp:
//line /usr/local/go/src/html/template/transition.go:286
			_go_fuzz_dep_.CoverTab[31691]++
										c.state = stateJSRegexp
//line /usr/local/go/src/html/template/transition.go:287
			// _ = "end of CoverTab[31691]"
		case c.jsCtx == jsCtxDivOp:
//line /usr/local/go/src/html/template/transition.go:288
			_go_fuzz_dep_.CoverTab[31692]++
										c.jsCtx = jsCtxRegexp
//line /usr/local/go/src/html/template/transition.go:289
			// _ = "end of CoverTab[31692]"
		default:
//line /usr/local/go/src/html/template/transition.go:290
			_go_fuzz_dep_.CoverTab[31693]++
										return context{
				state:	stateError,
				err:	errorf(ErrSlashAmbig, nil, 0, "'/' could start a division or regexp: %.32q", s[i:]),
			}, len(s)
//line /usr/local/go/src/html/template/transition.go:294
			// _ = "end of CoverTab[31693]"
		}
//line /usr/local/go/src/html/template/transition.go:295
		// _ = "end of CoverTab[31687]"
	default:
//line /usr/local/go/src/html/template/transition.go:296
		_go_fuzz_dep_.CoverTab[31688]++
									panic("unreachable")
//line /usr/local/go/src/html/template/transition.go:297
		// _ = "end of CoverTab[31688]"
	}
//line /usr/local/go/src/html/template/transition.go:298
	// _ = "end of CoverTab[31680]"
//line /usr/local/go/src/html/template/transition.go:298
	_go_fuzz_dep_.CoverTab[31681]++
								return c, i + 1
//line /usr/local/go/src/html/template/transition.go:299
	// _ = "end of CoverTab[31681]"
}

// tJSDelimited is the context transition function for the JS string and regexp
//line /usr/local/go/src/html/template/transition.go:302
// states.
//line /usr/local/go/src/html/template/transition.go:304
func tJSDelimited(c context, s []byte) (context, int) {
//line /usr/local/go/src/html/template/transition.go:304
	_go_fuzz_dep_.CoverTab[31696]++
								specials := `\"`
								switch c.state {
	case stateJSSqStr:
//line /usr/local/go/src/html/template/transition.go:307
		_go_fuzz_dep_.CoverTab[31700]++
									specials = `\'`
//line /usr/local/go/src/html/template/transition.go:308
		// _ = "end of CoverTab[31700]"
	case stateJSBqStr:
//line /usr/local/go/src/html/template/transition.go:309
		_go_fuzz_dep_.CoverTab[31701]++
									specials = "`\\"
//line /usr/local/go/src/html/template/transition.go:310
		// _ = "end of CoverTab[31701]"
	case stateJSRegexp:
//line /usr/local/go/src/html/template/transition.go:311
		_go_fuzz_dep_.CoverTab[31702]++
									specials = `\/[]`
//line /usr/local/go/src/html/template/transition.go:312
		// _ = "end of CoverTab[31702]"
//line /usr/local/go/src/html/template/transition.go:312
	default:
//line /usr/local/go/src/html/template/transition.go:312
		_go_fuzz_dep_.CoverTab[31703]++
//line /usr/local/go/src/html/template/transition.go:312
		// _ = "end of CoverTab[31703]"
	}
//line /usr/local/go/src/html/template/transition.go:313
	// _ = "end of CoverTab[31696]"
//line /usr/local/go/src/html/template/transition.go:313
	_go_fuzz_dep_.CoverTab[31697]++

								k, inCharset := 0, false
								for {
//line /usr/local/go/src/html/template/transition.go:316
		_go_fuzz_dep_.CoverTab[31704]++
									i := k + bytes.IndexAny(s[k:], specials)
									if i < k {
//line /usr/local/go/src/html/template/transition.go:318
			_go_fuzz_dep_.CoverTab[31707]++
										break
//line /usr/local/go/src/html/template/transition.go:319
			// _ = "end of CoverTab[31707]"
		} else {
//line /usr/local/go/src/html/template/transition.go:320
			_go_fuzz_dep_.CoverTab[31708]++
//line /usr/local/go/src/html/template/transition.go:320
			// _ = "end of CoverTab[31708]"
//line /usr/local/go/src/html/template/transition.go:320
		}
//line /usr/local/go/src/html/template/transition.go:320
		// _ = "end of CoverTab[31704]"
//line /usr/local/go/src/html/template/transition.go:320
		_go_fuzz_dep_.CoverTab[31705]++
									switch s[i] {
		case '\\':
//line /usr/local/go/src/html/template/transition.go:322
			_go_fuzz_dep_.CoverTab[31709]++
										i++
										if i == len(s) {
//line /usr/local/go/src/html/template/transition.go:324
				_go_fuzz_dep_.CoverTab[31713]++
											return context{
					state:	stateError,
					err:	errorf(ErrPartialEscape, nil, 0, "unfinished escape sequence in JS string: %q", s),
				}, len(s)
//line /usr/local/go/src/html/template/transition.go:328
				// _ = "end of CoverTab[31713]"
			} else {
//line /usr/local/go/src/html/template/transition.go:329
				_go_fuzz_dep_.CoverTab[31714]++
//line /usr/local/go/src/html/template/transition.go:329
				// _ = "end of CoverTab[31714]"
//line /usr/local/go/src/html/template/transition.go:329
			}
//line /usr/local/go/src/html/template/transition.go:329
			// _ = "end of CoverTab[31709]"
		case '[':
//line /usr/local/go/src/html/template/transition.go:330
			_go_fuzz_dep_.CoverTab[31710]++
										inCharset = true
//line /usr/local/go/src/html/template/transition.go:331
			// _ = "end of CoverTab[31710]"
		case ']':
//line /usr/local/go/src/html/template/transition.go:332
			_go_fuzz_dep_.CoverTab[31711]++
										inCharset = false
//line /usr/local/go/src/html/template/transition.go:333
			// _ = "end of CoverTab[31711]"
		default:
//line /usr/local/go/src/html/template/transition.go:334
			_go_fuzz_dep_.CoverTab[31712]++

										if !inCharset {
//line /usr/local/go/src/html/template/transition.go:336
				_go_fuzz_dep_.CoverTab[31715]++
											c.state, c.jsCtx = stateJS, jsCtxDivOp
											return c, i + 1
//line /usr/local/go/src/html/template/transition.go:338
				// _ = "end of CoverTab[31715]"
			} else {
//line /usr/local/go/src/html/template/transition.go:339
				_go_fuzz_dep_.CoverTab[31716]++
//line /usr/local/go/src/html/template/transition.go:339
				// _ = "end of CoverTab[31716]"
//line /usr/local/go/src/html/template/transition.go:339
			}
//line /usr/local/go/src/html/template/transition.go:339
			// _ = "end of CoverTab[31712]"
		}
//line /usr/local/go/src/html/template/transition.go:340
		// _ = "end of CoverTab[31705]"
//line /usr/local/go/src/html/template/transition.go:340
		_go_fuzz_dep_.CoverTab[31706]++
									k = i + 1
//line /usr/local/go/src/html/template/transition.go:341
		// _ = "end of CoverTab[31706]"
	}
//line /usr/local/go/src/html/template/transition.go:342
	// _ = "end of CoverTab[31697]"
//line /usr/local/go/src/html/template/transition.go:342
	_go_fuzz_dep_.CoverTab[31698]++

								if inCharset {
//line /usr/local/go/src/html/template/transition.go:344
		_go_fuzz_dep_.CoverTab[31717]++

//line /usr/local/go/src/html/template/transition.go:347
		return context{
			state:	stateError,
			err:	errorf(ErrPartialCharset, nil, 0, "unfinished JS regexp charset: %q", s),
		}, len(s)
//line /usr/local/go/src/html/template/transition.go:350
		// _ = "end of CoverTab[31717]"
	} else {
//line /usr/local/go/src/html/template/transition.go:351
		_go_fuzz_dep_.CoverTab[31718]++
//line /usr/local/go/src/html/template/transition.go:351
		// _ = "end of CoverTab[31718]"
//line /usr/local/go/src/html/template/transition.go:351
	}
//line /usr/local/go/src/html/template/transition.go:351
	// _ = "end of CoverTab[31698]"
//line /usr/local/go/src/html/template/transition.go:351
	_go_fuzz_dep_.CoverTab[31699]++

								return c, len(s)
//line /usr/local/go/src/html/template/transition.go:353
	// _ = "end of CoverTab[31699]"
}

var blockCommentEnd = []byte("*/")

// tBlockCmt is the context transition function for /*comment*/ states.
func tBlockCmt(c context, s []byte) (context, int) {
//line /usr/local/go/src/html/template/transition.go:359
	_go_fuzz_dep_.CoverTab[31719]++
								i := bytes.Index(s, blockCommentEnd)
								if i == -1 {
//line /usr/local/go/src/html/template/transition.go:361
		_go_fuzz_dep_.CoverTab[31722]++
									return c, len(s)
//line /usr/local/go/src/html/template/transition.go:362
		// _ = "end of CoverTab[31722]"
	} else {
//line /usr/local/go/src/html/template/transition.go:363
		_go_fuzz_dep_.CoverTab[31723]++
//line /usr/local/go/src/html/template/transition.go:363
		// _ = "end of CoverTab[31723]"
//line /usr/local/go/src/html/template/transition.go:363
	}
//line /usr/local/go/src/html/template/transition.go:363
	// _ = "end of CoverTab[31719]"
//line /usr/local/go/src/html/template/transition.go:363
	_go_fuzz_dep_.CoverTab[31720]++
								switch c.state {
	case stateJSBlockCmt:
//line /usr/local/go/src/html/template/transition.go:365
		_go_fuzz_dep_.CoverTab[31724]++
									c.state = stateJS
//line /usr/local/go/src/html/template/transition.go:366
		// _ = "end of CoverTab[31724]"
	case stateCSSBlockCmt:
//line /usr/local/go/src/html/template/transition.go:367
		_go_fuzz_dep_.CoverTab[31725]++
									c.state = stateCSS
//line /usr/local/go/src/html/template/transition.go:368
		// _ = "end of CoverTab[31725]"
	default:
//line /usr/local/go/src/html/template/transition.go:369
		_go_fuzz_dep_.CoverTab[31726]++
									panic(c.state.String())
//line /usr/local/go/src/html/template/transition.go:370
		// _ = "end of CoverTab[31726]"
	}
//line /usr/local/go/src/html/template/transition.go:371
	// _ = "end of CoverTab[31720]"
//line /usr/local/go/src/html/template/transition.go:371
	_go_fuzz_dep_.CoverTab[31721]++
								return c, i + 2
//line /usr/local/go/src/html/template/transition.go:372
	// _ = "end of CoverTab[31721]"
}

// tLineCmt is the context transition function for //comment states.
func tLineCmt(c context, s []byte) (context, int) {
//line /usr/local/go/src/html/template/transition.go:376
	_go_fuzz_dep_.CoverTab[31727]++
								var lineTerminators string
								var endState state
								switch c.state {
	case stateJSLineCmt:
//line /usr/local/go/src/html/template/transition.go:380
		_go_fuzz_dep_.CoverTab[31730]++
									lineTerminators, endState = "\n\r\u2028\u2029", stateJS
//line /usr/local/go/src/html/template/transition.go:381
		// _ = "end of CoverTab[31730]"
	case stateCSSLineCmt:
//line /usr/local/go/src/html/template/transition.go:382
		_go_fuzz_dep_.CoverTab[31731]++
									lineTerminators, endState = "\n\f\r", stateCSS
//line /usr/local/go/src/html/template/transition.go:383
		// _ = "end of CoverTab[31731]"

//line /usr/local/go/src/html/template/transition.go:391
	default:
//line /usr/local/go/src/html/template/transition.go:391
		_go_fuzz_dep_.CoverTab[31732]++
									panic(c.state.String())
//line /usr/local/go/src/html/template/transition.go:392
		// _ = "end of CoverTab[31732]"
	}
//line /usr/local/go/src/html/template/transition.go:393
	// _ = "end of CoverTab[31727]"
//line /usr/local/go/src/html/template/transition.go:393
	_go_fuzz_dep_.CoverTab[31728]++

								i := bytes.IndexAny(s, lineTerminators)
								if i == -1 {
//line /usr/local/go/src/html/template/transition.go:396
		_go_fuzz_dep_.CoverTab[31733]++
									return c, len(s)
//line /usr/local/go/src/html/template/transition.go:397
		// _ = "end of CoverTab[31733]"
	} else {
//line /usr/local/go/src/html/template/transition.go:398
		_go_fuzz_dep_.CoverTab[31734]++
//line /usr/local/go/src/html/template/transition.go:398
		// _ = "end of CoverTab[31734]"
//line /usr/local/go/src/html/template/transition.go:398
	}
//line /usr/local/go/src/html/template/transition.go:398
	// _ = "end of CoverTab[31728]"
//line /usr/local/go/src/html/template/transition.go:398
	_go_fuzz_dep_.CoverTab[31729]++
								c.state = endState

//line /usr/local/go/src/html/template/transition.go:405
	return c, i
//line /usr/local/go/src/html/template/transition.go:405
	// _ = "end of CoverTab[31729]"
}

// tCSS is the context transition function for the CSS state.
func tCSS(c context, s []byte) (context, int) {
//line /usr/local/go/src/html/template/transition.go:409
	_go_fuzz_dep_.CoverTab[31735]++

//line /usr/local/go/src/html/template/transition.go:437
	k := 0
	for {
//line /usr/local/go/src/html/template/transition.go:438
		_go_fuzz_dep_.CoverTab[31736]++
									i := k + bytes.IndexAny(s[k:], `("'/`)
									if i < k {
//line /usr/local/go/src/html/template/transition.go:440
			_go_fuzz_dep_.CoverTab[31739]++
										return c, len(s)
//line /usr/local/go/src/html/template/transition.go:441
			// _ = "end of CoverTab[31739]"
		} else {
//line /usr/local/go/src/html/template/transition.go:442
			_go_fuzz_dep_.CoverTab[31740]++
//line /usr/local/go/src/html/template/transition.go:442
			// _ = "end of CoverTab[31740]"
//line /usr/local/go/src/html/template/transition.go:442
		}
//line /usr/local/go/src/html/template/transition.go:442
		// _ = "end of CoverTab[31736]"
//line /usr/local/go/src/html/template/transition.go:442
		_go_fuzz_dep_.CoverTab[31737]++
									switch s[i] {
		case '(':
//line /usr/local/go/src/html/template/transition.go:444
			_go_fuzz_dep_.CoverTab[31741]++

										p := bytes.TrimRight(s[:i], "\t\n\f\r ")
										if endsWithCSSKeyword(p, "url") {
//line /usr/local/go/src/html/template/transition.go:447
				_go_fuzz_dep_.CoverTab[31746]++
											j := len(s) - len(bytes.TrimLeft(s[i+1:], "\t\n\f\r "))
											switch {
				case j != len(s) && func() bool {
//line /usr/local/go/src/html/template/transition.go:450
					_go_fuzz_dep_.CoverTab[31751]++
//line /usr/local/go/src/html/template/transition.go:450
					return s[j] == '"'
//line /usr/local/go/src/html/template/transition.go:450
					// _ = "end of CoverTab[31751]"
//line /usr/local/go/src/html/template/transition.go:450
				}():
//line /usr/local/go/src/html/template/transition.go:450
					_go_fuzz_dep_.CoverTab[31748]++
												c.state, j = stateCSSDqURL, j+1
//line /usr/local/go/src/html/template/transition.go:451
					// _ = "end of CoverTab[31748]"
				case j != len(s) && func() bool {
//line /usr/local/go/src/html/template/transition.go:452
					_go_fuzz_dep_.CoverTab[31752]++
//line /usr/local/go/src/html/template/transition.go:452
					return s[j] == '\''
//line /usr/local/go/src/html/template/transition.go:452
					// _ = "end of CoverTab[31752]"
//line /usr/local/go/src/html/template/transition.go:452
				}():
//line /usr/local/go/src/html/template/transition.go:452
					_go_fuzz_dep_.CoverTab[31749]++
												c.state, j = stateCSSSqURL, j+1
//line /usr/local/go/src/html/template/transition.go:453
					// _ = "end of CoverTab[31749]"
				default:
//line /usr/local/go/src/html/template/transition.go:454
					_go_fuzz_dep_.CoverTab[31750]++
												c.state = stateCSSURL
//line /usr/local/go/src/html/template/transition.go:455
					// _ = "end of CoverTab[31750]"
				}
//line /usr/local/go/src/html/template/transition.go:456
				// _ = "end of CoverTab[31746]"
//line /usr/local/go/src/html/template/transition.go:456
				_go_fuzz_dep_.CoverTab[31747]++
											return c, j
//line /usr/local/go/src/html/template/transition.go:457
				// _ = "end of CoverTab[31747]"
			} else {
//line /usr/local/go/src/html/template/transition.go:458
				_go_fuzz_dep_.CoverTab[31753]++
//line /usr/local/go/src/html/template/transition.go:458
				// _ = "end of CoverTab[31753]"
//line /usr/local/go/src/html/template/transition.go:458
			}
//line /usr/local/go/src/html/template/transition.go:458
			// _ = "end of CoverTab[31741]"
		case '/':
//line /usr/local/go/src/html/template/transition.go:459
			_go_fuzz_dep_.CoverTab[31742]++
										if i+1 < len(s) {
//line /usr/local/go/src/html/template/transition.go:460
				_go_fuzz_dep_.CoverTab[31754]++
											switch s[i+1] {
				case '/':
//line /usr/local/go/src/html/template/transition.go:462
					_go_fuzz_dep_.CoverTab[31755]++
												c.state = stateCSSLineCmt
												return c, i + 2
//line /usr/local/go/src/html/template/transition.go:464
					// _ = "end of CoverTab[31755]"
				case '*':
//line /usr/local/go/src/html/template/transition.go:465
					_go_fuzz_dep_.CoverTab[31756]++
												c.state = stateCSSBlockCmt
												return c, i + 2
//line /usr/local/go/src/html/template/transition.go:467
					// _ = "end of CoverTab[31756]"
//line /usr/local/go/src/html/template/transition.go:467
				default:
//line /usr/local/go/src/html/template/transition.go:467
					_go_fuzz_dep_.CoverTab[31757]++
//line /usr/local/go/src/html/template/transition.go:467
					// _ = "end of CoverTab[31757]"
				}
//line /usr/local/go/src/html/template/transition.go:468
				// _ = "end of CoverTab[31754]"
			} else {
//line /usr/local/go/src/html/template/transition.go:469
				_go_fuzz_dep_.CoverTab[31758]++
//line /usr/local/go/src/html/template/transition.go:469
				// _ = "end of CoverTab[31758]"
//line /usr/local/go/src/html/template/transition.go:469
			}
//line /usr/local/go/src/html/template/transition.go:469
			// _ = "end of CoverTab[31742]"
		case '"':
//line /usr/local/go/src/html/template/transition.go:470
			_go_fuzz_dep_.CoverTab[31743]++
										c.state = stateCSSDqStr
										return c, i + 1
//line /usr/local/go/src/html/template/transition.go:472
			// _ = "end of CoverTab[31743]"
		case '\'':
//line /usr/local/go/src/html/template/transition.go:473
			_go_fuzz_dep_.CoverTab[31744]++
										c.state = stateCSSSqStr
										return c, i + 1
//line /usr/local/go/src/html/template/transition.go:475
			// _ = "end of CoverTab[31744]"
//line /usr/local/go/src/html/template/transition.go:475
		default:
//line /usr/local/go/src/html/template/transition.go:475
			_go_fuzz_dep_.CoverTab[31745]++
//line /usr/local/go/src/html/template/transition.go:475
			// _ = "end of CoverTab[31745]"
		}
//line /usr/local/go/src/html/template/transition.go:476
		// _ = "end of CoverTab[31737]"
//line /usr/local/go/src/html/template/transition.go:476
		_go_fuzz_dep_.CoverTab[31738]++
									k = i + 1
//line /usr/local/go/src/html/template/transition.go:477
		// _ = "end of CoverTab[31738]"
	}
//line /usr/local/go/src/html/template/transition.go:478
	// _ = "end of CoverTab[31735]"
}

// tCSSStr is the context transition function for the CSS string and URL states.
func tCSSStr(c context, s []byte) (context, int) {
//line /usr/local/go/src/html/template/transition.go:482
	_go_fuzz_dep_.CoverTab[31759]++
								var endAndEsc string
								switch c.state {
	case stateCSSDqStr, stateCSSDqURL:
//line /usr/local/go/src/html/template/transition.go:485
		_go_fuzz_dep_.CoverTab[31761]++
									endAndEsc = `\"`
//line /usr/local/go/src/html/template/transition.go:486
		// _ = "end of CoverTab[31761]"
	case stateCSSSqStr, stateCSSSqURL:
//line /usr/local/go/src/html/template/transition.go:487
		_go_fuzz_dep_.CoverTab[31762]++
									endAndEsc = `\'`
//line /usr/local/go/src/html/template/transition.go:488
		// _ = "end of CoverTab[31762]"
	case stateCSSURL:
//line /usr/local/go/src/html/template/transition.go:489
		_go_fuzz_dep_.CoverTab[31763]++

//line /usr/local/go/src/html/template/transition.go:492
		endAndEsc = "\\\t\n\f\r )"
//line /usr/local/go/src/html/template/transition.go:492
		// _ = "end of CoverTab[31763]"
	default:
//line /usr/local/go/src/html/template/transition.go:493
		_go_fuzz_dep_.CoverTab[31764]++
									panic(c.state.String())
//line /usr/local/go/src/html/template/transition.go:494
		// _ = "end of CoverTab[31764]"
	}
//line /usr/local/go/src/html/template/transition.go:495
	// _ = "end of CoverTab[31759]"
//line /usr/local/go/src/html/template/transition.go:495
	_go_fuzz_dep_.CoverTab[31760]++

								k := 0
								for {
//line /usr/local/go/src/html/template/transition.go:498
		_go_fuzz_dep_.CoverTab[31765]++
									i := k + bytes.IndexAny(s[k:], endAndEsc)
									if i < k {
//line /usr/local/go/src/html/template/transition.go:500
			_go_fuzz_dep_.CoverTab[31768]++
										c, nread := tURL(c, decodeCSS(s[k:]))
										return c, k + nread
//line /usr/local/go/src/html/template/transition.go:502
			// _ = "end of CoverTab[31768]"
		} else {
//line /usr/local/go/src/html/template/transition.go:503
			_go_fuzz_dep_.CoverTab[31769]++
//line /usr/local/go/src/html/template/transition.go:503
			// _ = "end of CoverTab[31769]"
//line /usr/local/go/src/html/template/transition.go:503
		}
//line /usr/local/go/src/html/template/transition.go:503
		// _ = "end of CoverTab[31765]"
//line /usr/local/go/src/html/template/transition.go:503
		_go_fuzz_dep_.CoverTab[31766]++
									if s[i] == '\\' {
//line /usr/local/go/src/html/template/transition.go:504
			_go_fuzz_dep_.CoverTab[31770]++
										i++
										if i == len(s) {
//line /usr/local/go/src/html/template/transition.go:506
				_go_fuzz_dep_.CoverTab[31771]++
											return context{
					state:	stateError,
					err:	errorf(ErrPartialEscape, nil, 0, "unfinished escape sequence in CSS string: %q", s),
				}, len(s)
//line /usr/local/go/src/html/template/transition.go:510
				// _ = "end of CoverTab[31771]"
			} else {
//line /usr/local/go/src/html/template/transition.go:511
				_go_fuzz_dep_.CoverTab[31772]++
//line /usr/local/go/src/html/template/transition.go:511
				// _ = "end of CoverTab[31772]"
//line /usr/local/go/src/html/template/transition.go:511
			}
//line /usr/local/go/src/html/template/transition.go:511
			// _ = "end of CoverTab[31770]"
		} else {
//line /usr/local/go/src/html/template/transition.go:512
			_go_fuzz_dep_.CoverTab[31773]++
										c.state = stateCSS
										return c, i + 1
//line /usr/local/go/src/html/template/transition.go:514
			// _ = "end of CoverTab[31773]"
		}
//line /usr/local/go/src/html/template/transition.go:515
		// _ = "end of CoverTab[31766]"
//line /usr/local/go/src/html/template/transition.go:515
		_go_fuzz_dep_.CoverTab[31767]++
									c, _ = tURL(c, decodeCSS(s[:i+1]))
									k = i + 1
//line /usr/local/go/src/html/template/transition.go:517
		// _ = "end of CoverTab[31767]"
	}
//line /usr/local/go/src/html/template/transition.go:518
	// _ = "end of CoverTab[31760]"
}

// tError is the context transition function for the error state.
func tError(c context, s []byte) (context, int) {
//line /usr/local/go/src/html/template/transition.go:522
	_go_fuzz_dep_.CoverTab[31774]++
								return c, len(s)
//line /usr/local/go/src/html/template/transition.go:523
	// _ = "end of CoverTab[31774]"
}

// eatAttrName returns the largest j such that s[i:j] is an attribute name.
//line /usr/local/go/src/html/template/transition.go:526
// It returns an error if s[i:] does not look like it begins with an
//line /usr/local/go/src/html/template/transition.go:526
// attribute name, such as encountering a quote mark without a preceding
//line /usr/local/go/src/html/template/transition.go:526
// equals sign.
//line /usr/local/go/src/html/template/transition.go:530
func eatAttrName(s []byte, i int) (int, *Error) {
//line /usr/local/go/src/html/template/transition.go:530
	_go_fuzz_dep_.CoverTab[31775]++
								for j := i; j < len(s); j++ {
//line /usr/local/go/src/html/template/transition.go:531
		_go_fuzz_dep_.CoverTab[31777]++
									switch s[j] {
		case ' ', '\t', '\n', '\f', '\r', '=', '>':
//line /usr/local/go/src/html/template/transition.go:533
			_go_fuzz_dep_.CoverTab[31778]++
										return j, nil
//line /usr/local/go/src/html/template/transition.go:534
			// _ = "end of CoverTab[31778]"
		case '\'', '"', '<':
//line /usr/local/go/src/html/template/transition.go:535
			_go_fuzz_dep_.CoverTab[31779]++

//line /usr/local/go/src/html/template/transition.go:539
			return -1, errorf(ErrBadHTML, nil, 0, "%q in attribute name: %.32q", s[j:j+1], s)
//line /usr/local/go/src/html/template/transition.go:539
			// _ = "end of CoverTab[31779]"
		default:
//line /usr/local/go/src/html/template/transition.go:540
			_go_fuzz_dep_.CoverTab[31780]++
//line /usr/local/go/src/html/template/transition.go:540
			// _ = "end of CoverTab[31780]"

		}
//line /usr/local/go/src/html/template/transition.go:542
		// _ = "end of CoverTab[31777]"
	}
//line /usr/local/go/src/html/template/transition.go:543
	// _ = "end of CoverTab[31775]"
//line /usr/local/go/src/html/template/transition.go:543
	_go_fuzz_dep_.CoverTab[31776]++
								return len(s), nil
//line /usr/local/go/src/html/template/transition.go:544
	// _ = "end of CoverTab[31776]"
}

var elementNameMap = map[string]element{
	"script":	elementScript,
	"style":	elementStyle,
	"textarea":	elementTextarea,
	"title":	elementTitle,
}

// asciiAlpha reports whether c is an ASCII letter.
func asciiAlpha(c byte) bool {
//line /usr/local/go/src/html/template/transition.go:555
	_go_fuzz_dep_.CoverTab[31781]++
								return 'A' <= c && func() bool {
//line /usr/local/go/src/html/template/transition.go:556
		_go_fuzz_dep_.CoverTab[31782]++
//line /usr/local/go/src/html/template/transition.go:556
		return c <= 'Z'
//line /usr/local/go/src/html/template/transition.go:556
		// _ = "end of CoverTab[31782]"
//line /usr/local/go/src/html/template/transition.go:556
	}() || func() bool {
//line /usr/local/go/src/html/template/transition.go:556
		_go_fuzz_dep_.CoverTab[31783]++
//line /usr/local/go/src/html/template/transition.go:556
		return 'a' <= c && func() bool {
//line /usr/local/go/src/html/template/transition.go:556
			_go_fuzz_dep_.CoverTab[31784]++
//line /usr/local/go/src/html/template/transition.go:556
			return c <= 'z'
//line /usr/local/go/src/html/template/transition.go:556
			// _ = "end of CoverTab[31784]"
//line /usr/local/go/src/html/template/transition.go:556
		}()
//line /usr/local/go/src/html/template/transition.go:556
		// _ = "end of CoverTab[31783]"
//line /usr/local/go/src/html/template/transition.go:556
	}()
//line /usr/local/go/src/html/template/transition.go:556
	// _ = "end of CoverTab[31781]"
}

// asciiAlphaNum reports whether c is an ASCII letter or digit.
func asciiAlphaNum(c byte) bool {
//line /usr/local/go/src/html/template/transition.go:560
	_go_fuzz_dep_.CoverTab[31785]++
								return asciiAlpha(c) || func() bool {
//line /usr/local/go/src/html/template/transition.go:561
		_go_fuzz_dep_.CoverTab[31786]++
//line /usr/local/go/src/html/template/transition.go:561
		return '0' <= c && func() bool {
//line /usr/local/go/src/html/template/transition.go:561
			_go_fuzz_dep_.CoverTab[31787]++
//line /usr/local/go/src/html/template/transition.go:561
			return c <= '9'
//line /usr/local/go/src/html/template/transition.go:561
			// _ = "end of CoverTab[31787]"
//line /usr/local/go/src/html/template/transition.go:561
		}()
//line /usr/local/go/src/html/template/transition.go:561
		// _ = "end of CoverTab[31786]"
//line /usr/local/go/src/html/template/transition.go:561
	}()
//line /usr/local/go/src/html/template/transition.go:561
	// _ = "end of CoverTab[31785]"
}

// eatTagName returns the largest j such that s[i:j] is a tag name and the tag type.
func eatTagName(s []byte, i int) (int, element) {
//line /usr/local/go/src/html/template/transition.go:565
	_go_fuzz_dep_.CoverTab[31788]++
								if i == len(s) || func() bool {
//line /usr/local/go/src/html/template/transition.go:566
		_go_fuzz_dep_.CoverTab[31791]++
//line /usr/local/go/src/html/template/transition.go:566
		return !asciiAlpha(s[i])
//line /usr/local/go/src/html/template/transition.go:566
		// _ = "end of CoverTab[31791]"
//line /usr/local/go/src/html/template/transition.go:566
	}() {
//line /usr/local/go/src/html/template/transition.go:566
		_go_fuzz_dep_.CoverTab[31792]++
									return i, elementNone
//line /usr/local/go/src/html/template/transition.go:567
		// _ = "end of CoverTab[31792]"
	} else {
//line /usr/local/go/src/html/template/transition.go:568
		_go_fuzz_dep_.CoverTab[31793]++
//line /usr/local/go/src/html/template/transition.go:568
		// _ = "end of CoverTab[31793]"
//line /usr/local/go/src/html/template/transition.go:568
	}
//line /usr/local/go/src/html/template/transition.go:568
	// _ = "end of CoverTab[31788]"
//line /usr/local/go/src/html/template/transition.go:568
	_go_fuzz_dep_.CoverTab[31789]++
								j := i + 1
								for j < len(s) {
//line /usr/local/go/src/html/template/transition.go:570
		_go_fuzz_dep_.CoverTab[31794]++
									x := s[j]
									if asciiAlphaNum(x) {
//line /usr/local/go/src/html/template/transition.go:572
			_go_fuzz_dep_.CoverTab[31797]++
										j++
										continue
//line /usr/local/go/src/html/template/transition.go:574
			// _ = "end of CoverTab[31797]"
		} else {
//line /usr/local/go/src/html/template/transition.go:575
			_go_fuzz_dep_.CoverTab[31798]++
//line /usr/local/go/src/html/template/transition.go:575
			// _ = "end of CoverTab[31798]"
//line /usr/local/go/src/html/template/transition.go:575
		}
//line /usr/local/go/src/html/template/transition.go:575
		// _ = "end of CoverTab[31794]"
//line /usr/local/go/src/html/template/transition.go:575
		_go_fuzz_dep_.CoverTab[31795]++

									if (x == ':' || func() bool {
//line /usr/local/go/src/html/template/transition.go:577
			_go_fuzz_dep_.CoverTab[31799]++
//line /usr/local/go/src/html/template/transition.go:577
			return x == '-'
//line /usr/local/go/src/html/template/transition.go:577
			// _ = "end of CoverTab[31799]"
//line /usr/local/go/src/html/template/transition.go:577
		}()) && func() bool {
//line /usr/local/go/src/html/template/transition.go:577
			_go_fuzz_dep_.CoverTab[31800]++
//line /usr/local/go/src/html/template/transition.go:577
			return j+1 < len(s)
//line /usr/local/go/src/html/template/transition.go:577
			// _ = "end of CoverTab[31800]"
//line /usr/local/go/src/html/template/transition.go:577
		}() && func() bool {
//line /usr/local/go/src/html/template/transition.go:577
			_go_fuzz_dep_.CoverTab[31801]++
//line /usr/local/go/src/html/template/transition.go:577
			return asciiAlphaNum(s[j+1])
//line /usr/local/go/src/html/template/transition.go:577
			// _ = "end of CoverTab[31801]"
//line /usr/local/go/src/html/template/transition.go:577
		}() {
//line /usr/local/go/src/html/template/transition.go:577
			_go_fuzz_dep_.CoverTab[31802]++
										j += 2
										continue
//line /usr/local/go/src/html/template/transition.go:579
			// _ = "end of CoverTab[31802]"
		} else {
//line /usr/local/go/src/html/template/transition.go:580
			_go_fuzz_dep_.CoverTab[31803]++
//line /usr/local/go/src/html/template/transition.go:580
			// _ = "end of CoverTab[31803]"
//line /usr/local/go/src/html/template/transition.go:580
		}
//line /usr/local/go/src/html/template/transition.go:580
		// _ = "end of CoverTab[31795]"
//line /usr/local/go/src/html/template/transition.go:580
		_go_fuzz_dep_.CoverTab[31796]++
									break
//line /usr/local/go/src/html/template/transition.go:581
		// _ = "end of CoverTab[31796]"
	}
//line /usr/local/go/src/html/template/transition.go:582
	// _ = "end of CoverTab[31789]"
//line /usr/local/go/src/html/template/transition.go:582
	_go_fuzz_dep_.CoverTab[31790]++
								return j, elementNameMap[strings.ToLower(string(s[i:j]))]
//line /usr/local/go/src/html/template/transition.go:583
	// _ = "end of CoverTab[31790]"
}

// eatWhiteSpace returns the largest j such that s[i:j] is white space.
func eatWhiteSpace(s []byte, i int) int {
//line /usr/local/go/src/html/template/transition.go:587
	_go_fuzz_dep_.CoverTab[31804]++
								for j := i; j < len(s); j++ {
//line /usr/local/go/src/html/template/transition.go:588
		_go_fuzz_dep_.CoverTab[31806]++
									switch s[j] {
		case ' ', '\t', '\n', '\f', '\r':
//line /usr/local/go/src/html/template/transition.go:590
			_go_fuzz_dep_.CoverTab[31807]++
//line /usr/local/go/src/html/template/transition.go:590
			// _ = "end of CoverTab[31807]"

		default:
//line /usr/local/go/src/html/template/transition.go:592
			_go_fuzz_dep_.CoverTab[31808]++
										return j
//line /usr/local/go/src/html/template/transition.go:593
			// _ = "end of CoverTab[31808]"
		}
//line /usr/local/go/src/html/template/transition.go:594
		// _ = "end of CoverTab[31806]"
	}
//line /usr/local/go/src/html/template/transition.go:595
	// _ = "end of CoverTab[31804]"
//line /usr/local/go/src/html/template/transition.go:595
	_go_fuzz_dep_.CoverTab[31805]++
								return len(s)
//line /usr/local/go/src/html/template/transition.go:596
	// _ = "end of CoverTab[31805]"
}

//line /usr/local/go/src/html/template/transition.go:597
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/html/template/transition.go:597
var _ = _go_fuzz_dep_.CoverTab
