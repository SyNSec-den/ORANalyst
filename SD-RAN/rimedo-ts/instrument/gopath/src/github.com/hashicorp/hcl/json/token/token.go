//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:1
package token

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:1
)

import (
	"fmt"
	"strconv"

	hcltoken "github.com/hashicorp/hcl/hcl/token"
)

// Token defines a single HCL token which can be obtained via the Scanner
type Token struct {
	Type	Type
	Pos	Pos
	Text	string
}

// Type is the set of lexical tokens of the HCL (HashiCorp Configuration Language)
type Type int

const (
	// Special tokens
	ILLEGAL	Type	= iota
	EOF

	identifier_beg
	literal_beg
	NUMBER	// 12345
	FLOAT	// 123.45
	BOOL	// true,false
	STRING	// "abc"
	NULL	// null
	literal_end
	identifier_end

	operator_beg
	LBRACK	// [
	LBRACE	// {
	COMMA	// ,
	PERIOD	// .
	COLON	// :

	RBRACK	// ]
	RBRACE	// }

	operator_end
)

var tokens = [...]string{
	ILLEGAL:	"ILLEGAL",

	EOF:	"EOF",

	NUMBER:	"NUMBER",
	FLOAT:	"FLOAT",
	BOOL:	"BOOL",
	STRING:	"STRING",
	NULL:	"NULL",

	LBRACK:	"LBRACK",
	LBRACE:	"LBRACE",
	COMMA:	"COMMA",
	PERIOD:	"PERIOD",
	COLON:	"COLON",

	RBRACK:	"RBRACK",
	RBRACE:	"RBRACE",
}

// String returns the string corresponding to the token tok.
func (t Type) String() string {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:70
	_go_fuzz_dep_.CoverTab[121562]++
												s := ""
												if 0 <= t && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:72
		_go_fuzz_dep_.CoverTab[121565]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:72
		return t < Type(len(tokens))
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:72
		// _ = "end of CoverTab[121565]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:72
	}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:72
		_go_fuzz_dep_.CoverTab[121566]++
													s = tokens[t]
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:73
		// _ = "end of CoverTab[121566]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:74
		_go_fuzz_dep_.CoverTab[121567]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:74
		// _ = "end of CoverTab[121567]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:74
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:74
	// _ = "end of CoverTab[121562]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:74
	_go_fuzz_dep_.CoverTab[121563]++
												if s == "" {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:75
		_go_fuzz_dep_.CoverTab[121568]++
													s = "token(" + strconv.Itoa(int(t)) + ")"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:76
		// _ = "end of CoverTab[121568]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:77
		_go_fuzz_dep_.CoverTab[121569]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:77
		// _ = "end of CoverTab[121569]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:77
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:77
	// _ = "end of CoverTab[121563]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:77
	_go_fuzz_dep_.CoverTab[121564]++
												return s
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:78
	// _ = "end of CoverTab[121564]"
}

// IsIdentifier returns true for tokens corresponding to identifiers and basic
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:81
// type literals; it returns false otherwise.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:83
func (t Type) IsIdentifier() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:83
	_go_fuzz_dep_.CoverTab[121570]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:83
	return identifier_beg < t && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:83
		_go_fuzz_dep_.CoverTab[121571]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:83
		return t < identifier_end
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:83
		// _ = "end of CoverTab[121571]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:83
	}()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:83
	// _ = "end of CoverTab[121570]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:83
}

// IsLiteral returns true for tokens corresponding to basic type literals; it
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:85
// returns false otherwise.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:87
func (t Type) IsLiteral() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:87
	_go_fuzz_dep_.CoverTab[121572]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:87
	return literal_beg < t && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:87
		_go_fuzz_dep_.CoverTab[121573]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:87
		return t < literal_end
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:87
		// _ = "end of CoverTab[121573]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:87
	}()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:87
	// _ = "end of CoverTab[121572]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:87
}

// IsOperator returns true for tokens corresponding to operators and
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:89
// delimiters; it returns false otherwise.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:91
func (t Type) IsOperator() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:91
	_go_fuzz_dep_.CoverTab[121574]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:91
	return operator_beg < t && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:91
		_go_fuzz_dep_.CoverTab[121575]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:91
		return t < operator_end
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:91
		// _ = "end of CoverTab[121575]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:91
	}()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:91
	// _ = "end of CoverTab[121574]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:91
}

// String returns the token's literal text. Note that this is only
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:93
// applicable for certain token types, such as token.IDENT,
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:93
// token.STRING, etc..
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:96
func (t Token) String() string {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:96
	_go_fuzz_dep_.CoverTab[121576]++
												return fmt.Sprintf("%s %s %s", t.Pos.String(), t.Type.String(), t.Text)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:97
	// _ = "end of CoverTab[121576]"
}

// HCLToken converts this token to an HCL token.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:100
//
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:100
// The token type must be a literal type or this will panic.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:103
func (t Token) HCLToken() hcltoken.Token {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:103
	_go_fuzz_dep_.CoverTab[121577]++
												switch t.Type {
	case BOOL:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:105
		_go_fuzz_dep_.CoverTab[121578]++
													return hcltoken.Token{Type: hcltoken.BOOL, Text: t.Text}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:106
		// _ = "end of CoverTab[121578]"
	case FLOAT:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:107
		_go_fuzz_dep_.CoverTab[121579]++
													return hcltoken.Token{Type: hcltoken.FLOAT, Text: t.Text}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:108
		// _ = "end of CoverTab[121579]"
	case NULL:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:109
		_go_fuzz_dep_.CoverTab[121580]++
													return hcltoken.Token{Type: hcltoken.STRING, Text: ""}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:110
		// _ = "end of CoverTab[121580]"
	case NUMBER:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:111
		_go_fuzz_dep_.CoverTab[121581]++
													return hcltoken.Token{Type: hcltoken.NUMBER, Text: t.Text}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:112
		// _ = "end of CoverTab[121581]"
	case STRING:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:113
		_go_fuzz_dep_.CoverTab[121582]++
													return hcltoken.Token{Type: hcltoken.STRING, Text: t.Text, JSON: true}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:114
		// _ = "end of CoverTab[121582]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:115
		_go_fuzz_dep_.CoverTab[121583]++
													panic(fmt.Sprintf("unimplemented HCLToken for type: %s", t.Type))
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:116
		// _ = "end of CoverTab[121583]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:117
	// _ = "end of CoverTab[121577]"
}

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:118
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/token.go:118
var _ = _go_fuzz_dep_.CoverTab
