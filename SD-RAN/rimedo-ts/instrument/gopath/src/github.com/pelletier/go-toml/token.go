//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:1
package toml

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:1
)

import "fmt"

// Define tokens
type tokenType int

const (
	eof = -(iota + 1)
)

const (
	tokenError	tokenType	= iota
	tokenEOF
	tokenComment
	tokenKey
	tokenString
	tokenInteger
	tokenTrue
	tokenFalse
	tokenFloat
	tokenInf
	tokenNan
	tokenEqual
	tokenLeftBracket
	tokenRightBracket
	tokenLeftCurlyBrace
	tokenRightCurlyBrace
	tokenLeftParen
	tokenRightParen
	tokenDoubleLeftBracket
	tokenDoubleRightBracket
	tokenLocalDate
	tokenLocalTime
	tokenTimeOffset
	tokenKeyGroup
	tokenKeyGroupArray
	tokenComma
	tokenColon
	tokenDollar
	tokenStar
	tokenQuestion
	tokenDot
	tokenDotDot
	tokenEOL
)

var tokenTypeNames = []string{
	"Error",
	"EOF",
	"Comment",
	"Key",
	"String",
	"Integer",
	"True",
	"False",
	"Float",
	"Inf",
	"NaN",
	"=",
	"[",
	"]",
	"{",
	"}",
	"(",
	")",
	"]]",
	"[[",
	"LocalDate",
	"LocalTime",
	"TimeOffset",
	"KeyGroup",
	"KeyGroupArray",
	",",
	":",
	"$",
	"*",
	"?",
	".",
	"..",
	"EOL",
}

type token struct {
	Position
	typ	tokenType
	val	string
}

func (tt tokenType) String() string {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:90
	_go_fuzz_dep_.CoverTab[123862]++
											idx := int(tt)
											if idx < len(tokenTypeNames) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:92
		_go_fuzz_dep_.CoverTab[123864]++
												return tokenTypeNames[idx]
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:93
		// _ = "end of CoverTab[123864]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:94
		_go_fuzz_dep_.CoverTab[123865]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:94
		// _ = "end of CoverTab[123865]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:94
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:94
	// _ = "end of CoverTab[123862]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:94
	_go_fuzz_dep_.CoverTab[123863]++
											return "Unknown"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:95
	// _ = "end of CoverTab[123863]"
}

func (t token) String() string {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:98
	_go_fuzz_dep_.CoverTab[123866]++
											switch t.typ {
	case tokenEOF:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:100
		_go_fuzz_dep_.CoverTab[123868]++
												return "EOF"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:101
		// _ = "end of CoverTab[123868]"
	case tokenError:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:102
		_go_fuzz_dep_.CoverTab[123869]++
												return t.val
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:103
		// _ = "end of CoverTab[123869]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:103
	default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:103
		_go_fuzz_dep_.CoverTab[123870]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:103
		// _ = "end of CoverTab[123870]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:104
	// _ = "end of CoverTab[123866]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:104
	_go_fuzz_dep_.CoverTab[123867]++

											return fmt.Sprintf("%q", t.val)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:106
	// _ = "end of CoverTab[123867]"
}

func isSpace(r rune) bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:109
	_go_fuzz_dep_.CoverTab[123871]++
											return r == ' ' || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:110
		_go_fuzz_dep_.CoverTab[123872]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:110
		return r == '\t'
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:110
		// _ = "end of CoverTab[123872]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:110
	}()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:110
	// _ = "end of CoverTab[123871]"
}

func isAlphanumeric(r rune) bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:113
	_go_fuzz_dep_.CoverTab[123873]++
											return 'a' <= r && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:114
		_go_fuzz_dep_.CoverTab[123874]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:114
		return r <= 'z'
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:114
		// _ = "end of CoverTab[123874]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:114
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:114
		_go_fuzz_dep_.CoverTab[123875]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:114
		return 'A' <= r && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:114
			_go_fuzz_dep_.CoverTab[123876]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:114
			return r <= 'Z'
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:114
			// _ = "end of CoverTab[123876]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:114
		}()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:114
		// _ = "end of CoverTab[123875]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:114
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:114
		_go_fuzz_dep_.CoverTab[123877]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:114
		return r == '_'
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:114
		// _ = "end of CoverTab[123877]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:114
	}()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:114
	// _ = "end of CoverTab[123873]"
}

func isKeyChar(r rune) bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:117
	_go_fuzz_dep_.CoverTab[123878]++

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:121
	return !(r == '\r' || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:121
		_go_fuzz_dep_.CoverTab[123879]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:121
		return r == '\n'
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:121
		// _ = "end of CoverTab[123879]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:121
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:121
		_go_fuzz_dep_.CoverTab[123880]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:121
		return r == eof
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:121
		// _ = "end of CoverTab[123880]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:121
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:121
		_go_fuzz_dep_.CoverTab[123881]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:121
		return r == '='
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:121
		// _ = "end of CoverTab[123881]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:121
	}())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:121
	// _ = "end of CoverTab[123878]"
}

func isKeyStartChar(r rune) bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:124
	_go_fuzz_dep_.CoverTab[123882]++
											return !(isSpace(r) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:125
		_go_fuzz_dep_.CoverTab[123883]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:125
		return r == '\r'
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:125
		// _ = "end of CoverTab[123883]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:125
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:125
		_go_fuzz_dep_.CoverTab[123884]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:125
		return r == '\n'
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:125
		// _ = "end of CoverTab[123884]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:125
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:125
		_go_fuzz_dep_.CoverTab[123885]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:125
		return r == eof
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:125
		// _ = "end of CoverTab[123885]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:125
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:125
		_go_fuzz_dep_.CoverTab[123886]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:125
		return r == '['
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:125
		// _ = "end of CoverTab[123886]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:125
	}())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:125
	// _ = "end of CoverTab[123882]"
}

func isDigit(r rune) bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:128
	_go_fuzz_dep_.CoverTab[123887]++
											return '0' <= r && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:129
		_go_fuzz_dep_.CoverTab[123888]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:129
		return r <= '9'
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:129
		// _ = "end of CoverTab[123888]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:129
	}()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:129
	// _ = "end of CoverTab[123887]"
}

func isHexDigit(r rune) bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:132
	_go_fuzz_dep_.CoverTab[123889]++
											return isDigit(r) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:133
		_go_fuzz_dep_.CoverTab[123890]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:133
		return (r >= 'a' && func() bool {
													_go_fuzz_dep_.CoverTab[123891]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:134
			return r <= 'f'
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:134
			// _ = "end of CoverTab[123891]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:134
		}())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:134
		// _ = "end of CoverTab[123890]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:134
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:134
		_go_fuzz_dep_.CoverTab[123892]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:134
		return (r >= 'A' && func() bool {
													_go_fuzz_dep_.CoverTab[123893]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:135
			return r <= 'F'
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:135
			// _ = "end of CoverTab[123893]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:135
		}())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:135
		// _ = "end of CoverTab[123892]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:135
	}()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:135
	// _ = "end of CoverTab[123889]"
}

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:136
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/token.go:136
var _ = _go_fuzz_dep_.CoverTab
