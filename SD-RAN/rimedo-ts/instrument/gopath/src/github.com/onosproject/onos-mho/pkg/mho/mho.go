// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:5
package mho

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:5
)

import (
	"context"
	"reflect"
	"strconv"
	"sync"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-mho-go"
	e2sm_v2_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-v2-ies"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	appConfig "github.com/onosproject/onos-mho/pkg/config"
	"github.com/onosproject/onos-mho/pkg/store"
	"github.com/onosproject/rrm-son-lib/pkg/handover"
	rrmid "github.com/onosproject/rrm-son-lib/pkg/model/id"
	"google.golang.org/protobuf/proto"
)

const (
	ControlPriority = 10
)

var log = logging.GetLogger()

type UeData struct {
	UeID	string	// assuming that string carries decimal number
	//ToDo - stop ignoring it once we'veset up UeID treatment all over the SD-RAN
	//UeIDtype      string // represents Gnb, Enb, ngEnb and etc..
	E2NodeID	string
	CGI		*e2sm_v2_ies.Cgi
	CGIString	string
	RrcState	string
	FiveQI		int32
	RsrpServing	int32
	RsrpNeighbors	map[string]int32
}

type CellData struct {
	CGIString		string
	CumulativeHandoversIn	int
	CumulativeHandoversOut	int
	Ues			map[string]*UeData
}

type E2NodeIndication struct {
	NodeID		string
	TriggerType	e2sm_mho.MhoTriggerType
	IndMsg		e2api.Indication
}

// Ctrl is the controller for MHO
type Ctrl struct {
	IndChan		chan *E2NodeIndication
	CtrlReqChans	map[string]chan *e2api.ControlMessage
	HoCtrl		*HandOverController
	ueStore		store.Store
	cellStore	store.Store
	mu		sync.RWMutex
	cells		map[string]*CellData
}

// NewMhoController returns the struct for MHO logic
func NewMhoController(cfg appConfig.Config, indChan chan *E2NodeIndication, ctrlReqChs map[string]chan *e2api.ControlMessage, ueStore store.Store, cellStore store.Store) *Ctrl {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:68
	_go_fuzz_dep_.CoverTab[194501]++
												log.Info("Init MhoController")
												return &Ctrl{
		IndChan:	indChan,
		CtrlReqChans:	ctrlReqChs,
		HoCtrl:		NewHandOverController(cfg),
		ueStore:	ueStore,
		cellStore:	cellStore,
		cells:		make(map[string]*CellData),
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:77
	// _ = "end of CoverTab[194501]"
}

// Run starts to listen Indication message and then save the result to its struct
func (c *Ctrl) Run(ctx context.Context) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:81
	_go_fuzz_dep_.CoverTab[194502]++
												log.Info("Start MhoController")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:82
	_curRoutineNum180_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:82
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum180_)
												go c.HoCtrl.Run()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:83
	_curRoutineNum181_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:83
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum181_)
												go c.listenIndChan(ctx)
												c.listenHandOver(ctx)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:85
	// _ = "end of CoverTab[194502]"
}

func (c *Ctrl) listenIndChan(ctx context.Context) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:88
	_go_fuzz_dep_.CoverTab[194503]++
												var err error
												for indMsg := range c.IndChan {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:90
		_go_fuzz_dep_.CoverTab[194504]++

													indHeaderByte := indMsg.IndMsg.Header
													indMessageByte := indMsg.IndMsg.Payload
													e2NodeID := indMsg.NodeID

													indHeader := e2sm_mho.E2SmMhoIndicationHeader{}
													if err = proto.Unmarshal(indHeaderByte, &indHeader); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:97
			_go_fuzz_dep_.CoverTab[194506]++
														indMessage := e2sm_mho.E2SmMhoIndicationMessage{}
														if err = proto.Unmarshal(indMessageByte, &indMessage); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:99
				_go_fuzz_dep_.CoverTab[194507]++
															switch x := indMessage.E2SmMhoIndicationMessage.(type) {
				case *e2sm_mho.E2SmMhoIndicationMessage_IndicationMessageFormat1:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:101
					_go_fuzz_dep_.CoverTab[194508]++
																if indMsg.TriggerType == e2sm_mho.MhoTriggerType_MHO_TRIGGER_TYPE_UPON_RCV_MEAS_REPORT {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:102
						_go_fuzz_dep_.CoverTab[194511]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:102
						_curRoutineNum182_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:102
						_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum182_)
																	go c.handleMeasReport(ctx, indHeader.GetIndicationHeaderFormat1(), indMessage.GetIndicationMessageFormat1(), e2NodeID)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:103
						// _ = "end of CoverTab[194511]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:104
						_go_fuzz_dep_.CoverTab[194512]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:104
						if indMsg.TriggerType == e2sm_mho.MhoTriggerType_MHO_TRIGGER_TYPE_PERIODIC {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:104
							_go_fuzz_dep_.CoverTab[194513]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:104
							_curRoutineNum183_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:104
							_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum183_)
																		go c.handlePeriodicReport(ctx, indHeader.GetIndicationHeaderFormat1(), indMessage.GetIndicationMessageFormat1(), e2NodeID)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:105
							// _ = "end of CoverTab[194513]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:106
							_go_fuzz_dep_.CoverTab[194514]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:106
							// _ = "end of CoverTab[194514]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:106
						}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:106
						// _ = "end of CoverTab[194512]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:106
					}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:106
					// _ = "end of CoverTab[194508]"
				case *e2sm_mho.E2SmMhoIndicationMessage_IndicationMessageFormat2:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:107
					_go_fuzz_dep_.CoverTab[194509]++
																go c.handleRrcState(ctx, indHeader.GetIndicationHeaderFormat1(), indMessage.GetIndicationMessageFormat2())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:108
					// _ = "end of CoverTab[194509]"
				default:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:109
					_go_fuzz_dep_.CoverTab[194510]++
																log.Warnf("Unknown MHO indication message format, indication message: %v", x)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:110
					// _ = "end of CoverTab[194510]"
				}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:111
				// _ = "end of CoverTab[194507]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:112
				_go_fuzz_dep_.CoverTab[194515]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:112
				// _ = "end of CoverTab[194515]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:112
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:112
			// _ = "end of CoverTab[194506]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:113
			_go_fuzz_dep_.CoverTab[194516]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:113
			// _ = "end of CoverTab[194516]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:113
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:113
		// _ = "end of CoverTab[194504]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:113
		_go_fuzz_dep_.CoverTab[194505]++
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:114
			_go_fuzz_dep_.CoverTab[194517]++
														log.Error(err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:115
			// _ = "end of CoverTab[194517]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:116
			_go_fuzz_dep_.CoverTab[194518]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:116
			// _ = "end of CoverTab[194518]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:116
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:116
		// _ = "end of CoverTab[194505]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:117
	// _ = "end of CoverTab[194503]"
}

func (c *Ctrl) listenHandOver(ctx context.Context) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:120
	_go_fuzz_dep_.CoverTab[194519]++
												for hoDecision := range c.HoCtrl.HandoverHandler.Chans.OutputChan {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:121
		_go_fuzz_dep_.CoverTab[194520]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:121
		_curRoutineNum184_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:121
		_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum184_)
													go func(hoDecision handover.A3HandoverDecision) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:122
			_go_fuzz_dep_.CoverTab[194521]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:122
			defer func() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:122
				_go_fuzz_dep_.CoverTab[194522]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:122
				_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum184_)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:122
				// _ = "end of CoverTab[194522]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:122
			}()
														if err := c.control(ctx, hoDecision); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:123
				_go_fuzz_dep_.CoverTab[194523]++
															log.Error(err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:124
				// _ = "end of CoverTab[194523]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:125
				_go_fuzz_dep_.CoverTab[194524]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:125
				// _ = "end of CoverTab[194524]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:125
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:125
			// _ = "end of CoverTab[194521]"
		}(hoDecision)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:126
		// _ = "end of CoverTab[194520]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:127
	// _ = "end of CoverTab[194519]"
}

func (c *Ctrl) handlePeriodicReport(ctx context.Context, header *e2sm_mho.E2SmMhoIndicationHeaderFormat1, message *e2sm_mho.E2SmMhoIndicationMessageFormat1, e2NodeID string) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:130
	_go_fuzz_dep_.CoverTab[194525]++
												c.mu.Lock()
												defer c.mu.Unlock()
												ueID, err := GetUeID(message.GetUeId())
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:134
		_go_fuzz_dep_.CoverTab[194529]++
													log.Errorf("handlePeriodicReport() couldn't extract UeID: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:135
		// _ = "end of CoverTab[194529]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:136
		_go_fuzz_dep_.CoverTab[194530]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:136
		// _ = "end of CoverTab[194530]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:136
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:136
	// _ = "end of CoverTab[194525]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:136
	_go_fuzz_dep_.CoverTab[194526]++
												cgi := getCGIFromIndicationHeader(header)
												log.Infof("rx periodic ueID:%v cgi:%v", ueID, cgi)

	// get ue from store (create if it does not exist)
	var ueData *UeData
	newUe := false

	ueData = c.getUe(ctx, strconv.Itoa(int(ueID)))
	if ueData == nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:145
		_go_fuzz_dep_.CoverTab[194531]++

													ueData = c.createUe(ctx, strconv.Itoa(int(ueID)))
													c.attachUe(ctx, ueData, cgi)
													newUe = true
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:149
		// _ = "end of CoverTab[194531]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:150
		_go_fuzz_dep_.CoverTab[194532]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:150
		if ueData.CGIString != cgi {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:150
			_go_fuzz_dep_.CoverTab[194533]++
														return
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:151
			// _ = "end of CoverTab[194533]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:152
			_go_fuzz_dep_.CoverTab[194534]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:152
			// _ = "end of CoverTab[194534]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:152
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:152
		// _ = "end of CoverTab[194532]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:152
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:152
	// _ = "end of CoverTab[194526]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:152
	_go_fuzz_dep_.CoverTab[194527]++

												rsrpServing, rsrpNeighbors := getRsrpFromMeasReport(getNciFromCellGlobalID(header.GetCgi()), message.MeasReport)

												if !newUe && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:156
		_go_fuzz_dep_.CoverTab[194535]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:156
		return rsrpServing == ueData.RsrpServing
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:156
		// _ = "end of CoverTab[194535]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:156
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:156
		_go_fuzz_dep_.CoverTab[194536]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:156
		return reflect.DeepEqual(rsrpNeighbors, ueData.RsrpNeighbors)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:156
		// _ = "end of CoverTab[194536]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:156
	}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:156
		_go_fuzz_dep_.CoverTab[194537]++
													return
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:157
		// _ = "end of CoverTab[194537]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:158
		_go_fuzz_dep_.CoverTab[194538]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:158
		// _ = "end of CoverTab[194538]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:158
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:158
	// _ = "end of CoverTab[194527]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:158
	_go_fuzz_dep_.CoverTab[194528]++

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:161
	ueData.RsrpServing, ueData.RsrpNeighbors = rsrpServing, rsrpNeighbors
												c.setUe(ctx, ueData)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:162
	// _ = "end of CoverTab[194528]"

}

func (c *Ctrl) handleMeasReport(ctx context.Context, header *e2sm_mho.E2SmMhoIndicationHeaderFormat1, message *e2sm_mho.E2SmMhoIndicationMessageFormat1, e2NodeID string) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:166
	_go_fuzz_dep_.CoverTab[194539]++
												c.mu.Lock()
												defer c.mu.Unlock()
												ueID, err := GetUeID(message.GetUeId())
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:170
		_go_fuzz_dep_.CoverTab[194543]++
													log.Errorf("handleMeasReport() couldn't extract UeID: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:171
		// _ = "end of CoverTab[194543]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:172
		_go_fuzz_dep_.CoverTab[194544]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:172
		// _ = "end of CoverTab[194544]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:172
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:172
	// _ = "end of CoverTab[194539]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:172
	_go_fuzz_dep_.CoverTab[194540]++
												cgi := getCGIFromIndicationHeader(header)
												log.Infof("rx a3 ueID:%v cgi:%v", ueID, cgi)

	// get ue from store (create if it does not exist)
	var ueData *UeData

	ueData = c.getUe(ctx, strconv.Itoa(int(ueID)))
	if ueData == nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:180
		_go_fuzz_dep_.CoverTab[194545]++

													ueData = c.createUe(ctx, strconv.Itoa(int(ueID)))
													c.attachUe(ctx, ueData, cgi)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:183
		// _ = "end of CoverTab[194545]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:184
		_go_fuzz_dep_.CoverTab[194546]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:184
		if ueData.CGIString != cgi {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:184
			_go_fuzz_dep_.CoverTab[194547]++
														return
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:185
			// _ = "end of CoverTab[194547]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:186
			_go_fuzz_dep_.CoverTab[194548]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:186
			// _ = "end of CoverTab[194548]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:186
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:186
		// _ = "end of CoverTab[194546]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:186
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:186
	// _ = "end of CoverTab[194540]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:186
	_go_fuzz_dep_.CoverTab[194541]++

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:189
	ueData.CGI = header.GetCgi()
												ueData.E2NodeID = e2NodeID

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:193
	ueData.RsrpServing, ueData.RsrpNeighbors = getRsrpFromMeasReport(getNciFromCellGlobalID(header.GetCgi()), message.MeasReport)

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:196
	log.Debugf("Going to update 5QI (%v) for UE %v", ueData.FiveQI, ueID)
	for _, item := range message.GetMeasReport() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:197
		_go_fuzz_dep_.CoverTab[194549]++
													if item.GetFiveQi() != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:198
			_go_fuzz_dep_.CoverTab[194550]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:198
			return item.GetFiveQi().GetValue() > -1
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:198
			// _ = "end of CoverTab[194550]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:198
		}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:198
			_go_fuzz_dep_.CoverTab[194551]++
														ueData.FiveQI = item.GetFiveQi().GetValue()
														log.Debugf("Obtained 5QI value %v for UE %v", ueData.FiveQI, ueID)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:200
			// _ = "end of CoverTab[194551]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:201
			_go_fuzz_dep_.CoverTab[194552]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:201
			// _ = "end of CoverTab[194552]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:201
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:201
		// _ = "end of CoverTab[194549]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:202
	// _ = "end of CoverTab[194541]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:202
	_go_fuzz_dep_.CoverTab[194542]++

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:205
	c.setUe(ctx, ueData)

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:208
	c.HoCtrl.Input(ctx, header, message)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:208
	// _ = "end of CoverTab[194542]"

}

func (c *Ctrl) handleRrcState(ctx context.Context, header *e2sm_mho.E2SmMhoIndicationHeaderFormat1, message *e2sm_mho.E2SmMhoIndicationMessageFormat2) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:212
	_go_fuzz_dep_.CoverTab[194553]++
												c.mu.Lock()
												defer c.mu.Unlock()
												ueID, err := GetUeID(message.GetUeId())
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:216
		_go_fuzz_dep_.CoverTab[194556]++
													log.Errorf("handleRrcState() couldn't extract UeID: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:217
		// _ = "end of CoverTab[194556]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:218
		_go_fuzz_dep_.CoverTab[194557]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:218
		// _ = "end of CoverTab[194557]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:218
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:218
	// _ = "end of CoverTab[194553]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:218
	_go_fuzz_dep_.CoverTab[194554]++
												cgi := getCGIFromIndicationHeader(header)
												log.Infof("rx rrc ueID:%v cgi:%v", ueID, cgi)

	// get ue from store (create if it does not exist)
	var ueData *UeData

	ueData = c.getUe(ctx, strconv.Itoa(int(ueID)))
	if ueData == nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:226
		_go_fuzz_dep_.CoverTab[194558]++

													ueData = c.createUe(ctx, strconv.Itoa(int(ueID)))
													c.attachUe(ctx, ueData, cgi)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:229
		// _ = "end of CoverTab[194558]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:230
		_go_fuzz_dep_.CoverTab[194559]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:230
		if ueData.CGIString != cgi {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:230
			_go_fuzz_dep_.CoverTab[194560]++
														return
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:231
			// _ = "end of CoverTab[194560]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:232
			_go_fuzz_dep_.CoverTab[194561]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:232
			// _ = "end of CoverTab[194561]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:232
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:232
		// _ = "end of CoverTab[194559]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:232
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:232
	// _ = "end of CoverTab[194554]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:232
	_go_fuzz_dep_.CoverTab[194555]++

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:235
	newRrcState := message.GetRrcStatus().String()
												c.setUeRrcState(ctx, ueData, newRrcState, cgi)

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:239
	c.setUe(ctx, ueData)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:239
	// _ = "end of CoverTab[194555]"

}

func (c *Ctrl) control(ctx context.Context, ho handover.A3HandoverDecision) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:243
	_go_fuzz_dep_.CoverTab[194562]++
												c.mu.Lock()
												defer c.mu.Unlock()
												id := ho.UE.GetID().GetID().(rrmid.UEID).IMSI
												ueID := strconv.Itoa(int(id))

												ueData := c.getUe(ctx, ueID)
												if ueData == nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:250
		_go_fuzz_dep_.CoverTab[194564]++
													panic("bad data")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:251
		// _ = "end of CoverTab[194564]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:252
		_go_fuzz_dep_.CoverTab[194565]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:252
		// _ = "end of CoverTab[194565]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:252
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:252
	// _ = "end of CoverTab[194562]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:252
	_go_fuzz_dep_.CoverTab[194563]++

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:255
	targetCGIString := getCgiFromHO(ueData, ho)
												c.doHandover(ctx, ueData, targetCGIString)

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:259
	SendHORequest(ueData, ho, c.CtrlReqChans[ueData.E2NodeID])

												return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:261
	// _ = "end of CoverTab[194563]"

}

func (c *Ctrl) doHandover(ctx context.Context, ueData *UeData, targetCgi string) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:265
	_go_fuzz_dep_.CoverTab[194566]++
												servingCgi := ueData.CGIString
												c.attachUe(ctx, ueData, targetCgi)

												targetCell := c.getCell(ctx, targetCgi)
												targetCell.CumulativeHandoversOut++
												c.setCell(ctx, targetCell)

												servingCell := c.getCell(ctx, servingCgi)
												servingCell.CumulativeHandoversIn++
												c.setCell(ctx, servingCell)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:275
	// _ = "end of CoverTab[194566]"
}

func (c *Ctrl) createUe(ctx context.Context, ueID string) *UeData {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:278
	_go_fuzz_dep_.CoverTab[194567]++
												if len(ueID) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:279
		_go_fuzz_dep_.CoverTab[194570]++
													panic("bad data")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:280
		// _ = "end of CoverTab[194570]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:281
		_go_fuzz_dep_.CoverTab[194571]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:281
		// _ = "end of CoverTab[194571]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:281
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:281
	// _ = "end of CoverTab[194567]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:281
	_go_fuzz_dep_.CoverTab[194568]++
												ueData := &UeData{
		UeID:		ueID,
		CGIString:	"",
		RrcState:	e2sm_mho.Rrcstatus_name[int32(e2sm_mho.Rrcstatus_RRCSTATUS_CONNECTED)],
		RsrpNeighbors:	make(map[string]int32),
	}
	_, err := c.ueStore.Put(ctx, ueID, *ueData)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:289
		_go_fuzz_dep_.CoverTab[194572]++
													log.Warn(err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:290
		// _ = "end of CoverTab[194572]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:291
		_go_fuzz_dep_.CoverTab[194573]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:291
		// _ = "end of CoverTab[194573]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:291
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:291
	// _ = "end of CoverTab[194568]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:291
	_go_fuzz_dep_.CoverTab[194569]++

												return ueData
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:293
	// _ = "end of CoverTab[194569]"
}

func (c *Ctrl) getUe(ctx context.Context, ueID string) *UeData {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:296
	_go_fuzz_dep_.CoverTab[194574]++
												var ueData *UeData
												u, err := c.ueStore.Get(ctx, ueID)
												if err != nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:299
		_go_fuzz_dep_.CoverTab[194577]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:299
		return u == nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:299
		// _ = "end of CoverTab[194577]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:299
	}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:299
		_go_fuzz_dep_.CoverTab[194578]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:300
		// _ = "end of CoverTab[194578]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:301
		_go_fuzz_dep_.CoverTab[194579]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:301
		// _ = "end of CoverTab[194579]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:301
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:301
	// _ = "end of CoverTab[194574]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:301
	_go_fuzz_dep_.CoverTab[194575]++
												t := u.Value.(UeData)
												ueData = &t
												if ueData.UeID != ueID {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:304
		_go_fuzz_dep_.CoverTab[194580]++
													panic("bad data")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:305
		// _ = "end of CoverTab[194580]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:306
		_go_fuzz_dep_.CoverTab[194581]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:306
		// _ = "end of CoverTab[194581]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:306
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:306
	// _ = "end of CoverTab[194575]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:306
	_go_fuzz_dep_.CoverTab[194576]++

												return ueData
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:308
	// _ = "end of CoverTab[194576]"
}

func (c *Ctrl) setUe(ctx context.Context, ueData *UeData) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:311
	_go_fuzz_dep_.CoverTab[194582]++
												_, err := c.ueStore.Put(ctx, ueData.UeID, *ueData)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:313
		_go_fuzz_dep_.CoverTab[194583]++
													panic("bad data")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:314
		// _ = "end of CoverTab[194583]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:315
		_go_fuzz_dep_.CoverTab[194584]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:315
		// _ = "end of CoverTab[194584]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:315
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:315
	// _ = "end of CoverTab[194582]"
}

func (c *Ctrl) attachUe(ctx context.Context, ueData *UeData, cgi string) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:318
	_go_fuzz_dep_.CoverTab[194585]++

												c.detachUe(ctx, ueData)

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:323
	ueData.CGIString = cgi
	c.setUe(ctx, ueData)
	cell := c.getCell(ctx, cgi)
	if cell == nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:326
		_go_fuzz_dep_.CoverTab[194587]++
													cell = c.createCell(ctx, cgi)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:327
		// _ = "end of CoverTab[194587]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:328
		_go_fuzz_dep_.CoverTab[194588]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:328
		// _ = "end of CoverTab[194588]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:328
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:328
	// _ = "end of CoverTab[194585]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:328
	_go_fuzz_dep_.CoverTab[194586]++
												cell.Ues[ueData.UeID] = ueData
												c.setCell(ctx, cell)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:330
	// _ = "end of CoverTab[194586]"
}

func (c *Ctrl) detachUe(ctx context.Context, ueData *UeData) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:333
	_go_fuzz_dep_.CoverTab[194589]++
												for _, cell := range c.cells {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:334
		_go_fuzz_dep_.CoverTab[194590]++
													delete(cell.Ues, ueData.UeID)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:335
		// _ = "end of CoverTab[194590]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:336
	// _ = "end of CoverTab[194589]"
}

func (c *Ctrl) setUeRrcState(ctx context.Context, ueData *UeData, newRrcState string, cgi string) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:339
	_go_fuzz_dep_.CoverTab[194591]++
												oldRrcState := ueData.RrcState

												if oldRrcState == e2sm_mho.Rrcstatus_name[int32(e2sm_mho.Rrcstatus_RRCSTATUS_CONNECTED)] && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:342
		_go_fuzz_dep_.CoverTab[194593]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:342
		return newRrcState == e2sm_mho.Rrcstatus_name[int32(e2sm_mho.Rrcstatus_RRCSTATUS_IDLE)]
													// _ = "end of CoverTab[194593]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:343
	}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:343
		_go_fuzz_dep_.CoverTab[194594]++
													c.detachUe(ctx, ueData)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:344
		// _ = "end of CoverTab[194594]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:345
		_go_fuzz_dep_.CoverTab[194595]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:345
		if oldRrcState == e2sm_mho.Rrcstatus_name[int32(e2sm_mho.Rrcstatus_RRCSTATUS_IDLE)] && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:345
			_go_fuzz_dep_.CoverTab[194596]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:345
			return newRrcState == e2sm_mho.Rrcstatus_name[int32(e2sm_mho.Rrcstatus_RRCSTATUS_CONNECTED)]
														// _ = "end of CoverTab[194596]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:346
		}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:346
			_go_fuzz_dep_.CoverTab[194597]++
														c.attachUe(ctx, ueData, cgi)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:347
			// _ = "end of CoverTab[194597]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:348
			_go_fuzz_dep_.CoverTab[194598]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:348
			// _ = "end of CoverTab[194598]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:348
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:348
		// _ = "end of CoverTab[194595]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:348
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:348
	// _ = "end of CoverTab[194591]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:348
	_go_fuzz_dep_.CoverTab[194592]++
												ueData.RrcState = newRrcState
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:349
	// _ = "end of CoverTab[194592]"
}

func (c *Ctrl) createCell(ctx context.Context, cgi string) *CellData {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:352
	_go_fuzz_dep_.CoverTab[194599]++
												if len(cgi) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:353
		_go_fuzz_dep_.CoverTab[194602]++
													panic("bad data")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:354
		// _ = "end of CoverTab[194602]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:355
		_go_fuzz_dep_.CoverTab[194603]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:355
		// _ = "end of CoverTab[194603]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:355
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:355
	// _ = "end of CoverTab[194599]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:355
	_go_fuzz_dep_.CoverTab[194600]++
												cellData := &CellData{
		CGIString:	cgi,
		Ues:		make(map[string]*UeData),
	}
	_, err := c.cellStore.Put(ctx, cgi, *cellData)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:361
		_go_fuzz_dep_.CoverTab[194604]++
													panic("bad data")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:362
		// _ = "end of CoverTab[194604]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:363
		_go_fuzz_dep_.CoverTab[194605]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:363
		// _ = "end of CoverTab[194605]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:363
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:363
	// _ = "end of CoverTab[194600]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:363
	_go_fuzz_dep_.CoverTab[194601]++
												c.cells[cellData.CGIString] = cellData
												return cellData
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:365
	// _ = "end of CoverTab[194601]"
}

func (c *Ctrl) getCell(ctx context.Context, cgi string) *CellData {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:368
	_go_fuzz_dep_.CoverTab[194606]++
												var cellData *CellData
												cell, err := c.cellStore.Get(ctx, cgi)
												if err != nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:371
		_go_fuzz_dep_.CoverTab[194609]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:371
		return cell == nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:371
		// _ = "end of CoverTab[194609]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:371
	}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:371
		_go_fuzz_dep_.CoverTab[194610]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:372
		// _ = "end of CoverTab[194610]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:373
		_go_fuzz_dep_.CoverTab[194611]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:373
		// _ = "end of CoverTab[194611]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:373
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:373
	// _ = "end of CoverTab[194606]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:373
	_go_fuzz_dep_.CoverTab[194607]++
												t := cell.Value.(CellData)
												if t.CGIString != cgi {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:375
		_go_fuzz_dep_.CoverTab[194612]++
													panic("bad data")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:376
		// _ = "end of CoverTab[194612]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:377
		_go_fuzz_dep_.CoverTab[194613]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:377
		// _ = "end of CoverTab[194613]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:377
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:377
	// _ = "end of CoverTab[194607]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:377
	_go_fuzz_dep_.CoverTab[194608]++
												cellData = &t
												return cellData
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:379
	// _ = "end of CoverTab[194608]"
}

func (c *Ctrl) setCell(ctx context.Context, cellData *CellData) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:382
	_go_fuzz_dep_.CoverTab[194614]++
												if len(cellData.CGIString) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:383
		_go_fuzz_dep_.CoverTab[194616]++
													panic("bad data")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:384
		// _ = "end of CoverTab[194616]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:385
		_go_fuzz_dep_.CoverTab[194617]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:385
		// _ = "end of CoverTab[194617]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:385
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:385
	// _ = "end of CoverTab[194614]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:385
	_go_fuzz_dep_.CoverTab[194615]++
												_, err := c.cellStore.Put(ctx, cellData.CGIString, *cellData)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:387
		_go_fuzz_dep_.CoverTab[194618]++
													panic("bad data")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:388
		// _ = "end of CoverTab[194618]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:389
		_go_fuzz_dep_.CoverTab[194619]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:389
		// _ = "end of CoverTab[194619]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:389
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:389
	// _ = "end of CoverTab[194615]"
}

func plmnIDBytesToInt(b []byte) uint64 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:392
	_go_fuzz_dep_.CoverTab[194620]++
												return uint64(b[2])<<16 | uint64(b[1])<<8 | uint64(b[0])
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:393
	// _ = "end of CoverTab[194620]"
}

func plmnIDNciToCGI(plmnID uint64, nci uint64) string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:396
	_go_fuzz_dep_.CoverTab[194621]++
												return strconv.FormatInt(int64(plmnID<<36|(nci&0xfffffffff)), 16)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:397
	// _ = "end of CoverTab[194621]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:405
func getNciFromCellGlobalID(cellGlobalID *e2sm_v2_ies.Cgi) uint64 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:405
	_go_fuzz_dep_.CoverTab[194622]++
												return BitStringToUint64(cellGlobalID.GetNRCgi().GetNRcellIdentity().GetValue().GetValue(), int(cellGlobalID.GetNRCgi().GetNRcellIdentity().GetValue().GetLen()))
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:406
	// _ = "end of CoverTab[194622]"
}

func getPlmnIDBytesFromCellGlobalID(cellGlobalID *e2sm_v2_ies.Cgi) []byte {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:409
	_go_fuzz_dep_.CoverTab[194623]++
												return cellGlobalID.GetNRCgi().GetPLmnidentity().GetValue()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:410
	// _ = "end of CoverTab[194623]"
}

func getCGIFromIndicationHeader(header *e2sm_mho.E2SmMhoIndicationHeaderFormat1) string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:413
	_go_fuzz_dep_.CoverTab[194624]++
												nci := getNciFromCellGlobalID(header.GetCgi())
												plmnIDBytes := getPlmnIDBytesFromCellGlobalID(header.GetCgi())
												plmnID := plmnIDBytesToInt(plmnIDBytes)
												return plmnIDNciToCGI(plmnID, nci)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:417
	// _ = "end of CoverTab[194624]"
}

func getCGIFromMeasReportItem(measReport *e2sm_mho.E2SmMhoMeasurementReportItem) string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:420
	_go_fuzz_dep_.CoverTab[194625]++
												nci := getNciFromCellGlobalID(measReport.GetCgi())
												plmnIDBytes := getPlmnIDBytesFromCellGlobalID(measReport.GetCgi())
												plmnID := plmnIDBytesToInt(plmnIDBytes)
												return plmnIDNciToCGI(plmnID, nci)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:424
	// _ = "end of CoverTab[194625]"
}

func getCgiFromHO(ueData *UeData, ho handover.A3HandoverDecision) string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:427
	_go_fuzz_dep_.CoverTab[194626]++
												servingCGI := ueData.CGI
												servingPlmnIDBytes := servingCGI.GetNRCgi().GetPLmnidentity().GetValue()
												servingPlmnID := plmnIDBytesToInt(servingPlmnIDBytes)
												targetPlmnID := servingPlmnID
												targetNCI, err := strconv.Atoi(ho.TargetCell.GetID().String())
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:433
		_go_fuzz_dep_.CoverTab[194628]++
													panic("bad data")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:434
		// _ = "end of CoverTab[194628]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:435
		_go_fuzz_dep_.CoverTab[194629]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:435
		// _ = "end of CoverTab[194629]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:435
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:435
	// _ = "end of CoverTab[194626]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:435
	_go_fuzz_dep_.CoverTab[194627]++
												return plmnIDNciToCGI(targetPlmnID, uint64(targetNCI))
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:436
	// _ = "end of CoverTab[194627]"
}

func getRsrpFromMeasReport(servingNci uint64, measReport []*e2sm_mho.E2SmMhoMeasurementReportItem) (int32, map[string]int32) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:439
	_go_fuzz_dep_.CoverTab[194630]++
												var rsrpServing int32
												rsrpNeighbors := make(map[string]int32)

												for _, measReportItem := range measReport {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:443
		_go_fuzz_dep_.CoverTab[194632]++
													if getNciFromCellGlobalID(measReportItem.GetCgi()) == servingNci {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:444
			_go_fuzz_dep_.CoverTab[194633]++
														rsrpServing = measReportItem.GetRsrp().GetValue()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:445
			// _ = "end of CoverTab[194633]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:446
			_go_fuzz_dep_.CoverTab[194634]++
														CGIString := getCGIFromMeasReportItem(measReportItem)
														rsrpNeighbors[CGIString] = measReportItem.GetRsrp().GetValue()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:448
			// _ = "end of CoverTab[194634]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:449
		// _ = "end of CoverTab[194632]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:450
	// _ = "end of CoverTab[194630]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:450
	_go_fuzz_dep_.CoverTab[194631]++

												return rsrpServing, rsrpNeighbors
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:452
	// _ = "end of CoverTab[194631]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:453
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/mho.go:453
var _ = _go_fuzz_dep_.CoverTab
