package PdgGraph

import "fmt"

const (
	FunctionCallExp = iota
	Literal
	ChildrenExp
	VarExp
	RetExp
	FlowSensitiveExp
	ConditionalExp
)

type Flow struct {
	FlowOrder []*Block
}

func NewFlow() *Flow {
	return &Flow{}
}

func (f *Flow) DeepCopyFlow() *Flow {
	newFlow := NewFlow()
	for _, b := range f.FlowOrder {
		newFlow.FlowOrder = append(newFlow.FlowOrder, b)
	}
	return newFlow
}

func (f *Flow) AppendToFlow(b *Block) {
	f.FlowOrder = append(f.FlowOrder, b)
}

func (f *Flow) PrintFlow() string {
	str := "[ "
	for _, b := range f.FlowOrder {
		str += " " + b.Id + " "
	}
	str += " ] "
	fmt.Printf("[ ")
	for _, b := range f.FlowOrder {
		fmt.Printf(" %s  ", b.Id)
	}
	fmt.Printf(" ]\n")
	return str
}

type ConditionalExpression struct {
	Condition *Flow
	Value     *Expression
}

type ExpressionStack struct {
	StackMap map[*Expression]int
}

func (expStack *ExpressionStack) deepCopyStack() *ExpressionStack {
	newExpStack := NewExpressionStack()
	for k, v := range expStack.StackMap {
		newExpStack.StackMap[k] = v
	}
	return newExpStack
}

func NewExpressionStack() *ExpressionStack {
	return &ExpressionStack{
		StackMap: make(map[*Expression]int),
	}
}

func (expStack *ExpressionStack) SafeToCalculate(exp *Expression) bool {
	occurrences, ok := expStack.StackMap[exp]
	if ok {
		if occurrences < 2 {
			expStack.StackMap[exp] = occurrences + 1
			return true
		} else {
			return false
		}
	} else {
		occurrences = 0
		expStack.StackMap[exp] = occurrences + 1
		return true
	}
}

func (expStack *ExpressionStack) RemoveElementFromStack(exp *Expression) {
	occurrences, ok := expStack.StackMap[exp]
	if ok {
		expStack.StackMap[exp] = occurrences - 1
	}
}

type Expression struct {
	ExpressionType          int
	LiteralExpression       string
	FlowSensitiveExpression map[*Block]*Expression
	Children                []*Expression
	PointsToVar             *Variable
}

func NewExpression(expression string, expressionType int) *Expression {
	return &Expression{
		LiteralExpression:       expression,
		ExpressionType:          expressionType,
		FlowSensitiveExpression: make(map[*Block]*Expression),
		PointsToVar:             nil,
	}
}

func ResolveFlow(e *Expression, f *Flow) (*Expression, *Flow) {
	if f != nil {
		i := len(f.FlowOrder) - 1
		for i = len(f.FlowOrder) - 1; i > 0; i-- {
			exp, ok := e.FlowSensitiveExpression[f.FlowOrder[i]]
			if ok {
				return exp, &Flow{FlowOrder: f.FlowOrder[:i]}
			}
		}
		if i == 0 {
			var dummyBlock []*Block
			return e, &Flow{FlowOrder: dummyBlock}
		}
	}

	return e, nil

}

func (e *Expression) GetCallerExpression(c *CallingContext) *Expression {

	if e.ExpressionType == Literal || e.ExpressionType == FunctionCallExp {
		return e
	}

	NewExp := NewExpression("", e.ExpressionType)
	v, ok := c.ParamToArgs[e.LiteralExpression]
	if ok && e.ExpressionType == VarExp {
		return v.GetExpression()
	}

	if len(e.FlowSensitiveExpression) != 0 {
		for predBlock, exp := range e.FlowSensitiveExpression {
			NewExp.FlowSensitiveExpression[predBlock] = exp.GetCallerExpression(c)
		}
	} else if len(e.Children) != 0 {
		for _, child := range e.Children {
			NewExp.Children = append(NewExp.Children, child.GetCallerExpression(c))
		}
	} else {
		NewExp = e
	}
	return NewExp
}

func (e *Expression) GetStringStack(f *Flow, expStack *ExpressionStack) string {
	if len(e.FlowSensitiveExpression) != 0 {
		if f != nil {
			if len(f.FlowOrder) == 0 {
				exp, newFlow := ResolveFlow(e, f)
				if expStack.SafeToCalculate(exp) {
					expString := exp.GetStringStack(newFlow, expStack)
					expStack.RemoveElementFromStack(exp)
					return expString
				} else {
					return exp.LiteralExpression
				}
			} else {
				return e.LiteralExpression
			}
		} else if len(e.FlowSensitiveExpression) == 1 {
			for _, v := range e.FlowSensitiveExpression {
				if expStack.SafeToCalculate(v) {
					expString := v.GetStringStack(f, expStack)
					expStack.RemoveElementFromStack(v)
					return expString
				}
			}
			return e.LiteralExpression
		} else {
			return e.LiteralExpression
		}

	} else if len(e.Children) != 0 {
		fullString := ""
		for _, child := range e.Children {
			if expStack.SafeToCalculate(child) {
				fullString += child.GetStringStack(f, expStack)
				expStack.RemoveElementFromStack(child)
			} else {
				fullString += child.LiteralExpression
			}

		}
		if fullString != "" {
			return fullString
		} else {
			return e.LiteralExpression
		}
	}
	return e.LiteralExpression
}

func (e *Expression) GetString(f *Flow) string {
	expStack := NewExpressionStack()
	expressionString := e.GetStringStack(f, expStack)
	return expressionString
}

func (e *Expression) GetCallerString(f *Flow, c *CallingContext) string {
	expStack := NewExpressionStack()
	callerExp := e.GetCallerExpression(c)
	expString := callerExp.GetStringStack(f, expStack)
	return expString
}

func (e *Expression) GetFlowExpression(f *Flow) *Expression {
	expStack := NewExpressionStack()
	expFlow := e.GetFlowExpressionStack(f, expStack)
	return expFlow
}

func (e *Expression) GetFlowExpressionStack(f *Flow, expStack *ExpressionStack) *Expression {
	if len(e.FlowSensitiveExpression) != 0 {
		if f != nil {
			if len(f.FlowOrder) != 0 {
				exp, newFlow := ResolveFlow(e, f)
				if expStack.SafeToCalculate(exp) {
					expFlow := exp.GetFlowExpressionStack(newFlow, expStack)
					expStack.RemoveElementFromStack(exp)
					return expFlow
				} else {
					return exp
				}

			} else {
				return e
			}
		} else {
			return e
		}
	} else if len(e.Children) != 0 {
		newExp := NewExpression("", ChildrenExp)
		for _, child := range e.Children {
			if expStack.SafeToCalculate(child) {
				childFlowExp := child.GetFlowExpressionStack(f, expStack)
				newExp.Children = append(newExp.Children, childFlowExp)
				expStack.RemoveElementFromStack(childFlowExp)
			} else {
				newExp.Children = append(newExp.Children, child)
			}
		}
		return newExp
	}
	return e
}

func SetVariableExpression(variable *Variable, expression *Expression) {
	if len(expression.Children) == 0 && len(expression.FlowSensitiveExpression) == 0 &&
		expression.LiteralExpression == "" {
		return
	}
	variable.ValueExpression.Children = variable.ValueExpression.Children[:0]
	variable.ValueExpression.Children = append(variable.ValueExpression.Children, expression)
}

func CopyExpression(e *Expression) *Expression {
	n := NewExpression(e.LiteralExpression, e.ExpressionType)
	n.PointsToVar = e.PointsToVar
	for k, v := range e.FlowSensitiveExpression {
		n.FlowSensitiveExpression[k] = v
	}
	for _, v := range e.Children {
		n.Children = append(n.Children, v)
	}
	return n
}

func AddToVariableExpression(variable *Variable, expression *Expression) {
	expression = CopyExpression(expression)
	if len(expression.Children) == 0 && len(expression.FlowSensitiveExpression) == 0 &&
		expression.LiteralExpression == "" {
		return
	}
	if len(variable.ValueExpression.FlowSensitiveExpression) == 0 && len(variable.ValueExpression.Children) == 0 {
		variable.ValueExpression.Children = append(variable.ValueExpression.Children, expression)
	} else if len(variable.ValueExpression.Children) != 0 {

		var newChildren []*Expression
		newChildren = append(newChildren, NewExpression(" ( ", Literal))
		for _, child := range variable.ValueExpression.Children {
			newChildren = append(newChildren, child)
		}
		newChildren = append(newChildren, NewExpression(" ) ", Literal))
		newChildren = append(newChildren, NewExpression(" || ", Literal))
		newChildren = append(newChildren, NewExpression(" ( ", Literal))
		newChildren = append(newChildren, expression)
		newChildren = append(newChildren, NewExpression(" ) ", Literal))

		variable.ValueExpression.Children = variable.ValueExpression.Children[:0]
		for _, child := range newChildren {
			variable.ValueExpression.Children = append(variable.ValueExpression.Children, child)
		}

	} else if len(variable.ValueExpression.FlowSensitiveExpression) != 0 {

		FlowExpressionChild := NewExpression(variable.ValueExpression.LiteralExpression, FlowSensitiveExp)
		for k, v := range variable.ValueExpression.FlowSensitiveExpression {
			FlowExpressionChild.FlowSensitiveExpression[k] = v
		}
		variable.ValueExpression.FlowSensitiveExpression = make(map[*Block]*Expression)

		var newChildren []*Expression
		newChildren = append(newChildren, NewExpression(" ( ", Literal))
		newChildren = append(newChildren, FlowExpressionChild)

		newChildren = append(newChildren, NewExpression(" ) ", Literal))

		newChildren = append(newChildren, NewExpression(" || ", Literal))
		newChildren = append(newChildren, NewExpression(" ( ", Literal))
		newChildren = append(newChildren, expression)
		newChildren = append(newChildren, NewExpression(" ) ", Literal))

		variable.ValueExpression.Children = variable.ValueExpression.Children[:0]
		for _, child := range newChildren {
			variable.ValueExpression.Children = append(variable.ValueExpression.Children, child)
		}

	}
}

func AddChildrenToExpression(parent *Expression, childExp *Expression) {
	var newChildren []*Expression
	if len(parent.Children) != 0 {
		newChildren = append(newChildren, NewExpression(" ( ", Literal))
		for _, OriginalChild := range parent.Children {
			newChildren = append(newChildren, OriginalChild)
		}
		newChildren = append(newChildren, NewExpression(" ) ", Literal))
		newChildren = append(newChildren, NewExpression(" || ", Literal))
		newChildren = append(newChildren, NewExpression(" ( ", Literal))
		newChildren = append(newChildren, childExp)
		newChildren = append(newChildren, NewExpression(" ) ", Literal))
		parent.Children = newChildren
	} else {
		if !(len(childExp.Children) == 0 && len(childExp.FlowSensitiveExpression) == 0 &&
			childExp.LiteralExpression == "") {
			parent.Children = append(parent.Children, childExp)
		}

	}
}

func AppendCondition(firstCond *Expression, secondCond *Expression) *Expression {

	combinedExp := NewExpression("", ChildrenExp)
	combinedExp.Children = append(combinedExp.Children, firstCond)
	combinedExp.Children = append(combinedExp.Children, NewExpression(" && ", Literal))
	combinedExp.Children = append(combinedExp.Children, secondCond)
	return combinedExp
}

func (f *Flow) CalculateFlowCondition() *Expression {

	finalConditionExp := NewExpression(" true ", ConditionalExp)
	for i, block := range f.FlowOrder {
		if i == len(f.FlowOrder)-1 {
			break
		}

		if block.LastInstr != nil {
			if block.LastInstr.SSAType == "If" {
				var condVar *Variable
				for _, v := range block.LastInstr.AssignFrom {
					condVar = v
				}

				blockLastInstr := block.LastInstr
				conditionalExpression := NewExpression("", ChildrenExp)
				succBlock := f.FlowOrder[i+1]

				if block.TrueSuccessor == succBlock {
					if blockLastInstr.Misc == "true" {
						conditionalExpression.Children = append(conditionalExpression.Children,
							NewExpression(" true ", Literal))
					} else if blockLastInstr.Misc == "false" {
						conditionalExpression.Children = append(conditionalExpression.Children,
							NewExpression(" false ", Literal))
					} else {
						conditionalExpression.Children = append(conditionalExpression.Children, condVar.ValueExpression)
					}
				} else if block.FalseSuccessor == succBlock {
					//fmt.Printf("%s is False Successor of %s\n", block.FullName(), predBlock.FullName())
					if blockLastInstr.Misc == "false" {
						conditionalExpression.Children = append(conditionalExpression.Children,
							NewExpression(" true ", Literal))
					} else if blockLastInstr.Misc == "true" {
						conditionalExpression.Children = append(conditionalExpression.Children,
							NewExpression(" false ", Literal))
					} else {
						conditionalExpression.Children = append(conditionalExpression.Children,
							NewExpression(" ! ", Literal))
						conditionalExpression.Children = append(conditionalExpression.Children,
							NewExpression(" ( ", Literal))
						conditionalExpression.Children = append(conditionalExpression.Children, condVar.ValueExpression)
						conditionalExpression.Children = append(conditionalExpression.Children,
							NewExpression(" ) ", Literal))
					}
				}

				finalConditionExp = AppendCondition(finalConditionExp, conditionalExpression)

			}
		}
	}

	return finalConditionExp

}
