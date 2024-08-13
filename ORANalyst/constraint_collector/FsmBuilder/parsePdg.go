package FsmBuilder

import (
	"bufio"
	"constraint_collector/PdgGraph"
	"fmt"
	"os"
	"strings"
)

func ParseWhiteListFunctions(pdgGraph *PdgGraph.ProgramInfo,
	inFile string) map[string]*PdgGraph.Function {
	wlFunctions := make(map[string]*PdgGraph.Function)

	file, err := os.Open(inFile)
	if err != nil {
		return wlFunctions
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			return wlFunctions
		}
		pkgAndFunc := strings.TrimSpace(strings.Split(line, "\n")[0])
		if strings.Contains(pkgAndFunc, ":") {
			pkgPath := strings.TrimSpace(strings.Split(pkgAndFunc, ":")[0])
			fnName := strings.TrimSpace(strings.Split(pkgAndFunc, ":")[1])

			for pkgKey, pkg := range pdgGraph.ContainsPkg {
				if pkgKey == pkgPath {
					if fnName == "" {
						for _, wlfn := range pkg.ContainsFunction {
							//fmt.Printf("%s is WhiteListed\n", wlfn.FullName())
							wlFunctions[wlfn.FullName()] = wlfn
						}
					} else {
						for wlfnName, wlfn := range pkg.ContainsFunction {
							if wlfnName == fnName {
								//fmt.Printf("%s is WhiteListed\n", wlfn.FullName())
								wlFunctions[wlfn.FullName()] = wlfn
							}
						}
					}
				}
			}
		}
	}

	return wlFunctions
}

func MarkActionFunction(
	pdgGraph *PdgGraph.ProgramInfo,
	inFile string) {

	file, err := os.Open(inFile)
	if err != nil {
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			return
		}
		pkgAndFunc := strings.TrimSpace(strings.Split(line, "\n")[0])
		if strings.Contains(pkgAndFunc, ":") {
			pkgPath := strings.TrimSpace(strings.Split(pkgAndFunc, ":")[0])
			fnName := strings.TrimSpace(strings.Split(pkgAndFunc, ":")[1])

			for pkgKey, pkg := range pdgGraph.ContainsPkg {
				if pkgKey == pkgPath {
					if fnName == "" {
						for _, actionFn := range pkg.ContainsFunction {
							//fmt.Printf("%s is an action Function\n", actionFn.FullName())
							actionFn.IsAction = true
						}
					} else {
						for actionFnName, actionFn := range pkg.ContainsFunction {
							if actionFnName == fnName {
								//fmt.Printf("%s is an action Function\n", actionFn.FullName())
								actionFn.IsAction = true
							}
						}
					}
				}
			}
		}
	}
}

func ParsePdgGraph(pdgGraph *PdgGraph.ProgramInfo) {

	MarkActionFunction(pdgGraph, "AnnotationPos/ActionFunctions.txt")

	wlFunctionsMap := ParseWhiteListFunctions(pdgGraph, "AnnotationPos/WhiteListFunctions.txt")
	var mainFunction *PdgGraph.Function
	for k, v := range wlFunctionsMap {
		//fmt.Printf("%s\n", k)
		if strings.Contains(k, ":main") {
			mainFunction = v
		}
	}

	for _, fn := range wlFunctionsMap {
		fn.WhiteListed = true
	}

	for _, fn := range wlFunctionsMap {
		//fmt.Printf("Testing IC Function %s\n", fn.FullName())
		ProcessFunctionContextInsensitive(fn)
		//fmt.Printf("\n\n\n")
	}

	mainContext := PdgGraph.NewCallingContext()
	mainContext.Caller = nil
	mainContext.Callee = mainFunction
	mainContext.StateVariablesInCallee = make(map[string]*PdgGraph.Variable)
	for _, v := range mainFunction.LocalVariables {
		if v.AliasToStateVariable != nil {
			mainContext.StateVariablesInCallee[v.FullName()] = v.AliasToStateVariable
		}
	}

	callStack := PdgGraph.NewCallStack()
	callStack.Push(mainContext)

	functionFSM := ProcessFunctionContextSensitive(mainFunction, callStack)

	//for _, fn := range wlFunctionsMap {
	//	fmt.Printf("Function %s\n", fn.FullName())
	//	for _, v := range fn.LocalVariables {
	//		fmt.Printf("Variable %s : %s\n", v.FullName(), v.ValueExpression.GetString(nil))
	//	}
	//	fmt.Printf("\n")
	//}

	//for _, v := range mainFunction.LocalVariables {
	//	str := "No SV"
	//	if v.AliasToStateVariable != nil {
	//		str = v.AliasToStateVariable.FullName()
	//	}
	//
	//	fmt.Printf("Variable %s Expression Value Name %s Type %d Alias %s\n", v.FullName(),
	//		v.ValueExpression.LiteralExpression,
	//		v.ValueExpression.ExpressionType, str)
	//	fmt.Printf("Expression %s\n\n", v.ValueExpression.GetString(nil))
	//}

	functionFSM.WriteFSMInfoToFile()
	fmt.Printf("Total States %d\n", len(functionFSM.States))
	fmt.Printf("Total Transitions %d\n", len(functionFSM.Transitions))

	//functionFSM.PrintFSM()
}
