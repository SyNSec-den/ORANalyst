// SPDX-FileCopyrightText: 2019-present Open Networking Foundation <info@opennetworking.org>
// SPDX-FileCopyrightText: 2019-present Rimedo Labs
//
// SPDX-License-Identifier: Apache-2.0
// Copy from https://github.com/woojoong88/onos-kpimon/tree/sample-a1t-xapp/pkg/northbound/a1
// modified by RIMEDO Labs team

//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:8
package a1

//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:8
import (
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:8
)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:8
import (
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:8
	_atomic_ "sync/atomic"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:8
)

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	a1tapi "github.com/onosproject/onos-api/go/onos/a1t/a1"
	"github.com/onosproject/onos-lib-go/pkg/logging/service"
	"google.golang.org/grpc"
)

var SampleJSON1 = `
{
  "scope": {
    "ueId": "0000000000000001"
  },
  "tspResources": [
    {
      "cellIdList": [
        {"plmnId": {"mcc": "248","mnc": "35"},
          "cId": {"ncI": 39}},
        {"plmnId": {"mcc": "248","mnc": "35"},
         "cId": {"ncI": 40}}
      ], 
      "preference": "PREFER"
    },
    {
      "cellIdList": [
        {"plmnId": {"mcc": "248","mnc": "35"},
          "cId": {"ncI": 81}},
        {"plmnId": {"mcc": "248","mnc": "35"},
          "cId": {"ncI": 82}},
        {"plmnId": {"mcc": "248","mnc": "35"},
         "cId": {"ncI": 83}}
      ],
      "preference": "FORBID"
    }
  ]
}
`

var SampleJSON2 = `
{
  "scope": {
    "ueId": "0000000000000002"
  },
  "tspResources": [
    {
      "cellIdList": [
        {"plmnId": {"mcc": "248","mnc": "35"},
          "cId": {"ncI": 39}},
        {"plmnId": {"mcc": "248","mnc": "35"},
         "cId": {"ncI": 40}}
      ], 
      "preference": "PREFER"
    },
    {
      "cellIdList": [
        {"plmnId": {"mcc": "248","mnc": "35"},
          "cId": {"ncI": 81}},
        {"plmnId": {"mcc": "248","mnc": "35"},
          "cId": {"ncI": 82}},
        {"plmnId": {"mcc": "248","mnc": "35"},
         "cId": {"ncI": 83}}
      ],
      "preference": "FORBID"
    }
  ]
}
`

var SampleEnforcedStatus = `
{
  "enforceStatus": "ENFORCED"
}
`

var SampleNotEnforcedStatus = `
{
  "enforceStatus": "NOT_ENFORCED",
  "enforceReason": "SCOPE_NOT_APPLICABLE"
}
`

var SampleNotEnforcedPolicyID = "2"

func NewA1PService(policyMap *map[string][]byte, notifier chan bool) service.Service {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:98
	_go_fuzz_dep_.CoverTab[190661]++
											return &A1PService{
		TsPolicyTypeMap:	policyMap,
		notifier:		notifier,
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:102
	// _ = "end of CoverTab[190661]"
}

type A1PService struct {
	TsPolicyTypeMap	*map[string][]byte
	notifier	chan bool
}

func (a *A1PService) Register(s *grpc.Server) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:110
	_go_fuzz_dep_.CoverTab[190662]++
											server := &A1PServer{
		TsPolicyTypeMap:	*a.TsPolicyTypeMap,
		StatusUpdateCh:		make(chan *a1tapi.PolicyStatusMessage),
		notifier:		a.notifier,
	}
											a1tapi.RegisterPolicyServiceServer(s, server)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:116
	// _ = "end of CoverTab[190662]"
}

type A1PServer struct {
	TsPolicyTypeMap	map[string][]byte
	StatusUpdateCh	chan *a1tapi.PolicyStatusMessage
	notifier	chan bool
	mu		sync.RWMutex
}

func (a *A1PServer) PolicySetup(ctx context.Context, message *a1tapi.PolicyRequestMessage) (*a1tapi.PolicyResultMessage, error) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:126
	_go_fuzz_dep_.CoverTab[190663]++
											a.mu.Lock()
											defer a.mu.Unlock()
											var result map[string]interface{}
											json.Unmarshal(message.Message.Payload, &result)

											if message.PolicyType.Id != "ORAN_TrafficSteeringPreference_2.0.0" {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:132
		_go_fuzz_dep_.CoverTab[190667]++
												res := &a1tapi.PolicyResultMessage{
			PolicyId:	message.PolicyId,
			PolicyType:	message.PolicyType,
			Message: &a1tapi.ResultMessage{
				Header: &a1tapi.Header{
					PayloadType:	message.Message.Header.PayloadType,
					RequestId:	message.Message.Header.RequestId,
					Encoding:	message.Message.Header.Encoding,
					AppId:		message.Message.Header.AppId,
				}, Payload:	message.Message.Payload,
				Result: &a1tapi.Result{
					Success:	false,
					Reason:		"Policy type does not support",
				},
			},
		}
												return res, nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:149
		// _ = "end of CoverTab[190667]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:150
		_go_fuzz_dep_.CoverTab[190668]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:150
		// _ = "end of CoverTab[190668]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:150
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:150
	// _ = "end of CoverTab[190663]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:150
	_go_fuzz_dep_.CoverTab[190664]++

											if _, ok := a.TsPolicyTypeMap[message.PolicyId]; ok {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:152
		_go_fuzz_dep_.CoverTab[190669]++
												res := &a1tapi.PolicyResultMessage{
			PolicyId:	message.PolicyId,
			PolicyType:	message.PolicyType,
			Message: &a1tapi.ResultMessage{
				Header: &a1tapi.Header{
					PayloadType:	message.Message.Header.PayloadType,
					RequestId:	message.Message.Header.RequestId,
					Encoding:	message.Message.Header.Encoding,
					AppId:		message.Message.Header.AppId,
				}, Payload:	message.Message.Payload,
				Result: &a1tapi.Result{
					Success:	false,
					Reason:		"Policy ID already exists",
				},
			},
		}
												return res, nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:169
		// _ = "end of CoverTab[190669]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:170
		_go_fuzz_dep_.CoverTab[190670]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:170
		// _ = "end of CoverTab[190670]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:170
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:170
	// _ = "end of CoverTab[190664]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:170
	_go_fuzz_dep_.CoverTab[190665]++

											a.TsPolicyTypeMap[message.PolicyId] = message.Message.Payload
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:172
	_curRoutineNum168_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:172
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum168_)

											go func() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:174
		_go_fuzz_dep_.CoverTab[190671]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:174
		defer func() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:174
			_go_fuzz_dep_.CoverTab[190672]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:174
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum168_)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:174
			// _ = "end of CoverTab[190672]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:174
		}()
												if message.NotificationDestination != "" {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:175
			_go_fuzz_dep_.CoverTab[190673]++
													statusUpdateMsg := &a1tapi.PolicyStatusMessage{
				PolicyId:	message.PolicyId,
				PolicyType:	message.PolicyType,
				Message: &a1tapi.StatusMessage{
					Header: &a1tapi.Header{
						RequestId:	uuid.New().String(),
						AppId:		message.Message.Header.AppId,
						Encoding:	message.Message.Header.Encoding,
						PayloadType:	a1tapi.PayloadType_STATUS,
					},
				},
				NotificationDestination:	message.NotificationDestination,
			}

			if message.PolicyId == SampleNotEnforcedPolicyID {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:190
				_go_fuzz_dep_.CoverTab[190675]++
														statusUpdateMsg.Message.Payload = []byte(SampleNotEnforcedStatus)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:191
				// _ = "end of CoverTab[190675]"
			} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:192
				_go_fuzz_dep_.CoverTab[190676]++
														statusUpdateMsg.Message.Payload = []byte(SampleEnforcedStatus)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:193
				// _ = "end of CoverTab[190676]"
			}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:194
			// _ = "end of CoverTab[190673]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:194
			_go_fuzz_dep_.CoverTab[190674]++

													a.StatusUpdateCh <- statusUpdateMsg
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:196
			// _ = "end of CoverTab[190674]"
		} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:197
			_go_fuzz_dep_.CoverTab[190677]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:197
			// _ = "end of CoverTab[190677]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:197
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:197
		// _ = "end of CoverTab[190671]"
	}()
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:198
	// _ = "end of CoverTab[190665]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:198
	_go_fuzz_dep_.CoverTab[190666]++

											res := &a1tapi.PolicyResultMessage{
		PolicyId:	message.PolicyId,
		PolicyType:	message.PolicyType,
		Message: &a1tapi.ResultMessage{
			Header: &a1tapi.Header{
				PayloadType:	message.Message.Header.PayloadType,
				RequestId:	message.Message.Header.RequestId,
				Encoding:	message.Message.Header.Encoding,
				AppId:		message.Message.Header.AppId,
			},
			Payload:	a.TsPolicyTypeMap[message.PolicyId],
			Result: &a1tapi.Result{
				Success: true,
			},
		},
		NotificationDestination:	message.NotificationDestination,
	}
											a.notifier <- true
											return res, nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:218
	// _ = "end of CoverTab[190666]"
}

func (a *A1PServer) PolicyUpdate(ctx context.Context, message *a1tapi.PolicyRequestMessage) (*a1tapi.PolicyResultMessage, error) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:221
	_go_fuzz_dep_.CoverTab[190678]++
											a.mu.Lock()
											defer a.mu.Unlock()

											var result map[string]interface{}
											json.Unmarshal(message.Message.Payload, &result)

											if message.PolicyType.Id != "ORAN_TrafficSteeringPreference_2.0.0" {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:228
		_go_fuzz_dep_.CoverTab[190682]++
												res := &a1tapi.PolicyResultMessage{
			PolicyId:	message.PolicyId,
			PolicyType:	message.PolicyType,
			Message: &a1tapi.ResultMessage{
				Header: &a1tapi.Header{
					PayloadType:	message.Message.Header.PayloadType,
					RequestId:	message.Message.Header.RequestId,
					Encoding:	message.Message.Header.Encoding,
					AppId:		message.Message.Header.AppId,
				}, Payload:	message.Message.Payload,
				Result: &a1tapi.Result{
					Success:	false,
					Reason:		"Policy type does not support",
				},
			},
		}
												return res, nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:245
		// _ = "end of CoverTab[190682]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:246
		_go_fuzz_dep_.CoverTab[190683]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:246
		// _ = "end of CoverTab[190683]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:246
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:246
	// _ = "end of CoverTab[190678]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:246
	_go_fuzz_dep_.CoverTab[190679]++

											if _, ok := a.TsPolicyTypeMap[message.PolicyId]; !ok {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:248
		_go_fuzz_dep_.CoverTab[190684]++
												res := &a1tapi.PolicyResultMessage{
			PolicyId:	message.PolicyId,
			PolicyType:	message.PolicyType,
			Message: &a1tapi.ResultMessage{
				Header: &a1tapi.Header{
					PayloadType:	message.Message.Header.PayloadType,
					RequestId:	message.Message.Header.RequestId,
					Encoding:	message.Message.Header.Encoding,
					AppId:		message.Message.Header.AppId,
				}, Payload:	message.Message.Payload,
				Result: &a1tapi.Result{
					Success:	false,
					Reason:		"Policy ID does not exists",
				},
			},
		}
												return res, nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:265
		// _ = "end of CoverTab[190684]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:266
		_go_fuzz_dep_.CoverTab[190685]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:266
		// _ = "end of CoverTab[190685]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:266
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:266
	// _ = "end of CoverTab[190679]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:266
	_go_fuzz_dep_.CoverTab[190680]++

											a.TsPolicyTypeMap[message.PolicyId] = message.Message.Payload
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:268
	_curRoutineNum169_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:268
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum169_)

											go func() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:270
		_go_fuzz_dep_.CoverTab[190686]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:270
		defer func() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:270
			_go_fuzz_dep_.CoverTab[190687]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:270
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum169_)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:270
			// _ = "end of CoverTab[190687]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:270
		}()
												if message.NotificationDestination != "" {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:271
			_go_fuzz_dep_.CoverTab[190688]++
													statusUpdateMsg := &a1tapi.PolicyStatusMessage{
				PolicyId:	message.PolicyId,
				PolicyType:	message.PolicyType,
				Message: &a1tapi.StatusMessage{
					Header: &a1tapi.Header{
						RequestId:	uuid.New().String(),
						AppId:		message.Message.Header.AppId,
						Encoding:	message.Message.Header.Encoding,
						PayloadType:	a1tapi.PayloadType_STATUS,
					},
				},
				NotificationDestination:	message.NotificationDestination,
			}

			if message.PolicyId == SampleNotEnforcedPolicyID {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:286
				_go_fuzz_dep_.CoverTab[190690]++
														statusUpdateMsg.Message.Payload = []byte(SampleNotEnforcedStatus)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:287
				// _ = "end of CoverTab[190690]"
			} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:288
				_go_fuzz_dep_.CoverTab[190691]++
														statusUpdateMsg.Message.Payload = []byte(SampleEnforcedStatus)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:289
				// _ = "end of CoverTab[190691]"
			}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:290
			// _ = "end of CoverTab[190688]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:290
			_go_fuzz_dep_.CoverTab[190689]++

													a.StatusUpdateCh <- statusUpdateMsg
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:292
			// _ = "end of CoverTab[190689]"
		} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:293
			_go_fuzz_dep_.CoverTab[190692]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:293
			// _ = "end of CoverTab[190692]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:293
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:293
		// _ = "end of CoverTab[190686]"
	}()
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:294
	// _ = "end of CoverTab[190680]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:294
	_go_fuzz_dep_.CoverTab[190681]++

											res := &a1tapi.PolicyResultMessage{
		PolicyId:	message.PolicyId,
		PolicyType:	message.PolicyType,
		Message: &a1tapi.ResultMessage{
			Header: &a1tapi.Header{
				PayloadType:	message.Message.Header.PayloadType,
				RequestId:	message.Message.Header.RequestId,
				Encoding:	message.Message.Header.Encoding,
				AppId:		message.Message.Header.AppId,
			}, Payload:	a.TsPolicyTypeMap[message.PolicyId],
			Result: &a1tapi.Result{
				Success: true,
			},
		},
		NotificationDestination:	message.NotificationDestination,
	}
											a.notifier <- true
											return res, nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:313
	// _ = "end of CoverTab[190681]"
}

func (a *A1PServer) PolicyDelete(ctx context.Context, message *a1tapi.PolicyRequestMessage) (*a1tapi.PolicyResultMessage, error) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:316
	_go_fuzz_dep_.CoverTab[190693]++
											a.mu.Lock()
											defer a.mu.Unlock()

											var result map[string]interface{}
											json.Unmarshal(message.Message.Payload, &result)

											if message.PolicyType.Id != "ORAN_TrafficSteeringPreference_2.0.0" {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:323
		_go_fuzz_dep_.CoverTab[190696]++
												res := &a1tapi.PolicyResultMessage{
			PolicyId:	message.PolicyId,
			PolicyType:	message.PolicyType,
			Message: &a1tapi.ResultMessage{
				Header: &a1tapi.Header{
					PayloadType:	message.Message.Header.PayloadType,
					RequestId:	message.Message.Header.RequestId,
					Encoding:	message.Message.Header.Encoding,
					AppId:		message.Message.Header.AppId,
				}, Payload:	message.Message.Payload,
				Result: &a1tapi.Result{
					Success:	false,
					Reason:		"Policy type does not support",
				},
			},
		}
												return res, nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:340
		// _ = "end of CoverTab[190696]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:341
		_go_fuzz_dep_.CoverTab[190697]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:341
		// _ = "end of CoverTab[190697]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:341
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:341
	// _ = "end of CoverTab[190693]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:341
	_go_fuzz_dep_.CoverTab[190694]++

											if _, ok := a.TsPolicyTypeMap[message.PolicyId]; !ok {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:343
		_go_fuzz_dep_.CoverTab[190698]++
												res := &a1tapi.PolicyResultMessage{
			PolicyId:	message.PolicyId,
			PolicyType:	message.PolicyType,
			Message: &a1tapi.ResultMessage{
				Header: &a1tapi.Header{
					PayloadType:	message.Message.Header.PayloadType,
					RequestId:	message.Message.Header.RequestId,
					Encoding:	message.Message.Header.Encoding,
					AppId:		message.Message.Header.AppId,
				}, Payload:	message.Message.Payload,
				Result: &a1tapi.Result{
					Success:	false,
					Reason:		"Policy ID does not exists",
				},
			},
		}
												return res, nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:360
		// _ = "end of CoverTab[190698]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:361
		_go_fuzz_dep_.CoverTab[190699]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:361
		// _ = "end of CoverTab[190699]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:361
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:361
	// _ = "end of CoverTab[190694]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:361
	_go_fuzz_dep_.CoverTab[190695]++

											delete(a.TsPolicyTypeMap, message.PolicyId)

											res := &a1tapi.PolicyResultMessage{
		PolicyId:	message.PolicyId,
		PolicyType:	message.PolicyType,
		Message: &a1tapi.ResultMessage{
			Header: &a1tapi.Header{
				PayloadType:	message.Message.Header.PayloadType,
				RequestId:	message.Message.Header.RequestId,
				Encoding:	message.Message.Header.Encoding,
				AppId:		message.Message.Header.AppId,
			}, Payload:	a.TsPolicyTypeMap[message.PolicyId],
			Result: &a1tapi.Result{
				Success: true,
			},
		},
	}
											a.notifier <- true
											return res, nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:381
	// _ = "end of CoverTab[190695]"
}

func (a *A1PServer) PolicyQuery(ctx context.Context, message *a1tapi.PolicyRequestMessage) (*a1tapi.PolicyResultMessage, error) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:384
	_go_fuzz_dep_.CoverTab[190700]++
											a.mu.Lock()
											defer a.mu.Unlock()
											var result map[string]interface{}
											json.Unmarshal(message.Message.Payload, &result)

											if message.PolicyType.Id != "ORAN_TrafficSteeringPreference_2.0.0" {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:390
		_go_fuzz_dep_.CoverTab[190705]++
												res := &a1tapi.PolicyResultMessage{
			PolicyId:	message.PolicyId,
			PolicyType:	message.PolicyType,
			Message: &a1tapi.ResultMessage{
				Header: &a1tapi.Header{
					PayloadType:	message.Message.Header.PayloadType,
					RequestId:	message.Message.Header.RequestId,
					Encoding:	message.Message.Header.Encoding,
					AppId:		message.Message.Header.AppId,
				}, Payload:	message.Message.Payload,
				Result: &a1tapi.Result{
					Success:	false,
					Reason:		"Policy type does not support",
				},
			},
		}
												return res, nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:407
		// _ = "end of CoverTab[190705]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:408
		_go_fuzz_dep_.CoverTab[190706]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:408
		// _ = "end of CoverTab[190706]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:408
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:408
	// _ = "end of CoverTab[190700]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:408
	_go_fuzz_dep_.CoverTab[190701]++

											if message.PolicyId == "" {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:410
		_go_fuzz_dep_.CoverTab[190707]++

												listPolicies := make([]string, 0)
												for k := range a.TsPolicyTypeMap {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:413
			_go_fuzz_dep_.CoverTab[190710]++
													listPolicies = append(listPolicies, k)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:414
			// _ = "end of CoverTab[190710]"
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:415
		// _ = "end of CoverTab[190707]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:415
		_go_fuzz_dep_.CoverTab[190708]++

												listPoliciesJson, err := json.Marshal(listPolicies)
												if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:418
			_go_fuzz_dep_.CoverTab[190711]++
													log.Error(err)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:419
			// _ = "end of CoverTab[190711]"
		} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:420
			_go_fuzz_dep_.CoverTab[190712]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:420
			// _ = "end of CoverTab[190712]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:420
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:420
		// _ = "end of CoverTab[190708]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:420
		_go_fuzz_dep_.CoverTab[190709]++

												res := &a1tapi.PolicyResultMessage{
			PolicyType:	message.PolicyType,
			Message: &a1tapi.ResultMessage{
				Header: &a1tapi.Header{
					PayloadType:	message.Message.Header.PayloadType,
					RequestId:	message.Message.Header.RequestId,
					Encoding:	message.Message.Header.Encoding,
					AppId:		message.Message.Header.AppId,
				}, Payload:	listPoliciesJson,
				Result: &a1tapi.Result{
					Success: true,
				},
			},
		}
												return res, nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:436
		// _ = "end of CoverTab[190709]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:437
		_go_fuzz_dep_.CoverTab[190713]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:437
		// _ = "end of CoverTab[190713]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:437
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:437
	// _ = "end of CoverTab[190701]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:437
	_go_fuzz_dep_.CoverTab[190702]++

											if _, ok := a.TsPolicyTypeMap[message.PolicyId]; !ok {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:439
		_go_fuzz_dep_.CoverTab[190714]++
												res := &a1tapi.PolicyResultMessage{
			PolicyId:	message.PolicyId,
			PolicyType:	message.PolicyType,
			Message: &a1tapi.ResultMessage{
				Header: &a1tapi.Header{
					PayloadType:	message.Message.Header.PayloadType,
					RequestId:	message.Message.Header.RequestId,
					Encoding:	message.Message.Header.Encoding,
					AppId:		message.Message.Header.AppId,
				}, Payload:	message.Message.Payload,
				Result: &a1tapi.Result{
					Success:	false,
					Reason:		"Policy ID does not exists",
				},
			},
		}
												return res, nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:456
		// _ = "end of CoverTab[190714]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:457
		_go_fuzz_dep_.CoverTab[190715]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:457
		// _ = "end of CoverTab[190715]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:457
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:457
	// _ = "end of CoverTab[190702]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:457
	_go_fuzz_dep_.CoverTab[190703]++

											resultMsg := &a1tapi.PolicyResultMessage{
		PolicyId:	message.PolicyId,
		PolicyType:	message.PolicyType,
		Message: &a1tapi.ResultMessage{
			Header: &a1tapi.Header{
				PayloadType:	message.Message.Header.PayloadType,
				RequestId:	message.Message.Header.RequestId,
				Encoding:	message.Message.Header.Encoding,
				AppId:		message.Message.Header.AppId,
			},
			Result: &a1tapi.Result{
				Success: true,
			},
		},
	}

	switch message.Message.Header.PayloadType {
	case a1tapi.PayloadType_POLICY:
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:476
		_go_fuzz_dep_.CoverTab[190716]++
												resultMsg.Message.Payload = a.TsPolicyTypeMap[message.PolicyId]
												resultMsg.Message.Header.PayloadType = a1tapi.PayloadType_POLICY
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:478
		// _ = "end of CoverTab[190716]"
	case a1tapi.PayloadType_STATUS:
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:479
		_go_fuzz_dep_.CoverTab[190717]++
												resultMsg.Message.Header.PayloadType = a1tapi.PayloadType_STATUS
												if message.PolicyId == SampleNotEnforcedPolicyID {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:481
			_go_fuzz_dep_.CoverTab[190719]++
													resultMsg.Message.Payload = []byte(SampleNotEnforcedStatus)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:482
			// _ = "end of CoverTab[190719]"

		} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:484
			_go_fuzz_dep_.CoverTab[190720]++
													resultMsg.Message.Payload = []byte(SampleEnforcedStatus)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:485
			// _ = "end of CoverTab[190720]"
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:486
		// _ = "end of CoverTab[190717]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:486
	default:
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:486
		_go_fuzz_dep_.CoverTab[190718]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:486
		// _ = "end of CoverTab[190718]"
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:487
	// _ = "end of CoverTab[190703]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:487
	_go_fuzz_dep_.CoverTab[190704]++
											return resultMsg, nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:488
	// _ = "end of CoverTab[190704]"
}

func (a *A1PServer) PolicyStatus(server a1tapi.PolicyService_PolicyStatusServer) error {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:491
	_go_fuzz_dep_.CoverTab[190721]++

											watchers := make(map[uuid.UUID]chan *a1tapi.PolicyAckMessage)
											mu := sync.RWMutex{}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:494
	_curRoutineNum170_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:494
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum170_)

											go func(m *sync.RWMutex) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:496
		_go_fuzz_dep_.CoverTab[190724]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:496
		defer func() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:496
			_go_fuzz_dep_.CoverTab[190725]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:496
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum170_)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:496
			// _ = "end of CoverTab[190725]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:496
		}()
												for {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:497
			_go_fuzz_dep_.CoverTab[190726]++
													ack, err := server.Recv()
													if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:499
				_go_fuzz_dep_.CoverTab[190729]++
														log.Error(err)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:500
				// _ = "end of CoverTab[190729]"
			} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:501
				_go_fuzz_dep_.CoverTab[190730]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:501
				// _ = "end of CoverTab[190730]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:501
			}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:501
			// _ = "end of CoverTab[190726]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:501
			_go_fuzz_dep_.CoverTab[190727]++
													m.Lock()
													for _, v := range watchers {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:503
				_go_fuzz_dep_.CoverTab[190731]++
														select {
				case v <- ack:
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:505
					_go_fuzz_dep_.CoverTab[190732]++
															log.Debugf("Sent msg %v on %v", ack, v)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:506
					// _ = "end of CoverTab[190732]"
				default:
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:507
					_go_fuzz_dep_.CoverTab[190733]++
															log.Debugf("Failed to send msg %v on %v", ack, v)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:508
					// _ = "end of CoverTab[190733]"
				}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:509
				// _ = "end of CoverTab[190731]"
			}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:510
			// _ = "end of CoverTab[190727]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:510
			_go_fuzz_dep_.CoverTab[190728]++
													m.Unlock()
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:511
			// _ = "end of CoverTab[190728]"
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:512
		// _ = "end of CoverTab[190724]"
	}(&a.mu)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:513
	// _ = "end of CoverTab[190721]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:513
	_go_fuzz_dep_.CoverTab[190722]++

											for msg := range a.StatusUpdateCh {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:515
		_go_fuzz_dep_.CoverTab[190734]++
												watcherID := uuid.New()
												ackCh := make(chan *a1tapi.PolicyAckMessage)
												timerCh := make(chan bool, 1)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:518
		_curRoutineNum171_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:518
		_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum171_)
												go func(ch chan bool) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:519
			_go_fuzz_dep_.CoverTab[190737]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:519
			defer func() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:519
				_go_fuzz_dep_.CoverTab[190738]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:519
				_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum171_)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:519
				// _ = "end of CoverTab[190738]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:519
			}()
													time.Sleep(5 * time.Second)
													timerCh <- true
													close(timerCh)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:522
			// _ = "end of CoverTab[190737]"
		}(timerCh)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:523
		// _ = "end of CoverTab[190734]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:523
		_go_fuzz_dep_.CoverTab[190735]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:523
		_curRoutineNum172_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:523
		_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum172_)

												go func(m *sync.RWMutex) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:525
			_go_fuzz_dep_.CoverTab[190739]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:525
			defer func() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:525
				_go_fuzz_dep_.CoverTab[190740]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:525
				_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum172_)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:525
				// _ = "end of CoverTab[190740]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:525
			}()
													for {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:526
				_go_fuzz_dep_.CoverTab[190741]++
														select {
				case ack := <-ackCh:
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:528
					_go_fuzz_dep_.CoverTab[190742]++
															if ack.Message.Header.RequestId == msg.Message.Header.RequestId {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:529
						_go_fuzz_dep_.CoverTab[190744]++
																m.Lock()
																close(ackCh)
																delete(watchers, watcherID)
																m.Unlock()
																return
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:534
						// _ = "end of CoverTab[190744]"
					} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:535
						_go_fuzz_dep_.CoverTab[190745]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:535
						// _ = "end of CoverTab[190745]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:535
					}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:535
					// _ = "end of CoverTab[190742]"
				case <-timerCh:
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:536
					_go_fuzz_dep_.CoverTab[190743]++
															log.Error(fmt.Errorf("could not receive PolicyACKMessage in timer"))
															m.Lock()
															close(ackCh)
															delete(watchers, watcherID)
															m.Unlock()
															return
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:542
					// _ = "end of CoverTab[190743]"
				}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:543
				// _ = "end of CoverTab[190741]"
			}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:544
			// _ = "end of CoverTab[190739]"
		}(&a.mu)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:545
		// _ = "end of CoverTab[190735]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:545
		_go_fuzz_dep_.CoverTab[190736]++

												mu.Lock()
												watchers[watcherID] = ackCh
												mu.Unlock()

												err := server.Send(msg)
												if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:552
			_go_fuzz_dep_.CoverTab[190746]++
													log.Error(err)
													mu.Lock()
													close(ackCh)
													delete(watchers, watcherID)
													mu.Unlock()
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:557
			// _ = "end of CoverTab[190746]"
		} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:558
			_go_fuzz_dep_.CoverTab[190747]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:558
			// _ = "end of CoverTab[190747]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:558
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:558
		// _ = "end of CoverTab[190736]"
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:559
	// _ = "end of CoverTab[190722]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:559
	_go_fuzz_dep_.CoverTab[190723]++

											return nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:561
	// _ = "end of CoverTab[190723]"
}

//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:562
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/northbound/a1/a1p-service.go:562
var _ = _go_fuzz_dep_.CoverTab
