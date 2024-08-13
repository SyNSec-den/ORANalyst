//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:1
package parser

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:1
)

import "github.com/hashicorp/hcl/hcl/ast"

// flattenObjects takes an AST node, walks it, and flattens
func flattenObjects(node ast.Node) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:6
	_go_fuzz_dep_.CoverTab[121765]++
												ast.Walk(node, func(n ast.Node) (ast.Node, bool) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:7
		_go_fuzz_dep_.CoverTab[121766]++

													list, ok := n.(*ast.ObjectList)
													if !ok {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:10
			_go_fuzz_dep_.CoverTab[121770]++
														return n, true
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:11
			// _ = "end of CoverTab[121770]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:12
			_go_fuzz_dep_.CoverTab[121771]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:12
			// _ = "end of CoverTab[121771]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:12
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:12
		// _ = "end of CoverTab[121766]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:12
		_go_fuzz_dep_.CoverTab[121767]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:15
		items := make([]*ast.ObjectItem, 0, len(list.Items))
		frontier := make([]*ast.ObjectItem, len(list.Items))
		copy(frontier, list.Items)
		for len(frontier) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:18
			_go_fuzz_dep_.CoverTab[121772]++

														n := len(frontier)
														item := frontier[n-1]
														frontier = frontier[:n-1]

														switch v := item.Val.(type) {
			case *ast.ObjectType:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:25
				_go_fuzz_dep_.CoverTab[121773]++
															items, frontier = flattenObjectType(v, item, items, frontier)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:26
				// _ = "end of CoverTab[121773]"
			case *ast.ListType:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:27
				_go_fuzz_dep_.CoverTab[121774]++
															items, frontier = flattenListType(v, item, items, frontier)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:28
				// _ = "end of CoverTab[121774]"
			default:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:29
				_go_fuzz_dep_.CoverTab[121775]++
															items = append(items, item)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:30
				// _ = "end of CoverTab[121775]"
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:31
			// _ = "end of CoverTab[121772]"
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:32
		// _ = "end of CoverTab[121767]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:32
		_go_fuzz_dep_.CoverTab[121768]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:35
		for i := len(items)/2 - 1; i >= 0; i-- {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:35
			_go_fuzz_dep_.CoverTab[121776]++
														opp := len(items) - 1 - i
														items[i], items[opp] = items[opp], items[i]
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:37
			// _ = "end of CoverTab[121776]"
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:38
		// _ = "end of CoverTab[121768]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:38
		_go_fuzz_dep_.CoverTab[121769]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:41
		list.Items = items
													return n, true
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:42
		// _ = "end of CoverTab[121769]"
	})
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:43
	// _ = "end of CoverTab[121765]"
}

func flattenListType(
	ot *ast.ListType,
	item *ast.ObjectItem,
	items []*ast.ObjectItem,
	frontier []*ast.ObjectItem) ([]*ast.ObjectItem, []*ast.ObjectItem) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:50
	_go_fuzz_dep_.CoverTab[121777]++

												if len(ot.List) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:52
		_go_fuzz_dep_.CoverTab[121781]++
													items = append(items, item)
													return items, frontier
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:54
		// _ = "end of CoverTab[121781]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:55
		_go_fuzz_dep_.CoverTab[121782]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:55
		// _ = "end of CoverTab[121782]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:55
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:55
	// _ = "end of CoverTab[121777]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:55
	_go_fuzz_dep_.CoverTab[121778]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:58
	for _, subitem := range ot.List {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:58
		_go_fuzz_dep_.CoverTab[121783]++
													if _, ok := subitem.(*ast.ObjectType); !ok {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:59
			_go_fuzz_dep_.CoverTab[121784]++
														items = append(items, item)
														return items, frontier
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:61
			// _ = "end of CoverTab[121784]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:62
			_go_fuzz_dep_.CoverTab[121785]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:62
			// _ = "end of CoverTab[121785]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:62
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:62
		// _ = "end of CoverTab[121783]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:63
	// _ = "end of CoverTab[121778]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:63
	_go_fuzz_dep_.CoverTab[121779]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:66
	for _, elem := range ot.List {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:66
		_go_fuzz_dep_.CoverTab[121786]++

													frontier = append(frontier, &ast.ObjectItem{
			Keys:		item.Keys,
			Assign:		item.Assign,
			Val:		elem,
			LeadComment:	item.LeadComment,
			LineComment:	item.LineComment,
		})
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:74
		// _ = "end of CoverTab[121786]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:75
	// _ = "end of CoverTab[121779]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:75
	_go_fuzz_dep_.CoverTab[121780]++

												return items, frontier
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:77
	// _ = "end of CoverTab[121780]"
}

func flattenObjectType(
	ot *ast.ObjectType,
	item *ast.ObjectItem,
	items []*ast.ObjectItem,
	frontier []*ast.ObjectItem) ([]*ast.ObjectItem, []*ast.ObjectItem) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:84
	_go_fuzz_dep_.CoverTab[121787]++

												if ot.List.Items == nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:86
		_go_fuzz_dep_.CoverTab[121791]++
													items = append(items, item)
													return items, frontier
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:88
		// _ = "end of CoverTab[121791]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:89
		_go_fuzz_dep_.CoverTab[121792]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:89
		// _ = "end of CoverTab[121792]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:89
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:89
	// _ = "end of CoverTab[121787]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:89
	_go_fuzz_dep_.CoverTab[121788]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:92
	for _, subitem := range ot.List.Items {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:92
		_go_fuzz_dep_.CoverTab[121793]++
													if _, ok := subitem.Val.(*ast.ObjectType); !ok {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:93
			_go_fuzz_dep_.CoverTab[121794]++
														items = append(items, item)
														return items, frontier
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:95
			// _ = "end of CoverTab[121794]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:96
			_go_fuzz_dep_.CoverTab[121795]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:96
			// _ = "end of CoverTab[121795]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:96
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:96
		// _ = "end of CoverTab[121793]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:97
	// _ = "end of CoverTab[121788]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:97
	_go_fuzz_dep_.CoverTab[121789]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:100
	for _, subitem := range ot.List.Items {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:100
		_go_fuzz_dep_.CoverTab[121796]++

													keys := make([]*ast.ObjectKey, len(item.Keys)+len(subitem.Keys))
													copy(keys, item.Keys)
													copy(keys[len(item.Keys):], subitem.Keys)

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:107
		frontier = append(frontier, &ast.ObjectItem{
			Keys:		keys,
			Assign:		item.Assign,
			Val:		subitem.Val,
			LeadComment:	item.LeadComment,
			LineComment:	item.LineComment,
		})
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:113
		// _ = "end of CoverTab[121796]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:114
	// _ = "end of CoverTab[121789]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:114
	_go_fuzz_dep_.CoverTab[121790]++

												return items, frontier
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:116
	// _ = "end of CoverTab[121790]"
}

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:117
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/parser/flatten.go:117
var _ = _go_fuzz_dep_.CoverTab
