// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_a3offsetrange.go:5
package meastype

//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_a3offsetrange.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_a3offsetrange.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_a3offsetrange.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_a3offsetrange.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_a3offsetrange.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_a3offsetrange.go:5
)

import "fmt"

// NewA3OffsetRange returns A3OffsetRange object
func NewA3OffsetRange(a3Offset A3OffsetRange) MeasType {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_a3offsetrange.go:10
	_go_fuzz_dep_.CoverTab[194320]++
																	if a3Offset < -30 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_a3offsetrange.go:11
		_go_fuzz_dep_.CoverTab[194322]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_a3offsetrange.go:11
		return a3Offset > 30
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_a3offsetrange.go:11
		// _ = "end of CoverTab[194322]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_a3offsetrange.go:11
	}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_a3offsetrange.go:11
		_go_fuzz_dep_.CoverTab[194323]++
																		return &DefaultA3OffsetRange
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_a3offsetrange.go:12
		// _ = "end of CoverTab[194323]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_a3offsetrange.go:13
		_go_fuzz_dep_.CoverTab[194324]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_a3offsetrange.go:13
		// _ = "end of CoverTab[194324]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_a3offsetrange.go:13
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_a3offsetrange.go:13
	// _ = "end of CoverTab[194320]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_a3offsetrange.go:13
	_go_fuzz_dep_.CoverTab[194321]++
																	return &a3Offset
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_a3offsetrange.go:14
	// _ = "end of CoverTab[194321]"
}

// A3OffsetRange is the type for A3 offset - -30 dB to 30 dB integer
type A3OffsetRange int

// DefaultA3OffsetRange is the default value of A3 offset
var DefaultA3OffsetRange = A3OffsetRange(0)

// SetValue is the set function for value
func (a *A3OffsetRange) SetValue(i interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_a3offsetrange.go:24
	_go_fuzz_dep_.CoverTab[194325]++
																	if i.(A3OffsetRange) < -30 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_a3offsetrange.go:25
		_go_fuzz_dep_.CoverTab[194327]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_a3offsetrange.go:25
		return i.(A3OffsetRange) > 30
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_a3offsetrange.go:25
		// _ = "end of CoverTab[194327]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_a3offsetrange.go:25
	}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_a3offsetrange.go:25
		_go_fuzz_dep_.CoverTab[194328]++
																		*a = DefaultA3OffsetRange
																		return fmt.Errorf("A3-Offset should be in the range from -30 to 30; received %v - set to default 0 dB", i.(A3OffsetRange))
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_a3offsetrange.go:27
		// _ = "end of CoverTab[194328]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_a3offsetrange.go:28
		_go_fuzz_dep_.CoverTab[194329]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_a3offsetrange.go:28
		// _ = "end of CoverTab[194329]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_a3offsetrange.go:28
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_a3offsetrange.go:28
	// _ = "end of CoverTab[194325]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_a3offsetrange.go:28
	_go_fuzz_dep_.CoverTab[194326]++
																	*a = i.(A3OffsetRange)
																	return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_a3offsetrange.go:30
	// _ = "end of CoverTab[194326]"
}

// GetValue is the get function for value as interface type
func (a *A3OffsetRange) GetValue() interface{} {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_a3offsetrange.go:34
	_go_fuzz_dep_.CoverTab[194330]++
																	return *a
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_a3offsetrange.go:35
	// _ = "end of CoverTab[194330]"
}

// String returns value as string type
func (a *A3OffsetRange) String() string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_a3offsetrange.go:39
	_go_fuzz_dep_.CoverTab[194331]++
																	return fmt.Sprintf("%d", *a)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_a3offsetrange.go:40
	// _ = "end of CoverTab[194331]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_a3offsetrange.go:41
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_a3offsetrange.go:41
var _ = _go_fuzz_dep_.CoverTab
