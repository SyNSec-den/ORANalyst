package main

import (
	"constraint_collector/FsmBuilder"
	"constraint_collector/PdgGraph"
	"constraint_collector/module_constraint"
	"flag"
	"fmt"
	"strings"
)

func main() {
	//fmt.Printf("A")
	//PdgGraph.ComputeDomTree()
	//fmt.Printf("Done\n")
	printPDG := flag.Bool("print-pdg", false, "print PDG of the program to file")
	pathConstraint := flag.Bool("path-constraint", false, "print path constraint of the program to file")
	module := flag.String("module", "", "module to process")
	flag.Parse()
	if *pathConstraint {
		switch *module {
		case "onos-e2t-indication":
			module_constraint.OnosE2tIndication()
		case "onos-e2t-e2setup":
			module_constraint.OnosE2tE2Setup()
		case "rtmgr":
			module_constraint.RtMgr()
		case "rimedo":
			module_constraint.Rimedo()
		case "oran-sc-kpimon":
			module_constraint.OranScKpimon()
		default:
			fmt.Printf("module not found\n")
		}
		return
	}
	if *printPDG {
		progInfoGraph := PdgGraph.NewProgramInfo()
		fmt.Printf("------------- printing PDG -------------\n")
		for i := 0; i < flag.NArg(); i++ {
			path := flag.Arg(i)
			if path[len(path)-1] == []byte("/")[0] {
				path = path[0 : len(path)-1]
			}
			pathSplit := strings.Split(path, "/")
			module := pathSplit[len(pathSplit)-1]
			printer := PdgGraph.NewProgParser(module)
			progInfoGraph = printer.ProcessPackage(path, progInfoGraph)
			//FsmBuilder.ParsePdgGraph(progInfoGraph, fmt.Sprintf("%v/%v_vars.txt", path, module))
			//fmt.Printf("------------- output to %v/%v_pdg.txt -------------\n", path, module)
			FsmBuilder.GetStateVariableMap(progInfoGraph, "AnnotationPos/StateVariablePos.txt")
			FsmBuilder.ParsePdgGraph(progInfoGraph)

		}
		return
	}
	//exectest.Z3TestFunc()
}
