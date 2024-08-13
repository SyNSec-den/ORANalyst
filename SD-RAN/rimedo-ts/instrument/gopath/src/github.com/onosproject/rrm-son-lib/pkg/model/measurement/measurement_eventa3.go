// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/measurement_eventa3.go:5
package measurement

//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/measurement_eventa3.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/measurement_eventa3.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/measurement_eventa3.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/measurement_eventa3.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/measurement_eventa3.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/measurement_eventa3.go:5
)

import "github.com/onosproject/rrm-son-lib/pkg/model/id"

// RSRP is the type for RSRP
type RSRP float64

// NewMeasEventA3 returns MeasurementEventA3 object
func NewMeasEventA3(ecgi id.ID, rsrp RSRP) Measurement {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/measurement_eventa3.go:13
	_go_fuzz_dep_.CoverTab[194312]++
																return &MeasEventA3{
		ECGI:		ecgi,
		RSRP:		rsrp,
		MeasEventType:	EventA3,
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/measurement_eventa3.go:18
	// _ = "end of CoverTab[194312]"
}

// MeasEventA3 is the struct for measurement event A3
type MeasEventA3 struct {
	ECGI		id.ID
	RSRP		RSRP
	MeasEventType	MeasEventType
}

// GetMeasurementEventType returns measurement event a3
func (m *MeasEventA3) GetMeasurementEventType() MeasEventType {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/measurement_eventa3.go:29
	_go_fuzz_dep_.CoverTab[194313]++
																return m.MeasEventType
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/measurement_eventa3.go:30
	// _ = "end of CoverTab[194313]"
}

// GetCellID returns cell ID
func (m *MeasEventA3) GetCellID() id.ID {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/measurement_eventa3.go:34
	_go_fuzz_dep_.CoverTab[194314]++
																return m.ECGI
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/measurement_eventa3.go:35
	// _ = "end of CoverTab[194314]"
}

// GetMeasurement returns RSRP
func (m *MeasEventA3) GetMeasurement() interface{} {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/measurement_eventa3.go:39
	_go_fuzz_dep_.CoverTab[194315]++
																return m.RSRP
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/measurement_eventa3.go:40
	// _ = "end of CoverTab[194315]"
}

// SetMeasurementEventType sets measurement event type
func (m *MeasEventA3) SetMeasurementEventType(eventType MeasEventType) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/measurement_eventa3.go:44
	_go_fuzz_dep_.CoverTab[194316]++
																m.MeasEventType = eventType
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/measurement_eventa3.go:45
	// _ = "end of CoverTab[194316]"
}

// SetCellID sets cell ID
func (m *MeasEventA3) SetCellID(i id.ID) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/measurement_eventa3.go:49
	_go_fuzz_dep_.CoverTab[194317]++
																m.ECGI = i
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/measurement_eventa3.go:50
	// _ = "end of CoverTab[194317]"
}

// SetMeasurement sets measurement
func (m *MeasEventA3) SetMeasurement(i interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/measurement_eventa3.go:54
	_go_fuzz_dep_.CoverTab[194318]++
																m.RSRP = i.(RSRP)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/measurement_eventa3.go:55
	// _ = "end of CoverTab[194318]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/measurement_eventa3.go:56
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/measurement_eventa3.go:56
var _ = _go_fuzz_dep_.CoverTab
