// Copyright 2021-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:15
package v1beta1

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:15
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:15
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:15
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:15
)

import (
	"context"
	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-lib-go/pkg/northbound"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
)

var log = logging.GetLogger("onos", "proxy", "e2", "v1beta1")

const e2NodeIDHeader = "e2-node-id"

// NewProxyService creates a new E2T control and subscription proxy service
func NewProxyService(clientConn *grpc.ClientConn) northbound.Service {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:32
	_go_fuzz_dep_.CoverTab[190489]++
													return &SubscriptionService{
		conn: clientConn,
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:35
	// _ = "end of CoverTab[190489]"
}

// SubscriptionService is a Service implementation for E2 Subscription service.
type SubscriptionService struct {
	northbound.Service
	conn	*grpc.ClientConn
}

// Register registers the SubscriptionService with the gRPC server.
func (s SubscriptionService) Register(r *grpc.Server) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:45
	_go_fuzz_dep_.CoverTab[190490]++
													server := &ProxyServer{
		conn: s.conn,
	}
													e2api.RegisterSubscriptionServiceServer(r, server)
													e2api.RegisterControlServiceServer(r, server)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:50
	// _ = "end of CoverTab[190490]"
}

// ProxyServer implements the gRPC service for E2 Subscription related functions.
type ProxyServer struct {
	conn *grpc.ClientConn
}

func (s *ProxyServer) Control(ctx context.Context, request *e2api.ControlRequest) (*e2api.ControlResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:58
	_go_fuzz_dep_.CoverTab[190491]++
													log.Debugf("ControlRequest %+v", request)
													client := e2api.NewControlServiceClient(s.conn)
													ctx = metadata.AppendToOutgoingContext(ctx, e2NodeIDHeader, string(request.Headers.E2NodeID))
													response, err := client.Control(ctx, request)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:63
		_go_fuzz_dep_.CoverTab[190493]++
														log.Warnf("ControlRequest %+v error: %s", request, err)
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:65
		// _ = "end of CoverTab[190493]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:66
		_go_fuzz_dep_.CoverTab[190494]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:66
		// _ = "end of CoverTab[190494]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:66
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:66
	// _ = "end of CoverTab[190491]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:66
	_go_fuzz_dep_.CoverTab[190492]++
													log.Debugf("ControlResponse %+v", response)
													return response, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:68
	// _ = "end of CoverTab[190492]"
}

func (s *ProxyServer) Subscribe(request *e2api.SubscribeRequest, server e2api.SubscriptionService_SubscribeServer) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:71
	_go_fuzz_dep_.CoverTab[190495]++
													log.Debugf("SubscribeRequest %+v", request)
													client := e2api.NewSubscriptionServiceClient(s.conn)
													ctx := metadata.AppendToOutgoingContext(server.Context(), e2NodeIDHeader, string(request.Headers.E2NodeID))
													clientStream, err := client.Subscribe(ctx, request)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:76
		_go_fuzz_dep_.CoverTab[190497]++
														log.Warnf("SubscribeRequest %+v error: %s", request, err)
														return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:78
		// _ = "end of CoverTab[190497]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:79
		_go_fuzz_dep_.CoverTab[190498]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:79
		// _ = "end of CoverTab[190498]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:79
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:79
	// _ = "end of CoverTab[190495]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:79
	_go_fuzz_dep_.CoverTab[190496]++

													for {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:81
		_go_fuzz_dep_.CoverTab[190499]++
														response, err := clientStream.Recv()
														if err == io.EOF {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:83
			_go_fuzz_dep_.CoverTab[190502]++
															return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:84
			// _ = "end of CoverTab[190502]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:85
			_go_fuzz_dep_.CoverTab[190503]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:85
			// _ = "end of CoverTab[190503]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:85
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:85
		// _ = "end of CoverTab[190499]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:85
		_go_fuzz_dep_.CoverTab[190500]++
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:86
			_go_fuzz_dep_.CoverTab[190504]++
															log.Warnf("SubscribeRequest %+v error: %s", request, err)
															return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:88
			// _ = "end of CoverTab[190504]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:89
			_go_fuzz_dep_.CoverTab[190505]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:89
			// _ = "end of CoverTab[190505]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:89
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:89
		// _ = "end of CoverTab[190500]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:89
		_go_fuzz_dep_.CoverTab[190501]++
														log.Debugf("SubscribeResponse %+v", response)
														err = server.Send(response)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:92
			_go_fuzz_dep_.CoverTab[190506]++
															log.Warnf("SubscribeResponse %+v error: %s", response, err)
															return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:94
			// _ = "end of CoverTab[190506]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:95
			_go_fuzz_dep_.CoverTab[190507]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:95
			// _ = "end of CoverTab[190507]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:95
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:95
		// _ = "end of CoverTab[190501]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:96
	// _ = "end of CoverTab[190496]"
}

func (s *ProxyServer) Unsubscribe(ctx context.Context, request *e2api.UnsubscribeRequest) (*e2api.UnsubscribeResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:99
	_go_fuzz_dep_.CoverTab[190508]++
													log.Debugf("UnsubscribeRequest %+v", request)
													client := e2api.NewSubscriptionServiceClient(s.conn)
													ctx = metadata.AppendToOutgoingContext(ctx, e2NodeIDHeader, string(request.Headers.E2NodeID))
													response, err := client.Unsubscribe(ctx, request)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:104
		_go_fuzz_dep_.CoverTab[190510]++
														log.Warnf("UnsubscribeRequest %+v error: %s", request, err)
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:106
		// _ = "end of CoverTab[190510]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:107
		_go_fuzz_dep_.CoverTab[190511]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:107
		// _ = "end of CoverTab[190511]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:107
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:107
	// _ = "end of CoverTab[190508]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:107
	_go_fuzz_dep_.CoverTab[190509]++
													log.Debugf("UnsubscribeResponse %+v", response)
													return response, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:109
	// _ = "end of CoverTab[190509]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:110
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/proxy.go:110
var _ = _go_fuzz_dep_.CoverTab
