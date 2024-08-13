//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:1
// Package token defines constants representing the lexical tokens for HCL
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:1
// (HashiCorp Configuration Language)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:3
package token

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:3
import (
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:3
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:3
)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:3
import (
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:3
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:3
)

import (
	"fmt"
	"strconv"
	"strings"

	hclstrconv "github.com/hashicorp/hcl/hcl/strconv"
)

// Token defines a single HCL token which can be obtained via the Scanner
type Token struct {
	Type	Type
	Pos	Pos
	Text	string
	JSON	bool
}

// Type is the set of lexical tokens of the HCL (HashiCorp Configuration Language)
type Type int

const (
	// Special tokens
	ILLEGAL	Type	= iota
	EOF
	COMMENT

	identifier_beg
	IDENT	// literals
	literal_beg
	NUMBER	// 12345
	FLOAT	// 123.45
	BOOL	// true,false
	STRING	// "abc"
	HEREDOC	// <<FOO\nbar\nFOO
	literal_end
	identifier_end

	operator_beg
	LBRACK	// [
	LBRACE	// {
	COMMA	// ,
	PERIOD	// .

	RBRACK	// ]
	RBRACE	// }

	ASSIGN	// =
	ADD	// +
	SUB	// -
	operator_end
)

var tokens = [...]string{
	ILLEGAL:	"ILLEGAL",

	EOF:		"EOF",
	COMMENT:	"COMMENT",

	IDENT:	"IDENT",
	NUMBER:	"NUMBER",
	FLOAT:	"FLOAT",
	BOOL:	"BOOL",
	STRING:	"STRING",

	LBRACK:		"LBRACK",
	LBRACE:		"LBRACE",
	COMMA:		"COMMA",
	PERIOD:		"PERIOD",
	HEREDOC:	"HEREDOC",

	RBRACK:	"RBRACK",
	RBRACE:	"RBRACE",

	ASSIGN:	"ASSIGN",
	ADD:	"ADD",
	SUB:	"SUB",
}

// String returns the string corresponding to the token tok.
func (t Type) String() string {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:83
	_go_fuzz_dep_.CoverTab[120949]++
												s := ""
												if 0 <= t && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:85
		_go_fuzz_dep_.CoverTab[120952]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:85
		return t < Type(len(tokens))
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:85
		// _ = "end of CoverTab[120952]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:85
	}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:85
		_go_fuzz_dep_.CoverTab[120953]++
													s = tokens[t]
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:86
		// _ = "end of CoverTab[120953]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:87
		_go_fuzz_dep_.CoverTab[120954]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:87
		// _ = "end of CoverTab[120954]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:87
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:87
	// _ = "end of CoverTab[120949]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:87
	_go_fuzz_dep_.CoverTab[120950]++
												if s == "" {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:88
		_go_fuzz_dep_.CoverTab[120955]++
													s = "token(" + strconv.Itoa(int(t)) + ")"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:89
		// _ = "end of CoverTab[120955]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:90
		_go_fuzz_dep_.CoverTab[120956]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:90
		// _ = "end of CoverTab[120956]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:90
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:90
	// _ = "end of CoverTab[120950]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:90
	_go_fuzz_dep_.CoverTab[120951]++
												return s
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:91
	// _ = "end of CoverTab[120951]"
}

// IsIdentifier returns true for tokens corresponding to identifiers and basic
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:94
// type literals; it returns false otherwise.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:96
func (t Type) IsIdentifier() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:96
	_go_fuzz_dep_.CoverTab[120957]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:96
	return identifier_beg < t && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:96
		_go_fuzz_dep_.CoverTab[120958]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:96
		return t < identifier_end
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:96
		// _ = "end of CoverTab[120958]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:96
	}()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:96
	// _ = "end of CoverTab[120957]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:96
}

// IsLiteral returns true for tokens corresponding to basic type literals; it
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:98
// returns false otherwise.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:100
func (t Type) IsLiteral() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:100
	_go_fuzz_dep_.CoverTab[120959]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:100
	return literal_beg < t && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:100
		_go_fuzz_dep_.CoverTab[120960]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:100
		return t < literal_end
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:100
		// _ = "end of CoverTab[120960]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:100
	}()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:100
	// _ = "end of CoverTab[120959]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:100
}

// IsOperator returns true for tokens corresponding to operators and
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:102
// delimiters; it returns false otherwise.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:104
func (t Type) IsOperator() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:104
	_go_fuzz_dep_.CoverTab[120961]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:104
	return operator_beg < t && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:104
		_go_fuzz_dep_.CoverTab[120962]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:104
		return t < operator_end
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:104
		// _ = "end of CoverTab[120962]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:104
	}()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:104
	// _ = "end of CoverTab[120961]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:104
}

// String returns the token's literal text. Note that this is only
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:106
// applicable for certain token types, such as token.IDENT,
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:106
// token.STRING, etc..
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:109
func (t Token) String() string {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:109
	_go_fuzz_dep_.CoverTab[120963]++
												return fmt.Sprintf("%s %s %s", t.Pos.String(), t.Type.String(), t.Text)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:110
	// _ = "end of CoverTab[120963]"
}

// Value returns the properly typed value for this token. The type of
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:113
// the returned interface{} is guaranteed based on the Type field.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:113
//
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:113
// This can only be called for literal types. If it is called for any other
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:113
// type, this will panic.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:118
func (t Token) Value() interface{} {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:118
	_go_fuzz_dep_.CoverTab[120964]++
												switch t.Type {
	case BOOL:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:120
		_go_fuzz_dep_.CoverTab[120965]++
													if t.Text == "true" {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:121
			_go_fuzz_dep_.CoverTab[120978]++
														return true
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:122
			// _ = "end of CoverTab[120978]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:123
			_go_fuzz_dep_.CoverTab[120979]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:123
			if t.Text == "false" {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:123
				_go_fuzz_dep_.CoverTab[120980]++
															return false
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:124
				// _ = "end of CoverTab[120980]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:125
				_go_fuzz_dep_.CoverTab[120981]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:125
				// _ = "end of CoverTab[120981]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:125
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:125
			// _ = "end of CoverTab[120979]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:125
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:125
		// _ = "end of CoverTab[120965]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:125
		_go_fuzz_dep_.CoverTab[120966]++

													panic("unknown bool value: " + t.Text)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:127
		// _ = "end of CoverTab[120966]"
	case FLOAT:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:128
		_go_fuzz_dep_.CoverTab[120967]++
													v, err := strconv.ParseFloat(t.Text, 64)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:130
			_go_fuzz_dep_.CoverTab[120982]++
														panic(err)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:131
			// _ = "end of CoverTab[120982]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:132
			_go_fuzz_dep_.CoverTab[120983]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:132
			// _ = "end of CoverTab[120983]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:132
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:132
		// _ = "end of CoverTab[120967]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:132
		_go_fuzz_dep_.CoverTab[120968]++

													return float64(v)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:134
		// _ = "end of CoverTab[120968]"
	case NUMBER:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:135
		_go_fuzz_dep_.CoverTab[120969]++
													v, err := strconv.ParseInt(t.Text, 0, 64)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:137
			_go_fuzz_dep_.CoverTab[120984]++
														panic(err)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:138
			// _ = "end of CoverTab[120984]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:139
			_go_fuzz_dep_.CoverTab[120985]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:139
			// _ = "end of CoverTab[120985]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:139
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:139
		// _ = "end of CoverTab[120969]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:139
		_go_fuzz_dep_.CoverTab[120970]++

													return int64(v)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:141
		// _ = "end of CoverTab[120970]"
	case IDENT:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:142
		_go_fuzz_dep_.CoverTab[120971]++
													return t.Text
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:143
		// _ = "end of CoverTab[120971]"
	case HEREDOC:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:144
		_go_fuzz_dep_.CoverTab[120972]++
													return unindentHeredoc(t.Text)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:145
		// _ = "end of CoverTab[120972]"
	case STRING:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:146
		_go_fuzz_dep_.CoverTab[120973]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:150
		f := hclstrconv.Unquote
		if t.JSON {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:151
			_go_fuzz_dep_.CoverTab[120986]++
														f = strconv.Unquote
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:152
			// _ = "end of CoverTab[120986]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:153
			_go_fuzz_dep_.CoverTab[120987]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:153
			// _ = "end of CoverTab[120987]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:153
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:153
		// _ = "end of CoverTab[120973]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:153
		_go_fuzz_dep_.CoverTab[120974]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:156
		if t.Text == "" {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:156
			_go_fuzz_dep_.CoverTab[120988]++
														return ""
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:157
			// _ = "end of CoverTab[120988]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:158
			_go_fuzz_dep_.CoverTab[120989]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:158
			// _ = "end of CoverTab[120989]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:158
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:158
		// _ = "end of CoverTab[120974]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:158
		_go_fuzz_dep_.CoverTab[120975]++

													v, err := f(t.Text)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:161
			_go_fuzz_dep_.CoverTab[120990]++
														panic(fmt.Sprintf("unquote %s err: %s", t.Text, err))
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:162
			// _ = "end of CoverTab[120990]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:163
			_go_fuzz_dep_.CoverTab[120991]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:163
			// _ = "end of CoverTab[120991]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:163
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:163
		// _ = "end of CoverTab[120975]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:163
		_go_fuzz_dep_.CoverTab[120976]++

													return v
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:165
		// _ = "end of CoverTab[120976]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:166
		_go_fuzz_dep_.CoverTab[120977]++
													panic(fmt.Sprintf("unimplemented Value for type: %s", t.Type))
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:167
		// _ = "end of CoverTab[120977]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:168
	// _ = "end of CoverTab[120964]"
}

// unindentHeredoc returns the string content of a HEREDOC if it is started with <<
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:171
// and the content of a HEREDOC with the hanging indent removed if it is started with
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:171
// a <<-, and the terminating line is at least as indented as the least indented line.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:174
func unindentHeredoc(heredoc string) string {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:174
	_go_fuzz_dep_.CoverTab[120992]++

												idx := strings.IndexByte(heredoc, '\n')
												if idx == -1 {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:177
		_go_fuzz_dep_.CoverTab[120998]++
													panic("heredoc doesn't contain newline")
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:178
		// _ = "end of CoverTab[120998]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:179
		_go_fuzz_dep_.CoverTab[120999]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:179
		// _ = "end of CoverTab[120999]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:179
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:179
	// _ = "end of CoverTab[120992]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:179
	_go_fuzz_dep_.CoverTab[120993]++

												unindent := heredoc[2] == '-'

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:184
	if !unindent {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:184
		_go_fuzz_dep_.CoverTab[121000]++
													return string(heredoc[idx+1 : len(heredoc)-idx+1])
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:185
		// _ = "end of CoverTab[121000]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:186
		_go_fuzz_dep_.CoverTab[121001]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:186
		// _ = "end of CoverTab[121001]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:186
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:186
	// _ = "end of CoverTab[120993]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:186
	_go_fuzz_dep_.CoverTab[120994]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:189
	lines := strings.Split(string(heredoc[idx+1:len(heredoc)-idx+2]), "\n")
	whitespacePrefix := lines[len(lines)-1]

	isIndented := true
	for _, v := range lines {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:193
		_go_fuzz_dep_.CoverTab[121002]++
													if strings.HasPrefix(v, whitespacePrefix) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:194
			_go_fuzz_dep_.CoverTab[121004]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:195
			// _ = "end of CoverTab[121004]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:196
			_go_fuzz_dep_.CoverTab[121005]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:196
			// _ = "end of CoverTab[121005]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:196
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:196
		// _ = "end of CoverTab[121002]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:196
		_go_fuzz_dep_.CoverTab[121003]++

													isIndented = false
													break
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:199
		// _ = "end of CoverTab[121003]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:200
	// _ = "end of CoverTab[120994]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:200
	_go_fuzz_dep_.CoverTab[120995]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:204
	if !isIndented {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:204
		_go_fuzz_dep_.CoverTab[121006]++
													return strings.TrimRight(string(heredoc[idx+1:len(heredoc)-idx+1]), " \t")
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:205
		// _ = "end of CoverTab[121006]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:206
		_go_fuzz_dep_.CoverTab[121007]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:206
		// _ = "end of CoverTab[121007]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:206
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:206
	// _ = "end of CoverTab[120995]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:206
	_go_fuzz_dep_.CoverTab[120996]++

												unindentedLines := make([]string, len(lines))
												for k, v := range lines {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:209
		_go_fuzz_dep_.CoverTab[121008]++
													if k == len(lines)-1 {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:210
			_go_fuzz_dep_.CoverTab[121010]++
														unindentedLines[k] = ""
														break
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:212
			// _ = "end of CoverTab[121010]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:213
			_go_fuzz_dep_.CoverTab[121011]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:213
			// _ = "end of CoverTab[121011]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:213
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:213
		// _ = "end of CoverTab[121008]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:213
		_go_fuzz_dep_.CoverTab[121009]++

													unindentedLines[k] = strings.TrimPrefix(v, whitespacePrefix)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:215
		// _ = "end of CoverTab[121009]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:216
	// _ = "end of CoverTab[120996]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:216
	_go_fuzz_dep_.CoverTab[120997]++

												return strings.Join(unindentedLines, "\n")
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:218
	// _ = "end of CoverTab[120997]"
}

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:219
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/token/token.go:219
var _ = _go_fuzz_dep_.CoverTab
