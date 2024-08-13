//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:1
// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:1
// SPDX-License-Identifier: Apache-2.0
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:4
package pdubuilder

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:4
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:4
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:4
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:4
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:4
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:4
)

import (
	"fmt"
	e2sm_mho_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-mho-go"
	e2sm_v2_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-v2-ies"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
)

func CreateE2SmMhoControlMessage(servingCgi *e2sm_v2_ies.Cgi, uedID *e2sm_v2_ies.Ueid, targetCgi *e2sm_v2_ies.Cgi) (*e2sm_mho_go.E2SmMhoControlMessage, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:13
	_go_fuzz_dep_.CoverTab[192351]++

																		e2smMhoMsgFormat1 := e2sm_mho_go.E2SmMhoControlMessageFormat1{
		ServingCgi:	servingCgi,
		UedId:		uedID,
		TargetCgi:	targetCgi,
	}

	e2smMhoPdu := e2sm_mho_go.E2SmMhoControlMessage{
		E2SmMhoControlMessage: &e2sm_mho_go.E2SmMhoControlMessage_ControlMessageFormat1{
			ControlMessageFormat1: &e2smMhoMsgFormat1,
		},
	}

	if err := e2smMhoPdu.Validate(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:27
		_go_fuzz_dep_.CoverTab[192353]++
																			return nil, fmt.Errorf("CreateE2SmMhoControlMessage(): error validating E2SmPDU %s", err.Error())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:28
		// _ = "end of CoverTab[192353]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:29
		_go_fuzz_dep_.CoverTab[192354]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:29
		// _ = "end of CoverTab[192354]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:29
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:29
	// _ = "end of CoverTab[192351]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:29
	_go_fuzz_dep_.CoverTab[192352]++
																		return &e2smMhoPdu, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:30
	// _ = "end of CoverTab[192352]"
}

func CreateUeIDGNb(amf int64, plmnID []byte, amfRegionID []byte, amfSetID []byte, amfPointer []byte) (*e2sm_v2_ies.Ueid, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:33
	_go_fuzz_dep_.CoverTab[192355]++

																		if len(plmnID) != 3 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:35
		_go_fuzz_dep_.CoverTab[192363]++
																			return nil, fmt.Errorf("CreateUeIDGNb() PlmnID should contain only 3 bytes, got %v", len(plmnID))
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:36
		// _ = "end of CoverTab[192363]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:37
		_go_fuzz_dep_.CoverTab[192364]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:37
		// _ = "end of CoverTab[192364]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:37
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:37
	// _ = "end of CoverTab[192355]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:37
	_go_fuzz_dep_.CoverTab[192356]++
																		if len(amfRegionID) != 1 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:38
		_go_fuzz_dep_.CoverTab[192365]++
																			return nil, fmt.Errorf("CreateUeIDGNb() AMfRegionID should contain only 1 byte, got %v", len(amfRegionID))
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:39
		// _ = "end of CoverTab[192365]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:40
		_go_fuzz_dep_.CoverTab[192366]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:40
		// _ = "end of CoverTab[192366]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:40
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:40
	// _ = "end of CoverTab[192356]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:40
	_go_fuzz_dep_.CoverTab[192357]++
																		if len(amfSetID) != 2 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:41
		_go_fuzz_dep_.CoverTab[192367]++
																			return nil, fmt.Errorf("CreateUeIDGNb() AMfSetID should contain only 2 bytes, got %v", len(amfSetID))
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:42
		// _ = "end of CoverTab[192367]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:43
		_go_fuzz_dep_.CoverTab[192368]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:43
		// _ = "end of CoverTab[192368]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:43
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:43
	// _ = "end of CoverTab[192357]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:43
	_go_fuzz_dep_.CoverTab[192358]++
																		if amfSetID[1]&0x3F > 0 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:44
		_go_fuzz_dep_.CoverTab[192369]++
																			return nil, fmt.Errorf("CreateUeIDGNb() AMfSetID should contain only 10 bits, i.e., last 6 bits of last byte should be trailing zeros, got %v", amfSetID[1])
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:45
		// _ = "end of CoverTab[192369]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:46
		_go_fuzz_dep_.CoverTab[192370]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:46
		// _ = "end of CoverTab[192370]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:46
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:46
	// _ = "end of CoverTab[192358]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:46
	_go_fuzz_dep_.CoverTab[192359]++
																		if len(amfPointer) != 1 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:47
		_go_fuzz_dep_.CoverTab[192371]++
																			return nil, fmt.Errorf("CreateUeIDGNb() AMfPointer should contain only 1 byte, got %v", len(amfPointer))
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:48
		// _ = "end of CoverTab[192371]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:49
		_go_fuzz_dep_.CoverTab[192372]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:49
		// _ = "end of CoverTab[192372]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:49
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:49
	// _ = "end of CoverTab[192359]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:49
	_go_fuzz_dep_.CoverTab[192360]++
																		if amfPointer[0]&0x03 > 0 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:50
		_go_fuzz_dep_.CoverTab[192373]++
																			return nil, fmt.Errorf("CreateUeIDGNb() AMfSetID should contain only 6 bits, i.e., last 2 bits should be trailing zeros, got %v", amfPointer)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:51
		// _ = "end of CoverTab[192373]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:52
		_go_fuzz_dep_.CoverTab[192374]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:52
		// _ = "end of CoverTab[192374]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:52
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:52
	// _ = "end of CoverTab[192360]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:52
	_go_fuzz_dep_.CoverTab[192361]++

																		ueID := &e2sm_v2_ies.Ueid{
		Ueid: &e2sm_v2_ies.Ueid_GNbUeid{
			GNbUeid: &e2sm_v2_ies.UeidGnb{
				AmfUeNgapId: &e2sm_v2_ies.AmfUeNgapId{
					Value: amf,
				},
				Guami: &e2sm_v2_ies.Guami{
					PLmnidentity: &e2sm_v2_ies.PlmnIdentity{
						Value: plmnID,
					},
					AMfregionId: &e2sm_v2_ies.AmfregionId{
						Value: &asn1.BitString{
							Value:	amfRegionID,
							Len:	8,
						},
					},
					AMfsetId: &e2sm_v2_ies.AmfsetId{
						Value: &asn1.BitString{
							Value:	amfSetID,
							Len:	10,
						},
					},
					AMfpointer: &e2sm_v2_ies.Amfpointer{
						Value: &asn1.BitString{
							Value:	amfPointer,
							Len:	6,
						},
					},
				},
			},
		},
	}

	if err := ueID.Validate(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:87
		_go_fuzz_dep_.CoverTab[192375]++
																			return nil, fmt.Errorf("CreateUeIDGNb() validation of UeID failed with %v", err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:88
		// _ = "end of CoverTab[192375]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:89
		_go_fuzz_dep_.CoverTab[192376]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:89
		// _ = "end of CoverTab[192376]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:89
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:89
	// _ = "end of CoverTab[192361]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:89
	_go_fuzz_dep_.CoverTab[192362]++

																		return ueID, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:91
	// _ = "end of CoverTab[192362]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:92
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Message.go:92
var _ = _go_fuzz_dep_.CoverTab
