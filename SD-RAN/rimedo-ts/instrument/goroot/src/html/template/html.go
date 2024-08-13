// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/html/template/html.go:5
package template

//line /usr/local/go/src/html/template/html.go:5
import (
//line /usr/local/go/src/html/template/html.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/html/template/html.go:5
)
//line /usr/local/go/src/html/template/html.go:5
import (
//line /usr/local/go/src/html/template/html.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/html/template/html.go:5
)

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

// htmlNospaceEscaper escapes for inclusion in unquoted attribute values.
func htmlNospaceEscaper(args ...any) string {
//line /usr/local/go/src/html/template/html.go:15
	_go_fuzz_dep_.CoverTab[31224]++
							s, t := stringify(args...)
							if s == "" {
//line /usr/local/go/src/html/template/html.go:17
		_go_fuzz_dep_.CoverTab[31227]++
								return filterFailsafe
//line /usr/local/go/src/html/template/html.go:18
		// _ = "end of CoverTab[31227]"
	} else {
//line /usr/local/go/src/html/template/html.go:19
		_go_fuzz_dep_.CoverTab[31228]++
//line /usr/local/go/src/html/template/html.go:19
		// _ = "end of CoverTab[31228]"
//line /usr/local/go/src/html/template/html.go:19
	}
//line /usr/local/go/src/html/template/html.go:19
	// _ = "end of CoverTab[31224]"
//line /usr/local/go/src/html/template/html.go:19
	_go_fuzz_dep_.CoverTab[31225]++
							if t == contentTypeHTML {
//line /usr/local/go/src/html/template/html.go:20
		_go_fuzz_dep_.CoverTab[31229]++
								return htmlReplacer(stripTags(s), htmlNospaceNormReplacementTable, false)
//line /usr/local/go/src/html/template/html.go:21
		// _ = "end of CoverTab[31229]"
	} else {
//line /usr/local/go/src/html/template/html.go:22
		_go_fuzz_dep_.CoverTab[31230]++
//line /usr/local/go/src/html/template/html.go:22
		// _ = "end of CoverTab[31230]"
//line /usr/local/go/src/html/template/html.go:22
	}
//line /usr/local/go/src/html/template/html.go:22
	// _ = "end of CoverTab[31225]"
//line /usr/local/go/src/html/template/html.go:22
	_go_fuzz_dep_.CoverTab[31226]++
							return htmlReplacer(s, htmlNospaceReplacementTable, false)
//line /usr/local/go/src/html/template/html.go:23
	// _ = "end of CoverTab[31226]"
}

// attrEscaper escapes for inclusion in quoted attribute values.
func attrEscaper(args ...any) string {
//line /usr/local/go/src/html/template/html.go:27
	_go_fuzz_dep_.CoverTab[31231]++
							s, t := stringify(args...)
							if t == contentTypeHTML {
//line /usr/local/go/src/html/template/html.go:29
		_go_fuzz_dep_.CoverTab[31233]++
								return htmlReplacer(stripTags(s), htmlNormReplacementTable, true)
//line /usr/local/go/src/html/template/html.go:30
		// _ = "end of CoverTab[31233]"
	} else {
//line /usr/local/go/src/html/template/html.go:31
		_go_fuzz_dep_.CoverTab[31234]++
//line /usr/local/go/src/html/template/html.go:31
		// _ = "end of CoverTab[31234]"
//line /usr/local/go/src/html/template/html.go:31
	}
//line /usr/local/go/src/html/template/html.go:31
	// _ = "end of CoverTab[31231]"
//line /usr/local/go/src/html/template/html.go:31
	_go_fuzz_dep_.CoverTab[31232]++
							return htmlReplacer(s, htmlReplacementTable, true)
//line /usr/local/go/src/html/template/html.go:32
	// _ = "end of CoverTab[31232]"
}

// rcdataEscaper escapes for inclusion in an RCDATA element body.
func rcdataEscaper(args ...any) string {
//line /usr/local/go/src/html/template/html.go:36
	_go_fuzz_dep_.CoverTab[31235]++
							s, t := stringify(args...)
							if t == contentTypeHTML {
//line /usr/local/go/src/html/template/html.go:38
		_go_fuzz_dep_.CoverTab[31237]++
								return htmlReplacer(s, htmlNormReplacementTable, true)
//line /usr/local/go/src/html/template/html.go:39
		// _ = "end of CoverTab[31237]"
	} else {
//line /usr/local/go/src/html/template/html.go:40
		_go_fuzz_dep_.CoverTab[31238]++
//line /usr/local/go/src/html/template/html.go:40
		// _ = "end of CoverTab[31238]"
//line /usr/local/go/src/html/template/html.go:40
	}
//line /usr/local/go/src/html/template/html.go:40
	// _ = "end of CoverTab[31235]"
//line /usr/local/go/src/html/template/html.go:40
	_go_fuzz_dep_.CoverTab[31236]++
							return htmlReplacer(s, htmlReplacementTable, true)
//line /usr/local/go/src/html/template/html.go:41
	// _ = "end of CoverTab[31236]"
}

// htmlEscaper escapes for inclusion in HTML text.
func htmlEscaper(args ...any) string {
//line /usr/local/go/src/html/template/html.go:45
	_go_fuzz_dep_.CoverTab[31239]++
							s, t := stringify(args...)
							if t == contentTypeHTML {
//line /usr/local/go/src/html/template/html.go:47
		_go_fuzz_dep_.CoverTab[31241]++
								return s
//line /usr/local/go/src/html/template/html.go:48
		// _ = "end of CoverTab[31241]"
	} else {
//line /usr/local/go/src/html/template/html.go:49
		_go_fuzz_dep_.CoverTab[31242]++
//line /usr/local/go/src/html/template/html.go:49
		// _ = "end of CoverTab[31242]"
//line /usr/local/go/src/html/template/html.go:49
	}
//line /usr/local/go/src/html/template/html.go:49
	// _ = "end of CoverTab[31239]"
//line /usr/local/go/src/html/template/html.go:49
	_go_fuzz_dep_.CoverTab[31240]++
							return htmlReplacer(s, htmlReplacementTable, true)
//line /usr/local/go/src/html/template/html.go:50
	// _ = "end of CoverTab[31240]"
}

// htmlReplacementTable contains the runes that need to be escaped
//line /usr/local/go/src/html/template/html.go:53
// inside a quoted attribute value or in a text node.
//line /usr/local/go/src/html/template/html.go:55
var htmlReplacementTable = []string{

//line /usr/local/go/src/html/template/html.go:62
	0:	"\uFFFD",
							'"':	"&#34;",
							'&':	"&amp;",
							'\'':	"&#39;",
							'+':	"&#43;",
							'<':	"&lt;",
							'>':	"&gt;",
}

// htmlNormReplacementTable is like htmlReplacementTable but without '&' to
//line /usr/local/go/src/html/template/html.go:71
// avoid over-encoding existing entities.
//line /usr/local/go/src/html/template/html.go:73
var htmlNormReplacementTable = []string{
	0:	"\uFFFD",
	'"':	"&#34;",
	'\'':	"&#39;",
	'+':	"&#43;",
	'<':	"&lt;",
	'>':	"&gt;",
}

// htmlNospaceReplacementTable contains the runes that need to be escaped
//line /usr/local/go/src/html/template/html.go:82
// inside an unquoted attribute value.
//line /usr/local/go/src/html/template/html.go:82
// The set of runes escaped is the union of the HTML specials and
//line /usr/local/go/src/html/template/html.go:82
// those determined by running the JS below in browsers:
//line /usr/local/go/src/html/template/html.go:82
// <div id=d></div>
//line /usr/local/go/src/html/template/html.go:82
// <script>(function () {
//line /usr/local/go/src/html/template/html.go:82
// var a = [], d = document.getElementById("d"), i, c, s;
//line /usr/local/go/src/html/template/html.go:82
// for (i = 0; i < 0x10000; ++i) {
//line /usr/local/go/src/html/template/html.go:82
//
//line /usr/local/go/src/html/template/html.go:82
//	c = String.fromCharCode(i);
//line /usr/local/go/src/html/template/html.go:82
//	d.innerHTML = "<span title=" + c + "lt" + c + "></span>"
//line /usr/local/go/src/html/template/html.go:82
//	s = d.getElementsByTagName("SPAN")[0];
//line /usr/local/go/src/html/template/html.go:82
//	if (!s || s.title !== c + "lt" + c) { a.push(i.toString(16)); }
//line /usr/local/go/src/html/template/html.go:82
//
//line /usr/local/go/src/html/template/html.go:82
// }
//line /usr/local/go/src/html/template/html.go:82
// document.write(a.join(", "));
//line /usr/local/go/src/html/template/html.go:82
// })()</script>
//line /usr/local/go/src/html/template/html.go:99
var htmlNospaceReplacementTable = []string{
							0:	"&#xfffd;",
							'\t':	"&#9;",
							'\n':	"&#10;",
							'\v':	"&#11;",
							'\f':	"&#12;",
							'\r':	"&#13;",
							' ':	"&#32;",
							'"':	"&#34;",
							'&':	"&amp;",
							'\'':	"&#39;",
							'+':	"&#43;",
							'<':	"&lt;",
							'=':	"&#61;",
							'>':	"&gt;",

//line /usr/local/go/src/html/template/html.go:117
	'`':	"&#96;",
}

// htmlNospaceNormReplacementTable is like htmlNospaceReplacementTable but
//line /usr/local/go/src/html/template/html.go:120
// without '&' to avoid over-encoding existing entities.
//line /usr/local/go/src/html/template/html.go:122
var htmlNospaceNormReplacementTable = []string{
							0:	"&#xfffd;",
							'\t':	"&#9;",
							'\n':	"&#10;",
							'\v':	"&#11;",
							'\f':	"&#12;",
							'\r':	"&#13;",
							' ':	"&#32;",
							'"':	"&#34;",
							'\'':	"&#39;",
							'+':	"&#43;",
							'<':	"&lt;",
							'=':	"&#61;",
							'>':	"&gt;",

//line /usr/local/go/src/html/template/html.go:139
	'`':	"&#96;",
}

// htmlReplacer returns s with runes replaced according to replacementTable
//line /usr/local/go/src/html/template/html.go:142
// and when badRunes is true, certain bad runes are allowed through unescaped.
//line /usr/local/go/src/html/template/html.go:144
func htmlReplacer(s string, replacementTable []string, badRunes bool) string {
//line /usr/local/go/src/html/template/html.go:144
	_go_fuzz_dep_.CoverTab[31243]++
							written, b := 0, new(strings.Builder)
							r, w := rune(0), 0
							for i := 0; i < len(s); i += w {
//line /usr/local/go/src/html/template/html.go:147
		_go_fuzz_dep_.CoverTab[31246]++

//line /usr/local/go/src/html/template/html.go:151
		r, w = utf8.DecodeRuneInString(s[i:])
		if int(r) < len(replacementTable) {
//line /usr/local/go/src/html/template/html.go:152
			_go_fuzz_dep_.CoverTab[31247]++
									if repl := replacementTable[r]; len(repl) != 0 {
//line /usr/local/go/src/html/template/html.go:153
				_go_fuzz_dep_.CoverTab[31248]++
										if written == 0 {
//line /usr/local/go/src/html/template/html.go:154
					_go_fuzz_dep_.CoverTab[31250]++
											b.Grow(len(s))
//line /usr/local/go/src/html/template/html.go:155
					// _ = "end of CoverTab[31250]"
				} else {
//line /usr/local/go/src/html/template/html.go:156
					_go_fuzz_dep_.CoverTab[31251]++
//line /usr/local/go/src/html/template/html.go:156
					// _ = "end of CoverTab[31251]"
//line /usr/local/go/src/html/template/html.go:156
				}
//line /usr/local/go/src/html/template/html.go:156
				// _ = "end of CoverTab[31248]"
//line /usr/local/go/src/html/template/html.go:156
				_go_fuzz_dep_.CoverTab[31249]++
										b.WriteString(s[written:i])
										b.WriteString(repl)
										written = i + w
//line /usr/local/go/src/html/template/html.go:159
				// _ = "end of CoverTab[31249]"
			} else {
//line /usr/local/go/src/html/template/html.go:160
				_go_fuzz_dep_.CoverTab[31252]++
//line /usr/local/go/src/html/template/html.go:160
				// _ = "end of CoverTab[31252]"
//line /usr/local/go/src/html/template/html.go:160
			}
//line /usr/local/go/src/html/template/html.go:160
			// _ = "end of CoverTab[31247]"
		} else {
//line /usr/local/go/src/html/template/html.go:161
			_go_fuzz_dep_.CoverTab[31253]++
//line /usr/local/go/src/html/template/html.go:161
			if badRunes {
//line /usr/local/go/src/html/template/html.go:161
				_go_fuzz_dep_.CoverTab[31254]++
//line /usr/local/go/src/html/template/html.go:161
				// _ = "end of CoverTab[31254]"

//line /usr/local/go/src/html/template/html.go:164
			} else {
//line /usr/local/go/src/html/template/html.go:164
				_go_fuzz_dep_.CoverTab[31255]++
//line /usr/local/go/src/html/template/html.go:164
				if 0xfdd0 <= r && func() bool {
//line /usr/local/go/src/html/template/html.go:164
					_go_fuzz_dep_.CoverTab[31256]++
//line /usr/local/go/src/html/template/html.go:164
					return r <= 0xfdef
//line /usr/local/go/src/html/template/html.go:164
					// _ = "end of CoverTab[31256]"
//line /usr/local/go/src/html/template/html.go:164
				}() || func() bool {
//line /usr/local/go/src/html/template/html.go:164
					_go_fuzz_dep_.CoverTab[31257]++
//line /usr/local/go/src/html/template/html.go:164
					return 0xfff0 <= r && func() bool {
//line /usr/local/go/src/html/template/html.go:164
						_go_fuzz_dep_.CoverTab[31258]++
//line /usr/local/go/src/html/template/html.go:164
						return r <= 0xffff
//line /usr/local/go/src/html/template/html.go:164
						// _ = "end of CoverTab[31258]"
//line /usr/local/go/src/html/template/html.go:164
					}()
//line /usr/local/go/src/html/template/html.go:164
					// _ = "end of CoverTab[31257]"
//line /usr/local/go/src/html/template/html.go:164
				}() {
//line /usr/local/go/src/html/template/html.go:164
					_go_fuzz_dep_.CoverTab[31259]++
											if written == 0 {
//line /usr/local/go/src/html/template/html.go:165
						_go_fuzz_dep_.CoverTab[31261]++
												b.Grow(len(s))
//line /usr/local/go/src/html/template/html.go:166
						// _ = "end of CoverTab[31261]"
					} else {
//line /usr/local/go/src/html/template/html.go:167
						_go_fuzz_dep_.CoverTab[31262]++
//line /usr/local/go/src/html/template/html.go:167
						// _ = "end of CoverTab[31262]"
//line /usr/local/go/src/html/template/html.go:167
					}
//line /usr/local/go/src/html/template/html.go:167
					// _ = "end of CoverTab[31259]"
//line /usr/local/go/src/html/template/html.go:167
					_go_fuzz_dep_.CoverTab[31260]++
											fmt.Fprintf(b, "%s&#x%x;", s[written:i], r)
											written = i + w
//line /usr/local/go/src/html/template/html.go:169
					// _ = "end of CoverTab[31260]"
				} else {
//line /usr/local/go/src/html/template/html.go:170
					_go_fuzz_dep_.CoverTab[31263]++
//line /usr/local/go/src/html/template/html.go:170
					// _ = "end of CoverTab[31263]"
//line /usr/local/go/src/html/template/html.go:170
				}
//line /usr/local/go/src/html/template/html.go:170
				// _ = "end of CoverTab[31255]"
//line /usr/local/go/src/html/template/html.go:170
			}
//line /usr/local/go/src/html/template/html.go:170
			// _ = "end of CoverTab[31253]"
//line /usr/local/go/src/html/template/html.go:170
		}
//line /usr/local/go/src/html/template/html.go:170
		// _ = "end of CoverTab[31246]"
	}
//line /usr/local/go/src/html/template/html.go:171
	// _ = "end of CoverTab[31243]"
//line /usr/local/go/src/html/template/html.go:171
	_go_fuzz_dep_.CoverTab[31244]++
							if written == 0 {
//line /usr/local/go/src/html/template/html.go:172
		_go_fuzz_dep_.CoverTab[31264]++
								return s
//line /usr/local/go/src/html/template/html.go:173
		// _ = "end of CoverTab[31264]"
	} else {
//line /usr/local/go/src/html/template/html.go:174
		_go_fuzz_dep_.CoverTab[31265]++
//line /usr/local/go/src/html/template/html.go:174
		// _ = "end of CoverTab[31265]"
//line /usr/local/go/src/html/template/html.go:174
	}
//line /usr/local/go/src/html/template/html.go:174
	// _ = "end of CoverTab[31244]"
//line /usr/local/go/src/html/template/html.go:174
	_go_fuzz_dep_.CoverTab[31245]++
							b.WriteString(s[written:])
							return b.String()
//line /usr/local/go/src/html/template/html.go:176
	// _ = "end of CoverTab[31245]"
}

// stripTags takes a snippet of HTML and returns only the text content.
//line /usr/local/go/src/html/template/html.go:179
// For example, `<b>&iexcl;Hi!</b> <script>...</script>` -> `&iexcl;Hi! `.
//line /usr/local/go/src/html/template/html.go:181
func stripTags(html string) string {
//line /usr/local/go/src/html/template/html.go:181
	_go_fuzz_dep_.CoverTab[31266]++
							var b strings.Builder
							s, c, i, allText := []byte(html), context{}, 0, true

//line /usr/local/go/src/html/template/html.go:186
	for i != len(s) {
//line /usr/local/go/src/html/template/html.go:186
		_go_fuzz_dep_.CoverTab[31269]++
								if c.delim == delimNone {
//line /usr/local/go/src/html/template/html.go:187
			_go_fuzz_dep_.CoverTab[31273]++
									st := c.state

									if c.element != elementNone && func() bool {
//line /usr/local/go/src/html/template/html.go:190
				_go_fuzz_dep_.CoverTab[31276]++
//line /usr/local/go/src/html/template/html.go:190
				return !isInTag(st)
//line /usr/local/go/src/html/template/html.go:190
				// _ = "end of CoverTab[31276]"
//line /usr/local/go/src/html/template/html.go:190
			}() {
//line /usr/local/go/src/html/template/html.go:190
				_go_fuzz_dep_.CoverTab[31277]++
										st = stateRCDATA
//line /usr/local/go/src/html/template/html.go:191
				// _ = "end of CoverTab[31277]"
			} else {
//line /usr/local/go/src/html/template/html.go:192
				_go_fuzz_dep_.CoverTab[31278]++
//line /usr/local/go/src/html/template/html.go:192
				// _ = "end of CoverTab[31278]"
//line /usr/local/go/src/html/template/html.go:192
			}
//line /usr/local/go/src/html/template/html.go:192
			// _ = "end of CoverTab[31273]"
//line /usr/local/go/src/html/template/html.go:192
			_go_fuzz_dep_.CoverTab[31274]++
									d, nread := transitionFunc[st](c, s[i:])
									i1 := i + nread
									if c.state == stateText || func() bool {
//line /usr/local/go/src/html/template/html.go:195
				_go_fuzz_dep_.CoverTab[31279]++
//line /usr/local/go/src/html/template/html.go:195
				return c.state == stateRCDATA
//line /usr/local/go/src/html/template/html.go:195
				// _ = "end of CoverTab[31279]"
//line /usr/local/go/src/html/template/html.go:195
			}() {
//line /usr/local/go/src/html/template/html.go:195
				_go_fuzz_dep_.CoverTab[31280]++

										j := i1
										if d.state != c.state {
//line /usr/local/go/src/html/template/html.go:198
					_go_fuzz_dep_.CoverTab[31282]++
											for j1 := j - 1; j1 >= i; j1-- {
//line /usr/local/go/src/html/template/html.go:199
						_go_fuzz_dep_.CoverTab[31283]++
												if s[j1] == '<' {
//line /usr/local/go/src/html/template/html.go:200
							_go_fuzz_dep_.CoverTab[31284]++
													j = j1
													break
//line /usr/local/go/src/html/template/html.go:202
							// _ = "end of CoverTab[31284]"
						} else {
//line /usr/local/go/src/html/template/html.go:203
							_go_fuzz_dep_.CoverTab[31285]++
//line /usr/local/go/src/html/template/html.go:203
							// _ = "end of CoverTab[31285]"
//line /usr/local/go/src/html/template/html.go:203
						}
//line /usr/local/go/src/html/template/html.go:203
						// _ = "end of CoverTab[31283]"
					}
//line /usr/local/go/src/html/template/html.go:204
					// _ = "end of CoverTab[31282]"
				} else {
//line /usr/local/go/src/html/template/html.go:205
					_go_fuzz_dep_.CoverTab[31286]++
//line /usr/local/go/src/html/template/html.go:205
					// _ = "end of CoverTab[31286]"
//line /usr/local/go/src/html/template/html.go:205
				}
//line /usr/local/go/src/html/template/html.go:205
				// _ = "end of CoverTab[31280]"
//line /usr/local/go/src/html/template/html.go:205
				_go_fuzz_dep_.CoverTab[31281]++
										b.Write(s[i:j])
//line /usr/local/go/src/html/template/html.go:206
				// _ = "end of CoverTab[31281]"
			} else {
//line /usr/local/go/src/html/template/html.go:207
				_go_fuzz_dep_.CoverTab[31287]++
										allText = false
//line /usr/local/go/src/html/template/html.go:208
				// _ = "end of CoverTab[31287]"
			}
//line /usr/local/go/src/html/template/html.go:209
			// _ = "end of CoverTab[31274]"
//line /usr/local/go/src/html/template/html.go:209
			_go_fuzz_dep_.CoverTab[31275]++
									c, i = d, i1
									continue
//line /usr/local/go/src/html/template/html.go:211
			// _ = "end of CoverTab[31275]"
		} else {
//line /usr/local/go/src/html/template/html.go:212
			_go_fuzz_dep_.CoverTab[31288]++
//line /usr/local/go/src/html/template/html.go:212
			// _ = "end of CoverTab[31288]"
//line /usr/local/go/src/html/template/html.go:212
		}
//line /usr/local/go/src/html/template/html.go:212
		// _ = "end of CoverTab[31269]"
//line /usr/local/go/src/html/template/html.go:212
		_go_fuzz_dep_.CoverTab[31270]++
								i1 := i + bytes.IndexAny(s[i:], delimEnds[c.delim])
								if i1 < i {
//line /usr/local/go/src/html/template/html.go:214
			_go_fuzz_dep_.CoverTab[31289]++
									break
//line /usr/local/go/src/html/template/html.go:215
			// _ = "end of CoverTab[31289]"
		} else {
//line /usr/local/go/src/html/template/html.go:216
			_go_fuzz_dep_.CoverTab[31290]++
//line /usr/local/go/src/html/template/html.go:216
			// _ = "end of CoverTab[31290]"
//line /usr/local/go/src/html/template/html.go:216
		}
//line /usr/local/go/src/html/template/html.go:216
		// _ = "end of CoverTab[31270]"
//line /usr/local/go/src/html/template/html.go:216
		_go_fuzz_dep_.CoverTab[31271]++
								if c.delim != delimSpaceOrTagEnd {
//line /usr/local/go/src/html/template/html.go:217
			_go_fuzz_dep_.CoverTab[31291]++

									i1++
//line /usr/local/go/src/html/template/html.go:219
			// _ = "end of CoverTab[31291]"
		} else {
//line /usr/local/go/src/html/template/html.go:220
			_go_fuzz_dep_.CoverTab[31292]++
//line /usr/local/go/src/html/template/html.go:220
			// _ = "end of CoverTab[31292]"
//line /usr/local/go/src/html/template/html.go:220
		}
//line /usr/local/go/src/html/template/html.go:220
		// _ = "end of CoverTab[31271]"
//line /usr/local/go/src/html/template/html.go:220
		_go_fuzz_dep_.CoverTab[31272]++
								c, i = context{state: stateTag, element: c.element}, i1
//line /usr/local/go/src/html/template/html.go:221
		// _ = "end of CoverTab[31272]"
	}
//line /usr/local/go/src/html/template/html.go:222
	// _ = "end of CoverTab[31266]"
//line /usr/local/go/src/html/template/html.go:222
	_go_fuzz_dep_.CoverTab[31267]++
							if allText {
//line /usr/local/go/src/html/template/html.go:223
		_go_fuzz_dep_.CoverTab[31293]++
								return html
//line /usr/local/go/src/html/template/html.go:224
		// _ = "end of CoverTab[31293]"
	} else {
//line /usr/local/go/src/html/template/html.go:225
		_go_fuzz_dep_.CoverTab[31294]++
//line /usr/local/go/src/html/template/html.go:225
		if c.state == stateText || func() bool {
//line /usr/local/go/src/html/template/html.go:225
			_go_fuzz_dep_.CoverTab[31295]++
//line /usr/local/go/src/html/template/html.go:225
			return c.state == stateRCDATA
//line /usr/local/go/src/html/template/html.go:225
			// _ = "end of CoverTab[31295]"
//line /usr/local/go/src/html/template/html.go:225
		}() {
//line /usr/local/go/src/html/template/html.go:225
			_go_fuzz_dep_.CoverTab[31296]++
									b.Write(s[i:])
//line /usr/local/go/src/html/template/html.go:226
			// _ = "end of CoverTab[31296]"
		} else {
//line /usr/local/go/src/html/template/html.go:227
			_go_fuzz_dep_.CoverTab[31297]++
//line /usr/local/go/src/html/template/html.go:227
			// _ = "end of CoverTab[31297]"
//line /usr/local/go/src/html/template/html.go:227
		}
//line /usr/local/go/src/html/template/html.go:227
		// _ = "end of CoverTab[31294]"
//line /usr/local/go/src/html/template/html.go:227
	}
//line /usr/local/go/src/html/template/html.go:227
	// _ = "end of CoverTab[31267]"
//line /usr/local/go/src/html/template/html.go:227
	_go_fuzz_dep_.CoverTab[31268]++
							return b.String()
//line /usr/local/go/src/html/template/html.go:228
	// _ = "end of CoverTab[31268]"
}

// htmlNameFilter accepts valid parts of an HTML attribute or tag name or
//line /usr/local/go/src/html/template/html.go:231
// a known-safe HTML attribute.
//line /usr/local/go/src/html/template/html.go:233
func htmlNameFilter(args ...any) string {
//line /usr/local/go/src/html/template/html.go:233
	_go_fuzz_dep_.CoverTab[31298]++
							s, t := stringify(args...)
							if t == contentTypeHTMLAttr {
//line /usr/local/go/src/html/template/html.go:235
		_go_fuzz_dep_.CoverTab[31303]++
								return s
//line /usr/local/go/src/html/template/html.go:236
		// _ = "end of CoverTab[31303]"
	} else {
//line /usr/local/go/src/html/template/html.go:237
		_go_fuzz_dep_.CoverTab[31304]++
//line /usr/local/go/src/html/template/html.go:237
		// _ = "end of CoverTab[31304]"
//line /usr/local/go/src/html/template/html.go:237
	}
//line /usr/local/go/src/html/template/html.go:237
	// _ = "end of CoverTab[31298]"
//line /usr/local/go/src/html/template/html.go:237
	_go_fuzz_dep_.CoverTab[31299]++
							if len(s) == 0 {
//line /usr/local/go/src/html/template/html.go:238
		_go_fuzz_dep_.CoverTab[31305]++

//line /usr/local/go/src/html/template/html.go:244
		return filterFailsafe
//line /usr/local/go/src/html/template/html.go:244
		// _ = "end of CoverTab[31305]"
	} else {
//line /usr/local/go/src/html/template/html.go:245
		_go_fuzz_dep_.CoverTab[31306]++
//line /usr/local/go/src/html/template/html.go:245
		// _ = "end of CoverTab[31306]"
//line /usr/local/go/src/html/template/html.go:245
	}
//line /usr/local/go/src/html/template/html.go:245
	// _ = "end of CoverTab[31299]"
//line /usr/local/go/src/html/template/html.go:245
	_go_fuzz_dep_.CoverTab[31300]++
							s = strings.ToLower(s)
							if t := attrType(s); t != contentTypePlain {
//line /usr/local/go/src/html/template/html.go:247
		_go_fuzz_dep_.CoverTab[31307]++

								return filterFailsafe
//line /usr/local/go/src/html/template/html.go:249
		// _ = "end of CoverTab[31307]"
	} else {
//line /usr/local/go/src/html/template/html.go:250
		_go_fuzz_dep_.CoverTab[31308]++
//line /usr/local/go/src/html/template/html.go:250
		// _ = "end of CoverTab[31308]"
//line /usr/local/go/src/html/template/html.go:250
	}
//line /usr/local/go/src/html/template/html.go:250
	// _ = "end of CoverTab[31300]"
//line /usr/local/go/src/html/template/html.go:250
	_go_fuzz_dep_.CoverTab[31301]++
							for _, r := range s {
//line /usr/local/go/src/html/template/html.go:251
		_go_fuzz_dep_.CoverTab[31309]++
								switch {
		case '0' <= r && func() bool {
//line /usr/local/go/src/html/template/html.go:253
			_go_fuzz_dep_.CoverTab[31313]++
//line /usr/local/go/src/html/template/html.go:253
			return r <= '9'
//line /usr/local/go/src/html/template/html.go:253
			// _ = "end of CoverTab[31313]"
//line /usr/local/go/src/html/template/html.go:253
		}():
//line /usr/local/go/src/html/template/html.go:253
			_go_fuzz_dep_.CoverTab[31310]++
//line /usr/local/go/src/html/template/html.go:253
			// _ = "end of CoverTab[31310]"
		case 'a' <= r && func() bool {
//line /usr/local/go/src/html/template/html.go:254
			_go_fuzz_dep_.CoverTab[31314]++
//line /usr/local/go/src/html/template/html.go:254
			return r <= 'z'
//line /usr/local/go/src/html/template/html.go:254
			// _ = "end of CoverTab[31314]"
//line /usr/local/go/src/html/template/html.go:254
		}():
//line /usr/local/go/src/html/template/html.go:254
			_go_fuzz_dep_.CoverTab[31311]++
//line /usr/local/go/src/html/template/html.go:254
			// _ = "end of CoverTab[31311]"
		default:
//line /usr/local/go/src/html/template/html.go:255
			_go_fuzz_dep_.CoverTab[31312]++
									return filterFailsafe
//line /usr/local/go/src/html/template/html.go:256
			// _ = "end of CoverTab[31312]"
		}
//line /usr/local/go/src/html/template/html.go:257
		// _ = "end of CoverTab[31309]"
	}
//line /usr/local/go/src/html/template/html.go:258
	// _ = "end of CoverTab[31301]"
//line /usr/local/go/src/html/template/html.go:258
	_go_fuzz_dep_.CoverTab[31302]++
							return s
//line /usr/local/go/src/html/template/html.go:259
	// _ = "end of CoverTab[31302]"
}

// commentEscaper returns the empty string regardless of input.
//line /usr/local/go/src/html/template/html.go:262
// Comment content does not correspond to any parsed structure or
//line /usr/local/go/src/html/template/html.go:262
// human-readable content, so the simplest and most secure policy is to drop
//line /usr/local/go/src/html/template/html.go:262
// content interpolated into comments.
//line /usr/local/go/src/html/template/html.go:262
// This approach is equally valid whether or not static comment content is
//line /usr/local/go/src/html/template/html.go:262
// removed from the template.
//line /usr/local/go/src/html/template/html.go:268
func commentEscaper(args ...any) string {
//line /usr/local/go/src/html/template/html.go:268
	_go_fuzz_dep_.CoverTab[31315]++
							return ""
//line /usr/local/go/src/html/template/html.go:269
	// _ = "end of CoverTab[31315]"
}

//line /usr/local/go/src/html/template/html.go:270
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/html/template/html.go:270
var _ = _go_fuzz_dep_.CoverTab
