//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Event-Trigger-Definition.go:1
// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Event-Trigger-Definition.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Event-Trigger-Definition.go:1
// SPDX-License-Identifier: Apache-2.0
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Event-Trigger-Definition.go:4
package pdubuilder

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Event-Trigger-Definition.go:4
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Event-Trigger-Definition.go:4
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Event-Trigger-Definition.go:4
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Event-Trigger-Definition.go:4
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Event-Trigger-Definition.go:4
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Event-Trigger-Definition.go:4
)

import (
	"fmt"
	e2sm_mho_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-mho-go"
)

func CreateE2SmMhoEventTriggerDefinition(triggerType e2sm_mho_go.MhoTriggerType) (*e2sm_mho_go.E2SmMhoEventTriggerDefinition, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Event-Trigger-Definition.go:11
	_go_fuzz_dep_.CoverTab[192377]++

																				eventDefinitionFormat1 := &e2sm_mho_go.E2SmMhoEventTriggerDefinitionFormat1{
		TriggerType: triggerType,
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Event-Trigger-Definition.go:16
	}

	E2SmMhoPdu := e2sm_mho_go.E2SmMhoEventTriggerDefinition{
		EventDefinitionFormats: &e2sm_mho_go.MhoEventTriggerDefinitionFormats{
			E2SmMhoEventTriggerDefinition: &e2sm_mho_go.MhoEventTriggerDefinitionFormats_EventDefinitionFormat1{
				EventDefinitionFormat1: eventDefinitionFormat1,
			},
		},
	}

	if err := E2SmMhoPdu.Validate(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Event-Trigger-Definition.go:26
		_go_fuzz_dep_.CoverTab[192379]++
																					return nil, fmt.Errorf("CreateE2SmMhoEventTriggerDefinition(): error validating E2SmMhoPDU %s", err.Error())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Event-Trigger-Definition.go:27
		// _ = "end of CoverTab[192379]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Event-Trigger-Definition.go:28
		_go_fuzz_dep_.CoverTab[192380]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Event-Trigger-Definition.go:28
		// _ = "end of CoverTab[192380]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Event-Trigger-Definition.go:28
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Event-Trigger-Definition.go:28
	// _ = "end of CoverTab[192377]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Event-Trigger-Definition.go:28
	_go_fuzz_dep_.CoverTab[192378]++
																				return &E2SmMhoPdu, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Event-Trigger-Definition.go:29
	// _ = "end of CoverTab[192378]"
}

func CreateE2SmMhoEventTriggerDefinitionPeriodic(rtPeriod int32) (*e2sm_mho_go.E2SmMhoEventTriggerDefinition, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Event-Trigger-Definition.go:32
	_go_fuzz_dep_.CoverTab[192381]++

																				eventDefinitionFormat1 := &e2sm_mho_go.E2SmMhoEventTriggerDefinitionFormat1{
		TriggerType:		e2sm_mho_go.MhoTriggerType_MHO_TRIGGER_TYPE_PERIODIC,
		ReportingPeriodMs:	&rtPeriod,
	}

	E2SmMhoPdu := e2sm_mho_go.E2SmMhoEventTriggerDefinition{
		EventDefinitionFormats: &e2sm_mho_go.MhoEventTriggerDefinitionFormats{
			E2SmMhoEventTriggerDefinition: &e2sm_mho_go.MhoEventTriggerDefinitionFormats_EventDefinitionFormat1{
				EventDefinitionFormat1: eventDefinitionFormat1,
			},
		},
	}

	if err := E2SmMhoPdu.Validate(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Event-Trigger-Definition.go:47
		_go_fuzz_dep_.CoverTab[192383]++
																					return nil, fmt.Errorf("CreateE2SmMhoEventTriggerDefinitionPeriodic(): error validating E2SmMhoPDU %s", err.Error())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Event-Trigger-Definition.go:48
		// _ = "end of CoverTab[192383]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Event-Trigger-Definition.go:49
		_go_fuzz_dep_.CoverTab[192384]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Event-Trigger-Definition.go:49
		// _ = "end of CoverTab[192384]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Event-Trigger-Definition.go:49
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Event-Trigger-Definition.go:49
	// _ = "end of CoverTab[192381]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Event-Trigger-Definition.go:49
	_go_fuzz_dep_.CoverTab[192382]++
																				return &E2SmMhoPdu, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Event-Trigger-Definition.go:50
	// _ = "end of CoverTab[192382]"
}

func CreateE2SmMhoEventTriggerDefinitionUponRcvMeasReport() (*e2sm_mho_go.E2SmMhoEventTriggerDefinition, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Event-Trigger-Definition.go:53
	_go_fuzz_dep_.CoverTab[192385]++

																				eventDefinitionFormat1 := &e2sm_mho_go.E2SmMhoEventTriggerDefinitionFormat1{
		TriggerType: e2sm_mho_go.MhoTriggerType_MHO_TRIGGER_TYPE_UPON_RCV_MEAS_REPORT,
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Event-Trigger-Definition.go:58
	}

	E2SmMhoPdu := e2sm_mho_go.E2SmMhoEventTriggerDefinition{
		EventDefinitionFormats: &e2sm_mho_go.MhoEventTriggerDefinitionFormats{
			E2SmMhoEventTriggerDefinition: &e2sm_mho_go.MhoEventTriggerDefinitionFormats_EventDefinitionFormat1{
				EventDefinitionFormat1: eventDefinitionFormat1,
			},
		},
	}

	if err := E2SmMhoPdu.Validate(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Event-Trigger-Definition.go:68
		_go_fuzz_dep_.CoverTab[192387]++
																					return nil, fmt.Errorf("CreateE2SmMhoEventTriggerDefinitionUponRcvMeasReport(): error validating E2SmMhoPDU %s", err.Error())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Event-Trigger-Definition.go:69
		// _ = "end of CoverTab[192387]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Event-Trigger-Definition.go:70
		_go_fuzz_dep_.CoverTab[192388]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Event-Trigger-Definition.go:70
		// _ = "end of CoverTab[192388]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Event-Trigger-Definition.go:70
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Event-Trigger-Definition.go:70
	// _ = "end of CoverTab[192385]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Event-Trigger-Definition.go:70
	_go_fuzz_dep_.CoverTab[192386]++
																				return &E2SmMhoPdu, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Event-Trigger-Definition.go:71
	// _ = "end of CoverTab[192386]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Event-Trigger-Definition.go:72
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Event-Trigger-Definition.go:72
var _ = _go_fuzz_dep_.CoverTab
