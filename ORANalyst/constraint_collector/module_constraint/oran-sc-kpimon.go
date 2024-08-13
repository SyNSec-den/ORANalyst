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

func OranScKpimon() {
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
		var inputType types.Type
		var inputPackage packages.Package
		for name, pkg := range progInfoGraph.ContainsPkg {
			if inputType == nil {
				inputType, _ = condition_collection.LookUpTypeFromPackage(pkg.PackagePtr, "_Ctype_struct_E2SM_KPM_IndicationMessage")
				inputPackage = *pkg.PackagePtr
			}
			if name == "example.com/kpimon/control" {
				// if name == "github.com/onosproject/onos-e2t/pkg/southbound/e2ap/server" {
				for funcName, funcInfo := range pkg.ContainsFunction {
					// if funcName == "RICIndication" {
					if funcName == "handleIndication" {
						entry = funcInfo
					}
				}
			}
			if name == "example.com/kpimon/control" {
				// if name == "github.com/onosproject/onos-e2t/pkg/southbound/e2ap/stream" {
				for funcName, funcInfo := range pkg.ContainsFunction {
					// if funcName == "In" {
					if funcName == "writeCellMetrics_db" {
						exit = funcInfo
					}
				}
			}
		}
		fmt.Printf("input type: %v\n", inputType)
		cc := condition_collection.NewConditionCollector("routing-manager", []types.Type{inputType}, nil, inputPackage)
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
