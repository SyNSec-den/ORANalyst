// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_timetotrigger.go:5
package meastype

//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_timetotrigger.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_timetotrigger.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_timetotrigger.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_timetotrigger.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_timetotrigger.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_timetotrigger.go:5
)

import "fmt"

const (
	// TTT0ms is TimeToTrigger value 0ms
	TTT0ms	TimeToTriggerRange	= iota
	// TTT40ms is TimeToTrigger value 0ms
	TTT40ms
	// TTT64ms is TimeToTrigger value 0ms
	TTT64ms
	// TTT80ms is TimeToTrigger value 0ms
	TTT80ms
	// TTT100ms is TimeToTrigger value 0ms
	TTT100ms
	// TTT128ms is TimeToTrigger value 0ms
	TTT128ms
	// TTT160ms is TimeToTrigger value 0ms
	TTT160ms
	// TTT256ms is TimeToTrigger value 0ms
	TTT256ms
	// TTT320ms is TimeToTrigger value 0ms
	TTT320ms
	// TTT480ms is TimeToTrigger value 0ms
	TTT480ms
	// TTT512ms is TimeToTrigger value 0ms
	TTT512ms
	// TTT640ms is TimeToTrigger value 0ms
	TTT640ms
	// TTT1024ms is TimeToTrigger value 0ms
	TTT1024ms
	// TTT1280ms is TimeToTrigger value 0ms
	TTT1280ms
	// TTT2560ms is TimeToTrigger value 0ms
	TTT2560ms
	// TTT5120ms is TimeToTrigger value 0ms
	TTT5120ms
)

// NewTimeToTriggerRange returns TimeToTriggerRange object
func NewTimeToTriggerRange(ttt TimeToTriggerRange) MeasType {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_timetotrigger.go:45
	_go_fuzz_dep_.CoverTab[194356]++
																	if ttt < TTT0ms || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_timetotrigger.go:46
		_go_fuzz_dep_.CoverTab[194358]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_timetotrigger.go:46
		return ttt > TTT5120ms
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_timetotrigger.go:46
		// _ = "end of CoverTab[194358]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_timetotrigger.go:46
	}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_timetotrigger.go:46
		_go_fuzz_dep_.CoverTab[194359]++
																		return &DefaultTimeToTriggerRange
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_timetotrigger.go:47
		// _ = "end of CoverTab[194359]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_timetotrigger.go:48
		_go_fuzz_dep_.CoverTab[194360]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_timetotrigger.go:48
		// _ = "end of CoverTab[194360]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_timetotrigger.go:48
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_timetotrigger.go:48
	// _ = "end of CoverTab[194356]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_timetotrigger.go:48
	_go_fuzz_dep_.CoverTab[194357]++
																	return &ttt
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_timetotrigger.go:49
	// _ = "end of CoverTab[194357]"
}

// TimeToTriggerRange is the type for TimeToTrigger
type TimeToTriggerRange int

// SetValue is the set function for value
func (t *TimeToTriggerRange) SetValue(i interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_timetotrigger.go:56
	_go_fuzz_dep_.CoverTab[194361]++
																	if i.(TimeToTriggerRange) < TTT0ms || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_timetotrigger.go:57
		_go_fuzz_dep_.CoverTab[194363]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_timetotrigger.go:57
		return i.(TimeToTriggerRange) > TTT5120ms
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_timetotrigger.go:57
		// _ = "end of CoverTab[194363]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_timetotrigger.go:57
	}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_timetotrigger.go:57
		_go_fuzz_dep_.CoverTab[194364]++
																		*t = DefaultTimeToTriggerRange
																		return fmt.Errorf("TimeToTrigger should be set in the range TTT_0MS to TTT_5120MS; received %v - set to default TTT_0MS", i.(TimeToTriggerRange))
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_timetotrigger.go:59
		// _ = "end of CoverTab[194364]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_timetotrigger.go:60
		_go_fuzz_dep_.CoverTab[194365]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_timetotrigger.go:60
		// _ = "end of CoverTab[194365]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_timetotrigger.go:60
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_timetotrigger.go:60
	// _ = "end of CoverTab[194361]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_timetotrigger.go:60
	_go_fuzz_dep_.CoverTab[194362]++
																	*t = i.(TimeToTriggerRange)
																	return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_timetotrigger.go:62
	// _ = "end of CoverTab[194362]"
}

// GetValue is the get function for value as interface type
func (t *TimeToTriggerRange) GetValue() interface{} {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_timetotrigger.go:66
	_go_fuzz_dep_.CoverTab[194366]++
																	return valueListTimeToTriggerRange[*t]
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_timetotrigger.go:67
	// _ = "end of CoverTab[194366]"
}

// String returns value as string type
func (t *TimeToTriggerRange) String() string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_timetotrigger.go:71
	_go_fuzz_dep_.CoverTab[194367]++
																	return strListTimeToTriggerRange[*t]
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_timetotrigger.go:72
	// _ = "end of CoverTab[194367]"
}

// DefaultTimeToTriggerRange is the default value of TimeToTrigger
var DefaultTimeToTriggerRange = TTT0ms

var strListTimeToTriggerRange = []string{
	"TTT0ms",
	"TTT40ms",
	"TTT64ms",
	"TTT80ms",
	"TTT100ms",
	"TTT128ms",
	"TTT160ms",
	"TTT256ms",
	"TTT320ms",
	"TTT480ms",
	"TTT512ms",
	"TTT640ms",
	"TTT1024ms",
	"TTT1280ms",
	"TTT2560ms",
	"TTT5120ms",
}

var valueListTimeToTriggerRange = []int{
	0,
	40,
	64,
	80,
	100,
	128,
	160,
	256,
	320,
	480,
	512,
	640,
	1024,
	1280,
	2560,
	5120,
}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_timetotrigger.go:114
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_timetotrigger.go:114
var _ = _go_fuzz_dep_.CoverTab
