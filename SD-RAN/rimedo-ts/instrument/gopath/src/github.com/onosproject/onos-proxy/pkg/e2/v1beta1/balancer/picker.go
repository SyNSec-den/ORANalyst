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

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/picker.go:15
package balancer

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/picker.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/picker.go:15
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/picker.go:15
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/picker.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/picker.go:15
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/picker.go:15
)

import (
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
	"google.golang.org/grpc/metadata"
)

const e2NodeIDHeader = "e2-node-id"

func init() {
	balancer.Register(base.NewBalancerBuilder(ResolverName, &PickerBuilder{}, base.Config{}))
}

// PickerBuilder :
type PickerBuilder struct{}

// Build :
func (p *PickerBuilder) Build(info base.PickerBuildInfo) balancer.Picker {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/picker.go:33
	_go_fuzz_dep_.CoverTab[190512]++
															masters := make(map[string]balancer.SubConn)

															for sc, scInfo := range info.ReadySCs {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/picker.go:36
		_go_fuzz_dep_.CoverTab[190514]++
																nodes := scInfo.Address.Attributes.Value("nodes").([]string)
																for _, node := range nodes {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/picker.go:38
			_go_fuzz_dep_.CoverTab[190515]++
																	log.Debugf("E2 node %s is mastered by E2T %s; conn=%+v", node, scInfo.Address.Addr, sc)
																	masters[node] = sc
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/picker.go:40
			// _ = "end of CoverTab[190515]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/picker.go:41
		// _ = "end of CoverTab[190514]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/picker.go:42
	// _ = "end of CoverTab[190512]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/picker.go:42
	_go_fuzz_dep_.CoverTab[190513]++
															log.Infof("Built new picker for E2T instances: %+v", masters)
															return &Picker{
		masters: masters,
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/picker.go:46
	// _ = "end of CoverTab[190513]"
}

var _ base.PickerBuilder = (*PickerBuilder)(nil)

// Picker :
type Picker struct {
	masters map[string]balancer.SubConn	// NodeID string to connection mapping
}

// Pick :
func (p *Picker) Pick(info balancer.PickInfo) (balancer.PickResult, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/picker.go:57
	_go_fuzz_dep_.CoverTab[190516]++
															var result balancer.PickResult
															if md, ok := metadata.FromOutgoingContext(info.Ctx); ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/picker.go:59
		_go_fuzz_dep_.CoverTab[190518]++
																ids := md.Get(e2NodeIDHeader)
																if len(ids) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/picker.go:61
			_go_fuzz_dep_.CoverTab[190519]++
																	if subConn, ok := p.masters[ids[0]]; ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/picker.go:62
				_go_fuzz_dep_.CoverTab[190520]++
																		log.Debugf("Picked subconn for %s: %+v", ids[0], subConn)
																		result.SubConn = subConn
																		return result, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/picker.go:65
				// _ = "end of CoverTab[190520]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/picker.go:66
				_go_fuzz_dep_.CoverTab[190521]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/picker.go:66
				// _ = "end of CoverTab[190521]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/picker.go:66
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/picker.go:66
			// _ = "end of CoverTab[190519]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/picker.go:67
			_go_fuzz_dep_.CoverTab[190522]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/picker.go:67
			// _ = "end of CoverTab[190522]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/picker.go:67
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/picker.go:67
		// _ = "end of CoverTab[190518]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/picker.go:68
		_go_fuzz_dep_.CoverTab[190523]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/picker.go:68
		// _ = "end of CoverTab[190523]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/picker.go:68
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/picker.go:68
	// _ = "end of CoverTab[190516]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/picker.go:68
	_go_fuzz_dep_.CoverTab[190517]++
															log.Warn("No subconn available")
															return result, balancer.ErrNoSubConnAvailable
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/picker.go:70
	// _ = "end of CoverTab[190517]"
}

var _ balancer.Picker = (*Picker)(nil)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/picker.go:73
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/e2/v1beta1/balancer/picker.go:73
var _ = _go_fuzz_dep_.CoverTab
