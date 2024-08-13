//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:1
// Package ast declares the types used to represent syntax trees for HCL
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:1
// (HashiCorp Configuration Language)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:3
package ast

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:3
import (
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:3
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:3
)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:3
import (
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:3
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:3
)

import (
	"fmt"
	"strings"

	"github.com/hashicorp/hcl/hcl/token"
)

// Node is an element in the abstract syntax tree.
type Node interface {
	node()
	Pos() token.Pos
}

func (File) node()		{ _go_fuzz_dep_.CoverTab[121012]++; // _ = "end of CoverTab[121012]" }
func (ObjectList) node()	{ _go_fuzz_dep_.CoverTab[121013]++; // _ = "end of CoverTab[121013]" }
func (ObjectKey) node()		{ _go_fuzz_dep_.CoverTab[121014]++; // _ = "end of CoverTab[121014]" }
func (ObjectItem) node()	{ _go_fuzz_dep_.CoverTab[121015]++; // _ = "end of CoverTab[121015]" }
func (Comment) node()		{ _go_fuzz_dep_.CoverTab[121016]++; // _ = "end of CoverTab[121016]" }
func (CommentGroup) node()	{ _go_fuzz_dep_.CoverTab[121017]++; // _ = "end of CoverTab[121017]" }
func (ObjectType) node()	{ _go_fuzz_dep_.CoverTab[121018]++; // _ = "end of CoverTab[121018]" }
func (LiteralType) node()	{ _go_fuzz_dep_.CoverTab[121019]++; // _ = "end of CoverTab[121019]" }
func (ListType) node()		{ _go_fuzz_dep_.CoverTab[121020]++; // _ = "end of CoverTab[121020]" }

// File represents a single HCL file
type File struct {
	Node		Node		// usually a *ObjectList
	Comments	[]*CommentGroup	// list of all comments in the source
}

func (f *File) Pos() token.Pos {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:34
	_go_fuzz_dep_.CoverTab[121021]++
											return f.Node.Pos()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:35
	// _ = "end of CoverTab[121021]"
}

// ObjectList represents a list of ObjectItems. An HCL file itself is an
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:38
// ObjectList.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:40
type ObjectList struct {
	Items []*ObjectItem
}

func (o *ObjectList) Add(item *ObjectItem) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:44
	_go_fuzz_dep_.CoverTab[121022]++
											o.Items = append(o.Items, item)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:45
	// _ = "end of CoverTab[121022]"
}

// Filter filters out the objects with the given key list as a prefix.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:48
//
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:48
// The returned list of objects contain ObjectItems where the keys have
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:48
// this prefix already stripped off. This might result in objects with
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:48
// zero-length key lists if they have no children.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:48
//
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:48
// If no matches are found, an empty ObjectList (non-nil) is returned.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:55
func (o *ObjectList) Filter(keys ...string) *ObjectList {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:55
	_go_fuzz_dep_.CoverTab[121023]++
											var result ObjectList
											for _, item := range o.Items {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:57
		_go_fuzz_dep_.CoverTab[121025]++

												if len(item.Keys) < len(keys) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:59
			_go_fuzz_dep_.CoverTab[121029]++
													continue
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:60
			// _ = "end of CoverTab[121029]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:61
			_go_fuzz_dep_.CoverTab[121030]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:61
			// _ = "end of CoverTab[121030]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:61
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:61
		// _ = "end of CoverTab[121025]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:61
		_go_fuzz_dep_.CoverTab[121026]++

												match := true
												for i, key := range item.Keys[:len(keys)] {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:64
			_go_fuzz_dep_.CoverTab[121031]++
													key := key.Token.Value().(string)
													if key != keys[i] && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:66
				_go_fuzz_dep_.CoverTab[121032]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:66
				return !strings.EqualFold(key, keys[i])
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:66
				// _ = "end of CoverTab[121032]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:66
			}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:66
				_go_fuzz_dep_.CoverTab[121033]++
														match = false
														break
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:68
				// _ = "end of CoverTab[121033]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:69
				_go_fuzz_dep_.CoverTab[121034]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:69
				// _ = "end of CoverTab[121034]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:69
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:69
			// _ = "end of CoverTab[121031]"
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:70
		// _ = "end of CoverTab[121026]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:70
		_go_fuzz_dep_.CoverTab[121027]++
												if !match {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:71
			_go_fuzz_dep_.CoverTab[121035]++
													continue
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:72
			// _ = "end of CoverTab[121035]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:73
			_go_fuzz_dep_.CoverTab[121036]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:73
			// _ = "end of CoverTab[121036]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:73
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:73
		// _ = "end of CoverTab[121027]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:73
		_go_fuzz_dep_.CoverTab[121028]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:76
		newItem := *item
												newItem.Keys = newItem.Keys[len(keys):]
												result.Add(&newItem)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:78
		// _ = "end of CoverTab[121028]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:79
	// _ = "end of CoverTab[121023]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:79
	_go_fuzz_dep_.CoverTab[121024]++

											return &result
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:81
	// _ = "end of CoverTab[121024]"
}

// Children returns further nested objects (key length > 0) within this
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:84
// ObjectList. This should be used with Filter to get at child items.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:86
func (o *ObjectList) Children() *ObjectList {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:86
	_go_fuzz_dep_.CoverTab[121037]++
											var result ObjectList
											for _, item := range o.Items {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:88
		_go_fuzz_dep_.CoverTab[121039]++
												if len(item.Keys) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:89
			_go_fuzz_dep_.CoverTab[121040]++
													result.Add(item)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:90
			// _ = "end of CoverTab[121040]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:91
			_go_fuzz_dep_.CoverTab[121041]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:91
			// _ = "end of CoverTab[121041]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:91
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:91
		// _ = "end of CoverTab[121039]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:92
	// _ = "end of CoverTab[121037]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:92
	_go_fuzz_dep_.CoverTab[121038]++

											return &result
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:94
	// _ = "end of CoverTab[121038]"
}

// Elem returns items in the list that are direct element assignments
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:97
// (key length == 0). This should be used with Filter to get at elements.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:99
func (o *ObjectList) Elem() *ObjectList {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:99
	_go_fuzz_dep_.CoverTab[121042]++
											var result ObjectList
											for _, item := range o.Items {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:101
		_go_fuzz_dep_.CoverTab[121044]++
												if len(item.Keys) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:102
			_go_fuzz_dep_.CoverTab[121045]++
													result.Add(item)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:103
			// _ = "end of CoverTab[121045]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:104
			_go_fuzz_dep_.CoverTab[121046]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:104
			// _ = "end of CoverTab[121046]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:104
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:104
		// _ = "end of CoverTab[121044]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:105
	// _ = "end of CoverTab[121042]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:105
	_go_fuzz_dep_.CoverTab[121043]++

											return &result
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:107
	// _ = "end of CoverTab[121043]"
}

func (o *ObjectList) Pos() token.Pos {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:110
	_go_fuzz_dep_.CoverTab[121047]++

											return o.Items[0].Pos()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:112
	// _ = "end of CoverTab[121047]"
}

// ObjectItem represents a HCL Object Item. An item is represented with a key
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:115
// (or keys). It can be an assignment or an object (both normal and nested)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:117
type ObjectItem struct {
	// keys is only one length long if it's of type assignment. If it's a
	// nested object it can be larger than one. In that case "assign" is
	// invalid as there is no assignments for a nested object.
	Keys	[]*ObjectKey

	// assign contains the position of "=", if any
	Assign	token.Pos

	// val is the item itself. It can be an object,list, number, bool or a
	// string. If key length is larger than one, val can be only of type
	// Object.
	Val	Node

	LeadComment	*CommentGroup	// associated lead comment
	LineComment	*CommentGroup	// associated line comment
}

func (o *ObjectItem) Pos() token.Pos {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:135
	_go_fuzz_dep_.CoverTab[121048]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:138
	if len(o.Keys) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:138
		_go_fuzz_dep_.CoverTab[121050]++
												return token.Pos{}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:139
		// _ = "end of CoverTab[121050]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:140
		_go_fuzz_dep_.CoverTab[121051]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:140
		// _ = "end of CoverTab[121051]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:140
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:140
	// _ = "end of CoverTab[121048]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:140
	_go_fuzz_dep_.CoverTab[121049]++

											return o.Keys[0].Pos()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:142
	// _ = "end of CoverTab[121049]"
}

// ObjectKeys are either an identifier or of type string.
type ObjectKey struct {
	Token token.Token
}

func (o *ObjectKey) Pos() token.Pos {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:150
	_go_fuzz_dep_.CoverTab[121052]++
											return o.Token.Pos
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:151
	// _ = "end of CoverTab[121052]"
}

// LiteralType represents a literal of basic type. Valid types are:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:154
// token.NUMBER, token.FLOAT, token.BOOL and token.STRING
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:156
type LiteralType struct {
	Token	token.Token

	// comment types, only used when in a list
	LeadComment	*CommentGroup
	LineComment	*CommentGroup
}

func (l *LiteralType) Pos() token.Pos {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:164
	_go_fuzz_dep_.CoverTab[121053]++
											return l.Token.Pos
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:165
	// _ = "end of CoverTab[121053]"
}

// ListStatement represents a HCL List type
type ListType struct {
	Lbrack	token.Pos	// position of "["
	Rbrack	token.Pos	// position of "]"
	List	[]Node		// the elements in lexical order
}

func (l *ListType) Pos() token.Pos {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:175
	_go_fuzz_dep_.CoverTab[121054]++
											return l.Lbrack
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:176
	// _ = "end of CoverTab[121054]"
}

func (l *ListType) Add(node Node) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:179
	_go_fuzz_dep_.CoverTab[121055]++
											l.List = append(l.List, node)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:180
	// _ = "end of CoverTab[121055]"
}

// ObjectType represents a HCL Object Type
type ObjectType struct {
	Lbrace	token.Pos	// position of "{"
	Rbrace	token.Pos	// position of "}"
	List	*ObjectList	// the nodes in lexical order
}

func (o *ObjectType) Pos() token.Pos {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:190
	_go_fuzz_dep_.CoverTab[121056]++
											return o.Lbrace
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:191
	// _ = "end of CoverTab[121056]"
}

// Comment node represents a single //, # style or /*- style commment
type Comment struct {
	Start	token.Pos	// position of / or #
	Text	string
}

func (c *Comment) Pos() token.Pos {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:200
	_go_fuzz_dep_.CoverTab[121057]++
											return c.Start
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:201
	// _ = "end of CoverTab[121057]"
}

// CommentGroup node represents a sequence of comments with no other tokens and
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:204
// no empty lines between.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:206
type CommentGroup struct {
	List []*Comment	// len(List) > 0
}

func (c *CommentGroup) Pos() token.Pos {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:210
	_go_fuzz_dep_.CoverTab[121058]++
											return c.List[0].Pos()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:211
	// _ = "end of CoverTab[121058]"
}

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:218
func (o *ObjectKey) GoString() string {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:218
	_go_fuzz_dep_.CoverTab[121059]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:218
	return fmt.Sprintf("*%#v", *o)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:218
	// _ = "end of CoverTab[121059]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:218
}
func (o *ObjectList) GoString() string {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:219
	_go_fuzz_dep_.CoverTab[121060]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:219
	return fmt.Sprintf("*%#v", *o)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:219
	// _ = "end of CoverTab[121060]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:219
}

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:219
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/ast.go:219
var _ = _go_fuzz_dep_.CoverTab
