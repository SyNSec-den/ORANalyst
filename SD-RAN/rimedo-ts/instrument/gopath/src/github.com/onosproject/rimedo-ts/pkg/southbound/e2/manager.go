// SPDX-FileCopyrightText: 2019-present Open Networking Foundation <info@opennetworking.org>
// SPDX-FileCopyrightText: 2019-present Rimedo Labs
//
// SPDX-License-Identifier: Apache-2.0
// Created by RIMEDO-Labs team
// based on onosproject/onos-mho/pkg/southbound/e2/manager.go

//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:8
package e2

//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:8
import (
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:8
)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:8
import (
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:8
	_atomic_ "sync/atomic"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:8
)

import (
	"context"
	"fmt"
	"strings"

	prototypes "github.com/gogo/protobuf/types"
	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/pdubuilder"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-mho-go"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-mho/pkg/broker"
	e2client "github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1"
	"github.com/onosproject/rimedo-ts/pkg/mho"
	"github.com/onosproject/rimedo-ts/pkg/monitoring"
	"github.com/onosproject/rimedo-ts/pkg/rnib"
	"google.golang.org/protobuf/proto"
)

var log = logging.GetLogger("rimedo-ts", "e2", "manager")

const (
	oid = "1.3.6.1.4.1.53148.1.2.2.101"
)

type Options struct {
	AppID		string
	E2tAddress	string
	E2tPort		int
	TopoAddress	string
	TopoPort	int
	SMName		string
	SMVersion	string
}

func NewManager(options Options, indCh chan *mho.E2NodeIndication, ctrlReqChs map[string]chan *e2api.ControlMessage) (Manager, error) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:46
	_go_fuzz_dep_.CoverTab[196701]++

										smName := e2client.ServiceModelName(options.SMName)
										smVer := e2client.ServiceModelVersion(options.SMVersion)
										appID := e2client.AppID(options.AppID)
										e2Client := e2client.NewClient(
		e2client.WithAppID(appID),
		e2client.WithServiceModel(smName, smVer),
		e2client.WithE2TAddress(options.E2tAddress, options.E2tPort),
	)

	rnibOptions := rnib.Options{
		TopoAddress:	options.TopoAddress,
		TopoPort:	options.TopoPort,
	}

	rnibClient, err := rnib.NewClient(rnibOptions)
	if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:63
		_go_fuzz_dep_.CoverTab[196703]++
											return Manager{}, err
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:64
		// _ = "end of CoverTab[196703]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:65
		_go_fuzz_dep_.CoverTab[196704]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:65
		// _ = "end of CoverTab[196704]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:65
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:65
	// _ = "end of CoverTab[196701]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:65
	_go_fuzz_dep_.CoverTab[196702]++

										return Manager{
		e2client:	e2Client,
		rnibClient:	rnibClient,
		streams:	broker.NewBroker(),
		indCh:		indCh,
		ctrlReqChs:	ctrlReqChs,
		smModelName:	smName,
	}, nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:74
	// _ = "end of CoverTab[196702]"
}

type Manager struct {
	e2client	e2client.Client
	rnibClient	rnib.Client
	streams		broker.Broker
	indCh		chan *mho.E2NodeIndication
	ctrlReqChs	map[string]chan *e2api.ControlMessage
	smModelName	e2client.ServiceModelName
}

func (m *Manager) Start() error {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:86
	_go_fuzz_dep_.CoverTab[196705]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:86
	_curRoutineNum189_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:86
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum189_)
										go func() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:87
		_go_fuzz_dep_.CoverTab[196707]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:87
		defer func() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:87
			_go_fuzz_dep_.CoverTab[196708]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:87
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum189_)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:87
			// _ = "end of CoverTab[196708]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:87
		}()
											ctx, cancel := context.WithCancel(context.Background())
											defer cancel()
											err := m.watchE2Connections(ctx)
											if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:91
			_go_fuzz_dep_.CoverTab[196709]++
												return
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:92
			// _ = "end of CoverTab[196709]"
		} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:93
			_go_fuzz_dep_.CoverTab[196710]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:93
			// _ = "end of CoverTab[196710]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:93
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:93
		// _ = "end of CoverTab[196707]"
	}()
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:94
	// _ = "end of CoverTab[196705]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:94
	_go_fuzz_dep_.CoverTab[196706]++

										return nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:96
	// _ = "end of CoverTab[196706]"
}

func (m *Manager) watchE2Connections(ctx context.Context) error {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:99
	_go_fuzz_dep_.CoverTab[196711]++
										ch := make(chan topoapi.Event)
										err := m.rnibClient.WatchE2Connections(ctx, ch)
										if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:102
		_go_fuzz_dep_.CoverTab[196714]++
											log.Warn(err)
											return err
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:104
		// _ = "end of CoverTab[196714]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:105
		_go_fuzz_dep_.CoverTab[196715]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:105
		// _ = "end of CoverTab[196715]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:105
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:105
	// _ = "end of CoverTab[196711]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:105
	_go_fuzz_dep_.CoverTab[196712]++

										for topoEvent := range ch {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:107
		_go_fuzz_dep_.CoverTab[196716]++
											if topoEvent.Type == topoapi.EventType_ADDED || func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:108
			_go_fuzz_dep_.CoverTab[196717]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:108
			return topoEvent.Type == topoapi.EventType_NONE
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:108
			// _ = "end of CoverTab[196717]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:108
		}() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:108
			_go_fuzz_dep_.CoverTab[196718]++
												relation := topoEvent.Object.Obj.(*topoapi.Object_Relation)
												e2NodeID := relation.Relation.TgtEntityID
												m.ctrlReqChs[string(e2NodeID)] = make(chan *e2api.ControlMessage)

												triggers := make(map[e2sm_mho.MhoTriggerType]bool)
												triggers[e2sm_mho.MhoTriggerType_MHO_TRIGGER_TYPE_PERIODIC] = true
												triggers[e2sm_mho.MhoTriggerType_MHO_TRIGGER_TYPE_UPON_RCV_MEAS_REPORT] = true
												triggers[e2sm_mho.MhoTriggerType_MHO_TRIGGER_TYPE_UPON_CHANGE_RRC_STATUS] = true

												for triggerType, enabled := range triggers {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:118
				_go_fuzz_dep_.CoverTab[196720]++
													if enabled {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:119
					_go_fuzz_dep_.CoverTab[196721]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:119
					_curRoutineNum191_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:119
					_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum191_)
														go func(triggerType e2sm_mho.MhoTriggerType) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:120
						_go_fuzz_dep_.CoverTab[196722]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:120
						defer func() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:120
							_go_fuzz_dep_.CoverTab[196723]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:120
							_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum191_)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:120
							// _ = "end of CoverTab[196723]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:120
						}()
															_ = m.createSubscription(ctx, e2NodeID, triggerType)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:121
						// _ = "end of CoverTab[196722]"
					}(triggerType)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:122
					// _ = "end of CoverTab[196721]"
				} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:123
					_go_fuzz_dep_.CoverTab[196724]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:123
					// _ = "end of CoverTab[196724]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:123
				}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:123
				// _ = "end of CoverTab[196720]"
			}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:124
			// _ = "end of CoverTab[196718]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:124
			_go_fuzz_dep_.CoverTab[196719]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:124
			_curRoutineNum190_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:124
			_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum190_)
												go m.watchMHOChanges(ctx, e2NodeID)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:125
			// _ = "end of CoverTab[196719]"
		} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:126
			_go_fuzz_dep_.CoverTab[196725]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:126
			// _ = "end of CoverTab[196725]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:126
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:126
		// _ = "end of CoverTab[196716]"
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:127
	// _ = "end of CoverTab[196712]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:127
	_go_fuzz_dep_.CoverTab[196713]++

										return nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:129
	// _ = "end of CoverTab[196713]"
}

func (m *Manager) watchMHOChanges(ctx context.Context, e2nodeID topoapi.ID) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:132
	_go_fuzz_dep_.CoverTab[196726]++

										for ctrlReqMsg := range m.ctrlReqChs[string(e2nodeID)] {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:134
		_go_fuzz_dep_.CoverTab[196727]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:134
		_curRoutineNum192_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:134
		_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum192_)
											go func(ctrlReqMsg *e2api.ControlMessage) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:135
			_go_fuzz_dep_.CoverTab[196728]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:135
			defer func() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:135
				_go_fuzz_dep_.CoverTab[196729]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:135
				_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum192_)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:135
				// _ = "end of CoverTab[196729]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:135
			}()
												node := m.e2client.Node(e2client.NodeID(e2nodeID))
												_, _ = node.Control(ctx, ctrlReqMsg)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:137
			// _ = "end of CoverTab[196728]"
		}(ctrlReqMsg)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:138
		// _ = "end of CoverTab[196727]"
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:139
	// _ = "end of CoverTab[196726]"
}

func (m *Manager) createSubscription(ctx context.Context, e2nodeID topoapi.ID, triggerType e2sm_mho.MhoTriggerType) error {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:142
	_go_fuzz_dep_.CoverTab[196730]++
										eventTriggerData, err := m.createEventTrigger(triggerType)
										if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:144
		_go_fuzz_dep_.CoverTab[196737]++
											return err
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:145
		// _ = "end of CoverTab[196737]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:146
		_go_fuzz_dep_.CoverTab[196738]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:146
		// _ = "end of CoverTab[196738]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:146
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:146
	// _ = "end of CoverTab[196730]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:146
	_go_fuzz_dep_.CoverTab[196731]++

										actions := m.createSubscriptionActions()

										aspects, err := m.rnibClient.GetE2NodeAspects(ctx, e2nodeID)
										if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:151
		_go_fuzz_dep_.CoverTab[196739]++
											return err
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:152
		// _ = "end of CoverTab[196739]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:153
		_go_fuzz_dep_.CoverTab[196740]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:153
		// _ = "end of CoverTab[196740]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:153
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:153
	// _ = "end of CoverTab[196731]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:153
	_go_fuzz_dep_.CoverTab[196732]++

										_, err = m.getRanFunction(aspects.ServiceModels)
										if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:156
		_go_fuzz_dep_.CoverTab[196741]++
											return err
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:157
		// _ = "end of CoverTab[196741]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:158
		_go_fuzz_dep_.CoverTab[196742]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:158
		// _ = "end of CoverTab[196742]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:158
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:158
	// _ = "end of CoverTab[196732]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:158
	_go_fuzz_dep_.CoverTab[196733]++

										ch := make(chan e2api.Indication)
										node := m.e2client.Node(e2client.NodeID(e2nodeID))
										subName := fmt.Sprintf("rimedo-ts-subscription-%s", triggerType)
										subSpec := e2api.SubscriptionSpec{
		Actions:	actions,
		EventTrigger: e2api.EventTrigger{
			Payload: eventTriggerData,
		},
	}

	channelID, err := node.Subscribe(ctx, subName, subSpec, ch)
	if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:171
		_go_fuzz_dep_.CoverTab[196743]++
											log.Warn(err)
											return err
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:173
		// _ = "end of CoverTab[196743]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:174
		_go_fuzz_dep_.CoverTab[196744]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:174
		// _ = "end of CoverTab[196744]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:174
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:174
	// _ = "end of CoverTab[196733]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:174
	_go_fuzz_dep_.CoverTab[196734]++

										streamReader, err := m.streams.OpenReader(ctx, node, subName, channelID, subSpec)
										if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:177
		_go_fuzz_dep_.CoverTab[196745]++
											return err
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:178
		// _ = "end of CoverTab[196745]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:179
		_go_fuzz_dep_.CoverTab[196746]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:179
		// _ = "end of CoverTab[196746]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:179
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:179
	// _ = "end of CoverTab[196734]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:179
	_go_fuzz_dep_.CoverTab[196735]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:179
	_curRoutineNum193_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:179
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum193_)
										go m.sendIndicationOnStream(streamReader.StreamID(), ch)

										monitor := monitoring.NewMonitor(streamReader, e2nodeID, m.indCh, triggerType)

										err = monitor.Start(ctx)
										if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:185
		_go_fuzz_dep_.CoverTab[196747]++
											log.Warn(err)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:186
		// _ = "end of CoverTab[196747]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:187
		_go_fuzz_dep_.CoverTab[196748]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:187
		// _ = "end of CoverTab[196748]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:187
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:187
	// _ = "end of CoverTab[196735]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:187
	_go_fuzz_dep_.CoverTab[196736]++

										return nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:189
	// _ = "end of CoverTab[196736]"
}

func (m *Manager) getRanFunction(serviceModelsInfo map[string]*topoapi.ServiceModelInfo) (*topoapi.MHORanFunction, error) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:192
	_go_fuzz_dep_.CoverTab[196749]++
										for _, sm := range serviceModelsInfo {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:193
		_go_fuzz_dep_.CoverTab[196751]++
											smName := strings.ToLower(sm.Name)
											if smName == string(m.smModelName) && func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:195
			_go_fuzz_dep_.CoverTab[196752]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:195
			return sm.OID == oid
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:195
			// _ = "end of CoverTab[196752]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:195
		}() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:195
			_go_fuzz_dep_.CoverTab[196753]++
												mhoRanFunction := &topoapi.MHORanFunction{}
												for _, ranFunction := range sm.RanFunctions {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:197
				_go_fuzz_dep_.CoverTab[196754]++
													if ranFunction.TypeUrl == ranFunction.GetTypeUrl() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:198
					_go_fuzz_dep_.CoverTab[196755]++
														err := prototypes.UnmarshalAny(ranFunction, mhoRanFunction)
														if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:200
						_go_fuzz_dep_.CoverTab[196757]++
															return nil, err
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:201
						// _ = "end of CoverTab[196757]"
					} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:202
						_go_fuzz_dep_.CoverTab[196758]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:202
						// _ = "end of CoverTab[196758]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:202
					}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:202
					// _ = "end of CoverTab[196755]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:202
					_go_fuzz_dep_.CoverTab[196756]++
														return mhoRanFunction, nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:203
					// _ = "end of CoverTab[196756]"
				} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:204
					_go_fuzz_dep_.CoverTab[196759]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:204
					// _ = "end of CoverTab[196759]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:204
				}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:204
				// _ = "end of CoverTab[196754]"
			}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:205
			// _ = "end of CoverTab[196753]"
		} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:206
			_go_fuzz_dep_.CoverTab[196760]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:206
			// _ = "end of CoverTab[196760]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:206
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:206
		// _ = "end of CoverTab[196751]"
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:207
	// _ = "end of CoverTab[196749]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:207
	_go_fuzz_dep_.CoverTab[196750]++
										return nil, errors.New(errors.NotFound, "cannot retrieve ran functions")
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:208
	// _ = "end of CoverTab[196750]"

}

func (m *Manager) createEventTrigger(triggerType e2sm_mho.MhoTriggerType) ([]byte, error) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:212
	_go_fuzz_dep_.CoverTab[196761]++
										var reportPeriodMs int32
										reportingPeriod := 1000
										if triggerType == e2sm_mho.MhoTriggerType_MHO_TRIGGER_TYPE_PERIODIC {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:215
		_go_fuzz_dep_.CoverTab[196766]++
											reportPeriodMs = int32(reportingPeriod)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:216
		// _ = "end of CoverTab[196766]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:217
		_go_fuzz_dep_.CoverTab[196767]++
											reportPeriodMs = 0
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:218
		// _ = "end of CoverTab[196767]"
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:219
	// _ = "end of CoverTab[196761]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:219
	_go_fuzz_dep_.CoverTab[196762]++
										e2smRcEventTriggerDefinition, err := pdubuilder.CreateE2SmMhoEventTriggerDefinition(triggerType)
										if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:221
		_go_fuzz_dep_.CoverTab[196768]++
											return []byte{}, err
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:222
		// _ = "end of CoverTab[196768]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:223
		_go_fuzz_dep_.CoverTab[196769]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:223
		// _ = "end of CoverTab[196769]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:223
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:223
	// _ = "end of CoverTab[196762]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:223
	_go_fuzz_dep_.CoverTab[196763]++
										e2smRcEventTriggerDefinition.GetEventDefinitionFormats().GetEventDefinitionFormat1().SetReportingPeriodInMs(reportPeriodMs)

										err = e2smRcEventTriggerDefinition.Validate()
										if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:227
		_go_fuzz_dep_.CoverTab[196770]++
											return []byte{}, err
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:228
		// _ = "end of CoverTab[196770]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:229
		_go_fuzz_dep_.CoverTab[196771]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:229
		// _ = "end of CoverTab[196771]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:229
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:229
	// _ = "end of CoverTab[196763]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:229
	_go_fuzz_dep_.CoverTab[196764]++

										protoBytes, err := proto.Marshal(e2smRcEventTriggerDefinition)
										if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:232
		_go_fuzz_dep_.CoverTab[196772]++
											return []byte{}, err
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:233
		// _ = "end of CoverTab[196772]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:234
		_go_fuzz_dep_.CoverTab[196773]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:234
		// _ = "end of CoverTab[196773]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:234
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:234
	// _ = "end of CoverTab[196764]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:234
	_go_fuzz_dep_.CoverTab[196765]++

										return protoBytes, err
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:236
	// _ = "end of CoverTab[196765]"
}

func (m *Manager) createSubscriptionActions() []e2api.Action {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:239
	_go_fuzz_dep_.CoverTab[196774]++
										actions := make([]e2api.Action, 0)
										action := &e2api.Action{
		ID:	int32(0),
		Type:	e2api.ActionType_ACTION_TYPE_REPORT,
		SubsequentAction: &e2api.SubsequentAction{
			Type:		e2api.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
			TimeToWait:	e2api.TimeToWait_TIME_TO_WAIT_ZERO,
		},
	}
										actions = append(actions, *action)
										return actions
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:250
	// _ = "end of CoverTab[196774]"
}

func (m *Manager) sendIndicationOnStream(streamID broker.StreamID, ch chan e2api.Indication) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:253
	_go_fuzz_dep_.CoverTab[196775]++
										streamWriter, err := m.streams.GetWriter(streamID)
										if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:255
		_go_fuzz_dep_.CoverTab[196777]++
											log.Error(err)
											return
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:257
		// _ = "end of CoverTab[196777]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:258
		_go_fuzz_dep_.CoverTab[196778]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:258
		// _ = "end of CoverTab[196778]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:258
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:258
	// _ = "end of CoverTab[196775]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:258
	_go_fuzz_dep_.CoverTab[196776]++

										for msg := range ch {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:260
		_go_fuzz_dep_.CoverTab[196779]++
											err := streamWriter.Send(msg)
											if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:262
			_go_fuzz_dep_.CoverTab[196780]++
												log.Warn(err)
												return
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:264
			// _ = "end of CoverTab[196780]"
		} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:265
			_go_fuzz_dep_.CoverTab[196781]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:265
			// _ = "end of CoverTab[196781]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:265
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:265
		// _ = "end of CoverTab[196779]"
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:266
	// _ = "end of CoverTab[196776]"
}

func (m *Manager) GetCellTypes(ctx context.Context) map[string]rnib.Cell {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:269
	_go_fuzz_dep_.CoverTab[196782]++
										cellTypes, err := m.rnibClient.GetCellTypes(ctx)
										if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:271
		_go_fuzz_dep_.CoverTab[196784]++
											log.Warn(err)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:272
		// _ = "end of CoverTab[196784]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:273
		_go_fuzz_dep_.CoverTab[196785]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:273
		// _ = "end of CoverTab[196785]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:273
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:273
	// _ = "end of CoverTab[196782]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:273
	_go_fuzz_dep_.CoverTab[196783]++
										return cellTypes
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:274
	// _ = "end of CoverTab[196783]"
}

func (m *Manager) SetCellType(ctx context.Context, cellID string, cellType string) error {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:277
	_go_fuzz_dep_.CoverTab[196786]++
										err := m.rnibClient.SetCellType(ctx, cellID, cellType)
										if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:279
		_go_fuzz_dep_.CoverTab[196788]++
											log.Warn(err)
											return err
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:281
		// _ = "end of CoverTab[196788]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:282
		_go_fuzz_dep_.CoverTab[196789]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:282
		// _ = "end of CoverTab[196789]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:282
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:282
	// _ = "end of CoverTab[196786]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:282
	_go_fuzz_dep_.CoverTab[196787]++
										return nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:283
	// _ = "end of CoverTab[196787]"
}

//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:284
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/southbound/e2/manager.go:284
var _ = _go_fuzz_dep_.CoverTab
