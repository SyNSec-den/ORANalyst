// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/v2/e2sm-mho-go/builder.go:5
package e2sm_mho_go

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/v2/e2sm-mho-go/builder.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/v2/e2sm-mho-go/builder.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/v2/e2sm-mho-go/builder.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/v2/e2sm-mho-go/builder.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/v2/e2sm-mho-go/builder.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/v2/e2sm-mho-go/builder.go:5
)

import e2sm_v2_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-v2-ies"

func (ed *E2SmMhoEventTriggerDefinitionFormat1) SetReportingPeriodInMs(rp int32) *E2SmMhoEventTriggerDefinitionFormat1 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/v2/e2sm-mho-go/builder.go:9
	_go_fuzz_dep_.CoverTab[177535]++
																	ed.ReportingPeriodMs = &rp
																	return ed
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/v2/e2sm-mho-go/builder.go:11
	// _ = "end of CoverTab[177535]"
}

func (ed *E2SmMhoControlHeaderFormat1) SetRicControlMessagePriority(cmp int32) *E2SmMhoControlHeaderFormat1 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/v2/e2sm-mho-go/builder.go:14
	_go_fuzz_dep_.CoverTab[177536]++
																	ed.RicControlMessagePriority = &RicControlMessagePriority{
		Value: cmp,
	}
																	return ed
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/v2/e2sm-mho-go/builder.go:18
	// _ = "end of CoverTab[177536]"
}

func (rfd *E2SmMhoRanfunctionDescription) SetRicEventTriggerStyleList(retsl []*RicEventTriggerStyleList) *E2SmMhoRanfunctionDescription {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/v2/e2sm-mho-go/builder.go:21
	_go_fuzz_dep_.CoverTab[177537]++
																	rfd.GetE2SmMhoRanfunctionItem().RicEventTriggerStyleList = make([]*RicEventTriggerStyleList, 0)
																	rfd.GetE2SmMhoRanfunctionItem().RicEventTriggerStyleList = retsl
																	return rfd
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/v2/e2sm-mho-go/builder.go:24
	// _ = "end of CoverTab[177537]"
}

func (rfd *E2SmMhoRanfunctionDescription) SetRicReportStyleList(rrsl []*RicReportStyleList) *E2SmMhoRanfunctionDescription {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/v2/e2sm-mho-go/builder.go:27
	_go_fuzz_dep_.CoverTab[177538]++
																	rfd.GetE2SmMhoRanfunctionItem().RicReportStyleList = make([]*RicReportStyleList, 0)
																	rfd.GetE2SmMhoRanfunctionItem().RicReportStyleList = rrsl
																	return rfd
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/v2/e2sm-mho-go/builder.go:30
	// _ = "end of CoverTab[177538]"
}

func (mri *E2SmMhoMeasurementReportItem) SetFiveQi(fiveQI int32) *E2SmMhoMeasurementReportItem {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/v2/e2sm-mho-go/builder.go:33
	_go_fuzz_dep_.CoverTab[177539]++
																	mri.FiveQi = &e2sm_v2_ies.FiveQi{
		Value: fiveQI,
	}
																	return mri
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/v2/e2sm-mho-go/builder.go:37
	// _ = "end of CoverTab[177539]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/v2/e2sm-mho-go/builder.go:38
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go@v0.8.5/v2/e2sm-mho-go/builder.go:38
var _ = _go_fuzz_dep_.CoverTab
