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

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:15
package balancer

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:15
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:15
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:15
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:15
)

import (
	"context"
	"fmt"
	"github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-lib-go/pkg/grpc/retry"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/attributes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/serviceconfig"
)

var log = logging.GetLogger("onos", "proxy", "e2", "v1beta1", "balancer")

const ResolverName = "e2"
const topoAddress = "onos-topo:5150"

func init() {
	resolver.Register(&ResolverBuilder{})
}

// ResolverBuilder :
type ResolverBuilder struct{}

// Scheme :
func (b *ResolverBuilder) Scheme() string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:43
	_go_fuzz_dep_.CoverTab[190524]++
															return ResolverName
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:44
	// _ = "end of CoverTab[190524]"
}

// Build :
func (b *ResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:48
	_go_fuzz_dep_.CoverTab[190525]++
															var dialOpts []grpc.DialOption
															if opts.DialCreds != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:50
		_go_fuzz_dep_.CoverTab[190529]++
																dialOpts = append(
			dialOpts,
			grpc.WithTransportCredentials(opts.DialCreds),
		)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:54
		// _ = "end of CoverTab[190529]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:55
		_go_fuzz_dep_.CoverTab[190530]++
																dialOpts = append(dialOpts, grpc.WithInsecure())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:56
		// _ = "end of CoverTab[190530]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:57
	// _ = "end of CoverTab[190525]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:57
	_go_fuzz_dep_.CoverTab[190526]++
															dialOpts = append(dialOpts, grpc.WithUnaryInterceptor(retry.RetryingUnaryClientInterceptor(retry.WithRetryOn(codes.Unavailable))))
															dialOpts = append(dialOpts, grpc.WithStreamInterceptor(retry.RetryingStreamClientInterceptor(retry.WithRetryOn(codes.Unavailable))))
															dialOpts = append(dialOpts, grpc.WithContextDialer(opts.Dialer))

															topoConn, err := grpc.Dial(topoAddress, dialOpts...)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:63
		_go_fuzz_dep_.CoverTab[190531]++
																return nil, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:64
		// _ = "end of CoverTab[190531]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:65
		_go_fuzz_dep_.CoverTab[190532]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:65
		// _ = "end of CoverTab[190532]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:65
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:65
	// _ = "end of CoverTab[190526]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:65
	_go_fuzz_dep_.CoverTab[190527]++

															serviceConfig := cc.ParseServiceConfig(
		fmt.Sprintf(`{"loadBalancingConfig":[{"%s":{}}]}`, ResolverName),
	)

	log.Infof("Built new resolver")

	resolver := &Resolver{
		clientConn:	cc,
		topoConn:	topoConn,
		serviceConfig:	serviceConfig,
		masterships:	make(map[topo.ID]topo.MastershipState),
		controls:	make(map[topo.ID]topo.ID),
		addresses:	make(map[topo.ID]string),
	}
	err = resolver.start()
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:82
		_go_fuzz_dep_.CoverTab[190533]++
																return nil, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:83
		// _ = "end of CoverTab[190533]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:84
		_go_fuzz_dep_.CoverTab[190534]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:84
		// _ = "end of CoverTab[190534]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:84
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:84
	// _ = "end of CoverTab[190527]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:84
	_go_fuzz_dep_.CoverTab[190528]++
															return resolver, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:85
	// _ = "end of CoverTab[190528]"
}

var _ resolver.Builder = (*ResolverBuilder)(nil)

// Resolver :
type Resolver struct {
	clientConn	resolver.ClientConn
	topoConn	*grpc.ClientConn
	serviceConfig	*serviceconfig.ParseResult
	masterships	map[topo.ID]topo.MastershipState	// E2 node to mastership (controls relation ID)
	controls	map[topo.ID]topo.ID			// controls relation to E2T ID
	addresses	map[topo.ID]string			// E2T ID to address
}

func (r *Resolver) start() error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:100
	_go_fuzz_dep_.CoverTab[190535]++
															log.Infof("Starting resolver")

															client := topo.NewTopoClient(r.topoConn)
															request := &topo.WatchRequest{}
															stream, err := client.Watch(context.Background(), request)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:106
		_go_fuzz_dep_.CoverTab[190538]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:107
		// _ = "end of CoverTab[190538]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:108
		_go_fuzz_dep_.CoverTab[190539]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:108
		// _ = "end of CoverTab[190539]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:108
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:108
	// _ = "end of CoverTab[190535]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:108
	_go_fuzz_dep_.CoverTab[190536]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:108
	_curRoutineNum165_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:108
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum165_)
															go func() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:109
		_go_fuzz_dep_.CoverTab[190540]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:109
		defer func() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:109
			_go_fuzz_dep_.CoverTab[190541]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:109
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum165_)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:109
			// _ = "end of CoverTab[190541]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:109
		}()
																for {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:110
			_go_fuzz_dep_.CoverTab[190542]++
																	response, err := stream.Recv()
																	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:112
				_go_fuzz_dep_.CoverTab[190544]++
																		return
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:113
				// _ = "end of CoverTab[190544]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:114
				_go_fuzz_dep_.CoverTab[190545]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:114
				// _ = "end of CoverTab[190545]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:114
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:114
			// _ = "end of CoverTab[190542]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:114
			_go_fuzz_dep_.CoverTab[190543]++
																	r.handleEvent(response.Event)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:115
			// _ = "end of CoverTab[190543]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:116
		// _ = "end of CoverTab[190540]"
	}()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:117
	// _ = "end of CoverTab[190536]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:117
	_go_fuzz_dep_.CoverTab[190537]++
															return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:118
	// _ = "end of CoverTab[190537]"
}

func (r *Resolver) handleEvent(event topo.Event) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:121
	_go_fuzz_dep_.CoverTab[190546]++
															object := event.Object
															if entity, ok := object.Obj.(*topo.Object_Entity); ok && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:123
		_go_fuzz_dep_.CoverTab[190547]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:123
		return entity.Entity.KindID == topo.E2NODE
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:123
		// _ = "end of CoverTab[190547]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:123
	}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:123
		_go_fuzz_dep_.CoverTab[190548]++

																switch event.Type {
		case topo.EventType_REMOVED:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:126
			_go_fuzz_dep_.CoverTab[190550]++
																	delete(r.masterships, object.ID)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:127
			// _ = "end of CoverTab[190550]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:128
			_go_fuzz_dep_.CoverTab[190551]++
																	var mastership topo.MastershipState
																	_ = object.GetAspect(&mastership)
																	if mastership.Term > r.masterships[object.ID].Term {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:131
				_go_fuzz_dep_.CoverTab[190552]++
																		r.masterships[object.ID] = mastership
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:132
				// _ = "end of CoverTab[190552]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:133
				_go_fuzz_dep_.CoverTab[190553]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:133
				// _ = "end of CoverTab[190553]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:133
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:133
			// _ = "end of CoverTab[190551]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:134
		// _ = "end of CoverTab[190548]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:134
		_go_fuzz_dep_.CoverTab[190549]++
																r.updateState()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:135
		// _ = "end of CoverTab[190549]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:136
		_go_fuzz_dep_.CoverTab[190554]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:136
		if entity, ok := object.Obj.(*topo.Object_Entity); ok && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:136
			_go_fuzz_dep_.CoverTab[190555]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:136
			return entity.Entity.KindID == topo.E2T
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:136
			// _ = "end of CoverTab[190555]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:136
		}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:136
			_go_fuzz_dep_.CoverTab[190556]++

																	switch event.Type {
			case topo.EventType_REMOVED:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:139
				_go_fuzz_dep_.CoverTab[190557]++
																		delete(r.addresses, object.ID)
																		r.updateState()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:141
				// _ = "end of CoverTab[190557]"
			default:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:142
				_go_fuzz_dep_.CoverTab[190558]++
																		var info topo.E2TInfo
																		_ = object.GetAspect(&info)
																		for _, iface := range info.Interfaces {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:145
					_go_fuzz_dep_.CoverTab[190559]++
																			if iface.Type == topo.Interface_INTERFACE_E2T {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:146
						_go_fuzz_dep_.CoverTab[190560]++
																				address := fmt.Sprintf("%s:%d", iface.IP, iface.Port)
																				if r.addresses[object.ID] != address {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:148
							_go_fuzz_dep_.CoverTab[190561]++
																					r.addresses[object.ID] = address
																					r.updateState()
																					break
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:151
							// _ = "end of CoverTab[190561]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:152
							_go_fuzz_dep_.CoverTab[190562]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:152
							// _ = "end of CoverTab[190562]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:152
						}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:152
						// _ = "end of CoverTab[190560]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:153
						_go_fuzz_dep_.CoverTab[190563]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:153
						// _ = "end of CoverTab[190563]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:153
					}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:153
					// _ = "end of CoverTab[190559]"
				}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:154
				// _ = "end of CoverTab[190558]"
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:155
			// _ = "end of CoverTab[190556]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:156
			_go_fuzz_dep_.CoverTab[190564]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:156
			if relation, ok := object.Obj.(*topo.Object_Relation); ok && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:156
				_go_fuzz_dep_.CoverTab[190565]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:156
				return relation.Relation.KindID == topo.CONTROLS
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:156
				// _ = "end of CoverTab[190565]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:156
			}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:156
				_go_fuzz_dep_.CoverTab[190566]++

																		switch event.Type {
				case topo.EventType_REMOVED:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:159
					_go_fuzz_dep_.CoverTab[190568]++
																			delete(r.controls, object.ID)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:160
					// _ = "end of CoverTab[190568]"
				default:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:161
					_go_fuzz_dep_.CoverTab[190569]++
																			r.controls[object.ID] = relation.Relation.SrcEntityID
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:162
					// _ = "end of CoverTab[190569]"
				}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:163
				// _ = "end of CoverTab[190566]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:163
				_go_fuzz_dep_.CoverTab[190567]++
																		r.updateState()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:164
				// _ = "end of CoverTab[190567]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:165
				_go_fuzz_dep_.CoverTab[190570]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:165
				// _ = "end of CoverTab[190570]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:165
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:165
			// _ = "end of CoverTab[190564]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:165
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:165
		// _ = "end of CoverTab[190554]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:165
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:165
	// _ = "end of CoverTab[190546]"
}

func (r *Resolver) updateState() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:168
	_go_fuzz_dep_.CoverTab[190571]++

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:171
	e2tE2Nodes := make(map[topo.ID][]string)

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:174
	for nodeID, mastership := range r.masterships {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:174
		_go_fuzz_dep_.CoverTab[190574]++
																if e2tID, ok := r.controls[topo.ID(mastership.NodeId)]; ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:175
			_go_fuzz_dep_.CoverTab[190575]++
																	e2tE2Nodes[e2tID] = append(e2tE2Nodes[e2tID], string(nodeID))
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:176
			// _ = "end of CoverTab[190575]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:177
			_go_fuzz_dep_.CoverTab[190576]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:177
			// _ = "end of CoverTab[190576]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:177
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:177
		// _ = "end of CoverTab[190574]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:178
	// _ = "end of CoverTab[190571]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:178
	_go_fuzz_dep_.CoverTab[190572]++

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:181
	addresses := make([]resolver.Address, 0, len(r.addresses))
	for e2tID, addr := range r.addresses {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:182
		_go_fuzz_dep_.CoverTab[190577]++
																if nodes, ok := e2tE2Nodes[e2tID]; ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:183
			_go_fuzz_dep_.CoverTab[190578]++
																	addresses = append(addresses, resolver.Address{
				Addr:	addr,
				Attributes: attributes.New(
					"nodes",
					nodes,
				),
			})
																	log.Debugf("New resolver address: %s => %+v", addr, nodes)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:191
			// _ = "end of CoverTab[190578]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:192
			_go_fuzz_dep_.CoverTab[190579]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:192
			// _ = "end of CoverTab[190579]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:192
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:192
		// _ = "end of CoverTab[190577]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:193
	// _ = "end of CoverTab[190572]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:193
	_go_fuzz_dep_.CoverTab[190573]++

															log.Infof("New resolver addresses: %+v", addresses)

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:199
	_ = r.clientConn.UpdateState(resolver.State{
		Addresses:	addresses,
		ServiceConfig:	r.serviceConfig,
	})
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:202
	// _ = "end of CoverTab[190573]"
}

// ResolveNow :
func (r *Resolver) ResolveNow(resolver.ResolveNowOptions) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:206
	_go_fuzz_dep_.CoverTab[190580]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:206
	// _ = "end of CoverTab[190580]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:206
}

// Close :
func (r *Resolver) Close() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:209
	_go_fuzz_dep_.CoverTab[190581]++
															if err := r.topoConn.Close(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:210
		_go_fuzz_dep_.CoverTab[190582]++
																log.Error("failed to close conn", err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:211
		// _ = "end of CoverTab[190582]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:212
		_go_fuzz_dep_.CoverTab[190583]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:212
		// _ = "end of CoverTab[190583]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:212
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:212
	// _ = "end of CoverTab[190581]"
}

var _ resolver.Resolver = (*Resolver)(nil)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:215
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/resolver.go:215
var _ = _go_fuzz_dep_.CoverTab
