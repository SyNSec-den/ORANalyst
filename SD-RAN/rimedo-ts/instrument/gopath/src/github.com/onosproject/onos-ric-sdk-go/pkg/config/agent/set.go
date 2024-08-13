// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/set.go:5
// Package agent implements a gnmi server
package agent

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/set.go:6
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/set.go:6
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/set.go:6
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/set.go:6
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/set.go:6
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/set.go:6
)

import (
	"github.com/onosproject/onos-ric-sdk-go/pkg/config/configurable"
	pb "github.com/openconfig/gnmi/proto/gnmi"
	"golang.org/x/net/context"
)

// Set implements the Set RPC in gNMI spec.
func (s *Server) Set(ctx context.Context, req *pb.SetRequest) (*pb.SetResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/set.go:15
	_go_fuzz_dep_.CoverTab[194180]++
														log.Debugf("Processing Set Request:%+v", req)
														s.mu.Lock()
														defer s.mu.Unlock()

														setReq := configurable.SetRequest{
		DeletePaths:	req.GetDelete(),
		UpdatePaths:	req.GetUpdate(),
		ReplacePaths:	req.GetReplace(),
		Prefix:		req.GetPrefix(),
	}
	resp, err := s.configurable.Set(setReq)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/set.go:27
		_go_fuzz_dep_.CoverTab[194182]++
															return &pb.SetResponse{}, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/set.go:28
		// _ = "end of CoverTab[194182]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/set.go:29
		_go_fuzz_dep_.CoverTab[194183]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/set.go:29
		// _ = "end of CoverTab[194183]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/set.go:29
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/set.go:29
	// _ = "end of CoverTab[194180]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/set.go:29
	_go_fuzz_dep_.CoverTab[194181]++

														setResponse := &pb.SetResponse{
		Prefix:		req.GetPrefix(),
		Response:	resp.Results,
	}

														return setResponse, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/set.go:36
	// _ = "end of CoverTab[194181]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/set.go:37
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/agent/set.go:37
var _ = _go_fuzz_dep_.CoverTab
