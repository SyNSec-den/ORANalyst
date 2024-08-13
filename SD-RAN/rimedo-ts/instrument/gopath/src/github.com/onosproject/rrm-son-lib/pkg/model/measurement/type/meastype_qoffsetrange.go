// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_qoffsetrange.go:5
package meastype

//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_qoffsetrange.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_qoffsetrange.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_qoffsetrange.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_qoffsetrange.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_qoffsetrange.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_qoffsetrange.go:5
)

import "fmt"

const (
	// QOffsetMinus24dB is the Q-Offset value -24 dB
	QOffsetMinus24dB	QOffsetRange	= iota
	// QOffsetMinus22dB is the Q-Offset value -22 dB
	QOffsetMinus22dB
	// QOffsetMinus20dB is the Q-Offset value -20 dB
	QOffsetMinus20dB
	// QOffsetMinus18dB is the Q-Offset value -18 dB
	QOffsetMinus18dB
	// QOffsetMinus16dB is the Q-Offset value -16 dB
	QOffsetMinus16dB
	// QOffsetMinus14dB is the Q-Offset value -14 dB
	QOffsetMinus14dB
	// QOffsetMinus12dB is the Q-Offset value -12 dB
	QOffsetMinus12dB
	// QOffsetMinus10dB is the Q-Offset value -10 dB
	QOffsetMinus10dB
	// QOffsetMinus8dB is the Q-Offset value -8 dB
	QOffsetMinus8dB
	// QOffsetMinus6dB is the Q-Offset value -6 dB
	QOffsetMinus6dB
	// QOffsetMinus5dB is the Q-Offset value -5 dB
	QOffsetMinus5dB
	// QOffsetMinus4dB is the Q-Offset value -4 dB
	QOffsetMinus4dB
	// QOffsetMinus3dB is the Q-Offset value -3 dB
	QOffsetMinus3dB
	// QOffsetMinus2dB is the Q-Offset value -2 dB
	QOffsetMinus2dB
	// QOffsetMinus1dB is the Q-Offset value -1 dB
	QOffsetMinus1dB
	// QOffset0dB is the Q-Offset value 0 dB
	QOffset0dB
	// QOffset1dB is the Q-Offset value 1 dB
	QOffset1dB
	// QOffset2dB is the Q-Offset value 2 dB
	QOffset2dB
	// QOffset3dB is the Q-Offset value 3 dB
	QOffset3dB
	// QOffset4dB is the Q-Offset value 4 dB
	QOffset4dB
	// QOffset5dB is the Q-Offset value 5 dB
	QOffset5dB
	// QOffset6dB is the Q-Offset value 6 dB
	QOffset6dB
	// QOffset8dB is the Q-Offset value 8 dB
	QOffset8dB
	// QOffset10dB is the Q-Offset value 10 dB
	QOffset10dB
	// QOffset12dB is the Q-Offset value 12 dB
	QOffset12dB
	// QOffset14dB is the Q-Offset value 14 dB
	QOffset14dB
	// QOffset16dB is the Q-Offset value 16 dB
	QOffset16dB
	// QOffset18dB is the Q-Offset value 18 dB
	QOffset18dB
	// QOffset20dB is the Q-Offset value 20 dB
	QOffset20dB
	// QOffset22dB is the Q-Offset value 22 dB
	QOffset22dB
	// QOffset24dB is the Q-Offset value 24 dB
	QOffset24dB
)

// NewQOffsetRange returns NewQOffsetRange object
func NewQOffsetRange(qOffset QOffsetRange) MeasType {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_qoffsetrange.go:75
	_go_fuzz_dep_.CoverTab[194344]++
																	if qOffset < QOffsetMinus24dB || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_qoffsetrange.go:76
		_go_fuzz_dep_.CoverTab[194346]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_qoffsetrange.go:76
		return qOffset > QOffset24dB
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_qoffsetrange.go:76
		// _ = "end of CoverTab[194346]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_qoffsetrange.go:76
	}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_qoffsetrange.go:76
		_go_fuzz_dep_.CoverTab[194347]++
																		return &DefaultQOffsetRange
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_qoffsetrange.go:77
		// _ = "end of CoverTab[194347]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_qoffsetrange.go:78
		_go_fuzz_dep_.CoverTab[194348]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_qoffsetrange.go:78
		// _ = "end of CoverTab[194348]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_qoffsetrange.go:78
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_qoffsetrange.go:78
	// _ = "end of CoverTab[194344]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_qoffsetrange.go:78
	_go_fuzz_dep_.CoverTab[194345]++
																	return &qOffset
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_qoffsetrange.go:79
	// _ = "end of CoverTab[194345]"
}

// QOffsetRange is the type for Q-Offset range
type QOffsetRange int

// SetValue is the set function for value
func (q *QOffsetRange) SetValue(i interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_qoffsetrange.go:86
	_go_fuzz_dep_.CoverTab[194349]++
																	if i.(QOffsetRange) < QOffsetMinus24dB || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_qoffsetrange.go:87
		_go_fuzz_dep_.CoverTab[194351]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_qoffsetrange.go:87
		return i.(QOffsetRange) > QOffset24dB
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_qoffsetrange.go:87
		// _ = "end of CoverTab[194351]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_qoffsetrange.go:87
	}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_qoffsetrange.go:87
		_go_fuzz_dep_.CoverTab[194352]++
																		*q = DefaultQOffsetRange
																		return fmt.Errorf("Q-Offset should be set in the range QOFFSET_MINUS24DB to QOFFSET_24DB; received %v - set to default QOFFSET_0DB", i.(QOffsetRange))
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_qoffsetrange.go:89
		// _ = "end of CoverTab[194352]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_qoffsetrange.go:90
		_go_fuzz_dep_.CoverTab[194353]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_qoffsetrange.go:90
		// _ = "end of CoverTab[194353]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_qoffsetrange.go:90
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_qoffsetrange.go:90
	// _ = "end of CoverTab[194349]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_qoffsetrange.go:90
	_go_fuzz_dep_.CoverTab[194350]++
																	*q = i.(QOffsetRange)
																	return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_qoffsetrange.go:92
	// _ = "end of CoverTab[194350]"
}

// GetValue is the get function for value as interface type
func (q *QOffsetRange) GetValue() interface{} {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_qoffsetrange.go:96
	_go_fuzz_dep_.CoverTab[194354]++
																	return valueListQOffsetRange[*q]
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_qoffsetrange.go:97
	// _ = "end of CoverTab[194354]"
}

// String returns value as string type
func (q *QOffsetRange) String() string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_qoffsetrange.go:101
	_go_fuzz_dep_.CoverTab[194355]++
																	return strListQOffsetRange[*q]
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_qoffsetrange.go:102
	// _ = "end of CoverTab[194355]"
}

// DefaultQOffsetRange is the default value of Q-Offset range
var DefaultQOffsetRange = QOffset0dB

var strListQOffsetRange = []string{
	"QOffsetMinus24dB",
	"QOffsetMinus22dB",
	"QOffsetMinus20dB",
	"QOffsetMinus18dB",
	"QOffsetMinus16dB",
	"QOffsetMinus14dB",
	"QOffsetMinus12dB",
	"QOffsetMinus10dB",
	"QOffsetMinus8dB",
	"QOffsetMinus6dB",
	"QOffsetMinus5dB",
	"QOffsetMinus4dB",
	"QOffsetMinus3dB",
	"QOffsetMinus2dB",
	"QOffsetMinus1dB",
	"QOffset0dB",
	"QOffset1dB",
	"QOffset2dB",
	"QOffset3dB",
	"QOffset4dB",
	"QOffset5dB",
	"QOffset6dB",
	"QOffset8dB",
	"QOffset10dB",
	"QOffset12dB",
	"QOffset14dB",
	"QOffset16dB",
	"QOffset18dB",
	"QOffset20dB",
	"QOffset22dB",
	"QOffset24dB",
}

var valueListQOffsetRange = []int{
	-24,
	-22,
	-20,
	-18,
	-16,
	-14,
	-12,
	-10,
	-8,
	-6,
	-5,
	-4,
	-3,
	-2,
	-1,
	0,
	1,
	2,
	3,
	4,
	5,
	6,
	8,
	10,
	12,
	14,
	16,
	18,
	20,
	22,
	24,
}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_qoffsetrange.go:174
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/measurement/type/meastype_qoffsetrange.go:174
var _ = _go_fuzz_dep_.CoverTab
