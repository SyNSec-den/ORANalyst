// SPDX-FileCopyrightText: 2019-present Open Networking Foundation <info@opennetworking.org>
// SPDX-FileCopyrightText: 2019-present Rimedo Labs
//
// SPDX-License-Identifier: Apache-2.0
// Copy from https://github.com/woojoong88/onos-kpimon/tree/sample-a1t-xapp/pkg/northbound/a1

//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1ei-service.go:7
package a1

//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1ei-service.go:7
import (
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1ei-service.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1ei-service.go:7
)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1ei-service.go:7
import (
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1ei-service.go:7
	_atomic_ "sync/atomic"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1ei-service.go:7
)

import (
	"context"

	a1tapi "github.com/onosproject/onos-api/go/onos/a1t/a1"
	"github.com/onosproject/onos-lib-go/pkg/logging/service"
	"google.golang.org/grpc"
)

func NewA1EIService() service.Service {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1ei-service.go:17
	_go_fuzz_dep_.CoverTab[190652]++
											log.Debugf("A1EI service created")
											return &A1EIService{}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1ei-service.go:19
	// _ = "end of CoverTab[190652]"
}

type A1EIService struct {
}

func (a *A1EIService) Register(s *grpc.Server) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1ei-service.go:25
	_go_fuzz_dep_.CoverTab[190653]++
											server := &A1EIServer{}
											a1tapi.RegisterEIServiceServer(s, server)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1ei-service.go:27
	// _ = "end of CoverTab[190653]"
}

type A1EIServer struct {
}

func (a *A1EIServer) EIQuery(server a1tapi.EIService_EIQueryServer) error {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1ei-service.go:33
	_go_fuzz_dep_.CoverTab[190654]++
											log.Debug("EIQuery stream established")
											ch := make(chan bool)
											<-ch
											return nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1ei-service.go:37
	// _ = "end of CoverTab[190654]"
}

func (a *A1EIServer) EIJobSetup(server a1tapi.EIService_EIJobSetupServer) error {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1ei-service.go:40
	_go_fuzz_dep_.CoverTab[190655]++
											log.Debug("EIJobSetup stream established")
											ch := make(chan bool)
											<-ch
											return nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1ei-service.go:44
	// _ = "end of CoverTab[190655]"
}

func (a *A1EIServer) EIJobUpdate(server a1tapi.EIService_EIJobUpdateServer) error {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1ei-service.go:47
	_go_fuzz_dep_.CoverTab[190656]++
											log.Debug("EIJobUpdate stream established")
											ch := make(chan bool)
											<-ch
											return nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1ei-service.go:51
	// _ = "end of CoverTab[190656]"
}

func (a *A1EIServer) EIJobDelete(server a1tapi.EIService_EIJobDeleteServer) error {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1ei-service.go:54
	_go_fuzz_dep_.CoverTab[190657]++
											log.Debug("EIJobDelete stream established")
											ch := make(chan bool)
											<-ch
											return nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1ei-service.go:58
	// _ = "end of CoverTab[190657]"
}

func (a *A1EIServer) EIJobStatusQuery(server a1tapi.EIService_EIJobStatusQueryServer) error {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1ei-service.go:61
	_go_fuzz_dep_.CoverTab[190658]++
											log.Debug("EIJobStatusQuery stream established")
											ch := make(chan bool)
											<-ch
											return nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1ei-service.go:65
	// _ = "end of CoverTab[190658]"
}

func (a *A1EIServer) EIJobStatusNotify(ctx context.Context, message *a1tapi.EIStatusMessage) (*a1tapi.EIAckMessage, error) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1ei-service.go:68
	_go_fuzz_dep_.CoverTab[190659]++
											log.Debug("EIJobStatusNotify called %v", message)
											return nil, nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1ei-service.go:70
	// _ = "end of CoverTab[190659]"
}

func (a *A1EIServer) EIJobResultDelivery(ctx context.Context, message *a1tapi.EIResultMessage) (*a1tapi.EIAckMessage, error) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1ei-service.go:73
	_go_fuzz_dep_.CoverTab[190660]++
											log.Debug("EIJobResultDelivery called %v", message)
											return nil, nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1ei-service.go:75
	// _ = "end of CoverTab[190660]"
}

//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1ei-service.go:76
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1ei-service.go:76
var _ = _go_fuzz_dep_.CoverTab
