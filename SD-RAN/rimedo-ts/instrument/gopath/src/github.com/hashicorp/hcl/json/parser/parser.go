//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:1
package parser

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:1
)

import (
	"errors"
	"fmt"

	"github.com/hashicorp/hcl/hcl/ast"
	hcltoken "github.com/hashicorp/hcl/hcl/token"
	"github.com/hashicorp/hcl/json/scanner"
	"github.com/hashicorp/hcl/json/token"
)

type Parser struct {
	sc	*scanner.Scanner

	// Last read token
	tok		token.Token
	commaPrev	token.Token

	enableTrace	bool
	indent		int
	n		int	// buffer size (max = 1)
}

func newParser(src []byte) *Parser {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:25
	_go_fuzz_dep_.CoverTab[121797]++
												return &Parser{
		sc: scanner.New(src),
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:28
	// _ = "end of CoverTab[121797]"
}

// Parse returns the fully parsed source and returns the abstract syntax tree.
func Parse(src []byte) (*ast.File, error) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:32
	_go_fuzz_dep_.CoverTab[121798]++
												p := newParser(src)
												return p.Parse()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:34
	// _ = "end of CoverTab[121798]"
}

var errEofToken = errors.New("EOF token found")

// Parse returns the fully parsed source and returns the abstract syntax tree.
func (p *Parser) Parse() (*ast.File, error) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:40
	_go_fuzz_dep_.CoverTab[121799]++
												f := &ast.File{}
												var err, scerr error
												p.sc.Error = func(pos token.Pos, msg string) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:43
		_go_fuzz_dep_.CoverTab[121803]++
													scerr = fmt.Errorf("%s: %s", pos, msg)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:44
		// _ = "end of CoverTab[121803]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:45
	// _ = "end of CoverTab[121799]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:45
	_go_fuzz_dep_.CoverTab[121800]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:48
	object, err := p.object()
	if scerr != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:49
		_go_fuzz_dep_.CoverTab[121804]++
													return nil, scerr
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:50
		// _ = "end of CoverTab[121804]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:51
		_go_fuzz_dep_.CoverTab[121805]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:51
		// _ = "end of CoverTab[121805]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:51
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:51
	// _ = "end of CoverTab[121800]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:51
	_go_fuzz_dep_.CoverTab[121801]++
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:52
		_go_fuzz_dep_.CoverTab[121806]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:53
		// _ = "end of CoverTab[121806]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:54
		_go_fuzz_dep_.CoverTab[121807]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:54
		// _ = "end of CoverTab[121807]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:54
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:54
	// _ = "end of CoverTab[121801]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:54
	_go_fuzz_dep_.CoverTab[121802]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:57
	f.Node = object.List

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:61
	flattenObjects(f.Node)

												return f, nil
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:63
	// _ = "end of CoverTab[121802]"
}

func (p *Parser) objectList() (*ast.ObjectList, error) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:66
	_go_fuzz_dep_.CoverTab[121808]++
												defer un(trace(p, "ParseObjectList"))
												node := &ast.ObjectList{}

												for {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:70
		_go_fuzz_dep_.CoverTab[121810]++
													n, err := p.objectItem()
													if err == errEofToken {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:72
			_go_fuzz_dep_.CoverTab[121813]++
														break
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:73
			// _ = "end of CoverTab[121813]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:74
			_go_fuzz_dep_.CoverTab[121814]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:74
			// _ = "end of CoverTab[121814]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:74
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:74
		// _ = "end of CoverTab[121810]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:74
		_go_fuzz_dep_.CoverTab[121811]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:78
		if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:78
			_go_fuzz_dep_.CoverTab[121815]++
														return node, err
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:79
			// _ = "end of CoverTab[121815]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:80
			_go_fuzz_dep_.CoverTab[121816]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:80
			// _ = "end of CoverTab[121816]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:80
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:80
		// _ = "end of CoverTab[121811]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:80
		_go_fuzz_dep_.CoverTab[121812]++

													node.Add(n)

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:85
		if tok := p.scan(); tok.Type != token.COMMA {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:85
			_go_fuzz_dep_.CoverTab[121817]++
														break
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:86
			// _ = "end of CoverTab[121817]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:87
			_go_fuzz_dep_.CoverTab[121818]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:87
			// _ = "end of CoverTab[121818]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:87
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:87
		// _ = "end of CoverTab[121812]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:88
	// _ = "end of CoverTab[121808]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:88
	_go_fuzz_dep_.CoverTab[121809]++

												return node, nil
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:90
	// _ = "end of CoverTab[121809]"
}

// objectItem parses a single object item
func (p *Parser) objectItem() (*ast.ObjectItem, error) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:94
	_go_fuzz_dep_.CoverTab[121819]++
												defer un(trace(p, "ParseObjectItem"))

												keys, err := p.objectKey()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:98
		_go_fuzz_dep_.CoverTab[121822]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:99
		// _ = "end of CoverTab[121822]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:100
		_go_fuzz_dep_.CoverTab[121823]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:100
		// _ = "end of CoverTab[121823]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:100
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:100
	// _ = "end of CoverTab[121819]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:100
	_go_fuzz_dep_.CoverTab[121820]++

												o := &ast.ObjectItem{
		Keys: keys,
	}

	switch p.tok.Type {
	case token.COLON:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:107
		_go_fuzz_dep_.CoverTab[121824]++
													pos := p.tok.Pos
													o.Assign = hcltoken.Pos{
			Filename:	pos.Filename,
			Offset:		pos.Offset,
			Line:		pos.Line,
			Column:		pos.Column,
		}

		o.Val, err = p.objectValue()
		if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:117
			_go_fuzz_dep_.CoverTab[121826]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:118
			// _ = "end of CoverTab[121826]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:119
			_go_fuzz_dep_.CoverTab[121827]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:119
			// _ = "end of CoverTab[121827]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:119
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:119
		// _ = "end of CoverTab[121824]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:119
	default:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:119
		_go_fuzz_dep_.CoverTab[121825]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:119
		// _ = "end of CoverTab[121825]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:120
	// _ = "end of CoverTab[121820]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:120
	_go_fuzz_dep_.CoverTab[121821]++

												return o, nil
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:122
	// _ = "end of CoverTab[121821]"
}

// objectKey parses an object key and returns a ObjectKey AST
func (p *Parser) objectKey() ([]*ast.ObjectKey, error) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:126
	_go_fuzz_dep_.CoverTab[121828]++
												keyCount := 0
												keys := make([]*ast.ObjectKey, 0)

												for {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:130
		_go_fuzz_dep_.CoverTab[121829]++
													tok := p.scan()
													switch tok.Type {
		case token.EOF:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:133
			_go_fuzz_dep_.CoverTab[121830]++
														return nil, errEofToken
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:134
			// _ = "end of CoverTab[121830]"
		case token.STRING:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:135
			_go_fuzz_dep_.CoverTab[121831]++
														keyCount++
														keys = append(keys, &ast.ObjectKey{
				Token: p.tok.HCLToken(),
			})
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:139
			// _ = "end of CoverTab[121831]"
		case token.COLON:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:140
			_go_fuzz_dep_.CoverTab[121832]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:143
			if keyCount == 0 {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:143
				_go_fuzz_dep_.CoverTab[121836]++
															return nil, fmt.Errorf("expected: STRING got: %s", p.tok.Type)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:144
				// _ = "end of CoverTab[121836]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:145
				_go_fuzz_dep_.CoverTab[121837]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:145
				// _ = "end of CoverTab[121837]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:145
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:145
			// _ = "end of CoverTab[121832]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:145
			_go_fuzz_dep_.CoverTab[121833]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:148
			return keys, nil
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:148
			// _ = "end of CoverTab[121833]"
		case token.ILLEGAL:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:149
			_go_fuzz_dep_.CoverTab[121834]++
														return nil, errors.New("illegal")
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:150
			// _ = "end of CoverTab[121834]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:151
			_go_fuzz_dep_.CoverTab[121835]++
														return nil, fmt.Errorf("expected: STRING got: %s", p.tok.Type)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:152
			// _ = "end of CoverTab[121835]"
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:153
		// _ = "end of CoverTab[121829]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:154
	// _ = "end of CoverTab[121828]"
}

// object parses any type of object, such as number, bool, string, object or
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:157
// list.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:159
func (p *Parser) objectValue() (ast.Node, error) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:159
	_go_fuzz_dep_.CoverTab[121838]++
												defer un(trace(p, "ParseObjectValue"))
												tok := p.scan()

												switch tok.Type {
	case token.NUMBER, token.FLOAT, token.BOOL, token.NULL, token.STRING:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:164
		_go_fuzz_dep_.CoverTab[121840]++
													return p.literalType()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:165
		// _ = "end of CoverTab[121840]"
	case token.LBRACE:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:166
		_go_fuzz_dep_.CoverTab[121841]++
													return p.objectType()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:167
		// _ = "end of CoverTab[121841]"
	case token.LBRACK:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:168
		_go_fuzz_dep_.CoverTab[121842]++
													return p.listType()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:169
		// _ = "end of CoverTab[121842]"
	case token.EOF:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:170
		_go_fuzz_dep_.CoverTab[121843]++
													return nil, errEofToken
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:171
		// _ = "end of CoverTab[121843]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:171
	default:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:171
		_go_fuzz_dep_.CoverTab[121844]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:171
		// _ = "end of CoverTab[121844]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:172
	// _ = "end of CoverTab[121838]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:172
	_go_fuzz_dep_.CoverTab[121839]++

												return nil, fmt.Errorf("Expected object value, got unknown token: %+v", tok)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:174
	// _ = "end of CoverTab[121839]"
}

// object parses any type of object, such as number, bool, string, object or
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:177
// list.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:179
func (p *Parser) object() (*ast.ObjectType, error) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:179
	_go_fuzz_dep_.CoverTab[121845]++
												defer un(trace(p, "ParseType"))
												tok := p.scan()

												switch tok.Type {
	case token.LBRACE:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:184
		_go_fuzz_dep_.CoverTab[121847]++
													return p.objectType()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:185
		// _ = "end of CoverTab[121847]"
	case token.EOF:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:186
		_go_fuzz_dep_.CoverTab[121848]++
													return nil, errEofToken
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:187
		// _ = "end of CoverTab[121848]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:187
	default:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:187
		_go_fuzz_dep_.CoverTab[121849]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:187
		// _ = "end of CoverTab[121849]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:188
	// _ = "end of CoverTab[121845]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:188
	_go_fuzz_dep_.CoverTab[121846]++

												return nil, fmt.Errorf("Expected object, got unknown token: %+v", tok)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:190
	// _ = "end of CoverTab[121846]"
}

// objectType parses an object type and returns a ObjectType AST
func (p *Parser) objectType() (*ast.ObjectType, error) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:194
	_go_fuzz_dep_.CoverTab[121850]++
												defer un(trace(p, "ParseObjectType"))

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:198
	o := &ast.ObjectType{}

												l, err := p.objectList()

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:204
	if err != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:204
		_go_fuzz_dep_.CoverTab[121852]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:204
		return p.tok.Type != token.RBRACE
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:204
		// _ = "end of CoverTab[121852]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:204
	}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:204
		_go_fuzz_dep_.CoverTab[121853]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:205
		// _ = "end of CoverTab[121853]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:206
		_go_fuzz_dep_.CoverTab[121854]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:206
		// _ = "end of CoverTab[121854]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:206
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:206
	// _ = "end of CoverTab[121850]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:206
	_go_fuzz_dep_.CoverTab[121851]++

												o.List = l
												return o, nil
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:209
	// _ = "end of CoverTab[121851]"
}

// listType parses a list type and returns a ListType AST
func (p *Parser) listType() (*ast.ListType, error) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:213
	_go_fuzz_dep_.CoverTab[121855]++
												defer un(trace(p, "ParseListType"))

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:217
	l := &ast.ListType{}

	for {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:219
		_go_fuzz_dep_.CoverTab[121856]++
													tok := p.scan()
													switch tok.Type {
		case token.NUMBER, token.FLOAT, token.STRING:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:222
			_go_fuzz_dep_.CoverTab[121857]++
														node, err := p.literalType()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:224
				_go_fuzz_dep_.CoverTab[121866]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:225
				// _ = "end of CoverTab[121866]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:226
				_go_fuzz_dep_.CoverTab[121867]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:226
				// _ = "end of CoverTab[121867]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:226
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:226
			// _ = "end of CoverTab[121857]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:226
			_go_fuzz_dep_.CoverTab[121858]++

														l.Add(node)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:228
			// _ = "end of CoverTab[121858]"
		case token.COMMA:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:229
			_go_fuzz_dep_.CoverTab[121859]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:230
			// _ = "end of CoverTab[121859]"
		case token.LBRACE:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:231
			_go_fuzz_dep_.CoverTab[121860]++
														node, err := p.objectType()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:233
				_go_fuzz_dep_.CoverTab[121868]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:234
				// _ = "end of CoverTab[121868]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:235
				_go_fuzz_dep_.CoverTab[121869]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:235
				// _ = "end of CoverTab[121869]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:235
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:235
			// _ = "end of CoverTab[121860]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:235
			_go_fuzz_dep_.CoverTab[121861]++

														l.Add(node)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:237
			// _ = "end of CoverTab[121861]"
		case token.BOOL:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:238
			_go_fuzz_dep_.CoverTab[121862]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:238
			// _ = "end of CoverTab[121862]"

		case token.LBRACK:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:240
			_go_fuzz_dep_.CoverTab[121863]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:240
			// _ = "end of CoverTab[121863]"

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:244
		case token.RBRACK:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:244
			_go_fuzz_dep_.CoverTab[121864]++

														return l, nil
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:246
			// _ = "end of CoverTab[121864]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:247
			_go_fuzz_dep_.CoverTab[121865]++
														return nil, fmt.Errorf("unexpected token while parsing list: %s", tok.Type)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:248
			// _ = "end of CoverTab[121865]"
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:249
		// _ = "end of CoverTab[121856]"

	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:251
	// _ = "end of CoverTab[121855]"
}

// literalType parses a literal type and returns a LiteralType AST
func (p *Parser) literalType() (*ast.LiteralType, error) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:255
	_go_fuzz_dep_.CoverTab[121870]++
												defer un(trace(p, "ParseLiteral"))

												return &ast.LiteralType{
		Token: p.tok.HCLToken(),
	}, nil
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:260
	// _ = "end of CoverTab[121870]"
}

// scan returns the next token from the underlying scanner. If a token has
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:263
// been unscanned then read that instead.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:265
func (p *Parser) scan() token.Token {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:265
	_go_fuzz_dep_.CoverTab[121871]++

												if p.n != 0 {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:267
		_go_fuzz_dep_.CoverTab[121873]++
													p.n = 0
													return p.tok
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:269
		// _ = "end of CoverTab[121873]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:270
		_go_fuzz_dep_.CoverTab[121874]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:270
		// _ = "end of CoverTab[121874]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:270
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:270
	// _ = "end of CoverTab[121871]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:270
	_go_fuzz_dep_.CoverTab[121872]++

												p.tok = p.sc.Scan()
												return p.tok
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:273
	// _ = "end of CoverTab[121872]"
}

// unscan pushes the previously read token back onto the buffer.
func (p *Parser) unscan() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:277
	_go_fuzz_dep_.CoverTab[121875]++
												p.n = 1
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:278
	// _ = "end of CoverTab[121875]"
}

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:284
func (p *Parser) printTrace(a ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:284
	_go_fuzz_dep_.CoverTab[121876]++
												if !p.enableTrace {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:285
		_go_fuzz_dep_.CoverTab[121879]++
													return
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:286
		// _ = "end of CoverTab[121879]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:287
		_go_fuzz_dep_.CoverTab[121880]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:287
		// _ = "end of CoverTab[121880]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:287
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:287
	// _ = "end of CoverTab[121876]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:287
	_go_fuzz_dep_.CoverTab[121877]++

												const dots = ". . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . "
												const n = len(dots)
												fmt.Printf("%5d:%3d: ", p.tok.Pos.Line, p.tok.Pos.Column)

												i := 2 * p.indent
												for i > n {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:294
		_go_fuzz_dep_.CoverTab[121881]++
													fmt.Print(dots)
													i -= n
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:296
		// _ = "end of CoverTab[121881]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:297
	// _ = "end of CoverTab[121877]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:297
	_go_fuzz_dep_.CoverTab[121878]++

												fmt.Print(dots[0:i])
												fmt.Println(a...)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:300
	// _ = "end of CoverTab[121878]"
}

func trace(p *Parser, msg string) *Parser {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:303
	_go_fuzz_dep_.CoverTab[121882]++
												p.printTrace(msg, "(")
												p.indent++
												return p
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:306
	// _ = "end of CoverTab[121882]"
}

// Usage pattern: defer un(trace(p, "..."))
func un(p *Parser) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:310
	_go_fuzz_dep_.CoverTab[121883]++
												p.indent--
												p.printTrace(")")
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:312
	// _ = "end of CoverTab[121883]"
}

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:313
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/parser.go:313
var _ = _go_fuzz_dep_.CoverTab
