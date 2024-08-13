// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/cell.go:5
package device

//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/cell.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/cell.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/cell.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/cell.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/cell.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/cell.go:5
)

import (
	"github.com/onosproject/rrm-son-lib/pkg/model/id"
	meastype "github.com/onosproject/rrm-son-lib/pkg/model/measurement/type"
)

// NewCell returns cell object
func NewCell(ecgi id.ID, a3offset meastype.A3OffsetRange, hyst meastype.HysteresisRange,
	cellOffset meastype.QOffsetRange, freqOffset meastype.QOffsetRange, ttt meastype.TimeToTriggerRange) Cell {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/cell.go:14
	_go_fuzz_dep_.CoverTab[194368]++
														return &CellImpl{
		ECGI:			ecgi,
		DeviceType:		DeviceCell,
		A3Offset:		a3offset,
		Hysteresis:		hyst,
		CellIndividualOffset:	cellOffset,
		FrequencyOffset:	freqOffset,
		TimeToTrigger:		ttt,
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/cell.go:23
	// _ = "end of CoverTab[194368]"
}

// Cell is the interface for Cell device
type Cell interface {
	Device
	GetA3Offset() meastype.A3OffsetRange
	GetHysteresis() meastype.HysteresisRange
	GetCellIndividualOffset() meastype.QOffsetRange
	GetFrequencyOffset() meastype.QOffsetRange
	GetTimeToTrigger() meastype.TimeToTriggerRange
	SetA3Offset(offsetRange meastype.A3OffsetRange)
	SetHysteresis(hysteresisRange meastype.HysteresisRange)
	SetCellIndividualOffset(offsetRange meastype.QOffsetRange)
	SetFrequencyOffset(offsetRange meastype.QOffsetRange)
	SetTimeToTrigger(triggerRange meastype.TimeToTriggerRange)
}

// CellImpl is the struct for Cell implementation
type CellImpl struct {
	ECGI			id.ID
	DeviceType		Type
	A3Offset		meastype.A3OffsetRange
	Hysteresis		meastype.HysteresisRange
	CellIndividualOffset	meastype.QOffsetRange
	FrequencyOffset		meastype.QOffsetRange
	TimeToTrigger		meastype.TimeToTriggerRange
}

// GetID returns ID
func (c *CellImpl) GetID() id.ID {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/cell.go:53
	_go_fuzz_dep_.CoverTab[194369]++
														return c.ECGI
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/cell.go:54
	// _ = "end of CoverTab[194369]"
}

// GetType returns device type
func (c *CellImpl) GetType() Type {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/cell.go:58
	_go_fuzz_dep_.CoverTab[194370]++
														return c.DeviceType
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/cell.go:59
	// _ = "end of CoverTab[194370]"
}

// SetID sets ID
func (c *CellImpl) SetID(i id.ID) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/cell.go:63
	_go_fuzz_dep_.CoverTab[194371]++
														c.ECGI = i
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/cell.go:64
	// _ = "end of CoverTab[194371]"
}

// SetType sets device type
func (c *CellImpl) SetType(deviceType Type) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/cell.go:68
	_go_fuzz_dep_.CoverTab[194372]++
														c.DeviceType = deviceType
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/cell.go:69
	// _ = "end of CoverTab[194372]"
}

// GetA3Offset returns a3 offset
func (c *CellImpl) GetA3Offset() meastype.A3OffsetRange {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/cell.go:73
	_go_fuzz_dep_.CoverTab[194373]++
														return c.A3Offset
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/cell.go:74
	// _ = "end of CoverTab[194373]"
}

// GetHysteresis returns hysteresis
func (c *CellImpl) GetHysteresis() meastype.HysteresisRange {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/cell.go:78
	_go_fuzz_dep_.CoverTab[194374]++
														return c.Hysteresis
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/cell.go:79
	// _ = "end of CoverTab[194374]"
}

// GetCellIndividualOffset returns cell individual offset
func (c *CellImpl) GetCellIndividualOffset() meastype.QOffsetRange {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/cell.go:83
	_go_fuzz_dep_.CoverTab[194375]++
														return c.CellIndividualOffset
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/cell.go:84
	// _ = "end of CoverTab[194375]"
}

// GetFrequencyOffset returns frqeuency offset
func (c *CellImpl) GetFrequencyOffset() meastype.QOffsetRange {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/cell.go:88
	_go_fuzz_dep_.CoverTab[194376]++
														return c.FrequencyOffset
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/cell.go:89
	// _ = "end of CoverTab[194376]"
}

// GetTimeToTrigger returns time to trigger
func (c *CellImpl) GetTimeToTrigger() meastype.TimeToTriggerRange {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/cell.go:93
	_go_fuzz_dep_.CoverTab[194377]++
														return c.TimeToTrigger
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/cell.go:94
	// _ = "end of CoverTab[194377]"
}

// SetA3Offset sets a3 offset
func (c *CellImpl) SetA3Offset(offsetRange meastype.A3OffsetRange) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/cell.go:98
	_go_fuzz_dep_.CoverTab[194378]++
														c.A3Offset = offsetRange
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/cell.go:99
	// _ = "end of CoverTab[194378]"
}

// SetHysteresis sets hysteresis
func (c *CellImpl) SetHysteresis(hysteresisRange meastype.HysteresisRange) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/cell.go:103
	_go_fuzz_dep_.CoverTab[194379]++
														c.Hysteresis = hysteresisRange
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/cell.go:104
	// _ = "end of CoverTab[194379]"
}

// SetCellIndividualOffset sets cell individual offset
func (c *CellImpl) SetCellIndividualOffset(offsetRange meastype.QOffsetRange) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/cell.go:108
	_go_fuzz_dep_.CoverTab[194380]++
														c.CellIndividualOffset = offsetRange
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/cell.go:109
	// _ = "end of CoverTab[194380]"
}

// SetFrequencyOffset sets frequency offset
func (c *CellImpl) SetFrequencyOffset(offsetRange meastype.QOffsetRange) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/cell.go:113
	_go_fuzz_dep_.CoverTab[194381]++
														c.FrequencyOffset = offsetRange
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/cell.go:114
	// _ = "end of CoverTab[194381]"
}

// SetTimeToTrigger sets time to trigger
func (c *CellImpl) SetTimeToTrigger(triggerRange meastype.TimeToTriggerRange) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/cell.go:118
	_go_fuzz_dep_.CoverTab[194382]++
														c.TimeToTrigger = triggerRange
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/cell.go:119
	// _ = "end of CoverTab[194382]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/cell.go:120
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/cell.go:120
var _ = _go_fuzz_dep_.CoverTab
