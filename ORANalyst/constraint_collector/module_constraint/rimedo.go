package module_constraint

import (
	"constraint_collector/PdgGraph"
	"constraint_collector/condition_collection"
	"flag"
	"fmt"
	"go/types"
	"strings"

	"golang.org/x/tools/go/packages"
)

func Rimedo() {
	progInfoGraph := PdgGraph.NewProgramInfo()
	fmt.Printf("------------- processing path constraint -------------\n")
	for i := 0; i < flag.NArg(); i++ {
		path := flag.Arg(i)
		if path[len(path)-1] == []byte("/")[0] {
			path = path[0 : len(path)-1]
		}
		pathSplit := strings.Split(path, "/")
		module := pathSplit[len(pathSplit)-1]
		fmt.Printf("------------- processing %v -------------\n", module)
		parser := PdgGraph.NewProgParser(module)
		progInfoGraph = parser.ProcessPackage(path, progInfoGraph)
		var entry *PdgGraph.Function
		var exit *PdgGraph.Function
		var msgType types.Type
		var hdrType types.Type
		var inputPackage packages.Package
		for name, pkg := range progInfoGraph.ContainsPkg {
			if msgType == nil {
				msgType, _ = condition_collection.LookUpTypeFromPackage(pkg.PackagePtr, "E2SmMhoIndicationMessage")
				inputPackage = *pkg.PackagePtr
			}
			if hdrType == nil {
				hdrType, _ = condition_collection.LookUpTypeFromPackage(pkg.PackagePtr, "E2SmMhoIndicationHeader")
				// inputPackage = *pkg.PackagePtr
			}

			fmt.Printf("------------- pkg name: %v -------------\n", name)
			for funcName := range pkg.ContainsFunction {
				fmt.Printf("funcName: %v\n", funcName)
			}

			if name == "github.com/onosproject/rimedo-ts/pkg/mho" {
				// if name == "github.com/onosproject/onos-e2t/pkg/southbound/e2ap/server" {
				for funcName, funcInfo := range pkg.ContainsFunction {
					// if funcName == "RICIndication" {
					if funcName == "listenIndChan" {
						entry = funcInfo
						fmt.Printf("------------- entry -------------\n")
						condition_collection.PrintFuncSSA(funcInfo)
						fmt.Printf("------------- end of entry -------------\n")
						// condition_collection.PrintFuncSSAToDot(funcInfo, "rimedo_listenIndChan.dot")
					}
				}
			}
			if name == "github.com/onosproject/rimedo-ts/pkg/mho" {
				// if name == "github.com/onosproject/onos-e2t/pkg/southbound/e2ap/stream" {
				for funcName, funcInfo := range pkg.ContainsFunction {
					// if funcName == "In" {
					if funcName == "SetUe" {
						exit = funcInfo
					}
				}
			}
		}
		fmt.Printf("msg type: %v\n", msgType)
		fmt.Printf("hdr type: %v\n", hdrType)
		fmt.Printf("entry: %+v, exit: %+v\n", *entry, *exit)

		cc := condition_collection.NewConditionCollector("rimedo-ts", []types.Type{msgType, hdrType}, nil, inputPackage)
		callPaths, conditions := cc.CollectPathConditions(entry, exit)
		resStr := ""
		for i, path := range conditions {
			resStr += fmt.Sprintf("-------------------- path %v --------------------\n", i)
			fmt.Printf("-------------------- path %v --------------------\n", i)
			resStr += condition_collection.PrintCallPath(callPaths[i])
			condMap := make(map[condition_collection.Condition]bool)
			for cond := range path {
				if _, ok := condMap[cond]; ok {
					continue
				}
				condMap[cond] = true
				fmt.Printf("solving condition: %v - %v\n", cond.Instr.Text, cond.Cond)
				resStr += fmt.Sprintf("solving condition: %v - %v\n", cond.Instr.Text, cond.Cond)
				constraints := cc.SolveCondition(cond)
				for _, c := range constraints {
					fmt.Print(condition_collection.PrintConstraint(c))
					resStr += condition_collection.PrintConstraint(c)
				}
			}
			resStr += fmt.Sprintf("-------------------- end path %v --------------------\n", i)
			fmt.Printf("-------------------- end path %v --------------------\n", i)
		}
		fmt.Printf("%v\n", resStr)
		return
	}
}
