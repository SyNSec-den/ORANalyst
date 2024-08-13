// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/measurement_eventtype.go:5
package measurement

//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/measurement_eventtype.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/measurement_eventtype.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/measurement_eventtype.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/measurement_eventtype.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/measurement_eventtype.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/measurement_eventtype.go:5
)

const (
	// EventA1 is the Measurement Event A1
	EventA1	MeasEventType	= iota
	// EventA2 is the Measurement Event A2
	EventA2
	// EventA3 is the Measurement Event A3
	EventA3
	// EventA4 is the Measurement Event A4
	EventA4
	// EventA5 is the Measurement Event A5
	EventA5
)

// MeasEventType is the type for Measurement Event
type MeasEventType int

var strListMeasEventType = []string{
	"EventA1",
	"EventA2",
	"EventA3",
	"EventA4",
	"EventA5",
}

// String returns value as string type
func (e *MeasEventType) String() string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/measurement_eventtype.go:32
	_go_fuzz_dep_.CoverTab[194319]++
																return strListMeasEventType[*e]
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/measurement_eventtype.go:33
	// _ = "end of CoverTab[194319]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/measurement_eventtype.go:34
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/measurement_eventtype.go:34
var _ = _go_fuzz_dep_.CoverTab
