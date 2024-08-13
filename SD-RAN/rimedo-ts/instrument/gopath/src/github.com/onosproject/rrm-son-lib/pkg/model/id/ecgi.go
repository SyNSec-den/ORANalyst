// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/id/ecgi.go:5
package id

//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/id/ecgi.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/id/ecgi.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/id/ecgi.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/id/ecgi.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/id/ecgi.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/id/ecgi.go:5
)

import "fmt"

// NewECGI returns ECGI object
func NewECGI(e uint64) ID {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/id/ecgi.go:10
	_go_fuzz_dep_.CoverTab[194303]++
													id := ECGI(e)
													return &id
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/id/ecgi.go:12
	// _ = "end of CoverTab[194303]"
}

// ECGI is the type for the ID ECGI
type ECGI uint64

// String returns ECGI string type
func (E *ECGI) String() string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/id/ecgi.go:19
	_go_fuzz_dep_.CoverTab[194304]++
													return fmt.Sprintf("%d", *E)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/id/ecgi.go:20
	// _ = "end of CoverTab[194304]"
}

// GetType returns ID type
func (E *ECGI) GetType() Type {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/id/ecgi.go:24
	_go_fuzz_dep_.CoverTab[194305]++
													return TypeIDECGI
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/id/ecgi.go:25
	// _ = "end of CoverTab[194305]"
}

// GetID returns ID as interface type
func (E *ECGI) GetID() interface{} {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/id/ecgi.go:29
	_go_fuzz_dep_.CoverTab[194306]++
													return *E
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/id/ecgi.go:30
	// _ = "end of CoverTab[194306]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/id/ecgi.go:31
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/id/ecgi.go:31
var _ = _go_fuzz_dep_.CoverTab
