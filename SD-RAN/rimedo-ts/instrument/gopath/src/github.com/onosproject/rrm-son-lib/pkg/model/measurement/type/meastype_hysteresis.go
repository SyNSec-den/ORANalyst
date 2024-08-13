// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_hysteresis.go:5
package meastype

//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_hysteresis.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_hysteresis.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_hysteresis.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_hysteresis.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_hysteresis.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_hysteresis.go:5
)

import "fmt"

// NewHysteresisRange returns NewHysteresisRange object
func NewHysteresisRange(hysteresis HysteresisRange) MeasType {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_hysteresis.go:10
	_go_fuzz_dep_.CoverTab[194332]++
																	if hysteresis < 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_hysteresis.go:11
		_go_fuzz_dep_.CoverTab[194334]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_hysteresis.go:11
		return hysteresis > 30
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_hysteresis.go:11
		// _ = "end of CoverTab[194334]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_hysteresis.go:11
	}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_hysteresis.go:11
		_go_fuzz_dep_.CoverTab[194335]++
																		return &DefaultHysteresisRange
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_hysteresis.go:12
		// _ = "end of CoverTab[194335]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_hysteresis.go:13
		_go_fuzz_dep_.CoverTab[194336]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_hysteresis.go:13
		// _ = "end of CoverTab[194336]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_hysteresis.go:13
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_hysteresis.go:13
	// _ = "end of CoverTab[194332]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_hysteresis.go:13
	_go_fuzz_dep_.CoverTab[194333]++
																	return &hysteresis
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_hysteresis.go:14
	// _ = "end of CoverTab[194333]"
}

// HysteresisRange is the type for hysteresis 0 dB to 30 dB integer
type HysteresisRange int

// DefaultHysteresisRange is the default value of Hysteresis
var DefaultHysteresisRange = HysteresisRange(0)

// SetValue is the set function for value
func (h *HysteresisRange) SetValue(i interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_hysteresis.go:24
	_go_fuzz_dep_.CoverTab[194337]++
																	if i.(HysteresisRange) < 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_hysteresis.go:25
		_go_fuzz_dep_.CoverTab[194339]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_hysteresis.go:25
		return i.(HysteresisRange) > 30
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_hysteresis.go:25
		// _ = "end of CoverTab[194339]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_hysteresis.go:25
	}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_hysteresis.go:25
		_go_fuzz_dep_.CoverTab[194340]++
																		*h = DefaultHysteresisRange
																		return fmt.Errorf("Hysteresis should be in the range from 0 to 30; received %v - set to default 0 dB", i.(HysteresisRange))
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_hysteresis.go:27
		// _ = "end of CoverTab[194340]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_hysteresis.go:28
		_go_fuzz_dep_.CoverTab[194341]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_hysteresis.go:28
		// _ = "end of CoverTab[194341]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_hysteresis.go:28
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_hysteresis.go:28
	// _ = "end of CoverTab[194337]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_hysteresis.go:28
	_go_fuzz_dep_.CoverTab[194338]++
																	*h = i.(HysteresisRange)
																	return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_hysteresis.go:30
	// _ = "end of CoverTab[194338]"
}

// GetValue is the get function for value as interface type
func (h *HysteresisRange) GetValue() interface{} {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_hysteresis.go:34
	_go_fuzz_dep_.CoverTab[194342]++
																	return *h
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_hysteresis.go:35
	// _ = "end of CoverTab[194342]"
}

// String returns value as string type
func (h *HysteresisRange) String() string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_hysteresis.go:39
	_go_fuzz_dep_.CoverTab[194343]++
																	return fmt.Sprintf("%d", *h)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_hysteresis.go:40
	// _ = "end of CoverTab[194343]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_hysteresis.go:41
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_hysteresis.go:41
var _ = _go_fuzz_dep_.CoverTab
