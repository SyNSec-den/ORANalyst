// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/devicetype.go:5
package device

//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/devicetype.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/devicetype.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/devicetype.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/devicetype.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/devicetype.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/devicetype.go:5
)

const (
	// DeviceUE is the type for UE
	DeviceUE	Type	= iota
	// DeviceCell is the type for Cell
	DeviceCell
)

// Type is the type of device
type Type int

var strListDeviceType = []string{
	"DeviceUE",
	"DeviceCell",
}

// String returns value as string type
func (t *Type) String() string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/devicetype.go:23
	_go_fuzz_dep_.CoverTab[194383]++
														return strListDeviceType[*t]
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/devicetype.go:24
	// _ = "end of CoverTab[194383]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/devicetype.go:25
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/device/devicetype.go:25
var _ = _go_fuzz_dep_.CoverTab
