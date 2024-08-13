// SPDX-FileCopyrightText: 2019-present Open Networking Foundation <info@opennetworking.org>
// SPDX-FileCopyrightText: 2019-present Rimedo Labs
//
// SPDX-License-Identifier: Apache-2.0

//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:6
package mho

//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:6
import (
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:6
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:6
)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:6
import (
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:6
	_atomic_ "sync/atomic"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:6
)

import (
	"context"
	"reflect"
	"strconv"
	"sync"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"

	gofuzzdep "github.com/dvyukov/go-fuzz/go-fuzz-dep"
	policyAPI "github.com/onosproject/onos-a1-dm/go/policy_schemas/traffic_steering_preference/v2"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-mho-go"
	e2sm_v2_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-v2-ies"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-mho/pkg/store"
	"google.golang.org/protobuf/proto"
)

var log = logging.GetLogger("rimedo-ts", "mho")

type E2NodeIndication struct {
	NodeID		string
	TriggerType	e2sm_mho.MhoTriggerType
	IndMsg		e2api.Indication
}

func NewController(indChan chan *E2NodeIndication, ueStore store.Store, cellStore store.Store, onosPolicyStore store.Store, policies map[string]*PolicyData, flag bool) *Controller {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:33
	_go_fuzz_dep_.CoverTab[179586]++

									return &Controller{
		IndChan:		indChan,
		ueStore:		ueStore,
		cellStore:		cellStore,
		onosPolicyStore:	onosPolicyStore,
		mu:			sync.RWMutex{},
		cells:			make(map[string]*CellData),
		policies:		policies,
		topoIDsEnabled:		flag,
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:44
	// _ = "end of CoverTab[179586]"
}

type Controller struct {
	IndChan		chan *E2NodeIndication
	ueStore		store.Store
	cellStore	store.Store
	onosPolicyStore	store.Store
	mu		sync.RWMutex
	cells		map[string]*CellData
	policies	map[string]*PolicyData
	topoIDsEnabled	bool
}

func (c *Controller) Run(ctx context.Context, flag *bool) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:58
	_go_fuzz_dep_.CoverTab[179587]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:58
	_curRoutineNum159_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:58
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum159_)
									go c.listenIndChan(ctx, flag)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:59
	// _ = "end of CoverTab[179587]"
}

func (c *Controller) listenIndChan(ctx context.Context, flag *bool) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:62
	_go_fuzz_dep_.CoverTab[179588]++
									var err error
									for {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:64
		_go_fuzz_dep_.CoverTab[179589]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:64
		_go_fuzz_dep_.LoopFunc()
										gofuzzdep.LoopPos()
										indMsg := <-c.IndChan

										indHeaderByte := indMsg.IndMsg.Header
										indMessageByte := indMsg.IndMsg.Payload
										e2NodeID := indMsg.NodeID

										indHeader := e2sm_mho.E2SmMhoIndicationHeader{}
										if err = proto.Unmarshal(indHeaderByte, &indHeader); err == nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:73
			_go_fuzz_dep_.CoverTab[179591]++
											indMessage := e2sm_mho.E2SmMhoIndicationMessage{}
											if err = proto.Unmarshal(indMessageByte, &indMessage); err == nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:75
				_go_fuzz_dep_.CoverTab[179592]++
												log.Debugf("Received MHO indication message header: %+v, message: %+v, e2NodeID: %v", indHeader.E2SmMhoIndicationHeader, indMessage.E2SmMhoIndicationMessage, e2NodeID)
												switch x := indMessage.E2SmMhoIndicationMessage.(type) {
				case *e2sm_mho.E2SmMhoIndicationMessage_IndicationMessageFormat1:
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:78
					_go_fuzz_dep_.CoverTab[179593]++
													if indMsg.TriggerType == e2sm_mho.MhoTriggerType_MHO_TRIGGER_TYPE_UPON_RCV_MEAS_REPORT {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:79
						_go_fuzz_dep_.CoverTab[179596]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:79
						_curRoutineNum160_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:79
						_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum160_)
														go c.handleMeasReport(ctx, indHeader.GetIndicationHeaderFormat1(), indMessage.GetIndicationMessageFormat1(), e2NodeID, flag)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:80
						// _ = "end of CoverTab[179596]"
					} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:81
						_go_fuzz_dep_.CoverTab[179597]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:81
						if indMsg.TriggerType == e2sm_mho.MhoTriggerType_MHO_TRIGGER_TYPE_PERIODIC {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:81
							_go_fuzz_dep_.CoverTab[179598]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:81
							_curRoutineNum161_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:81
							_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum161_)
															go c.handlePeriodicReport(ctx, indHeader.GetIndicationHeaderFormat1(), indMessage.GetIndicationMessageFormat1(), e2NodeID, flag)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:82
							// _ = "end of CoverTab[179598]"
						} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:83
							_go_fuzz_dep_.CoverTab[179599]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:83
							// _ = "end of CoverTab[179599]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:83
						}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:83
						// _ = "end of CoverTab[179597]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:83
					}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:83
					// _ = "end of CoverTab[179593]"
				case *e2sm_mho.E2SmMhoIndicationMessage_IndicationMessageFormat2:
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:84
					_go_fuzz_dep_.CoverTab[179594]++
													go c.handleRrcState(ctx, indHeader.GetIndicationHeaderFormat1(), indMessage.GetIndicationMessageFormat2(), e2NodeID)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:85
					// _ = "end of CoverTab[179594]"
				default:
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:86
					_go_fuzz_dep_.CoverTab[179595]++
													log.Warnf("Unknown MHO indication message format, indication message: %v", x)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:87
					// _ = "end of CoverTab[179595]"
				}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:88
				// _ = "end of CoverTab[179592]"
			} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:89
				_go_fuzz_dep_.CoverTab[179600]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:89
				// _ = "end of CoverTab[179600]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:89
			}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:89
			// _ = "end of CoverTab[179591]"
		} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:90
			_go_fuzz_dep_.CoverTab[179601]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:90
			// _ = "end of CoverTab[179601]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:90
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:90
		// _ = "end of CoverTab[179589]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:90
		_go_fuzz_dep_.CoverTab[179590]++
										if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:91
			_go_fuzz_dep_.CoverTab[179602]++
											log.Error(err)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:92
			// _ = "end of CoverTab[179602]"
		} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:93
			_go_fuzz_dep_.CoverTab[179603]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:93
			// _ = "end of CoverTab[179603]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:93
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:93
		// _ = "end of CoverTab[179590]"
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:94
	// _ = "end of CoverTab[179588]"
}

func (c *Controller) handlePeriodicReport(ctx context.Context, header *e2sm_mho.E2SmMhoIndicationHeaderFormat1, message *e2sm_mho.E2SmMhoIndicationMessageFormat1, e2NodeID string, flag *bool) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:97
	_go_fuzz_dep_.CoverTab[179604]++
									c.mu.Lock()
									defer c.mu.Unlock()
									ueID, err := GetUeID(message.GetUeId())
									if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:101
		_go_fuzz_dep_.CoverTab[179610]++
										log.Errorf("handlePeriodicReport() couldn't extract UeID: %v", err)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:102
		// _ = "end of CoverTab[179610]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:103
		_go_fuzz_dep_.CoverTab[179611]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:103
		// _ = "end of CoverTab[179611]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:103
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:103
	// _ = "end of CoverTab[179604]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:103
	_go_fuzz_dep_.CoverTab[179605]++
									cgi := GetCGIFromIndicationHeader(header)
									cgi = c.ConvertCgiToTheRightForm(cgi)
									cgiObject := header.GetCgi()

									ueIdString := strconv.Itoa(int(ueID))
									n := (16 - len(ueIdString))
									for i := 0; i < n; i++ {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:110
		_go_fuzz_dep_.CoverTab[179612]++
										ueIdString = "0" + ueIdString
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:111
		// _ = "end of CoverTab[179612]"
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:112
	// _ = "end of CoverTab[179605]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:112
	_go_fuzz_dep_.CoverTab[179606]++
									var ueData *UeData
									newUe := false
									ueData = c.GetUe(ctx, ueIdString)
									if ueData == nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:116
		_go_fuzz_dep_.CoverTab[179613]++
										ueData = c.CreateUe(ctx, ueIdString)
										c.AttachUe(ctx, ueData, cgi, cgiObject)
										newUe = true
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:119
		// _ = "end of CoverTab[179613]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:120
		_go_fuzz_dep_.CoverTab[179614]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:120
		if ueData.CGIString != cgi {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:120
			_go_fuzz_dep_.CoverTab[179615]++
											return
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:121
			// _ = "end of CoverTab[179615]"
		} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:122
			_go_fuzz_dep_.CoverTab[179616]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:122
			// _ = "end of CoverTab[179616]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:122
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:122
		// _ = "end of CoverTab[179614]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:122
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:122
	// _ = "end of CoverTab[179606]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:122
	_go_fuzz_dep_.CoverTab[179607]++

									ueData.E2NodeID = e2NodeID

									rsrpServing, rsrpNeighbors, rsrpTable, cgiTable := c.GetRsrpFromMeasReport(ctx, GetNciFromCellGlobalID(header.GetCgi()), message.MeasReport)

									old5qi := ueData.FiveQi
									ueData.FiveQi = c.GetFiveQiFromMeasReport(ctx, GetNciFromCellGlobalID(header.GetCgi()), message.MeasReport)

									if *flag && func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:131
		_go_fuzz_dep_.CoverTab[179617]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:131
		return (old5qi != ueData.FiveQi)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:131
		// _ = "end of CoverTab[179617]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:131
	}() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:131
		_go_fuzz_dep_.CoverTab[179618]++
										log.Infof("\t\tQUALITY MESSAGE: 5QI for UE [ID:%v] changed [5QI:%v]\n", ueData.UeID, ueData.FiveQi)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:132
		// _ = "end of CoverTab[179618]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:133
		_go_fuzz_dep_.CoverTab[179619]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:133
		// _ = "end of CoverTab[179619]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:133
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:133
	// _ = "end of CoverTab[179607]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:133
	_go_fuzz_dep_.CoverTab[179608]++

									if !newUe && func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:135
		_go_fuzz_dep_.CoverTab[179620]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:135
		return rsrpServing == ueData.RsrpServing
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:135
		// _ = "end of CoverTab[179620]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:135
	}() && func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:135
		_go_fuzz_dep_.CoverTab[179621]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:135
		return reflect.DeepEqual(rsrpNeighbors, ueData.RsrpNeighbors)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:135
		// _ = "end of CoverTab[179621]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:135
	}() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:135
		_go_fuzz_dep_.CoverTab[179622]++
										return
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:136
		// _ = "end of CoverTab[179622]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:137
		_go_fuzz_dep_.CoverTab[179623]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:137
		// _ = "end of CoverTab[179623]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:137
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:137
	// _ = "end of CoverTab[179608]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:137
	_go_fuzz_dep_.CoverTab[179609]++

									ueData.RsrpServing, ueData.RsrpNeighbors, ueData.RsrpTable, ueData.CgiTable = rsrpServing, rsrpNeighbors, rsrpTable, cgiTable
									c.SetUe(ctx, ueData)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:140
	// _ = "end of CoverTab[179609]"

}

func (c *Controller) handleMeasReport(ctx context.Context, header *e2sm_mho.E2SmMhoIndicationHeaderFormat1, message *e2sm_mho.E2SmMhoIndicationMessageFormat1, e2NodeID string, flag *bool) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:144
	_go_fuzz_dep_.CoverTab[179624]++
									c.mu.Lock()
									defer c.mu.Unlock()
									ueID, err := GetUeID(message.GetUeId())
									if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:148
		_go_fuzz_dep_.CoverTab[179629]++
										log.Errorf("handleMeasReport() couldn't extract UeID: %v", err)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:149
		// _ = "end of CoverTab[179629]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:150
		_go_fuzz_dep_.CoverTab[179630]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:150
		// _ = "end of CoverTab[179630]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:150
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:150
	// _ = "end of CoverTab[179624]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:150
	_go_fuzz_dep_.CoverTab[179625]++
									cgi := GetCGIFromIndicationHeader(header)
									cgi = c.ConvertCgiToTheRightForm(cgi)
									cgiObject := header.GetCgi()

									ueIdString := strconv.Itoa(int(ueID))
									n := (16 - len(ueIdString))
									for i := 0; i < n; i++ {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:157
		_go_fuzz_dep_.CoverTab[179631]++
										ueIdString = "0" + ueIdString
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:158
		// _ = "end of CoverTab[179631]"
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:159
	// _ = "end of CoverTab[179625]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:159
	_go_fuzz_dep_.CoverTab[179626]++
									var ueData *UeData
									ueData = c.GetUe(ctx, ueIdString)
									if ueData == nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:162
		_go_fuzz_dep_.CoverTab[179632]++
										ueData = c.CreateUe(ctx, ueIdString)
										c.AttachUe(ctx, ueData, cgi, cgiObject)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:164
		// _ = "end of CoverTab[179632]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:165
		_go_fuzz_dep_.CoverTab[179633]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:165
		if ueData.CGIString != cgi {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:165
			_go_fuzz_dep_.CoverTab[179634]++
											return
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:166
			// _ = "end of CoverTab[179634]"
		} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:167
			_go_fuzz_dep_.CoverTab[179635]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:167
			// _ = "end of CoverTab[179635]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:167
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:167
		// _ = "end of CoverTab[179633]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:167
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:167
	// _ = "end of CoverTab[179626]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:167
	_go_fuzz_dep_.CoverTab[179627]++

									ueData.E2NodeID = e2NodeID

									ueData.RsrpServing, ueData.RsrpNeighbors, ueData.RsrpTable, ueData.CgiTable = c.GetRsrpFromMeasReport(ctx, GetNciFromCellGlobalID(header.GetCgi()), message.MeasReport)

									old5qi := ueData.FiveQi
									ueData.FiveQi = c.GetFiveQiFromMeasReport(ctx, GetNciFromCellGlobalID(header.GetCgi()), message.MeasReport)
									if *flag && func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:175
		_go_fuzz_dep_.CoverTab[179636]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:175
		return (old5qi != ueData.FiveQi)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:175
		// _ = "end of CoverTab[179636]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:175
	}() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:175
		_go_fuzz_dep_.CoverTab[179637]++
										log.Infof("\t\tQUALITY MESSAGE: 5QI for UE [ID:%v] changed [5QI:%v]\n", ueData.UeID, ueData.FiveQi)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:176
		// _ = "end of CoverTab[179637]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:177
		_go_fuzz_dep_.CoverTab[179638]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:177
		// _ = "end of CoverTab[179638]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:177
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:177
	// _ = "end of CoverTab[179627]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:177
	_go_fuzz_dep_.CoverTab[179628]++

									c.SetUe(ctx, ueData)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:179
	// _ = "end of CoverTab[179628]"

}

func (c *Controller) handleRrcState(ctx context.Context, header *e2sm_mho.E2SmMhoIndicationHeaderFormat1, message *e2sm_mho.E2SmMhoIndicationMessageFormat2, e2NodeID string) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:183
	_go_fuzz_dep_.CoverTab[179639]++
									c.mu.Lock()
									defer c.mu.Unlock()
									ueID, err := GetUeID(message.GetUeId())
									if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:187
		_go_fuzz_dep_.CoverTab[179643]++
										log.Errorf("handleRrcState() couldn't extract UeID: %v", err)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:188
		// _ = "end of CoverTab[179643]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:189
		_go_fuzz_dep_.CoverTab[179644]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:189
		// _ = "end of CoverTab[179644]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:189
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:189
	// _ = "end of CoverTab[179639]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:189
	_go_fuzz_dep_.CoverTab[179640]++
									cgi := GetCGIFromIndicationHeader(header)
									cgi = c.ConvertCgiToTheRightForm(cgi)
									cgiObject := header.GetCgi()

									ueIdString := strconv.Itoa(int(ueID))
									n := (16 - len(ueIdString))
									for i := 0; i < n; i++ {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:196
		_go_fuzz_dep_.CoverTab[179645]++
										ueIdString = "0" + ueIdString
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:197
		// _ = "end of CoverTab[179645]"
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:198
	// _ = "end of CoverTab[179640]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:198
	_go_fuzz_dep_.CoverTab[179641]++
									var ueData *UeData
									ueData = c.GetUe(ctx, ueIdString)
									if ueData == nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:201
		_go_fuzz_dep_.CoverTab[179646]++
										ueData = c.CreateUe(ctx, ueIdString)
										c.AttachUe(ctx, ueData, cgi, cgiObject)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:203
		// _ = "end of CoverTab[179646]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:204
		_go_fuzz_dep_.CoverTab[179647]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:204
		if ueData.CGIString != cgi {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:204
			_go_fuzz_dep_.CoverTab[179648]++
											return
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:205
			// _ = "end of CoverTab[179648]"
		} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:206
			_go_fuzz_dep_.CoverTab[179649]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:206
			// _ = "end of CoverTab[179649]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:206
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:206
		// _ = "end of CoverTab[179647]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:206
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:206
	// _ = "end of CoverTab[179641]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:206
	_go_fuzz_dep_.CoverTab[179642]++

									ueData.E2NodeID = e2NodeID

									newRrcState := message.GetRrcStatus().String()
									c.SetUeRrcState(ctx, ueData, newRrcState, cgi, cgiObject)

									c.SetUe(ctx, ueData)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:213
	// _ = "end of CoverTab[179642]"

}

func (c *Controller) CreateUe(ctx context.Context, ueID string) *UeData {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:217
	_go_fuzz_dep_.CoverTab[179650]++
									if len(ueID) == 0 {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:218
		_go_fuzz_dep_.CoverTab[179653]++
										panic("bad data")
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:219
		// _ = "end of CoverTab[179653]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:220
		_go_fuzz_dep_.CoverTab[179654]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:220
		// _ = "end of CoverTab[179654]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:220
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:220
	// _ = "end of CoverTab[179650]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:220
	_go_fuzz_dep_.CoverTab[179651]++
									ueData := &UeData{
		UeID:		ueID,
		CGIString:	"",
		RrcState:	e2sm_mho.Rrcstatus_name[int32(e2sm_mho.Rrcstatus_RRCSTATUS_CONNECTED)],
		RsrpNeighbors:	make(map[string]int32),
		Idle:		false,
	}
	_, err := c.ueStore.Put(ctx, ueID, *ueData)
	if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:229
		_go_fuzz_dep_.CoverTab[179655]++
										log.Warn(err)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:230
		// _ = "end of CoverTab[179655]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:231
		_go_fuzz_dep_.CoverTab[179656]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:231
		// _ = "end of CoverTab[179656]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:231
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:231
	// _ = "end of CoverTab[179651]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:231
	_go_fuzz_dep_.CoverTab[179652]++

									return ueData
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:233
	// _ = "end of CoverTab[179652]"
}

func (c *Controller) GetUe(ctx context.Context, ueID string) *UeData {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:236
	_go_fuzz_dep_.CoverTab[179657]++
									var ueData *UeData
									u, err := c.ueStore.Get(ctx, ueID)
									if err != nil || func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:239
		_go_fuzz_dep_.CoverTab[179660]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:239
		return u == nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:239
		// _ = "end of CoverTab[179660]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:239
	}() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:239
		_go_fuzz_dep_.CoverTab[179661]++
										return nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:240
		// _ = "end of CoverTab[179661]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:241
		_go_fuzz_dep_.CoverTab[179662]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:241
		// _ = "end of CoverTab[179662]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:241
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:241
	// _ = "end of CoverTab[179657]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:241
	_go_fuzz_dep_.CoverTab[179658]++
									t := u.Value.(UeData)
									ueData = &t
									if ueData.UeID != ueID {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:244
		_go_fuzz_dep_.CoverTab[179663]++
										panic("bad data")
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:245
		// _ = "end of CoverTab[179663]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:246
		_go_fuzz_dep_.CoverTab[179664]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:246
		// _ = "end of CoverTab[179664]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:246
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:246
	// _ = "end of CoverTab[179658]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:246
	_go_fuzz_dep_.CoverTab[179659]++

									return ueData
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:248
	// _ = "end of CoverTab[179659]"
}

func (c *Controller) SetUe(ctx context.Context, ueData *UeData) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:251
	_go_fuzz_dep_.CoverTab[179665]++
									_, err := c.ueStore.Put(ctx, ueData.UeID, *ueData)
									if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:253
		_go_fuzz_dep_.CoverTab[179666]++
										panic("bad data")
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:254
		// _ = "end of CoverTab[179666]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:255
		_go_fuzz_dep_.CoverTab[179667]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:255
		// _ = "end of CoverTab[179667]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:255
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:255
	// _ = "end of CoverTab[179665]"
}

func (c *Controller) AttachUe(ctx context.Context, ueData *UeData, cgi string, cgiObject *e2sm_v2_ies.Cgi) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:258
	_go_fuzz_dep_.CoverTab[179668]++

									c.DetachUe(ctx, ueData)

									ueData.CGIString = cgi
									ueData.CGI = cgiObject
									c.SetUe(ctx, ueData)
									cell := c.GetCell(ctx, cgi)
									if cell == nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:266
		_go_fuzz_dep_.CoverTab[179670]++
										cell = c.CreateCell(ctx, cgi, cgiObject)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:267
		// _ = "end of CoverTab[179670]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:268
		_go_fuzz_dep_.CoverTab[179671]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:268
		// _ = "end of CoverTab[179671]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:268
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:268
	// _ = "end of CoverTab[179668]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:268
	_go_fuzz_dep_.CoverTab[179669]++
									cell.Ues[ueData.UeID] = ueData
									c.SetCell(ctx, cell)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:270
	// _ = "end of CoverTab[179669]"
}

func (c *Controller) DetachUe(ctx context.Context, ueData *UeData) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:273
	_go_fuzz_dep_.CoverTab[179672]++
									for _, cell := range c.cells {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:274
		_go_fuzz_dep_.CoverTab[179673]++
										delete(cell.Ues, ueData.UeID)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:275
		// _ = "end of CoverTab[179673]"
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:276
	// _ = "end of CoverTab[179672]"
}

func (c *Controller) SetUeRrcState(ctx context.Context, ueData *UeData, newRrcState string, cgi string, cgiObject *e2sm_v2_ies.Cgi) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:279
	_go_fuzz_dep_.CoverTab[179674]++
									oldRrcState := ueData.RrcState

									if oldRrcState == e2sm_mho.Rrcstatus_name[int32(e2sm_mho.Rrcstatus_RRCSTATUS_CONNECTED)] && func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:282
		_go_fuzz_dep_.CoverTab[179676]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:282
		return newRrcState == e2sm_mho.Rrcstatus_name[int32(e2sm_mho.Rrcstatus_RRCSTATUS_IDLE)]
										// _ = "end of CoverTab[179676]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:283
	}() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:283
		_go_fuzz_dep_.CoverTab[179677]++
										ueData.Idle = true
										c.DetachUe(ctx, ueData)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:285
		// _ = "end of CoverTab[179677]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:286
		_go_fuzz_dep_.CoverTab[179678]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:286
		if oldRrcState == e2sm_mho.Rrcstatus_name[int32(e2sm_mho.Rrcstatus_RRCSTATUS_IDLE)] && func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:286
			_go_fuzz_dep_.CoverTab[179679]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:286
			return newRrcState == e2sm_mho.Rrcstatus_name[int32(e2sm_mho.Rrcstatus_RRCSTATUS_CONNECTED)]
											// _ = "end of CoverTab[179679]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:287
		}() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:287
			_go_fuzz_dep_.CoverTab[179680]++
											ueData.Idle = false
											c.AttachUe(ctx, ueData, cgi, cgiObject)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:289
			// _ = "end of CoverTab[179680]"
		} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:290
			_go_fuzz_dep_.CoverTab[179681]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:290
			// _ = "end of CoverTab[179681]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:290
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:290
		// _ = "end of CoverTab[179678]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:290
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:290
	// _ = "end of CoverTab[179674]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:290
	_go_fuzz_dep_.CoverTab[179675]++
									ueData.RrcState = newRrcState
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:291
	// _ = "end of CoverTab[179675]"
}

func (c *Controller) CreateCell(ctx context.Context, cgi string, cgiObject *e2sm_v2_ies.Cgi) *CellData {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:294
	_go_fuzz_dep_.CoverTab[179682]++
									if len(cgi) == 0 {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:295
		_go_fuzz_dep_.CoverTab[179685]++
										panic("bad data")
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:296
		// _ = "end of CoverTab[179685]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:297
		_go_fuzz_dep_.CoverTab[179686]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:297
		// _ = "end of CoverTab[179686]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:297
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:297
	// _ = "end of CoverTab[179682]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:297
	_go_fuzz_dep_.CoverTab[179683]++
									cellData := &CellData{
		CGI:		cgiObject,
		CGIString:	cgi,
		Ues:		make(map[string]*UeData),
	}
	_, err := c.cellStore.Put(ctx, cgi, *cellData)
	if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:304
		_go_fuzz_dep_.CoverTab[179687]++
										panic("bad data")
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:305
		// _ = "end of CoverTab[179687]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:306
		_go_fuzz_dep_.CoverTab[179688]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:306
		// _ = "end of CoverTab[179688]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:306
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:306
	// _ = "end of CoverTab[179683]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:306
	_go_fuzz_dep_.CoverTab[179684]++
									c.cells[cellData.CGIString] = cellData
									return cellData
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:308
	// _ = "end of CoverTab[179684]"
}

func (c *Controller) GetCell(ctx context.Context, cgi string) *CellData {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:311
	_go_fuzz_dep_.CoverTab[179689]++
									var cellData *CellData
									cell, err := c.cellStore.Get(ctx, cgi)
									if err != nil || func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:314
		_go_fuzz_dep_.CoverTab[179692]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:314
		return cell == nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:314
		// _ = "end of CoverTab[179692]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:314
	}() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:314
		_go_fuzz_dep_.CoverTab[179693]++
										return nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:315
		// _ = "end of CoverTab[179693]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:316
		_go_fuzz_dep_.CoverTab[179694]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:316
		// _ = "end of CoverTab[179694]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:316
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:316
	// _ = "end of CoverTab[179689]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:316
	_go_fuzz_dep_.CoverTab[179690]++
									t := cell.Value.(CellData)
									if t.CGIString != cgi {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:318
		_go_fuzz_dep_.CoverTab[179695]++
										panic("bad data")
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:319
		// _ = "end of CoverTab[179695]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:320
		_go_fuzz_dep_.CoverTab[179696]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:320
		// _ = "end of CoverTab[179696]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:320
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:320
	// _ = "end of CoverTab[179690]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:320
	_go_fuzz_dep_.CoverTab[179691]++
									cellData = &t
									return cellData
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:322
	// _ = "end of CoverTab[179691]"
}

func (c *Controller) SetCell(ctx context.Context, cellData *CellData) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:325
	_go_fuzz_dep_.CoverTab[179697]++
									if len(cellData.CGIString) == 0 {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:326
		_go_fuzz_dep_.CoverTab[179700]++
										panic("bad data")
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:327
		// _ = "end of CoverTab[179700]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:328
		_go_fuzz_dep_.CoverTab[179701]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:328
		// _ = "end of CoverTab[179701]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:328
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:328
	// _ = "end of CoverTab[179697]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:328
	_go_fuzz_dep_.CoverTab[179698]++
									_, err := c.cellStore.Put(ctx, cellData.CGIString, *cellData)
									if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:330
		_go_fuzz_dep_.CoverTab[179702]++
										panic("bad data")
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:331
		// _ = "end of CoverTab[179702]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:332
		_go_fuzz_dep_.CoverTab[179703]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:332
		// _ = "end of CoverTab[179703]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:332
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:332
	// _ = "end of CoverTab[179698]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:332
	_go_fuzz_dep_.CoverTab[179699]++
									c.cells[cellData.CGIString] = cellData
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:333
	// _ = "end of CoverTab[179699]"
}

func (c *Controller) GetFiveQiFromMeasReport(ctx context.Context, servingNci uint64, measReport []*e2sm_mho.E2SmMhoMeasurementReportItem) int64 {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:336
	_go_fuzz_dep_.CoverTab[179704]++
									var fiveQiServing int64

									for _, measReportItem := range measReport {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:339
		_go_fuzz_dep_.CoverTab[179706]++

										if GetNciFromCellGlobalID(measReportItem.GetCgi()) == servingNci {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:341
			_go_fuzz_dep_.CoverTab[179707]++
											fiveQi := measReportItem.GetFiveQi()
											if fiveQi != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:343
				_go_fuzz_dep_.CoverTab[179708]++
												fiveQiServing = int64(fiveQi.GetValue())
												if fiveQiServing > 127 {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:345
					_go_fuzz_dep_.CoverTab[179709]++
													fiveQiServing = 2
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:346
					// _ = "end of CoverTab[179709]"
				} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:347
					_go_fuzz_dep_.CoverTab[179710]++
													fiveQiServing = 1
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:348
					// _ = "end of CoverTab[179710]"
				}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:349
				// _ = "end of CoverTab[179708]"
			} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:350
				_go_fuzz_dep_.CoverTab[179711]++
												fiveQiServing = -1
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:351
				// _ = "end of CoverTab[179711]"
			}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:352
			// _ = "end of CoverTab[179707]"
		} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:353
			_go_fuzz_dep_.CoverTab[179712]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:353
			// _ = "end of CoverTab[179712]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:353
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:353
		// _ = "end of CoverTab[179706]"
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:354
	// _ = "end of CoverTab[179704]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:354
	_go_fuzz_dep_.CoverTab[179705]++

									return fiveQiServing
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:356
	// _ = "end of CoverTab[179705]"
}

func (c *Controller) GetRsrpFromMeasReport(ctx context.Context, servingNci uint64, measReport []*e2sm_mho.E2SmMhoMeasurementReportItem) (int32, map[string]int32, map[string]int32, map[string]*e2sm_v2_ies.Cgi) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:359
	_go_fuzz_dep_.CoverTab[179713]++
									var rsrpServing int32
									rsrpNeighbors := make(map[string]int32)
									rsrpTable := make(map[string]int32)
									cgiTable := make(map[string]*e2sm_v2_ies.Cgi)

									for _, measReportItem := range measReport {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:365
		_go_fuzz_dep_.CoverTab[179715]++

										if GetNciFromCellGlobalID(measReportItem.GetCgi()) == servingNci {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:367
			_go_fuzz_dep_.CoverTab[179716]++
											CGIString := GetCGIFromMeasReportItem(measReportItem)
											CGIString = c.ConvertCgiToTheRightForm(CGIString)
											rsrpServing = measReportItem.GetRsrp().GetValue()
											rsrpTable[CGIString] = measReportItem.GetRsrp().GetValue()
											cgiTable[CGIString] = measReportItem.GetCgi()
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:372
			// _ = "end of CoverTab[179716]"
		} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:373
			_go_fuzz_dep_.CoverTab[179717]++
											CGIString := GetCGIFromMeasReportItem(measReportItem)
											CGIString = c.ConvertCgiToTheRightForm(CGIString)
											rsrpNeighbors[CGIString] = measReportItem.GetRsrp().GetValue()
											rsrpTable[CGIString] = measReportItem.GetRsrp().GetValue()
											cgiTable[CGIString] = measReportItem.GetCgi()
											cell := c.GetCell(ctx, CGIString)
											if cell == nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:380
				_go_fuzz_dep_.CoverTab[179718]++
												_ = c.CreateCell(ctx, CGIString, measReportItem.GetCgi())
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:381
				// _ = "end of CoverTab[179718]"
			} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:382
				_go_fuzz_dep_.CoverTab[179719]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:382
				// _ = "end of CoverTab[179719]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:382
			}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:382
			// _ = "end of CoverTab[179717]"
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:383
		// _ = "end of CoverTab[179715]"
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:384
	// _ = "end of CoverTab[179713]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:384
	_go_fuzz_dep_.CoverTab[179714]++

									return rsrpServing, rsrpNeighbors, rsrpTable, cgiTable
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:386
	// _ = "end of CoverTab[179714]"
}

func (c *Controller) CreatePolicy(ctx context.Context, key string, policy *policyAPI.API) *PolicyData {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:389
	_go_fuzz_dep_.CoverTab[179720]++
									if len(key) == 0 {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:390
		_go_fuzz_dep_.CoverTab[179723]++
										panic("bad data")
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:391
		// _ = "end of CoverTab[179723]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:392
		_go_fuzz_dep_.CoverTab[179724]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:392
		// _ = "end of CoverTab[179724]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:392
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:392
	// _ = "end of CoverTab[179720]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:392
	_go_fuzz_dep_.CoverTab[179721]++
									policyData := &PolicyData{
		Key:		key,
		API:		policy,
		IsEnforced:	true,
	}
	_, err := c.onosPolicyStore.Put(ctx, key, *policyData)
	if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:399
		_go_fuzz_dep_.CoverTab[179725]++
										log.Panic("bad data")
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:400
		// _ = "end of CoverTab[179725]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:401
		_go_fuzz_dep_.CoverTab[179726]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:401
		// _ = "end of CoverTab[179726]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:401
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:401
	// _ = "end of CoverTab[179721]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:401
	_go_fuzz_dep_.CoverTab[179722]++
									c.policies[policyData.Key] = policyData
									return policyData
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:403
	// _ = "end of CoverTab[179722]"
}

func (c *Controller) GetPolicy(ctx context.Context, key string) *PolicyData {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:406
	_go_fuzz_dep_.CoverTab[179727]++
									var policy *PolicyData
									p, err := c.onosPolicyStore.Get(ctx, key)
									if err != nil || func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:409
		_go_fuzz_dep_.CoverTab[179730]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:409
		return p == nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:409
		// _ = "end of CoverTab[179730]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:409
	}() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:409
		_go_fuzz_dep_.CoverTab[179731]++
										return nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:410
		// _ = "end of CoverTab[179731]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:411
		_go_fuzz_dep_.CoverTab[179732]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:411
		// _ = "end of CoverTab[179732]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:411
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:411
	// _ = "end of CoverTab[179727]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:411
	_go_fuzz_dep_.CoverTab[179728]++
									t := p.Value.(PolicyData)
									if t.Key != key {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:413
		_go_fuzz_dep_.CoverTab[179733]++
										panic("bad data")
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:414
		// _ = "end of CoverTab[179733]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:415
		_go_fuzz_dep_.CoverTab[179734]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:415
		// _ = "end of CoverTab[179734]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:415
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:415
	// _ = "end of CoverTab[179728]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:415
	_go_fuzz_dep_.CoverTab[179729]++
									policy = &t

									return policy
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:418
	// _ = "end of CoverTab[179729]"
}

func (c *Controller) SetPolicy(ctx context.Context, key string, policy *PolicyData) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:421
	_go_fuzz_dep_.CoverTab[179735]++
									_, err := c.onosPolicyStore.Put(ctx, key, *policy)
									if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:423
		_go_fuzz_dep_.CoverTab[179737]++
										panic("bad data")
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:424
		// _ = "end of CoverTab[179737]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:425
		_go_fuzz_dep_.CoverTab[179738]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:425
		// _ = "end of CoverTab[179738]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:425
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:425
	// _ = "end of CoverTab[179735]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:425
	_go_fuzz_dep_.CoverTab[179736]++
									c.policies[policy.Key] = policy
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:426
	// _ = "end of CoverTab[179736]"
}

func (c *Controller) DeletePolicy(ctx context.Context, key string) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:429
	_go_fuzz_dep_.CoverTab[179739]++
									if err := c.onosPolicyStore.Delete(ctx, key); err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:430
		_go_fuzz_dep_.CoverTab[179740]++
										panic("bad data")
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:431
		// _ = "end of CoverTab[179740]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:432
		_go_fuzz_dep_.CoverTab[179741]++
										delete(c.policies, key)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:433
		// _ = "end of CoverTab[179741]"
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:434
	// _ = "end of CoverTab[179739]"
}

func (c *Controller) GetPolicyStore() *store.Store {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:437
	_go_fuzz_dep_.CoverTab[179742]++
									return &c.onosPolicyStore
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:438
	// _ = "end of CoverTab[179742]"
}

func (c *Controller) ConvertCgiToTheRightForm(cgi string) string {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:441
	_go_fuzz_dep_.CoverTab[179743]++
									if c.topoIDsEnabled {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:442
		_go_fuzz_dep_.CoverTab[179745]++
										return cgi[0:6] + cgi[14:15] + cgi[12:14] + cgi[10:12] + cgi[8:10] + cgi[6:8]
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:443
		// _ = "end of CoverTab[179745]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:444
		_go_fuzz_dep_.CoverTab[179746]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:444
		// _ = "end of CoverTab[179746]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:444
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:444
	// _ = "end of CoverTab[179743]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:444
	_go_fuzz_dep_.CoverTab[179744]++
									return cgi
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:445
	// _ = "end of CoverTab[179744]"
}

//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:446
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/mho.go:446
var _ = _go_fuzz_dep_.CoverTab
