//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:1
package strconv

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:1
)

import (
	"errors"
	"unicode/utf8"
)

// ErrSyntax indicates that a value does not have the right syntax for the target type.
var ErrSyntax = errors.New("invalid syntax")

// Unquote interprets s as a single-quoted, double-quoted,
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:11
// or backquoted Go string literal, returning the string value
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:11
// that s quotes.  (If s is single-quoted, it would be a Go
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:11
// character literal; Unquote returns the corresponding
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:11
// one-character string.)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:16
func Unquote(s string) (t string, err error) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:16
	_go_fuzz_dep_.CoverTab[120801]++
												n := len(s)
												if n < 2 {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:18
		_go_fuzz_dep_.CoverTab[120808]++
													return "", ErrSyntax
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:19
		// _ = "end of CoverTab[120808]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:20
		_go_fuzz_dep_.CoverTab[120809]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:20
		// _ = "end of CoverTab[120809]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:20
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:20
	// _ = "end of CoverTab[120801]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:20
	_go_fuzz_dep_.CoverTab[120802]++
												quote := s[0]
												if quote != s[n-1] {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:22
		_go_fuzz_dep_.CoverTab[120810]++
													return "", ErrSyntax
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:23
		// _ = "end of CoverTab[120810]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:24
		_go_fuzz_dep_.CoverTab[120811]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:24
		// _ = "end of CoverTab[120811]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:24
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:24
	// _ = "end of CoverTab[120802]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:24
	_go_fuzz_dep_.CoverTab[120803]++
												s = s[1 : n-1]

												if quote != '"' {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:27
		_go_fuzz_dep_.CoverTab[120812]++
													return "", ErrSyntax
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:28
		// _ = "end of CoverTab[120812]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:29
		_go_fuzz_dep_.CoverTab[120813]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:29
		// _ = "end of CoverTab[120813]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:29
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:29
	// _ = "end of CoverTab[120803]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:29
	_go_fuzz_dep_.CoverTab[120804]++
												if !contains(s, '$') && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:30
		_go_fuzz_dep_.CoverTab[120814]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:30
		return !contains(s, '{')
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:30
		// _ = "end of CoverTab[120814]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:30
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:30
		_go_fuzz_dep_.CoverTab[120815]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:30
		return contains(s, '\n')
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:30
		// _ = "end of CoverTab[120815]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:30
	}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:30
		_go_fuzz_dep_.CoverTab[120816]++
													return "", ErrSyntax
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:31
		// _ = "end of CoverTab[120816]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:32
		_go_fuzz_dep_.CoverTab[120817]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:32
		// _ = "end of CoverTab[120817]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:32
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:32
	// _ = "end of CoverTab[120804]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:32
	_go_fuzz_dep_.CoverTab[120805]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:35
	if !contains(s, '\\') && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:35
		_go_fuzz_dep_.CoverTab[120818]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:35
		return !contains(s, quote)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:35
		// _ = "end of CoverTab[120818]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:35
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:35
		_go_fuzz_dep_.CoverTab[120819]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:35
		return !contains(s, '$')
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:35
		// _ = "end of CoverTab[120819]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:35
	}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:35
		_go_fuzz_dep_.CoverTab[120820]++
													switch quote {
		case '"':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:37
			_go_fuzz_dep_.CoverTab[120821]++
														return s, nil
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:38
			// _ = "end of CoverTab[120821]"
		case '\'':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:39
			_go_fuzz_dep_.CoverTab[120822]++
														r, size := utf8.DecodeRuneInString(s)
														if size == len(s) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:41
				_go_fuzz_dep_.CoverTab[120824]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:41
				return (r != utf8.RuneError || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:41
					_go_fuzz_dep_.CoverTab[120825]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:41
					return size != 1
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:41
					// _ = "end of CoverTab[120825]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:41
				}())
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:41
				// _ = "end of CoverTab[120824]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:41
			}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:41
				_go_fuzz_dep_.CoverTab[120826]++
															return s, nil
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:42
				// _ = "end of CoverTab[120826]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:43
				_go_fuzz_dep_.CoverTab[120827]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:43
				// _ = "end of CoverTab[120827]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:43
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:43
			// _ = "end of CoverTab[120822]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:43
		default:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:43
			_go_fuzz_dep_.CoverTab[120823]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:43
			// _ = "end of CoverTab[120823]"
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:44
		// _ = "end of CoverTab[120820]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:45
		_go_fuzz_dep_.CoverTab[120828]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:45
		// _ = "end of CoverTab[120828]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:45
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:45
	// _ = "end of CoverTab[120805]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:45
	_go_fuzz_dep_.CoverTab[120806]++

												var runeTmp [utf8.UTFMax]byte
												buf := make([]byte, 0, 3*len(s)/2)
												for len(s) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:49
		_go_fuzz_dep_.CoverTab[120829]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:53
		if s[0] == '$' && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:53
			_go_fuzz_dep_.CoverTab[120834]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:53
			return len(s) > 1
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:53
			// _ = "end of CoverTab[120834]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:53
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:53
			_go_fuzz_dep_.CoverTab[120835]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:53
			return s[1] == '{'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:53
			// _ = "end of CoverTab[120835]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:53
		}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:53
			_go_fuzz_dep_.CoverTab[120836]++
														buf = append(buf, '$', '{')
														s = s[2:]

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:58
			braces := 1
			for len(s) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:59
				_go_fuzz_dep_.CoverTab[120839]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:59
				return braces > 0
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:59
				// _ = "end of CoverTab[120839]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:59
			}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:59
				_go_fuzz_dep_.CoverTab[120840]++
															r, size := utf8.DecodeRuneInString(s)
															if r == utf8.RuneError {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:61
					_go_fuzz_dep_.CoverTab[120842]++
																return "", ErrSyntax
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:62
					// _ = "end of CoverTab[120842]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:63
					_go_fuzz_dep_.CoverTab[120843]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:63
					// _ = "end of CoverTab[120843]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:63
				}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:63
				// _ = "end of CoverTab[120840]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:63
				_go_fuzz_dep_.CoverTab[120841]++

															s = s[size:]

															n := utf8.EncodeRune(runeTmp[:], r)
															buf = append(buf, runeTmp[:n]...)

															switch r {
				case '{':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:71
					_go_fuzz_dep_.CoverTab[120844]++
																braces++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:72
					// _ = "end of CoverTab[120844]"
				case '}':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:73
					_go_fuzz_dep_.CoverTab[120845]++
																braces--
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:74
					// _ = "end of CoverTab[120845]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:74
				default:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:74
					_go_fuzz_dep_.CoverTab[120846]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:74
					// _ = "end of CoverTab[120846]"
				}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:75
				// _ = "end of CoverTab[120841]"
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:76
			// _ = "end of CoverTab[120836]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:76
			_go_fuzz_dep_.CoverTab[120837]++
														if braces != 0 {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:77
				_go_fuzz_dep_.CoverTab[120847]++
															return "", ErrSyntax
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:78
				// _ = "end of CoverTab[120847]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:79
				_go_fuzz_dep_.CoverTab[120848]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:79
				// _ = "end of CoverTab[120848]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:79
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:79
			// _ = "end of CoverTab[120837]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:79
			_go_fuzz_dep_.CoverTab[120838]++
														if len(s) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:80
				_go_fuzz_dep_.CoverTab[120849]++

															break
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:82
				// _ = "end of CoverTab[120849]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:83
				_go_fuzz_dep_.CoverTab[120850]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:86
				continue
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:86
				// _ = "end of CoverTab[120850]"
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:87
			// _ = "end of CoverTab[120838]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:88
			_go_fuzz_dep_.CoverTab[120851]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:88
			// _ = "end of CoverTab[120851]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:88
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:88
		// _ = "end of CoverTab[120829]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:88
		_go_fuzz_dep_.CoverTab[120830]++

													if s[0] == '\n' {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:90
			_go_fuzz_dep_.CoverTab[120852]++
														return "", ErrSyntax
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:91
			// _ = "end of CoverTab[120852]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:92
			_go_fuzz_dep_.CoverTab[120853]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:92
			// _ = "end of CoverTab[120853]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:92
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:92
		// _ = "end of CoverTab[120830]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:92
		_go_fuzz_dep_.CoverTab[120831]++

													c, multibyte, ss, err := unquoteChar(s, quote)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:95
			_go_fuzz_dep_.CoverTab[120854]++
														return "", err
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:96
			// _ = "end of CoverTab[120854]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:97
			_go_fuzz_dep_.CoverTab[120855]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:97
			// _ = "end of CoverTab[120855]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:97
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:97
		// _ = "end of CoverTab[120831]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:97
		_go_fuzz_dep_.CoverTab[120832]++
													s = ss
													if c < utf8.RuneSelf || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:99
			_go_fuzz_dep_.CoverTab[120856]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:99
			return !multibyte
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:99
			// _ = "end of CoverTab[120856]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:99
		}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:99
			_go_fuzz_dep_.CoverTab[120857]++
														buf = append(buf, byte(c))
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:100
			// _ = "end of CoverTab[120857]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:101
			_go_fuzz_dep_.CoverTab[120858]++
														n := utf8.EncodeRune(runeTmp[:], c)
														buf = append(buf, runeTmp[:n]...)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:103
			// _ = "end of CoverTab[120858]"
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:104
		// _ = "end of CoverTab[120832]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:104
		_go_fuzz_dep_.CoverTab[120833]++
													if quote == '\'' && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:105
			_go_fuzz_dep_.CoverTab[120859]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:105
			return len(s) != 0
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:105
			// _ = "end of CoverTab[120859]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:105
		}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:105
			_go_fuzz_dep_.CoverTab[120860]++

														return "", ErrSyntax
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:107
			// _ = "end of CoverTab[120860]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:108
			_go_fuzz_dep_.CoverTab[120861]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:108
			// _ = "end of CoverTab[120861]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:108
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:108
		// _ = "end of CoverTab[120833]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:109
	// _ = "end of CoverTab[120806]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:109
	_go_fuzz_dep_.CoverTab[120807]++
												return string(buf), nil
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:110
	// _ = "end of CoverTab[120807]"
}

// contains reports whether the string contains the byte c.
func contains(s string, c byte) bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:114
	_go_fuzz_dep_.CoverTab[120862]++
												for i := 0; i < len(s); i++ {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:115
		_go_fuzz_dep_.CoverTab[120864]++
													if s[i] == c {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:116
			_go_fuzz_dep_.CoverTab[120865]++
														return true
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:117
			// _ = "end of CoverTab[120865]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:118
			_go_fuzz_dep_.CoverTab[120866]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:118
			// _ = "end of CoverTab[120866]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:118
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:118
		// _ = "end of CoverTab[120864]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:119
	// _ = "end of CoverTab[120862]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:119
	_go_fuzz_dep_.CoverTab[120863]++
												return false
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:120
	// _ = "end of CoverTab[120863]"
}

func unhex(b byte) (v rune, ok bool) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:123
	_go_fuzz_dep_.CoverTab[120867]++
												c := rune(b)
												switch {
	case '0' <= c && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:126
		_go_fuzz_dep_.CoverTab[120873]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:126
		return c <= '9'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:126
		// _ = "end of CoverTab[120873]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:126
	}():
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:126
		_go_fuzz_dep_.CoverTab[120869]++
													return c - '0', true
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:127
		// _ = "end of CoverTab[120869]"
	case 'a' <= c && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:128
		_go_fuzz_dep_.CoverTab[120874]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:128
		return c <= 'f'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:128
		// _ = "end of CoverTab[120874]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:128
	}():
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:128
		_go_fuzz_dep_.CoverTab[120870]++
													return c - 'a' + 10, true
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:129
		// _ = "end of CoverTab[120870]"
	case 'A' <= c && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:130
		_go_fuzz_dep_.CoverTab[120875]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:130
		return c <= 'F'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:130
		// _ = "end of CoverTab[120875]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:130
	}():
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:130
		_go_fuzz_dep_.CoverTab[120871]++
													return c - 'A' + 10, true
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:131
		// _ = "end of CoverTab[120871]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:131
	default:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:131
		_go_fuzz_dep_.CoverTab[120872]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:131
		// _ = "end of CoverTab[120872]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:132
	// _ = "end of CoverTab[120867]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:132
	_go_fuzz_dep_.CoverTab[120868]++
												return
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:133
	// _ = "end of CoverTab[120868]"
}

func unquoteChar(s string, quote byte) (value rune, multibyte bool, tail string, err error) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:136
	_go_fuzz_dep_.CoverTab[120876]++

												switch c := s[0]; {
	case c == quote && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:139
		_go_fuzz_dep_.CoverTab[120884]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:139
		return (quote == '\'' || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:139
			_go_fuzz_dep_.CoverTab[120885]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:139
			return quote == '"'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:139
			// _ = "end of CoverTab[120885]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:139
		}())
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:139
		// _ = "end of CoverTab[120884]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:139
	}():
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:139
		_go_fuzz_dep_.CoverTab[120880]++
													err = ErrSyntax
													return
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:141
		// _ = "end of CoverTab[120880]"
	case c >= utf8.RuneSelf:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:142
		_go_fuzz_dep_.CoverTab[120881]++
													r, size := utf8.DecodeRuneInString(s)
													return r, true, s[size:], nil
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:144
		// _ = "end of CoverTab[120881]"
	case c != '\\':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:145
		_go_fuzz_dep_.CoverTab[120882]++
													return rune(s[0]), false, s[1:], nil
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:146
		// _ = "end of CoverTab[120882]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:146
	default:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:146
		_go_fuzz_dep_.CoverTab[120883]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:146
		// _ = "end of CoverTab[120883]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:147
	// _ = "end of CoverTab[120876]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:147
	_go_fuzz_dep_.CoverTab[120877]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:150
	if len(s) <= 1 {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:150
		_go_fuzz_dep_.CoverTab[120886]++
													err = ErrSyntax
													return
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:152
		// _ = "end of CoverTab[120886]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:153
		_go_fuzz_dep_.CoverTab[120887]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:153
		// _ = "end of CoverTab[120887]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:153
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:153
	// _ = "end of CoverTab[120877]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:153
	_go_fuzz_dep_.CoverTab[120878]++
												c := s[1]
												s = s[2:]

												switch c {
	case 'a':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:158
		_go_fuzz_dep_.CoverTab[120888]++
													value = '\a'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:159
		// _ = "end of CoverTab[120888]"
	case 'b':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:160
		_go_fuzz_dep_.CoverTab[120889]++
													value = '\b'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:161
		// _ = "end of CoverTab[120889]"
	case 'f':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:162
		_go_fuzz_dep_.CoverTab[120890]++
													value = '\f'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:163
		// _ = "end of CoverTab[120890]"
	case 'n':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:164
		_go_fuzz_dep_.CoverTab[120891]++
													value = '\n'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:165
		// _ = "end of CoverTab[120891]"
	case 'r':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:166
		_go_fuzz_dep_.CoverTab[120892]++
													value = '\r'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:167
		// _ = "end of CoverTab[120892]"
	case 't':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:168
		_go_fuzz_dep_.CoverTab[120893]++
													value = '\t'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:169
		// _ = "end of CoverTab[120893]"
	case 'v':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:170
		_go_fuzz_dep_.CoverTab[120894]++
													value = '\v'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:171
		// _ = "end of CoverTab[120894]"
	case 'x', 'u', 'U':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:172
		_go_fuzz_dep_.CoverTab[120895]++
													n := 0
													switch c {
		case 'x':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:175
			_go_fuzz_dep_.CoverTab[120909]++
														n = 2
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:176
			// _ = "end of CoverTab[120909]"
		case 'u':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:177
			_go_fuzz_dep_.CoverTab[120910]++
														n = 4
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:178
			// _ = "end of CoverTab[120910]"
		case 'U':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:179
			_go_fuzz_dep_.CoverTab[120911]++
														n = 8
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:180
			// _ = "end of CoverTab[120911]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:180
		default:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:180
			_go_fuzz_dep_.CoverTab[120912]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:180
			// _ = "end of CoverTab[120912]"
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:181
		// _ = "end of CoverTab[120895]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:181
		_go_fuzz_dep_.CoverTab[120896]++
													var v rune
													if len(s) < n {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:183
			_go_fuzz_dep_.CoverTab[120913]++
														err = ErrSyntax
														return
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:185
			// _ = "end of CoverTab[120913]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:186
			_go_fuzz_dep_.CoverTab[120914]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:186
			// _ = "end of CoverTab[120914]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:186
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:186
		// _ = "end of CoverTab[120896]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:186
		_go_fuzz_dep_.CoverTab[120897]++
													for j := 0; j < n; j++ {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:187
			_go_fuzz_dep_.CoverTab[120915]++
														x, ok := unhex(s[j])
														if !ok {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:189
				_go_fuzz_dep_.CoverTab[120917]++
															err = ErrSyntax
															return
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:191
				// _ = "end of CoverTab[120917]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:192
				_go_fuzz_dep_.CoverTab[120918]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:192
				// _ = "end of CoverTab[120918]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:192
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:192
			// _ = "end of CoverTab[120915]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:192
			_go_fuzz_dep_.CoverTab[120916]++
														v = v<<4 | x
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:193
			// _ = "end of CoverTab[120916]"
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:194
		// _ = "end of CoverTab[120897]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:194
		_go_fuzz_dep_.CoverTab[120898]++
													s = s[n:]
													if c == 'x' {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:196
			_go_fuzz_dep_.CoverTab[120919]++

														value = v
														break
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:199
			// _ = "end of CoverTab[120919]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:200
			_go_fuzz_dep_.CoverTab[120920]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:200
			// _ = "end of CoverTab[120920]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:200
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:200
		// _ = "end of CoverTab[120898]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:200
		_go_fuzz_dep_.CoverTab[120899]++
													if v > utf8.MaxRune {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:201
			_go_fuzz_dep_.CoverTab[120921]++
														err = ErrSyntax
														return
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:203
			// _ = "end of CoverTab[120921]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:204
			_go_fuzz_dep_.CoverTab[120922]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:204
			// _ = "end of CoverTab[120922]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:204
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:204
		// _ = "end of CoverTab[120899]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:204
		_go_fuzz_dep_.CoverTab[120900]++
													value = v
													multibyte = true
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:206
		// _ = "end of CoverTab[120900]"
	case '0', '1', '2', '3', '4', '5', '6', '7':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:207
		_go_fuzz_dep_.CoverTab[120901]++
													v := rune(c) - '0'
													if len(s) < 2 {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:209
			_go_fuzz_dep_.CoverTab[120923]++
														err = ErrSyntax
														return
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:211
			// _ = "end of CoverTab[120923]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:212
			_go_fuzz_dep_.CoverTab[120924]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:212
			// _ = "end of CoverTab[120924]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:212
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:212
		// _ = "end of CoverTab[120901]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:212
		_go_fuzz_dep_.CoverTab[120902]++
													for j := 0; j < 2; j++ {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:213
			_go_fuzz_dep_.CoverTab[120925]++
														x := rune(s[j]) - '0'
														if x < 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:215
				_go_fuzz_dep_.CoverTab[120927]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:215
				return x > 7
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:215
				// _ = "end of CoverTab[120927]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:215
			}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:215
				_go_fuzz_dep_.CoverTab[120928]++
															err = ErrSyntax
															return
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:217
				// _ = "end of CoverTab[120928]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:218
				_go_fuzz_dep_.CoverTab[120929]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:218
				// _ = "end of CoverTab[120929]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:218
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:218
			// _ = "end of CoverTab[120925]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:218
			_go_fuzz_dep_.CoverTab[120926]++
														v = (v << 3) | x
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:219
			// _ = "end of CoverTab[120926]"
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:220
		// _ = "end of CoverTab[120902]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:220
		_go_fuzz_dep_.CoverTab[120903]++
													s = s[2:]
													if v > 255 {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:222
			_go_fuzz_dep_.CoverTab[120930]++
														err = ErrSyntax
														return
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:224
			// _ = "end of CoverTab[120930]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:225
			_go_fuzz_dep_.CoverTab[120931]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:225
			// _ = "end of CoverTab[120931]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:225
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:225
		// _ = "end of CoverTab[120903]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:225
		_go_fuzz_dep_.CoverTab[120904]++
													value = v
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:226
		// _ = "end of CoverTab[120904]"
	case '\\':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:227
		_go_fuzz_dep_.CoverTab[120905]++
													value = '\\'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:228
		// _ = "end of CoverTab[120905]"
	case '\'', '"':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:229
		_go_fuzz_dep_.CoverTab[120906]++
													if c != quote {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:230
			_go_fuzz_dep_.CoverTab[120932]++
														err = ErrSyntax
														return
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:232
			// _ = "end of CoverTab[120932]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:233
			_go_fuzz_dep_.CoverTab[120933]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:233
			// _ = "end of CoverTab[120933]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:233
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:233
		// _ = "end of CoverTab[120906]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:233
		_go_fuzz_dep_.CoverTab[120907]++
													value = rune(c)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:234
		// _ = "end of CoverTab[120907]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:235
		_go_fuzz_dep_.CoverTab[120908]++
													err = ErrSyntax
													return
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:237
		// _ = "end of CoverTab[120908]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:238
	// _ = "end of CoverTab[120878]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:238
	_go_fuzz_dep_.CoverTab[120879]++
												tail = s
												return
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:240
	// _ = "end of CoverTab[120879]"
}

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:241
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/strconv/quote.go:241
var _ = _go_fuzz_dep_.CoverTab
