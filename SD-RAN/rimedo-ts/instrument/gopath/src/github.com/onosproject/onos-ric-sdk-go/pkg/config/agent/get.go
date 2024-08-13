// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/get.go:5
// Package agent implements a gnmi server to mock a device with YANG models.
package agent

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/get.go:6
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/get.go:6
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/get.go:6
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/get.go:6
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/get.go:6
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/get.go:6
)

import (
	"github.com/onosproject/onos-ric-sdk-go/pkg/config/configurable"
	pb "github.com/openconfig/gnmi/proto/gnmi"
	"golang.org/x/net/context"
)

// Get implements the Get RPC in gNMI spec.
func (s *Server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/get.go:15
	_go_fuzz_dep_.CoverTab[194173]++
														log.Debugf("Processing Get Request: %+v", req)
														s.mu.RLock()
														defer s.mu.RUnlock()
														getResponse, err := s.configurable.Get(configurable.GetRequest{
		Paths:		req.Path,
		EncodingType:	req.GetEncoding(),
		Prefix:		req.GetPrefix(),
		DataType:	req.Type.String(),
	})

	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/get.go:26
		_go_fuzz_dep_.CoverTab[194175]++
															return &pb.GetResponse{}, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/get.go:27
		// _ = "end of CoverTab[194175]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/get.go:28
		_go_fuzz_dep_.CoverTab[194176]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/get.go:28
		// _ = "end of CoverTab[194176]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/get.go:28
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/get.go:28
	// _ = "end of CoverTab[194173]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/get.go:28
	_go_fuzz_dep_.CoverTab[194174]++

														resp := &pb.GetResponse{Notification: getResponse.Notifications}
														return resp, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/get.go:31
	// _ = "end of CoverTab[194174]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/get.go:32
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/get.go:32
var _ = _go_fuzz_dep_.CoverTab
