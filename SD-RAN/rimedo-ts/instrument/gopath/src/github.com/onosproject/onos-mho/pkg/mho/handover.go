// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/handover.go:5
package mho

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/handover.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/handover.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/handover.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/handover.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/handover.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/handover.go:5
)

import (
	"bytes"
	"context"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-mho-go"
	appConfig "github.com/onosproject/onos-mho/pkg/config"
	"github.com/onosproject/rrm-son-lib/pkg/handover"
	measurement2 "github.com/onosproject/rrm-son-lib/pkg/measurement"
	"github.com/onosproject/rrm-son-lib/pkg/model/device"
	rrmid "github.com/onosproject/rrm-son-lib/pkg/model/id"
	"github.com/onosproject/rrm-son-lib/pkg/model/measurement"
	meastype "github.com/onosproject/rrm-son-lib/pkg/model/measurement/type"
)

// HandOverController is the handover controller
type HandOverController struct {
	A3OffsetRange		uint64
	HysteresisRange		uint64
	CellIndividualOffset	uint64
	FrequencyOffset		uint64
	TimeToTrigger		uint64
	UeChan			chan device.UE
	A3Handler		*measurement2.MeasEventA3Handler
	HandoverHandler		*handover.A3HandoverHandler
}

func NewHandOverController(cfg appConfig.Config) *HandOverController {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/handover.go:32
	_go_fuzz_dep_.CoverTab[194490]++
													log.Info("Init HandOverController")
													return &HandOverController{
		A3OffsetRange:		cfg.GetA3OffsetRange(),
		HysteresisRange:	cfg.GetHysteresisRange(),
		CellIndividualOffset:	cfg.GetCellIndividualOffset(),
		FrequencyOffset:	cfg.GetFrequencyOffset(),
		TimeToTrigger:		cfg.GetTimeToTrigger(),
		UeChan:			make(chan device.UE),
		A3Handler:		measurement2.NewMeasEventA3Handler(),
		HandoverHandler:	handover.NewA3HandoverHandler(),
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/handover.go:43
	// _ = "end of CoverTab[194490]"
}

func (h *HandOverController) Run() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/handover.go:46
	_go_fuzz_dep_.CoverTab[194491]++
													log.Info("Start A3Handler")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/handover.go:47
	_curRoutineNum178_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/handover.go:47
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum178_)
													go h.A3Handler.Run()

													log.Info("Start HandoverHandler")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/handover.go:50
	_curRoutineNum179_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/handover.go:50
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum179_)
													go h.HandoverHandler.Run()

													log.Info("Start forwarding A3Handler events to HandoverHandler")
													for ue := range h.A3Handler.Chans.OutputChan {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/handover.go:54
		_go_fuzz_dep_.CoverTab[194492]++
														h.HandoverHandler.Chans.InputChan <- ue
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/handover.go:55
		// _ = "end of CoverTab[194492]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/handover.go:56
	// _ = "end of CoverTab[194491]"

}

func (h *HandOverController) Input(ctx context.Context, header *e2sm_mho.E2SmMhoIndicationHeaderFormat1, message *e2sm_mho.E2SmMhoIndicationMessageFormat1) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/handover.go:60
	_go_fuzz_dep_.CoverTab[194493]++
													ueID, err := GetUeID(message.GetUeId())
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/handover.go:62
		_go_fuzz_dep_.CoverTab[194496]++
														log.Errorf("handlePeriodicReport() couldn't extract UeID: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/handover.go:63
		// _ = "end of CoverTab[194496]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/handover.go:64
		_go_fuzz_dep_.CoverTab[194497]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/handover.go:64
		// _ = "end of CoverTab[194497]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/handover.go:64
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/handover.go:64
	// _ = "end of CoverTab[194493]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/handover.go:64
	_go_fuzz_dep_.CoverTab[194494]++

													ecgiSCell := rrmid.NewECGI(BitStringToUint64(header.GetCgi().GetNRCgi().GetNRcellIdentity().GetValue().GetValue(), int(header.GetCgi().GetNRCgi().GetNRcellIdentity().GetValue().GetLen())))
													scell := device.NewCell(
		ecgiSCell,
		meastype.A3OffsetRange(h.A3OffsetRange),
		meastype.HysteresisRange(h.HysteresisRange),
		meastype.QOffsetRange(h.CellIndividualOffset),
		meastype.QOffsetRange(h.FrequencyOffset),
		meastype.TimeToTriggerRange(h.TimeToTrigger))
	cscellList := make([]device.Cell, 0)

	ue := device.NewUE(rrmid.NewUEID(uint64(ueID), uint32(0),
		BitStringToUint64(header.GetCgi().GetNRCgi().GetNRcellIdentity().GetValue().GetValue(), int(header.GetCgi().GetNRCgi().GetNRcellIdentity().GetValue().GetLen()))),
		scell, nil)

	for _, measReport := range message.MeasReport {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/handover.go:80
		_go_fuzz_dep_.CoverTab[194498]++
														if bytes.Equal(measReport.GetCgi().GetNRCgi().GetNRcellIdentity().GetValue().GetValue(), header.GetCgi().GetNRCgi().GetNRcellIdentity().GetValue().GetValue()) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/handover.go:81
			_go_fuzz_dep_.CoverTab[194499]++
															ue.GetMeasurements()[ecgiSCell.String()] = measurement.NewMeasEventA3(ecgiSCell, measurement.RSRP(measReport.GetRsrp().GetValue()))
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/handover.go:82
			// _ = "end of CoverTab[194499]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/handover.go:83
			_go_fuzz_dep_.CoverTab[194500]++
															ecgiCSCell := rrmid.NewECGI(BitStringToUint64(measReport.GetCgi().GetNRCgi().GetNRcellIdentity().GetValue().GetValue(), int(measReport.GetCgi().GetNRCgi().GetNRcellIdentity().GetValue().GetLen())))
															cscell := device.NewCell(
				ecgiCSCell,
				meastype.A3OffsetRange(h.A3OffsetRange),
				meastype.HysteresisRange(h.HysteresisRange),
				meastype.QOffsetRange(h.CellIndividualOffset),
				meastype.QOffsetRange(h.FrequencyOffset),
				meastype.TimeToTriggerRange(h.TimeToTrigger))
															cscellList = append(cscellList, cscell)
															ue.GetMeasurements()[ecgiCSCell.String()] = measurement.NewMeasEventA3(ecgiCSCell, measurement.RSRP(measReport.GetRsrp().GetValue()))
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/handover.go:93
			// _ = "end of CoverTab[194500]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/handover.go:94
		// _ = "end of CoverTab[194498]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/handover.go:95
	// _ = "end of CoverTab[194494]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/handover.go:95
	_go_fuzz_dep_.CoverTab[194495]++
													ue.SetCSCells(cscellList)

													h.A3Handler.Chans.InputChan <- ue
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/handover.go:98
	// _ = "end of CoverTab[194495]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/handover.go:99
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/handover.go:99
var _ = _go_fuzz_dep_.CoverTab
