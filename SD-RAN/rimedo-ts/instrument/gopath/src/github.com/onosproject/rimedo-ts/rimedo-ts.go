// SPDX-FileCopyrightText: 2019-present Open Networking Foundation <info@opennetworking.org>
// SPDX-FileCopyrightText: 2019-present Rimedo Labs
//
// SPDX-License-Identifier: Apache-2.0
// Created by RIMEDO-Labs team

//line /home/tianchang/Desktop/proj/rimedo-ts/rimedo-ts.go:7
package main

//line /home/tianchang/Desktop/proj/rimedo-ts/rimedo-ts.go:7
import (
//line /home/tianchang/Desktop/proj/rimedo-ts/rimedo-ts.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/Desktop/proj/rimedo-ts/rimedo-ts.go:7
)
//line /home/tianchang/Desktop/proj/rimedo-ts/rimedo-ts.go:7
import (
//line /home/tianchang/Desktop/proj/rimedo-ts/rimedo-ts.go:7
	_atomic_ "sync/atomic"
//line /home/tianchang/Desktop/proj/rimedo-ts/rimedo-ts.go:7
)

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/onosproject/onos-lib-go/pkg/certs"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/rimedo-ts/pkg/manager"
	"github.com/onosproject/rimedo-ts/pkg/northbound/a1"
	"github.com/onosproject/rimedo-ts/pkg/sdran"
)

var log = logging.GetLogger("rimedo-ts")

func main() {
//line /home/tianchang/Desktop/proj/rimedo-ts/rimedo-ts.go:23
	_go_fuzz_dep_.InstrumentMain()
//line /home/tianchang/Desktop/proj/rimedo-ts/rimedo-ts.go:23
	_go_fuzz_dep_.CoverTab[197217]++

								log.SetLevel(logging.DebugLevel)
								log.Info("Starting RIMEDO Labs Traffic Steering xAPP")

								sdranConfig := sdran.Config{
		AppID:			"rimedo-ts",
		E2tAddress:		"onos-e2t",
		E2tPort:		5150,
		TopoAddress:		"onos-topo",
		TopoPort:		5150,
		SMName:			"oran-e2sm-mho",
		SMVersion:		"v2",
		TSPolicySchemePath:	"/data/schemas/ORAN_TrafficSteeringPreference_v102.json",
	}

	a1Config := a1.Config{
		PolicyName:		"ORAN_TrafficSteeringPreference",
		PolicyVersion:		"2.0.0",
		PolicyID:		"ORAN_TrafficSteeringPreference_2.0.0",
		PolicyDescription:	"O-RAN traffic steering",
		A1tPort:		5150,
	}

	_, err := certs.HandleCertPaths("", "", "", true)
	if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/rimedo-ts.go:48
		_go_fuzz_dep_.CoverTab[197219]++
									log.Fatal(err)
//line /home/tianchang/Desktop/proj/rimedo-ts/rimedo-ts.go:49
		// _ = "end of CoverTab[197219]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/rimedo-ts.go:50
		_go_fuzz_dep_.CoverTab[197220]++
//line /home/tianchang/Desktop/proj/rimedo-ts/rimedo-ts.go:50
		// _ = "end of CoverTab[197220]"
//line /home/tianchang/Desktop/proj/rimedo-ts/rimedo-ts.go:50
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/rimedo-ts.go:50
	// _ = "end of CoverTab[197217]"
//line /home/tianchang/Desktop/proj/rimedo-ts/rimedo-ts.go:50
	_go_fuzz_dep_.CoverTab[197218]++

								mgr := manager.NewManager(sdranConfig, a1Config, false)
								mgr.Run()

								killSignal := make(chan os.Signal, 1)
								signal.Notify(killSignal, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
								log.Debug("app: received a shutdown signal:", <-killSignal)
								mgr.Close()
//line /home/tianchang/Desktop/proj/rimedo-ts/rimedo-ts.go:58
	// _ = "end of CoverTab[197218]"
}

//line /home/tianchang/Desktop/proj/rimedo-ts/rimedo-ts.go:59
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/Desktop/proj/rimedo-ts/rimedo-ts.go:59
var _ = _go_fuzz_dep_.CoverTab
