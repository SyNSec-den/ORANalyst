package condition_collection

import (
	"constraint_collector/PdgGraph"
	"fmt"
	"go/types"

	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa"
)

type ConditionCollector struct {
	conditionToBlock map[*PdgGraph.Block]map[Condition]bool // a set of conditions to reach a block (AND relations)
	constraintMap    map[Condition][]Constraint
	typeTracker      *typeTracker
	exploreFuncs     map[*PdgGraph.Function]bool
	moduleName       string
}

type Condition struct {
	Instr *PdgGraph.Instruction // The instruction that contains the condition - must be a if statement
	Cond  bool                  // The condition to satisfy
}

func NewConditionCollector(module string, trackedTypes []types.Type,
	exploreFuncs map[*PdgGraph.Function]bool, typePkg packages.Package) *ConditionCollector {
	if exploreFuncs == nil {
		exploreFuncs = make(map[*PdgGraph.Function]bool)
	}
	return &ConditionCollector{
		conditionToBlock: make(map[*PdgGraph.Block]map[Condition]bool),
		constraintMap:    make(map[Condition][]Constraint),
		typeTracker:      NewTypeTracker(trackedTypes, typePkg),
		exploreFuncs:     exploreFuncs,
		moduleName:       module,
	}
}

func (cc *ConditionCollector) CollectPathConditions(entry *PdgGraph.Function,
	exit *PdgGraph.Function) ([][]*PdgGraph.Function, []map[Condition]bool) {
	callPaths := cc.CollectCallPaths(entry, exit)
	if len(callPaths) == 0 {
		return callPaths, nil
	}
	printCallPaths(callPaths)
	conditions := make([]map[Condition]bool, 0)
	for j, path := range callPaths {
		fmt.Printf("-------------------- path %v --------------------\n", j)
		pathConditions := make([]map[Condition]bool, 0)
		for i := 0; i < len(path)-1; i++ {
			callConds := cc.FindCallConds(path[i], path[i+1])
			fmt.Printf("\tconditions for func call %v -> %v: \n", path[i].Name, path[i+1].Name)
			for _, conds := range callConds {
				fmt.Printf("-------------------- conds --------------------\n")
				for cond := range conds {
					fmt.Printf("\t\t%v %v\n", cond.Instr.Text, cond.Cond)
				}
			}
			pathConditions = combineConditions(pathConditions, callConds)
		}
		conditions = append(conditions, pathConditions...)
	}

	for i, cond := range conditions {
		fmt.Printf("-------------------- path %v --------------------\n", i)
		for c := range cond {
			fmt.Printf(" %v %v\n", c.Instr.Text, c.Cond)
		}
	}

	return callPaths, conditions
}

// combineConditions return the cartesian product of two sets of conditions
func combineConditions(c1 []map[Condition]bool, c2 []map[Condition]bool) []map[Condition]bool {
	if len(c1) == 0 {
		return c2
	}
	if len(c2) == 0 {
		return c1
	}

	conditions := make([]map[Condition]bool, 0)
	for i := 0; i < len(c1); i++ {
		for j := 0; j < len(c2); j++ {
			cond := make(map[Condition]bool)
			for c := range c1[i] {
				cond[c] = true
			}
			for c := range c2[j] {
				cond[c] = true
			}
			conditions = append(conditions, cond)
		}
	}
	return conditions
}

func (cc *ConditionCollector) CollectCallPaths(entry *PdgGraph.Function, exit *PdgGraph.Function) [][]*PdgGraph.Function {
	return cc.findPaths(entry, exit, make([]*PdgGraph.Function, 0), make([][]*PdgGraph.Function, 0))
}

func printCallPaths(paths [][]*PdgGraph.Function) {
	res := ""
	for i, path := range paths {
		res += fmt.Sprintf("-------------------- func path %v --------------------\n", i)
		for _, funcInfo := range path {
			res += fmt.Sprintf("%v -> ", funcInfo.Name)
		}
		res = res[0:len(res)-4] + "\n"
	}
	fmt.Printf("%v\n", res)
}

func (cc *ConditionCollector) findPaths(current *PdgGraph.Function, exit *PdgGraph.Function,
	path []*PdgGraph.Function, allPaths [][]*PdgGraph.Function) [][]*PdgGraph.Function {
	fmt.Printf("current: %v, exit: %v\n", current.Name, exit.Name)
	if current.Name == exit.Name && current != exit {
		fmt.Printf("current package: %v, exit package: %v\n", current.Package.FullName(), exit.Package.FullName())
	}
	path = append(path, current)
	// TODO: Figure out why we need the second part for function matching
	if current == exit || (current.Name == exit.Name && current.FullName() == exit.FullName()) {
		newPath := make([]*PdgGraph.Function, len(path))
		copy(newPath, path)
		allPaths = append(allPaths, newPath)
		return allPaths
	}

	for _, call := range current.Calls {
		// if pe.processFunc(call) {
		allPaths = cc.findPaths(call, exit, path, allPaths)
		// }
	}

	path = path[:len(path)-1]
	return allPaths
}

// FindCallConds context-insentive function-wise exploration.
// Find conditions to reach a function.
// The function has to contain a call to the target function.
func (cc *ConditionCollector) FindCallConds(cur *PdgGraph.Function,
	target *PdgGraph.Function) []map[Condition]bool {
	fmt.Printf("cur: %v, target: %v\n", cur.Name, target.Name)
	found := false
	for _, call := range cur.Calls {
		if call == target {
			found = true
			break
		}
	}
	if !found {
		return nil
	}

	conditions := make([]map[Condition]bool, 0)
	for _, b := range cur.ContainsBlock {
		fmt.Printf("cur block: %v\n", b.Id)
		curInstr := b.FirstInstr
		for {
			if curInstr == nil {
				break
			}
			fmt.Printf("curInstr: %v\n", curInstr.Text)
			if _, ok := (*curInstr.SsaInstrPtr).(ssa.CallInstruction); ok {
				fmt.Printf("curInstr is call: %v\n", curInstr.Text)
				for _, call := range curInstr.InstrCalls {
					fmt.Printf("call name: %v, target name: %v\n", call.Name, target.Name)
					if call == target {
						fmt.Printf("Processing %v in block %v for condition\n", call.Name, b.Id)
						blockConds := cc.FindConditionsToBlock(b)
						conditions = append(conditions, blockConds)
					}
				}
			}
			curInstr = curInstr.InstrNext
		}
	}
	return conditions
}

// Function to find conditions to reach a block in the CFG
func (cc *ConditionCollector) FindConditionsToBlock(targetBlock *PdgGraph.Block) map[Condition]bool {
	conditions := make(map[Condition]bool)
	visitedBlocks := make(map[*PdgGraph.Block]bool)
	fmt.Printf("finding conditions to block: %v %v\n", targetBlock.Function.Name, targetBlock.Id)

	// Function to recursively find conditions
	var findConditions func(block *PdgGraph.Block)
	findConditions = func(block *PdgGraph.Block) {
		if visitedBlocks[block] {
			return
		}
		visitedBlocks[block] = true
		fmt.Printf("target block: %v\n", block.Id)
		// if cached, ok := cc.conditionToBlock[block]; ok {
		// 	for cond := range cached {
		// 		conditions[cond] = true
		// 	}
		// 	return
		// }

		for controller := range block.ControlDepBlocks {
			fmt.Printf("controller: %v ", controller.Id)
		}
		fmt.Printf("\n")

		for controller := range block.ControlDepBlocks {
			if controller.Id == "-1" || controller.Id == block.Id {
				continue
			}
			if _, ok := (*controller.LastInstr.SsaInstrPtr).(*ssa.If); !ok {
				panic(fmt.Sprintf("CFG controller last statement is not if: %v", controller.LastInstr.Text))
			}

			var findTarget = func(orig *PdgGraph.Block, next *PdgGraph.Block, target *PdgGraph.Block) bool {
				visited := make(map[*PdgGraph.Block]bool)
				predecessor := make(map[*PdgGraph.Block]*PdgGraph.Block) // Map to track the predecessor of each block
				visited[next] = true
				visited[orig] = true
				if next == target {
					printPath(predecessor, next, target) // Print the path
					return true
				}
				var stack []*PdgGraph.Block
				stack = append(stack, next)

				for len(stack) > 0 {
					current := stack[len(stack)-1]
					stack = stack[:len(stack)-1]

					if current == target {
						printPath(predecessor, next, target) // Print the path
						return true
					}

					for _, succ := range current.Successors {
						if succ == target {
							predecessor[target] = current        // Set the predecessor for the target
							printPath(predecessor, next, target) // Print the path
							return true
						}
						if !visited[succ] {
							visited[succ] = true
							predecessor[succ] = current // Set the predecessor for the successor
							stack = append(stack, succ)
						}
					}
				}
				return false
			}

			fmt.Printf("findTarget func: %v, source: %v, target: %v\n", controller.Function.Name, controller.Id, targetBlock.Id)
			cond := true
			fmt.Printf("findTarget TrueSuccessor: %v, FalseSuccessor: %v\n", controller.TrueSuccessor.Id, controller.FalseSuccessor.Id)
			if findTarget(controller, controller.FalseSuccessor, targetBlock) {
				cond = false
				fmt.Printf("findTarget FalseSuccessor: %v, target: %v\n", controller.FalseSuccessor.Id, targetBlock.Id)
			} else if !findTarget(controller, controller.TrueSuccessor, targetBlock) {
				panic(fmt.Sprintf("target block %v is not in the true or false branch of %v:%v",
					targetBlock.Id, controller.Id, controller.LastInstr.Text))
			}

			fmt.Printf("found condition: %v %v\n", controller.LastInstr.Text, cond)
			conditions[Condition{
				Instr: controller.LastInstr,
				Cond:  cond,
			}] = true
			findConditions(controller)
		}
		cc.conditionToBlock[block] = conditions
	}

	if targetBlock != nil {
		findConditions(targetBlock)
	}

	for c := range conditions {
		fmt.Printf("condition: %v %v\n", c.Instr.Text, c.Cond)
	}
	return conditions
}
