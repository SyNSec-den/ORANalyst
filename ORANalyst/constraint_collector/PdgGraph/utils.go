package PdgGraph

import (
	"fmt"
	"go/token"
	"go/types"
	"reflect"
	"strconv"
	"strings"

	git "github.com/go-git/go-git/v5"
	"golang.org/x/tools/go/callgraph"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/types/typeutil"
)

// Extract info from ssa.Instruction
func ExtractInstrInfo(fset *token.FileSet,
	instr ssa.Instruction) (text string, file string, pos int) {
	dest := ""
	_, ok := reflect.TypeOf(instr).MethodByName("Name")
	if ok {
		r := reflect.ValueOf(instr).MethodByName("Name").Call([]reflect.Value{})[0]
		dest += r.String()
	}
	if dest == "" {
		text = instr.String()
	} else {
		text = dest + " = " + instr.String()
	}
	if fset == nil {
		return text, "", 0
	}
	fsetPos := fset.Position(instr.Pos())
	file = fsetPos.Filename
	posStr := buildPos(fsetPos)
	posInt := 0
	if strings.Contains(posStr, ":") {
		posStrSp := strings.Split(posStr, ":")[0]
		posInt, _ = strconv.Atoi(posStrSp)
	} else {
		return text, file, 0
	}
	return text, file, posInt
}

// Extract info from ssa.BasicBlock
func extractBlockInfo(fset *token.FileSet,
	block *ssa.BasicBlock) (name string, file string, pos int) {
	if block == nil || fset == nil {
		return "", "", 0
	}

	name = GetBlockName(fset, block)
	parentPos := fset.Position(block.Parent().Pos())
	file = parentPos.Filename
	posStr := buildPos(parentPos)
	posInt := 0
	if strings.Contains(posStr, ":") {
		posStrSp := strings.Split(posStr, ":")[0]
		posInt, _ = strconv.Atoi(posStrSp)
	} else {
		return name, file, 0
	}
	return name, file, posInt
}

// Extract info from ssa.Function
func extractFnInfo(fset *token.FileSet,
	fn *ssa.Function) (name string, file string, pos int) {
	if fn == nil || fset == nil {
		return "", "", 0
	}

	// name = GetFnName(fset, fn)
	name = fn.Name()
	fnPos := fset.Position(fn.Pos())
	file = fnPos.Filename
	posStr := buildPos(fnPos)
	posInt := 0
	if strings.Contains(posStr, ":") {
		posStrSp := strings.Split(posStr, ":")[0]
		posInt, _ = strconv.Atoi(posStrSp)
	} else {
		return name, file, 0
	}
	return name, file, posInt
}

func buildPos(p token.Position) string {
	return fmt.Sprintf("%v:%v", p.Line, p.Column)
}

func GetBlockName(fset *token.FileSet, b *ssa.BasicBlock) string {
	if b == nil {
		return ""
	}
	if b.Parent() == nil {
		return fmt.Sprintf(":::%v", b.Index)
	}
	file := ""
	if fset != nil {
		file = fset.Position(b.Parent().Pos()).Filename
	}
	// pkg:filename:function:index
	return fmt.Sprintf("%v:%v:%v:%v",
		b.Parent().Pkg.Pkg.Name(), file, b.Parent().Name(), b.Index)
}

func ParseBlockName(bName string) (pkgName string, fileName string, funcName string, blockNum string) {
	bStrip := strings.Split(bName, ":")
	if len(bStrip) != 4 {
		return "", "", "", ""
	}
	return bStrip[0], bStrip[1], bStrip[2], bStrip[3]
}

func GetFnName(fset *token.FileSet, fn *ssa.Function) string {
	if fn == nil {
		return ""
	}
	file := ""
	if fset != nil {
		file = fset.Position(fn.Pos()).Filename
	}
	// pkg:function
	return fmt.Sprintf("%v:%v",
		file, fn.Name())
}

// this doesn't really belong here but might as well put it with the rest of the PackageTuple stuff
func MakeFallbackVersion(pkgDir string) string {
	repo, err := git.PlainOpenWithOptions(pkgDir, &git.PlainOpenOptions{DetectDotGit: true})
	if err != nil {
		return pkgDir
	}
	remotes, err := repo.Remotes()
	if err != nil || len(remotes) == 0 {
		return pkgDir
	}
	head, err := repo.Head()
	if err != nil {
		return pkgDir
	}
	url := remotes[0].Config().URLs[0]
	url = strings.TrimPrefix("https://", url)
	hash := head.Hash().String()[:8]
	return fmt.Sprintf("%s@%s", url, hash)
}

// Process SSA variable
func ParseSSAVar(v ssa.Value) (isVar bool, isConst bool, name string, t string) {
	if v == nil {
		return false, false, "", ""
	}
	if strings.Contains(v.Name(), ":") {
		nameSplit := strings.Split(v.Name(), ":")
		return false, false, nameSplit[len(nameSplit)-1], v.Type().String()
	}
	if v.Parent() == nil {
		return true, true, v.Name(), v.Type().String()
	}
	return true, false, v.Name(), v.Type().String()
}

func InstrsToString(instrs *[]ssa.Instruction) string {
	if instrs == nil {
		return "[]"
	}
	res := "["
	for i, instr := range *instrs {
		if i != 0 {
			res += ", "
		}
		res += instr.String()
	}
	res += "]"
	return res
}

// chaLazyCallees returns a function that maps a call site (in a function in fns)
// to its callees within fns.
//
// The resulting function is not concurrency safe.
func chaLazyCallees(fns map[*ssa.Function]bool) func(site ssa.CallInstruction) []*ssa.Function {
	// funcsBySig contains all functions, keyed by signature.  It is
	// the effective set of address-taken functions used to resolve
	// a dynamic call of a particular signature.
	var funcsBySig typeutil.Map // value is []*ssa.Function

	// methodsByName contains all methods,
	// grouped by name for efficient lookup.
	// (methodsById would be better but not every SSA method has a go/types ID.)
	methodsByName := make(map[string][]*ssa.Function)

	// An imethod represents an interface method I.m.
	// (There's no go/types object for it;
	// a *types.Func may be shared by many interfaces due to interface embedding.)
	type imethod struct {
		I  *types.Interface
		id string
	}
	// methodsMemo records, for every abstract method call I.m on
	// interface type I, the set of concrete methods C.m of all
	// types C that satisfy interface I.
	//
	// Abstract methods may be shared by several interfaces,
	// hence we must pass I explicitly, not guess from m.
	//
	// methodsMemo is just a cache, so it needn't be a typeutil.Map.
	methodsMemo := make(map[imethod][]*ssa.Function)
	lookupMethods := func(I *types.Interface, m *types.Func) []*ssa.Function {
		id := m.Id()
		methods, ok := methodsMemo[imethod{I, id}]
		if !ok {
			for _, f := range methodsByName[m.Name()] {
				C := f.Signature.Recv().Type() // named or *named
				if types.Implements(C, I) {
					methods = append(methods, f)
				}
			}
			methodsMemo[imethod{I, id}] = methods
		}
		return methods
	}

	for f := range fns {
		if f.Signature.Recv() == nil {
			// Package initializers can never be address-taken.
			if f.Name() == "init" && f.Synthetic == "package initializer" {
				continue
			}
			funcs, _ := funcsBySig.At(f.Signature).([]*ssa.Function)
			funcs = append(funcs, f)
			funcsBySig.Set(f.Signature, funcs)
		} else {
			methodsByName[f.Name()] = append(methodsByName[f.Name()], f)
		}
	}

	return func(site ssa.CallInstruction) []*ssa.Function {
		call := site.Common()
		if call.IsInvoke() {
			tiface := call.Value.Type().Underlying().(*types.Interface)
			return lookupMethods(tiface, call.Method)
		} else if g := call.StaticCallee(); g != nil {
			return []*ssa.Function{g}
		} else if _, ok := call.Value.(*ssa.Builtin); !ok {
			fns, _ := funcsBySig.At(call.Signature()).([]*ssa.Function)
			return fns
		}
		return nil
	}
}

// getMainPackages returns the main packages to analyze.
// Each resulting package is named "main" and has a main function.
func getMainPackages(pkgs []*ssa.Package) ([]*ssa.Package, error) {
	var mains []*ssa.Package
	for _, p := range pkgs {
		if p != nil && p.Pkg.Name() == "main" && p.Func("main") != nil {
			mains = append(mains, p)
		}
	}
	if len(mains) == 0 {
		return nil, fmt.Errorf("no main packages")
	}
	return mains, nil
}

// TODO: implement this to allow choice of callgraph algorithms
func callgraphToCallsite(cg *callgraph.Graph) func(site ssa.CallInstruction) []*ssa.Function {
	panic("TODO: not implemented")
}

func possibleChan(t string) bool {
	tokenize := strings.Split(t, " ")
	if len(tokenize) < 2 {
		return false
	}
	if strings.ReplaceAll(tokenize[0], "*", "") != "chan" {
		return false
	}
	return true
}

func matchType(t1 string, t2 string) bool {
	if t1 == "" || t2 == "" {
		return false
	}
	t1Map := tokenizeType(t1)
	t2Map := tokenizeType(t2)
	for k := range t1Map {
		if !t2Map[k] {
			return false
		}
	}
	for k := range t2Map {
		if !t1Map[k] {
			return false
		}
	}
	return true
}

func tokenizeType(t string) map[string]bool {
	t = strings.ReplaceAll(t, "*", "")
	t = strings.ReplaceAll(t, "&", "")
	tSplit := strings.Split(t, " ")
	tMap := make(map[string]bool)
	for _, token := range tSplit {
		tMap[token] = true
	}
	return tMap
}

func TopologicalSort(f *Function) []*Block {

	WhiteBlocks := make(map[string]*Block)
	GreyBlocks := make(map[string]*Block)
	BlackBlocks := make(map[string]*Block)

	startTimes := make(map[string]int)
	endTimes := make(map[string]int)

	var ReverseTopSortedBlocks []*Block
	var TopSortedBlocks []*Block

	var dfsStack []*Block

	for _, b := range f.ContainsBlock {
		WhiteBlocks[b.FullName()] = b
	}

	currentTime := 0

	startNode := f.CfgGraph.StartNode

	dfsStack = append(dfsStack, startNode)

	for len(dfsStack) != 0 {
		currNode := dfsStack[len(dfsStack)-1]
		_, ok := WhiteBlocks[currNode.FullName()]

		if ok {
			delete(WhiteBlocks, currNode.FullName())
			GreyBlocks[currNode.FullName()] = currNode
			startTimes[currNode.FullName()] = currentTime
			currentTime += 1

			for _, b := range currNode.Successors {
				_, ok2 := WhiteBlocks[b.FullName()]
				if ok2 {
					dfsStack = append(dfsStack, b)
				}
			}

		} else {
			_, ok2 := GreyBlocks[currNode.FullName()]
			if ok2 {
				delete(GreyBlocks, currNode.FullName())
				BlackBlocks[currNode.FullName()] = currNode
				endTimes[currNode.FullName()] = currentTime
				currentTime += 1
				ReverseTopSortedBlocks = append(ReverseTopSortedBlocks, currNode)
			}
			dfsStack = dfsStack[:len(dfsStack)-1]
		}
	}

	for len(WhiteBlocks) != 0 {
		for _, b := range WhiteBlocks {
			dfsStack = append(dfsStack, b)

			for len(dfsStack) != 0 {
				currNode := dfsStack[len(dfsStack)-1]
				_, ok := WhiteBlocks[currNode.FullName()]

				if ok {
					delete(WhiteBlocks, currNode.FullName())
					GreyBlocks[currNode.FullName()] = currNode
					startTimes[currNode.FullName()] = currentTime
					currentTime += 1

					for _, b := range currNode.Successors {
						_, ok2 := WhiteBlocks[b.FullName()]
						if ok2 {
							dfsStack = append(dfsStack, b)
						}
					}

				} else {
					_, ok2 := GreyBlocks[currNode.FullName()]
					if ok2 {
						delete(GreyBlocks, currNode.FullName())
						BlackBlocks[currNode.FullName()] = currNode
						endTimes[currNode.FullName()] = currentTime
						currentTime += 1
						ReverseTopSortedBlocks = append(ReverseTopSortedBlocks, currNode)
					}
					dfsStack = dfsStack[:len(dfsStack)-1]
				}
			}

		}
	}

	for i := len(ReverseTopSortedBlocks) - 1; i >= 0; i-- {
		TopSortedBlocks = append(TopSortedBlocks, ReverseTopSortedBlocks[i])
	}

	//for _, b := range TopSortedBlocks {
	//	fmt.Printf("Topsorted Block %s\n", b.FullName())
	//	for _, cb := range b.ControlDepBlocks {
	//		fmt.Printf("\t Controldep block %s ", cb.Name)
	//	}
	//	fmt.Printf("\n")
	//}

	return TopSortedBlocks
}
