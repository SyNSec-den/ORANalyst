// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:5
package handover

//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:5
)

import (
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/rrm-son-lib/pkg/model/device"
	"github.com/onosproject/rrm-son-lib/pkg/model/measurement"
	"sync"
	"time"
)

var log = logging.GetLogger("rrm-son-lib", "handover", "a3")

// NewA3HandoverHandler returns A3HandoverHandler object
func NewA3HandoverHandler() *A3HandoverHandler {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:18
	_go_fuzz_dep_.CoverTab[194395]++
														return &A3HandoverHandler{
		HandoverMap:	make(map[string]A3HandoverTimer),
		Chans: A3HandoverChannel{
			InputChan:	make(chan device.UE),
			OutputChan:	make(chan A3HandoverDecision),
		},
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:25
	// _ = "end of CoverTab[194395]"
}

// A3HandoverHandler is A3 handover handler
type A3HandoverHandler struct {
	HandoverMap	map[string]A3HandoverTimer
	Chans		A3HandoverChannel
	HandlerMutex	sync.RWMutex
}

// A3HandoverChannel struct has channels used in A3 handover handler
type A3HandoverChannel struct {
	InputChan	chan device.UE
	OutputChan	chan A3HandoverDecision
}

// Run starts A3 handover handler
func (h *A3HandoverHandler) Run() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:42
	_go_fuzz_dep_.CoverTab[194396]++
														for ue := range h.Chans.InputChan {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:43
		_go_fuzz_dep_.CoverTab[194397]++
															ttt := ue.GetSCell().GetTimeToTrigger()

															if ttt.GetValue().(int) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:46
			_go_fuzz_dep_.CoverTab[194398]++
																h.runWithZeroTTT(ue)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:47
			// _ = "end of CoverTab[194398]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:48
			_go_fuzz_dep_.CoverTab[194399]++
																h.runWithNonZeroTTT(ue, ttt.GetValue().(int))
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:49
			// _ = "end of CoverTab[194399]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:50
		// _ = "end of CoverTab[194397]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:51
	// _ = "end of CoverTab[194396]"
}

func (h *A3HandoverHandler) runWithZeroTTT(ue device.UE) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:54
	_go_fuzz_dep_.CoverTab[194400]++
														hoDecision := NewA3HandoverDecision(ue, h.getTargetCell(ue))
														h.Chans.OutputChan <- hoDecision
														log.Debugf("Handover - %v", hoDecision)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:57
	// _ = "end of CoverTab[194400]"
}

func (h *A3HandoverHandler) runWithNonZeroTTT(ue device.UE, ttt int) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:60
	_go_fuzz_dep_.CoverTab[194401]++

														h.HandlerMutex.Lock()
														if _, ok := h.HandoverMap[ue.GetID().String()]; !ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:63
		_go_fuzz_dep_.CoverTab[194403]++
															h.HandoverMap[ue.GetID().String()] = NewA3HandoverTimer(ue)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:64
		_curRoutineNum176_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:64
		_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum176_)
															go h.a3HandoverTimerProc(ue, ttt)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:65
		// _ = "end of CoverTab[194403]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:66
		_go_fuzz_dep_.CoverTab[194404]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:66
		// _ = "end of CoverTab[194404]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:66
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:66
	// _ = "end of CoverTab[194401]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:66
	_go_fuzz_dep_.CoverTab[194402]++

														h.HandoverMap[ue.GetID().String()].TimerChan <- ue
														h.HandlerMutex.Unlock()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:69
	// _ = "end of CoverTab[194402]"
}

func (h *A3HandoverHandler) a3HandoverTimerProc(ue device.UE, ttt int) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:72
	_go_fuzz_dep_.CoverTab[194405]++
														startTime := time.Now()
														targetCellID := h.getTargetCell(ue).GetID().String()
														for {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:75
		_go_fuzz_dep_.CoverTab[194406]++
															select {
		case <-time.After(time.Duration(ttt) * time.Millisecond):
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:77
			_go_fuzz_dep_.CoverTab[194407]++

																h.HandlerMutex.Lock()
																delete(h.HandoverMap, ue.GetID().String())
																h.HandlerMutex.Unlock()
																return
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:82
			// _ = "end of CoverTab[194407]"
		case ue := <-h.HandoverMap[ue.GetID().String()].TimerChan:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:83
			_go_fuzz_dep_.CoverTab[194408]++
																tmpTime := time.Now()
																eTime := tmpTime.Sub(startTime).Milliseconds()
																tmpTargetCell := h.getTargetCell(ue)
																tmpTargetCellID := tmpTargetCell.GetID().String()

//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:90
			if tmpTargetCellID != targetCellID {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:90
				_go_fuzz_dep_.CoverTab[194410]++
																	startTime = time.Now()
																	targetCellID = tmpTargetCellID
																	continue
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:93
				// _ = "end of CoverTab[194410]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:94
				_go_fuzz_dep_.CoverTab[194411]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:94
				// _ = "end of CoverTab[194411]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:94
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:94
			// _ = "end of CoverTab[194408]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:94
			_go_fuzz_dep_.CoverTab[194409]++

//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:97
			if tmpTargetCellID == targetCellID && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:97
				_go_fuzz_dep_.CoverTab[194412]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:97
				return eTime >= int64(ttt)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:97
				// _ = "end of CoverTab[194412]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:97
			}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:97
				_go_fuzz_dep_.CoverTab[194413]++

																	hoDecision := NewA3HandoverDecision(ue, tmpTargetCell)
																	h.Chans.OutputChan <- hoDecision
																	log.Debugf("Handover - %v", hoDecision)
																	h.HandlerMutex.Lock()
																	delete(h.HandoverMap, ue.GetID().String())
																	h.HandlerMutex.Unlock()
																	return
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:105
				// _ = "end of CoverTab[194413]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:106
				_go_fuzz_dep_.CoverTab[194414]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:106
				// _ = "end of CoverTab[194414]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:106
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:106
			// _ = "end of CoverTab[194409]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:107
		// _ = "end of CoverTab[194406]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:108
	// _ = "end of CoverTab[194405]"
}

func (h *A3HandoverHandler) getTargetCell(ue device.UE) device.Cell {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:111
	_go_fuzz_dep_.CoverTab[194415]++
														var targetCell device.Cell
														var bestRSRP measurement.RSRP
														flag := false
														for _, cscell := range ue.GetCSCells() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:115
		_go_fuzz_dep_.CoverTab[194417]++
															tmpRSRP := ue.GetMeasurements()[cscell.GetID().String()].GetMeasurement().(measurement.RSRP)
															if !flag {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:117
			_go_fuzz_dep_.CoverTab[194419]++
																targetCell = cscell
																bestRSRP = tmpRSRP
																flag = true
																continue
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:121
			// _ = "end of CoverTab[194419]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:122
			_go_fuzz_dep_.CoverTab[194420]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:122
			// _ = "end of CoverTab[194420]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:122
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:122
		// _ = "end of CoverTab[194417]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:122
		_go_fuzz_dep_.CoverTab[194418]++

															if tmpRSRP > bestRSRP {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:124
			_go_fuzz_dep_.CoverTab[194421]++
																targetCell = cscell
																bestRSRP = tmpRSRP
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:126
			// _ = "end of CoverTab[194421]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:127
			_go_fuzz_dep_.CoverTab[194422]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:127
			// _ = "end of CoverTab[194422]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:127
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:127
		// _ = "end of CoverTab[194418]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:128
	// _ = "end of CoverTab[194415]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:128
	_go_fuzz_dep_.CoverTab[194416]++
														return targetCell
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:129
	// _ = "end of CoverTab[194416]"
}

// NewA3HandoverTimer returns A3HandoverTimer object
func NewA3HandoverTimer(ue device.UE) A3HandoverTimer {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:133
	_go_fuzz_dep_.CoverTab[194423]++
														return A3HandoverTimer{
		UE:		ue,
		TimerChan:	make(chan device.UE),
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:137
	// _ = "end of CoverTab[194423]"
}

// A3HandoverTimer struct is for A3 handover timer
type A3HandoverTimer struct {
	UE		device.UE
	TimerChan	chan device.UE
}

// NewA3HandoverDecision returns A3HandoverDecision object
func NewA3HandoverDecision(ue device.UE, targetCell device.Cell) A3HandoverDecision {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:147
	_go_fuzz_dep_.CoverTab[194424]++
														return A3HandoverDecision{
		UE:		ue,
		ServingCell:	ue.GetSCell(),
		TargetCell:	targetCell,
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:152
	// _ = "end of CoverTab[194424]"
}

// A3HandoverDecision struct has A3 handover decision information
type A3HandoverDecision struct {
	UE		device.UE
	ServingCell	device.Cell
	TargetCell	device.Cell
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:160
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/handover/a3handover.go:160
var _ = _go_fuzz_dep_.CoverTab
