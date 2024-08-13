// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/html/template/context.go:5
package template

//line /usr/local/go/src/html/template/context.go:5
import (
//line /usr/local/go/src/html/template/context.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/html/template/context.go:5
)
//line /usr/local/go/src/html/template/context.go:5
import (
//line /usr/local/go/src/html/template/context.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/html/template/context.go:5
)

import (
	"fmt"
	"text/template/parse"
)

//line /usr/local/go/src/html/template/context.go:19
type context struct {
	state	state
	delim	delim
	urlPart	urlPart
	jsCtx	jsCtx
	attr	attr
	element	element
	n	parse.Node
	err	*Error
}

func (c context) String() string {
//line /usr/local/go/src/html/template/context.go:30
	_go_fuzz_dep_.CoverTab[30694]++
							var err error
							if c.err != nil {
//line /usr/local/go/src/html/template/context.go:32
		_go_fuzz_dep_.CoverTab[30696]++
								err = c.err
//line /usr/local/go/src/html/template/context.go:33
		// _ = "end of CoverTab[30696]"
	} else {
//line /usr/local/go/src/html/template/context.go:34
		_go_fuzz_dep_.CoverTab[30697]++
//line /usr/local/go/src/html/template/context.go:34
		// _ = "end of CoverTab[30697]"
//line /usr/local/go/src/html/template/context.go:34
	}
//line /usr/local/go/src/html/template/context.go:34
	// _ = "end of CoverTab[30694]"
//line /usr/local/go/src/html/template/context.go:34
	_go_fuzz_dep_.CoverTab[30695]++
							return fmt.Sprintf("{%v %v %v %v %v %v %v}", c.state, c.delim, c.urlPart, c.jsCtx, c.attr, c.element, err)
//line /usr/local/go/src/html/template/context.go:35
	// _ = "end of CoverTab[30695]"
}

//line /usr/local/go/src/html/template/context.go:39
func (c context) eq(d context) bool {
//line /usr/local/go/src/html/template/context.go:39
	_go_fuzz_dep_.CoverTab[30698]++
							return c.state == d.state && func() bool {
//line /usr/local/go/src/html/template/context.go:40
		_go_fuzz_dep_.CoverTab[30699]++
//line /usr/local/go/src/html/template/context.go:40
		return c.delim == d.delim
								// _ = "end of CoverTab[30699]"
//line /usr/local/go/src/html/template/context.go:41
	}() && func() bool {
//line /usr/local/go/src/html/template/context.go:41
		_go_fuzz_dep_.CoverTab[30700]++
//line /usr/local/go/src/html/template/context.go:41
		return c.urlPart == d.urlPart
								// _ = "end of CoverTab[30700]"
//line /usr/local/go/src/html/template/context.go:42
	}() && func() bool {
//line /usr/local/go/src/html/template/context.go:42
		_go_fuzz_dep_.CoverTab[30701]++
//line /usr/local/go/src/html/template/context.go:42
		return c.jsCtx == d.jsCtx
								// _ = "end of CoverTab[30701]"
//line /usr/local/go/src/html/template/context.go:43
	}() && func() bool {
//line /usr/local/go/src/html/template/context.go:43
		_go_fuzz_dep_.CoverTab[30702]++
//line /usr/local/go/src/html/template/context.go:43
		return c.attr == d.attr
								// _ = "end of CoverTab[30702]"
//line /usr/local/go/src/html/template/context.go:44
	}() && func() bool {
//line /usr/local/go/src/html/template/context.go:44
		_go_fuzz_dep_.CoverTab[30703]++
//line /usr/local/go/src/html/template/context.go:44
		return c.element == d.element
								// _ = "end of CoverTab[30703]"
//line /usr/local/go/src/html/template/context.go:45
	}() && func() bool {
//line /usr/local/go/src/html/template/context.go:45
		_go_fuzz_dep_.CoverTab[30704]++
//line /usr/local/go/src/html/template/context.go:45
		return c.err == d.err
								// _ = "end of CoverTab[30704]"
//line /usr/local/go/src/html/template/context.go:46
	}()
//line /usr/local/go/src/html/template/context.go:46
	// _ = "end of CoverTab[30698]"
}

//line /usr/local/go/src/html/template/context.go:51
func (c context) mangle(templateName string) string {
//line /usr/local/go/src/html/template/context.go:51
	_go_fuzz_dep_.CoverTab[30705]++

							if c.state == stateText {
//line /usr/local/go/src/html/template/context.go:53
		_go_fuzz_dep_.CoverTab[30712]++
								return templateName
//line /usr/local/go/src/html/template/context.go:54
		// _ = "end of CoverTab[30712]"
	} else {
//line /usr/local/go/src/html/template/context.go:55
		_go_fuzz_dep_.CoverTab[30713]++
//line /usr/local/go/src/html/template/context.go:55
		// _ = "end of CoverTab[30713]"
//line /usr/local/go/src/html/template/context.go:55
	}
//line /usr/local/go/src/html/template/context.go:55
	// _ = "end of CoverTab[30705]"
//line /usr/local/go/src/html/template/context.go:55
	_go_fuzz_dep_.CoverTab[30706]++
							s := templateName + "$htmltemplate_" + c.state.String()
							if c.delim != delimNone {
//line /usr/local/go/src/html/template/context.go:57
		_go_fuzz_dep_.CoverTab[30714]++
								s += "_" + c.delim.String()
//line /usr/local/go/src/html/template/context.go:58
		// _ = "end of CoverTab[30714]"
	} else {
//line /usr/local/go/src/html/template/context.go:59
		_go_fuzz_dep_.CoverTab[30715]++
//line /usr/local/go/src/html/template/context.go:59
		// _ = "end of CoverTab[30715]"
//line /usr/local/go/src/html/template/context.go:59
	}
//line /usr/local/go/src/html/template/context.go:59
	// _ = "end of CoverTab[30706]"
//line /usr/local/go/src/html/template/context.go:59
	_go_fuzz_dep_.CoverTab[30707]++
							if c.urlPart != urlPartNone {
//line /usr/local/go/src/html/template/context.go:60
		_go_fuzz_dep_.CoverTab[30716]++
								s += "_" + c.urlPart.String()
//line /usr/local/go/src/html/template/context.go:61
		// _ = "end of CoverTab[30716]"
	} else {
//line /usr/local/go/src/html/template/context.go:62
		_go_fuzz_dep_.CoverTab[30717]++
//line /usr/local/go/src/html/template/context.go:62
		// _ = "end of CoverTab[30717]"
//line /usr/local/go/src/html/template/context.go:62
	}
//line /usr/local/go/src/html/template/context.go:62
	// _ = "end of CoverTab[30707]"
//line /usr/local/go/src/html/template/context.go:62
	_go_fuzz_dep_.CoverTab[30708]++
							if c.jsCtx != jsCtxRegexp {
//line /usr/local/go/src/html/template/context.go:63
		_go_fuzz_dep_.CoverTab[30718]++
								s += "_" + c.jsCtx.String()
//line /usr/local/go/src/html/template/context.go:64
		// _ = "end of CoverTab[30718]"
	} else {
//line /usr/local/go/src/html/template/context.go:65
		_go_fuzz_dep_.CoverTab[30719]++
//line /usr/local/go/src/html/template/context.go:65
		// _ = "end of CoverTab[30719]"
//line /usr/local/go/src/html/template/context.go:65
	}
//line /usr/local/go/src/html/template/context.go:65
	// _ = "end of CoverTab[30708]"
//line /usr/local/go/src/html/template/context.go:65
	_go_fuzz_dep_.CoverTab[30709]++
							if c.attr != attrNone {
//line /usr/local/go/src/html/template/context.go:66
		_go_fuzz_dep_.CoverTab[30720]++
								s += "_" + c.attr.String()
//line /usr/local/go/src/html/template/context.go:67
		// _ = "end of CoverTab[30720]"
	} else {
//line /usr/local/go/src/html/template/context.go:68
		_go_fuzz_dep_.CoverTab[30721]++
//line /usr/local/go/src/html/template/context.go:68
		// _ = "end of CoverTab[30721]"
//line /usr/local/go/src/html/template/context.go:68
	}
//line /usr/local/go/src/html/template/context.go:68
	// _ = "end of CoverTab[30709]"
//line /usr/local/go/src/html/template/context.go:68
	_go_fuzz_dep_.CoverTab[30710]++
							if c.element != elementNone {
//line /usr/local/go/src/html/template/context.go:69
		_go_fuzz_dep_.CoverTab[30722]++
								s += "_" + c.element.String()
//line /usr/local/go/src/html/template/context.go:70
		// _ = "end of CoverTab[30722]"
	} else {
//line /usr/local/go/src/html/template/context.go:71
		_go_fuzz_dep_.CoverTab[30723]++
//line /usr/local/go/src/html/template/context.go:71
		// _ = "end of CoverTab[30723]"
//line /usr/local/go/src/html/template/context.go:71
	}
//line /usr/local/go/src/html/template/context.go:71
	// _ = "end of CoverTab[30710]"
//line /usr/local/go/src/html/template/context.go:71
	_go_fuzz_dep_.CoverTab[30711]++
							return s
//line /usr/local/go/src/html/template/context.go:72
	// _ = "end of CoverTab[30711]"
}

//line /usr/local/go/src/html/template/context.go:86
type state uint8

//go:generate stringer -type state

const (
//line /usr/local/go/src/html/template/context.go:94
	stateText	state	= iota

							stateTag

//line /usr/local/go/src/html/template/context.go:99
	stateAttrName

//line /usr/local/go/src/html/template/context.go:102
	stateAfterName

//line /usr/local/go/src/html/template/context.go:105
	stateBeforeValue

							stateHTMLCmt

//line /usr/local/go/src/html/template/context.go:110
	stateRCDATA

							stateAttr

							stateURL

							stateSrcset

							stateJS

							stateJSDqStr

							stateJSSqStr

							stateJSBqStr

							stateJSRegexp

							stateJSBlockCmt

							stateJSLineCmt

							stateCSS

							stateCSSDqStr

							stateCSSSqStr

							stateCSSDqURL

							stateCSSSqURL

							stateCSSURL

							stateCSSBlockCmt

							stateCSSLineCmt

//line /usr/local/go/src/html/template/context.go:149
	stateError

	stateDead
)

//line /usr/local/go/src/html/template/context.go:156
func isComment(s state) bool {
//line /usr/local/go/src/html/template/context.go:156
	_go_fuzz_dep_.CoverTab[30724]++
							switch s {
	case stateHTMLCmt, stateJSBlockCmt, stateJSLineCmt, stateCSSBlockCmt, stateCSSLineCmt:
//line /usr/local/go/src/html/template/context.go:158
		_go_fuzz_dep_.CoverTab[30726]++
								return true
//line /usr/local/go/src/html/template/context.go:159
		// _ = "end of CoverTab[30726]"
//line /usr/local/go/src/html/template/context.go:159
	default:
//line /usr/local/go/src/html/template/context.go:159
		_go_fuzz_dep_.CoverTab[30727]++
//line /usr/local/go/src/html/template/context.go:159
		// _ = "end of CoverTab[30727]"
	}
//line /usr/local/go/src/html/template/context.go:160
	// _ = "end of CoverTab[30724]"
//line /usr/local/go/src/html/template/context.go:160
	_go_fuzz_dep_.CoverTab[30725]++
							return false
//line /usr/local/go/src/html/template/context.go:161
	// _ = "end of CoverTab[30725]"
}

//line /usr/local/go/src/html/template/context.go:165
func isInTag(s state) bool {
//line /usr/local/go/src/html/template/context.go:165
	_go_fuzz_dep_.CoverTab[30728]++
							switch s {
	case stateTag, stateAttrName, stateAfterName, stateBeforeValue, stateAttr:
//line /usr/local/go/src/html/template/context.go:167
		_go_fuzz_dep_.CoverTab[30730]++
								return true
//line /usr/local/go/src/html/template/context.go:168
		// _ = "end of CoverTab[30730]"
//line /usr/local/go/src/html/template/context.go:168
	default:
//line /usr/local/go/src/html/template/context.go:168
		_go_fuzz_dep_.CoverTab[30731]++
//line /usr/local/go/src/html/template/context.go:168
		// _ = "end of CoverTab[30731]"
	}
//line /usr/local/go/src/html/template/context.go:169
	// _ = "end of CoverTab[30728]"
//line /usr/local/go/src/html/template/context.go:169
	_go_fuzz_dep_.CoverTab[30729]++
							return false
//line /usr/local/go/src/html/template/context.go:170
	// _ = "end of CoverTab[30729]"
}

//line /usr/local/go/src/html/template/context.go:174
type delim uint8

//go:generate stringer -type delim

const (
//line /usr/local/go/src/html/template/context.go:180
	delimNone	delim	= iota

							delimDoubleQuote

							delimSingleQuote

//line /usr/local/go/src/html/template/context.go:187
	delimSpaceOrTagEnd
)

//line /usr/local/go/src/html/template/context.go:192
type urlPart uint8

//go:generate stringer -type urlPart

const (
//line /usr/local/go/src/html/template/context.go:199
	urlPartNone	urlPart	= iota

//line /usr/local/go/src/html/template/context.go:202
	urlPartPreQuery

//line /usr/local/go/src/html/template/context.go:205
	urlPartQueryOrFrag

//line /usr/local/go/src/html/template/context.go:208
	urlPartUnknown
)

//line /usr/local/go/src/html/template/context.go:213
type jsCtx uint8

//go:generate stringer -type jsCtx

const (
//line /usr/local/go/src/html/template/context.go:219
	jsCtxRegexp	jsCtx	= iota

	jsCtxDivOp

	jsCtxUnknown
)

//line /usr/local/go/src/html/template/context.go:231
type element uint8

//go:generate stringer -type element

const (
//line /usr/local/go/src/html/template/context.go:237
	elementNone	element	= iota

//line /usr/local/go/src/html/template/context.go:240
	elementScript

	elementStyle

	elementTextarea

	elementTitle
)

//go:generate stringer -type attr

//line /usr/local/go/src/html/template/context.go:253
type attr uint8

const (
//line /usr/local/go/src/html/template/context.go:257
	attrNone	attr	= iota

	attrScript

	attrScriptType

	attrStyle

	attrURL

	attrSrcset
)

//line /usr/local/go/src/html/template/context.go:268
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/html/template/context.go:268
var _ = _go_fuzz_dep_.CoverTab
