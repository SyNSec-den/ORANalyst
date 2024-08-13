// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/html/template/attr.go:5
package template

//line /usr/local/go/src/html/template/attr.go:5
import (
//line /usr/local/go/src/html/template/attr.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/html/template/attr.go:5
)
//line /usr/local/go/src/html/template/attr.go:5
import (
//line /usr/local/go/src/html/template/attr.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/html/template/attr.go:5
)

import (
	"strings"
)

// attrTypeMap[n] describes the value of the given attribute.
//line /usr/local/go/src/html/template/attr.go:11
// If an attribute affects (or can mask) the encoding or interpretation of
//line /usr/local/go/src/html/template/attr.go:11
// other content, or affects the contents, idempotency, or credentials of a
//line /usr/local/go/src/html/template/attr.go:11
// network message, then the value in this map is contentTypeUnsafe.
//line /usr/local/go/src/html/template/attr.go:11
// This map is derived from HTML5, specifically
//line /usr/local/go/src/html/template/attr.go:11
// https://www.w3.org/TR/html5/Overview.html#attributes-1
//line /usr/local/go/src/html/template/attr.go:11
// as well as "%URI"-typed attributes from
//line /usr/local/go/src/html/template/attr.go:11
// https://www.w3.org/TR/html4/index/attributes.html
//line /usr/local/go/src/html/template/attr.go:19
var attrTypeMap = map[string]contentType{
							"accept":		contentTypePlain,
							"accept-charset":	contentTypeUnsafe,
							"action":		contentTypeURL,
							"alt":			contentTypePlain,
							"archive":		contentTypeURL,
							"async":		contentTypeUnsafe,
							"autocomplete":		contentTypePlain,
							"autofocus":		contentTypePlain,
							"autoplay":		contentTypePlain,
							"background":		contentTypeURL,
							"border":		contentTypePlain,
							"checked":		contentTypePlain,
							"cite":			contentTypeURL,
							"challenge":		contentTypeUnsafe,
							"charset":		contentTypeUnsafe,
							"class":		contentTypePlain,
							"classid":		contentTypeURL,
							"codebase":		contentTypeURL,
							"cols":			contentTypePlain,
							"colspan":		contentTypePlain,
							"content":		contentTypeUnsafe,
							"contenteditable":	contentTypePlain,
							"contextmenu":		contentTypePlain,
							"controls":		contentTypePlain,
							"coords":		contentTypePlain,
							"crossorigin":		contentTypeUnsafe,
							"data":			contentTypeURL,
							"datetime":		contentTypePlain,
							"default":		contentTypePlain,
							"defer":		contentTypeUnsafe,
							"dir":			contentTypePlain,
							"dirname":		contentTypePlain,
							"disabled":		contentTypePlain,
							"draggable":		contentTypePlain,
							"dropzone":		contentTypePlain,
							"enctype":		contentTypeUnsafe,
							"for":			contentTypePlain,
							"form":			contentTypeUnsafe,
							"formaction":		contentTypeURL,
							"formenctype":		contentTypeUnsafe,
							"formmethod":		contentTypeUnsafe,
							"formnovalidate":	contentTypeUnsafe,
							"formtarget":		contentTypePlain,
							"headers":		contentTypePlain,
							"height":		contentTypePlain,
							"hidden":		contentTypePlain,
							"high":			contentTypePlain,
							"href":			contentTypeURL,
							"hreflang":		contentTypePlain,
							"http-equiv":		contentTypeUnsafe,
							"icon":			contentTypeURL,
							"id":			contentTypePlain,
							"ismap":		contentTypePlain,
							"keytype":		contentTypeUnsafe,
							"kind":			contentTypePlain,
							"label":		contentTypePlain,
							"lang":			contentTypePlain,
							"language":		contentTypeUnsafe,
							"list":			contentTypePlain,
							"longdesc":		contentTypeURL,
							"loop":			contentTypePlain,
							"low":			contentTypePlain,
							"manifest":		contentTypeURL,
							"max":			contentTypePlain,
							"maxlength":		contentTypePlain,
							"media":		contentTypePlain,
							"mediagroup":		contentTypePlain,
							"method":		contentTypeUnsafe,
							"min":			contentTypePlain,
							"multiple":		contentTypePlain,
							"name":			contentTypePlain,
							"novalidate":		contentTypeUnsafe,

//line /usr/local/go/src/html/template/attr.go:95
	"open":		contentTypePlain,
							"optimum":	contentTypePlain,
							"pattern":	contentTypeUnsafe,
							"placeholder":	contentTypePlain,
							"poster":	contentTypeURL,
							"profile":	contentTypeURL,
							"preload":	contentTypePlain,
							"pubdate":	contentTypePlain,
							"radiogroup":	contentTypePlain,
							"readonly":	contentTypePlain,
							"rel":		contentTypeUnsafe,
							"required":	contentTypePlain,
							"reversed":	contentTypePlain,
							"rows":		contentTypePlain,
							"rowspan":	contentTypePlain,
							"sandbox":	contentTypeUnsafe,
							"spellcheck":	contentTypePlain,
							"scope":	contentTypePlain,
							"scoped":	contentTypePlain,
							"seamless":	contentTypePlain,
							"selected":	contentTypePlain,
							"shape":	contentTypePlain,
							"size":		contentTypePlain,
							"sizes":	contentTypePlain,
							"span":		contentTypePlain,
							"src":		contentTypeURL,
							"srcdoc":	contentTypeHTML,
							"srclang":	contentTypePlain,
							"srcset":	contentTypeSrcset,
							"start":	contentTypePlain,
							"step":		contentTypePlain,
							"style":	contentTypeCSS,
							"tabindex":	contentTypePlain,
							"target":	contentTypePlain,
							"title":	contentTypePlain,
							"type":		contentTypeUnsafe,
							"usemap":	contentTypeURL,
							"value":	contentTypeUnsafe,
							"width":	contentTypePlain,
							"wrap":		contentTypePlain,
							"xmlns":	contentTypeURL,
}

// attrType returns a conservative (upper-bound on authority) guess at the
//line /usr/local/go/src/html/template/attr.go:138
// type of the lowercase named attribute.
//line /usr/local/go/src/html/template/attr.go:140
func attrType(name string) contentType {
//line /usr/local/go/src/html/template/attr.go:140
	_go_fuzz_dep_.CoverTab[30634]++
							if strings.HasPrefix(name, "data-") {
//line /usr/local/go/src/html/template/attr.go:141
		_go_fuzz_dep_.CoverTab[30639]++

//line /usr/local/go/src/html/template/attr.go:145
		name = name[5:]
//line /usr/local/go/src/html/template/attr.go:145
		// _ = "end of CoverTab[30639]"
	} else {
//line /usr/local/go/src/html/template/attr.go:146
		_go_fuzz_dep_.CoverTab[30640]++
//line /usr/local/go/src/html/template/attr.go:146
		if prefix, short, ok := strings.Cut(name, ":"); ok {
//line /usr/local/go/src/html/template/attr.go:146
			_go_fuzz_dep_.CoverTab[30641]++
									if prefix == "xmlns" {
//line /usr/local/go/src/html/template/attr.go:147
				_go_fuzz_dep_.CoverTab[30643]++
										return contentTypeURL
//line /usr/local/go/src/html/template/attr.go:148
				// _ = "end of CoverTab[30643]"
			} else {
//line /usr/local/go/src/html/template/attr.go:149
				_go_fuzz_dep_.CoverTab[30644]++
//line /usr/local/go/src/html/template/attr.go:149
				// _ = "end of CoverTab[30644]"
//line /usr/local/go/src/html/template/attr.go:149
			}
//line /usr/local/go/src/html/template/attr.go:149
			// _ = "end of CoverTab[30641]"
//line /usr/local/go/src/html/template/attr.go:149
			_go_fuzz_dep_.CoverTab[30642]++

									name = short
//line /usr/local/go/src/html/template/attr.go:151
			// _ = "end of CoverTab[30642]"
		} else {
//line /usr/local/go/src/html/template/attr.go:152
			_go_fuzz_dep_.CoverTab[30645]++
//line /usr/local/go/src/html/template/attr.go:152
			// _ = "end of CoverTab[30645]"
//line /usr/local/go/src/html/template/attr.go:152
		}
//line /usr/local/go/src/html/template/attr.go:152
		// _ = "end of CoverTab[30640]"
//line /usr/local/go/src/html/template/attr.go:152
	}
//line /usr/local/go/src/html/template/attr.go:152
	// _ = "end of CoverTab[30634]"
//line /usr/local/go/src/html/template/attr.go:152
	_go_fuzz_dep_.CoverTab[30635]++
							if t, ok := attrTypeMap[name]; ok {
//line /usr/local/go/src/html/template/attr.go:153
		_go_fuzz_dep_.CoverTab[30646]++
								return t
//line /usr/local/go/src/html/template/attr.go:154
		// _ = "end of CoverTab[30646]"
	} else {
//line /usr/local/go/src/html/template/attr.go:155
		_go_fuzz_dep_.CoverTab[30647]++
//line /usr/local/go/src/html/template/attr.go:155
		// _ = "end of CoverTab[30647]"
//line /usr/local/go/src/html/template/attr.go:155
	}
//line /usr/local/go/src/html/template/attr.go:155
	// _ = "end of CoverTab[30635]"
//line /usr/local/go/src/html/template/attr.go:155
	_go_fuzz_dep_.CoverTab[30636]++

							if strings.HasPrefix(name, "on") {
//line /usr/local/go/src/html/template/attr.go:157
		_go_fuzz_dep_.CoverTab[30648]++
								return contentTypeJS
//line /usr/local/go/src/html/template/attr.go:158
		// _ = "end of CoverTab[30648]"
	} else {
//line /usr/local/go/src/html/template/attr.go:159
		_go_fuzz_dep_.CoverTab[30649]++
//line /usr/local/go/src/html/template/attr.go:159
		// _ = "end of CoverTab[30649]"
//line /usr/local/go/src/html/template/attr.go:159
	}
//line /usr/local/go/src/html/template/attr.go:159
	// _ = "end of CoverTab[30636]"
//line /usr/local/go/src/html/template/attr.go:159
	_go_fuzz_dep_.CoverTab[30637]++

//line /usr/local/go/src/html/template/attr.go:169
	if strings.Contains(name, "src") || func() bool {
//line /usr/local/go/src/html/template/attr.go:169
		_go_fuzz_dep_.CoverTab[30650]++
//line /usr/local/go/src/html/template/attr.go:169
		return strings.Contains(name, "uri")
								// _ = "end of CoverTab[30650]"
//line /usr/local/go/src/html/template/attr.go:170
	}() || func() bool {
//line /usr/local/go/src/html/template/attr.go:170
		_go_fuzz_dep_.CoverTab[30651]++
//line /usr/local/go/src/html/template/attr.go:170
		return strings.Contains(name, "url")
								// _ = "end of CoverTab[30651]"
//line /usr/local/go/src/html/template/attr.go:171
	}() {
//line /usr/local/go/src/html/template/attr.go:171
		_go_fuzz_dep_.CoverTab[30652]++
								return contentTypeURL
//line /usr/local/go/src/html/template/attr.go:172
		// _ = "end of CoverTab[30652]"
	} else {
//line /usr/local/go/src/html/template/attr.go:173
		_go_fuzz_dep_.CoverTab[30653]++
//line /usr/local/go/src/html/template/attr.go:173
		// _ = "end of CoverTab[30653]"
//line /usr/local/go/src/html/template/attr.go:173
	}
//line /usr/local/go/src/html/template/attr.go:173
	// _ = "end of CoverTab[30637]"
//line /usr/local/go/src/html/template/attr.go:173
	_go_fuzz_dep_.CoverTab[30638]++
							return contentTypePlain
//line /usr/local/go/src/html/template/attr.go:174
	// _ = "end of CoverTab[30638]"
}

//line /usr/local/go/src/html/template/attr.go:175
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/html/template/attr.go:175
var _ = _go_fuzz_dep_.CoverTab
