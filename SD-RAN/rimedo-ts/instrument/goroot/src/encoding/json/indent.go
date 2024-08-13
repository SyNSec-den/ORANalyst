// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/encoding/json/indent.go:5
package json

//line /usr/local/go/src/encoding/json/indent.go:5
import (
//line /usr/local/go/src/encoding/json/indent.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/encoding/json/indent.go:5
)
//line /usr/local/go/src/encoding/json/indent.go:5
import (
//line /usr/local/go/src/encoding/json/indent.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/encoding/json/indent.go:5
)

import (
	"bytes"
)

// Compact appends to dst the JSON-encoded src with
//line /usr/local/go/src/encoding/json/indent.go:11
// insignificant space characters elided.
//line /usr/local/go/src/encoding/json/indent.go:13
func Compact(dst *bytes.Buffer, src []byte) error {
//line /usr/local/go/src/encoding/json/indent.go:13
	_go_fuzz_dep_.CoverTab[28049]++
							return compact(dst, src, false)
//line /usr/local/go/src/encoding/json/indent.go:14
	// _ = "end of CoverTab[28049]"
}

func compact(dst *bytes.Buffer, src []byte, escape bool) error {
//line /usr/local/go/src/encoding/json/indent.go:17
	_go_fuzz_dep_.CoverTab[28050]++
							origLen := dst.Len()
							scan := newScanner()
							defer freeScanner(scan)
							start := 0
							for i, c := range src {
//line /usr/local/go/src/encoding/json/indent.go:22
		_go_fuzz_dep_.CoverTab[28054]++
								if escape && func() bool {
//line /usr/local/go/src/encoding/json/indent.go:23
			_go_fuzz_dep_.CoverTab[28057]++
//line /usr/local/go/src/encoding/json/indent.go:23
			return (c == '<' || func() bool {
//line /usr/local/go/src/encoding/json/indent.go:23
				_go_fuzz_dep_.CoverTab[28058]++
//line /usr/local/go/src/encoding/json/indent.go:23
				return c == '>'
//line /usr/local/go/src/encoding/json/indent.go:23
				// _ = "end of CoverTab[28058]"
//line /usr/local/go/src/encoding/json/indent.go:23
			}() || func() bool {
//line /usr/local/go/src/encoding/json/indent.go:23
				_go_fuzz_dep_.CoverTab[28059]++
//line /usr/local/go/src/encoding/json/indent.go:23
				return c == '&'
//line /usr/local/go/src/encoding/json/indent.go:23
				// _ = "end of CoverTab[28059]"
//line /usr/local/go/src/encoding/json/indent.go:23
			}())
//line /usr/local/go/src/encoding/json/indent.go:23
			// _ = "end of CoverTab[28057]"
//line /usr/local/go/src/encoding/json/indent.go:23
		}() {
//line /usr/local/go/src/encoding/json/indent.go:23
			_go_fuzz_dep_.CoverTab[28060]++
									if start < i {
//line /usr/local/go/src/encoding/json/indent.go:24
				_go_fuzz_dep_.CoverTab[28062]++
										dst.Write(src[start:i])
//line /usr/local/go/src/encoding/json/indent.go:25
				// _ = "end of CoverTab[28062]"
			} else {
//line /usr/local/go/src/encoding/json/indent.go:26
				_go_fuzz_dep_.CoverTab[28063]++
//line /usr/local/go/src/encoding/json/indent.go:26
				// _ = "end of CoverTab[28063]"
//line /usr/local/go/src/encoding/json/indent.go:26
			}
//line /usr/local/go/src/encoding/json/indent.go:26
			// _ = "end of CoverTab[28060]"
//line /usr/local/go/src/encoding/json/indent.go:26
			_go_fuzz_dep_.CoverTab[28061]++
									dst.WriteString(`\u00`)
									dst.WriteByte(hex[c>>4])
									dst.WriteByte(hex[c&0xF])
									start = i + 1
//line /usr/local/go/src/encoding/json/indent.go:30
			// _ = "end of CoverTab[28061]"
		} else {
//line /usr/local/go/src/encoding/json/indent.go:31
			_go_fuzz_dep_.CoverTab[28064]++
//line /usr/local/go/src/encoding/json/indent.go:31
			// _ = "end of CoverTab[28064]"
//line /usr/local/go/src/encoding/json/indent.go:31
		}
//line /usr/local/go/src/encoding/json/indent.go:31
		// _ = "end of CoverTab[28054]"
//line /usr/local/go/src/encoding/json/indent.go:31
		_go_fuzz_dep_.CoverTab[28055]++

								if escape && func() bool {
//line /usr/local/go/src/encoding/json/indent.go:33
			_go_fuzz_dep_.CoverTab[28065]++
//line /usr/local/go/src/encoding/json/indent.go:33
			return c == 0xE2
//line /usr/local/go/src/encoding/json/indent.go:33
			// _ = "end of CoverTab[28065]"
//line /usr/local/go/src/encoding/json/indent.go:33
		}() && func() bool {
//line /usr/local/go/src/encoding/json/indent.go:33
			_go_fuzz_dep_.CoverTab[28066]++
//line /usr/local/go/src/encoding/json/indent.go:33
			return i+2 < len(src)
//line /usr/local/go/src/encoding/json/indent.go:33
			// _ = "end of CoverTab[28066]"
//line /usr/local/go/src/encoding/json/indent.go:33
		}() && func() bool {
//line /usr/local/go/src/encoding/json/indent.go:33
			_go_fuzz_dep_.CoverTab[28067]++
//line /usr/local/go/src/encoding/json/indent.go:33
			return src[i+1] == 0x80
//line /usr/local/go/src/encoding/json/indent.go:33
			// _ = "end of CoverTab[28067]"
//line /usr/local/go/src/encoding/json/indent.go:33
		}() && func() bool {
//line /usr/local/go/src/encoding/json/indent.go:33
			_go_fuzz_dep_.CoverTab[28068]++
//line /usr/local/go/src/encoding/json/indent.go:33
			return src[i+2]&^1 == 0xA8
//line /usr/local/go/src/encoding/json/indent.go:33
			// _ = "end of CoverTab[28068]"
//line /usr/local/go/src/encoding/json/indent.go:33
		}() {
//line /usr/local/go/src/encoding/json/indent.go:33
			_go_fuzz_dep_.CoverTab[28069]++
									if start < i {
//line /usr/local/go/src/encoding/json/indent.go:34
				_go_fuzz_dep_.CoverTab[28071]++
										dst.Write(src[start:i])
//line /usr/local/go/src/encoding/json/indent.go:35
				// _ = "end of CoverTab[28071]"
			} else {
//line /usr/local/go/src/encoding/json/indent.go:36
				_go_fuzz_dep_.CoverTab[28072]++
//line /usr/local/go/src/encoding/json/indent.go:36
				// _ = "end of CoverTab[28072]"
//line /usr/local/go/src/encoding/json/indent.go:36
			}
//line /usr/local/go/src/encoding/json/indent.go:36
			// _ = "end of CoverTab[28069]"
//line /usr/local/go/src/encoding/json/indent.go:36
			_go_fuzz_dep_.CoverTab[28070]++
									dst.WriteString(`\u202`)
									dst.WriteByte(hex[src[i+2]&0xF])
									start = i + 3
//line /usr/local/go/src/encoding/json/indent.go:39
			// _ = "end of CoverTab[28070]"
		} else {
//line /usr/local/go/src/encoding/json/indent.go:40
			_go_fuzz_dep_.CoverTab[28073]++
//line /usr/local/go/src/encoding/json/indent.go:40
			// _ = "end of CoverTab[28073]"
//line /usr/local/go/src/encoding/json/indent.go:40
		}
//line /usr/local/go/src/encoding/json/indent.go:40
		// _ = "end of CoverTab[28055]"
//line /usr/local/go/src/encoding/json/indent.go:40
		_go_fuzz_dep_.CoverTab[28056]++
								v := scan.step(scan, c)
								if v >= scanSkipSpace {
//line /usr/local/go/src/encoding/json/indent.go:42
			_go_fuzz_dep_.CoverTab[28074]++
									if v == scanError {
//line /usr/local/go/src/encoding/json/indent.go:43
				_go_fuzz_dep_.CoverTab[28077]++
										break
//line /usr/local/go/src/encoding/json/indent.go:44
				// _ = "end of CoverTab[28077]"
			} else {
//line /usr/local/go/src/encoding/json/indent.go:45
				_go_fuzz_dep_.CoverTab[28078]++
//line /usr/local/go/src/encoding/json/indent.go:45
				// _ = "end of CoverTab[28078]"
//line /usr/local/go/src/encoding/json/indent.go:45
			}
//line /usr/local/go/src/encoding/json/indent.go:45
			// _ = "end of CoverTab[28074]"
//line /usr/local/go/src/encoding/json/indent.go:45
			_go_fuzz_dep_.CoverTab[28075]++
									if start < i {
//line /usr/local/go/src/encoding/json/indent.go:46
				_go_fuzz_dep_.CoverTab[28079]++
										dst.Write(src[start:i])
//line /usr/local/go/src/encoding/json/indent.go:47
				// _ = "end of CoverTab[28079]"
			} else {
//line /usr/local/go/src/encoding/json/indent.go:48
				_go_fuzz_dep_.CoverTab[28080]++
//line /usr/local/go/src/encoding/json/indent.go:48
				// _ = "end of CoverTab[28080]"
//line /usr/local/go/src/encoding/json/indent.go:48
			}
//line /usr/local/go/src/encoding/json/indent.go:48
			// _ = "end of CoverTab[28075]"
//line /usr/local/go/src/encoding/json/indent.go:48
			_go_fuzz_dep_.CoverTab[28076]++
									start = i + 1
//line /usr/local/go/src/encoding/json/indent.go:49
			// _ = "end of CoverTab[28076]"
		} else {
//line /usr/local/go/src/encoding/json/indent.go:50
			_go_fuzz_dep_.CoverTab[28081]++
//line /usr/local/go/src/encoding/json/indent.go:50
			// _ = "end of CoverTab[28081]"
//line /usr/local/go/src/encoding/json/indent.go:50
		}
//line /usr/local/go/src/encoding/json/indent.go:50
		// _ = "end of CoverTab[28056]"
	}
//line /usr/local/go/src/encoding/json/indent.go:51
	// _ = "end of CoverTab[28050]"
//line /usr/local/go/src/encoding/json/indent.go:51
	_go_fuzz_dep_.CoverTab[28051]++
							if scan.eof() == scanError {
//line /usr/local/go/src/encoding/json/indent.go:52
		_go_fuzz_dep_.CoverTab[28082]++
								dst.Truncate(origLen)
								return scan.err
//line /usr/local/go/src/encoding/json/indent.go:54
		// _ = "end of CoverTab[28082]"
	} else {
//line /usr/local/go/src/encoding/json/indent.go:55
		_go_fuzz_dep_.CoverTab[28083]++
//line /usr/local/go/src/encoding/json/indent.go:55
		// _ = "end of CoverTab[28083]"
//line /usr/local/go/src/encoding/json/indent.go:55
	}
//line /usr/local/go/src/encoding/json/indent.go:55
	// _ = "end of CoverTab[28051]"
//line /usr/local/go/src/encoding/json/indent.go:55
	_go_fuzz_dep_.CoverTab[28052]++
							if start < len(src) {
//line /usr/local/go/src/encoding/json/indent.go:56
		_go_fuzz_dep_.CoverTab[28084]++
								dst.Write(src[start:])
//line /usr/local/go/src/encoding/json/indent.go:57
		// _ = "end of CoverTab[28084]"
	} else {
//line /usr/local/go/src/encoding/json/indent.go:58
		_go_fuzz_dep_.CoverTab[28085]++
//line /usr/local/go/src/encoding/json/indent.go:58
		// _ = "end of CoverTab[28085]"
//line /usr/local/go/src/encoding/json/indent.go:58
	}
//line /usr/local/go/src/encoding/json/indent.go:58
	// _ = "end of CoverTab[28052]"
//line /usr/local/go/src/encoding/json/indent.go:58
	_go_fuzz_dep_.CoverTab[28053]++
							return nil
//line /usr/local/go/src/encoding/json/indent.go:59
	// _ = "end of CoverTab[28053]"
}

func newline(dst *bytes.Buffer, prefix, indent string, depth int) {
//line /usr/local/go/src/encoding/json/indent.go:62
	_go_fuzz_dep_.CoverTab[28086]++
							dst.WriteByte('\n')
							dst.WriteString(prefix)
							for i := 0; i < depth; i++ {
//line /usr/local/go/src/encoding/json/indent.go:65
		_go_fuzz_dep_.CoverTab[28087]++
								dst.WriteString(indent)
//line /usr/local/go/src/encoding/json/indent.go:66
		// _ = "end of CoverTab[28087]"
	}
//line /usr/local/go/src/encoding/json/indent.go:67
	// _ = "end of CoverTab[28086]"
}

// Indent appends to dst an indented form of the JSON-encoded src.
//line /usr/local/go/src/encoding/json/indent.go:70
// Each element in a JSON object or array begins on a new,
//line /usr/local/go/src/encoding/json/indent.go:70
// indented line beginning with prefix followed by one or more
//line /usr/local/go/src/encoding/json/indent.go:70
// copies of indent according to the indentation nesting.
//line /usr/local/go/src/encoding/json/indent.go:70
// The data appended to dst does not begin with the prefix nor
//line /usr/local/go/src/encoding/json/indent.go:70
// any indentation, to make it easier to embed inside other formatted JSON data.
//line /usr/local/go/src/encoding/json/indent.go:70
// Although leading space characters (space, tab, carriage return, newline)
//line /usr/local/go/src/encoding/json/indent.go:70
// at the beginning of src are dropped, trailing space characters
//line /usr/local/go/src/encoding/json/indent.go:70
// at the end of src are preserved and copied to dst.
//line /usr/local/go/src/encoding/json/indent.go:70
// For example, if src has no trailing spaces, neither will dst;
//line /usr/local/go/src/encoding/json/indent.go:70
// if src ends in a trailing newline, so will dst.
//line /usr/local/go/src/encoding/json/indent.go:81
func Indent(dst *bytes.Buffer, src []byte, prefix, indent string) error {
//line /usr/local/go/src/encoding/json/indent.go:81
	_go_fuzz_dep_.CoverTab[28088]++
							origLen := dst.Len()
							scan := newScanner()
							defer freeScanner(scan)
							needIndent := false
							depth := 0
							for _, c := range src {
//line /usr/local/go/src/encoding/json/indent.go:87
		_go_fuzz_dep_.CoverTab[28091]++
								scan.bytes++
								v := scan.step(scan, c)
								if v == scanSkipSpace {
//line /usr/local/go/src/encoding/json/indent.go:90
			_go_fuzz_dep_.CoverTab[28096]++
									continue
//line /usr/local/go/src/encoding/json/indent.go:91
			// _ = "end of CoverTab[28096]"
		} else {
//line /usr/local/go/src/encoding/json/indent.go:92
			_go_fuzz_dep_.CoverTab[28097]++
//line /usr/local/go/src/encoding/json/indent.go:92
			// _ = "end of CoverTab[28097]"
//line /usr/local/go/src/encoding/json/indent.go:92
		}
//line /usr/local/go/src/encoding/json/indent.go:92
		// _ = "end of CoverTab[28091]"
//line /usr/local/go/src/encoding/json/indent.go:92
		_go_fuzz_dep_.CoverTab[28092]++
								if v == scanError {
//line /usr/local/go/src/encoding/json/indent.go:93
			_go_fuzz_dep_.CoverTab[28098]++
									break
//line /usr/local/go/src/encoding/json/indent.go:94
			// _ = "end of CoverTab[28098]"
		} else {
//line /usr/local/go/src/encoding/json/indent.go:95
			_go_fuzz_dep_.CoverTab[28099]++
//line /usr/local/go/src/encoding/json/indent.go:95
			// _ = "end of CoverTab[28099]"
//line /usr/local/go/src/encoding/json/indent.go:95
		}
//line /usr/local/go/src/encoding/json/indent.go:95
		// _ = "end of CoverTab[28092]"
//line /usr/local/go/src/encoding/json/indent.go:95
		_go_fuzz_dep_.CoverTab[28093]++
								if needIndent && func() bool {
//line /usr/local/go/src/encoding/json/indent.go:96
			_go_fuzz_dep_.CoverTab[28100]++
//line /usr/local/go/src/encoding/json/indent.go:96
			return v != scanEndObject
//line /usr/local/go/src/encoding/json/indent.go:96
			// _ = "end of CoverTab[28100]"
//line /usr/local/go/src/encoding/json/indent.go:96
		}() && func() bool {
//line /usr/local/go/src/encoding/json/indent.go:96
			_go_fuzz_dep_.CoverTab[28101]++
//line /usr/local/go/src/encoding/json/indent.go:96
			return v != scanEndArray
//line /usr/local/go/src/encoding/json/indent.go:96
			// _ = "end of CoverTab[28101]"
//line /usr/local/go/src/encoding/json/indent.go:96
		}() {
//line /usr/local/go/src/encoding/json/indent.go:96
			_go_fuzz_dep_.CoverTab[28102]++
									needIndent = false
									depth++
									newline(dst, prefix, indent, depth)
//line /usr/local/go/src/encoding/json/indent.go:99
			// _ = "end of CoverTab[28102]"
		} else {
//line /usr/local/go/src/encoding/json/indent.go:100
			_go_fuzz_dep_.CoverTab[28103]++
//line /usr/local/go/src/encoding/json/indent.go:100
			// _ = "end of CoverTab[28103]"
//line /usr/local/go/src/encoding/json/indent.go:100
		}
//line /usr/local/go/src/encoding/json/indent.go:100
		// _ = "end of CoverTab[28093]"
//line /usr/local/go/src/encoding/json/indent.go:100
		_go_fuzz_dep_.CoverTab[28094]++

//line /usr/local/go/src/encoding/json/indent.go:104
		if v == scanContinue {
//line /usr/local/go/src/encoding/json/indent.go:104
			_go_fuzz_dep_.CoverTab[28104]++
									dst.WriteByte(c)
									continue
//line /usr/local/go/src/encoding/json/indent.go:106
			// _ = "end of CoverTab[28104]"
		} else {
//line /usr/local/go/src/encoding/json/indent.go:107
			_go_fuzz_dep_.CoverTab[28105]++
//line /usr/local/go/src/encoding/json/indent.go:107
			// _ = "end of CoverTab[28105]"
//line /usr/local/go/src/encoding/json/indent.go:107
		}
//line /usr/local/go/src/encoding/json/indent.go:107
		// _ = "end of CoverTab[28094]"
//line /usr/local/go/src/encoding/json/indent.go:107
		_go_fuzz_dep_.CoverTab[28095]++

//line /usr/local/go/src/encoding/json/indent.go:110
		switch c {
		case '{', '[':
//line /usr/local/go/src/encoding/json/indent.go:111
			_go_fuzz_dep_.CoverTab[28106]++

									needIndent = true
									dst.WriteByte(c)
//line /usr/local/go/src/encoding/json/indent.go:114
			// _ = "end of CoverTab[28106]"

		case ',':
//line /usr/local/go/src/encoding/json/indent.go:116
			_go_fuzz_dep_.CoverTab[28107]++
									dst.WriteByte(c)
									newline(dst, prefix, indent, depth)
//line /usr/local/go/src/encoding/json/indent.go:118
			// _ = "end of CoverTab[28107]"

		case ':':
//line /usr/local/go/src/encoding/json/indent.go:120
			_go_fuzz_dep_.CoverTab[28108]++
									dst.WriteByte(c)
									dst.WriteByte(' ')
//line /usr/local/go/src/encoding/json/indent.go:122
			// _ = "end of CoverTab[28108]"

		case '}', ']':
//line /usr/local/go/src/encoding/json/indent.go:124
			_go_fuzz_dep_.CoverTab[28109]++
									if needIndent {
//line /usr/local/go/src/encoding/json/indent.go:125
				_go_fuzz_dep_.CoverTab[28112]++

										needIndent = false
//line /usr/local/go/src/encoding/json/indent.go:127
				// _ = "end of CoverTab[28112]"
			} else {
//line /usr/local/go/src/encoding/json/indent.go:128
				_go_fuzz_dep_.CoverTab[28113]++
										depth--
										newline(dst, prefix, indent, depth)
//line /usr/local/go/src/encoding/json/indent.go:130
				// _ = "end of CoverTab[28113]"
			}
//line /usr/local/go/src/encoding/json/indent.go:131
			// _ = "end of CoverTab[28109]"
//line /usr/local/go/src/encoding/json/indent.go:131
			_go_fuzz_dep_.CoverTab[28110]++
									dst.WriteByte(c)
//line /usr/local/go/src/encoding/json/indent.go:132
			// _ = "end of CoverTab[28110]"

		default:
//line /usr/local/go/src/encoding/json/indent.go:134
			_go_fuzz_dep_.CoverTab[28111]++
									dst.WriteByte(c)
//line /usr/local/go/src/encoding/json/indent.go:135
			// _ = "end of CoverTab[28111]"
		}
//line /usr/local/go/src/encoding/json/indent.go:136
		// _ = "end of CoverTab[28095]"
	}
//line /usr/local/go/src/encoding/json/indent.go:137
	// _ = "end of CoverTab[28088]"
//line /usr/local/go/src/encoding/json/indent.go:137
	_go_fuzz_dep_.CoverTab[28089]++
							if scan.eof() == scanError {
//line /usr/local/go/src/encoding/json/indent.go:138
		_go_fuzz_dep_.CoverTab[28114]++
								dst.Truncate(origLen)
								return scan.err
//line /usr/local/go/src/encoding/json/indent.go:140
		// _ = "end of CoverTab[28114]"
	} else {
//line /usr/local/go/src/encoding/json/indent.go:141
		_go_fuzz_dep_.CoverTab[28115]++
//line /usr/local/go/src/encoding/json/indent.go:141
		// _ = "end of CoverTab[28115]"
//line /usr/local/go/src/encoding/json/indent.go:141
	}
//line /usr/local/go/src/encoding/json/indent.go:141
	// _ = "end of CoverTab[28089]"
//line /usr/local/go/src/encoding/json/indent.go:141
	_go_fuzz_dep_.CoverTab[28090]++
							return nil
//line /usr/local/go/src/encoding/json/indent.go:142
	// _ = "end of CoverTab[28090]"
}

//line /usr/local/go/src/encoding/json/indent.go:143
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/encoding/json/indent.go:143
var _ = _go_fuzz_dep_.CoverTab
