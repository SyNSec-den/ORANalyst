package FsmBuilder

import "constraint_collector/PdgGraph"

func CalculateBlockReachingCondition(block *PdgGraph.Block) {
	currblockExp := PdgGraph.NewExpression("", PdgGraph.ChildrenExp)
	for _, predBlock := range block.Predecessors {
		predBlockLastInstr := predBlock.LastInstr

		if predBlock != block.Function.CfgGraph.StartNode {
			if predBlock.LastInstr != nil && predBlockLastInstr.SSAType == "If" {
				var condVar *PdgGraph.Variable
				for _, v := range predBlockLastInstr.AssignFrom {
					condVar = v
				}
				predExpression := PdgGraph.NewExpression("", PdgGraph.ChildrenExp)

				if predBlock.TrueSuccessor == block {
					if predBlockLastInstr.Misc == "true" {
						predExpression.Children = append(predExpression.Children,
							PdgGraph.NewExpression(" true ", PdgGraph.Literal))
					} else if predBlockLastInstr.Misc == "false" {
						predExpression.Children = append(predExpression.Children,
							PdgGraph.NewExpression(" false ", PdgGraph.Literal))
					} else {
						predExpression.Children = append(predExpression.Children, condVar.ValueExpression)
					}
				} else if predBlock.FalseSuccessor == block {
					//fmt.Printf("%s is False Successor of %s\n", block.FullName(), predBlock.FullName())
					if predBlockLastInstr.Misc == "false" {
						predExpression.Children = append(predExpression.Children,
							PdgGraph.NewExpression(" true ", PdgGraph.Literal))
					} else if predBlockLastInstr.Misc == "true" {
						predExpression.Children = append(predExpression.Children,
							PdgGraph.NewExpression(" false ", PdgGraph.Literal))
					} else {
						predExpression.Children = append(predExpression.Children,
							PdgGraph.NewExpression(" ! ", PdgGraph.Literal))
						predExpression.Children = append(predExpression.Children,
							PdgGraph.NewExpression(" ( ", PdgGraph.Literal))
						predExpression.Children = append(predExpression.Children, condVar.ValueExpression)
						predExpression.Children = append(predExpression.Children,
							PdgGraph.NewExpression(" ) ", PdgGraph.Literal))
					}
				}

				if len(predBlock.ReachingConditionExpression.Children) != 0 {
					if len(predExpression.Children) != 0 {
						appendExp := PdgGraph.NewExpression("", PdgGraph.ChildrenExp)
						appendExp.Children = append(appendExp.Children,
							PdgGraph.NewExpression(" ( ", PdgGraph.Literal))
						appendExp.Children = append(appendExp.Children,
							predBlock.ReachingConditionExpression)
						appendExp.Children = append(appendExp.Children,
							PdgGraph.NewExpression(" ) ", PdgGraph.Literal))
						appendExp.Children = append(appendExp.Children,
							PdgGraph.NewExpression(" && ", PdgGraph.Literal))
						appendExp.Children = append(appendExp.Children, predExpression)
						PdgGraph.AddChildrenToExpression(currblockExp, appendExp)
					} else {
						appendExp := PdgGraph.NewExpression("", PdgGraph.ChildrenExp)
						appendExp.Children = append(appendExp.Children, predBlock.ReachingConditionExpression)
						PdgGraph.AddChildrenToExpression(currblockExp, appendExp)
					}
				} else {
					if len(predExpression.Children) != 0 {
						appendExp := PdgGraph.NewExpression("", PdgGraph.ChildrenExp)
						appendExp.Children = append(appendExp.Children, predExpression)
						PdgGraph.AddChildrenToExpression(currblockExp, appendExp)
					}
				}

			} else if predBlock.LastInstr != nil && predBlock.LastInstr.SSAType == "Jump" {

				if len(currblockExp.Children) == 0 {
					if len(predBlock.ReachingConditionExpression.Children) != 0 {
						appendExp := PdgGraph.NewExpression("", PdgGraph.ChildrenExp)
						appendExp.Children = append(appendExp.Children, predBlock.ReachingConditionExpression)
						PdgGraph.AddChildrenToExpression(currblockExp, appendExp)
					}

				} else {
					if len(predBlock.ReachingConditionExpression.Children) != 0 {
						PdgGraph.AddChildrenToExpression(currblockExp, predBlock.ReachingConditionExpression)
					}
				}
			}
		}
	}

	if len(currblockExp.Children) == 0 && len(currblockExp.FlowSensitiveExpression) == 0 && currblockExp.LiteralExpression == "" {
		currblockExp = PdgGraph.NewExpression(" true ", PdgGraph.Literal)
	}
	block.ReachingConditionExpression = currblockExp
}
