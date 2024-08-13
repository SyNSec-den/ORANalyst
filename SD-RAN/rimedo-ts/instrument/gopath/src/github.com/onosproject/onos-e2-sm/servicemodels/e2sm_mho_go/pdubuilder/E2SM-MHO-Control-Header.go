//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Header.go:1
// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Header.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Header.go:1
// SPDX-License-Identifier: Apache-2.0
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Header.go:4
package pdubuilder

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Header.go:4
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Header.go:4
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Header.go:4
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Header.go:4
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Header.go:4
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Header.go:4
)

import (
	"fmt"
	e2sm_mho_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-mho-go"
)

func CreateE2SmMhoControlHeader(controlMessagePriority int32) (*e2sm_mho_go.E2SmMhoControlHeader, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Header.go:11
	_go_fuzz_dep_.CoverTab[192347]++

																		e2smMhoFormat1 := e2sm_mho_go.E2SmMhoControlHeaderFormat1{
		RcCommand:	e2sm_mho_go.MhoCommand_MHO_COMMAND_INITIATE_HANDOVER,
		RicControlMessagePriority: &e2sm_mho_go.RicControlMessagePriority{
			Value: controlMessagePriority,
		},
	}
	e2smMhoPdu := e2sm_mho_go.E2SmMhoControlHeader{
		E2SmMhoControlHeader: &e2sm_mho_go.E2SmMhoControlHeader_ControlHeaderFormat1{
			ControlHeaderFormat1: &e2smMhoFormat1,
		},
	}

	if err := e2smMhoPdu.Validate(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Header.go:25
		_go_fuzz_dep_.CoverTab[192349]++
																			return nil, fmt.Errorf("CreateE2SmMhoControlHeader(): error validating E2SmPDU %s", err.Error())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Header.go:26
		// _ = "end of CoverTab[192349]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Header.go:27
		_go_fuzz_dep_.CoverTab[192350]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Header.go:27
		// _ = "end of CoverTab[192350]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Header.go:27
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Header.go:27
	// _ = "end of CoverTab[192347]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Header.go:27
	_go_fuzz_dep_.CoverTab[192348]++
																		return &e2smMhoPdu, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Header.go:28
	// _ = "end of CoverTab[192348]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Header.go:29
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/pdubuilder/E2SM-MHO-Control-Header.go:29
var _ = _go_fuzz_dep_.CoverTab
