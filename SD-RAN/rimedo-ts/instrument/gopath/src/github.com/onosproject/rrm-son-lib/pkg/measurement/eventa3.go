// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:5
package measurement

//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:5
)

import (
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/rrm-son-lib/pkg/model/device"
	"github.com/onosproject/rrm-son-lib/pkg/model/measurement"
	meastype "github.com/onosproject/rrm-son-lib/pkg/model/measurement/type"
)

var log = logging.GetLogger("rrm-son-lib", "measurement", "eventa3")

// A3Status is the status if the UE is in the event a3 status
type A3Status bool

// NewMeasEventA3Handler returns MeasEventA3Handler object
func NewMeasEventA3Handler() *MeasEventA3Handler {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:20
	_go_fuzz_dep_.CoverTab[194425]++
														return &MeasEventA3Handler{
		EventMap:	make(map[string]MeasEventA3Obj),
		Chans: MeasEventA3Channel{
			InputChan:	make(chan device.UE),
			OutputChan:	make(chan device.UE),
		},
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:27
	// _ = "end of CoverTab[194425]"
}

// MeasEventA3Handler is Event a3 handler
type MeasEventA3Handler struct {
	Chans		MeasEventA3Channel
	EventMap	map[string]MeasEventA3Obj
}

// MeasEventA3Channel has all channels used in Event A3 handler
type MeasEventA3Channel struct {
	InputChan	chan device.UE
	OutputChan	chan device.UE
}

// Run starts handler
func (h *MeasEventA3Handler) Run() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:43
	_go_fuzz_dep_.CoverTab[194426]++
														for ue := range h.Chans.InputChan {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:44
		_go_fuzz_dep_.CoverTab[194427]++
															h.updateA3EventMap(ue)
															obj := h.EventMap[ue.GetID().String()]
															if obj.isInA3Event() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:47
			_go_fuzz_dep_.CoverTab[194428]++
																log.Debugf("UE %v is in A3 event - report through output channel: UE info - %v", ue.GetID().String(), ue)
																h.Chans.OutputChan <- ue
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:49
			// _ = "end of CoverTab[194428]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:50
			_go_fuzz_dep_.CoverTab[194429]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:50
			// _ = "end of CoverTab[194429]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:50
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:50
		// _ = "end of CoverTab[194427]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:51
	// _ = "end of CoverTab[194426]"
}

func (h *MeasEventA3Handler) updateA3EventMap(ue device.UE) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:54
	_go_fuzz_dep_.CoverTab[194430]++
														if _, ok := h.EventMap[ue.GetID().String()]; !ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:55
		_go_fuzz_dep_.CoverTab[194432]++
															h.EventMap[ue.GetID().String()] = NewMeasEventA3Obj(ue)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:56
		// _ = "end of CoverTab[194432]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:57
		_go_fuzz_dep_.CoverTab[194433]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:57
		// _ = "end of CoverTab[194433]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:57
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:57
	// _ = "end of CoverTab[194430]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:57
	_go_fuzz_dep_.CoverTab[194431]++

														obj := h.EventMap[ue.GetID().String()]
														log.Debugf("A3 Object [id: %v]: %v", ue.GetID().String(), ue)

//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:63
	tmpA3StatusMap := obj.A3StatusMap
	for _, cscell := range ue.GetCSCells() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:64
		_go_fuzz_dep_.CoverTab[194434]++
															if _, ok := tmpA3StatusMap[cscell.GetID().String()]; !ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:65
			_go_fuzz_dep_.CoverTab[194436]++
																obj.A3StatusMap[cscell.GetID().String()] = h.isEnteringA3Condition(ue, cscell)
																continue
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:67
			// _ = "end of CoverTab[194436]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:68
			_go_fuzz_dep_.CoverTab[194437]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:68
			// _ = "end of CoverTab[194437]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:68
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:68
		// _ = "end of CoverTab[194434]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:68
		_go_fuzz_dep_.CoverTab[194435]++

															if tmpA3StatusMap[cscell.GetID().String()] {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:70
			_go_fuzz_dep_.CoverTab[194438]++
																obj.A3StatusMap[cscell.GetID().String()] = !h.isLeavingA3Condition(ue, cscell)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:71
			// _ = "end of CoverTab[194438]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:72
			_go_fuzz_dep_.CoverTab[194439]++
																obj.A3StatusMap[cscell.GetID().String()] = h.isEnteringA3Condition(ue, cscell)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:73
			// _ = "end of CoverTab[194439]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:74
		// _ = "end of CoverTab[194435]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:75
	// _ = "end of CoverTab[194431]"
}

func (h *MeasEventA3Handler) isEnteringA3Condition(ue device.UE, cell device.Cell) A3Status {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:78
	_go_fuzz_dep_.CoverTab[194440]++
														mp := ue.GetMeasurements()[ue.GetSCell().GetID().String()].GetMeasurement().(measurement.RSRP)
														ofp := ue.GetSCell().GetFrequencyOffset()
														ocp := ue.GetSCell().GetCellIndividualOffset()
														a3offset := ue.GetSCell().GetA3Offset()
														hyst := ue.GetSCell().GetHysteresis()

														mn := ue.GetMeasurements()[cell.GetID().String()].GetMeasurement().(measurement.RSRP)
														ofn := cell.GetFrequencyOffset()
														ocn := cell.GetCellIndividualOffset()

														return float64(mn)+float64(ofn.GetValue().(int))+float64(ocn.GetValue().(int))-float64(hyst.GetValue().(meastype.HysteresisRange)) >
		float64(mp)+float64(ofp.GetValue().(int))+float64(ocp.GetValue().(int))+float64(a3offset.GetValue().(meastype.A3OffsetRange))
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:90
	// _ = "end of CoverTab[194440]"
}

func (h *MeasEventA3Handler) isLeavingA3Condition(ue device.UE, cell device.Cell) A3Status {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:93
	_go_fuzz_dep_.CoverTab[194441]++
														mp := ue.GetMeasurements()[ue.GetSCell().GetID().String()].GetMeasurement().(measurement.RSRP)
														ofp := ue.GetSCell().GetFrequencyOffset()
														ocp := ue.GetSCell().GetCellIndividualOffset()
														a3offset := ue.GetSCell().GetA3Offset()
														hyst := ue.GetSCell().GetHysteresis()

														mn := ue.GetMeasurements()[cell.GetID().String()].GetMeasurement().(measurement.RSRP)
														ofn := cell.GetFrequencyOffset()
														ocn := cell.GetCellIndividualOffset()

														return float64(mn)+float64(ofn.GetValue().(int))+float64(ocn.GetValue().(int))+float64(hyst.GetValue().(meastype.HysteresisRange)) <
		float64(mp)+float64(ofp.GetValue().(int))+float64(ocp.GetValue().(int))+float64(a3offset.GetValue().(meastype.A3OffsetRange))
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:105
	// _ = "end of CoverTab[194441]"
}

// NewMeasEventA3Obj returns MeasEventA3Obj
func NewMeasEventA3Obj(ue device.UE) MeasEventA3Obj {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:109
	_go_fuzz_dep_.CoverTab[194442]++
														a3StatusMap := make(map[string]A3Status)

														for _, cell := range ue.GetCSCells() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:112
		_go_fuzz_dep_.CoverTab[194444]++
															a3StatusMap[cell.GetID().String()] = false
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:113
		// _ = "end of CoverTab[194444]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:114
	// _ = "end of CoverTab[194442]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:114
	_go_fuzz_dep_.CoverTab[194443]++
														return MeasEventA3Obj{
		UE:		ue,
		A3StatusMap:	a3StatusMap,
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:118
	// _ = "end of CoverTab[194443]"
}

// MeasEventA3Obj is the struct for Event A3 record
type MeasEventA3Obj struct {
	UE		device.UE
	A3StatusMap	map[string]A3Status
}

func (o *MeasEventA3Obj) isInA3Event() A3Status {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:127
	_go_fuzz_dep_.CoverTab[194445]++
														for _, v := range o.A3StatusMap {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:128
		_go_fuzz_dep_.CoverTab[194447]++
															if v {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:129
			_go_fuzz_dep_.CoverTab[194448]++
																return true
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:130
			// _ = "end of CoverTab[194448]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:131
			_go_fuzz_dep_.CoverTab[194449]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:131
			// _ = "end of CoverTab[194449]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:131
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:131
		// _ = "end of CoverTab[194447]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:132
	// _ = "end of CoverTab[194445]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:132
	_go_fuzz_dep_.CoverTab[194446]++
														return false
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:133
	// _ = "end of CoverTab[194446]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:134
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/measurement/eventa3.go:134
var _ = _go_fuzz_dep_.CoverTab
