//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:1
// Package parser implements a parser for HCL (HashiCorp Configuration
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:1
// Language)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:3
package parser

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:3
import (
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:3
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:3
)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:3
import (
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:3
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:3
)

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"github.com/hashicorp/hcl/hcl/ast"
	"github.com/hashicorp/hcl/hcl/scanner"
	"github.com/hashicorp/hcl/hcl/token"
)

type Parser struct {
	sc	*scanner.Scanner

	// Last read token
	tok		token.Token
	commaPrev	token.Token

	comments	[]*ast.CommentGroup
	leadComment	*ast.CommentGroup	// last lead comment
	lineComment	*ast.CommentGroup	// last line comment

	enableTrace	bool
	indent		int
	n		int	// buffer size (max = 1)
}

func newParser(src []byte) *Parser {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:32
	_go_fuzz_dep_.CoverTab[121381]++
												return &Parser{
		sc: scanner.New(src),
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:35
	// _ = "end of CoverTab[121381]"
}

// Parse returns the fully parsed source and returns the abstract syntax tree.
func Parse(src []byte) (*ast.File, error) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:39
	_go_fuzz_dep_.CoverTab[121382]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:43
	src = bytes.Replace(src, []byte("\r\n"), []byte("\n"), -1)

												p := newParser(src)
												return p.Parse()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:46
	// _ = "end of CoverTab[121382]"
}

var errEofToken = errors.New("EOF token found")

// Parse returns the fully parsed source and returns the abstract syntax tree.
func (p *Parser) Parse() (*ast.File, error) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:52
	_go_fuzz_dep_.CoverTab[121383]++
												f := &ast.File{}
												var err, scerr error
												p.sc.Error = func(pos token.Pos, msg string) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:55
		_go_fuzz_dep_.CoverTab[121387]++
													scerr = &PosError{Pos: pos, Err: errors.New(msg)}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:56
		// _ = "end of CoverTab[121387]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:57
	// _ = "end of CoverTab[121383]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:57
	_go_fuzz_dep_.CoverTab[121384]++

												f.Node, err = p.objectList(false)
												if scerr != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:60
		_go_fuzz_dep_.CoverTab[121388]++
													return nil, scerr
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:61
		// _ = "end of CoverTab[121388]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:62
		_go_fuzz_dep_.CoverTab[121389]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:62
		// _ = "end of CoverTab[121389]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:62
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:62
	// _ = "end of CoverTab[121384]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:62
	_go_fuzz_dep_.CoverTab[121385]++
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:63
		_go_fuzz_dep_.CoverTab[121390]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:64
		// _ = "end of CoverTab[121390]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:65
		_go_fuzz_dep_.CoverTab[121391]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:65
		// _ = "end of CoverTab[121391]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:65
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:65
	// _ = "end of CoverTab[121385]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:65
	_go_fuzz_dep_.CoverTab[121386]++

												f.Comments = p.comments
												return f, nil
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:68
	// _ = "end of CoverTab[121386]"
}

// objectList parses a list of items within an object (generally k/v pairs).
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:71
// The parameter" obj" tells this whether to we are within an object (braces:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:71
// '{', '}') or just at the top level. If we're within an object, we end
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:71
// at an RBRACE.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:75
func (p *Parser) objectList(obj bool) (*ast.ObjectList, error) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:75
	_go_fuzz_dep_.CoverTab[121392]++
												defer un(trace(p, "ParseObjectList"))
												node := &ast.ObjectList{}

												for {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:79
		_go_fuzz_dep_.CoverTab[121394]++
													if obj {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:80
			_go_fuzz_dep_.CoverTab[121398]++
														tok := p.scan()
														p.unscan()
														if tok.Type == token.RBRACE {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:83
				_go_fuzz_dep_.CoverTab[121399]++
															break
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:84
				// _ = "end of CoverTab[121399]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:85
				_go_fuzz_dep_.CoverTab[121400]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:85
				// _ = "end of CoverTab[121400]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:85
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:85
			// _ = "end of CoverTab[121398]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:86
			_go_fuzz_dep_.CoverTab[121401]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:86
			// _ = "end of CoverTab[121401]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:86
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:86
		// _ = "end of CoverTab[121394]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:86
		_go_fuzz_dep_.CoverTab[121395]++

													n, err := p.objectItem()
													if err == errEofToken {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:89
			_go_fuzz_dep_.CoverTab[121402]++
														break
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:90
			// _ = "end of CoverTab[121402]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:91
			_go_fuzz_dep_.CoverTab[121403]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:91
			// _ = "end of CoverTab[121403]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:91
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:91
		// _ = "end of CoverTab[121395]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:91
		_go_fuzz_dep_.CoverTab[121396]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:95
		if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:95
			_go_fuzz_dep_.CoverTab[121404]++
														return node, err
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:96
			// _ = "end of CoverTab[121404]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:97
			_go_fuzz_dep_.CoverTab[121405]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:97
			// _ = "end of CoverTab[121405]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:97
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:97
		// _ = "end of CoverTab[121396]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:97
		_go_fuzz_dep_.CoverTab[121397]++

													node.Add(n)

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:103
		tok := p.scan()
		if tok.Type != token.COMMA {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:104
			_go_fuzz_dep_.CoverTab[121406]++
														p.unscan()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:105
			// _ = "end of CoverTab[121406]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:106
			_go_fuzz_dep_.CoverTab[121407]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:106
			// _ = "end of CoverTab[121407]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:106
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:106
		// _ = "end of CoverTab[121397]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:107
	// _ = "end of CoverTab[121392]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:107
	_go_fuzz_dep_.CoverTab[121393]++
												return node, nil
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:108
	// _ = "end of CoverTab[121393]"
}

func (p *Parser) consumeComment() (comment *ast.Comment, endline int) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:111
	_go_fuzz_dep_.CoverTab[121408]++
												endline = p.tok.Pos.Line

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:115
	if len(p.tok.Text) > 1 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:115
		_go_fuzz_dep_.CoverTab[121410]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:115
		return p.tok.Text[1] == '*'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:115
		// _ = "end of CoverTab[121410]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:115
	}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:115
		_go_fuzz_dep_.CoverTab[121411]++

													for i := 0; i < len(p.tok.Text); i++ {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:117
			_go_fuzz_dep_.CoverTab[121412]++
														if p.tok.Text[i] == '\n' {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:118
				_go_fuzz_dep_.CoverTab[121413]++
															endline++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:119
				// _ = "end of CoverTab[121413]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:120
				_go_fuzz_dep_.CoverTab[121414]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:120
				// _ = "end of CoverTab[121414]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:120
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:120
			// _ = "end of CoverTab[121412]"
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:121
		// _ = "end of CoverTab[121411]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:122
		_go_fuzz_dep_.CoverTab[121415]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:122
		// _ = "end of CoverTab[121415]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:122
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:122
	// _ = "end of CoverTab[121408]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:122
	_go_fuzz_dep_.CoverTab[121409]++

												comment = &ast.Comment{Start: p.tok.Pos, Text: p.tok.Text}
												p.tok = p.sc.Scan()
												return
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:126
	// _ = "end of CoverTab[121409]"
}

func (p *Parser) consumeCommentGroup(n int) (comments *ast.CommentGroup, endline int) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:129
	_go_fuzz_dep_.CoverTab[121416]++
												var list []*ast.Comment
												endline = p.tok.Pos.Line

												for p.tok.Type == token.COMMENT && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:133
		_go_fuzz_dep_.CoverTab[121418]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:133
		return p.tok.Pos.Line <= endline+n
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:133
		// _ = "end of CoverTab[121418]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:133
	}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:133
		_go_fuzz_dep_.CoverTab[121419]++
													var comment *ast.Comment
													comment, endline = p.consumeComment()
													list = append(list, comment)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:136
		// _ = "end of CoverTab[121419]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:137
	// _ = "end of CoverTab[121416]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:137
	_go_fuzz_dep_.CoverTab[121417]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:140
	comments = &ast.CommentGroup{List: list}
												p.comments = append(p.comments, comments)

												return
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:143
	// _ = "end of CoverTab[121417]"
}

// objectItem parses a single object item
func (p *Parser) objectItem() (*ast.ObjectItem, error) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:147
	_go_fuzz_dep_.CoverTab[121420]++
												defer un(trace(p, "ParseObjectItem"))

												keys, err := p.objectKey()
												if len(keys) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:151
		_go_fuzz_dep_.CoverTab[121428]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:151
		return err == errEofToken
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:151
		// _ = "end of CoverTab[121428]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:151
	}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:151
		_go_fuzz_dep_.CoverTab[121429]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:154
		err = nil
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:154
		// _ = "end of CoverTab[121429]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:155
		_go_fuzz_dep_.CoverTab[121430]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:155
		// _ = "end of CoverTab[121430]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:155
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:155
	// _ = "end of CoverTab[121420]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:155
	_go_fuzz_dep_.CoverTab[121421]++
												if len(keys) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:156
		_go_fuzz_dep_.CoverTab[121431]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:156
		return err != nil
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:156
		// _ = "end of CoverTab[121431]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:156
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:156
		_go_fuzz_dep_.CoverTab[121432]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:156
		return p.tok.Type == token.RBRACE
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:156
		// _ = "end of CoverTab[121432]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:156
	}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:156
		_go_fuzz_dep_.CoverTab[121433]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:162
		err = nil

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:167
		p.tok.Type = token.EOF
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:167
		// _ = "end of CoverTab[121433]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:168
		_go_fuzz_dep_.CoverTab[121434]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:168
		// _ = "end of CoverTab[121434]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:168
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:168
	// _ = "end of CoverTab[121421]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:168
	_go_fuzz_dep_.CoverTab[121422]++
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:169
		_go_fuzz_dep_.CoverTab[121435]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:170
		// _ = "end of CoverTab[121435]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:171
		_go_fuzz_dep_.CoverTab[121436]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:171
		// _ = "end of CoverTab[121436]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:171
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:171
	// _ = "end of CoverTab[121422]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:171
	_go_fuzz_dep_.CoverTab[121423]++

												o := &ast.ObjectItem{
		Keys: keys,
	}

	if p.leadComment != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:177
		_go_fuzz_dep_.CoverTab[121437]++
													o.LeadComment = p.leadComment
													p.leadComment = nil
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:179
		// _ = "end of CoverTab[121437]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:180
		_go_fuzz_dep_.CoverTab[121438]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:180
		// _ = "end of CoverTab[121438]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:180
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:180
	// _ = "end of CoverTab[121423]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:180
	_go_fuzz_dep_.CoverTab[121424]++

												switch p.tok.Type {
	case token.ASSIGN:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:183
		_go_fuzz_dep_.CoverTab[121439]++
													o.Assign = p.tok.Pos
													o.Val, err = p.object()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:186
			_go_fuzz_dep_.CoverTab[121443]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:187
			// _ = "end of CoverTab[121443]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:188
			_go_fuzz_dep_.CoverTab[121444]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:188
			// _ = "end of CoverTab[121444]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:188
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:188
		// _ = "end of CoverTab[121439]"
	case token.LBRACE:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:189
		_go_fuzz_dep_.CoverTab[121440]++
													o.Val, err = p.objectType()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:191
			_go_fuzz_dep_.CoverTab[121445]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:192
			// _ = "end of CoverTab[121445]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:193
			_go_fuzz_dep_.CoverTab[121446]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:193
			// _ = "end of CoverTab[121446]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:193
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:193
		// _ = "end of CoverTab[121440]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:194
		_go_fuzz_dep_.CoverTab[121441]++
													keyStr := make([]string, 0, len(keys))
													for _, k := range keys {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:196
			_go_fuzz_dep_.CoverTab[121447]++
														keyStr = append(keyStr, k.Token.Text)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:197
			// _ = "end of CoverTab[121447]"
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:198
		// _ = "end of CoverTab[121441]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:198
		_go_fuzz_dep_.CoverTab[121442]++

													return nil, &PosError{
			Pos:	p.tok.Pos,
			Err: fmt.Errorf(
				"key '%s' expected start of object ('{') or assignment ('=')",
				strings.Join(keyStr, " ")),
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:205
		// _ = "end of CoverTab[121442]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:206
	// _ = "end of CoverTab[121424]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:206
	_go_fuzz_dep_.CoverTab[121425]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:210
	if p.lineComment != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:210
		_go_fuzz_dep_.CoverTab[121448]++
													o.LineComment, p.lineComment = p.lineComment, nil
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:211
		// _ = "end of CoverTab[121448]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:212
		_go_fuzz_dep_.CoverTab[121449]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:212
		// _ = "end of CoverTab[121449]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:212
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:212
	// _ = "end of CoverTab[121425]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:212
	_go_fuzz_dep_.CoverTab[121426]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:215
	p.scan()
	if len(keys) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:216
		_go_fuzz_dep_.CoverTab[121450]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:216
		return o.Val.Pos().Line == keys[0].Pos().Line
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:216
		// _ = "end of CoverTab[121450]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:216
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:216
		_go_fuzz_dep_.CoverTab[121451]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:216
		return p.lineComment != nil
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:216
		// _ = "end of CoverTab[121451]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:216
	}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:216
		_go_fuzz_dep_.CoverTab[121452]++
													o.LineComment = p.lineComment
													p.lineComment = nil
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:218
		// _ = "end of CoverTab[121452]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:219
		_go_fuzz_dep_.CoverTab[121453]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:219
		// _ = "end of CoverTab[121453]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:219
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:219
	// _ = "end of CoverTab[121426]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:219
	_go_fuzz_dep_.CoverTab[121427]++
												p.unscan()
												return o, nil
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:221
	// _ = "end of CoverTab[121427]"
}

// objectKey parses an object key and returns a ObjectKey AST
func (p *Parser) objectKey() ([]*ast.ObjectKey, error) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:225
	_go_fuzz_dep_.CoverTab[121454]++
												keyCount := 0
												keys := make([]*ast.ObjectKey, 0)

												for {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:229
		_go_fuzz_dep_.CoverTab[121455]++
													tok := p.scan()
													switch tok.Type {
		case token.EOF:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:232
			_go_fuzz_dep_.CoverTab[121456]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:237
			return keys, errEofToken
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:237
			// _ = "end of CoverTab[121456]"
		case token.ASSIGN:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:238
			_go_fuzz_dep_.CoverTab[121457]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:241
			if keyCount > 1 {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:241
				_go_fuzz_dep_.CoverTab[121465]++
															return nil, &PosError{
					Pos:	p.tok.Pos,
					Err:	fmt.Errorf("nested object expected: LBRACE got: %s", p.tok.Type),
				}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:245
				// _ = "end of CoverTab[121465]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:246
				_go_fuzz_dep_.CoverTab[121466]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:246
				// _ = "end of CoverTab[121466]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:246
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:246
			// _ = "end of CoverTab[121457]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:246
			_go_fuzz_dep_.CoverTab[121458]++

														if keyCount == 0 {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:248
				_go_fuzz_dep_.CoverTab[121467]++
															return nil, &PosError{
					Pos:	p.tok.Pos,
					Err:	errors.New("no object keys found!"),
				}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:252
				// _ = "end of CoverTab[121467]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:253
				_go_fuzz_dep_.CoverTab[121468]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:253
				// _ = "end of CoverTab[121468]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:253
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:253
			// _ = "end of CoverTab[121458]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:253
			_go_fuzz_dep_.CoverTab[121459]++

														return keys, nil
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:255
			// _ = "end of CoverTab[121459]"
		case token.LBRACE:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:256
			_go_fuzz_dep_.CoverTab[121460]++
														var err error

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:261
			if len(keys) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:261
				_go_fuzz_dep_.CoverTab[121469]++
															err = &PosError{
					Pos:	p.tok.Pos,
					Err:	fmt.Errorf("expected: IDENT | STRING got: %s", p.tok.Type),
				}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:265
				// _ = "end of CoverTab[121469]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:266
				_go_fuzz_dep_.CoverTab[121470]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:266
				// _ = "end of CoverTab[121470]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:266
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:266
			// _ = "end of CoverTab[121460]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:266
			_go_fuzz_dep_.CoverTab[121461]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:269
			return keys, err
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:269
			// _ = "end of CoverTab[121461]"
		case token.IDENT, token.STRING:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:270
			_go_fuzz_dep_.CoverTab[121462]++
														keyCount++
														keys = append(keys, &ast.ObjectKey{Token: p.tok})
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:272
			// _ = "end of CoverTab[121462]"
		case token.ILLEGAL:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:273
			_go_fuzz_dep_.CoverTab[121463]++
														return keys, &PosError{
				Pos:	p.tok.Pos,
				Err:	fmt.Errorf("illegal character"),
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:277
			// _ = "end of CoverTab[121463]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:278
			_go_fuzz_dep_.CoverTab[121464]++
														return keys, &PosError{
				Pos:	p.tok.Pos,
				Err:	fmt.Errorf("expected: IDENT | STRING | ASSIGN | LBRACE got: %s", p.tok.Type),
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:282
			// _ = "end of CoverTab[121464]"
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:283
		// _ = "end of CoverTab[121455]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:284
	// _ = "end of CoverTab[121454]"
}

// object parses any type of object, such as number, bool, string, object or
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:287
// list.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:289
func (p *Parser) object() (ast.Node, error) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:289
	_go_fuzz_dep_.CoverTab[121471]++
												defer un(trace(p, "ParseType"))
												tok := p.scan()

												switch tok.Type {
	case token.NUMBER, token.FLOAT, token.BOOL, token.STRING, token.HEREDOC:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:294
		_go_fuzz_dep_.CoverTab[121473]++
													return p.literalType()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:295
		// _ = "end of CoverTab[121473]"
	case token.LBRACE:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:296
		_go_fuzz_dep_.CoverTab[121474]++
													return p.objectType()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:297
		// _ = "end of CoverTab[121474]"
	case token.LBRACK:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:298
		_go_fuzz_dep_.CoverTab[121475]++
													return p.listType()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:299
		// _ = "end of CoverTab[121475]"
	case token.COMMENT:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:300
		_go_fuzz_dep_.CoverTab[121476]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:300
		// _ = "end of CoverTab[121476]"

	case token.EOF:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:302
		_go_fuzz_dep_.CoverTab[121477]++
													return nil, errEofToken
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:303
		// _ = "end of CoverTab[121477]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:303
	default:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:303
		_go_fuzz_dep_.CoverTab[121478]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:303
		// _ = "end of CoverTab[121478]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:304
	// _ = "end of CoverTab[121471]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:304
	_go_fuzz_dep_.CoverTab[121472]++

												return nil, &PosError{
		Pos:	tok.Pos,
		Err:	fmt.Errorf("Unknown token: %+v", tok),
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:309
	// _ = "end of CoverTab[121472]"
}

// objectType parses an object type and returns a ObjectType AST
func (p *Parser) objectType() (*ast.ObjectType, error) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:313
	_go_fuzz_dep_.CoverTab[121479]++
												defer un(trace(p, "ParseObjectType"))

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:317
	o := &ast.ObjectType{
		Lbrace: p.tok.Pos,
	}

												l, err := p.objectList(true)

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:325
	if err != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:325
		_go_fuzz_dep_.CoverTab[121482]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:325
		return p.tok.Type != token.RBRACE
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:325
		// _ = "end of CoverTab[121482]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:325
	}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:325
		_go_fuzz_dep_.CoverTab[121483]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:326
		// _ = "end of CoverTab[121483]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:327
		_go_fuzz_dep_.CoverTab[121484]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:327
		// _ = "end of CoverTab[121484]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:327
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:327
	// _ = "end of CoverTab[121479]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:327
	_go_fuzz_dep_.CoverTab[121480]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:330
	if tok := p.scan(); tok.Type != token.RBRACE {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:330
		_go_fuzz_dep_.CoverTab[121485]++
													return nil, &PosError{
			Pos:	tok.Pos,
			Err:	fmt.Errorf("object expected closing RBRACE got: %s", tok.Type),
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:334
		// _ = "end of CoverTab[121485]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:335
		_go_fuzz_dep_.CoverTab[121486]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:335
		// _ = "end of CoverTab[121486]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:335
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:335
	// _ = "end of CoverTab[121480]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:335
	_go_fuzz_dep_.CoverTab[121481]++

												o.List = l
												o.Rbrace = p.tok.Pos
												return o, nil
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:339
	// _ = "end of CoverTab[121481]"
}

// listType parses a list type and returns a ListType AST
func (p *Parser) listType() (*ast.ListType, error) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:343
	_go_fuzz_dep_.CoverTab[121487]++
												defer un(trace(p, "ParseListType"))

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:347
	l := &ast.ListType{
		Lbrack: p.tok.Pos,
	}

	needComma := false
	for {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:352
		_go_fuzz_dep_.CoverTab[121488]++
													tok := p.scan()
													if needComma {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:354
			_go_fuzz_dep_.CoverTab[121490]++
														switch tok.Type {
			case token.COMMA, token.RBRACK:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:356
				_go_fuzz_dep_.CoverTab[121491]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:356
				// _ = "end of CoverTab[121491]"
			default:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:357
				_go_fuzz_dep_.CoverTab[121492]++
															return nil, &PosError{
					Pos:	tok.Pos,
					Err: fmt.Errorf(
						"error parsing list, expected comma or list end, got: %s",
						tok.Type),
				}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:363
				// _ = "end of CoverTab[121492]"
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:364
			// _ = "end of CoverTab[121490]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:365
			_go_fuzz_dep_.CoverTab[121493]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:365
			// _ = "end of CoverTab[121493]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:365
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:365
		// _ = "end of CoverTab[121488]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:365
		_go_fuzz_dep_.CoverTab[121489]++
													switch tok.Type {
		case token.BOOL, token.NUMBER, token.FLOAT, token.STRING, token.HEREDOC:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:367
			_go_fuzz_dep_.CoverTab[121494]++
														node, err := p.literalType()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:369
				_go_fuzz_dep_.CoverTab[121505]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:370
				// _ = "end of CoverTab[121505]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:371
				_go_fuzz_dep_.CoverTab[121506]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:371
				// _ = "end of CoverTab[121506]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:371
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:371
			// _ = "end of CoverTab[121494]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:371
			_go_fuzz_dep_.CoverTab[121495]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:374
			if p.leadComment != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:374
				_go_fuzz_dep_.CoverTab[121507]++
															node.LeadComment = p.leadComment
															p.leadComment = nil
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:376
				// _ = "end of CoverTab[121507]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:377
				_go_fuzz_dep_.CoverTab[121508]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:377
				// _ = "end of CoverTab[121508]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:377
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:377
			// _ = "end of CoverTab[121495]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:377
			_go_fuzz_dep_.CoverTab[121496]++

														l.Add(node)
														needComma = true
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:380
			// _ = "end of CoverTab[121496]"
		case token.COMMA:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:381
			_go_fuzz_dep_.CoverTab[121497]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:384
			p.scan()
			if p.lineComment != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:385
				_go_fuzz_dep_.CoverTab[121509]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:385
				return len(l.List) > 0
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:385
				// _ = "end of CoverTab[121509]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:385
			}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:385
				_go_fuzz_dep_.CoverTab[121510]++
															lit, ok := l.List[len(l.List)-1].(*ast.LiteralType)
															if ok {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:387
					_go_fuzz_dep_.CoverTab[121511]++
																lit.LineComment = p.lineComment
																l.List[len(l.List)-1] = lit
																p.lineComment = nil
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:390
					// _ = "end of CoverTab[121511]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:391
					_go_fuzz_dep_.CoverTab[121512]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:391
					// _ = "end of CoverTab[121512]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:391
				}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:391
				// _ = "end of CoverTab[121510]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:392
				_go_fuzz_dep_.CoverTab[121513]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:392
				// _ = "end of CoverTab[121513]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:392
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:392
			// _ = "end of CoverTab[121497]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:392
			_go_fuzz_dep_.CoverTab[121498]++
														p.unscan()

														needComma = false
														continue
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:396
			// _ = "end of CoverTab[121498]"
		case token.LBRACE:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:397
			_go_fuzz_dep_.CoverTab[121499]++

														node, err := p.objectType()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:400
				_go_fuzz_dep_.CoverTab[121514]++
															return nil, &PosError{
					Pos:	tok.Pos,
					Err: fmt.Errorf(
						"error while trying to parse object within list: %s", err),
				}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:405
				// _ = "end of CoverTab[121514]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:406
				_go_fuzz_dep_.CoverTab[121515]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:406
				// _ = "end of CoverTab[121515]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:406
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:406
			// _ = "end of CoverTab[121499]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:406
			_go_fuzz_dep_.CoverTab[121500]++
														l.Add(node)
														needComma = true
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:408
			// _ = "end of CoverTab[121500]"
		case token.LBRACK:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:409
			_go_fuzz_dep_.CoverTab[121501]++
														node, err := p.listType()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:411
				_go_fuzz_dep_.CoverTab[121516]++
															return nil, &PosError{
					Pos:	tok.Pos,
					Err: fmt.Errorf(
						"error while trying to parse list within list: %s", err),
				}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:416
				// _ = "end of CoverTab[121516]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:417
				_go_fuzz_dep_.CoverTab[121517]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:417
				// _ = "end of CoverTab[121517]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:417
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:417
			// _ = "end of CoverTab[121501]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:417
			_go_fuzz_dep_.CoverTab[121502]++
														l.Add(node)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:418
			// _ = "end of CoverTab[121502]"
		case token.RBRACK:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:419
			_go_fuzz_dep_.CoverTab[121503]++

														l.Rbrack = p.tok.Pos
														return l, nil
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:422
			// _ = "end of CoverTab[121503]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:423
			_go_fuzz_dep_.CoverTab[121504]++
														return nil, &PosError{
				Pos:	tok.Pos,
				Err:	fmt.Errorf("unexpected token while parsing list: %s", tok.Type),
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:427
			// _ = "end of CoverTab[121504]"
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:428
		// _ = "end of CoverTab[121489]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:429
	// _ = "end of CoverTab[121487]"
}

// literalType parses a literal type and returns a LiteralType AST
func (p *Parser) literalType() (*ast.LiteralType, error) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:433
	_go_fuzz_dep_.CoverTab[121518]++
												defer un(trace(p, "ParseLiteral"))

												return &ast.LiteralType{
		Token: p.tok,
	}, nil
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:438
	// _ = "end of CoverTab[121518]"
}

// scan returns the next token from the underlying scanner. If a token has
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:441
// been unscanned then read that instead. In the process, it collects any
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:441
// comment groups encountered, and remembers the last lead and line comments.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:444
func (p *Parser) scan() token.Token {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:444
	_go_fuzz_dep_.CoverTab[121519]++

												if p.n != 0 {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:446
		_go_fuzz_dep_.CoverTab[121522]++
													p.n = 0
													return p.tok
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:448
		// _ = "end of CoverTab[121522]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:449
		_go_fuzz_dep_.CoverTab[121523]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:449
		// _ = "end of CoverTab[121523]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:449
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:449
	// _ = "end of CoverTab[121519]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:449
	_go_fuzz_dep_.CoverTab[121520]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:453
	prev := p.tok
	p.tok = p.sc.Scan()

	if p.tok.Type == token.COMMENT {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:456
		_go_fuzz_dep_.CoverTab[121524]++
													var comment *ast.CommentGroup
													var endline int

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:462
		if p.tok.Pos.Line == prev.Pos.Line {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:462
			_go_fuzz_dep_.CoverTab[121527]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:465
			comment, endline = p.consumeCommentGroup(0)
			if p.tok.Pos.Line != endline {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:466
				_go_fuzz_dep_.CoverTab[121528]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:469
				p.lineComment = comment
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:469
				// _ = "end of CoverTab[121528]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:470
				_go_fuzz_dep_.CoverTab[121529]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:470
				// _ = "end of CoverTab[121529]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:470
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:470
			// _ = "end of CoverTab[121527]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:471
			_go_fuzz_dep_.CoverTab[121530]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:471
			// _ = "end of CoverTab[121530]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:471
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:471
		// _ = "end of CoverTab[121524]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:471
		_go_fuzz_dep_.CoverTab[121525]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:474
		endline = -1
		for p.tok.Type == token.COMMENT {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:475
			_go_fuzz_dep_.CoverTab[121531]++
														comment, endline = p.consumeCommentGroup(1)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:476
			// _ = "end of CoverTab[121531]"
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:477
		// _ = "end of CoverTab[121525]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:477
		_go_fuzz_dep_.CoverTab[121526]++

													if endline+1 == p.tok.Pos.Line && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:479
			_go_fuzz_dep_.CoverTab[121532]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:479
			return p.tok.Type != token.RBRACE
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:479
			// _ = "end of CoverTab[121532]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:479
		}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:479
			_go_fuzz_dep_.CoverTab[121533]++
														switch p.tok.Type {
			case token.RBRACE, token.RBRACK:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:481
				_go_fuzz_dep_.CoverTab[121534]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:481
				// _ = "end of CoverTab[121534]"

			default:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:483
				_go_fuzz_dep_.CoverTab[121535]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:486
				p.leadComment = comment
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:486
				// _ = "end of CoverTab[121535]"
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:487
			// _ = "end of CoverTab[121533]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:488
			_go_fuzz_dep_.CoverTab[121536]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:488
			// _ = "end of CoverTab[121536]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:488
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:488
		// _ = "end of CoverTab[121526]"

	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:490
		_go_fuzz_dep_.CoverTab[121537]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:490
		// _ = "end of CoverTab[121537]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:490
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:490
	// _ = "end of CoverTab[121520]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:490
	_go_fuzz_dep_.CoverTab[121521]++

												return p.tok
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:492
	// _ = "end of CoverTab[121521]"
}

// unscan pushes the previously read token back onto the buffer.
func (p *Parser) unscan() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:496
	_go_fuzz_dep_.CoverTab[121538]++
												p.n = 1
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:497
	// _ = "end of CoverTab[121538]"
}

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:503
func (p *Parser) printTrace(a ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:503
	_go_fuzz_dep_.CoverTab[121539]++
												if !p.enableTrace {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:504
		_go_fuzz_dep_.CoverTab[121542]++
													return
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:505
		// _ = "end of CoverTab[121542]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:506
		_go_fuzz_dep_.CoverTab[121543]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:506
		// _ = "end of CoverTab[121543]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:506
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:506
	// _ = "end of CoverTab[121539]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:506
	_go_fuzz_dep_.CoverTab[121540]++

												const dots = ". . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . "
												const n = len(dots)
												fmt.Printf("%5d:%3d: ", p.tok.Pos.Line, p.tok.Pos.Column)

												i := 2 * p.indent
												for i > n {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:513
		_go_fuzz_dep_.CoverTab[121544]++
													fmt.Print(dots)
													i -= n
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:515
		// _ = "end of CoverTab[121544]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:516
	// _ = "end of CoverTab[121540]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:516
	_go_fuzz_dep_.CoverTab[121541]++

												fmt.Print(dots[0:i])
												fmt.Println(a...)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:519
	// _ = "end of CoverTab[121541]"
}

func trace(p *Parser, msg string) *Parser {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:522
	_go_fuzz_dep_.CoverTab[121545]++
												p.printTrace(msg, "(")
												p.indent++
												return p
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:525
	// _ = "end of CoverTab[121545]"
}

// Usage pattern: defer un(trace(p, "..."))
func un(p *Parser) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:529
	_go_fuzz_dep_.CoverTab[121546]++
												p.indent--
												p.printTrace(")")
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:531
	// _ = "end of CoverTab[121546]"
}

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:532
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/parser.go:532
var _ = _go_fuzz_dep_.CoverTab
