// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/mime/mediatype.go:5
package mime

//line /usr/local/go/src/mime/mediatype.go:5
import (
//line /usr/local/go/src/mime/mediatype.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/mime/mediatype.go:5
)
//line /usr/local/go/src/mime/mediatype.go:5
import (
//line /usr/local/go/src/mime/mediatype.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/mime/mediatype.go:5
)

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"unicode"
)

// FormatMediaType serializes mediatype t and the parameters
//line /usr/local/go/src/mime/mediatype.go:15
// param as a media type conforming to RFC 2045 and RFC 2616.
//line /usr/local/go/src/mime/mediatype.go:15
// The type and parameter names are written in lower-case.
//line /usr/local/go/src/mime/mediatype.go:15
// When any of the arguments result in a standard violation then
//line /usr/local/go/src/mime/mediatype.go:15
// FormatMediaType returns the empty string.
//line /usr/local/go/src/mime/mediatype.go:20
func FormatMediaType(t string, param map[string]string) string {
//line /usr/local/go/src/mime/mediatype.go:20
	_go_fuzz_dep_.CoverTab[35608]++
						var b strings.Builder
						if major, sub, ok := strings.Cut(t, "/"); !ok {
//line /usr/local/go/src/mime/mediatype.go:22
		_go_fuzz_dep_.CoverTab[35612]++
							if !isToken(t) {
//line /usr/local/go/src/mime/mediatype.go:23
			_go_fuzz_dep_.CoverTab[35614]++
								return ""
//line /usr/local/go/src/mime/mediatype.go:24
			// _ = "end of CoverTab[35614]"
		} else {
//line /usr/local/go/src/mime/mediatype.go:25
			_go_fuzz_dep_.CoverTab[35615]++
//line /usr/local/go/src/mime/mediatype.go:25
			// _ = "end of CoverTab[35615]"
//line /usr/local/go/src/mime/mediatype.go:25
		}
//line /usr/local/go/src/mime/mediatype.go:25
		// _ = "end of CoverTab[35612]"
//line /usr/local/go/src/mime/mediatype.go:25
		_go_fuzz_dep_.CoverTab[35613]++
							b.WriteString(strings.ToLower(t))
//line /usr/local/go/src/mime/mediatype.go:26
		// _ = "end of CoverTab[35613]"
	} else {
//line /usr/local/go/src/mime/mediatype.go:27
		_go_fuzz_dep_.CoverTab[35616]++
							if !isToken(major) || func() bool {
//line /usr/local/go/src/mime/mediatype.go:28
			_go_fuzz_dep_.CoverTab[35618]++
//line /usr/local/go/src/mime/mediatype.go:28
			return !isToken(sub)
//line /usr/local/go/src/mime/mediatype.go:28
			// _ = "end of CoverTab[35618]"
//line /usr/local/go/src/mime/mediatype.go:28
		}() {
//line /usr/local/go/src/mime/mediatype.go:28
			_go_fuzz_dep_.CoverTab[35619]++
								return ""
//line /usr/local/go/src/mime/mediatype.go:29
			// _ = "end of CoverTab[35619]"
		} else {
//line /usr/local/go/src/mime/mediatype.go:30
			_go_fuzz_dep_.CoverTab[35620]++
//line /usr/local/go/src/mime/mediatype.go:30
			// _ = "end of CoverTab[35620]"
//line /usr/local/go/src/mime/mediatype.go:30
		}
//line /usr/local/go/src/mime/mediatype.go:30
		// _ = "end of CoverTab[35616]"
//line /usr/local/go/src/mime/mediatype.go:30
		_go_fuzz_dep_.CoverTab[35617]++
							b.WriteString(strings.ToLower(major))
							b.WriteByte('/')
							b.WriteString(strings.ToLower(sub))
//line /usr/local/go/src/mime/mediatype.go:33
		// _ = "end of CoverTab[35617]"
	}
//line /usr/local/go/src/mime/mediatype.go:34
	// _ = "end of CoverTab[35608]"
//line /usr/local/go/src/mime/mediatype.go:34
	_go_fuzz_dep_.CoverTab[35609]++

						attrs := make([]string, 0, len(param))
						for a := range param {
//line /usr/local/go/src/mime/mediatype.go:37
		_go_fuzz_dep_.CoverTab[35621]++
							attrs = append(attrs, a)
//line /usr/local/go/src/mime/mediatype.go:38
		// _ = "end of CoverTab[35621]"
	}
//line /usr/local/go/src/mime/mediatype.go:39
	// _ = "end of CoverTab[35609]"
//line /usr/local/go/src/mime/mediatype.go:39
	_go_fuzz_dep_.CoverTab[35610]++
						sort.Strings(attrs)

						for _, attribute := range attrs {
//line /usr/local/go/src/mime/mediatype.go:42
		_go_fuzz_dep_.CoverTab[35622]++
							value := param[attribute]
							b.WriteByte(';')
							b.WriteByte(' ')
							if !isToken(attribute) {
//line /usr/local/go/src/mime/mediatype.go:46
			_go_fuzz_dep_.CoverTab[35628]++
								return ""
//line /usr/local/go/src/mime/mediatype.go:47
			// _ = "end of CoverTab[35628]"
		} else {
//line /usr/local/go/src/mime/mediatype.go:48
			_go_fuzz_dep_.CoverTab[35629]++
//line /usr/local/go/src/mime/mediatype.go:48
			// _ = "end of CoverTab[35629]"
//line /usr/local/go/src/mime/mediatype.go:48
		}
//line /usr/local/go/src/mime/mediatype.go:48
		// _ = "end of CoverTab[35622]"
//line /usr/local/go/src/mime/mediatype.go:48
		_go_fuzz_dep_.CoverTab[35623]++
							b.WriteString(strings.ToLower(attribute))

							needEnc := needsEncoding(value)
							if needEnc {
//line /usr/local/go/src/mime/mediatype.go:52
			_go_fuzz_dep_.CoverTab[35630]++

								b.WriteByte('*')
//line /usr/local/go/src/mime/mediatype.go:54
			// _ = "end of CoverTab[35630]"
		} else {
//line /usr/local/go/src/mime/mediatype.go:55
			_go_fuzz_dep_.CoverTab[35631]++
//line /usr/local/go/src/mime/mediatype.go:55
			// _ = "end of CoverTab[35631]"
//line /usr/local/go/src/mime/mediatype.go:55
		}
//line /usr/local/go/src/mime/mediatype.go:55
		// _ = "end of CoverTab[35623]"
//line /usr/local/go/src/mime/mediatype.go:55
		_go_fuzz_dep_.CoverTab[35624]++
							b.WriteByte('=')

							if needEnc {
//line /usr/local/go/src/mime/mediatype.go:58
			_go_fuzz_dep_.CoverTab[35632]++
								b.WriteString("utf-8''")

								offset := 0
								for index := 0; index < len(value); index++ {
//line /usr/local/go/src/mime/mediatype.go:62
				_go_fuzz_dep_.CoverTab[35634]++
									ch := value[index]

//line /usr/local/go/src/mime/mediatype.go:66
				if ch <= ' ' || func() bool {
//line /usr/local/go/src/mime/mediatype.go:66
					_go_fuzz_dep_.CoverTab[35635]++
//line /usr/local/go/src/mime/mediatype.go:66
					return ch >= 0x7F
//line /usr/local/go/src/mime/mediatype.go:66
					// _ = "end of CoverTab[35635]"
//line /usr/local/go/src/mime/mediatype.go:66
				}() || func() bool {
//line /usr/local/go/src/mime/mediatype.go:66
					_go_fuzz_dep_.CoverTab[35636]++
//line /usr/local/go/src/mime/mediatype.go:66
					return ch == '*'
										// _ = "end of CoverTab[35636]"
//line /usr/local/go/src/mime/mediatype.go:67
				}() || func() bool {
//line /usr/local/go/src/mime/mediatype.go:67
					_go_fuzz_dep_.CoverTab[35637]++
//line /usr/local/go/src/mime/mediatype.go:67
					return ch == '\''
//line /usr/local/go/src/mime/mediatype.go:67
					// _ = "end of CoverTab[35637]"
//line /usr/local/go/src/mime/mediatype.go:67
				}() || func() bool {
//line /usr/local/go/src/mime/mediatype.go:67
					_go_fuzz_dep_.CoverTab[35638]++
//line /usr/local/go/src/mime/mediatype.go:67
					return ch == '%'
//line /usr/local/go/src/mime/mediatype.go:67
					// _ = "end of CoverTab[35638]"
//line /usr/local/go/src/mime/mediatype.go:67
				}() || func() bool {
//line /usr/local/go/src/mime/mediatype.go:67
					_go_fuzz_dep_.CoverTab[35639]++
//line /usr/local/go/src/mime/mediatype.go:67
					return isTSpecial(rune(ch))
										// _ = "end of CoverTab[35639]"
//line /usr/local/go/src/mime/mediatype.go:68
				}() {
//line /usr/local/go/src/mime/mediatype.go:68
					_go_fuzz_dep_.CoverTab[35640]++

										b.WriteString(value[offset:index])
										offset = index + 1

										b.WriteByte('%')
										b.WriteByte(upperhex[ch>>4])
										b.WriteByte(upperhex[ch&0x0F])
//line /usr/local/go/src/mime/mediatype.go:75
					// _ = "end of CoverTab[35640]"
				} else {
//line /usr/local/go/src/mime/mediatype.go:76
					_go_fuzz_dep_.CoverTab[35641]++
//line /usr/local/go/src/mime/mediatype.go:76
					// _ = "end of CoverTab[35641]"
//line /usr/local/go/src/mime/mediatype.go:76
				}
//line /usr/local/go/src/mime/mediatype.go:76
				// _ = "end of CoverTab[35634]"
			}
//line /usr/local/go/src/mime/mediatype.go:77
			// _ = "end of CoverTab[35632]"
//line /usr/local/go/src/mime/mediatype.go:77
			_go_fuzz_dep_.CoverTab[35633]++
								b.WriteString(value[offset:])
								continue
//line /usr/local/go/src/mime/mediatype.go:79
			// _ = "end of CoverTab[35633]"
		} else {
//line /usr/local/go/src/mime/mediatype.go:80
			_go_fuzz_dep_.CoverTab[35642]++
//line /usr/local/go/src/mime/mediatype.go:80
			// _ = "end of CoverTab[35642]"
//line /usr/local/go/src/mime/mediatype.go:80
		}
//line /usr/local/go/src/mime/mediatype.go:80
		// _ = "end of CoverTab[35624]"
//line /usr/local/go/src/mime/mediatype.go:80
		_go_fuzz_dep_.CoverTab[35625]++

							if isToken(value) {
//line /usr/local/go/src/mime/mediatype.go:82
			_go_fuzz_dep_.CoverTab[35643]++
								b.WriteString(value)
								continue
//line /usr/local/go/src/mime/mediatype.go:84
			// _ = "end of CoverTab[35643]"
		} else {
//line /usr/local/go/src/mime/mediatype.go:85
			_go_fuzz_dep_.CoverTab[35644]++
//line /usr/local/go/src/mime/mediatype.go:85
			// _ = "end of CoverTab[35644]"
//line /usr/local/go/src/mime/mediatype.go:85
		}
//line /usr/local/go/src/mime/mediatype.go:85
		// _ = "end of CoverTab[35625]"
//line /usr/local/go/src/mime/mediatype.go:85
		_go_fuzz_dep_.CoverTab[35626]++

							b.WriteByte('"')
							offset := 0
							for index := 0; index < len(value); index++ {
//line /usr/local/go/src/mime/mediatype.go:89
			_go_fuzz_dep_.CoverTab[35645]++
								character := value[index]
								if character == '"' || func() bool {
//line /usr/local/go/src/mime/mediatype.go:91
				_go_fuzz_dep_.CoverTab[35646]++
//line /usr/local/go/src/mime/mediatype.go:91
				return character == '\\'
//line /usr/local/go/src/mime/mediatype.go:91
				// _ = "end of CoverTab[35646]"
//line /usr/local/go/src/mime/mediatype.go:91
			}() {
//line /usr/local/go/src/mime/mediatype.go:91
				_go_fuzz_dep_.CoverTab[35647]++
									b.WriteString(value[offset:index])
									offset = index
									b.WriteByte('\\')
//line /usr/local/go/src/mime/mediatype.go:94
				// _ = "end of CoverTab[35647]"
			} else {
//line /usr/local/go/src/mime/mediatype.go:95
				_go_fuzz_dep_.CoverTab[35648]++
//line /usr/local/go/src/mime/mediatype.go:95
				// _ = "end of CoverTab[35648]"
//line /usr/local/go/src/mime/mediatype.go:95
			}
//line /usr/local/go/src/mime/mediatype.go:95
			// _ = "end of CoverTab[35645]"
		}
//line /usr/local/go/src/mime/mediatype.go:96
		// _ = "end of CoverTab[35626]"
//line /usr/local/go/src/mime/mediatype.go:96
		_go_fuzz_dep_.CoverTab[35627]++
							b.WriteString(value[offset:])
							b.WriteByte('"')
//line /usr/local/go/src/mime/mediatype.go:98
		// _ = "end of CoverTab[35627]"
	}
//line /usr/local/go/src/mime/mediatype.go:99
	// _ = "end of CoverTab[35610]"
//line /usr/local/go/src/mime/mediatype.go:99
	_go_fuzz_dep_.CoverTab[35611]++
						return b.String()
//line /usr/local/go/src/mime/mediatype.go:100
	// _ = "end of CoverTab[35611]"
}

func checkMediaTypeDisposition(s string) error {
//line /usr/local/go/src/mime/mediatype.go:103
	_go_fuzz_dep_.CoverTab[35649]++
						typ, rest := consumeToken(s)
						if typ == "" {
//line /usr/local/go/src/mime/mediatype.go:105
		_go_fuzz_dep_.CoverTab[35655]++
							return errors.New("mime: no media type")
//line /usr/local/go/src/mime/mediatype.go:106
		// _ = "end of CoverTab[35655]"
	} else {
//line /usr/local/go/src/mime/mediatype.go:107
		_go_fuzz_dep_.CoverTab[35656]++
//line /usr/local/go/src/mime/mediatype.go:107
		// _ = "end of CoverTab[35656]"
//line /usr/local/go/src/mime/mediatype.go:107
	}
//line /usr/local/go/src/mime/mediatype.go:107
	// _ = "end of CoverTab[35649]"
//line /usr/local/go/src/mime/mediatype.go:107
	_go_fuzz_dep_.CoverTab[35650]++
						if rest == "" {
//line /usr/local/go/src/mime/mediatype.go:108
		_go_fuzz_dep_.CoverTab[35657]++
							return nil
//line /usr/local/go/src/mime/mediatype.go:109
		// _ = "end of CoverTab[35657]"
	} else {
//line /usr/local/go/src/mime/mediatype.go:110
		_go_fuzz_dep_.CoverTab[35658]++
//line /usr/local/go/src/mime/mediatype.go:110
		// _ = "end of CoverTab[35658]"
//line /usr/local/go/src/mime/mediatype.go:110
	}
//line /usr/local/go/src/mime/mediatype.go:110
	// _ = "end of CoverTab[35650]"
//line /usr/local/go/src/mime/mediatype.go:110
	_go_fuzz_dep_.CoverTab[35651]++
						if !strings.HasPrefix(rest, "/") {
//line /usr/local/go/src/mime/mediatype.go:111
		_go_fuzz_dep_.CoverTab[35659]++
							return errors.New("mime: expected slash after first token")
//line /usr/local/go/src/mime/mediatype.go:112
		// _ = "end of CoverTab[35659]"
	} else {
//line /usr/local/go/src/mime/mediatype.go:113
		_go_fuzz_dep_.CoverTab[35660]++
//line /usr/local/go/src/mime/mediatype.go:113
		// _ = "end of CoverTab[35660]"
//line /usr/local/go/src/mime/mediatype.go:113
	}
//line /usr/local/go/src/mime/mediatype.go:113
	// _ = "end of CoverTab[35651]"
//line /usr/local/go/src/mime/mediatype.go:113
	_go_fuzz_dep_.CoverTab[35652]++
						subtype, rest := consumeToken(rest[1:])
						if subtype == "" {
//line /usr/local/go/src/mime/mediatype.go:115
		_go_fuzz_dep_.CoverTab[35661]++
							return errors.New("mime: expected token after slash")
//line /usr/local/go/src/mime/mediatype.go:116
		// _ = "end of CoverTab[35661]"
	} else {
//line /usr/local/go/src/mime/mediatype.go:117
		_go_fuzz_dep_.CoverTab[35662]++
//line /usr/local/go/src/mime/mediatype.go:117
		// _ = "end of CoverTab[35662]"
//line /usr/local/go/src/mime/mediatype.go:117
	}
//line /usr/local/go/src/mime/mediatype.go:117
	// _ = "end of CoverTab[35652]"
//line /usr/local/go/src/mime/mediatype.go:117
	_go_fuzz_dep_.CoverTab[35653]++
						if rest != "" {
//line /usr/local/go/src/mime/mediatype.go:118
		_go_fuzz_dep_.CoverTab[35663]++
							return errors.New("mime: unexpected content after media subtype")
//line /usr/local/go/src/mime/mediatype.go:119
		// _ = "end of CoverTab[35663]"
	} else {
//line /usr/local/go/src/mime/mediatype.go:120
		_go_fuzz_dep_.CoverTab[35664]++
//line /usr/local/go/src/mime/mediatype.go:120
		// _ = "end of CoverTab[35664]"
//line /usr/local/go/src/mime/mediatype.go:120
	}
//line /usr/local/go/src/mime/mediatype.go:120
	// _ = "end of CoverTab[35653]"
//line /usr/local/go/src/mime/mediatype.go:120
	_go_fuzz_dep_.CoverTab[35654]++
						return nil
//line /usr/local/go/src/mime/mediatype.go:121
	// _ = "end of CoverTab[35654]"
}

// ErrInvalidMediaParameter is returned by ParseMediaType if
//line /usr/local/go/src/mime/mediatype.go:124
// the media type value was found but there was an error parsing
//line /usr/local/go/src/mime/mediatype.go:124
// the optional parameters
//line /usr/local/go/src/mime/mediatype.go:127
var ErrInvalidMediaParameter = errors.New("mime: invalid media parameter")

// ParseMediaType parses a media type value and any optional
//line /usr/local/go/src/mime/mediatype.go:129
// parameters, per RFC 1521.  Media types are the values in
//line /usr/local/go/src/mime/mediatype.go:129
// Content-Type and Content-Disposition headers (RFC 2183).
//line /usr/local/go/src/mime/mediatype.go:129
// On success, ParseMediaType returns the media type converted
//line /usr/local/go/src/mime/mediatype.go:129
// to lowercase and trimmed of white space and a non-nil map.
//line /usr/local/go/src/mime/mediatype.go:129
// If there is an error parsing the optional parameter,
//line /usr/local/go/src/mime/mediatype.go:129
// the media type will be returned along with the error
//line /usr/local/go/src/mime/mediatype.go:129
// ErrInvalidMediaParameter.
//line /usr/local/go/src/mime/mediatype.go:129
// The returned map, params, maps from the lowercase
//line /usr/local/go/src/mime/mediatype.go:129
// attribute to the attribute value with its case preserved.
//line /usr/local/go/src/mime/mediatype.go:139
func ParseMediaType(v string) (mediatype string, params map[string]string, err error) {
//line /usr/local/go/src/mime/mediatype.go:139
	_go_fuzz_dep_.CoverTab[35665]++
						base, _, _ := strings.Cut(v, ";")
						mediatype = strings.TrimSpace(strings.ToLower(base))

						err = checkMediaTypeDisposition(mediatype)
						if err != nil {
//line /usr/local/go/src/mime/mediatype.go:144
		_go_fuzz_dep_.CoverTab[35669]++
							return "", nil, err
//line /usr/local/go/src/mime/mediatype.go:145
		// _ = "end of CoverTab[35669]"
	} else {
//line /usr/local/go/src/mime/mediatype.go:146
		_go_fuzz_dep_.CoverTab[35670]++
//line /usr/local/go/src/mime/mediatype.go:146
		// _ = "end of CoverTab[35670]"
//line /usr/local/go/src/mime/mediatype.go:146
	}
//line /usr/local/go/src/mime/mediatype.go:146
	// _ = "end of CoverTab[35665]"
//line /usr/local/go/src/mime/mediatype.go:146
	_go_fuzz_dep_.CoverTab[35666]++

						params = make(map[string]string)

	// Map of base parameter name -> parameter name -> value
	// for parameters containing a '*' character.
	// Lazily initialized.
	var continuation map[string]map[string]string

	v = v[len(base):]
	for len(v) > 0 {
//line /usr/local/go/src/mime/mediatype.go:156
		_go_fuzz_dep_.CoverTab[35671]++
							v = strings.TrimLeftFunc(v, unicode.IsSpace)
							if len(v) == 0 {
//line /usr/local/go/src/mime/mediatype.go:158
			_go_fuzz_dep_.CoverTab[35676]++
								break
//line /usr/local/go/src/mime/mediatype.go:159
			// _ = "end of CoverTab[35676]"
		} else {
//line /usr/local/go/src/mime/mediatype.go:160
			_go_fuzz_dep_.CoverTab[35677]++
//line /usr/local/go/src/mime/mediatype.go:160
			// _ = "end of CoverTab[35677]"
//line /usr/local/go/src/mime/mediatype.go:160
		}
//line /usr/local/go/src/mime/mediatype.go:160
		// _ = "end of CoverTab[35671]"
//line /usr/local/go/src/mime/mediatype.go:160
		_go_fuzz_dep_.CoverTab[35672]++
							key, value, rest := consumeMediaParam(v)
							if key == "" {
//line /usr/local/go/src/mime/mediatype.go:162
			_go_fuzz_dep_.CoverTab[35678]++
								if strings.TrimSpace(rest) == ";" {
//line /usr/local/go/src/mime/mediatype.go:163
				_go_fuzz_dep_.CoverTab[35680]++

//line /usr/local/go/src/mime/mediatype.go:166
				break
//line /usr/local/go/src/mime/mediatype.go:166
				// _ = "end of CoverTab[35680]"
			} else {
//line /usr/local/go/src/mime/mediatype.go:167
				_go_fuzz_dep_.CoverTab[35681]++
//line /usr/local/go/src/mime/mediatype.go:167
				// _ = "end of CoverTab[35681]"
//line /usr/local/go/src/mime/mediatype.go:167
			}
//line /usr/local/go/src/mime/mediatype.go:167
			// _ = "end of CoverTab[35678]"
//line /usr/local/go/src/mime/mediatype.go:167
			_go_fuzz_dep_.CoverTab[35679]++

								return mediatype, nil, ErrInvalidMediaParameter
//line /usr/local/go/src/mime/mediatype.go:169
			// _ = "end of CoverTab[35679]"
		} else {
//line /usr/local/go/src/mime/mediatype.go:170
			_go_fuzz_dep_.CoverTab[35682]++
//line /usr/local/go/src/mime/mediatype.go:170
			// _ = "end of CoverTab[35682]"
//line /usr/local/go/src/mime/mediatype.go:170
		}
//line /usr/local/go/src/mime/mediatype.go:170
		// _ = "end of CoverTab[35672]"
//line /usr/local/go/src/mime/mediatype.go:170
		_go_fuzz_dep_.CoverTab[35673]++

							pmap := params
							if baseName, _, ok := strings.Cut(key, "*"); ok {
//line /usr/local/go/src/mime/mediatype.go:173
			_go_fuzz_dep_.CoverTab[35683]++
								if continuation == nil {
//line /usr/local/go/src/mime/mediatype.go:174
				_go_fuzz_dep_.CoverTab[35685]++
									continuation = make(map[string]map[string]string)
//line /usr/local/go/src/mime/mediatype.go:175
				// _ = "end of CoverTab[35685]"
			} else {
//line /usr/local/go/src/mime/mediatype.go:176
				_go_fuzz_dep_.CoverTab[35686]++
//line /usr/local/go/src/mime/mediatype.go:176
				// _ = "end of CoverTab[35686]"
//line /usr/local/go/src/mime/mediatype.go:176
			}
//line /usr/local/go/src/mime/mediatype.go:176
			// _ = "end of CoverTab[35683]"
//line /usr/local/go/src/mime/mediatype.go:176
			_go_fuzz_dep_.CoverTab[35684]++
								var ok bool
								if pmap, ok = continuation[baseName]; !ok {
//line /usr/local/go/src/mime/mediatype.go:178
				_go_fuzz_dep_.CoverTab[35687]++
									continuation[baseName] = make(map[string]string)
									pmap = continuation[baseName]
//line /usr/local/go/src/mime/mediatype.go:180
				// _ = "end of CoverTab[35687]"
			} else {
//line /usr/local/go/src/mime/mediatype.go:181
				_go_fuzz_dep_.CoverTab[35688]++
//line /usr/local/go/src/mime/mediatype.go:181
				// _ = "end of CoverTab[35688]"
//line /usr/local/go/src/mime/mediatype.go:181
			}
//line /usr/local/go/src/mime/mediatype.go:181
			// _ = "end of CoverTab[35684]"
		} else {
//line /usr/local/go/src/mime/mediatype.go:182
			_go_fuzz_dep_.CoverTab[35689]++
//line /usr/local/go/src/mime/mediatype.go:182
			// _ = "end of CoverTab[35689]"
//line /usr/local/go/src/mime/mediatype.go:182
		}
//line /usr/local/go/src/mime/mediatype.go:182
		// _ = "end of CoverTab[35673]"
//line /usr/local/go/src/mime/mediatype.go:182
		_go_fuzz_dep_.CoverTab[35674]++
							if v, exists := pmap[key]; exists && func() bool {
//line /usr/local/go/src/mime/mediatype.go:183
			_go_fuzz_dep_.CoverTab[35690]++
//line /usr/local/go/src/mime/mediatype.go:183
			return v != value
//line /usr/local/go/src/mime/mediatype.go:183
			// _ = "end of CoverTab[35690]"
//line /usr/local/go/src/mime/mediatype.go:183
		}() {
//line /usr/local/go/src/mime/mediatype.go:183
			_go_fuzz_dep_.CoverTab[35691]++

								return "", nil, errors.New("mime: duplicate parameter name")
//line /usr/local/go/src/mime/mediatype.go:185
			// _ = "end of CoverTab[35691]"
		} else {
//line /usr/local/go/src/mime/mediatype.go:186
			_go_fuzz_dep_.CoverTab[35692]++
//line /usr/local/go/src/mime/mediatype.go:186
			// _ = "end of CoverTab[35692]"
//line /usr/local/go/src/mime/mediatype.go:186
		}
//line /usr/local/go/src/mime/mediatype.go:186
		// _ = "end of CoverTab[35674]"
//line /usr/local/go/src/mime/mediatype.go:186
		_go_fuzz_dep_.CoverTab[35675]++
							pmap[key] = value
							v = rest
//line /usr/local/go/src/mime/mediatype.go:188
		// _ = "end of CoverTab[35675]"
	}
//line /usr/local/go/src/mime/mediatype.go:189
	// _ = "end of CoverTab[35666]"
//line /usr/local/go/src/mime/mediatype.go:189
	_go_fuzz_dep_.CoverTab[35667]++

	// Stitch together any continuations or things with stars
	// (i.e. RFC 2231 things with stars: "foo*0" or "foo*")
	var buf strings.Builder
	for key, pieceMap := range continuation {
//line /usr/local/go/src/mime/mediatype.go:194
		_go_fuzz_dep_.CoverTab[35693]++
							singlePartKey := key + "*"
							if v, ok := pieceMap[singlePartKey]; ok {
//line /usr/local/go/src/mime/mediatype.go:196
			_go_fuzz_dep_.CoverTab[35696]++
								if decv, ok := decode2231Enc(v); ok {
//line /usr/local/go/src/mime/mediatype.go:197
				_go_fuzz_dep_.CoverTab[35698]++
									params[key] = decv
//line /usr/local/go/src/mime/mediatype.go:198
				// _ = "end of CoverTab[35698]"
			} else {
//line /usr/local/go/src/mime/mediatype.go:199
				_go_fuzz_dep_.CoverTab[35699]++
//line /usr/local/go/src/mime/mediatype.go:199
				// _ = "end of CoverTab[35699]"
//line /usr/local/go/src/mime/mediatype.go:199
			}
//line /usr/local/go/src/mime/mediatype.go:199
			// _ = "end of CoverTab[35696]"
//line /usr/local/go/src/mime/mediatype.go:199
			_go_fuzz_dep_.CoverTab[35697]++
								continue
//line /usr/local/go/src/mime/mediatype.go:200
			// _ = "end of CoverTab[35697]"
		} else {
//line /usr/local/go/src/mime/mediatype.go:201
			_go_fuzz_dep_.CoverTab[35700]++
//line /usr/local/go/src/mime/mediatype.go:201
			// _ = "end of CoverTab[35700]"
//line /usr/local/go/src/mime/mediatype.go:201
		}
//line /usr/local/go/src/mime/mediatype.go:201
		// _ = "end of CoverTab[35693]"
//line /usr/local/go/src/mime/mediatype.go:201
		_go_fuzz_dep_.CoverTab[35694]++

							buf.Reset()
							valid := false
							for n := 0; ; n++ {
//line /usr/local/go/src/mime/mediatype.go:205
			_go_fuzz_dep_.CoverTab[35701]++
								simplePart := fmt.Sprintf("%s*%d", key, n)
								if v, ok := pieceMap[simplePart]; ok {
//line /usr/local/go/src/mime/mediatype.go:207
				_go_fuzz_dep_.CoverTab[35704]++
									valid = true
									buf.WriteString(v)
									continue
//line /usr/local/go/src/mime/mediatype.go:210
				// _ = "end of CoverTab[35704]"
			} else {
//line /usr/local/go/src/mime/mediatype.go:211
				_go_fuzz_dep_.CoverTab[35705]++
//line /usr/local/go/src/mime/mediatype.go:211
				// _ = "end of CoverTab[35705]"
//line /usr/local/go/src/mime/mediatype.go:211
			}
//line /usr/local/go/src/mime/mediatype.go:211
			// _ = "end of CoverTab[35701]"
//line /usr/local/go/src/mime/mediatype.go:211
			_go_fuzz_dep_.CoverTab[35702]++
								encodedPart := simplePart + "*"
								v, ok := pieceMap[encodedPart]
								if !ok {
//line /usr/local/go/src/mime/mediatype.go:214
				_go_fuzz_dep_.CoverTab[35706]++
									break
//line /usr/local/go/src/mime/mediatype.go:215
				// _ = "end of CoverTab[35706]"
			} else {
//line /usr/local/go/src/mime/mediatype.go:216
				_go_fuzz_dep_.CoverTab[35707]++
//line /usr/local/go/src/mime/mediatype.go:216
				// _ = "end of CoverTab[35707]"
//line /usr/local/go/src/mime/mediatype.go:216
			}
//line /usr/local/go/src/mime/mediatype.go:216
			// _ = "end of CoverTab[35702]"
//line /usr/local/go/src/mime/mediatype.go:216
			_go_fuzz_dep_.CoverTab[35703]++
								valid = true
								if n == 0 {
//line /usr/local/go/src/mime/mediatype.go:218
				_go_fuzz_dep_.CoverTab[35708]++
									if decv, ok := decode2231Enc(v); ok {
//line /usr/local/go/src/mime/mediatype.go:219
					_go_fuzz_dep_.CoverTab[35709]++
										buf.WriteString(decv)
//line /usr/local/go/src/mime/mediatype.go:220
					// _ = "end of CoverTab[35709]"
				} else {
//line /usr/local/go/src/mime/mediatype.go:221
					_go_fuzz_dep_.CoverTab[35710]++
//line /usr/local/go/src/mime/mediatype.go:221
					// _ = "end of CoverTab[35710]"
//line /usr/local/go/src/mime/mediatype.go:221
				}
//line /usr/local/go/src/mime/mediatype.go:221
				// _ = "end of CoverTab[35708]"
			} else {
//line /usr/local/go/src/mime/mediatype.go:222
				_go_fuzz_dep_.CoverTab[35711]++
									decv, _ := percentHexUnescape(v)
									buf.WriteString(decv)
//line /usr/local/go/src/mime/mediatype.go:224
				// _ = "end of CoverTab[35711]"
			}
//line /usr/local/go/src/mime/mediatype.go:225
			// _ = "end of CoverTab[35703]"
		}
//line /usr/local/go/src/mime/mediatype.go:226
		// _ = "end of CoverTab[35694]"
//line /usr/local/go/src/mime/mediatype.go:226
		_go_fuzz_dep_.CoverTab[35695]++
							if valid {
//line /usr/local/go/src/mime/mediatype.go:227
			_go_fuzz_dep_.CoverTab[35712]++
								params[key] = buf.String()
//line /usr/local/go/src/mime/mediatype.go:228
			// _ = "end of CoverTab[35712]"
		} else {
//line /usr/local/go/src/mime/mediatype.go:229
			_go_fuzz_dep_.CoverTab[35713]++
//line /usr/local/go/src/mime/mediatype.go:229
			// _ = "end of CoverTab[35713]"
//line /usr/local/go/src/mime/mediatype.go:229
		}
//line /usr/local/go/src/mime/mediatype.go:229
		// _ = "end of CoverTab[35695]"
	}
//line /usr/local/go/src/mime/mediatype.go:230
	// _ = "end of CoverTab[35667]"
//line /usr/local/go/src/mime/mediatype.go:230
	_go_fuzz_dep_.CoverTab[35668]++

						return
//line /usr/local/go/src/mime/mediatype.go:232
	// _ = "end of CoverTab[35668]"
}

func decode2231Enc(v string) (string, bool) {
//line /usr/local/go/src/mime/mediatype.go:235
	_go_fuzz_dep_.CoverTab[35714]++
						sv := strings.SplitN(v, "'", 3)
						if len(sv) != 3 {
//line /usr/local/go/src/mime/mediatype.go:237
		_go_fuzz_dep_.CoverTab[35719]++
							return "", false
//line /usr/local/go/src/mime/mediatype.go:238
		// _ = "end of CoverTab[35719]"
	} else {
//line /usr/local/go/src/mime/mediatype.go:239
		_go_fuzz_dep_.CoverTab[35720]++
//line /usr/local/go/src/mime/mediatype.go:239
		// _ = "end of CoverTab[35720]"
//line /usr/local/go/src/mime/mediatype.go:239
	}
//line /usr/local/go/src/mime/mediatype.go:239
	// _ = "end of CoverTab[35714]"
//line /usr/local/go/src/mime/mediatype.go:239
	_go_fuzz_dep_.CoverTab[35715]++

//line /usr/local/go/src/mime/mediatype.go:243
	charset := strings.ToLower(sv[0])
	if len(charset) == 0 {
//line /usr/local/go/src/mime/mediatype.go:244
		_go_fuzz_dep_.CoverTab[35721]++
							return "", false
//line /usr/local/go/src/mime/mediatype.go:245
		// _ = "end of CoverTab[35721]"
	} else {
//line /usr/local/go/src/mime/mediatype.go:246
		_go_fuzz_dep_.CoverTab[35722]++
//line /usr/local/go/src/mime/mediatype.go:246
		// _ = "end of CoverTab[35722]"
//line /usr/local/go/src/mime/mediatype.go:246
	}
//line /usr/local/go/src/mime/mediatype.go:246
	// _ = "end of CoverTab[35715]"
//line /usr/local/go/src/mime/mediatype.go:246
	_go_fuzz_dep_.CoverTab[35716]++
						if charset != "us-ascii" && func() bool {
//line /usr/local/go/src/mime/mediatype.go:247
		_go_fuzz_dep_.CoverTab[35723]++
//line /usr/local/go/src/mime/mediatype.go:247
		return charset != "utf-8"
//line /usr/local/go/src/mime/mediatype.go:247
		// _ = "end of CoverTab[35723]"
//line /usr/local/go/src/mime/mediatype.go:247
	}() {
//line /usr/local/go/src/mime/mediatype.go:247
		_go_fuzz_dep_.CoverTab[35724]++

							return "", false
//line /usr/local/go/src/mime/mediatype.go:249
		// _ = "end of CoverTab[35724]"
	} else {
//line /usr/local/go/src/mime/mediatype.go:250
		_go_fuzz_dep_.CoverTab[35725]++
//line /usr/local/go/src/mime/mediatype.go:250
		// _ = "end of CoverTab[35725]"
//line /usr/local/go/src/mime/mediatype.go:250
	}
//line /usr/local/go/src/mime/mediatype.go:250
	// _ = "end of CoverTab[35716]"
//line /usr/local/go/src/mime/mediatype.go:250
	_go_fuzz_dep_.CoverTab[35717]++
						encv, err := percentHexUnescape(sv[2])
						if err != nil {
//line /usr/local/go/src/mime/mediatype.go:252
		_go_fuzz_dep_.CoverTab[35726]++
							return "", false
//line /usr/local/go/src/mime/mediatype.go:253
		// _ = "end of CoverTab[35726]"
	} else {
//line /usr/local/go/src/mime/mediatype.go:254
		_go_fuzz_dep_.CoverTab[35727]++
//line /usr/local/go/src/mime/mediatype.go:254
		// _ = "end of CoverTab[35727]"
//line /usr/local/go/src/mime/mediatype.go:254
	}
//line /usr/local/go/src/mime/mediatype.go:254
	// _ = "end of CoverTab[35717]"
//line /usr/local/go/src/mime/mediatype.go:254
	_go_fuzz_dep_.CoverTab[35718]++
						return encv, true
//line /usr/local/go/src/mime/mediatype.go:255
	// _ = "end of CoverTab[35718]"
}

func isNotTokenChar(r rune) bool {
//line /usr/local/go/src/mime/mediatype.go:258
	_go_fuzz_dep_.CoverTab[35728]++
						return !isTokenChar(r)
//line /usr/local/go/src/mime/mediatype.go:259
	// _ = "end of CoverTab[35728]"
}

// consumeToken consumes a token from the beginning of provided
//line /usr/local/go/src/mime/mediatype.go:262
// string, per RFC 2045 section 5.1 (referenced from 2183), and return
//line /usr/local/go/src/mime/mediatype.go:262
// the token consumed and the rest of the string. Returns ("", v) on
//line /usr/local/go/src/mime/mediatype.go:262
// failure to consume at least one character.
//line /usr/local/go/src/mime/mediatype.go:266
func consumeToken(v string) (token, rest string) {
//line /usr/local/go/src/mime/mediatype.go:266
	_go_fuzz_dep_.CoverTab[35729]++
						notPos := strings.IndexFunc(v, isNotTokenChar)
						if notPos == -1 {
//line /usr/local/go/src/mime/mediatype.go:268
		_go_fuzz_dep_.CoverTab[35732]++
							return v, ""
//line /usr/local/go/src/mime/mediatype.go:269
		// _ = "end of CoverTab[35732]"
	} else {
//line /usr/local/go/src/mime/mediatype.go:270
		_go_fuzz_dep_.CoverTab[35733]++
//line /usr/local/go/src/mime/mediatype.go:270
		// _ = "end of CoverTab[35733]"
//line /usr/local/go/src/mime/mediatype.go:270
	}
//line /usr/local/go/src/mime/mediatype.go:270
	// _ = "end of CoverTab[35729]"
//line /usr/local/go/src/mime/mediatype.go:270
	_go_fuzz_dep_.CoverTab[35730]++
						if notPos == 0 {
//line /usr/local/go/src/mime/mediatype.go:271
		_go_fuzz_dep_.CoverTab[35734]++
							return "", v
//line /usr/local/go/src/mime/mediatype.go:272
		// _ = "end of CoverTab[35734]"
	} else {
//line /usr/local/go/src/mime/mediatype.go:273
		_go_fuzz_dep_.CoverTab[35735]++
//line /usr/local/go/src/mime/mediatype.go:273
		// _ = "end of CoverTab[35735]"
//line /usr/local/go/src/mime/mediatype.go:273
	}
//line /usr/local/go/src/mime/mediatype.go:273
	// _ = "end of CoverTab[35730]"
//line /usr/local/go/src/mime/mediatype.go:273
	_go_fuzz_dep_.CoverTab[35731]++
						return v[0:notPos], v[notPos:]
//line /usr/local/go/src/mime/mediatype.go:274
	// _ = "end of CoverTab[35731]"
}

// consumeValue consumes a "value" per RFC 2045, where a value is
//line /usr/local/go/src/mime/mediatype.go:277
// either a 'token' or a 'quoted-string'.  On success, consumeValue
//line /usr/local/go/src/mime/mediatype.go:277
// returns the value consumed (and de-quoted/escaped, if a
//line /usr/local/go/src/mime/mediatype.go:277
// quoted-string) and the rest of the string. On failure, returns
//line /usr/local/go/src/mime/mediatype.go:277
// ("", v).
//line /usr/local/go/src/mime/mediatype.go:282
func consumeValue(v string) (value, rest string) {
//line /usr/local/go/src/mime/mediatype.go:282
	_go_fuzz_dep_.CoverTab[35736]++
						if v == "" {
//line /usr/local/go/src/mime/mediatype.go:283
		_go_fuzz_dep_.CoverTab[35740]++
							return
//line /usr/local/go/src/mime/mediatype.go:284
		// _ = "end of CoverTab[35740]"
	} else {
//line /usr/local/go/src/mime/mediatype.go:285
		_go_fuzz_dep_.CoverTab[35741]++
//line /usr/local/go/src/mime/mediatype.go:285
		// _ = "end of CoverTab[35741]"
//line /usr/local/go/src/mime/mediatype.go:285
	}
//line /usr/local/go/src/mime/mediatype.go:285
	// _ = "end of CoverTab[35736]"
//line /usr/local/go/src/mime/mediatype.go:285
	_go_fuzz_dep_.CoverTab[35737]++
						if v[0] != '"' {
//line /usr/local/go/src/mime/mediatype.go:286
		_go_fuzz_dep_.CoverTab[35742]++
							return consumeToken(v)
//line /usr/local/go/src/mime/mediatype.go:287
		// _ = "end of CoverTab[35742]"
	} else {
//line /usr/local/go/src/mime/mediatype.go:288
		_go_fuzz_dep_.CoverTab[35743]++
//line /usr/local/go/src/mime/mediatype.go:288
		// _ = "end of CoverTab[35743]"
//line /usr/local/go/src/mime/mediatype.go:288
	}
//line /usr/local/go/src/mime/mediatype.go:288
	// _ = "end of CoverTab[35737]"
//line /usr/local/go/src/mime/mediatype.go:288
	_go_fuzz_dep_.CoverTab[35738]++

//line /usr/local/go/src/mime/mediatype.go:291
	buffer := new(strings.Builder)
	for i := 1; i < len(v); i++ {
//line /usr/local/go/src/mime/mediatype.go:292
		_go_fuzz_dep_.CoverTab[35744]++
							r := v[i]
							if r == '"' {
//line /usr/local/go/src/mime/mediatype.go:294
			_go_fuzz_dep_.CoverTab[35748]++
								return buffer.String(), v[i+1:]
//line /usr/local/go/src/mime/mediatype.go:295
			// _ = "end of CoverTab[35748]"
		} else {
//line /usr/local/go/src/mime/mediatype.go:296
			_go_fuzz_dep_.CoverTab[35749]++
//line /usr/local/go/src/mime/mediatype.go:296
			// _ = "end of CoverTab[35749]"
//line /usr/local/go/src/mime/mediatype.go:296
		}
//line /usr/local/go/src/mime/mediatype.go:296
		// _ = "end of CoverTab[35744]"
//line /usr/local/go/src/mime/mediatype.go:296
		_go_fuzz_dep_.CoverTab[35745]++

//line /usr/local/go/src/mime/mediatype.go:307
		if r == '\\' && func() bool {
//line /usr/local/go/src/mime/mediatype.go:307
			_go_fuzz_dep_.CoverTab[35750]++
//line /usr/local/go/src/mime/mediatype.go:307
			return i+1 < len(v)
//line /usr/local/go/src/mime/mediatype.go:307
			// _ = "end of CoverTab[35750]"
//line /usr/local/go/src/mime/mediatype.go:307
		}() && func() bool {
//line /usr/local/go/src/mime/mediatype.go:307
			_go_fuzz_dep_.CoverTab[35751]++
//line /usr/local/go/src/mime/mediatype.go:307
			return isTSpecial(rune(v[i+1]))
//line /usr/local/go/src/mime/mediatype.go:307
			// _ = "end of CoverTab[35751]"
//line /usr/local/go/src/mime/mediatype.go:307
		}() {
//line /usr/local/go/src/mime/mediatype.go:307
			_go_fuzz_dep_.CoverTab[35752]++
								buffer.WriteByte(v[i+1])
								i++
								continue
//line /usr/local/go/src/mime/mediatype.go:310
			// _ = "end of CoverTab[35752]"
		} else {
//line /usr/local/go/src/mime/mediatype.go:311
			_go_fuzz_dep_.CoverTab[35753]++
//line /usr/local/go/src/mime/mediatype.go:311
			// _ = "end of CoverTab[35753]"
//line /usr/local/go/src/mime/mediatype.go:311
		}
//line /usr/local/go/src/mime/mediatype.go:311
		// _ = "end of CoverTab[35745]"
//line /usr/local/go/src/mime/mediatype.go:311
		_go_fuzz_dep_.CoverTab[35746]++
							if r == '\r' || func() bool {
//line /usr/local/go/src/mime/mediatype.go:312
			_go_fuzz_dep_.CoverTab[35754]++
//line /usr/local/go/src/mime/mediatype.go:312
			return r == '\n'
//line /usr/local/go/src/mime/mediatype.go:312
			// _ = "end of CoverTab[35754]"
//line /usr/local/go/src/mime/mediatype.go:312
		}() {
//line /usr/local/go/src/mime/mediatype.go:312
			_go_fuzz_dep_.CoverTab[35755]++
								return "", v
//line /usr/local/go/src/mime/mediatype.go:313
			// _ = "end of CoverTab[35755]"
		} else {
//line /usr/local/go/src/mime/mediatype.go:314
			_go_fuzz_dep_.CoverTab[35756]++
//line /usr/local/go/src/mime/mediatype.go:314
			// _ = "end of CoverTab[35756]"
//line /usr/local/go/src/mime/mediatype.go:314
		}
//line /usr/local/go/src/mime/mediatype.go:314
		// _ = "end of CoverTab[35746]"
//line /usr/local/go/src/mime/mediatype.go:314
		_go_fuzz_dep_.CoverTab[35747]++
							buffer.WriteByte(v[i])
//line /usr/local/go/src/mime/mediatype.go:315
		// _ = "end of CoverTab[35747]"
	}
//line /usr/local/go/src/mime/mediatype.go:316
	// _ = "end of CoverTab[35738]"
//line /usr/local/go/src/mime/mediatype.go:316
	_go_fuzz_dep_.CoverTab[35739]++

						return "", v
//line /usr/local/go/src/mime/mediatype.go:318
	// _ = "end of CoverTab[35739]"
}

func consumeMediaParam(v string) (param, value, rest string) {
//line /usr/local/go/src/mime/mediatype.go:321
	_go_fuzz_dep_.CoverTab[35757]++
						rest = strings.TrimLeftFunc(v, unicode.IsSpace)
						if !strings.HasPrefix(rest, ";") {
//line /usr/local/go/src/mime/mediatype.go:323
		_go_fuzz_dep_.CoverTab[35762]++
							return "", "", v
//line /usr/local/go/src/mime/mediatype.go:324
		// _ = "end of CoverTab[35762]"
	} else {
//line /usr/local/go/src/mime/mediatype.go:325
		_go_fuzz_dep_.CoverTab[35763]++
//line /usr/local/go/src/mime/mediatype.go:325
		// _ = "end of CoverTab[35763]"
//line /usr/local/go/src/mime/mediatype.go:325
	}
//line /usr/local/go/src/mime/mediatype.go:325
	// _ = "end of CoverTab[35757]"
//line /usr/local/go/src/mime/mediatype.go:325
	_go_fuzz_dep_.CoverTab[35758]++

						rest = rest[1:]
						rest = strings.TrimLeftFunc(rest, unicode.IsSpace)
						param, rest = consumeToken(rest)
						param = strings.ToLower(param)
						if param == "" {
//line /usr/local/go/src/mime/mediatype.go:331
		_go_fuzz_dep_.CoverTab[35764]++
							return "", "", v
//line /usr/local/go/src/mime/mediatype.go:332
		// _ = "end of CoverTab[35764]"
	} else {
//line /usr/local/go/src/mime/mediatype.go:333
		_go_fuzz_dep_.CoverTab[35765]++
//line /usr/local/go/src/mime/mediatype.go:333
		// _ = "end of CoverTab[35765]"
//line /usr/local/go/src/mime/mediatype.go:333
	}
//line /usr/local/go/src/mime/mediatype.go:333
	// _ = "end of CoverTab[35758]"
//line /usr/local/go/src/mime/mediatype.go:333
	_go_fuzz_dep_.CoverTab[35759]++

						rest = strings.TrimLeftFunc(rest, unicode.IsSpace)
						if !strings.HasPrefix(rest, "=") {
//line /usr/local/go/src/mime/mediatype.go:336
		_go_fuzz_dep_.CoverTab[35766]++
							return "", "", v
//line /usr/local/go/src/mime/mediatype.go:337
		// _ = "end of CoverTab[35766]"
	} else {
//line /usr/local/go/src/mime/mediatype.go:338
		_go_fuzz_dep_.CoverTab[35767]++
//line /usr/local/go/src/mime/mediatype.go:338
		// _ = "end of CoverTab[35767]"
//line /usr/local/go/src/mime/mediatype.go:338
	}
//line /usr/local/go/src/mime/mediatype.go:338
	// _ = "end of CoverTab[35759]"
//line /usr/local/go/src/mime/mediatype.go:338
	_go_fuzz_dep_.CoverTab[35760]++
						rest = rest[1:]
						rest = strings.TrimLeftFunc(rest, unicode.IsSpace)
						value, rest2 := consumeValue(rest)
						if value == "" && func() bool {
//line /usr/local/go/src/mime/mediatype.go:342
		_go_fuzz_dep_.CoverTab[35768]++
//line /usr/local/go/src/mime/mediatype.go:342
		return rest2 == rest
//line /usr/local/go/src/mime/mediatype.go:342
		// _ = "end of CoverTab[35768]"
//line /usr/local/go/src/mime/mediatype.go:342
	}() {
//line /usr/local/go/src/mime/mediatype.go:342
		_go_fuzz_dep_.CoverTab[35769]++
							return "", "", v
//line /usr/local/go/src/mime/mediatype.go:343
		// _ = "end of CoverTab[35769]"
	} else {
//line /usr/local/go/src/mime/mediatype.go:344
		_go_fuzz_dep_.CoverTab[35770]++
//line /usr/local/go/src/mime/mediatype.go:344
		// _ = "end of CoverTab[35770]"
//line /usr/local/go/src/mime/mediatype.go:344
	}
//line /usr/local/go/src/mime/mediatype.go:344
	// _ = "end of CoverTab[35760]"
//line /usr/local/go/src/mime/mediatype.go:344
	_go_fuzz_dep_.CoverTab[35761]++
						rest = rest2
						return param, value, rest
//line /usr/local/go/src/mime/mediatype.go:346
	// _ = "end of CoverTab[35761]"
}

func percentHexUnescape(s string) (string, error) {
//line /usr/local/go/src/mime/mediatype.go:349
	_go_fuzz_dep_.CoverTab[35771]++

						percents := 0
						for i := 0; i < len(s); {
//line /usr/local/go/src/mime/mediatype.go:352
		_go_fuzz_dep_.CoverTab[35775]++
							if s[i] != '%' {
//line /usr/local/go/src/mime/mediatype.go:353
			_go_fuzz_dep_.CoverTab[35778]++
								i++
								continue
//line /usr/local/go/src/mime/mediatype.go:355
			// _ = "end of CoverTab[35778]"
		} else {
//line /usr/local/go/src/mime/mediatype.go:356
			_go_fuzz_dep_.CoverTab[35779]++
//line /usr/local/go/src/mime/mediatype.go:356
			// _ = "end of CoverTab[35779]"
//line /usr/local/go/src/mime/mediatype.go:356
		}
//line /usr/local/go/src/mime/mediatype.go:356
		// _ = "end of CoverTab[35775]"
//line /usr/local/go/src/mime/mediatype.go:356
		_go_fuzz_dep_.CoverTab[35776]++
							percents++
							if i+2 >= len(s) || func() bool {
//line /usr/local/go/src/mime/mediatype.go:358
			_go_fuzz_dep_.CoverTab[35780]++
//line /usr/local/go/src/mime/mediatype.go:358
			return !ishex(s[i+1])
//line /usr/local/go/src/mime/mediatype.go:358
			// _ = "end of CoverTab[35780]"
//line /usr/local/go/src/mime/mediatype.go:358
		}() || func() bool {
//line /usr/local/go/src/mime/mediatype.go:358
			_go_fuzz_dep_.CoverTab[35781]++
//line /usr/local/go/src/mime/mediatype.go:358
			return !ishex(s[i+2])
//line /usr/local/go/src/mime/mediatype.go:358
			// _ = "end of CoverTab[35781]"
//line /usr/local/go/src/mime/mediatype.go:358
		}() {
//line /usr/local/go/src/mime/mediatype.go:358
			_go_fuzz_dep_.CoverTab[35782]++
								s = s[i:]
								if len(s) > 3 {
//line /usr/local/go/src/mime/mediatype.go:360
				_go_fuzz_dep_.CoverTab[35784]++
									s = s[0:3]
//line /usr/local/go/src/mime/mediatype.go:361
				// _ = "end of CoverTab[35784]"
			} else {
//line /usr/local/go/src/mime/mediatype.go:362
				_go_fuzz_dep_.CoverTab[35785]++
//line /usr/local/go/src/mime/mediatype.go:362
				// _ = "end of CoverTab[35785]"
//line /usr/local/go/src/mime/mediatype.go:362
			}
//line /usr/local/go/src/mime/mediatype.go:362
			// _ = "end of CoverTab[35782]"
//line /usr/local/go/src/mime/mediatype.go:362
			_go_fuzz_dep_.CoverTab[35783]++
								return "", fmt.Errorf("mime: bogus characters after %%: %q", s)
//line /usr/local/go/src/mime/mediatype.go:363
			// _ = "end of CoverTab[35783]"
		} else {
//line /usr/local/go/src/mime/mediatype.go:364
			_go_fuzz_dep_.CoverTab[35786]++
//line /usr/local/go/src/mime/mediatype.go:364
			// _ = "end of CoverTab[35786]"
//line /usr/local/go/src/mime/mediatype.go:364
		}
//line /usr/local/go/src/mime/mediatype.go:364
		// _ = "end of CoverTab[35776]"
//line /usr/local/go/src/mime/mediatype.go:364
		_go_fuzz_dep_.CoverTab[35777]++
							i += 3
//line /usr/local/go/src/mime/mediatype.go:365
		// _ = "end of CoverTab[35777]"
	}
//line /usr/local/go/src/mime/mediatype.go:366
	// _ = "end of CoverTab[35771]"
//line /usr/local/go/src/mime/mediatype.go:366
	_go_fuzz_dep_.CoverTab[35772]++
						if percents == 0 {
//line /usr/local/go/src/mime/mediatype.go:367
		_go_fuzz_dep_.CoverTab[35787]++
							return s, nil
//line /usr/local/go/src/mime/mediatype.go:368
		// _ = "end of CoverTab[35787]"
	} else {
//line /usr/local/go/src/mime/mediatype.go:369
		_go_fuzz_dep_.CoverTab[35788]++
//line /usr/local/go/src/mime/mediatype.go:369
		// _ = "end of CoverTab[35788]"
//line /usr/local/go/src/mime/mediatype.go:369
	}
//line /usr/local/go/src/mime/mediatype.go:369
	// _ = "end of CoverTab[35772]"
//line /usr/local/go/src/mime/mediatype.go:369
	_go_fuzz_dep_.CoverTab[35773]++

						t := make([]byte, len(s)-2*percents)
						j := 0
						for i := 0; i < len(s); {
//line /usr/local/go/src/mime/mediatype.go:373
		_go_fuzz_dep_.CoverTab[35789]++
							switch s[i] {
		case '%':
//line /usr/local/go/src/mime/mediatype.go:375
			_go_fuzz_dep_.CoverTab[35790]++
								t[j] = unhex(s[i+1])<<4 | unhex(s[i+2])
								j++
								i += 3
//line /usr/local/go/src/mime/mediatype.go:378
			// _ = "end of CoverTab[35790]"
		default:
//line /usr/local/go/src/mime/mediatype.go:379
			_go_fuzz_dep_.CoverTab[35791]++
								t[j] = s[i]
								j++
								i++
//line /usr/local/go/src/mime/mediatype.go:382
			// _ = "end of CoverTab[35791]"
		}
//line /usr/local/go/src/mime/mediatype.go:383
		// _ = "end of CoverTab[35789]"
	}
//line /usr/local/go/src/mime/mediatype.go:384
	// _ = "end of CoverTab[35773]"
//line /usr/local/go/src/mime/mediatype.go:384
	_go_fuzz_dep_.CoverTab[35774]++
						return string(t), nil
//line /usr/local/go/src/mime/mediatype.go:385
	// _ = "end of CoverTab[35774]"
}

func ishex(c byte) bool {
//line /usr/local/go/src/mime/mediatype.go:388
	_go_fuzz_dep_.CoverTab[35792]++
						switch {
	case '0' <= c && func() bool {
//line /usr/local/go/src/mime/mediatype.go:390
		_go_fuzz_dep_.CoverTab[35798]++
//line /usr/local/go/src/mime/mediatype.go:390
		return c <= '9'
//line /usr/local/go/src/mime/mediatype.go:390
		// _ = "end of CoverTab[35798]"
//line /usr/local/go/src/mime/mediatype.go:390
	}():
//line /usr/local/go/src/mime/mediatype.go:390
		_go_fuzz_dep_.CoverTab[35794]++
							return true
//line /usr/local/go/src/mime/mediatype.go:391
		// _ = "end of CoverTab[35794]"
	case 'a' <= c && func() bool {
//line /usr/local/go/src/mime/mediatype.go:392
		_go_fuzz_dep_.CoverTab[35799]++
//line /usr/local/go/src/mime/mediatype.go:392
		return c <= 'f'
//line /usr/local/go/src/mime/mediatype.go:392
		// _ = "end of CoverTab[35799]"
//line /usr/local/go/src/mime/mediatype.go:392
	}():
//line /usr/local/go/src/mime/mediatype.go:392
		_go_fuzz_dep_.CoverTab[35795]++
							return true
//line /usr/local/go/src/mime/mediatype.go:393
		// _ = "end of CoverTab[35795]"
	case 'A' <= c && func() bool {
//line /usr/local/go/src/mime/mediatype.go:394
		_go_fuzz_dep_.CoverTab[35800]++
//line /usr/local/go/src/mime/mediatype.go:394
		return c <= 'F'
//line /usr/local/go/src/mime/mediatype.go:394
		// _ = "end of CoverTab[35800]"
//line /usr/local/go/src/mime/mediatype.go:394
	}():
//line /usr/local/go/src/mime/mediatype.go:394
		_go_fuzz_dep_.CoverTab[35796]++
							return true
//line /usr/local/go/src/mime/mediatype.go:395
		// _ = "end of CoverTab[35796]"
//line /usr/local/go/src/mime/mediatype.go:395
	default:
//line /usr/local/go/src/mime/mediatype.go:395
		_go_fuzz_dep_.CoverTab[35797]++
//line /usr/local/go/src/mime/mediatype.go:395
		// _ = "end of CoverTab[35797]"
	}
//line /usr/local/go/src/mime/mediatype.go:396
	// _ = "end of CoverTab[35792]"
//line /usr/local/go/src/mime/mediatype.go:396
	_go_fuzz_dep_.CoverTab[35793]++
						return false
//line /usr/local/go/src/mime/mediatype.go:397
	// _ = "end of CoverTab[35793]"
}

func unhex(c byte) byte {
//line /usr/local/go/src/mime/mediatype.go:400
	_go_fuzz_dep_.CoverTab[35801]++
						switch {
	case '0' <= c && func() bool {
//line /usr/local/go/src/mime/mediatype.go:402
		_go_fuzz_dep_.CoverTab[35807]++
//line /usr/local/go/src/mime/mediatype.go:402
		return c <= '9'
//line /usr/local/go/src/mime/mediatype.go:402
		// _ = "end of CoverTab[35807]"
//line /usr/local/go/src/mime/mediatype.go:402
	}():
//line /usr/local/go/src/mime/mediatype.go:402
		_go_fuzz_dep_.CoverTab[35803]++
							return c - '0'
//line /usr/local/go/src/mime/mediatype.go:403
		// _ = "end of CoverTab[35803]"
	case 'a' <= c && func() bool {
//line /usr/local/go/src/mime/mediatype.go:404
		_go_fuzz_dep_.CoverTab[35808]++
//line /usr/local/go/src/mime/mediatype.go:404
		return c <= 'f'
//line /usr/local/go/src/mime/mediatype.go:404
		// _ = "end of CoverTab[35808]"
//line /usr/local/go/src/mime/mediatype.go:404
	}():
//line /usr/local/go/src/mime/mediatype.go:404
		_go_fuzz_dep_.CoverTab[35804]++
							return c - 'a' + 10
//line /usr/local/go/src/mime/mediatype.go:405
		// _ = "end of CoverTab[35804]"
	case 'A' <= c && func() bool {
//line /usr/local/go/src/mime/mediatype.go:406
		_go_fuzz_dep_.CoverTab[35809]++
//line /usr/local/go/src/mime/mediatype.go:406
		return c <= 'F'
//line /usr/local/go/src/mime/mediatype.go:406
		// _ = "end of CoverTab[35809]"
//line /usr/local/go/src/mime/mediatype.go:406
	}():
//line /usr/local/go/src/mime/mediatype.go:406
		_go_fuzz_dep_.CoverTab[35805]++
							return c - 'A' + 10
//line /usr/local/go/src/mime/mediatype.go:407
		// _ = "end of CoverTab[35805]"
//line /usr/local/go/src/mime/mediatype.go:407
	default:
//line /usr/local/go/src/mime/mediatype.go:407
		_go_fuzz_dep_.CoverTab[35806]++
//line /usr/local/go/src/mime/mediatype.go:407
		// _ = "end of CoverTab[35806]"
	}
//line /usr/local/go/src/mime/mediatype.go:408
	// _ = "end of CoverTab[35801]"
//line /usr/local/go/src/mime/mediatype.go:408
	_go_fuzz_dep_.CoverTab[35802]++
						return 0
//line /usr/local/go/src/mime/mediatype.go:409
	// _ = "end of CoverTab[35802]"
}

//line /usr/local/go/src/mime/mediatype.go:410
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/mime/mediatype.go:410
var _ = _go_fuzz_dep_.CoverTab
