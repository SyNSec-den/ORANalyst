package main

import (
	"go/ast"
	"go/token"
	"strconv"

	. "github.com/dvyukov/go-fuzz/internal/go-fuzz-types"
)

// TODO:  break, continue, goto
// Do we need to treat edge and block counters differently?

// EdgeVisitor struct
type EdgeVisitor struct {
	file *File
}

type LoopVisitor struct {
	file *File
}

// Visit method for LoopVisitor
func (lv *LoopVisitor) Visit(n ast.Node) ast.Visitor {
	switch node := n.(type) {
	case *ast.BlockStmt:
		node.List = lv.file.addLoopCounters(node.Pos(), node.End(), node.List)
	}

	return lv
}

// Visit method for EdgeVisitor
func (ev *EdgeVisitor) Visit(n ast.Node) ast.Visitor {
	switch node := n.(type) {
	case *ast.IfStmt:
		// Instrument the 'if' statement
		stmt, _ := ev.file.newEdgeCounter(node.Pos(), node.End(), 0)
		node.Body.List = append([]ast.Stmt{stmt}, node.Body.List...)

		// Instrument 'else if' and 'else' statements
		elseStmt := node.Else
		for elseStmt != nil {
			switch es := elseStmt.(type) {
			case *ast.IfStmt:
				// 'else if' statement
				stmt, _ := ev.file.newEdgeCounter(es.Pos(), es.End(), 0)
				es.Body.List = append([]ast.Stmt{stmt}, es.Body.List...)
				elseStmt = es.Else
			case *ast.BlockStmt:
				// 'else' statement
				stmt, _ := ev.file.newEdgeCounter(es.Pos(), es.End(), 0)
				es.List = append([]ast.Stmt{stmt}, es.List...)
				elseStmt = nil
			default:
				elseStmt = nil
			}
		}
	case *ast.SwitchStmt:
		// Instrument the 'switch' statement
		for _, stmt := range node.Body.List {
			caseClause, ok := stmt.(*ast.CaseClause)
			if !ok {
				continue
			}
			counterStmt, _ := ev.file.newEdgeCounter(caseClause.Pos(), caseClause.End(), 0)
			caseClause.Body = append([]ast.Stmt{counterStmt}, caseClause.Body...)
		}

	case *ast.TypeSwitchStmt:
		// Instrument the 'type switch' statement
		for _, stmt := range node.Body.List {
			caseClause, ok := stmt.(*ast.CaseClause)
			if !ok {
				continue
			}
			counterStmt, _ := ev.file.newEdgeCounter(caseClause.Pos(), caseClause.End(), 0)
			caseClause.Body = append([]ast.Stmt{counterStmt}, caseClause.Body...)
		}
	}

	return ev
}

func (f *File) addLoopCounters(pos, blockEnd token.Pos, list []ast.Stmt) []ast.Stmt {
	var newList []ast.Stmt

	for i := 0; i < len(list); i++ {
		stmt := list[i]
		switch typedStmt := stmt.(type) {
		case *ast.ForStmt:
			loopExecId := genCounter(loopRun)

			// Boolean variable declaration for first run check
			firstRunInit := &ast.AssignStmt{
				Lhs: []ast.Expr{
					&ast.IndexExpr{
						X: &ast.SelectorExpr{
							X:   ast.NewIdent(fuzzdepPkg),
							Sel: ast.NewIdent("CoverTab"),
						},
						Index: &ast.BasicLit{
							Kind:  token.INT,
							Value: strconv.Itoa(loopExecId),
						},
					},
				},
				Tok: token.ASSIGN,
				Rhs: []ast.Expr{
					&ast.BasicLit{
						Kind:  token.INT,
						Value: "0",
					},
				},
			}

			// // Create the if statement inside the loop
			ifStmt := &ast.IfStmt{
				Cond: &ast.BinaryExpr{
					X: &ast.IndexExpr{
						X: &ast.SelectorExpr{
							X:   ast.NewIdent(fuzzdepPkg),
							Sel: ast.NewIdent("CoverTab"),
						},
						Index: &ast.BasicLit{
							Kind:  token.INT,
							Value: strconv.Itoa(loopExecId),
						},
					},
					Op: token.EQL,
					Y: &ast.BasicLit{
						Kind:  token.INT,
						Value: "0",
					},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{f.newCounterStmt()},
				},
				Else: &ast.BlockStmt{
					List: []ast.Stmt{f.newCounterStmt()},
				},
			}

			// Set firstRun variable to false after first run
			setFirstRunFalse := &ast.AssignStmt{
				Lhs: []ast.Expr{
					&ast.IndexExpr{
						X: &ast.SelectorExpr{
							X:   ast.NewIdent(fuzzdepPkg),
							Sel: ast.NewIdent("CoverTab"),
						},
						Index: &ast.BasicLit{
							Kind:  token.INT,
							Value: strconv.Itoa(loopExecId),
						},
					},
				},
				Tok: token.ASSIGN,
				Rhs: []ast.Expr{&ast.BasicLit{Kind: token.INT, Value: "1"}},
			}

			// Create the post-loop counter
			postLoopStmt := &ast.IfStmt{
				Cond: &ast.BinaryExpr{
					X: &ast.IndexExpr{
						X: &ast.SelectorExpr{
							X:   ast.NewIdent(fuzzdepPkg),
							Sel: ast.NewIdent("CoverTab"),
						},
						Index: &ast.BasicLit{
							Kind:  token.INT,
							Value: strconv.Itoa(loopExecId),
						},
					},
					Op: token.EQL,
					Y: &ast.BasicLit{
						Kind:  token.INT,
						Value: "0",
					},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{f.newCounterStmt()},
				},
				Else: &ast.BlockStmt{
					List: []ast.Stmt{f.newCounterStmt()},
				},
			}

			newList = append(newList, firstRunInit)
			typedStmt.Body.List = append([]ast.Stmt{ifStmt, setFirstRunFalse}, typedStmt.Body.List...)
			newList = append(newList, typedStmt)
			if typedStmt.Cond != nil {
				newList = append(newList, postLoopStmt)
			}
		case *ast.RangeStmt:
			loopExecId := genCounter(loopRun)

			// Boolean variable declaration for first run check
			firstRunInit := &ast.AssignStmt{
				Lhs: []ast.Expr{
					&ast.IndexExpr{
						X: &ast.SelectorExpr{
							X:   ast.NewIdent(fuzzdepPkg),
							Sel: ast.NewIdent("CoverTab"),
						},
						Index: &ast.BasicLit{
							Kind:  token.INT,
							Value: strconv.Itoa(loopExecId),
						},
					},
				},
				Tok: token.ASSIGN,
				Rhs: []ast.Expr{
					&ast.BasicLit{
						Kind:  token.INT,
						Value: "0",
					},
				},
			}

			// // Create the if statement inside the loop
			ifStmt := &ast.IfStmt{
				Cond: &ast.BinaryExpr{
					X: &ast.IndexExpr{
						X: &ast.SelectorExpr{
							X:   ast.NewIdent(fuzzdepPkg),
							Sel: ast.NewIdent("CoverTab"),
						},
						Index: &ast.BasicLit{
							Kind:  token.INT,
							Value: strconv.Itoa(loopExecId),
						},
					},
					Op: token.EQL,
					Y: &ast.BasicLit{
						Kind:  token.INT,
						Value: "0",
					},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{f.newCounterStmt()},
				},
				Else: &ast.BlockStmt{
					List: []ast.Stmt{f.newCounterStmt()},
				},
			}

			// Set firstRun variable to false after first run
			setFirstRunFalse := &ast.AssignStmt{
				Lhs: []ast.Expr{
					&ast.IndexExpr{
						X: &ast.SelectorExpr{
							X:   ast.NewIdent(fuzzdepPkg),
							Sel: ast.NewIdent("CoverTab"),
						},
						Index: &ast.BasicLit{
							Kind:  token.INT,
							Value: strconv.Itoa(loopExecId),
						},
					},
				},
				Tok: token.ASSIGN,
				Rhs: []ast.Expr{&ast.BasicLit{Kind: token.INT, Value: "1"}},
			}

			// Create the post-loop counter
			postLoopStmt := &ast.IfStmt{
				Cond: &ast.BinaryExpr{
					X: &ast.IndexExpr{
						X: &ast.SelectorExpr{
							X:   ast.NewIdent(fuzzdepPkg),
							Sel: ast.NewIdent("CoverTab"),
						},
						Index: &ast.BasicLit{
							Kind:  token.INT,
							Value: strconv.Itoa(loopExecId),
						},
					},
					Op: token.EQL,
					Y: &ast.BasicLit{
						Kind:  token.INT,
						Value: "0",
					},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{f.newCounterStmt()},
				},
				Else: &ast.BlockStmt{
					List: []ast.Stmt{f.newCounterStmt()},
				},
			}

			newList = append(newList, firstRunInit)
			typedStmt.Body.List = append([]ast.Stmt{ifStmt, setFirstRunFalse}, typedStmt.Body.List...)
			newList = append(newList, typedStmt)
			newList = append(newList, postLoopStmt)

		default:
			newList = append(newList, stmt)
		}
	}

	return newList
}

func (f *File) newEdgeCounter(start, end token.Pos, numStmt int) (ast.Stmt, int) {
	cnt := genCounter(edge)

	if f.blocks != nil {
		s := f.fset.Position(start)
		e := f.fset.Position(end)
		*f.blocks = append(*f.blocks, CoverBlock{cnt, f.fullName, s.Line, s.Column, e.Line, e.Column, numStmt})
	}

	idx := &ast.BasicLit{
		Kind:  token.INT,
		Value: strconv.Itoa(cnt),
	}
	counter := &ast.IndexExpr{
		X: &ast.SelectorExpr{
			X:   ast.NewIdent(fuzzdepPkg),
			Sel: ast.NewIdent("CoverTab"),
		},
		Index: idx,
	}
	return &ast.IncDecStmt{
		X:   counter,
		Tok: token.INC,
	}, cnt
}

func (f *File) newCounterStmt() ast.Stmt {
	cnt := genCounter(edge)
	// Convert the counter ID to a string
	counterIDStr := strconv.Itoa(cnt)

	// Create the AST node for the counter index
	idx := &ast.BasicLit{
		Kind:  token.INT,
		Value: counterIDStr,
	}

	// Create the AST node for accessing the counter array
	counter := &ast.IndexExpr{
		X: &ast.SelectorExpr{
			X:   ast.NewIdent(fuzzdepPkg),
			Sel: ast.NewIdent("CoverTab"),
		},
		Index: idx,
	}

	// Create the AST node for incrementing the counter
	counterStmt := &ast.IncDecStmt{
		X:   counter,
		Tok: token.INC,
	}

	return counterStmt
}
