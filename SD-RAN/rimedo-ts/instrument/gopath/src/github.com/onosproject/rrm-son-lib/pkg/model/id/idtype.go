// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/id/idtype.go:5
package id

//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/id/idtype.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/id/idtype.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/id/idtype.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/id/idtype.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/id/idtype.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/id/idtype.go:5
)

// Type is the ID type
type Type int

const (
	// TypeIDECGI is the ID type for ECGI
	TypeIDECGI	Type	= iota
	// TypeIDUEID is the ID type for UEID
	TypeIDUEID
)

var typeList = []string{
	"TypeIDECGI",
	"TypeIDUEID",
}

// String is the function to return type name
func (t *Type) String() string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/id/idtype.go:23
	_go_fuzz_dep_.CoverTab[194307]++
													return typeList[*t]
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/id/idtype.go:24
	// _ = "end of CoverTab[194307]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/id/idtype.go:25
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/rrm-son-lib@v0.0.2/pkg/model/id/idtype.go:25
var _ = _go_fuzz_dep_.CoverTab
