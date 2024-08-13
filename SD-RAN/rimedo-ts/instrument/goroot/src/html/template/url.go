// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/html/template/url.go:5
package template

//line /usr/local/go/src/html/template/url.go:5
import (
//line /usr/local/go/src/html/template/url.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/html/template/url.go:5
)
//line /usr/local/go/src/html/template/url.go:5
import (
//line /usr/local/go/src/html/template/url.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/html/template/url.go:5
)

import (
	"fmt"
	"strings"
)

// urlFilter returns its input unless it contains an unsafe scheme in which
//line /usr/local/go/src/html/template/url.go:12
// case it defangs the entire URL.
//line /usr/local/go/src/html/template/url.go:12
//
//line /usr/local/go/src/html/template/url.go:12
// Schemes that cause unintended side effects that are irreversible without user
//line /usr/local/go/src/html/template/url.go:12
// interaction are considered unsafe. For example, clicking on a "javascript:"
//line /usr/local/go/src/html/template/url.go:12
// link can immediately trigger JavaScript code execution.
//line /usr/local/go/src/html/template/url.go:12
//
//line /usr/local/go/src/html/template/url.go:12
// This filter conservatively assumes that all schemes other than the following
//line /usr/local/go/src/html/template/url.go:12
// are unsafe:
//line /usr/local/go/src/html/template/url.go:12
//   - http:   Navigates to a new website, and may open a new window or tab.
//line /usr/local/go/src/html/template/url.go:12
//     These side effects can be reversed by navigating back to the
//line /usr/local/go/src/html/template/url.go:12
//     previous website, or closing the window or tab. No irreversible
//line /usr/local/go/src/html/template/url.go:12
//     changes will take place without further user interaction with
//line /usr/local/go/src/html/template/url.go:12
//     the new website.
//line /usr/local/go/src/html/template/url.go:12
//   - https:  Same as http.
//line /usr/local/go/src/html/template/url.go:12
//   - mailto: Opens an email program and starts a new draft. This side effect
//line /usr/local/go/src/html/template/url.go:12
//     is not irreversible until the user explicitly clicks send; it
//line /usr/local/go/src/html/template/url.go:12
//     can be undone by closing the email program.
//line /usr/local/go/src/html/template/url.go:12
//
//line /usr/local/go/src/html/template/url.go:12
// To allow URLs containing other schemes to bypass this filter, developers must
//line /usr/local/go/src/html/template/url.go:12
// explicitly indicate that such a URL is expected and safe by encapsulating it
//line /usr/local/go/src/html/template/url.go:12
// in a template.URL value.
//line /usr/local/go/src/html/template/url.go:34
func urlFilter(args ...any) string {
//line /usr/local/go/src/html/template/url.go:34
	_go_fuzz_dep_.CoverTab[31809]++
							s, t := stringify(args...)
							if t == contentTypeURL {
//line /usr/local/go/src/html/template/url.go:36
		_go_fuzz_dep_.CoverTab[31812]++
								return s
//line /usr/local/go/src/html/template/url.go:37
		// _ = "end of CoverTab[31812]"
	} else {
//line /usr/local/go/src/html/template/url.go:38
		_go_fuzz_dep_.CoverTab[31813]++
//line /usr/local/go/src/html/template/url.go:38
		// _ = "end of CoverTab[31813]"
//line /usr/local/go/src/html/template/url.go:38
	}
//line /usr/local/go/src/html/template/url.go:38
	// _ = "end of CoverTab[31809]"
//line /usr/local/go/src/html/template/url.go:38
	_go_fuzz_dep_.CoverTab[31810]++
							if !isSafeURL(s) {
//line /usr/local/go/src/html/template/url.go:39
		_go_fuzz_dep_.CoverTab[31814]++
								return "#" + filterFailsafe
//line /usr/local/go/src/html/template/url.go:40
		// _ = "end of CoverTab[31814]"
	} else {
//line /usr/local/go/src/html/template/url.go:41
		_go_fuzz_dep_.CoverTab[31815]++
//line /usr/local/go/src/html/template/url.go:41
		// _ = "end of CoverTab[31815]"
//line /usr/local/go/src/html/template/url.go:41
	}
//line /usr/local/go/src/html/template/url.go:41
	// _ = "end of CoverTab[31810]"
//line /usr/local/go/src/html/template/url.go:41
	_go_fuzz_dep_.CoverTab[31811]++
							return s
//line /usr/local/go/src/html/template/url.go:42
	// _ = "end of CoverTab[31811]"
}

// isSafeURL is true if s is a relative URL or if URL has a protocol in
//line /usr/local/go/src/html/template/url.go:45
// (http, https, mailto).
//line /usr/local/go/src/html/template/url.go:47
func isSafeURL(s string) bool {
//line /usr/local/go/src/html/template/url.go:47
	_go_fuzz_dep_.CoverTab[31816]++
							if protocol, _, ok := strings.Cut(s, ":"); ok && func() bool {
//line /usr/local/go/src/html/template/url.go:48
		_go_fuzz_dep_.CoverTab[31818]++
//line /usr/local/go/src/html/template/url.go:48
		return !strings.Contains(protocol, "/")
//line /usr/local/go/src/html/template/url.go:48
		// _ = "end of CoverTab[31818]"
//line /usr/local/go/src/html/template/url.go:48
	}() {
//line /usr/local/go/src/html/template/url.go:48
		_go_fuzz_dep_.CoverTab[31819]++
								if !strings.EqualFold(protocol, "http") && func() bool {
//line /usr/local/go/src/html/template/url.go:49
			_go_fuzz_dep_.CoverTab[31820]++
//line /usr/local/go/src/html/template/url.go:49
			return !strings.EqualFold(protocol, "https")
//line /usr/local/go/src/html/template/url.go:49
			// _ = "end of CoverTab[31820]"
//line /usr/local/go/src/html/template/url.go:49
		}() && func() bool {
//line /usr/local/go/src/html/template/url.go:49
			_go_fuzz_dep_.CoverTab[31821]++
//line /usr/local/go/src/html/template/url.go:49
			return !strings.EqualFold(protocol, "mailto")
//line /usr/local/go/src/html/template/url.go:49
			// _ = "end of CoverTab[31821]"
//line /usr/local/go/src/html/template/url.go:49
		}() {
//line /usr/local/go/src/html/template/url.go:49
			_go_fuzz_dep_.CoverTab[31822]++
									return false
//line /usr/local/go/src/html/template/url.go:50
			// _ = "end of CoverTab[31822]"
		} else {
//line /usr/local/go/src/html/template/url.go:51
			_go_fuzz_dep_.CoverTab[31823]++
//line /usr/local/go/src/html/template/url.go:51
			// _ = "end of CoverTab[31823]"
//line /usr/local/go/src/html/template/url.go:51
		}
//line /usr/local/go/src/html/template/url.go:51
		// _ = "end of CoverTab[31819]"
	} else {
//line /usr/local/go/src/html/template/url.go:52
		_go_fuzz_dep_.CoverTab[31824]++
//line /usr/local/go/src/html/template/url.go:52
		// _ = "end of CoverTab[31824]"
//line /usr/local/go/src/html/template/url.go:52
	}
//line /usr/local/go/src/html/template/url.go:52
	// _ = "end of CoverTab[31816]"
//line /usr/local/go/src/html/template/url.go:52
	_go_fuzz_dep_.CoverTab[31817]++
							return true
//line /usr/local/go/src/html/template/url.go:53
	// _ = "end of CoverTab[31817]"
}

// urlEscaper produces an output that can be embedded in a URL query.
//line /usr/local/go/src/html/template/url.go:56
// The output can be embedded in an HTML attribute without further escaping.
//line /usr/local/go/src/html/template/url.go:58
func urlEscaper(args ...any) string {
//line /usr/local/go/src/html/template/url.go:58
	_go_fuzz_dep_.CoverTab[31825]++
							return urlProcessor(false, args...)
//line /usr/local/go/src/html/template/url.go:59
	// _ = "end of CoverTab[31825]"
}

// urlNormalizer normalizes URL content so it can be embedded in a quote-delimited
//line /usr/local/go/src/html/template/url.go:62
// string or parenthesis delimited url(...).
//line /usr/local/go/src/html/template/url.go:62
// The normalizer does not encode all HTML specials. Specifically, it does not
//line /usr/local/go/src/html/template/url.go:62
// encode '&' so correct embedding in an HTML attribute requires escaping of
//line /usr/local/go/src/html/template/url.go:62
// '&' to '&amp;'.
//line /usr/local/go/src/html/template/url.go:67
func urlNormalizer(args ...any) string {
//line /usr/local/go/src/html/template/url.go:67
	_go_fuzz_dep_.CoverTab[31826]++
							return urlProcessor(true, args...)
//line /usr/local/go/src/html/template/url.go:68
	// _ = "end of CoverTab[31826]"
}

// urlProcessor normalizes (when norm is true) or escapes its input to produce
//line /usr/local/go/src/html/template/url.go:71
// a valid hierarchical or opaque URL part.
//line /usr/local/go/src/html/template/url.go:73
func urlProcessor(norm bool, args ...any) string {
//line /usr/local/go/src/html/template/url.go:73
	_go_fuzz_dep_.CoverTab[31827]++
							s, t := stringify(args...)
							if t == contentTypeURL {
//line /usr/local/go/src/html/template/url.go:75
		_go_fuzz_dep_.CoverTab[31830]++
								norm = true
//line /usr/local/go/src/html/template/url.go:76
		// _ = "end of CoverTab[31830]"
	} else {
//line /usr/local/go/src/html/template/url.go:77
		_go_fuzz_dep_.CoverTab[31831]++
//line /usr/local/go/src/html/template/url.go:77
		// _ = "end of CoverTab[31831]"
//line /usr/local/go/src/html/template/url.go:77
	}
//line /usr/local/go/src/html/template/url.go:77
	// _ = "end of CoverTab[31827]"
//line /usr/local/go/src/html/template/url.go:77
	_go_fuzz_dep_.CoverTab[31828]++
							var b strings.Builder
							if processURLOnto(s, norm, &b) {
//line /usr/local/go/src/html/template/url.go:79
		_go_fuzz_dep_.CoverTab[31832]++
								return b.String()
//line /usr/local/go/src/html/template/url.go:80
		// _ = "end of CoverTab[31832]"
	} else {
//line /usr/local/go/src/html/template/url.go:81
		_go_fuzz_dep_.CoverTab[31833]++
//line /usr/local/go/src/html/template/url.go:81
		// _ = "end of CoverTab[31833]"
//line /usr/local/go/src/html/template/url.go:81
	}
//line /usr/local/go/src/html/template/url.go:81
	// _ = "end of CoverTab[31828]"
//line /usr/local/go/src/html/template/url.go:81
	_go_fuzz_dep_.CoverTab[31829]++
							return s
//line /usr/local/go/src/html/template/url.go:82
	// _ = "end of CoverTab[31829]"
}

// processURLOnto appends a normalized URL corresponding to its input to b
//line /usr/local/go/src/html/template/url.go:85
// and reports whether the appended content differs from s.
//line /usr/local/go/src/html/template/url.go:87
func processURLOnto(s string, norm bool, b *strings.Builder) bool {
//line /usr/local/go/src/html/template/url.go:87
	_go_fuzz_dep_.CoverTab[31834]++
							b.Grow(len(s) + 16)
							written := 0

//line /usr/local/go/src/html/template/url.go:96
	for i, n := 0, len(s); i < n; i++ {
//line /usr/local/go/src/html/template/url.go:96
		_go_fuzz_dep_.CoverTab[31836]++
								c := s[i]
								switch c {

//line /usr/local/go/src/html/template/url.go:105
		case '!', '#', '$', '&', '*', '+', ',', '/', ':', ';', '=', '?', '@', '[', ']':
//line /usr/local/go/src/html/template/url.go:105
			_go_fuzz_dep_.CoverTab[31838]++
									if norm {
//line /usr/local/go/src/html/template/url.go:106
				_go_fuzz_dep_.CoverTab[31844]++
										continue
//line /usr/local/go/src/html/template/url.go:107
				// _ = "end of CoverTab[31844]"
			} else {
//line /usr/local/go/src/html/template/url.go:108
				_go_fuzz_dep_.CoverTab[31845]++
//line /usr/local/go/src/html/template/url.go:108
				// _ = "end of CoverTab[31845]"
//line /usr/local/go/src/html/template/url.go:108
			}
//line /usr/local/go/src/html/template/url.go:108
			// _ = "end of CoverTab[31838]"

//line /usr/local/go/src/html/template/url.go:114
		case '-', '.', '_', '~':
//line /usr/local/go/src/html/template/url.go:114
			_go_fuzz_dep_.CoverTab[31839]++
									continue
//line /usr/local/go/src/html/template/url.go:115
			// _ = "end of CoverTab[31839]"
		case '%':
//line /usr/local/go/src/html/template/url.go:116
			_go_fuzz_dep_.CoverTab[31840]++

									if norm && func() bool {
//line /usr/local/go/src/html/template/url.go:118
				_go_fuzz_dep_.CoverTab[31846]++
//line /usr/local/go/src/html/template/url.go:118
				return i+2 < len(s)
//line /usr/local/go/src/html/template/url.go:118
				// _ = "end of CoverTab[31846]"
//line /usr/local/go/src/html/template/url.go:118
			}() && func() bool {
//line /usr/local/go/src/html/template/url.go:118
				_go_fuzz_dep_.CoverTab[31847]++
//line /usr/local/go/src/html/template/url.go:118
				return isHex(s[i+1])
//line /usr/local/go/src/html/template/url.go:118
				// _ = "end of CoverTab[31847]"
//line /usr/local/go/src/html/template/url.go:118
			}() && func() bool {
//line /usr/local/go/src/html/template/url.go:118
				_go_fuzz_dep_.CoverTab[31848]++
//line /usr/local/go/src/html/template/url.go:118
				return isHex(s[i+2])
//line /usr/local/go/src/html/template/url.go:118
				// _ = "end of CoverTab[31848]"
//line /usr/local/go/src/html/template/url.go:118
			}() {
//line /usr/local/go/src/html/template/url.go:118
				_go_fuzz_dep_.CoverTab[31849]++
										continue
//line /usr/local/go/src/html/template/url.go:119
				// _ = "end of CoverTab[31849]"
			} else {
//line /usr/local/go/src/html/template/url.go:120
				_go_fuzz_dep_.CoverTab[31850]++
//line /usr/local/go/src/html/template/url.go:120
				// _ = "end of CoverTab[31850]"
//line /usr/local/go/src/html/template/url.go:120
			}
//line /usr/local/go/src/html/template/url.go:120
			// _ = "end of CoverTab[31840]"
		default:
//line /usr/local/go/src/html/template/url.go:121
			_go_fuzz_dep_.CoverTab[31841]++

									if 'a' <= c && func() bool {
//line /usr/local/go/src/html/template/url.go:123
				_go_fuzz_dep_.CoverTab[31851]++
//line /usr/local/go/src/html/template/url.go:123
				return c <= 'z'
//line /usr/local/go/src/html/template/url.go:123
				// _ = "end of CoverTab[31851]"
//line /usr/local/go/src/html/template/url.go:123
			}() {
//line /usr/local/go/src/html/template/url.go:123
				_go_fuzz_dep_.CoverTab[31852]++
										continue
//line /usr/local/go/src/html/template/url.go:124
				// _ = "end of CoverTab[31852]"
			} else {
//line /usr/local/go/src/html/template/url.go:125
				_go_fuzz_dep_.CoverTab[31853]++
//line /usr/local/go/src/html/template/url.go:125
				// _ = "end of CoverTab[31853]"
//line /usr/local/go/src/html/template/url.go:125
			}
//line /usr/local/go/src/html/template/url.go:125
			// _ = "end of CoverTab[31841]"
//line /usr/local/go/src/html/template/url.go:125
			_go_fuzz_dep_.CoverTab[31842]++
									if 'A' <= c && func() bool {
//line /usr/local/go/src/html/template/url.go:126
				_go_fuzz_dep_.CoverTab[31854]++
//line /usr/local/go/src/html/template/url.go:126
				return c <= 'Z'
//line /usr/local/go/src/html/template/url.go:126
				// _ = "end of CoverTab[31854]"
//line /usr/local/go/src/html/template/url.go:126
			}() {
//line /usr/local/go/src/html/template/url.go:126
				_go_fuzz_dep_.CoverTab[31855]++
										continue
//line /usr/local/go/src/html/template/url.go:127
				// _ = "end of CoverTab[31855]"
			} else {
//line /usr/local/go/src/html/template/url.go:128
				_go_fuzz_dep_.CoverTab[31856]++
//line /usr/local/go/src/html/template/url.go:128
				// _ = "end of CoverTab[31856]"
//line /usr/local/go/src/html/template/url.go:128
			}
//line /usr/local/go/src/html/template/url.go:128
			// _ = "end of CoverTab[31842]"
//line /usr/local/go/src/html/template/url.go:128
			_go_fuzz_dep_.CoverTab[31843]++
									if '0' <= c && func() bool {
//line /usr/local/go/src/html/template/url.go:129
				_go_fuzz_dep_.CoverTab[31857]++
//line /usr/local/go/src/html/template/url.go:129
				return c <= '9'
//line /usr/local/go/src/html/template/url.go:129
				// _ = "end of CoverTab[31857]"
//line /usr/local/go/src/html/template/url.go:129
			}() {
//line /usr/local/go/src/html/template/url.go:129
				_go_fuzz_dep_.CoverTab[31858]++
										continue
//line /usr/local/go/src/html/template/url.go:130
				// _ = "end of CoverTab[31858]"
			} else {
//line /usr/local/go/src/html/template/url.go:131
				_go_fuzz_dep_.CoverTab[31859]++
//line /usr/local/go/src/html/template/url.go:131
				// _ = "end of CoverTab[31859]"
//line /usr/local/go/src/html/template/url.go:131
			}
//line /usr/local/go/src/html/template/url.go:131
			// _ = "end of CoverTab[31843]"
		}
//line /usr/local/go/src/html/template/url.go:132
		// _ = "end of CoverTab[31836]"
//line /usr/local/go/src/html/template/url.go:132
		_go_fuzz_dep_.CoverTab[31837]++
								b.WriteString(s[written:i])
								fmt.Fprintf(b, "%%%02x", c)
								written = i + 1
//line /usr/local/go/src/html/template/url.go:135
		// _ = "end of CoverTab[31837]"
	}
//line /usr/local/go/src/html/template/url.go:136
	// _ = "end of CoverTab[31834]"
//line /usr/local/go/src/html/template/url.go:136
	_go_fuzz_dep_.CoverTab[31835]++
							b.WriteString(s[written:])
							return written != 0
//line /usr/local/go/src/html/template/url.go:138
	// _ = "end of CoverTab[31835]"
}

// Filters and normalizes srcset values which are comma separated
//line /usr/local/go/src/html/template/url.go:141
// URLs followed by metadata.
//line /usr/local/go/src/html/template/url.go:143
func srcsetFilterAndEscaper(args ...any) string {
//line /usr/local/go/src/html/template/url.go:143
	_go_fuzz_dep_.CoverTab[31860]++
							s, t := stringify(args...)
							switch t {
	case contentTypeSrcset:
//line /usr/local/go/src/html/template/url.go:146
		_go_fuzz_dep_.CoverTab[31863]++
								return s
//line /usr/local/go/src/html/template/url.go:147
		// _ = "end of CoverTab[31863]"
	case contentTypeURL:
//line /usr/local/go/src/html/template/url.go:148
		_go_fuzz_dep_.CoverTab[31864]++
		// Normalizing gets rid of all HTML whitespace
		// which separate the image URL from its metadata.
		var b strings.Builder
		if processURLOnto(s, true, &b) {
//line /usr/local/go/src/html/template/url.go:152
			_go_fuzz_dep_.CoverTab[31867]++
									s = b.String()
//line /usr/local/go/src/html/template/url.go:153
			// _ = "end of CoverTab[31867]"
		} else {
//line /usr/local/go/src/html/template/url.go:154
			_go_fuzz_dep_.CoverTab[31868]++
//line /usr/local/go/src/html/template/url.go:154
			// _ = "end of CoverTab[31868]"
//line /usr/local/go/src/html/template/url.go:154
		}
//line /usr/local/go/src/html/template/url.go:154
		// _ = "end of CoverTab[31864]"
//line /usr/local/go/src/html/template/url.go:154
		_go_fuzz_dep_.CoverTab[31865]++

								return strings.ReplaceAll(s, ",", "%2c")
//line /usr/local/go/src/html/template/url.go:156
		// _ = "end of CoverTab[31865]"
//line /usr/local/go/src/html/template/url.go:156
	default:
//line /usr/local/go/src/html/template/url.go:156
		_go_fuzz_dep_.CoverTab[31866]++
//line /usr/local/go/src/html/template/url.go:156
		// _ = "end of CoverTab[31866]"
	}
//line /usr/local/go/src/html/template/url.go:157
	// _ = "end of CoverTab[31860]"
//line /usr/local/go/src/html/template/url.go:157
	_go_fuzz_dep_.CoverTab[31861]++

							var b strings.Builder
							written := 0
							for i := 0; i < len(s); i++ {
//line /usr/local/go/src/html/template/url.go:161
		_go_fuzz_dep_.CoverTab[31869]++
								if s[i] == ',' {
//line /usr/local/go/src/html/template/url.go:162
			_go_fuzz_dep_.CoverTab[31870]++
									filterSrcsetElement(s, written, i, &b)
									b.WriteString(",")
									written = i + 1
//line /usr/local/go/src/html/template/url.go:165
			// _ = "end of CoverTab[31870]"
		} else {
//line /usr/local/go/src/html/template/url.go:166
			_go_fuzz_dep_.CoverTab[31871]++
//line /usr/local/go/src/html/template/url.go:166
			// _ = "end of CoverTab[31871]"
//line /usr/local/go/src/html/template/url.go:166
		}
//line /usr/local/go/src/html/template/url.go:166
		// _ = "end of CoverTab[31869]"
	}
//line /usr/local/go/src/html/template/url.go:167
	// _ = "end of CoverTab[31861]"
//line /usr/local/go/src/html/template/url.go:167
	_go_fuzz_dep_.CoverTab[31862]++
							filterSrcsetElement(s, written, len(s), &b)
							return b.String()
//line /usr/local/go/src/html/template/url.go:169
	// _ = "end of CoverTab[31862]"
}

// Derived from https://play.golang.org/p/Dhmj7FORT5
const htmlSpaceAndASCIIAlnumBytes = "\x00\x36\x00\x00\x01\x00\xff\x03\xfe\xff\xff\x07\xfe\xff\xff\x07"

// isHTMLSpace is true iff c is a whitespace character per
//line /usr/local/go/src/html/template/url.go:175
// https://infra.spec.whatwg.org/#ascii-whitespace
//line /usr/local/go/src/html/template/url.go:177
func isHTMLSpace(c byte) bool {
//line /usr/local/go/src/html/template/url.go:177
	_go_fuzz_dep_.CoverTab[31872]++
							return (c <= 0x20) && func() bool {
//line /usr/local/go/src/html/template/url.go:178
		_go_fuzz_dep_.CoverTab[31873]++
//line /usr/local/go/src/html/template/url.go:178
		return 0 != (htmlSpaceAndASCIIAlnumBytes[c>>3] & (1 << uint(c&0x7)))
//line /usr/local/go/src/html/template/url.go:178
		// _ = "end of CoverTab[31873]"
//line /usr/local/go/src/html/template/url.go:178
	}()
//line /usr/local/go/src/html/template/url.go:178
	// _ = "end of CoverTab[31872]"
}

func isHTMLSpaceOrASCIIAlnum(c byte) bool {
//line /usr/local/go/src/html/template/url.go:181
	_go_fuzz_dep_.CoverTab[31874]++
							return (c < 0x80) && func() bool {
//line /usr/local/go/src/html/template/url.go:182
		_go_fuzz_dep_.CoverTab[31875]++
//line /usr/local/go/src/html/template/url.go:182
		return 0 != (htmlSpaceAndASCIIAlnumBytes[c>>3] & (1 << uint(c&0x7)))
//line /usr/local/go/src/html/template/url.go:182
		// _ = "end of CoverTab[31875]"
//line /usr/local/go/src/html/template/url.go:182
	}()
//line /usr/local/go/src/html/template/url.go:182
	// _ = "end of CoverTab[31874]"
}

func filterSrcsetElement(s string, left int, right int, b *strings.Builder) {
//line /usr/local/go/src/html/template/url.go:185
	_go_fuzz_dep_.CoverTab[31876]++
							start := left
							for start < right && func() bool {
//line /usr/local/go/src/html/template/url.go:187
		_go_fuzz_dep_.CoverTab[31880]++
//line /usr/local/go/src/html/template/url.go:187
		return isHTMLSpace(s[start])
//line /usr/local/go/src/html/template/url.go:187
		// _ = "end of CoverTab[31880]"
//line /usr/local/go/src/html/template/url.go:187
	}() {
//line /usr/local/go/src/html/template/url.go:187
		_go_fuzz_dep_.CoverTab[31881]++
								start++
//line /usr/local/go/src/html/template/url.go:188
		// _ = "end of CoverTab[31881]"
	}
//line /usr/local/go/src/html/template/url.go:189
	// _ = "end of CoverTab[31876]"
//line /usr/local/go/src/html/template/url.go:189
	_go_fuzz_dep_.CoverTab[31877]++
							end := right
							for i := start; i < right; i++ {
//line /usr/local/go/src/html/template/url.go:191
		_go_fuzz_dep_.CoverTab[31882]++
								if isHTMLSpace(s[i]) {
//line /usr/local/go/src/html/template/url.go:192
			_go_fuzz_dep_.CoverTab[31883]++
									end = i
									break
//line /usr/local/go/src/html/template/url.go:194
			// _ = "end of CoverTab[31883]"
		} else {
//line /usr/local/go/src/html/template/url.go:195
			_go_fuzz_dep_.CoverTab[31884]++
//line /usr/local/go/src/html/template/url.go:195
			// _ = "end of CoverTab[31884]"
//line /usr/local/go/src/html/template/url.go:195
		}
//line /usr/local/go/src/html/template/url.go:195
		// _ = "end of CoverTab[31882]"
	}
//line /usr/local/go/src/html/template/url.go:196
	// _ = "end of CoverTab[31877]"
//line /usr/local/go/src/html/template/url.go:196
	_go_fuzz_dep_.CoverTab[31878]++
							if url := s[start:end]; isSafeURL(url) {
//line /usr/local/go/src/html/template/url.go:197
		_go_fuzz_dep_.CoverTab[31885]++

//line /usr/local/go/src/html/template/url.go:200
		metadataOk := true
		for i := end; i < right; i++ {
//line /usr/local/go/src/html/template/url.go:201
			_go_fuzz_dep_.CoverTab[31887]++
									if !isHTMLSpaceOrASCIIAlnum(s[i]) {
//line /usr/local/go/src/html/template/url.go:202
				_go_fuzz_dep_.CoverTab[31888]++
										metadataOk = false
										break
//line /usr/local/go/src/html/template/url.go:204
				// _ = "end of CoverTab[31888]"
			} else {
//line /usr/local/go/src/html/template/url.go:205
				_go_fuzz_dep_.CoverTab[31889]++
//line /usr/local/go/src/html/template/url.go:205
				// _ = "end of CoverTab[31889]"
//line /usr/local/go/src/html/template/url.go:205
			}
//line /usr/local/go/src/html/template/url.go:205
			// _ = "end of CoverTab[31887]"
		}
//line /usr/local/go/src/html/template/url.go:206
		// _ = "end of CoverTab[31885]"
//line /usr/local/go/src/html/template/url.go:206
		_go_fuzz_dep_.CoverTab[31886]++
								if metadataOk {
//line /usr/local/go/src/html/template/url.go:207
			_go_fuzz_dep_.CoverTab[31890]++
									b.WriteString(s[left:start])
									processURLOnto(url, true, b)
									b.WriteString(s[end:right])
									return
//line /usr/local/go/src/html/template/url.go:211
			// _ = "end of CoverTab[31890]"
		} else {
//line /usr/local/go/src/html/template/url.go:212
			_go_fuzz_dep_.CoverTab[31891]++
//line /usr/local/go/src/html/template/url.go:212
			// _ = "end of CoverTab[31891]"
//line /usr/local/go/src/html/template/url.go:212
		}
//line /usr/local/go/src/html/template/url.go:212
		// _ = "end of CoverTab[31886]"
	} else {
//line /usr/local/go/src/html/template/url.go:213
		_go_fuzz_dep_.CoverTab[31892]++
//line /usr/local/go/src/html/template/url.go:213
		// _ = "end of CoverTab[31892]"
//line /usr/local/go/src/html/template/url.go:213
	}
//line /usr/local/go/src/html/template/url.go:213
	// _ = "end of CoverTab[31878]"
//line /usr/local/go/src/html/template/url.go:213
	_go_fuzz_dep_.CoverTab[31879]++
							b.WriteString("#")
							b.WriteString(filterFailsafe)
//line /usr/local/go/src/html/template/url.go:215
	// _ = "end of CoverTab[31879]"
}

//line /usr/local/go/src/html/template/url.go:216
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/html/template/url.go:216
var _ = _go_fuzz_dep_.CoverTab
