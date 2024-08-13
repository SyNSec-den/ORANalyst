// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:5
package utils

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:5
)

import (
	"fmt"
	"strconv"

	"github.com/onosproject/onos-lib-go/pkg/errors"

	gnmi "github.com/openconfig/gnmi/proto/gnmi"
)

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:18
// ToUint64 converts an interface value to uint64
func ToUint64(value interface{}) (uint64, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:19
	_go_fuzz_dep_.CoverTab[193725]++
															switch v := value.(type) {
	case *gnmi.TypedValue:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:21
		_go_fuzz_dep_.CoverTab[193726]++
																return toGnmiTypedValue(value).GetUintVal(), nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:22
		// _ = "end of CoverTab[193726]"

	case float64:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:24
		_go_fuzz_dep_.CoverTab[193727]++
																val, err := strconv.ParseUint(fmt.Sprintf("%v", value), 10, 64)
																if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:26
			_go_fuzz_dep_.CoverTab[193732]++
																	return 0, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:27
			// _ = "end of CoverTab[193732]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:28
			_go_fuzz_dep_.CoverTab[193733]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:28
			// _ = "end of CoverTab[193733]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:28
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:28
		// _ = "end of CoverTab[193727]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:28
		_go_fuzz_dep_.CoverTab[193728]++
																return val, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:29
		// _ = "end of CoverTab[193728]"
	case uint64:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:30
		_go_fuzz_dep_.CoverTab[193729]++
																val, err := strconv.ParseUint(fmt.Sprintf("%v", value), 10, 64)
																if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:32
			_go_fuzz_dep_.CoverTab[193734]++
																	return 0, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:33
			// _ = "end of CoverTab[193734]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:34
			_go_fuzz_dep_.CoverTab[193735]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:34
			// _ = "end of CoverTab[193735]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:34
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:34
		// _ = "end of CoverTab[193729]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:34
		_go_fuzz_dep_.CoverTab[193730]++
																return val, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:35
		// _ = "end of CoverTab[193730]"

	default:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:37
		_go_fuzz_dep_.CoverTab[193731]++
																return 0, errors.New(errors.NotSupported, "Not supported type %v", v)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:38
		// _ = "end of CoverTab[193731]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:39
	// _ = "end of CoverTab[193725]"

}

// ToFloat converts an interface value to float
func ToFloat(value interface{}) (float32, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:44
	_go_fuzz_dep_.CoverTab[193736]++
															switch v := value.(type) {
	case *gnmi.TypedValue:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:46
		_go_fuzz_dep_.CoverTab[193737]++
																return toGnmiTypedValue(value).GetFloatVal(), nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:47
		// _ = "end of CoverTab[193737]"

	case float32:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:49
		_go_fuzz_dep_.CoverTab[193738]++
																return v, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:50
		// _ = "end of CoverTab[193738]"

	default:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:52
		_go_fuzz_dep_.CoverTab[193739]++
																return 0, errors.New(errors.NotSupported, "Not supported type %v", v)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:53
		// _ = "end of CoverTab[193739]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:54
	// _ = "end of CoverTab[193736]"

}

// ToString converts value to string
func ToString(value interface{}) (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:59
	_go_fuzz_dep_.CoverTab[193740]++
															switch v := value.(type) {
	case *gnmi.TypedValue:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:61
		_go_fuzz_dep_.CoverTab[193741]++
																return toGnmiTypedValue(value).GetStringVal(), nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:62
		// _ = "end of CoverTab[193741]"

	case string:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:64
		_go_fuzz_dep_.CoverTab[193742]++
																return value.(string), nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:65
		// _ = "end of CoverTab[193742]"

	default:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:67
		_go_fuzz_dep_.CoverTab[193743]++
																return "", errors.New(errors.NotSupported, "Not supported type %v", v)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:68
		// _ = "end of CoverTab[193743]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:69
	// _ = "end of CoverTab[193740]"

}

// toGnmiTypedValue
func toGnmiTypedValue(value interface{}) *gnmi.TypedValue {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:74
	_go_fuzz_dep_.CoverTab[193744]++
															return value.(*gnmi.TypedValue)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:75
	// _ = "end of CoverTab[193744]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:76
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/typedvalue.go:76
var _ = _go_fuzz_dep_.CoverTab
