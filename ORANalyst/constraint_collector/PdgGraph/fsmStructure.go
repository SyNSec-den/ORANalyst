package PdgGraph

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

type FSMState struct {
	//combination of state variable value
	Name               string
	SvValueExpressions map[string]*Expression
}

func NewFSMState(svValues map[string]*Expression) *FSMState {
	svValueExpressions := make(map[string]*Expression)
	for k, v := range svValues {
		svValueExpressions[k] = v
	}
	return &FSMState{
		SvValueExpressions: svValueExpressions,
	}
}

func (state *FSMState) PrintFSMState() {
	fmt.Printf("----\n")
	fmt.Printf("State Name : %s\n", state.Name)
	// fmt.Printf("Reaching Condition : %s\n", state.ReachingCondition.GetString(nil))
	fmt.Printf("SV Values: ")

	var keys []string
	for k := range state.SvValueExpressions {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("var %s : val %s\n", k, state.SvValueExpressions[k].GetString(nil))
	}
	fmt.Printf("\n\n")
	fmt.Printf("----\n")
}

func (state *FSMState) StateInfo() string {

	str := "-----------------\n"
	str += "State Name : " + state.Name + "\n"
	str += "State Variable & Values : \n"
	for k, v := range state.SvValueExpressions {
		str += k + " :-> " + v.GetString(nil) + "\n"
	}
	str += "-----------------\n"
	return str
}

func PrintSVVAlues(svValues map[string]*Expression) {
	for k, v := range svValues {
		fmt.Printf("; %s : %s ; ", k, v.GetString(nil))
	}
	fmt.Printf("\n")
}

type FSMTransition struct {
	Id                string
	Name              string
	FromState         *FSMState
	ToState           *FSMState
	Condition         *Expression
	IsFinalTransition bool
	Actions           []string
}

func NewFSMTransition(fromState *FSMState, toState *FSMState, condition *Expression) *FSMTransition {
	return &FSMTransition{
		FromState:         fromState,
		ToState:           toState,
		Condition:         condition,
		IsFinalTransition: false,
	}
}

func (t *FSMTransition) PrintTransition() {
	fmt.Printf("Transition %s -> %s\n", t.FromState.Name, t.ToState.Name)
}

func (t *FSMTransition) TxInfo() string {
	str := "<transition label " + t.Id + " >\n"
	str += "<start> " + t.FromState.Name + " </start> \n"
	str += "<end> " + t.ToState.Name + " </end> \n"
	//fmt.Printf("Transition %s Children %d\n", t.Id, len(t.Condition.Children))
	str += "<condition> " + t.Condition.GetString(nil) + " </condition> \n"
	str += "<actions> \n"

	for _, a := range t.Actions {
		str += "<action label = " + a + " ></action>\n"
	}

	str += "</actions> \n\n"

	return str
}

func (t *FSMTransition) GetName() string {
	return t.FromState.Name + " --> " + t.ToState.Name
}

type BlockStateIO struct {
	InputTransitions   map[string]*FSMTransition
	OutputTransitions  map[string]*FSMTransition
	VisitedTransitions map[string]*FSMTransition
	VisitedStates      map[string]*FSMState
	BlockVisitCount    map[string]int
}

func NewBlockStateIO() *BlockStateIO {
	return &BlockStateIO{
		InputTransitions:   make(map[string]*FSMTransition),
		OutputTransitions:  make(map[string]*FSMTransition),
		VisitedTransitions: make(map[string]*FSMTransition),
		VisitedStates:      make(map[string]*FSMState),
		BlockVisitCount:    make(map[string]int),
	}
}

type FSM struct {
	ParentFunction   *Function
	ContextSignature string

	InitState        *FSMState
	InitTransition   *FSMTransition
	FinalTransitions map[string]*FSMTransition

	States              map[string]*FSMState
	Transitions         map[string]*FSMTransition
	StateCount          int
	TransitionCount     int
	StatesTrieStructure *TrieStructure

	BlockInfo          map[*Block]*BlockStateIO
	MaxBlockVisitCount int
}

func NewFSM() *FSM {
	newFSM := &FSM{
		States:              make(map[string]*FSMState),
		Transitions:         make(map[string]*FSMTransition),
		BlockInfo:           make(map[*Block]*BlockStateIO),
		FinalTransitions:    make(map[string]*FSMTransition),
		StatesTrieStructure: NewTrieStructure(),
		StateCount:          0,
		TransitionCount:     0,
		MaxBlockVisitCount:  2,
	}
	newInitState := NewFSMState(make(map[string]*Expression))
	// newInitState.ReachingCondition = NewExpression(" true ", Literal)
	newFSM.AddNewFSMState(newInitState)

	newInitTx := NewFSMTransition(newInitState, newInitState, NewExpression(" true ", Literal))
	newFSM.AddNewFSMTransition(newInitTx)
	newFSM.InitState = newInitState
	newFSM.InitTransition = newInitTx

	return newFSM
}

func (fsmInstance *FSM) WriteFSMInfoToFile() {

	file, err := os.Create("fsmOutputs/states.txt")
	if err != nil {
		fmt.Println("Failed to create file:", err)
		return
	}

	var keys []string
	for k := range fsmInstance.States {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		state := fsmInstance.States[k]
		stateInfoStr := state.StateInfo() + "\n"
		_, err = file.WriteString(stateInfoStr)
		if err != nil {
			fmt.Println("Failed to write to file:", err)
			return
		}
	}

	file.Close()

	file, err = os.Create("fsmOutputs/transitions.txt")
	if err != nil {
		fmt.Println("Failed to create file:", err)
		return
	}

	keys = keys[:0]
	for k := range fsmInstance.Transitions {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		tx := fsmInstance.Transitions[k]
		txInfoStr := tx.TxInfo()
		_, err = file.WriteString(txInfoStr)
		if err != nil {
			fmt.Println("Failed to write to file:", err)
			return
		}
	}

	file.Close()

	fmt.Println("FSM written successfully!")
}

func (fsmInstance *FSM) AddNewFSMState(newState *FSMState) *FSMState {
	fsmState, isNew := fsmInstance.StatesTrieStructure.ExtractState(newState)
	if isNew {
		fsmInstance.StateCount += 1
		newStateName := "s_" + strconv.Itoa(fsmInstance.StateCount)
		fsmState.Name = newStateName
	}
	fsmInstance.States[fsmState.Name] = fsmState
	return fsmState
}

func (fsmInstance *FSM) AddNewFSMTransition(transition *FSMTransition) *FSMTransition {

	fsmInstance.TransitionCount += 1
	newTxName := transition.FromState.Name + " --> " + transition.ToState.Name
	newTxId := "T_" + strconv.Itoa(fsmInstance.TransitionCount)
	transition.Name = newTxName
	transition.Id = newTxId
	fsmInstance.Transitions[newTxId] = transition
	return transition
}

func (fsmInstance *FSM) AddToBlockInputs(b *Block, transition *FSMTransition) {
	_, ok := fsmInstance.BlockInfo[b]
	if !ok {
		fsmInstance.BlockInfo[b] = NewBlockStateIO()
		fsmInstance.BlockInfo[b].InputTransitions[transition.GetName()] = transition
	}
	fsmInstance.BlockInfo[b].InputTransitions[transition.GetName()] = transition
}

func (fsmInstance *FSM) AddToBlockOutputs(b *Block, transition *FSMTransition) {
	_, ok := fsmInstance.BlockInfo[b]
	if !ok {
		fsmInstance.BlockInfo[b] = NewBlockStateIO()
		fsmInstance.BlockInfo[b].OutputTransitions[transition.GetName()] = transition
	}
	fsmInstance.BlockInfo[b].OutputTransitions[transition.GetName()] = transition
}

func (fsmInstance *FSM) AddToBlockVisited(b *Block, transition *FSMTransition) {
	_, ok := fsmInstance.BlockInfo[b]
	if !ok {
		fsmInstance.BlockInfo[b] = NewBlockStateIO()
		fsmInstance.BlockInfo[b].VisitedTransitions[transition.GetName()] = transition
		fsmInstance.BlockInfo[b].VisitedStates[transition.ToState.Name] = transition.ToState
	}
	fsmInstance.BlockInfo[b].VisitedTransitions[transition.GetName()] = transition
	fsmInstance.BlockInfo[b].VisitedStates[transition.ToState.Name] = transition.ToState
}

func (fsmInstance *FSM) CheckBlockVisit(predBlock *Block, b *Block) bool {
	_, ok := fsmInstance.BlockInfo[b]
	if !ok {
		fsmInstance.BlockInfo[b] = NewBlockStateIO()
		fsmInstance.BlockInfo[b].BlockVisitCount[predBlock.Name] = 1
		return true
	} else {
		visitCount, ok2 := fsmInstance.BlockInfo[b].BlockVisitCount[predBlock.Name]
		if ok2 {
			if visitCount < fsmInstance.MaxBlockVisitCount {
				fsmInstance.BlockInfo[b].BlockVisitCount[predBlock.Name] += 1
				return true
			} else {
				return false
			}
		} else {
			fsmInstance.BlockInfo[b].BlockVisitCount[predBlock.Name] = 1
			return true
		}
	}
}

func (fsmInstance *FSM) PrintFSM() {
	for k, v := range fsmInstance.Transitions {
		fmt.Printf("Transition %s\n", k)
		v.PrintTransition()
	}
}

func (fsmInstance *FSM) ObtainTxWithNewStates(b *Block) map[string]*FSMTransition {

	newInputTransitions := make(map[string]*FSMTransition)
	for inputTxName, inputTx := range fsmInstance.BlockInfo[b].InputTransitions {
		newInputTransitions[inputTxName] = inputTx
	}

	return newInputTransitions
}

func DeepCopyTransitions(transitions map[string]*FSMTransition) map[string]*FSMTransition {
	newTxs := make(map[string]*FSMTransition)
	for k, v := range transitions {
		newTxs[k] = v
	}
	return newTxs
}

func ConvertSVMapToSignature(SVMap map[string]*Variable) string {
	var keys []string
	for key := range SVMap {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	signature := "####"
	for _, key := range keys {
		signature += key + "####"
	}
	return signature
}

func (fsmInstance *FSM) AddCalleeFuncFSM(calleeFSM *FSM,
	callingState *FSMState,
	callFrame *CallingContext,
	callerFlow *Flow) map[string]*FSMTransition {

	calleeTocallerMap := make(map[string]string)

	currStateCount := fsmInstance.StateCount

	for k := range calleeFSM.States {
		if k == calleeFSM.InitState.Name {
			calleeTocallerMap[k] = callingState.Name
		} else {
			calleeTocallerMap[k] = "s_" + strconv.Itoa(currStateCount)
			currStateCount += 1
		}
	}
	for k, v := range calleeFSM.States {
		if k != calleeFSM.InitState.Name {
			newFSMState := NewFSMState(make(map[string]*Expression))
			for k2, v2 := range v.SvValueExpressions {
				newFSMState.SvValueExpressions[callFrame.GetArgumentValue(k2, nil)] = v2.GetCallerExpression(callFrame)
			}
			newFSMState = fsmInstance.AddNewFSMState(newFSMState)
			calleeTocallerMap[v.Name] = newFSMState.Name
		}
	}

	finalFSMMap := make(map[string]*FSMTransition)

	for k := range fsmInstance.States {
		fmt.Printf("%s\n", k)
	}

	for _, v := range calleeFSM.Transitions {
		fmt.Printf("A %s %s %s\n", v.FromState.Name, v.ToState.Name, v.IsFinalTransition)
		if !(v.FromState.Name == v.ToState.Name &&
			v.FromState.Name == calleeFSM.InitState.Name) || len(v.Actions) != 0 {
			fmt.Printf("B %s %s\n", calleeTocallerMap[v.FromState.Name], calleeTocallerMap[v.ToState.Name])
			newTransition := NewFSMTransition(fsmInstance.States[calleeTocallerMap[v.FromState.Name]],
				fsmInstance.States[calleeTocallerMap[v.ToState.Name]],
				v.Condition.GetCallerExpression(callFrame).GetFlowExpression(callerFlow))
			newTransition.Name = calleeTocallerMap[v.FromState.Name] + " --> " + calleeTocallerMap[v.ToState.Name]
			newTransition = fsmInstance.AddNewFSMTransition(newTransition)

			if v.IsFinalTransition == true {
				finalFSMMap[newTransition.Name] = newTransition
			}
		}
	}

	return finalFSMMap
}
