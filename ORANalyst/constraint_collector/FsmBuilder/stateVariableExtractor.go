package FsmBuilder

import (
	"bufio"
	"constraint_collector/PdgGraph"
	"fmt"
	"os"
	"strings"
)

//func GetAllVariables(pdgGraph *PdgGraph.ProgramInfo, outFile string) {
//	file, err := os.Create(outFile)
//	if err != nil {
//		fmt.Printf("Error opening file %s\n", outFile)
//		return
//	}
//	fmt.Printf("Dumping All Variables")
//	for pkgKey, pkg := range pdgGraph.ContainsPkg {
//		if !strings.HasPrefix(pkgKey, "main") {
//			continue
//		}
//		//fmt.Printf("<Package Name : %v>\n", pkgKey)
//		_, err := file.WriteString(fmt.Sprintf("<Package Name : %s>\n", pkgKey))
//		if err != nil {
//			return
//		}
//		for fnKey, fn := range pkg.ContainsFunction {
//			if fnKey == "init" {
//				continue
//			}
//			_, err := file.WriteString(fmt.Sprintf("\t<Function Name : %s>\n", fnKey))
//			if err != nil {
//				return
//			}
//			for varName, varObj := range fn.LocalVariables {
//				if varObj.DefiningInstruction != nil {
//					_, err := file.WriteString(fmt.Sprintf(
//						"\t\t<Variable Name : %s, File : %s, Pos : %s>\n",
//						varName, varObj.DefiningInstruction.File, varObj.DefiningInstruction.Pos))
//					if err != nil {
//						return
//					}
//				} else {
//					_, err := file.WriteString(fmt.Sprintf(
//						"\t\t<Variable Name : %s has no defining Instruction\n",
//						varName))
//					if err != nil {
//						return
//					}
//				}
//
//			}
//		}
//	}
//}

func GetStateVariableMap(pdgGraph *PdgGraph.ProgramInfo, inFile string) *PdgGraph.ProgramInfo {
	// TODO : If state variables are in multiple files
	// TODO : If state variables are structure fields

	file, err := os.Open(inFile)
	if err != nil {
		fmt.Printf("Error opening file %s\n", inFile)
		return pdgGraph
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	scanner := bufio.NewScanner(file)
	var svFunction *PdgGraph.Function

	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			return pdgGraph
		}
		line = strings.TrimSpace(strings.Split(line, "\n")[0])

		if strings.HasPrefix(line, "<") {
			pkgAndFunc := strings.Split(strings.Split(line, "<")[1], ">")[0]
			pkgPath := strings.Split(pkgAndFunc, ":")[0]
			fnName := strings.Split(pkgAndFunc, ":")[1]

			for pkgKey, pkg := range pdgGraph.ContainsPkg {
				if pkgKey == pkgPath {
					for fnKey, fn := range pkg.ContainsFunction {
						if fnKey == fnName {
							svFunction = fn
						}
					}
					break
				}
			}
		} else {
			line = strings.TrimSpace(strings.Split(line, "\n")[0])

			for varName, varObj := range svFunction.LocalVariables {
				if varName == svFunction.FullName()+":"+line {
					fmt.Printf("variable %s in Function %s is SV\n", varObj.FullName(), varObj.ContainingFunction.FullName())
					varObj.AliasToStateVariable = varObj
				}
			}
		}
	}

	return pdgGraph
}
