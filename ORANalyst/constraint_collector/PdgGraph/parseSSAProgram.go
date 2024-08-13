package PdgGraph

import (
	"constraint_collector/coordination"
	"fmt"
	"go/parser"
	"go/token"
	"go/types"
	"reflect"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"golang.org/x/tools/go/callgraph"
	"golang.org/x/tools/go/callgraph/cha"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

type ProgParser struct {
	ProgressBar bool
	ModuleName  string

	// values that will be built up during package walking, and used in callgraph processing
	pkgMap       map[string]*Package
	graphFuncMap map[*ssa.Function]*Function
	fsetMap      map[string]*token.FileSet
	blockMap     map[string]*Block
	fset         token.FileSet
	instrMap     map[ssa.Instruction]*Instruction
	varMap       map[string]*Variable
	calleesOf    func(site ssa.CallInstruction) []*ssa.Function
	processedPkg map[string]bool
	globalVarMap map[string]map[string]*Variable
	chanMap      map[*ssa.Value]*Variable
	// chanMap      *om.OrderedMap[*ssa.Value, *Variable]
}

func CanPoint(T types.Type) bool {
	switch T := T.(type) {
	case *types.Named:
		if obj := T.Obj(); obj.Name() == "Value" && obj.Pkg().Path() == "reflect" {
			return true // treat reflect.Value like interface{}
		}
		return CanPoint(T.Underlying())
	case *types.Pointer, *types.Interface, *types.Map, *types.Chan, *types.Signature, *types.Slice:
		return true
	}

	return false // array struct tuple builtin basic
}

func NewProgParser(moduleName string) ProgParser {
	return ProgParser{
		ProgressBar:  true,
		ModuleName:   moduleName,
		pkgMap:       map[string]*Package{},
		graphFuncMap: map[*ssa.Function]*Function{},
		fsetMap:      map[string]*token.FileSet{},
		blockMap:     map[string]*Block{},
		instrMap:     map[ssa.Instruction]*Instruction{},
		varMap:       map[string]*Variable{},
		processedPkg: map[string]bool{},
		globalVarMap: map[string]map[string]*Variable{},
		chanMap:      map[*ssa.Value]*Variable{},
		// chanMap:      om.New[*ssa.Value, *Variable](),
	}
}

func (pp *ProgParser) ProcessPackage(pkgDir string, progInfo *ProgramInfo) *ProgramInfo {
	// Minimal load to check if this is already done and get imports
	config := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedModule | packages.NeedImports |
			packages.NeedDeps | packages.NeedTypes | packages.NeedSyntax | packages.NeedTypesInfo,
		// packages.LoadAllSyntax,
		Logf: func(format string, args ...interface{}) {
			logrus.Tracef(format, args...)
		},
		Dir: pkgDir,
	}

	pkgs, err := packages.Load(config, "./...")
	if err != nil {
		logrus.Errorf("Error loading %q: %v", pkgDir, err)
		return nil
	}

	logrus.Infof("Processing %q", pkgDir)
	fallbackVersion := MakeFallbackVersion(pkgDir)

	// GlobalDebug allows us to go from ssa function to ast funcdecl
	fset := token.NewFileSet()
	_, err = parser.ParseDir(fset, pkgDir, nil, parser.ParseComments)
	if err != nil {
		logrus.Fatalf("Error lparsing package %v: %v", pkgDir, err)
	}
	ssaProg := ssa.NewProgram(fset, ssa.GlobalDebug)

	packages.Visit(pkgs, nil, func(pkg *packages.Package) {
		if pkg.Types == nil || pkg.IllTyped {
			fmt.Printf("Found ill-typed package %v\n", pkg.PkgPath)
			for _, err := range pkg.Errors {
				fmt.Printf("Error: %v\n", err)
			}
			return
		}
		pp.fsetMap[pkg.PkgPath] = pkg.Fset
		if _, ok := pp.pkgMap[pkg.PkgPath]; !ok {
			_ = ssaProg.CreatePackage(pkg.Types, pkg.Syntax, pkg.TypesInfo, true)
			gPkg := NewPackage()
			gPkg.SourceURL = pkg.Name
			gPkg.Path = pkg.PkgPath
			gPkg.PkgID = pkg.ID
			gPkg.Version = coordination.PkgVersion(pkg, fallbackVersion)
			gPkg.Program = progInfo
			gPkg.PackagePtr = pkg

			progInfo.ContainsPkg[gPkg.PkgID] = gPkg
			pp.pkgMap[pkg.PkgPath] = gPkg
		}
	})

	ssaProg.Build()

	allFuncs := ssautil.AllFunctions(ssaProg)
	pp.calleesOf = chaLazyCallees(allFuncs)
	// parse all functions first to resolve all future function and block linkage when parsing instrs
	for fn := range allFuncs {
		if fn.Package() == nil || fn.Package().Pkg == nil || strings.HasPrefix(fn.Name(), "init") {
			continue
		}

		_ = fn.Package().Pkg.Name()
		pkgPath := fn.Package().Pkg.Path()

		if !strings.Contains(pkgPath, pp.ModuleName) {
			continue
		}

		if !pp.processedPkg[pkgPath] {
			fmt.Printf("processing new pkg %v\n", pkgPath)
			pp.processedPkg[pkgPath] = true
		}

		name, file, pos := extractFnInfo(pp.fsetMap[pkgPath], fn)
		gFn := NewFunction()
		gFn.Name = name
		gFn.File = file
		gFn.Pos = pos
		gFn.SsaFuncPointer = fn

		gFn.CfgGraph = NewCFGGraph()

		cfgRoot := NewBlock()
		cfgRoot.Id = "-1"
		cfgRoot.Function = gFn
		gFn.CfgGraph.Nodes["-1"] = cfgRoot
		gFn.CfgGraph.StartNode = cfgRoot
		gFn.CfgGraph.Function = gFn
		gFn.ContainsBlock[cfgRoot.Id] = cfgRoot

		cfgEndNode := NewBlock()
		endBlockIdx := strconv.Itoa(len(fn.Blocks))
		cfgEndNode.Id = endBlockIdx
		cfgEndNode.Function = gFn
		gFn.CfgGraph.Nodes[endBlockIdx] = cfgEndNode
		gFn.CfgGraph.EndNode = cfgEndNode
		gFn.ContainsBlock[cfgEndNode.Id] = cfgEndNode

		pp.graphFuncMap[fn] = gFn
		gPkg := pp.pkgMap[pkgPath]

		gPkg.ContainsFunction[gFn.Name] = gFn
		gFn.Package = gPkg

		for _, b := range fn.Blocks {
			pp.handleBlocks(b, gFn)
			bName := GetBlockName(pp.fsetMap[pkgPath], b)
			gBlock := pp.blockMap[bName]
			if b.Index == 0 {
				gFn.FirstBlock = gBlock
			}
		}
	}

	fmt.Printf("Parsing functions of %s\n", pkgDir)
	for fn := range allFuncs {
		// fmt.Printf("Parsing function %s\n", fn.String())
		if fn.Package() == nil || fn.Package().Pkg == nil || strings.HasPrefix(fn.Name(), "init") {
			continue
		}

		_ = fn.Package().Pkg.Name()
		pkgPath := fn.Package().Pkg.Path()
		if !strings.Contains(pkgPath, pp.ModuleName) {
			continue
		}

		if !pp.processedPkg[pkgPath] {
			fmt.Printf("processing new pkg %v\n", pkgPath)
			pp.processedPkg[pkgPath] = true
		}

		gFn, ok := pp.graphFuncMap[fn]
		if !ok {
			logrus.Fatalf("cannot find fn: %+v\n", fn)
			return nil
		}

		pp.instrMap = map[ssa.Instruction]*Instruction{}
		pp.varMap = make(map[string]*Variable)
		for i, param := range fn.Params {
			gVar := NewVariable()
			gVar.Name = param.Name()
			gVar.Type = param.Type().String()
			gVar.ContainingFunction = gFn

			gParam := NewParameter()
			gParam.Name = param.Name()
			gParam.Type = param.Type().String()
			gParam.ParamNum = i

			pp.varMap[param.Name()] = gVar
			gFn.Parameter[gParam.ParamNum] = gParam
			gParam.ParamIsVar = gVar
			gFn.ParamIsVar[gParam] = gVar

		}
		// need to iterate twice to make sure blocks have all been parsed
		// -> so that we can get access to all parsed blocks for our instruction analysis
		// fmt.Printf("Parsing All blocks of %s\n", fn.String())
		for _, b := range fn.Blocks {
			pp.handleBlocks(b, gFn)
			bName := GetBlockName(pp.fsetMap[pkgPath], b)
			gBlock := pp.blockMap[bName]

			gBlock.Id = b.String()

			// fmt.Printf("Parsing block %s\n", b.String())

			gFn.ContainsBlock[gBlock.Id] = gBlock
			gBlock.Function = gFn
			gFn.CfgGraph.Nodes[gBlock.Id] = gBlock

			var prevInstr *Instruction
			var lastInstr *Instruction
			for _, instr := range b.Instrs {

				if reflect.TypeOf(instr).Elem().Name() == "DebugRef" {
					continue
				}

				gInstr := pp.processInstrByType(instr, pp.fsetMap[pkgPath])
				gInstr.Block = gBlock
				if gInstr.SSAType == "Phi" {
					gInstr.Block.PhiInstr = gInstr
				}

				if prevInstr != nil {
					prevInstr.InstrNext = gInstr
					gInstr.InstrPrev = prevInstr

				} else {
					if gBlock != nil {
						gBlock.FirstInstr = gInstr
					}
				}

				prevInstr = gInstr
				lastInstr = gInstr
			}
			if gBlock != nil && lastInstr != nil {
				gBlock.LastInstr = lastInstr
			}

			gBlock.PredecessorIds = []string{}
			for _, val := range b.Preds {
				gBlock.PredecessorIds = append(gBlock.PredecessorIds, val.String())
			}

			gBlock.SuccessorIds = []string{}
			for _, val := range b.Succs {
				gBlock.SuccessorIds = append(gBlock.SuccessorIds, val.String())
			}

			if len(b.Preds) == 0 {
				gBlock.Function.CfgGraph.StartNode.SuccessorIds =
					append(gBlock.Function.CfgGraph.StartNode.SuccessorIds, b.String())
				gBlock.PredecessorIds = append(gBlock.PredecessorIds, "-1")
			}

			endBlockIdx := strconv.Itoa(len(fn.Blocks))
			if len(b.Succs) == 0 {
				gBlock.Function.CfgGraph.Nodes[endBlockIdx].PredecessorIds =
					append(gBlock.Function.CfgGraph.Nodes[endBlockIdx].PredecessorIds, b.String())
				gBlock.SuccessorIds = append(gBlock.SuccessorIds, endBlockIdx)
			}

		}

		gFn.LocalVariables = pp.varMap

		// // disable pointer analysis because it's too slow
		// ssaMains, err := getMainPackages(ssaProg.AllPackages())
		// if err != nil {
		// 	logrus.Fatal(err)
		// }
		// ptaConfig2 := &pointer.Config{
		// 	Mains:          ssaMains,
		// 	BuildCallGraph: true,
		// }
		// for _, varVal := range gFn.LocalVariables {
		// 	ssaVal := varVal.SSAValue
		// 	if ssaVal != nil && CanPoint((*ssaVal).Type()) {
		// 		ptaConfig2.AddQuery(*ssaVal)
		// 	}
		// }
		// ptares, err := pointer.Analyze(ptaConfig2)
		// if err != nil {
		// 	logrus.Fatal(err)
		// }

		// for _, varVal1 := range gFn.LocalVariables {
		// 	ssaVal1 := varVal1.SSAValue
		// 	if ssaVal1 != nil && CanPoint((*ssaVal1).Type()) {
		// 		//for _, l := range ptares.Queries[*ssaVal1].PointsTo().Labels() {
		// 		//	fmt.Printf("%s Possibly Points to %s\n", varVal1.Name, l.String())
		// 		//}
		// 		for _, varVal2 := range gFn.LocalVariables {
		// 			ssaVal2 := varVal2.SSAValue
		// 			if ssaVal2 != nil && ssaVal2 != ssaVal1 {
		// 				if ptares.Queries[*ssaVal1].MayAlias(ptares.Queries[*ssaVal2]) {
		// 					//fmt.Printf("%s Possibly Aliases with %s\n", varVal1.Name, varVal2.Name)
		// 					varVal1.MayAlias[varVal2.FullName()] = varVal2
		// 					varVal2.MayAlias[varVal1.FullName()] = varVal1
		// 				}
		// 			}
		// 		}

		// 	} else {
		// 		//fmt.Printf("SSAValue of %s is nil\n", varVal1.Name)
		// 	}
		// }

		fullNameVarMap := make(map[string]*Variable)
		for _, v := range pp.varMap {
			v.ContainingFunction = gFn
			fullNameVarMap[v.FullName()] = v
			v.ValueExpression = NewExpression(v.FullName(), VarExp)
			v.ValueExpression.PointsToVar = v
		}

		gFn.LocalVariables = fullNameVarMap

		fullNameBlockMap := make(map[string]*Block)
		for _, b := range gFn.ContainsBlock {
			fullNameBlockMap[b.FullName()] = b
		}

		gFn.ContainsBlock = fullNameBlockMap

		for _, b := range gFn.ContainsBlock {
			b.Successors = []*Block{}
			b.Predecessors = []*Block{}
			for _, succ := range b.SuccessorIds {
				succFullName := b.Function.FullName() + ":" + succ
				b.Successors = append(b.Successors, gFn.ContainsBlock[succFullName])
			}
			for _, pred := range b.PredecessorIds {
				predFullName := b.Function.FullName() + ":" + pred
				b.Predecessors = append(b.Predecessors, gFn.ContainsBlock[predFullName])
			}

			if b.LastInstr != nil && b.LastInstr.SSAType == "If" {

				trueBlockName := gFn.FullName() + ":" + b.LastInstr.TrueJump
				trueBlock := gFn.ContainsBlock[trueBlockName]
				b.TrueSuccessor = trueBlock

				falseBlockName := gFn.FullName() + ":" + b.LastInstr.FalseJump
				falseBLock := gFn.ContainsBlock[falseBlockName]
				b.FalseSuccessor = falseBLock

			}
			if b.LastInstr != nil && b.LastInstr.SSAType == "Jump" {
				uncondBlockName := gFn.FullName() + ":" + b.LastInstr.UnCondJump
				uncondBlock := gFn.ContainsBlock[uncondBlockName]

				//fmt.Printf("block %s : %s\n", b.FullName(), uncondBlockName)

				b.UnconditionalSuccessor = uncondBlock
			}
		}

		gFn.CfgGraph.BuildCDGFromCFG()

		//topologicalSort(gFn)

		for _, LocalVariable := range gFn.LocalVariables {
			if LocalVariable.DefiningInstruction == nil {
				// TODO : If parameters are state variables
				continue
			}
			LocalVariablePos := LocalVariable.DefiningInstruction.Pos
			if LocalVariablePos != 0 {
				gFn.PosMap[LocalVariablePos] = append(gFn.PosMap[LocalVariablePos], LocalVariable)
			}
		}

	}

	// use RTA instead? would require the "global graph" to be almost flow-sensitive, but would remove unreachable
	// interface calls
	logrus.Trace("Computing callgraph")
	cg := cha.CallGraph(ssaProg)
	logrus.Trace("Created callgraph")

	callgraph.GraphVisitEdges(cg, func(edge *callgraph.Edge) error {
		callerNode := edge.Caller.Func
		calleeNode := edge.Callee.Func
		// pkgName := edge.Site.Parent().Pkg.Pkg.Path()
		// if !pkgIsNew[pkgName] {
		// 	return nil
		// }

		if caller, found := pp.graphFuncMap[callerNode]; found {
			if callee, found := pp.graphFuncMap[calleeNode]; found {
				caller.Calls[callee.Name] = callee
				callee.Caller[caller.Name] = caller
			}
		}

		return nil
	})
	logrus.Trace("Callgraph nodes created")

	//-------

	//-------
	return progInfo
}

func (pp ProgParser) processVar(v ssa.Value, pkg string) *Variable {
	if v == nil {
		return nil
	}
	isVar, isGlobal, name, t := ParseSSAVar(v)
	if !isVar {
		return nil
	}

	if isGlobal {
		gVar, ok := pp.globalVarMap[pkg][name]
		if !ok {
			gVar = NewVariable()
			gVar.Name = name
			gVar.Type = t
			gVar.IsGlobal = isGlobal
			gVar.SSAValue = &v

			if pp.globalVarMap[pkg] == nil {
				pp.globalVarMap[pkg] = make(map[string]*Variable)
			}
			pp.globalVarMap[pkg][name] = gVar
		}
		return gVar
	}

	gVar, ok := pp.varMap[name]

	if !ok {
		gVar = NewVariable()
		gVar.Name = name
		gVar.Type = t
		gVar.IsGlobal = isGlobal
		gVar.SSAValue = &v

		pp.varMap[name] = gVar
	} else {
		if gVar.SSAValue == nil {
			gVar.SSAValue = &v
		}
	}

	if possibleChan(gVar.Type) {
		logrus.Tracef("found possible chan variable: %v", *gVar)
		// pp.chanMap.Set(&v, gVar)
		pp.chanMap[&v] = gVar
	}

	//for k, v := range pp.varMap {
	//	fmt.Printf("After VarMap %s\n", k)
	//	if v.SSAValue == nil {
	//		fmt.Printf("SSAValue of %s is nil\n", k)
	//	}
	//}

	return gVar
}

func (pp ProgParser) processParam(v ssa.Value, i int) *Parameter {
	if v == nil {
		return nil
	}
	isVar, _, name, t := ParseSSAVar(v)
	if !isVar {
		return nil
	}
	gParam := NewParameter()
	gParam.Name = name
	gParam.Type = t
	gParam.ParamNum = i

	return gParam
}

func (pp ProgParser) processRegister(name string, t string) *Variable {
	gVar, ok := pp.varMap[name]
	if !ok {
		gVar = NewVariable()
		gVar.Name = name
		gVar.Type = t

		pp.varMap[name] = gVar
	}
	return gVar
}

func (pp ProgParser) processInstrByType(instr ssa.Instruction,
	fset *token.FileSet) *Instruction {
	gInstr, ok := pp.instrMap[instr]
	if !ok {
		text, file, pos := ExtractInstrInfo(fset, instr)
		gInstr = NewInstruction()
		gInstr.SsaInstrPtr = &instr
		gInstr.File = file
		gInstr.Pos = pos
		gInstr.Text = text
		gInstr.SSAType = reflect.TypeOf(instr).Elem().Name()

		bName := GetBlockName(fset, instr.Block())
		gBlock := pp.blockMap[bName]
		gBlock.ContainsInstr[gInstr.Text] = gInstr
		gInstr.Block = gBlock

		currentFunction := gInstr.Block.Function

		switch instr := instr.(type) {
		case ssa.CallInstruction: // *Go, *Defer, *Call
			if instr.Value() != nil {
				varR := pp.processRegister(instr.Value().Name(), instr.Value().Type().String())
				currentFunction.LocalVariables[varR.Name] = varR

				gInstr.FunctionCallResult = varR
				gInstr.AssignTo["varR::"+varR.Name] = varR
				varR.DefiningInstruction = gInstr
			}
			// link callee function
			if g := instr.Common().StaticCallee(); g != nil {
				if gFn, ok := pp.graphFuncMap[g]; ok {
					gInstr.InstrCalls[gFn.Name] = gFn
				} else {
					// TODO: figure out why some functions aren't resolved!
					// fmt.Printf("error retrieving function %+v\n", g)
				}
			} else {
				for _, g := range pp.calleesOf(instr) {
					if gFn, ok := pp.graphFuncMap[g]; ok {
						gInstr.InstrCalls[gFn.Name] = gFn
					} else {
						// fmt.Printf("error retrieving function %+v\n", g)
					}
				}
			}
			if instr.Common().Value != nil {
				varVal := pp.processVar(instr.Common().Value, instr.Parent().Pkg.Pkg.Name())

				if varVal != nil {
					varVal.ContainingFunction = currentFunction
					currentFunction.LocalVariables[varVal.Name] = varVal
					gInstr.AssignFrom["varVal::"+varVal.Name] = varVal
				}
			}
			// link arguments

			for i, arg := range instr.Common().Args {
				newArg := NewArgument()

				varArg := pp.processVar(arg, instr.Parent().Pkg.Pkg.Name())

				if varArg != nil {
					varArg.ContainingFunction = currentFunction
					currentFunction.LocalVariables[varArg.Name] = varArg
					gInstr.AssignFrom["varArg::"+varArg.Name] = varArg

					newArg.CallerVariable = varArg
					newArg.Position = i
					gInstr.CallerArguments[i] = newArg
				} else {
					newArg.ConstantValue = arg.String()
					newArg.Position = i
					gInstr.CallerArguments[i] = newArg
				}

				gInstr.CalledFnArguments = append(gInstr.CalledFnArguments, newArg)
			}
		case *ssa.UnOp:
			varR := pp.processRegister(instr.Name(), instr.Type().String())
			currentFunction.LocalVariables[varR.Name] = varR
			gInstr.AssignTo["varR::"+varR.Name] = varR
			varR.DefiningInstruction = gInstr

			varX := pp.processVar(instr.X, instr.Parent().Pkg.Pkg.Name())

			if varX != nil {
				varX.ContainingFunction = currentFunction
				currentFunction.LocalVariables[varX.Name] = varX
				gInstr.AssignFrom["varX::"+varX.Name] = varX
			}
			if instr.Op == token.ARROW {
				if varX != nil {
					gInstr.ChannelRecv = varX
				}
			}

			gInstr.Operator = instr.Op.String()

		case *ssa.BinOp:
			//fmt.Printf("BinOp instr string: %v, op: %v, X: %v, Y: %v\n",
			//	instr.String(), instr.Op, instr.X, instr.Y)
			//fmt.Printf("X name: %v, string: %v, type: %v, referrers: %v\n",
			//	instr.X.Name(), instr.X.String(), instr.X.Type(),
			//	InstrsToString(instr.X.Referrers()))
			//fmt.Printf("Y name: %v, string: %v, type: %v, referrers: %v\n",
			//	instr.Y.Name(), instr.Y.String(), instr.Y.Type(),
			//	InstrsToString(instr.Y.Referrers()))
			//fmt.Printf("BinOp register string: %v, type: %v\n",
			//	instr.Name(), instr.Type())

			varR := pp.processRegister(instr.Name(), instr.Type().String())

			currentFunction.LocalVariables[varR.Name] = varR
			gInstr.AssignTo["varR::"+varR.Name] = varR
			varR.DefiningInstruction = gInstr

			varX := pp.processVar(instr.X, instr.Parent().Pkg.Pkg.Name())

			if varX != nil {
				varX.ContainingFunction = currentFunction
				currentFunction.LocalVariables[varX.Name] = varX
				gInstr.AssignFrom["varX::"+varX.Name] = varX
			} else {
				gInstr.OperandX = instr.X.String()
			}
			varY := pp.processVar(instr.Y, instr.Parent().Pkg.Pkg.Name())

			if varY != nil {
				varY.ContainingFunction = currentFunction
				currentFunction.LocalVariables[varY.Name] = varY
				gInstr.AssignFrom["varY::"+varY.Name] = varY
			} else {
				gInstr.OperandY = instr.Y.String()
			}

			gInstr.Operator = instr.Op.String()

		case *ssa.Store:
			// fmt.Printf("Store instr string: %v, Addr: %v, Val: %v\n",
			// 	instr.String(), instr.Addr, instr.Val)
			// fmt.Printf("Addr name: %v, string: %v, type: %v, referrers: %v\n",
			// 	instr.Addr.Name(), instr.Addr.String(), instr.Addr.Type(),
			// 	InstrsToString(instr.Addr.Referrers()))
			// fmt.Printf("Val name: %v, string: %v, type: %v, referrers: %v\n",
			// 	instr.Val.Name(), instr.Val.String(), instr.Val.Type(),
			// 	InstrsToString(instr.Val.Referrers()))
			varAddr := pp.processVar(instr.Addr, instr.Parent().Pkg.Pkg.Name())

			if varAddr != nil {
				varAddr.ContainingFunction = currentFunction
				currentFunction.LocalVariables[varAddr.Name] = varAddr
				gInstr.AssignTo["varAddr::"+varAddr.Name] = varAddr
			}

			varVal := pp.processVar(instr.Val, instr.Parent().Pkg.Pkg.Name())

			if varVal != nil {
				varVal.ContainingFunction = currentFunction
				currentFunction.LocalVariables[varVal.Name] = varVal
				gInstr.AssignFrom["varVal::"+varVal.Name] = varVal
			} else {
				gInstr.Misc = instr.Val.Name()
			}
		case *ssa.Phi:
			//fmt.Printf("Phi instr string: %v\n",
			//	instr.String())
			//for _, edge := range instr.Edges {
			//	fmt.Printf("Phi edge name: %v, string: %v, type: %v, referrers: %v\n",
			//		edge.Name(), edge.String(), edge.Type(),
			//		InstrsToString(edge.Referrers()))
			//}

			varR := pp.processRegister(instr.Name(), instr.Type().String())

			currentFunction.LocalVariables[varR.Name] = varR
			gInstr.AssignTo["varR::"+varR.Name] = varR
			varR.DefiningInstruction = gInstr

			for i, edge := range instr.Edges {
				varEdge := pp.processVar(edge, instr.Parent().Pkg.Pkg.Name())

				if varEdge != nil {
					varEdge.ContainingFunction = currentFunction
					currentFunction.LocalVariables[varEdge.FullName()] = varEdge
					gInstr.AssignFrom["varEdge::"+varEdge.FullName()] = varEdge
					gInstr.PhiMap[instr.Block().Preds[i].String()] = varEdge.FullName()
				} else {
					gInstr.PhiMap[instr.Block().Preds[i].String()] = instr.Edges[i].Name()
				}
			}

		case *ssa.ChangeType: // cannot fail dynamically
			varR := pp.processRegister(instr.Name(), instr.Type().String())

			currentFunction.LocalVariables[varR.Name] = varR
			gInstr.AssignTo["varR::"+varR.Name] = varR

			varR.DefiningInstruction = gInstr

			varX := pp.processVar(instr.X, instr.Parent().Pkg.Pkg.Name())

			if varX != nil {
				varX.ContainingFunction = currentFunction
				currentFunction.LocalVariables[varX.Name] = varX
				gInstr.AssignFrom["varX::"+varX.Name] = varX
			}

		case *ssa.Convert: // cannot fail dynamically
			varR := pp.processRegister(instr.Name(), instr.Type().String())

			currentFunction.LocalVariables[varR.Name] = varR
			gInstr.AssignTo["varR::"+varR.Name] = varR

			varR.DefiningInstruction = gInstr

			varX := pp.processVar(instr.X, instr.Parent().Pkg.Pkg.Name())

			if varX != nil {
				varX.ContainingFunction = currentFunction
				currentFunction.LocalVariables[varX.Name] = varX
				gInstr.AssignFrom["varX::"+varX.Name] = varX
			}

		case *ssa.MultiConvert: // can fail dynamically (see SliceToArrayPointer)
			varR := pp.processRegister(instr.Name(), instr.Type().String())

			currentFunction.LocalVariables[varR.Name] = varR
			gInstr.AssignTo["varR::"+varR.Name] = varR
			varR.DefiningInstruction = gInstr

			varX := pp.processVar(instr.X, instr.Parent().Pkg.Pkg.Name())

			if varX != nil {
				varX.ContainingFunction = currentFunction
				currentFunction.LocalVariables[varX.Name] = varX
				gInstr.AssignFrom["varX::"+varX.Name] = varX
			}

		case *ssa.ChangeInterface: // cannot fail
			varR := pp.processRegister(instr.Name(), instr.Type().String())

			currentFunction.LocalVariables[varR.Name] = varR
			gInstr.AssignTo["varR::"+varR.Name] = varR

			varR.DefiningInstruction = gInstr

			varX := pp.processVar(instr.X, instr.Parent().Pkg.Pkg.Name())

			if varX != nil {
				varX.ContainingFunction = currentFunction
				currentFunction.LocalVariables[varX.Name] = varX
				gInstr.AssignFrom["varX::"+varX.Name] = varX
			}

		case *ssa.SliceToArrayPointer: // can fail
			varR := pp.processRegister(instr.Name(), instr.Type().String())

			currentFunction.LocalVariables[varR.Name] = varR
			gInstr.AssignTo["varR::"+varR.Name] = varR

			varR.DefiningInstruction = gInstr

			varX := pp.processVar(instr.X, instr.Parent().Pkg.Pkg.Name())

			if varX != nil {
				varX.ContainingFunction = currentFunction
				currentFunction.LocalVariables[varX.Name] = varX
				gInstr.AssignFrom["varX::"+varX.Name] = varX
			}

		case *ssa.MakeInterface:
			varR := pp.processRegister(instr.Name(), instr.Type().String())

			currentFunction.LocalVariables[varR.Name] = varR
			gInstr.AssignTo["varR::"+varR.Name] = varR

			varR.DefiningInstruction = gInstr

			varX := pp.processVar(instr.X, instr.Parent().Pkg.Pkg.Name())

			if varX != nil {
				varX.ContainingFunction = currentFunction
				currentFunction.LocalVariables[varX.Name] = varX
				gInstr.AssignFrom["varX::"+varX.Name] = varX
			}

		case *ssa.MakeClosure:
			varR := pp.processRegister(instr.Name(), instr.Type().String())

			currentFunction.LocalVariables[varR.Name] = varR
			gInstr.AssignTo["varR::"+varR.Name] = varR

			varR.DefiningInstruction = gInstr

			for _, freeVar := range instr.Bindings {
				varX := pp.processVar(freeVar, instr.Parent().Pkg.Pkg.Name())
				if varX != nil {
					varX.ContainingFunction = currentFunction
					currentFunction.LocalVariables[varX.Name] = varX
					gInstr.AssignFrom["varX::"+varX.Name] = varX
				}
			}

		case *ssa.MakeMap:
		case *ssa.MakeChan:
		case *ssa.MakeSlice:
			varR := pp.processRegister(instr.Name(), instr.Type().String())

			currentFunction.LocalVariables[varR.Name] = varR
			gInstr.AssignTo["varR::"+varR.Name] = varR
			varR.DefiningInstruction = gInstr

			varLen := pp.processVar(instr.Len, instr.Parent().Pkg.Pkg.Name())

			if varLen != nil {
				varLen.ContainingFunction = currentFunction
				currentFunction.LocalVariables[varLen.Name] = varLen
				gInstr.AssignFrom["varLen::"+varLen.Name] = varLen
			}

			varCap := pp.processVar(instr.Cap, instr.Parent().Pkg.Pkg.Name())

			if varCap != nil {
				varCap.ContainingFunction = currentFunction
				currentFunction.LocalVariables[varCap.Name] = varCap
				gInstr.AssignFrom["varCap::"+varCap.Name] = varCap
			}

		case *ssa.Slice:
			varR := pp.processRegister(instr.Name(), instr.Type().String())

			currentFunction.LocalVariables[varR.Name] = varR
			gInstr.AssignTo["varR::"+varR.Name] = varR
			varR.DefiningInstruction = gInstr

			varX := pp.processVar(instr.X, instr.Parent().Pkg.Pkg.Name())

			if varX != nil {
				varX.ContainingFunction = currentFunction
				currentFunction.LocalVariables[varX.Name] = varX
				gInstr.AssignFrom["varX::"+varX.Name] = varX
			}

			varLow := pp.processVar(instr.Low, instr.Parent().Pkg.Pkg.Name())

			if varLow != nil {
				varLow.ContainingFunction = currentFunction
				currentFunction.LocalVariables[varLow.Name] = varLow
				gInstr.AssignFrom["varLow::"+varLow.Name] = varLow
			} else if instr.Low != nil {
				gInstr.LowVal = instr.Low.String()
			}

			varHigh := pp.processVar(instr.High, instr.Parent().Pkg.Pkg.Name())

			if varHigh != nil {
				varHigh.ContainingFunction = currentFunction
				currentFunction.LocalVariables[varHigh.Name] = varHigh
				gInstr.AssignFrom["varHigh::"+varHigh.Name] = varHigh
			} else if instr.High != nil {
				gInstr.HighVal = instr.High.String()
			}

			varMax := pp.processVar(instr.Max, instr.Parent().Pkg.Pkg.Name())

			if varMax != nil {
				varMax.ContainingFunction = currentFunction
				currentFunction.LocalVariables[varMax.Name] = varMax
				gInstr.AssignFrom["varMax::"+varMax.Name] = varMax
			}

		case *ssa.FieldAddr:
			// fmt.Printf("FieldAddr instr string: %v\n", instr.String())
			// fmt.Printf("FieldAddr X name: %v, string: %v, type: %v, referrers: %v\n",
			// 	instr.X.Name(), instr.X.String(), instr.X.Type(),
			// 	InstrsToString(instr.X.Referrers()))
			varR := pp.processRegister(instr.Name(), instr.Type().String())
			currentFunction.LocalVariables[varR.Name] = varR
			gInstr.AssignTo["varR::"+varR.Name] = varR
			varR.DefiningInstruction = gInstr

			varX := pp.processVar(instr.X, instr.Parent().Pkg.Pkg.Name())

			if varX != nil {
				varX.ContainingFunction = currentFunction
				currentFunction.LocalVariables[varX.Name] = varX
				gInstr.AssignFrom["varX::"+varX.Name] = varX
			}

			gInstr.Misc = fmt.Sprint(instr.Field + 1)
		case *ssa.Field:
			varR := pp.processRegister(instr.Name(), instr.Type().String())

			currentFunction.LocalVariables[varR.Name] = varR
			gInstr.AssignTo["varR::"+varR.Name] = varR
			varR.DefiningInstruction = gInstr

			varX := pp.processVar(instr.X, instr.Parent().Pkg.Pkg.Name())

			if varX != nil {
				varX.ContainingFunction = currentFunction
				currentFunction.LocalVariables[varX.Name] = varX
				gInstr.AssignFrom["varX::"+varX.Name] = varX
			}

			gInstr.Misc = fmt.Sprint(instr.Field + 1)
		case *ssa.IndexAddr:
			varR := pp.processRegister(instr.Name(), instr.Type().String())

			currentFunction.LocalVariables[varR.Name] = varR
			gInstr.AssignTo["varR::"+varR.Name] = varR
			varR.DefiningInstruction = gInstr

			varX := pp.processVar(instr.X, instr.Parent().Pkg.Pkg.Name())

			if varX != nil {
				varX.ContainingFunction = currentFunction
				currentFunction.LocalVariables[varX.Name] = varX
				gInstr.AssignFrom["varX::"+varX.Name] = varX
			}

			varIdx := pp.processVar(instr.Index, instr.Parent().Pkg.Pkg.Name())
			if varIdx != nil {
				varIdx.ContainingFunction = currentFunction
				currentFunction.LocalVariables[varIdx.Name] = varIdx
				gInstr.AssignFrom["varIdx::"+varIdx.Name] = varIdx
			}

			if isVar, _, _, t := ParseSSAVar(instr.Index); !isVar {
				if t == "int" {
					gInstr.Misc = instr.Index.String()
				}
			}
		case *ssa.Index:
			varR := pp.processRegister(instr.Name(), instr.Type().String())

			currentFunction.LocalVariables[varR.Name] = varR
			gInstr.AssignTo["varR::"+varR.Name] = varR
			varR.DefiningInstruction = gInstr

			varX := pp.processVar(instr.X, instr.Parent().Pkg.Pkg.Name())
			if varX != nil {
				varX.ContainingFunction = currentFunction
				currentFunction.LocalVariables[varX.Name] = varX
				gInstr.AssignFrom["varX::"+varX.Name] = varX
			}

			varIdx := pp.processVar(instr.Index, instr.Parent().Pkg.Pkg.Name())
			if varIdx != nil {
				varIdx.ContainingFunction = currentFunction
				currentFunction.LocalVariables[varIdx.Name] = varIdx
				gInstr.AssignFrom["varIdx::"+varIdx.Name] = varIdx
			}

			if isVar, _, _, t := ParseSSAVar(instr.Index); !isVar {
				if t == "int" {
					gInstr.Misc = instr.Index.String()
				}
			}

		case *ssa.Lookup:
			varR := pp.processRegister(instr.Name(), instr.Type().String())

			currentFunction.LocalVariables[varR.Name] = varR
			gInstr.AssignTo["varR::"+varR.Name] = varR
			varR.DefiningInstruction = gInstr

			varX := pp.processVar(instr.X, instr.Parent().Pkg.Pkg.Name())
			if varX != nil {
				varX.ContainingFunction = currentFunction
				currentFunction.LocalVariables[varX.Name] = varX
				gInstr.AssignFrom["varX::"+varX.Name] = varX
			}

			varIdx := pp.processVar(instr.Index, instr.Parent().Pkg.Pkg.Name())
			if varIdx != nil {
				varIdx.ContainingFunction = currentFunction
				currentFunction.LocalVariables[varIdx.Name] = varIdx
				gInstr.AssignFrom["varIdx::"+varIdx.Name] = varIdx
			}

		case *ssa.Select:
			varR := pp.processRegister(instr.Name(), instr.Type().String())

			currentFunction.LocalVariables[varR.Name] = varR
			gInstr.AssignTo["varR::"+varR.Name] = varR
			varR.DefiningInstruction = gInstr

			for _, st := range instr.States {
				varChan := pp.processVar(st.Chan, instr.Parent().Pkg.Pkg.Name())

				if st.Dir == types.RecvOnly {
					if varChan != nil {
						gInstr.ChannelRecv = varChan
						//TODO : CHECK
					}

				} else if st.Dir == types.SendOnly {
					if varChan != nil {
						gInstr.ChannelSend = varChan
						//TODO : CHECK
					}

					varSend := pp.processVar(st.Send, instr.Parent().Pkg.Pkg.Name())

					if varSend != nil {
						varSend.ContainingFunction = currentFunction
						currentFunction.LocalVariables[varSend.Name] = varSend
						gInstr.AssignFrom["varSend::"+varSend.Name] = varSend
					}
				}
			}

		case *ssa.Range:
			varR := pp.processRegister(instr.Name(), instr.Type().String())

			currentFunction.LocalVariables[varR.Name] = varR
			gInstr.AssignTo["varR::"+varR.Name] = varR
			varR.DefiningInstruction = gInstr

			varX := pp.processVar(instr.X, instr.Parent().Pkg.Pkg.Name())
			if varX != nil {
				varX.ContainingFunction = currentFunction
				currentFunction.LocalVariables[varX.Name] = varX
				gInstr.AssignFrom["varX::"+varX.Name] = varX
			}

		case *ssa.Next:
			varR := pp.processRegister(instr.Name(), instr.Type().String())

			currentFunction.LocalVariables[varR.Name] = varR
			gInstr.AssignTo["varR::"+varR.Name] = varR
			varR.DefiningInstruction = gInstr

			varIter := pp.processVar(instr.Iter, instr.Parent().Pkg.Pkg.Name())
			if varIter != nil {
				varIter.ContainingFunction = currentFunction
				currentFunction.LocalVariables[varIter.Name] = varIter
				gInstr.AssignFrom["varIter::"+varIter.Name] = varIter
			}

		case *ssa.TypeAssert:
			varR := pp.processRegister(instr.Name(), instr.Type().String())

			currentFunction.LocalVariables[varR.Name] = varR
			gInstr.AssignTo["varR::"+varR.Name] = varR
			varR.DefiningInstruction = gInstr

			varX := pp.processVar(instr.X, instr.Parent().Pkg.Pkg.Name())
			if varX != nil {
				varX.ContainingFunction = currentFunction
				currentFunction.LocalVariables[varX.Name] = varX
				gInstr.AssignFrom["varX::"+varX.Name] = varX
			}

		case *ssa.Extract:
			varR := pp.processRegister(instr.Name(), instr.Type().String())

			currentFunction.LocalVariables[varR.Name] = varR
			gInstr.AssignTo["varR::"+varR.Name] = varR
			varR.DefiningInstruction = gInstr

			varTuple := pp.processVar(instr.Tuple, instr.Parent().Pkg.Pkg.Name())
			if varTuple != nil {
				varTuple.ContainingFunction = currentFunction
				currentFunction.LocalVariables[varTuple.Name] = varTuple
				gInstr.AssignFrom["varTuple::"+varTuple.Name] = varTuple
			}

			gInstr.Misc = fmt.Sprint(instr.Index)
		case *ssa.Jump:
			gInstr.UnCondJump = strings.TrimSpace(strings.Split(gInstr.Text, "jump ")[1])

		case *ssa.If:
			varCond := pp.processVar(instr.Cond, instr.Parent().Pkg.Pkg.Name())
			if varCond != nil {
				varCond.ContainingFunction = currentFunction
				currentFunction.LocalVariables[varCond.Name] = varCond
				gInstr.AssignFrom["varCond::"+varCond.Name] = varCond
			} else {
				gInstr.Misc = strings.TrimSpace(strings.Split(instr.Cond.String(), ":")[0])
			}

			gInstr.TrueJump = strings.TrimSpace(strings.Split(strings.Split(gInstr.Text, "goto ")[1], " else")[0])
			gInstr.FalseJump = strings.TrimSpace(strings.Split(gInstr.Text, "else ")[1])

		case *ssa.Return:
			for i, res := range instr.Results {

				newRes := NewArgument()

				varRes := pp.processVar(res, instr.Parent().Pkg.Pkg.Name())
				if varRes != nil {

					newRes.CallerVariable = varRes
					newRes.Position = i

					varRes.ContainingFunction = currentFunction
					currentFunction.LocalVariables[varRes.Name] = varRes
					gInstr.AssignFrom["varRes::"+varRes.Name] = varRes
					gInstr.ReturnValues[i] = newRes
				}
				//paramRes := pp.processParam(res, i)
				//if paramRes != nil {
				//	gInstr.CallParameter[paramRes.Name] = paramRes
				//}
				//
				//if paramRes != nil && varRes != nil {
				//	paramRes.ParamIsVar[varRes.Name] = varRes
				//}
			}

		case *ssa.RunDefers:
		case *ssa.Panic:
			varX := pp.processVar(instr.X, instr.Parent().Pkg.Pkg.Name())

			if varX != nil {
				varX.ContainingFunction = currentFunction
				currentFunction.LocalVariables[varX.Name] = varX
				gInstr.AssignFrom["varX::"+varX.Name] = varX
			}

		case *ssa.Send:
			varChan := pp.processVar(instr.Chan, instr.Parent().Pkg.Pkg.Name())
			if varChan != nil {
				gInstr.ChannelSend = varChan
			}

			varX := pp.processVar(instr.X, instr.Parent().Pkg.Pkg.Name())
			if varX != nil {
				varX.ContainingFunction = currentFunction
				currentFunction.LocalVariables[varX.Name] = varX
				gInstr.AssignFrom["varX::"+varX.Name] = varX
			}

		case *ssa.MapUpdate:
			varM := pp.processVar(instr.Map, instr.Parent().Pkg.Pkg.Name())
			if varM != nil {
				varM.ContainingFunction = currentFunction
				currentFunction.LocalVariables[varM.Name] = varM
				gInstr.AssignTo["varM::"+varM.Name] = varM
				varM.DefiningInstruction = gInstr
			}

			varK := pp.processVar(instr.Key, instr.Parent().Pkg.Pkg.Name())
			if varK != nil {
				varK.ContainingFunction = currentFunction
				currentFunction.LocalVariables[varK.Name] = varK
				gInstr.AssignTo["varK::"+varK.Name] = varK
				varK.DefiningInstruction = gInstr
			}

			varVal := pp.processVar(instr.Value, instr.Parent().Pkg.Pkg.Name())
			if varVal != nil {
				varVal.ContainingFunction = currentFunction
				currentFunction.LocalVariables[varVal.Name] = varVal
				gInstr.AssignFrom["varVal::"+varVal.Name] = varVal
			}

		case *ssa.DebugRef:
			varX := pp.processVar(instr.X, instr.Parent().Pkg.Pkg.Name())
			if varX != nil {
				varX.ContainingFunction = currentFunction
				currentFunction.LocalVariables[varX.Name] = varX
				gInstr.AssignFrom["varX::"+varX.Name] = varX
			}

		case *ssa.Alloc:
			varR := pp.processRegister(instr.Name(), instr.Type().String())

			currentFunction.LocalVariables[varR.Name] = varR
			gInstr.AssignTo["varR::"+varR.Name] = varR
			varR.DefiningInstruction = gInstr

		default:
			fmt.Printf("Missing handling routing for instr value: "+
				"%v,  type: %v, file: %v,block: %v, operands: %v, parent: %v\n",
				instr.String(), reflect.TypeOf(instr).Elem().Name(),
				fset.PositionFor(instr.Pos(), false),
				instr.Block(), operandsToString(instr.Operands(nil)),
				instr.Parent())
		}

	}
	return gInstr
}

func (pp ProgParser) handleBlocks(b *ssa.BasicBlock, containingFunction *Function) {
	fn := b.Parent()
	pkgPath := fn.Package().Pkg.Path()
	bName := GetBlockName(pp.fsetMap[pkgPath], b)
	curBlock, ok := pp.blockMap[bName]
	if !ok {
		name, file, pos := extractBlockInfo(pp.fsetMap[pkgPath], b)
		//fmt.Printf("905 Creating Block %s\n", b.String())
		gBlock := NewBlock()
		gBlock.Name = name
		gBlock.File = file
		gBlock.Pos = pos
		gBlock.Id = b.String()

		gBlock.Function = containingFunction
		containingFunction.ContainsBlock[gBlock.Id] = gBlock
		containingFunction.CfgGraph.Nodes[gBlock.Id] = gBlock

		pp.blockMap[bName] = gBlock
		curBlock = gBlock
	}

	for _, preB := range b.Preds {
		bName = GetBlockName(pp.fsetMap[pkgPath], preB)
		preBlock, ok := pp.blockMap[bName]
		if !ok {
			name, file, pos := extractBlockInfo(pp.fsetMap[pkgPath], preB)
			//fmt.Printf("921 Creating Block %s\n", b.String())
			gBlock := NewBlock()
			gBlock.Name = name
			gBlock.File = file
			gBlock.Pos = pos
			gBlock.Id = b.String()

			gBlock.Function = containingFunction
			containingFunction.ContainsBlock[gBlock.Id] = gBlock
			containingFunction.CfgGraph.Nodes[gBlock.Id] = gBlock

			pp.blockMap[bName] = gBlock
			preBlock = gBlock
		}
		curBlock.BlockPrev = preBlock

	}

	for _, succB := range b.Succs {
		bName = GetBlockName(pp.fsetMap[pkgPath], succB)
		succBlock, ok := pp.blockMap[bName]
		if !ok {
			name, file, pos := extractBlockInfo(pp.fsetMap[pkgPath], succB)
			//fmt.Printf("940 Creating Block %s\n", b.String())
			gBlock := NewBlock()
			gBlock.Name = name
			gBlock.File = file
			gBlock.Pos = pos
			gBlock.Id = succB.String()

			gBlock.Function = containingFunction
			containingFunction.ContainsBlock[gBlock.Id] = gBlock
			containingFunction.CfgGraph.Nodes[gBlock.Id] = gBlock

			pp.blockMap[bName] = gBlock
			succBlock = gBlock
		}
		curBlock.BlockNext = succBlock

	}
}

func operandsToString(vals []*ssa.Value) string {
	res := "["
	for _, val := range vals {
		if val != nil && *val != nil {
			res += fmt.Sprintf("%v, ", (*val).String())
		}
	}
	if len(res) > 1 {
		res = res[:len(res)-2]
	}
	res += "]"
	return res
}
