package FsmBuilder

import "constraint_collector/PdgGraph"

func ProcessFunctionContextInsensitive(fn *PdgGraph.Function) {

	topSortedBlocks := PdgGraph.TopologicalSort(fn)

	for _, block := range topSortedBlocks {
		CalculateBlockReachingCondition(block)
		ProcessBlockContextInsensitive(block)
	}

	//for _, block := range topSortedBlocks {
	//	fmt.Printf("Block %s \n Reaching Condition %s\n", block.FullName(),
	//		block.ReachingConditionExpression.GetString(nil))
	//	fmt.Printf("\n")
	//}
}
