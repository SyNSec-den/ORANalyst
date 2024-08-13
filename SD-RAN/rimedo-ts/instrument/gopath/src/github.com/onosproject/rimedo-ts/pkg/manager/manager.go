//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:1
// SPDX-FileCopyrightText: 2019-present Open Networking Foundation <info@opennetworking.org>
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:1
// SPDX-FileCopyrightText: 2019-present Rimedo Labs
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:1
//
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:1
// SPDX-License-Identifier: Apache-2.0
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:1
// Created by RIMEDO-Labs team
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:1
// based on any onosproject manager
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:7
package manager

//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:7
import (
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:7
)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:7
import (
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:7
	_atomic_ "sync/atomic"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:7
)

import (
	"context"
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"sync"
	"time"

	policyAPI "github.com/onosproject/onos-a1-dm/go/policy_schemas/traffic_steering_preference/v2"
	topoAPI "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/rimedo-ts/pkg/mho"
	"github.com/onosproject/rimedo-ts/pkg/northbound/a1"
	"github.com/onosproject/rimedo-ts/pkg/sdran"
)

var log = logging.GetLogger("rimedo-ts", "ts-manager")
var logLength = 150
var nodesLogLen = 0
var policiesLogLen = 0

type Config struct {
	AppID		string
	E2tAddress	string
	E2tPort		int
	TopoAddress	string
	TopoPort	int
	SMName		string
	SMVersion	string
}

func NewManager(sdranConfig sdran.Config, a1Config a1.Config, flag bool) *Manager {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:41
	_go_fuzz_dep_.CoverTab[196865]++

										sdranManager := sdran.NewManager(sdranConfig, flag)

										a1PolicyTypes := make([]*topoAPI.A1PolicyType, 0)
										a1Policy := &topoAPI.A1PolicyType{
		Name:		topoAPI.PolicyTypeName(a1Config.PolicyName),
		Version:	topoAPI.PolicyTypeVersion(a1Config.PolicyVersion),
		ID:		topoAPI.PolicyTypeID(a1Config.PolicyID),
		Description:	topoAPI.PolicyTypeDescription(a1Config.PolicyDescription),
	}
	a1PolicyTypes = append(a1PolicyTypes, a1Policy)

	a1Manager, err := a1.NewManager("", "", "", a1Config.A1tPort, sdranConfig.AppID, a1PolicyTypes)
	if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:55
		_go_fuzz_dep_.CoverTab[196867]++
											log.Warn(err)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:56
		// _ = "end of CoverTab[196867]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:57
		_go_fuzz_dep_.CoverTab[196868]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:57
		// _ = "end of CoverTab[196868]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:57
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:57
	// _ = "end of CoverTab[196865]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:57
	_go_fuzz_dep_.CoverTab[196866]++

										manager := &Manager{
		sdranManager:	sdranManager,
		a1Manager:	*a1Manager,
		topoIDsEnabled:	flag,
		mutex:		sync.RWMutex{},
	}
										return manager
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:65
	// _ = "end of CoverTab[196866]"
}

type Manager struct {
	sdranManager	*sdran.Manager
	a1Manager	a1.Manager
	topoIDsEnabled	bool
	mutex		sync.RWMutex
}

func (m *Manager) Run() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:75
	_go_fuzz_dep_.CoverTab[196869]++

										if err := m.start(); err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:77
		_go_fuzz_dep_.CoverTab[196870]++
											log.Fatal("Unable to run Manager", err)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:78
		// _ = "end of CoverTab[196870]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:79
		_go_fuzz_dep_.CoverTab[196871]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:79
		// _ = "end of CoverTab[196871]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:79
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:79
	// _ = "end of CoverTab[196869]"

}

func (m *Manager) Close() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:83
	_go_fuzz_dep_.CoverTab[196872]++
										m.a1Manager.Close(context.Background())
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:84
	// _ = "end of CoverTab[196872]"
}

func (m *Manager) start() error {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:87
	_go_fuzz_dep_.CoverTab[196873]++

										ctx := context.Background()

										policyMap := make(map[string][]byte)

										policyChange := make(chan bool)

										m.sdranManager.AddService(a1.NewA1EIService())
										m.sdranManager.AddService(a1.NewA1PService(&policyMap, policyChange))

										handleFlag := false

										m.sdranManager.Run(&handleFlag)

										m.a1Manager.Start()
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:102
	_curRoutineNum197_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:102
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum197_)

										go func() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:104
		_go_fuzz_dep_.CoverTab[196876]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:104
		defer func() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:104
			_go_fuzz_dep_.CoverTab[196877]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:104
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum197_)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:104
			// _ = "end of CoverTab[196877]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:104
		}()
											for range policyChange {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:105
			_go_fuzz_dep_.CoverTab[196878]++
												log.Debug("")
												drawWithLine("POLICY STORE CHANGED!", logLength)
												log.Debug("")
												if err := m.updatePolicies(ctx, policyMap); err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:109
				_go_fuzz_dep_.CoverTab[196880]++
													log.Warn("Some problems occured when updating Policy store!")
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:110
				// _ = "end of CoverTab[196880]"
			} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:111
				_go_fuzz_dep_.CoverTab[196881]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:111
				// _ = "end of CoverTab[196881]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:111
			}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:111
			// _ = "end of CoverTab[196878]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:111
			_go_fuzz_dep_.CoverTab[196879]++
												log.Debug("")
												m.checkPolicies(ctx, true, true, true)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:113
			// _ = "end of CoverTab[196879]"
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:114
		// _ = "end of CoverTab[196876]"

	}()
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:116
	// _ = "end of CoverTab[196873]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:116
	_go_fuzz_dep_.CoverTab[196874]++
										flag := true
										show := false
										prepare := false
										counter := 0
										delay := 3
										time.Sleep(5 * time.Second)
										log.Info("\n\n\n\n\n\n\n\n\n\n")
										handleFlag = true
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:124
	_curRoutineNum198_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:124
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum198_)
										go func() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:125
		_go_fuzz_dep_.CoverTab[196882]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:125
		defer func() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:125
			_go_fuzz_dep_.CoverTab[196883]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:125
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum198_)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:125
			// _ = "end of CoverTab[196883]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:125
		}()
											for {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:126
			_go_fuzz_dep_.CoverTab[196884]++
												time.Sleep(1 * time.Second)
												counter++
												if counter == delay {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:129
				_go_fuzz_dep_.CoverTab[196886]++
													compareLengths()
													counter = 0
													show = true
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:132
				// _ = "end of CoverTab[196886]"
			} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:133
				_go_fuzz_dep_.CoverTab[196887]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:133
				if counter == delay-1 {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:133
					_go_fuzz_dep_.CoverTab[196888]++
														prepare = true
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:134
					// _ = "end of CoverTab[196888]"
				} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:135
					_go_fuzz_dep_.CoverTab[196889]++
														show = false
														prepare = false
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:137
					// _ = "end of CoverTab[196889]"
				}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:138
				// _ = "end of CoverTab[196887]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:138
			}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:138
			// _ = "end of CoverTab[196884]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:138
			_go_fuzz_dep_.CoverTab[196885]++
												m.checkPolicies(ctx, flag, show, prepare)
												m.showAvailableNodes(ctx, show, prepare)
												flag = false
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:141
			// _ = "end of CoverTab[196885]"
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:142
		// _ = "end of CoverTab[196882]"
	}()
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:143
	// _ = "end of CoverTab[196874]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:143
	_go_fuzz_dep_.CoverTab[196875]++

										return nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:145
	// _ = "end of CoverTab[196875]"
}

func (m *Manager) updatePolicies(ctx context.Context, policyMap map[string][]byte) error {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:148
	_go_fuzz_dep_.CoverTab[196890]++
										m.mutex.Lock()
										defer m.mutex.Unlock()
										policies := m.sdranManager.GetPolicies(ctx)
										for k := range policies {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:152
		_go_fuzz_dep_.CoverTab[196893]++
											if _, ok := policyMap[k]; !ok {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:153
			_go_fuzz_dep_.CoverTab[196894]++
												m.sdranManager.DeletePolicy(ctx, k)
												log.Infof("POLICY MESSAGE: Policy [ID:%v] deleted\n", k)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:155
			// _ = "end of CoverTab[196894]"
		} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:156
			_go_fuzz_dep_.CoverTab[196895]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:156
			// _ = "end of CoverTab[196895]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:156
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:156
		// _ = "end of CoverTab[196893]"
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:157
	// _ = "end of CoverTab[196890]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:157
	_go_fuzz_dep_.CoverTab[196891]++
										for i := range policyMap {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:158
		_go_fuzz_dep_.CoverTab[196896]++
											r, err := policyAPI.UnmarshalAPI(policyMap[i])
											if err == nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:160
			_go_fuzz_dep_.CoverTab[196897]++
												policyObject := m.sdranManager.CreatePolicy(ctx, i, &r)
												info := fmt.Sprintf("POLICY MESSAGE: Policy [ID:%v] applied -> ", policyObject.Key)
												previous := false
												if policyObject.API.Scope.SliceID != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:164
				_go_fuzz_dep_.CoverTab[196903]++
													info = info + fmt.Sprintf("Slice [SD:%v, SST:%v, PLMN:(MCC:%v, MNC:%v)]", *policyObject.API.Scope.SliceID.SD, policyObject.API.Scope.SliceID.Sst, policyObject.API.Scope.SliceID.PlmnID.Mcc, policyObject.API.Scope.SliceID.PlmnID.Mnc)
													previous = true
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:166
				// _ = "end of CoverTab[196903]"
			} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:167
				_go_fuzz_dep_.CoverTab[196904]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:167
				// _ = "end of CoverTab[196904]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:167
			}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:167
			// _ = "end of CoverTab[196897]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:167
			_go_fuzz_dep_.CoverTab[196898]++
												if policyObject.API.Scope.UeID != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:168
				_go_fuzz_dep_.CoverTab[196905]++
													if previous {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:169
					_go_fuzz_dep_.CoverTab[196908]++
														info = info + ", "
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:170
					// _ = "end of CoverTab[196908]"
				} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:171
					_go_fuzz_dep_.CoverTab[196909]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:171
					// _ = "end of CoverTab[196909]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:171
				}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:171
				// _ = "end of CoverTab[196905]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:171
				_go_fuzz_dep_.CoverTab[196906]++
													ue := *policyObject.API.Scope.UeID
													new_ue := ue
													for i := 0; i < len(ue); i++ {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:174
					_go_fuzz_dep_.CoverTab[196910]++
														if ue[i:i+1] == "0" {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:175
						_go_fuzz_dep_.CoverTab[196911]++
															new_ue = ue[i+1:]
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:176
						// _ = "end of CoverTab[196911]"
					} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:177
						_go_fuzz_dep_.CoverTab[196912]++
															break
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:178
						// _ = "end of CoverTab[196912]"
					}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:179
					// _ = "end of CoverTab[196910]"
				}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:180
				// _ = "end of CoverTab[196906]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:180
				_go_fuzz_dep_.CoverTab[196907]++
													info = info + fmt.Sprintf("UE [ID:%v]", new_ue)
													previous = true
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:182
				// _ = "end of CoverTab[196907]"
			} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:183
				_go_fuzz_dep_.CoverTab[196913]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:183
				// _ = "end of CoverTab[196913]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:183
			}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:183
			// _ = "end of CoverTab[196898]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:183
			_go_fuzz_dep_.CoverTab[196899]++
												if policyObject.API.Scope.QosID != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:184
				_go_fuzz_dep_.CoverTab[196914]++
													if previous {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:185
					_go_fuzz_dep_.CoverTab[196917]++
														info = info + ", "
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:186
					// _ = "end of CoverTab[196917]"
				} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:187
					_go_fuzz_dep_.CoverTab[196918]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:187
					// _ = "end of CoverTab[196918]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:187
				}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:187
				// _ = "end of CoverTab[196914]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:187
				_go_fuzz_dep_.CoverTab[196915]++
													if policyObject.API.Scope.QosID.QcI != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:188
					_go_fuzz_dep_.CoverTab[196919]++
														info = info + fmt.Sprintf("QoS [QCI:%v]", *policyObject.API.Scope.QosID.QcI)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:189
					// _ = "end of CoverTab[196919]"
				} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:190
					_go_fuzz_dep_.CoverTab[196920]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:190
					// _ = "end of CoverTab[196920]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:190
				}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:190
				// _ = "end of CoverTab[196915]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:190
				_go_fuzz_dep_.CoverTab[196916]++
													if policyObject.API.Scope.QosID.The5QI != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:191
					_go_fuzz_dep_.CoverTab[196921]++
														info = info + fmt.Sprintf("QoS [5QI:%v]", *policyObject.API.Scope.QosID.The5QI)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:192
					// _ = "end of CoverTab[196921]"
				} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:193
					_go_fuzz_dep_.CoverTab[196922]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:193
					// _ = "end of CoverTab[196922]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:193
				}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:193
				// _ = "end of CoverTab[196916]"
			} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:194
				_go_fuzz_dep_.CoverTab[196923]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:194
				// _ = "end of CoverTab[196923]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:194
			}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:194
			// _ = "end of CoverTab[196899]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:194
			_go_fuzz_dep_.CoverTab[196900]++
												if policyObject.API.Scope.CellID != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:195
				_go_fuzz_dep_.CoverTab[196924]++
													if previous {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:196
					_go_fuzz_dep_.CoverTab[196928]++
														info = info + ", "
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:197
					// _ = "end of CoverTab[196928]"
				} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:198
					_go_fuzz_dep_.CoverTab[196929]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:198
					// _ = "end of CoverTab[196929]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:198
				}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:198
				// _ = "end of CoverTab[196924]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:198
				_go_fuzz_dep_.CoverTab[196925]++
													info = info + "CELL ["
													if policyObject.API.Scope.CellID.CID.NcI != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:200
					_go_fuzz_dep_.CoverTab[196930]++

														info = info + fmt.Sprintf("NCI:%v, ", *policyObject.API.Scope.CellID.CID.NcI)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:202
					// _ = "end of CoverTab[196930]"
				} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:203
					_go_fuzz_dep_.CoverTab[196931]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:203
					// _ = "end of CoverTab[196931]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:203
				}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:203
				// _ = "end of CoverTab[196925]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:203
				_go_fuzz_dep_.CoverTab[196926]++
													if policyObject.API.Scope.CellID.CID.EcI != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:204
					_go_fuzz_dep_.CoverTab[196932]++

														info = info + fmt.Sprintf("ECI:%v, ", *policyObject.API.Scope.CellID.CID.EcI)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:206
					// _ = "end of CoverTab[196932]"
				} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:207
					_go_fuzz_dep_.CoverTab[196933]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:207
					// _ = "end of CoverTab[196933]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:207
				}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:207
				// _ = "end of CoverTab[196926]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:207
				_go_fuzz_dep_.CoverTab[196927]++
													info = info + fmt.Sprintf("PLMN:(MCC:%v, MNC:%v)]", policyObject.API.Scope.CellID.PlmnID.Mcc, policyObject.API.Scope.CellID.PlmnID.Mnc)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:208
				// _ = "end of CoverTab[196927]"
			} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:209
				_go_fuzz_dep_.CoverTab[196934]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:209
				// _ = "end of CoverTab[196934]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:209
			}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:209
			// _ = "end of CoverTab[196900]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:209
			_go_fuzz_dep_.CoverTab[196901]++
												for i := range policyObject.API.TSPResources {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:210
				_go_fuzz_dep_.CoverTab[196935]++
													info = info + fmt.Sprintf(" - (%v) -", policyObject.API.TSPResources[i].Preference)
													for j := range policyObject.API.TSPResources[i].CellIDList {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:212
					_go_fuzz_dep_.CoverTab[196937]++
														nci := *policyObject.API.TSPResources[i].CellIDList[j].CID.NcI
														plmnId, _ := mho.GetPlmnIdFromMccMnc(policyObject.API.TSPResources[i].CellIDList[j].PlmnID.Mcc, policyObject.API.TSPResources[i].CellIDList[j].PlmnID.Mnc)
														cgi := m.PlmnIDNciToCGI(plmnId, uint64(nci))
														info = info + fmt.Sprintf(" CELL [CGI:%v],", cgi)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:216
					// _ = "end of CoverTab[196937]"
				}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:217
				// _ = "end of CoverTab[196935]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:217
				_go_fuzz_dep_.CoverTab[196936]++
													info = info[0 : len(info)-1]
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:218
				// _ = "end of CoverTab[196936]"

			}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:220
			// _ = "end of CoverTab[196901]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:220
			_go_fuzz_dep_.CoverTab[196902]++
												info = info + "\n"
												log.Info(info)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:222
			// _ = "end of CoverTab[196902]"
		} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:223
			_go_fuzz_dep_.CoverTab[196938]++
												log.Warn("Can't unmarshal the JSON file!")
												return err
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:225
			// _ = "end of CoverTab[196938]"
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:226
		// _ = "end of CoverTab[196896]"
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:227
	// _ = "end of CoverTab[196891]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:227
	_go_fuzz_dep_.CoverTab[196892]++
										return nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:228
	// _ = "end of CoverTab[196892]"
}

func (m *Manager) deployPolicies(ctx context.Context) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:231
	_go_fuzz_dep_.CoverTab[196939]++
										policyManager := m.sdranManager.GetPolicyManager()
										ues := m.sdranManager.GetUEs(ctx)
										keys := make([]string, 0, len(ues))
										for k := range ues {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:235
		_go_fuzz_dep_.CoverTab[196941]++
											keys = append(keys, k)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:236
		// _ = "end of CoverTab[196941]"
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:237
	// _ = "end of CoverTab[196939]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:237
	_go_fuzz_dep_.CoverTab[196940]++
										sort.Strings(keys)

										for i := range keys {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:240
		_go_fuzz_dep_.CoverTab[196942]++
											var cellIDs []policyAPI.CellID
											var rsrps []int
											fiveQi := ues[keys[i]].FiveQi
											sd := "456DEF"
											scopeUe := policyAPI.Scope{

			SliceID: &policyAPI.SliceID{
				SD:	&sd,
				Sst:	1,
				PlmnID: policyAPI.PlmnID{
					Mcc:	"314",
					Mnc:	"628",
				},
			},
			UeID:	&keys[i],
			QosID: &policyAPI.QosID{
				The5QI: &fiveQi,
			},
		}

		cgiKeys := make([]string, 0, len(ues[keys[i]].CgiTable))
		for cgi := range ues[keys[i]].CgiTable {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:262
			_go_fuzz_dep_.CoverTab[196946]++
												cgiKeys = append(cgiKeys, cgi)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:263
			// _ = "end of CoverTab[196946]"
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:264
		// _ = "end of CoverTab[196942]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:264
		_go_fuzz_dep_.CoverTab[196943]++
											inside := false
											for j := range cgiKeys {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:266
			_go_fuzz_dep_.CoverTab[196947]++

												inside = true
												cgi := ues[keys[i]].CgiTable[cgiKeys[j]]
												nci := int64(mho.GetNciFromCellGlobalID(cgi))
												plmnIdBytes := mho.GetPlmnIDBytesFromCellGlobalID(cgi)
												plmnId := mho.PlmnIDBytesToInt(plmnIdBytes)
												mcc, mnc := mho.GetMccMncFromPlmnID(plmnId)
												cellID := policyAPI.CellID{
				CID: policyAPI.CID{
					NcI: &nci,
				},
				PlmnID: policyAPI.PlmnID{
					Mcc:	mcc,
					Mnc:	mnc,
				},
			}

												cellIDs = append(cellIDs, cellID)
												rsrps = append(rsrps, int(ues[keys[i]].RsrpTable[cgiKeys[j]]))
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:285
			// _ = "end of CoverTab[196947]"

		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:287
		// _ = "end of CoverTab[196943]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:287
		_go_fuzz_dep_.CoverTab[196944]++

											if inside {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:289
			_go_fuzz_dep_.CoverTab[196948]++

												tsResult := policyManager.GetTsResultForUEV2(scopeUe, rsrps, cellIDs)
												plmnId, err := mho.GetPlmnIdFromMccMnc(tsResult.PlmnID.Mcc, tsResult.PlmnID.Mnc)

												if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:294
				_go_fuzz_dep_.CoverTab[196949]++
													log.Warnf("Cannot get PLMN ID from these MCC and MNC parameters:%v,%v.", tsResult.PlmnID.Mcc, tsResult.PlmnID.Mnc)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:295
				// _ = "end of CoverTab[196949]"
			} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:296
				_go_fuzz_dep_.CoverTab[196950]++
													targetCellCGI := m.PlmnIDNciToCGI(plmnId, uint64(*tsResult.CID.NcI))
													m.sdranManager.SwitchUeBetweenCells(ctx, keys[i], targetCellCGI)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:298
				// _ = "end of CoverTab[196950]"
			}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:299
			// _ = "end of CoverTab[196948]"

		} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:301
			_go_fuzz_dep_.CoverTab[196951]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:301
			// _ = "end of CoverTab[196951]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:301
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:301
		// _ = "end of CoverTab[196944]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:301
		_go_fuzz_dep_.CoverTab[196945]++

											cellIDs = nil
											rsrps = nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:304
		// _ = "end of CoverTab[196945]"

	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:306
	// _ = "end of CoverTab[196940]"

}

func (m *Manager) checkPolicies(ctx context.Context, defaultFlag bool, showFlag bool, prepareFlag bool) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:310
	_go_fuzz_dep_.CoverTab[196952]++
										m.mutex.Lock()
										defer m.mutex.Unlock()
										policyLen := 0
										policies := m.sdranManager.GetPolicies(ctx)
										keys := make([]string, 0, len(policies))
										for k := range policies {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:316
		_go_fuzz_dep_.CoverTab[196956]++
											keys = append(keys, k)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:317
		// _ = "end of CoverTab[196956]"
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:318
	// _ = "end of CoverTab[196952]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:318
	_go_fuzz_dep_.CoverTab[196953]++
										sort.Strings(keys)
										if defaultFlag && func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:320
		_go_fuzz_dep_.CoverTab[196957]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:320
		return (len(policies) == 0)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:320
		// _ = "end of CoverTab[196957]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:320
	}() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:320
		_go_fuzz_dep_.CoverTab[196958]++
											log.Infof("POLICY MESSAGE: Default policy applied\n")
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:321
		// _ = "end of CoverTab[196958]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:322
		_go_fuzz_dep_.CoverTab[196959]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:322
		// _ = "end of CoverTab[196959]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:322
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:322
	// _ = "end of CoverTab[196953]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:322
	_go_fuzz_dep_.CoverTab[196954]++
										if prepareFlag && func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:323
		_go_fuzz_dep_.CoverTab[196960]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:323
		return len(policies) != 0
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:323
		// _ = "end of CoverTab[196960]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:323
	}() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:323
		_go_fuzz_dep_.CoverTab[196961]++
											if showFlag {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:324
			_go_fuzz_dep_.CoverTab[196964]++
												log.Debug("")
												drawWithLine("POLICIES", logLength)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:326
			// _ = "end of CoverTab[196964]"
		} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:327
			_go_fuzz_dep_.CoverTab[196965]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:327
			// _ = "end of CoverTab[196965]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:327
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:327
		// _ = "end of CoverTab[196961]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:327
		_go_fuzz_dep_.CoverTab[196962]++
											for _, key := range keys {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:328
			_go_fuzz_dep_.CoverTab[196966]++
												policyObject := policies[key]
												info := fmt.Sprintf("ID:%v POLICY: {", policyObject.Key)
												previous := false
												if policyObject.API.Scope.SliceID != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:332
				_go_fuzz_dep_.CoverTab[196974]++
													info = info + fmt.Sprintf("Slice [SD:%v, SST:%v, PLMN:(MCC:%v, MNC:%v)]", *policyObject.API.Scope.SliceID.SD, policyObject.API.Scope.SliceID.Sst, policyObject.API.Scope.SliceID.PlmnID.Mcc, policyObject.API.Scope.SliceID.PlmnID.Mnc)
													previous = true
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:334
				// _ = "end of CoverTab[196974]"
			} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:335
				_go_fuzz_dep_.CoverTab[196975]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:335
				// _ = "end of CoverTab[196975]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:335
			}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:335
			// _ = "end of CoverTab[196966]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:335
			_go_fuzz_dep_.CoverTab[196967]++
												if policyObject.API.Scope.UeID != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:336
				_go_fuzz_dep_.CoverTab[196976]++
													if previous {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:337
					_go_fuzz_dep_.CoverTab[196979]++
														info = info + ", "
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:338
					// _ = "end of CoverTab[196979]"
				} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:339
					_go_fuzz_dep_.CoverTab[196980]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:339
					// _ = "end of CoverTab[196980]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:339
				}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:339
				// _ = "end of CoverTab[196976]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:339
				_go_fuzz_dep_.CoverTab[196977]++
													ue := *policyObject.API.Scope.UeID
													new_ue := ue
													for i := 0; i < len(ue); i++ {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:342
					_go_fuzz_dep_.CoverTab[196981]++
														if ue[i:i+1] == "0" {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:343
						_go_fuzz_dep_.CoverTab[196982]++
															new_ue = ue[i+1:]
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:344
						// _ = "end of CoverTab[196982]"
					} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:345
						_go_fuzz_dep_.CoverTab[196983]++
															break
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:346
						// _ = "end of CoverTab[196983]"
					}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:347
					// _ = "end of CoverTab[196981]"
				}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:348
				// _ = "end of CoverTab[196977]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:348
				_go_fuzz_dep_.CoverTab[196978]++
													info = info + fmt.Sprintf("UE [ID:%v]", new_ue)
													previous = true
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:350
				// _ = "end of CoverTab[196978]"
			} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:351
				_go_fuzz_dep_.CoverTab[196984]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:351
				// _ = "end of CoverTab[196984]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:351
			}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:351
			// _ = "end of CoverTab[196967]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:351
			_go_fuzz_dep_.CoverTab[196968]++
												if policyObject.API.Scope.QosID != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:352
				_go_fuzz_dep_.CoverTab[196985]++
													if previous {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:353
					_go_fuzz_dep_.CoverTab[196988]++
														info = info + ", "
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:354
					// _ = "end of CoverTab[196988]"
				} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:355
					_go_fuzz_dep_.CoverTab[196989]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:355
					// _ = "end of CoverTab[196989]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:355
				}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:355
				// _ = "end of CoverTab[196985]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:355
				_go_fuzz_dep_.CoverTab[196986]++
													if policyObject.API.Scope.QosID.QcI != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:356
					_go_fuzz_dep_.CoverTab[196990]++
														info = info + fmt.Sprintf("QoS [QCI:%v]", *policyObject.API.Scope.QosID.QcI)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:357
					// _ = "end of CoverTab[196990]"
				} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:358
					_go_fuzz_dep_.CoverTab[196991]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:358
					// _ = "end of CoverTab[196991]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:358
				}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:358
				// _ = "end of CoverTab[196986]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:358
				_go_fuzz_dep_.CoverTab[196987]++
													if policyObject.API.Scope.QosID.The5QI != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:359
					_go_fuzz_dep_.CoverTab[196992]++
														info = info + fmt.Sprintf("QoS [5QI:%v]", *policyObject.API.Scope.QosID.The5QI)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:360
					// _ = "end of CoverTab[196992]"
				} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:361
					_go_fuzz_dep_.CoverTab[196993]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:361
					// _ = "end of CoverTab[196993]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:361
				}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:361
				// _ = "end of CoverTab[196987]"
			} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:362
				_go_fuzz_dep_.CoverTab[196994]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:362
				// _ = "end of CoverTab[196994]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:362
			}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:362
			// _ = "end of CoverTab[196968]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:362
			_go_fuzz_dep_.CoverTab[196969]++
												if policyObject.API.Scope.CellID != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:363
				_go_fuzz_dep_.CoverTab[196995]++
													if previous {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:364
					_go_fuzz_dep_.CoverTab[196999]++
														info = info + ", "
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:365
					// _ = "end of CoverTab[196999]"
				} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:366
					_go_fuzz_dep_.CoverTab[197000]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:366
					// _ = "end of CoverTab[197000]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:366
				}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:366
				// _ = "end of CoverTab[196995]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:366
				_go_fuzz_dep_.CoverTab[196996]++
													info = info + "CELL ["
													if policyObject.API.Scope.CellID.CID.NcI != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:368
					_go_fuzz_dep_.CoverTab[197001]++

														info = info + fmt.Sprintf("NCI:%v, ", *policyObject.API.Scope.CellID.CID.NcI)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:370
					// _ = "end of CoverTab[197001]"
				} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:371
					_go_fuzz_dep_.CoverTab[197002]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:371
					// _ = "end of CoverTab[197002]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:371
				}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:371
				// _ = "end of CoverTab[196996]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:371
				_go_fuzz_dep_.CoverTab[196997]++
													if policyObject.API.Scope.CellID.CID.EcI != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:372
					_go_fuzz_dep_.CoverTab[197003]++

														info = info + fmt.Sprintf("ECI:%v, ", *policyObject.API.Scope.CellID.CID.EcI)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:374
					// _ = "end of CoverTab[197003]"
				} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:375
					_go_fuzz_dep_.CoverTab[197004]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:375
					// _ = "end of CoverTab[197004]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:375
				}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:375
				// _ = "end of CoverTab[196997]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:375
				_go_fuzz_dep_.CoverTab[196998]++
													info = info + fmt.Sprintf("PLMN:(MCC:%v, MNC:%v)]", policyObject.API.Scope.CellID.PlmnID.Mcc, policyObject.API.Scope.CellID.PlmnID.Mnc)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:376
				// _ = "end of CoverTab[196998]"
			} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:377
				_go_fuzz_dep_.CoverTab[197005]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:377
				// _ = "end of CoverTab[197005]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:377
			}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:377
			// _ = "end of CoverTab[196969]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:377
			_go_fuzz_dep_.CoverTab[196970]++
												for i := range policyObject.API.TSPResources {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:378
				_go_fuzz_dep_.CoverTab[197006]++
													info = info + fmt.Sprintf(" - (%v) -", policyObject.API.TSPResources[i].Preference)
													for j := range policyObject.API.TSPResources[i].CellIDList {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:380
					_go_fuzz_dep_.CoverTab[197008]++
														nci := *policyObject.API.TSPResources[i].CellIDList[j].CID.NcI
														plmnId, _ := mho.GetPlmnIdFromMccMnc(policyObject.API.TSPResources[i].CellIDList[j].PlmnID.Mcc, policyObject.API.TSPResources[i].CellIDList[j].PlmnID.Mnc)
														cgi := m.PlmnIDNciToCGI(plmnId, uint64(nci))
														info = info + fmt.Sprintf(" CELL [CGI:%v],", cgi)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:384
					// _ = "end of CoverTab[197008]"
				}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:385
				// _ = "end of CoverTab[197006]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:385
				_go_fuzz_dep_.CoverTab[197007]++
													info = info[0 : len(info)-1]
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:386
				// _ = "end of CoverTab[197007]"

			}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:388
			// _ = "end of CoverTab[196970]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:388
			_go_fuzz_dep_.CoverTab[196971]++
												info = info + "} STATUS: "
												if policyObject.IsEnforced {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:390
				_go_fuzz_dep_.CoverTab[197009]++
													info = info + "ENFORCED"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:391
				// _ = "end of CoverTab[197009]"
			} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:392
				_go_fuzz_dep_.CoverTab[197010]++
													info = info + "NOT ENFORCED"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:393
				// _ = "end of CoverTab[197010]"
			}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:394
			// _ = "end of CoverTab[196971]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:394
			_go_fuzz_dep_.CoverTab[196972]++
												if policyLen < len(info) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:395
				_go_fuzz_dep_.CoverTab[197011]++
													policyLen = len(info)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:396
				// _ = "end of CoverTab[197011]"
			} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:397
				_go_fuzz_dep_.CoverTab[197012]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:397
				// _ = "end of CoverTab[197012]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:397
			}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:397
			// _ = "end of CoverTab[196972]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:397
			_go_fuzz_dep_.CoverTab[196973]++
												if showFlag {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:398
				_go_fuzz_dep_.CoverTab[197013]++
													log.Debug(info)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:399
				// _ = "end of CoverTab[197013]"
			} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:400
				_go_fuzz_dep_.CoverTab[197014]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:400
				// _ = "end of CoverTab[197014]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:400
			}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:400
			// _ = "end of CoverTab[196973]"
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:401
		// _ = "end of CoverTab[196962]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:401
		_go_fuzz_dep_.CoverTab[196963]++
											if showFlag {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:402
			_go_fuzz_dep_.CoverTab[197015]++
												log.Debug("")
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:403
			// _ = "end of CoverTab[197015]"
		} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:404
			_go_fuzz_dep_.CoverTab[197016]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:404
			// _ = "end of CoverTab[197016]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:404
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:404
		// _ = "end of CoverTab[196963]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:405
		_go_fuzz_dep_.CoverTab[197017]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:405
		// _ = "end of CoverTab[197017]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:405
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:405
	// _ = "end of CoverTab[196954]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:405
	_go_fuzz_dep_.CoverTab[196955]++
										m.deployPolicies(ctx)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:406
	// _ = "end of CoverTab[196955]"
}

func (m *Manager) showAvailableNodes(ctx context.Context, showFlag bool, prepareFlag bool) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:409
	_go_fuzz_dep_.CoverTab[197018]++
										m.mutex.Lock()
										defer m.mutex.Unlock()
										cellLen := 0
										ueLen := 0
										cells := m.sdranManager.GetCellTypes(ctx)
										cellsObjects := m.sdranManager.GetCells(ctx)
										keys := make([]string, 0, len(cells))
										for k := range cells {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:417
		_go_fuzz_dep_.CoverTab[197023]++
											keys = append(keys, k)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:418
		// _ = "end of CoverTab[197023]"
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:419
	// _ = "end of CoverTab[197018]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:419
	_go_fuzz_dep_.CoverTab[197019]++
										sort.Strings(keys)
										if prepareFlag && func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:421
		_go_fuzz_dep_.CoverTab[197024]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:421
		return len(cells) > 0
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:421
		// _ = "end of CoverTab[197024]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:421
	}() && func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:421
		_go_fuzz_dep_.CoverTab[197025]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:421
		return len(cellsObjects) > 0
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:421
		// _ = "end of CoverTab[197025]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:421
	}() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:421
		_go_fuzz_dep_.CoverTab[197026]++
											if showFlag {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:422
			_go_fuzz_dep_.CoverTab[197028]++
												log.Debug("")
												drawWithLine("CELLS", logLength)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:424
			// _ = "end of CoverTab[197028]"
		} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:425
			_go_fuzz_dep_.CoverTab[197029]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:425
			// _ = "end of CoverTab[197029]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:425
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:425
		// _ = "end of CoverTab[197026]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:425
		_go_fuzz_dep_.CoverTab[197027]++
											for _, key := range keys {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:426
			_go_fuzz_dep_.CoverTab[197030]++
												cgi_str := m.CgiFromTopoToIndicationFormat(cells[key].CGI)
												info := fmt.Sprintf("ID:%v CGI:%v UEs:[", key, cgi_str)
												cellObject := m.sdranManager.GetCell(ctx, cgi_str)
												inside := false
												if cellObject != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:431
				_go_fuzz_dep_.CoverTab[197034]++
													for ue := range cellObject.Ues {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:432
					_go_fuzz_dep_.CoverTab[197035]++
														inside = true
														new_ue := ue
														for i := 0; i < len(ue); i++ {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:435
						_go_fuzz_dep_.CoverTab[197037]++
															if ue[i:i+1] == "0" {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:436
							_go_fuzz_dep_.CoverTab[197038]++
																new_ue = ue[i+1:]
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:437
							// _ = "end of CoverTab[197038]"
						} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:438
							_go_fuzz_dep_.CoverTab[197039]++
																break
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:439
							// _ = "end of CoverTab[197039]"
						}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:440
						// _ = "end of CoverTab[197037]"
					}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:441
					// _ = "end of CoverTab[197035]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:441
					_go_fuzz_dep_.CoverTab[197036]++
														info = info + new_ue + " "
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:442
					// _ = "end of CoverTab[197036]"
				}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:443
				// _ = "end of CoverTab[197034]"
			} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:444
				_go_fuzz_dep_.CoverTab[197040]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:444
				// _ = "end of CoverTab[197040]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:444
			}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:444
			// _ = "end of CoverTab[197030]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:444
			_go_fuzz_dep_.CoverTab[197031]++
												if inside {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:445
				_go_fuzz_dep_.CoverTab[197041]++
													info = info[:len(info)-1]
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:446
				// _ = "end of CoverTab[197041]"
			} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:447
				_go_fuzz_dep_.CoverTab[197042]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:447
				// _ = "end of CoverTab[197042]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:447
			}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:447
			// _ = "end of CoverTab[197031]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:447
			_go_fuzz_dep_.CoverTab[197032]++
												info = info + "]"
												if cellLen < len(info) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:449
				_go_fuzz_dep_.CoverTab[197043]++
													cellLen = len(info)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:450
				// _ = "end of CoverTab[197043]"
			} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:451
				_go_fuzz_dep_.CoverTab[197044]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:451
				// _ = "end of CoverTab[197044]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:451
			}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:451
			// _ = "end of CoverTab[197032]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:451
			_go_fuzz_dep_.CoverTab[197033]++
												if showFlag {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:452
				_go_fuzz_dep_.CoverTab[197045]++
													log.Debug(info)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:453
				// _ = "end of CoverTab[197045]"
			} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:454
				_go_fuzz_dep_.CoverTab[197046]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:454
				// _ = "end of CoverTab[197046]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:454
			}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:454
			// _ = "end of CoverTab[197033]"
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:455
		// _ = "end of CoverTab[197027]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:456
		_go_fuzz_dep_.CoverTab[197047]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:456
		// _ = "end of CoverTab[197047]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:456
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:456
	// _ = "end of CoverTab[197019]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:456
	_go_fuzz_dep_.CoverTab[197020]++

										ues := m.sdranManager.GetUEs(ctx)
										keys = make([]string, 0, len(ues))
										for k := range ues {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:460
		_go_fuzz_dep_.CoverTab[197048]++
											keys = append(keys, k)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:461
		// _ = "end of CoverTab[197048]"
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:462
	// _ = "end of CoverTab[197020]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:462
	_go_fuzz_dep_.CoverTab[197021]++
										sort.Strings(keys)
										if prepareFlag && func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:464
		_go_fuzz_dep_.CoverTab[197049]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:464
		return len(ues) > 0
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:464
		// _ = "end of CoverTab[197049]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:464
	}() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:464
		_go_fuzz_dep_.CoverTab[197050]++
											if showFlag {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:465
			_go_fuzz_dep_.CoverTab[197053]++
												log.Debug("")
												drawWithLine("UES", logLength)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:467
			// _ = "end of CoverTab[197053]"
		} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:468
			_go_fuzz_dep_.CoverTab[197054]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:468
			// _ = "end of CoverTab[197054]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:468
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:468
		// _ = "end of CoverTab[197050]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:468
		_go_fuzz_dep_.CoverTab[197051]++
											for _, key := range keys {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:469
			_go_fuzz_dep_.CoverTab[197055]++
												ueIdString, _ := strconv.Atoi(key)
												cgiString := ues[key].CGIString
												if cgiString == "" {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:472
				_go_fuzz_dep_.CoverTab[197062]++
													cgiString = "NONE"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:473
				// _ = "end of CoverTab[197062]"
			} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:474
				_go_fuzz_dep_.CoverTab[197063]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:474
				// _ = "end of CoverTab[197063]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:474
			}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:474
			// _ = "end of CoverTab[197055]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:474
			_go_fuzz_dep_.CoverTab[197056]++
												status := "CONNECTED"
												if ues[key].Idle {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:476
				_go_fuzz_dep_.CoverTab[197064]++
													status = "IDLE     "
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:477
				// _ = "end of CoverTab[197064]"
			} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:478
				_go_fuzz_dep_.CoverTab[197065]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:478
				// _ = "end of CoverTab[197065]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:478
			}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:478
			// _ = "end of CoverTab[197056]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:478
			_go_fuzz_dep_.CoverTab[197057]++
												info := fmt.Sprintf("ID:%v STATUS:%v 5QI: %v CGI:%v CGIs(RSRP): [", ueIdString, status, ues[key].FiveQi, cgiString)

												cgi_keys := make([]string, 0, len(ues[key].RsrpTable))
												for k := range ues[key].RsrpTable {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:482
				_go_fuzz_dep_.CoverTab[197066]++
													cgi_keys = append(cgi_keys, k)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:483
				// _ = "end of CoverTab[197066]"
			}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:484
			// _ = "end of CoverTab[197057]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:484
			_go_fuzz_dep_.CoverTab[197058]++
												sort.Strings(cgi_keys)
												inside := false
												for _, cgi := range cgi_keys {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:487
				_go_fuzz_dep_.CoverTab[197067]++
													inside = true
													info += fmt.Sprintf("%v (%v) ", cgi, ues[key].RsrpTable[cgi])
													rsrp_str := strconv.Itoa(int(ues[key].RsrpTable[cgi]))
													if len(rsrp_str) < 4 {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:491
					_go_fuzz_dep_.CoverTab[197068]++
														diff := 4 - len(rsrp_str)
														for i := 0; i < diff; i++ {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:493
						_go_fuzz_dep_.CoverTab[197069]++
															info = info + " "
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:494
						// _ = "end of CoverTab[197069]"
					}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:495
					// _ = "end of CoverTab[197068]"
				} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:496
					_go_fuzz_dep_.CoverTab[197070]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:496
					// _ = "end of CoverTab[197070]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:496
				}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:496
				// _ = "end of CoverTab[197067]"
			}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:497
			// _ = "end of CoverTab[197058]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:497
			_go_fuzz_dep_.CoverTab[197059]++
												if inside {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:498
				_go_fuzz_dep_.CoverTab[197071]++
													info = info[:len(info)-1]
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:499
				// _ = "end of CoverTab[197071]"
			} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:500
				_go_fuzz_dep_.CoverTab[197072]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:500
				// _ = "end of CoverTab[197072]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:500
			}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:500
			// _ = "end of CoverTab[197059]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:500
			_go_fuzz_dep_.CoverTab[197060]++
												info = info + "]"
												if ueLen < len(info) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:502
				_go_fuzz_dep_.CoverTab[197073]++
													ueLen = len(info)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:503
				// _ = "end of CoverTab[197073]"
			} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:504
				_go_fuzz_dep_.CoverTab[197074]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:504
				// _ = "end of CoverTab[197074]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:504
			}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:504
			// _ = "end of CoverTab[197060]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:504
			_go_fuzz_dep_.CoverTab[197061]++
												if showFlag {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:505
				_go_fuzz_dep_.CoverTab[197075]++
													log.Debug(info)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:506
				// _ = "end of CoverTab[197075]"
			} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:507
				_go_fuzz_dep_.CoverTab[197076]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:507
				// _ = "end of CoverTab[197076]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:507
			}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:507
			// _ = "end of CoverTab[197061]"
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:508
		// _ = "end of CoverTab[197051]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:508
		_go_fuzz_dep_.CoverTab[197052]++
											if showFlag {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:509
			_go_fuzz_dep_.CoverTab[197077]++
												log.Debug("")
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:510
			// _ = "end of CoverTab[197077]"
		} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:511
			_go_fuzz_dep_.CoverTab[197078]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:511
			// _ = "end of CoverTab[197078]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:511
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:511
		// _ = "end of CoverTab[197052]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:512
		_go_fuzz_dep_.CoverTab[197079]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:512
		// _ = "end of CoverTab[197079]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:512
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:512
	// _ = "end of CoverTab[197021]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:512
	_go_fuzz_dep_.CoverTab[197022]++
										if cellLen > ueLen {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:513
		_go_fuzz_dep_.CoverTab[197080]++
											nodesLogLen = cellLen
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:514
		// _ = "end of CoverTab[197080]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:515
		_go_fuzz_dep_.CoverTab[197081]++
											nodesLogLen = ueLen
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:516
		// _ = "end of CoverTab[197081]"
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:517
	// _ = "end of CoverTab[197022]"
}

func (m *Manager) changeCellsTypes(ctx context.Context) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:520
	_go_fuzz_dep_.CoverTab[197082]++
										m.mutex.Lock()
										defer m.mutex.Unlock()
										cellTypes := make(map[int]string)
										cellTypes[0] = "Macro"
										cellTypes[1] = "SmallCell"
										for {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:526
		_go_fuzz_dep_.CoverTab[197083]++
											time.Sleep(10 * time.Second)
											cells := m.sdranManager.GetCellTypes(ctx)
											type_id := rand.Intn(len(cellTypes))
											for key, val := range cells {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:530
			_go_fuzz_dep_.CoverTab[197084]++
												_ = val
												err := m.sdranManager.SetCellType(ctx, key, cellTypes[type_id])
												if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:533
				_go_fuzz_dep_.CoverTab[197086]++
													log.Warn(err)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:534
				// _ = "end of CoverTab[197086]"
			} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:535
				_go_fuzz_dep_.CoverTab[197087]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:535
				// _ = "end of CoverTab[197087]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:535
			}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:535
			// _ = "end of CoverTab[197084]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:535
			_go_fuzz_dep_.CoverTab[197085]++
												break
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:536
			// _ = "end of CoverTab[197085]"
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:537
		// _ = "end of CoverTab[197083]"

	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:539
	// _ = "end of CoverTab[197082]"
}

func (m *Manager) PlmnIDNciToCGI(plmnID uint64, nci uint64) string {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:542
	_go_fuzz_dep_.CoverTab[197088]++
										cgi := strconv.FormatInt(int64(plmnID<<36|(nci&0xfffffffff)), 16)
										if m.topoIDsEnabled {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:544
		_go_fuzz_dep_.CoverTab[197090]++
											cgi = cgi[0:6] + cgi[14:15] + cgi[12:14] + cgi[10:12] + cgi[8:10] + cgi[6:8]
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:545
		// _ = "end of CoverTab[197090]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:546
		_go_fuzz_dep_.CoverTab[197091]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:546
		// _ = "end of CoverTab[197091]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:546
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:546
	// _ = "end of CoverTab[197088]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:546
	_go_fuzz_dep_.CoverTab[197089]++
										return cgi
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:547
	// _ = "end of CoverTab[197089]"
}

func (m *Manager) CgiFromTopoToIndicationFormat(cgi string) string {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:550
	_go_fuzz_dep_.CoverTab[197092]++
										if !m.topoIDsEnabled {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:551
		_go_fuzz_dep_.CoverTab[197094]++
											cgi = cgi[0:6] + cgi[13:15] + cgi[11:13] + cgi[9:11] + cgi[7:9] + cgi[6:7]
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:552
		// _ = "end of CoverTab[197094]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:553
		_go_fuzz_dep_.CoverTab[197095]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:553
		// _ = "end of CoverTab[197095]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:553
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:553
	// _ = "end of CoverTab[197092]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:553
	_go_fuzz_dep_.CoverTab[197093]++
										return cgi
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:554
	// _ = "end of CoverTab[197093]"
}

func drawWithLine(word string, length int) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:557
	_go_fuzz_dep_.CoverTab[197096]++
										wordLength := len(word)
										diff := length - wordLength
										info := ""
										if diff == length {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:561
		_go_fuzz_dep_.CoverTab[197098]++
											for i := 0; i < diff; i++ {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:562
			_go_fuzz_dep_.CoverTab[197099]++
												info = info + "-"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:563
			// _ = "end of CoverTab[197099]"
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:564
		// _ = "end of CoverTab[197098]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:565
		_go_fuzz_dep_.CoverTab[197100]++
											info = " " + word + " "
											diff -= 2
											for i := 0; i < diff/2; i++ {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:568
			_go_fuzz_dep_.CoverTab[197102]++
												info = "-" + info + "-"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:569
			// _ = "end of CoverTab[197102]"
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:570
		// _ = "end of CoverTab[197100]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:570
		_go_fuzz_dep_.CoverTab[197101]++
											if diff%2 != 0 {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:571
			_go_fuzz_dep_.CoverTab[197103]++
												info = info + "-"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:572
			// _ = "end of CoverTab[197103]"
		} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:573
			_go_fuzz_dep_.CoverTab[197104]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:573
			// _ = "end of CoverTab[197104]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:573
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:573
		// _ = "end of CoverTab[197101]"
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:574
	// _ = "end of CoverTab[197096]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:574
	_go_fuzz_dep_.CoverTab[197097]++
										log.Debug(info)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:575
	// _ = "end of CoverTab[197097]"
}

func compareLengths() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:578
	_go_fuzz_dep_.CoverTab[197105]++
										temp := nodesLogLen
										if nodesLogLen < policiesLogLen {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:580
		_go_fuzz_dep_.CoverTab[197107]++
											temp = policiesLogLen
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:581
		// _ = "end of CoverTab[197107]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:582
		_go_fuzz_dep_.CoverTab[197108]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:582
		// _ = "end of CoverTab[197108]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:582
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:582
	// _ = "end of CoverTab[197105]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:582
	_go_fuzz_dep_.CoverTab[197106]++
										logLength = temp
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:583
	// _ = "end of CoverTab[197106]"
}

//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:584
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/manager/manager.go:584
var _ = _go_fuzz_dep_.CoverTab
