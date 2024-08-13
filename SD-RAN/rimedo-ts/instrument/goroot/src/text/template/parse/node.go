// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Parse nodes.

//line /usr/local/go/src/text/template/parse/node.go:7
package parse

//line /usr/local/go/src/text/template/parse/node.go:7
import (
//line /usr/local/go/src/text/template/parse/node.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/text/template/parse/node.go:7
)
//line /usr/local/go/src/text/template/parse/node.go:7
import (
//line /usr/local/go/src/text/template/parse/node.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/text/template/parse/node.go:7
)

import (
	"fmt"
	"strconv"
	"strings"
)

var textFormat = "%s"	// Changed to "%q" in tests for better error messages.

// A Node is an element in the parse tree. The interface is trivial.
//line /usr/local/go/src/text/template/parse/node.go:17
// The interface contains an unexported method so that only
//line /usr/local/go/src/text/template/parse/node.go:17
// types local to this package can satisfy it.
//line /usr/local/go/src/text/template/parse/node.go:20
type Node interface {
	Type() NodeType
	String() string
	// Copy does a deep copy of the Node and all its components.
	// To avoid type assertions, some XxxNodes also have specialized
	// CopyXxx methods that return *XxxNode.
	Copy() Node
	Position() Pos	// byte position of start of node in full original input string
	// tree returns the containing *Tree.
	// It is unexported so all implementations of Node are in this package.
	tree() *Tree
	// writeTo writes the String output to the builder.
	writeTo(*strings.Builder)
}

// NodeType identifies the type of a parse tree node.
type NodeType int

// Pos represents a byte position in the original input text from which
//line /usr/local/go/src/text/template/parse/node.go:38
// this template was parsed.
//line /usr/local/go/src/text/template/parse/node.go:40
type Pos int

func (p Pos) Position() Pos {
//line /usr/local/go/src/text/template/parse/node.go:42
	_go_fuzz_dep_.CoverTab[29187]++
								return p
//line /usr/local/go/src/text/template/parse/node.go:43
	// _ = "end of CoverTab[29187]"
}

// Type returns itself and provides an easy default implementation
//line /usr/local/go/src/text/template/parse/node.go:46
// for embedding in a Node. Embedded in all non-trivial Nodes.
//line /usr/local/go/src/text/template/parse/node.go:48
func (t NodeType) Type() NodeType {
//line /usr/local/go/src/text/template/parse/node.go:48
	_go_fuzz_dep_.CoverTab[29188]++
								return t
//line /usr/local/go/src/text/template/parse/node.go:49
	// _ = "end of CoverTab[29188]"
}

const (
	NodeText	NodeType	= iota	// Plain text.
	NodeAction				// A non-control action such as a field evaluation.
	NodeBool				// A boolean constant.
	NodeChain				// A sequence of field accesses.
	NodeCommand				// An element of a pipeline.
	NodeDot					// The cursor, dot.
	nodeElse				// An else action. Not added to tree.
	nodeEnd					// An end action. Not added to tree.
	NodeField				// A field or method name.
	NodeIdentifier				// An identifier; always a function name.
	NodeIf					// An if action.
	NodeList				// A list of Nodes.
	NodeNil					// An untyped nil constant.
	NodeNumber				// A numerical constant.
	NodePipe				// A pipeline of commands.
	NodeRange				// A range action.
	NodeString				// A string constant.
	NodeTemplate				// A template invocation action.
	NodeVariable				// A $ variable.
	NodeWith				// A with action.
	NodeComment				// A comment.
	NodeBreak				// A break action.
	NodeContinue				// A continue action.
)

//line /usr/local/go/src/text/template/parse/node.go:80
// ListNode holds a sequence of nodes.
type ListNode struct {
	NodeType
	Pos
	tr	*Tree
	Nodes	[]Node	// The element nodes in lexical order.
}

func (t *Tree) newList(pos Pos) *ListNode {
//line /usr/local/go/src/text/template/parse/node.go:88
	_go_fuzz_dep_.CoverTab[29189]++
								return &ListNode{tr: t, NodeType: NodeList, Pos: pos}
//line /usr/local/go/src/text/template/parse/node.go:89
	// _ = "end of CoverTab[29189]"
}

func (l *ListNode) append(n Node) {
//line /usr/local/go/src/text/template/parse/node.go:92
	_go_fuzz_dep_.CoverTab[29190]++
								l.Nodes = append(l.Nodes, n)
//line /usr/local/go/src/text/template/parse/node.go:93
	// _ = "end of CoverTab[29190]"
}

func (l *ListNode) tree() *Tree {
//line /usr/local/go/src/text/template/parse/node.go:96
	_go_fuzz_dep_.CoverTab[29191]++
								return l.tr
//line /usr/local/go/src/text/template/parse/node.go:97
	// _ = "end of CoverTab[29191]"
}

func (l *ListNode) String() string {
//line /usr/local/go/src/text/template/parse/node.go:100
	_go_fuzz_dep_.CoverTab[29192]++
								var sb strings.Builder
								l.writeTo(&sb)
								return sb.String()
//line /usr/local/go/src/text/template/parse/node.go:103
	// _ = "end of CoverTab[29192]"
}

func (l *ListNode) writeTo(sb *strings.Builder) {
//line /usr/local/go/src/text/template/parse/node.go:106
	_go_fuzz_dep_.CoverTab[29193]++
								for _, n := range l.Nodes {
//line /usr/local/go/src/text/template/parse/node.go:107
		_go_fuzz_dep_.CoverTab[29194]++
									n.writeTo(sb)
//line /usr/local/go/src/text/template/parse/node.go:108
		// _ = "end of CoverTab[29194]"
	}
//line /usr/local/go/src/text/template/parse/node.go:109
	// _ = "end of CoverTab[29193]"
}

func (l *ListNode) CopyList() *ListNode {
//line /usr/local/go/src/text/template/parse/node.go:112
	_go_fuzz_dep_.CoverTab[29195]++
								if l == nil {
//line /usr/local/go/src/text/template/parse/node.go:113
		_go_fuzz_dep_.CoverTab[29198]++
									return l
//line /usr/local/go/src/text/template/parse/node.go:114
		// _ = "end of CoverTab[29198]"
	} else {
//line /usr/local/go/src/text/template/parse/node.go:115
		_go_fuzz_dep_.CoverTab[29199]++
//line /usr/local/go/src/text/template/parse/node.go:115
		// _ = "end of CoverTab[29199]"
//line /usr/local/go/src/text/template/parse/node.go:115
	}
//line /usr/local/go/src/text/template/parse/node.go:115
	// _ = "end of CoverTab[29195]"
//line /usr/local/go/src/text/template/parse/node.go:115
	_go_fuzz_dep_.CoverTab[29196]++
								n := l.tr.newList(l.Pos)
								for _, elem := range l.Nodes {
//line /usr/local/go/src/text/template/parse/node.go:117
		_go_fuzz_dep_.CoverTab[29200]++
									n.append(elem.Copy())
//line /usr/local/go/src/text/template/parse/node.go:118
		// _ = "end of CoverTab[29200]"
	}
//line /usr/local/go/src/text/template/parse/node.go:119
	// _ = "end of CoverTab[29196]"
//line /usr/local/go/src/text/template/parse/node.go:119
	_go_fuzz_dep_.CoverTab[29197]++
								return n
//line /usr/local/go/src/text/template/parse/node.go:120
	// _ = "end of CoverTab[29197]"
}

func (l *ListNode) Copy() Node {
//line /usr/local/go/src/text/template/parse/node.go:123
	_go_fuzz_dep_.CoverTab[29201]++
								return l.CopyList()
//line /usr/local/go/src/text/template/parse/node.go:124
	// _ = "end of CoverTab[29201]"
}

// TextNode holds plain text.
type TextNode struct {
	NodeType
	Pos
	tr	*Tree
	Text	[]byte	// The text; may span newlines.
}

func (t *Tree) newText(pos Pos, text string) *TextNode {
//line /usr/local/go/src/text/template/parse/node.go:135
	_go_fuzz_dep_.CoverTab[29202]++
								return &TextNode{tr: t, NodeType: NodeText, Pos: pos, Text: []byte(text)}
//line /usr/local/go/src/text/template/parse/node.go:136
	// _ = "end of CoverTab[29202]"
}

func (t *TextNode) String() string {
//line /usr/local/go/src/text/template/parse/node.go:139
	_go_fuzz_dep_.CoverTab[29203]++
								return fmt.Sprintf(textFormat, t.Text)
//line /usr/local/go/src/text/template/parse/node.go:140
	// _ = "end of CoverTab[29203]"
}

func (t *TextNode) writeTo(sb *strings.Builder) {
//line /usr/local/go/src/text/template/parse/node.go:143
	_go_fuzz_dep_.CoverTab[29204]++
								sb.WriteString(t.String())
//line /usr/local/go/src/text/template/parse/node.go:144
	// _ = "end of CoverTab[29204]"
}

func (t *TextNode) tree() *Tree {
//line /usr/local/go/src/text/template/parse/node.go:147
	_go_fuzz_dep_.CoverTab[29205]++
								return t.tr
//line /usr/local/go/src/text/template/parse/node.go:148
	// _ = "end of CoverTab[29205]"
}

func (t *TextNode) Copy() Node {
//line /usr/local/go/src/text/template/parse/node.go:151
	_go_fuzz_dep_.CoverTab[29206]++
								return &TextNode{tr: t.tr, NodeType: NodeText, Pos: t.Pos, Text: append([]byte{}, t.Text...)}
//line /usr/local/go/src/text/template/parse/node.go:152
	// _ = "end of CoverTab[29206]"
}

// CommentNode holds a comment.
type CommentNode struct {
	NodeType
	Pos
	tr	*Tree
	Text	string	// Comment text.
}

func (t *Tree) newComment(pos Pos, text string) *CommentNode {
//line /usr/local/go/src/text/template/parse/node.go:163
	_go_fuzz_dep_.CoverTab[29207]++
								return &CommentNode{tr: t, NodeType: NodeComment, Pos: pos, Text: text}
//line /usr/local/go/src/text/template/parse/node.go:164
	// _ = "end of CoverTab[29207]"
}

func (c *CommentNode) String() string {
//line /usr/local/go/src/text/template/parse/node.go:167
	_go_fuzz_dep_.CoverTab[29208]++
								var sb strings.Builder
								c.writeTo(&sb)
								return sb.String()
//line /usr/local/go/src/text/template/parse/node.go:170
	// _ = "end of CoverTab[29208]"
}

func (c *CommentNode) writeTo(sb *strings.Builder) {
//line /usr/local/go/src/text/template/parse/node.go:173
	_go_fuzz_dep_.CoverTab[29209]++
								sb.WriteString("{{")
								sb.WriteString(c.Text)
								sb.WriteString("}}")
//line /usr/local/go/src/text/template/parse/node.go:176
	// _ = "end of CoverTab[29209]"
}

func (c *CommentNode) tree() *Tree {
//line /usr/local/go/src/text/template/parse/node.go:179
	_go_fuzz_dep_.CoverTab[29210]++
								return c.tr
//line /usr/local/go/src/text/template/parse/node.go:180
	// _ = "end of CoverTab[29210]"
}

func (c *CommentNode) Copy() Node {
//line /usr/local/go/src/text/template/parse/node.go:183
	_go_fuzz_dep_.CoverTab[29211]++
								return &CommentNode{tr: c.tr, NodeType: NodeComment, Pos: c.Pos, Text: c.Text}
//line /usr/local/go/src/text/template/parse/node.go:184
	// _ = "end of CoverTab[29211]"
}

// PipeNode holds a pipeline with optional declaration
type PipeNode struct {
	NodeType
	Pos
	tr		*Tree
	Line		int		// The line number in the input. Deprecated: Kept for compatibility.
	IsAssign	bool		// The variables are being assigned, not declared.
	Decl		[]*VariableNode	// Variables in lexical order.
	Cmds		[]*CommandNode	// The commands in lexical order.
}

func (t *Tree) newPipeline(pos Pos, line int, vars []*VariableNode) *PipeNode {
//line /usr/local/go/src/text/template/parse/node.go:198
	_go_fuzz_dep_.CoverTab[29212]++
								return &PipeNode{tr: t, NodeType: NodePipe, Pos: pos, Line: line, Decl: vars}
//line /usr/local/go/src/text/template/parse/node.go:199
	// _ = "end of CoverTab[29212]"
}

func (p *PipeNode) append(command *CommandNode) {
//line /usr/local/go/src/text/template/parse/node.go:202
	_go_fuzz_dep_.CoverTab[29213]++
								p.Cmds = append(p.Cmds, command)
//line /usr/local/go/src/text/template/parse/node.go:203
	// _ = "end of CoverTab[29213]"
}

func (p *PipeNode) String() string {
//line /usr/local/go/src/text/template/parse/node.go:206
	_go_fuzz_dep_.CoverTab[29214]++
								var sb strings.Builder
								p.writeTo(&sb)
								return sb.String()
//line /usr/local/go/src/text/template/parse/node.go:209
	// _ = "end of CoverTab[29214]"
}

func (p *PipeNode) writeTo(sb *strings.Builder) {
//line /usr/local/go/src/text/template/parse/node.go:212
	_go_fuzz_dep_.CoverTab[29215]++
								if len(p.Decl) > 0 {
//line /usr/local/go/src/text/template/parse/node.go:213
		_go_fuzz_dep_.CoverTab[29217]++
									for i, v := range p.Decl {
//line /usr/local/go/src/text/template/parse/node.go:214
			_go_fuzz_dep_.CoverTab[29219]++
										if i > 0 {
//line /usr/local/go/src/text/template/parse/node.go:215
				_go_fuzz_dep_.CoverTab[29221]++
											sb.WriteString(", ")
//line /usr/local/go/src/text/template/parse/node.go:216
				// _ = "end of CoverTab[29221]"
			} else {
//line /usr/local/go/src/text/template/parse/node.go:217
				_go_fuzz_dep_.CoverTab[29222]++
//line /usr/local/go/src/text/template/parse/node.go:217
				// _ = "end of CoverTab[29222]"
//line /usr/local/go/src/text/template/parse/node.go:217
			}
//line /usr/local/go/src/text/template/parse/node.go:217
			// _ = "end of CoverTab[29219]"
//line /usr/local/go/src/text/template/parse/node.go:217
			_go_fuzz_dep_.CoverTab[29220]++
										v.writeTo(sb)
//line /usr/local/go/src/text/template/parse/node.go:218
			// _ = "end of CoverTab[29220]"
		}
//line /usr/local/go/src/text/template/parse/node.go:219
		// _ = "end of CoverTab[29217]"
//line /usr/local/go/src/text/template/parse/node.go:219
		_go_fuzz_dep_.CoverTab[29218]++
									sb.WriteString(" := ")
//line /usr/local/go/src/text/template/parse/node.go:220
		// _ = "end of CoverTab[29218]"
	} else {
//line /usr/local/go/src/text/template/parse/node.go:221
		_go_fuzz_dep_.CoverTab[29223]++
//line /usr/local/go/src/text/template/parse/node.go:221
		// _ = "end of CoverTab[29223]"
//line /usr/local/go/src/text/template/parse/node.go:221
	}
//line /usr/local/go/src/text/template/parse/node.go:221
	// _ = "end of CoverTab[29215]"
//line /usr/local/go/src/text/template/parse/node.go:221
	_go_fuzz_dep_.CoverTab[29216]++
								for i, c := range p.Cmds {
//line /usr/local/go/src/text/template/parse/node.go:222
		_go_fuzz_dep_.CoverTab[29224]++
									if i > 0 {
//line /usr/local/go/src/text/template/parse/node.go:223
			_go_fuzz_dep_.CoverTab[29226]++
										sb.WriteString(" | ")
//line /usr/local/go/src/text/template/parse/node.go:224
			// _ = "end of CoverTab[29226]"
		} else {
//line /usr/local/go/src/text/template/parse/node.go:225
			_go_fuzz_dep_.CoverTab[29227]++
//line /usr/local/go/src/text/template/parse/node.go:225
			// _ = "end of CoverTab[29227]"
//line /usr/local/go/src/text/template/parse/node.go:225
		}
//line /usr/local/go/src/text/template/parse/node.go:225
		// _ = "end of CoverTab[29224]"
//line /usr/local/go/src/text/template/parse/node.go:225
		_go_fuzz_dep_.CoverTab[29225]++
									c.writeTo(sb)
//line /usr/local/go/src/text/template/parse/node.go:226
		// _ = "end of CoverTab[29225]"
	}
//line /usr/local/go/src/text/template/parse/node.go:227
	// _ = "end of CoverTab[29216]"
}

func (p *PipeNode) tree() *Tree {
//line /usr/local/go/src/text/template/parse/node.go:230
	_go_fuzz_dep_.CoverTab[29228]++
								return p.tr
//line /usr/local/go/src/text/template/parse/node.go:231
	// _ = "end of CoverTab[29228]"
}

func (p *PipeNode) CopyPipe() *PipeNode {
//line /usr/local/go/src/text/template/parse/node.go:234
	_go_fuzz_dep_.CoverTab[29229]++
								if p == nil {
//line /usr/local/go/src/text/template/parse/node.go:235
		_go_fuzz_dep_.CoverTab[29233]++
									return p
//line /usr/local/go/src/text/template/parse/node.go:236
		// _ = "end of CoverTab[29233]"
	} else {
//line /usr/local/go/src/text/template/parse/node.go:237
		_go_fuzz_dep_.CoverTab[29234]++
//line /usr/local/go/src/text/template/parse/node.go:237
		// _ = "end of CoverTab[29234]"
//line /usr/local/go/src/text/template/parse/node.go:237
	}
//line /usr/local/go/src/text/template/parse/node.go:237
	// _ = "end of CoverTab[29229]"
//line /usr/local/go/src/text/template/parse/node.go:237
	_go_fuzz_dep_.CoverTab[29230]++
								vars := make([]*VariableNode, len(p.Decl))
								for i, d := range p.Decl {
//line /usr/local/go/src/text/template/parse/node.go:239
		_go_fuzz_dep_.CoverTab[29235]++
									vars[i] = d.Copy().(*VariableNode)
//line /usr/local/go/src/text/template/parse/node.go:240
		// _ = "end of CoverTab[29235]"
	}
//line /usr/local/go/src/text/template/parse/node.go:241
	// _ = "end of CoverTab[29230]"
//line /usr/local/go/src/text/template/parse/node.go:241
	_go_fuzz_dep_.CoverTab[29231]++
								n := p.tr.newPipeline(p.Pos, p.Line, vars)
								n.IsAssign = p.IsAssign
								for _, c := range p.Cmds {
//line /usr/local/go/src/text/template/parse/node.go:244
		_go_fuzz_dep_.CoverTab[29236]++
									n.append(c.Copy().(*CommandNode))
//line /usr/local/go/src/text/template/parse/node.go:245
		// _ = "end of CoverTab[29236]"
	}
//line /usr/local/go/src/text/template/parse/node.go:246
	// _ = "end of CoverTab[29231]"
//line /usr/local/go/src/text/template/parse/node.go:246
	_go_fuzz_dep_.CoverTab[29232]++
								return n
//line /usr/local/go/src/text/template/parse/node.go:247
	// _ = "end of CoverTab[29232]"
}

func (p *PipeNode) Copy() Node {
//line /usr/local/go/src/text/template/parse/node.go:250
	_go_fuzz_dep_.CoverTab[29237]++
								return p.CopyPipe()
//line /usr/local/go/src/text/template/parse/node.go:251
	// _ = "end of CoverTab[29237]"
}

// ActionNode holds an action (something bounded by delimiters).
//line /usr/local/go/src/text/template/parse/node.go:254
// Control actions have their own nodes; ActionNode represents simple
//line /usr/local/go/src/text/template/parse/node.go:254
// ones such as field evaluations and parenthesized pipelines.
//line /usr/local/go/src/text/template/parse/node.go:257
type ActionNode struct {
	NodeType
	Pos
	tr	*Tree
	Line	int		// The line number in the input. Deprecated: Kept for compatibility.
	Pipe	*PipeNode	// The pipeline in the action.
}

func (t *Tree) newAction(pos Pos, line int, pipe *PipeNode) *ActionNode {
//line /usr/local/go/src/text/template/parse/node.go:265
	_go_fuzz_dep_.CoverTab[29238]++
								return &ActionNode{tr: t, NodeType: NodeAction, Pos: pos, Line: line, Pipe: pipe}
//line /usr/local/go/src/text/template/parse/node.go:266
	// _ = "end of CoverTab[29238]"
}

func (a *ActionNode) String() string {
//line /usr/local/go/src/text/template/parse/node.go:269
	_go_fuzz_dep_.CoverTab[29239]++
								var sb strings.Builder
								a.writeTo(&sb)
								return sb.String()
//line /usr/local/go/src/text/template/parse/node.go:272
	// _ = "end of CoverTab[29239]"
}

func (a *ActionNode) writeTo(sb *strings.Builder) {
//line /usr/local/go/src/text/template/parse/node.go:275
	_go_fuzz_dep_.CoverTab[29240]++
								sb.WriteString("{{")
								a.Pipe.writeTo(sb)
								sb.WriteString("}}")
//line /usr/local/go/src/text/template/parse/node.go:278
	// _ = "end of CoverTab[29240]"
}

func (a *ActionNode) tree() *Tree {
//line /usr/local/go/src/text/template/parse/node.go:281
	_go_fuzz_dep_.CoverTab[29241]++
								return a.tr
//line /usr/local/go/src/text/template/parse/node.go:282
	// _ = "end of CoverTab[29241]"
}

func (a *ActionNode) Copy() Node {
//line /usr/local/go/src/text/template/parse/node.go:285
	_go_fuzz_dep_.CoverTab[29242]++
								return a.tr.newAction(a.Pos, a.Line, a.Pipe.CopyPipe())
//line /usr/local/go/src/text/template/parse/node.go:286
	// _ = "end of CoverTab[29242]"

}

// CommandNode holds a command (a pipeline inside an evaluating action).
type CommandNode struct {
	NodeType
	Pos
	tr	*Tree
	Args	[]Node	// Arguments in lexical order: Identifier, field, or constant.
}

func (t *Tree) newCommand(pos Pos) *CommandNode {
//line /usr/local/go/src/text/template/parse/node.go:298
	_go_fuzz_dep_.CoverTab[29243]++
								return &CommandNode{tr: t, NodeType: NodeCommand, Pos: pos}
//line /usr/local/go/src/text/template/parse/node.go:299
	// _ = "end of CoverTab[29243]"
}

func (c *CommandNode) append(arg Node) {
//line /usr/local/go/src/text/template/parse/node.go:302
	_go_fuzz_dep_.CoverTab[29244]++
								c.Args = append(c.Args, arg)
//line /usr/local/go/src/text/template/parse/node.go:303
	// _ = "end of CoverTab[29244]"
}

func (c *CommandNode) String() string {
//line /usr/local/go/src/text/template/parse/node.go:306
	_go_fuzz_dep_.CoverTab[29245]++
								var sb strings.Builder
								c.writeTo(&sb)
								return sb.String()
//line /usr/local/go/src/text/template/parse/node.go:309
	// _ = "end of CoverTab[29245]"
}

func (c *CommandNode) writeTo(sb *strings.Builder) {
//line /usr/local/go/src/text/template/parse/node.go:312
	_go_fuzz_dep_.CoverTab[29246]++
								for i, arg := range c.Args {
//line /usr/local/go/src/text/template/parse/node.go:313
		_go_fuzz_dep_.CoverTab[29247]++
									if i > 0 {
//line /usr/local/go/src/text/template/parse/node.go:314
			_go_fuzz_dep_.CoverTab[29250]++
										sb.WriteByte(' ')
//line /usr/local/go/src/text/template/parse/node.go:315
			// _ = "end of CoverTab[29250]"
		} else {
//line /usr/local/go/src/text/template/parse/node.go:316
			_go_fuzz_dep_.CoverTab[29251]++
//line /usr/local/go/src/text/template/parse/node.go:316
			// _ = "end of CoverTab[29251]"
//line /usr/local/go/src/text/template/parse/node.go:316
		}
//line /usr/local/go/src/text/template/parse/node.go:316
		// _ = "end of CoverTab[29247]"
//line /usr/local/go/src/text/template/parse/node.go:316
		_go_fuzz_dep_.CoverTab[29248]++
									if arg, ok := arg.(*PipeNode); ok {
//line /usr/local/go/src/text/template/parse/node.go:317
			_go_fuzz_dep_.CoverTab[29252]++
										sb.WriteByte('(')
										arg.writeTo(sb)
										sb.WriteByte(')')
										continue
//line /usr/local/go/src/text/template/parse/node.go:321
			// _ = "end of CoverTab[29252]"
		} else {
//line /usr/local/go/src/text/template/parse/node.go:322
			_go_fuzz_dep_.CoverTab[29253]++
//line /usr/local/go/src/text/template/parse/node.go:322
			// _ = "end of CoverTab[29253]"
//line /usr/local/go/src/text/template/parse/node.go:322
		}
//line /usr/local/go/src/text/template/parse/node.go:322
		// _ = "end of CoverTab[29248]"
//line /usr/local/go/src/text/template/parse/node.go:322
		_go_fuzz_dep_.CoverTab[29249]++
									arg.writeTo(sb)
//line /usr/local/go/src/text/template/parse/node.go:323
		// _ = "end of CoverTab[29249]"
	}
//line /usr/local/go/src/text/template/parse/node.go:324
	// _ = "end of CoverTab[29246]"
}

func (c *CommandNode) tree() *Tree {
//line /usr/local/go/src/text/template/parse/node.go:327
	_go_fuzz_dep_.CoverTab[29254]++
								return c.tr
//line /usr/local/go/src/text/template/parse/node.go:328
	// _ = "end of CoverTab[29254]"
}

func (c *CommandNode) Copy() Node {
//line /usr/local/go/src/text/template/parse/node.go:331
	_go_fuzz_dep_.CoverTab[29255]++
								if c == nil {
//line /usr/local/go/src/text/template/parse/node.go:332
		_go_fuzz_dep_.CoverTab[29258]++
									return c
//line /usr/local/go/src/text/template/parse/node.go:333
		// _ = "end of CoverTab[29258]"
	} else {
//line /usr/local/go/src/text/template/parse/node.go:334
		_go_fuzz_dep_.CoverTab[29259]++
//line /usr/local/go/src/text/template/parse/node.go:334
		// _ = "end of CoverTab[29259]"
//line /usr/local/go/src/text/template/parse/node.go:334
	}
//line /usr/local/go/src/text/template/parse/node.go:334
	// _ = "end of CoverTab[29255]"
//line /usr/local/go/src/text/template/parse/node.go:334
	_go_fuzz_dep_.CoverTab[29256]++
								n := c.tr.newCommand(c.Pos)
								for _, c := range c.Args {
//line /usr/local/go/src/text/template/parse/node.go:336
		_go_fuzz_dep_.CoverTab[29260]++
									n.append(c.Copy())
//line /usr/local/go/src/text/template/parse/node.go:337
		// _ = "end of CoverTab[29260]"
	}
//line /usr/local/go/src/text/template/parse/node.go:338
	// _ = "end of CoverTab[29256]"
//line /usr/local/go/src/text/template/parse/node.go:338
	_go_fuzz_dep_.CoverTab[29257]++
								return n
//line /usr/local/go/src/text/template/parse/node.go:339
	// _ = "end of CoverTab[29257]"
}

// IdentifierNode holds an identifier.
type IdentifierNode struct {
	NodeType
	Pos
	tr	*Tree
	Ident	string	// The identifier's name.
}

// NewIdentifier returns a new IdentifierNode with the given identifier name.
func NewIdentifier(ident string) *IdentifierNode {
//line /usr/local/go/src/text/template/parse/node.go:351
	_go_fuzz_dep_.CoverTab[29261]++
								return &IdentifierNode{NodeType: NodeIdentifier, Ident: ident}
//line /usr/local/go/src/text/template/parse/node.go:352
	// _ = "end of CoverTab[29261]"
}

// SetPos sets the position. NewIdentifier is a public method so we can't modify its signature.
//line /usr/local/go/src/text/template/parse/node.go:355
// Chained for convenience.
//line /usr/local/go/src/text/template/parse/node.go:355
// TODO: fix one day?
//line /usr/local/go/src/text/template/parse/node.go:358
func (i *IdentifierNode) SetPos(pos Pos) *IdentifierNode {
//line /usr/local/go/src/text/template/parse/node.go:358
	_go_fuzz_dep_.CoverTab[29262]++
								i.Pos = pos
								return i
//line /usr/local/go/src/text/template/parse/node.go:360
	// _ = "end of CoverTab[29262]"
}

// SetTree sets the parent tree for the node. NewIdentifier is a public method so we can't modify its signature.
//line /usr/local/go/src/text/template/parse/node.go:363
// Chained for convenience.
//line /usr/local/go/src/text/template/parse/node.go:363
// TODO: fix one day?
//line /usr/local/go/src/text/template/parse/node.go:366
func (i *IdentifierNode) SetTree(t *Tree) *IdentifierNode {
//line /usr/local/go/src/text/template/parse/node.go:366
	_go_fuzz_dep_.CoverTab[29263]++
								i.tr = t
								return i
//line /usr/local/go/src/text/template/parse/node.go:368
	// _ = "end of CoverTab[29263]"
}

func (i *IdentifierNode) String() string {
//line /usr/local/go/src/text/template/parse/node.go:371
	_go_fuzz_dep_.CoverTab[29264]++
								return i.Ident
//line /usr/local/go/src/text/template/parse/node.go:372
	// _ = "end of CoverTab[29264]"
}

func (i *IdentifierNode) writeTo(sb *strings.Builder) {
//line /usr/local/go/src/text/template/parse/node.go:375
	_go_fuzz_dep_.CoverTab[29265]++
								sb.WriteString(i.String())
//line /usr/local/go/src/text/template/parse/node.go:376
	// _ = "end of CoverTab[29265]"
}

func (i *IdentifierNode) tree() *Tree {
//line /usr/local/go/src/text/template/parse/node.go:379
	_go_fuzz_dep_.CoverTab[29266]++
								return i.tr
//line /usr/local/go/src/text/template/parse/node.go:380
	// _ = "end of CoverTab[29266]"
}

func (i *IdentifierNode) Copy() Node {
//line /usr/local/go/src/text/template/parse/node.go:383
	_go_fuzz_dep_.CoverTab[29267]++
								return NewIdentifier(i.Ident).SetTree(i.tr).SetPos(i.Pos)
//line /usr/local/go/src/text/template/parse/node.go:384
	// _ = "end of CoverTab[29267]"
}

// VariableNode holds a list of variable names, possibly with chained field
//line /usr/local/go/src/text/template/parse/node.go:387
// accesses. The dollar sign is part of the (first) name.
//line /usr/local/go/src/text/template/parse/node.go:389
type VariableNode struct {
	NodeType
	Pos
	tr	*Tree
	Ident	[]string	// Variable name and fields in lexical order.
}

func (t *Tree) newVariable(pos Pos, ident string) *VariableNode {
//line /usr/local/go/src/text/template/parse/node.go:396
	_go_fuzz_dep_.CoverTab[29268]++
								return &VariableNode{tr: t, NodeType: NodeVariable, Pos: pos, Ident: strings.Split(ident, ".")}
//line /usr/local/go/src/text/template/parse/node.go:397
	// _ = "end of CoverTab[29268]"
}

func (v *VariableNode) String() string {
//line /usr/local/go/src/text/template/parse/node.go:400
	_go_fuzz_dep_.CoverTab[29269]++
								var sb strings.Builder
								v.writeTo(&sb)
								return sb.String()
//line /usr/local/go/src/text/template/parse/node.go:403
	// _ = "end of CoverTab[29269]"
}

func (v *VariableNode) writeTo(sb *strings.Builder) {
//line /usr/local/go/src/text/template/parse/node.go:406
	_go_fuzz_dep_.CoverTab[29270]++
								for i, id := range v.Ident {
//line /usr/local/go/src/text/template/parse/node.go:407
		_go_fuzz_dep_.CoverTab[29271]++
									if i > 0 {
//line /usr/local/go/src/text/template/parse/node.go:408
			_go_fuzz_dep_.CoverTab[29273]++
										sb.WriteByte('.')
//line /usr/local/go/src/text/template/parse/node.go:409
			// _ = "end of CoverTab[29273]"
		} else {
//line /usr/local/go/src/text/template/parse/node.go:410
			_go_fuzz_dep_.CoverTab[29274]++
//line /usr/local/go/src/text/template/parse/node.go:410
			// _ = "end of CoverTab[29274]"
//line /usr/local/go/src/text/template/parse/node.go:410
		}
//line /usr/local/go/src/text/template/parse/node.go:410
		// _ = "end of CoverTab[29271]"
//line /usr/local/go/src/text/template/parse/node.go:410
		_go_fuzz_dep_.CoverTab[29272]++
									sb.WriteString(id)
//line /usr/local/go/src/text/template/parse/node.go:411
		// _ = "end of CoverTab[29272]"
	}
//line /usr/local/go/src/text/template/parse/node.go:412
	// _ = "end of CoverTab[29270]"
}

func (v *VariableNode) tree() *Tree {
//line /usr/local/go/src/text/template/parse/node.go:415
	_go_fuzz_dep_.CoverTab[29275]++
								return v.tr
//line /usr/local/go/src/text/template/parse/node.go:416
	// _ = "end of CoverTab[29275]"
}

func (v *VariableNode) Copy() Node {
//line /usr/local/go/src/text/template/parse/node.go:419
	_go_fuzz_dep_.CoverTab[29276]++
								return &VariableNode{tr: v.tr, NodeType: NodeVariable, Pos: v.Pos, Ident: append([]string{}, v.Ident...)}
//line /usr/local/go/src/text/template/parse/node.go:420
	// _ = "end of CoverTab[29276]"
}

// DotNode holds the special identifier '.'.
type DotNode struct {
	NodeType
	Pos
	tr	*Tree
}

func (t *Tree) newDot(pos Pos) *DotNode {
//line /usr/local/go/src/text/template/parse/node.go:430
	_go_fuzz_dep_.CoverTab[29277]++
								return &DotNode{tr: t, NodeType: NodeDot, Pos: pos}
//line /usr/local/go/src/text/template/parse/node.go:431
	// _ = "end of CoverTab[29277]"
}

func (d *DotNode) Type() NodeType {
//line /usr/local/go/src/text/template/parse/node.go:434
	_go_fuzz_dep_.CoverTab[29278]++

//line /usr/local/go/src/text/template/parse/node.go:438
	return NodeDot
//line /usr/local/go/src/text/template/parse/node.go:438
	// _ = "end of CoverTab[29278]"
}

func (d *DotNode) String() string {
//line /usr/local/go/src/text/template/parse/node.go:441
	_go_fuzz_dep_.CoverTab[29279]++
								return "."
//line /usr/local/go/src/text/template/parse/node.go:442
	// _ = "end of CoverTab[29279]"
}

func (d *DotNode) writeTo(sb *strings.Builder) {
//line /usr/local/go/src/text/template/parse/node.go:445
	_go_fuzz_dep_.CoverTab[29280]++
								sb.WriteString(d.String())
//line /usr/local/go/src/text/template/parse/node.go:446
	// _ = "end of CoverTab[29280]"
}

func (d *DotNode) tree() *Tree {
//line /usr/local/go/src/text/template/parse/node.go:449
	_go_fuzz_dep_.CoverTab[29281]++
								return d.tr
//line /usr/local/go/src/text/template/parse/node.go:450
	// _ = "end of CoverTab[29281]"
}

func (d *DotNode) Copy() Node {
//line /usr/local/go/src/text/template/parse/node.go:453
	_go_fuzz_dep_.CoverTab[29282]++
								return d.tr.newDot(d.Pos)
//line /usr/local/go/src/text/template/parse/node.go:454
	// _ = "end of CoverTab[29282]"
}

// NilNode holds the special identifier 'nil' representing an untyped nil constant.
type NilNode struct {
	NodeType
	Pos
	tr	*Tree
}

func (t *Tree) newNil(pos Pos) *NilNode {
//line /usr/local/go/src/text/template/parse/node.go:464
	_go_fuzz_dep_.CoverTab[29283]++
								return &NilNode{tr: t, NodeType: NodeNil, Pos: pos}
//line /usr/local/go/src/text/template/parse/node.go:465
	// _ = "end of CoverTab[29283]"
}

func (n *NilNode) Type() NodeType {
//line /usr/local/go/src/text/template/parse/node.go:468
	_go_fuzz_dep_.CoverTab[29284]++

//line /usr/local/go/src/text/template/parse/node.go:472
	return NodeNil
//line /usr/local/go/src/text/template/parse/node.go:472
	// _ = "end of CoverTab[29284]"
}

func (n *NilNode) String() string {
//line /usr/local/go/src/text/template/parse/node.go:475
	_go_fuzz_dep_.CoverTab[29285]++
								return "nil"
//line /usr/local/go/src/text/template/parse/node.go:476
	// _ = "end of CoverTab[29285]"
}

func (n *NilNode) writeTo(sb *strings.Builder) {
//line /usr/local/go/src/text/template/parse/node.go:479
	_go_fuzz_dep_.CoverTab[29286]++
								sb.WriteString(n.String())
//line /usr/local/go/src/text/template/parse/node.go:480
	// _ = "end of CoverTab[29286]"
}

func (n *NilNode) tree() *Tree {
//line /usr/local/go/src/text/template/parse/node.go:483
	_go_fuzz_dep_.CoverTab[29287]++
								return n.tr
//line /usr/local/go/src/text/template/parse/node.go:484
	// _ = "end of CoverTab[29287]"
}

func (n *NilNode) Copy() Node {
//line /usr/local/go/src/text/template/parse/node.go:487
	_go_fuzz_dep_.CoverTab[29288]++
								return n.tr.newNil(n.Pos)
//line /usr/local/go/src/text/template/parse/node.go:488
	// _ = "end of CoverTab[29288]"
}

// FieldNode holds a field (identifier starting with '.').
//line /usr/local/go/src/text/template/parse/node.go:491
// The names may be chained ('.x.y').
//line /usr/local/go/src/text/template/parse/node.go:491
// The period is dropped from each ident.
//line /usr/local/go/src/text/template/parse/node.go:494
type FieldNode struct {
	NodeType
	Pos
	tr	*Tree
	Ident	[]string	// The identifiers in lexical order.
}

func (t *Tree) newField(pos Pos, ident string) *FieldNode {
//line /usr/local/go/src/text/template/parse/node.go:501
	_go_fuzz_dep_.CoverTab[29289]++
								return &FieldNode{tr: t, NodeType: NodeField, Pos: pos, Ident: strings.Split(ident[1:], ".")}
//line /usr/local/go/src/text/template/parse/node.go:502
	// _ = "end of CoverTab[29289]"
}

func (f *FieldNode) String() string {
//line /usr/local/go/src/text/template/parse/node.go:505
	_go_fuzz_dep_.CoverTab[29290]++
								var sb strings.Builder
								f.writeTo(&sb)
								return sb.String()
//line /usr/local/go/src/text/template/parse/node.go:508
	// _ = "end of CoverTab[29290]"
}

func (f *FieldNode) writeTo(sb *strings.Builder) {
//line /usr/local/go/src/text/template/parse/node.go:511
	_go_fuzz_dep_.CoverTab[29291]++
								for _, id := range f.Ident {
//line /usr/local/go/src/text/template/parse/node.go:512
		_go_fuzz_dep_.CoverTab[29292]++
									sb.WriteByte('.')
									sb.WriteString(id)
//line /usr/local/go/src/text/template/parse/node.go:514
		// _ = "end of CoverTab[29292]"
	}
//line /usr/local/go/src/text/template/parse/node.go:515
	// _ = "end of CoverTab[29291]"
}

func (f *FieldNode) tree() *Tree {
//line /usr/local/go/src/text/template/parse/node.go:518
	_go_fuzz_dep_.CoverTab[29293]++
								return f.tr
//line /usr/local/go/src/text/template/parse/node.go:519
	// _ = "end of CoverTab[29293]"
}

func (f *FieldNode) Copy() Node {
//line /usr/local/go/src/text/template/parse/node.go:522
	_go_fuzz_dep_.CoverTab[29294]++
								return &FieldNode{tr: f.tr, NodeType: NodeField, Pos: f.Pos, Ident: append([]string{}, f.Ident...)}
//line /usr/local/go/src/text/template/parse/node.go:523
	// _ = "end of CoverTab[29294]"
}

// ChainNode holds a term followed by a chain of field accesses (identifier starting with '.').
//line /usr/local/go/src/text/template/parse/node.go:526
// The names may be chained ('.x.y').
//line /usr/local/go/src/text/template/parse/node.go:526
// The periods are dropped from each ident.
//line /usr/local/go/src/text/template/parse/node.go:529
type ChainNode struct {
	NodeType
	Pos
	tr	*Tree
	Node	Node
	Field	[]string	// The identifiers in lexical order.
}

func (t *Tree) newChain(pos Pos, node Node) *ChainNode {
//line /usr/local/go/src/text/template/parse/node.go:537
	_go_fuzz_dep_.CoverTab[29295]++
								return &ChainNode{tr: t, NodeType: NodeChain, Pos: pos, Node: node}
//line /usr/local/go/src/text/template/parse/node.go:538
	// _ = "end of CoverTab[29295]"
}

// Add adds the named field (which should start with a period) to the end of the chain.
func (c *ChainNode) Add(field string) {
//line /usr/local/go/src/text/template/parse/node.go:542
	_go_fuzz_dep_.CoverTab[29296]++
								if len(field) == 0 || func() bool {
//line /usr/local/go/src/text/template/parse/node.go:543
		_go_fuzz_dep_.CoverTab[29299]++
//line /usr/local/go/src/text/template/parse/node.go:543
		return field[0] != '.'
//line /usr/local/go/src/text/template/parse/node.go:543
		// _ = "end of CoverTab[29299]"
//line /usr/local/go/src/text/template/parse/node.go:543
	}() {
//line /usr/local/go/src/text/template/parse/node.go:543
		_go_fuzz_dep_.CoverTab[29300]++
									panic("no dot in field")
//line /usr/local/go/src/text/template/parse/node.go:544
		// _ = "end of CoverTab[29300]"
	} else {
//line /usr/local/go/src/text/template/parse/node.go:545
		_go_fuzz_dep_.CoverTab[29301]++
//line /usr/local/go/src/text/template/parse/node.go:545
		// _ = "end of CoverTab[29301]"
//line /usr/local/go/src/text/template/parse/node.go:545
	}
//line /usr/local/go/src/text/template/parse/node.go:545
	// _ = "end of CoverTab[29296]"
//line /usr/local/go/src/text/template/parse/node.go:545
	_go_fuzz_dep_.CoverTab[29297]++
								field = field[1:]
								if field == "" {
//line /usr/local/go/src/text/template/parse/node.go:547
		_go_fuzz_dep_.CoverTab[29302]++
									panic("empty field")
//line /usr/local/go/src/text/template/parse/node.go:548
		// _ = "end of CoverTab[29302]"
	} else {
//line /usr/local/go/src/text/template/parse/node.go:549
		_go_fuzz_dep_.CoverTab[29303]++
//line /usr/local/go/src/text/template/parse/node.go:549
		// _ = "end of CoverTab[29303]"
//line /usr/local/go/src/text/template/parse/node.go:549
	}
//line /usr/local/go/src/text/template/parse/node.go:549
	// _ = "end of CoverTab[29297]"
//line /usr/local/go/src/text/template/parse/node.go:549
	_go_fuzz_dep_.CoverTab[29298]++
								c.Field = append(c.Field, field)
//line /usr/local/go/src/text/template/parse/node.go:550
	// _ = "end of CoverTab[29298]"
}

func (c *ChainNode) String() string {
//line /usr/local/go/src/text/template/parse/node.go:553
	_go_fuzz_dep_.CoverTab[29304]++
								var sb strings.Builder
								c.writeTo(&sb)
								return sb.String()
//line /usr/local/go/src/text/template/parse/node.go:556
	// _ = "end of CoverTab[29304]"
}

func (c *ChainNode) writeTo(sb *strings.Builder) {
//line /usr/local/go/src/text/template/parse/node.go:559
	_go_fuzz_dep_.CoverTab[29305]++
								if _, ok := c.Node.(*PipeNode); ok {
//line /usr/local/go/src/text/template/parse/node.go:560
		_go_fuzz_dep_.CoverTab[29307]++
									sb.WriteByte('(')
									c.Node.writeTo(sb)
									sb.WriteByte(')')
//line /usr/local/go/src/text/template/parse/node.go:563
		// _ = "end of CoverTab[29307]"
	} else {
//line /usr/local/go/src/text/template/parse/node.go:564
		_go_fuzz_dep_.CoverTab[29308]++
									c.Node.writeTo(sb)
//line /usr/local/go/src/text/template/parse/node.go:565
		// _ = "end of CoverTab[29308]"
	}
//line /usr/local/go/src/text/template/parse/node.go:566
	// _ = "end of CoverTab[29305]"
//line /usr/local/go/src/text/template/parse/node.go:566
	_go_fuzz_dep_.CoverTab[29306]++
								for _, field := range c.Field {
//line /usr/local/go/src/text/template/parse/node.go:567
		_go_fuzz_dep_.CoverTab[29309]++
									sb.WriteByte('.')
									sb.WriteString(field)
//line /usr/local/go/src/text/template/parse/node.go:569
		// _ = "end of CoverTab[29309]"
	}
//line /usr/local/go/src/text/template/parse/node.go:570
	// _ = "end of CoverTab[29306]"
}

func (c *ChainNode) tree() *Tree {
//line /usr/local/go/src/text/template/parse/node.go:573
	_go_fuzz_dep_.CoverTab[29310]++
								return c.tr
//line /usr/local/go/src/text/template/parse/node.go:574
	// _ = "end of CoverTab[29310]"
}

func (c *ChainNode) Copy() Node {
//line /usr/local/go/src/text/template/parse/node.go:577
	_go_fuzz_dep_.CoverTab[29311]++
								return &ChainNode{tr: c.tr, NodeType: NodeChain, Pos: c.Pos, Node: c.Node, Field: append([]string{}, c.Field...)}
//line /usr/local/go/src/text/template/parse/node.go:578
	// _ = "end of CoverTab[29311]"
}

// BoolNode holds a boolean constant.
type BoolNode struct {
	NodeType
	Pos
	tr	*Tree
	True	bool	// The value of the boolean constant.
}

func (t *Tree) newBool(pos Pos, true bool) *BoolNode {
//line /usr/local/go/src/text/template/parse/node.go:589
	_go_fuzz_dep_.CoverTab[29312]++
								return &BoolNode{tr: t, NodeType: NodeBool, Pos: pos, True: true}
//line /usr/local/go/src/text/template/parse/node.go:590
	// _ = "end of CoverTab[29312]"
}

func (b *BoolNode) String() string {
//line /usr/local/go/src/text/template/parse/node.go:593
	_go_fuzz_dep_.CoverTab[29313]++
								if b.True {
//line /usr/local/go/src/text/template/parse/node.go:594
		_go_fuzz_dep_.CoverTab[29315]++
									return "true"
//line /usr/local/go/src/text/template/parse/node.go:595
		// _ = "end of CoverTab[29315]"
	} else {
//line /usr/local/go/src/text/template/parse/node.go:596
		_go_fuzz_dep_.CoverTab[29316]++
//line /usr/local/go/src/text/template/parse/node.go:596
		// _ = "end of CoverTab[29316]"
//line /usr/local/go/src/text/template/parse/node.go:596
	}
//line /usr/local/go/src/text/template/parse/node.go:596
	// _ = "end of CoverTab[29313]"
//line /usr/local/go/src/text/template/parse/node.go:596
	_go_fuzz_dep_.CoverTab[29314]++
								return "false"
//line /usr/local/go/src/text/template/parse/node.go:597
	// _ = "end of CoverTab[29314]"
}

func (b *BoolNode) writeTo(sb *strings.Builder) {
//line /usr/local/go/src/text/template/parse/node.go:600
	_go_fuzz_dep_.CoverTab[29317]++
								sb.WriteString(b.String())
//line /usr/local/go/src/text/template/parse/node.go:601
	// _ = "end of CoverTab[29317]"
}

func (b *BoolNode) tree() *Tree {
//line /usr/local/go/src/text/template/parse/node.go:604
	_go_fuzz_dep_.CoverTab[29318]++
								return b.tr
//line /usr/local/go/src/text/template/parse/node.go:605
	// _ = "end of CoverTab[29318]"
}

func (b *BoolNode) Copy() Node {
//line /usr/local/go/src/text/template/parse/node.go:608
	_go_fuzz_dep_.CoverTab[29319]++
								return b.tr.newBool(b.Pos, b.True)
//line /usr/local/go/src/text/template/parse/node.go:609
	// _ = "end of CoverTab[29319]"
}

// NumberNode holds a number: signed or unsigned integer, float, or complex.
//line /usr/local/go/src/text/template/parse/node.go:612
// The value is parsed and stored under all the types that can represent the value.
//line /usr/local/go/src/text/template/parse/node.go:612
// This simulates in a small amount of code the behavior of Go's ideal constants.
//line /usr/local/go/src/text/template/parse/node.go:615
type NumberNode struct {
	NodeType
	Pos
	tr		*Tree
	IsInt		bool		// Number has an integral value.
	IsUint		bool		// Number has an unsigned integral value.
	IsFloat		bool		// Number has a floating-point value.
	IsComplex	bool		// Number is complex.
	Int64		int64		// The signed integer value.
	Uint64		uint64		// The unsigned integer value.
	Float64		float64		// The floating-point value.
	Complex128	complex128	// The complex value.
	Text		string		// The original textual representation from the input.
}

func (t *Tree) newNumber(pos Pos, text string, typ itemType) (*NumberNode, error) {
//line /usr/local/go/src/text/template/parse/node.go:630
	_go_fuzz_dep_.CoverTab[29320]++
								n := &NumberNode{tr: t, NodeType: NodeNumber, Pos: pos, Text: text}
								switch typ {
	case itemCharConstant:
//line /usr/local/go/src/text/template/parse/node.go:633
		_go_fuzz_dep_.CoverTab[29327]++
									rune, _, tail, err := strconv.UnquoteChar(text[1:], text[0])
									if err != nil {
//line /usr/local/go/src/text/template/parse/node.go:635
			_go_fuzz_dep_.CoverTab[29333]++
										return nil, err
//line /usr/local/go/src/text/template/parse/node.go:636
			// _ = "end of CoverTab[29333]"
		} else {
//line /usr/local/go/src/text/template/parse/node.go:637
			_go_fuzz_dep_.CoverTab[29334]++
//line /usr/local/go/src/text/template/parse/node.go:637
			// _ = "end of CoverTab[29334]"
//line /usr/local/go/src/text/template/parse/node.go:637
		}
//line /usr/local/go/src/text/template/parse/node.go:637
		// _ = "end of CoverTab[29327]"
//line /usr/local/go/src/text/template/parse/node.go:637
		_go_fuzz_dep_.CoverTab[29328]++
									if tail != "'" {
//line /usr/local/go/src/text/template/parse/node.go:638
			_go_fuzz_dep_.CoverTab[29335]++
										return nil, fmt.Errorf("malformed character constant: %s", text)
//line /usr/local/go/src/text/template/parse/node.go:639
			// _ = "end of CoverTab[29335]"
		} else {
//line /usr/local/go/src/text/template/parse/node.go:640
			_go_fuzz_dep_.CoverTab[29336]++
//line /usr/local/go/src/text/template/parse/node.go:640
			// _ = "end of CoverTab[29336]"
//line /usr/local/go/src/text/template/parse/node.go:640
		}
//line /usr/local/go/src/text/template/parse/node.go:640
		// _ = "end of CoverTab[29328]"
//line /usr/local/go/src/text/template/parse/node.go:640
		_go_fuzz_dep_.CoverTab[29329]++
									n.Int64 = int64(rune)
									n.IsInt = true
									n.Uint64 = uint64(rune)
									n.IsUint = true
									n.Float64 = float64(rune)
									n.IsFloat = true
									return n, nil
//line /usr/local/go/src/text/template/parse/node.go:647
		// _ = "end of CoverTab[29329]"
	case itemComplex:
//line /usr/local/go/src/text/template/parse/node.go:648
		_go_fuzz_dep_.CoverTab[29330]++

									if _, err := fmt.Sscan(text, &n.Complex128); err != nil {
//line /usr/local/go/src/text/template/parse/node.go:650
			_go_fuzz_dep_.CoverTab[29337]++
										return nil, err
//line /usr/local/go/src/text/template/parse/node.go:651
			// _ = "end of CoverTab[29337]"
		} else {
//line /usr/local/go/src/text/template/parse/node.go:652
			_go_fuzz_dep_.CoverTab[29338]++
//line /usr/local/go/src/text/template/parse/node.go:652
			// _ = "end of CoverTab[29338]"
//line /usr/local/go/src/text/template/parse/node.go:652
		}
//line /usr/local/go/src/text/template/parse/node.go:652
		// _ = "end of CoverTab[29330]"
//line /usr/local/go/src/text/template/parse/node.go:652
		_go_fuzz_dep_.CoverTab[29331]++
									n.IsComplex = true
									n.simplifyComplex()
									return n, nil
//line /usr/local/go/src/text/template/parse/node.go:655
		// _ = "end of CoverTab[29331]"
//line /usr/local/go/src/text/template/parse/node.go:655
	default:
//line /usr/local/go/src/text/template/parse/node.go:655
		_go_fuzz_dep_.CoverTab[29332]++
//line /usr/local/go/src/text/template/parse/node.go:655
		// _ = "end of CoverTab[29332]"
	}
//line /usr/local/go/src/text/template/parse/node.go:656
	// _ = "end of CoverTab[29320]"
//line /usr/local/go/src/text/template/parse/node.go:656
	_go_fuzz_dep_.CoverTab[29321]++

								if len(text) > 0 && func() bool {
//line /usr/local/go/src/text/template/parse/node.go:658
		_go_fuzz_dep_.CoverTab[29339]++
//line /usr/local/go/src/text/template/parse/node.go:658
		return text[len(text)-1] == 'i'
//line /usr/local/go/src/text/template/parse/node.go:658
		// _ = "end of CoverTab[29339]"
//line /usr/local/go/src/text/template/parse/node.go:658
	}() {
//line /usr/local/go/src/text/template/parse/node.go:658
		_go_fuzz_dep_.CoverTab[29340]++
									f, err := strconv.ParseFloat(text[:len(text)-1], 64)
									if err == nil {
//line /usr/local/go/src/text/template/parse/node.go:660
			_go_fuzz_dep_.CoverTab[29341]++
										n.IsComplex = true
										n.Complex128 = complex(0, f)
										n.simplifyComplex()
										return n, nil
//line /usr/local/go/src/text/template/parse/node.go:664
			// _ = "end of CoverTab[29341]"
		} else {
//line /usr/local/go/src/text/template/parse/node.go:665
			_go_fuzz_dep_.CoverTab[29342]++
//line /usr/local/go/src/text/template/parse/node.go:665
			// _ = "end of CoverTab[29342]"
//line /usr/local/go/src/text/template/parse/node.go:665
		}
//line /usr/local/go/src/text/template/parse/node.go:665
		// _ = "end of CoverTab[29340]"
	} else {
//line /usr/local/go/src/text/template/parse/node.go:666
		_go_fuzz_dep_.CoverTab[29343]++
//line /usr/local/go/src/text/template/parse/node.go:666
		// _ = "end of CoverTab[29343]"
//line /usr/local/go/src/text/template/parse/node.go:666
	}
//line /usr/local/go/src/text/template/parse/node.go:666
	// _ = "end of CoverTab[29321]"
//line /usr/local/go/src/text/template/parse/node.go:666
	_go_fuzz_dep_.CoverTab[29322]++

								u, err := strconv.ParseUint(text, 0, 64)
								if err == nil {
//line /usr/local/go/src/text/template/parse/node.go:669
		_go_fuzz_dep_.CoverTab[29344]++
									n.IsUint = true
									n.Uint64 = u
//line /usr/local/go/src/text/template/parse/node.go:671
		// _ = "end of CoverTab[29344]"
	} else {
//line /usr/local/go/src/text/template/parse/node.go:672
		_go_fuzz_dep_.CoverTab[29345]++
//line /usr/local/go/src/text/template/parse/node.go:672
		// _ = "end of CoverTab[29345]"
//line /usr/local/go/src/text/template/parse/node.go:672
	}
//line /usr/local/go/src/text/template/parse/node.go:672
	// _ = "end of CoverTab[29322]"
//line /usr/local/go/src/text/template/parse/node.go:672
	_go_fuzz_dep_.CoverTab[29323]++
								i, err := strconv.ParseInt(text, 0, 64)
								if err == nil {
//line /usr/local/go/src/text/template/parse/node.go:674
		_go_fuzz_dep_.CoverTab[29346]++
									n.IsInt = true
									n.Int64 = i
									if i == 0 {
//line /usr/local/go/src/text/template/parse/node.go:677
			_go_fuzz_dep_.CoverTab[29347]++
										n.IsUint = true
										n.Uint64 = u
//line /usr/local/go/src/text/template/parse/node.go:679
			// _ = "end of CoverTab[29347]"
		} else {
//line /usr/local/go/src/text/template/parse/node.go:680
			_go_fuzz_dep_.CoverTab[29348]++
//line /usr/local/go/src/text/template/parse/node.go:680
			// _ = "end of CoverTab[29348]"
//line /usr/local/go/src/text/template/parse/node.go:680
		}
//line /usr/local/go/src/text/template/parse/node.go:680
		// _ = "end of CoverTab[29346]"
	} else {
//line /usr/local/go/src/text/template/parse/node.go:681
		_go_fuzz_dep_.CoverTab[29349]++
//line /usr/local/go/src/text/template/parse/node.go:681
		// _ = "end of CoverTab[29349]"
//line /usr/local/go/src/text/template/parse/node.go:681
	}
//line /usr/local/go/src/text/template/parse/node.go:681
	// _ = "end of CoverTab[29323]"
//line /usr/local/go/src/text/template/parse/node.go:681
	_go_fuzz_dep_.CoverTab[29324]++

								if n.IsInt {
//line /usr/local/go/src/text/template/parse/node.go:683
		_go_fuzz_dep_.CoverTab[29350]++
									n.IsFloat = true
									n.Float64 = float64(n.Int64)
//line /usr/local/go/src/text/template/parse/node.go:685
		// _ = "end of CoverTab[29350]"
	} else {
//line /usr/local/go/src/text/template/parse/node.go:686
		_go_fuzz_dep_.CoverTab[29351]++
//line /usr/local/go/src/text/template/parse/node.go:686
		if n.IsUint {
//line /usr/local/go/src/text/template/parse/node.go:686
			_go_fuzz_dep_.CoverTab[29352]++
										n.IsFloat = true
										n.Float64 = float64(n.Uint64)
//line /usr/local/go/src/text/template/parse/node.go:688
			// _ = "end of CoverTab[29352]"
		} else {
//line /usr/local/go/src/text/template/parse/node.go:689
			_go_fuzz_dep_.CoverTab[29353]++
										f, err := strconv.ParseFloat(text, 64)
										if err == nil {
//line /usr/local/go/src/text/template/parse/node.go:691
				_go_fuzz_dep_.CoverTab[29354]++

//line /usr/local/go/src/text/template/parse/node.go:694
				if !strings.ContainsAny(text, ".eEpP") {
//line /usr/local/go/src/text/template/parse/node.go:694
					_go_fuzz_dep_.CoverTab[29357]++
												return nil, fmt.Errorf("integer overflow: %q", text)
//line /usr/local/go/src/text/template/parse/node.go:695
					// _ = "end of CoverTab[29357]"
				} else {
//line /usr/local/go/src/text/template/parse/node.go:696
					_go_fuzz_dep_.CoverTab[29358]++
//line /usr/local/go/src/text/template/parse/node.go:696
					// _ = "end of CoverTab[29358]"
//line /usr/local/go/src/text/template/parse/node.go:696
				}
//line /usr/local/go/src/text/template/parse/node.go:696
				// _ = "end of CoverTab[29354]"
//line /usr/local/go/src/text/template/parse/node.go:696
				_go_fuzz_dep_.CoverTab[29355]++
											n.IsFloat = true
											n.Float64 = f

											if !n.IsInt && func() bool {
//line /usr/local/go/src/text/template/parse/node.go:700
					_go_fuzz_dep_.CoverTab[29359]++
//line /usr/local/go/src/text/template/parse/node.go:700
					return float64(int64(f)) == f
//line /usr/local/go/src/text/template/parse/node.go:700
					// _ = "end of CoverTab[29359]"
//line /usr/local/go/src/text/template/parse/node.go:700
				}() {
//line /usr/local/go/src/text/template/parse/node.go:700
					_go_fuzz_dep_.CoverTab[29360]++
												n.IsInt = true
												n.Int64 = int64(f)
//line /usr/local/go/src/text/template/parse/node.go:702
					// _ = "end of CoverTab[29360]"
				} else {
//line /usr/local/go/src/text/template/parse/node.go:703
					_go_fuzz_dep_.CoverTab[29361]++
//line /usr/local/go/src/text/template/parse/node.go:703
					// _ = "end of CoverTab[29361]"
//line /usr/local/go/src/text/template/parse/node.go:703
				}
//line /usr/local/go/src/text/template/parse/node.go:703
				// _ = "end of CoverTab[29355]"
//line /usr/local/go/src/text/template/parse/node.go:703
				_go_fuzz_dep_.CoverTab[29356]++
											if !n.IsUint && func() bool {
//line /usr/local/go/src/text/template/parse/node.go:704
					_go_fuzz_dep_.CoverTab[29362]++
//line /usr/local/go/src/text/template/parse/node.go:704
					return float64(uint64(f)) == f
//line /usr/local/go/src/text/template/parse/node.go:704
					// _ = "end of CoverTab[29362]"
//line /usr/local/go/src/text/template/parse/node.go:704
				}() {
//line /usr/local/go/src/text/template/parse/node.go:704
					_go_fuzz_dep_.CoverTab[29363]++
												n.IsUint = true
												n.Uint64 = uint64(f)
//line /usr/local/go/src/text/template/parse/node.go:706
					// _ = "end of CoverTab[29363]"
				} else {
//line /usr/local/go/src/text/template/parse/node.go:707
					_go_fuzz_dep_.CoverTab[29364]++
//line /usr/local/go/src/text/template/parse/node.go:707
					// _ = "end of CoverTab[29364]"
//line /usr/local/go/src/text/template/parse/node.go:707
				}
//line /usr/local/go/src/text/template/parse/node.go:707
				// _ = "end of CoverTab[29356]"
			} else {
//line /usr/local/go/src/text/template/parse/node.go:708
				_go_fuzz_dep_.CoverTab[29365]++
//line /usr/local/go/src/text/template/parse/node.go:708
				// _ = "end of CoverTab[29365]"
//line /usr/local/go/src/text/template/parse/node.go:708
			}
//line /usr/local/go/src/text/template/parse/node.go:708
			// _ = "end of CoverTab[29353]"
		}
//line /usr/local/go/src/text/template/parse/node.go:709
		// _ = "end of CoverTab[29351]"
//line /usr/local/go/src/text/template/parse/node.go:709
	}
//line /usr/local/go/src/text/template/parse/node.go:709
	// _ = "end of CoverTab[29324]"
//line /usr/local/go/src/text/template/parse/node.go:709
	_go_fuzz_dep_.CoverTab[29325]++
								if !n.IsInt && func() bool {
//line /usr/local/go/src/text/template/parse/node.go:710
		_go_fuzz_dep_.CoverTab[29366]++
//line /usr/local/go/src/text/template/parse/node.go:710
		return !n.IsUint
//line /usr/local/go/src/text/template/parse/node.go:710
		// _ = "end of CoverTab[29366]"
//line /usr/local/go/src/text/template/parse/node.go:710
	}() && func() bool {
//line /usr/local/go/src/text/template/parse/node.go:710
		_go_fuzz_dep_.CoverTab[29367]++
//line /usr/local/go/src/text/template/parse/node.go:710
		return !n.IsFloat
//line /usr/local/go/src/text/template/parse/node.go:710
		// _ = "end of CoverTab[29367]"
//line /usr/local/go/src/text/template/parse/node.go:710
	}() {
//line /usr/local/go/src/text/template/parse/node.go:710
		_go_fuzz_dep_.CoverTab[29368]++
									return nil, fmt.Errorf("illegal number syntax: %q", text)
//line /usr/local/go/src/text/template/parse/node.go:711
		// _ = "end of CoverTab[29368]"
	} else {
//line /usr/local/go/src/text/template/parse/node.go:712
		_go_fuzz_dep_.CoverTab[29369]++
//line /usr/local/go/src/text/template/parse/node.go:712
		// _ = "end of CoverTab[29369]"
//line /usr/local/go/src/text/template/parse/node.go:712
	}
//line /usr/local/go/src/text/template/parse/node.go:712
	// _ = "end of CoverTab[29325]"
//line /usr/local/go/src/text/template/parse/node.go:712
	_go_fuzz_dep_.CoverTab[29326]++
								return n, nil
//line /usr/local/go/src/text/template/parse/node.go:713
	// _ = "end of CoverTab[29326]"
}

// simplifyComplex pulls out any other types that are represented by the complex number.
//line /usr/local/go/src/text/template/parse/node.go:716
// These all require that the imaginary part be zero.
//line /usr/local/go/src/text/template/parse/node.go:718
func (n *NumberNode) simplifyComplex() {
//line /usr/local/go/src/text/template/parse/node.go:718
	_go_fuzz_dep_.CoverTab[29370]++
								n.IsFloat = imag(n.Complex128) == 0
								if n.IsFloat {
//line /usr/local/go/src/text/template/parse/node.go:720
		_go_fuzz_dep_.CoverTab[29371]++
									n.Float64 = real(n.Complex128)
									n.IsInt = float64(int64(n.Float64)) == n.Float64
									if n.IsInt {
//line /usr/local/go/src/text/template/parse/node.go:723
			_go_fuzz_dep_.CoverTab[29373]++
										n.Int64 = int64(n.Float64)
//line /usr/local/go/src/text/template/parse/node.go:724
			// _ = "end of CoverTab[29373]"
		} else {
//line /usr/local/go/src/text/template/parse/node.go:725
			_go_fuzz_dep_.CoverTab[29374]++
//line /usr/local/go/src/text/template/parse/node.go:725
			// _ = "end of CoverTab[29374]"
//line /usr/local/go/src/text/template/parse/node.go:725
		}
//line /usr/local/go/src/text/template/parse/node.go:725
		// _ = "end of CoverTab[29371]"
//line /usr/local/go/src/text/template/parse/node.go:725
		_go_fuzz_dep_.CoverTab[29372]++
									n.IsUint = float64(uint64(n.Float64)) == n.Float64
									if n.IsUint {
//line /usr/local/go/src/text/template/parse/node.go:727
			_go_fuzz_dep_.CoverTab[29375]++
										n.Uint64 = uint64(n.Float64)
//line /usr/local/go/src/text/template/parse/node.go:728
			// _ = "end of CoverTab[29375]"
		} else {
//line /usr/local/go/src/text/template/parse/node.go:729
			_go_fuzz_dep_.CoverTab[29376]++
//line /usr/local/go/src/text/template/parse/node.go:729
			// _ = "end of CoverTab[29376]"
//line /usr/local/go/src/text/template/parse/node.go:729
		}
//line /usr/local/go/src/text/template/parse/node.go:729
		// _ = "end of CoverTab[29372]"
	} else {
//line /usr/local/go/src/text/template/parse/node.go:730
		_go_fuzz_dep_.CoverTab[29377]++
//line /usr/local/go/src/text/template/parse/node.go:730
		// _ = "end of CoverTab[29377]"
//line /usr/local/go/src/text/template/parse/node.go:730
	}
//line /usr/local/go/src/text/template/parse/node.go:730
	// _ = "end of CoverTab[29370]"
}

func (n *NumberNode) String() string {
//line /usr/local/go/src/text/template/parse/node.go:733
	_go_fuzz_dep_.CoverTab[29378]++
								return n.Text
//line /usr/local/go/src/text/template/parse/node.go:734
	// _ = "end of CoverTab[29378]"
}

func (n *NumberNode) writeTo(sb *strings.Builder) {
//line /usr/local/go/src/text/template/parse/node.go:737
	_go_fuzz_dep_.CoverTab[29379]++
								sb.WriteString(n.String())
//line /usr/local/go/src/text/template/parse/node.go:738
	// _ = "end of CoverTab[29379]"
}

func (n *NumberNode) tree() *Tree {
//line /usr/local/go/src/text/template/parse/node.go:741
	_go_fuzz_dep_.CoverTab[29380]++
								return n.tr
//line /usr/local/go/src/text/template/parse/node.go:742
	// _ = "end of CoverTab[29380]"
}

func (n *NumberNode) Copy() Node {
//line /usr/local/go/src/text/template/parse/node.go:745
	_go_fuzz_dep_.CoverTab[29381]++
								nn := new(NumberNode)
								*nn = *n
								return nn
//line /usr/local/go/src/text/template/parse/node.go:748
	// _ = "end of CoverTab[29381]"
}

// StringNode holds a string constant. The value has been "unquoted".
type StringNode struct {
	NodeType
	Pos
	tr	*Tree
	Quoted	string	// The original text of the string, with quotes.
	Text	string	// The string, after quote processing.
}

func (t *Tree) newString(pos Pos, orig, text string) *StringNode {
//line /usr/local/go/src/text/template/parse/node.go:760
	_go_fuzz_dep_.CoverTab[29382]++
								return &StringNode{tr: t, NodeType: NodeString, Pos: pos, Quoted: orig, Text: text}
//line /usr/local/go/src/text/template/parse/node.go:761
	// _ = "end of CoverTab[29382]"
}

func (s *StringNode) String() string {
//line /usr/local/go/src/text/template/parse/node.go:764
	_go_fuzz_dep_.CoverTab[29383]++
								return s.Quoted
//line /usr/local/go/src/text/template/parse/node.go:765
	// _ = "end of CoverTab[29383]"
}

func (s *StringNode) writeTo(sb *strings.Builder) {
//line /usr/local/go/src/text/template/parse/node.go:768
	_go_fuzz_dep_.CoverTab[29384]++
								sb.WriteString(s.String())
//line /usr/local/go/src/text/template/parse/node.go:769
	// _ = "end of CoverTab[29384]"
}

func (s *StringNode) tree() *Tree {
//line /usr/local/go/src/text/template/parse/node.go:772
	_go_fuzz_dep_.CoverTab[29385]++
								return s.tr
//line /usr/local/go/src/text/template/parse/node.go:773
	// _ = "end of CoverTab[29385]"
}

func (s *StringNode) Copy() Node {
//line /usr/local/go/src/text/template/parse/node.go:776
	_go_fuzz_dep_.CoverTab[29386]++
								return s.tr.newString(s.Pos, s.Quoted, s.Text)
//line /usr/local/go/src/text/template/parse/node.go:777
	// _ = "end of CoverTab[29386]"
}

// endNode represents an {{end}} action.
//line /usr/local/go/src/text/template/parse/node.go:780
// It does not appear in the final parse tree.
//line /usr/local/go/src/text/template/parse/node.go:782
type endNode struct {
	NodeType
	Pos
	tr	*Tree
}

func (t *Tree) newEnd(pos Pos) *endNode {
//line /usr/local/go/src/text/template/parse/node.go:788
	_go_fuzz_dep_.CoverTab[29387]++
								return &endNode{tr: t, NodeType: nodeEnd, Pos: pos}
//line /usr/local/go/src/text/template/parse/node.go:789
	// _ = "end of CoverTab[29387]"
}

func (e *endNode) String() string {
//line /usr/local/go/src/text/template/parse/node.go:792
	_go_fuzz_dep_.CoverTab[29388]++
								return "{{end}}"
//line /usr/local/go/src/text/template/parse/node.go:793
	// _ = "end of CoverTab[29388]"
}

func (e *endNode) writeTo(sb *strings.Builder) {
//line /usr/local/go/src/text/template/parse/node.go:796
	_go_fuzz_dep_.CoverTab[29389]++
								sb.WriteString(e.String())
//line /usr/local/go/src/text/template/parse/node.go:797
	// _ = "end of CoverTab[29389]"
}

func (e *endNode) tree() *Tree {
//line /usr/local/go/src/text/template/parse/node.go:800
	_go_fuzz_dep_.CoverTab[29390]++
								return e.tr
//line /usr/local/go/src/text/template/parse/node.go:801
	// _ = "end of CoverTab[29390]"
}

func (e *endNode) Copy() Node {
//line /usr/local/go/src/text/template/parse/node.go:804
	_go_fuzz_dep_.CoverTab[29391]++
								return e.tr.newEnd(e.Pos)
//line /usr/local/go/src/text/template/parse/node.go:805
	// _ = "end of CoverTab[29391]"
}

// elseNode represents an {{else}} action. Does not appear in the final tree.
type elseNode struct {
	NodeType
	Pos
	tr	*Tree
	Line	int	// The line number in the input. Deprecated: Kept for compatibility.
}

func (t *Tree) newElse(pos Pos, line int) *elseNode {
//line /usr/local/go/src/text/template/parse/node.go:816
	_go_fuzz_dep_.CoverTab[29392]++
								return &elseNode{tr: t, NodeType: nodeElse, Pos: pos, Line: line}
//line /usr/local/go/src/text/template/parse/node.go:817
	// _ = "end of CoverTab[29392]"
}

func (e *elseNode) Type() NodeType {
//line /usr/local/go/src/text/template/parse/node.go:820
	_go_fuzz_dep_.CoverTab[29393]++
								return nodeElse
//line /usr/local/go/src/text/template/parse/node.go:821
	// _ = "end of CoverTab[29393]"
}

func (e *elseNode) String() string {
//line /usr/local/go/src/text/template/parse/node.go:824
	_go_fuzz_dep_.CoverTab[29394]++
								return "{{else}}"
//line /usr/local/go/src/text/template/parse/node.go:825
	// _ = "end of CoverTab[29394]"
}

func (e *elseNode) writeTo(sb *strings.Builder) {
//line /usr/local/go/src/text/template/parse/node.go:828
	_go_fuzz_dep_.CoverTab[29395]++
								sb.WriteString(e.String())
//line /usr/local/go/src/text/template/parse/node.go:829
	// _ = "end of CoverTab[29395]"
}

func (e *elseNode) tree() *Tree {
//line /usr/local/go/src/text/template/parse/node.go:832
	_go_fuzz_dep_.CoverTab[29396]++
								return e.tr
//line /usr/local/go/src/text/template/parse/node.go:833
	// _ = "end of CoverTab[29396]"
}

func (e *elseNode) Copy() Node {
//line /usr/local/go/src/text/template/parse/node.go:836
	_go_fuzz_dep_.CoverTab[29397]++
								return e.tr.newElse(e.Pos, e.Line)
//line /usr/local/go/src/text/template/parse/node.go:837
	// _ = "end of CoverTab[29397]"
}

// BranchNode is the common representation of if, range, and with.
type BranchNode struct {
	NodeType
	Pos
	tr		*Tree
	Line		int		// The line number in the input. Deprecated: Kept for compatibility.
	Pipe		*PipeNode	// The pipeline to be evaluated.
	List		*ListNode	// What to execute if the value is non-empty.
	ElseList	*ListNode	// What to execute if the value is empty (nil if absent).
}

func (b *BranchNode) String() string {
//line /usr/local/go/src/text/template/parse/node.go:851
	_go_fuzz_dep_.CoverTab[29398]++
								var sb strings.Builder
								b.writeTo(&sb)
								return sb.String()
//line /usr/local/go/src/text/template/parse/node.go:854
	// _ = "end of CoverTab[29398]"
}

func (b *BranchNode) writeTo(sb *strings.Builder) {
//line /usr/local/go/src/text/template/parse/node.go:857
	_go_fuzz_dep_.CoverTab[29399]++
								name := ""
								switch b.NodeType {
	case NodeIf:
//line /usr/local/go/src/text/template/parse/node.go:860
		_go_fuzz_dep_.CoverTab[29402]++
									name = "if"
//line /usr/local/go/src/text/template/parse/node.go:861
		// _ = "end of CoverTab[29402]"
	case NodeRange:
//line /usr/local/go/src/text/template/parse/node.go:862
		_go_fuzz_dep_.CoverTab[29403]++
									name = "range"
//line /usr/local/go/src/text/template/parse/node.go:863
		// _ = "end of CoverTab[29403]"
	case NodeWith:
//line /usr/local/go/src/text/template/parse/node.go:864
		_go_fuzz_dep_.CoverTab[29404]++
									name = "with"
//line /usr/local/go/src/text/template/parse/node.go:865
		// _ = "end of CoverTab[29404]"
	default:
//line /usr/local/go/src/text/template/parse/node.go:866
		_go_fuzz_dep_.CoverTab[29405]++
									panic("unknown branch type")
//line /usr/local/go/src/text/template/parse/node.go:867
		// _ = "end of CoverTab[29405]"
	}
//line /usr/local/go/src/text/template/parse/node.go:868
	// _ = "end of CoverTab[29399]"
//line /usr/local/go/src/text/template/parse/node.go:868
	_go_fuzz_dep_.CoverTab[29400]++
								sb.WriteString("{{")
								sb.WriteString(name)
								sb.WriteByte(' ')
								b.Pipe.writeTo(sb)
								sb.WriteString("}}")
								b.List.writeTo(sb)
								if b.ElseList != nil {
//line /usr/local/go/src/text/template/parse/node.go:875
		_go_fuzz_dep_.CoverTab[29406]++
									sb.WriteString("{{else}}")
									b.ElseList.writeTo(sb)
//line /usr/local/go/src/text/template/parse/node.go:877
		// _ = "end of CoverTab[29406]"
	} else {
//line /usr/local/go/src/text/template/parse/node.go:878
		_go_fuzz_dep_.CoverTab[29407]++
//line /usr/local/go/src/text/template/parse/node.go:878
		// _ = "end of CoverTab[29407]"
//line /usr/local/go/src/text/template/parse/node.go:878
	}
//line /usr/local/go/src/text/template/parse/node.go:878
	// _ = "end of CoverTab[29400]"
//line /usr/local/go/src/text/template/parse/node.go:878
	_go_fuzz_dep_.CoverTab[29401]++
								sb.WriteString("{{end}}")
//line /usr/local/go/src/text/template/parse/node.go:879
	// _ = "end of CoverTab[29401]"
}

func (b *BranchNode) tree() *Tree {
//line /usr/local/go/src/text/template/parse/node.go:882
	_go_fuzz_dep_.CoverTab[29408]++
								return b.tr
//line /usr/local/go/src/text/template/parse/node.go:883
	// _ = "end of CoverTab[29408]"
}

func (b *BranchNode) Copy() Node {
//line /usr/local/go/src/text/template/parse/node.go:886
	_go_fuzz_dep_.CoverTab[29409]++
								switch b.NodeType {
	case NodeIf:
//line /usr/local/go/src/text/template/parse/node.go:888
		_go_fuzz_dep_.CoverTab[29410]++
									return b.tr.newIf(b.Pos, b.Line, b.Pipe, b.List, b.ElseList)
//line /usr/local/go/src/text/template/parse/node.go:889
		// _ = "end of CoverTab[29410]"
	case NodeRange:
//line /usr/local/go/src/text/template/parse/node.go:890
		_go_fuzz_dep_.CoverTab[29411]++
									return b.tr.newRange(b.Pos, b.Line, b.Pipe, b.List, b.ElseList)
//line /usr/local/go/src/text/template/parse/node.go:891
		// _ = "end of CoverTab[29411]"
	case NodeWith:
//line /usr/local/go/src/text/template/parse/node.go:892
		_go_fuzz_dep_.CoverTab[29412]++
									return b.tr.newWith(b.Pos, b.Line, b.Pipe, b.List, b.ElseList)
//line /usr/local/go/src/text/template/parse/node.go:893
		// _ = "end of CoverTab[29412]"
	default:
//line /usr/local/go/src/text/template/parse/node.go:894
		_go_fuzz_dep_.CoverTab[29413]++
									panic("unknown branch type")
//line /usr/local/go/src/text/template/parse/node.go:895
		// _ = "end of CoverTab[29413]"
	}
//line /usr/local/go/src/text/template/parse/node.go:896
	// _ = "end of CoverTab[29409]"
}

// IfNode represents an {{if}} action and its commands.
type IfNode struct {
	BranchNode
}

func (t *Tree) newIf(pos Pos, line int, pipe *PipeNode, list, elseList *ListNode) *IfNode {
//line /usr/local/go/src/text/template/parse/node.go:904
	_go_fuzz_dep_.CoverTab[29414]++
								return &IfNode{BranchNode{tr: t, NodeType: NodeIf, Pos: pos, Line: line, Pipe: pipe, List: list, ElseList: elseList}}
//line /usr/local/go/src/text/template/parse/node.go:905
	// _ = "end of CoverTab[29414]"
}

func (i *IfNode) Copy() Node {
//line /usr/local/go/src/text/template/parse/node.go:908
	_go_fuzz_dep_.CoverTab[29415]++
								return i.tr.newIf(i.Pos, i.Line, i.Pipe.CopyPipe(), i.List.CopyList(), i.ElseList.CopyList())
//line /usr/local/go/src/text/template/parse/node.go:909
	// _ = "end of CoverTab[29415]"
}

// BreakNode represents a {{break}} action.
type BreakNode struct {
	tr	*Tree
	NodeType
	Pos
	Line	int
}

func (t *Tree) newBreak(pos Pos, line int) *BreakNode {
//line /usr/local/go/src/text/template/parse/node.go:920
	_go_fuzz_dep_.CoverTab[29416]++
								return &BreakNode{tr: t, NodeType: NodeBreak, Pos: pos, Line: line}
//line /usr/local/go/src/text/template/parse/node.go:921
	// _ = "end of CoverTab[29416]"
}

func (b *BreakNode) Copy() Node {
//line /usr/local/go/src/text/template/parse/node.go:924
	_go_fuzz_dep_.CoverTab[29417]++
//line /usr/local/go/src/text/template/parse/node.go:924
	return b.tr.newBreak(b.Pos, b.Line)
//line /usr/local/go/src/text/template/parse/node.go:924
	// _ = "end of CoverTab[29417]"
//line /usr/local/go/src/text/template/parse/node.go:924
}
func (b *BreakNode) String() string {
//line /usr/local/go/src/text/template/parse/node.go:925
	_go_fuzz_dep_.CoverTab[29418]++
//line /usr/local/go/src/text/template/parse/node.go:925
	return "{{break}}"
//line /usr/local/go/src/text/template/parse/node.go:925
	// _ = "end of CoverTab[29418]"
//line /usr/local/go/src/text/template/parse/node.go:925
}
func (b *BreakNode) tree() *Tree {
//line /usr/local/go/src/text/template/parse/node.go:926
	_go_fuzz_dep_.CoverTab[29419]++
//line /usr/local/go/src/text/template/parse/node.go:926
	return b.tr
//line /usr/local/go/src/text/template/parse/node.go:926
	// _ = "end of CoverTab[29419]"
//line /usr/local/go/src/text/template/parse/node.go:926
}
func (b *BreakNode) writeTo(sb *strings.Builder) {
//line /usr/local/go/src/text/template/parse/node.go:927
	_go_fuzz_dep_.CoverTab[29420]++
//line /usr/local/go/src/text/template/parse/node.go:927
	sb.WriteString("{{break}}")
//line /usr/local/go/src/text/template/parse/node.go:927
	// _ = "end of CoverTab[29420]"
//line /usr/local/go/src/text/template/parse/node.go:927
}

// ContinueNode represents a {{continue}} action.
type ContinueNode struct {
	tr	*Tree
	NodeType
	Pos
	Line	int
}

func (t *Tree) newContinue(pos Pos, line int) *ContinueNode {
//line /usr/local/go/src/text/template/parse/node.go:937
	_go_fuzz_dep_.CoverTab[29421]++
								return &ContinueNode{tr: t, NodeType: NodeContinue, Pos: pos, Line: line}
//line /usr/local/go/src/text/template/parse/node.go:938
	// _ = "end of CoverTab[29421]"
}

func (c *ContinueNode) Copy() Node {
//line /usr/local/go/src/text/template/parse/node.go:941
	_go_fuzz_dep_.CoverTab[29422]++
//line /usr/local/go/src/text/template/parse/node.go:941
	return c.tr.newContinue(c.Pos, c.Line)
//line /usr/local/go/src/text/template/parse/node.go:941
	// _ = "end of CoverTab[29422]"
//line /usr/local/go/src/text/template/parse/node.go:941
}
func (c *ContinueNode) String() string {
//line /usr/local/go/src/text/template/parse/node.go:942
	_go_fuzz_dep_.CoverTab[29423]++
//line /usr/local/go/src/text/template/parse/node.go:942
	return "{{continue}}"
//line /usr/local/go/src/text/template/parse/node.go:942
	// _ = "end of CoverTab[29423]"
//line /usr/local/go/src/text/template/parse/node.go:942
}
func (c *ContinueNode) tree() *Tree {
//line /usr/local/go/src/text/template/parse/node.go:943
	_go_fuzz_dep_.CoverTab[29424]++
//line /usr/local/go/src/text/template/parse/node.go:943
	return c.tr
//line /usr/local/go/src/text/template/parse/node.go:943
	// _ = "end of CoverTab[29424]"
//line /usr/local/go/src/text/template/parse/node.go:943
}
func (c *ContinueNode) writeTo(sb *strings.Builder) {
//line /usr/local/go/src/text/template/parse/node.go:944
	_go_fuzz_dep_.CoverTab[29425]++
//line /usr/local/go/src/text/template/parse/node.go:944
	sb.WriteString("{{continue}}")
//line /usr/local/go/src/text/template/parse/node.go:944
	// _ = "end of CoverTab[29425]"
//line /usr/local/go/src/text/template/parse/node.go:944
}

// RangeNode represents a {{range}} action and its commands.
type RangeNode struct {
	BranchNode
}

func (t *Tree) newRange(pos Pos, line int, pipe *PipeNode, list, elseList *ListNode) *RangeNode {
//line /usr/local/go/src/text/template/parse/node.go:951
	_go_fuzz_dep_.CoverTab[29426]++
								return &RangeNode{BranchNode{tr: t, NodeType: NodeRange, Pos: pos, Line: line, Pipe: pipe, List: list, ElseList: elseList}}
//line /usr/local/go/src/text/template/parse/node.go:952
	// _ = "end of CoverTab[29426]"
}

func (r *RangeNode) Copy() Node {
//line /usr/local/go/src/text/template/parse/node.go:955
	_go_fuzz_dep_.CoverTab[29427]++
								return r.tr.newRange(r.Pos, r.Line, r.Pipe.CopyPipe(), r.List.CopyList(), r.ElseList.CopyList())
//line /usr/local/go/src/text/template/parse/node.go:956
	// _ = "end of CoverTab[29427]"
}

// WithNode represents a {{with}} action and its commands.
type WithNode struct {
	BranchNode
}

func (t *Tree) newWith(pos Pos, line int, pipe *PipeNode, list, elseList *ListNode) *WithNode {
//line /usr/local/go/src/text/template/parse/node.go:964
	_go_fuzz_dep_.CoverTab[29428]++
								return &WithNode{BranchNode{tr: t, NodeType: NodeWith, Pos: pos, Line: line, Pipe: pipe, List: list, ElseList: elseList}}
//line /usr/local/go/src/text/template/parse/node.go:965
	// _ = "end of CoverTab[29428]"
}

func (w *WithNode) Copy() Node {
//line /usr/local/go/src/text/template/parse/node.go:968
	_go_fuzz_dep_.CoverTab[29429]++
								return w.tr.newWith(w.Pos, w.Line, w.Pipe.CopyPipe(), w.List.CopyList(), w.ElseList.CopyList())
//line /usr/local/go/src/text/template/parse/node.go:969
	// _ = "end of CoverTab[29429]"
}

// TemplateNode represents a {{template}} action.
type TemplateNode struct {
	NodeType
	Pos
	tr	*Tree
	Line	int		// The line number in the input. Deprecated: Kept for compatibility.
	Name	string		// The name of the template (unquoted).
	Pipe	*PipeNode	// The command to evaluate as dot for the template.
}

func (t *Tree) newTemplate(pos Pos, line int, name string, pipe *PipeNode) *TemplateNode {
//line /usr/local/go/src/text/template/parse/node.go:982
	_go_fuzz_dep_.CoverTab[29430]++
								return &TemplateNode{tr: t, NodeType: NodeTemplate, Pos: pos, Line: line, Name: name, Pipe: pipe}
//line /usr/local/go/src/text/template/parse/node.go:983
	// _ = "end of CoverTab[29430]"
}

func (t *TemplateNode) String() string {
//line /usr/local/go/src/text/template/parse/node.go:986
	_go_fuzz_dep_.CoverTab[29431]++
								var sb strings.Builder
								t.writeTo(&sb)
								return sb.String()
//line /usr/local/go/src/text/template/parse/node.go:989
	// _ = "end of CoverTab[29431]"
}

func (t *TemplateNode) writeTo(sb *strings.Builder) {
//line /usr/local/go/src/text/template/parse/node.go:992
	_go_fuzz_dep_.CoverTab[29432]++
								sb.WriteString("{{template ")
								sb.WriteString(strconv.Quote(t.Name))
								if t.Pipe != nil {
//line /usr/local/go/src/text/template/parse/node.go:995
		_go_fuzz_dep_.CoverTab[29434]++
									sb.WriteByte(' ')
									t.Pipe.writeTo(sb)
//line /usr/local/go/src/text/template/parse/node.go:997
		// _ = "end of CoverTab[29434]"
	} else {
//line /usr/local/go/src/text/template/parse/node.go:998
		_go_fuzz_dep_.CoverTab[29435]++
//line /usr/local/go/src/text/template/parse/node.go:998
		// _ = "end of CoverTab[29435]"
//line /usr/local/go/src/text/template/parse/node.go:998
	}
//line /usr/local/go/src/text/template/parse/node.go:998
	// _ = "end of CoverTab[29432]"
//line /usr/local/go/src/text/template/parse/node.go:998
	_go_fuzz_dep_.CoverTab[29433]++
								sb.WriteString("}}")
//line /usr/local/go/src/text/template/parse/node.go:999
	// _ = "end of CoverTab[29433]"
}

func (t *TemplateNode) tree() *Tree {
//line /usr/local/go/src/text/template/parse/node.go:1002
	_go_fuzz_dep_.CoverTab[29436]++
								return t.tr
//line /usr/local/go/src/text/template/parse/node.go:1003
	// _ = "end of CoverTab[29436]"
}

func (t *TemplateNode) Copy() Node {
//line /usr/local/go/src/text/template/parse/node.go:1006
	_go_fuzz_dep_.CoverTab[29437]++
								return t.tr.newTemplate(t.Pos, t.Line, t.Name, t.Pipe.CopyPipe())
//line /usr/local/go/src/text/template/parse/node.go:1007
	// _ = "end of CoverTab[29437]"
}

//line /usr/local/go/src/text/template/parse/node.go:1008
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/text/template/parse/node.go:1008
var _ = _go_fuzz_dep_.CoverTab
