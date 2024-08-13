//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:1
// SPDX-FileCopyrightText: 2019-present Open Networking Foundation <info@opennetworking.org>
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:1
// SPDX-FileCopyrightText: 2019-present Rimedo Labs
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:1
//
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:1
// SPDX-License-Identifier: Apache-2.0
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:1
// Copy from onosproject/onos-mho/pkg/monitoring/monitor.go
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:1
// modified by RIMEDO-Labs team
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:7
package mho

//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:7
import (
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:7
)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:7
import (
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:7
	_atomic_ "sync/atomic"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:7
)

import (
	"fmt"
	"strconv"

	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-mho-go"
	e2sm_v2_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-v2-ies"
)

func PlmnIDBytesToInt(b []byte) uint64 {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:17
	_go_fuzz_dep_.CoverTab[179747]++
									if len(b) > 3 {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:18
		_go_fuzz_dep_.CoverTab[179749]++
										return uint64(b[2])<<16 | uint64(b[1])<<8 | uint64(b[0])
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:19
		// _ = "end of CoverTab[179749]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:20
		_go_fuzz_dep_.CoverTab[179750]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:20
		// _ = "end of CoverTab[179750]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:20
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:20
	// _ = "end of CoverTab[179747]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:20
	_go_fuzz_dep_.CoverTab[179748]++
									return 0
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:21
	// _ = "end of CoverTab[179748]"
}

func PlmnIDNciToCGI(plmnID uint64, nci uint64) string {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:24
	_go_fuzz_dep_.CoverTab[179751]++
									cgi := strconv.FormatInt(int64(plmnID<<36|(nci&0xfffffffff)), 16)
									return cgi
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:26
	// _ = "end of CoverTab[179751]"
}

func GetNciFromCellGlobalID(cellGlobalID *e2sm_v2_ies.Cgi) uint64 {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:29
	_go_fuzz_dep_.CoverTab[179752]++
									return BitStringToUint64(cellGlobalID.GetNRCgi().GetNRcellIdentity().GetValue().GetValue(), int(cellGlobalID.GetNRCgi().GetNRcellIdentity().GetValue().GetLen()))
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:30
	// _ = "end of CoverTab[179752]"
}

func GetPlmnIDBytesFromCellGlobalID(cellGlobalID *e2sm_v2_ies.Cgi) []byte {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:33
	_go_fuzz_dep_.CoverTab[179753]++
									return cellGlobalID.GetNRCgi().GetPLmnidentity().GetValue()
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:34
	// _ = "end of CoverTab[179753]"
}

func GetMccMncFromPlmnID(plmnId uint64) (string, string) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:37
	_go_fuzz_dep_.CoverTab[179754]++
									plmnIdString := strconv.FormatUint(plmnId, 16)
									if len(plmnIdString) > 3 {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:39
		_go_fuzz_dep_.CoverTab[179756]++
										return plmnIdString[0:3], plmnIdString[3:]
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:40
		// _ = "end of CoverTab[179756]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:41
		_go_fuzz_dep_.CoverTab[179757]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:41
		// _ = "end of CoverTab[179757]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:41
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:41
	// _ = "end of CoverTab[179754]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:41
	_go_fuzz_dep_.CoverTab[179755]++
									return "", ""
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:42
	// _ = "end of CoverTab[179755]"
}

func GetPlmnIdFromMccMnc(mcc string, mnc string) (uint64, error) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:45
	_go_fuzz_dep_.CoverTab[179758]++
									combined := mcc + mnc
									plmnId, err := strconv.ParseUint(combined, 16, 64)
									if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:48
		_go_fuzz_dep_.CoverTab[179760]++
										log.Warn("Cannot convert PLMN ID string into uint64 type!")
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:49
		// _ = "end of CoverTab[179760]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:50
		_go_fuzz_dep_.CoverTab[179761]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:50
		// _ = "end of CoverTab[179761]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:50
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:50
	// _ = "end of CoverTab[179758]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:50
	_go_fuzz_dep_.CoverTab[179759]++
									return plmnId, err
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:51
	// _ = "end of CoverTab[179759]"
}

func GetCGIFromIndicationHeader(header *e2sm_mho.E2SmMhoIndicationHeaderFormat1) string {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:54
	_go_fuzz_dep_.CoverTab[179762]++
									nci := GetNciFromCellGlobalID(header.GetCgi())
									plmnIDBytes := GetPlmnIDBytesFromCellGlobalID(header.GetCgi())
									plmnID := PlmnIDBytesToInt(plmnIDBytes)
									return PlmnIDNciToCGI(plmnID, nci)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:58
	// _ = "end of CoverTab[179762]"
}

func GetCGIFromMeasReportItem(measReport *e2sm_mho.E2SmMhoMeasurementReportItem) string {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:61
	_go_fuzz_dep_.CoverTab[179763]++
									nci := GetNciFromCellGlobalID(measReport.GetCgi())
									plmnIDBytes := GetPlmnIDBytesFromCellGlobalID(measReport.GetCgi())
									plmnID := PlmnIDBytesToInt(plmnIDBytes)
									return PlmnIDNciToCGI(plmnID, nci)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:65
	// _ = "end of CoverTab[179763]"
}

func BitStringToUint64(bitString []byte, bitCount int) uint64 {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:68
	_go_fuzz_dep_.CoverTab[179764]++
									var result uint64
									for i, b := range bitString {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:70
		_go_fuzz_dep_.CoverTab[179767]++
										result += uint64(b) << ((len(bitString) - i - 1) * 8)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:71
		// _ = "end of CoverTab[179767]"
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:72
	// _ = "end of CoverTab[179764]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:72
	_go_fuzz_dep_.CoverTab[179765]++
									if bitCount%8 != 0 {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:73
		_go_fuzz_dep_.CoverTab[179768]++
										return result >> (8 - bitCount%8)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:74
		// _ = "end of CoverTab[179768]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:75
		_go_fuzz_dep_.CoverTab[179769]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:75
		// _ = "end of CoverTab[179769]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:75
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:75
	// _ = "end of CoverTab[179765]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:75
	_go_fuzz_dep_.CoverTab[179766]++
									return result
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:76
	// _ = "end of CoverTab[179766]"
}

func GetUeID(ueID *e2sm_v2_ies.Ueid) (int64, error) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:79
	_go_fuzz_dep_.CoverTab[179770]++

									switch ue := ueID.Ueid.(type) {
	case *e2sm_v2_ies.Ueid_GNbUeid:
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:82
		_go_fuzz_dep_.CoverTab[179771]++
										return ue.GNbUeid.GetAmfUeNgapId().GetValue(), nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:83
		// _ = "end of CoverTab[179771]"
	case *e2sm_v2_ies.Ueid_ENbUeid:
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:84
		_go_fuzz_dep_.CoverTab[179772]++
										return ue.ENbUeid.GetMMeUeS1ApId().GetValue(), nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:85
		// _ = "end of CoverTab[179772]"
	case *e2sm_v2_ies.Ueid_EnGNbUeid:
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:86
		_go_fuzz_dep_.CoverTab[179773]++
										return int64(ue.EnGNbUeid.GetMENbUeX2ApId().GetValue()), nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:87
		// _ = "end of CoverTab[179773]"
	case *e2sm_v2_ies.Ueid_NgENbUeid:
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:88
		_go_fuzz_dep_.CoverTab[179774]++
										return ue.NgENbUeid.GetAmfUeNgapId().GetValue(), nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:89
		// _ = "end of CoverTab[179774]"
	default:
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:90
		_go_fuzz_dep_.CoverTab[179775]++
										return -1, fmt.Errorf("GetUeID() couldn't extract UeID - obtained unexpected type %v", ue)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:91
		// _ = "end of CoverTab[179775]"
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:92
	// _ = "end of CoverTab[179770]"
}

//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:93
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/mho/reader.go:93
var _ = _go_fuzz_dep_.CoverTab
