//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:1
// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:1
// SPDX-License-Identifier: Apache-2.0
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:4
package pdubuilder

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:4
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:4
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:4
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:4
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:4
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:4
)

import (
	"fmt"
	e2sm_mho_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-mho-go"
	e2sm_v2_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-v2-ies"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
)

func CreateE2SmMhoIndicationMsgFormat1(ueID *e2sm_v2_ies.Ueid, measReport []*e2sm_mho_go.E2SmMhoMeasurementReportItem) (*e2sm_mho_go.E2SmMhoIndicationMessage, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:13
	_go_fuzz_dep_.CoverTab[192393]++

																			E2SmMhoPdu := e2sm_mho_go.E2SmMhoIndicationMessage{
		E2SmMhoIndicationMessage: &e2sm_mho_go.E2SmMhoIndicationMessage_IndicationMessageFormat1{
			IndicationMessageFormat1: &e2sm_mho_go.E2SmMhoIndicationMessageFormat1{
				UeId:		ueID,
				MeasReport:	measReport,
			},
		},
	}

	if err := E2SmMhoPdu.Validate(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:24
		_go_fuzz_dep_.CoverTab[192395]++
																				return nil, fmt.Errorf("CreateE2SmMhoIndicationMsgFormat1(): error validating E2SmPDU %s", err.Error())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:25
		// _ = "end of CoverTab[192395]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:26
		_go_fuzz_dep_.CoverTab[192396]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:26
		// _ = "end of CoverTab[192396]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:26
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:26
	// _ = "end of CoverTab[192393]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:26
	_go_fuzz_dep_.CoverTab[192394]++
																			return &E2SmMhoPdu, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:27
	// _ = "end of CoverTab[192394]"
}

func CreateE2SmMhoIndicationMsgFormat2(ueID *e2sm_v2_ies.Ueid, rrcStatus e2sm_mho_go.Rrcstatus) (*e2sm_mho_go.E2SmMhoIndicationMessage, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:30
	_go_fuzz_dep_.CoverTab[192397]++
																			E2SmMhoPdu := e2sm_mho_go.E2SmMhoIndicationMessage{
		E2SmMhoIndicationMessage: &e2sm_mho_go.E2SmMhoIndicationMessage_IndicationMessageFormat2{
			IndicationMessageFormat2: &e2sm_mho_go.E2SmMhoIndicationMessageFormat2{
				UeId:		ueID,
				RrcStatus:	rrcStatus,
			},
		},
	}

	if err := E2SmMhoPdu.Validate(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:40
		_go_fuzz_dep_.CoverTab[192399]++
																				return nil, fmt.Errorf("CreateE2SmMhoIndicationMsgFormat2(): error validating E2SmPDU %s", err.Error())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:41
		// _ = "end of CoverTab[192399]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:42
		_go_fuzz_dep_.CoverTab[192400]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:42
		// _ = "end of CoverTab[192400]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:42
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:42
	// _ = "end of CoverTab[192397]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:42
	_go_fuzz_dep_.CoverTab[192398]++
																			return &E2SmMhoPdu, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:43
	// _ = "end of CoverTab[192398]"
}

func CreateMeasurementRecordItem(cgi *e2sm_v2_ies.Cgi, rsrp *e2sm_mho_go.Rsrp) (*e2sm_mho_go.E2SmMhoMeasurementReportItem, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:46
	_go_fuzz_dep_.CoverTab[192401]++
																			res := &e2sm_mho_go.E2SmMhoMeasurementReportItem{
		Cgi:	cgi,
		Rsrp:	rsrp,
	}

	if err := res.Validate(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:52
		_go_fuzz_dep_.CoverTab[192403]++
																				return nil, fmt.Errorf("CreateMeasurementRecordItem(): error validationg E2SmPDU %s", err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:53
		// _ = "end of CoverTab[192403]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:54
		_go_fuzz_dep_.CoverTab[192404]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:54
		// _ = "end of CoverTab[192404]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:54
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:54
	// _ = "end of CoverTab[192401]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:54
	_go_fuzz_dep_.CoverTab[192402]++

																			return res, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:56
	// _ = "end of CoverTab[192402]"
}

func CreateCgiNrCGI(plmnID []byte, nrCGI *asn1.BitString) (*e2sm_v2_ies.Cgi, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:59
	_go_fuzz_dep_.CoverTab[192405]++

																			if len(plmnID) != 3 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:61
		_go_fuzz_dep_.CoverTab[192410]++
																				return nil, fmt.Errorf("CreateCellGlobalIDNrCGI(): PlmnID should contain only 3 bytes, got %v", len(plmnID))
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:62
		// _ = "end of CoverTab[192410]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:63
		_go_fuzz_dep_.CoverTab[192411]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:63
		// _ = "end of CoverTab[192411]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:63
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:63
	// _ = "end of CoverTab[192405]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:63
	_go_fuzz_dep_.CoverTab[192406]++

																			if nrCGI.Len != uint32(36) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:65
		_go_fuzz_dep_.CoverTab[192412]++
																				return nil, fmt.Errorf("CreateCellGlobalIDNrCGI(): EutraCellIdentity should be of length 36 bits, got %v", nrCGI.Len)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:66
		// _ = "end of CoverTab[192412]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:67
		_go_fuzz_dep_.CoverTab[192413]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:67
		// _ = "end of CoverTab[192413]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:67
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:67
	// _ = "end of CoverTab[192406]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:67
	_go_fuzz_dep_.CoverTab[192407]++
																			if len(nrCGI.Value) != 5 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:68
		_go_fuzz_dep_.CoverTab[192414]++
																				return nil, fmt.Errorf("CreateCellGlobalIDNrCGI(): EutraCellIdentity should be of length 36 bits (5 bytes), got %v bytes", len(nrCGI.Value))
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:69
		// _ = "end of CoverTab[192414]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:70
		_go_fuzz_dep_.CoverTab[192415]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:70
		// _ = "end of CoverTab[192415]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:70
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:70
	// _ = "end of CoverTab[192407]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:70
	_go_fuzz_dep_.CoverTab[192408]++

																			cgi := &e2sm_v2_ies.Cgi{
		Cgi: &e2sm_v2_ies.Cgi_NRCgi{
			NRCgi: &e2sm_v2_ies.NrCgi{
				PLmnidentity: &e2sm_v2_ies.PlmnIdentity{
					Value: plmnID,
				},
				NRcellIdentity: &e2sm_v2_ies.NrcellIdentity{
					Value: nrCGI,
				},
			},
		},
	}

	if err := cgi.Validate(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:85
		_go_fuzz_dep_.CoverTab[192416]++
																				return nil, fmt.Errorf("CreateCellGlobalIDNrCGI(): error validationg E2SmPDU %s", err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:86
		// _ = "end of CoverTab[192416]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:87
		_go_fuzz_dep_.CoverTab[192417]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:87
		// _ = "end of CoverTab[192417]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:87
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:87
	// _ = "end of CoverTab[192408]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:87
	_go_fuzz_dep_.CoverTab[192409]++

																			return cgi, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:89
	// _ = "end of CoverTab[192409]"
}

func CreateCgiEutraCGI(plmnID []byte, eutraCGI *asn1.BitString) (*e2sm_v2_ies.Cgi, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:92
	_go_fuzz_dep_.CoverTab[192418]++

																			if len(plmnID) != 3 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:94
		_go_fuzz_dep_.CoverTab[192423]++
																				return nil, fmt.Errorf("CreateCellGlobalIDEutraCGI(): PlmnID should contain only 3 bytes, got %v", len(plmnID))
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:95
		// _ = "end of CoverTab[192423]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:96
		_go_fuzz_dep_.CoverTab[192424]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:96
		// _ = "end of CoverTab[192424]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:96
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:96
	// _ = "end of CoverTab[192418]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:96
	_go_fuzz_dep_.CoverTab[192419]++
																			if eutraCGI.Len != uint32(28) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:97
		_go_fuzz_dep_.CoverTab[192425]++
																				return nil, fmt.Errorf("CreateCellGlobalIDEutraCGI(): EutraCellIdentity should be of length 28 bits, got %v", eutraCGI.Len)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:98
		// _ = "end of CoverTab[192425]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:99
		_go_fuzz_dep_.CoverTab[192426]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:99
		// _ = "end of CoverTab[192426]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:99
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:99
	// _ = "end of CoverTab[192419]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:99
	_go_fuzz_dep_.CoverTab[192420]++
																			if len(eutraCGI.Value) != 4 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:100
		_go_fuzz_dep_.CoverTab[192427]++
																				return nil, fmt.Errorf("CreateCellGlobalIDEutraCGI(): EutraCellIdentity should be of length 28 bits (4 bytes), got %v bytes", len(eutraCGI.Value))
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:101
		// _ = "end of CoverTab[192427]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:102
		_go_fuzz_dep_.CoverTab[192428]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:102
		// _ = "end of CoverTab[192428]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:102
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:102
	// _ = "end of CoverTab[192420]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:102
	_go_fuzz_dep_.CoverTab[192421]++

																			cgi := &e2sm_v2_ies.Cgi{
		Cgi: &e2sm_v2_ies.Cgi_EUtraCgi{
			EUtraCgi: &e2sm_v2_ies.EutraCgi{
				PLmnidentity: &e2sm_v2_ies.PlmnIdentity{
					Value: plmnID,
				},
				EUtracellIdentity: &e2sm_v2_ies.EutracellIdentity{
					Value: eutraCGI,
				},
			},
		},
	}

	if err := cgi.Validate(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:117
		_go_fuzz_dep_.CoverTab[192429]++
																				return nil, fmt.Errorf("CreateCellGlobalIDEutraCGI(): error validationg E2SmPDU %s", err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:118
		// _ = "end of CoverTab[192429]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:119
		_go_fuzz_dep_.CoverTab[192430]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:119
		// _ = "end of CoverTab[192430]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:119
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:119
	// _ = "end of CoverTab[192421]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:119
	_go_fuzz_dep_.CoverTab[192422]++

																			return cgi, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:121
	// _ = "end of CoverTab[192422]"
}

func CreateRrcStatusConnected() e2sm_mho_go.Rrcstatus {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:124
	_go_fuzz_dep_.CoverTab[192431]++
																			return e2sm_mho_go.Rrcstatus_RRCSTATUS_CONNECTED
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:125
	// _ = "end of CoverTab[192431]"
}

func CreateRrcStatusInactive() e2sm_mho_go.Rrcstatus {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:128
	_go_fuzz_dep_.CoverTab[192432]++
																			return e2sm_mho_go.Rrcstatus_RRCSTATUS_INACTIVE
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:129
	// _ = "end of CoverTab[192432]"
}

func CreateRrcStatusIdle() e2sm_mho_go.Rrcstatus {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:132
	_go_fuzz_dep_.CoverTab[192433]++
																			return e2sm_mho_go.Rrcstatus_RRCSTATUS_IDLE
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:133
	// _ = "end of CoverTab[192433]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:134
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Indication-Message.go:134
var _ = _go_fuzz_dep_.CoverTab
