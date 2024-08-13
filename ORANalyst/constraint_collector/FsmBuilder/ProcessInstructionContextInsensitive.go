package FsmBuilder

import (
	"constraint_collector/PdgGraph"
	"strings"
)

func ProcessInstructionContextInsensitive(instruction *PdgGraph.Instruction) {
	switch instruction.SSAType {
	case "Call":
		var varR *PdgGraph.Variable

		for k, v := range instruction.AssignTo {
			if strings.HasPrefix(k, "varR") == true {
				varR = v
			}
		}

		var callee *PdgGraph.Function
		for _, v := range instruction.InstrCalls {
			callee = v
		}

		whiteListed := false
		if callee != nil {
			if callee.WhiteListed == true {
				whiteListed = true
			}
		}

		if whiteListed == false {
			functionCallString := ""

			for _, v := range instruction.InstrCalls {
				functionCallString += v.FullName()
			}

			functionCallString += "("
			for i, v := range instruction.CallerArguments {
				functionCallString += v.GetExpression().GetString(nil)
				if i != len(instruction.CallerArguments)-1 {
					functionCallString += " , "
				}
			}
			functionCallString += ")"

			newExp := PdgGraph.NewExpression(functionCallString, PdgGraph.FunctionCallExp)

			PdgGraph.AddToVariableExpression(varR, newExp)
		}

	case "Return":
		//TODO : Multiple Returns
		if len(instruction.ReturnValues) == 1 {
			var varRes *PdgGraph.Variable
			for k, v := range instruction.AssignFrom {
				if strings.HasPrefix(k, "varRes") == true {
					varRes = v
				}
			}
			retExp := instruction.Block.Function.ReturnValueExpression
			currBlock := instruction.Block
			retExp.FlowSensitiveExpression[currBlock] = varRes.GetExpression()
		}
	case "Alloc":
	case "FieldAddr":
		var varR *PdgGraph.Variable
		var varX *PdgGraph.Variable
		for k, v := range instruction.AssignTo {
			if strings.HasPrefix(k, "varR") == true {
				varR = v
			}
		}
		for k, v := range instruction.AssignFrom {
			if strings.HasPrefix(k, "varX") == true {
				varX = v
			}
		}
		fieldVarName := varX.FullName() + ".field" + instruction.Misc
		fieldVar, ok := varX.ContainingFunction.LocalVariables[fieldVarName]
		if !ok {
			fieldVar = PdgGraph.NewVariable()
			fieldVar.Name = varX.Name + ".field" + instruction.Misc
			fieldVar.IsGlobal = varX.IsGlobal
			fieldVar.DefiningInstruction = varX.DefiningInstruction
			fieldVar.ValueExpression = PdgGraph.NewExpression(fieldVarName, PdgGraph.VarExp)
			fieldVar.ValueExpression.PointsToVar = fieldVar
			fieldPart := ".field" + instruction.Misc

			if varX.AliasToStateVariable == varX {
				fieldVar.AliasToStateVariable = fieldVar
			} else if varX.AliasToStateVariable != nil {
				originalSV, okElement :=
					varX.AliasToStateVariable.ArrayElements[varX.AliasToStateVariable.FullName()+fieldPart]
				if okElement {
					fieldVar.AliasToStateVariable = originalSV
				}
			}

			fieldVar.ContainingFunction = varX.ContainingFunction
			varX.ContainingFunction.LocalVariables[fieldVarName] = fieldVar

		}

		varX.IsStructVariable = true
		varX.FieldElements[fieldVarName] = fieldVar

		varR.PointsTo[fieldVarName] = fieldVar
		fieldVar.PointedBy[varR.FullName()] = varR

		var keysToRemove []string
		for _, aliasVar := range varR.MayAlias {
			_, fieldVarInAlias := aliasVar.PointsTo[fieldVarName]
			if len(aliasVar.PointsTo) != 0 && !fieldVarInAlias {
				keysToRemove = append(keysToRemove, aliasVar.FullName())
			}
		}

		for _, k := range keysToRemove {
			v, ok2 := varR.MayAlias[k]
			if ok2 {
				delete(varR.MayAlias, k)
				_, ok3 := v.MayAlias[varR.FullName()]
				if ok3 {
					delete(v.MayAlias, varR.FullName())
				}
			}
		}

		expTree := PdgGraph.NewExpression("", PdgGraph.ChildrenExp)
		expTree.Children = append(expTree.Children, PdgGraph.NewExpression(" ( ", PdgGraph.Literal))
		expTree.Children = append(expTree.Children, PdgGraph.NewExpression(" & ", PdgGraph.Literal))
		expTree.Children = append(expTree.Children, fieldVar.GetExpression())
		expTree.Children = append(expTree.Children, PdgGraph.NewExpression(" ) ", PdgGraph.Literal))

		PdgGraph.AddToVariableExpression(varR, expTree)
	case "IndexAddr":
		var varR *PdgGraph.Variable
		var varX *PdgGraph.Variable
		var varIdx *PdgGraph.Variable

		for k, v := range instruction.AssignTo {
			if strings.HasPrefix(k, "varR") == true {
				varR = v
			}
		}

		for k, v := range instruction.AssignFrom {
			if strings.HasPrefix(k, "varX") == true {
				varX = v
			}
		}

		for k, v := range instruction.AssignFrom {
			if strings.HasPrefix(k, "varIdx") == true {
				varIdx = v
			}
		}

		elementIdxExpression := ""
		if varIdx != nil {
			elementIdxExpression = varIdx.GetExpression().GetString(nil)
		} else {
			elementIdxExpression = instruction.Misc
		}

		elementVarName := varX.FullName() + "[" + elementIdxExpression + "]"
		elementVar, ok := varX.ContainingFunction.LocalVariables[elementVarName]

		if !ok {
			elementVar = PdgGraph.NewVariable()
			elementVar.Name = varX.Name + "[" + elementIdxExpression + "]"

			elementVar.IsGlobal = varX.IsGlobal
			elementVar.DefiningInstruction = varX.DefiningInstruction
			elementVar.ValueExpression = PdgGraph.NewExpression(elementVarName, PdgGraph.VarExp)
			elementVar.ValueExpression.PointsToVar = elementVar

			elementIdxPart := "[" + elementIdxExpression + "]"

			if varX.AliasToStateVariable == varX {
				elementVar.AliasToStateVariable = elementVar
			} else if varX.AliasToStateVariable != nil {
				originalSV, okElement :=
					varX.AliasToStateVariable.ArrayElements[varX.AliasToStateVariable.FullName()+elementIdxPart]
				if okElement {
					elementVar.AliasToStateVariable = originalSV
				}
			}

			elementVar.ContainingFunction = varX.ContainingFunction
			varX.ContainingFunction.LocalVariables[elementVarName] = elementVar

		}

		varX.IsArrayVariable = true
		varX.ArrayElements[elementVarName] = elementVar

		varR.PointsTo[elementVarName] = elementVar
		elementVar.PointedBy[varR.FullName()] = varR

		var keysToRemove []string
		for _, aliasVar := range varR.MayAlias {
			_, elementVarInAlias := aliasVar.PointsTo[elementVarName]
			if len(aliasVar.PointsTo) != 0 && !elementVarInAlias {
				keysToRemove = append(keysToRemove, aliasVar.FullName())
			}
		}

		for _, k := range keysToRemove {
			v, ok2 := varR.MayAlias[k]
			if ok2 {
				delete(varR.MayAlias, k)
				_, ok3 := v.MayAlias[varR.FullName()]
				if ok3 {
					delete(v.MayAlias, varR.FullName())
				}
			}
		}

		expTree := PdgGraph.NewExpression("", PdgGraph.ChildrenExp)
		expTree.Children = append(expTree.Children, PdgGraph.NewExpression(" ( ", PdgGraph.Literal))
		expTree.Children = append(expTree.Children, PdgGraph.NewExpression(" & ", PdgGraph.Literal))
		expTree.Children = append(expTree.Children, elementVar.GetExpression())
		expTree.Children = append(expTree.Children, PdgGraph.NewExpression(" ) ", PdgGraph.Literal))

		PdgGraph.AddToVariableExpression(varR, expTree)
	case "Store":
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

		addrExpression := PdgGraph.NewExpression(varAddr.FullName(), PdgGraph.VarExp)
		addrExpression.PointsToVar = varAddr

		if len(varAddr.PointsTo) == 0 && varVal != nil && varVal.GetExpression() != nil {
			derefVar := PdgGraph.NewVariable()
			derefVar.Name = "*" + varAddr.Name
			derefVar.ContainingFunction = varAddr.ContainingFunction
			derefVar.DefiningInstruction = varAddr.DefiningInstruction
			derefVar.ValueExpression = varVal.GetExpression()
			varAddr.PointsTo[derefVar.FullName()] = derefVar

			derefVar.AliasToStateVariable = varVal.AliasToStateVariable
			if varAddr.AliasToStateVariable != nil && varVal.AliasToStateVariable == nil {
				derefVar.AliasToStateVariable = derefVar
			}

			addrExpression.Children = append(addrExpression.Children, PdgGraph.NewExpression(" & ", PdgGraph.Literal))
			addrExpression.Children = append(addrExpression.Children, PdgGraph.NewExpression(" ( ", PdgGraph.Literal))
			addrExpression.Children = append(addrExpression.Children, derefVar.GetExpression())
			addrExpression.Children = append(addrExpression.Children, PdgGraph.NewExpression(" ) ", PdgGraph.Literal))

		} else if len(varAddr.PointsTo) == 0 {
			derefVar := PdgGraph.NewVariable()
			derefVar.Name = "*" + varAddr.Name
			derefVar.ContainingFunction = varAddr.ContainingFunction
			derefVar.DefiningInstruction = varAddr.DefiningInstruction
			derefVar.ValueExpression = PdgGraph.NewExpression(instruction.Misc, PdgGraph.Literal)
			varAddr.PointsTo[derefVar.FullName()] = derefVar

			if varAddr.AliasToStateVariable != nil {
				derefVar.AliasToStateVariable = derefVar
			}

			addrExpression.Children = append(addrExpression.Children, PdgGraph.NewExpression(" & ", PdgGraph.Literal))
			addrExpression.Children = append(addrExpression.Children, PdgGraph.NewExpression(" ( ", PdgGraph.Literal))
			addrExpression.Children = append(addrExpression.Children, derefVar.GetExpression())
			addrExpression.Children = append(addrExpression.Children, PdgGraph.NewExpression(" ) ", PdgGraph.Literal))

		} else if len(varAddr.PointsTo) != 0 && varVal != nil && varVal.GetExpression() != nil {
			for derefVarName, derefVar := range varAddr.PointsTo {
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

				derefVar.AliasToStateVariable = varVal.AliasToStateVariable
				if varAddr.AliasToStateVariable != nil && varVal.AliasToStateVariable == nil {
					derefVar.AliasToStateVariable = derefVar
				}

				if len(addrExpression.Children) == 0 {
					addrExpression.Children = append(addrExpression.Children, PdgGraph.NewExpression(" & ", PdgGraph.Literal))
					addrExpression.Children = append(addrExpression.Children, PdgGraph.NewExpression(" ( ", PdgGraph.Literal))
					addrExpression.Children = append(addrExpression.Children, derefVar.GetExpression())
					addrExpression.Children = append(addrExpression.Children, PdgGraph.NewExpression(" ) ", PdgGraph.Literal))

				} else {
					var newChildren []*PdgGraph.Expression
					newChildren = append(newChildren, PdgGraph.NewExpression(" ( ", PdgGraph.Literal))
					for _, child := range addrExpression.Children {
						newChildren = append(newChildren, child)
					}
					newChildren = append(newChildren, PdgGraph.NewExpression(" ) ", PdgGraph.Literal))
					newChildren = append(newChildren, PdgGraph.NewExpression(" || ", PdgGraph.Literal))
					newChildren = append(newChildren, PdgGraph.NewExpression(" ( ", PdgGraph.Literal))

					newChildren = append(newChildren, PdgGraph.NewExpression(" & ", PdgGraph.Literal))
					newChildren = append(newChildren, PdgGraph.NewExpression(" ( ", PdgGraph.Literal))
					newChildren = append(newChildren, derefVar.GetExpression())
					newChildren = append(newChildren, PdgGraph.NewExpression(" ) ", PdgGraph.Literal))

					addrExpression.Children = newChildren
				}
			}
		} else if len(varAddr.PointsTo) != 0 {
			for _, derefVar := range varAddr.PointsTo {

				if varAddr.AliasToStateVariable != nil {
					derefVar.AliasToStateVariable = derefVar
				}

				//Issue : Flow insensitive expressions
				PdgGraph.AddToVariableExpression(derefVar, PdgGraph.NewExpression(instruction.Misc, PdgGraph.Literal))

				if len(addrExpression.Children) == 0 {
					addrExpression.Children = append(addrExpression.Children, PdgGraph.NewExpression(" & ", PdgGraph.Literal))
					addrExpression.Children = append(addrExpression.Children, PdgGraph.NewExpression(" ( ", PdgGraph.Literal))
					addrExpression.Children = append(addrExpression.Children, derefVar.GetExpression())
					addrExpression.Children = append(addrExpression.Children, PdgGraph.NewExpression(" ) ", PdgGraph.Literal))

				} else {
					var newChildren []*PdgGraph.Expression
					newChildren = append(newChildren, PdgGraph.NewExpression(" ( ", PdgGraph.Literal))
					for _, child := range addrExpression.Children {
						newChildren = append(newChildren, child)
					}
					newChildren = append(newChildren, PdgGraph.NewExpression(" ) ", PdgGraph.Literal))
					newChildren = append(newChildren, PdgGraph.NewExpression(" || ", PdgGraph.Literal))
					newChildren = append(newChildren, PdgGraph.NewExpression(" ( ", PdgGraph.Literal))

					newChildren = append(newChildren, PdgGraph.NewExpression(" & ", PdgGraph.Literal))
					newChildren = append(newChildren, PdgGraph.NewExpression(" ( ", PdgGraph.Literal))
					newChildren = append(newChildren, derefVar.GetExpression())
					newChildren = append(newChildren, PdgGraph.NewExpression(" ) ", PdgGraph.Literal))

					addrExpression.Children = newChildren
				}
			}
		}

		varAddr.ValueExpression = addrExpression

	case "MakeInterface":
		var varR *PdgGraph.Variable
		var varX *PdgGraph.Variable

		for k, v := range instruction.AssignTo {
			if strings.HasPrefix(k, "varR") == true {
				varR = v
			}
		}

		for k, v := range instruction.AssignFrom {
			if strings.HasPrefix(k, "varX") == true {
				varX = v
			}
		}

		for k, v := range varX.MayAlias {
			varR.MayAlias[k] = v
			v.MayAlias[varR.FullName()] = varR
		}
		for k, v := range varX.PointsTo {
			varR.PointsTo[k] = v
			v.PointedBy[varR.FullName()] = varR
		}
		for k, v := range varX.PointedBy {
			varR.PointedBy[k] = v
			v.PointsTo[varR.FullName()] = varR
		}

		PdgGraph.AddToVariableExpression(varR, varX.GetExpression())
		varR.AliasToStateVariable = varX.AliasToStateVariable
	case "Phi":
		var varR *PdgGraph.Variable
		for k, v := range instruction.AssignTo {
			if strings.HasPrefix(k, "varR") == true {
				varR = v
			}
		}

		currBlock := instruction.Block
		phiExpression := PdgGraph.NewExpression("", PdgGraph.FlowSensitiveExp)

		for _, predBlock := range currBlock.Predecessors {
			v, ok := instruction.PhiMap[predBlock.Id]
			if ok {
				expVar, ok2 := instruction.AssignFrom["varEdge::"+v]
				if ok2 {
					phiExpression.FlowSensitiveExpression[predBlock] = expVar.GetExpression()
				} else {
					phiExpression.FlowSensitiveExpression[predBlock] = PdgGraph.NewExpression(v, PdgGraph.Literal)
				}
			}
		}

		//fmt.Printf("Calculating %s\n", varR.FullName())
		//fmt.Printf("Calculated Phi Expression %s\n", phiExpression.GetString(nil))
		PdgGraph.AddToVariableExpression(varR, phiExpression)
		//fmt.Printf("%s Expression is %s\n", varR.FullName(), varR.ValueExpression.GetString(nil))
	case "Slice":
		var varR *PdgGraph.Variable
		var varX *PdgGraph.Variable
		var varLow *PdgGraph.Variable
		var varHigh *PdgGraph.Variable

		for k, v := range instruction.AssignTo {
			if strings.HasPrefix(k, "varR") == true {
				varR = v
			}
		}

		for k, v := range instruction.AssignFrom {
			if strings.HasPrefix(k, "varX") == true {
				varX = v
			}

			if strings.HasPrefix(k, "varHigh") == true {
				varHigh = v
			}

			if strings.HasPrefix(k, "varLow") == true {
				varLow = v
			}
		}

		LowIdxExpression := ""
		if varLow != nil {
			LowIdxExpression = varLow.GetExpression().GetString(nil)
		} else if instruction.LowVal != "" {
			LowIdxExpression = instruction.LowVal
		}

		HighIdxExpression := ""
		if varHigh != nil {
			HighIdxExpression = varHigh.GetExpression().GetString(nil)
		} else if instruction.HighVal != "" {
			HighIdxExpression = instruction.HighVal
		}

		sliceVarName := ""
		if LowIdxExpression == "" && HighIdxExpression == "" {
			sliceVarName = varX.FullName()
		} else if LowIdxExpression == "" {
			sliceVarName = varX.FullName() + "[" + LowIdxExpression + ":]"
		} else if HighIdxExpression == "" {
			sliceVarName = varX.FullName() + "[:" + HighIdxExpression + "]"
		} else {
			sliceVarName =
				varX.FullName() + "[" + LowIdxExpression + ":" + HighIdxExpression + "]"
		}

		sliceVar, ok := varX.ContainingFunction.LocalVariables[sliceVarName]

		if !ok {
			sliceVar = PdgGraph.NewVariable()
			if LowIdxExpression == "" && HighIdxExpression == "" {
				sliceVar.Name = varX.Name
			} else if LowIdxExpression == "" {
				sliceVar.Name = varX.Name + "[" + LowIdxExpression + ":]"
			} else if HighIdxExpression == "" {
				sliceVar.Name = varX.Name + "[:" + HighIdxExpression + "]"
			} else {
				sliceVar.Name =
					varX.Name + "[" + LowIdxExpression + ":" + HighIdxExpression + "]"
			}

			sliceVar.IsGlobal = varX.IsGlobal
			sliceVar.DefiningInstruction = instruction

			sliceVar.ValueExpression = PdgGraph.NewExpression(sliceVarName, PdgGraph.ChildrenExp)
			sliceVar.ValueExpression.PointsToVar = sliceVar

			slicePart := ""
			if LowIdxExpression == "" && HighIdxExpression == "" {
				slicePart = ""
			} else if LowIdxExpression == "" {
				slicePart = "[" + LowIdxExpression + ":]"
			} else if HighIdxExpression == "" {
				slicePart = "[:" + HighIdxExpression + "]"
			} else {
				slicePart = "[" + LowIdxExpression + ":" + HighIdxExpression + "]"
			}

			if varX.AliasToStateVariable == varX {
				sliceVar.AliasToStateVariable = sliceVar
			} else if varX.AliasToStateVariable != nil {
				if slicePart == "" {
					sliceVar.AliasToStateVariable = varX.AliasToStateVariable
				} else {
					originalSV, okElement := varX.AliasToStateVariable.ArrayElements[sliceVarName]
					if okElement {
						sliceVar.AliasToStateVariable = originalSV
					}
				}
			}

			sliceVar.ContainingFunction = varX.ContainingFunction
			varX.ContainingFunction.LocalVariables[sliceVarName] = sliceVar

		}

		varX.IsArrayVariable = true
		//varX.ArrayElements[sliceVarName] = sliceVar

		//TODO : Discuss
		expTree := PdgGraph.NewExpression("", PdgGraph.ChildrenExp)
		expTree.Children = append(expTree.Children, PdgGraph.NewExpression(" ( ", PdgGraph.Literal))
		expTree.Children = append(expTree.Children, PdgGraph.NewExpression(" & ", PdgGraph.Literal))
		expTree.Children = append(expTree.Children, sliceVar.GetExpression())
		expTree.Children = append(expTree.Children, PdgGraph.NewExpression(" ) ", PdgGraph.Literal))

		PdgGraph.AddToVariableExpression(varR, expTree)
	case "UnOp":
		var varR *PdgGraph.Variable
		var varX *PdgGraph.Variable

		for k, v := range instruction.AssignTo {
			if strings.HasPrefix(k, "varR") == true {
				varR = v
			}
		}

		for k, v := range instruction.AssignFrom {
			if strings.HasPrefix(k, "varX") == true {
				varX = v
			}
		}

		if strings.TrimSpace(instruction.Operator) == "*" {
			for _, derefVar := range varX.PointsTo {
				PdgGraph.AddToVariableExpression(varR, derefVar.GetExpression())
			}
		} else {
			unOpExpression := PdgGraph.NewExpression("", PdgGraph.ChildrenExp)
			unOpExpression.Children = append(unOpExpression.Children, PdgGraph.NewExpression(instruction.Operator, PdgGraph.Literal))
			unOpExpression.Children = append(unOpExpression.Children, PdgGraph.NewExpression(" ", PdgGraph.Literal))
			unOpExpression.Children = append(unOpExpression.Children, varX.GetExpression())
			PdgGraph.AddToVariableExpression(varR, unOpExpression)
		}
	case "BinOp":
		var varR *PdgGraph.Variable
		var varX *PdgGraph.Variable
		var varY *PdgGraph.Variable

		for k, v := range instruction.AssignTo {
			if strings.HasPrefix(k, "varR") == true {
				varR = v
			}
		}

		for k, v := range instruction.AssignFrom {
			if strings.HasPrefix(k, "varX") == true {
				varX = v
			}
		}

		for k, v := range instruction.AssignFrom {
			if strings.HasPrefix(k, "varY") == true {
				varY = v
			}
		}

		BinOpExpression := PdgGraph.NewExpression("", PdgGraph.ChildrenExp)
		if varX != nil && varY != nil {
			BinOpExpression.Children = append(BinOpExpression.Children, varX.GetExpression())
			BinOpExpression.Children = append(BinOpExpression.Children, PdgGraph.NewExpression(" ", PdgGraph.Literal))
			BinOpExpression.Children = append(BinOpExpression.Children, PdgGraph.NewExpression(instruction.Operator, PdgGraph.Literal))
			BinOpExpression.Children = append(BinOpExpression.Children, PdgGraph.NewExpression(" ", PdgGraph.Literal))
			BinOpExpression.Children = append(BinOpExpression.Children, varY.GetExpression())
		} else if varX != nil && varY == nil {
			BinOpExpression.Children = append(BinOpExpression.Children, varX.GetExpression())
			BinOpExpression.Children = append(BinOpExpression.Children, PdgGraph.NewExpression(" ", PdgGraph.Literal))
			BinOpExpression.Children = append(BinOpExpression.Children, PdgGraph.NewExpression(instruction.Operator, PdgGraph.Literal))
			BinOpExpression.Children = append(BinOpExpression.Children, PdgGraph.NewExpression(" ", PdgGraph.Literal))
			BinOpExpression.Children = append(BinOpExpression.Children, PdgGraph.NewExpression(instruction.OperandY, PdgGraph.Literal))
		} else if varX == nil && varY != nil {
			BinOpExpression.Children = append(BinOpExpression.Children, PdgGraph.NewExpression(instruction.OperandX, PdgGraph.Literal))
			BinOpExpression.Children = append(BinOpExpression.Children, PdgGraph.NewExpression(" ", PdgGraph.Literal))
			BinOpExpression.Children = append(BinOpExpression.Children, PdgGraph.NewExpression(instruction.Operator, PdgGraph.Literal))
			BinOpExpression.Children = append(BinOpExpression.Children, PdgGraph.NewExpression(" ", PdgGraph.Literal))
			BinOpExpression.Children = append(BinOpExpression.Children, varY.GetExpression())
		} else {
			BinOpExpression.Children = append(BinOpExpression.Children, PdgGraph.NewExpression(instruction.OperandX, PdgGraph.Literal))
			BinOpExpression.Children = append(BinOpExpression.Children, PdgGraph.NewExpression(" ", PdgGraph.Literal))
			BinOpExpression.Children = append(BinOpExpression.Children, PdgGraph.NewExpression(instruction.Operator, PdgGraph.Literal))
			BinOpExpression.Children = append(BinOpExpression.Children, PdgGraph.NewExpression(" ", PdgGraph.Literal))
			BinOpExpression.Children = append(BinOpExpression.Children, PdgGraph.NewExpression(instruction.OperandY, PdgGraph.Literal))
		}
		PdgGraph.AddToVariableExpression(varR, BinOpExpression)
		//		varR.LocalExpressionExpansion = varX.GetExpression() + " " + instruction.Operator + " " + varY.Name
	}
}
