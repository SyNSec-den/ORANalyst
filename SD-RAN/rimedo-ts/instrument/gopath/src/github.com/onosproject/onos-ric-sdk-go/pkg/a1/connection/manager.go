// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:5
package a1connection

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:5
)

import (
	"context"
	gogotypes "github.com/gogo/protobuf/types"
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-lib-go/pkg/env"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	a1endpoint "github.com/onosproject/onos-ric-sdk-go/pkg/a1/endpoint"
	"github.com/onosproject/onos-ric-sdk-go/pkg/topo"
	"github.com/onosproject/onos-ric-sdk-go/pkg/utils"
)

var log = logging.GetLogger("a1", "manager")

// NewManager creates a new A1 manager
func NewManager(caPath string, keyPath string, certPath string, grpcPort int, a1PolicyTypes []*topoapi.A1PolicyType) (*Manager, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:22
	_go_fuzz_dep_.CoverTab[190621]++
															topoClient, err := topo.NewClient()
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:24
		_go_fuzz_dep_.CoverTab[190623]++
																return nil, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:25
		// _ = "end of CoverTab[190623]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:26
		_go_fuzz_dep_.CoverTab[190624]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:26
		// _ = "end of CoverTab[190624]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:26
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:26
	// _ = "end of CoverTab[190621]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:26
	_go_fuzz_dep_.CoverTab[190622]++
															return &Manager{
		id:		utils.GetXappTopoID(),
		server:		a1endpoint.NewServer(caPath, keyPath, certPath, grpcPort),
		topoClient:	topoClient,
		a1PolicyTypes:	a1PolicyTypes,
	}, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:32
	// _ = "end of CoverTab[190622]"
}

// Manager is a struct of A1 interface
type Manager struct {
	id		topoapi.ID
	server		a1endpoint.Server
	topoClient	topo.Client
	a1PolicyTypes	[]*topoapi.A1PolicyType
}

// Start inits and starts A1 server
func (m *Manager) Start(ctx context.Context) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:44
	_go_fuzz_dep_.CoverTab[190625]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:44
	_curRoutineNum167_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:44
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum167_)
															go func(ctx context.Context) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:45
		_go_fuzz_dep_.CoverTab[190626]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:45
		defer func() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:45
			_go_fuzz_dep_.CoverTab[190627]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:45
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum167_)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:45
			// _ = "end of CoverTab[190627]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:45
		}()
																log.Infof("Start (or restart) A1 connection manager")
																err := m.AddXAppElementOnTopo(ctx)
																if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:48
			_go_fuzz_dep_.CoverTab[190628]++
																	log.Warn(err)
																	log.Warn()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:50
			// _ = "end of CoverTab[190628]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:51
			_go_fuzz_dep_.CoverTab[190629]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:51
			// _ = "end of CoverTab[190629]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:51
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:51
		// _ = "end of CoverTab[190626]"
	}(ctx)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:52
	// _ = "end of CoverTab[190625]"
}

// GetID returns ID
func (m *Manager) GetID() topoapi.ID {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:56
	_go_fuzz_dep_.CoverTab[190630]++
															return m.id
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:57
	// _ = "end of CoverTab[190630]"
}

// AddXAppElementOnTopo adds XApp type on topo
func (m *Manager) AddXAppElementOnTopo(ctx context.Context) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:61
	_go_fuzz_dep_.CoverTab[190631]++
															object := &topoapi.Object{
		ID:	m.id,
		Type:	topoapi.Object_ENTITY,
		Obj: &topoapi.Object_Entity{
			Entity: &topoapi.Entity{
				KindID: topoapi.XAPP,
			},
		},
		Aspects:	make(map[string]*gogotypes.Any),
		Labels:		map[string]string{},
	}
	interfaces := make([]*topoapi.Interface, 1)
	interfaces[0] = &topoapi.Interface{
		IP:	env.GetPodIP(),
		Port:	uint32(m.server.GRPCPort),
		Type:	topoapi.Interface_INTERFACE_A1_XAPP,
	}

	aspect := &topoapi.XAppInfo{
		Interfaces:	interfaces,
		A1PolicyTypes:	m.a1PolicyTypes,
	}

	err := object.SetAspect(aspect)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:86
		_go_fuzz_dep_.CoverTab[190635]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:87
		// _ = "end of CoverTab[190635]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:88
		_go_fuzz_dep_.CoverTab[190636]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:88
		// _ = "end of CoverTab[190636]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:88
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:88
	// _ = "end of CoverTab[190631]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:88
	_go_fuzz_dep_.CoverTab[190632]++
															obj, err := m.topoClient.Get(ctx, m.id)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:90
		_go_fuzz_dep_.CoverTab[190637]++
																err = m.topoClient.Create(ctx, object)
																return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:92
		// _ = "end of CoverTab[190637]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:93
		_go_fuzz_dep_.CoverTab[190638]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:93
		// _ = "end of CoverTab[190638]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:93
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:93
	// _ = "end of CoverTab[190632]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:93
	_go_fuzz_dep_.CoverTab[190633]++
															log.Warn("Update topo A1 XApp entity to have the latest aspects since there is already exist")
															err = obj.SetAspect(aspect)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:96
		_go_fuzz_dep_.CoverTab[190639]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:97
		// _ = "end of CoverTab[190639]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:98
		_go_fuzz_dep_.CoverTab[190640]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:98
		// _ = "end of CoverTab[190640]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:98
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:98
	// _ = "end of CoverTab[190633]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:98
	_go_fuzz_dep_.CoverTab[190634]++
															err = m.topoClient.Update(ctx, obj)
															return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:100
	// _ = "end of CoverTab[190634]"
}

// DeleteXAppElementOnTopo removes all aspects on topo
func (m *Manager) DeleteXAppElementOnTopo(ctx context.Context) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:104
	_go_fuzz_dep_.CoverTab[190641]++
															obj, err := m.topoClient.Get(ctx, m.id)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:106
		_go_fuzz_dep_.CoverTab[190645]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:107
		// _ = "end of CoverTab[190645]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:108
		_go_fuzz_dep_.CoverTab[190646]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:108
		// _ = "end of CoverTab[190646]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:108
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:108
	// _ = "end of CoverTab[190641]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:108
	_go_fuzz_dep_.CoverTab[190642]++
															aspect := &topoapi.XAppInfo{}
															err = obj.SetAspect(aspect)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:111
		_go_fuzz_dep_.CoverTab[190647]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:112
		// _ = "end of CoverTab[190647]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:113
		_go_fuzz_dep_.CoverTab[190648]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:113
		// _ = "end of CoverTab[190648]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:113
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:113
	// _ = "end of CoverTab[190642]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:113
	_go_fuzz_dep_.CoverTab[190643]++
															err = m.topoClient.Update(ctx, obj)
															if err != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:115
		_go_fuzz_dep_.CoverTab[190649]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:115
		return !errors.IsAlreadyExists(err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:115
		// _ = "end of CoverTab[190649]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:115
	}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:115
		_go_fuzz_dep_.CoverTab[190650]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:116
		// _ = "end of CoverTab[190650]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:117
		_go_fuzz_dep_.CoverTab[190651]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:117
		// _ = "end of CoverTab[190651]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:117
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:117
	// _ = "end of CoverTab[190643]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:117
	_go_fuzz_dep_.CoverTab[190644]++
															return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:118
	// _ = "end of CoverTab[190644]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:119
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/connection/manager.go:119
var _ = _go_fuzz_dep_.CoverTab
