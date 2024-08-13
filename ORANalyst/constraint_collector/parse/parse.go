package parse

import (
	"fmt"
	"go/parser"
	"go/token"
	"go/types"
	"reflect"
	"strings"
	"sync"

	"constraint_collector/coordination"
	"constraint_collector/schema"

	"github.com/sirupsen/logrus"
	"golang.org/x/tools/go/callgraph"
	"golang.org/x/tools/go/callgraph/cha"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/pointer"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

type ProgParser struct {
	ProgressBar bool
	ModuleName  string

	// values that will be built up during package walking, and used in callgraph processing
	pkgMap       map[string]*schema.Package
	graphFuncMap map[*ssa.Function]*schema.Function
	fsetMap      map[string]*token.FileSet
	blockMap     map[string]*schema.Block
	fset         token.FileSet
	vertices     chan schema.Vertex
	instrMap     map[ssa.Instruction]*schema.Instruction
	varMap       map[string]*schema.Variable
	edges        []schema.Edge
	calleesOf    func(site ssa.CallInstruction) []*ssa.Function
	processedPkg map[string]bool
	globalVarMap map[string]map[string]*schema.Variable
	chanMap      map[*ssa.Value]*schema.Variable
	// chanMap      *om.OrderedMap[*ssa.Value, *schema.Variable]
}

func NewProgParser(progressBar bool, moduleName string) ProgParser {
	return ProgParser{
		ProgressBar:  true,
		ModuleName:   moduleName,
		pkgMap:       map[string]*schema.Package{},
		graphFuncMap: map[*ssa.Function]*schema.Function{},
		fsetMap:      map[string]*token.FileSet{},
		blockMap:     map[string]*schema.Block{},
		vertices:     make(chan schema.Vertex),
		instrMap:     map[ssa.Instruction]*schema.Instruction{},
		varMap:       map[string]*schema.Variable{},
		processedPkg: map[string]bool{},
		globalVarMap: map[string]map[string]*schema.Variable{},
		chanMap:      map[*ssa.Value]*schema.Variable{},
		// chanMap:      om.New[*ssa.Value, *schema.Variable](),
	}
}

func (pp ProgParser) ProcessPackage(pkgDir string) {
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
		return
	}

	logrus.Infof("Processing %q", pkgDir)
	pp.vertices = make(chan schema.Vertex, 32768)
	verticesCount := 0
	var vertexWG *sync.WaitGroup

	//if !pp.ProgressBar {
	//	vertexWG = backend.AddVStream(pp.vertices, func(done []schema.Vertex) {
	//		verticesCount += len(done)
	//	})
	//} else {
	//	progress := progressbar.Default(-1)
	//	vertexWG = backend.AddVStream(pp.vertices, func(done []schema.Vertex) {
	//		verticesCount += len(done)
	//		progress.Add(len(done))
	//	})
	//}
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
			return
		}
		pp.fsetMap[pkg.PkgPath] = pkg.Fset
		if _, ok := pp.pkgMap[pkg.PkgPath]; !ok {
			_ = ssaProg.CreatePackage(pkg.Types, pkg.Syntax, pkg.TypesInfo, true)
			gPkg := &schema.Package{
				SourceURL: pkg.Name,
				Path:      pkg.PkgPath,
				PkgID:     pkg.ID,
				Version:   coordination.PkgVersion(pkg, fallbackVersion),
			}
			pp.vertices <- gPkg
			pp.pkgMap[pkg.PkgPath] = gPkg
		}
	})

	ssaProg.Build()

	allFuncs := ssautil.AllFunctions(ssaProg)
	pp.calleesOf = chaLazyCallees(allFuncs)
	// parse all functions first to resolve all future function and block linkage when parsing instrs
	for fn := range allFuncs {
		if fn.Package() == nil || fn.Package().Pkg == nil {
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
		gFn := &schema.Function{
			Name: name,
			File: file,
			Pos:  pos,
		}
		pp.graphFuncMap[fn] = gFn
		pp.vertices <- gFn
		gPkg := pp.pkgMap[pkgPath]
		pp.edges = append(pp.edges, schema.Edge{
			Source: gPkg,
			Label:  "ContainsFunction",
			Target: gFn,
		})

		for _, b := range fn.Blocks {
			pp.edges = pp.handleBlocks(b)
			bName := GetBlockName(pp.fsetMap[pkgPath], b)
			gBlock := pp.blockMap[bName]
			pp.edges = append(pp.edges, schema.Edge{
				Source: gFn,
				Label:  "ContainsBlock",
				Target: gBlock,
			})
			if b.Index == 0 {
				pp.edges = append(pp.edges, schema.Edge{
					Source: gFn,
					Label:  "FirstBlock",
					Target: gBlock,
				})
			}
		}
	}

	for fn := range allFuncs {
		if fn.Package() == nil || fn.Package().Pkg == nil {
			continue
		}

		// if !strings.Contains(fn.Name(), "processMid") && !strings.Contains(fn.Name(), "factorial") &&
		// 	!strings.Contains(fn.Name(), "processMessage") && !strings.Contains(fn.Name(), "processNegative") &&
		// 	!strings.Contains(fn.Name(), "processLarge") && !strings.Contains(fn.Name(), "grpcSayHello") {
		// 	continue
		// }
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
			return
		}
		pp.instrMap = map[ssa.Instruction]*schema.Instruction{}
		pp.varMap = map[string]*schema.Variable{}
		for i, param := range fn.Params {
			gVar := &schema.Variable{
				Name: param.Name(),
				Type: param.Type().String(),
			}
			gParam := &schema.Parameter{
				Name:     param.Name(),
				Type:     param.Type().String(),
				ParamNum: i,
			}
			pp.vertices <- gParam
			pp.vertices <- gVar
			pp.varMap[param.Name()] = gVar
			pp.edges = append(pp.edges, schema.Edge{
				Source: gFn,
				Label:  "Parameter",
				Target: gParam,
			})
			pp.edges = append(pp.edges, schema.Edge{
				Source: gParam,
				Label:  "ParamIsVar",
				Target: gVar,
			})
		}
		// need to iterate twice to make sure blocks have all been parsed
		// -> so that we can get access to all parsed blocks for our instruction analysis
		for _, b := range fn.Blocks {
			bName := GetBlockName(pp.fsetMap[pkgPath], b)
			gBlock := pp.blockMap[bName]
			var prevInstr *schema.Instruction
			var lastInstr *schema.Instruction
			for _, instr := range b.Instrs {
				gInstr, newEdges := pp.processInstrByType(instr, pp.fsetMap[pkgPath])
				pp.edges = newEdges

				if prevInstr != nil {
					pp.edges = append(pp.edges, schema.Edge{
						Source: prevInstr,
						Label:  "InstrNext",
						Target: gInstr,
					}, schema.Edge{
						Source: gInstr,
						Label:  "InstrPrev",
						Target: prevInstr,
					})
				} else {
					if gBlock != nil {
						pp.edges = append(pp.edges, schema.Edge{
							Source: gBlock,
							Label:  "FirstInstr",
							Target: gInstr,
						})
					}
				}

				prevInstr = gInstr
				lastInstr = gInstr
			}
			if gBlock != nil && lastInstr != nil {
				pp.edges = append(pp.edges, schema.Edge{
					Source: gBlock,
					Label:  "LastInstr",
					Target: lastInstr,
				})
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
				pp.edges = append(
					pp.edges,
					schema.Edge{Source: caller, Label: "Calls", Target: callee},
					schema.Edge{Source: callee, Label: "Caller", Target: caller},
				)
			}
		}

		return nil
	})
	logrus.Trace("Callgraph nodes created")

	// TODO: points to analysis
	logrus.Info("performing points to analysis")
	// logrus.Infof("chan map length: %v", pp.chanMap.Len())
	logrus.Infof("chan map length: %v", len(pp.chanMap))
	ssaMains, err := getMainPackages(ssaProg.AllPackages())
	if err != nil {
		logrus.Fatal(err)
	}
	ptaConfig := &pointer.Config{
		Mains:          ssaMains,
		BuildCallGraph: true,
	}
	for val := range pp.chanMap {
		ptaConfig.AddQuery(*val)
	}
	ptares, err := pointer.Analyze(ptaConfig)
	if err != nil {
		logrus.Fatal(err)
	}

	for val1, val1Vertex := range pp.chanMap {
		logrus.Tracef("for val: %v\n", val1Vertex.Name)
		for _, label := range ptares.Queries[*val1].PointsTo().Labels() {
			logrus.Tracef("got label: %v\n", label.String())
		}
		for _, label := range ptares.IndirectQueries[*val1].PointsTo().Labels() {
			logrus.Tracef("got indirect label: %v\n", label.String())
		}
	}

	for val1, val1Vertex := range pp.chanMap {
		for val2, val2Vertex := range pp.chanMap {
			if ptares.Queries[*val1].MayAlias(ptares.Queries[*val2]) {
				pp.edges = append(pp.edges, schema.Edge{
					Source: val1Vertex,
					Label:  "PossibleSameChan",
					Target: val2Vertex,
				})
			}
		}
	}

	// processedMap := make(map[*ssa.Value]bool)
	// var ptaProgress *progressbar.ProgressBar
	// if pp.ProgressBar {
	// 	ptaProgress = progressbar.Default(int64(len(pp.chanMap)))
	// }
	// for {
	// 	if len(processedMap) == len(pp.chanMap) {
	// 		break
	// 	}
	// 	ptaConfig := &pointer.Config{
	// 		Mains:          ssaMains,
	// 		BuildCallGraph: true,
	// 	}
	// 	// for pair := pp.chanMap.Oldest(); pair != nil; pair = pair.Next() {
	// 	// 	ptaConfig.AddQuery(*pair.Key)
	// 	// }
	// 	chanType := ""
	// 	processingMap := make(map[*ssa.Value]*schema.Variable)
	// 	for val, valVertex := range pp.chanMap {
	// 		if chanType == "" && !processedMap[val] {
	// 			chanType = valVertex.Type
	// 			logrus.Tracef("points to analysis processing type %v", chanType)
	// 		}
	// 		if matchType(chanType, valVertex.Type) {
	// 			ptaConfig.AddQuery(*val)
	// 			processingMap[val] = valVertex
	// 			processedMap[val] = true
	// 		}
	// 	}
	// 	ptares, err := pointer.Analyze(ptaConfig)
	// 	if err != nil {
	// 		logrus.Fatal(err)
	// 	}

	// 	// for pair1 := pp.chanMap.Oldest(); pair1 != nil; pair1 = pair1.Next() {
	// 	// 	for pair2 := pp.chanMap.Oldest(); pair2 != nil; pair2 = pair1.Next() {
	// 	// 		if ptares.Queries[*pair1.Key].MayAlias(ptares.Queries[*pair2.Key]) {
	// 	// 			pp.edges = append(pp.edges, schema.Edge{
	// 	// 				Source: pair1.Value,
	// 	// 				Label:  "PossibleSameChan",
	// 	// 				Target: pair2.Value,
	// 	// 			})
	// 	// 		}
	// 	// 	}
	// 	// }

	// 	for val1, val1Vertex := range processingMap {
	// 		logrus.Tracef("for val: %v\n", val1Vertex.Name)
	// 		for _, label := range ptares.Queries[*val1].PointsTo().Labels() {
	// 			logrus.Tracef("got label: %v\n", label.String())
	// 		}
	// 		for _, label := range ptares.IndirectQueries[*val1].PointsTo().Labels() {
	// 			logrus.Tracef("got indirect label: %v\n", label.String())
	// 		}
	// 	}

	// 	for val1, val1Vertex := range processingMap {
	// 		for val2, val2Vertex := range processingMap {
	// 			if ptares.Queries[*val1].MayAlias(ptares.Queries[*val2]) {
	// 				pp.edges = append(pp.edges, schema.Edge{
	// 					Source: val1Vertex,
	// 					Label:  "PossibleSameChan",
	// 					Target: val2Vertex,
	// 				})
	// 			}
	// 		}
	// 	}

	// 	if pp.ProgressBar {
	// 		ptaProgress.Add(len(processingMap))
	// 	}
	// }
	logrus.Info("points to analysis finished")

	close(pp.vertices)
	vertexWG.Wait()

	logrus.Infof("Added %d vertices", verticesCount)

	//if !pp.ProgressBar {
	//	backend.AddEBulk(pp.edges, func(doneEdges []schema.Edge) {})
	//} else {
	//	progress := progressbar.Default(int64(len(pp.edges)))
	//	backend.AddEBulk(pp.edges, func(doneEdges []schema.Edge) {
	//		progress.Add(len(doneEdges))
	//	})
	//}
}

func (pp ProgParser) processVar(v ssa.Value, pkg string) *schema.Variable {
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
			gVar = &schema.Variable{
				Name:     name,
				Type:     t,
				IsGlobal: isGlobal,
			}
			pp.vertices <- gVar
			if pp.globalVarMap[pkg] == nil {
				pp.globalVarMap[pkg] = make(map[string]*schema.Variable)
			}
			pp.globalVarMap[pkg][name] = gVar
		}
		return gVar
	}

	gVar, ok := pp.varMap[name]
	if !ok {
		gVar = &schema.Variable{
			Name:     name,
			Type:     t,
			IsGlobal: isGlobal,
		}
		pp.vertices <- gVar
		pp.varMap[name] = gVar
	}

	if possibleChan(gVar.Type) {
		logrus.Tracef("found possible chan variable: %v", *gVar)
		// pp.chanMap.Set(&v, gVar)
		pp.chanMap[&v] = gVar
	}
	return gVar
}

func (pp ProgParser) processParam(v ssa.Value, i int) *schema.Parameter {
	if v == nil {
		return nil
	}
	isVar, _, name, t := ParseSSAVar(v)
	if !isVar {
		return nil
	}

	gParam := &schema.Parameter{
		Name:     name,
		Type:     t,
		ParamNum: i,
	}
	pp.vertices <- gParam
	return gParam
}

func (pp ProgParser) processRegister(name string, t string) *schema.Variable {
	gVar, ok := pp.varMap[name]
	if !ok {
		gVar = &schema.Variable{
			Name: name,
			Type: t,
		}
		pp.vertices <- gVar
		pp.varMap[name] = gVar
	}
	return gVar
}

func (pp ProgParser) processInstrByType(instr ssa.Instruction,
	fset *token.FileSet) (*schema.Instruction, []schema.Edge) {
	gInstr, ok := pp.instrMap[instr]
	if !ok {
		text, file, pos := ExtractInstrInfo(fset, instr)
		gInstr = &schema.Instruction{
			File:    file,
			Pos:     pos,
			Text:    text,
			SSAType: reflect.TypeOf(instr).Elem().Name(),
		}
		bName := GetBlockName(fset, instr.Block())
		gBlock := pp.blockMap[bName]
		pp.edges = append(pp.edges, schema.Edge{
			Source: gBlock,
			Label:  "ContainsInstr",
			Target: gInstr,
		})

		switch instr := instr.(type) {
		case ssa.CallInstruction: // *Go, *Defer, *Call
			if instr.Value() != nil {
				varR := pp.processRegister(instr.Value().Name(), instr.Value().Type().String())
				pp.edges = append(pp.edges, schema.Edge{
					Source: gInstr,
					Label:  "FunctionCallResult",
					Target: varR,
				})
				pp.edges = append(pp.edges, schema.Edge{
					Source: gInstr,
					Label:  "AssignTo",
					Target: varR,
				})
			}
			// link callee function
			if g := instr.Common().StaticCallee(); g != nil {
				if gFn, ok := pp.graphFuncMap[g]; ok {
					pp.edges = append(pp.edges, schema.Edge{
						Source: gInstr,
						Label:  "InstrCalls",
						Target: gFn,
					})
				} else {
					// TODO: figure out why some functions aren't resolved!
					// fmt.Printf("error retrieving function %+v\n", g)
				}
			} else {
				for _, g := range pp.calleesOf(instr) {
					if gFn, ok := pp.graphFuncMap[g]; ok {
						pp.edges = append(pp.edges, schema.Edge{
							Source: gInstr,
							Label:  "InstrCalls",
							Target: gFn,
						})
					} else {
						// fmt.Printf("error retrieving function %+v\n", g)
					}
				}
			}
			if instr.Common().Value != nil {
				varVal := pp.processVar(instr.Common().Value, instr.Parent().Pkg.Pkg.Name())
				if varVal != nil {
					pp.edges = append(pp.edges, schema.Edge{
						Source: gInstr,
						Label:  "AssignFrom",
						Target: varVal,
					})
				}
			}
			// link arguments
			for i, arg := range instr.Common().Args {
				varArg := pp.processVar(arg, instr.Parent().Pkg.Pkg.Name())
				if varArg != nil {
					pp.edges = append(pp.edges, schema.Edge{
						Source: gInstr,
						Label:  "AssignFrom",
						Target: varArg,
					})
				}
				paramArg := pp.processParam(arg, i)
				if paramArg != nil {
					pp.edges = append(pp.edges, schema.Edge{
						Source: gInstr,
						Label:  "CallParameter",
						Target: paramArg,
					})
				}
				if paramArg != nil && varArg != nil {
					pp.edges = append(pp.edges, schema.Edge{
						Source: paramArg,
						Label:  "ParamIsVar",
						Target: varArg,
					})
				}
			}
		case *ssa.UnOp:
			varR := pp.processRegister(instr.Name(), instr.Type().String())
			pp.edges = append(pp.edges, schema.Edge{
				Source: gInstr,
				Label:  "AssignTo",
				Target: varR,
			})
			varX := pp.processVar(instr.X, instr.Parent().Pkg.Pkg.Name())
			if varX != nil {
				pp.edges = append(pp.edges, schema.Edge{
					Source: gInstr,
					Label:  "AssignFrom",
					Target: varX,
				})
			}
			if instr.Op == token.ARROW {
				if varX != nil {
					pp.edges = append(pp.edges, schema.Edge{
						Source: gInstr,
						Label:  "ChannelRecv",
						Target: varX,
					})
				}
			}
		case *ssa.BinOp:
			// fmt.Printf("BinOp instr string: %v, op: %v, X: %v, Y: %v\n",
			// 	instr.String(), instr.Op, instr.X, instr.Y)
			// fmt.Printf("X name: %v, string: %v, type: %v, referrers: %v\n",
			// 	instr.X.Name(), instr.X.String(), instr.X.Type(),
			// 	InstrsToString(instr.X.Referrers()))
			// fmt.Printf("Y name: %v, string: %v, type: %v, referrers: %v\n",
			// 	instr.Y.Name(), instr.Y.String(), instr.Y.Type(),
			// 	InstrsToString(instr.Y.Referrers()))
			// fmt.Printf("BinOp register string: %v, type: %v\n",
			// 	instr.Name(), instr.Type())
			varR := pp.processRegister(instr.Name(), instr.Type().String())
			pp.edges = append(pp.edges, schema.Edge{
				Source: gInstr,
				Label:  "AssignTo",
				Target: varR,
			})
			varX := pp.processVar(instr.X, instr.Parent().Pkg.Pkg.Name())
			if varX != nil {
				pp.edges = append(pp.edges, schema.Edge{
					Source: gInstr,
					Label:  "AssignFrom",
					Target: varX,
				})
			}
			varY := pp.processVar(instr.Y, instr.Parent().Pkg.Pkg.Name())
			if varY != nil {
				pp.edges = append(pp.edges, schema.Edge{
					Source: gInstr,
					Label:  "AssignFrom",
					Target: varY,
				})
			}
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
				pp.edges = append(pp.edges, schema.Edge{
					Source: gInstr,
					Label:  "AssignTo",
					Target: varAddr,
				})
			}
			varVal := pp.processVar(instr.Val, instr.Parent().Pkg.Pkg.Name())
			if varVal != nil {
				pp.edges = append(pp.edges, schema.Edge{
					Source: gInstr,
					Label:  "AssignFrom",
					Target: varVal,
				})
			}
		case *ssa.Phi:
			// fmt.Printf("Phi instr string: %v\n",
			// 	instr.String())
			// for _, edge := range instr.Edges {
			// 	fmt.Printf("Phi edge name: %v, string: %v, type: %v, referrers: %v\n",
			// 		edge.Name(), edge.String(), edge.Type(),
			// 		InstrsToString(edge.Referrers()))
			// }
			varR := pp.processRegister(instr.Name(), instr.Type().String())
			pp.edges = append(pp.edges, schema.Edge{
				Source: gInstr,
				Label:  "AssignTo",
				Target: varR,
			})
			for _, edge := range instr.Edges {
				varEdge := pp.processVar(edge, instr.Parent().Pkg.Pkg.Name())
				if varEdge != nil {
					pp.edges = append(pp.edges, schema.Edge{
						Source: gInstr,
						Label:  "AssignFrom",
						Target: varEdge,
					})
				}
			}
		case *ssa.ChangeType: // cannot fail dynamically
			varR := pp.processRegister(instr.Name(), instr.Type().String())
			pp.edges = append(pp.edges, schema.Edge{
				Source: gInstr,
				Label:  "AssignTo",
				Target: varR,
			})
			varX := pp.processVar(instr.X, instr.Parent().Pkg.Pkg.Name())
			if varX != nil {
				pp.edges = append(pp.edges, schema.Edge{
					Source: gInstr,
					Label:  "AssignFrom",
					Target: varX,
				})
			}
		case *ssa.Convert: // cannot fail dynamically
			varR := pp.processRegister(instr.Name(), instr.Type().String())
			pp.edges = append(pp.edges, schema.Edge{
				Source: gInstr,
				Label:  "AssignTo",
				Target: varR,
			})
			varX := pp.processVar(instr.X, instr.Parent().Pkg.Pkg.Name())
			if varX != nil {
				pp.edges = append(pp.edges, schema.Edge{
					Source: gInstr,
					Label:  "AssignFrom",
					Target: varX,
				})
			}
		case *ssa.MultiConvert: // can fail dynamically (see SliceToArrayPointer)
			varR := pp.processRegister(instr.Name(), instr.Type().String())
			pp.edges = append(pp.edges, schema.Edge{
				Source: gInstr,
				Label:  "AssignTo",
				Target: varR,
			})
			varX := pp.processVar(instr.X, instr.Parent().Pkg.Pkg.Name())
			if varX != nil {
				pp.edges = append(pp.edges, schema.Edge{
					Source: gInstr,
					Label:  "AssignFrom",
					Target: varX,
				})
			}
		case *ssa.ChangeInterface: // cannot fail
			varR := pp.processRegister(instr.Name(), instr.Type().String())
			pp.edges = append(pp.edges, schema.Edge{
				Source: gInstr,
				Label:  "AssignTo",
				Target: varR,
			})
			varX := pp.processVar(instr.X, instr.Parent().Pkg.Pkg.Name())
			if varX != nil {
				pp.edges = append(pp.edges, schema.Edge{
					Source: gInstr,
					Label:  "AssignFrom",
					Target: varX,
				})
			}
		case *ssa.SliceToArrayPointer: // can fail
			varR := pp.processRegister(instr.Name(), instr.Type().String())
			pp.edges = append(pp.edges, schema.Edge{
				Source: gInstr,
				Label:  "AssignTo",
				Target: varR,
			})
			varX := pp.processVar(instr.X, instr.Parent().Pkg.Pkg.Name())
			if varX != nil {
				pp.edges = append(pp.edges, schema.Edge{
					Source: gInstr,
					Label:  "AssignFrom",
					Target: varX,
				})
			}
		case *ssa.MakeInterface:
			varR := pp.processRegister(instr.Name(), instr.Type().String())
			pp.edges = append(pp.edges, schema.Edge{
				Source: gInstr,
				Label:  "AssignTo",
				Target: varR,
			})
			varX := pp.processVar(instr.X, instr.Parent().Pkg.Pkg.Name())
			if varX != nil {
				pp.edges = append(pp.edges, schema.Edge{
					Source: gInstr,
					Label:  "AssignFrom",
					Target: varX,
				})
			}
		case *ssa.MakeClosure:
			// TODO
		case *ssa.MakeMap:
		case *ssa.MakeChan:
		case *ssa.MakeSlice:
			varR := pp.processRegister(instr.Name(), instr.Type().String())
			pp.edges = append(pp.edges, schema.Edge{
				Source: gInstr,
				Label:  "AssignTo",
				Target: varR,
			})
			varLen := pp.processVar(instr.Len, instr.Parent().Pkg.Pkg.Name())
			if varLen != nil {
				pp.edges = append(pp.edges, schema.Edge{
					Source: gInstr,
					Label:  "AssignFrom",
					Target: varLen,
				})
			}
			varCap := pp.processVar(instr.Cap, instr.Parent().Pkg.Pkg.Name())
			if varCap != nil {
				pp.edges = append(pp.edges, schema.Edge{
					Source: gInstr,
					Label:  "AssignFrom",
					Target: varCap,
				})
			}
		case *ssa.Slice:
			// 	fmt.Printf("Slice instr string: %v\n",
			// 		instr.String())
			// 	fmt.Printf("Slice X name: %v, string: %v, type: %v, referrers: %v\n",
			// 		instr.X.Name(), instr.X.String(), instr.X.Type(),
			// 		InstrsToString(instr.X.Referrers()))
			// 	if instr.Low != nil {
			// 		fmt.Printf("Slice Low name: %v, string: %v, type: %v, referrers: %v\n",
			// 			instr.Low.Name(), instr.Low.String(), instr.Low.Type(),
			// 			InstrsToString(instr.Low.Referrers()))
			// 	}
			// 	if instr.High != nil {
			// 		fmt.Printf("Slice High name: %v, string: %v, type: %v, referrers: %v\n",
			// 			instr.High.Name(), instr.High.String(), instr.High.Type(),
			// 			InstrsToString(instr.High.Referrers()))
			// 	}
			// 	if instr.Max != nil {
			// 		fmt.Printf("Slice Max name: %v, string: %v, type: %v, referrers: %v\n",
			// 			instr.Max.Name(), instr.Max.String(), instr.Max.Type(),
			// 			InstrsToString(instr.Max.Referrers()))
			// 	}
			varR := pp.processRegister(instr.Name(), instr.Type().String())
			pp.edges = append(pp.edges, schema.Edge{
				Source: gInstr,
				Label:  "AssignTo",
				Target: varR,
			})
			varX := pp.processVar(instr.X, instr.Parent().Pkg.Pkg.Name())
			if varX != nil {
				pp.edges = append(pp.edges, schema.Edge{
					Source: gInstr,
					Label:  "AssignFrom",
					Target: varX,
				})
			}
			varLow := pp.processVar(instr.Low, instr.Parent().Pkg.Pkg.Name())
			if varLow != nil {
				pp.edges = append(pp.edges, schema.Edge{
					Source: gInstr,
					Label:  "AssignFrom",
					Target: varLow,
				})
			}
			varHigh := pp.processVar(instr.High, instr.Parent().Pkg.Pkg.Name())
			if varHigh != nil {
				pp.edges = append(pp.edges, schema.Edge{
					Source: gInstr,
					Label:  "AssignFrom",
					Target: varHigh,
				})
			}
			varMax := pp.processVar(instr.Max, instr.Parent().Pkg.Pkg.Name())
			if varMax != nil {
				pp.edges = append(pp.edges, schema.Edge{
					Source: gInstr,
					Label:  "AssignFrom",
					Target: varMax,
				})
			}
		case *ssa.FieldAddr:
			// fmt.Printf("FieldAddr instr string: %v\n", instr.String())
			// fmt.Printf("FieldAddr X name: %v, string: %v, type: %v, referrers: %v\n",
			// 	instr.X.Name(), instr.X.String(), instr.X.Type(),
			// 	InstrsToString(instr.X.Referrers()))
			varR := pp.processRegister(instr.Name(), instr.Type().String())
			pp.edges = append(pp.edges, schema.Edge{
				Source: gInstr,
				Label:  "AssignTo",
				Target: varR,
			})
			varX := pp.processVar(instr.X, instr.Parent().Pkg.Pkg.Name())
			if varX != nil {
				pp.edges = append(pp.edges, schema.Edge{
					Source: gInstr,
					Label:  "AssignFrom",
					Target: varX,
				})
			}
			gInstr.Misc = fmt.Sprint(instr.Field)
		case *ssa.Field:
			varR := pp.processRegister(instr.Name(), instr.Type().String())
			pp.edges = append(pp.edges, schema.Edge{
				Source: gInstr,
				Label:  "AssignTo",
				Target: varR,
			})
			varX := pp.processVar(instr.X, instr.Parent().Pkg.Pkg.Name())
			if varX != nil {
				pp.edges = append(pp.edges, schema.Edge{
					Source: gInstr,
					Label:  "AssignFrom",
					Target: varX,
				})
			}
			gInstr.Misc = fmt.Sprint(instr.Field)
		case *ssa.IndexAddr:
			varR := pp.processRegister(instr.Name(), instr.Type().String())
			pp.edges = append(pp.edges, schema.Edge{
				Source: gInstr,
				Label:  "AssignTo",
				Target: varR,
			})
			varX := pp.processVar(instr.X, instr.Parent().Pkg.Pkg.Name())
			if varX != nil {
				pp.edges = append(pp.edges, schema.Edge{
					Source: gInstr,
					Label:  "AssignFrom",
					Target: varX,
				})
			}
			varIdx := pp.processVar(instr.Index, instr.Parent().Pkg.Pkg.Name())
			if varIdx != nil {
				pp.edges = append(pp.edges, schema.Edge{
					Source: gInstr,
					Label:  "AssignFrom",
					Target: varIdx,
				})
			}
			if isVar, _, val, t := ParseSSAVar(instr.Index); !isVar {
				if t == "int" {
					gInstr.Misc = val
				}
			}
		case *ssa.Index:
			varR := pp.processRegister(instr.Name(), instr.Type().String())
			pp.edges = append(pp.edges, schema.Edge{
				Source: gInstr,
				Label:  "AssignTo",
				Target: varR,
			})
			varX := pp.processVar(instr.X, instr.Parent().Pkg.Pkg.Name())
			if varX != nil {
				pp.edges = append(pp.edges, schema.Edge{
					Source: gInstr,
					Label:  "AssignFrom",
					Target: varX,
				})
			}
			varIdx := pp.processVar(instr.Index, instr.Parent().Pkg.Pkg.Name())
			if varIdx != nil {
				pp.edges = append(pp.edges, schema.Edge{
					Source: gInstr,
					Label:  "AssignFrom",
					Target: varIdx,
				})
			}
			if isVar, _, val, t := ParseSSAVar(instr.Index); !isVar {
				if t == "int" {
					gInstr.Misc = val
				}
			}
		case *ssa.Lookup:
			varR := pp.processRegister(instr.Name(), instr.Type().String())
			pp.edges = append(pp.edges, schema.Edge{
				Source: gInstr,
				Label:  "AssignTo",
				Target: varR,
			})
			varX := pp.processVar(instr.X, instr.Parent().Pkg.Pkg.Name())
			if varX != nil {
				pp.edges = append(pp.edges, schema.Edge{
					Source: gInstr,
					Label:  "AssignFrom",
					Target: varX,
				})
			}
			varIdx := pp.processVar(instr.Index, instr.Parent().Pkg.Pkg.Name())
			if varIdx != nil {
				pp.edges = append(pp.edges, schema.Edge{
					Source: gInstr,
					Label:  "AssignFrom",
					Target: varIdx,
				})
			}
		case *ssa.Select:
			varR := pp.processRegister(instr.Name(), instr.Type().String())
			pp.edges = append(pp.edges, schema.Edge{
				Source: gInstr,
				Label:  "AssignTo",
				Target: varR,
			})
			for _, st := range instr.States {
				varChan := pp.processVar(st.Chan, instr.Parent().Pkg.Pkg.Name())
				if st.Dir == types.RecvOnly {
					if varChan != nil {
						pp.edges = append(pp.edges, schema.Edge{
							Source: gInstr,
							Label:  "ChannelRecv",
							Target: varChan,
						})
					}
				} else if st.Dir == types.SendOnly {
					if varChan != nil {
						pp.edges = append(pp.edges, schema.Edge{
							Source: gInstr,
							Label:  "ChannelSend",
							Target: varChan,
						})
					}
					varSend := pp.processVar(st.Send, instr.Parent().Pkg.Pkg.Name())
					if varSend != nil {
						pp.edges = append(pp.edges, schema.Edge{
							Source: gInstr,
							Label:  "AssignFrom",
							Target: varSend,
						})
					}
				}
			}
		case *ssa.Range:
			varR := pp.processRegister(instr.Name(), instr.Type().String())
			pp.edges = append(pp.edges, schema.Edge{
				Source: gInstr,
				Label:  "AssignTo",
				Target: varR,
			})
			varX := pp.processVar(instr.X, instr.Parent().Pkg.Pkg.Name())
			if varX != nil {
				pp.edges = append(pp.edges, schema.Edge{
					Source: gInstr,
					Label:  "AssignFrom",
					Target: varX,
				})
			}
		case *ssa.Next:
			varR := pp.processRegister(instr.Name(), instr.Type().String())
			pp.edges = append(pp.edges, schema.Edge{
				Source: gInstr,
				Label:  "AssignTo",
				Target: varR,
			})
			varIter := pp.processVar(instr.Iter, instr.Parent().Pkg.Pkg.Name())
			if varIter != nil {
				pp.edges = append(pp.edges, schema.Edge{
					Source: gInstr,
					Label:  "AssignFrom",
					Target: varIter,
				})
			}
		case *ssa.TypeAssert:
			varR := pp.processRegister(instr.Name(), instr.Type().String())
			pp.edges = append(pp.edges, schema.Edge{
				Source: gInstr,
				Label:  "AssignTo",
				Target: varR,
			})
			varX := pp.processVar(instr.X, instr.Parent().Pkg.Pkg.Name())
			if varX != nil {
				pp.edges = append(pp.edges, schema.Edge{
					Source: gInstr,
					Label:  "AssignFrom",
					Target: varX,
				})
			}
		case *ssa.Extract:
			varR := pp.processRegister(instr.Name(), instr.Type().String())
			pp.edges = append(pp.edges, schema.Edge{
				Source: gInstr,
				Label:  "AssignTo",
				Target: varR,
			})
			varTuple := pp.processVar(instr.Tuple, instr.Parent().Pkg.Pkg.Name())
			if varTuple != nil {
				pp.edges = append(pp.edges, schema.Edge{
					Source: gInstr,
					Label:  "AssignFrom",
					Target: varTuple,
				})
			}
			gInstr.Misc = fmt.Sprint(instr.Index)
		case *ssa.Jump:
		case *ssa.If:
			varCond := pp.processVar(instr.Cond, instr.Parent().Pkg.Pkg.Name())
			if varCond != nil {
				pp.edges = append(pp.edges, schema.Edge{
					Source: gInstr,
					Label:  "AssignFrom",
					Target: varCond,
				})
			}
		case *ssa.Return:
			for i, res := range instr.Results {
				varRes := pp.processVar(res, instr.Parent().Pkg.Pkg.Name())
				if varRes != nil {
					pp.edges = append(pp.edges, schema.Edge{
						Source: gInstr,
						Label:  "AssignFrom",
						Target: varRes,
					})
				}
				paramRes := pp.processParam(res, i)
				if paramRes != nil {
					pp.edges = append(pp.edges, schema.Edge{
						Source: gInstr,
						Label:  "CallParameter",
						Target: paramRes,
					})
				}
				if paramRes != nil && varRes != nil {
					pp.edges = append(pp.edges, schema.Edge{
						Source: paramRes,
						Label:  "ParamIsVar",
						Target: varRes,
					})
				}
			}
		case *ssa.RunDefers:
		case *ssa.Panic:
			varX := pp.processVar(instr.X, instr.Parent().Pkg.Pkg.Name())
			if varX != nil {
				pp.edges = append(pp.edges, schema.Edge{
					Source: gInstr,
					Label:  "AssignFrom",
					Target: varX,
				})
			}
		case *ssa.Send:
			varChan := pp.processVar(instr.Chan, instr.Parent().Pkg.Pkg.Name())
			if varChan != nil {
				pp.edges = append(pp.edges, schema.Edge{
					Source: gInstr,
					Label:  "ChannelSend",
					Target: varChan,
				})
			}
			varX := pp.processVar(instr.X, instr.Parent().Pkg.Pkg.Name())
			if varX != nil {
				pp.edges = append(pp.edges, schema.Edge{
					Source: gInstr,
					Label:  "AssignFrom",
					Target: varX,
				})
			}
		case *ssa.MapUpdate:
			varM := pp.processVar(instr.Map, instr.Parent().Pkg.Pkg.Name())
			if varM != nil {
				pp.edges = append(pp.edges, schema.Edge{
					Source: gInstr,
					Label:  "AssignTo",
					Target: varM,
				})
			}
			varK := pp.processVar(instr.Key, instr.Parent().Pkg.Pkg.Name())
			if varK != nil {
				pp.edges = append(pp.edges, schema.Edge{
					Source: gInstr,
					Label:  "AssignTo",
					Target: varK,
				})
			}
			varVal := pp.processVar(instr.Value, instr.Parent().Pkg.Pkg.Name())
			if varVal != nil {
				pp.edges = append(pp.edges, schema.Edge{
					Source: gInstr,
					Label:  "AssignFrom",
					Target: varVal,
				})
			}
		case *ssa.DebugRef:
			varX := pp.processVar(instr.X, instr.Parent().Pkg.Pkg.Name())
			if varX != nil {
				pp.edges = append(pp.edges, schema.Edge{
					Source: gInstr,
					Label:  "AssignFrom",
					Target: varX,
				})
			}
		case *ssa.Alloc:
			varR := pp.processRegister(instr.Name(), instr.Type().String())
			pp.edges = append(pp.edges, schema.Edge{
				Source: gInstr,
				Label:  "AssignTo",
				Target: varR,
			})
		default:
			fmt.Printf("Missing handling routing for instr value: %v,  type: %v, file: %v,block: %v, operands: %v, parent: %v\n",
				instr.String(), reflect.TypeOf(instr).Elem().Name(),
				fset.PositionFor(instr.Pos(), false),
				instr.Block(), operandsToString(instr.Operands(nil)),
				instr.Parent())
		}

		pp.vertices <- gInstr
	}
	return gInstr, pp.edges
}

func (pp ProgParser) handleBlocks(b *ssa.BasicBlock) []schema.Edge {
	fn := b.Parent()
	pkgPath := fn.Package().Pkg.Path()
	bName := GetBlockName(pp.fsetMap[pkgPath], b)
	curBlock, ok := pp.blockMap[bName]
	if !ok {
		name, file, pos := extractBlockInfo(pp.fsetMap[pkgPath], b)
		gBlock := &schema.Block{
			Name: name,
			File: file,
			Pos:  pos,
		}
		pp.vertices <- gBlock
		pp.blockMap[bName] = gBlock
		curBlock = gBlock
	}

	for _, preB := range b.Preds {
		bName = GetBlockName(pp.fsetMap[pkgPath], preB)
		preBlock, ok := pp.blockMap[bName]
		if !ok {
			name, file, pos := extractBlockInfo(pp.fsetMap[pkgPath], preB)
			gBlock := &schema.Block{
				Name: name,
				File: file,
				Pos:  pos,
			}
			pp.vertices <- gBlock
			pp.blockMap[bName] = gBlock
			preBlock = gBlock
		}
		pp.edges = append(pp.edges, schema.Edge{
			Source: curBlock,
			Label:  "BlockPrev",
			Target: preBlock,
		})
	}

	for _, succB := range b.Succs {
		bName = GetBlockName(pp.fsetMap[pkgPath], succB)
		succBlock, ok := pp.blockMap[bName]
		if !ok {
			name, file, pos := extractBlockInfo(pp.fsetMap[pkgPath], succB)
			gBlock := &schema.Block{
				Name: name,
				File: file,
				Pos:  pos,
			}
			pp.vertices <- gBlock
			pp.blockMap[bName] = gBlock
			succBlock = gBlock
		}
		pp.edges = append(pp.edges, schema.Edge{
			Source: curBlock,
			Label:  "BlockNext",
			Target: succBlock,
		})
	}
	return pp.edges
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
