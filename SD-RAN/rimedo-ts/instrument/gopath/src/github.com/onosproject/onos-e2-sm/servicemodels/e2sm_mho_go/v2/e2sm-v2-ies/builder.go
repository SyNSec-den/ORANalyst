// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/v2/e2sm-v2-ies/builder.go:5
package e2sm_v2_ies

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/v2/e2sm-v2-ies/builder.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/v2/e2sm-v2-ies/builder.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/v2/e2sm-v2-ies/builder.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/v2/e2sm-v2-ies/builder.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/v2/e2sm-v2-ies/builder.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/v2/e2sm-v2-ies/builder.go:5
)

import "github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"

func (m *RanfunctionName) SetRanFunctionInstance(rfi int32) *RanfunctionName {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/v2/e2sm-v2-ies/builder.go:9
	_go_fuzz_dep_.CoverTab[171436]++
																	m.RanFunctionInstance = &rfi
																	return m
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/v2/e2sm-v2-ies/builder.go:11
	// _ = "end of CoverTab[171436]"
}

func (m *UeidGnb) SetGNbCuUeF1ApIDList(list []int64) *UeidGnb {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/v2/e2sm-v2-ies/builder.go:14
	_go_fuzz_dep_.CoverTab[171437]++
																	m.GNbCuUeF1ApIdList = &UeidGnbCuF1ApIdList{
		Value: make([]*UeidGnbCuCpF1ApIdItem, 0),
	}

	for _, val := range list {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/v2/e2sm-v2-ies/builder.go:19
		_go_fuzz_dep_.CoverTab[171439]++
																		item := &UeidGnbCuCpF1ApIdItem{
			GNbCuUeF1ApId: &GnbCuUeF1ApId{
				Value: val,
			},
		}
																		m.GNbCuUeF1ApIdList.Value = append(m.GNbCuUeF1ApIdList.Value, item)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/v2/e2sm-v2-ies/builder.go:25
		// _ = "end of CoverTab[171439]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/v2/e2sm-v2-ies/builder.go:26
	// _ = "end of CoverTab[171437]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/v2/e2sm-v2-ies/builder.go:26
	_go_fuzz_dep_.CoverTab[171438]++

																	return m
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/v2/e2sm-v2-ies/builder.go:28
	// _ = "end of CoverTab[171438]"
}

func (m *UeidGnb) SetGNbCuCpUeE1ApIDList(list []int64) *UeidGnb {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/v2/e2sm-v2-ies/builder.go:31
	_go_fuzz_dep_.CoverTab[171440]++
																	m.GNbCuCpUeE1ApIdList = &UeidGnbCuCpE1ApIdList{
		Value: make([]*UeidGnbCuCpE1ApIdItem, 0),
	}

	for _, val := range list {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/v2/e2sm-v2-ies/builder.go:36
		_go_fuzz_dep_.CoverTab[171442]++
																		item := &UeidGnbCuCpE1ApIdItem{
			GNbCuCpUeE1ApId: &GnbCuCpUeE1ApId{
				Value: val,
			},
		}
																		m.GNbCuCpUeE1ApIdList.Value = append(m.GNbCuCpUeE1ApIdList.Value, item)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/v2/e2sm-v2-ies/builder.go:42
		// _ = "end of CoverTab[171442]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/v2/e2sm-v2-ies/builder.go:43
	// _ = "end of CoverTab[171440]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/v2/e2sm-v2-ies/builder.go:43
	_go_fuzz_dep_.CoverTab[171441]++

																	return m
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/v2/e2sm-v2-ies/builder.go:45
	// _ = "end of CoverTab[171441]"
}

func (m *UeidGnb) SetRanUeID(ranUeID []byte) *UeidGnb {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/v2/e2sm-v2-ies/builder.go:48
	_go_fuzz_dep_.CoverTab[171443]++
																	m.RanUeid = &Ranueid{
		Value: ranUeID,
	}

																	return m
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/v2/e2sm-v2-ies/builder.go:53
	// _ = "end of CoverTab[171443]"
}

func (m *UeidGnb) SetMNgRanUeXnApID(ngRanNodeID int64) *UeidGnb {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/v2/e2sm-v2-ies/builder.go:56
	_go_fuzz_dep_.CoverTab[171444]++
																	m.MNgRanUeXnApId = &NgRannodeUexnApid{
		Value: ngRanNodeID,
	}

																	return m
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/v2/e2sm-v2-ies/builder.go:61
	// _ = "end of CoverTab[171444]"
}

func (m *UeidGnb) SetGlobalGnbID(plmnID []byte, gnbID *asn1.BitString) *UeidGnb {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/v2/e2sm-v2-ies/builder.go:64
	_go_fuzz_dep_.CoverTab[171445]++
																	m.GlobalGnbId = &GlobalGnbId{
		PLmnidentity: &PlmnIdentity{
			Value: plmnID,
		},
		GNbId: &GnbId{
			GnbId: &GnbId_GNbId{
				GNbId: gnbID,
			},
		},
	}

																	return m
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/v2/e2sm-v2-ies/builder.go:76
	// _ = "end of CoverTab[171445]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/v2/e2sm-v2-ies/builder.go:77
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/v2/e2sm-v2-ies/builder.go:77
var _ = _go_fuzz_dep_.CoverTab
