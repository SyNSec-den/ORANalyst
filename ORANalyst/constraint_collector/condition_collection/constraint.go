package condition_collection

import (
	"constraint_collector/PdgGraph"
	"fmt"
	"go/types"
	"strings"

	"golang.org/x/tools/go/ssa"
)

type Constraint struct {
	AssignStatement       *PdgGraph.Instruction
	AssignStatementString string
	IfStatement           *PdgGraph.Instruction
	Condition             string
	ConditionAssign       *PdgGraph.Instruction
	ConditionAssignString string
	BoolCondition         bool
	TV                    TrackedVar
	ExpandConstraints     [][]Constraint
	// expand the current constraint
	// for example, if the constraint is validate(input.Field) != nil,
	// then the expandConstraints contains the constraints in validate(input.Field) function
}

type edge struct {
	from *PdgGraph.Block
	to   *PdgGraph.Block
}

type TrackedVar struct {
	TrackedVars  map[*PdgGraph.Variable]bool
	Instrs       []*PdgGraph.Instruction
	visitedEdges map[edge]bool
	condition    bool
}

func copyTrackedVar(tv TrackedVar) TrackedVar {
	res := TrackedVar{
		TrackedVars:  make(map[*PdgGraph.Variable]bool),
		Instrs:       make([]*PdgGraph.Instruction, 0),
		visitedEdges: make(map[edge]bool),
		condition:    tv.condition,
	}
	for k, v := range tv.TrackedVars {
		res.TrackedVars[k] = v
	}
	for _, instr := range tv.Instrs {
		res.Instrs = append(res.Instrs, instr)
	}
	for k, v := range tv.visitedEdges {
		res.visitedEdges[k] = v
	}
	return res
}

func (cc *ConditionCollector) SolveCondition(condition Condition) []Constraint {
	if res, ok := cc.constraintMap[condition]; ok {
		return res
	}

	fmt.Printf("solving condition: %v\n", condition.Instr.Text)
	instr := condition.Instr
	if _, ok := (*instr.SsaInstrPtr).(*ssa.If); !ok {
		panic("not an if statement")
	}
	// instr = makeInstrCopy(instr)
	// instr.Text = cond
	trackedVars := make(map[*PdgGraph.Variable]bool)
	instrs := make([]*PdgGraph.Instruction, 0)
	instrs = append(instrs, instr)
	for _, assignFrom := range instr.AssignFrom {
		trackedVars[assignFrom] = true
	}
	trackedVar := TrackedVar{
		TrackedVars:  trackedVars,
		Instrs:       instrs,
		visitedEdges: make(map[edge]bool),
		condition:    condition.Cond,
	}
	prevInstr := instr.InstrPrev

	var backwardTaint func(tv TrackedVar, curInstr *PdgGraph.Instruction) []Constraint
	backwardTaint = func(tv TrackedVar, curInstr *PdgGraph.Instruction) []Constraint {
		prevBlocks := make(map[*PdgGraph.Block]bool)
		curBlock := instr.Block
		for {
			if curInstr == nil {
				for _, pb := range curBlock.Predecessors {
					fmt.Printf("previous blocks: %v\n", pb.Id)
				}
				break
			}
			var tainted bool
			var assignToTainted bool
			tv, tainted, assignToTainted = taintVarBackward(curInstr, tv)
			_ = tainted
			if assignToTainted {
				fmt.Printf("tainted: %v, type: %v\n", curInstr.Text, curInstr.SSAType)
				switch ins := (*curInstr.SsaInstrPtr).(type) {
				case ssa.CallInstruction:
					if ins.Value().Call.Value != nil && cc.typeTracker.isTrackedType(ins.Value().Call.Value.Type()) {
						return cc.saveConstraintAndReturn(condition, cc.GenConstraint(tv))
					}
					for _, arg := range ins.Value().Call.Args {
						if cc.typeTracker.isTrackedType(arg.Type()) {
							return cc.saveConstraintAndReturn(condition, cc.GenConstraint(tv))
						}
					}
				case *ssa.MakeInterface:
					if cc.typeTracker.isTrackedType(ins.X.Type()) {
						return cc.saveConstraintAndReturn(condition, cc.GenConstraint(tv))
					}
				case *ssa.FieldAddr:
					if cc.typeTracker.isTrackedType(ins.X.Type()) {
						return cc.saveConstraintAndReturn(condition, cc.GenConstraint(tv))
					}
				case *ssa.Phi:
					for i, v := range ins.Edges {
						if _, ok := v.(*ssa.Const); ok {
							// constant is not assigned from the input, skip
							fmt.Printf("phi node edge is const: %v, skipping\n", v)
							continue
						}
						predBlock := ins.Block().Preds[i] // Get the predecessor block corresponding to the phi edge
						blockNumber := predBlock.Index
						keyPrefix := ""
						for k := range instr.Block.Function.ContainsBlock {
							if kSplit := strings.Split(k, ":"); len(kSplit) > 0 {
								keyPrefix = k[0 : len(k)-len(kSplit[len(kSplit)-1])]
								break
							}
						}
						// if phiB := instr.Block.Function.ContainsBlock[fmt.Sprintf("%v%v",
						// 	keyPrefix, blockNumber)]; phiB != nil {
						phiB := instr.Block.Function.ContainsBlock[fmt.Sprintf("%v%v", keyPrefix, blockNumber)]
						prevBlocks[phiB] = true
						fmt.Printf("added phi block: %v\n", phiB.Id)
						// }
					}
				}
			}
			curInstr = curInstr.InstrPrev
		}
		res := make([]Constraint, 0)
		// values are assigned from phi instructions, follow that flow
		if len(prevBlocks) > 0 {
			for pb := range prevBlocks {
				if visited := tv.visitedEdges[edge{from: curBlock, to: pb}]; visited {
					continue
				}
				fmt.Printf("phi pb lastInstr: %v, pb Id: %v\n", pb.LastInstr.Text, pb.Id)
				cpTv := copyTrackedVar(tv)
				cpTv.visitedEdges[edge{from: curBlock, to: pb}] = true
				tempRes := backwardTaint(cpTv, pb.LastInstr)
				res = append(res, tempRes...)
			}
			return cc.saveConstraintAndReturn(condition, res)
		} else if len(curBlock.Predecessors) > 0 {
			for _, pb := range curBlock.Predecessors {
				if visited := tv.visitedEdges[edge{from: curBlock, to: pb}]; visited {
					continue
				}
				// fmt.Printf("pb lastInstr: %v, pb Id: %v\n", pb.LastInstr.Text, pb.Id)
				fmt.Printf(" pb Id: %v\n", pb.Id)
				cpTv := copyTrackedVar(tv)
				cpTv.visitedEdges[edge{from: curBlock, to: pb}] = true
				return cc.saveConstraintAndReturn(condition, backwardTaint(cpTv, pb.LastInstr))
			}
		}
		return cc.saveConstraintAndReturn(condition, res)
	}
	re := backwardTaint(trackedVar, prevInstr)
	if len(re) == 0 {
		fmt.Printf("no constraint generated\n")
	}
	return backwardTaint(trackedVar, prevInstr)
}

func (cc *ConditionCollector) GenConstraint(trackedVar TrackedVar) (constraint []Constraint) {
	fmt.Printf("generating constraint for tracked variable\n")
	res := make([]Constraint, 0)
	instrs := trackedVar.Instrs
	assignInstr := instrs[0]
	var assignStr string
	funcConstraints := make([][]Constraint, 0)
	switch ins := (*assignInstr.SsaInstrPtr).(type) {
	case ssa.CallInstruction:
		process := false
		if ins.Value().Call.Value != nil && cc.typeTracker.isTrackedType(ins.Value().Call.Value.Type()) {
			process = true
			t := ins.Value().Call.Value.Type()
			if sigT, ok := t.(*types.Signature); ok {
				t = sigT.Recv().Type()
			}
			// fmt.Printf("variable: %v\n", t)
			// fmt.Printf("generating constraint for input function: %v.%v\n", t, ins.Call.Value.Name())
			assignStr = fmt.Sprintf("%v.%v", t, ins.Value().Call.Value.Name())
		} else {
			for _, arg := range ins.Value().Call.Args {
				if cc.typeTracker.isTrackedType(arg.Type()) {
					process = true
					assignStr = fmt.Sprintf("%v(%v)", ins.Value().Call.Value.Name(), arg.Type())
				}
			}
		}
		if process {
			for _, call := range assignInstr.InstrCalls {
				funcConstraints = append(funcConstraints, cc.genFunctionConstraints(call)...)
			}
		}
	case *ssa.FieldAddr:
		if cc.typeTracker.isTrackedType(ins.X.Type()) {
			// fmt.Printf("variable: %v\n", ins.X.Type())
			_, fieldName, err := getField(ins.X.Type(), ins.Field)
			if err != nil {
				panic(err)
			}
			// fmt.Printf("generating constraint for input field: %v.%v\n", ins.X.Type(), fieldName)
			assignStr = fmt.Sprintf("%v.%v", ins.X.Type(), fieldName)
		}
	case *ssa.MakeInterface:
		if cc.typeTracker.isTrackedType(ins.X.Type()) {
			assignStr = fmt.Sprintf("%v", ins.X.Type())
		}
	}

	for i, instr := range instrs {
		switch ssaInstr := (*instr.SsaInstrPtr).(type) {
		case *ssa.If:
			fmt.Printf("the condition is: %v\n", instr.Text)
			var assignFrom *PdgGraph.Variable
			for _, af := range instr.AssignFrom {
				assignFrom = af
			}
			if len(instr.AssignFrom) != 1 {
				panic(fmt.Sprintf("if instruction does not assign from one variable: %v, len assign from: %v\n",
					instr.Text, len(instr.AssignFrom)))
			}

			// do a back data-flow to find the possible assignment for the condition
			tainted := make(map[*PdgGraph.Variable]bool)
			tainted[assignFrom] = true
			var conditionAssign *PdgGraph.Instruction
			var conditionAssignStr string
			foundConditionAssign := false
			for j := i - 1; j >= 0; j-- {
				if foundConditionAssign {
					break
				}
				prevInstr := instrs[j]
				for _, assignTo := range prevInstr.AssignTo {
					if tainted[assignTo] {
						for _, assignFrom := range prevInstr.AssignFrom {
							tainted[assignFrom] = true
						}
						switch (*prevInstr.SsaInstrPtr).(type) {
						case ssa.CallInstruction, *ssa.BinOp, *ssa.TypeAssert:
							conditionAssign = prevInstr
							conditionAssignStr = prevInstr.Text
							res = append(res, Constraint{
								AssignStatement:       assignInstr,
								AssignStatementString: assignStr,
								IfStatement:           instr,
								Condition:             instr.Text,
								ConditionAssign:       conditionAssign,
								ConditionAssignString: conditionAssignStr,
								BoolCondition:         trackedVar.condition,
								TV:                    trackedVar,
								ExpandConstraints:     funcConstraints,
							})
							foundConditionAssign = true
						}
					}
				}
			}
		case ssa.CallInstruction:
			if builtin, ok := ssaInstr.Common().Value.(*ssa.Builtin); ok {
				if builtin.Name() == "len" {
					// a for loop iterating the slice
					// return make([]Constraint, 0)
				}
			}
		}
	}
	return res
}

// return val: updated trackedVar, tainted, if the assignTo is tainted
func taintVarBackward(instr *PdgGraph.Instruction,
	trackedVar TrackedVar) (TrackedVar, bool, bool) {
	assignToTainted := false
	assignFromTainted := false
	for _, assignTo := range instr.AssignTo {
		if !assignToTainted && trackedVar.TrackedVars[assignTo] {
			fmt.Printf("tainted because of assignTo: %v\n", assignTo.Name)
			assignToTainted = true
			trackedVar.Instrs = append([]*PdgGraph.Instruction{instr}, trackedVar.Instrs...) //prepend
			for _, assignTo := range instr.AssignTo {
				trackedVar.TrackedVars[assignTo] = true
			}
			for _, AssignFrom := range instr.AssignFrom {
				trackedVar.TrackedVars[AssignFrom] = true
			}
		}
	}

	if !assignToTainted {
		for _, assignFrom := range instr.AssignFrom {
			if !assignFromTainted && trackedVar.TrackedVars[assignFrom] {
				fmt.Printf("tainted because of assignFrom: %v\n", assignFrom.Name)
				assignFromTainted = true
				trackedVar.Instrs = append([]*PdgGraph.Instruction{instr}, trackedVar.Instrs...) //prepend
				for _, assignTo := range instr.AssignTo {
					trackedVar.TrackedVars[assignTo] = true
				}
				for _, AssignFrom := range instr.AssignFrom {
					trackedVar.TrackedVars[AssignFrom] = true
				}
			}
		}
	}

	return trackedVar, assignFromTainted || assignToTainted, assignToTainted
}
