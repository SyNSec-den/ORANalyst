package condition_collection

import (
	"constraint_collector/PdgGraph"
	"fmt"
	"go/types"
	"strings"

	"golang.org/x/tools/go/ssa"
)

// markErrorBlock Given a function, mark the blocks that only lead to error
func (cc *ConditionCollector) markErrorBlock(f *PdgGraph.Function) map[*PdgGraph.Block]bool {
	res := make(map[*PdgGraph.Block]bool)
	if f.SsaFuncPointer == nil {
		return res
	}

	signature := f.SsaFuncPointer.Signature
	if signature == nil {
		return res
	}

	results := signature.Results()
	if results == nil {
		return res
	}

	errorIdx := -1
	errorType := types.Universe.Lookup("error").Type()
	for i := 0; i < results.Len(); i++ {
		if results.At(i).Type() == errorType {
			errorIdx = i
			break
		}
	}
	if errorIdx == -1 {
		return res
	}

	for _, b := range f.ContainsBlock {
		if b.LastInstr == nil {
			continue
		}
		// find blocks that only return error
		if retInstr, ok := (*b.LastInstr.SsaInstrPtr).(*ssa.Return); ok {
			if constVal, ok := retInstr.Results[errorIdx].(*ssa.Const); ok && constVal.IsNil() {
				// fmt.Printf("non-error block: %v %v, last instruction: %v\n", b.Function.Name, b.Id, b.LastInstr.Text)
				res[b] = false
			} else {
				isErrorBlock := false
				retVar := retInstr.Results[errorIdx]
				fmt.Printf("retVar: %v\n", retVar.Name())
				fmt.Printf("error block: %v %v, last instruction: %v\n", b.Function.Name, b.Id, b.LastInstr.Text)
				instr := b.LastInstr
				for {
					if instr == nil {
						break
					}
					if call, ok := (*instr.SsaInstrPtr).(*ssa.Call); ok && call == retVar {
						fmt.Printf("call: %v\n", call.String())
						switch v := call.Common().Value.(type) {
						case *ssa.Builtin:
							if v.Name() == "panic" {
								isErrorBlock = true
							}
						case *ssa.Function:
							fullName := v.String()

							if fullName == "errors.New" || fullName == "fmt.Errorf" {
								isErrorBlock = true
							}

							if strings.Contains(fullName, "errors.New") {
								isErrorBlock = true
							}
						}

					}
					instr = instr.InstrPrev
				}
				if !isErrorBlock {
					fmt.Printf("cannot decide if error block: %v\n", b.Id)
					printBlock(b)
				}
				res[b] = isErrorBlock
			}
		}
	}
	return res
}

func printBlock(b *PdgGraph.Block) {
	fmt.Printf("block: %v\n", b.Id)
	instr := b.FirstInstr
	for instr != nil {
		fmt.Printf("instr: %v\n", instr.Text)
		instr = instr.InstrNext
	}
}

// processFunc returns true if the function should be explored
// heuristic:
// the function has to be in the package that we are processing
// only explore functions that are called by no more than two different callers
// and the callers are from the same package
// otherwise consider it a harnessing function
func (cc *ConditionCollector) processFunc(f *PdgGraph.Function) bool {
	if b, ok := cc.exploreFuncs[f]; ok {
		return b
	}

	fmt.Printf("package name: %v, module name: %v\n", f.Package.FullName(), cc.moduleName)
	if !strings.Contains(f.Package.FullName(), cc.moduleName) {
		cc.exploreFuncs[f] = false
		return false
	}

	if len(f.Caller) > 2 {
		cc.exploreFuncs[f] = false
		return false
	}

	pkgMap := make(map[*PdgGraph.Package]bool)
	for _, info := range f.Caller {
		pkgMap[info.Package] = true
	}
	if len(pkgMap) > 1 {
		cc.exploreFuncs[f] = false
		return false
	}

	cc.exploreFuncs[f] = true
	return true

	// fmt.Printf("\n\nfunction: %v\n", f.Name)
	// for name, info := range f.Caller {
	// 	fmt.Printf("caller: %v %+v\n", name, *info)
	// }
	// return f.Name == "AssociateRanToE2THandlerImpl" || f.Name == "validateE2TAddressRANListData" ||
	// 	f.Name == "checkValidaE2TAddress"
	// return f.Name != "Warn" && f.Name != "Debug" && f.Name != "Info" && f.Name != "Error"
}

func (cc *ConditionCollector) genFunctionConstraints(f *PdgGraph.Function) [][]Constraint {
	// if !cc.processFunc(f) {
	// 	return nil
	// }
	return cc.genErrorConstraints(f)
}

// genErrorConstraints generate a set of constraints so the function doesn't return error
func (cc *ConditionCollector) genErrorConstraints(f *PdgGraph.Function) [][]Constraint {
	targetBlocks := cc.markErrorBlock(f)
	// allConditions := make(map[Condition]bool)
	constraints := make([][]Constraint, 0)
	for b, isErrorBlock := range targetBlocks {
		// if isErrorBlock {
		// 	conditions := cc.FindConditionsToBlock(b)
		// 	for _, cond := range conditions {
		// 		constraints := cc.SolveCondition(cond)
		// 		for _, c := range constraints {
		// 			c.BoolCondition = !c.BoolCondition
		// 			constraints = append(constraints, c)
		// 		}
		// 	}
		// }
		if !isErrorBlock {
			blockConstraint := make([]Constraint, 0)
			conditions := cc.FindConditionsToBlock(b)
			for c := range conditions {
				constraint := cc.SolveCondition(c)
				blockConstraint = append(blockConstraint, constraint...)
			}
			constraints = append(constraints, blockConstraint)
		}
	}

	return constraints
}
