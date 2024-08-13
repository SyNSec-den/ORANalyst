//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:1
package ast

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:1
)

import "fmt"

// WalkFunc describes a function to be called for each node during a Walk. The
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:5
// returned node can be used to rewrite the AST. Walking stops the returned
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:5
// bool is false.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:8
type WalkFunc func(Node) (Node, bool)

// Walk traverses an AST in depth-first order: It starts by calling fn(node);
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:10
// node must not be nil. If fn returns true, Walk invokes fn recursively for
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:10
// each of the non-nil children of node, followed by a call of fn(nil). The
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:10
// returned node of fn can be used to rewrite the passed node to fn.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:14
func Walk(node Node, fn WalkFunc) Node {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:14
	_go_fuzz_dep_.CoverTab[121061]++
											rewritten, ok := fn(node)
											if !ok {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:16
		_go_fuzz_dep_.CoverTab[121064]++
												return rewritten
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:17
		// _ = "end of CoverTab[121064]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:18
		_go_fuzz_dep_.CoverTab[121065]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:18
		// _ = "end of CoverTab[121065]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:18
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:18
	// _ = "end of CoverTab[121061]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:18
	_go_fuzz_dep_.CoverTab[121062]++

											switch n := node.(type) {
	case *File:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:21
		_go_fuzz_dep_.CoverTab[121066]++
												n.Node = Walk(n.Node, fn)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:22
		// _ = "end of CoverTab[121066]"
	case *ObjectList:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:23
		_go_fuzz_dep_.CoverTab[121067]++
												for i, item := range n.Items {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:24
			_go_fuzz_dep_.CoverTab[121075]++
													n.Items[i] = Walk(item, fn).(*ObjectItem)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:25
			// _ = "end of CoverTab[121075]"
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:26
		// _ = "end of CoverTab[121067]"
	case *ObjectKey:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:27
		_go_fuzz_dep_.CoverTab[121068]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:27
		// _ = "end of CoverTab[121068]"

	case *ObjectItem:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:29
		_go_fuzz_dep_.CoverTab[121069]++
												for i, k := range n.Keys {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:30
			_go_fuzz_dep_.CoverTab[121076]++
													n.Keys[i] = Walk(k, fn).(*ObjectKey)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:31
			// _ = "end of CoverTab[121076]"
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:32
		// _ = "end of CoverTab[121069]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:32
		_go_fuzz_dep_.CoverTab[121070]++

												if n.Val != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:34
			_go_fuzz_dep_.CoverTab[121077]++
													n.Val = Walk(n.Val, fn)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:35
			// _ = "end of CoverTab[121077]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:36
			_go_fuzz_dep_.CoverTab[121078]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:36
			// _ = "end of CoverTab[121078]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:36
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:36
		// _ = "end of CoverTab[121070]"
	case *LiteralType:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:37
		_go_fuzz_dep_.CoverTab[121071]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:37
		// _ = "end of CoverTab[121071]"

	case *ListType:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:39
		_go_fuzz_dep_.CoverTab[121072]++
												for i, l := range n.List {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:40
			_go_fuzz_dep_.CoverTab[121079]++
													n.List[i] = Walk(l, fn)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:41
			// _ = "end of CoverTab[121079]"
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:42
		// _ = "end of CoverTab[121072]"
	case *ObjectType:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:43
		_go_fuzz_dep_.CoverTab[121073]++
												n.List = Walk(n.List, fn).(*ObjectList)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:44
		// _ = "end of CoverTab[121073]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:45
		_go_fuzz_dep_.CoverTab[121074]++

												fmt.Printf("unknown type: %T\n", n)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:47
		// _ = "end of CoverTab[121074]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:48
	// _ = "end of CoverTab[121062]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:48
	_go_fuzz_dep_.CoverTab[121063]++

											fn(nil)
											return rewritten
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:51
	// _ = "end of CoverTab[121063]"
}

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:52
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/ast/walk.go:52
var _ = _go_fuzz_dep_.CoverTab
