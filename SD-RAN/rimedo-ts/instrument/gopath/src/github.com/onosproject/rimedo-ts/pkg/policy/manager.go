//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:1
// SPDX-FileCopyrightText: 2019-present Open Networking Foundation <info@opennetworking.org>
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:1
// SPDX-FileCopyrightText: 2019-present Rimedo Labs
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:1
//
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:1
// SPDX-License-Identifier: Apache-2.0
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:1
// Created by RIMEDO-Labs team
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:6
package policy

//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:6
import (
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:6
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:6
)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:6
import (
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:6
	_atomic_ "sync/atomic"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:6
)

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"os"

	policyAPI "github.com/onosproject/onos-a1-dm/go/policy_schemas/traffic_steering_preference/v2"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/rimedo-ts/pkg/mho"
	"github.com/xeipuuv/gojsonschema"
)

var log = logging.GetLogger("rimedo-ts", "policy")

func NewPolicySchemaValidatorV2(path string) *PolicySchemaValidatorV2 {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:24
	_go_fuzz_dep_.CoverTab[196125]++

									return &PolicySchemaValidatorV2{
		schemePath: path,
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:28
	// _ = "end of CoverTab[196125]"

}

type PolicySchemaValidatorV2 struct {
	schemePath string
}

func NewPolicyManager(policyMap *map[string]*mho.PolicyData) *PolicyManager {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:36
	_go_fuzz_dep_.CoverTab[196126]++

									var POLICY_WEIGHTS = map[string]int{
		"DEFAULT":	0.0,
		"PREFER":	16.0,
		"AVOID":	-16.0,
		"SHALL":	1000.0,
		"FORBID":	-1000.0,
	}

	return &PolicyManager{
		validator:	NewPolicySchemaValidatorV2("schemePath"),
		policyMap:	policyMap,
		preferenceMap:	POLICY_WEIGHTS,
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:50
	// _ = "end of CoverTab[196126]"

}

type PolicyManager struct {
	validator	*PolicySchemaValidatorV2
	policyMap	*map[string]*mho.PolicyData
	preferenceMap	map[string]int
}

func (m *PolicyManager) ReadPolicyObjectFromFileV2(jsonPath string, policyObject *mho.PolicyData) error {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:60
	_go_fuzz_dep_.CoverTab[196127]++

									jsonFile, err := m.LoadPolicyJsonFromFileV2(jsonPath)
									if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:63
		_go_fuzz_dep_.CoverTab[196132]++
										log.Error("Couldn't read PolicyObject from file")
										return err
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:65
		// _ = "end of CoverTab[196132]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:66
		_go_fuzz_dep_.CoverTab[196133]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:66
		// _ = "end of CoverTab[196133]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:66
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:66
	// _ = "end of CoverTab[196127]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:66
	_go_fuzz_dep_.CoverTab[196128]++

									var ok bool
									ok, err = m.ValidatePolicyJsonSchemaV2(jsonPath)
									if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:70
		_go_fuzz_dep_.CoverTab[196134]++
										log.Error("Error validating json scheme")
										return err
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:72
		// _ = "end of CoverTab[196134]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:73
		_go_fuzz_dep_.CoverTab[196135]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:73
		// _ = "end of CoverTab[196135]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:73
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:73
	// _ = "end of CoverTab[196128]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:73
	_go_fuzz_dep_.CoverTab[196129]++
									if !ok {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:74
		_go_fuzz_dep_.CoverTab[196136]++
										return errors.New("the json file is invalid")
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:75
		// _ = "end of CoverTab[196136]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:76
		_go_fuzz_dep_.CoverTab[196137]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:76
		// _ = "end of CoverTab[196137]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:76
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:76
	// _ = "end of CoverTab[196129]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:76
	_go_fuzz_dep_.CoverTab[196130]++
									if err = m.UnmarshalPolicyJsonV2(jsonFile, policyObject); err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:77
		_go_fuzz_dep_.CoverTab[196138]++
										log.Error("Error unmarshaling json file")
										return err
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:79
		// _ = "end of CoverTab[196138]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:80
		_go_fuzz_dep_.CoverTab[196139]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:80
		// _ = "end of CoverTab[196139]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:80
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:80
	// _ = "end of CoverTab[196130]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:80
	_go_fuzz_dep_.CoverTab[196131]++

									return nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:82
	// _ = "end of CoverTab[196131]"
}

func (m *PolicyManager) CheckPerUePolicyV2(ueScope policyAPI.Scope, policyObject *mho.PolicyData) bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:85
	_go_fuzz_dep_.CoverTab[196140]++

									if policyObject.API.Scope.UeID == nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:87
		_go_fuzz_dep_.CoverTab[196147]++
										return false
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:88
		// _ = "end of CoverTab[196147]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:89
		_go_fuzz_dep_.CoverTab[196148]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:89
		// _ = "end of CoverTab[196148]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:89
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:89
	// _ = "end of CoverTab[196140]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:89
	_go_fuzz_dep_.CoverTab[196141]++

									if *policyObject.API.Scope.UeID == "" {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:91
		_go_fuzz_dep_.CoverTab[196149]++
										return false
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:92
		// _ = "end of CoverTab[196149]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:93
		_go_fuzz_dep_.CoverTab[196150]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:93
		// _ = "end of CoverTab[196150]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:93
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:93
	// _ = "end of CoverTab[196141]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:93
	_go_fuzz_dep_.CoverTab[196142]++

									if *policyObject.API.Scope.UeID != *ueScope.UeID {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:95
		_go_fuzz_dep_.CoverTab[196151]++
										return false
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:96
		// _ = "end of CoverTab[196151]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:97
		_go_fuzz_dep_.CoverTab[196152]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:97
		// _ = "end of CoverTab[196152]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:97
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:97
	// _ = "end of CoverTab[196142]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:97
	_go_fuzz_dep_.CoverTab[196143]++

									if (policyObject.API.Scope.SliceID != nil) && func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:99
		_go_fuzz_dep_.CoverTab[196153]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:99
		return (((policyObject.API.Scope.SliceID.SD == nil || func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:99
			_go_fuzz_dep_.CoverTab[196154]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:99
			return (policyObject.API.Scope.SliceID.SD != nil && func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:99
				_go_fuzz_dep_.CoverTab[196155]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:99
				return *policyObject.API.Scope.SliceID.SD == ""
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:99
				// _ = "end of CoverTab[196155]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:99
			}())
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:99
			// _ = "end of CoverTab[196154]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:99
		}()) || func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:99
			_go_fuzz_dep_.CoverTab[196156]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:99
			return policyObject.API.Scope.SliceID.Sst <= 0
											// _ = "end of CoverTab[196156]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:100
		}() || func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:100
			_go_fuzz_dep_.CoverTab[196157]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:100
			return policyObject.API.Scope.SliceID.PlmnID.Mcc == ""
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:100
			// _ = "end of CoverTab[196157]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:100
		}() || func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:100
			_go_fuzz_dep_.CoverTab[196158]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:100
			return policyObject.API.Scope.SliceID.PlmnID.Mnc == ""
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:100
			// _ = "end of CoverTab[196158]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:100
		}()) || func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:100
			_go_fuzz_dep_.CoverTab[196159]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:100
			return (*policyObject.API.Scope.SliceID.SD != *ueScope.SliceID.SD || func() bool {
													_go_fuzz_dep_.CoverTab[196160]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:101
				return policyObject.API.Scope.SliceID.Sst != ueScope.SliceID.Sst
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:101
				// _ = "end of CoverTab[196160]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:101
			}() || func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:101
				_go_fuzz_dep_.CoverTab[196161]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:101
				return policyObject.API.Scope.SliceID.PlmnID.Mcc != ueScope.SliceID.PlmnID.Mcc
													// _ = "end of CoverTab[196161]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:102
			}() || func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:102
				_go_fuzz_dep_.CoverTab[196162]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:102
				return policyObject.API.Scope.SliceID.PlmnID.Mnc != ueScope.SliceID.PlmnID.Mnc
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:102
				// _ = "end of CoverTab[196162]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:102
			}())
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:102
			// _ = "end of CoverTab[196159]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:102
		}())
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:102
		// _ = "end of CoverTab[196153]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:102
	}() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:102
		_go_fuzz_dep_.CoverTab[196163]++
											return false
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:103
		// _ = "end of CoverTab[196163]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:104
		_go_fuzz_dep_.CoverTab[196164]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:104
		// _ = "end of CoverTab[196164]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:104
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:104
	// _ = "end of CoverTab[196143]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:104
	_go_fuzz_dep_.CoverTab[196144]++

										if (policyObject.API.Scope.QosID != nil) && func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:106
		_go_fuzz_dep_.CoverTab[196165]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:106
		return ((policyObject.API.Scope.QosID.QcI == nil && func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:106
			_go_fuzz_dep_.CoverTab[196166]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:106
			return policyObject.API.Scope.QosID.The5QI == nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:106
			// _ = "end of CoverTab[196166]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:106
		}()) || func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:106
			_go_fuzz_dep_.CoverTab[196167]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:106
			return (policyObject.API.Scope.QosID.QcI != nil && func() bool {
													_go_fuzz_dep_.CoverTab[196168]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:107
				return *policyObject.API.Scope.QosID.QcI != *ueScope.QosID.QcI
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:107
				// _ = "end of CoverTab[196168]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:107
			}())
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:107
			// _ = "end of CoverTab[196167]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:107
		}() || func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:107
			_go_fuzz_dep_.CoverTab[196169]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:107
			return (policyObject.API.Scope.QosID.The5QI != nil && func() bool {
													_go_fuzz_dep_.CoverTab[196170]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:108
				return *policyObject.API.Scope.QosID.The5QI != *ueScope.QosID.The5QI
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:108
				// _ = "end of CoverTab[196170]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:108
			}())
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:108
			// _ = "end of CoverTab[196169]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:108
		}())
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:108
		// _ = "end of CoverTab[196165]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:108
	}() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:108
		_go_fuzz_dep_.CoverTab[196171]++
											return false
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:109
		// _ = "end of CoverTab[196171]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:110
		_go_fuzz_dep_.CoverTab[196172]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:110
		// _ = "end of CoverTab[196172]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:110
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:110
	// _ = "end of CoverTab[196144]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:110
	_go_fuzz_dep_.CoverTab[196145]++

										if (policyObject.API.Scope.CellID != nil) && func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:112
		_go_fuzz_dep_.CoverTab[196173]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:112
		return (((policyObject.API.Scope.CellID.CID.NcI == nil && func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:112
			_go_fuzz_dep_.CoverTab[196174]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:112
			return policyObject.API.Scope.CellID.CID.EcI == nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:112
			// _ = "end of CoverTab[196174]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:112
		}()) || func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:112
			_go_fuzz_dep_.CoverTab[196175]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:112
			return (policyObject.API.Scope.CellID.CID.NcI != nil && func() bool {
													_go_fuzz_dep_.CoverTab[196176]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:113
				return *policyObject.API.Scope.CellID.CID.NcI != *ueScope.CellID.CID.NcI
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:113
				// _ = "end of CoverTab[196176]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:113
			}())
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:113
			// _ = "end of CoverTab[196175]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:113
		}() || func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:113
			_go_fuzz_dep_.CoverTab[196177]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:113
			return (policyObject.API.Scope.CellID.CID.EcI != nil && func() bool {
													_go_fuzz_dep_.CoverTab[196178]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:114
				return *policyObject.API.Scope.CellID.CID.EcI != *ueScope.CellID.CID.EcI
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:114
				// _ = "end of CoverTab[196178]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:114
			}())
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:114
			// _ = "end of CoverTab[196177]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:114
		}()) || func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:114
			_go_fuzz_dep_.CoverTab[196179]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:114
			return ((policyObject.API.Scope.CellID.PlmnID.Mcc == "" || func() bool {
													_go_fuzz_dep_.CoverTab[196180]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:115
				return policyObject.API.Scope.CellID.PlmnID.Mnc == ""
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:115
				// _ = "end of CoverTab[196180]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:115
			}()) || func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:115
				_go_fuzz_dep_.CoverTab[196181]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:115
				return (policyObject.API.Scope.CellID.PlmnID.Mcc != ueScope.CellID.PlmnID.Mcc || func() bool {
														_go_fuzz_dep_.CoverTab[196182]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:116
					return policyObject.API.Scope.CellID.PlmnID.Mnc != ueScope.CellID.PlmnID.Mnc
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:116
					// _ = "end of CoverTab[196182]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:116
				}())
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:116
				// _ = "end of CoverTab[196181]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:116
			}())
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:116
			// _ = "end of CoverTab[196179]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:116
		}())
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:116
		// _ = "end of CoverTab[196173]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:116
	}() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:116
		_go_fuzz_dep_.CoverTab[196183]++
											return false
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:117
		// _ = "end of CoverTab[196183]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:118
		_go_fuzz_dep_.CoverTab[196184]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:118
		// _ = "end of CoverTab[196184]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:118
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:118
	// _ = "end of CoverTab[196145]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:118
	_go_fuzz_dep_.CoverTab[196146]++

										return true
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:120
	// _ = "end of CoverTab[196146]"
}

func (m *PolicyManager) CheckPerSlicePolicyV2(ueScope policyAPI.Scope, policyObject *mho.PolicyData) bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:123
	_go_fuzz_dep_.CoverTab[196185]++

										if policyObject.API.Scope.SliceID == nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:125
		_go_fuzz_dep_.CoverTab[196191]++
											return false
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:126
		// _ = "end of CoverTab[196191]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:127
		_go_fuzz_dep_.CoverTab[196192]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:127
		// _ = "end of CoverTab[196192]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:127
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:127
	// _ = "end of CoverTab[196185]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:127
	_go_fuzz_dep_.CoverTab[196186]++

										if (policyObject.API.Scope.SliceID != nil && func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:129
		_go_fuzz_dep_.CoverTab[196193]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:129
		return *policyObject.API.Scope.SliceID.SD == ""
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:129
		// _ = "end of CoverTab[196193]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:129
	}()) || func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:129
		_go_fuzz_dep_.CoverTab[196194]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:129
		return policyObject.API.Scope.SliceID.Sst <= 0
											// _ = "end of CoverTab[196194]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:130
	}() || func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:130
		_go_fuzz_dep_.CoverTab[196195]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:130
		return policyObject.API.Scope.SliceID.PlmnID.Mcc == ""
											// _ = "end of CoverTab[196195]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:131
	}() || func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:131
		_go_fuzz_dep_.CoverTab[196196]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:131
		return policyObject.API.Scope.SliceID.PlmnID.Mnc == ""
											// _ = "end of CoverTab[196196]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:132
	}() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:132
		_go_fuzz_dep_.CoverTab[196197]++
											return false
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:133
		// _ = "end of CoverTab[196197]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:134
		_go_fuzz_dep_.CoverTab[196198]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:134
		// _ = "end of CoverTab[196198]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:134
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:134
	// _ = "end of CoverTab[196186]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:134
	_go_fuzz_dep_.CoverTab[196187]++

										if (policyObject.API.Scope.UeID != nil) && func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:136
		_go_fuzz_dep_.CoverTab[196199]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:136
		return !((*policyObject.API.Scope.UeID == "") || func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:136
			_go_fuzz_dep_.CoverTab[196200]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:136
			return (*policyObject.API.Scope.UeID == *ueScope.UeID)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:136
			// _ = "end of CoverTab[196200]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:136
		}())
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:136
		// _ = "end of CoverTab[196199]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:136
	}() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:136
		_go_fuzz_dep_.CoverTab[196201]++
											return false
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:137
		// _ = "end of CoverTab[196201]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:138
		_go_fuzz_dep_.CoverTab[196202]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:138
		// _ = "end of CoverTab[196202]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:138
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:138
	// _ = "end of CoverTab[196187]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:138
	_go_fuzz_dep_.CoverTab[196188]++

										if (policyObject.API.Scope.QosID != nil) && func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:140
		_go_fuzz_dep_.CoverTab[196203]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:140
		return ((policyObject.API.Scope.QosID.QcI == nil && func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:140
			_go_fuzz_dep_.CoverTab[196204]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:140
			return policyObject.API.Scope.QosID.The5QI == nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:140
			// _ = "end of CoverTab[196204]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:140
		}()) || func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:140
			_go_fuzz_dep_.CoverTab[196205]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:140
			return (policyObject.API.Scope.QosID.QcI != nil && func() bool {
													_go_fuzz_dep_.CoverTab[196206]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:141
				return *policyObject.API.Scope.QosID.QcI != *ueScope.QosID.QcI
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:141
				// _ = "end of CoverTab[196206]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:141
			}())
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:141
			// _ = "end of CoverTab[196205]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:141
		}() || func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:141
			_go_fuzz_dep_.CoverTab[196207]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:141
			return (policyObject.API.Scope.QosID.The5QI != nil && func() bool {
													_go_fuzz_dep_.CoverTab[196208]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:142
				return *policyObject.API.Scope.QosID.The5QI != *ueScope.QosID.The5QI
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:142
				// _ = "end of CoverTab[196208]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:142
			}())
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:142
			// _ = "end of CoverTab[196207]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:142
		}())
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:142
		// _ = "end of CoverTab[196203]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:142
	}() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:142
		_go_fuzz_dep_.CoverTab[196209]++
											return false
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:143
		// _ = "end of CoverTab[196209]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:144
		_go_fuzz_dep_.CoverTab[196210]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:144
		// _ = "end of CoverTab[196210]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:144
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:144
	// _ = "end of CoverTab[196188]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:144
	_go_fuzz_dep_.CoverTab[196189]++

										if (policyObject.API.Scope.CellID != nil) && func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:146
		_go_fuzz_dep_.CoverTab[196211]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:146
		return (((policyObject.API.Scope.CellID.CID.NcI == nil && func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:146
			_go_fuzz_dep_.CoverTab[196212]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:146
			return policyObject.API.Scope.CellID.CID.EcI == nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:146
			// _ = "end of CoverTab[196212]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:146
		}()) || func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:146
			_go_fuzz_dep_.CoverTab[196213]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:146
			return (policyObject.API.Scope.CellID.CID.NcI != nil && func() bool {
													_go_fuzz_dep_.CoverTab[196214]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:147
				return *policyObject.API.Scope.CellID.CID.NcI != *ueScope.CellID.CID.NcI
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:147
				// _ = "end of CoverTab[196214]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:147
			}())
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:147
			// _ = "end of CoverTab[196213]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:147
		}() || func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:147
			_go_fuzz_dep_.CoverTab[196215]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:147
			return (policyObject.API.Scope.CellID.CID.EcI != nil && func() bool {
													_go_fuzz_dep_.CoverTab[196216]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:148
				return *policyObject.API.Scope.CellID.CID.EcI != *ueScope.CellID.CID.EcI
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:148
				// _ = "end of CoverTab[196216]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:148
			}())
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:148
			// _ = "end of CoverTab[196215]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:148
		}()) || func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:148
			_go_fuzz_dep_.CoverTab[196217]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:148
			return ((policyObject.API.Scope.CellID.PlmnID.Mcc == "" || func() bool {
													_go_fuzz_dep_.CoverTab[196218]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:149
				return policyObject.API.Scope.CellID.PlmnID.Mnc == ""
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:149
				// _ = "end of CoverTab[196218]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:149
			}()) || func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:149
				_go_fuzz_dep_.CoverTab[196219]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:149
				return (policyObject.API.Scope.CellID.PlmnID.Mcc != ueScope.CellID.PlmnID.Mcc || func() bool {
														_go_fuzz_dep_.CoverTab[196220]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:150
					return policyObject.API.Scope.CellID.PlmnID.Mnc != ueScope.CellID.PlmnID.Mnc
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:150
					// _ = "end of CoverTab[196220]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:150
				}())
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:150
				// _ = "end of CoverTab[196219]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:150
			}())
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:150
			// _ = "end of CoverTab[196217]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:150
		}())
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:150
		// _ = "end of CoverTab[196211]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:150
	}() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:150
		_go_fuzz_dep_.CoverTab[196221]++
											return false
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:151
		// _ = "end of CoverTab[196221]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:152
		_go_fuzz_dep_.CoverTab[196222]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:152
		// _ = "end of CoverTab[196222]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:152
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:152
	// _ = "end of CoverTab[196189]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:152
	_go_fuzz_dep_.CoverTab[196190]++

										return true
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:154
	// _ = "end of CoverTab[196190]"
}

func (m *PolicyManager) GetTsResultForUEV2(ueScope policyAPI.Scope, rsrps []int, cellIds []policyAPI.CellID) policyAPI.CellID {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:157
	_go_fuzz_dep_.CoverTab[196223]++

										var bestCell policyAPI.CellID
										bestScore := -math.MaxFloat64
										for i := 0; i < len(rsrps); i++ {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:161
		_go_fuzz_dep_.CoverTab[196225]++
											preferece := m.GetPreferenceV2(ueScope, cellIds[i])
											score := m.GetPreferenceScoresV2(preferece, rsrps[i])
											if score > bestScore {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:164
			_go_fuzz_dep_.CoverTab[196226]++
												bestCell = cellIds[i]
												bestScore = score
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:166
			// _ = "end of CoverTab[196226]"
		} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:167
			_go_fuzz_dep_.CoverTab[196227]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:167
			// _ = "end of CoverTab[196227]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:167
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:167
		// _ = "end of CoverTab[196225]"
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:168
	// _ = "end of CoverTab[196223]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:168
	_go_fuzz_dep_.CoverTab[196224]++
										return bestCell
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:169
	// _ = "end of CoverTab[196224]"
}

func (m *PolicyManager) GetPreferenceScoresV2(preference string, rsrp int) float64 {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:172
	_go_fuzz_dep_.CoverTab[196228]++
										return float64(rsrp) + float64(m.preferenceMap[preference])
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:173
	// _ = "end of CoverTab[196228]"
}

func (m *PolicyManager) GetPreferenceV2(ueScope policyAPI.Scope, queryCellId policyAPI.CellID) string {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:176
	_go_fuzz_dep_.CoverTab[196229]++

										var preference string = "DEFAULT"
										for _, policy := range *m.policyMap {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:179
		_go_fuzz_dep_.CoverTab[196231]++
											if policy.IsEnforced {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:180
			_go_fuzz_dep_.CoverTab[196232]++
												if m.CheckPerSlicePolicyV2(ueScope, policy) || func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:181
				_go_fuzz_dep_.CoverTab[196233]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:181
				return m.CheckPerUePolicyV2(ueScope, policy)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:181
				// _ = "end of CoverTab[196233]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:181
			}() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:181
				_go_fuzz_dep_.CoverTab[196234]++
													for _, tspResource := range policy.API.TSPResources {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:182
					_go_fuzz_dep_.CoverTab[196235]++

														for _, cellId := range tspResource.CellIDList {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:184
						_go_fuzz_dep_.CoverTab[196236]++
															if ((cellId.CID.NcI != nil && func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:185
							_go_fuzz_dep_.CoverTab[196237]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:185
							return queryCellId.CID.NcI != nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:185
							// _ = "end of CoverTab[196237]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:185
						}() && func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:185
							_go_fuzz_dep_.CoverTab[196238]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:185
							return *cellId.CID.NcI == *queryCellId.CID.NcI
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:185
							// _ = "end of CoverTab[196238]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:185
						}()) || func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:185
							_go_fuzz_dep_.CoverTab[196239]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:185
							return (cellId.CID.EcI != nil && func() bool {
																	_go_fuzz_dep_.CoverTab[196240]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:186
								return queryCellId.CID.EcI != nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:186
								// _ = "end of CoverTab[196240]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:186
							}() && func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:186
								_go_fuzz_dep_.CoverTab[196241]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:186
								return *cellId.CID.EcI == *queryCellId.CID.EcI
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:186
								// _ = "end of CoverTab[196241]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:186
							}())
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:186
							// _ = "end of CoverTab[196239]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:186
						}()) && func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:186
							_go_fuzz_dep_.CoverTab[196242]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:186
							return (cellId.PlmnID.Mcc == queryCellId.PlmnID.Mcc && func() bool {
																	_go_fuzz_dep_.CoverTab[196243]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:187
								return cellId.PlmnID.Mnc == queryCellId.PlmnID.Mnc
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:187
								// _ = "end of CoverTab[196243]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:187
							}())
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:187
							// _ = "end of CoverTab[196242]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:187
						}() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:187
							_go_fuzz_dep_.CoverTab[196244]++
																preference = string(tspResource.Preference)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:188
							// _ = "end of CoverTab[196244]"
						} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:189
							_go_fuzz_dep_.CoverTab[196245]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:189
							// _ = "end of CoverTab[196245]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:189
						}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:189
						// _ = "end of CoverTab[196236]"
					}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:190
					// _ = "end of CoverTab[196235]"
				}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:191
				// _ = "end of CoverTab[196234]"
			} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:192
				_go_fuzz_dep_.CoverTab[196246]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:192
				// _ = "end of CoverTab[196246]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:192
			}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:192
			// _ = "end of CoverTab[196232]"
		} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:193
			_go_fuzz_dep_.CoverTab[196247]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:193
			// _ = "end of CoverTab[196247]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:193
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:193
		// _ = "end of CoverTab[196231]"
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:194
	// _ = "end of CoverTab[196229]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:194
	_go_fuzz_dep_.CoverTab[196230]++
										return preference
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:195
	// _ = "end of CoverTab[196230]"
}

func (m *PolicyManager) AddPolicyV2(policyId string, policyDir string, policyObject *mho.PolicyData) (*mho.PolicyData, error) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:198
	_go_fuzz_dep_.CoverTab[196248]++

										policyPath := policyDir + policyId
										err := m.ReadPolicyObjectFromFileV2(policyPath, policyObject)
										if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:202
		_go_fuzz_dep_.CoverTab[196250]++
											log.Error(fmt.Sprintf("Couldn't read PolicyObject from file \n policyId: %s from: %s", policyId, policyPath))
											return nil, err
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:204
		// _ = "end of CoverTab[196250]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:205
		_go_fuzz_dep_.CoverTab[196251]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:205
		// _ = "end of CoverTab[196251]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:205
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:205
	// _ = "end of CoverTab[196248]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:205
	_go_fuzz_dep_.CoverTab[196249]++
										return policyObject, nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:206
	// _ = "end of CoverTab[196249]"
}

func (m *PolicyManager) EnforcePolicyV2(policyId string) bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:209
	_go_fuzz_dep_.CoverTab[196252]++

										if _, ok := (*m.policyMap)[policyId]; ok {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:211
		_go_fuzz_dep_.CoverTab[196254]++
											(*m.policyMap)[policyId].IsEnforced = true
											return true
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:213
		// _ = "end of CoverTab[196254]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:214
		_go_fuzz_dep_.CoverTab[196255]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:214
		// _ = "end of CoverTab[196255]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:214
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:214
	// _ = "end of CoverTab[196252]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:214
	_go_fuzz_dep_.CoverTab[196253]++
										log.Error(fmt.Sprintf("Policy with policyId: %s, not enforced", policyId))
										return false
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:216
	// _ = "end of CoverTab[196253]"
}

func (m *PolicyManager) DisablePolicyV2(policyId string) bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:219
	_go_fuzz_dep_.CoverTab[196256]++

										if _, ok := (*m.policyMap)[policyId]; ok {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:221
		_go_fuzz_dep_.CoverTab[196258]++
											(*m.policyMap)[policyId].IsEnforced = false
											return true
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:223
		// _ = "end of CoverTab[196258]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:224
		_go_fuzz_dep_.CoverTab[196259]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:224
		// _ = "end of CoverTab[196259]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:224
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:224
	// _ = "end of CoverTab[196256]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:224
	_go_fuzz_dep_.CoverTab[196257]++
										log.Error(fmt.Sprintf("Policy with policyId: %s, not enforced", policyId))
										return false
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:226
	// _ = "end of CoverTab[196257]"
}

func (m *PolicyManager) GetPolicyV2(policyId string) (*mho.PolicyData, bool) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:229
	_go_fuzz_dep_.CoverTab[196260]++

										if val, ok := (*m.policyMap)[policyId]; ok {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:231
		_go_fuzz_dep_.CoverTab[196262]++
											return val, ok
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:232
		// _ = "end of CoverTab[196262]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:233
		_go_fuzz_dep_.CoverTab[196263]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:233
		// _ = "end of CoverTab[196263]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:233
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:233
	// _ = "end of CoverTab[196260]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:233
	_go_fuzz_dep_.CoverTab[196261]++
										log.Error(fmt.Sprintf("Policy with policyId: %s, not enforced", policyId))
										return nil, false
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:235
	// _ = "end of CoverTab[196261]"
}

func (m *PolicyManager) ValidatePolicyJsonSchemaV2(jsonPath string) (bool, error) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:238
	_go_fuzz_dep_.CoverTab[196264]++

										schemaLoader := gojsonschema.NewReferenceLoader("file://" + m.validator.schemePath)
										documentLoader := gojsonschema.NewReferenceLoader("file://" + jsonPath)

										result, err := gojsonschema.Validate(schemaLoader, documentLoader)
										if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:244
		_go_fuzz_dep_.CoverTab[196266]++
											return false, err
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:245
		// _ = "end of CoverTab[196266]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:246
		_go_fuzz_dep_.CoverTab[196267]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:246
		// _ = "end of CoverTab[196267]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:246
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:246
	// _ = "end of CoverTab[196264]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:246
	_go_fuzz_dep_.CoverTab[196265]++
										return result.Valid(), nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:247
	// _ = "end of CoverTab[196265]"
}

func (m *PolicyManager) UnmarshalPolicyJsonV2(jsonFile []byte, policyObject *mho.PolicyData) error {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:250
	_go_fuzz_dep_.CoverTab[196268]++

										if err := json.Unmarshal(jsonFile, policyObject.API); err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:252
		_go_fuzz_dep_.CoverTab[196270]++
											log.Error("Couldn't read PolicyObject from file")
											return err
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:254
		// _ = "end of CoverTab[196270]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:255
		_go_fuzz_dep_.CoverTab[196271]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:255
		// _ = "end of CoverTab[196271]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:255
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:255
	// _ = "end of CoverTab[196268]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:255
	_go_fuzz_dep_.CoverTab[196269]++
										policyObject.IsEnforced = false
										return nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:257
	// _ = "end of CoverTab[196269]"

}

func (m *PolicyManager) LoadPolicyJsonFromFileV2(path string) ([]byte, error) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:261
	_go_fuzz_dep_.CoverTab[196272]++

										jsonFile, err := os.Open(path)
										if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:264
		_go_fuzz_dep_.CoverTab[196275]++
											log.Error("Failed to open policy JSON File")
											return nil, err
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:266
		// _ = "end of CoverTab[196275]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:267
		_go_fuzz_dep_.CoverTab[196276]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:267
		// _ = "end of CoverTab[196276]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:267
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:267
	// _ = "end of CoverTab[196272]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:267
	_go_fuzz_dep_.CoverTab[196273]++

										defer jsonFile.Close()

										byteValue, err := ioutil.ReadAll(jsonFile)
										if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:272
		_go_fuzz_dep_.CoverTab[196277]++
											log.Error("Failed to read data from policy JSON File")
											return nil, err
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:274
		// _ = "end of CoverTab[196277]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:275
		_go_fuzz_dep_.CoverTab[196278]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:275
		// _ = "end of CoverTab[196278]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:275
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:275
	// _ = "end of CoverTab[196273]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:275
	_go_fuzz_dep_.CoverTab[196274]++
										return byteValue, nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:276
	// _ = "end of CoverTab[196274]"

}

func (m *PolicyManager) isSimilarEnforced(policyData *mho.PolicyData) bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:280
	_go_fuzz_dep_.CoverTab[196279]++
										for _, policy := range *m.policyMap {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:281
		_go_fuzz_dep_.CoverTab[196281]++

											sameSlice := false
											sameUE := false
											sameQoS := false
											sameCellID := false

											if (policyData.API.Scope.SliceID == nil && func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:288
			_go_fuzz_dep_.CoverTab[196287]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:288
			return policy.API.Scope.SliceID == nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:288
			// _ = "end of CoverTab[196287]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:288
		}()) || func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:288
			_go_fuzz_dep_.CoverTab[196288]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:288
			return (policyData.API.Scope.SliceID != nil && func() bool {
													_go_fuzz_dep_.CoverTab[196289]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:289
				return policy.API.Scope.SliceID != nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:289
				// _ = "end of CoverTab[196289]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:289
			}() && func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:289
				_go_fuzz_dep_.CoverTab[196290]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:289
				return policy.API.Scope.SliceID.Sst == policyData.API.Scope.SliceID.Sst
													// _ = "end of CoverTab[196290]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:290
			}() && func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:290
				_go_fuzz_dep_.CoverTab[196291]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:290
				return policy.API.Scope.SliceID.SD != nil
													// _ = "end of CoverTab[196291]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:291
			}() && func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:291
				_go_fuzz_dep_.CoverTab[196292]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:291
				return policyData.API.Scope.SliceID.SD != nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:291
				// _ = "end of CoverTab[196292]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:291
			}() && func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:291
				_go_fuzz_dep_.CoverTab[196293]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:291
				return *policy.API.Scope.SliceID.SD == *policyData.API.Scope.SliceID.SD
													// _ = "end of CoverTab[196293]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:292
			}() && func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:292
				_go_fuzz_dep_.CoverTab[196294]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:292
				return policy.API.Scope.SliceID.PlmnID.Mcc == policyData.API.Scope.SliceID.PlmnID.Mcc
													// _ = "end of CoverTab[196294]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:293
			}() && func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:293
				_go_fuzz_dep_.CoverTab[196295]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:293
				return policy.API.Scope.SliceID.PlmnID.Mnc == policyData.API.Scope.SliceID.PlmnID.Mnc
													// _ = "end of CoverTab[196295]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:294
			}())
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:294
			// _ = "end of CoverTab[196288]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:294
		}() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:294
			_go_fuzz_dep_.CoverTab[196296]++
												sameSlice = true
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:295
			// _ = "end of CoverTab[196296]"
		} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:296
			_go_fuzz_dep_.CoverTab[196297]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:296
			// _ = "end of CoverTab[196297]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:296
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:296
		// _ = "end of CoverTab[196281]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:296
		_go_fuzz_dep_.CoverTab[196282]++

											if (policyData.API.Scope.UeID == nil && func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:298
			_go_fuzz_dep_.CoverTab[196298]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:298
			return policy.API.Scope.UeID == nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:298
			// _ = "end of CoverTab[196298]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:298
		}()) || func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:298
			_go_fuzz_dep_.CoverTab[196299]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:298
			return (policyData.API.Scope.UeID != nil && func() bool {
													_go_fuzz_dep_.CoverTab[196300]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:299
				return policy.API.Scope.UeID != nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:299
				// _ = "end of CoverTab[196300]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:299
			}() && func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:299
				_go_fuzz_dep_.CoverTab[196301]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:299
				return *policy.API.Scope.UeID == *policyData.API.Scope.UeID
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:299
				// _ = "end of CoverTab[196301]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:299
			}())
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:299
			// _ = "end of CoverTab[196299]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:299
		}() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:299
			_go_fuzz_dep_.CoverTab[196302]++
												sameUE = true
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:300
			// _ = "end of CoverTab[196302]"
		} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:301
			_go_fuzz_dep_.CoverTab[196303]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:301
			// _ = "end of CoverTab[196303]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:301
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:301
		// _ = "end of CoverTab[196282]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:301
		_go_fuzz_dep_.CoverTab[196283]++

											if (policyData.API.Scope.QosID == nil && func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:303
			_go_fuzz_dep_.CoverTab[196304]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:303
			return policy.API.Scope.QosID == nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:303
			// _ = "end of CoverTab[196304]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:303
		}()) || func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:303
			_go_fuzz_dep_.CoverTab[196305]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:303
			return (policyData.API.Scope.QosID != nil && func() bool {
													_go_fuzz_dep_.CoverTab[196306]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:304
				return policy.API.Scope.QosID != nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:304
				// _ = "end of CoverTab[196306]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:304
			}() && func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:304
				_go_fuzz_dep_.CoverTab[196307]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:304
				return ((policy.API.Scope.QosID.QcI != nil && func() bool {
														_go_fuzz_dep_.CoverTab[196308]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:305
					return policyData.API.Scope.QosID.QcI != nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:305
					// _ = "end of CoverTab[196308]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:305
				}() && func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:305
					_go_fuzz_dep_.CoverTab[196309]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:305
					return *policy.API.Scope.QosID.QcI == *policyData.API.Scope.QosID.QcI
														// _ = "end of CoverTab[196309]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:306
				}()) || func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:306
					_go_fuzz_dep_.CoverTab[196310]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:306
					return (policy.API.Scope.QosID.The5QI != nil && func() bool {
															_go_fuzz_dep_.CoverTab[196311]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:307
						return policyData.API.Scope.QosID.The5QI != nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:307
						// _ = "end of CoverTab[196311]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:307
					}() && func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:307
						_go_fuzz_dep_.CoverTab[196312]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:307
						return *policy.API.Scope.QosID.The5QI == *policyData.API.Scope.QosID.The5QI
															// _ = "end of CoverTab[196312]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:308
					}())
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:308
					// _ = "end of CoverTab[196310]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:308
				}())
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:308
				// _ = "end of CoverTab[196307]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:308
			}())
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:308
			// _ = "end of CoverTab[196305]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:308
		}() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:308
			_go_fuzz_dep_.CoverTab[196313]++
												sameQoS = true
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:309
			// _ = "end of CoverTab[196313]"
		} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:310
			_go_fuzz_dep_.CoverTab[196314]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:310
			// _ = "end of CoverTab[196314]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:310
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:310
		// _ = "end of CoverTab[196283]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:310
		_go_fuzz_dep_.CoverTab[196284]++

											if (policyData.API.Scope.CellID == nil && func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:312
			_go_fuzz_dep_.CoverTab[196315]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:312
			return policy.API.Scope.CellID == nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:312
			// _ = "end of CoverTab[196315]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:312
		}()) || func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:312
			_go_fuzz_dep_.CoverTab[196316]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:312
			return (policyData.API.Scope.CellID != nil && func() bool {
													_go_fuzz_dep_.CoverTab[196317]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:313
				return policy.API.Scope.CellID != nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:313
				// _ = "end of CoverTab[196317]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:313
			}() && func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:313
				_go_fuzz_dep_.CoverTab[196318]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:313
				return ((policy.API.Scope.CellID.CID.NcI != nil && func() bool {
														_go_fuzz_dep_.CoverTab[196319]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:314
					return policyData.API.Scope.CellID.CID.NcI != nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:314
					// _ = "end of CoverTab[196319]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:314
				}() && func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:314
					_go_fuzz_dep_.CoverTab[196320]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:314
					return *policy.API.Scope.CellID.CID.NcI == *policyData.API.Scope.CellID.CID.NcI
														// _ = "end of CoverTab[196320]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:315
				}()) || func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:315
					_go_fuzz_dep_.CoverTab[196321]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:315
					return (policy.API.Scope.CellID.CID.EcI != nil && func() bool {
															_go_fuzz_dep_.CoverTab[196322]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:316
						return policyData.API.Scope.CellID.CID.EcI != nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:316
						// _ = "end of CoverTab[196322]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:316
					}() && func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:316
						_go_fuzz_dep_.CoverTab[196323]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:316
						return *policy.API.Scope.CellID.CID.EcI == *policyData.API.Scope.CellID.CID.EcI
															// _ = "end of CoverTab[196323]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:317
					}())
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:317
					// _ = "end of CoverTab[196321]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:317
				}())
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:317
				// _ = "end of CoverTab[196318]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:317
			}() && func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:317
				_go_fuzz_dep_.CoverTab[196324]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:317
				return policy.API.Scope.CellID.PlmnID.Mcc == policyData.API.Scope.CellID.PlmnID.Mcc
													// _ = "end of CoverTab[196324]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:318
			}() && func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:318
				_go_fuzz_dep_.CoverTab[196325]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:318
				return policy.API.Scope.CellID.PlmnID.Mnc == policyData.API.Scope.CellID.PlmnID.Mnc
													// _ = "end of CoverTab[196325]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:319
			}())
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:319
			// _ = "end of CoverTab[196316]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:319
		}() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:319
			_go_fuzz_dep_.CoverTab[196326]++
												sameCellID = true
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:320
			// _ = "end of CoverTab[196326]"
		} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:321
			_go_fuzz_dep_.CoverTab[196327]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:321
			// _ = "end of CoverTab[196327]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:321
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:321
		// _ = "end of CoverTab[196284]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:321
		_go_fuzz_dep_.CoverTab[196285]++

											if sameSlice {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:323
			_go_fuzz_dep_.CoverTab[196328]++
												if sameUE {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:324
				_go_fuzz_dep_.CoverTab[196329]++
													if sameQoS {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:325
					_go_fuzz_dep_.CoverTab[196330]++
														if !sameCellID {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:326
						_go_fuzz_dep_.CoverTab[196331]++
															continue
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:327
						// _ = "end of CoverTab[196331]"
					} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:328
						_go_fuzz_dep_.CoverTab[196332]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:328
						// _ = "end of CoverTab[196332]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:328
					}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:328
					// _ = "end of CoverTab[196330]"
				} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:329
					_go_fuzz_dep_.CoverTab[196333]++
														continue
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:330
					// _ = "end of CoverTab[196333]"
				}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:331
				// _ = "end of CoverTab[196329]"
			} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:332
				_go_fuzz_dep_.CoverTab[196334]++
													if sameQoS {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:333
					_go_fuzz_dep_.CoverTab[196335]++
														if !sameCellID {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:334
						_go_fuzz_dep_.CoverTab[196336]++
															continue
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:335
						// _ = "end of CoverTab[196336]"
					} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:336
						_go_fuzz_dep_.CoverTab[196337]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:336
						// _ = "end of CoverTab[196337]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:336
					}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:336
					// _ = "end of CoverTab[196335]"
				} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:337
					_go_fuzz_dep_.CoverTab[196338]++
														continue
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:338
					// _ = "end of CoverTab[196338]"
				}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:339
				// _ = "end of CoverTab[196334]"
			}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:340
			// _ = "end of CoverTab[196328]"
		} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:341
			_go_fuzz_dep_.CoverTab[196339]++
												if sameUE {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:342
				_go_fuzz_dep_.CoverTab[196340]++
													if sameQoS {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:343
					_go_fuzz_dep_.CoverTab[196341]++
														if !sameCellID {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:344
						_go_fuzz_dep_.CoverTab[196342]++
															continue
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:345
						// _ = "end of CoverTab[196342]"
					} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:346
						_go_fuzz_dep_.CoverTab[196343]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:346
						// _ = "end of CoverTab[196343]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:346
					}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:346
					// _ = "end of CoverTab[196341]"
				} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:347
					_go_fuzz_dep_.CoverTab[196344]++
														continue
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:348
					// _ = "end of CoverTab[196344]"
				}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:349
				// _ = "end of CoverTab[196340]"
			} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:350
				_go_fuzz_dep_.CoverTab[196345]++
													continue
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:351
				// _ = "end of CoverTab[196345]"
			}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:352
			// _ = "end of CoverTab[196339]"
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:353
		// _ = "end of CoverTab[196285]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:353
		_go_fuzz_dep_.CoverTab[196286]++

											if policy.IsEnforced {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:355
			_go_fuzz_dep_.CoverTab[196346]++
												policy.IsEnforced = false
												return true
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:357
			// _ = "end of CoverTab[196346]"
		} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:358
			_go_fuzz_dep_.CoverTab[196347]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:358
			// _ = "end of CoverTab[196347]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:358
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:358
		// _ = "end of CoverTab[196286]"

	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:360
	// _ = "end of CoverTab[196279]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:360
	_go_fuzz_dep_.CoverTab[196280]++
										return false
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:361
	// _ = "end of CoverTab[196280]"
}

//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:362
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/policy/manager.go:362
var _ = _go_fuzz_dep_.CoverTab
