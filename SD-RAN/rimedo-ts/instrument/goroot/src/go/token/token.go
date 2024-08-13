// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/go/token/token.go:5
// Package token defines constants representing the lexical tokens of the Go
//line /usr/local/go/src/go/token/token.go:5
// programming language and basic operations on tokens (printing, predicates).
//line /usr/local/go/src/go/token/token.go:7
package token

//line /usr/local/go/src/go/token/token.go:7
import (
//line /usr/local/go/src/go/token/token.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/go/token/token.go:7
)
//line /usr/local/go/src/go/token/token.go:7
import (
//line /usr/local/go/src/go/token/token.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/go/token/token.go:7
)

import (
	"strconv"
	"unicode"
	"unicode/utf8"
)

// Token is the set of lexical tokens of the Go programming language.
type Token int

// The list of tokens.
const (
	// Special tokens
	ILLEGAL	Token	= iota
	EOF
	COMMENT

	literal_beg
	// Identifiers and basic type literals
	// (these tokens stand for classes of literals)
	IDENT	// main
	INT	// 12345
	FLOAT	// 123.45
	IMAG	// 123.45i
	CHAR	// 'a'
	STRING	// "abc"
	literal_end

	operator_beg
	// Operators and delimiters
	ADD	// +
	SUB	// -
	MUL	// *
	QUO	// /
	REM	// %

	AND	// &
	OR	// |
	XOR	// ^
	SHL	// <<
	SHR	// >>
	AND_NOT	// &^

	ADD_ASSIGN	// +=
	SUB_ASSIGN	// -=
	MUL_ASSIGN	// *=
	QUO_ASSIGN	// /=
	REM_ASSIGN	// %=

	AND_ASSIGN	// &=
	OR_ASSIGN	// |=
	XOR_ASSIGN	// ^=
	SHL_ASSIGN	// <<=
	SHR_ASSIGN	// >>=
	AND_NOT_ASSIGN	// &^=

	LAND	// &&
	LOR	// ||
	ARROW	// <-
	INC	// ++
	DEC	// --

	EQL	// ==
	LSS	// <
	GTR	// >
	ASSIGN	// =
	NOT	// !

	NEQ		// !=
	LEQ		// <=
	GEQ		// >=
	DEFINE		// :=
	ELLIPSIS	// ...

	LPAREN	// (
	LBRACK	// [
	LBRACE	// {
	COMMA	// ,
	PERIOD	// .

	RPAREN		// )
	RBRACK		// ]
	RBRACE		// }
	SEMICOLON	// ;
	COLON		// :
	operator_end

	keyword_beg
	// Keywords
	BREAK
	CASE
	CHAN
	CONST
	CONTINUE

	DEFAULT
	DEFER
	ELSE
	FALLTHROUGH
	FOR

	FUNC
	GO
	GOTO
	IF
	IMPORT

	INTERFACE
	MAP
	PACKAGE
	RANGE
	RETURN

	SELECT
	STRUCT
	SWITCH
	TYPE
	VAR
	keyword_end

	additional_beg
	// additional tokens, handled in an ad-hoc manner
	TILDE
	additional_end
)

var tokens = [...]string{
	ILLEGAL:	"ILLEGAL",

	EOF:		"EOF",
	COMMENT:	"COMMENT",

	IDENT:	"IDENT",
	INT:	"INT",
	FLOAT:	"FLOAT",
	IMAG:	"IMAG",
	CHAR:	"CHAR",
	STRING:	"STRING",

	ADD:	"+",
	SUB:	"-",
	MUL:	"*",
	QUO:	"/",
	REM:	"%",

	AND:		"&",
	OR:		"|",
	XOR:		"^",
	SHL:		"<<",
	SHR:		">>",
	AND_NOT:	"&^",

	ADD_ASSIGN:	"+=",
	SUB_ASSIGN:	"-=",
	MUL_ASSIGN:	"*=",
	QUO_ASSIGN:	"/=",
	REM_ASSIGN:	"%=",

	AND_ASSIGN:	"&=",
	OR_ASSIGN:	"|=",
	XOR_ASSIGN:	"^=",
	SHL_ASSIGN:	"<<=",
	SHR_ASSIGN:	">>=",
	AND_NOT_ASSIGN:	"&^=",

	LAND:	"&&",
	LOR:	"||",
	ARROW:	"<-",
	INC:	"++",
	DEC:	"--",

	EQL:	"==",
	LSS:	"<",
	GTR:	">",
	ASSIGN:	"=",
	NOT:	"!",

	NEQ:		"!=",
	LEQ:		"<=",
	GEQ:		">=",
	DEFINE:		":=",
	ELLIPSIS:	"...",

	LPAREN:	"(",
	LBRACK:	"[",
	LBRACE:	"{",
	COMMA:	",",
	PERIOD:	".",

	RPAREN:		")",
	RBRACK:		"]",
	RBRACE:		"}",
	SEMICOLON:	";",
	COLON:		":",

	BREAK:		"break",
	CASE:		"case",
	CHAN:		"chan",
	CONST:		"const",
	CONTINUE:	"continue",

	DEFAULT:	"default",
	DEFER:		"defer",
	ELSE:		"else",
	FALLTHROUGH:	"fallthrough",
	FOR:		"for",

	FUNC:	"func",
	GO:	"go",
	GOTO:	"goto",
	IF:	"if",
	IMPORT:	"import",

	INTERFACE:	"interface",
	MAP:		"map",
	PACKAGE:	"package",
	RANGE:		"range",
	RETURN:		"return",

	SELECT:	"select",
	STRUCT:	"struct",
	SWITCH:	"switch",
	TYPE:	"type",
	VAR:	"var",

	TILDE:	"~",
}

// String returns the string corresponding to the token tok.
//line /usr/local/go/src/go/token/token.go:236
// For operators, delimiters, and keywords the string is the actual
//line /usr/local/go/src/go/token/token.go:236
// token character sequence (e.g., for the token ADD, the string is
//line /usr/local/go/src/go/token/token.go:236
// "+"). For all other tokens the string corresponds to the token
//line /usr/local/go/src/go/token/token.go:236
// constant name (e.g. for the token IDENT, the string is "IDENT").
//line /usr/local/go/src/go/token/token.go:241
func (tok Token) String() string {
//line /usr/local/go/src/go/token/token.go:241
	_go_fuzz_dep_.CoverTab[49263]++
						s := ""
						if 0 <= tok && func() bool {
//line /usr/local/go/src/go/token/token.go:243
		_go_fuzz_dep_.CoverTab[49266]++
//line /usr/local/go/src/go/token/token.go:243
		return tok < Token(len(tokens))
//line /usr/local/go/src/go/token/token.go:243
		// _ = "end of CoverTab[49266]"
//line /usr/local/go/src/go/token/token.go:243
	}() {
//line /usr/local/go/src/go/token/token.go:243
		_go_fuzz_dep_.CoverTab[49267]++
							s = tokens[tok]
//line /usr/local/go/src/go/token/token.go:244
		// _ = "end of CoverTab[49267]"
	} else {
//line /usr/local/go/src/go/token/token.go:245
		_go_fuzz_dep_.CoverTab[49268]++
//line /usr/local/go/src/go/token/token.go:245
		// _ = "end of CoverTab[49268]"
//line /usr/local/go/src/go/token/token.go:245
	}
//line /usr/local/go/src/go/token/token.go:245
	// _ = "end of CoverTab[49263]"
//line /usr/local/go/src/go/token/token.go:245
	_go_fuzz_dep_.CoverTab[49264]++
						if s == "" {
//line /usr/local/go/src/go/token/token.go:246
		_go_fuzz_dep_.CoverTab[49269]++
							s = "token(" + strconv.Itoa(int(tok)) + ")"
//line /usr/local/go/src/go/token/token.go:247
		// _ = "end of CoverTab[49269]"
	} else {
//line /usr/local/go/src/go/token/token.go:248
		_go_fuzz_dep_.CoverTab[49270]++
//line /usr/local/go/src/go/token/token.go:248
		// _ = "end of CoverTab[49270]"
//line /usr/local/go/src/go/token/token.go:248
	}
//line /usr/local/go/src/go/token/token.go:248
	// _ = "end of CoverTab[49264]"
//line /usr/local/go/src/go/token/token.go:248
	_go_fuzz_dep_.CoverTab[49265]++
						return s
//line /usr/local/go/src/go/token/token.go:249
	// _ = "end of CoverTab[49265]"
}

// A set of constants for precedence-based expression parsing.
//line /usr/local/go/src/go/token/token.go:252
// Non-operators have lowest precedence, followed by operators
//line /usr/local/go/src/go/token/token.go:252
// starting with precedence 1 up to unary operators. The highest
//line /usr/local/go/src/go/token/token.go:252
// precedence serves as "catch-all" precedence for selector,
//line /usr/local/go/src/go/token/token.go:252
// indexing, and other operator and delimiter tokens.
//line /usr/local/go/src/go/token/token.go:257
const (
	LowestPrec	= 0	// non-operators
	UnaryPrec	= 6
	HighestPrec	= 7
)

// Precedence returns the operator precedence of the binary
//line /usr/local/go/src/go/token/token.go:263
// operator op. If op is not a binary operator, the result
//line /usr/local/go/src/go/token/token.go:263
// is LowestPrecedence.
//line /usr/local/go/src/go/token/token.go:266
func (op Token) Precedence() int {
//line /usr/local/go/src/go/token/token.go:266
	_go_fuzz_dep_.CoverTab[49271]++
						switch op {
	case LOR:
//line /usr/local/go/src/go/token/token.go:268
		_go_fuzz_dep_.CoverTab[49273]++
							return 1
//line /usr/local/go/src/go/token/token.go:269
		// _ = "end of CoverTab[49273]"
	case LAND:
//line /usr/local/go/src/go/token/token.go:270
		_go_fuzz_dep_.CoverTab[49274]++
							return 2
//line /usr/local/go/src/go/token/token.go:271
		// _ = "end of CoverTab[49274]"
	case EQL, NEQ, LSS, LEQ, GTR, GEQ:
//line /usr/local/go/src/go/token/token.go:272
		_go_fuzz_dep_.CoverTab[49275]++
							return 3
//line /usr/local/go/src/go/token/token.go:273
		// _ = "end of CoverTab[49275]"
	case ADD, SUB, OR, XOR:
//line /usr/local/go/src/go/token/token.go:274
		_go_fuzz_dep_.CoverTab[49276]++
							return 4
//line /usr/local/go/src/go/token/token.go:275
		// _ = "end of CoverTab[49276]"
	case MUL, QUO, REM, SHL, SHR, AND, AND_NOT:
//line /usr/local/go/src/go/token/token.go:276
		_go_fuzz_dep_.CoverTab[49277]++
							return 5
//line /usr/local/go/src/go/token/token.go:277
		// _ = "end of CoverTab[49277]"
//line /usr/local/go/src/go/token/token.go:277
	default:
//line /usr/local/go/src/go/token/token.go:277
		_go_fuzz_dep_.CoverTab[49278]++
//line /usr/local/go/src/go/token/token.go:277
		// _ = "end of CoverTab[49278]"
	}
//line /usr/local/go/src/go/token/token.go:278
	// _ = "end of CoverTab[49271]"
//line /usr/local/go/src/go/token/token.go:278
	_go_fuzz_dep_.CoverTab[49272]++
						return LowestPrec
//line /usr/local/go/src/go/token/token.go:279
	// _ = "end of CoverTab[49272]"
}

var keywords map[string]Token

func init() {
	keywords = make(map[string]Token, keyword_end-(keyword_beg+1))
	for i := keyword_beg + 1; i < keyword_end; i++ {
		keywords[tokens[i]] = i
	}
}

// Lookup maps an identifier to its keyword token or IDENT (if not a keyword).
func Lookup(ident string) Token {
//line /usr/local/go/src/go/token/token.go:292
	_go_fuzz_dep_.CoverTab[49279]++
						if tok, is_keyword := keywords[ident]; is_keyword {
//line /usr/local/go/src/go/token/token.go:293
		_go_fuzz_dep_.CoverTab[49281]++
							return tok
//line /usr/local/go/src/go/token/token.go:294
		// _ = "end of CoverTab[49281]"
	} else {
//line /usr/local/go/src/go/token/token.go:295
		_go_fuzz_dep_.CoverTab[49282]++
//line /usr/local/go/src/go/token/token.go:295
		// _ = "end of CoverTab[49282]"
//line /usr/local/go/src/go/token/token.go:295
	}
//line /usr/local/go/src/go/token/token.go:295
	// _ = "end of CoverTab[49279]"
//line /usr/local/go/src/go/token/token.go:295
	_go_fuzz_dep_.CoverTab[49280]++
						return IDENT
//line /usr/local/go/src/go/token/token.go:296
	// _ = "end of CoverTab[49280]"
}

//line /usr/local/go/src/go/token/token.go:301
// IsLiteral returns true for tokens corresponding to identifiers
//line /usr/local/go/src/go/token/token.go:301
// and basic type literals; it returns false otherwise.
//line /usr/local/go/src/go/token/token.go:303
func (tok Token) IsLiteral() bool {
//line /usr/local/go/src/go/token/token.go:303
	_go_fuzz_dep_.CoverTab[49283]++
//line /usr/local/go/src/go/token/token.go:303
	return literal_beg < tok && func() bool {
//line /usr/local/go/src/go/token/token.go:303
		_go_fuzz_dep_.CoverTab[49284]++
//line /usr/local/go/src/go/token/token.go:303
		return tok < literal_end
//line /usr/local/go/src/go/token/token.go:303
		// _ = "end of CoverTab[49284]"
//line /usr/local/go/src/go/token/token.go:303
	}()
//line /usr/local/go/src/go/token/token.go:303
	// _ = "end of CoverTab[49283]"
//line /usr/local/go/src/go/token/token.go:303
}

// IsOperator returns true for tokens corresponding to operators and
//line /usr/local/go/src/go/token/token.go:305
// delimiters; it returns false otherwise.
//line /usr/local/go/src/go/token/token.go:307
func (tok Token) IsOperator() bool {
//line /usr/local/go/src/go/token/token.go:307
	_go_fuzz_dep_.CoverTab[49285]++
						return (operator_beg < tok && func() bool {
//line /usr/local/go/src/go/token/token.go:308
		_go_fuzz_dep_.CoverTab[49286]++
//line /usr/local/go/src/go/token/token.go:308
		return tok < operator_end
//line /usr/local/go/src/go/token/token.go:308
		// _ = "end of CoverTab[49286]"
//line /usr/local/go/src/go/token/token.go:308
	}()) || func() bool {
//line /usr/local/go/src/go/token/token.go:308
		_go_fuzz_dep_.CoverTab[49287]++
//line /usr/local/go/src/go/token/token.go:308
		return tok == TILDE
//line /usr/local/go/src/go/token/token.go:308
		// _ = "end of CoverTab[49287]"
//line /usr/local/go/src/go/token/token.go:308
	}()
//line /usr/local/go/src/go/token/token.go:308
	// _ = "end of CoverTab[49285]"
}

// IsKeyword returns true for tokens corresponding to keywords;
//line /usr/local/go/src/go/token/token.go:311
// it returns false otherwise.
//line /usr/local/go/src/go/token/token.go:313
func (tok Token) IsKeyword() bool {
//line /usr/local/go/src/go/token/token.go:313
	_go_fuzz_dep_.CoverTab[49288]++
//line /usr/local/go/src/go/token/token.go:313
	return keyword_beg < tok && func() bool {
//line /usr/local/go/src/go/token/token.go:313
		_go_fuzz_dep_.CoverTab[49289]++
//line /usr/local/go/src/go/token/token.go:313
		return tok < keyword_end
//line /usr/local/go/src/go/token/token.go:313
		// _ = "end of CoverTab[49289]"
//line /usr/local/go/src/go/token/token.go:313
	}()
//line /usr/local/go/src/go/token/token.go:313
	// _ = "end of CoverTab[49288]"
//line /usr/local/go/src/go/token/token.go:313
}

// IsExported reports whether name starts with an upper-case letter.
func IsExported(name string) bool {
//line /usr/local/go/src/go/token/token.go:316
	_go_fuzz_dep_.CoverTab[49290]++
						ch, _ := utf8.DecodeRuneInString(name)
						return unicode.IsUpper(ch)
//line /usr/local/go/src/go/token/token.go:318
	// _ = "end of CoverTab[49290]"
}

// IsKeyword reports whether name is a Go keyword, such as "func" or "return".
func IsKeyword(name string) bool {
//line /usr/local/go/src/go/token/token.go:322
	_go_fuzz_dep_.CoverTab[49291]++

						_, ok := keywords[name]
						return ok
//line /usr/local/go/src/go/token/token.go:325
	// _ = "end of CoverTab[49291]"
}

// IsIdentifier reports whether name is a Go identifier, that is, a non-empty
//line /usr/local/go/src/go/token/token.go:328
// string made up of letters, digits, and underscores, where the first character
//line /usr/local/go/src/go/token/token.go:328
// is not a digit. Keywords are not identifiers.
//line /usr/local/go/src/go/token/token.go:331
func IsIdentifier(name string) bool {
//line /usr/local/go/src/go/token/token.go:331
	_go_fuzz_dep_.CoverTab[49292]++
						if name == "" || func() bool {
//line /usr/local/go/src/go/token/token.go:332
		_go_fuzz_dep_.CoverTab[49295]++
//line /usr/local/go/src/go/token/token.go:332
		return IsKeyword(name)
//line /usr/local/go/src/go/token/token.go:332
		// _ = "end of CoverTab[49295]"
//line /usr/local/go/src/go/token/token.go:332
	}() {
//line /usr/local/go/src/go/token/token.go:332
		_go_fuzz_dep_.CoverTab[49296]++
							return false
//line /usr/local/go/src/go/token/token.go:333
		// _ = "end of CoverTab[49296]"
	} else {
//line /usr/local/go/src/go/token/token.go:334
		_go_fuzz_dep_.CoverTab[49297]++
//line /usr/local/go/src/go/token/token.go:334
		// _ = "end of CoverTab[49297]"
//line /usr/local/go/src/go/token/token.go:334
	}
//line /usr/local/go/src/go/token/token.go:334
	// _ = "end of CoverTab[49292]"
//line /usr/local/go/src/go/token/token.go:334
	_go_fuzz_dep_.CoverTab[49293]++
						for i, c := range name {
//line /usr/local/go/src/go/token/token.go:335
		_go_fuzz_dep_.CoverTab[49298]++
							if !unicode.IsLetter(c) && func() bool {
//line /usr/local/go/src/go/token/token.go:336
			_go_fuzz_dep_.CoverTab[49299]++
//line /usr/local/go/src/go/token/token.go:336
			return c != '_'
//line /usr/local/go/src/go/token/token.go:336
			// _ = "end of CoverTab[49299]"
//line /usr/local/go/src/go/token/token.go:336
		}() && func() bool {
//line /usr/local/go/src/go/token/token.go:336
			_go_fuzz_dep_.CoverTab[49300]++
//line /usr/local/go/src/go/token/token.go:336
			return (i == 0 || func() bool {
//line /usr/local/go/src/go/token/token.go:336
				_go_fuzz_dep_.CoverTab[49301]++
//line /usr/local/go/src/go/token/token.go:336
				return !unicode.IsDigit(c)
//line /usr/local/go/src/go/token/token.go:336
				// _ = "end of CoverTab[49301]"
//line /usr/local/go/src/go/token/token.go:336
			}())
//line /usr/local/go/src/go/token/token.go:336
			// _ = "end of CoverTab[49300]"
//line /usr/local/go/src/go/token/token.go:336
		}() {
//line /usr/local/go/src/go/token/token.go:336
			_go_fuzz_dep_.CoverTab[49302]++
								return false
//line /usr/local/go/src/go/token/token.go:337
			// _ = "end of CoverTab[49302]"
		} else {
//line /usr/local/go/src/go/token/token.go:338
			_go_fuzz_dep_.CoverTab[49303]++
//line /usr/local/go/src/go/token/token.go:338
			// _ = "end of CoverTab[49303]"
//line /usr/local/go/src/go/token/token.go:338
		}
//line /usr/local/go/src/go/token/token.go:338
		// _ = "end of CoverTab[49298]"
	}
//line /usr/local/go/src/go/token/token.go:339
	// _ = "end of CoverTab[49293]"
//line /usr/local/go/src/go/token/token.go:339
	_go_fuzz_dep_.CoverTab[49294]++
						return true
//line /usr/local/go/src/go/token/token.go:340
	// _ = "end of CoverTab[49294]"
}

//line /usr/local/go/src/go/token/token.go:341
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/go/token/token.go:341
var _ = _go_fuzz_dep_.CoverTab
