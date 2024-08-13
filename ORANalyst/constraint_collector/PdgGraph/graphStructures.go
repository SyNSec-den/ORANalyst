package PdgGraph

import (
	"fmt"
	"sort"
	"strconv"

	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa"
)

type Vertex interface {
	Label() string
	Properties() map[string]interface{}
}

type Package struct {
	SourceURL string
	Version   string
	Path      string
	PkgID     string

	ContainsFunction map[string]*Function
	Program          *ProgramInfo
	PackagePtr       *packages.Package
}

func NewPackage() *Package {
	return &Package{
		ContainsFunction: make(map[string]*Function),
	}
}

func (_ *Package) Label() string {
	return "package"
}

func (p *Package) Properties() map[string]interface{} {
	return map[string]interface{}{
		"SourceURL": p.SourceURL,
		"Version":   p.Version,
		"Path":      p.Path,
		"PkgID":     p.PkgID,
	}
}

func (p *Package) FullName() string {
	return p.PkgID
}

type ProgramInfo struct {
	ContainsPkg         map[string]*Package
	WhiteListedFunction map[string]*Function
}

func NewProgramInfo() *ProgramInfo {
	return &ProgramInfo{
		ContainsPkg:         make(map[string]*Package),
		WhiteListedFunction: make(map[string]*Function),
	}
}

type Function struct {
	Id   string
	Key  string
	Name string
	File string
	Pos  int

	WhiteListed bool
	IsAction    bool

	Package           *Package
	FirstStatement    *Instruction
	Statements        []*Instruction
	CallsFunctionCall map[string]*CallingContext

	ContainsBlock map[string]*Block
	FirstBlock    *Block
	Parameter     map[int]*Parameter
	Calls         map[string]*Function
	Caller        map[string]*Function

	LocalVariables map[string]*Variable
	StateVariables map[string]*Variable
	ParamIsVar     map[*Parameter]*Variable //Parameter to Function Local Variable Mapping

	PosMap map[int][]*Variable

	SsaFuncPointer *ssa.Function
	CfgGraph       *CFGGraph

	ReturnValueExpression *Expression

	//TODO: Context Sensitive mapping to FSM
	ContextBasedFSMs map[string]*FSM
}

func NewFunction() *Function {
	return &Function{
		WhiteListed:           false,
		IsAction:              false,
		Parameter:             make(map[int]*Parameter),
		CallsFunctionCall:     make(map[string]*CallingContext),
		ContainsBlock:         make(map[string]*Block),
		Calls:                 make(map[string]*Function),
		Caller:                make(map[string]*Function),
		PosMap:                make(map[int][]*Variable),
		StateVariables:        make(map[string]*Variable),
		LocalVariables:        make(map[string]*Variable),
		ParamIsVar:            make(map[*Parameter]*Variable),
		ContextBasedFSMs:      make(map[string]*FSM),
		ReturnValueExpression: NewExpression("", RetExp),
	}
}

func (_ *Function) Label() string {
	return "function"
}

func (f *Function) Properties() map[string]interface{} {
	return map[string]interface{}{
		"Name": f.Name,
		"File": f.File,
		"Pos":  f.Pos,
	}
}

func (f *Function) FullName() string {
	return f.Package.FullName() + ":" + f.Name
}

type Block struct {
	Id   string
	Key  string
	Name string
	File string
	Pos  int

	PredecessorIds []string
	SuccessorIds   []string

	Function      *Function
	FirstInstr    *Instruction
	LastInstr     *Instruction
	ContainsInstr map[string]*Instruction
	BlockPrev     *Block
	BlockNext     *Block

	ControlDepBlocks   map[*Block]string
	ControllerOfBlocks []*Block
	BFSVisitController bool

	PhiInstr *Instruction ``

	Predecessors []*Block
	Successors   []*Block

	UnconditionalSuccessor *Block
	TrueSuccessor          *Block
	FalseSuccessor         *Block

	ReachingConditionExpression *Expression
}

func NewBlock() *Block {
	return &Block{
		ContainsInstr:               make(map[string]*Instruction),
		ControlDepBlocks:            make(map[*Block]string),
		ReachingConditionExpression: NewExpression("", ConditionalExp),
	}
}

func (b *Block) FullName() string {
	return b.Function.FullName() + ":" + b.Id
}

func (_ *Block) Label() string {
	return "block"
}

func (b *Block) Properties() map[string]interface{} {
	return map[string]interface{}{
		"File": b.File,
		"Pos":  b.Pos,
		"Name": b.Name,
	}
}

type CFGGraph struct {
	StartNode *Block
	EndNode   *Block
	Function  *Function
	Nodes     map[string]*Block
}

func (cfgGraph *CFGGraph) getNodeById(id int) *Block {
	idxStr := strconv.Itoa(id)
	node, ok := cfgGraph.Nodes[idxStr]
	if !ok {
		return nil
	}
	return node
}

func NewCFGGraph() *CFGGraph {
	return &CFGGraph{
		Nodes: make(map[string]*Block),
	}
}

func (graph *CFGGraph) getNodeWithIdx(idx int) *Block {
	strIdx := strconv.Itoa(idx)
	node, _ := graph.Nodes[strIdx]
	return node
}

func (graph *CFGGraph) ParseCfg() {

	for nodeName := range graph.Nodes {
		node := graph.Nodes[nodeName]
		fmt.Printf("Block Name : %s\n", node.Id)
		fmt.Printf("PredecessorIds: %v\n", node.PredecessorIds)
		fmt.Printf("SuccessorIds: %v\n\n", node.SuccessorIds)
	}

}

type Instruction struct {

	//condition expansion expression

	Key        string
	Id         string
	File       string
	Offset     int
	Text       string
	SSAType    string
	ASTType    string
	Pos        int
	Misc       string
	Next       []*Instruction
	References []*Variable
	Assigns    []*Variable

	Block       *Block
	SsaInstrPtr *ssa.Instruction
	InstrType   string

	InstrNext *Instruction
	InstrPrev *Instruction

	FunctionCallResult *Variable
	CalledFnArguments  []*Argument

	AssignTo      map[string]*Variable
	InstrCalls    map[string]*Function
	AssignFrom    map[string]*Variable
	CallParameter map[string]*Parameter

	CallerArguments map[int]*Argument
	ReturnValues    map[int]*Argument

	PhiMap      map[string]string
	ChannelRecv *Variable
	ChannelSend *Variable

	UsesDefinition map[string]*Instruction

	TrueJump   string
	FalseJump  string
	UnCondJump string

	LowVal  string
	HighVal string

	Operator string

	OperandX string
	OperandY string
}

func NewInstruction() *Instruction {
	return &Instruction{
		InstrCalls:      make(map[string]*Function),
		AssignTo:        make(map[string]*Variable),
		AssignFrom:      make(map[string]*Variable),
		UsesDefinition:  make(map[string]*Instruction),
		CallParameter:   make(map[string]*Parameter),
		PhiMap:          make(map[string]string),
		CallerArguments: make(map[int]*Argument),
		ReturnValues:    make(map[int]*Argument),

		LowVal:  "",
		HighVal: "",

		TrueJump:   "",
		FalseJump:  "",
		UnCondJump: "",
	}
}

func (_ *Instruction) Label() string {
	return "instruction"
}

func (s *Instruction) Properties() map[string]interface{} {
	return map[string]interface{}{
		"File":    s.File,
		"Offset":  s.Offset,
		"Text":    s.Text,
		"ASTType": s.ASTType,
		"SSAType": s.SSAType,
		"Pos":     s.Pos,
		"Misc":    s.Misc,
	}
}

type Variable struct {
	Id       string
	Key      string
	Name     string
	Type     string
	IsGlobal bool

	//pointer level
	//possibleSamePointingVar

	PossibleSameChan map[string]*Variable

	DefiningInstruction *Instruction
	ContainingFunction  *Function
	SSAValue            *ssa.Value

	ValueExpression *Expression

	IsArrayVariable  bool
	IsStructVariable bool

	MayAlias  map[string]*Variable
	PointsTo  map[string]*Variable
	PointedBy map[string]*Variable

	ArrayElements map[string]*Variable // Key is Index
	FieldElements map[string]*Variable // Key is Field

	ParentVariable       *Variable //If element of Array or Struct
	AliasToStateVariable *Variable
}

func NewVariable() *Variable {
	return &Variable{
		PossibleSameChan: make(map[string]*Variable),
		MayAlias:         make(map[string]*Variable),
		ArrayElements:    make(map[string]*Variable),
		FieldElements:    make(map[string]*Variable),
		PointsTo:         make(map[string]*Variable),
		PointedBy:        make(map[string]*Variable),
		IsArrayVariable:  false,
		IsStructVariable: false,
	}
}

func GetAllSVFromVariable(currVariable *Variable) map[string]*Variable {
	allSV := make(map[string]*Variable)
	if currVariable.AliasToStateVariable != nil {
		allSV[currVariable.FullName()] = currVariable.AliasToStateVariable
	} else {
		if len(currVariable.ArrayElements) != 0 {
			for _, ElemVar := range currVariable.ArrayElements {
				elemSVs := GetAllSVFromVariable(ElemVar)
				for elemSVName, elemSV := range elemSVs {
					allSV[elemSVName] = elemSV
				}
			}
		}
		if len(currVariable.FieldElements) != 0 {
			for _, ElemVar := range currVariable.FieldElements {
				elemSVs := GetAllSVFromVariable(ElemVar)
				for elemSVName, elemSV := range elemSVs {
					allSV[elemSVName] = elemSV
				}
			}
		}
	}
	return allSV
}

func CheckForStateVariable(varArg *Variable, varParam *Variable) map[string]*Variable {
	allSV := make(map[string]*Variable)
	if varArg.AliasToStateVariable != nil {
		allSV[varParam.FullName()] = varArg.AliasToStateVariable
		varParam.AliasToStateVariable = varArg.AliasToStateVariable
	} else {

		if len(varArg.ArrayElements) != 0 {
			for argElemIdx, argElemVar := range varArg.ArrayElements {
				val, ok := varParam.ArrayElements[argElemIdx]
				if ok {
					elemSVs := CheckForStateVariable(argElemVar, val)
					for elemSVName, elemSV := range elemSVs {
						allSV[elemSVName] = elemSV
					}
				}
			}
		}

		if len(varArg.FieldElements) != 0 {
			for argElemIdx, argElemVar := range varArg.FieldElements {
				val, ok := varParam.FieldElements[argElemIdx]
				if ok {
					elemSVs := CheckForStateVariable(argElemVar, val)
					for elemSVName, elemSV := range elemSVs {
						allSV[elemSVName] = elemSV
					}
				}
			}
		}
	}
	return allSV
}

func (v *Variable) GetExpression() *Expression {
	if len(v.ArrayElements) == 0 && len(v.FieldElements) == 0 {
		return v.ValueExpression
	} else if len(v.ArrayElements) != 0 {
		arrayExpansion := NewExpression(v.FullName(), VarExp)
		arrayExpansion.Children = append(arrayExpansion.Children, NewExpression(" ( ", Literal))

		var keys []string
		for key := range v.ArrayElements {
			keys = append(keys, key)
		}

		sort.Strings(keys)

		elementCount := 0
		for _, key := range keys {
			value := v.ArrayElements[key].GetExpression()
			if elementCount == 0 {
				arrayExpansion.Children = append(arrayExpansion.Children, value)
			} else {
				arrayExpansion.Children = append(arrayExpansion.Children, NewExpression(" , ", Literal))
				arrayExpansion.Children = append(arrayExpansion.Children, value)
			}
			elementCount += 1
		}
		arrayExpansion.Children = append(arrayExpansion.Children, NewExpression(" ) ", Literal))
		return arrayExpansion

	} else if len(v.FieldElements) != 0 {

		FieldExpansion := NewExpression(v.FullName(), VarExp)
		FieldExpansion.Children = append(FieldExpansion.Children, NewExpression(" ( ", Literal))

		var keys []string
		for key := range v.FieldElements {
			keys = append(keys, key)
		}

		sort.Strings(keys)

		elementCount := 0
		for _, key := range keys {
			value := v.FieldElements[key].GetExpression()
			if elementCount == 0 {
				FieldExpansion.Children = append(FieldExpansion.Children, value)
			} else {
				FieldExpansion.Children = append(FieldExpansion.Children, NewExpression(" , ", Literal))
				FieldExpansion.Children = append(FieldExpansion.Children, value)
			}
			elementCount += 1
		}
		FieldExpansion.Children = append(FieldExpansion.Children, NewExpression(" ) ", Literal))
		return FieldExpansion
	}

	return v.ValueExpression
}

func (_ *Variable) Label() string {
	return "variable"
}

func (v *Variable) FullName() string {
	return v.ContainingFunction.FullName() + ":" + v.Name
}

func (v *Variable) Properties() map[string]interface{} {
	return map[string]interface{}{
		"Name":     v.Name,
		"Type":     v.Type,
		"IsGlobal": v.IsGlobal,
	}
}

type Parameter struct {
	Id       string
	Key      string
	Name     string
	Type     string
	ParamNum int // used for ordering params

	ParamIsVar *Variable
}

func NewParameter() *Parameter {
	return &Parameter{}
}

func (_ *Parameter) Label() string {
	return "parameter"
}

func (v *Parameter) Properties() map[string]interface{} {
	return map[string]interface{}{
		"Name":     v.Name,
		"Type":     v.Type,
		"ParamNum": v.ParamNum,
	}
}

func (v *Parameter) FullName() string {
	return v.ParamIsVar.FullName()
}

type CallingContext struct {
	Caller                 *Function
	Callee                 *Function
	CallingInstruction     *Instruction
	Args                   map[int]*Argument
	ParamToArgs            map[string]*Argument
	StateVariablesInCallee map[string]*Variable
	CallerFlow             *Flow
	CalleeParamVarPointsTo map[string]map[string]*Variable
	ParsedOnce             bool

	//TODO:Add also element and field var maps from caller as well

}

func NewCallingContext() *CallingContext {
	return &CallingContext{
		Args:                   make(map[int]*Argument),
		ParamToArgs:            make(map[string]*Argument),
		StateVariablesInCallee: make(map[string]*Variable),
		CallerFlow:             NewFlow(),
		CalleeParamVarPointsTo: make(map[string]map[string]*Variable),
		ParsedOnce:             false,
	}
}

func GetAllStateVariables(calleeFunction *Function,
	currentContext *CallingContext) map[string]*Variable {
	allSVs := make(map[string]*Variable)
	if currentContext.Callee.FullName() == calleeFunction.FullName() {
		for _, param := range calleeFunction.Parameter {
			v, ok := calleeFunction.ParamIsVar[param]
			if ok {
				if currentContext.ParamToArgs[v.FullName()].CallerVariable != nil {

					for k2, v2 := range currentContext.ParamToArgs[v.FullName()].CallerVariable.PointsTo {
						_, ok2 := currentContext.CalleeParamVarPointsTo[v.FullName()]
						if !ok2 {
							currentContext.CalleeParamVarPointsTo[v.FullName()] = make(map[string]*Variable)
						}
						currentContext.CalleeParamVarPointsTo[v.FullName()][k2] = v2
					}

					allSVsFromVar :=
						CheckForStateVariable(currentContext.
							ParamToArgs[v.FullName()].
							CallerVariable,
							v)

					for k2, v2 := range allSVsFromVar {
						allSVs[k2] = v2
					}
				}
			} else {
				fmt.Printf("Not found\n")
			}
		}
	}

	for _, v := range calleeFunction.LocalVariables {
		allLocalSV := GetAllSVFromVariable(v)
		for k2, v2 := range allLocalSV {
			allSVs[k2] = v2
		}
	}
	return allSVs
}

func CreateContextFromCallerInstruction(callInst *Instruction, callInstFlow *Flow) (*CallingContext, bool) {
	callContext := NewCallingContext()
	callContext.Caller = callInst.Block.Function
	callContext.CallerFlow = callInstFlow

	if len(callInst.InstrCalls) == 0 {
		return nil, false
	}

	for _, v := range callInst.InstrCalls {
		callContext.Callee = v
		//TODO : Function Pointers
	}

	if callContext.Callee == nil {
		return nil, false
	}

	if callContext.Callee.WhiteListed == false {
		return nil, false
	}

	callContext.CallingInstruction = callInst
	for k, v := range callInst.CallerArguments {
		callContext.Args[k] = v
		param, ok := callContext.Callee.Parameter[k]
		if ok {
			calleeParamVar := param.ParamIsVar
			callContext.ParamToArgs[calleeParamVar.FullName()] = v
			//TODO : Add For Elements/Fields of Argument
		}
	}
	callContext.StateVariablesInCallee = GetAllStateVariables(callContext.Callee, callContext)
	return callContext, true
}

func (c *CallingContext) GetArgumentValue(paramName string, f *Flow) string {
	arg, ok := c.ParamToArgs[paramName]
	if ok {
		return arg.GetExpression().GetString(f)
	} else {
		return paramName
	}
}

func (_ *CallingContext) Label() string {
	return "CallingContext"
}

type CallStack struct {
	data []*CallingContext
}

func NewCallStack() *CallStack {
	return &CallStack{}
}

func (s *CallStack) Push(obj *CallingContext) bool {
	if len(s.data) < 10 {
		s.data = append(s.data, obj)
		return true
	}
	return false
}

func (s *CallStack) Top() *CallingContext {
	if len(s.data) > 0 {
		return s.data[len(s.data)-1]
	}
	return nil
}

func (s *CallStack) Pop() *CallingContext {
	if s.IsEmpty() {
		return nil
	}
	index := len(s.data) - 1
	obj := s.data[index]
	s.data = s.data[:index]
	return obj
}

func (s *CallStack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *CallStack) getCalleeRecursionLevel(callContext *CallingContext) int {
	calleeFuncName := callContext.Callee.FullName()
	numFunc := 0
	for _, existingCallFrames := range s.data {
		if existingCallFrames.Callee.FullName() == calleeFuncName {
			numFunc += 1
		}
	}
	return numFunc
}

func (s *CallStack) deepCopyCallStack() *CallStack {
	copiedCallStack := &CallStack{}
	for _, frame := range s.data {
		copiedCallStack.Push(frame)
	}
	return copiedCallStack
}

type Argument struct {
	ConstantValue  string
	CallerVariable *Variable
	Position       int
}

func NewArgument() *Argument {
	return &Argument{}
}

func (a *Argument) FullName() string {
	if a.CallerVariable == nil {
		return a.ConstantValue
	} else {
		return a.CallerVariable.FullName()
	}
}

func (a *Argument) GetExpression() *Expression {
	if a.CallerVariable == nil {
		return NewExpression(a.ConstantValue, Literal)
	} else {
		return a.CallerVariable.GetExpression()
	}
}
