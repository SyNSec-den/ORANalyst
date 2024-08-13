// SPDX-FileCopyrightText: 2019-present Open Networking Foundation <info@opennetworking.org>
// SPDX-FileCopyrightText: 2019-present Rimedo Labs
//
// SPDX-License-Identifier: Apache-2.0
// Copy from https://github.com/woojoong88/onos-kpimon/tree/sample-a1t-xapp/pkg/northbound/a1
// modified by RIMEDO Labs team

//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/manager.go:8
package a1

//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/manager.go:8
import (
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/manager.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/manager.go:8
)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/manager.go:8
import (
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/manager.go:8
	_atomic_ "sync/atomic"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/manager.go:8
)

import (
	"context"

	"github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	a1connection "github.com/onosproject/onos-ric-sdk-go/pkg/a1/connection"
)

var log = logging.GetLogger("rimedo-ts", "a1")

type Config struct {
	PolicyName		string
	PolicyVersion		string
	PolicyID		string
	PolicyDescription	string
	A1tPort			int
}

func NewManager(caPath string, keyPath string, certPath string, grpcPort int, xAppName string, a1PolicyTypes []*topo.A1PolicyType) (*Manager, error) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/manager.go:28
	_go_fuzz_dep_.CoverTab[190748]++
										a1ConnManager, err := a1connection.NewManager(caPath, keyPath, certPath, grpcPort, a1PolicyTypes)
										if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/manager.go:30
		_go_fuzz_dep_.CoverTab[190750]++
											return nil, err
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/manager.go:31
		// _ = "end of CoverTab[190750]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/manager.go:32
		_go_fuzz_dep_.CoverTab[190751]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/manager.go:32
		// _ = "end of CoverTab[190751]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/manager.go:32
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/manager.go:32
	// _ = "end of CoverTab[190748]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/manager.go:32
	_go_fuzz_dep_.CoverTab[190749]++
										return &Manager{
		a1ConnManager: a1ConnManager,
	}, nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/manager.go:35
	// _ = "end of CoverTab[190749]"
}

type Manager struct {
	a1ConnManager *a1connection.Manager
}

func (m *Manager) Start() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/manager.go:42
	_go_fuzz_dep_.CoverTab[190752]++
										m.a1ConnManager.Start(context.Background())
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/manager.go:43
	// _ = "end of CoverTab[190752]"
}

func (m *Manager) Close(ctx context.Context) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/manager.go:46
	_go_fuzz_dep_.CoverTab[190753]++
										err := m.a1ConnManager.DeleteXAppElementOnTopo(ctx)
										if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/manager.go:48
		_go_fuzz_dep_.CoverTab[190754]++
											log.Error(err)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/manager.go:49
		// _ = "end of CoverTab[190754]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/manager.go:50
		_go_fuzz_dep_.CoverTab[190755]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/manager.go:50
		// _ = "end of CoverTab[190755]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/manager.go:50
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/manager.go:50
	// _ = "end of CoverTab[190753]"
}

//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/manager.go:51
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/manager.go:51
var _ = _go_fuzz_dep_.CoverTab
