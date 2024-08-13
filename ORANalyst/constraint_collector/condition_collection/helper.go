package condition_collection

import (
	"constraint_collector/PdgGraph"
	"fmt"
	"os"
)

func PrintConstraint(c Constraint) string {
	return DescribeConstraint(c, "")
}

func printPath(predecessor map[*PdgGraph.Block]*PdgGraph.Block, source, target *PdgGraph.Block) {
	path := []*PdgGraph.Block{}
	for at := target; at != nil; at = predecessor[at] {
		path = append(path, at)
		if at == source {
			break
		}
	}
	// Reverse the path to start from the source
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	// Print the path
	fmt.Printf("The CFG path from %v to %v is:\n", source.Id, target.Id)
	for _, block := range path {
		fmt.Printf("%v -> ", block.Id)
	}
	fmt.Printf("\n")
}

func DescribeConstraint(c Constraint, indent string) string {
	resStr := ""
	resStr += fmt.Sprintf("%v%v\n", indent, c.AssignStatement.Text)
	resStr += fmt.Sprintf("%v%v\n", indent, c.AssignStatementString)
	resStr += fmt.Sprintf("%v%v\n", indent, c.ConditionAssignString)
	resStr += fmt.Sprintf("%v%v - %v\n\n", indent, c.Condition, c.BoolCondition)
	// resStr += fmt.Sprintf("%v%v - %v\n\n", indent, c.ConditionAssign.Block.Function.Name, c.ConditionAssign.Block.Id)
	if len(c.ExpandConstraints) > 0 {
		resStr += fmt.Sprintf("%vconstraint can be expanded:\n", indent)
		for i, constraints := range c.ExpandConstraints {
			for _, c := range constraints {
				resStr += DescribeConstraint(c, indent+"\t")
			}
			if i != len(c.ExpandConstraints)-1 {
				resStr += fmt.Sprintf("%vOR\n", indent)
			}
		}
	}
	return resStr
}

func PrintCallPath(path []*PdgGraph.Function) string {
	res := ""
	for _, f := range path {
		res += fmt.Sprintf("%v -> ", f.Name)
	}
	res = res[0:len(res)-4] + "\n"
	return res
}

func (cc *ConditionCollector) saveConstraintAndReturn(condition Condition,
	constraint []Constraint) []Constraint {
	cc.constraintMap[condition] = constraint
	return constraint
}

func PrintFuncSSA(f *PdgGraph.Function) {
	for _, block := range f.ContainsBlock {
		fmt.Printf("\n\nblock %v\n", block.Id)
		fmt.Printf("predecessors: ")
		for _, pred := range block.Predecessors {
			fmt.Printf("%v ", pred.Id)
		}
		fmt.Printf("\nsuccessors: ")
		for _, succ := range block.Successors {
			fmt.Printf("%v ", succ.Id)
		}
		fmt.Printf("\n")
		instr := block.FirstInstr
		for instr != nil {
			fmt.Printf("%v\n", instr.Text)
			instr = instr.InstrNext
		}
	}
}

// PrintFuncSSAToDot generates a DOT file representation of the control flow graph of a function.
func PrintFuncSSAToDot(f *PdgGraph.Function, dest string) {
	file, err := os.Create(dest)
	if err != nil {
		fmt.Printf("Error creating file: %s\n", err)
		return
	}
	defer file.Close()

	fmt.Fprintf(file, "digraph G {\n")

	for _, block := range f.ContainsBlock {
		fmt.Fprintf(file, "    \"%v\" [label=\"Block %v\"];\n", block.Id, block.Id)
		for _, succ := range block.Successors {
			fmt.Fprintf(file, "    \"%v\" -> \"%v\";\n", block.Id, succ.Id)
		}
	}

	fmt.Fprintf(file, "}\n")
}

// makeInstrCopy creates a deep copy of an Instruction, except for pointer fields.
func makeInstrCopy(orig *PdgGraph.Instruction) *PdgGraph.Instruction {
	if orig == nil {
		return nil
	}

	copy := &PdgGraph.Instruction{
		Key:                orig.Key,
		Id:                 orig.Id,
		File:               orig.File,
		Offset:             orig.Offset,
		Text:               orig.Text,
		SSAType:            orig.SSAType,
		ASTType:            orig.ASTType,
		Pos:                orig.Pos,
		Misc:               orig.Misc,
		Next:               orig.Next,        // shallow copy
		References:         orig.References,  // shallow copy
		Assigns:            orig.Assigns,     // shallow copy
		Block:              orig.Block,       // shallow copy
		SsaInstrPtr:        orig.SsaInstrPtr, // shallow copy
		InstrType:          orig.InstrType,
		InstrNext:          orig.InstrNext,          // shallow copy
		InstrPrev:          orig.InstrPrev,          // shallow copy
		FunctionCallResult: orig.FunctionCallResult, // shallow copy
		CalledFnArguments:  orig.CalledFnArguments,  // shallow copy
		AssignTo:           orig.AssignTo,           // shallow copy
		InstrCalls:         orig.InstrCalls,         // shallow copy
		AssignFrom:         orig.AssignFrom,         // shallow copy
		CallParameter:      orig.CallParameter,      // shallow copy
		CallerArguments:    orig.CallerArguments,    // shallow copy
		ReturnValues:       orig.ReturnValues,       // shallow copy
		PhiMap:             orig.PhiMap,             // shallow copy
		ChannelRecv:        orig.ChannelRecv,        // shallow copy
		ChannelSend:        orig.ChannelSend,        // shallow copy
		UsesDefinition:     orig.UsesDefinition,     // shallow copy
		TrueJump:           orig.TrueJump,
		FalseJump:          orig.FalseJump,
		UnCondJump:         orig.UnCondJump,
		LowVal:             orig.LowVal,
		HighVal:            orig.HighVal,
		Operator:           orig.Operator,
		OperandX:           orig.OperandX,
		OperandY:           orig.OperandY,
	}

	return copy
}
