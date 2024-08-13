// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/server.go:5
// Package agent implements a gnmi server to mock a device with YANG models.
package agent

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/server.go:6
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/server.go:6
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/server.go:6
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/server.go:6
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/server.go:6
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/server.go:6
)

import (
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-lib-go/pkg/northbound"
	"github.com/onosproject/onos-ric-sdk-go/pkg/config/configurable"
	api "github.com/openconfig/gnmi/proto/gnmi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var log = logging.GetLogger("gnmi", "agent")

// NewService creates a new gnmi service
func NewService(configurable configurable.Configurable) GnmiService {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/server.go:20
	_go_fuzz_dep_.CoverTab[194177]++
														return newService(configurable)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/server.go:21
	// _ = "end of CoverTab[194177]"
}

// GnmiService :
type GnmiService interface {
	northbound.Service
}

func newService(configurable configurable.Configurable) GnmiService {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/server.go:29
	_go_fuzz_dep_.CoverTab[194178]++
														server := &Server{
		configurable: configurable,
	}

	return &Service{
		server: server,
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/server.go:36
	// _ = "end of CoverTab[194178]"
}

// Service is a Service implementation for gnmi service.
type Service struct {
	server *Server
}

// Register registers the Service with the gRPC server.
func (s *Service) Register(r *grpc.Server) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/server.go:45
	_go_fuzz_dep_.CoverTab[194179]++
														server := s.server
														api.RegisterGNMIServer(r, server)
														reflection.Register(r)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/server.go:48
	// _ = "end of CoverTab[194179]"
}

// GnmiService :
var _ GnmiService = &Service{}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/server.go:52
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/server.go:52
var _ = _go_fuzz_dep_.CoverTab
