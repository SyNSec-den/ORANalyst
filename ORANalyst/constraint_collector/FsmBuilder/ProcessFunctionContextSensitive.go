package FsmBuilder

import "constraint_collector/PdgGraph"

func ProcessFunctionContextSensitive(fn *PdgGraph.Function,
	callStack *PdgGraph.CallStack) *PdgGraph.FSM {

	callContext := callStack.Top()
	contextSVSignature := PdgGraph.ConvertSVMapToSignature(callContext.StateVariablesInCallee)
	functionFSM, ok := fn.ContextBasedFSMs[contextSVSignature]
	if ok {
		return functionFSM
	}

	functionFSM = PdgGraph.NewFSM()

	functionFSM.ParentFunction = fn
	functionFSM.ContextSignature = contextSVSignature
	fn.ContextBasedFSMs[contextSVSignature] = functionFSM

	functionFSM.AddToBlockInputs(fn.CfgGraph.StartNode, functionFSM.InitTransition)

	newFlow := PdgGraph.NewFlow()
	newFlow.AppendToFlow(fn.CfgGraph.StartNode)
	newInputTransitions := functionFSM.ObtainTxWithNewStates(fn.CfgGraph.StartNode)

	ProcessBlockContextSensitive(fn.CfgGraph.StartNode, newInputTransitions, callStack, functionFSM, newFlow)

	callStack.Pop()

	return functionFSM
}
