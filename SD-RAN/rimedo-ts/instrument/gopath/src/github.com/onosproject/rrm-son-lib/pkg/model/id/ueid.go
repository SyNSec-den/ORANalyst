// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/id/ueid.go:5
package id

//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/id/ueid.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/id/ueid.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/id/ueid.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/id/ueid.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/id/ueid.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/id/ueid.go:5
)

import "fmt"

// IMSI is the type for UE's ID
type IMSI uint64

// CRNTI is the other type for UE's ID
type CRNTI uint32

// NewUEID returns UEID object
func NewUEID(imsi uint64, crnti uint32, ecgi uint64) ID {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/id/ueid.go:16
	_go_fuzz_dep_.CoverTab[194308]++
													return &UEID{
		IMSI:	IMSI(imsi),
		CRNTI:	CRNTI(crnti),
		ECGI:	ECGI(ecgi),
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/id/ueid.go:21
	// _ = "end of CoverTab[194308]"
}

// UEID is the struct for UE's ID
type UEID struct {
	IMSI	IMSI
	CRNTI	CRNTI
	ECGI	ECGI
}

// String returns UEID string type
func (U *UEID) String() string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/id/ueid.go:32
	_go_fuzz_dep_.CoverTab[194309]++
													return fmt.Sprintf("%d", *U)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/id/ueid.go:33
	// _ = "end of CoverTab[194309]"
}

// GetType returns ID type
func (U *UEID) GetType() Type {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/id/ueid.go:37
	_go_fuzz_dep_.CoverTab[194310]++
													return TypeIDUEID
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/id/ueid.go:38
	// _ = "end of CoverTab[194310]"
}

// GetID returns ID as interface type
func (U *UEID) GetID() interface{} {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/id/ueid.go:42
	_go_fuzz_dep_.CoverTab[194311]++
													return *U
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/id/ueid.go:43
	// _ = "end of CoverTab[194311]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/id/ueid.go:44
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/id/ueid.go:44
var _ = _go_fuzz_dep_.CoverTab
