// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/html/template/content.go:5
package template

//line /usr/local/go/src/html/template/content.go:5
import (
//line /usr/local/go/src/html/template/content.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/html/template/content.go:5
)
//line /usr/local/go/src/html/template/content.go:5
import (
//line /usr/local/go/src/html/template/content.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/html/template/content.go:5
)

import (
	"fmt"
	"reflect"
)

// Strings of content from a trusted source.
type (
	// CSS encapsulates known safe content that matches any of:
	//   1. The CSS3 stylesheet production, such as `p { color: purple }`.
	//   2. The CSS3 rule production, such as `a[href=~"https:"].foo#bar`.
	//   3. CSS3 declaration productions, such as `color: red; margin: 2px`.
	//   4. The CSS3 value production, such as `rgba(0, 0, 255, 127)`.
	// See https://www.w3.org/TR/css3-syntax/#parsing and
	// https://web.archive.org/web/20090211114933/http://w3.org/TR/css3-syntax#style
	//
	// Use of this type presents a security risk:
	// the encapsulated content should come from a trusted source,
	// as it will be included verbatim in the template output.
	CSS	string

	// HTML encapsulates a known safe HTML document fragment.
	// It should not be used for HTML from a third-party, or HTML with
	// unclosed tags or comments. The outputs of a sound HTML sanitizer
	// and a template escaped by this package are fine for use with HTML.
	//
	// Use of this type presents a security risk:
	// the encapsulated content should come from a trusted source,
	// as it will be included verbatim in the template output.
	HTML	string

	// HTMLAttr encapsulates an HTML attribute from a trusted source,
	// for example, ` dir="ltr"`.
	//
	// Use of this type presents a security risk:
	// the encapsulated content should come from a trusted source,
	// as it will be included verbatim in the template output.
	HTMLAttr	string

	// JS encapsulates a known safe EcmaScript5 Expression, for example,
	// `(x + y * z())`.
	// Template authors are responsible for ensuring that typed expressions
	// do not break the intended precedence and that there is no
	// statement/expression ambiguity as when passing an expression like
	// "{ foo: bar() }\n['foo']()", which is both a valid Expression and a
	// valid Program with a very different meaning.
	//
	// Use of this type presents a security risk:
	// the encapsulated content should come from a trusted source,
	// as it will be included verbatim in the template output.
	//
	// Using JS to include valid but untrusted JSON is not safe.
	// A safe alternative is to parse the JSON with json.Unmarshal and then
	// pass the resultant object into the template, where it will be
	// converted to sanitized JSON when presented in a JavaScript context.
	JS	string

	// JSStr encapsulates a sequence of characters meant to be embedded
	// between quotes in a JavaScript expression.
	// The string must match a series of StringCharacters:
	//   StringCharacter :: SourceCharacter but not `\` or LineTerminator
	//                    | EscapeSequence
	// Note that LineContinuations are not allowed.
	// JSStr("foo\\nbar") is fine, but JSStr("foo\\\nbar") is not.
	//
	// Use of this type presents a security risk:
	// the encapsulated content should come from a trusted source,
	// as it will be included verbatim in the template output.
	JSStr	string

	// URL encapsulates a known safe URL or URL substring (see RFC 3986).
	// A URL like `javascript:checkThatFormNotEditedBeforeLeavingPage()`
	// from a trusted source should go in the page, but by default dynamic
	// `javascript:` URLs are filtered out since they are a frequently
	// exploited injection vector.
	//
	// Use of this type presents a security risk:
	// the encapsulated content should come from a trusted source,
	// as it will be included verbatim in the template output.
	URL	string

	// Srcset encapsulates a known safe srcset attribute
	// (see https://w3c.github.io/html/semantics-embedded-content.html#element-attrdef-img-srcset).
	//
	// Use of this type presents a security risk:
	// the encapsulated content should come from a trusted source,
	// as it will be included verbatim in the template output.
	Srcset	string
)

type contentType uint8

const (
	contentTypePlain	contentType	= iota
	contentTypeCSS
	contentTypeHTML
	contentTypeHTMLAttr
	contentTypeJS
	contentTypeJSStr
	contentTypeURL
	contentTypeSrcset
	// contentTypeUnsafe is used in attr.go for values that affect how
	// embedded content and network messages are formed, vetted,
	// or interpreted; or which credentials network messages carry.
	contentTypeUnsafe
)

// indirect returns the value, after dereferencing as many times
//line /usr/local/go/src/html/template/content.go:113
// as necessary to reach the base type (or nil).
//line /usr/local/go/src/html/template/content.go:115
func indirect(a any) any {
//line /usr/local/go/src/html/template/content.go:115
	_go_fuzz_dep_.CoverTab[30658]++
							if a == nil {
//line /usr/local/go/src/html/template/content.go:116
		_go_fuzz_dep_.CoverTab[30662]++
								return nil
//line /usr/local/go/src/html/template/content.go:117
		// _ = "end of CoverTab[30662]"
	} else {
//line /usr/local/go/src/html/template/content.go:118
		_go_fuzz_dep_.CoverTab[30663]++
//line /usr/local/go/src/html/template/content.go:118
		// _ = "end of CoverTab[30663]"
//line /usr/local/go/src/html/template/content.go:118
	}
//line /usr/local/go/src/html/template/content.go:118
	// _ = "end of CoverTab[30658]"
//line /usr/local/go/src/html/template/content.go:118
	_go_fuzz_dep_.CoverTab[30659]++
							if t := reflect.TypeOf(a); t.Kind() != reflect.Pointer {
//line /usr/local/go/src/html/template/content.go:119
		_go_fuzz_dep_.CoverTab[30664]++

								return a
//line /usr/local/go/src/html/template/content.go:121
		// _ = "end of CoverTab[30664]"
	} else {
//line /usr/local/go/src/html/template/content.go:122
		_go_fuzz_dep_.CoverTab[30665]++
//line /usr/local/go/src/html/template/content.go:122
		// _ = "end of CoverTab[30665]"
//line /usr/local/go/src/html/template/content.go:122
	}
//line /usr/local/go/src/html/template/content.go:122
	// _ = "end of CoverTab[30659]"
//line /usr/local/go/src/html/template/content.go:122
	_go_fuzz_dep_.CoverTab[30660]++
							v := reflect.ValueOf(a)
							for v.Kind() == reflect.Pointer && func() bool {
//line /usr/local/go/src/html/template/content.go:124
		_go_fuzz_dep_.CoverTab[30666]++
//line /usr/local/go/src/html/template/content.go:124
		return !v.IsNil()
//line /usr/local/go/src/html/template/content.go:124
		// _ = "end of CoverTab[30666]"
//line /usr/local/go/src/html/template/content.go:124
	}() {
//line /usr/local/go/src/html/template/content.go:124
		_go_fuzz_dep_.CoverTab[30667]++
								v = v.Elem()
//line /usr/local/go/src/html/template/content.go:125
		// _ = "end of CoverTab[30667]"
	}
//line /usr/local/go/src/html/template/content.go:126
	// _ = "end of CoverTab[30660]"
//line /usr/local/go/src/html/template/content.go:126
	_go_fuzz_dep_.CoverTab[30661]++
							return v.Interface()
//line /usr/local/go/src/html/template/content.go:127
	// _ = "end of CoverTab[30661]"
}

var (
	errorType	= reflect.TypeOf((*error)(nil)).Elem()
	fmtStringerType	= reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
)

// indirectToStringerOrError returns the value, after dereferencing as many times
//line /usr/local/go/src/html/template/content.go:135
// as necessary to reach the base type (or nil) or an implementation of fmt.Stringer
//line /usr/local/go/src/html/template/content.go:135
// or error.
//line /usr/local/go/src/html/template/content.go:138
func indirectToStringerOrError(a any) any {
//line /usr/local/go/src/html/template/content.go:138
	_go_fuzz_dep_.CoverTab[30668]++
							if a == nil {
//line /usr/local/go/src/html/template/content.go:139
		_go_fuzz_dep_.CoverTab[30671]++
								return nil
//line /usr/local/go/src/html/template/content.go:140
		// _ = "end of CoverTab[30671]"
	} else {
//line /usr/local/go/src/html/template/content.go:141
		_go_fuzz_dep_.CoverTab[30672]++
//line /usr/local/go/src/html/template/content.go:141
		// _ = "end of CoverTab[30672]"
//line /usr/local/go/src/html/template/content.go:141
	}
//line /usr/local/go/src/html/template/content.go:141
	// _ = "end of CoverTab[30668]"
//line /usr/local/go/src/html/template/content.go:141
	_go_fuzz_dep_.CoverTab[30669]++
							v := reflect.ValueOf(a)
							for !v.Type().Implements(fmtStringerType) && func() bool {
//line /usr/local/go/src/html/template/content.go:143
		_go_fuzz_dep_.CoverTab[30673]++
//line /usr/local/go/src/html/template/content.go:143
		return !v.Type().Implements(errorType)
//line /usr/local/go/src/html/template/content.go:143
		// _ = "end of CoverTab[30673]"
//line /usr/local/go/src/html/template/content.go:143
	}() && func() bool {
//line /usr/local/go/src/html/template/content.go:143
		_go_fuzz_dep_.CoverTab[30674]++
//line /usr/local/go/src/html/template/content.go:143
		return v.Kind() == reflect.Pointer
//line /usr/local/go/src/html/template/content.go:143
		// _ = "end of CoverTab[30674]"
//line /usr/local/go/src/html/template/content.go:143
	}() && func() bool {
//line /usr/local/go/src/html/template/content.go:143
		_go_fuzz_dep_.CoverTab[30675]++
//line /usr/local/go/src/html/template/content.go:143
		return !v.IsNil()
//line /usr/local/go/src/html/template/content.go:143
		// _ = "end of CoverTab[30675]"
//line /usr/local/go/src/html/template/content.go:143
	}() {
//line /usr/local/go/src/html/template/content.go:143
		_go_fuzz_dep_.CoverTab[30676]++
								v = v.Elem()
//line /usr/local/go/src/html/template/content.go:144
		// _ = "end of CoverTab[30676]"
	}
//line /usr/local/go/src/html/template/content.go:145
	// _ = "end of CoverTab[30669]"
//line /usr/local/go/src/html/template/content.go:145
	_go_fuzz_dep_.CoverTab[30670]++
							return v.Interface()
//line /usr/local/go/src/html/template/content.go:146
	// _ = "end of CoverTab[30670]"
}

// stringify converts its arguments to a string and the type of the content.
//line /usr/local/go/src/html/template/content.go:149
// All pointers are dereferenced, as in the text/template package.
//line /usr/local/go/src/html/template/content.go:151
func stringify(args ...any) (string, contentType) {
//line /usr/local/go/src/html/template/content.go:151
	_go_fuzz_dep_.CoverTab[30677]++
							if len(args) == 1 {
//line /usr/local/go/src/html/template/content.go:152
		_go_fuzz_dep_.CoverTab[30680]++
								switch s := indirect(args[0]).(type) {
		case string:
//line /usr/local/go/src/html/template/content.go:154
			_go_fuzz_dep_.CoverTab[30681]++
									return s, contentTypePlain
//line /usr/local/go/src/html/template/content.go:155
			// _ = "end of CoverTab[30681]"
		case CSS:
//line /usr/local/go/src/html/template/content.go:156
			_go_fuzz_dep_.CoverTab[30682]++
									return string(s), contentTypeCSS
//line /usr/local/go/src/html/template/content.go:157
			// _ = "end of CoverTab[30682]"
		case HTML:
//line /usr/local/go/src/html/template/content.go:158
			_go_fuzz_dep_.CoverTab[30683]++
									return string(s), contentTypeHTML
//line /usr/local/go/src/html/template/content.go:159
			// _ = "end of CoverTab[30683]"
		case HTMLAttr:
//line /usr/local/go/src/html/template/content.go:160
			_go_fuzz_dep_.CoverTab[30684]++
									return string(s), contentTypeHTMLAttr
//line /usr/local/go/src/html/template/content.go:161
			// _ = "end of CoverTab[30684]"
		case JS:
//line /usr/local/go/src/html/template/content.go:162
			_go_fuzz_dep_.CoverTab[30685]++
									return string(s), contentTypeJS
//line /usr/local/go/src/html/template/content.go:163
			// _ = "end of CoverTab[30685]"
		case JSStr:
//line /usr/local/go/src/html/template/content.go:164
			_go_fuzz_dep_.CoverTab[30686]++
									return string(s), contentTypeJSStr
//line /usr/local/go/src/html/template/content.go:165
			// _ = "end of CoverTab[30686]"
		case URL:
//line /usr/local/go/src/html/template/content.go:166
			_go_fuzz_dep_.CoverTab[30687]++
									return string(s), contentTypeURL
//line /usr/local/go/src/html/template/content.go:167
			// _ = "end of CoverTab[30687]"
		case Srcset:
//line /usr/local/go/src/html/template/content.go:168
			_go_fuzz_dep_.CoverTab[30688]++
									return string(s), contentTypeSrcset
//line /usr/local/go/src/html/template/content.go:169
			// _ = "end of CoverTab[30688]"
		}
//line /usr/local/go/src/html/template/content.go:170
		// _ = "end of CoverTab[30680]"
	} else {
//line /usr/local/go/src/html/template/content.go:171
		_go_fuzz_dep_.CoverTab[30689]++
//line /usr/local/go/src/html/template/content.go:171
		// _ = "end of CoverTab[30689]"
//line /usr/local/go/src/html/template/content.go:171
	}
//line /usr/local/go/src/html/template/content.go:171
	// _ = "end of CoverTab[30677]"
//line /usr/local/go/src/html/template/content.go:171
	_go_fuzz_dep_.CoverTab[30678]++
							i := 0
							for _, arg := range args {
//line /usr/local/go/src/html/template/content.go:173
		_go_fuzz_dep_.CoverTab[30690]++

//line /usr/local/go/src/html/template/content.go:177
		if arg == nil {
//line /usr/local/go/src/html/template/content.go:177
			_go_fuzz_dep_.CoverTab[30692]++
									continue
//line /usr/local/go/src/html/template/content.go:178
			// _ = "end of CoverTab[30692]"
		} else {
//line /usr/local/go/src/html/template/content.go:179
			_go_fuzz_dep_.CoverTab[30693]++
//line /usr/local/go/src/html/template/content.go:179
			// _ = "end of CoverTab[30693]"
//line /usr/local/go/src/html/template/content.go:179
		}
//line /usr/local/go/src/html/template/content.go:179
		// _ = "end of CoverTab[30690]"
//line /usr/local/go/src/html/template/content.go:179
		_go_fuzz_dep_.CoverTab[30691]++

								args[i] = indirectToStringerOrError(arg)
								i++
//line /usr/local/go/src/html/template/content.go:182
		// _ = "end of CoverTab[30691]"
	}
//line /usr/local/go/src/html/template/content.go:183
	// _ = "end of CoverTab[30678]"
//line /usr/local/go/src/html/template/content.go:183
	_go_fuzz_dep_.CoverTab[30679]++
							return fmt.Sprint(args[:i]...), contentTypePlain
//line /usr/local/go/src/html/template/content.go:184
	// _ = "end of CoverTab[30679]"
}

//line /usr/local/go/src/html/template/content.go:185
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/html/template/content.go:185
var _ = _go_fuzz_dep_.CoverTab
