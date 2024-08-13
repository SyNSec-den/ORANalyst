// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/ue.go:5
package device

//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/ue.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/ue.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/ue.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/ue.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/ue.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/ue.go:5
)

import (
	"github.com/onosproject/rrm-son-lib/pkg/model/id"
	"github.com/onosproject/rrm-son-lib/pkg/model/measurement"
)

// NewUE returns UE object
func NewUE(ueid id.ID, scell Cell, cscells []Cell) UE {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/ue.go:13
	_go_fuzz_dep_.CoverTab[194384]++
													return &UEImpl{
		UEID:		ueid,
		DeviceType:	DeviceUE,
		SCell:		scell,
		CSCells:	cscells,
		Measurements:	make(map[string]measurement.Measurement),
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/ue.go:20
	// _ = "end of CoverTab[194384]"
}

// UE is the interface for UE device
type UE interface {
	Device
	GetSCell() Cell
	GetCSCells() []Cell
	GetMeasurements() map[string]measurement.Measurement
	SetScell(Cell)
	SetCSCells([]Cell)
	SetMeasurements(map[string]measurement.Measurement)
}

// UEImpl is the struct for UE implementation
type UEImpl struct {
	UEID		id.ID
	DeviceType	Type
	SCell		Cell
	CSCells		[]Cell
	Measurements	map[string]measurement.Measurement
}

// GetID returns ID
func (u *UEImpl) GetID() id.ID {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/ue.go:44
	_go_fuzz_dep_.CoverTab[194385]++
													return u.UEID
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/ue.go:45
	// _ = "end of CoverTab[194385]"
}

// GetType returns device type
func (u *UEImpl) GetType() Type {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/ue.go:49
	_go_fuzz_dep_.CoverTab[194386]++
													return u.DeviceType
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/ue.go:50
	// _ = "end of CoverTab[194386]"
}

// SetID sets ID
func (u *UEImpl) SetID(i id.ID) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/ue.go:54
	_go_fuzz_dep_.CoverTab[194387]++
													u.UEID = i
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/ue.go:55
	// _ = "end of CoverTab[194387]"
}

// SetType sets device type
func (u *UEImpl) SetType(t Type) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/ue.go:59
	_go_fuzz_dep_.CoverTab[194388]++
													u.DeviceType = t
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/ue.go:60
	// _ = "end of CoverTab[194388]"
}

// GetSCell returns serving cell
func (u *UEImpl) GetSCell() Cell {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/ue.go:64
	_go_fuzz_dep_.CoverTab[194389]++
													return u.SCell
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/ue.go:65
	// _ = "end of CoverTab[194389]"
}

// GetCSCells returns the list of candidate serving cells
func (u *UEImpl) GetCSCells() []Cell {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/ue.go:69
	_go_fuzz_dep_.CoverTab[194390]++
													return u.CSCells
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/ue.go:70
	// _ = "end of CoverTab[194390]"
}

// GetMeasurements returns measurement map
func (u *UEImpl) GetMeasurements() map[string]measurement.Measurement {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/ue.go:74
	_go_fuzz_dep_.CoverTab[194391]++
													return u.Measurements
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/ue.go:75
	// _ = "end of CoverTab[194391]"
}

// SetScell sets serving cell
func (u *UEImpl) SetScell(cell Cell) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/ue.go:79
	_go_fuzz_dep_.CoverTab[194392]++
													u.SCell = cell
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/ue.go:80
	// _ = "end of CoverTab[194392]"
}

// SetCSCells sets the list of candidate serving cells
func (u *UEImpl) SetCSCells(cells []Cell) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/ue.go:84
	_go_fuzz_dep_.CoverTab[194393]++
													u.CSCells = cells
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/ue.go:85
	// _ = "end of CoverTab[194393]"
}

// SetMeasurements sets the measurement map
func (u *UEImpl) SetMeasurements(m map[string]measurement.Measurement) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/ue.go:89
	_go_fuzz_dep_.CoverTab[194394]++
													u.Measurements = m
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/ue.go:90
	// _ = "end of CoverTab[194394]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/ue.go:91
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/ue.go:91
var _ = _go_fuzz_dep_.CoverTab
