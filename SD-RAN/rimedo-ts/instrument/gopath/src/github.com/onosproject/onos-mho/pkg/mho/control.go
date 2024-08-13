// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:5
package mho

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:5
)

import (
	e2tapi "github.com/onosproject/onos-api/go/onos/e2t/e2"
	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/pdubuilder"
	e2sm_v2_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-v2-ies"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"github.com/onosproject/rrm-son-lib/pkg/handover"
	"google.golang.org/protobuf/proto"
	"strconv"
)

type E2SmMhoControlHandler struct {
	NodeID			string
	ControlMessage		[]byte
	ControlHeader		[]byte
	ControlAckRequest	e2tapi.ControlAckRequest
}

func (c *E2SmMhoControlHandler) CreateMhoControlRequest() (*e2api.ControlMessage, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:25
	_go_fuzz_dep_.CoverTab[194450]++
												return &e2api.ControlMessage{
		Header:		c.ControlHeader,
		Payload:	c.ControlMessage,
	}, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:29
	// _ = "end of CoverTab[194450]"
}

func (c *E2SmMhoControlHandler) CreateMhoControlHeader(cellID []byte, cellIDLen uint32, priority int32, plmnID []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:32
	_go_fuzz_dep_.CoverTab[194451]++
												eci := &asn1.BitString{
		Value:	cellID,
		Len:	cellIDLen,
	}
	cgi, err := pdubuilder.CreateCgiNrCGI(plmnID, eci)
	log.Debugf("eci: %v", eci)
	log.Debugf("cgi: %v", cgi)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:40
		_go_fuzz_dep_.CoverTab[194456]++
													return []byte{}, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:41
		// _ = "end of CoverTab[194456]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:42
		_go_fuzz_dep_.CoverTab[194457]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:42
		// _ = "end of CoverTab[194457]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:42
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:42
	// _ = "end of CoverTab[194451]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:42
	_go_fuzz_dep_.CoverTab[194452]++

												newE2SmMhoPdu, err := pdubuilder.CreateE2SmMhoControlHeader(priority)

												log.Debugf("newE2SmMhoPdu (ControlHeader): %v", newE2SmMhoPdu)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:47
		_go_fuzz_dep_.CoverTab[194458]++
													return []byte{}, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:48
		// _ = "end of CoverTab[194458]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:49
		_go_fuzz_dep_.CoverTab[194459]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:49
		// _ = "end of CoverTab[194459]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:49
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:49
	// _ = "end of CoverTab[194452]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:49
	_go_fuzz_dep_.CoverTab[194453]++

												err = newE2SmMhoPdu.Validate()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:52
		_go_fuzz_dep_.CoverTab[194460]++
													return []byte{}, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:53
		// _ = "end of CoverTab[194460]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:54
		_go_fuzz_dep_.CoverTab[194461]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:54
		// _ = "end of CoverTab[194461]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:54
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:54
	// _ = "end of CoverTab[194453]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:54
	_go_fuzz_dep_.CoverTab[194454]++

												protoBytes, err := proto.Marshal(newE2SmMhoPdu)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:57
		_go_fuzz_dep_.CoverTab[194462]++
													return []byte{}, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:58
		// _ = "end of CoverTab[194462]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:59
		_go_fuzz_dep_.CoverTab[194463]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:59
		// _ = "end of CoverTab[194463]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:59
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:59
	// _ = "end of CoverTab[194454]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:59
	_go_fuzz_dep_.CoverTab[194455]++

												return protoBytes, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:61
	// _ = "end of CoverTab[194455]"
}

func (c *E2SmMhoControlHandler) CreateMhoControlMessage(servingCgi *e2sm_v2_ies.Cgi, uedID *e2sm_v2_ies.Ueid, targetCgi *e2sm_v2_ies.Cgi) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:64
	_go_fuzz_dep_.CoverTab[194464]++

												var err error

												if newE2SmMhoPdu, err := pdubuilder.CreateE2SmMhoControlMessage(servingCgi, uedID, targetCgi); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:68
		_go_fuzz_dep_.CoverTab[194466]++
													if err = newE2SmMhoPdu.Validate(); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:69
			_go_fuzz_dep_.CoverTab[194467]++
														log.Debugf("newE2SmMhoPdu (ControlMessage): %v", newE2SmMhoPdu)
														if protoBytes, err := proto.Marshal(newE2SmMhoPdu); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:71
				_go_fuzz_dep_.CoverTab[194468]++
															return protoBytes, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:72
				// _ = "end of CoverTab[194468]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:73
				_go_fuzz_dep_.CoverTab[194469]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:73
				// _ = "end of CoverTab[194469]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:73
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:73
			// _ = "end of CoverTab[194467]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:74
			_go_fuzz_dep_.CoverTab[194470]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:74
			// _ = "end of CoverTab[194470]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:74
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:74
		// _ = "end of CoverTab[194466]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:75
		_go_fuzz_dep_.CoverTab[194471]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:75
		// _ = "end of CoverTab[194471]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:75
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:75
	// _ = "end of CoverTab[194464]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:75
	_go_fuzz_dep_.CoverTab[194465]++

												return []byte{}, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:77
	// _ = "end of CoverTab[194465]"

}

func SendHORequest(ueData *UeData, ho handover.A3HandoverDecision, ctrlReqChan chan *e2api.ControlMessage) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:81
	_go_fuzz_dep_.CoverTab[194472]++
												e2NodeID := ueData.E2NodeID
												servingCGI := ueData.CGI
												servingPlmnIDBytes := servingCGI.GetNRCgi().GetPLmnidentity().GetValue()
												servingNCI := servingCGI.GetNRCgi().GetNRcellIdentity().GetValue().GetValue()
												servingNCILen := servingCGI.GetNRCgi().GetNRcellIdentity().GetValue().GetLen()
												targetPlmnIDBytes := servingPlmnIDBytes
												targetNCI, err := strconv.Atoi(ho.TargetCell.GetID().String())
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:89
		_go_fuzz_dep_.CoverTab[194476]++
													panic("bad data")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:90
		// _ = "end of CoverTab[194476]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:91
		_go_fuzz_dep_.CoverTab[194477]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:91
		// _ = "end of CoverTab[194477]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:91
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:91
	// _ = "end of CoverTab[194472]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:91
	_go_fuzz_dep_.CoverTab[194473]++
												targetNCILen := 36

												e2smMhoControlHandler := &E2SmMhoControlHandler{
		NodeID:			e2NodeID,
		ControlAckRequest:	e2tapi.ControlAckRequest_NO_ACK,
	}

	targetCGI := &e2sm_v2_ies.Cgi{
		Cgi: &e2sm_v2_ies.Cgi_NRCgi{
			NRCgi: &e2sm_v2_ies.NrCgi{
				PLmnidentity: &e2sm_v2_ies.PlmnIdentity{
					Value: targetPlmnIDBytes,
				},
				NRcellIdentity: &e2sm_v2_ies.NrcellIdentity{
					Value: &asn1.BitString{
						Value:	Uint64ToBitString(uint64(targetNCI), targetNCILen),
						Len:	uint32(targetNCILen),
					},
				},
			},
		},
	}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:116
	ueIDnum, err := strconv.Atoi(ueData.UeID)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:117
		_go_fuzz_dep_.CoverTab[194478]++
														log.Errorf("SendHORequest() failed to convert string %v to decimal number - assumption is not satisfied (UEID is a decimal number): %v", ueData.UeID, err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:118
		// _ = "end of CoverTab[194478]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:119
		_go_fuzz_dep_.CoverTab[194479]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:119
		// _ = "end of CoverTab[194479]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:119
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:119
	// _ = "end of CoverTab[194473]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:119
	_go_fuzz_dep_.CoverTab[194474]++

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:122
	ueIdentity, err := pdubuilder.CreateUeIDGNb(int64(ueIDnum), []byte{0xAA, 0xBB, 0xCC}, []byte{0xDD}, []byte{0xCC, 0xC0}, []byte{0xFC})
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:123
		_go_fuzz_dep_.CoverTab[194480]++
														log.Errorf("SendHORequest() Failed to create UEID: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:124
		// _ = "end of CoverTab[194480]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:125
		_go_fuzz_dep_.CoverTab[194481]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:125
		// _ = "end of CoverTab[194481]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:125
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:125
	// _ = "end of CoverTab[194474]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:125
	_go_fuzz_dep_.CoverTab[194475]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:125
	_curRoutineNum177_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:125
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum177_)

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:130
	go func() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:130
		_go_fuzz_dep_.CoverTab[194482]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:130
		defer func() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:130
			_go_fuzz_dep_.CoverTab[194483]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:130
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum177_)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:130
			// _ = "end of CoverTab[194483]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:130
		}()
														if e2smMhoControlHandler.ControlHeader, err = e2smMhoControlHandler.CreateMhoControlHeader(servingNCI, servingNCILen, int32(ControlPriority), servingPlmnIDBytes); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:131
			_go_fuzz_dep_.CoverTab[194484]++
															if e2smMhoControlHandler.ControlMessage, err = e2smMhoControlHandler.CreateMhoControlMessage(servingCGI, ueIdentity, targetCGI); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:132
				_go_fuzz_dep_.CoverTab[194485]++
																if controlRequest, err := e2smMhoControlHandler.CreateMhoControlRequest(); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:133
					_go_fuzz_dep_.CoverTab[194486]++
																	ctrlReqChan <- controlRequest
																	log.Infof("tx control, e2NodeID:%v, ueID:%v", e2NodeID, ueData.UeID)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:135
					// _ = "end of CoverTab[194486]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:136
					_go_fuzz_dep_.CoverTab[194487]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:136
					// _ = "end of CoverTab[194487]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:136
				}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:136
				// _ = "end of CoverTab[194485]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:137
				_go_fuzz_dep_.CoverTab[194488]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:137
				// _ = "end of CoverTab[194488]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:137
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:137
			// _ = "end of CoverTab[194484]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:138
			_go_fuzz_dep_.CoverTab[194489]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:138
			// _ = "end of CoverTab[194489]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:138
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:138
		// _ = "end of CoverTab[194482]"
	}()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:139
	// _ = "end of CoverTab[194475]"

}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:141
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/control.go:141
var _ = _go_fuzz_dep_.CoverTab
