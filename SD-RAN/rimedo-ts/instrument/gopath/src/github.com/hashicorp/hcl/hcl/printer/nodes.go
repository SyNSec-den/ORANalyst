//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:1
package printer

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:1
)

import (
	"bytes"
	"fmt"
	"sort"

	"github.com/hashicorp/hcl/hcl/ast"
	"github.com/hashicorp/hcl/hcl/token"
)

const (
	blank		= byte(' ')
	newline		= byte('\n')
	tab		= byte('\t')
	infinity	= 1 << 30	// offset or line
)

var (
	unindent = []byte("\uE123")	// in the private use space
)

type printer struct {
	cfg	Config
	prev	token.Pos

	comments		[]*ast.CommentGroup	// may be nil, contains all comments
	standaloneComments	[]*ast.CommentGroup	// contains all standalone comments (not assigned to any node)

	enableTrace	bool
	indentTrace	int
}

type ByPosition []*ast.CommentGroup

func (b ByPosition) Len() int {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:36
	_go_fuzz_dep_.CoverTab[122155]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:36
	return len(b)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:36
	// _ = "end of CoverTab[122155]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:36
}
func (b ByPosition) Swap(i, j int) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:37
	_go_fuzz_dep_.CoverTab[122156]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:37
	b[i], b[j] = b[j], b[i]
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:37
	// _ = "end of CoverTab[122156]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:37
}
func (b ByPosition) Less(i, j int) bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:38
	_go_fuzz_dep_.CoverTab[122157]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:38
	return b[i].Pos().Before(b[j].Pos())
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:38
	// _ = "end of CoverTab[122157]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:38
}

// collectComments comments all standalone comments which are not lead or line
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:40
// comment
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:42
func (p *printer) collectComments(node ast.Node) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:42
	_go_fuzz_dep_.CoverTab[122158]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:45
	ast.Walk(node, func(nn ast.Node) (ast.Node, bool) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:45
		_go_fuzz_dep_.CoverTab[122163]++
													switch t := nn.(type) {
		case *ast.File:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:47
			_go_fuzz_dep_.CoverTab[122165]++
														p.comments = t.Comments
														return nn, false
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:49
			// _ = "end of CoverTab[122165]"
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:50
		// _ = "end of CoverTab[122163]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:50
		_go_fuzz_dep_.CoverTab[122164]++
													return nn, true
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:51
		// _ = "end of CoverTab[122164]"
	})
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:52
	// _ = "end of CoverTab[122158]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:52
	_go_fuzz_dep_.CoverTab[122159]++

												standaloneComments := make(map[token.Pos]*ast.CommentGroup, 0)
												for _, c := range p.comments {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:55
		_go_fuzz_dep_.CoverTab[122166]++
													standaloneComments[c.Pos()] = c
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:56
		// _ = "end of CoverTab[122166]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:57
	// _ = "end of CoverTab[122159]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:57
	_go_fuzz_dep_.CoverTab[122160]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:62
	ast.Walk(node, func(nn ast.Node) (ast.Node, bool) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:62
		_go_fuzz_dep_.CoverTab[122167]++
													switch t := nn.(type) {
		case *ast.LiteralType:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:64
			_go_fuzz_dep_.CoverTab[122169]++
														if t.LeadComment != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:65
				_go_fuzz_dep_.CoverTab[122173]++
															for _, comment := range t.LeadComment.List {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:66
					_go_fuzz_dep_.CoverTab[122174]++
																if _, ok := standaloneComments[comment.Pos()]; ok {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:67
						_go_fuzz_dep_.CoverTab[122175]++
																	delete(standaloneComments, comment.Pos())
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:68
						// _ = "end of CoverTab[122175]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:69
						_go_fuzz_dep_.CoverTab[122176]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:69
						// _ = "end of CoverTab[122176]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:69
					}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:69
					// _ = "end of CoverTab[122174]"
				}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:70
				// _ = "end of CoverTab[122173]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:71
				_go_fuzz_dep_.CoverTab[122177]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:71
				// _ = "end of CoverTab[122177]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:71
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:71
			// _ = "end of CoverTab[122169]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:71
			_go_fuzz_dep_.CoverTab[122170]++

														if t.LineComment != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:73
				_go_fuzz_dep_.CoverTab[122178]++
															for _, comment := range t.LineComment.List {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:74
					_go_fuzz_dep_.CoverTab[122179]++
																if _, ok := standaloneComments[comment.Pos()]; ok {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:75
						_go_fuzz_dep_.CoverTab[122180]++
																	delete(standaloneComments, comment.Pos())
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:76
						// _ = "end of CoverTab[122180]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:77
						_go_fuzz_dep_.CoverTab[122181]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:77
						// _ = "end of CoverTab[122181]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:77
					}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:77
					// _ = "end of CoverTab[122179]"
				}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:78
				// _ = "end of CoverTab[122178]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:79
				_go_fuzz_dep_.CoverTab[122182]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:79
				// _ = "end of CoverTab[122182]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:79
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:79
			// _ = "end of CoverTab[122170]"
		case *ast.ObjectItem:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:80
			_go_fuzz_dep_.CoverTab[122171]++
														if t.LeadComment != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:81
				_go_fuzz_dep_.CoverTab[122183]++
															for _, comment := range t.LeadComment.List {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:82
					_go_fuzz_dep_.CoverTab[122184]++
																if _, ok := standaloneComments[comment.Pos()]; ok {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:83
						_go_fuzz_dep_.CoverTab[122185]++
																	delete(standaloneComments, comment.Pos())
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:84
						// _ = "end of CoverTab[122185]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:85
						_go_fuzz_dep_.CoverTab[122186]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:85
						// _ = "end of CoverTab[122186]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:85
					}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:85
					// _ = "end of CoverTab[122184]"
				}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:86
				// _ = "end of CoverTab[122183]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:87
				_go_fuzz_dep_.CoverTab[122187]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:87
				// _ = "end of CoverTab[122187]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:87
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:87
			// _ = "end of CoverTab[122171]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:87
			_go_fuzz_dep_.CoverTab[122172]++

														if t.LineComment != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:89
				_go_fuzz_dep_.CoverTab[122188]++
															for _, comment := range t.LineComment.List {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:90
					_go_fuzz_dep_.CoverTab[122189]++
																if _, ok := standaloneComments[comment.Pos()]; ok {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:91
						_go_fuzz_dep_.CoverTab[122190]++
																	delete(standaloneComments, comment.Pos())
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:92
						// _ = "end of CoverTab[122190]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:93
						_go_fuzz_dep_.CoverTab[122191]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:93
						// _ = "end of CoverTab[122191]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:93
					}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:93
					// _ = "end of CoverTab[122189]"
				}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:94
				// _ = "end of CoverTab[122188]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:95
				_go_fuzz_dep_.CoverTab[122192]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:95
				// _ = "end of CoverTab[122192]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:95
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:95
			// _ = "end of CoverTab[122172]"
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:96
		// _ = "end of CoverTab[122167]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:96
		_go_fuzz_dep_.CoverTab[122168]++

													return nn, true
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:98
		// _ = "end of CoverTab[122168]"
	})
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:99
	// _ = "end of CoverTab[122160]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:99
	_go_fuzz_dep_.CoverTab[122161]++

												for _, c := range standaloneComments {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:101
		_go_fuzz_dep_.CoverTab[122193]++
													p.standaloneComments = append(p.standaloneComments, c)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:102
		// _ = "end of CoverTab[122193]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:103
	// _ = "end of CoverTab[122161]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:103
	_go_fuzz_dep_.CoverTab[122162]++

												sort.Sort(ByPosition(p.standaloneComments))
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:105
	// _ = "end of CoverTab[122162]"
}

// output prints creates b printable HCL output and returns it.
func (p *printer) output(n interface{}) []byte {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:109
	_go_fuzz_dep_.CoverTab[122194]++
												var buf bytes.Buffer

												switch t := n.(type) {
	case *ast.File:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:113
		_go_fuzz_dep_.CoverTab[122196]++

													defer un(trace(p, "File"))
													return p.output(t.Node)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:116
		// _ = "end of CoverTab[122196]"
	case *ast.ObjectList:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:117
		_go_fuzz_dep_.CoverTab[122197]++
													defer un(trace(p, "ObjectList"))

													var index int
													for {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:121
			_go_fuzz_dep_.CoverTab[122204]++
			// Determine the location of the next actual non-comment
			// item. If we're at the end, the next item is at "infinity"
			var nextItem token.Pos
			if index != len(t.Items) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:125
				_go_fuzz_dep_.CoverTab[122209]++
															nextItem = t.Items[index].Pos()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:126
				// _ = "end of CoverTab[122209]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:127
				_go_fuzz_dep_.CoverTab[122210]++
															nextItem = token.Pos{Offset: infinity, Line: infinity}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:128
				// _ = "end of CoverTab[122210]"
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:129
			// _ = "end of CoverTab[122204]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:129
			_go_fuzz_dep_.CoverTab[122205]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:133
			for _, c := range p.standaloneComments {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:133
				_go_fuzz_dep_.CoverTab[122211]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:136
				printed := false
				newlinePrinted := false
				for _, comment := range c.List {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:138
					_go_fuzz_dep_.CoverTab[122213]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:143
					if comment.Pos().After(p.prev) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:143
						_go_fuzz_dep_.CoverTab[122214]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:143
						return comment.Pos().Before(nextItem)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:143
						// _ = "end of CoverTab[122214]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:143
					}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:143
						_go_fuzz_dep_.CoverTab[122215]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:148
						if !newlinePrinted && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:148
							_go_fuzz_dep_.CoverTab[122217]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:148
							return p.prev.IsValid()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:148
							// _ = "end of CoverTab[122217]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:148
						}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:148
							_go_fuzz_dep_.CoverTab[122218]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:148
							return index == len(t.Items)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:148
							// _ = "end of CoverTab[122218]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:148
						}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:148
							_go_fuzz_dep_.CoverTab[122219]++
																		buf.Write([]byte{newline, newline})
																		newlinePrinted = true
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:150
							// _ = "end of CoverTab[122219]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:151
							_go_fuzz_dep_.CoverTab[122220]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:151
							// _ = "end of CoverTab[122220]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:151
						}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:151
						// _ = "end of CoverTab[122215]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:151
						_go_fuzz_dep_.CoverTab[122216]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:154
						buf.WriteString(comment.Text)
																	buf.WriteByte(newline)

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:158
						printed = true
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:158
						// _ = "end of CoverTab[122216]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:159
						_go_fuzz_dep_.CoverTab[122221]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:159
						// _ = "end of CoverTab[122221]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:159
					}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:159
					// _ = "end of CoverTab[122213]"
				}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:160
				// _ = "end of CoverTab[122211]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:160
				_go_fuzz_dep_.CoverTab[122212]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:165
				if printed && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:165
					_go_fuzz_dep_.CoverTab[122222]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:165
					return index != len(t.Items)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:165
					// _ = "end of CoverTab[122222]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:165
				}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:165
					_go_fuzz_dep_.CoverTab[122223]++
																buf.WriteByte(newline)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:166
					// _ = "end of CoverTab[122223]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:167
					_go_fuzz_dep_.CoverTab[122224]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:167
					// _ = "end of CoverTab[122224]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:167
				}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:167
				// _ = "end of CoverTab[122212]"
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:168
			// _ = "end of CoverTab[122205]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:168
			_go_fuzz_dep_.CoverTab[122206]++

														if index == len(t.Items) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:170
				_go_fuzz_dep_.CoverTab[122225]++
															break
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:171
				// _ = "end of CoverTab[122225]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:172
				_go_fuzz_dep_.CoverTab[122226]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:172
				// _ = "end of CoverTab[122226]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:172
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:172
			// _ = "end of CoverTab[122206]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:172
			_go_fuzz_dep_.CoverTab[122207]++

														buf.Write(p.output(t.Items[index]))
														if index != len(t.Items)-1 {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:175
				_go_fuzz_dep_.CoverTab[122227]++

															buf.WriteByte(newline)

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:191
				current := t.Items[index]
				next := t.Items[index+1]
				if next.Pos().Line != t.Items[index].Pos().Line+1 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:193
					_go_fuzz_dep_.CoverTab[122228]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:193
					return !p.isSingleLineObject(next)
																// _ = "end of CoverTab[122228]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:194
				}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:194
					_go_fuzz_dep_.CoverTab[122229]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:194
					return !p.isSingleLineObject(current)
																// _ = "end of CoverTab[122229]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:195
				}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:195
					_go_fuzz_dep_.CoverTab[122230]++
																buf.WriteByte(newline)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:196
					// _ = "end of CoverTab[122230]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:197
					_go_fuzz_dep_.CoverTab[122231]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:197
					// _ = "end of CoverTab[122231]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:197
				}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:197
				// _ = "end of CoverTab[122227]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:198
				_go_fuzz_dep_.CoverTab[122232]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:198
				// _ = "end of CoverTab[122232]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:198
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:198
			// _ = "end of CoverTab[122207]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:198
			_go_fuzz_dep_.CoverTab[122208]++
														index++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:199
			// _ = "end of CoverTab[122208]"
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:200
		// _ = "end of CoverTab[122197]"
	case *ast.ObjectKey:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:201
		_go_fuzz_dep_.CoverTab[122198]++
													buf.WriteString(t.Token.Text)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:202
		// _ = "end of CoverTab[122198]"
	case *ast.ObjectItem:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:203
		_go_fuzz_dep_.CoverTab[122199]++
													p.prev = t.Pos()
													buf.Write(p.objectItem(t))
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:205
		// _ = "end of CoverTab[122199]"
	case *ast.LiteralType:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:206
		_go_fuzz_dep_.CoverTab[122200]++
													buf.Write(p.literalType(t))
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:207
		// _ = "end of CoverTab[122200]"
	case *ast.ListType:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:208
		_go_fuzz_dep_.CoverTab[122201]++
													buf.Write(p.list(t))
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:209
		// _ = "end of CoverTab[122201]"
	case *ast.ObjectType:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:210
		_go_fuzz_dep_.CoverTab[122202]++
													buf.Write(p.objectType(t))
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:211
		// _ = "end of CoverTab[122202]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:212
		_go_fuzz_dep_.CoverTab[122203]++
													fmt.Printf(" unknown type: %T\n", n)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:213
		// _ = "end of CoverTab[122203]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:214
	// _ = "end of CoverTab[122194]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:214
	_go_fuzz_dep_.CoverTab[122195]++

												return buf.Bytes()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:216
	// _ = "end of CoverTab[122195]"
}

func (p *printer) literalType(lit *ast.LiteralType) []byte {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:219
	_go_fuzz_dep_.CoverTab[122233]++
												result := []byte(lit.Token.Text)
												switch lit.Token.Type {
	case token.HEREDOC:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:222
		_go_fuzz_dep_.CoverTab[122235]++

													if result[len(result)-1] == '\n' {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:224
			_go_fuzz_dep_.CoverTab[122239]++
														result = result[:len(result)-1]
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:225
			// _ = "end of CoverTab[122239]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:226
			_go_fuzz_dep_.CoverTab[122240]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:226
			// _ = "end of CoverTab[122240]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:226
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:226
		// _ = "end of CoverTab[122235]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:226
		_go_fuzz_dep_.CoverTab[122236]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:229
		result = p.heredocIndent(result)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:229
		// _ = "end of CoverTab[122236]"
	case token.STRING:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:230
		_go_fuzz_dep_.CoverTab[122237]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:233
		if bytes.IndexRune(result, '\n') >= 0 {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:233
			_go_fuzz_dep_.CoverTab[122241]++
														result = p.heredocIndent(result)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:234
			// _ = "end of CoverTab[122241]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:235
			_go_fuzz_dep_.CoverTab[122242]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:235
			// _ = "end of CoverTab[122242]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:235
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:235
		// _ = "end of CoverTab[122237]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:235
	default:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:235
		_go_fuzz_dep_.CoverTab[122238]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:235
		// _ = "end of CoverTab[122238]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:236
	// _ = "end of CoverTab[122233]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:236
	_go_fuzz_dep_.CoverTab[122234]++

												return result
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:238
	// _ = "end of CoverTab[122234]"
}

// objectItem returns the printable HCL form of an object item. An object type
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:241
// starts with one/multiple keys and has a value. The value might be of any
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:241
// type.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:244
func (p *printer) objectItem(o *ast.ObjectItem) []byte {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:244
	_go_fuzz_dep_.CoverTab[122243]++
												defer un(trace(p, fmt.Sprintf("ObjectItem: %s", o.Keys[0].Token.Text)))
												var buf bytes.Buffer

												if o.LeadComment != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:248
		_go_fuzz_dep_.CoverTab[122248]++
													for _, comment := range o.LeadComment.List {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:249
			_go_fuzz_dep_.CoverTab[122249]++
														buf.WriteString(comment.Text)
														buf.WriteByte(newline)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:251
			// _ = "end of CoverTab[122249]"
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:252
		// _ = "end of CoverTab[122248]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:253
		_go_fuzz_dep_.CoverTab[122250]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:253
		// _ = "end of CoverTab[122250]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:253
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:253
	// _ = "end of CoverTab[122243]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:253
	_go_fuzz_dep_.CoverTab[122244]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:256
	if o.LineComment != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:256
		_go_fuzz_dep_.CoverTab[122251]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:256
		return o.Val.Pos().Line != o.Keys[0].Pos().Line
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:256
		// _ = "end of CoverTab[122251]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:256
	}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:256
		_go_fuzz_dep_.CoverTab[122252]++
													for _, comment := range o.LineComment.List {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:257
			_go_fuzz_dep_.CoverTab[122253]++
														buf.WriteString(comment.Text)
														buf.WriteByte(newline)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:259
			// _ = "end of CoverTab[122253]"
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:260
		// _ = "end of CoverTab[122252]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:261
		_go_fuzz_dep_.CoverTab[122254]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:261
		// _ = "end of CoverTab[122254]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:261
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:261
	// _ = "end of CoverTab[122244]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:261
	_go_fuzz_dep_.CoverTab[122245]++

												for i, k := range o.Keys {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:263
		_go_fuzz_dep_.CoverTab[122255]++
													buf.WriteString(k.Token.Text)
													buf.WriteByte(blank)

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:268
		if o.Assign.IsValid() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:268
			_go_fuzz_dep_.CoverTab[122256]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:268
			return i == len(o.Keys)-1
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:268
			// _ = "end of CoverTab[122256]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:268
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:268
			_go_fuzz_dep_.CoverTab[122257]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:268
			return len(o.Keys) == 1
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:268
			// _ = "end of CoverTab[122257]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:268
		}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:268
			_go_fuzz_dep_.CoverTab[122258]++
														buf.WriteString("=")
														buf.WriteByte(blank)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:270
			// _ = "end of CoverTab[122258]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:271
			_go_fuzz_dep_.CoverTab[122259]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:271
			// _ = "end of CoverTab[122259]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:271
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:271
		// _ = "end of CoverTab[122255]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:272
	// _ = "end of CoverTab[122245]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:272
	_go_fuzz_dep_.CoverTab[122246]++

												buf.Write(p.output(o.Val))

												if o.LineComment != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:276
		_go_fuzz_dep_.CoverTab[122260]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:276
		return o.Val.Pos().Line == o.Keys[0].Pos().Line
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:276
		// _ = "end of CoverTab[122260]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:276
	}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:276
		_go_fuzz_dep_.CoverTab[122261]++
													buf.WriteByte(blank)
													for _, comment := range o.LineComment.List {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:278
			_go_fuzz_dep_.CoverTab[122262]++
														buf.WriteString(comment.Text)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:279
			// _ = "end of CoverTab[122262]"
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:280
		// _ = "end of CoverTab[122261]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:281
		_go_fuzz_dep_.CoverTab[122263]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:281
		// _ = "end of CoverTab[122263]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:281
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:281
	// _ = "end of CoverTab[122246]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:281
	_go_fuzz_dep_.CoverTab[122247]++

												return buf.Bytes()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:283
	// _ = "end of CoverTab[122247]"
}

// objectType returns the printable HCL form of an object type. An object type
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:286
// begins with a brace and ends with a brace.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:288
func (p *printer) objectType(o *ast.ObjectType) []byte {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:288
	_go_fuzz_dep_.CoverTab[122264]++
												defer un(trace(p, "ObjectType"))
												var buf bytes.Buffer
												buf.WriteString("{")

												var index int
												var nextItem token.Pos
												var commented, newlinePrinted bool
												for {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:296
		_go_fuzz_dep_.CoverTab[122266]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:299
		if index != len(o.List.Items) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:299
			_go_fuzz_dep_.CoverTab[122274]++
														nextItem = o.List.Items[index].Pos()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:300
			// _ = "end of CoverTab[122274]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:301
			_go_fuzz_dep_.CoverTab[122275]++
														nextItem = o.Rbrace
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:302
			// _ = "end of CoverTab[122275]"
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:303
		// _ = "end of CoverTab[122266]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:303
		_go_fuzz_dep_.CoverTab[122267]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:307
		for _, c := range p.standaloneComments {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:307
			_go_fuzz_dep_.CoverTab[122276]++
														printed := false
														var lastCommentPos token.Pos
														for _, comment := range c.List {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:310
				_go_fuzz_dep_.CoverTab[122278]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:315
				if comment.Pos().After(p.prev) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:315
					_go_fuzz_dep_.CoverTab[122279]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:315
					return comment.Pos().Before(nextItem)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:315
					// _ = "end of CoverTab[122279]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:315
				}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:315
					_go_fuzz_dep_.CoverTab[122280]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:318
					if !newlinePrinted {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:318
						_go_fuzz_dep_.CoverTab[122283]++
																	newlinePrinted = true
																	buf.WriteByte(newline)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:320
						// _ = "end of CoverTab[122283]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:321
						_go_fuzz_dep_.CoverTab[122284]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:321
						// _ = "end of CoverTab[122284]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:321
					}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:321
					// _ = "end of CoverTab[122280]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:321
					_go_fuzz_dep_.CoverTab[122281]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:324
					if index > 0 {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:324
						_go_fuzz_dep_.CoverTab[122285]++
																	commented = true
																	buf.WriteByte(newline)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:326
						// _ = "end of CoverTab[122285]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:327
						_go_fuzz_dep_.CoverTab[122286]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:327
						// _ = "end of CoverTab[122286]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:327
					}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:327
					// _ = "end of CoverTab[122281]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:327
					_go_fuzz_dep_.CoverTab[122282]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:330
					lastCommentPos = comment.Pos()

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:333
					buf.Write(p.indent(p.heredocIndent([]byte(comment.Text))))

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:336
					printed = true
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:336
					// _ = "end of CoverTab[122282]"

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:343
				} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:343
					_go_fuzz_dep_.CoverTab[122287]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:343
					// _ = "end of CoverTab[122287]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:343
				}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:343
				// _ = "end of CoverTab[122278]"
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:344
			// _ = "end of CoverTab[122276]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:344
			_go_fuzz_dep_.CoverTab[122277]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:347
			if printed {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:347
				_go_fuzz_dep_.CoverTab[122288]++

															buf.WriteByte(newline)

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:354
				if nextItem != o.Rbrace && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:354
					_go_fuzz_dep_.CoverTab[122289]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:354
					return nextItem.Line != lastCommentPos.Line+1
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:354
					// _ = "end of CoverTab[122289]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:354
				}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:354
					_go_fuzz_dep_.CoverTab[122290]++
																buf.WriteByte(newline)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:355
					// _ = "end of CoverTab[122290]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:356
					_go_fuzz_dep_.CoverTab[122291]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:356
					// _ = "end of CoverTab[122291]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:356
				}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:356
				// _ = "end of CoverTab[122288]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:357
				_go_fuzz_dep_.CoverTab[122292]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:357
				// _ = "end of CoverTab[122292]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:357
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:357
			// _ = "end of CoverTab[122277]"
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:358
		// _ = "end of CoverTab[122267]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:358
		_go_fuzz_dep_.CoverTab[122268]++

													if index == len(o.List.Items) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:360
			_go_fuzz_dep_.CoverTab[122293]++
														p.prev = o.Rbrace
														break
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:362
			// _ = "end of CoverTab[122293]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:363
			_go_fuzz_dep_.CoverTab[122294]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:363
			// _ = "end of CoverTab[122294]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:363
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:363
		// _ = "end of CoverTab[122268]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:363
		_go_fuzz_dep_.CoverTab[122269]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:368
		if !newlinePrinted {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:368
			_go_fuzz_dep_.CoverTab[122295]++
														buf.WriteByte(newline)
														newlinePrinted = true
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:370
			// _ = "end of CoverTab[122295]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:371
			_go_fuzz_dep_.CoverTab[122296]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:371
			// _ = "end of CoverTab[122296]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:371
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:371
		// _ = "end of CoverTab[122269]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:371
		_go_fuzz_dep_.CoverTab[122270]++

		// check if we have adjacent one liner items. If yes we'll going to align
		// the comments.
		var aligned []*ast.ObjectItem
		for _, item := range o.List.Items[index:] {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:376
			_go_fuzz_dep_.CoverTab[122297]++

														if len(o.List.Items) == 1 {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:378
				_go_fuzz_dep_.CoverTab[122304]++
															break
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:379
				// _ = "end of CoverTab[122304]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:380
				_go_fuzz_dep_.CoverTab[122305]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:380
				// _ = "end of CoverTab[122305]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:380
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:380
			// _ = "end of CoverTab[122297]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:380
			_go_fuzz_dep_.CoverTab[122298]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:385
			cur := lines(string(p.objectItem(item)))
			if cur > 2 {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:386
				_go_fuzz_dep_.CoverTab[122306]++
															break
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:387
				// _ = "end of CoverTab[122306]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:388
				_go_fuzz_dep_.CoverTab[122307]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:388
				// _ = "end of CoverTab[122307]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:388
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:388
			// _ = "end of CoverTab[122298]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:388
			_go_fuzz_dep_.CoverTab[122299]++

														curPos := item.Pos()

														nextPos := token.Pos{}
														if index != len(o.List.Items)-1 {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:393
				_go_fuzz_dep_.CoverTab[122308]++
															nextPos = o.List.Items[index+1].Pos()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:394
				// _ = "end of CoverTab[122308]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:395
				_go_fuzz_dep_.CoverTab[122309]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:395
				// _ = "end of CoverTab[122309]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:395
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:395
			// _ = "end of CoverTab[122299]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:395
			_go_fuzz_dep_.CoverTab[122300]++

														prevPos := token.Pos{}
														if index != 0 {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:398
				_go_fuzz_dep_.CoverTab[122310]++
															prevPos = o.List.Items[index-1].Pos()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:399
				// _ = "end of CoverTab[122310]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:400
				_go_fuzz_dep_.CoverTab[122311]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:400
				// _ = "end of CoverTab[122311]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:400
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:400
			// _ = "end of CoverTab[122300]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:400
			_go_fuzz_dep_.CoverTab[122301]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:407
			if curPos.Line+1 == nextPos.Line {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:407
				_go_fuzz_dep_.CoverTab[122312]++
															aligned = append(aligned, item)
															index++
															continue
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:410
				// _ = "end of CoverTab[122312]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:411
				_go_fuzz_dep_.CoverTab[122313]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:411
				// _ = "end of CoverTab[122313]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:411
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:411
			// _ = "end of CoverTab[122301]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:411
			_go_fuzz_dep_.CoverTab[122302]++

														if curPos.Line-1 == prevPos.Line {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:413
				_go_fuzz_dep_.CoverTab[122314]++
															aligned = append(aligned, item)
															index++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:419
				if curPos.Line+1 != nextPos.Line {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:419
					_go_fuzz_dep_.CoverTab[122316]++
																break
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:420
					// _ = "end of CoverTab[122316]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:421
					_go_fuzz_dep_.CoverTab[122317]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:421
					// _ = "end of CoverTab[122317]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:421
				}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:421
				// _ = "end of CoverTab[122314]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:421
				_go_fuzz_dep_.CoverTab[122315]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:422
				// _ = "end of CoverTab[122315]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:423
				_go_fuzz_dep_.CoverTab[122318]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:423
				// _ = "end of CoverTab[122318]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:423
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:423
			// _ = "end of CoverTab[122302]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:423
			_go_fuzz_dep_.CoverTab[122303]++

														break
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:425
			// _ = "end of CoverTab[122303]"
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:426
		// _ = "end of CoverTab[122270]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:426
		_go_fuzz_dep_.CoverTab[122271]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:431
		if !commented && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:431
			_go_fuzz_dep_.CoverTab[122319]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:431
			return index != len(aligned)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:431
			// _ = "end of CoverTab[122319]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:431
		}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:431
			_go_fuzz_dep_.CoverTab[122320]++
														buf.WriteByte(newline)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:432
			// _ = "end of CoverTab[122320]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:433
			_go_fuzz_dep_.CoverTab[122321]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:433
			// _ = "end of CoverTab[122321]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:433
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:433
		// _ = "end of CoverTab[122271]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:433
		_go_fuzz_dep_.CoverTab[122272]++

													if len(aligned) >= 1 {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:435
			_go_fuzz_dep_.CoverTab[122322]++
														p.prev = aligned[len(aligned)-1].Pos()

														items := p.alignedItems(aligned)
														buf.Write(p.indent(items))
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:439
			// _ = "end of CoverTab[122322]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:440
			_go_fuzz_dep_.CoverTab[122323]++
														p.prev = o.List.Items[index].Pos()

														buf.Write(p.indent(p.objectItem(o.List.Items[index])))
														index++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:444
			// _ = "end of CoverTab[122323]"
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:445
		// _ = "end of CoverTab[122272]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:445
		_go_fuzz_dep_.CoverTab[122273]++

													buf.WriteByte(newline)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:447
		// _ = "end of CoverTab[122273]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:448
	// _ = "end of CoverTab[122264]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:448
	_go_fuzz_dep_.CoverTab[122265]++

												buf.WriteString("}")
												return buf.Bytes()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:451
	// _ = "end of CoverTab[122265]"
}

func (p *printer) alignedItems(items []*ast.ObjectItem) []byte {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:454
	_go_fuzz_dep_.CoverTab[122324]++
												var buf bytes.Buffer

	// find the longest key and value length, needed for alignment
	var longestKeyLen int	// longest key length
	var longestValLen int	// longest value length
	for _, item := range items {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:460
		_go_fuzz_dep_.CoverTab[122327]++
													key := len(item.Keys[0].Token.Text)
													val := len(p.output(item.Val))

													if key > longestKeyLen {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:464
			_go_fuzz_dep_.CoverTab[122329]++
														longestKeyLen = key
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:465
			// _ = "end of CoverTab[122329]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:466
			_go_fuzz_dep_.CoverTab[122330]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:466
			// _ = "end of CoverTab[122330]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:466
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:466
		// _ = "end of CoverTab[122327]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:466
		_go_fuzz_dep_.CoverTab[122328]++

													if val > longestValLen {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:468
			_go_fuzz_dep_.CoverTab[122331]++
														longestValLen = val
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:469
			// _ = "end of CoverTab[122331]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:470
			_go_fuzz_dep_.CoverTab[122332]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:470
			// _ = "end of CoverTab[122332]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:470
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:470
		// _ = "end of CoverTab[122328]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:471
	// _ = "end of CoverTab[122324]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:471
	_go_fuzz_dep_.CoverTab[122325]++

												for i, item := range items {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:473
		_go_fuzz_dep_.CoverTab[122333]++
													if item.LeadComment != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:474
			_go_fuzz_dep_.CoverTab[122337]++
														for _, comment := range item.LeadComment.List {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:475
				_go_fuzz_dep_.CoverTab[122338]++
															buf.WriteString(comment.Text)
															buf.WriteByte(newline)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:477
				// _ = "end of CoverTab[122338]"
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:478
			// _ = "end of CoverTab[122337]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:479
			_go_fuzz_dep_.CoverTab[122339]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:479
			// _ = "end of CoverTab[122339]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:479
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:479
		// _ = "end of CoverTab[122333]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:479
		_go_fuzz_dep_.CoverTab[122334]++

													for i, k := range item.Keys {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:481
			_go_fuzz_dep_.CoverTab[122340]++
														keyLen := len(k.Token.Text)
														buf.WriteString(k.Token.Text)
														for i := 0; i < longestKeyLen-keyLen+1; i++ {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:484
				_go_fuzz_dep_.CoverTab[122342]++
															buf.WriteByte(blank)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:485
				// _ = "end of CoverTab[122342]"
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:486
			// _ = "end of CoverTab[122340]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:486
			_go_fuzz_dep_.CoverTab[122341]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:489
			if i == len(item.Keys)-1 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:489
				_go_fuzz_dep_.CoverTab[122343]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:489
				return len(item.Keys) == 1
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:489
				// _ = "end of CoverTab[122343]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:489
			}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:489
				_go_fuzz_dep_.CoverTab[122344]++
															buf.WriteString("=")
															buf.WriteByte(blank)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:491
				// _ = "end of CoverTab[122344]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:492
				_go_fuzz_dep_.CoverTab[122345]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:492
				// _ = "end of CoverTab[122345]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:492
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:492
			// _ = "end of CoverTab[122341]"
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:493
		// _ = "end of CoverTab[122334]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:493
		_go_fuzz_dep_.CoverTab[122335]++

													val := p.output(item.Val)
													valLen := len(val)
													buf.Write(val)

													if item.Val.Pos().Line == item.Keys[0].Pos().Line && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:499
			_go_fuzz_dep_.CoverTab[122346]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:499
			return item.LineComment != nil
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:499
			// _ = "end of CoverTab[122346]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:499
		}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:499
			_go_fuzz_dep_.CoverTab[122347]++
														for i := 0; i < longestValLen-valLen+1; i++ {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:500
				_go_fuzz_dep_.CoverTab[122349]++
															buf.WriteByte(blank)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:501
				// _ = "end of CoverTab[122349]"
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:502
			// _ = "end of CoverTab[122347]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:502
			_go_fuzz_dep_.CoverTab[122348]++

														for _, comment := range item.LineComment.List {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:504
				_go_fuzz_dep_.CoverTab[122350]++
															buf.WriteString(comment.Text)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:505
				// _ = "end of CoverTab[122350]"
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:506
			// _ = "end of CoverTab[122348]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:507
			_go_fuzz_dep_.CoverTab[122351]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:507
			// _ = "end of CoverTab[122351]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:507
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:507
		// _ = "end of CoverTab[122335]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:507
		_go_fuzz_dep_.CoverTab[122336]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:510
		if i != len(items)-1 {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:510
			_go_fuzz_dep_.CoverTab[122352]++
														buf.WriteByte(newline)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:511
			// _ = "end of CoverTab[122352]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:512
			_go_fuzz_dep_.CoverTab[122353]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:512
			// _ = "end of CoverTab[122353]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:512
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:512
		// _ = "end of CoverTab[122336]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:513
	// _ = "end of CoverTab[122325]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:513
	_go_fuzz_dep_.CoverTab[122326]++

												return buf.Bytes()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:515
	// _ = "end of CoverTab[122326]"
}

// list returns the printable HCL form of an list type.
func (p *printer) list(l *ast.ListType) []byte {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:519
	_go_fuzz_dep_.CoverTab[122354]++
												if p.isSingleLineList(l) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:520
		_go_fuzz_dep_.CoverTab[122358]++
													return p.singleLineList(l)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:521
		// _ = "end of CoverTab[122358]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:522
		_go_fuzz_dep_.CoverTab[122359]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:522
		// _ = "end of CoverTab[122359]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:522
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:522
	// _ = "end of CoverTab[122354]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:522
	_go_fuzz_dep_.CoverTab[122355]++

												var buf bytes.Buffer
												buf.WriteString("[")
												buf.WriteByte(newline)

												var longestLine int
												for _, item := range l.List {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:529
		_go_fuzz_dep_.CoverTab[122360]++

													if lit, ok := item.(*ast.LiteralType); ok {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:531
			_go_fuzz_dep_.CoverTab[122361]++
														lineLen := len(lit.Token.Text)
														if lineLen > longestLine {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:533
				_go_fuzz_dep_.CoverTab[122362]++
															longestLine = lineLen
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:534
				// _ = "end of CoverTab[122362]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:535
				_go_fuzz_dep_.CoverTab[122363]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:535
				// _ = "end of CoverTab[122363]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:535
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:535
			// _ = "end of CoverTab[122361]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:536
			_go_fuzz_dep_.CoverTab[122364]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:536
			// _ = "end of CoverTab[122364]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:536
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:536
		// _ = "end of CoverTab[122360]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:537
	// _ = "end of CoverTab[122355]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:537
	_go_fuzz_dep_.CoverTab[122356]++

												haveEmptyLine := false
												for i, item := range l.List {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:540
		_go_fuzz_dep_.CoverTab[122365]++

													leadComment := false
													if lit, ok := item.(*ast.LiteralType); ok && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:543
			_go_fuzz_dep_.CoverTab[122369]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:543
			return lit.LeadComment != nil
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:543
			// _ = "end of CoverTab[122369]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:543
		}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:543
			_go_fuzz_dep_.CoverTab[122370]++
														leadComment = true

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:548
			if !haveEmptyLine && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:548
				_go_fuzz_dep_.CoverTab[122372]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:548
				return i != 0
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:548
				// _ = "end of CoverTab[122372]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:548
			}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:548
				_go_fuzz_dep_.CoverTab[122373]++
															buf.WriteByte(newline)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:549
				// _ = "end of CoverTab[122373]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:550
				_go_fuzz_dep_.CoverTab[122374]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:550
				// _ = "end of CoverTab[122374]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:550
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:550
			// _ = "end of CoverTab[122370]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:550
			_go_fuzz_dep_.CoverTab[122371]++

														for _, comment := range lit.LeadComment.List {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:552
				_go_fuzz_dep_.CoverTab[122375]++
															buf.Write(p.indent([]byte(comment.Text)))
															buf.WriteByte(newline)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:554
				// _ = "end of CoverTab[122375]"
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:555
			// _ = "end of CoverTab[122371]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:556
			_go_fuzz_dep_.CoverTab[122376]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:556
			// _ = "end of CoverTab[122376]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:556
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:556
		// _ = "end of CoverTab[122365]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:556
		_go_fuzz_dep_.CoverTab[122366]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:559
		val := p.output(item)
													curLen := len(val)
													buf.Write(p.indent(val))

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:565
		comma := []byte{','}
		if lit, ok := item.(*ast.LiteralType); ok && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:566
			_go_fuzz_dep_.CoverTab[122377]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:566
			return lit.Token.Type == token.HEREDOC
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:566
			// _ = "end of CoverTab[122377]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:566
		}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:566
			_go_fuzz_dep_.CoverTab[122378]++
														buf.WriteByte(newline)
														comma = p.indent(comma)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:568
			// _ = "end of CoverTab[122378]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:569
			_go_fuzz_dep_.CoverTab[122379]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:569
			// _ = "end of CoverTab[122379]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:569
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:569
		// _ = "end of CoverTab[122366]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:569
		_go_fuzz_dep_.CoverTab[122367]++

													buf.Write(comma)

													if lit, ok := item.(*ast.LiteralType); ok && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:573
			_go_fuzz_dep_.CoverTab[122380]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:573
			return lit.LineComment != nil
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:573
			// _ = "end of CoverTab[122380]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:573
		}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:573
			_go_fuzz_dep_.CoverTab[122381]++

														buf.WriteByte(blank)
														for i := 0; i < longestLine-curLen; i++ {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:576
				_go_fuzz_dep_.CoverTab[122383]++
															buf.WriteByte(blank)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:577
				// _ = "end of CoverTab[122383]"
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:578
			// _ = "end of CoverTab[122381]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:578
			_go_fuzz_dep_.CoverTab[122382]++

														for _, comment := range lit.LineComment.List {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:580
				_go_fuzz_dep_.CoverTab[122384]++
															buf.WriteString(comment.Text)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:581
				// _ = "end of CoverTab[122384]"
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:582
			// _ = "end of CoverTab[122382]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:583
			_go_fuzz_dep_.CoverTab[122385]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:583
			// _ = "end of CoverTab[122385]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:583
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:583
		// _ = "end of CoverTab[122367]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:583
		_go_fuzz_dep_.CoverTab[122368]++

													buf.WriteByte(newline)

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:589
		haveEmptyLine = leadComment && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:589
			_go_fuzz_dep_.CoverTab[122386]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:589
			return i != len(l.List)-1
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:589
			// _ = "end of CoverTab[122386]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:589
		}()
		if haveEmptyLine {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:590
			_go_fuzz_dep_.CoverTab[122387]++
														buf.WriteByte(newline)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:591
			// _ = "end of CoverTab[122387]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:592
			_go_fuzz_dep_.CoverTab[122388]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:592
			// _ = "end of CoverTab[122388]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:592
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:592
		// _ = "end of CoverTab[122368]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:593
	// _ = "end of CoverTab[122356]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:593
	_go_fuzz_dep_.CoverTab[122357]++

												buf.WriteString("]")
												return buf.Bytes()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:596
	// _ = "end of CoverTab[122357]"
}

// isSingleLineList returns true if:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:599
// * they were previously formatted entirely on one line
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:599
// * they consist entirely of literals
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:599
// * there are either no heredoc strings or the list has exactly one element
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:599
// * there are no line comments
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:604
func (printer) isSingleLineList(l *ast.ListType) bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:604
	_go_fuzz_dep_.CoverTab[122389]++
												for _, item := range l.List {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:605
		_go_fuzz_dep_.CoverTab[122391]++
													if item.Pos().Line != l.Lbrack.Line {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:606
			_go_fuzz_dep_.CoverTab[122395]++
														return false
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:607
			// _ = "end of CoverTab[122395]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:608
			_go_fuzz_dep_.CoverTab[122396]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:608
			// _ = "end of CoverTab[122396]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:608
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:608
		// _ = "end of CoverTab[122391]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:608
		_go_fuzz_dep_.CoverTab[122392]++

													lit, ok := item.(*ast.LiteralType)
													if !ok {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:611
			_go_fuzz_dep_.CoverTab[122397]++
														return false
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:612
			// _ = "end of CoverTab[122397]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:613
			_go_fuzz_dep_.CoverTab[122398]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:613
			// _ = "end of CoverTab[122398]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:613
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:613
		// _ = "end of CoverTab[122392]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:613
		_go_fuzz_dep_.CoverTab[122393]++

													if lit.Token.Type == token.HEREDOC && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:615
			_go_fuzz_dep_.CoverTab[122399]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:615
			return len(l.List) != 1
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:615
			// _ = "end of CoverTab[122399]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:615
		}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:615
			_go_fuzz_dep_.CoverTab[122400]++
														return false
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:616
			// _ = "end of CoverTab[122400]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:617
			_go_fuzz_dep_.CoverTab[122401]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:617
			// _ = "end of CoverTab[122401]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:617
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:617
		// _ = "end of CoverTab[122393]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:617
		_go_fuzz_dep_.CoverTab[122394]++

													if lit.LineComment != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:619
			_go_fuzz_dep_.CoverTab[122402]++
														return false
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:620
			// _ = "end of CoverTab[122402]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:621
			_go_fuzz_dep_.CoverTab[122403]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:621
			// _ = "end of CoverTab[122403]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:621
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:621
		// _ = "end of CoverTab[122394]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:622
	// _ = "end of CoverTab[122389]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:622
	_go_fuzz_dep_.CoverTab[122390]++

												return true
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:624
	// _ = "end of CoverTab[122390]"
}

// singleLineList prints a simple single line list.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:627
// For a definition of "simple", see isSingleLineList above.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:629
func (p *printer) singleLineList(l *ast.ListType) []byte {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:629
	_go_fuzz_dep_.CoverTab[122404]++
												buf := &bytes.Buffer{}

												buf.WriteString("[")
												for i, item := range l.List {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:633
		_go_fuzz_dep_.CoverTab[122406]++
													if i != 0 {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:634
			_go_fuzz_dep_.CoverTab[122408]++
														buf.WriteString(", ")
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:635
			// _ = "end of CoverTab[122408]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:636
			_go_fuzz_dep_.CoverTab[122409]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:636
			// _ = "end of CoverTab[122409]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:636
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:636
		// _ = "end of CoverTab[122406]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:636
		_go_fuzz_dep_.CoverTab[122407]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:639
		buf.Write(p.output(item))

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:642
		if lit, ok := item.(*ast.LiteralType); ok && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:642
			_go_fuzz_dep_.CoverTab[122410]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:642
			return lit.Token.Type == token.HEREDOC
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:642
			// _ = "end of CoverTab[122410]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:642
		}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:642
			_go_fuzz_dep_.CoverTab[122411]++
														buf.WriteByte(newline)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:643
			// _ = "end of CoverTab[122411]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:644
			_go_fuzz_dep_.CoverTab[122412]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:644
			// _ = "end of CoverTab[122412]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:644
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:644
		// _ = "end of CoverTab[122407]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:645
	// _ = "end of CoverTab[122404]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:645
	_go_fuzz_dep_.CoverTab[122405]++

												buf.WriteString("]")
												return buf.Bytes()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:648
	// _ = "end of CoverTab[122405]"
}

// indent indents the lines of the given buffer for each non-empty line
func (p *printer) indent(buf []byte) []byte {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:652
	_go_fuzz_dep_.CoverTab[122413]++
												var prefix []byte
												if p.cfg.SpacesWidth != 0 {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:654
		_go_fuzz_dep_.CoverTab[122416]++
													for i := 0; i < p.cfg.SpacesWidth; i++ {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:655
			_go_fuzz_dep_.CoverTab[122417]++
														prefix = append(prefix, blank)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:656
			// _ = "end of CoverTab[122417]"
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:657
		// _ = "end of CoverTab[122416]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:658
		_go_fuzz_dep_.CoverTab[122418]++
													prefix = []byte{tab}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:659
		// _ = "end of CoverTab[122418]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:660
	// _ = "end of CoverTab[122413]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:660
	_go_fuzz_dep_.CoverTab[122414]++

												var res []byte
												bol := true
												for _, c := range buf {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:664
		_go_fuzz_dep_.CoverTab[122419]++
													if bol && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:665
			_go_fuzz_dep_.CoverTab[122421]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:665
			return c != '\n'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:665
			// _ = "end of CoverTab[122421]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:665
		}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:665
			_go_fuzz_dep_.CoverTab[122422]++
														res = append(res, prefix...)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:666
			// _ = "end of CoverTab[122422]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:667
			_go_fuzz_dep_.CoverTab[122423]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:667
			// _ = "end of CoverTab[122423]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:667
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:667
		// _ = "end of CoverTab[122419]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:667
		_go_fuzz_dep_.CoverTab[122420]++

													res = append(res, c)
													bol = c == '\n'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:670
		// _ = "end of CoverTab[122420]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:671
	// _ = "end of CoverTab[122414]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:671
	_go_fuzz_dep_.CoverTab[122415]++
												return res
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:672
	// _ = "end of CoverTab[122415]"
}

// unindent removes all the indentation from the tombstoned lines
func (p *printer) unindent(buf []byte) []byte {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:676
	_go_fuzz_dep_.CoverTab[122424]++
												var res []byte
												for i := 0; i < len(buf); i++ {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:678
		_go_fuzz_dep_.CoverTab[122426]++
													skip := len(buf)-i <= len(unindent)
													if !skip {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:680
			_go_fuzz_dep_.CoverTab[122430]++
														skip = !bytes.Equal(unindent, buf[i:i+len(unindent)])
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:681
			// _ = "end of CoverTab[122430]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:682
			_go_fuzz_dep_.CoverTab[122431]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:682
			// _ = "end of CoverTab[122431]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:682
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:682
		// _ = "end of CoverTab[122426]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:682
		_go_fuzz_dep_.CoverTab[122427]++
													if skip {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:683
			_go_fuzz_dep_.CoverTab[122432]++
														res = append(res, buf[i])
														continue
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:685
			// _ = "end of CoverTab[122432]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:686
			_go_fuzz_dep_.CoverTab[122433]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:686
			// _ = "end of CoverTab[122433]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:686
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:686
		// _ = "end of CoverTab[122427]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:686
		_go_fuzz_dep_.CoverTab[122428]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:690
		for j := len(res) - 1; j >= 0; j-- {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:690
			_go_fuzz_dep_.CoverTab[122434]++
														if res[j] == '\n' {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:691
				_go_fuzz_dep_.CoverTab[122436]++
															break
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:692
				// _ = "end of CoverTab[122436]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:693
				_go_fuzz_dep_.CoverTab[122437]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:693
				// _ = "end of CoverTab[122437]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:693
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:693
			// _ = "end of CoverTab[122434]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:693
			_go_fuzz_dep_.CoverTab[122435]++

														res = res[:j]
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:695
			// _ = "end of CoverTab[122435]"
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:696
		// _ = "end of CoverTab[122428]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:696
		_go_fuzz_dep_.CoverTab[122429]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:699
		i += len(unindent) - 1
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:699
		// _ = "end of CoverTab[122429]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:700
	// _ = "end of CoverTab[122424]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:700
	_go_fuzz_dep_.CoverTab[122425]++

												return res
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:702
	// _ = "end of CoverTab[122425]"
}

// heredocIndent marks all the 2nd and further lines as unindentable
func (p *printer) heredocIndent(buf []byte) []byte {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:706
	_go_fuzz_dep_.CoverTab[122438]++
												var res []byte
												bol := false
												for _, c := range buf {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:709
		_go_fuzz_dep_.CoverTab[122440]++
													if bol && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:710
			_go_fuzz_dep_.CoverTab[122442]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:710
			return c != '\n'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:710
			// _ = "end of CoverTab[122442]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:710
		}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:710
			_go_fuzz_dep_.CoverTab[122443]++
														res = append(res, unindent...)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:711
			// _ = "end of CoverTab[122443]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:712
			_go_fuzz_dep_.CoverTab[122444]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:712
			// _ = "end of CoverTab[122444]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:712
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:712
		// _ = "end of CoverTab[122440]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:712
		_go_fuzz_dep_.CoverTab[122441]++
													res = append(res, c)
													bol = c == '\n'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:714
		// _ = "end of CoverTab[122441]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:715
	// _ = "end of CoverTab[122438]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:715
	_go_fuzz_dep_.CoverTab[122439]++
												return res
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:716
	// _ = "end of CoverTab[122439]"
}

// isSingleLineObject tells whether the given object item is a single
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:719
// line object such as "obj {}".
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:719
//
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:719
// A single line object:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:719
//
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:719
//   - has no lead comments (hence multi-line)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:719
//   - has no assignment
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:719
//   - has no values in the stanza (within {})
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:728
func (p *printer) isSingleLineObject(val *ast.ObjectItem) bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:728
	_go_fuzz_dep_.CoverTab[122445]++

												if val.LeadComment != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:730
		_go_fuzz_dep_.CoverTab[122449]++
													return false
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:731
		// _ = "end of CoverTab[122449]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:732
		_go_fuzz_dep_.CoverTab[122450]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:732
		// _ = "end of CoverTab[122450]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:732
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:732
	// _ = "end of CoverTab[122445]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:732
	_go_fuzz_dep_.CoverTab[122446]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:735
	if val.Assign.IsValid() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:735
		_go_fuzz_dep_.CoverTab[122451]++
													return false
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:736
		// _ = "end of CoverTab[122451]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:737
		_go_fuzz_dep_.CoverTab[122452]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:737
		// _ = "end of CoverTab[122452]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:737
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:737
	// _ = "end of CoverTab[122446]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:737
	_go_fuzz_dep_.CoverTab[122447]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:740
	ot, ok := val.Val.(*ast.ObjectType)
	if !ok {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:741
		_go_fuzz_dep_.CoverTab[122453]++
													return false
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:742
		// _ = "end of CoverTab[122453]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:743
		_go_fuzz_dep_.CoverTab[122454]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:743
		// _ = "end of CoverTab[122454]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:743
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:743
	// _ = "end of CoverTab[122447]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:743
	_go_fuzz_dep_.CoverTab[122448]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:746
	return len(ot.List.Items) == 0
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:746
	// _ = "end of CoverTab[122448]"
}

func lines(txt string) int {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:749
	_go_fuzz_dep_.CoverTab[122455]++
												endline := 1
												for i := 0; i < len(txt); i++ {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:751
		_go_fuzz_dep_.CoverTab[122457]++
													if txt[i] == '\n' {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:752
			_go_fuzz_dep_.CoverTab[122458]++
														endline++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:753
			// _ = "end of CoverTab[122458]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:754
			_go_fuzz_dep_.CoverTab[122459]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:754
			// _ = "end of CoverTab[122459]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:754
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:754
		// _ = "end of CoverTab[122457]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:755
	// _ = "end of CoverTab[122455]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:755
	_go_fuzz_dep_.CoverTab[122456]++
												return endline
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:756
	// _ = "end of CoverTab[122456]"
}

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:762
func (p *printer) printTrace(a ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:762
	_go_fuzz_dep_.CoverTab[122460]++
												if !p.enableTrace {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:763
		_go_fuzz_dep_.CoverTab[122463]++
													return
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:764
		// _ = "end of CoverTab[122463]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:765
		_go_fuzz_dep_.CoverTab[122464]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:765
		// _ = "end of CoverTab[122464]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:765
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:765
	// _ = "end of CoverTab[122460]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:765
	_go_fuzz_dep_.CoverTab[122461]++

												const dots = ". . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . "
												const n = len(dots)
												i := 2 * p.indentTrace
												for i > n {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:770
		_go_fuzz_dep_.CoverTab[122465]++
													fmt.Print(dots)
													i -= n
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:772
		// _ = "end of CoverTab[122465]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:773
	// _ = "end of CoverTab[122461]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:773
	_go_fuzz_dep_.CoverTab[122462]++

												fmt.Print(dots[0:i])
												fmt.Println(a...)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:776
	// _ = "end of CoverTab[122462]"
}

func trace(p *printer, msg string) *printer {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:779
	_go_fuzz_dep_.CoverTab[122466]++
												p.printTrace(msg, "(")
												p.indentTrace++
												return p
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:782
	// _ = "end of CoverTab[122466]"
}

// Usage pattern: defer un(trace(p, "..."))
func un(p *printer) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:786
	_go_fuzz_dep_.CoverTab[122467]++
												p.indentTrace--
												p.printTrace(")")
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:788
	// _ = "end of CoverTab[122467]"
}

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:789
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/nodes.go:789
var _ = _go_fuzz_dep_.CoverTab
