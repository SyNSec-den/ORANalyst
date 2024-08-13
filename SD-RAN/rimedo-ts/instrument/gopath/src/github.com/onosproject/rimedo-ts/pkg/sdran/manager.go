// SPDX-FileCopyrightText: 2019-present Open Networking Foundation <info@opennetworking.org>
// SPDX-FileCopyrightText: 2019-present Rimedo Labs
//
// SPDX-License-Identifier: Apache-2.0

//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:6
package sdran

//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:6
import (
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:6
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:6
)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:6
import (
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:6
	_atomic_ "sync/atomic"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:6
)

import (
	"context"
	"strconv"
	"sync"

	policyAPI "github.com/onosproject/onos-a1-dm/go/policy_schemas/traffic_steering_preference/v2"
	e2tAPI "github.com/onosproject/onos-api/go/onos/e2t/e2"
	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/pdubuilder"
	e2sm_v2_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-v2-ies"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-lib-go/pkg/logging/service"
	"github.com/onosproject/onos-lib-go/pkg/northbound"
	control "github.com/onosproject/onos-mho/pkg/mho"
	"github.com/onosproject/onos-mho/pkg/store"
	"github.com/onosproject/rimedo-ts/pkg/mho"
	"github.com/onosproject/rimedo-ts/pkg/policy"
	"github.com/onosproject/rimedo-ts/pkg/rnib"
	"github.com/onosproject/rimedo-ts/pkg/southbound/e2"
)

var log = logging.GetLogger("rimedo-ts", "sdran", "manager")

type Config struct {
	AppID			string
	E2tAddress		string
	E2tPort			int
	TopoAddress		string
	TopoPort		int
	SMName			string
	SMVersion		string
	TSPolicySchemePath	string
}

func NewManager(config Config, flag bool) *Manager {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:42
	_go_fuzz_dep_.CoverTab[196790]++

									ueStore := store.NewStore()
									cellStore := store.NewStore()
									onosPolicyStore := store.NewStore()

									policyMap := make(map[string]*mho.PolicyData)

									indCh := make(chan *mho.E2NodeIndication)
									ctrlReqChs := make(map[string]chan *e2api.ControlMessage)

									options := e2.Options{
		AppID:		config.AppID,
		E2tAddress:	config.E2tAddress,
		E2tPort:	config.E2tPort,
		TopoAddress:	config.TopoAddress,
		TopoPort:	config.TopoPort,
		SMName:		config.SMName,
		SMVersion:	config.SMVersion,
	}

	e2Manager, err := e2.NewManager(options, indCh, ctrlReqChs)
	if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:64
		_go_fuzz_dep_.CoverTab[196792]++
										log.Warn(err)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:65
		// _ = "end of CoverTab[196792]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:66
		_go_fuzz_dep_.CoverTab[196793]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:66
		// _ = "end of CoverTab[196793]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:66
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:66
	// _ = "end of CoverTab[196790]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:66
	_go_fuzz_dep_.CoverTab[196791]++

									manager := &Manager{
		e2Manager:		e2Manager,
		mhoCtrl:		mho.NewController(indCh, ueStore, cellStore, onosPolicyStore, policyMap, flag),
		policyManager:		policy.NewPolicyManager(&policyMap),
		ueStore:		ueStore,
		cellStore:		cellStore,
		onosPolicyStore:	onosPolicyStore,
		ctrlReqChs:		ctrlReqChs,
		services:		[]service.Service{},
		mutex:			sync.RWMutex{},
	}
									return manager
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:79
	// _ = "end of CoverTab[196791]"
}

type Manager struct {
	e2Manager	e2.Manager
	mhoCtrl		*mho.Controller
	policyManager	*policy.PolicyManager
	ueStore		store.Store
	cellStore	store.Store
	onosPolicyStore	store.Store
	ctrlReqChs	map[string]chan *e2api.ControlMessage
	services	[]service.Service
	mutex		sync.RWMutex
}

func (m *Manager) Run(flag *bool) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:94
	_go_fuzz_dep_.CoverTab[196794]++
									if err := m.start(flag); err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:95
		_go_fuzz_dep_.CoverTab[196795]++
										log.Fatal("Unable to run Manager", err)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:96
		// _ = "end of CoverTab[196795]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:97
		_go_fuzz_dep_.CoverTab[196796]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:97
		// _ = "end of CoverTab[196796]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:97
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:97
	// _ = "end of CoverTab[196794]"
}

func (m *Manager) start(flag *bool) error {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:100
	_go_fuzz_dep_.CoverTab[196797]++
									m.startNorthboundServer()
									err := m.e2Manager.Start()
									if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:103
		_go_fuzz_dep_.CoverTab[196799]++
										log.Warn(err)
										return err
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:105
		// _ = "end of CoverTab[196799]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:106
		_go_fuzz_dep_.CoverTab[196800]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:106
		// _ = "end of CoverTab[196800]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:106
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:106
	// _ = "end of CoverTab[196797]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:106
	_go_fuzz_dep_.CoverTab[196798]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:106
	_curRoutineNum194_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:106
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum194_)

									go m.mhoCtrl.Run(context.Background(), flag)

									return nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:110
	// _ = "end of CoverTab[196798]"
}

func (m *Manager) startNorthboundServer() error {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:113
	_go_fuzz_dep_.CoverTab[196801]++

									s := northbound.NewServer(northbound.NewServerCfg(
		"",
		"",
		"",
		int16(5150),
		true,
		northbound.SecurityConfig{}))

	for i := range m.services {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:123
		_go_fuzz_dep_.CoverTab[196804]++
										s.AddService(m.services[i])
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:124
		// _ = "end of CoverTab[196804]"
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:125
	// _ = "end of CoverTab[196801]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:125
	_go_fuzz_dep_.CoverTab[196802]++

									doneCh := make(chan error)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:127
	_curRoutineNum195_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:127
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum195_)
									go func() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:128
		_go_fuzz_dep_.CoverTab[196805]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:128
		defer func() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:128
			_go_fuzz_dep_.CoverTab[196807]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:128
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum195_)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:128
			// _ = "end of CoverTab[196807]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:128
		}()
										err := s.Serve(func(started string) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:129
			_go_fuzz_dep_.CoverTab[196808]++
											close(doneCh)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:130
			// _ = "end of CoverTab[196808]"
		})
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:131
		// _ = "end of CoverTab[196805]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:131
		_go_fuzz_dep_.CoverTab[196806]++
										if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:132
			_go_fuzz_dep_.CoverTab[196809]++
											doneCh <- err
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:133
			// _ = "end of CoverTab[196809]"
		} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:134
			_go_fuzz_dep_.CoverTab[196810]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:134
			// _ = "end of CoverTab[196810]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:134
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:134
		// _ = "end of CoverTab[196806]"
	}()
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:135
	// _ = "end of CoverTab[196802]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:135
	_go_fuzz_dep_.CoverTab[196803]++
									return <-doneCh
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:136
	// _ = "end of CoverTab[196803]"
}

func (m *Manager) AddService(service service.Service) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:139
	_go_fuzz_dep_.CoverTab[196811]++

									m.services = append(m.services, service)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:141
	// _ = "end of CoverTab[196811]"

}

func (m *Manager) GetUEs(ctx context.Context) map[string]mho.UeData {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:145
	_go_fuzz_dep_.CoverTab[196812]++
									output := make(map[string]mho.UeData)
									chEntries := make(chan *store.Entry, 1024)
									err := m.ueStore.Entries(ctx, chEntries)
									if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:149
		_go_fuzz_dep_.CoverTab[196815]++
										log.Warn(err)
										return output
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:151
		// _ = "end of CoverTab[196815]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:152
		_go_fuzz_dep_.CoverTab[196816]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:152
		// _ = "end of CoverTab[196816]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:152
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:152
	// _ = "end of CoverTab[196812]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:152
	_go_fuzz_dep_.CoverTab[196813]++
									for entry := range chEntries {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:153
		_go_fuzz_dep_.CoverTab[196817]++
										ueData := entry.Value.(mho.UeData)
										output[ueData.UeID] = ueData
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:155
		// _ = "end of CoverTab[196817]"
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:156
	// _ = "end of CoverTab[196813]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:156
	_go_fuzz_dep_.CoverTab[196814]++
									return output
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:157
	// _ = "end of CoverTab[196814]"
}

func (m *Manager) GetCells(ctx context.Context) map[string]mho.CellData {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:160
	_go_fuzz_dep_.CoverTab[196818]++
									output := make(map[string]mho.CellData)
									chEntries := make(chan *store.Entry, 1024)
									err := m.cellStore.Entries(ctx, chEntries)
									if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:164
		_go_fuzz_dep_.CoverTab[196821]++
										log.Warn(err)
										return output
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:166
		// _ = "end of CoverTab[196821]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:167
		_go_fuzz_dep_.CoverTab[196822]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:167
		// _ = "end of CoverTab[196822]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:167
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:167
	// _ = "end of CoverTab[196818]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:167
	_go_fuzz_dep_.CoverTab[196819]++
									for entry := range chEntries {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:168
		_go_fuzz_dep_.CoverTab[196823]++
										cellData := entry.Value.(mho.CellData)
										output[cellData.CGIString] = cellData
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:170
		// _ = "end of CoverTab[196823]"
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:171
	// _ = "end of CoverTab[196819]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:171
	_go_fuzz_dep_.CoverTab[196820]++
									return output
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:172
	// _ = "end of CoverTab[196820]"
}

func (m *Manager) GetPolicies(ctx context.Context) map[string]mho.PolicyData {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:175
	_go_fuzz_dep_.CoverTab[196824]++
									output := make(map[string]mho.PolicyData)
									chEntries := make(chan *store.Entry, 1024)
									err := m.onosPolicyStore.Entries(ctx, chEntries)
									if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:179
		_go_fuzz_dep_.CoverTab[196827]++
										log.Warn(err)
										return output
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:181
		// _ = "end of CoverTab[196827]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:182
		_go_fuzz_dep_.CoverTab[196828]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:182
		// _ = "end of CoverTab[196828]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:182
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:182
	// _ = "end of CoverTab[196824]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:182
	_go_fuzz_dep_.CoverTab[196825]++
									for entry := range chEntries {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:183
		_go_fuzz_dep_.CoverTab[196829]++
										policyData := entry.Value.(mho.PolicyData)
										output[policyData.Key] = policyData
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:185
		// _ = "end of CoverTab[196829]"
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:186
	// _ = "end of CoverTab[196825]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:186
	_go_fuzz_dep_.CoverTab[196826]++
									return output
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:187
	// _ = "end of CoverTab[196826]"
}

func (m *Manager) GetCellTypes(ctx context.Context) map[string]rnib.Cell {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:190
	_go_fuzz_dep_.CoverTab[196830]++
									return m.e2Manager.GetCellTypes(ctx)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:191
	// _ = "end of CoverTab[196830]"
}

func (m *Manager) SetCellType(ctx context.Context, cellID string, cellType string) error {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:194
	_go_fuzz_dep_.CoverTab[196831]++
									return m.e2Manager.SetCellType(ctx, cellID, cellType)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:195
	// _ = "end of CoverTab[196831]"
}

func (m *Manager) GetCell(ctx context.Context, CGI string) *mho.CellData {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:198
	_go_fuzz_dep_.CoverTab[196832]++

									return m.mhoCtrl.GetCell(ctx, CGI)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:200
	// _ = "end of CoverTab[196832]"

}

func (m *Manager) SetCell(ctx context.Context, cell *mho.CellData) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:204
	_go_fuzz_dep_.CoverTab[196833]++

									m.mhoCtrl.SetCell(ctx, cell)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:206
	// _ = "end of CoverTab[196833]"

}

func (m *Manager) AttachUe(ctx context.Context, ue *mho.UeData, CGI string, cgiObject *e2sm_v2_ies.Cgi) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:210
	_go_fuzz_dep_.CoverTab[196834]++

									m.mhoCtrl.AttachUe(ctx, ue, CGI, cgiObject)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:212
	// _ = "end of CoverTab[196834]"

}

func (m *Manager) GetUe(ctx context.Context, ueID string) *mho.UeData {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:216
	_go_fuzz_dep_.CoverTab[196835]++

									return m.mhoCtrl.GetUe(ctx, ueID)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:218
	// _ = "end of CoverTab[196835]"

}

func (m *Manager) SetUe(ctx context.Context, ueData *mho.UeData) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:222
	_go_fuzz_dep_.CoverTab[196836]++

									m.mhoCtrl.SetUe(ctx, ueData)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:224
	// _ = "end of CoverTab[196836]"

}

func (m *Manager) CreatePolicy(ctx context.Context, key string, policy *policyAPI.API) *mho.PolicyData {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:228
	_go_fuzz_dep_.CoverTab[196837]++

									return m.mhoCtrl.CreatePolicy(ctx, key, policy)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:230
	// _ = "end of CoverTab[196837]"

}

func (m *Manager) GetPolicy(ctx context.Context, key string) *mho.PolicyData {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:234
	_go_fuzz_dep_.CoverTab[196838]++

									return m.mhoCtrl.GetPolicy(ctx, key)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:236
	// _ = "end of CoverTab[196838]"

}

func (m *Manager) SetPolicy(ctx context.Context, key string, policy *mho.PolicyData) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:240
	_go_fuzz_dep_.CoverTab[196839]++

									m.mhoCtrl.SetPolicy(ctx, key, policy)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:242
	// _ = "end of CoverTab[196839]"

}

func (m *Manager) DeletePolicy(ctx context.Context, key string) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:246
	_go_fuzz_dep_.CoverTab[196840]++

									m.mhoCtrl.DeletePolicy(ctx, key)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:248
	// _ = "end of CoverTab[196840]"

}

func (m *Manager) GetPolicyStore() *store.Store {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:252
	_go_fuzz_dep_.CoverTab[196841]++
									return m.mhoCtrl.GetPolicyStore()
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:253
	// _ = "end of CoverTab[196841]"
}

func (m *Manager) GetControlChannelsMap(ctx context.Context) map[string]chan *e2api.ControlMessage {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:256
	_go_fuzz_dep_.CoverTab[196842]++
									return m.ctrlReqChs
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:257
	// _ = "end of CoverTab[196842]"
}

func (m *Manager) GetPolicyManager() *policy.PolicyManager {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:260
	_go_fuzz_dep_.CoverTab[196843]++
									return m.policyManager
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:261
	// _ = "end of CoverTab[196843]"
}

func (m *Manager) SwitchUeBetweenCells(ctx context.Context, ueID string, targetCellCGI string) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:264
	_go_fuzz_dep_.CoverTab[196844]++

									m.mutex.Lock()
									defer m.mutex.Unlock()

									availableUes := m.GetUEs(ctx)
									chosenUe := availableUes[ueID]

									if shouldBeSwitched(chosenUe, targetCellCGI) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:272
		_go_fuzz_dep_.CoverTab[196845]++

										targetCell := m.GetCell(ctx, targetCellCGI)
										servingCell := m.GetCell(ctx, chosenUe.CGIString)

										targetCell.CumulativeHandoversOut++
										servingCell.CumulativeHandoversIn++

										chosenUe.Idle = false
										m.AttachUe(ctx, &chosenUe, targetCellCGI, targetCell.CGI)

										m.SetCell(ctx, targetCell)
										m.SetCell(ctx, servingCell)

										controlChannel := m.ctrlReqChs[chosenUe.E2NodeID]

										controlHandler := &control.E2SmMhoControlHandler{
			NodeID:			chosenUe.E2NodeID,
			ControlAckRequest:	e2tAPI.ControlAckRequest_NO_ACK,
		}

		ueIDnum, err := strconv.Atoi(chosenUe.UeID)
		if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:294
			_go_fuzz_dep_.CoverTab[196848]++
											log.Errorf("SendHORequest() failed to convert string %v to decimal number - assumption is not satisfied (UEID is a decimal number): %v", chosenUe.UeID, err)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:295
			// _ = "end of CoverTab[196848]"
		} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:296
			_go_fuzz_dep_.CoverTab[196849]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:296
			// _ = "end of CoverTab[196849]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:296
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:296
		// _ = "end of CoverTab[196845]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:296
		_go_fuzz_dep_.CoverTab[196846]++

										ueIdentity, err := pdubuilder.CreateUeIDGNb(int64(ueIDnum), []byte{0xAA, 0xBB, 0xCC}, []byte{0xDD}, []byte{0xCC, 0xC0}, []byte{0xFC})
										if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:299
			_go_fuzz_dep_.CoverTab[196850]++
											log.Errorf("SendHORequest() Failed to create UEID: %v", err)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:300
			// _ = "end of CoverTab[196850]"
		} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:301
			_go_fuzz_dep_.CoverTab[196851]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:301
			// _ = "end of CoverTab[196851]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:301
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:301
		// _ = "end of CoverTab[196846]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:301
		_go_fuzz_dep_.CoverTab[196847]++

										servingPlmnIDBytes := servingCell.CGI.GetNRCgi().GetPLmnidentity().GetValue()
										servingNCI := servingCell.CGI.GetNRCgi().GetNRcellIdentity().GetValue().GetValue()
										servingNCILen := servingCell.CGI.GetNRCgi().GetNRcellIdentity().GetValue().GetLen()
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:305
		_curRoutineNum196_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:305
		_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum196_)

										go func() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:307
			_go_fuzz_dep_.CoverTab[196852]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:307
			defer func() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:307
				_go_fuzz_dep_.CoverTab[196853]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:307
				_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum196_)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:307
				// _ = "end of CoverTab[196853]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:307
			}()
											if controlHandler.ControlHeader, err = controlHandler.CreateMhoControlHeader(servingNCI, servingNCILen, 1, servingPlmnIDBytes); err == nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:308
				_go_fuzz_dep_.CoverTab[196854]++

												if controlHandler.ControlMessage, err = controlHandler.CreateMhoControlMessage(servingCell.CGI, ueIdentity, targetCell.CGI); err == nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:310
					_go_fuzz_dep_.CoverTab[196855]++

													if controlRequest, err := controlHandler.CreateMhoControlRequest(); err == nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:312
						_go_fuzz_dep_.CoverTab[196856]++

														controlChannel <- controlRequest
														log.Infof("CONTROL MESSAGE: UE [ID:%v, 5QI:%v] switched between CELLs [CGI:%v -> CGI:%v]\n", chosenUe.UeID, chosenUe.FiveQi, servingCell.CGIString, targetCell.CGIString)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:315
						// _ = "end of CoverTab[196856]"

					} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:317
						_go_fuzz_dep_.CoverTab[196857]++
														log.Warn("Control request problem!", err)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:318
						// _ = "end of CoverTab[196857]"
					}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:319
					// _ = "end of CoverTab[196855]"
				} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:320
					_go_fuzz_dep_.CoverTab[196858]++
													log.Warn("Control message problem!", err)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:321
					// _ = "end of CoverTab[196858]"
				}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:322
				// _ = "end of CoverTab[196854]"
			} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:323
				_go_fuzz_dep_.CoverTab[196859]++
												log.Warn("Control header problem!", err)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:324
				// _ = "end of CoverTab[196859]"
			}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:325
			// _ = "end of CoverTab[196852]"
		}()
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:326
		// _ = "end of CoverTab[196847]"

	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:328
		_go_fuzz_dep_.CoverTab[196860]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:328
		// _ = "end of CoverTab[196860]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:328
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:328
	// _ = "end of CoverTab[196844]"

}

func shouldBeSwitched(ue mho.UeData, cgi string) bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:332
	_go_fuzz_dep_.CoverTab[196861]++

									servingCgi := ue.CGIString
									if servingCgi == cgi {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:335
		_go_fuzz_dep_.CoverTab[196863]++
										return false
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:336
		// _ = "end of CoverTab[196863]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:337
		_go_fuzz_dep_.CoverTab[196864]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:337
		// _ = "end of CoverTab[196864]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:337
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:337
	// _ = "end of CoverTab[196861]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:337
	_go_fuzz_dep_.CoverTab[196862]++
									return true
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:338
	// _ = "end of CoverTab[196862]"

}

//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:340
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/sdran/manager.go:340
var _ = _go_fuzz_dep_.CoverTab
