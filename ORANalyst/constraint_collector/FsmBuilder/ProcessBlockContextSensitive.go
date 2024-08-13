package FsmBuilder

import "constraint_collector/PdgGraph"

func ProcessBlockContextSensitive(block *PdgGraph.Block,
	inputTransitions map[string]*PdgGraph.FSMTransition,
	callStack *PdgGraph.CallStack,
	functionFSM *PdgGraph.FSM,
	controlFlow *PdgGraph.Flow) map[string]*PdgGraph.FSMTransition {

	//fmt.Printf("Current Block %s ControlFlow : ", block.FullName())
	//controlFlow.PrintFlow()
	for _, v := range inputTransitions {
		functionFSM.AddToBlockVisited(block, v)
	}

	currentInputTransitions := make(map[string]*PdgGraph.FSMTransition)
	for k, v := range inputTransitions {
		currentInputTransitions[k] = v
	}

	if block != block.Function.CfgGraph.StartNode && block != block.Function.CfgGraph.EndNode {
		currInst := block.FirstInstr

		for currInst.InstrNext != nil {
			newInputTransitions := ProcessInstructionContextSensitive(currInst,
				currentInputTransitions,
				callStack,
				functionFSM,
				controlFlow)
			for k := range currentInputTransitions {
				delete(currentInputTransitions, k)
			}
			for k, newTransition := range newInputTransitions {
				currentInputTransitions[k] = newTransition
			}
			currInst = currInst.InstrNext
		}
		if currInst != nil {
			newInputTransitions := ProcessInstructionContextSensitive(currInst,
				currentInputTransitions,
				callStack,
				functionFSM,
				controlFlow)
			for k := range currentInputTransitions {
				delete(currentInputTransitions, k)
			}
			for k, newTransition := range newInputTransitions {
				currentInputTransitions[k] = newTransition
			}
		}

		for _, v := range currentInputTransitions {
			functionFSM.AddToBlockOutputs(block, v)
		}

		for _, succBlock := range block.Successors {
			for _, v := range currentInputTransitions {
				functionFSM.AddToBlockInputs(succBlock, v)
			}
		}
	}

	if block == block.Function.CfgGraph.StartNode {

		for _, v := range inputTransitions {
			functionFSM.AddToBlockOutputs(block, v)
		}

		for _, succBlock := range block.Successors {
			for _, v := range functionFSM.BlockInfo[block].OutputTransitions {
				functionFSM.AddToBlockInputs(succBlock, v)
			}
		}

	} else if block == block.Function.CfgGraph.EndNode {
		for _, v := range functionFSM.BlockInfo[block].InputTransitions {
			v.IsFinalTransition = true
		}
	}

	for _, succBlock := range block.Successors {
		if functionFSM.CheckBlockVisit(block, succBlock) == true {
			//fmt.Printf("Block %s SuccBlock %s\n", block.Id, succBlock.Id)
			newInputTransitions := PdgGraph.DeepCopyTransitions(currentInputTransitions)

			newFlow := controlFlow.DeepCopyFlow()
			newFlow.AppendToFlow(succBlock)
			ProcessBlockContextSensitive(succBlock, newInputTransitions, callStack, functionFSM, newFlow)
		}
	}

	return currentInputTransitions

}
