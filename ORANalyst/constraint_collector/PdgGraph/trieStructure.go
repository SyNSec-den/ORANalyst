package PdgGraph

import (
	"fmt"
	"sort"
)

type TrieNode struct {
	IsLeaf        bool
	StateVariable string
	Children      map[string]*TrieNode // SV value, Children Node
	Parent        *TrieNode
	FSMState      *FSMState
}

func NewTrieNode(isLeaf bool, stateVariable string, state *FSMState) *TrieNode {
	trieNode := &TrieNode{
		Children:      make(map[string]*TrieNode),
		IsLeaf:        isLeaf,
		StateVariable: stateVariable,
	}
	if isLeaf == true {
		trieNode.FSMState = state
	}
	return trieNode
}

type TrieStructure struct {
	root  *TrieNode
	Dummy int
}

func NewTrieStructure() *TrieStructure {
	return &TrieStructure{
		root:  NewTrieNode(true, "", nil),
		Dummy: 0,
	}
}

func (t *TrieNode) PrintTrieNodeDFS() {
	if t == nil {
		fmt.Printf("node is nil\n")
		return
	}
	fmt.Printf("TrieNode %v , SVar %s, ", t.IsLeaf, t.StateVariable)
	if t.FSMState != nil {
		fmt.Printf("State %s\n", t.FSMState.Name)
	} else {
		fmt.Printf("Null State\n")

	}
	fmt.Printf("Children: \n")
	for k, v := range t.Children {
		if v != nil {
			fmt.Printf("TrieNode %v,%s Value %s : %v,%s\n", t.IsLeaf, t.StateVariable, k, v.IsLeaf, v.StateVariable)
		} else {
			fmt.Printf("TrieNode %v,%s Value %s : nil\n", t.IsLeaf, t.StateVariable, k)

		}
	}
	for _, v := range t.Children {
		v.PrintTrieNodeDFS()
	}
}

func (trie *TrieStructure) PrintTrieStructure() {
	fmt.Printf("-----\n")
	currNode := trie.root
	if currNode != nil {
		currNode.PrintTrieNodeDFS()
	} else {
		fmt.Printf("Root is nil\n")
	}
	fmt.Printf("\n-----\n\n")
}

func (trie *TrieStructure) ExtractState(createdState *FSMState) (*FSMState, bool) {
	svValuesStr := make(map[string]string)
	for k, v := range createdState.SvValueExpressions {
		svValuesStr[k] = v.GetString(nil)
	}

	var stateVariables []string
	for k := range svValuesStr {
		stateVariables = append(stateVariables, k)
	}

	sort.Strings(stateVariables)

	checkedSVs := make(map[string]int)
	for _, sv := range stateVariables {
		checkedSVs[sv] = 0
	}

	currentTrieNode := trie.root
	parentTrieNode := trie.root.Parent
	parentStateVariableValue := "UNDEFINED"

	//fmt.Printf("Root %v\n", currentTrieNode.IsLeaf)

	for currentTrieNode.IsLeaf != true {
		currentStateVar := currentTrieNode.StateVariable
		//fmt.Printf("Not Leaf, CurrStateVar %s\n", currentStateVar)

		svValue, ok := svValuesStr[currentStateVar]
		if ok {
			//fmt.Printf("svValue %s\n", svValue)
			newTrieNode, ok2 := currentTrieNode.Children[svValue]
			checkedSVs[currentStateVar] = 1
			if ok2 {
				currentTrieNode = newTrieNode
				parentTrieNode = newTrieNode.Parent
				parentStateVariableValue = svValue

			} else {
				newTrieNode = NewTrieNode(true, "", nil)
				newTrieNode.Parent = currentTrieNode
				currentTrieNode.Children[svValue] = newTrieNode

				currentTrieNode = currentTrieNode.Children[svValue]
				parentTrieNode = currentTrieNode.Parent
				parentStateVariableValue = svValue
			}
		} else {

			newTrieNode, ok2 := currentTrieNode.Children["UNDEFINED"]
			if ok2 {
				currentTrieNode = newTrieNode
				parentTrieNode = newTrieNode.Parent
				parentStateVariableValue = "UNDEFINED"

			} else {
				newTrieNode = NewTrieNode(true, "", nil)
				newTrieNode.Parent = currentTrieNode
				currentTrieNode.Children["UNDEFINED"] = newTrieNode

				currentTrieNode = currentTrieNode.Children["UNDEFINED"]
				parentTrieNode = currentTrieNode.Parent
				parentStateVariableValue = "UNDEFINED"

			}
		}
	}

	checkedAllSVs := true
	for _, v := range checkedSVs {
		if v == 0 {
			checkedAllSVs = false
		}
	}

	if currentTrieNode.FSMState != nil && checkedAllSVs == true {
		//fmt.Printf("Case 1\n")
		//AddChildrenToExpression(currentTrieNode.FSMState.ReachingCondition, createdState.ReachingCondition)
		return currentTrieNode.FSMState, false

	} else if currentTrieNode.FSMState != nil && checkedAllSVs == false {

		originalChild := currentTrieNode

		firstSV := ""
		for k, v := range checkedSVs {
			if v == 0 {
				firstSV = k
				break
			}
		}

		currentTrieNode = NewTrieNode(false, firstSV, nil)
		currentTrieNode.Parent = parentTrieNode

		if parentTrieNode != nil {
			parentTrieNode.Children[parentStateVariableValue] = currentTrieNode
		} else if originalChild == trie.root {
			trie.root = currentTrieNode
		}

		currentTrieNode.Children["UNDEFINED"] = originalChild
		originalChild.Parent = currentTrieNode

		newTrieNode := NewTrieNode(true, "", nil)
		newTrieNode.Parent = currentTrieNode
		currentTrieNode.Children[svValuesStr[firstSV]] = newTrieNode

		checkedSVs[firstSV] = 1
		currentTrieNode = currentTrieNode.Children[svValuesStr[firstSV]]

		for k, v := range checkedSVs {
			if v == 0 {
				currentTrieNode.IsLeaf = false
				currentTrieNode.StateVariable = k

				newChild := NewTrieNode(true, "", nil)

				newChild.Parent = currentTrieNode
				currentTrieNode.Children[svValuesStr[k]] = newChild
				currentTrieNode = newChild
			}
		}

		currentTrieNode.FSMState = createdState
		//fmt.Printf("Case 2\n\n")

		return currentTrieNode.FSMState, true

	} else if currentTrieNode.FSMState == nil && checkedAllSVs == true {
		//fmt.Printf("Case 3\n")

		currentTrieNode.FSMState = createdState
		return currentTrieNode.FSMState, true
	} else {
		firstSV := ""
		for k, v := range checkedSVs {
			if v == 0 {
				firstSV = k
				break
			}
		}

		currentTrieNode.IsLeaf = false
		currentTrieNode.StateVariable = firstSV

		if parentTrieNode != nil {
			parentTrieNode.Children[parentStateVariableValue] = currentTrieNode
		}
		currentTrieNode.Parent = parentTrieNode

		newTrieNode := NewTrieNode(true, "", nil)
		newTrieNode.Parent = currentTrieNode
		currentTrieNode.Children[svValuesStr[firstSV]] = newTrieNode
		checkedSVs[firstSV] = 1

		currentTrieNode = currentTrieNode.Children[svValuesStr[firstSV]]

		for k, v := range checkedSVs {
			if v == 0 {
				currentTrieNode.IsLeaf = false
				currentTrieNode.StateVariable = k

				newChild := NewTrieNode(true, "", nil)
				newChild.Parent = currentTrieNode

				currentTrieNode.Children[svValuesStr[k]] = newChild
				currentTrieNode = newChild
			}
		}

		currentTrieNode.FSMState = createdState
		//fmt.Printf("Case 4\n")

		return currentTrieNode.FSMState, true
	}
}
