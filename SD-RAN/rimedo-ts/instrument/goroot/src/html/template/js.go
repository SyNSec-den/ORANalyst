// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/html/template/js.go:5
package template

//line /usr/local/go/src/html/template/js.go:5
import (
//line /usr/local/go/src/html/template/js.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/html/template/js.go:5
)
//line /usr/local/go/src/html/template/js.go:5
import (
//line /usr/local/go/src/html/template/js.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/html/template/js.go:5
)

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"unicode/utf8"
)

// jsWhitespace contains all of the JS whitespace characters, as defined
//line /usr/local/go/src/html/template/js.go:16
// by the \s character class.
//line /usr/local/go/src/html/template/js.go:16
// See https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Regular_expressions/Character_classes.
//line /usr/local/go/src/html/template/js.go:19
const jsWhitespace = "\f\n\r\t\v\u0020\u00a0\u1680\u2000\u2001\u2002\u2003\u2004\u2005\u2006\u2007\u2008\u2009\u200a\u2028\u2029\u202f\u205f\u3000\ufeff"

// nextJSCtx returns the context that determines whether a slash after the
//line /usr/local/go/src/html/template/js.go:21
// given run of tokens starts a regular expression instead of a division
//line /usr/local/go/src/html/template/js.go:21
// operator: / or /=.
//line /usr/local/go/src/html/template/js.go:21
//
//line /usr/local/go/src/html/template/js.go:21
// This assumes that the token run does not include any string tokens, comment
//line /usr/local/go/src/html/template/js.go:21
// tokens, regular expression literal tokens, or division operators.
//line /usr/local/go/src/html/template/js.go:21
//
//line /usr/local/go/src/html/template/js.go:21
// This fails on some valid but nonsensical JavaScript programs like
//line /usr/local/go/src/html/template/js.go:21
// "x = ++/foo/i" which is quite different than "x++/foo/i", but is not known to
//line /usr/local/go/src/html/template/js.go:21
// fail on any known useful programs. It is based on the draft
//line /usr/local/go/src/html/template/js.go:21
// JavaScript 2.0 lexical grammar and requires one token of lookbehind:
//line /usr/local/go/src/html/template/js.go:21
// https://www.mozilla.org/js/language/js20-2000-07/rationale/syntax.html
//line /usr/local/go/src/html/template/js.go:33
func nextJSCtx(s []byte, preceding jsCtx) jsCtx {
//line /usr/local/go/src/html/template/js.go:33
	_go_fuzz_dep_.CoverTab[31316]++

							s = bytes.TrimRight(s, jsWhitespace)
							if len(s) == 0 {
//line /usr/local/go/src/html/template/js.go:36
		_go_fuzz_dep_.CoverTab[31319]++
								return preceding
//line /usr/local/go/src/html/template/js.go:37
		// _ = "end of CoverTab[31319]"
	} else {
//line /usr/local/go/src/html/template/js.go:38
		_go_fuzz_dep_.CoverTab[31320]++
//line /usr/local/go/src/html/template/js.go:38
		// _ = "end of CoverTab[31320]"
//line /usr/local/go/src/html/template/js.go:38
	}
//line /usr/local/go/src/html/template/js.go:38
	// _ = "end of CoverTab[31316]"
//line /usr/local/go/src/html/template/js.go:38
	_go_fuzz_dep_.CoverTab[31317]++

//line /usr/local/go/src/html/template/js.go:41
	switch c, n := s[len(s)-1], len(s); c {
	case '+', '-':
//line /usr/local/go/src/html/template/js.go:42
		_go_fuzz_dep_.CoverTab[31321]++

//line /usr/local/go/src/html/template/js.go:45
		start := n - 1

		for start > 0 && func() bool {
//line /usr/local/go/src/html/template/js.go:47
			_go_fuzz_dep_.CoverTab[31333]++
//line /usr/local/go/src/html/template/js.go:47
			return s[start-1] == c
//line /usr/local/go/src/html/template/js.go:47
			// _ = "end of CoverTab[31333]"
//line /usr/local/go/src/html/template/js.go:47
		}() {
//line /usr/local/go/src/html/template/js.go:47
			_go_fuzz_dep_.CoverTab[31334]++
									start--
//line /usr/local/go/src/html/template/js.go:48
			// _ = "end of CoverTab[31334]"
		}
//line /usr/local/go/src/html/template/js.go:49
		// _ = "end of CoverTab[31321]"
//line /usr/local/go/src/html/template/js.go:49
		_go_fuzz_dep_.CoverTab[31322]++
								if (n-start)&1 == 1 {
//line /usr/local/go/src/html/template/js.go:50
			_go_fuzz_dep_.CoverTab[31335]++

//line /usr/local/go/src/html/template/js.go:53
			return jsCtxRegexp
//line /usr/local/go/src/html/template/js.go:53
			// _ = "end of CoverTab[31335]"
		} else {
//line /usr/local/go/src/html/template/js.go:54
			_go_fuzz_dep_.CoverTab[31336]++
//line /usr/local/go/src/html/template/js.go:54
			// _ = "end of CoverTab[31336]"
//line /usr/local/go/src/html/template/js.go:54
		}
//line /usr/local/go/src/html/template/js.go:54
		// _ = "end of CoverTab[31322]"
//line /usr/local/go/src/html/template/js.go:54
		_go_fuzz_dep_.CoverTab[31323]++
								return jsCtxDivOp
//line /usr/local/go/src/html/template/js.go:55
		// _ = "end of CoverTab[31323]"
	case '.':
//line /usr/local/go/src/html/template/js.go:56
		_go_fuzz_dep_.CoverTab[31324]++

								if n != 1 && func() bool {
//line /usr/local/go/src/html/template/js.go:58
			_go_fuzz_dep_.CoverTab[31337]++
//line /usr/local/go/src/html/template/js.go:58
			return '0' <= s[n-2]
//line /usr/local/go/src/html/template/js.go:58
			// _ = "end of CoverTab[31337]"
//line /usr/local/go/src/html/template/js.go:58
		}() && func() bool {
//line /usr/local/go/src/html/template/js.go:58
			_go_fuzz_dep_.CoverTab[31338]++
//line /usr/local/go/src/html/template/js.go:58
			return s[n-2] <= '9'
//line /usr/local/go/src/html/template/js.go:58
			// _ = "end of CoverTab[31338]"
//line /usr/local/go/src/html/template/js.go:58
		}() {
//line /usr/local/go/src/html/template/js.go:58
			_go_fuzz_dep_.CoverTab[31339]++
									return jsCtxDivOp
//line /usr/local/go/src/html/template/js.go:59
			// _ = "end of CoverTab[31339]"
		} else {
//line /usr/local/go/src/html/template/js.go:60
			_go_fuzz_dep_.CoverTab[31340]++
//line /usr/local/go/src/html/template/js.go:60
			// _ = "end of CoverTab[31340]"
//line /usr/local/go/src/html/template/js.go:60
		}
//line /usr/local/go/src/html/template/js.go:60
		// _ = "end of CoverTab[31324]"
//line /usr/local/go/src/html/template/js.go:60
		_go_fuzz_dep_.CoverTab[31325]++
								return jsCtxRegexp
//line /usr/local/go/src/html/template/js.go:61
		// _ = "end of CoverTab[31325]"

//line /usr/local/go/src/html/template/js.go:64
	case ',', '<', '>', '=', '*', '%', '&', '|', '^', '?':
//line /usr/local/go/src/html/template/js.go:64
		_go_fuzz_dep_.CoverTab[31326]++
								return jsCtxRegexp
//line /usr/local/go/src/html/template/js.go:65
		// _ = "end of CoverTab[31326]"

//line /usr/local/go/src/html/template/js.go:68
	case '!', '~':
//line /usr/local/go/src/html/template/js.go:68
		_go_fuzz_dep_.CoverTab[31327]++
								return jsCtxRegexp
//line /usr/local/go/src/html/template/js.go:69
		// _ = "end of CoverTab[31327]"

//line /usr/local/go/src/html/template/js.go:72
	case '(', '[':
//line /usr/local/go/src/html/template/js.go:72
		_go_fuzz_dep_.CoverTab[31328]++
								return jsCtxRegexp
//line /usr/local/go/src/html/template/js.go:73
		// _ = "end of CoverTab[31328]"

//line /usr/local/go/src/html/template/js.go:76
	case ':', ';', '{':
//line /usr/local/go/src/html/template/js.go:76
		_go_fuzz_dep_.CoverTab[31329]++
								return jsCtxRegexp
//line /usr/local/go/src/html/template/js.go:77
		// _ = "end of CoverTab[31329]"

//line /usr/local/go/src/html/template/js.go:89
	case '}':
//line /usr/local/go/src/html/template/js.go:89
		_go_fuzz_dep_.CoverTab[31330]++
								return jsCtxRegexp
//line /usr/local/go/src/html/template/js.go:90
		// _ = "end of CoverTab[31330]"
	default:
//line /usr/local/go/src/html/template/js.go:91
		_go_fuzz_dep_.CoverTab[31331]++

//line /usr/local/go/src/html/template/js.go:94
		j := n
		for j > 0 && func() bool {
//line /usr/local/go/src/html/template/js.go:95
			_go_fuzz_dep_.CoverTab[31341]++
//line /usr/local/go/src/html/template/js.go:95
			return isJSIdentPart(rune(s[j-1]))
//line /usr/local/go/src/html/template/js.go:95
			// _ = "end of CoverTab[31341]"
//line /usr/local/go/src/html/template/js.go:95
		}() {
//line /usr/local/go/src/html/template/js.go:95
			_go_fuzz_dep_.CoverTab[31342]++
									j--
//line /usr/local/go/src/html/template/js.go:96
			// _ = "end of CoverTab[31342]"
		}
//line /usr/local/go/src/html/template/js.go:97
		// _ = "end of CoverTab[31331]"
//line /usr/local/go/src/html/template/js.go:97
		_go_fuzz_dep_.CoverTab[31332]++
								if regexpPrecederKeywords[string(s[j:])] {
//line /usr/local/go/src/html/template/js.go:98
			_go_fuzz_dep_.CoverTab[31343]++
									return jsCtxRegexp
//line /usr/local/go/src/html/template/js.go:99
			// _ = "end of CoverTab[31343]"
		} else {
//line /usr/local/go/src/html/template/js.go:100
			_go_fuzz_dep_.CoverTab[31344]++
//line /usr/local/go/src/html/template/js.go:100
			// _ = "end of CoverTab[31344]"
//line /usr/local/go/src/html/template/js.go:100
		}
//line /usr/local/go/src/html/template/js.go:100
		// _ = "end of CoverTab[31332]"
	}
//line /usr/local/go/src/html/template/js.go:101
	// _ = "end of CoverTab[31317]"
//line /usr/local/go/src/html/template/js.go:101
	_go_fuzz_dep_.CoverTab[31318]++

//line /usr/local/go/src/html/template/js.go:105
	return jsCtxDivOp
//line /usr/local/go/src/html/template/js.go:105
	// _ = "end of CoverTab[31318]"
}

// regexpPrecederKeywords is a set of reserved JS keywords that can precede a
//line /usr/local/go/src/html/template/js.go:108
// regular expression in JS source.
//line /usr/local/go/src/html/template/js.go:110
var regexpPrecederKeywords = map[string]bool{
	"break":	true,
	"case":		true,
	"continue":	true,
	"delete":	true,
	"do":		true,
	"else":		true,
	"finally":	true,
	"in":		true,
	"instanceof":	true,
	"return":	true,
	"throw":	true,
	"try":		true,
	"typeof":	true,
	"void":		true,
}

var jsonMarshalType = reflect.TypeOf((*json.Marshaler)(nil)).Elem()

// indirectToJSONMarshaler returns the value, after dereferencing as many times
//line /usr/local/go/src/html/template/js.go:129
// as necessary to reach the base type (or nil) or an implementation of json.Marshal.
//line /usr/local/go/src/html/template/js.go:131
func indirectToJSONMarshaler(a any) any {
//line /usr/local/go/src/html/template/js.go:131
	_go_fuzz_dep_.CoverTab[31345]++

//line /usr/local/go/src/html/template/js.go:136
	if a == nil {
//line /usr/local/go/src/html/template/js.go:136
		_go_fuzz_dep_.CoverTab[31348]++
								return nil
//line /usr/local/go/src/html/template/js.go:137
		// _ = "end of CoverTab[31348]"
	} else {
//line /usr/local/go/src/html/template/js.go:138
		_go_fuzz_dep_.CoverTab[31349]++
//line /usr/local/go/src/html/template/js.go:138
		// _ = "end of CoverTab[31349]"
//line /usr/local/go/src/html/template/js.go:138
	}
//line /usr/local/go/src/html/template/js.go:138
	// _ = "end of CoverTab[31345]"
//line /usr/local/go/src/html/template/js.go:138
	_go_fuzz_dep_.CoverTab[31346]++

							v := reflect.ValueOf(a)
							for !v.Type().Implements(jsonMarshalType) && func() bool {
//line /usr/local/go/src/html/template/js.go:141
		_go_fuzz_dep_.CoverTab[31350]++
//line /usr/local/go/src/html/template/js.go:141
		return v.Kind() == reflect.Pointer
//line /usr/local/go/src/html/template/js.go:141
		// _ = "end of CoverTab[31350]"
//line /usr/local/go/src/html/template/js.go:141
	}() && func() bool {
//line /usr/local/go/src/html/template/js.go:141
		_go_fuzz_dep_.CoverTab[31351]++
//line /usr/local/go/src/html/template/js.go:141
		return !v.IsNil()
//line /usr/local/go/src/html/template/js.go:141
		// _ = "end of CoverTab[31351]"
//line /usr/local/go/src/html/template/js.go:141
	}() {
//line /usr/local/go/src/html/template/js.go:141
		_go_fuzz_dep_.CoverTab[31352]++
								v = v.Elem()
//line /usr/local/go/src/html/template/js.go:142
		// _ = "end of CoverTab[31352]"
	}
//line /usr/local/go/src/html/template/js.go:143
	// _ = "end of CoverTab[31346]"
//line /usr/local/go/src/html/template/js.go:143
	_go_fuzz_dep_.CoverTab[31347]++
							return v.Interface()
//line /usr/local/go/src/html/template/js.go:144
	// _ = "end of CoverTab[31347]"
}

// jsValEscaper escapes its inputs to a JS Expression (section 11.14) that has
//line /usr/local/go/src/html/template/js.go:147
// neither side-effects nor free variables outside (NaN, Infinity).
//line /usr/local/go/src/html/template/js.go:149
func jsValEscaper(args ...any) string {
//line /usr/local/go/src/html/template/js.go:149
	_go_fuzz_dep_.CoverTab[31353]++
							var a any
							if len(args) == 1 {
//line /usr/local/go/src/html/template/js.go:151
		_go_fuzz_dep_.CoverTab[31360]++
								a = indirectToJSONMarshaler(args[0])
								switch t := a.(type) {
		case JS:
//line /usr/local/go/src/html/template/js.go:154
			_go_fuzz_dep_.CoverTab[31361]++
									return string(t)
//line /usr/local/go/src/html/template/js.go:155
			// _ = "end of CoverTab[31361]"
		case JSStr:
//line /usr/local/go/src/html/template/js.go:156
			_go_fuzz_dep_.CoverTab[31362]++

									return `"` + string(t) + `"`
//line /usr/local/go/src/html/template/js.go:158
			// _ = "end of CoverTab[31362]"
		case json.Marshaler:
//line /usr/local/go/src/html/template/js.go:159
			_go_fuzz_dep_.CoverTab[31363]++
//line /usr/local/go/src/html/template/js.go:159
			// _ = "end of CoverTab[31363]"

		case fmt.Stringer:
//line /usr/local/go/src/html/template/js.go:161
			_go_fuzz_dep_.CoverTab[31364]++
									a = t.String()
//line /usr/local/go/src/html/template/js.go:162
			// _ = "end of CoverTab[31364]"
		}
//line /usr/local/go/src/html/template/js.go:163
		// _ = "end of CoverTab[31360]"
	} else {
//line /usr/local/go/src/html/template/js.go:164
		_go_fuzz_dep_.CoverTab[31365]++
								for i, arg := range args {
//line /usr/local/go/src/html/template/js.go:165
			_go_fuzz_dep_.CoverTab[31367]++
									args[i] = indirectToJSONMarshaler(arg)
//line /usr/local/go/src/html/template/js.go:166
			// _ = "end of CoverTab[31367]"
		}
//line /usr/local/go/src/html/template/js.go:167
		// _ = "end of CoverTab[31365]"
//line /usr/local/go/src/html/template/js.go:167
		_go_fuzz_dep_.CoverTab[31366]++
								a = fmt.Sprint(args...)
//line /usr/local/go/src/html/template/js.go:168
		// _ = "end of CoverTab[31366]"
	}
//line /usr/local/go/src/html/template/js.go:169
	// _ = "end of CoverTab[31353]"
//line /usr/local/go/src/html/template/js.go:169
	_go_fuzz_dep_.CoverTab[31354]++

//line /usr/local/go/src/html/template/js.go:172
	b, err := json.Marshal(a)
	if err != nil {
//line /usr/local/go/src/html/template/js.go:173
		_go_fuzz_dep_.CoverTab[31368]++

//line /usr/local/go/src/html/template/js.go:180
		return fmt.Sprintf(" /* %s */null ", strings.ReplaceAll(err.Error(), "*/", "* /"))
//line /usr/local/go/src/html/template/js.go:180
		// _ = "end of CoverTab[31368]"
	} else {
//line /usr/local/go/src/html/template/js.go:181
		_go_fuzz_dep_.CoverTab[31369]++
//line /usr/local/go/src/html/template/js.go:181
		// _ = "end of CoverTab[31369]"
//line /usr/local/go/src/html/template/js.go:181
	}
//line /usr/local/go/src/html/template/js.go:181
	// _ = "end of CoverTab[31354]"
//line /usr/local/go/src/html/template/js.go:181
	_go_fuzz_dep_.CoverTab[31355]++

//line /usr/local/go/src/html/template/js.go:188
	if len(b) == 0 {
//line /usr/local/go/src/html/template/js.go:188
		_go_fuzz_dep_.CoverTab[31370]++

//line /usr/local/go/src/html/template/js.go:191
		return " null "
//line /usr/local/go/src/html/template/js.go:191
		// _ = "end of CoverTab[31370]"
	} else {
//line /usr/local/go/src/html/template/js.go:192
		_go_fuzz_dep_.CoverTab[31371]++
//line /usr/local/go/src/html/template/js.go:192
		// _ = "end of CoverTab[31371]"
//line /usr/local/go/src/html/template/js.go:192
	}
//line /usr/local/go/src/html/template/js.go:192
	// _ = "end of CoverTab[31355]"
//line /usr/local/go/src/html/template/js.go:192
	_go_fuzz_dep_.CoverTab[31356]++
							first, _ := utf8.DecodeRune(b)
							last, _ := utf8.DecodeLastRune(b)
							var buf strings.Builder

//line /usr/local/go/src/html/template/js.go:198
	pad := isJSIdentPart(first) || func() bool {
//line /usr/local/go/src/html/template/js.go:198
		_go_fuzz_dep_.CoverTab[31372]++
//line /usr/local/go/src/html/template/js.go:198
		return isJSIdentPart(last)
//line /usr/local/go/src/html/template/js.go:198
		// _ = "end of CoverTab[31372]"
//line /usr/local/go/src/html/template/js.go:198
	}()
	if pad {
//line /usr/local/go/src/html/template/js.go:199
		_go_fuzz_dep_.CoverTab[31373]++
								buf.WriteByte(' ')
//line /usr/local/go/src/html/template/js.go:200
		// _ = "end of CoverTab[31373]"
	} else {
//line /usr/local/go/src/html/template/js.go:201
		_go_fuzz_dep_.CoverTab[31374]++
//line /usr/local/go/src/html/template/js.go:201
		// _ = "end of CoverTab[31374]"
//line /usr/local/go/src/html/template/js.go:201
	}
//line /usr/local/go/src/html/template/js.go:201
	// _ = "end of CoverTab[31356]"
//line /usr/local/go/src/html/template/js.go:201
	_go_fuzz_dep_.CoverTab[31357]++
							written := 0

//line /usr/local/go/src/html/template/js.go:205
	for i := 0; i < len(b); {
//line /usr/local/go/src/html/template/js.go:205
		_go_fuzz_dep_.CoverTab[31375]++
								rune, n := utf8.DecodeRune(b[i:])
								repl := ""
								if rune == 0x2028 {
//line /usr/local/go/src/html/template/js.go:208
			_go_fuzz_dep_.CoverTab[31378]++
									repl = `\u2028`
//line /usr/local/go/src/html/template/js.go:209
			// _ = "end of CoverTab[31378]"
		} else {
//line /usr/local/go/src/html/template/js.go:210
			_go_fuzz_dep_.CoverTab[31379]++
//line /usr/local/go/src/html/template/js.go:210
			if rune == 0x2029 {
//line /usr/local/go/src/html/template/js.go:210
				_go_fuzz_dep_.CoverTab[31380]++
										repl = `\u2029`
//line /usr/local/go/src/html/template/js.go:211
				// _ = "end of CoverTab[31380]"
			} else {
//line /usr/local/go/src/html/template/js.go:212
				_go_fuzz_dep_.CoverTab[31381]++
//line /usr/local/go/src/html/template/js.go:212
				// _ = "end of CoverTab[31381]"
//line /usr/local/go/src/html/template/js.go:212
			}
//line /usr/local/go/src/html/template/js.go:212
			// _ = "end of CoverTab[31379]"
//line /usr/local/go/src/html/template/js.go:212
		}
//line /usr/local/go/src/html/template/js.go:212
		// _ = "end of CoverTab[31375]"
//line /usr/local/go/src/html/template/js.go:212
		_go_fuzz_dep_.CoverTab[31376]++
								if repl != "" {
//line /usr/local/go/src/html/template/js.go:213
			_go_fuzz_dep_.CoverTab[31382]++
									buf.Write(b[written:i])
									buf.WriteString(repl)
									written = i + n
//line /usr/local/go/src/html/template/js.go:216
			// _ = "end of CoverTab[31382]"
		} else {
//line /usr/local/go/src/html/template/js.go:217
			_go_fuzz_dep_.CoverTab[31383]++
//line /usr/local/go/src/html/template/js.go:217
			// _ = "end of CoverTab[31383]"
//line /usr/local/go/src/html/template/js.go:217
		}
//line /usr/local/go/src/html/template/js.go:217
		// _ = "end of CoverTab[31376]"
//line /usr/local/go/src/html/template/js.go:217
		_go_fuzz_dep_.CoverTab[31377]++
								i += n
//line /usr/local/go/src/html/template/js.go:218
		// _ = "end of CoverTab[31377]"
	}
//line /usr/local/go/src/html/template/js.go:219
	// _ = "end of CoverTab[31357]"
//line /usr/local/go/src/html/template/js.go:219
	_go_fuzz_dep_.CoverTab[31358]++
							if buf.Len() != 0 {
//line /usr/local/go/src/html/template/js.go:220
		_go_fuzz_dep_.CoverTab[31384]++
								buf.Write(b[written:])
								if pad {
//line /usr/local/go/src/html/template/js.go:222
			_go_fuzz_dep_.CoverTab[31386]++
									buf.WriteByte(' ')
//line /usr/local/go/src/html/template/js.go:223
			// _ = "end of CoverTab[31386]"
		} else {
//line /usr/local/go/src/html/template/js.go:224
			_go_fuzz_dep_.CoverTab[31387]++
//line /usr/local/go/src/html/template/js.go:224
			// _ = "end of CoverTab[31387]"
//line /usr/local/go/src/html/template/js.go:224
		}
//line /usr/local/go/src/html/template/js.go:224
		// _ = "end of CoverTab[31384]"
//line /usr/local/go/src/html/template/js.go:224
		_go_fuzz_dep_.CoverTab[31385]++
								return buf.String()
//line /usr/local/go/src/html/template/js.go:225
		// _ = "end of CoverTab[31385]"
	} else {
//line /usr/local/go/src/html/template/js.go:226
		_go_fuzz_dep_.CoverTab[31388]++
//line /usr/local/go/src/html/template/js.go:226
		// _ = "end of CoverTab[31388]"
//line /usr/local/go/src/html/template/js.go:226
	}
//line /usr/local/go/src/html/template/js.go:226
	// _ = "end of CoverTab[31358]"
//line /usr/local/go/src/html/template/js.go:226
	_go_fuzz_dep_.CoverTab[31359]++
							return string(b)
//line /usr/local/go/src/html/template/js.go:227
	// _ = "end of CoverTab[31359]"
}

// jsStrEscaper produces a string that can be included between quotes in
//line /usr/local/go/src/html/template/js.go:230
// JavaScript source, in JavaScript embedded in an HTML5 <script> element,
//line /usr/local/go/src/html/template/js.go:230
// or in an HTML5 event handler attribute such as onclick.
//line /usr/local/go/src/html/template/js.go:233
func jsStrEscaper(args ...any) string {
//line /usr/local/go/src/html/template/js.go:233
	_go_fuzz_dep_.CoverTab[31389]++
							s, t := stringify(args...)
							if t == contentTypeJSStr {
//line /usr/local/go/src/html/template/js.go:235
		_go_fuzz_dep_.CoverTab[31391]++
								return replace(s, jsStrNormReplacementTable)
//line /usr/local/go/src/html/template/js.go:236
		// _ = "end of CoverTab[31391]"
	} else {
//line /usr/local/go/src/html/template/js.go:237
		_go_fuzz_dep_.CoverTab[31392]++
//line /usr/local/go/src/html/template/js.go:237
		// _ = "end of CoverTab[31392]"
//line /usr/local/go/src/html/template/js.go:237
	}
//line /usr/local/go/src/html/template/js.go:237
	// _ = "end of CoverTab[31389]"
//line /usr/local/go/src/html/template/js.go:237
	_go_fuzz_dep_.CoverTab[31390]++
							return replace(s, jsStrReplacementTable)
//line /usr/local/go/src/html/template/js.go:238
	// _ = "end of CoverTab[31390]"
}

// jsRegexpEscaper behaves like jsStrEscaper but escapes regular expression
//line /usr/local/go/src/html/template/js.go:241
// specials so the result is treated literally when included in a regular
//line /usr/local/go/src/html/template/js.go:241
// expression literal. /foo{{.X}}bar/ matches the string "foo" followed by
//line /usr/local/go/src/html/template/js.go:241
// the literal text of {{.X}} followed by the string "bar".
//line /usr/local/go/src/html/template/js.go:245
func jsRegexpEscaper(args ...any) string {
//line /usr/local/go/src/html/template/js.go:245
	_go_fuzz_dep_.CoverTab[31393]++
							s, _ := stringify(args...)
							s = replace(s, jsRegexpReplacementTable)
							if s == "" {
//line /usr/local/go/src/html/template/js.go:248
		_go_fuzz_dep_.CoverTab[31395]++

								return "(?:)"
//line /usr/local/go/src/html/template/js.go:250
		// _ = "end of CoverTab[31395]"
	} else {
//line /usr/local/go/src/html/template/js.go:251
		_go_fuzz_dep_.CoverTab[31396]++
//line /usr/local/go/src/html/template/js.go:251
		// _ = "end of CoverTab[31396]"
//line /usr/local/go/src/html/template/js.go:251
	}
//line /usr/local/go/src/html/template/js.go:251
	// _ = "end of CoverTab[31393]"
//line /usr/local/go/src/html/template/js.go:251
	_go_fuzz_dep_.CoverTab[31394]++
							return s
//line /usr/local/go/src/html/template/js.go:252
	// _ = "end of CoverTab[31394]"
}

// replace replaces each rune r of s with replacementTable[r], provided that
//line /usr/local/go/src/html/template/js.go:255
// r < len(replacementTable). If replacementTable[r] is the empty string then
//line /usr/local/go/src/html/template/js.go:255
// no replacement is made.
//line /usr/local/go/src/html/template/js.go:255
// It also replaces runes U+2028 and U+2029 with the raw strings `\u2028` and
//line /usr/local/go/src/html/template/js.go:255
// `\u2029`.
//line /usr/local/go/src/html/template/js.go:260
func replace(s string, replacementTable []string) string {
//line /usr/local/go/src/html/template/js.go:260
	_go_fuzz_dep_.CoverTab[31397]++
							var b strings.Builder
							r, w, written := rune(0), 0, 0
							for i := 0; i < len(s); i += w {
//line /usr/local/go/src/html/template/js.go:263
		_go_fuzz_dep_.CoverTab[31400]++

								r, w = utf8.DecodeRuneInString(s[i:])
								var repl string
								switch {
		case int(r) < len(lowUnicodeReplacementTable):
//line /usr/local/go/src/html/template/js.go:268
			_go_fuzz_dep_.CoverTab[31403]++
									repl = lowUnicodeReplacementTable[r]
//line /usr/local/go/src/html/template/js.go:269
			// _ = "end of CoverTab[31403]"
		case int(r) < len(replacementTable) && func() bool {
//line /usr/local/go/src/html/template/js.go:270
			_go_fuzz_dep_.CoverTab[31408]++
//line /usr/local/go/src/html/template/js.go:270
			return replacementTable[r] != ""
//line /usr/local/go/src/html/template/js.go:270
			// _ = "end of CoverTab[31408]"
//line /usr/local/go/src/html/template/js.go:270
		}():
//line /usr/local/go/src/html/template/js.go:270
			_go_fuzz_dep_.CoverTab[31404]++
									repl = replacementTable[r]
//line /usr/local/go/src/html/template/js.go:271
			// _ = "end of CoverTab[31404]"
		case r == '\u2028':
//line /usr/local/go/src/html/template/js.go:272
			_go_fuzz_dep_.CoverTab[31405]++
									repl = `\u2028`
//line /usr/local/go/src/html/template/js.go:273
			// _ = "end of CoverTab[31405]"
		case r == '\u2029':
//line /usr/local/go/src/html/template/js.go:274
			_go_fuzz_dep_.CoverTab[31406]++
									repl = `\u2029`
//line /usr/local/go/src/html/template/js.go:275
			// _ = "end of CoverTab[31406]"
		default:
//line /usr/local/go/src/html/template/js.go:276
			_go_fuzz_dep_.CoverTab[31407]++
									continue
//line /usr/local/go/src/html/template/js.go:277
			// _ = "end of CoverTab[31407]"
		}
//line /usr/local/go/src/html/template/js.go:278
		// _ = "end of CoverTab[31400]"
//line /usr/local/go/src/html/template/js.go:278
		_go_fuzz_dep_.CoverTab[31401]++
								if written == 0 {
//line /usr/local/go/src/html/template/js.go:279
			_go_fuzz_dep_.CoverTab[31409]++
									b.Grow(len(s))
//line /usr/local/go/src/html/template/js.go:280
			// _ = "end of CoverTab[31409]"
		} else {
//line /usr/local/go/src/html/template/js.go:281
			_go_fuzz_dep_.CoverTab[31410]++
//line /usr/local/go/src/html/template/js.go:281
			// _ = "end of CoverTab[31410]"
//line /usr/local/go/src/html/template/js.go:281
		}
//line /usr/local/go/src/html/template/js.go:281
		// _ = "end of CoverTab[31401]"
//line /usr/local/go/src/html/template/js.go:281
		_go_fuzz_dep_.CoverTab[31402]++
								b.WriteString(s[written:i])
								b.WriteString(repl)
								written = i + w
//line /usr/local/go/src/html/template/js.go:284
		// _ = "end of CoverTab[31402]"
	}
//line /usr/local/go/src/html/template/js.go:285
	// _ = "end of CoverTab[31397]"
//line /usr/local/go/src/html/template/js.go:285
	_go_fuzz_dep_.CoverTab[31398]++
							if written == 0 {
//line /usr/local/go/src/html/template/js.go:286
		_go_fuzz_dep_.CoverTab[31411]++
								return s
//line /usr/local/go/src/html/template/js.go:287
		// _ = "end of CoverTab[31411]"
	} else {
//line /usr/local/go/src/html/template/js.go:288
		_go_fuzz_dep_.CoverTab[31412]++
//line /usr/local/go/src/html/template/js.go:288
		// _ = "end of CoverTab[31412]"
//line /usr/local/go/src/html/template/js.go:288
	}
//line /usr/local/go/src/html/template/js.go:288
	// _ = "end of CoverTab[31398]"
//line /usr/local/go/src/html/template/js.go:288
	_go_fuzz_dep_.CoverTab[31399]++
							b.WriteString(s[written:])
							return b.String()
//line /usr/local/go/src/html/template/js.go:290
	// _ = "end of CoverTab[31399]"
}

var lowUnicodeReplacementTable = []string{
	0:	`\u0000`, 1: `\u0001`, 2: `\u0002`, 3: `\u0003`, 4: `\u0004`, 5: `\u0005`, 6: `\u0006`,
	'\a':	`\u0007`,
	'\b':	`\u0008`,
	'\t':	`\t`,
	'\n':	`\n`,
	'\v':	`\u000b`,
	'\f':	`\f`,
	'\r':	`\r`,
	0xe:	`\u000e`, 0xf: `\u000f`, 0x10: `\u0010`, 0x11: `\u0011`, 0x12: `\u0012`, 0x13: `\u0013`,
	0x14:	`\u0014`, 0x15: `\u0015`, 0x16: `\u0016`, 0x17: `\u0017`, 0x18: `\u0018`, 0x19: `\u0019`,
	0x1a:	`\u001a`, 0x1b: `\u001b`, 0x1c: `\u001c`, 0x1d: `\u001d`, 0x1e: `\u001e`, 0x1f: `\u001f`,
}

var jsStrReplacementTable = []string{
							0:	`\u0000`,
							'\t':	`\t`,
							'\n':	`\n`,
							'\v':	`\u000b`,
							'\f':	`\f`,
							'\r':	`\r`,

//line /usr/local/go/src/html/template/js.go:316
	'"':	`\u0022`,
							'`':	`\u0060`,
							'&':	`\u0026`,
							'\'':	`\u0027`,
							'+':	`\u002b`,
							'/':	`\/`,
							'<':	`\u003c`,
							'>':	`\u003e`,
							'\\':	`\\`,
}

// jsStrNormReplacementTable is like jsStrReplacementTable but does not
//line /usr/local/go/src/html/template/js.go:327
// overencode existing escapes since this table has no entry for `\`.
//line /usr/local/go/src/html/template/js.go:329
var jsStrNormReplacementTable = []string{
							0:	`\u0000`,
							'\t':	`\t`,
							'\n':	`\n`,
							'\v':	`\u000b`,
							'\f':	`\f`,
							'\r':	`\r`,

//line /usr/local/go/src/html/template/js.go:338
	'"':	`\u0022`,
							'&':	`\u0026`,
							'\'':	`\u0027`,
							'`':	`\u0060`,
							'+':	`\u002b`,
							'/':	`\/`,
							'<':	`\u003c`,
							'>':	`\u003e`,
}
var jsRegexpReplacementTable = []string{
							0:	`\u0000`,
							'\t':	`\t`,
							'\n':	`\n`,
							'\v':	`\u000b`,
							'\f':	`\f`,
							'\r':	`\r`,

//line /usr/local/go/src/html/template/js.go:356
	'"':	`\u0022`,
							'$':	`\$`,
							'&':	`\u0026`,
							'\'':	`\u0027`,
							'(':	`\(`,
							')':	`\)`,
							'*':	`\*`,
							'+':	`\u002b`,
							'-':	`\-`,
							'.':	`\.`,
							'/':	`\/`,
							'<':	`\u003c`,
							'>':	`\u003e`,
							'?':	`\?`,
							'[':	`\[`,
							'\\':	`\\`,
							']':	`\]`,
							'^':	`\^`,
							'{':	`\{`,
							'|':	`\|`,
							'}':	`\}`,
}

// isJSIdentPart reports whether the given rune is a JS identifier part.
//line /usr/local/go/src/html/template/js.go:379
// It does not handle all the non-Latin letters, joiners, and combining marks,
//line /usr/local/go/src/html/template/js.go:379
// but it does handle every codepoint that can occur in a numeric literal or
//line /usr/local/go/src/html/template/js.go:379
// a keyword.
//line /usr/local/go/src/html/template/js.go:383
func isJSIdentPart(r rune) bool {
//line /usr/local/go/src/html/template/js.go:383
	_go_fuzz_dep_.CoverTab[31413]++
							switch {
	case r == '$':
//line /usr/local/go/src/html/template/js.go:385
		_go_fuzz_dep_.CoverTab[31415]++
								return true
//line /usr/local/go/src/html/template/js.go:386
		// _ = "end of CoverTab[31415]"
	case '0' <= r && func() bool {
//line /usr/local/go/src/html/template/js.go:387
		_go_fuzz_dep_.CoverTab[31421]++
//line /usr/local/go/src/html/template/js.go:387
		return r <= '9'
//line /usr/local/go/src/html/template/js.go:387
		// _ = "end of CoverTab[31421]"
//line /usr/local/go/src/html/template/js.go:387
	}():
//line /usr/local/go/src/html/template/js.go:387
		_go_fuzz_dep_.CoverTab[31416]++
								return true
//line /usr/local/go/src/html/template/js.go:388
		// _ = "end of CoverTab[31416]"
	case 'A' <= r && func() bool {
//line /usr/local/go/src/html/template/js.go:389
		_go_fuzz_dep_.CoverTab[31422]++
//line /usr/local/go/src/html/template/js.go:389
		return r <= 'Z'
//line /usr/local/go/src/html/template/js.go:389
		// _ = "end of CoverTab[31422]"
//line /usr/local/go/src/html/template/js.go:389
	}():
//line /usr/local/go/src/html/template/js.go:389
		_go_fuzz_dep_.CoverTab[31417]++
								return true
//line /usr/local/go/src/html/template/js.go:390
		// _ = "end of CoverTab[31417]"
	case r == '_':
//line /usr/local/go/src/html/template/js.go:391
		_go_fuzz_dep_.CoverTab[31418]++
								return true
//line /usr/local/go/src/html/template/js.go:392
		// _ = "end of CoverTab[31418]"
	case 'a' <= r && func() bool {
//line /usr/local/go/src/html/template/js.go:393
		_go_fuzz_dep_.CoverTab[31423]++
//line /usr/local/go/src/html/template/js.go:393
		return r <= 'z'
//line /usr/local/go/src/html/template/js.go:393
		// _ = "end of CoverTab[31423]"
//line /usr/local/go/src/html/template/js.go:393
	}():
//line /usr/local/go/src/html/template/js.go:393
		_go_fuzz_dep_.CoverTab[31419]++
								return true
//line /usr/local/go/src/html/template/js.go:394
		// _ = "end of CoverTab[31419]"
//line /usr/local/go/src/html/template/js.go:394
	default:
//line /usr/local/go/src/html/template/js.go:394
		_go_fuzz_dep_.CoverTab[31420]++
//line /usr/local/go/src/html/template/js.go:394
		// _ = "end of CoverTab[31420]"
	}
//line /usr/local/go/src/html/template/js.go:395
	// _ = "end of CoverTab[31413]"
//line /usr/local/go/src/html/template/js.go:395
	_go_fuzz_dep_.CoverTab[31414]++
							return false
//line /usr/local/go/src/html/template/js.go:396
	// _ = "end of CoverTab[31414]"
}

// isJSType reports whether the given MIME type should be considered JavaScript.
//line /usr/local/go/src/html/template/js.go:399
//
//line /usr/local/go/src/html/template/js.go:399
// It is used to determine whether a script tag with a type attribute is a javascript container.
//line /usr/local/go/src/html/template/js.go:402
func isJSType(mimeType string) bool {
//line /usr/local/go/src/html/template/js.go:402
	_go_fuzz_dep_.CoverTab[31424]++

//line /usr/local/go/src/html/template/js.go:409
	mimeType, _, _ = strings.Cut(mimeType, ";")
	mimeType = strings.ToLower(mimeType)
	mimeType = strings.TrimSpace(mimeType)
	switch mimeType {
	case
		"application/ecmascript",
		"application/javascript",
		"application/json",
		"application/ld+json",
		"application/x-ecmascript",
		"application/x-javascript",
		"module",
		"text/ecmascript",
		"text/javascript",
		"text/javascript1.0",
		"text/javascript1.1",
		"text/javascript1.2",
		"text/javascript1.3",
		"text/javascript1.4",
		"text/javascript1.5",
		"text/jscript",
		"text/livescript",
		"text/x-ecmascript",
		"text/x-javascript":
//line /usr/local/go/src/html/template/js.go:432
		_go_fuzz_dep_.CoverTab[31425]++
								return true
//line /usr/local/go/src/html/template/js.go:433
		// _ = "end of CoverTab[31425]"
	default:
//line /usr/local/go/src/html/template/js.go:434
		_go_fuzz_dep_.CoverTab[31426]++
								return false
//line /usr/local/go/src/html/template/js.go:435
		// _ = "end of CoverTab[31426]"
	}
//line /usr/local/go/src/html/template/js.go:436
	// _ = "end of CoverTab[31424]"
}

//line /usr/local/go/src/html/template/js.go:437
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/html/template/js.go:437
var _ = _go_fuzz_dep_.CoverTab
