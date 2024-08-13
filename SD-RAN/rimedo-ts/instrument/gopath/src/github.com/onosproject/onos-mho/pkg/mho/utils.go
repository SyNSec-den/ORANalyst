// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

// Copied over onos-e2t/test/utils/bitstring.go

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/utils.go:7
package mho

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/utils.go:7
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/utils.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/utils.go:7
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/utils.go:7
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/utils.go:7
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/utils.go:7
)

import (
	"fmt"
	e2sm_v2_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-v2-ies"
)

// Uint64ToBitString converts uint64 to a bit string byte array
func Uint64ToBitString(value uint64, bitCount int) []byte {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/utils.go:15
	_go_fuzz_dep_.CoverTab[194635]++
												result := make([]byte, bitCount/8+1)
												if bitCount%8 > 0 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/utils.go:17
		_go_fuzz_dep_.CoverTab[194638]++
													value = value << (8 - bitCount%8)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/utils.go:18
		// _ = "end of CoverTab[194638]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/utils.go:19
		_go_fuzz_dep_.CoverTab[194639]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/utils.go:19
		// _ = "end of CoverTab[194639]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/utils.go:19
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/utils.go:19
	// _ = "end of CoverTab[194635]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/utils.go:19
	_go_fuzz_dep_.CoverTab[194636]++

												for i := 0; i <= (bitCount / 8); i++ {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/utils.go:21
		_go_fuzz_dep_.CoverTab[194640]++
													result[i] = byte(value >> (((bitCount / 8) - i) * 8) & 0xFF)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/utils.go:22
		// _ = "end of CoverTab[194640]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/utils.go:23
	// _ = "end of CoverTab[194636]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/utils.go:23
	_go_fuzz_dep_.CoverTab[194637]++

												return result
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/utils.go:25
	// _ = "end of CoverTab[194637]"
}

// BitStringToUint64 converts bit string to uint 64
func BitStringToUint64(bitString []byte, bitCount int) uint64 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/utils.go:29
	_go_fuzz_dep_.CoverTab[194641]++
												var result uint64
												for i, b := range bitString {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/utils.go:31
		_go_fuzz_dep_.CoverTab[194644]++
													result += uint64(b) << ((len(bitString) - i - 1) * 8)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/utils.go:32
		// _ = "end of CoverTab[194644]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/utils.go:33
	// _ = "end of CoverTab[194641]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/utils.go:33
	_go_fuzz_dep_.CoverTab[194642]++
												if bitCount%8 != 0 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/utils.go:34
		_go_fuzz_dep_.CoverTab[194645]++
													return result >> (8 - bitCount%8)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/utils.go:35
		// _ = "end of CoverTab[194645]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/utils.go:36
		_go_fuzz_dep_.CoverTab[194646]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/utils.go:36
		// _ = "end of CoverTab[194646]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/utils.go:36
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/utils.go:36
	// _ = "end of CoverTab[194642]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/utils.go:36
	_go_fuzz_dep_.CoverTab[194643]++
												return result
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/utils.go:37
	// _ = "end of CoverTab[194643]"
}

// ToDo - assuming that UeID is represented with an integer. Change it once treatment of UeID in SD-RAN is defined
func GetUeID(ueID *e2sm_v2_ies.Ueid) (int64, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/utils.go:41
	_go_fuzz_dep_.CoverTab[194647]++

												switch ue := ueID.Ueid.(type) {
	case *e2sm_v2_ies.Ueid_GNbUeid:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/utils.go:44
		_go_fuzz_dep_.CoverTab[194648]++
													return ue.GNbUeid.GetAmfUeNgapId().GetValue(), nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/utils.go:45
		// _ = "end of CoverTab[194648]"
	case *e2sm_v2_ies.Ueid_ENbUeid:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/utils.go:46
		_go_fuzz_dep_.CoverTab[194649]++
													return ue.ENbUeid.GetMMeUeS1ApId().GetValue(), nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/utils.go:47
		// _ = "end of CoverTab[194649]"
	case *e2sm_v2_ies.Ueid_EnGNbUeid:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/utils.go:48
		_go_fuzz_dep_.CoverTab[194650]++
													return int64(ue.EnGNbUeid.GetMENbUeX2ApId().GetValue()), nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/utils.go:49
		// _ = "end of CoverTab[194650]"
	case *e2sm_v2_ies.Ueid_NgENbUeid:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/utils.go:50
		_go_fuzz_dep_.CoverTab[194651]++
													return ue.NgENbUeid.GetAmfUeNgapId().GetValue(), nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/utils.go:51
		// _ = "end of CoverTab[194651]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/utils.go:52
		_go_fuzz_dep_.CoverTab[194652]++
													return -1, fmt.Errorf("GetUeID() couldn't extract UeID - obtained unexpected type %v", ue)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/utils.go:53
		// _ = "end of CoverTab[194652]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/utils.go:54
	// _ = "end of CoverTab[194647]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/utils.go:55
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/mho/utils.go:55
var _ = _go_fuzz_dep_.CoverTab
