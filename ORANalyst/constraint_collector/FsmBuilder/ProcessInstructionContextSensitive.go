package FsmBuilder

import (
	"constraint_collector/PdgGraph"
	"fmt"
	"strings"
)

func ProcessInstructionContextSensitive(instruction *PdgGraph.Instruction,
	inputTransitions map[string]*PdgGraph.FSMTransition,
	callStack *PdgGraph.CallStack,
	functionFSM *PdgGraph.FSM,
	controlFlow *PdgGraph.Flow) map[string]*PdgGraph.FSMTransition {

	currentContext := callStack.Top()
	currentInputTransitions := make(map[string]*PdgGraph.FSMTransition)
	newInputTransitions := make(map[string]*PdgGraph.FSMTransition)

	for k, v := range inputTransitions {
		currentInputTransitions[k] = v
	}
	//fmt.Printf("Control Flow :")
	//controlFlow.PrintFlow()
	fmt.Printf("Instruction : %s\n", instruction.Text)
	fmt.Printf("Input Tarnsitions : ")
	for k := range currentInputTransitions {
		fmt.Printf(" %s ", k)
	}
	fmt.Printf("\n")

	if instruction.SSAType == "Store" && currentContext.ParsedOnce == false {
		currentContext.ParsedOnce = true

		var varAddr *PdgGraph.Variable
		var varVal *PdgGraph.Variable

		for k, v := range instruction.AssignTo {
			if strings.HasPrefix(k, "varAddr") == true {
				varAddr = v
			}
		}

		for k, v := range instruction.AssignFrom {
			if strings.HasPrefix(k, "varVal") == true {
				varVal = v
			}
		}

		ParamPointsTo := false
		if currentContext.CalleeParamVarPointsTo[varAddr.FullName()] != nil && len(currentContext.
			CalleeParamVarPointsTo[varAddr.FullName()]) != 0 {
			ParamPointsTo = true
		}

		if len(varAddr.PointsTo) == 0 && !ParamPointsTo && varVal != nil && varVal.GetExpression() != nil {
			derefVar := PdgGraph.NewVariable()
			derefVar.Name = "*" + varAddr.Name
			derefVar.ContainingFunction = varAddr.ContainingFunction
			derefVar.DefiningInstruction = varAddr.DefiningInstruction
			derefVar.ValueExpression = varVal.GetExpression()
			varAddr.PointsTo[derefVar.FullName()] = derefVar

			if varVal.AliasToStateVariable != nil {
				derefVar.AliasToStateVariable = varVal.AliasToStateVariable
				currentContext.StateVariablesInCallee[derefVar.FullName()] = varVal
			}
			if varAddr.AliasToStateVariable != nil && varVal.AliasToStateVariable == nil {
				derefVar.AliasToStateVariable = derefVar
				currentContext.StateVariablesInCallee[derefVar.FullName()] = derefVar
			}

		} else if len(varAddr.PointsTo) == 0 && !ParamPointsTo {
			derefVar := PdgGraph.NewVariable()
			derefVar.Name = "*" + varAddr.Name
			derefVar.ContainingFunction = varAddr.ContainingFunction
			derefVar.DefiningInstruction = varAddr.DefiningInstruction
			derefVar.ValueExpression = PdgGraph.NewExpression(instruction.Misc, PdgGraph.Literal)
			varAddr.PointsTo[derefVar.FullName()] = derefVar

			if varAddr.AliasToStateVariable != nil {
				derefVar.AliasToStateVariable = derefVar
				currentContext.StateVariablesInCallee[derefVar.FullName()] = derefVar
			}

		} else if (len(varAddr.PointsTo) != 0 || ParamPointsTo) && varVal != nil && varVal.GetExpression() != nil {

			allPointsTo := make(map[string]*PdgGraph.Variable)
			for k, v := range varAddr.PointsTo {
				allPointsTo[k] = v
			}
			if currentContext.CalleeParamVarPointsTo[varAddr.FullName()] != nil {
				for k, v := range currentContext.CalleeParamVarPointsTo[varAddr.FullName()] {
					allPointsTo[k] = v
				}
			}

			for derefVarName, derefVar := range allPointsTo {
				//Issue : Flow insensitive expressions
				PdgGraph.AddToVariableExpression(derefVar, varVal.GetExpression())
				for k, v := range varVal.MayAlias {
					derefVar.MayAlias[k] = v
					v.MayAlias[derefVarName] = derefVar
				}
				for k, v := range varVal.PointsTo {
					derefVar.PointsTo[k] = v
					v.PointedBy[derefVarName] = derefVar
				}
				for k, v := range varVal.PointedBy {
					derefVar.PointedBy[k] = v
					v.PointsTo[derefVarName] = derefVar
				}

				if varVal.AliasToStateVariable != nil {
					derefVar.AliasToStateVariable = varVal.AliasToStateVariable
					currentContext.StateVariablesInCallee[derefVar.FullName()] = varVal
				}
				if varAddr.AliasToStateVariable != nil && varVal.AliasToStateVariable == nil {
					derefVar.AliasToStateVariable = derefVar
					currentContext.StateVariablesInCallee[derefVar.FullName()] = derefVar
				}

			}
		} else if len(varAddr.PointsTo) != 0 || ParamPointsTo {

			allPointsTo := make(map[string]*PdgGraph.Variable)
			for k, v := range varAddr.PointsTo {
				allPointsTo[k] = v
			}
			if currentContext.CalleeParamVarPointsTo[varAddr.FullName()] != nil {
				for k, v := range currentContext.CalleeParamVarPointsTo[varAddr.FullName()] {
					allPointsTo[k] = v
				}
			}

			for _, derefVar := range allPointsTo {
				//Issue : Flow insensitive expressions

				if varAddr.AliasToStateVariable != nil {
					derefVar.AliasToStateVariable = derefVar
					currentContext.StateVariablesInCallee[derefVar.FullName()] = derefVar
				}
				fmt.Printf("DerefVar %s\n", derefVar.FullName())
				PdgGraph.AddToVariableExpression(derefVar, PdgGraph.NewExpression(instruction.Misc, PdgGraph.Literal))

			}
		}
	}

	if instruction.SSAType == "Store" {
		var varR *PdgGraph.Variable
		for _, v := range instruction.AssignTo {
			varR = v
		}
		if varR != nil {
			allPointsTo := make(map[string]*PdgGraph.Variable)
			for k, v := range varR.PointsTo {
				allPointsTo[k] = v
			}
			if currentContext.CalleeParamVarPointsTo[varR.FullName()] != nil {
				for k, v := range currentContext.CalleeParamVarPointsTo[varR.FullName()] {
					allPointsTo[k] = v
				}
			}

			for _, derefVar := range allPointsTo {

				aliasedSV, ok := currentContext.StateVariablesInCallee[derefVar.FullName()]

				if ok {
					//fmt.Printf(" Block %s : %s\n", instruction.Block.FullName(), instruction.Text)
					svName := aliasedSV.FullName()
					svValue := derefVar.GetExpression().GetFlowExpression(controlFlow)
					fmt.Printf(" SV Variable %s Value %s\n", svName, svValue.GetString(nil))

					for k := range newInputTransitions {
						delete(newInputTransitions, k)
					}

					currentFromStates := make(map[string]*PdgGraph.FSMState)
					for _, v := range currentInputTransitions {
						fmt.Printf("Current State %s\n", v.ToState.Name)
						currentFromStates[v.ToState.Name] = v.ToState
					}

					for _, fromState := range currentFromStates {
						fromStateSVValues := make(map[string]*PdgGraph.Expression)
						for k, v := range fromState.SvValueExpressions {
							fromStateSVValues[k] = v
						}

						fromStateSVValues[svName] = svValue

						newState := PdgGraph.NewFSMState(fromStateSVValues)
						//condition := instruction.Block.ReachingConditionExpression.GetFlowExpression(controlFlow)
						//condition := PdgGraph.NewExpression(controlFlow.PrintFlow(), PdgGraph.Literal)
						//newState.ReachingCondition = condition

						newState = functionFSM.AddNewFSMState(newState)
						fmt.Printf("NewState Created %s %s\n", newState.Name, fromState.Name)
						if newState.Name != fromState.Name {
							condition := controlFlow.CalculateFlowCondition().GetFlowExpression(controlFlow)
							newTransition := PdgGraph.NewFSMTransition(fromState, newState, condition)
							newTransition = functionFSM.AddNewFSMTransition(newTransition)
							newTransition.Actions = append(newTransition.Actions, svName+" = "+svValue.GetString(controlFlow))
							newInputTransitions[newTransition.GetName()] = newTransition
							//fmt.Printf(" %s %s %d\n", newTransition.Id, newTransition.Condition.LiteralExpression, len(newTransition.Condition.Children))

						} else {
							for k, v := range currentInputTransitions {
								if v.ToState.Name == fromState.Name {
									newInputTransitions[k] = v
								}
							}
						}
					}

					for k := range currentInputTransitions {
						delete(currentInputTransitions, k)
					}

					for k, newState := range newInputTransitions {
						currentInputTransitions[k] = newState
					}
				}

			}
		}

	}

	if instruction.SSAType == "UnOp" ||
		instruction.SSAType == "BinOp" ||
		instruction.SSAType == "Phi" {

		var varR *PdgGraph.Variable

		for _, v := range instruction.AssignTo {
			varR = v
		}

		if varR != nil {
			aliasedSV, ok := currentContext.StateVariablesInCallee[varR.FullName()]
			if ok {
				//fmt.Printf(" Block %s : %s\n", instruction.Block.FullName(), instruction.Text)
				svName := aliasedSV.FullName()
				svValue := varR.GetExpression().GetFlowExpression(controlFlow)
				//fmt.Printf(" SV Variable %s Value %s\n", svName, svValue.GetString(nil))

				for k := range newInputTransitions {
					delete(newInputTransitions, k)
				}

				currentFromStates := make(map[string]*PdgGraph.FSMState)
				for _, v := range currentInputTransitions {
					currentFromStates[v.ToState.Name] = v.ToState
				}

				for _, fromState := range currentFromStates {
					fromStateSVValues := make(map[string]*PdgGraph.Expression)
					for k, v := range fromState.SvValueExpressions {
						fromStateSVValues[k] = v
					}

					fromStateSVValues[svName] = svValue

					newState := PdgGraph.NewFSMState(fromStateSVValues)
					//condition := instruction.Block.ReachingConditionExpression.GetFlowExpression(controlFlow)
					//condition := PdgGraph.NewExpression(controlFlow.PrintFlow(), PdgGraph.Literal)
					//newState.ReachingCondition = condition

					newState = functionFSM.AddNewFSMState(newState)

					if newState.Name != fromState.Name {
						condition := controlFlow.CalculateFlowCondition().GetFlowExpression(controlFlow)
						newTransition := PdgGraph.NewFSMTransition(fromState, newState, condition)
						newTransition = functionFSM.AddNewFSMTransition(newTransition)
						newTransition.Actions = append(newTransition.Actions, svName+" = "+svValue.GetString(controlFlow))
						newInputTransitions[newTransition.GetName()] = newTransition
						//fmt.Printf(" %s %s %d\n", newTransition.Id, newTransition.Condition.LiteralExpression, len(newTransition.Condition.Children))

					} else {
						for k, v := range currentInputTransitions {
							if v.ToState.Name == fromState.Name {
								newInputTransitions[k] = v
							}
						}
					}
				}

				for k := range currentInputTransitions {
					delete(currentInputTransitions, k)
				}

				for k, newState := range newInputTransitions {
					currentInputTransitions[k] = newState
				}
			}
		}
	}

	if instruction.SSAType == "Call" {

		//fmt.Printf("Instruction %s\n", instruction.Text)
		newContext, ok := PdgGraph.CreateContextFromCallerInstruction(instruction, controlFlow)
		if ok {
			//for k, v := range newContext.StateVariablesInCallee {
			//	fmt.Printf("SV %s : %s\n", k, v.GetExpression().GetString(nil))
			//}

			if callStack.Push(newContext) == true {

				calleeFunctionFSM := ProcessFunctionContextSensitive(newContext.Callee, callStack)

				var varR *PdgGraph.Variable
				for k, v := range instruction.AssignTo {
					if strings.HasPrefix(k, "varR") == true {
						varR = v
					}
				}
				if varR != nil && newContext.Callee.ReturnValueExpression != nil {
					PdgGraph.SetVariableExpression(varR, newContext.Callee.ReturnValueExpression.
						GetCallerExpression(newContext))
				}

				for k := range newInputTransitions {
					delete(newInputTransitions, k)
				}

				if calleeFunctionFSM.InitTransition.IsFinalTransition == true {
					for k, v := range currentInputTransitions {
						newInputTransitions[k] = v
					}
				}

				currentFromStates := make(map[string]*PdgGraph.FSMState)
				for _, v := range currentInputTransitions {
					currentFromStates[v.ToState.Name] = v.ToState
				}

				for _, currFromState := range currentFromStates {
					outTransitions := functionFSM.AddCalleeFuncFSM(
						calleeFunctionFSM, currFromState, newContext, controlFlow)
					fmt.Printf("%d\n", len(outTransitions))
					for k, v := range outTransitions {
						newInputTransitions[k] = v
					}
				}

				for k := range currentInputTransitions {
					delete(currentInputTransitions, k)
				}

				for k, newState := range newInputTransitions {
					currentInputTransitions[k] = newState
				}
			}
		} else {
			var callee *PdgGraph.Function
			for _, v := range instruction.InstrCalls {
				callee = v
			}
			if callee != nil && callee.IsAction == true {
				newAction := callee.FullName()

				currentToStates := make(map[string]*PdgGraph.FSMState)
				for _, v := range currentInputTransitions {
					currentToStates[v.ToState.Name] = v.ToState
				}

				for _, currState := range currentToStates {
					//controlFlow.PrintFlow()
					//fmt.Printf("--> %s\n", controlFlow.CalculateFlowCondition().GetString(nil))
					newCondition := controlFlow.CalculateFlowCondition().GetFlowExpression(controlFlow)
					newTx := PdgGraph.NewFSMTransition(currState, currState, newCondition)
					newTx.Actions = append(newTx.Actions, newAction)
					functionFSM.AddNewFSMTransition(newTx)
				}
			}
		}
	}

	//fmt.Printf("Input ")
	//for k := range inputTransitions {
	//	fmt.Printf("%s ", k)
	//}
	//fmt.Printf("\n")
	//
	//fmt.Printf("Output ")
	//for k := range currentInputTransitions {
	//	fmt.Printf("%s ", k)
	//}
	//fmt.Printf("\n")

	fmt.Printf("Output Transitions : ")
	for k := range currentInputTransitions {
		fmt.Printf(" %s ", k)
	}
	fmt.Printf("\n\n")
	//if len(currentInputTransitions) == 0 {
	//	fmt.Printf("Error\n")
	//	os.Exit(1)
	//}

	return currentInputTransitions
}
