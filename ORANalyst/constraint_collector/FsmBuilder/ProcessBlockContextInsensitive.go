package FsmBuilder

import "constraint_collector/PdgGraph"

func ProcessBlockContextInsensitive(block *PdgGraph.Block) {
	if block != block.Function.CfgGraph.StartNode && block != block.Function.CfgGraph.EndNode {
		currInst := block.FirstInstr
		for currInst.InstrNext != nil {
			ProcessInstructionContextInsensitive(currInst)
			currInst = currInst.InstrNext
		}
		if currInst != nil {
			ProcessInstructionContextInsensitive(currInst)
		}
	}
}
