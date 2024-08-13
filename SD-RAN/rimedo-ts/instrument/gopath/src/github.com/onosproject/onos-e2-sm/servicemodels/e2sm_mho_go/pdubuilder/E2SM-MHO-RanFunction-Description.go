//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-RanFunction-Description.go:1
// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-RanFunction-Description.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-RanFunction-Description.go:1
// SPDX-License-Identifier: Apache-2.0
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-RanFunction-Description.go:4
package pdubuilder

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-RanFunction-Description.go:4
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-RanFunction-Description.go:4
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-RanFunction-Description.go:4
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-RanFunction-Description.go:4
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-RanFunction-Description.go:4
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-RanFunction-Description.go:4
)

import (
	"fmt"
	e2sm_mho_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-mho-go"
	e2sm_v2_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-v2-ies"
)

func CreateE2SmMhoRanfunctionDescriptionMsg(ranFunctionShortName string, ranFunctionE2SmOid string, ranFunctionDescription string) (*e2sm_mho_go.E2SmMhoRanfunctionDescription, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-RanFunction-Description.go:12
	_go_fuzz_dep_.CoverTab[192434]++

																			e2smMhoPdu := e2sm_mho_go.E2SmMhoRanfunctionDescription{
		RanFunctionName: &e2sm_v2_ies.RanfunctionName{
			RanFunctionShortName:	ranFunctionShortName,
			RanFunctionE2SmOid:	ranFunctionE2SmOid,
			RanFunctionDescription:	ranFunctionDescription,
		},
		E2SmMhoRanfunctionItem:	&e2sm_mho_go.E2SmMhoRanfunctionDescription_E2SmMhoRanfunctionItem001{},
	}

	if err := e2smMhoPdu.Validate(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-RanFunction-Description.go:23
		_go_fuzz_dep_.CoverTab[192436]++
																				return nil, fmt.Errorf("CreateE2SmMhoRanfunctionDescriptionMsg(): error validating E2SmPDU %s", err.Error())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-RanFunction-Description.go:24
		// _ = "end of CoverTab[192436]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-RanFunction-Description.go:25
		_go_fuzz_dep_.CoverTab[192437]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-RanFunction-Description.go:25
		// _ = "end of CoverTab[192437]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-RanFunction-Description.go:25
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-RanFunction-Description.go:25
	// _ = "end of CoverTab[192434]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-RanFunction-Description.go:25
	_go_fuzz_dep_.CoverTab[192435]++
																			return &e2smMhoPdu, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-RanFunction-Description.go:26
	// _ = "end of CoverTab[192435]"
}

func CreateRicEventTriggerStyleItem(ricStyleType int32, ricStyleName string, ricFormatType int32) (*e2sm_mho_go.RicEventTriggerStyleList, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-RanFunction-Description.go:29
	_go_fuzz_dep_.CoverTab[192438]++

																			res := &e2sm_mho_go.RicEventTriggerStyleList{
		RicEventTriggerStyleType: &e2sm_v2_ies.RicStyleType{
			Value: ricStyleType,
		},
		RicEventTriggerStyleName: &e2sm_v2_ies.RicStyleName{
			Value: ricStyleName,
		},
		RicEventTriggerFormatType: &e2sm_v2_ies.RicFormatType{
			Value: ricFormatType,
		},
	}

	if err := res.Validate(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-RanFunction-Description.go:43
		_go_fuzz_dep_.CoverTab[192440]++
																				return nil, fmt.Errorf("CreateRicEventTriggerStyleItem(): error validationg E2SmPDU %s", err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-RanFunction-Description.go:44
		// _ = "end of CoverTab[192440]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-RanFunction-Description.go:45
		_go_fuzz_dep_.CoverTab[192441]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-RanFunction-Description.go:45
		// _ = "end of CoverTab[192441]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-RanFunction-Description.go:45
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-RanFunction-Description.go:45
	// _ = "end of CoverTab[192438]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-RanFunction-Description.go:45
	_go_fuzz_dep_.CoverTab[192439]++

																			return res, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-RanFunction-Description.go:47
	// _ = "end of CoverTab[192439]"
}

func CreateRicReportStyleItem(ricStyleType int32, ricStyleName string, indHdrFormatType int32, indMsgFormatType int32) (*e2sm_mho_go.RicReportStyleList, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-RanFunction-Description.go:50
	_go_fuzz_dep_.CoverTab[192442]++

																			res := &e2sm_mho_go.RicReportStyleList{
		RicReportStyleType: &e2sm_v2_ies.RicStyleType{
			Value: ricStyleType,
		},
		RicReportStyleName: &e2sm_v2_ies.RicStyleName{
			Value: ricStyleName,
		},
		RicIndicationHeaderFormatType: &e2sm_v2_ies.RicFormatType{
			Value: indHdrFormatType,
		},
		RicIndicationMessageFormatType: &e2sm_v2_ies.RicFormatType{
			Value: indMsgFormatType,
		},
	}

	if err := res.Validate(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-RanFunction-Description.go:67
		_go_fuzz_dep_.CoverTab[192444]++
																				return nil, fmt.Errorf("CreateRicReportStyleItem(): error validationg E2SmPDU %s", err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-RanFunction-Description.go:68
		// _ = "end of CoverTab[192444]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-RanFunction-Description.go:69
		_go_fuzz_dep_.CoverTab[192445]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-RanFunction-Description.go:69
		// _ = "end of CoverTab[192445]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-RanFunction-Description.go:69
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-RanFunction-Description.go:69
	// _ = "end of CoverTab[192442]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-RanFunction-Description.go:69
	_go_fuzz_dep_.CoverTab[192443]++

																			return res, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-RanFunction-Description.go:71
	// _ = "end of CoverTab[192443]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-RanFunction-Description.go:72
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-RanFunction-Description.go:72
var _ = _go_fuzz_dep_.CoverTab
