// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/text/template/parse/parse.go:5
// Package parse builds parse trees for templates as defined by text/template
//line /usr/local/go/src/text/template/parse/parse.go:5
// and html/template. Clients should use those packages to construct templates
//line /usr/local/go/src/text/template/parse/parse.go:5
// rather than this one, which provides shared internal data structures not
//line /usr/local/go/src/text/template/parse/parse.go:5
// intended for general use.
//line /usr/local/go/src/text/template/parse/parse.go:9
package parse

//line /usr/local/go/src/text/template/parse/parse.go:9
import (
//line /usr/local/go/src/text/template/parse/parse.go:9
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/text/template/parse/parse.go:9
)
//line /usr/local/go/src/text/template/parse/parse.go:9
import (
//line /usr/local/go/src/text/template/parse/parse.go:9
	_atomic_ "sync/atomic"
//line /usr/local/go/src/text/template/parse/parse.go:9
)

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

// Tree is the representation of a single parsed template.
type Tree struct {
	Name		string		// name of the template represented by the tree.
	ParseName	string		// name of the top-level template during parsing, for error messages.
	Root		*ListNode	// top-level root of the tree.
	Mode		Mode		// parsing mode.
	text		string		// text parsed to create the template (or its parent)
	// Parsing only; cleared after parse.
	funcs		[]map[string]any
	lex		*lexer
	token		[3]item	// three-token lookahead for parser.
	peekCount	int
	vars		[]string	// variables defined at the moment.
	treeSet		map[string]*Tree
	actionLine	int	// line of left delim starting action
	rangeDepth	int
}

// A mode value is a set of flags (or 0). Modes control parser behavior.
type Mode uint

const (
	ParseComments	Mode	= 1 << iota	// parse comments and add them to AST
	SkipFuncCheck				// do not check that functions are defined
)

// Copy returns a copy of the Tree. Any parsing state is discarded.
func (t *Tree) Copy() *Tree {
//line /usr/local/go/src/text/template/parse/parse.go:46
	_go_fuzz_dep_.CoverTab[29438]++
								if t == nil {
//line /usr/local/go/src/text/template/parse/parse.go:47
		_go_fuzz_dep_.CoverTab[29440]++
									return nil
//line /usr/local/go/src/text/template/parse/parse.go:48
		// _ = "end of CoverTab[29440]"
	} else {
//line /usr/local/go/src/text/template/parse/parse.go:49
		_go_fuzz_dep_.CoverTab[29441]++
//line /usr/local/go/src/text/template/parse/parse.go:49
		// _ = "end of CoverTab[29441]"
//line /usr/local/go/src/text/template/parse/parse.go:49
	}
//line /usr/local/go/src/text/template/parse/parse.go:49
	// _ = "end of CoverTab[29438]"
//line /usr/local/go/src/text/template/parse/parse.go:49
	_go_fuzz_dep_.CoverTab[29439]++
								return &Tree{
		Name:		t.Name,
		ParseName:	t.ParseName,
		Root:		t.Root.CopyList(),
		text:		t.text,
	}
//line /usr/local/go/src/text/template/parse/parse.go:55
	// _ = "end of CoverTab[29439]"
}

// Parse returns a map from template name to parse.Tree, created by parsing the
//line /usr/local/go/src/text/template/parse/parse.go:58
// templates described in the argument string. The top-level template will be
//line /usr/local/go/src/text/template/parse/parse.go:58
// given the specified name. If an error is encountered, parsing stops and an
//line /usr/local/go/src/text/template/parse/parse.go:58
// empty map is returned with the error.
//line /usr/local/go/src/text/template/parse/parse.go:62
func Parse(name, text, leftDelim, rightDelim string, funcs ...map[string]any) (map[string]*Tree, error) {
//line /usr/local/go/src/text/template/parse/parse.go:62
	_go_fuzz_dep_.CoverTab[29442]++
								treeSet := make(map[string]*Tree)
								t := New(name)
								t.text = text
								_, err := t.Parse(text, leftDelim, rightDelim, treeSet, funcs...)
								return treeSet, err
//line /usr/local/go/src/text/template/parse/parse.go:67
	// _ = "end of CoverTab[29442]"
}

// next returns the next token.
func (t *Tree) next() item {
//line /usr/local/go/src/text/template/parse/parse.go:71
	_go_fuzz_dep_.CoverTab[29443]++
								if t.peekCount > 0 {
//line /usr/local/go/src/text/template/parse/parse.go:72
		_go_fuzz_dep_.CoverTab[29445]++
									t.peekCount--
//line /usr/local/go/src/text/template/parse/parse.go:73
		// _ = "end of CoverTab[29445]"
	} else {
//line /usr/local/go/src/text/template/parse/parse.go:74
		_go_fuzz_dep_.CoverTab[29446]++
									t.token[0] = t.lex.nextItem()
//line /usr/local/go/src/text/template/parse/parse.go:75
		// _ = "end of CoverTab[29446]"
	}
//line /usr/local/go/src/text/template/parse/parse.go:76
	// _ = "end of CoverTab[29443]"
//line /usr/local/go/src/text/template/parse/parse.go:76
	_go_fuzz_dep_.CoverTab[29444]++
								return t.token[t.peekCount]
//line /usr/local/go/src/text/template/parse/parse.go:77
	// _ = "end of CoverTab[29444]"
}

// backup backs the input stream up one token.
func (t *Tree) backup() {
//line /usr/local/go/src/text/template/parse/parse.go:81
	_go_fuzz_dep_.CoverTab[29447]++
								t.peekCount++
//line /usr/local/go/src/text/template/parse/parse.go:82
	// _ = "end of CoverTab[29447]"
}

// backup2 backs the input stream up two tokens.
//line /usr/local/go/src/text/template/parse/parse.go:85
// The zeroth token is already there.
//line /usr/local/go/src/text/template/parse/parse.go:87
func (t *Tree) backup2(t1 item) {
//line /usr/local/go/src/text/template/parse/parse.go:87
	_go_fuzz_dep_.CoverTab[29448]++
								t.token[1] = t1
								t.peekCount = 2
//line /usr/local/go/src/text/template/parse/parse.go:89
	// _ = "end of CoverTab[29448]"
}

// backup3 backs the input stream up three tokens
//line /usr/local/go/src/text/template/parse/parse.go:92
// The zeroth token is already there.
//line /usr/local/go/src/text/template/parse/parse.go:94
func (t *Tree) backup3(t2, t1 item) {
//line /usr/local/go/src/text/template/parse/parse.go:94
	_go_fuzz_dep_.CoverTab[29449]++
								t.token[1] = t1
								t.token[2] = t2
								t.peekCount = 3
//line /usr/local/go/src/text/template/parse/parse.go:97
	// _ = "end of CoverTab[29449]"
}

// peek returns but does not consume the next token.
func (t *Tree) peek() item {
//line /usr/local/go/src/text/template/parse/parse.go:101
	_go_fuzz_dep_.CoverTab[29450]++
								if t.peekCount > 0 {
//line /usr/local/go/src/text/template/parse/parse.go:102
		_go_fuzz_dep_.CoverTab[29452]++
									return t.token[t.peekCount-1]
//line /usr/local/go/src/text/template/parse/parse.go:103
		// _ = "end of CoverTab[29452]"
	} else {
//line /usr/local/go/src/text/template/parse/parse.go:104
		_go_fuzz_dep_.CoverTab[29453]++
//line /usr/local/go/src/text/template/parse/parse.go:104
		// _ = "end of CoverTab[29453]"
//line /usr/local/go/src/text/template/parse/parse.go:104
	}
//line /usr/local/go/src/text/template/parse/parse.go:104
	// _ = "end of CoverTab[29450]"
//line /usr/local/go/src/text/template/parse/parse.go:104
	_go_fuzz_dep_.CoverTab[29451]++
								t.peekCount = 1
								t.token[0] = t.lex.nextItem()
								return t.token[0]
//line /usr/local/go/src/text/template/parse/parse.go:107
	// _ = "end of CoverTab[29451]"
}

// nextNonSpace returns the next non-space token.
func (t *Tree) nextNonSpace() (token item) {
//line /usr/local/go/src/text/template/parse/parse.go:111
	_go_fuzz_dep_.CoverTab[29454]++
								for {
//line /usr/local/go/src/text/template/parse/parse.go:112
		_go_fuzz_dep_.CoverTab[29456]++
									token = t.next()
									if token.typ != itemSpace {
//line /usr/local/go/src/text/template/parse/parse.go:114
			_go_fuzz_dep_.CoverTab[29457]++
										break
//line /usr/local/go/src/text/template/parse/parse.go:115
			// _ = "end of CoverTab[29457]"
		} else {
//line /usr/local/go/src/text/template/parse/parse.go:116
			_go_fuzz_dep_.CoverTab[29458]++
//line /usr/local/go/src/text/template/parse/parse.go:116
			// _ = "end of CoverTab[29458]"
//line /usr/local/go/src/text/template/parse/parse.go:116
		}
//line /usr/local/go/src/text/template/parse/parse.go:116
		// _ = "end of CoverTab[29456]"
	}
//line /usr/local/go/src/text/template/parse/parse.go:117
	// _ = "end of CoverTab[29454]"
//line /usr/local/go/src/text/template/parse/parse.go:117
	_go_fuzz_dep_.CoverTab[29455]++
								return token
//line /usr/local/go/src/text/template/parse/parse.go:118
	// _ = "end of CoverTab[29455]"
}

// peekNonSpace returns but does not consume the next non-space token.
func (t *Tree) peekNonSpace() item {
//line /usr/local/go/src/text/template/parse/parse.go:122
	_go_fuzz_dep_.CoverTab[29459]++
								token := t.nextNonSpace()
								t.backup()
								return token
//line /usr/local/go/src/text/template/parse/parse.go:125
	// _ = "end of CoverTab[29459]"
}

//line /usr/local/go/src/text/template/parse/parse.go:130
// New allocates a new parse tree with the given name.
func New(name string, funcs ...map[string]any) *Tree {
//line /usr/local/go/src/text/template/parse/parse.go:131
	_go_fuzz_dep_.CoverTab[29460]++
								return &Tree{
		Name:	name,
		funcs:	funcs,
	}
//line /usr/local/go/src/text/template/parse/parse.go:135
	// _ = "end of CoverTab[29460]"
}

// ErrorContext returns a textual representation of the location of the node in the input text.
//line /usr/local/go/src/text/template/parse/parse.go:138
// The receiver is only used when the node does not have a pointer to the tree inside,
//line /usr/local/go/src/text/template/parse/parse.go:138
// which can occur in old code.
//line /usr/local/go/src/text/template/parse/parse.go:141
func (t *Tree) ErrorContext(n Node) (location, context string) {
//line /usr/local/go/src/text/template/parse/parse.go:141
	_go_fuzz_dep_.CoverTab[29461]++
								pos := int(n.Position())
								tree := n.tree()
								if tree == nil {
//line /usr/local/go/src/text/template/parse/parse.go:144
		_go_fuzz_dep_.CoverTab[29464]++
									tree = t
//line /usr/local/go/src/text/template/parse/parse.go:145
		// _ = "end of CoverTab[29464]"
	} else {
//line /usr/local/go/src/text/template/parse/parse.go:146
		_go_fuzz_dep_.CoverTab[29465]++
//line /usr/local/go/src/text/template/parse/parse.go:146
		// _ = "end of CoverTab[29465]"
//line /usr/local/go/src/text/template/parse/parse.go:146
	}
//line /usr/local/go/src/text/template/parse/parse.go:146
	// _ = "end of CoverTab[29461]"
//line /usr/local/go/src/text/template/parse/parse.go:146
	_go_fuzz_dep_.CoverTab[29462]++
								text := tree.text[:pos]
								byteNum := strings.LastIndex(text, "\n")
								if byteNum == -1 {
//line /usr/local/go/src/text/template/parse/parse.go:149
		_go_fuzz_dep_.CoverTab[29466]++
									byteNum = pos
//line /usr/local/go/src/text/template/parse/parse.go:150
		// _ = "end of CoverTab[29466]"
	} else {
//line /usr/local/go/src/text/template/parse/parse.go:151
		_go_fuzz_dep_.CoverTab[29467]++
									byteNum++
									byteNum = pos - byteNum
//line /usr/local/go/src/text/template/parse/parse.go:153
		// _ = "end of CoverTab[29467]"
	}
//line /usr/local/go/src/text/template/parse/parse.go:154
	// _ = "end of CoverTab[29462]"
//line /usr/local/go/src/text/template/parse/parse.go:154
	_go_fuzz_dep_.CoverTab[29463]++
								lineNum := 1 + strings.Count(text, "\n")
								context = n.String()
								return fmt.Sprintf("%s:%d:%d", tree.ParseName, lineNum, byteNum), context
//line /usr/local/go/src/text/template/parse/parse.go:157
	// _ = "end of CoverTab[29463]"
}

// errorf formats the error and terminates processing.
func (t *Tree) errorf(format string, args ...any) {
//line /usr/local/go/src/text/template/parse/parse.go:161
	_go_fuzz_dep_.CoverTab[29468]++
								t.Root = nil
								format = fmt.Sprintf("template: %s:%d: %s", t.ParseName, t.token[0].line, format)
								panic(fmt.Errorf(format, args...))
//line /usr/local/go/src/text/template/parse/parse.go:164
	// _ = "end of CoverTab[29468]"
}

// error terminates processing.
func (t *Tree) error(err error) {
//line /usr/local/go/src/text/template/parse/parse.go:168
	_go_fuzz_dep_.CoverTab[29469]++
								t.errorf("%s", err)
//line /usr/local/go/src/text/template/parse/parse.go:169
	// _ = "end of CoverTab[29469]"
}

// expect consumes the next token and guarantees it has the required type.
func (t *Tree) expect(expected itemType, context string) item {
//line /usr/local/go/src/text/template/parse/parse.go:173
	_go_fuzz_dep_.CoverTab[29470]++
								token := t.nextNonSpace()
								if token.typ != expected {
//line /usr/local/go/src/text/template/parse/parse.go:175
		_go_fuzz_dep_.CoverTab[29472]++
									t.unexpected(token, context)
//line /usr/local/go/src/text/template/parse/parse.go:176
		// _ = "end of CoverTab[29472]"
	} else {
//line /usr/local/go/src/text/template/parse/parse.go:177
		_go_fuzz_dep_.CoverTab[29473]++
//line /usr/local/go/src/text/template/parse/parse.go:177
		// _ = "end of CoverTab[29473]"
//line /usr/local/go/src/text/template/parse/parse.go:177
	}
//line /usr/local/go/src/text/template/parse/parse.go:177
	// _ = "end of CoverTab[29470]"
//line /usr/local/go/src/text/template/parse/parse.go:177
	_go_fuzz_dep_.CoverTab[29471]++
								return token
//line /usr/local/go/src/text/template/parse/parse.go:178
	// _ = "end of CoverTab[29471]"
}

// expectOneOf consumes the next token and guarantees it has one of the required types.
func (t *Tree) expectOneOf(expected1, expected2 itemType, context string) item {
//line /usr/local/go/src/text/template/parse/parse.go:182
	_go_fuzz_dep_.CoverTab[29474]++
								token := t.nextNonSpace()
								if token.typ != expected1 && func() bool {
//line /usr/local/go/src/text/template/parse/parse.go:184
		_go_fuzz_dep_.CoverTab[29476]++
//line /usr/local/go/src/text/template/parse/parse.go:184
		return token.typ != expected2
//line /usr/local/go/src/text/template/parse/parse.go:184
		// _ = "end of CoverTab[29476]"
//line /usr/local/go/src/text/template/parse/parse.go:184
	}() {
//line /usr/local/go/src/text/template/parse/parse.go:184
		_go_fuzz_dep_.CoverTab[29477]++
									t.unexpected(token, context)
//line /usr/local/go/src/text/template/parse/parse.go:185
		// _ = "end of CoverTab[29477]"
	} else {
//line /usr/local/go/src/text/template/parse/parse.go:186
		_go_fuzz_dep_.CoverTab[29478]++
//line /usr/local/go/src/text/template/parse/parse.go:186
		// _ = "end of CoverTab[29478]"
//line /usr/local/go/src/text/template/parse/parse.go:186
	}
//line /usr/local/go/src/text/template/parse/parse.go:186
	// _ = "end of CoverTab[29474]"
//line /usr/local/go/src/text/template/parse/parse.go:186
	_go_fuzz_dep_.CoverTab[29475]++
								return token
//line /usr/local/go/src/text/template/parse/parse.go:187
	// _ = "end of CoverTab[29475]"
}

// unexpected complains about the token and terminates processing.
func (t *Tree) unexpected(token item, context string) {
//line /usr/local/go/src/text/template/parse/parse.go:191
	_go_fuzz_dep_.CoverTab[29479]++
								if token.typ == itemError {
//line /usr/local/go/src/text/template/parse/parse.go:192
		_go_fuzz_dep_.CoverTab[29481]++
									extra := ""
									if t.actionLine != 0 && func() bool {
//line /usr/local/go/src/text/template/parse/parse.go:194
			_go_fuzz_dep_.CoverTab[29483]++
//line /usr/local/go/src/text/template/parse/parse.go:194
			return t.actionLine != token.line
//line /usr/local/go/src/text/template/parse/parse.go:194
			// _ = "end of CoverTab[29483]"
//line /usr/local/go/src/text/template/parse/parse.go:194
		}() {
//line /usr/local/go/src/text/template/parse/parse.go:194
			_go_fuzz_dep_.CoverTab[29484]++
										extra = fmt.Sprintf(" in action started at %s:%d", t.ParseName, t.actionLine)
										if strings.HasSuffix(token.val, " action") {
//line /usr/local/go/src/text/template/parse/parse.go:196
				_go_fuzz_dep_.CoverTab[29485]++
											extra = extra[len(" in action"):]
//line /usr/local/go/src/text/template/parse/parse.go:197
				// _ = "end of CoverTab[29485]"
			} else {
//line /usr/local/go/src/text/template/parse/parse.go:198
				_go_fuzz_dep_.CoverTab[29486]++
//line /usr/local/go/src/text/template/parse/parse.go:198
				// _ = "end of CoverTab[29486]"
//line /usr/local/go/src/text/template/parse/parse.go:198
			}
//line /usr/local/go/src/text/template/parse/parse.go:198
			// _ = "end of CoverTab[29484]"
		} else {
//line /usr/local/go/src/text/template/parse/parse.go:199
			_go_fuzz_dep_.CoverTab[29487]++
//line /usr/local/go/src/text/template/parse/parse.go:199
			// _ = "end of CoverTab[29487]"
//line /usr/local/go/src/text/template/parse/parse.go:199
		}
//line /usr/local/go/src/text/template/parse/parse.go:199
		// _ = "end of CoverTab[29481]"
//line /usr/local/go/src/text/template/parse/parse.go:199
		_go_fuzz_dep_.CoverTab[29482]++
									t.errorf("%s%s", token, extra)
//line /usr/local/go/src/text/template/parse/parse.go:200
		// _ = "end of CoverTab[29482]"
	} else {
//line /usr/local/go/src/text/template/parse/parse.go:201
		_go_fuzz_dep_.CoverTab[29488]++
//line /usr/local/go/src/text/template/parse/parse.go:201
		// _ = "end of CoverTab[29488]"
//line /usr/local/go/src/text/template/parse/parse.go:201
	}
//line /usr/local/go/src/text/template/parse/parse.go:201
	// _ = "end of CoverTab[29479]"
//line /usr/local/go/src/text/template/parse/parse.go:201
	_go_fuzz_dep_.CoverTab[29480]++
								t.errorf("unexpected %s in %s", token, context)
//line /usr/local/go/src/text/template/parse/parse.go:202
	// _ = "end of CoverTab[29480]"
}

// recover is the handler that turns panics into returns from the top level of Parse.
func (t *Tree) recover(errp *error) {
//line /usr/local/go/src/text/template/parse/parse.go:206
	_go_fuzz_dep_.CoverTab[29489]++
								e := recover()
								if e != nil {
//line /usr/local/go/src/text/template/parse/parse.go:208
		_go_fuzz_dep_.CoverTab[29490]++
									if _, ok := e.(runtime.Error); ok {
//line /usr/local/go/src/text/template/parse/parse.go:209
			_go_fuzz_dep_.CoverTab[29493]++
										panic(e)
//line /usr/local/go/src/text/template/parse/parse.go:210
			// _ = "end of CoverTab[29493]"
		} else {
//line /usr/local/go/src/text/template/parse/parse.go:211
			_go_fuzz_dep_.CoverTab[29494]++
//line /usr/local/go/src/text/template/parse/parse.go:211
			// _ = "end of CoverTab[29494]"
//line /usr/local/go/src/text/template/parse/parse.go:211
		}
//line /usr/local/go/src/text/template/parse/parse.go:211
		// _ = "end of CoverTab[29490]"
//line /usr/local/go/src/text/template/parse/parse.go:211
		_go_fuzz_dep_.CoverTab[29491]++
									if t != nil {
//line /usr/local/go/src/text/template/parse/parse.go:212
			_go_fuzz_dep_.CoverTab[29495]++
										t.stopParse()
//line /usr/local/go/src/text/template/parse/parse.go:213
			// _ = "end of CoverTab[29495]"
		} else {
//line /usr/local/go/src/text/template/parse/parse.go:214
			_go_fuzz_dep_.CoverTab[29496]++
//line /usr/local/go/src/text/template/parse/parse.go:214
			// _ = "end of CoverTab[29496]"
//line /usr/local/go/src/text/template/parse/parse.go:214
		}
//line /usr/local/go/src/text/template/parse/parse.go:214
		// _ = "end of CoverTab[29491]"
//line /usr/local/go/src/text/template/parse/parse.go:214
		_go_fuzz_dep_.CoverTab[29492]++
									*errp = e.(error)
//line /usr/local/go/src/text/template/parse/parse.go:215
		// _ = "end of CoverTab[29492]"
	} else {
//line /usr/local/go/src/text/template/parse/parse.go:216
		_go_fuzz_dep_.CoverTab[29497]++
//line /usr/local/go/src/text/template/parse/parse.go:216
		// _ = "end of CoverTab[29497]"
//line /usr/local/go/src/text/template/parse/parse.go:216
	}
//line /usr/local/go/src/text/template/parse/parse.go:216
	// _ = "end of CoverTab[29489]"
}

// startParse initializes the parser, using the lexer.
func (t *Tree) startParse(funcs []map[string]any, lex *lexer, treeSet map[string]*Tree) {
//line /usr/local/go/src/text/template/parse/parse.go:220
	_go_fuzz_dep_.CoverTab[29498]++
								t.Root = nil
								t.lex = lex
								t.vars = []string{"$"}
								t.funcs = funcs
								t.treeSet = treeSet
								lex.options = lexOptions{
		emitComment:	t.Mode&ParseComments != 0,
		breakOK:	!t.hasFunction("break"),
		continueOK:	!t.hasFunction("continue"),
	}
//line /usr/local/go/src/text/template/parse/parse.go:230
	// _ = "end of CoverTab[29498]"
}

// stopParse terminates parsing.
func (t *Tree) stopParse() {
//line /usr/local/go/src/text/template/parse/parse.go:234
	_go_fuzz_dep_.CoverTab[29499]++
								t.lex = nil
								t.vars = nil
								t.funcs = nil
								t.treeSet = nil
//line /usr/local/go/src/text/template/parse/parse.go:238
	// _ = "end of CoverTab[29499]"
}

// Parse parses the template definition string to construct a representation of
//line /usr/local/go/src/text/template/parse/parse.go:241
// the template for execution. If either action delimiter string is empty, the
//line /usr/local/go/src/text/template/parse/parse.go:241
// default ("{{" or "}}") is used. Embedded template definitions are added to
//line /usr/local/go/src/text/template/parse/parse.go:241
// the treeSet map.
//line /usr/local/go/src/text/template/parse/parse.go:245
func (t *Tree) Parse(text, leftDelim, rightDelim string, treeSet map[string]*Tree, funcs ...map[string]any) (tree *Tree, err error) {
//line /usr/local/go/src/text/template/parse/parse.go:245
	_go_fuzz_dep_.CoverTab[29500]++
								defer t.recover(&err)
								t.ParseName = t.Name
								lexer := lex(t.Name, text, leftDelim, rightDelim)
								t.startParse(funcs, lexer, treeSet)
								t.text = text
								t.parse()
								t.add()
								t.stopParse()
								return t, nil
//line /usr/local/go/src/text/template/parse/parse.go:254
	// _ = "end of CoverTab[29500]"
}

// add adds tree to t.treeSet.
func (t *Tree) add() {
//line /usr/local/go/src/text/template/parse/parse.go:258
	_go_fuzz_dep_.CoverTab[29501]++
								tree := t.treeSet[t.Name]
								if tree == nil || func() bool {
//line /usr/local/go/src/text/template/parse/parse.go:260
		_go_fuzz_dep_.CoverTab[29503]++
//line /usr/local/go/src/text/template/parse/parse.go:260
		return IsEmptyTree(tree.Root)
//line /usr/local/go/src/text/template/parse/parse.go:260
		// _ = "end of CoverTab[29503]"
//line /usr/local/go/src/text/template/parse/parse.go:260
	}() {
//line /usr/local/go/src/text/template/parse/parse.go:260
		_go_fuzz_dep_.CoverTab[29504]++
									t.treeSet[t.Name] = t
									return
//line /usr/local/go/src/text/template/parse/parse.go:262
		// _ = "end of CoverTab[29504]"
	} else {
//line /usr/local/go/src/text/template/parse/parse.go:263
		_go_fuzz_dep_.CoverTab[29505]++
//line /usr/local/go/src/text/template/parse/parse.go:263
		// _ = "end of CoverTab[29505]"
//line /usr/local/go/src/text/template/parse/parse.go:263
	}
//line /usr/local/go/src/text/template/parse/parse.go:263
	// _ = "end of CoverTab[29501]"
//line /usr/local/go/src/text/template/parse/parse.go:263
	_go_fuzz_dep_.CoverTab[29502]++
								if !IsEmptyTree(t.Root) {
//line /usr/local/go/src/text/template/parse/parse.go:264
		_go_fuzz_dep_.CoverTab[29506]++
									t.errorf("template: multiple definition of template %q", t.Name)
//line /usr/local/go/src/text/template/parse/parse.go:265
		// _ = "end of CoverTab[29506]"
	} else {
//line /usr/local/go/src/text/template/parse/parse.go:266
		_go_fuzz_dep_.CoverTab[29507]++
//line /usr/local/go/src/text/template/parse/parse.go:266
		// _ = "end of CoverTab[29507]"
//line /usr/local/go/src/text/template/parse/parse.go:266
	}
//line /usr/local/go/src/text/template/parse/parse.go:266
	// _ = "end of CoverTab[29502]"
}

// IsEmptyTree reports whether this tree (node) is empty of everything but space or comments.
func IsEmptyTree(n Node) bool {
//line /usr/local/go/src/text/template/parse/parse.go:270
	_go_fuzz_dep_.CoverTab[29508]++
								switch n := n.(type) {
	case nil:
//line /usr/local/go/src/text/template/parse/parse.go:272
		_go_fuzz_dep_.CoverTab[29510]++
									return true
//line /usr/local/go/src/text/template/parse/parse.go:273
		// _ = "end of CoverTab[29510]"
	case *ActionNode:
//line /usr/local/go/src/text/template/parse/parse.go:274
		_go_fuzz_dep_.CoverTab[29511]++
//line /usr/local/go/src/text/template/parse/parse.go:274
		// _ = "end of CoverTab[29511]"
	case *CommentNode:
//line /usr/local/go/src/text/template/parse/parse.go:275
		_go_fuzz_dep_.CoverTab[29512]++
									return true
//line /usr/local/go/src/text/template/parse/parse.go:276
		// _ = "end of CoverTab[29512]"
	case *IfNode:
//line /usr/local/go/src/text/template/parse/parse.go:277
		_go_fuzz_dep_.CoverTab[29513]++
//line /usr/local/go/src/text/template/parse/parse.go:277
		// _ = "end of CoverTab[29513]"
	case *ListNode:
//line /usr/local/go/src/text/template/parse/parse.go:278
		_go_fuzz_dep_.CoverTab[29514]++
									for _, node := range n.Nodes {
//line /usr/local/go/src/text/template/parse/parse.go:279
			_go_fuzz_dep_.CoverTab[29521]++
										if !IsEmptyTree(node) {
//line /usr/local/go/src/text/template/parse/parse.go:280
				_go_fuzz_dep_.CoverTab[29522]++
											return false
//line /usr/local/go/src/text/template/parse/parse.go:281
				// _ = "end of CoverTab[29522]"
			} else {
//line /usr/local/go/src/text/template/parse/parse.go:282
				_go_fuzz_dep_.CoverTab[29523]++
//line /usr/local/go/src/text/template/parse/parse.go:282
				// _ = "end of CoverTab[29523]"
//line /usr/local/go/src/text/template/parse/parse.go:282
			}
//line /usr/local/go/src/text/template/parse/parse.go:282
			// _ = "end of CoverTab[29521]"
		}
//line /usr/local/go/src/text/template/parse/parse.go:283
		// _ = "end of CoverTab[29514]"
//line /usr/local/go/src/text/template/parse/parse.go:283
		_go_fuzz_dep_.CoverTab[29515]++
									return true
//line /usr/local/go/src/text/template/parse/parse.go:284
		// _ = "end of CoverTab[29515]"
	case *RangeNode:
//line /usr/local/go/src/text/template/parse/parse.go:285
		_go_fuzz_dep_.CoverTab[29516]++
//line /usr/local/go/src/text/template/parse/parse.go:285
		// _ = "end of CoverTab[29516]"
	case *TemplateNode:
//line /usr/local/go/src/text/template/parse/parse.go:286
		_go_fuzz_dep_.CoverTab[29517]++
//line /usr/local/go/src/text/template/parse/parse.go:286
		// _ = "end of CoverTab[29517]"
	case *TextNode:
//line /usr/local/go/src/text/template/parse/parse.go:287
		_go_fuzz_dep_.CoverTab[29518]++
									return len(bytes.TrimSpace(n.Text)) == 0
//line /usr/local/go/src/text/template/parse/parse.go:288
		// _ = "end of CoverTab[29518]"
	case *WithNode:
//line /usr/local/go/src/text/template/parse/parse.go:289
		_go_fuzz_dep_.CoverTab[29519]++
//line /usr/local/go/src/text/template/parse/parse.go:289
		// _ = "end of CoverTab[29519]"
	default:
//line /usr/local/go/src/text/template/parse/parse.go:290
		_go_fuzz_dep_.CoverTab[29520]++
									panic("unknown node: " + n.String())
//line /usr/local/go/src/text/template/parse/parse.go:291
		// _ = "end of CoverTab[29520]"
	}
//line /usr/local/go/src/text/template/parse/parse.go:292
	// _ = "end of CoverTab[29508]"
//line /usr/local/go/src/text/template/parse/parse.go:292
	_go_fuzz_dep_.CoverTab[29509]++
								return false
//line /usr/local/go/src/text/template/parse/parse.go:293
	// _ = "end of CoverTab[29509]"
}

// parse is the top-level parser for a template, essentially the same
//line /usr/local/go/src/text/template/parse/parse.go:296
// as itemList except it also parses {{define}} actions.
//line /usr/local/go/src/text/template/parse/parse.go:296
// It runs to EOF.
//line /usr/local/go/src/text/template/parse/parse.go:299
func (t *Tree) parse() {
//line /usr/local/go/src/text/template/parse/parse.go:299
	_go_fuzz_dep_.CoverTab[29524]++
								t.Root = t.newList(t.peek().pos)
								for t.peek().typ != itemEOF {
//line /usr/local/go/src/text/template/parse/parse.go:301
		_go_fuzz_dep_.CoverTab[29525]++
									if t.peek().typ == itemLeftDelim {
//line /usr/local/go/src/text/template/parse/parse.go:302
			_go_fuzz_dep_.CoverTab[29527]++
										delim := t.next()
										if t.nextNonSpace().typ == itemDefine {
//line /usr/local/go/src/text/template/parse/parse.go:304
				_go_fuzz_dep_.CoverTab[29529]++
											newT := New("definition")
											newT.text = t.text
											newT.Mode = t.Mode
											newT.ParseName = t.ParseName
											newT.startParse(t.funcs, t.lex, t.treeSet)
											newT.parseDefinition()
											continue
//line /usr/local/go/src/text/template/parse/parse.go:311
				// _ = "end of CoverTab[29529]"
			} else {
//line /usr/local/go/src/text/template/parse/parse.go:312
				_go_fuzz_dep_.CoverTab[29530]++
//line /usr/local/go/src/text/template/parse/parse.go:312
				// _ = "end of CoverTab[29530]"
//line /usr/local/go/src/text/template/parse/parse.go:312
			}
//line /usr/local/go/src/text/template/parse/parse.go:312
			// _ = "end of CoverTab[29527]"
//line /usr/local/go/src/text/template/parse/parse.go:312
			_go_fuzz_dep_.CoverTab[29528]++
										t.backup2(delim)
//line /usr/local/go/src/text/template/parse/parse.go:313
			// _ = "end of CoverTab[29528]"
		} else {
//line /usr/local/go/src/text/template/parse/parse.go:314
			_go_fuzz_dep_.CoverTab[29531]++
//line /usr/local/go/src/text/template/parse/parse.go:314
			// _ = "end of CoverTab[29531]"
//line /usr/local/go/src/text/template/parse/parse.go:314
		}
//line /usr/local/go/src/text/template/parse/parse.go:314
		// _ = "end of CoverTab[29525]"
//line /usr/local/go/src/text/template/parse/parse.go:314
		_go_fuzz_dep_.CoverTab[29526]++
									switch n := t.textOrAction(); n.Type() {
		case nodeEnd, nodeElse:
//line /usr/local/go/src/text/template/parse/parse.go:316
			_go_fuzz_dep_.CoverTab[29532]++
										t.errorf("unexpected %s", n)
//line /usr/local/go/src/text/template/parse/parse.go:317
			// _ = "end of CoverTab[29532]"
		default:
//line /usr/local/go/src/text/template/parse/parse.go:318
			_go_fuzz_dep_.CoverTab[29533]++
										t.Root.append(n)
//line /usr/local/go/src/text/template/parse/parse.go:319
			// _ = "end of CoverTab[29533]"
		}
//line /usr/local/go/src/text/template/parse/parse.go:320
		// _ = "end of CoverTab[29526]"
	}
//line /usr/local/go/src/text/template/parse/parse.go:321
	// _ = "end of CoverTab[29524]"
}

// parseDefinition parses a {{define}} ...  {{end}} template definition and
//line /usr/local/go/src/text/template/parse/parse.go:324
// installs the definition in t.treeSet. The "define" keyword has already
//line /usr/local/go/src/text/template/parse/parse.go:324
// been scanned.
//line /usr/local/go/src/text/template/parse/parse.go:327
func (t *Tree) parseDefinition() {
//line /usr/local/go/src/text/template/parse/parse.go:327
	_go_fuzz_dep_.CoverTab[29534]++
								const context = "define clause"
								name := t.expectOneOf(itemString, itemRawString, context)
								var err error
								t.Name, err = strconv.Unquote(name.val)
								if err != nil {
//line /usr/local/go/src/text/template/parse/parse.go:332
		_go_fuzz_dep_.CoverTab[29537]++
									t.error(err)
//line /usr/local/go/src/text/template/parse/parse.go:333
		// _ = "end of CoverTab[29537]"
	} else {
//line /usr/local/go/src/text/template/parse/parse.go:334
		_go_fuzz_dep_.CoverTab[29538]++
//line /usr/local/go/src/text/template/parse/parse.go:334
		// _ = "end of CoverTab[29538]"
//line /usr/local/go/src/text/template/parse/parse.go:334
	}
//line /usr/local/go/src/text/template/parse/parse.go:334
	// _ = "end of CoverTab[29534]"
//line /usr/local/go/src/text/template/parse/parse.go:334
	_go_fuzz_dep_.CoverTab[29535]++
								t.expect(itemRightDelim, context)
								var end Node
								t.Root, end = t.itemList()
								if end.Type() != nodeEnd {
//line /usr/local/go/src/text/template/parse/parse.go:338
		_go_fuzz_dep_.CoverTab[29539]++
									t.errorf("unexpected %s in %s", end, context)
//line /usr/local/go/src/text/template/parse/parse.go:339
		// _ = "end of CoverTab[29539]"
	} else {
//line /usr/local/go/src/text/template/parse/parse.go:340
		_go_fuzz_dep_.CoverTab[29540]++
//line /usr/local/go/src/text/template/parse/parse.go:340
		// _ = "end of CoverTab[29540]"
//line /usr/local/go/src/text/template/parse/parse.go:340
	}
//line /usr/local/go/src/text/template/parse/parse.go:340
	// _ = "end of CoverTab[29535]"
//line /usr/local/go/src/text/template/parse/parse.go:340
	_go_fuzz_dep_.CoverTab[29536]++
								t.add()
								t.stopParse()
//line /usr/local/go/src/text/template/parse/parse.go:342
	// _ = "end of CoverTab[29536]"
}

// itemList:
//line /usr/local/go/src/text/template/parse/parse.go:345
//
//line /usr/local/go/src/text/template/parse/parse.go:345
//	textOrAction*
//line /usr/local/go/src/text/template/parse/parse.go:345
//
//line /usr/local/go/src/text/template/parse/parse.go:345
// Terminates at {{end}} or {{else}}, returned separately.
//line /usr/local/go/src/text/template/parse/parse.go:350
func (t *Tree) itemList() (list *ListNode, next Node) {
//line /usr/local/go/src/text/template/parse/parse.go:350
	_go_fuzz_dep_.CoverTab[29541]++
								list = t.newList(t.peekNonSpace().pos)
								for t.peekNonSpace().typ != itemEOF {
//line /usr/local/go/src/text/template/parse/parse.go:352
		_go_fuzz_dep_.CoverTab[29543]++
									n := t.textOrAction()
									switch n.Type() {
		case nodeEnd, nodeElse:
//line /usr/local/go/src/text/template/parse/parse.go:355
			_go_fuzz_dep_.CoverTab[29545]++
										return list, n
//line /usr/local/go/src/text/template/parse/parse.go:356
			// _ = "end of CoverTab[29545]"
//line /usr/local/go/src/text/template/parse/parse.go:356
		default:
//line /usr/local/go/src/text/template/parse/parse.go:356
			_go_fuzz_dep_.CoverTab[29546]++
//line /usr/local/go/src/text/template/parse/parse.go:356
			// _ = "end of CoverTab[29546]"
		}
//line /usr/local/go/src/text/template/parse/parse.go:357
		// _ = "end of CoverTab[29543]"
//line /usr/local/go/src/text/template/parse/parse.go:357
		_go_fuzz_dep_.CoverTab[29544]++
									list.append(n)
//line /usr/local/go/src/text/template/parse/parse.go:358
		// _ = "end of CoverTab[29544]"
	}
//line /usr/local/go/src/text/template/parse/parse.go:359
	// _ = "end of CoverTab[29541]"
//line /usr/local/go/src/text/template/parse/parse.go:359
	_go_fuzz_dep_.CoverTab[29542]++
								t.errorf("unexpected EOF")
								return
//line /usr/local/go/src/text/template/parse/parse.go:361
	// _ = "end of CoverTab[29542]"
}

// textOrAction:
//line /usr/local/go/src/text/template/parse/parse.go:364
//
//line /usr/local/go/src/text/template/parse/parse.go:364
//	text | comment | action
//line /usr/local/go/src/text/template/parse/parse.go:367
func (t *Tree) textOrAction() Node {
//line /usr/local/go/src/text/template/parse/parse.go:367
	_go_fuzz_dep_.CoverTab[29547]++
								switch token := t.nextNonSpace(); token.typ {
	case itemText:
//line /usr/local/go/src/text/template/parse/parse.go:369
		_go_fuzz_dep_.CoverTab[29549]++
									return t.newText(token.pos, token.val)
//line /usr/local/go/src/text/template/parse/parse.go:370
		// _ = "end of CoverTab[29549]"
	case itemLeftDelim:
//line /usr/local/go/src/text/template/parse/parse.go:371
		_go_fuzz_dep_.CoverTab[29550]++
									t.actionLine = token.line
									defer t.clearActionLine()
									return t.action()
//line /usr/local/go/src/text/template/parse/parse.go:374
		// _ = "end of CoverTab[29550]"
	case itemComment:
//line /usr/local/go/src/text/template/parse/parse.go:375
		_go_fuzz_dep_.CoverTab[29551]++
									return t.newComment(token.pos, token.val)
//line /usr/local/go/src/text/template/parse/parse.go:376
		// _ = "end of CoverTab[29551]"
	default:
//line /usr/local/go/src/text/template/parse/parse.go:377
		_go_fuzz_dep_.CoverTab[29552]++
									t.unexpected(token, "input")
//line /usr/local/go/src/text/template/parse/parse.go:378
		// _ = "end of CoverTab[29552]"
	}
//line /usr/local/go/src/text/template/parse/parse.go:379
	// _ = "end of CoverTab[29547]"
//line /usr/local/go/src/text/template/parse/parse.go:379
	_go_fuzz_dep_.CoverTab[29548]++
								return nil
//line /usr/local/go/src/text/template/parse/parse.go:380
	// _ = "end of CoverTab[29548]"
}

func (t *Tree) clearActionLine() {
//line /usr/local/go/src/text/template/parse/parse.go:383
	_go_fuzz_dep_.CoverTab[29553]++
								t.actionLine = 0
//line /usr/local/go/src/text/template/parse/parse.go:384
	// _ = "end of CoverTab[29553]"
}

// Action:
//line /usr/local/go/src/text/template/parse/parse.go:387
//
//line /usr/local/go/src/text/template/parse/parse.go:387
//	control
//line /usr/local/go/src/text/template/parse/parse.go:387
//	command ("|" command)*
//line /usr/local/go/src/text/template/parse/parse.go:387
//
//line /usr/local/go/src/text/template/parse/parse.go:387
// Left delim is past. Now get actions.
//line /usr/local/go/src/text/template/parse/parse.go:387
// First word could be a keyword such as range.
//line /usr/local/go/src/text/template/parse/parse.go:394
func (t *Tree) action() (n Node) {
//line /usr/local/go/src/text/template/parse/parse.go:394
	_go_fuzz_dep_.CoverTab[29554]++
								switch token := t.nextNonSpace(); token.typ {
	case itemBlock:
//line /usr/local/go/src/text/template/parse/parse.go:396
		_go_fuzz_dep_.CoverTab[29556]++
									return t.blockControl()
//line /usr/local/go/src/text/template/parse/parse.go:397
		// _ = "end of CoverTab[29556]"
	case itemBreak:
//line /usr/local/go/src/text/template/parse/parse.go:398
		_go_fuzz_dep_.CoverTab[29557]++
									return t.breakControl(token.pos, token.line)
//line /usr/local/go/src/text/template/parse/parse.go:399
		// _ = "end of CoverTab[29557]"
	case itemContinue:
//line /usr/local/go/src/text/template/parse/parse.go:400
		_go_fuzz_dep_.CoverTab[29558]++
									return t.continueControl(token.pos, token.line)
//line /usr/local/go/src/text/template/parse/parse.go:401
		// _ = "end of CoverTab[29558]"
	case itemElse:
//line /usr/local/go/src/text/template/parse/parse.go:402
		_go_fuzz_dep_.CoverTab[29559]++
									return t.elseControl()
//line /usr/local/go/src/text/template/parse/parse.go:403
		// _ = "end of CoverTab[29559]"
	case itemEnd:
//line /usr/local/go/src/text/template/parse/parse.go:404
		_go_fuzz_dep_.CoverTab[29560]++
									return t.endControl()
//line /usr/local/go/src/text/template/parse/parse.go:405
		// _ = "end of CoverTab[29560]"
	case itemIf:
//line /usr/local/go/src/text/template/parse/parse.go:406
		_go_fuzz_dep_.CoverTab[29561]++
									return t.ifControl()
//line /usr/local/go/src/text/template/parse/parse.go:407
		// _ = "end of CoverTab[29561]"
	case itemRange:
//line /usr/local/go/src/text/template/parse/parse.go:408
		_go_fuzz_dep_.CoverTab[29562]++
									return t.rangeControl()
//line /usr/local/go/src/text/template/parse/parse.go:409
		// _ = "end of CoverTab[29562]"
	case itemTemplate:
//line /usr/local/go/src/text/template/parse/parse.go:410
		_go_fuzz_dep_.CoverTab[29563]++
									return t.templateControl()
//line /usr/local/go/src/text/template/parse/parse.go:411
		// _ = "end of CoverTab[29563]"
	case itemWith:
//line /usr/local/go/src/text/template/parse/parse.go:412
		_go_fuzz_dep_.CoverTab[29564]++
									return t.withControl()
//line /usr/local/go/src/text/template/parse/parse.go:413
		// _ = "end of CoverTab[29564]"
//line /usr/local/go/src/text/template/parse/parse.go:413
	default:
//line /usr/local/go/src/text/template/parse/parse.go:413
		_go_fuzz_dep_.CoverTab[29565]++
//line /usr/local/go/src/text/template/parse/parse.go:413
		// _ = "end of CoverTab[29565]"
	}
//line /usr/local/go/src/text/template/parse/parse.go:414
	// _ = "end of CoverTab[29554]"
//line /usr/local/go/src/text/template/parse/parse.go:414
	_go_fuzz_dep_.CoverTab[29555]++
								t.backup()
								token := t.peek()

								return t.newAction(token.pos, token.line, t.pipeline("command", itemRightDelim))
//line /usr/local/go/src/text/template/parse/parse.go:418
	// _ = "end of CoverTab[29555]"
}

// Break:
//line /usr/local/go/src/text/template/parse/parse.go:421
//
//line /usr/local/go/src/text/template/parse/parse.go:421
//	{{break}}
//line /usr/local/go/src/text/template/parse/parse.go:421
//
//line /usr/local/go/src/text/template/parse/parse.go:421
// Break keyword is past.
//line /usr/local/go/src/text/template/parse/parse.go:426
func (t *Tree) breakControl(pos Pos, line int) Node {
//line /usr/local/go/src/text/template/parse/parse.go:426
	_go_fuzz_dep_.CoverTab[29566]++
								if token := t.nextNonSpace(); token.typ != itemRightDelim {
//line /usr/local/go/src/text/template/parse/parse.go:427
		_go_fuzz_dep_.CoverTab[29569]++
									t.unexpected(token, "{{break}}")
//line /usr/local/go/src/text/template/parse/parse.go:428
		// _ = "end of CoverTab[29569]"
	} else {
//line /usr/local/go/src/text/template/parse/parse.go:429
		_go_fuzz_dep_.CoverTab[29570]++
//line /usr/local/go/src/text/template/parse/parse.go:429
		// _ = "end of CoverTab[29570]"
//line /usr/local/go/src/text/template/parse/parse.go:429
	}
//line /usr/local/go/src/text/template/parse/parse.go:429
	// _ = "end of CoverTab[29566]"
//line /usr/local/go/src/text/template/parse/parse.go:429
	_go_fuzz_dep_.CoverTab[29567]++
								if t.rangeDepth == 0 {
//line /usr/local/go/src/text/template/parse/parse.go:430
		_go_fuzz_dep_.CoverTab[29571]++
									t.errorf("{{break}} outside {{range}}")
//line /usr/local/go/src/text/template/parse/parse.go:431
		// _ = "end of CoverTab[29571]"
	} else {
//line /usr/local/go/src/text/template/parse/parse.go:432
		_go_fuzz_dep_.CoverTab[29572]++
//line /usr/local/go/src/text/template/parse/parse.go:432
		// _ = "end of CoverTab[29572]"
//line /usr/local/go/src/text/template/parse/parse.go:432
	}
//line /usr/local/go/src/text/template/parse/parse.go:432
	// _ = "end of CoverTab[29567]"
//line /usr/local/go/src/text/template/parse/parse.go:432
	_go_fuzz_dep_.CoverTab[29568]++
								return t.newBreak(pos, line)
//line /usr/local/go/src/text/template/parse/parse.go:433
	// _ = "end of CoverTab[29568]"
}

// Continue:
//line /usr/local/go/src/text/template/parse/parse.go:436
//
//line /usr/local/go/src/text/template/parse/parse.go:436
//	{{continue}}
//line /usr/local/go/src/text/template/parse/parse.go:436
//
//line /usr/local/go/src/text/template/parse/parse.go:436
// Continue keyword is past.
//line /usr/local/go/src/text/template/parse/parse.go:441
func (t *Tree) continueControl(pos Pos, line int) Node {
//line /usr/local/go/src/text/template/parse/parse.go:441
	_go_fuzz_dep_.CoverTab[29573]++
								if token := t.nextNonSpace(); token.typ != itemRightDelim {
//line /usr/local/go/src/text/template/parse/parse.go:442
		_go_fuzz_dep_.CoverTab[29576]++
									t.unexpected(token, "{{continue}}")
//line /usr/local/go/src/text/template/parse/parse.go:443
		// _ = "end of CoverTab[29576]"
	} else {
//line /usr/local/go/src/text/template/parse/parse.go:444
		_go_fuzz_dep_.CoverTab[29577]++
//line /usr/local/go/src/text/template/parse/parse.go:444
		// _ = "end of CoverTab[29577]"
//line /usr/local/go/src/text/template/parse/parse.go:444
	}
//line /usr/local/go/src/text/template/parse/parse.go:444
	// _ = "end of CoverTab[29573]"
//line /usr/local/go/src/text/template/parse/parse.go:444
	_go_fuzz_dep_.CoverTab[29574]++
								if t.rangeDepth == 0 {
//line /usr/local/go/src/text/template/parse/parse.go:445
		_go_fuzz_dep_.CoverTab[29578]++
									t.errorf("{{continue}} outside {{range}}")
//line /usr/local/go/src/text/template/parse/parse.go:446
		// _ = "end of CoverTab[29578]"
	} else {
//line /usr/local/go/src/text/template/parse/parse.go:447
		_go_fuzz_dep_.CoverTab[29579]++
//line /usr/local/go/src/text/template/parse/parse.go:447
		// _ = "end of CoverTab[29579]"
//line /usr/local/go/src/text/template/parse/parse.go:447
	}
//line /usr/local/go/src/text/template/parse/parse.go:447
	// _ = "end of CoverTab[29574]"
//line /usr/local/go/src/text/template/parse/parse.go:447
	_go_fuzz_dep_.CoverTab[29575]++
								return t.newContinue(pos, line)
//line /usr/local/go/src/text/template/parse/parse.go:448
	// _ = "end of CoverTab[29575]"
}

// Pipeline:
//line /usr/local/go/src/text/template/parse/parse.go:451
//
//line /usr/local/go/src/text/template/parse/parse.go:451
//	declarations? command ('|' command)*
//line /usr/local/go/src/text/template/parse/parse.go:454
func (t *Tree) pipeline(context string, end itemType) (pipe *PipeNode) {
//line /usr/local/go/src/text/template/parse/parse.go:454
	_go_fuzz_dep_.CoverTab[29580]++
								token := t.peekNonSpace()
								pipe = t.newPipeline(token.pos, token.line, nil)

decls:
	if v := t.peekNonSpace(); v.typ == itemVariable {
//line /usr/local/go/src/text/template/parse/parse.go:459
		_go_fuzz_dep_.CoverTab[29582]++
									t.next()

//line /usr/local/go/src/text/template/parse/parse.go:465
		tokenAfterVariable := t.peek()
		next := t.peekNonSpace()
		switch {
		case next.typ == itemAssign, next.typ == itemDeclare:
//line /usr/local/go/src/text/template/parse/parse.go:468
			_go_fuzz_dep_.CoverTab[29583]++
										pipe.IsAssign = next.typ == itemAssign
										t.nextNonSpace()
										pipe.Decl = append(pipe.Decl, t.newVariable(v.pos, v.val))
										t.vars = append(t.vars, v.val)
//line /usr/local/go/src/text/template/parse/parse.go:472
			// _ = "end of CoverTab[29583]"
		case next.typ == itemChar && func() bool {
//line /usr/local/go/src/text/template/parse/parse.go:473
			_go_fuzz_dep_.CoverTab[29588]++
//line /usr/local/go/src/text/template/parse/parse.go:473
			return next.val == ","
//line /usr/local/go/src/text/template/parse/parse.go:473
			// _ = "end of CoverTab[29588]"
//line /usr/local/go/src/text/template/parse/parse.go:473
		}():
//line /usr/local/go/src/text/template/parse/parse.go:473
			_go_fuzz_dep_.CoverTab[29584]++
										t.nextNonSpace()
										pipe.Decl = append(pipe.Decl, t.newVariable(v.pos, v.val))
										t.vars = append(t.vars, v.val)
										if context == "range" && func() bool {
//line /usr/local/go/src/text/template/parse/parse.go:477
				_go_fuzz_dep_.CoverTab[29589]++
//line /usr/local/go/src/text/template/parse/parse.go:477
				return len(pipe.Decl) < 2
//line /usr/local/go/src/text/template/parse/parse.go:477
				// _ = "end of CoverTab[29589]"
//line /usr/local/go/src/text/template/parse/parse.go:477
			}() {
//line /usr/local/go/src/text/template/parse/parse.go:477
				_go_fuzz_dep_.CoverTab[29590]++
											switch t.peekNonSpace().typ {
				case itemVariable, itemRightDelim, itemRightParen:
//line /usr/local/go/src/text/template/parse/parse.go:479
					_go_fuzz_dep_.CoverTab[29591]++

												goto decls
//line /usr/local/go/src/text/template/parse/parse.go:481
					// _ = "end of CoverTab[29591]"
				default:
//line /usr/local/go/src/text/template/parse/parse.go:482
					_go_fuzz_dep_.CoverTab[29592]++
												t.errorf("range can only initialize variables")
//line /usr/local/go/src/text/template/parse/parse.go:483
					// _ = "end of CoverTab[29592]"
				}
//line /usr/local/go/src/text/template/parse/parse.go:484
				// _ = "end of CoverTab[29590]"
			} else {
//line /usr/local/go/src/text/template/parse/parse.go:485
				_go_fuzz_dep_.CoverTab[29593]++
//line /usr/local/go/src/text/template/parse/parse.go:485
				// _ = "end of CoverTab[29593]"
//line /usr/local/go/src/text/template/parse/parse.go:485
			}
//line /usr/local/go/src/text/template/parse/parse.go:485
			// _ = "end of CoverTab[29584]"
//line /usr/local/go/src/text/template/parse/parse.go:485
			_go_fuzz_dep_.CoverTab[29585]++
										t.errorf("too many declarations in %s", context)
//line /usr/local/go/src/text/template/parse/parse.go:486
			// _ = "end of CoverTab[29585]"
		case tokenAfterVariable.typ == itemSpace:
//line /usr/local/go/src/text/template/parse/parse.go:487
			_go_fuzz_dep_.CoverTab[29586]++
										t.backup3(v, tokenAfterVariable)
//line /usr/local/go/src/text/template/parse/parse.go:488
			// _ = "end of CoverTab[29586]"
		default:
//line /usr/local/go/src/text/template/parse/parse.go:489
			_go_fuzz_dep_.CoverTab[29587]++
										t.backup2(v)
//line /usr/local/go/src/text/template/parse/parse.go:490
			// _ = "end of CoverTab[29587]"
		}
//line /usr/local/go/src/text/template/parse/parse.go:491
		// _ = "end of CoverTab[29582]"
	} else {
//line /usr/local/go/src/text/template/parse/parse.go:492
		_go_fuzz_dep_.CoverTab[29594]++
//line /usr/local/go/src/text/template/parse/parse.go:492
		// _ = "end of CoverTab[29594]"
//line /usr/local/go/src/text/template/parse/parse.go:492
	}
//line /usr/local/go/src/text/template/parse/parse.go:492
	// _ = "end of CoverTab[29580]"
//line /usr/local/go/src/text/template/parse/parse.go:492
	_go_fuzz_dep_.CoverTab[29581]++
								for {
//line /usr/local/go/src/text/template/parse/parse.go:493
		_go_fuzz_dep_.CoverTab[29595]++
									switch token := t.nextNonSpace(); token.typ {
		case end:
//line /usr/local/go/src/text/template/parse/parse.go:495
			_go_fuzz_dep_.CoverTab[29596]++

										t.checkPipeline(pipe, context)
										return
//line /usr/local/go/src/text/template/parse/parse.go:498
			// _ = "end of CoverTab[29596]"
		case itemBool, itemCharConstant, itemComplex, itemDot, itemField, itemIdentifier,
			itemNumber, itemNil, itemRawString, itemString, itemVariable, itemLeftParen:
//line /usr/local/go/src/text/template/parse/parse.go:500
			_go_fuzz_dep_.CoverTab[29597]++
										t.backup()
										pipe.append(t.command())
//line /usr/local/go/src/text/template/parse/parse.go:502
			// _ = "end of CoverTab[29597]"
		default:
//line /usr/local/go/src/text/template/parse/parse.go:503
			_go_fuzz_dep_.CoverTab[29598]++
										t.unexpected(token, context)
//line /usr/local/go/src/text/template/parse/parse.go:504
			// _ = "end of CoverTab[29598]"
		}
//line /usr/local/go/src/text/template/parse/parse.go:505
		// _ = "end of CoverTab[29595]"
	}
//line /usr/local/go/src/text/template/parse/parse.go:506
	// _ = "end of CoverTab[29581]"
}

func (t *Tree) checkPipeline(pipe *PipeNode, context string) {
//line /usr/local/go/src/text/template/parse/parse.go:509
	_go_fuzz_dep_.CoverTab[29599]++

								if len(pipe.Cmds) == 0 {
//line /usr/local/go/src/text/template/parse/parse.go:511
		_go_fuzz_dep_.CoverTab[29601]++
									t.errorf("missing value for %s", context)
//line /usr/local/go/src/text/template/parse/parse.go:512
		// _ = "end of CoverTab[29601]"
	} else {
//line /usr/local/go/src/text/template/parse/parse.go:513
		_go_fuzz_dep_.CoverTab[29602]++
//line /usr/local/go/src/text/template/parse/parse.go:513
		// _ = "end of CoverTab[29602]"
//line /usr/local/go/src/text/template/parse/parse.go:513
	}
//line /usr/local/go/src/text/template/parse/parse.go:513
	// _ = "end of CoverTab[29599]"
//line /usr/local/go/src/text/template/parse/parse.go:513
	_go_fuzz_dep_.CoverTab[29600]++

								for i, c := range pipe.Cmds[1:] {
//line /usr/local/go/src/text/template/parse/parse.go:515
		_go_fuzz_dep_.CoverTab[29603]++
									switch c.Args[0].Type() {
		case NodeBool, NodeDot, NodeNil, NodeNumber, NodeString:
//line /usr/local/go/src/text/template/parse/parse.go:517
			_go_fuzz_dep_.CoverTab[29604]++

										t.errorf("non executable command in pipeline stage %d", i+2)
//line /usr/local/go/src/text/template/parse/parse.go:519
			// _ = "end of CoverTab[29604]"
//line /usr/local/go/src/text/template/parse/parse.go:519
		default:
//line /usr/local/go/src/text/template/parse/parse.go:519
			_go_fuzz_dep_.CoverTab[29605]++
//line /usr/local/go/src/text/template/parse/parse.go:519
			// _ = "end of CoverTab[29605]"
		}
//line /usr/local/go/src/text/template/parse/parse.go:520
		// _ = "end of CoverTab[29603]"
	}
//line /usr/local/go/src/text/template/parse/parse.go:521
	// _ = "end of CoverTab[29600]"
}

func (t *Tree) parseControl(allowElseIf bool, context string) (pos Pos, line int, pipe *PipeNode, list, elseList *ListNode) {
//line /usr/local/go/src/text/template/parse/parse.go:524
	_go_fuzz_dep_.CoverTab[29606]++
								defer t.popVars(len(t.vars))
								pipe = t.pipeline(context, itemRightDelim)
								if context == "range" {
//line /usr/local/go/src/text/template/parse/parse.go:527
		_go_fuzz_dep_.CoverTab[29610]++
									t.rangeDepth++
//line /usr/local/go/src/text/template/parse/parse.go:528
		// _ = "end of CoverTab[29610]"
	} else {
//line /usr/local/go/src/text/template/parse/parse.go:529
		_go_fuzz_dep_.CoverTab[29611]++
//line /usr/local/go/src/text/template/parse/parse.go:529
		// _ = "end of CoverTab[29611]"
//line /usr/local/go/src/text/template/parse/parse.go:529
	}
//line /usr/local/go/src/text/template/parse/parse.go:529
	// _ = "end of CoverTab[29606]"
//line /usr/local/go/src/text/template/parse/parse.go:529
	_go_fuzz_dep_.CoverTab[29607]++
								var next Node
								list, next = t.itemList()
								if context == "range" {
//line /usr/local/go/src/text/template/parse/parse.go:532
		_go_fuzz_dep_.CoverTab[29612]++
									t.rangeDepth--
//line /usr/local/go/src/text/template/parse/parse.go:533
		// _ = "end of CoverTab[29612]"
	} else {
//line /usr/local/go/src/text/template/parse/parse.go:534
		_go_fuzz_dep_.CoverTab[29613]++
//line /usr/local/go/src/text/template/parse/parse.go:534
		// _ = "end of CoverTab[29613]"
//line /usr/local/go/src/text/template/parse/parse.go:534
	}
//line /usr/local/go/src/text/template/parse/parse.go:534
	// _ = "end of CoverTab[29607]"
//line /usr/local/go/src/text/template/parse/parse.go:534
	_go_fuzz_dep_.CoverTab[29608]++
								switch next.Type() {
	case nodeEnd:
//line /usr/local/go/src/text/template/parse/parse.go:536
		_go_fuzz_dep_.CoverTab[29614]++
//line /usr/local/go/src/text/template/parse/parse.go:536
		// _ = "end of CoverTab[29614]"
	case nodeElse:
//line /usr/local/go/src/text/template/parse/parse.go:537
		_go_fuzz_dep_.CoverTab[29615]++
									if allowElseIf {
//line /usr/local/go/src/text/template/parse/parse.go:538
			_go_fuzz_dep_.CoverTab[29618]++

//line /usr/local/go/src/text/template/parse/parse.go:547
			if t.peek().typ == itemIf {
//line /usr/local/go/src/text/template/parse/parse.go:547
				_go_fuzz_dep_.CoverTab[29619]++
											t.next()
											elseList = t.newList(next.Position())
											elseList.append(t.ifControl())

											break
//line /usr/local/go/src/text/template/parse/parse.go:552
				// _ = "end of CoverTab[29619]"
			} else {
//line /usr/local/go/src/text/template/parse/parse.go:553
				_go_fuzz_dep_.CoverTab[29620]++
//line /usr/local/go/src/text/template/parse/parse.go:553
				// _ = "end of CoverTab[29620]"
//line /usr/local/go/src/text/template/parse/parse.go:553
			}
//line /usr/local/go/src/text/template/parse/parse.go:553
			// _ = "end of CoverTab[29618]"
		} else {
//line /usr/local/go/src/text/template/parse/parse.go:554
			_go_fuzz_dep_.CoverTab[29621]++
//line /usr/local/go/src/text/template/parse/parse.go:554
			// _ = "end of CoverTab[29621]"
//line /usr/local/go/src/text/template/parse/parse.go:554
		}
//line /usr/local/go/src/text/template/parse/parse.go:554
		// _ = "end of CoverTab[29615]"
//line /usr/local/go/src/text/template/parse/parse.go:554
		_go_fuzz_dep_.CoverTab[29616]++
									elseList, next = t.itemList()
									if next.Type() != nodeEnd {
//line /usr/local/go/src/text/template/parse/parse.go:556
			_go_fuzz_dep_.CoverTab[29622]++
										t.errorf("expected end; found %s", next)
//line /usr/local/go/src/text/template/parse/parse.go:557
			// _ = "end of CoverTab[29622]"
		} else {
//line /usr/local/go/src/text/template/parse/parse.go:558
			_go_fuzz_dep_.CoverTab[29623]++
//line /usr/local/go/src/text/template/parse/parse.go:558
			// _ = "end of CoverTab[29623]"
//line /usr/local/go/src/text/template/parse/parse.go:558
		}
//line /usr/local/go/src/text/template/parse/parse.go:558
		// _ = "end of CoverTab[29616]"
//line /usr/local/go/src/text/template/parse/parse.go:558
	default:
//line /usr/local/go/src/text/template/parse/parse.go:558
		_go_fuzz_dep_.CoverTab[29617]++
//line /usr/local/go/src/text/template/parse/parse.go:558
		// _ = "end of CoverTab[29617]"
	}
//line /usr/local/go/src/text/template/parse/parse.go:559
	// _ = "end of CoverTab[29608]"
//line /usr/local/go/src/text/template/parse/parse.go:559
	_go_fuzz_dep_.CoverTab[29609]++
								return pipe.Position(), pipe.Line, pipe, list, elseList
//line /usr/local/go/src/text/template/parse/parse.go:560
	// _ = "end of CoverTab[29609]"
}

// If:
//line /usr/local/go/src/text/template/parse/parse.go:563
//
//line /usr/local/go/src/text/template/parse/parse.go:563
//	{{if pipeline}} itemList {{end}}
//line /usr/local/go/src/text/template/parse/parse.go:563
//	{{if pipeline}} itemList {{else}} itemList {{end}}
//line /usr/local/go/src/text/template/parse/parse.go:563
//
//line /usr/local/go/src/text/template/parse/parse.go:563
// If keyword is past.
//line /usr/local/go/src/text/template/parse/parse.go:569
func (t *Tree) ifControl() Node {
//line /usr/local/go/src/text/template/parse/parse.go:569
	_go_fuzz_dep_.CoverTab[29624]++
								return t.newIf(t.parseControl(true, "if"))
//line /usr/local/go/src/text/template/parse/parse.go:570
	// _ = "end of CoverTab[29624]"
}

// Range:
//line /usr/local/go/src/text/template/parse/parse.go:573
//
//line /usr/local/go/src/text/template/parse/parse.go:573
//	{{range pipeline}} itemList {{end}}
//line /usr/local/go/src/text/template/parse/parse.go:573
//	{{range pipeline}} itemList {{else}} itemList {{end}}
//line /usr/local/go/src/text/template/parse/parse.go:573
//
//line /usr/local/go/src/text/template/parse/parse.go:573
// Range keyword is past.
//line /usr/local/go/src/text/template/parse/parse.go:579
func (t *Tree) rangeControl() Node {
//line /usr/local/go/src/text/template/parse/parse.go:579
	_go_fuzz_dep_.CoverTab[29625]++
								r := t.newRange(t.parseControl(false, "range"))
								return r
//line /usr/local/go/src/text/template/parse/parse.go:581
	// _ = "end of CoverTab[29625]"
}

// With:
//line /usr/local/go/src/text/template/parse/parse.go:584
//
//line /usr/local/go/src/text/template/parse/parse.go:584
//	{{with pipeline}} itemList {{end}}
//line /usr/local/go/src/text/template/parse/parse.go:584
//	{{with pipeline}} itemList {{else}} itemList {{end}}
//line /usr/local/go/src/text/template/parse/parse.go:584
//
//line /usr/local/go/src/text/template/parse/parse.go:584
// If keyword is past.
//line /usr/local/go/src/text/template/parse/parse.go:590
func (t *Tree) withControl() Node {
//line /usr/local/go/src/text/template/parse/parse.go:590
	_go_fuzz_dep_.CoverTab[29626]++
								return t.newWith(t.parseControl(false, "with"))
//line /usr/local/go/src/text/template/parse/parse.go:591
	// _ = "end of CoverTab[29626]"
}

// End:
//line /usr/local/go/src/text/template/parse/parse.go:594
//
//line /usr/local/go/src/text/template/parse/parse.go:594
//	{{end}}
//line /usr/local/go/src/text/template/parse/parse.go:594
//
//line /usr/local/go/src/text/template/parse/parse.go:594
// End keyword is past.
//line /usr/local/go/src/text/template/parse/parse.go:599
func (t *Tree) endControl() Node {
//line /usr/local/go/src/text/template/parse/parse.go:599
	_go_fuzz_dep_.CoverTab[29627]++
								return t.newEnd(t.expect(itemRightDelim, "end").pos)
//line /usr/local/go/src/text/template/parse/parse.go:600
	// _ = "end of CoverTab[29627]"
}

// Else:
//line /usr/local/go/src/text/template/parse/parse.go:603
//
//line /usr/local/go/src/text/template/parse/parse.go:603
//	{{else}}
//line /usr/local/go/src/text/template/parse/parse.go:603
//
//line /usr/local/go/src/text/template/parse/parse.go:603
// Else keyword is past.
//line /usr/local/go/src/text/template/parse/parse.go:608
func (t *Tree) elseControl() Node {
//line /usr/local/go/src/text/template/parse/parse.go:608
	_go_fuzz_dep_.CoverTab[29628]++

								peek := t.peekNonSpace()
								if peek.typ == itemIf {
//line /usr/local/go/src/text/template/parse/parse.go:611
		_go_fuzz_dep_.CoverTab[29630]++

									return t.newElse(peek.pos, peek.line)
//line /usr/local/go/src/text/template/parse/parse.go:613
		// _ = "end of CoverTab[29630]"
	} else {
//line /usr/local/go/src/text/template/parse/parse.go:614
		_go_fuzz_dep_.CoverTab[29631]++
//line /usr/local/go/src/text/template/parse/parse.go:614
		// _ = "end of CoverTab[29631]"
//line /usr/local/go/src/text/template/parse/parse.go:614
	}
//line /usr/local/go/src/text/template/parse/parse.go:614
	// _ = "end of CoverTab[29628]"
//line /usr/local/go/src/text/template/parse/parse.go:614
	_go_fuzz_dep_.CoverTab[29629]++
								token := t.expect(itemRightDelim, "else")
								return t.newElse(token.pos, token.line)
//line /usr/local/go/src/text/template/parse/parse.go:616
	// _ = "end of CoverTab[29629]"
}

// Block:
//line /usr/local/go/src/text/template/parse/parse.go:619
//
//line /usr/local/go/src/text/template/parse/parse.go:619
//	{{block stringValue pipeline}}
//line /usr/local/go/src/text/template/parse/parse.go:619
//
//line /usr/local/go/src/text/template/parse/parse.go:619
// Block keyword is past.
//line /usr/local/go/src/text/template/parse/parse.go:619
// The name must be something that can evaluate to a string.
//line /usr/local/go/src/text/template/parse/parse.go:619
// The pipeline is mandatory.
//line /usr/local/go/src/text/template/parse/parse.go:626
func (t *Tree) blockControl() Node {
//line /usr/local/go/src/text/template/parse/parse.go:626
	_go_fuzz_dep_.CoverTab[29632]++
								const context = "block clause"

								token := t.nextNonSpace()
								name := t.parseTemplateName(token, context)
								pipe := t.pipeline(context, itemRightDelim)

								block := New(name)
								block.text = t.text
								block.Mode = t.Mode
								block.ParseName = t.ParseName
								block.startParse(t.funcs, t.lex, t.treeSet)
								var end Node
								block.Root, end = block.itemList()
								if end.Type() != nodeEnd {
//line /usr/local/go/src/text/template/parse/parse.go:640
		_go_fuzz_dep_.CoverTab[29634]++
									t.errorf("unexpected %s in %s", end, context)
//line /usr/local/go/src/text/template/parse/parse.go:641
		// _ = "end of CoverTab[29634]"
	} else {
//line /usr/local/go/src/text/template/parse/parse.go:642
		_go_fuzz_dep_.CoverTab[29635]++
//line /usr/local/go/src/text/template/parse/parse.go:642
		// _ = "end of CoverTab[29635]"
//line /usr/local/go/src/text/template/parse/parse.go:642
	}
//line /usr/local/go/src/text/template/parse/parse.go:642
	// _ = "end of CoverTab[29632]"
//line /usr/local/go/src/text/template/parse/parse.go:642
	_go_fuzz_dep_.CoverTab[29633]++
								block.add()
								block.stopParse()

								return t.newTemplate(token.pos, token.line, name, pipe)
//line /usr/local/go/src/text/template/parse/parse.go:646
	// _ = "end of CoverTab[29633]"
}

// Template:
//line /usr/local/go/src/text/template/parse/parse.go:649
//
//line /usr/local/go/src/text/template/parse/parse.go:649
//	{{template stringValue pipeline}}
//line /usr/local/go/src/text/template/parse/parse.go:649
//
//line /usr/local/go/src/text/template/parse/parse.go:649
// Template keyword is past. The name must be something that can evaluate
//line /usr/local/go/src/text/template/parse/parse.go:649
// to a string.
//line /usr/local/go/src/text/template/parse/parse.go:655
func (t *Tree) templateControl() Node {
//line /usr/local/go/src/text/template/parse/parse.go:655
	_go_fuzz_dep_.CoverTab[29636]++
								const context = "template clause"
								token := t.nextNonSpace()
								name := t.parseTemplateName(token, context)
								var pipe *PipeNode
								if t.nextNonSpace().typ != itemRightDelim {
//line /usr/local/go/src/text/template/parse/parse.go:660
		_go_fuzz_dep_.CoverTab[29638]++
									t.backup()

									pipe = t.pipeline(context, itemRightDelim)
//line /usr/local/go/src/text/template/parse/parse.go:663
		// _ = "end of CoverTab[29638]"
	} else {
//line /usr/local/go/src/text/template/parse/parse.go:664
		_go_fuzz_dep_.CoverTab[29639]++
//line /usr/local/go/src/text/template/parse/parse.go:664
		// _ = "end of CoverTab[29639]"
//line /usr/local/go/src/text/template/parse/parse.go:664
	}
//line /usr/local/go/src/text/template/parse/parse.go:664
	// _ = "end of CoverTab[29636]"
//line /usr/local/go/src/text/template/parse/parse.go:664
	_go_fuzz_dep_.CoverTab[29637]++
								return t.newTemplate(token.pos, token.line, name, pipe)
//line /usr/local/go/src/text/template/parse/parse.go:665
	// _ = "end of CoverTab[29637]"
}

func (t *Tree) parseTemplateName(token item, context string) (name string) {
//line /usr/local/go/src/text/template/parse/parse.go:668
	_go_fuzz_dep_.CoverTab[29640]++
								switch token.typ {
	case itemString, itemRawString:
//line /usr/local/go/src/text/template/parse/parse.go:670
		_go_fuzz_dep_.CoverTab[29642]++
									s, err := strconv.Unquote(token.val)
									if err != nil {
//line /usr/local/go/src/text/template/parse/parse.go:672
			_go_fuzz_dep_.CoverTab[29645]++
										t.error(err)
//line /usr/local/go/src/text/template/parse/parse.go:673
			// _ = "end of CoverTab[29645]"
		} else {
//line /usr/local/go/src/text/template/parse/parse.go:674
			_go_fuzz_dep_.CoverTab[29646]++
//line /usr/local/go/src/text/template/parse/parse.go:674
			// _ = "end of CoverTab[29646]"
//line /usr/local/go/src/text/template/parse/parse.go:674
		}
//line /usr/local/go/src/text/template/parse/parse.go:674
		// _ = "end of CoverTab[29642]"
//line /usr/local/go/src/text/template/parse/parse.go:674
		_go_fuzz_dep_.CoverTab[29643]++
									name = s
//line /usr/local/go/src/text/template/parse/parse.go:675
		// _ = "end of CoverTab[29643]"
	default:
//line /usr/local/go/src/text/template/parse/parse.go:676
		_go_fuzz_dep_.CoverTab[29644]++
									t.unexpected(token, context)
//line /usr/local/go/src/text/template/parse/parse.go:677
		// _ = "end of CoverTab[29644]"
	}
//line /usr/local/go/src/text/template/parse/parse.go:678
	// _ = "end of CoverTab[29640]"
//line /usr/local/go/src/text/template/parse/parse.go:678
	_go_fuzz_dep_.CoverTab[29641]++
								return
//line /usr/local/go/src/text/template/parse/parse.go:679
	// _ = "end of CoverTab[29641]"
}

// command:
//line /usr/local/go/src/text/template/parse/parse.go:682
//
//line /usr/local/go/src/text/template/parse/parse.go:682
//	operand (space operand)*
//line /usr/local/go/src/text/template/parse/parse.go:682
//
//line /usr/local/go/src/text/template/parse/parse.go:682
// space-separated arguments up to a pipeline character or right delimiter.
//line /usr/local/go/src/text/template/parse/parse.go:682
// we consume the pipe character but leave the right delim to terminate the action.
//line /usr/local/go/src/text/template/parse/parse.go:688
func (t *Tree) command() *CommandNode {
//line /usr/local/go/src/text/template/parse/parse.go:688
	_go_fuzz_dep_.CoverTab[29647]++
								cmd := t.newCommand(t.peekNonSpace().pos)
								for {
//line /usr/local/go/src/text/template/parse/parse.go:690
		_go_fuzz_dep_.CoverTab[29650]++
									t.peekNonSpace()
									operand := t.operand()
									if operand != nil {
//line /usr/local/go/src/text/template/parse/parse.go:693
			_go_fuzz_dep_.CoverTab[29653]++
										cmd.append(operand)
//line /usr/local/go/src/text/template/parse/parse.go:694
			// _ = "end of CoverTab[29653]"
		} else {
//line /usr/local/go/src/text/template/parse/parse.go:695
			_go_fuzz_dep_.CoverTab[29654]++
//line /usr/local/go/src/text/template/parse/parse.go:695
			// _ = "end of CoverTab[29654]"
//line /usr/local/go/src/text/template/parse/parse.go:695
		}
//line /usr/local/go/src/text/template/parse/parse.go:695
		// _ = "end of CoverTab[29650]"
//line /usr/local/go/src/text/template/parse/parse.go:695
		_go_fuzz_dep_.CoverTab[29651]++
									switch token := t.next(); token.typ {
		case itemSpace:
//line /usr/local/go/src/text/template/parse/parse.go:697
			_go_fuzz_dep_.CoverTab[29655]++
										continue
//line /usr/local/go/src/text/template/parse/parse.go:698
			// _ = "end of CoverTab[29655]"
		case itemRightDelim, itemRightParen:
//line /usr/local/go/src/text/template/parse/parse.go:699
			_go_fuzz_dep_.CoverTab[29656]++
										t.backup()
//line /usr/local/go/src/text/template/parse/parse.go:700
			// _ = "end of CoverTab[29656]"
		case itemPipe:
//line /usr/local/go/src/text/template/parse/parse.go:701
			_go_fuzz_dep_.CoverTab[29657]++
//line /usr/local/go/src/text/template/parse/parse.go:701
			// _ = "end of CoverTab[29657]"

		default:
//line /usr/local/go/src/text/template/parse/parse.go:703
			_go_fuzz_dep_.CoverTab[29658]++
										t.unexpected(token, "operand")
//line /usr/local/go/src/text/template/parse/parse.go:704
			// _ = "end of CoverTab[29658]"
		}
//line /usr/local/go/src/text/template/parse/parse.go:705
		// _ = "end of CoverTab[29651]"
//line /usr/local/go/src/text/template/parse/parse.go:705
		_go_fuzz_dep_.CoverTab[29652]++
									break
//line /usr/local/go/src/text/template/parse/parse.go:706
		// _ = "end of CoverTab[29652]"
	}
//line /usr/local/go/src/text/template/parse/parse.go:707
	// _ = "end of CoverTab[29647]"
//line /usr/local/go/src/text/template/parse/parse.go:707
	_go_fuzz_dep_.CoverTab[29648]++
								if len(cmd.Args) == 0 {
//line /usr/local/go/src/text/template/parse/parse.go:708
		_go_fuzz_dep_.CoverTab[29659]++
									t.errorf("empty command")
//line /usr/local/go/src/text/template/parse/parse.go:709
		// _ = "end of CoverTab[29659]"
	} else {
//line /usr/local/go/src/text/template/parse/parse.go:710
		_go_fuzz_dep_.CoverTab[29660]++
//line /usr/local/go/src/text/template/parse/parse.go:710
		// _ = "end of CoverTab[29660]"
//line /usr/local/go/src/text/template/parse/parse.go:710
	}
//line /usr/local/go/src/text/template/parse/parse.go:710
	// _ = "end of CoverTab[29648]"
//line /usr/local/go/src/text/template/parse/parse.go:710
	_go_fuzz_dep_.CoverTab[29649]++
								return cmd
//line /usr/local/go/src/text/template/parse/parse.go:711
	// _ = "end of CoverTab[29649]"
}

// operand:
//line /usr/local/go/src/text/template/parse/parse.go:714
//
//line /usr/local/go/src/text/template/parse/parse.go:714
//	term .Field*
//line /usr/local/go/src/text/template/parse/parse.go:714
//
//line /usr/local/go/src/text/template/parse/parse.go:714
// An operand is a space-separated component of a command,
//line /usr/local/go/src/text/template/parse/parse.go:714
// a term possibly followed by field accesses.
//line /usr/local/go/src/text/template/parse/parse.go:714
// A nil return means the next item is not an operand.
//line /usr/local/go/src/text/template/parse/parse.go:721
func (t *Tree) operand() Node {
//line /usr/local/go/src/text/template/parse/parse.go:721
	_go_fuzz_dep_.CoverTab[29661]++
								node := t.term()
								if node == nil {
//line /usr/local/go/src/text/template/parse/parse.go:723
		_go_fuzz_dep_.CoverTab[29664]++
									return nil
//line /usr/local/go/src/text/template/parse/parse.go:724
		// _ = "end of CoverTab[29664]"
	} else {
//line /usr/local/go/src/text/template/parse/parse.go:725
		_go_fuzz_dep_.CoverTab[29665]++
//line /usr/local/go/src/text/template/parse/parse.go:725
		// _ = "end of CoverTab[29665]"
//line /usr/local/go/src/text/template/parse/parse.go:725
	}
//line /usr/local/go/src/text/template/parse/parse.go:725
	// _ = "end of CoverTab[29661]"
//line /usr/local/go/src/text/template/parse/parse.go:725
	_go_fuzz_dep_.CoverTab[29662]++
								if t.peek().typ == itemField {
//line /usr/local/go/src/text/template/parse/parse.go:726
		_go_fuzz_dep_.CoverTab[29666]++
									chain := t.newChain(t.peek().pos, node)
									for t.peek().typ == itemField {
//line /usr/local/go/src/text/template/parse/parse.go:728
			_go_fuzz_dep_.CoverTab[29668]++
										chain.Add(t.next().val)
//line /usr/local/go/src/text/template/parse/parse.go:729
			// _ = "end of CoverTab[29668]"
		}
//line /usr/local/go/src/text/template/parse/parse.go:730
		// _ = "end of CoverTab[29666]"
//line /usr/local/go/src/text/template/parse/parse.go:730
		_go_fuzz_dep_.CoverTab[29667]++

//line /usr/local/go/src/text/template/parse/parse.go:736
		switch node.Type() {
		case NodeField:
//line /usr/local/go/src/text/template/parse/parse.go:737
			_go_fuzz_dep_.CoverTab[29669]++
										node = t.newField(chain.Position(), chain.String())
//line /usr/local/go/src/text/template/parse/parse.go:738
			// _ = "end of CoverTab[29669]"
		case NodeVariable:
//line /usr/local/go/src/text/template/parse/parse.go:739
			_go_fuzz_dep_.CoverTab[29670]++
										node = t.newVariable(chain.Position(), chain.String())
//line /usr/local/go/src/text/template/parse/parse.go:740
			// _ = "end of CoverTab[29670]"
		case NodeBool, NodeString, NodeNumber, NodeNil, NodeDot:
//line /usr/local/go/src/text/template/parse/parse.go:741
			_go_fuzz_dep_.CoverTab[29671]++
										t.errorf("unexpected . after term %q", node.String())
//line /usr/local/go/src/text/template/parse/parse.go:742
			// _ = "end of CoverTab[29671]"
		default:
//line /usr/local/go/src/text/template/parse/parse.go:743
			_go_fuzz_dep_.CoverTab[29672]++
										node = chain
//line /usr/local/go/src/text/template/parse/parse.go:744
			// _ = "end of CoverTab[29672]"
		}
//line /usr/local/go/src/text/template/parse/parse.go:745
		// _ = "end of CoverTab[29667]"
	} else {
//line /usr/local/go/src/text/template/parse/parse.go:746
		_go_fuzz_dep_.CoverTab[29673]++
//line /usr/local/go/src/text/template/parse/parse.go:746
		// _ = "end of CoverTab[29673]"
//line /usr/local/go/src/text/template/parse/parse.go:746
	}
//line /usr/local/go/src/text/template/parse/parse.go:746
	// _ = "end of CoverTab[29662]"
//line /usr/local/go/src/text/template/parse/parse.go:746
	_go_fuzz_dep_.CoverTab[29663]++
								return node
//line /usr/local/go/src/text/template/parse/parse.go:747
	// _ = "end of CoverTab[29663]"
}

// term:
//line /usr/local/go/src/text/template/parse/parse.go:750
//
//line /usr/local/go/src/text/template/parse/parse.go:750
//	literal (number, string, nil, boolean)
//line /usr/local/go/src/text/template/parse/parse.go:750
//	function (identifier)
//line /usr/local/go/src/text/template/parse/parse.go:750
//	.
//line /usr/local/go/src/text/template/parse/parse.go:750
//	.Field
//line /usr/local/go/src/text/template/parse/parse.go:750
//	$
//line /usr/local/go/src/text/template/parse/parse.go:750
//	'(' pipeline ')'
//line /usr/local/go/src/text/template/parse/parse.go:750
//
//line /usr/local/go/src/text/template/parse/parse.go:750
// A term is a simple "expression".
//line /usr/local/go/src/text/template/parse/parse.go:750
// A nil return means the next item is not a term.
//line /usr/local/go/src/text/template/parse/parse.go:761
func (t *Tree) term() Node {
//line /usr/local/go/src/text/template/parse/parse.go:761
	_go_fuzz_dep_.CoverTab[29674]++
								switch token := t.nextNonSpace(); token.typ {
	case itemIdentifier:
//line /usr/local/go/src/text/template/parse/parse.go:763
		_go_fuzz_dep_.CoverTab[29676]++
									checkFunc := t.Mode&SkipFuncCheck == 0
									if checkFunc && func() bool {
//line /usr/local/go/src/text/template/parse/parse.go:765
			_go_fuzz_dep_.CoverTab[29689]++
//line /usr/local/go/src/text/template/parse/parse.go:765
			return !t.hasFunction(token.val)
//line /usr/local/go/src/text/template/parse/parse.go:765
			// _ = "end of CoverTab[29689]"
//line /usr/local/go/src/text/template/parse/parse.go:765
		}() {
//line /usr/local/go/src/text/template/parse/parse.go:765
			_go_fuzz_dep_.CoverTab[29690]++
										t.errorf("function %q not defined", token.val)
//line /usr/local/go/src/text/template/parse/parse.go:766
			// _ = "end of CoverTab[29690]"
		} else {
//line /usr/local/go/src/text/template/parse/parse.go:767
			_go_fuzz_dep_.CoverTab[29691]++
//line /usr/local/go/src/text/template/parse/parse.go:767
			// _ = "end of CoverTab[29691]"
//line /usr/local/go/src/text/template/parse/parse.go:767
		}
//line /usr/local/go/src/text/template/parse/parse.go:767
		// _ = "end of CoverTab[29676]"
//line /usr/local/go/src/text/template/parse/parse.go:767
		_go_fuzz_dep_.CoverTab[29677]++
									return NewIdentifier(token.val).SetTree(t).SetPos(token.pos)
//line /usr/local/go/src/text/template/parse/parse.go:768
		// _ = "end of CoverTab[29677]"
	case itemDot:
//line /usr/local/go/src/text/template/parse/parse.go:769
		_go_fuzz_dep_.CoverTab[29678]++
									return t.newDot(token.pos)
//line /usr/local/go/src/text/template/parse/parse.go:770
		// _ = "end of CoverTab[29678]"
	case itemNil:
//line /usr/local/go/src/text/template/parse/parse.go:771
		_go_fuzz_dep_.CoverTab[29679]++
									return t.newNil(token.pos)
//line /usr/local/go/src/text/template/parse/parse.go:772
		// _ = "end of CoverTab[29679]"
	case itemVariable:
//line /usr/local/go/src/text/template/parse/parse.go:773
		_go_fuzz_dep_.CoverTab[29680]++
									return t.useVar(token.pos, token.val)
//line /usr/local/go/src/text/template/parse/parse.go:774
		// _ = "end of CoverTab[29680]"
	case itemField:
//line /usr/local/go/src/text/template/parse/parse.go:775
		_go_fuzz_dep_.CoverTab[29681]++
									return t.newField(token.pos, token.val)
//line /usr/local/go/src/text/template/parse/parse.go:776
		// _ = "end of CoverTab[29681]"
	case itemBool:
//line /usr/local/go/src/text/template/parse/parse.go:777
		_go_fuzz_dep_.CoverTab[29682]++
									return t.newBool(token.pos, token.val == "true")
//line /usr/local/go/src/text/template/parse/parse.go:778
		// _ = "end of CoverTab[29682]"
	case itemCharConstant, itemComplex, itemNumber:
//line /usr/local/go/src/text/template/parse/parse.go:779
		_go_fuzz_dep_.CoverTab[29683]++
									number, err := t.newNumber(token.pos, token.val, token.typ)
									if err != nil {
//line /usr/local/go/src/text/template/parse/parse.go:781
			_go_fuzz_dep_.CoverTab[29692]++
										t.error(err)
//line /usr/local/go/src/text/template/parse/parse.go:782
			// _ = "end of CoverTab[29692]"
		} else {
//line /usr/local/go/src/text/template/parse/parse.go:783
			_go_fuzz_dep_.CoverTab[29693]++
//line /usr/local/go/src/text/template/parse/parse.go:783
			// _ = "end of CoverTab[29693]"
//line /usr/local/go/src/text/template/parse/parse.go:783
		}
//line /usr/local/go/src/text/template/parse/parse.go:783
		// _ = "end of CoverTab[29683]"
//line /usr/local/go/src/text/template/parse/parse.go:783
		_go_fuzz_dep_.CoverTab[29684]++
									return number
//line /usr/local/go/src/text/template/parse/parse.go:784
		// _ = "end of CoverTab[29684]"
	case itemLeftParen:
//line /usr/local/go/src/text/template/parse/parse.go:785
		_go_fuzz_dep_.CoverTab[29685]++
									return t.pipeline("parenthesized pipeline", itemRightParen)
//line /usr/local/go/src/text/template/parse/parse.go:786
		// _ = "end of CoverTab[29685]"
	case itemString, itemRawString:
//line /usr/local/go/src/text/template/parse/parse.go:787
		_go_fuzz_dep_.CoverTab[29686]++
									s, err := strconv.Unquote(token.val)
									if err != nil {
//line /usr/local/go/src/text/template/parse/parse.go:789
			_go_fuzz_dep_.CoverTab[29694]++
										t.error(err)
//line /usr/local/go/src/text/template/parse/parse.go:790
			// _ = "end of CoverTab[29694]"
		} else {
//line /usr/local/go/src/text/template/parse/parse.go:791
			_go_fuzz_dep_.CoverTab[29695]++
//line /usr/local/go/src/text/template/parse/parse.go:791
			// _ = "end of CoverTab[29695]"
//line /usr/local/go/src/text/template/parse/parse.go:791
		}
//line /usr/local/go/src/text/template/parse/parse.go:791
		// _ = "end of CoverTab[29686]"
//line /usr/local/go/src/text/template/parse/parse.go:791
		_go_fuzz_dep_.CoverTab[29687]++
									return t.newString(token.pos, token.val, s)
//line /usr/local/go/src/text/template/parse/parse.go:792
		// _ = "end of CoverTab[29687]"
//line /usr/local/go/src/text/template/parse/parse.go:792
	default:
//line /usr/local/go/src/text/template/parse/parse.go:792
		_go_fuzz_dep_.CoverTab[29688]++
//line /usr/local/go/src/text/template/parse/parse.go:792
		// _ = "end of CoverTab[29688]"
	}
//line /usr/local/go/src/text/template/parse/parse.go:793
	// _ = "end of CoverTab[29674]"
//line /usr/local/go/src/text/template/parse/parse.go:793
	_go_fuzz_dep_.CoverTab[29675]++
								t.backup()
								return nil
//line /usr/local/go/src/text/template/parse/parse.go:795
	// _ = "end of CoverTab[29675]"
}

// hasFunction reports if a function name exists in the Tree's maps.
func (t *Tree) hasFunction(name string) bool {
//line /usr/local/go/src/text/template/parse/parse.go:799
	_go_fuzz_dep_.CoverTab[29696]++
								for _, funcMap := range t.funcs {
//line /usr/local/go/src/text/template/parse/parse.go:800
		_go_fuzz_dep_.CoverTab[29698]++
									if funcMap == nil {
//line /usr/local/go/src/text/template/parse/parse.go:801
			_go_fuzz_dep_.CoverTab[29700]++
										continue
//line /usr/local/go/src/text/template/parse/parse.go:802
			// _ = "end of CoverTab[29700]"
		} else {
//line /usr/local/go/src/text/template/parse/parse.go:803
			_go_fuzz_dep_.CoverTab[29701]++
//line /usr/local/go/src/text/template/parse/parse.go:803
			// _ = "end of CoverTab[29701]"
//line /usr/local/go/src/text/template/parse/parse.go:803
		}
//line /usr/local/go/src/text/template/parse/parse.go:803
		// _ = "end of CoverTab[29698]"
//line /usr/local/go/src/text/template/parse/parse.go:803
		_go_fuzz_dep_.CoverTab[29699]++
									if funcMap[name] != nil {
//line /usr/local/go/src/text/template/parse/parse.go:804
			_go_fuzz_dep_.CoverTab[29702]++
										return true
//line /usr/local/go/src/text/template/parse/parse.go:805
			// _ = "end of CoverTab[29702]"
		} else {
//line /usr/local/go/src/text/template/parse/parse.go:806
			_go_fuzz_dep_.CoverTab[29703]++
//line /usr/local/go/src/text/template/parse/parse.go:806
			// _ = "end of CoverTab[29703]"
//line /usr/local/go/src/text/template/parse/parse.go:806
		}
//line /usr/local/go/src/text/template/parse/parse.go:806
		// _ = "end of CoverTab[29699]"
	}
//line /usr/local/go/src/text/template/parse/parse.go:807
	// _ = "end of CoverTab[29696]"
//line /usr/local/go/src/text/template/parse/parse.go:807
	_go_fuzz_dep_.CoverTab[29697]++
								return false
//line /usr/local/go/src/text/template/parse/parse.go:808
	// _ = "end of CoverTab[29697]"
}

// popVars trims the variable list to the specified length
func (t *Tree) popVars(n int) {
//line /usr/local/go/src/text/template/parse/parse.go:812
	_go_fuzz_dep_.CoverTab[29704]++
								t.vars = t.vars[:n]
//line /usr/local/go/src/text/template/parse/parse.go:813
	// _ = "end of CoverTab[29704]"
}

// useVar returns a node for a variable reference. It errors if the
//line /usr/local/go/src/text/template/parse/parse.go:816
// variable is not defined.
//line /usr/local/go/src/text/template/parse/parse.go:818
func (t *Tree) useVar(pos Pos, name string) Node {
//line /usr/local/go/src/text/template/parse/parse.go:818
	_go_fuzz_dep_.CoverTab[29705]++
								v := t.newVariable(pos, name)
								for _, varName := range t.vars {
//line /usr/local/go/src/text/template/parse/parse.go:820
		_go_fuzz_dep_.CoverTab[29707]++
									if varName == v.Ident[0] {
//line /usr/local/go/src/text/template/parse/parse.go:821
			_go_fuzz_dep_.CoverTab[29708]++
										return v
//line /usr/local/go/src/text/template/parse/parse.go:822
			// _ = "end of CoverTab[29708]"
		} else {
//line /usr/local/go/src/text/template/parse/parse.go:823
			_go_fuzz_dep_.CoverTab[29709]++
//line /usr/local/go/src/text/template/parse/parse.go:823
			// _ = "end of CoverTab[29709]"
//line /usr/local/go/src/text/template/parse/parse.go:823
		}
//line /usr/local/go/src/text/template/parse/parse.go:823
		// _ = "end of CoverTab[29707]"
	}
//line /usr/local/go/src/text/template/parse/parse.go:824
	// _ = "end of CoverTab[29705]"
//line /usr/local/go/src/text/template/parse/parse.go:824
	_go_fuzz_dep_.CoverTab[29706]++
								t.errorf("undefined variable %q", v.Ident[0])
								return nil
//line /usr/local/go/src/text/template/parse/parse.go:826
	// _ = "end of CoverTab[29706]"
}

//line /usr/local/go/src/text/template/parse/parse.go:827
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/text/template/parse/parse.go:827
var _ = _go_fuzz_dep_.CoverTab
