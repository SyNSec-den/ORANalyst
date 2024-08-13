// Copyright © 2014 Steve Francia <spf@spf13.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:6
package cast

//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:6
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:6
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:6
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:6
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:6
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:6
)

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var errNegativeNotAllowed = errors.New("unable to cast negative value")

// ToTimeE casts an interface to a time.Time type.
func ToTimeE(i interface{}) (tim time.Time, err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:22
	_go_fuzz_dep_.CoverTab[118554]++
										return ToTimeInDefaultLocationE(i, time.UTC)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:23
	// _ = "end of CoverTab[118554]"
}

// ToTimeInDefaultLocationE casts an empty interface to time.Time,
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:26
// interpreting inputs without a timezone to be in the given location,
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:26
// or the local timezone if nil.
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:29
func ToTimeInDefaultLocationE(i interface{}, location *time.Location) (tim time.Time, err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:29
	_go_fuzz_dep_.CoverTab[118555]++
										i = indirect(i)

										switch v := i.(type) {
	case time.Time:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:33
		_go_fuzz_dep_.CoverTab[118556]++
											return v, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:34
		// _ = "end of CoverTab[118556]"
	case string:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:35
		_go_fuzz_dep_.CoverTab[118557]++
											return StringToDateInDefaultLocation(v, location)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:36
		// _ = "end of CoverTab[118557]"
	case int:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:37
		_go_fuzz_dep_.CoverTab[118558]++
											return time.Unix(int64(v), 0), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:38
		// _ = "end of CoverTab[118558]"
	case int64:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:39
		_go_fuzz_dep_.CoverTab[118559]++
											return time.Unix(v, 0), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:40
		// _ = "end of CoverTab[118559]"
	case int32:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:41
		_go_fuzz_dep_.CoverTab[118560]++
											return time.Unix(int64(v), 0), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:42
		// _ = "end of CoverTab[118560]"
	case uint:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:43
		_go_fuzz_dep_.CoverTab[118561]++
											return time.Unix(int64(v), 0), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:44
		// _ = "end of CoverTab[118561]"
	case uint64:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:45
		_go_fuzz_dep_.CoverTab[118562]++
											return time.Unix(int64(v), 0), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:46
		// _ = "end of CoverTab[118562]"
	case uint32:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:47
		_go_fuzz_dep_.CoverTab[118563]++
											return time.Unix(int64(v), 0), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:48
		// _ = "end of CoverTab[118563]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:49
		_go_fuzz_dep_.CoverTab[118564]++
											return time.Time{}, fmt.Errorf("unable to cast %#v of type %T to Time", i, i)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:50
		// _ = "end of CoverTab[118564]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:51
	// _ = "end of CoverTab[118555]"
}

// ToDurationE casts an interface to a time.Duration type.
func ToDurationE(i interface{}) (d time.Duration, err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:55
	_go_fuzz_dep_.CoverTab[118565]++
										i = indirect(i)

										switch s := i.(type) {
	case time.Duration:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:59
		_go_fuzz_dep_.CoverTab[118566]++
											return s, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:60
		// _ = "end of CoverTab[118566]"
	case int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:61
		_go_fuzz_dep_.CoverTab[118567]++
											d = time.Duration(ToInt64(s))
											return
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:63
		// _ = "end of CoverTab[118567]"
	case float32, float64:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:64
		_go_fuzz_dep_.CoverTab[118568]++
											d = time.Duration(ToFloat64(s))
											return
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:66
		// _ = "end of CoverTab[118568]"
	case string:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:67
		_go_fuzz_dep_.CoverTab[118569]++
											if strings.ContainsAny(s, "nsuµmh") {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:68
			_go_fuzz_dep_.CoverTab[118572]++
												d, err = time.ParseDuration(s)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:69
			// _ = "end of CoverTab[118572]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:70
			_go_fuzz_dep_.CoverTab[118573]++
												d, err = time.ParseDuration(s + "ns")
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:71
			// _ = "end of CoverTab[118573]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:72
		// _ = "end of CoverTab[118569]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:72
		_go_fuzz_dep_.CoverTab[118570]++
											return
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:73
		// _ = "end of CoverTab[118570]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:74
		_go_fuzz_dep_.CoverTab[118571]++
											err = fmt.Errorf("unable to cast %#v of type %T to Duration", i, i)
											return
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:76
		// _ = "end of CoverTab[118571]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:77
	// _ = "end of CoverTab[118565]"
}

// ToBoolE casts an interface to a bool type.
func ToBoolE(i interface{}) (bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:81
	_go_fuzz_dep_.CoverTab[118574]++
										i = indirect(i)

										switch b := i.(type) {
	case bool:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:85
		_go_fuzz_dep_.CoverTab[118575]++
											return b, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:86
		// _ = "end of CoverTab[118575]"
	case nil:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:87
		_go_fuzz_dep_.CoverTab[118576]++
											return false, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:88
		// _ = "end of CoverTab[118576]"
	case int:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:89
		_go_fuzz_dep_.CoverTab[118577]++
											if i.(int) != 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:90
			_go_fuzz_dep_.CoverTab[118581]++
												return true, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:91
			// _ = "end of CoverTab[118581]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:92
			_go_fuzz_dep_.CoverTab[118582]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:92
			// _ = "end of CoverTab[118582]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:92
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:92
		// _ = "end of CoverTab[118577]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:92
		_go_fuzz_dep_.CoverTab[118578]++
											return false, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:93
		// _ = "end of CoverTab[118578]"
	case string:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:94
		_go_fuzz_dep_.CoverTab[118579]++
											return strconv.ParseBool(i.(string))
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:95
		// _ = "end of CoverTab[118579]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:96
		_go_fuzz_dep_.CoverTab[118580]++
											return false, fmt.Errorf("unable to cast %#v of type %T to bool", i, i)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:97
		// _ = "end of CoverTab[118580]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:98
	// _ = "end of CoverTab[118574]"
}

// ToFloat64E casts an interface to a float64 type.
func ToFloat64E(i interface{}) (float64, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:102
	_go_fuzz_dep_.CoverTab[118583]++
										i = indirect(i)

										switch s := i.(type) {
	case float64:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:106
		_go_fuzz_dep_.CoverTab[118584]++
											return s, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:107
		// _ = "end of CoverTab[118584]"
	case float32:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:108
		_go_fuzz_dep_.CoverTab[118585]++
											return float64(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:109
		// _ = "end of CoverTab[118585]"
	case int:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:110
		_go_fuzz_dep_.CoverTab[118586]++
											return float64(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:111
		// _ = "end of CoverTab[118586]"
	case int64:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:112
		_go_fuzz_dep_.CoverTab[118587]++
											return float64(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:113
		// _ = "end of CoverTab[118587]"
	case int32:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:114
		_go_fuzz_dep_.CoverTab[118588]++
											return float64(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:115
		// _ = "end of CoverTab[118588]"
	case int16:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:116
		_go_fuzz_dep_.CoverTab[118589]++
											return float64(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:117
		// _ = "end of CoverTab[118589]"
	case int8:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:118
		_go_fuzz_dep_.CoverTab[118590]++
											return float64(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:119
		// _ = "end of CoverTab[118590]"
	case uint:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:120
		_go_fuzz_dep_.CoverTab[118591]++
											return float64(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:121
		// _ = "end of CoverTab[118591]"
	case uint64:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:122
		_go_fuzz_dep_.CoverTab[118592]++
											return float64(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:123
		// _ = "end of CoverTab[118592]"
	case uint32:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:124
		_go_fuzz_dep_.CoverTab[118593]++
											return float64(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:125
		// _ = "end of CoverTab[118593]"
	case uint16:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:126
		_go_fuzz_dep_.CoverTab[118594]++
											return float64(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:127
		// _ = "end of CoverTab[118594]"
	case uint8:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:128
		_go_fuzz_dep_.CoverTab[118595]++
											return float64(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:129
		// _ = "end of CoverTab[118595]"
	case string:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:130
		_go_fuzz_dep_.CoverTab[118596]++
											v, err := strconv.ParseFloat(s, 64)
											if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:132
			_go_fuzz_dep_.CoverTab[118601]++
												return v, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:133
			// _ = "end of CoverTab[118601]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:134
			_go_fuzz_dep_.CoverTab[118602]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:134
			// _ = "end of CoverTab[118602]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:134
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:134
		// _ = "end of CoverTab[118596]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:134
		_go_fuzz_dep_.CoverTab[118597]++
											return 0, fmt.Errorf("unable to cast %#v of type %T to float64", i, i)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:135
		// _ = "end of CoverTab[118597]"
	case bool:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:136
		_go_fuzz_dep_.CoverTab[118598]++
											if s {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:137
			_go_fuzz_dep_.CoverTab[118603]++
												return 1, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:138
			// _ = "end of CoverTab[118603]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:139
			_go_fuzz_dep_.CoverTab[118604]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:139
			// _ = "end of CoverTab[118604]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:139
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:139
		// _ = "end of CoverTab[118598]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:139
		_go_fuzz_dep_.CoverTab[118599]++
											return 0, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:140
		// _ = "end of CoverTab[118599]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:141
		_go_fuzz_dep_.CoverTab[118600]++
											return 0, fmt.Errorf("unable to cast %#v of type %T to float64", i, i)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:142
		// _ = "end of CoverTab[118600]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:143
	// _ = "end of CoverTab[118583]"
}

// ToFloat32E casts an interface to a float32 type.
func ToFloat32E(i interface{}) (float32, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:147
	_go_fuzz_dep_.CoverTab[118605]++
										i = indirect(i)

										switch s := i.(type) {
	case float64:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:151
		_go_fuzz_dep_.CoverTab[118606]++
											return float32(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:152
		// _ = "end of CoverTab[118606]"
	case float32:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:153
		_go_fuzz_dep_.CoverTab[118607]++
											return s, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:154
		// _ = "end of CoverTab[118607]"
	case int:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:155
		_go_fuzz_dep_.CoverTab[118608]++
											return float32(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:156
		// _ = "end of CoverTab[118608]"
	case int64:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:157
		_go_fuzz_dep_.CoverTab[118609]++
											return float32(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:158
		// _ = "end of CoverTab[118609]"
	case int32:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:159
		_go_fuzz_dep_.CoverTab[118610]++
											return float32(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:160
		// _ = "end of CoverTab[118610]"
	case int16:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:161
		_go_fuzz_dep_.CoverTab[118611]++
											return float32(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:162
		// _ = "end of CoverTab[118611]"
	case int8:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:163
		_go_fuzz_dep_.CoverTab[118612]++
											return float32(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:164
		// _ = "end of CoverTab[118612]"
	case uint:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:165
		_go_fuzz_dep_.CoverTab[118613]++
											return float32(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:166
		// _ = "end of CoverTab[118613]"
	case uint64:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:167
		_go_fuzz_dep_.CoverTab[118614]++
											return float32(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:168
		// _ = "end of CoverTab[118614]"
	case uint32:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:169
		_go_fuzz_dep_.CoverTab[118615]++
											return float32(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:170
		// _ = "end of CoverTab[118615]"
	case uint16:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:171
		_go_fuzz_dep_.CoverTab[118616]++
											return float32(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:172
		// _ = "end of CoverTab[118616]"
	case uint8:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:173
		_go_fuzz_dep_.CoverTab[118617]++
											return float32(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:174
		// _ = "end of CoverTab[118617]"
	case string:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:175
		_go_fuzz_dep_.CoverTab[118618]++
											v, err := strconv.ParseFloat(s, 32)
											if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:177
			_go_fuzz_dep_.CoverTab[118623]++
												return float32(v), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:178
			// _ = "end of CoverTab[118623]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:179
			_go_fuzz_dep_.CoverTab[118624]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:179
			// _ = "end of CoverTab[118624]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:179
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:179
		// _ = "end of CoverTab[118618]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:179
		_go_fuzz_dep_.CoverTab[118619]++
											return 0, fmt.Errorf("unable to cast %#v of type %T to float32", i, i)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:180
		// _ = "end of CoverTab[118619]"
	case bool:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:181
		_go_fuzz_dep_.CoverTab[118620]++
											if s {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:182
			_go_fuzz_dep_.CoverTab[118625]++
												return 1, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:183
			// _ = "end of CoverTab[118625]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:184
			_go_fuzz_dep_.CoverTab[118626]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:184
			// _ = "end of CoverTab[118626]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:184
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:184
		// _ = "end of CoverTab[118620]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:184
		_go_fuzz_dep_.CoverTab[118621]++
											return 0, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:185
		// _ = "end of CoverTab[118621]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:186
		_go_fuzz_dep_.CoverTab[118622]++
											return 0, fmt.Errorf("unable to cast %#v of type %T to float32", i, i)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:187
		// _ = "end of CoverTab[118622]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:188
	// _ = "end of CoverTab[118605]"
}

// ToInt64E casts an interface to an int64 type.
func ToInt64E(i interface{}) (int64, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:192
	_go_fuzz_dep_.CoverTab[118627]++
										i = indirect(i)

										switch s := i.(type) {
	case int:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:196
		_go_fuzz_dep_.CoverTab[118628]++
											return int64(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:197
		// _ = "end of CoverTab[118628]"
	case int64:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:198
		_go_fuzz_dep_.CoverTab[118629]++
											return s, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:199
		// _ = "end of CoverTab[118629]"
	case int32:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:200
		_go_fuzz_dep_.CoverTab[118630]++
											return int64(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:201
		// _ = "end of CoverTab[118630]"
	case int16:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:202
		_go_fuzz_dep_.CoverTab[118631]++
											return int64(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:203
		// _ = "end of CoverTab[118631]"
	case int8:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:204
		_go_fuzz_dep_.CoverTab[118632]++
											return int64(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:205
		// _ = "end of CoverTab[118632]"
	case uint:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:206
		_go_fuzz_dep_.CoverTab[118633]++
											return int64(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:207
		// _ = "end of CoverTab[118633]"
	case uint64:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:208
		_go_fuzz_dep_.CoverTab[118634]++
											return int64(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:209
		// _ = "end of CoverTab[118634]"
	case uint32:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:210
		_go_fuzz_dep_.CoverTab[118635]++
											return int64(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:211
		// _ = "end of CoverTab[118635]"
	case uint16:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:212
		_go_fuzz_dep_.CoverTab[118636]++
											return int64(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:213
		// _ = "end of CoverTab[118636]"
	case uint8:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:214
		_go_fuzz_dep_.CoverTab[118637]++
											return int64(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:215
		// _ = "end of CoverTab[118637]"
	case float64:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:216
		_go_fuzz_dep_.CoverTab[118638]++
											return int64(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:217
		// _ = "end of CoverTab[118638]"
	case float32:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:218
		_go_fuzz_dep_.CoverTab[118639]++
											return int64(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:219
		// _ = "end of CoverTab[118639]"
	case string:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:220
		_go_fuzz_dep_.CoverTab[118640]++
											v, err := strconv.ParseInt(s, 0, 0)
											if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:222
			_go_fuzz_dep_.CoverTab[118646]++
												return v, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:223
			// _ = "end of CoverTab[118646]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:224
			_go_fuzz_dep_.CoverTab[118647]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:224
			// _ = "end of CoverTab[118647]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:224
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:224
		// _ = "end of CoverTab[118640]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:224
		_go_fuzz_dep_.CoverTab[118641]++
											return 0, fmt.Errorf("unable to cast %#v of type %T to int64", i, i)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:225
		// _ = "end of CoverTab[118641]"
	case bool:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:226
		_go_fuzz_dep_.CoverTab[118642]++
											if s {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:227
			_go_fuzz_dep_.CoverTab[118648]++
												return 1, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:228
			// _ = "end of CoverTab[118648]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:229
			_go_fuzz_dep_.CoverTab[118649]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:229
			// _ = "end of CoverTab[118649]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:229
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:229
		// _ = "end of CoverTab[118642]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:229
		_go_fuzz_dep_.CoverTab[118643]++
											return 0, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:230
		// _ = "end of CoverTab[118643]"
	case nil:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:231
		_go_fuzz_dep_.CoverTab[118644]++
											return 0, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:232
		// _ = "end of CoverTab[118644]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:233
		_go_fuzz_dep_.CoverTab[118645]++
											return 0, fmt.Errorf("unable to cast %#v of type %T to int64", i, i)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:234
		// _ = "end of CoverTab[118645]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:235
	// _ = "end of CoverTab[118627]"
}

// ToInt32E casts an interface to an int32 type.
func ToInt32E(i interface{}) (int32, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:239
	_go_fuzz_dep_.CoverTab[118650]++
										i = indirect(i)

										switch s := i.(type) {
	case int:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:243
		_go_fuzz_dep_.CoverTab[118651]++
											return int32(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:244
		// _ = "end of CoverTab[118651]"
	case int64:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:245
		_go_fuzz_dep_.CoverTab[118652]++
											return int32(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:246
		// _ = "end of CoverTab[118652]"
	case int32:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:247
		_go_fuzz_dep_.CoverTab[118653]++
											return s, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:248
		// _ = "end of CoverTab[118653]"
	case int16:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:249
		_go_fuzz_dep_.CoverTab[118654]++
											return int32(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:250
		// _ = "end of CoverTab[118654]"
	case int8:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:251
		_go_fuzz_dep_.CoverTab[118655]++
											return int32(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:252
		// _ = "end of CoverTab[118655]"
	case uint:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:253
		_go_fuzz_dep_.CoverTab[118656]++
											return int32(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:254
		// _ = "end of CoverTab[118656]"
	case uint64:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:255
		_go_fuzz_dep_.CoverTab[118657]++
											return int32(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:256
		// _ = "end of CoverTab[118657]"
	case uint32:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:257
		_go_fuzz_dep_.CoverTab[118658]++
											return int32(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:258
		// _ = "end of CoverTab[118658]"
	case uint16:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:259
		_go_fuzz_dep_.CoverTab[118659]++
											return int32(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:260
		// _ = "end of CoverTab[118659]"
	case uint8:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:261
		_go_fuzz_dep_.CoverTab[118660]++
											return int32(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:262
		// _ = "end of CoverTab[118660]"
	case float64:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:263
		_go_fuzz_dep_.CoverTab[118661]++
											return int32(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:264
		// _ = "end of CoverTab[118661]"
	case float32:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:265
		_go_fuzz_dep_.CoverTab[118662]++
											return int32(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:266
		// _ = "end of CoverTab[118662]"
	case string:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:267
		_go_fuzz_dep_.CoverTab[118663]++
											v, err := strconv.ParseInt(s, 0, 0)
											if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:269
			_go_fuzz_dep_.CoverTab[118669]++
												return int32(v), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:270
			// _ = "end of CoverTab[118669]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:271
			_go_fuzz_dep_.CoverTab[118670]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:271
			// _ = "end of CoverTab[118670]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:271
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:271
		// _ = "end of CoverTab[118663]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:271
		_go_fuzz_dep_.CoverTab[118664]++
											return 0, fmt.Errorf("unable to cast %#v of type %T to int32", i, i)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:272
		// _ = "end of CoverTab[118664]"
	case bool:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:273
		_go_fuzz_dep_.CoverTab[118665]++
											if s {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:274
			_go_fuzz_dep_.CoverTab[118671]++
												return 1, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:275
			// _ = "end of CoverTab[118671]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:276
			_go_fuzz_dep_.CoverTab[118672]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:276
			// _ = "end of CoverTab[118672]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:276
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:276
		// _ = "end of CoverTab[118665]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:276
		_go_fuzz_dep_.CoverTab[118666]++
											return 0, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:277
		// _ = "end of CoverTab[118666]"
	case nil:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:278
		_go_fuzz_dep_.CoverTab[118667]++
											return 0, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:279
		// _ = "end of CoverTab[118667]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:280
		_go_fuzz_dep_.CoverTab[118668]++
											return 0, fmt.Errorf("unable to cast %#v of type %T to int32", i, i)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:281
		// _ = "end of CoverTab[118668]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:282
	// _ = "end of CoverTab[118650]"
}

// ToInt16E casts an interface to an int16 type.
func ToInt16E(i interface{}) (int16, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:286
	_go_fuzz_dep_.CoverTab[118673]++
										i = indirect(i)

										switch s := i.(type) {
	case int:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:290
		_go_fuzz_dep_.CoverTab[118674]++
											return int16(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:291
		// _ = "end of CoverTab[118674]"
	case int64:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:292
		_go_fuzz_dep_.CoverTab[118675]++
											return int16(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:293
		// _ = "end of CoverTab[118675]"
	case int32:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:294
		_go_fuzz_dep_.CoverTab[118676]++
											return int16(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:295
		// _ = "end of CoverTab[118676]"
	case int16:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:296
		_go_fuzz_dep_.CoverTab[118677]++
											return s, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:297
		// _ = "end of CoverTab[118677]"
	case int8:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:298
		_go_fuzz_dep_.CoverTab[118678]++
											return int16(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:299
		// _ = "end of CoverTab[118678]"
	case uint:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:300
		_go_fuzz_dep_.CoverTab[118679]++
											return int16(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:301
		// _ = "end of CoverTab[118679]"
	case uint64:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:302
		_go_fuzz_dep_.CoverTab[118680]++
											return int16(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:303
		// _ = "end of CoverTab[118680]"
	case uint32:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:304
		_go_fuzz_dep_.CoverTab[118681]++
											return int16(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:305
		// _ = "end of CoverTab[118681]"
	case uint16:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:306
		_go_fuzz_dep_.CoverTab[118682]++
											return int16(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:307
		// _ = "end of CoverTab[118682]"
	case uint8:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:308
		_go_fuzz_dep_.CoverTab[118683]++
											return int16(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:309
		// _ = "end of CoverTab[118683]"
	case float64:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:310
		_go_fuzz_dep_.CoverTab[118684]++
											return int16(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:311
		// _ = "end of CoverTab[118684]"
	case float32:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:312
		_go_fuzz_dep_.CoverTab[118685]++
											return int16(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:313
		// _ = "end of CoverTab[118685]"
	case string:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:314
		_go_fuzz_dep_.CoverTab[118686]++
											v, err := strconv.ParseInt(s, 0, 0)
											if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:316
			_go_fuzz_dep_.CoverTab[118692]++
												return int16(v), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:317
			// _ = "end of CoverTab[118692]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:318
			_go_fuzz_dep_.CoverTab[118693]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:318
			// _ = "end of CoverTab[118693]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:318
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:318
		// _ = "end of CoverTab[118686]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:318
		_go_fuzz_dep_.CoverTab[118687]++
											return 0, fmt.Errorf("unable to cast %#v of type %T to int16", i, i)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:319
		// _ = "end of CoverTab[118687]"
	case bool:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:320
		_go_fuzz_dep_.CoverTab[118688]++
											if s {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:321
			_go_fuzz_dep_.CoverTab[118694]++
												return 1, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:322
			// _ = "end of CoverTab[118694]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:323
			_go_fuzz_dep_.CoverTab[118695]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:323
			// _ = "end of CoverTab[118695]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:323
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:323
		// _ = "end of CoverTab[118688]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:323
		_go_fuzz_dep_.CoverTab[118689]++
											return 0, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:324
		// _ = "end of CoverTab[118689]"
	case nil:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:325
		_go_fuzz_dep_.CoverTab[118690]++
											return 0, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:326
		// _ = "end of CoverTab[118690]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:327
		_go_fuzz_dep_.CoverTab[118691]++
											return 0, fmt.Errorf("unable to cast %#v of type %T to int16", i, i)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:328
		// _ = "end of CoverTab[118691]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:329
	// _ = "end of CoverTab[118673]"
}

// ToInt8E casts an interface to an int8 type.
func ToInt8E(i interface{}) (int8, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:333
	_go_fuzz_dep_.CoverTab[118696]++
										i = indirect(i)

										switch s := i.(type) {
	case int:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:337
		_go_fuzz_dep_.CoverTab[118697]++
											return int8(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:338
		// _ = "end of CoverTab[118697]"
	case int64:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:339
		_go_fuzz_dep_.CoverTab[118698]++
											return int8(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:340
		// _ = "end of CoverTab[118698]"
	case int32:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:341
		_go_fuzz_dep_.CoverTab[118699]++
											return int8(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:342
		// _ = "end of CoverTab[118699]"
	case int16:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:343
		_go_fuzz_dep_.CoverTab[118700]++
											return int8(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:344
		// _ = "end of CoverTab[118700]"
	case int8:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:345
		_go_fuzz_dep_.CoverTab[118701]++
											return s, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:346
		// _ = "end of CoverTab[118701]"
	case uint:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:347
		_go_fuzz_dep_.CoverTab[118702]++
											return int8(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:348
		// _ = "end of CoverTab[118702]"
	case uint64:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:349
		_go_fuzz_dep_.CoverTab[118703]++
											return int8(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:350
		// _ = "end of CoverTab[118703]"
	case uint32:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:351
		_go_fuzz_dep_.CoverTab[118704]++
											return int8(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:352
		// _ = "end of CoverTab[118704]"
	case uint16:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:353
		_go_fuzz_dep_.CoverTab[118705]++
											return int8(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:354
		// _ = "end of CoverTab[118705]"
	case uint8:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:355
		_go_fuzz_dep_.CoverTab[118706]++
											return int8(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:356
		// _ = "end of CoverTab[118706]"
	case float64:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:357
		_go_fuzz_dep_.CoverTab[118707]++
											return int8(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:358
		// _ = "end of CoverTab[118707]"
	case float32:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:359
		_go_fuzz_dep_.CoverTab[118708]++
											return int8(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:360
		// _ = "end of CoverTab[118708]"
	case string:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:361
		_go_fuzz_dep_.CoverTab[118709]++
											v, err := strconv.ParseInt(s, 0, 0)
											if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:363
			_go_fuzz_dep_.CoverTab[118715]++
												return int8(v), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:364
			// _ = "end of CoverTab[118715]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:365
			_go_fuzz_dep_.CoverTab[118716]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:365
			// _ = "end of CoverTab[118716]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:365
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:365
		// _ = "end of CoverTab[118709]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:365
		_go_fuzz_dep_.CoverTab[118710]++
											return 0, fmt.Errorf("unable to cast %#v of type %T to int8", i, i)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:366
		// _ = "end of CoverTab[118710]"
	case bool:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:367
		_go_fuzz_dep_.CoverTab[118711]++
											if s {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:368
			_go_fuzz_dep_.CoverTab[118717]++
												return 1, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:369
			// _ = "end of CoverTab[118717]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:370
			_go_fuzz_dep_.CoverTab[118718]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:370
			// _ = "end of CoverTab[118718]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:370
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:370
		// _ = "end of CoverTab[118711]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:370
		_go_fuzz_dep_.CoverTab[118712]++
											return 0, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:371
		// _ = "end of CoverTab[118712]"
	case nil:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:372
		_go_fuzz_dep_.CoverTab[118713]++
											return 0, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:373
		// _ = "end of CoverTab[118713]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:374
		_go_fuzz_dep_.CoverTab[118714]++
											return 0, fmt.Errorf("unable to cast %#v of type %T to int8", i, i)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:375
		// _ = "end of CoverTab[118714]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:376
	// _ = "end of CoverTab[118696]"
}

// ToIntE casts an interface to an int type.
func ToIntE(i interface{}) (int, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:380
	_go_fuzz_dep_.CoverTab[118719]++
										i = indirect(i)

										switch s := i.(type) {
	case int:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:384
		_go_fuzz_dep_.CoverTab[118720]++
											return s, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:385
		// _ = "end of CoverTab[118720]"
	case int64:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:386
		_go_fuzz_dep_.CoverTab[118721]++
											return int(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:387
		// _ = "end of CoverTab[118721]"
	case int32:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:388
		_go_fuzz_dep_.CoverTab[118722]++
											return int(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:389
		// _ = "end of CoverTab[118722]"
	case int16:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:390
		_go_fuzz_dep_.CoverTab[118723]++
											return int(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:391
		// _ = "end of CoverTab[118723]"
	case int8:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:392
		_go_fuzz_dep_.CoverTab[118724]++
											return int(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:393
		// _ = "end of CoverTab[118724]"
	case uint:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:394
		_go_fuzz_dep_.CoverTab[118725]++
											return int(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:395
		// _ = "end of CoverTab[118725]"
	case uint64:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:396
		_go_fuzz_dep_.CoverTab[118726]++
											return int(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:397
		// _ = "end of CoverTab[118726]"
	case uint32:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:398
		_go_fuzz_dep_.CoverTab[118727]++
											return int(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:399
		// _ = "end of CoverTab[118727]"
	case uint16:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:400
		_go_fuzz_dep_.CoverTab[118728]++
											return int(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:401
		// _ = "end of CoverTab[118728]"
	case uint8:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:402
		_go_fuzz_dep_.CoverTab[118729]++
											return int(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:403
		// _ = "end of CoverTab[118729]"
	case float64:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:404
		_go_fuzz_dep_.CoverTab[118730]++
											return int(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:405
		// _ = "end of CoverTab[118730]"
	case float32:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:406
		_go_fuzz_dep_.CoverTab[118731]++
											return int(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:407
		// _ = "end of CoverTab[118731]"
	case string:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:408
		_go_fuzz_dep_.CoverTab[118732]++
											v, err := strconv.ParseInt(s, 0, 0)
											if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:410
			_go_fuzz_dep_.CoverTab[118738]++
												return int(v), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:411
			// _ = "end of CoverTab[118738]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:412
			_go_fuzz_dep_.CoverTab[118739]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:412
			// _ = "end of CoverTab[118739]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:412
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:412
		// _ = "end of CoverTab[118732]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:412
		_go_fuzz_dep_.CoverTab[118733]++
											return 0, fmt.Errorf("unable to cast %#v of type %T to int", i, i)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:413
		// _ = "end of CoverTab[118733]"
	case bool:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:414
		_go_fuzz_dep_.CoverTab[118734]++
											if s {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:415
			_go_fuzz_dep_.CoverTab[118740]++
												return 1, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:416
			// _ = "end of CoverTab[118740]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:417
			_go_fuzz_dep_.CoverTab[118741]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:417
			// _ = "end of CoverTab[118741]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:417
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:417
		// _ = "end of CoverTab[118734]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:417
		_go_fuzz_dep_.CoverTab[118735]++
											return 0, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:418
		// _ = "end of CoverTab[118735]"
	case nil:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:419
		_go_fuzz_dep_.CoverTab[118736]++
											return 0, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:420
		// _ = "end of CoverTab[118736]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:421
		_go_fuzz_dep_.CoverTab[118737]++
											return 0, fmt.Errorf("unable to cast %#v of type %T to int", i, i)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:422
		// _ = "end of CoverTab[118737]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:423
	// _ = "end of CoverTab[118719]"
}

// ToUintE casts an interface to a uint type.
func ToUintE(i interface{}) (uint, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:427
	_go_fuzz_dep_.CoverTab[118742]++
										i = indirect(i)

										switch s := i.(type) {
	case string:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:431
		_go_fuzz_dep_.CoverTab[118743]++
											v, err := strconv.ParseUint(s, 0, 0)
											if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:433
			_go_fuzz_dep_.CoverTab[118768]++
												return uint(v), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:434
			// _ = "end of CoverTab[118768]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:435
			_go_fuzz_dep_.CoverTab[118769]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:435
			// _ = "end of CoverTab[118769]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:435
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:435
		// _ = "end of CoverTab[118743]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:435
		_go_fuzz_dep_.CoverTab[118744]++
											return 0, fmt.Errorf("unable to cast %#v to uint: %s", i, err)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:436
		// _ = "end of CoverTab[118744]"
	case int:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:437
		_go_fuzz_dep_.CoverTab[118745]++
											if s < 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:438
			_go_fuzz_dep_.CoverTab[118770]++
												return 0, errNegativeNotAllowed
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:439
			// _ = "end of CoverTab[118770]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:440
			_go_fuzz_dep_.CoverTab[118771]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:440
			// _ = "end of CoverTab[118771]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:440
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:440
		// _ = "end of CoverTab[118745]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:440
		_go_fuzz_dep_.CoverTab[118746]++
											return uint(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:441
		// _ = "end of CoverTab[118746]"
	case int64:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:442
		_go_fuzz_dep_.CoverTab[118747]++
											if s < 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:443
			_go_fuzz_dep_.CoverTab[118772]++
												return 0, errNegativeNotAllowed
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:444
			// _ = "end of CoverTab[118772]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:445
			_go_fuzz_dep_.CoverTab[118773]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:445
			// _ = "end of CoverTab[118773]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:445
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:445
		// _ = "end of CoverTab[118747]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:445
		_go_fuzz_dep_.CoverTab[118748]++
											return uint(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:446
		// _ = "end of CoverTab[118748]"
	case int32:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:447
		_go_fuzz_dep_.CoverTab[118749]++
											if s < 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:448
			_go_fuzz_dep_.CoverTab[118774]++
												return 0, errNegativeNotAllowed
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:449
			// _ = "end of CoverTab[118774]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:450
			_go_fuzz_dep_.CoverTab[118775]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:450
			// _ = "end of CoverTab[118775]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:450
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:450
		// _ = "end of CoverTab[118749]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:450
		_go_fuzz_dep_.CoverTab[118750]++
											return uint(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:451
		// _ = "end of CoverTab[118750]"
	case int16:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:452
		_go_fuzz_dep_.CoverTab[118751]++
											if s < 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:453
			_go_fuzz_dep_.CoverTab[118776]++
												return 0, errNegativeNotAllowed
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:454
			// _ = "end of CoverTab[118776]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:455
			_go_fuzz_dep_.CoverTab[118777]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:455
			// _ = "end of CoverTab[118777]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:455
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:455
		// _ = "end of CoverTab[118751]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:455
		_go_fuzz_dep_.CoverTab[118752]++
											return uint(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:456
		// _ = "end of CoverTab[118752]"
	case int8:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:457
		_go_fuzz_dep_.CoverTab[118753]++
											if s < 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:458
			_go_fuzz_dep_.CoverTab[118778]++
												return 0, errNegativeNotAllowed
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:459
			// _ = "end of CoverTab[118778]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:460
			_go_fuzz_dep_.CoverTab[118779]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:460
			// _ = "end of CoverTab[118779]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:460
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:460
		// _ = "end of CoverTab[118753]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:460
		_go_fuzz_dep_.CoverTab[118754]++
											return uint(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:461
		// _ = "end of CoverTab[118754]"
	case uint:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:462
		_go_fuzz_dep_.CoverTab[118755]++
											return s, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:463
		// _ = "end of CoverTab[118755]"
	case uint64:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:464
		_go_fuzz_dep_.CoverTab[118756]++
											return uint(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:465
		// _ = "end of CoverTab[118756]"
	case uint32:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:466
		_go_fuzz_dep_.CoverTab[118757]++
											return uint(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:467
		// _ = "end of CoverTab[118757]"
	case uint16:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:468
		_go_fuzz_dep_.CoverTab[118758]++
											return uint(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:469
		// _ = "end of CoverTab[118758]"
	case uint8:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:470
		_go_fuzz_dep_.CoverTab[118759]++
											return uint(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:471
		// _ = "end of CoverTab[118759]"
	case float64:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:472
		_go_fuzz_dep_.CoverTab[118760]++
											if s < 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:473
			_go_fuzz_dep_.CoverTab[118780]++
												return 0, errNegativeNotAllowed
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:474
			// _ = "end of CoverTab[118780]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:475
			_go_fuzz_dep_.CoverTab[118781]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:475
			// _ = "end of CoverTab[118781]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:475
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:475
		// _ = "end of CoverTab[118760]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:475
		_go_fuzz_dep_.CoverTab[118761]++
											return uint(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:476
		// _ = "end of CoverTab[118761]"
	case float32:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:477
		_go_fuzz_dep_.CoverTab[118762]++
											if s < 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:478
			_go_fuzz_dep_.CoverTab[118782]++
												return 0, errNegativeNotAllowed
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:479
			// _ = "end of CoverTab[118782]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:480
			_go_fuzz_dep_.CoverTab[118783]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:480
			// _ = "end of CoverTab[118783]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:480
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:480
		// _ = "end of CoverTab[118762]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:480
		_go_fuzz_dep_.CoverTab[118763]++
											return uint(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:481
		// _ = "end of CoverTab[118763]"
	case bool:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:482
		_go_fuzz_dep_.CoverTab[118764]++
											if s {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:483
			_go_fuzz_dep_.CoverTab[118784]++
												return 1, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:484
			// _ = "end of CoverTab[118784]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:485
			_go_fuzz_dep_.CoverTab[118785]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:485
			// _ = "end of CoverTab[118785]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:485
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:485
		// _ = "end of CoverTab[118764]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:485
		_go_fuzz_dep_.CoverTab[118765]++
											return 0, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:486
		// _ = "end of CoverTab[118765]"
	case nil:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:487
		_go_fuzz_dep_.CoverTab[118766]++
											return 0, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:488
		// _ = "end of CoverTab[118766]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:489
		_go_fuzz_dep_.CoverTab[118767]++
											return 0, fmt.Errorf("unable to cast %#v of type %T to uint", i, i)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:490
		// _ = "end of CoverTab[118767]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:491
	// _ = "end of CoverTab[118742]"
}

// ToUint64E casts an interface to a uint64 type.
func ToUint64E(i interface{}) (uint64, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:495
	_go_fuzz_dep_.CoverTab[118786]++
										i = indirect(i)

										switch s := i.(type) {
	case string:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:499
		_go_fuzz_dep_.CoverTab[118787]++
											v, err := strconv.ParseUint(s, 0, 64)
											if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:501
			_go_fuzz_dep_.CoverTab[118812]++
												return v, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:502
			// _ = "end of CoverTab[118812]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:503
			_go_fuzz_dep_.CoverTab[118813]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:503
			// _ = "end of CoverTab[118813]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:503
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:503
		// _ = "end of CoverTab[118787]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:503
		_go_fuzz_dep_.CoverTab[118788]++
											return 0, fmt.Errorf("unable to cast %#v to uint64: %s", i, err)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:504
		// _ = "end of CoverTab[118788]"
	case int:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:505
		_go_fuzz_dep_.CoverTab[118789]++
											if s < 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:506
			_go_fuzz_dep_.CoverTab[118814]++
												return 0, errNegativeNotAllowed
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:507
			// _ = "end of CoverTab[118814]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:508
			_go_fuzz_dep_.CoverTab[118815]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:508
			// _ = "end of CoverTab[118815]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:508
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:508
		// _ = "end of CoverTab[118789]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:508
		_go_fuzz_dep_.CoverTab[118790]++
											return uint64(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:509
		// _ = "end of CoverTab[118790]"
	case int64:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:510
		_go_fuzz_dep_.CoverTab[118791]++
											if s < 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:511
			_go_fuzz_dep_.CoverTab[118816]++
												return 0, errNegativeNotAllowed
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:512
			// _ = "end of CoverTab[118816]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:513
			_go_fuzz_dep_.CoverTab[118817]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:513
			// _ = "end of CoverTab[118817]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:513
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:513
		// _ = "end of CoverTab[118791]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:513
		_go_fuzz_dep_.CoverTab[118792]++
											return uint64(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:514
		// _ = "end of CoverTab[118792]"
	case int32:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:515
		_go_fuzz_dep_.CoverTab[118793]++
											if s < 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:516
			_go_fuzz_dep_.CoverTab[118818]++
												return 0, errNegativeNotAllowed
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:517
			// _ = "end of CoverTab[118818]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:518
			_go_fuzz_dep_.CoverTab[118819]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:518
			// _ = "end of CoverTab[118819]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:518
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:518
		// _ = "end of CoverTab[118793]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:518
		_go_fuzz_dep_.CoverTab[118794]++
											return uint64(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:519
		// _ = "end of CoverTab[118794]"
	case int16:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:520
		_go_fuzz_dep_.CoverTab[118795]++
											if s < 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:521
			_go_fuzz_dep_.CoverTab[118820]++
												return 0, errNegativeNotAllowed
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:522
			// _ = "end of CoverTab[118820]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:523
			_go_fuzz_dep_.CoverTab[118821]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:523
			// _ = "end of CoverTab[118821]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:523
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:523
		// _ = "end of CoverTab[118795]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:523
		_go_fuzz_dep_.CoverTab[118796]++
											return uint64(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:524
		// _ = "end of CoverTab[118796]"
	case int8:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:525
		_go_fuzz_dep_.CoverTab[118797]++
											if s < 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:526
			_go_fuzz_dep_.CoverTab[118822]++
												return 0, errNegativeNotAllowed
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:527
			// _ = "end of CoverTab[118822]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:528
			_go_fuzz_dep_.CoverTab[118823]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:528
			// _ = "end of CoverTab[118823]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:528
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:528
		// _ = "end of CoverTab[118797]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:528
		_go_fuzz_dep_.CoverTab[118798]++
											return uint64(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:529
		// _ = "end of CoverTab[118798]"
	case uint:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:530
		_go_fuzz_dep_.CoverTab[118799]++
											return uint64(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:531
		// _ = "end of CoverTab[118799]"
	case uint64:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:532
		_go_fuzz_dep_.CoverTab[118800]++
											return s, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:533
		// _ = "end of CoverTab[118800]"
	case uint32:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:534
		_go_fuzz_dep_.CoverTab[118801]++
											return uint64(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:535
		// _ = "end of CoverTab[118801]"
	case uint16:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:536
		_go_fuzz_dep_.CoverTab[118802]++
											return uint64(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:537
		// _ = "end of CoverTab[118802]"
	case uint8:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:538
		_go_fuzz_dep_.CoverTab[118803]++
											return uint64(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:539
		// _ = "end of CoverTab[118803]"
	case float32:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:540
		_go_fuzz_dep_.CoverTab[118804]++
											if s < 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:541
			_go_fuzz_dep_.CoverTab[118824]++
												return 0, errNegativeNotAllowed
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:542
			// _ = "end of CoverTab[118824]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:543
			_go_fuzz_dep_.CoverTab[118825]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:543
			// _ = "end of CoverTab[118825]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:543
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:543
		// _ = "end of CoverTab[118804]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:543
		_go_fuzz_dep_.CoverTab[118805]++
											return uint64(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:544
		// _ = "end of CoverTab[118805]"
	case float64:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:545
		_go_fuzz_dep_.CoverTab[118806]++
											if s < 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:546
			_go_fuzz_dep_.CoverTab[118826]++
												return 0, errNegativeNotAllowed
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:547
			// _ = "end of CoverTab[118826]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:548
			_go_fuzz_dep_.CoverTab[118827]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:548
			// _ = "end of CoverTab[118827]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:548
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:548
		// _ = "end of CoverTab[118806]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:548
		_go_fuzz_dep_.CoverTab[118807]++
											return uint64(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:549
		// _ = "end of CoverTab[118807]"
	case bool:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:550
		_go_fuzz_dep_.CoverTab[118808]++
											if s {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:551
			_go_fuzz_dep_.CoverTab[118828]++
												return 1, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:552
			// _ = "end of CoverTab[118828]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:553
			_go_fuzz_dep_.CoverTab[118829]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:553
			// _ = "end of CoverTab[118829]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:553
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:553
		// _ = "end of CoverTab[118808]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:553
		_go_fuzz_dep_.CoverTab[118809]++
											return 0, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:554
		// _ = "end of CoverTab[118809]"
	case nil:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:555
		_go_fuzz_dep_.CoverTab[118810]++
											return 0, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:556
		// _ = "end of CoverTab[118810]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:557
		_go_fuzz_dep_.CoverTab[118811]++
											return 0, fmt.Errorf("unable to cast %#v of type %T to uint64", i, i)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:558
		// _ = "end of CoverTab[118811]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:559
	// _ = "end of CoverTab[118786]"
}

// ToUint32E casts an interface to a uint32 type.
func ToUint32E(i interface{}) (uint32, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:563
	_go_fuzz_dep_.CoverTab[118830]++
										i = indirect(i)

										switch s := i.(type) {
	case string:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:567
		_go_fuzz_dep_.CoverTab[118831]++
											v, err := strconv.ParseUint(s, 0, 32)
											if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:569
			_go_fuzz_dep_.CoverTab[118856]++
												return uint32(v), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:570
			// _ = "end of CoverTab[118856]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:571
			_go_fuzz_dep_.CoverTab[118857]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:571
			// _ = "end of CoverTab[118857]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:571
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:571
		// _ = "end of CoverTab[118831]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:571
		_go_fuzz_dep_.CoverTab[118832]++
											return 0, fmt.Errorf("unable to cast %#v to uint32: %s", i, err)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:572
		// _ = "end of CoverTab[118832]"
	case int:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:573
		_go_fuzz_dep_.CoverTab[118833]++
											if s < 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:574
			_go_fuzz_dep_.CoverTab[118858]++
												return 0, errNegativeNotAllowed
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:575
			// _ = "end of CoverTab[118858]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:576
			_go_fuzz_dep_.CoverTab[118859]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:576
			// _ = "end of CoverTab[118859]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:576
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:576
		// _ = "end of CoverTab[118833]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:576
		_go_fuzz_dep_.CoverTab[118834]++
											return uint32(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:577
		// _ = "end of CoverTab[118834]"
	case int64:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:578
		_go_fuzz_dep_.CoverTab[118835]++
											if s < 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:579
			_go_fuzz_dep_.CoverTab[118860]++
												return 0, errNegativeNotAllowed
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:580
			// _ = "end of CoverTab[118860]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:581
			_go_fuzz_dep_.CoverTab[118861]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:581
			// _ = "end of CoverTab[118861]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:581
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:581
		// _ = "end of CoverTab[118835]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:581
		_go_fuzz_dep_.CoverTab[118836]++
											return uint32(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:582
		// _ = "end of CoverTab[118836]"
	case int32:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:583
		_go_fuzz_dep_.CoverTab[118837]++
											if s < 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:584
			_go_fuzz_dep_.CoverTab[118862]++
												return 0, errNegativeNotAllowed
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:585
			// _ = "end of CoverTab[118862]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:586
			_go_fuzz_dep_.CoverTab[118863]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:586
			// _ = "end of CoverTab[118863]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:586
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:586
		// _ = "end of CoverTab[118837]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:586
		_go_fuzz_dep_.CoverTab[118838]++
											return uint32(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:587
		// _ = "end of CoverTab[118838]"
	case int16:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:588
		_go_fuzz_dep_.CoverTab[118839]++
											if s < 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:589
			_go_fuzz_dep_.CoverTab[118864]++
												return 0, errNegativeNotAllowed
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:590
			// _ = "end of CoverTab[118864]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:591
			_go_fuzz_dep_.CoverTab[118865]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:591
			// _ = "end of CoverTab[118865]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:591
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:591
		// _ = "end of CoverTab[118839]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:591
		_go_fuzz_dep_.CoverTab[118840]++
											return uint32(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:592
		// _ = "end of CoverTab[118840]"
	case int8:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:593
		_go_fuzz_dep_.CoverTab[118841]++
											if s < 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:594
			_go_fuzz_dep_.CoverTab[118866]++
												return 0, errNegativeNotAllowed
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:595
			// _ = "end of CoverTab[118866]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:596
			_go_fuzz_dep_.CoverTab[118867]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:596
			// _ = "end of CoverTab[118867]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:596
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:596
		// _ = "end of CoverTab[118841]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:596
		_go_fuzz_dep_.CoverTab[118842]++
											return uint32(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:597
		// _ = "end of CoverTab[118842]"
	case uint:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:598
		_go_fuzz_dep_.CoverTab[118843]++
											return uint32(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:599
		// _ = "end of CoverTab[118843]"
	case uint64:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:600
		_go_fuzz_dep_.CoverTab[118844]++
											return uint32(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:601
		// _ = "end of CoverTab[118844]"
	case uint32:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:602
		_go_fuzz_dep_.CoverTab[118845]++
											return s, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:603
		// _ = "end of CoverTab[118845]"
	case uint16:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:604
		_go_fuzz_dep_.CoverTab[118846]++
											return uint32(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:605
		// _ = "end of CoverTab[118846]"
	case uint8:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:606
		_go_fuzz_dep_.CoverTab[118847]++
											return uint32(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:607
		// _ = "end of CoverTab[118847]"
	case float64:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:608
		_go_fuzz_dep_.CoverTab[118848]++
											if s < 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:609
			_go_fuzz_dep_.CoverTab[118868]++
												return 0, errNegativeNotAllowed
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:610
			// _ = "end of CoverTab[118868]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:611
			_go_fuzz_dep_.CoverTab[118869]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:611
			// _ = "end of CoverTab[118869]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:611
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:611
		// _ = "end of CoverTab[118848]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:611
		_go_fuzz_dep_.CoverTab[118849]++
											return uint32(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:612
		// _ = "end of CoverTab[118849]"
	case float32:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:613
		_go_fuzz_dep_.CoverTab[118850]++
											if s < 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:614
			_go_fuzz_dep_.CoverTab[118870]++
												return 0, errNegativeNotAllowed
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:615
			// _ = "end of CoverTab[118870]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:616
			_go_fuzz_dep_.CoverTab[118871]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:616
			// _ = "end of CoverTab[118871]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:616
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:616
		// _ = "end of CoverTab[118850]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:616
		_go_fuzz_dep_.CoverTab[118851]++
											return uint32(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:617
		// _ = "end of CoverTab[118851]"
	case bool:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:618
		_go_fuzz_dep_.CoverTab[118852]++
											if s {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:619
			_go_fuzz_dep_.CoverTab[118872]++
												return 1, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:620
			// _ = "end of CoverTab[118872]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:621
			_go_fuzz_dep_.CoverTab[118873]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:621
			// _ = "end of CoverTab[118873]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:621
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:621
		// _ = "end of CoverTab[118852]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:621
		_go_fuzz_dep_.CoverTab[118853]++
											return 0, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:622
		// _ = "end of CoverTab[118853]"
	case nil:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:623
		_go_fuzz_dep_.CoverTab[118854]++
											return 0, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:624
		// _ = "end of CoverTab[118854]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:625
		_go_fuzz_dep_.CoverTab[118855]++
											return 0, fmt.Errorf("unable to cast %#v of type %T to uint32", i, i)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:626
		// _ = "end of CoverTab[118855]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:627
	// _ = "end of CoverTab[118830]"
}

// ToUint16E casts an interface to a uint16 type.
func ToUint16E(i interface{}) (uint16, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:631
	_go_fuzz_dep_.CoverTab[118874]++
										i = indirect(i)

										switch s := i.(type) {
	case string:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:635
		_go_fuzz_dep_.CoverTab[118875]++
											v, err := strconv.ParseUint(s, 0, 16)
											if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:637
			_go_fuzz_dep_.CoverTab[118900]++
												return uint16(v), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:638
			// _ = "end of CoverTab[118900]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:639
			_go_fuzz_dep_.CoverTab[118901]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:639
			// _ = "end of CoverTab[118901]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:639
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:639
		// _ = "end of CoverTab[118875]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:639
		_go_fuzz_dep_.CoverTab[118876]++
											return 0, fmt.Errorf("unable to cast %#v to uint16: %s", i, err)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:640
		// _ = "end of CoverTab[118876]"
	case int:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:641
		_go_fuzz_dep_.CoverTab[118877]++
											if s < 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:642
			_go_fuzz_dep_.CoverTab[118902]++
												return 0, errNegativeNotAllowed
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:643
			// _ = "end of CoverTab[118902]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:644
			_go_fuzz_dep_.CoverTab[118903]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:644
			// _ = "end of CoverTab[118903]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:644
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:644
		// _ = "end of CoverTab[118877]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:644
		_go_fuzz_dep_.CoverTab[118878]++
											return uint16(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:645
		// _ = "end of CoverTab[118878]"
	case int64:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:646
		_go_fuzz_dep_.CoverTab[118879]++
											if s < 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:647
			_go_fuzz_dep_.CoverTab[118904]++
												return 0, errNegativeNotAllowed
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:648
			// _ = "end of CoverTab[118904]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:649
			_go_fuzz_dep_.CoverTab[118905]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:649
			// _ = "end of CoverTab[118905]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:649
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:649
		// _ = "end of CoverTab[118879]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:649
		_go_fuzz_dep_.CoverTab[118880]++
											return uint16(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:650
		// _ = "end of CoverTab[118880]"
	case int32:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:651
		_go_fuzz_dep_.CoverTab[118881]++
											if s < 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:652
			_go_fuzz_dep_.CoverTab[118906]++
												return 0, errNegativeNotAllowed
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:653
			// _ = "end of CoverTab[118906]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:654
			_go_fuzz_dep_.CoverTab[118907]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:654
			// _ = "end of CoverTab[118907]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:654
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:654
		// _ = "end of CoverTab[118881]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:654
		_go_fuzz_dep_.CoverTab[118882]++
											return uint16(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:655
		// _ = "end of CoverTab[118882]"
	case int16:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:656
		_go_fuzz_dep_.CoverTab[118883]++
											if s < 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:657
			_go_fuzz_dep_.CoverTab[118908]++
												return 0, errNegativeNotAllowed
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:658
			// _ = "end of CoverTab[118908]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:659
			_go_fuzz_dep_.CoverTab[118909]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:659
			// _ = "end of CoverTab[118909]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:659
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:659
		// _ = "end of CoverTab[118883]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:659
		_go_fuzz_dep_.CoverTab[118884]++
											return uint16(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:660
		// _ = "end of CoverTab[118884]"
	case int8:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:661
		_go_fuzz_dep_.CoverTab[118885]++
											if s < 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:662
			_go_fuzz_dep_.CoverTab[118910]++
												return 0, errNegativeNotAllowed
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:663
			// _ = "end of CoverTab[118910]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:664
			_go_fuzz_dep_.CoverTab[118911]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:664
			// _ = "end of CoverTab[118911]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:664
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:664
		// _ = "end of CoverTab[118885]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:664
		_go_fuzz_dep_.CoverTab[118886]++
											return uint16(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:665
		// _ = "end of CoverTab[118886]"
	case uint:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:666
		_go_fuzz_dep_.CoverTab[118887]++
											return uint16(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:667
		// _ = "end of CoverTab[118887]"
	case uint64:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:668
		_go_fuzz_dep_.CoverTab[118888]++
											return uint16(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:669
		// _ = "end of CoverTab[118888]"
	case uint32:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:670
		_go_fuzz_dep_.CoverTab[118889]++
											return uint16(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:671
		// _ = "end of CoverTab[118889]"
	case uint16:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:672
		_go_fuzz_dep_.CoverTab[118890]++
											return s, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:673
		// _ = "end of CoverTab[118890]"
	case uint8:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:674
		_go_fuzz_dep_.CoverTab[118891]++
											return uint16(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:675
		// _ = "end of CoverTab[118891]"
	case float64:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:676
		_go_fuzz_dep_.CoverTab[118892]++
											if s < 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:677
			_go_fuzz_dep_.CoverTab[118912]++
												return 0, errNegativeNotAllowed
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:678
			// _ = "end of CoverTab[118912]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:679
			_go_fuzz_dep_.CoverTab[118913]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:679
			// _ = "end of CoverTab[118913]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:679
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:679
		// _ = "end of CoverTab[118892]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:679
		_go_fuzz_dep_.CoverTab[118893]++
											return uint16(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:680
		// _ = "end of CoverTab[118893]"
	case float32:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:681
		_go_fuzz_dep_.CoverTab[118894]++
											if s < 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:682
			_go_fuzz_dep_.CoverTab[118914]++
												return 0, errNegativeNotAllowed
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:683
			// _ = "end of CoverTab[118914]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:684
			_go_fuzz_dep_.CoverTab[118915]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:684
			// _ = "end of CoverTab[118915]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:684
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:684
		// _ = "end of CoverTab[118894]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:684
		_go_fuzz_dep_.CoverTab[118895]++
											return uint16(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:685
		// _ = "end of CoverTab[118895]"
	case bool:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:686
		_go_fuzz_dep_.CoverTab[118896]++
											if s {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:687
			_go_fuzz_dep_.CoverTab[118916]++
												return 1, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:688
			// _ = "end of CoverTab[118916]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:689
			_go_fuzz_dep_.CoverTab[118917]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:689
			// _ = "end of CoverTab[118917]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:689
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:689
		// _ = "end of CoverTab[118896]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:689
		_go_fuzz_dep_.CoverTab[118897]++
											return 0, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:690
		// _ = "end of CoverTab[118897]"
	case nil:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:691
		_go_fuzz_dep_.CoverTab[118898]++
											return 0, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:692
		// _ = "end of CoverTab[118898]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:693
		_go_fuzz_dep_.CoverTab[118899]++
											return 0, fmt.Errorf("unable to cast %#v of type %T to uint16", i, i)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:694
		// _ = "end of CoverTab[118899]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:695
	// _ = "end of CoverTab[118874]"
}

// ToUint8E casts an interface to a uint type.
func ToUint8E(i interface{}) (uint8, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:699
	_go_fuzz_dep_.CoverTab[118918]++
										i = indirect(i)

										switch s := i.(type) {
	case string:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:703
		_go_fuzz_dep_.CoverTab[118919]++
											v, err := strconv.ParseUint(s, 0, 8)
											if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:705
			_go_fuzz_dep_.CoverTab[118944]++
												return uint8(v), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:706
			// _ = "end of CoverTab[118944]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:707
			_go_fuzz_dep_.CoverTab[118945]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:707
			// _ = "end of CoverTab[118945]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:707
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:707
		// _ = "end of CoverTab[118919]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:707
		_go_fuzz_dep_.CoverTab[118920]++
											return 0, fmt.Errorf("unable to cast %#v to uint8: %s", i, err)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:708
		// _ = "end of CoverTab[118920]"
	case int:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:709
		_go_fuzz_dep_.CoverTab[118921]++
											if s < 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:710
			_go_fuzz_dep_.CoverTab[118946]++
												return 0, errNegativeNotAllowed
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:711
			// _ = "end of CoverTab[118946]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:712
			_go_fuzz_dep_.CoverTab[118947]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:712
			// _ = "end of CoverTab[118947]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:712
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:712
		// _ = "end of CoverTab[118921]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:712
		_go_fuzz_dep_.CoverTab[118922]++
											return uint8(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:713
		// _ = "end of CoverTab[118922]"
	case int64:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:714
		_go_fuzz_dep_.CoverTab[118923]++
											if s < 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:715
			_go_fuzz_dep_.CoverTab[118948]++
												return 0, errNegativeNotAllowed
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:716
			// _ = "end of CoverTab[118948]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:717
			_go_fuzz_dep_.CoverTab[118949]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:717
			// _ = "end of CoverTab[118949]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:717
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:717
		// _ = "end of CoverTab[118923]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:717
		_go_fuzz_dep_.CoverTab[118924]++
											return uint8(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:718
		// _ = "end of CoverTab[118924]"
	case int32:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:719
		_go_fuzz_dep_.CoverTab[118925]++
											if s < 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:720
			_go_fuzz_dep_.CoverTab[118950]++
												return 0, errNegativeNotAllowed
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:721
			// _ = "end of CoverTab[118950]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:722
			_go_fuzz_dep_.CoverTab[118951]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:722
			// _ = "end of CoverTab[118951]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:722
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:722
		// _ = "end of CoverTab[118925]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:722
		_go_fuzz_dep_.CoverTab[118926]++
											return uint8(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:723
		// _ = "end of CoverTab[118926]"
	case int16:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:724
		_go_fuzz_dep_.CoverTab[118927]++
											if s < 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:725
			_go_fuzz_dep_.CoverTab[118952]++
												return 0, errNegativeNotAllowed
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:726
			// _ = "end of CoverTab[118952]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:727
			_go_fuzz_dep_.CoverTab[118953]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:727
			// _ = "end of CoverTab[118953]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:727
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:727
		// _ = "end of CoverTab[118927]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:727
		_go_fuzz_dep_.CoverTab[118928]++
											return uint8(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:728
		// _ = "end of CoverTab[118928]"
	case int8:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:729
		_go_fuzz_dep_.CoverTab[118929]++
											if s < 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:730
			_go_fuzz_dep_.CoverTab[118954]++
												return 0, errNegativeNotAllowed
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:731
			// _ = "end of CoverTab[118954]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:732
			_go_fuzz_dep_.CoverTab[118955]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:732
			// _ = "end of CoverTab[118955]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:732
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:732
		// _ = "end of CoverTab[118929]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:732
		_go_fuzz_dep_.CoverTab[118930]++
											return uint8(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:733
		// _ = "end of CoverTab[118930]"
	case uint:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:734
		_go_fuzz_dep_.CoverTab[118931]++
											return uint8(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:735
		// _ = "end of CoverTab[118931]"
	case uint64:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:736
		_go_fuzz_dep_.CoverTab[118932]++
											return uint8(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:737
		// _ = "end of CoverTab[118932]"
	case uint32:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:738
		_go_fuzz_dep_.CoverTab[118933]++
											return uint8(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:739
		// _ = "end of CoverTab[118933]"
	case uint16:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:740
		_go_fuzz_dep_.CoverTab[118934]++
											return uint8(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:741
		// _ = "end of CoverTab[118934]"
	case uint8:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:742
		_go_fuzz_dep_.CoverTab[118935]++
											return s, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:743
		// _ = "end of CoverTab[118935]"
	case float64:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:744
		_go_fuzz_dep_.CoverTab[118936]++
											if s < 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:745
			_go_fuzz_dep_.CoverTab[118956]++
												return 0, errNegativeNotAllowed
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:746
			// _ = "end of CoverTab[118956]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:747
			_go_fuzz_dep_.CoverTab[118957]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:747
			// _ = "end of CoverTab[118957]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:747
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:747
		// _ = "end of CoverTab[118936]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:747
		_go_fuzz_dep_.CoverTab[118937]++
											return uint8(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:748
		// _ = "end of CoverTab[118937]"
	case float32:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:749
		_go_fuzz_dep_.CoverTab[118938]++
											if s < 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:750
			_go_fuzz_dep_.CoverTab[118958]++
												return 0, errNegativeNotAllowed
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:751
			// _ = "end of CoverTab[118958]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:752
			_go_fuzz_dep_.CoverTab[118959]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:752
			// _ = "end of CoverTab[118959]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:752
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:752
		// _ = "end of CoverTab[118938]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:752
		_go_fuzz_dep_.CoverTab[118939]++
											return uint8(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:753
		// _ = "end of CoverTab[118939]"
	case bool:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:754
		_go_fuzz_dep_.CoverTab[118940]++
											if s {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:755
			_go_fuzz_dep_.CoverTab[118960]++
												return 1, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:756
			// _ = "end of CoverTab[118960]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:757
			_go_fuzz_dep_.CoverTab[118961]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:757
			// _ = "end of CoverTab[118961]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:757
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:757
		// _ = "end of CoverTab[118940]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:757
		_go_fuzz_dep_.CoverTab[118941]++
											return 0, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:758
		// _ = "end of CoverTab[118941]"
	case nil:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:759
		_go_fuzz_dep_.CoverTab[118942]++
											return 0, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:760
		// _ = "end of CoverTab[118942]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:761
		_go_fuzz_dep_.CoverTab[118943]++
											return 0, fmt.Errorf("unable to cast %#v of type %T to uint8", i, i)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:762
		// _ = "end of CoverTab[118943]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:763
	// _ = "end of CoverTab[118918]"
}

// From html/template/content.go
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:766
// Copyright 2011 The Go Authors. All rights reserved.
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:766
// indirect returns the value, after dereferencing as many times
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:766
// as necessary to reach the base type (or nil).
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:770
func indirect(a interface{}) interface{} {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:770
	_go_fuzz_dep_.CoverTab[118962]++
										if a == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:771
		_go_fuzz_dep_.CoverTab[118966]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:772
		// _ = "end of CoverTab[118966]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:773
		_go_fuzz_dep_.CoverTab[118967]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:773
		// _ = "end of CoverTab[118967]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:773
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:773
	// _ = "end of CoverTab[118962]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:773
	_go_fuzz_dep_.CoverTab[118963]++
										if t := reflect.TypeOf(a); t.Kind() != reflect.Ptr {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:774
		_go_fuzz_dep_.CoverTab[118968]++

											return a
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:776
		// _ = "end of CoverTab[118968]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:777
		_go_fuzz_dep_.CoverTab[118969]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:777
		// _ = "end of CoverTab[118969]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:777
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:777
	// _ = "end of CoverTab[118963]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:777
	_go_fuzz_dep_.CoverTab[118964]++
										v := reflect.ValueOf(a)
										for v.Kind() == reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:779
		_go_fuzz_dep_.CoverTab[118970]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:779
		return !v.IsNil()
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:779
		// _ = "end of CoverTab[118970]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:779
	}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:779
		_go_fuzz_dep_.CoverTab[118971]++
											v = v.Elem()
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:780
		// _ = "end of CoverTab[118971]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:781
	// _ = "end of CoverTab[118964]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:781
	_go_fuzz_dep_.CoverTab[118965]++
										return v.Interface()
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:782
	// _ = "end of CoverTab[118965]"
}

// From html/template/content.go
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:785
// Copyright 2011 The Go Authors. All rights reserved.
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:785
// indirectToStringerOrError returns the value, after dereferencing as many times
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:785
// as necessary to reach the base type (or nil) or an implementation of fmt.Stringer
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:785
// or error,
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:790
func indirectToStringerOrError(a interface{}) interface{} {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:790
	_go_fuzz_dep_.CoverTab[118972]++
										if a == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:791
		_go_fuzz_dep_.CoverTab[118975]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:792
		// _ = "end of CoverTab[118975]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:793
		_go_fuzz_dep_.CoverTab[118976]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:793
		// _ = "end of CoverTab[118976]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:793
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:793
	// _ = "end of CoverTab[118972]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:793
	_go_fuzz_dep_.CoverTab[118973]++

										var errorType = reflect.TypeOf((*error)(nil)).Elem()
										var fmtStringerType = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()

										v := reflect.ValueOf(a)
										for !v.Type().Implements(fmtStringerType) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:799
		_go_fuzz_dep_.CoverTab[118977]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:799
		return !v.Type().Implements(errorType)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:799
		// _ = "end of CoverTab[118977]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:799
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:799
		_go_fuzz_dep_.CoverTab[118978]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:799
		return v.Kind() == reflect.Ptr
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:799
		// _ = "end of CoverTab[118978]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:799
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:799
		_go_fuzz_dep_.CoverTab[118979]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:799
		return !v.IsNil()
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:799
		// _ = "end of CoverTab[118979]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:799
	}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:799
		_go_fuzz_dep_.CoverTab[118980]++
											v = v.Elem()
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:800
		// _ = "end of CoverTab[118980]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:801
	// _ = "end of CoverTab[118973]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:801
	_go_fuzz_dep_.CoverTab[118974]++
										return v.Interface()
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:802
	// _ = "end of CoverTab[118974]"
}

// ToStringE casts an interface to a string type.
func ToStringE(i interface{}) (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:806
	_go_fuzz_dep_.CoverTab[118981]++
										i = indirectToStringerOrError(i)

										switch s := i.(type) {
	case string:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:810
		_go_fuzz_dep_.CoverTab[118982]++
											return s, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:811
		// _ = "end of CoverTab[118982]"
	case bool:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:812
		_go_fuzz_dep_.CoverTab[118983]++
											return strconv.FormatBool(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:813
		// _ = "end of CoverTab[118983]"
	case float64:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:814
		_go_fuzz_dep_.CoverTab[118984]++
											return strconv.FormatFloat(s, 'f', -1, 64), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:815
		// _ = "end of CoverTab[118984]"
	case float32:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:816
		_go_fuzz_dep_.CoverTab[118985]++
											return strconv.FormatFloat(float64(s), 'f', -1, 32), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:817
		// _ = "end of CoverTab[118985]"
	case int:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:818
		_go_fuzz_dep_.CoverTab[118986]++
											return strconv.Itoa(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:819
		// _ = "end of CoverTab[118986]"
	case int64:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:820
		_go_fuzz_dep_.CoverTab[118987]++
											return strconv.FormatInt(s, 10), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:821
		// _ = "end of CoverTab[118987]"
	case int32:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:822
		_go_fuzz_dep_.CoverTab[118988]++
											return strconv.Itoa(int(s)), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:823
		// _ = "end of CoverTab[118988]"
	case int16:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:824
		_go_fuzz_dep_.CoverTab[118989]++
											return strconv.FormatInt(int64(s), 10), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:825
		// _ = "end of CoverTab[118989]"
	case int8:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:826
		_go_fuzz_dep_.CoverTab[118990]++
											return strconv.FormatInt(int64(s), 10), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:827
		// _ = "end of CoverTab[118990]"
	case uint:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:828
		_go_fuzz_dep_.CoverTab[118991]++
											return strconv.FormatUint(uint64(s), 10), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:829
		// _ = "end of CoverTab[118991]"
	case uint64:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:830
		_go_fuzz_dep_.CoverTab[118992]++
											return strconv.FormatUint(uint64(s), 10), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:831
		// _ = "end of CoverTab[118992]"
	case uint32:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:832
		_go_fuzz_dep_.CoverTab[118993]++
											return strconv.FormatUint(uint64(s), 10), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:833
		// _ = "end of CoverTab[118993]"
	case uint16:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:834
		_go_fuzz_dep_.CoverTab[118994]++
											return strconv.FormatUint(uint64(s), 10), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:835
		// _ = "end of CoverTab[118994]"
	case uint8:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:836
		_go_fuzz_dep_.CoverTab[118995]++
											return strconv.FormatUint(uint64(s), 10), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:837
		// _ = "end of CoverTab[118995]"
	case []byte:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:838
		_go_fuzz_dep_.CoverTab[118996]++
											return string(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:839
		// _ = "end of CoverTab[118996]"
	case template.HTML:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:840
		_go_fuzz_dep_.CoverTab[118997]++
											return string(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:841
		// _ = "end of CoverTab[118997]"
	case template.URL:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:842
		_go_fuzz_dep_.CoverTab[118998]++
											return string(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:843
		// _ = "end of CoverTab[118998]"
	case template.JS:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:844
		_go_fuzz_dep_.CoverTab[118999]++
											return string(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:845
		// _ = "end of CoverTab[118999]"
	case template.CSS:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:846
		_go_fuzz_dep_.CoverTab[119000]++
											return string(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:847
		// _ = "end of CoverTab[119000]"
	case template.HTMLAttr:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:848
		_go_fuzz_dep_.CoverTab[119001]++
											return string(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:849
		// _ = "end of CoverTab[119001]"
	case nil:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:850
		_go_fuzz_dep_.CoverTab[119002]++
											return "", nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:851
		// _ = "end of CoverTab[119002]"
	case fmt.Stringer:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:852
		_go_fuzz_dep_.CoverTab[119003]++
											return s.String(), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:853
		// _ = "end of CoverTab[119003]"
	case error:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:854
		_go_fuzz_dep_.CoverTab[119004]++
											return s.Error(), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:855
		// _ = "end of CoverTab[119004]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:856
		_go_fuzz_dep_.CoverTab[119005]++
											return "", fmt.Errorf("unable to cast %#v of type %T to string", i, i)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:857
		// _ = "end of CoverTab[119005]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:858
	// _ = "end of CoverTab[118981]"
}

// ToStringMapStringE casts an interface to a map[string]string type.
func ToStringMapStringE(i interface{}) (map[string]string, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:862
	_go_fuzz_dep_.CoverTab[119006]++
										var m = map[string]string{}

										switch v := i.(type) {
	case map[string]string:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:866
		_go_fuzz_dep_.CoverTab[119007]++
											return v, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:867
		// _ = "end of CoverTab[119007]"
	case map[string]interface{}:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:868
		_go_fuzz_dep_.CoverTab[119008]++
											for k, val := range v {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:869
			_go_fuzz_dep_.CoverTab[119016]++
												m[ToString(k)] = ToString(val)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:870
			// _ = "end of CoverTab[119016]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:871
		// _ = "end of CoverTab[119008]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:871
		_go_fuzz_dep_.CoverTab[119009]++
											return m, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:872
		// _ = "end of CoverTab[119009]"
	case map[interface{}]string:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:873
		_go_fuzz_dep_.CoverTab[119010]++
											for k, val := range v {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:874
			_go_fuzz_dep_.CoverTab[119017]++
												m[ToString(k)] = ToString(val)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:875
			// _ = "end of CoverTab[119017]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:876
		// _ = "end of CoverTab[119010]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:876
		_go_fuzz_dep_.CoverTab[119011]++
											return m, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:877
		// _ = "end of CoverTab[119011]"
	case map[interface{}]interface{}:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:878
		_go_fuzz_dep_.CoverTab[119012]++
											for k, val := range v {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:879
			_go_fuzz_dep_.CoverTab[119018]++
												m[ToString(k)] = ToString(val)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:880
			// _ = "end of CoverTab[119018]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:881
		// _ = "end of CoverTab[119012]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:881
		_go_fuzz_dep_.CoverTab[119013]++
											return m, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:882
		// _ = "end of CoverTab[119013]"
	case string:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:883
		_go_fuzz_dep_.CoverTab[119014]++
											err := jsonStringToObject(v, &m)
											return m, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:885
		// _ = "end of CoverTab[119014]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:886
		_go_fuzz_dep_.CoverTab[119015]++
											return m, fmt.Errorf("unable to cast %#v of type %T to map[string]string", i, i)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:887
		// _ = "end of CoverTab[119015]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:888
	// _ = "end of CoverTab[119006]"
}

// ToStringMapStringSliceE casts an interface to a map[string][]string type.
func ToStringMapStringSliceE(i interface{}) (map[string][]string, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:892
	_go_fuzz_dep_.CoverTab[119019]++
										var m = map[string][]string{}

										switch v := i.(type) {
	case map[string][]string:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:896
		_go_fuzz_dep_.CoverTab[119021]++
											return v, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:897
		// _ = "end of CoverTab[119021]"
	case map[string][]interface{}:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:898
		_go_fuzz_dep_.CoverTab[119022]++
											for k, val := range v {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:899
			_go_fuzz_dep_.CoverTab[119036]++
												m[ToString(k)] = ToStringSlice(val)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:900
			// _ = "end of CoverTab[119036]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:901
		// _ = "end of CoverTab[119022]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:901
		_go_fuzz_dep_.CoverTab[119023]++
											return m, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:902
		// _ = "end of CoverTab[119023]"
	case map[string]string:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:903
		_go_fuzz_dep_.CoverTab[119024]++
											for k, val := range v {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:904
			_go_fuzz_dep_.CoverTab[119037]++
												m[ToString(k)] = []string{val}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:905
			// _ = "end of CoverTab[119037]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:906
		// _ = "end of CoverTab[119024]"
	case map[string]interface{}:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:907
		_go_fuzz_dep_.CoverTab[119025]++
											for k, val := range v {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:908
			_go_fuzz_dep_.CoverTab[119038]++
												switch vt := val.(type) {
			case []interface{}:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:910
				_go_fuzz_dep_.CoverTab[119039]++
													m[ToString(k)] = ToStringSlice(vt)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:911
				// _ = "end of CoverTab[119039]"
			case []string:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:912
				_go_fuzz_dep_.CoverTab[119040]++
													m[ToString(k)] = vt
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:913
				// _ = "end of CoverTab[119040]"
			default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:914
				_go_fuzz_dep_.CoverTab[119041]++
													m[ToString(k)] = []string{ToString(val)}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:915
				// _ = "end of CoverTab[119041]"
			}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:916
			// _ = "end of CoverTab[119038]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:917
		// _ = "end of CoverTab[119025]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:917
		_go_fuzz_dep_.CoverTab[119026]++
											return m, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:918
		// _ = "end of CoverTab[119026]"
	case map[interface{}][]string:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:919
		_go_fuzz_dep_.CoverTab[119027]++
											for k, val := range v {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:920
			_go_fuzz_dep_.CoverTab[119042]++
												m[ToString(k)] = ToStringSlice(val)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:921
			// _ = "end of CoverTab[119042]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:922
		// _ = "end of CoverTab[119027]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:922
		_go_fuzz_dep_.CoverTab[119028]++
											return m, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:923
		// _ = "end of CoverTab[119028]"
	case map[interface{}]string:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:924
		_go_fuzz_dep_.CoverTab[119029]++
											for k, val := range v {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:925
			_go_fuzz_dep_.CoverTab[119043]++
												m[ToString(k)] = ToStringSlice(val)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:926
			// _ = "end of CoverTab[119043]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:927
		// _ = "end of CoverTab[119029]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:927
		_go_fuzz_dep_.CoverTab[119030]++
											return m, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:928
		// _ = "end of CoverTab[119030]"
	case map[interface{}][]interface{}:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:929
		_go_fuzz_dep_.CoverTab[119031]++
											for k, val := range v {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:930
			_go_fuzz_dep_.CoverTab[119044]++
												m[ToString(k)] = ToStringSlice(val)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:931
			// _ = "end of CoverTab[119044]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:932
		// _ = "end of CoverTab[119031]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:932
		_go_fuzz_dep_.CoverTab[119032]++
											return m, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:933
		// _ = "end of CoverTab[119032]"
	case map[interface{}]interface{}:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:934
		_go_fuzz_dep_.CoverTab[119033]++
											for k, val := range v {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:935
			_go_fuzz_dep_.CoverTab[119045]++
												key, err := ToStringE(k)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:937
				_go_fuzz_dep_.CoverTab[119048]++
													return m, fmt.Errorf("unable to cast %#v of type %T to map[string][]string", i, i)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:938
				// _ = "end of CoverTab[119048]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:939
				_go_fuzz_dep_.CoverTab[119049]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:939
				// _ = "end of CoverTab[119049]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:939
			}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:939
			// _ = "end of CoverTab[119045]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:939
			_go_fuzz_dep_.CoverTab[119046]++
												value, err := ToStringSliceE(val)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:941
				_go_fuzz_dep_.CoverTab[119050]++
													return m, fmt.Errorf("unable to cast %#v of type %T to map[string][]string", i, i)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:942
				// _ = "end of CoverTab[119050]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:943
				_go_fuzz_dep_.CoverTab[119051]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:943
				// _ = "end of CoverTab[119051]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:943
			}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:943
			// _ = "end of CoverTab[119046]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:943
			_go_fuzz_dep_.CoverTab[119047]++
												m[key] = value
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:944
			// _ = "end of CoverTab[119047]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:945
		// _ = "end of CoverTab[119033]"
	case string:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:946
		_go_fuzz_dep_.CoverTab[119034]++
											err := jsonStringToObject(v, &m)
											return m, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:948
		// _ = "end of CoverTab[119034]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:949
		_go_fuzz_dep_.CoverTab[119035]++
											return m, fmt.Errorf("unable to cast %#v of type %T to map[string][]string", i, i)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:950
		// _ = "end of CoverTab[119035]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:951
	// _ = "end of CoverTab[119019]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:951
	_go_fuzz_dep_.CoverTab[119020]++
										return m, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:952
	// _ = "end of CoverTab[119020]"
}

// ToStringMapBoolE casts an interface to a map[string]bool type.
func ToStringMapBoolE(i interface{}) (map[string]bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:956
	_go_fuzz_dep_.CoverTab[119052]++
										var m = map[string]bool{}

										switch v := i.(type) {
	case map[interface{}]interface{}:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:960
		_go_fuzz_dep_.CoverTab[119053]++
											for k, val := range v {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:961
			_go_fuzz_dep_.CoverTab[119060]++
												m[ToString(k)] = ToBool(val)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:962
			// _ = "end of CoverTab[119060]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:963
		// _ = "end of CoverTab[119053]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:963
		_go_fuzz_dep_.CoverTab[119054]++
											return m, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:964
		// _ = "end of CoverTab[119054]"
	case map[string]interface{}:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:965
		_go_fuzz_dep_.CoverTab[119055]++
											for k, val := range v {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:966
			_go_fuzz_dep_.CoverTab[119061]++
												m[ToString(k)] = ToBool(val)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:967
			// _ = "end of CoverTab[119061]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:968
		// _ = "end of CoverTab[119055]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:968
		_go_fuzz_dep_.CoverTab[119056]++
											return m, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:969
		// _ = "end of CoverTab[119056]"
	case map[string]bool:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:970
		_go_fuzz_dep_.CoverTab[119057]++
											return v, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:971
		// _ = "end of CoverTab[119057]"
	case string:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:972
		_go_fuzz_dep_.CoverTab[119058]++
											err := jsonStringToObject(v, &m)
											return m, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:974
		// _ = "end of CoverTab[119058]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:975
		_go_fuzz_dep_.CoverTab[119059]++
											return m, fmt.Errorf("unable to cast %#v of type %T to map[string]bool", i, i)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:976
		// _ = "end of CoverTab[119059]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:977
	// _ = "end of CoverTab[119052]"
}

// ToStringMapE casts an interface to a map[string]interface{} type.
func ToStringMapE(i interface{}) (map[string]interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:981
	_go_fuzz_dep_.CoverTab[119062]++
										var m = map[string]interface{}{}

										switch v := i.(type) {
	case map[interface{}]interface{}:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:985
		_go_fuzz_dep_.CoverTab[119063]++
											for k, val := range v {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:986
			_go_fuzz_dep_.CoverTab[119068]++
												m[ToString(k)] = val
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:987
			// _ = "end of CoverTab[119068]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:988
		// _ = "end of CoverTab[119063]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:988
		_go_fuzz_dep_.CoverTab[119064]++
											return m, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:989
		// _ = "end of CoverTab[119064]"
	case map[string]interface{}:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:990
		_go_fuzz_dep_.CoverTab[119065]++
											return v, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:991
		// _ = "end of CoverTab[119065]"
	case string:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:992
		_go_fuzz_dep_.CoverTab[119066]++
											err := jsonStringToObject(v, &m)
											return m, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:994
		// _ = "end of CoverTab[119066]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:995
		_go_fuzz_dep_.CoverTab[119067]++
											return m, fmt.Errorf("unable to cast %#v of type %T to map[string]interface{}", i, i)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:996
		// _ = "end of CoverTab[119067]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:997
	// _ = "end of CoverTab[119062]"
}

// ToStringMapIntE casts an interface to a map[string]int{} type.
func ToStringMapIntE(i interface{}) (map[string]int, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1001
	_go_fuzz_dep_.CoverTab[119069]++
										var m = map[string]int{}
										if i == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1003
		_go_fuzz_dep_.CoverTab[119074]++
											return m, fmt.Errorf("unable to cast %#v of type %T to map[string]int", i, i)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1004
		// _ = "end of CoverTab[119074]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1005
		_go_fuzz_dep_.CoverTab[119075]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1005
		// _ = "end of CoverTab[119075]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1005
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1005
	// _ = "end of CoverTab[119069]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1005
	_go_fuzz_dep_.CoverTab[119070]++

										switch v := i.(type) {
	case map[interface{}]interface{}:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1008
		_go_fuzz_dep_.CoverTab[119076]++
											for k, val := range v {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1009
			_go_fuzz_dep_.CoverTab[119082]++
												m[ToString(k)] = ToInt(val)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1010
			// _ = "end of CoverTab[119082]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1011
		// _ = "end of CoverTab[119076]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1011
		_go_fuzz_dep_.CoverTab[119077]++
											return m, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1012
		// _ = "end of CoverTab[119077]"
	case map[string]interface{}:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1013
		_go_fuzz_dep_.CoverTab[119078]++
											for k, val := range v {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1014
			_go_fuzz_dep_.CoverTab[119083]++
												m[k] = ToInt(val)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1015
			// _ = "end of CoverTab[119083]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1016
		// _ = "end of CoverTab[119078]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1016
		_go_fuzz_dep_.CoverTab[119079]++
											return m, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1017
		// _ = "end of CoverTab[119079]"
	case map[string]int:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1018
		_go_fuzz_dep_.CoverTab[119080]++
											return v, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1019
		// _ = "end of CoverTab[119080]"
	case string:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1020
		_go_fuzz_dep_.CoverTab[119081]++
											err := jsonStringToObject(v, &m)
											return m, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1022
		// _ = "end of CoverTab[119081]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1023
	// _ = "end of CoverTab[119070]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1023
	_go_fuzz_dep_.CoverTab[119071]++

										if reflect.TypeOf(i).Kind() != reflect.Map {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1025
		_go_fuzz_dep_.CoverTab[119084]++
											return m, fmt.Errorf("unable to cast %#v of type %T to map[string]int", i, i)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1026
		// _ = "end of CoverTab[119084]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1027
		_go_fuzz_dep_.CoverTab[119085]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1027
		// _ = "end of CoverTab[119085]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1027
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1027
	// _ = "end of CoverTab[119071]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1027
	_go_fuzz_dep_.CoverTab[119072]++

										mVal := reflect.ValueOf(m)
										v := reflect.ValueOf(i)
										for _, keyVal := range v.MapKeys() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1031
		_go_fuzz_dep_.CoverTab[119086]++
											val, err := ToIntE(v.MapIndex(keyVal).Interface())
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1033
			_go_fuzz_dep_.CoverTab[119088]++
												return m, fmt.Errorf("unable to cast %#v of type %T to map[string]int", i, i)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1034
			// _ = "end of CoverTab[119088]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1035
			_go_fuzz_dep_.CoverTab[119089]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1035
			// _ = "end of CoverTab[119089]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1035
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1035
		// _ = "end of CoverTab[119086]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1035
		_go_fuzz_dep_.CoverTab[119087]++
											mVal.SetMapIndex(keyVal, reflect.ValueOf(val))
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1036
		// _ = "end of CoverTab[119087]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1037
	// _ = "end of CoverTab[119072]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1037
	_go_fuzz_dep_.CoverTab[119073]++
										return m, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1038
	// _ = "end of CoverTab[119073]"
}

// ToStringMapInt64E casts an interface to a map[string]int64{} type.
func ToStringMapInt64E(i interface{}) (map[string]int64, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1042
	_go_fuzz_dep_.CoverTab[119090]++
										var m = map[string]int64{}
										if i == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1044
		_go_fuzz_dep_.CoverTab[119095]++
											return m, fmt.Errorf("unable to cast %#v of type %T to map[string]int64", i, i)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1045
		// _ = "end of CoverTab[119095]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1046
		_go_fuzz_dep_.CoverTab[119096]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1046
		// _ = "end of CoverTab[119096]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1046
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1046
	// _ = "end of CoverTab[119090]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1046
	_go_fuzz_dep_.CoverTab[119091]++

										switch v := i.(type) {
	case map[interface{}]interface{}:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1049
		_go_fuzz_dep_.CoverTab[119097]++
											for k, val := range v {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1050
			_go_fuzz_dep_.CoverTab[119103]++
												m[ToString(k)] = ToInt64(val)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1051
			// _ = "end of CoverTab[119103]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1052
		// _ = "end of CoverTab[119097]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1052
		_go_fuzz_dep_.CoverTab[119098]++
											return m, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1053
		// _ = "end of CoverTab[119098]"
	case map[string]interface{}:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1054
		_go_fuzz_dep_.CoverTab[119099]++
											for k, val := range v {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1055
			_go_fuzz_dep_.CoverTab[119104]++
												m[k] = ToInt64(val)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1056
			// _ = "end of CoverTab[119104]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1057
		// _ = "end of CoverTab[119099]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1057
		_go_fuzz_dep_.CoverTab[119100]++
											return m, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1058
		// _ = "end of CoverTab[119100]"
	case map[string]int64:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1059
		_go_fuzz_dep_.CoverTab[119101]++
											return v, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1060
		// _ = "end of CoverTab[119101]"
	case string:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1061
		_go_fuzz_dep_.CoverTab[119102]++
											err := jsonStringToObject(v, &m)
											return m, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1063
		// _ = "end of CoverTab[119102]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1064
	// _ = "end of CoverTab[119091]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1064
	_go_fuzz_dep_.CoverTab[119092]++

										if reflect.TypeOf(i).Kind() != reflect.Map {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1066
		_go_fuzz_dep_.CoverTab[119105]++
											return m, fmt.Errorf("unable to cast %#v of type %T to map[string]int64", i, i)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1067
		// _ = "end of CoverTab[119105]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1068
		_go_fuzz_dep_.CoverTab[119106]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1068
		// _ = "end of CoverTab[119106]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1068
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1068
	// _ = "end of CoverTab[119092]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1068
	_go_fuzz_dep_.CoverTab[119093]++
										mVal := reflect.ValueOf(m)
										v := reflect.ValueOf(i)
										for _, keyVal := range v.MapKeys() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1071
		_go_fuzz_dep_.CoverTab[119107]++
											val, err := ToInt64E(v.MapIndex(keyVal).Interface())
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1073
			_go_fuzz_dep_.CoverTab[119109]++
												return m, fmt.Errorf("unable to cast %#v of type %T to map[string]int64", i, i)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1074
			// _ = "end of CoverTab[119109]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1075
			_go_fuzz_dep_.CoverTab[119110]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1075
			// _ = "end of CoverTab[119110]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1075
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1075
		// _ = "end of CoverTab[119107]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1075
		_go_fuzz_dep_.CoverTab[119108]++
											mVal.SetMapIndex(keyVal, reflect.ValueOf(val))
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1076
		// _ = "end of CoverTab[119108]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1077
	// _ = "end of CoverTab[119093]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1077
	_go_fuzz_dep_.CoverTab[119094]++
										return m, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1078
	// _ = "end of CoverTab[119094]"
}

// ToSliceE casts an interface to a []interface{} type.
func ToSliceE(i interface{}) ([]interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1082
	_go_fuzz_dep_.CoverTab[119111]++
										var s []interface{}

										switch v := i.(type) {
	case []interface{}:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1086
		_go_fuzz_dep_.CoverTab[119112]++
											return append(s, v...), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1087
		// _ = "end of CoverTab[119112]"
	case []map[string]interface{}:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1088
		_go_fuzz_dep_.CoverTab[119113]++
											for _, u := range v {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1089
			_go_fuzz_dep_.CoverTab[119116]++
												s = append(s, u)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1090
			// _ = "end of CoverTab[119116]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1091
		// _ = "end of CoverTab[119113]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1091
		_go_fuzz_dep_.CoverTab[119114]++
											return s, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1092
		// _ = "end of CoverTab[119114]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1093
		_go_fuzz_dep_.CoverTab[119115]++
											return s, fmt.Errorf("unable to cast %#v of type %T to []interface{}", i, i)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1094
		// _ = "end of CoverTab[119115]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1095
	// _ = "end of CoverTab[119111]"
}

// ToBoolSliceE casts an interface to a []bool type.
func ToBoolSliceE(i interface{}) ([]bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1099
	_go_fuzz_dep_.CoverTab[119117]++
										if i == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1100
		_go_fuzz_dep_.CoverTab[119120]++
											return []bool{}, fmt.Errorf("unable to cast %#v of type %T to []bool", i, i)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1101
		// _ = "end of CoverTab[119120]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1102
		_go_fuzz_dep_.CoverTab[119121]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1102
		// _ = "end of CoverTab[119121]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1102
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1102
	// _ = "end of CoverTab[119117]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1102
	_go_fuzz_dep_.CoverTab[119118]++

										switch v := i.(type) {
	case []bool:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1105
		_go_fuzz_dep_.CoverTab[119122]++
											return v, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1106
		// _ = "end of CoverTab[119122]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1107
	// _ = "end of CoverTab[119118]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1107
	_go_fuzz_dep_.CoverTab[119119]++

										kind := reflect.TypeOf(i).Kind()
										switch kind {
	case reflect.Slice, reflect.Array:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1111
		_go_fuzz_dep_.CoverTab[119123]++
											s := reflect.ValueOf(i)
											a := make([]bool, s.Len())
											for j := 0; j < s.Len(); j++ {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1114
			_go_fuzz_dep_.CoverTab[119126]++
												val, err := ToBoolE(s.Index(j).Interface())
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1116
				_go_fuzz_dep_.CoverTab[119128]++
													return []bool{}, fmt.Errorf("unable to cast %#v of type %T to []bool", i, i)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1117
				// _ = "end of CoverTab[119128]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1118
				_go_fuzz_dep_.CoverTab[119129]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1118
				// _ = "end of CoverTab[119129]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1118
			}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1118
			// _ = "end of CoverTab[119126]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1118
			_go_fuzz_dep_.CoverTab[119127]++
												a[j] = val
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1119
			// _ = "end of CoverTab[119127]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1120
		// _ = "end of CoverTab[119123]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1120
		_go_fuzz_dep_.CoverTab[119124]++
											return a, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1121
		// _ = "end of CoverTab[119124]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1122
		_go_fuzz_dep_.CoverTab[119125]++
											return []bool{}, fmt.Errorf("unable to cast %#v of type %T to []bool", i, i)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1123
		// _ = "end of CoverTab[119125]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1124
	// _ = "end of CoverTab[119119]"
}

// ToStringSliceE casts an interface to a []string type.
func ToStringSliceE(i interface{}) ([]string, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1128
	_go_fuzz_dep_.CoverTab[119130]++
										var a []string

										switch v := i.(type) {
	case []interface{}:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1132
		_go_fuzz_dep_.CoverTab[119131]++
											for _, u := range v {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1133
			_go_fuzz_dep_.CoverTab[119152]++
												a = append(a, ToString(u))
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1134
			// _ = "end of CoverTab[119152]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1135
		// _ = "end of CoverTab[119131]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1135
		_go_fuzz_dep_.CoverTab[119132]++
											return a, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1136
		// _ = "end of CoverTab[119132]"
	case []string:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1137
		_go_fuzz_dep_.CoverTab[119133]++
											return v, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1138
		// _ = "end of CoverTab[119133]"
	case []int8:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1139
		_go_fuzz_dep_.CoverTab[119134]++
											for _, u := range v {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1140
			_go_fuzz_dep_.CoverTab[119153]++
												a = append(a, ToString(u))
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1141
			// _ = "end of CoverTab[119153]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1142
		// _ = "end of CoverTab[119134]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1142
		_go_fuzz_dep_.CoverTab[119135]++
											return a, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1143
		// _ = "end of CoverTab[119135]"
	case []int:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1144
		_go_fuzz_dep_.CoverTab[119136]++
											for _, u := range v {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1145
			_go_fuzz_dep_.CoverTab[119154]++
												a = append(a, ToString(u))
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1146
			// _ = "end of CoverTab[119154]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1147
		// _ = "end of CoverTab[119136]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1147
		_go_fuzz_dep_.CoverTab[119137]++
											return a, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1148
		// _ = "end of CoverTab[119137]"
	case []int32:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1149
		_go_fuzz_dep_.CoverTab[119138]++
											for _, u := range v {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1150
			_go_fuzz_dep_.CoverTab[119155]++
												a = append(a, ToString(u))
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1151
			// _ = "end of CoverTab[119155]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1152
		// _ = "end of CoverTab[119138]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1152
		_go_fuzz_dep_.CoverTab[119139]++
											return a, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1153
		// _ = "end of CoverTab[119139]"
	case []int64:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1154
		_go_fuzz_dep_.CoverTab[119140]++
											for _, u := range v {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1155
			_go_fuzz_dep_.CoverTab[119156]++
												a = append(a, ToString(u))
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1156
			// _ = "end of CoverTab[119156]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1157
		// _ = "end of CoverTab[119140]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1157
		_go_fuzz_dep_.CoverTab[119141]++
											return a, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1158
		// _ = "end of CoverTab[119141]"
	case []float32:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1159
		_go_fuzz_dep_.CoverTab[119142]++
											for _, u := range v {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1160
			_go_fuzz_dep_.CoverTab[119157]++
												a = append(a, ToString(u))
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1161
			// _ = "end of CoverTab[119157]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1162
		// _ = "end of CoverTab[119142]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1162
		_go_fuzz_dep_.CoverTab[119143]++
											return a, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1163
		// _ = "end of CoverTab[119143]"
	case []float64:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1164
		_go_fuzz_dep_.CoverTab[119144]++
											for _, u := range v {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1165
			_go_fuzz_dep_.CoverTab[119158]++
												a = append(a, ToString(u))
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1166
			// _ = "end of CoverTab[119158]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1167
		// _ = "end of CoverTab[119144]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1167
		_go_fuzz_dep_.CoverTab[119145]++
											return a, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1168
		// _ = "end of CoverTab[119145]"
	case string:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1169
		_go_fuzz_dep_.CoverTab[119146]++
											return strings.Fields(v), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1170
		// _ = "end of CoverTab[119146]"
	case []error:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1171
		_go_fuzz_dep_.CoverTab[119147]++
											for _, err := range i.([]error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1172
			_go_fuzz_dep_.CoverTab[119159]++
												a = append(a, err.Error())
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1173
			// _ = "end of CoverTab[119159]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1174
		// _ = "end of CoverTab[119147]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1174
		_go_fuzz_dep_.CoverTab[119148]++
											return a, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1175
		// _ = "end of CoverTab[119148]"
	case interface{}:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1176
		_go_fuzz_dep_.CoverTab[119149]++
											str, err := ToStringE(v)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1178
			_go_fuzz_dep_.CoverTab[119160]++
												return a, fmt.Errorf("unable to cast %#v of type %T to []string", i, i)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1179
			// _ = "end of CoverTab[119160]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1180
			_go_fuzz_dep_.CoverTab[119161]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1180
			// _ = "end of CoverTab[119161]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1180
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1180
		// _ = "end of CoverTab[119149]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1180
		_go_fuzz_dep_.CoverTab[119150]++
											return []string{str}, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1181
		// _ = "end of CoverTab[119150]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1182
		_go_fuzz_dep_.CoverTab[119151]++
											return a, fmt.Errorf("unable to cast %#v of type %T to []string", i, i)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1183
		// _ = "end of CoverTab[119151]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1184
	// _ = "end of CoverTab[119130]"
}

// ToIntSliceE casts an interface to a []int type.
func ToIntSliceE(i interface{}) ([]int, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1188
	_go_fuzz_dep_.CoverTab[119162]++
										if i == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1189
		_go_fuzz_dep_.CoverTab[119165]++
											return []int{}, fmt.Errorf("unable to cast %#v of type %T to []int", i, i)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1190
		// _ = "end of CoverTab[119165]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1191
		_go_fuzz_dep_.CoverTab[119166]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1191
		// _ = "end of CoverTab[119166]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1191
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1191
	// _ = "end of CoverTab[119162]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1191
	_go_fuzz_dep_.CoverTab[119163]++

										switch v := i.(type) {
	case []int:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1194
		_go_fuzz_dep_.CoverTab[119167]++
											return v, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1195
		// _ = "end of CoverTab[119167]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1196
	// _ = "end of CoverTab[119163]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1196
	_go_fuzz_dep_.CoverTab[119164]++

										kind := reflect.TypeOf(i).Kind()
										switch kind {
	case reflect.Slice, reflect.Array:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1200
		_go_fuzz_dep_.CoverTab[119168]++
											s := reflect.ValueOf(i)
											a := make([]int, s.Len())
											for j := 0; j < s.Len(); j++ {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1203
			_go_fuzz_dep_.CoverTab[119171]++
												val, err := ToIntE(s.Index(j).Interface())
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1205
				_go_fuzz_dep_.CoverTab[119173]++
													return []int{}, fmt.Errorf("unable to cast %#v of type %T to []int", i, i)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1206
				// _ = "end of CoverTab[119173]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1207
				_go_fuzz_dep_.CoverTab[119174]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1207
				// _ = "end of CoverTab[119174]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1207
			}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1207
			// _ = "end of CoverTab[119171]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1207
			_go_fuzz_dep_.CoverTab[119172]++
												a[j] = val
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1208
			// _ = "end of CoverTab[119172]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1209
		// _ = "end of CoverTab[119168]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1209
		_go_fuzz_dep_.CoverTab[119169]++
											return a, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1210
		// _ = "end of CoverTab[119169]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1211
		_go_fuzz_dep_.CoverTab[119170]++
											return []int{}, fmt.Errorf("unable to cast %#v of type %T to []int", i, i)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1212
		// _ = "end of CoverTab[119170]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1213
	// _ = "end of CoverTab[119164]"
}

// ToDurationSliceE casts an interface to a []time.Duration type.
func ToDurationSliceE(i interface{}) ([]time.Duration, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1217
	_go_fuzz_dep_.CoverTab[119175]++
										if i == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1218
		_go_fuzz_dep_.CoverTab[119178]++
											return []time.Duration{}, fmt.Errorf("unable to cast %#v of type %T to []time.Duration", i, i)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1219
		// _ = "end of CoverTab[119178]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1220
		_go_fuzz_dep_.CoverTab[119179]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1220
		// _ = "end of CoverTab[119179]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1220
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1220
	// _ = "end of CoverTab[119175]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1220
	_go_fuzz_dep_.CoverTab[119176]++

										switch v := i.(type) {
	case []time.Duration:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1223
		_go_fuzz_dep_.CoverTab[119180]++
											return v, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1224
		// _ = "end of CoverTab[119180]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1225
	// _ = "end of CoverTab[119176]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1225
	_go_fuzz_dep_.CoverTab[119177]++

										kind := reflect.TypeOf(i).Kind()
										switch kind {
	case reflect.Slice, reflect.Array:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1229
		_go_fuzz_dep_.CoverTab[119181]++
											s := reflect.ValueOf(i)
											a := make([]time.Duration, s.Len())
											for j := 0; j < s.Len(); j++ {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1232
			_go_fuzz_dep_.CoverTab[119184]++
												val, err := ToDurationE(s.Index(j).Interface())
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1234
				_go_fuzz_dep_.CoverTab[119186]++
													return []time.Duration{}, fmt.Errorf("unable to cast %#v of type %T to []time.Duration", i, i)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1235
				// _ = "end of CoverTab[119186]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1236
				_go_fuzz_dep_.CoverTab[119187]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1236
				// _ = "end of CoverTab[119187]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1236
			}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1236
			// _ = "end of CoverTab[119184]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1236
			_go_fuzz_dep_.CoverTab[119185]++
												a[j] = val
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1237
			// _ = "end of CoverTab[119185]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1238
		// _ = "end of CoverTab[119181]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1238
		_go_fuzz_dep_.CoverTab[119182]++
											return a, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1239
		// _ = "end of CoverTab[119182]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1240
		_go_fuzz_dep_.CoverTab[119183]++
											return []time.Duration{}, fmt.Errorf("unable to cast %#v of type %T to []time.Duration", i, i)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1241
		// _ = "end of CoverTab[119183]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1242
	// _ = "end of CoverTab[119177]"
}

// StringToDate attempts to parse a string into a time.Time type using a
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1245
// predefined list of formats.  If no suitable format is found, an error is
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1245
// returned.
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1248
func StringToDate(s string) (time.Time, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1248
	_go_fuzz_dep_.CoverTab[119188]++
										return parseDateWith(s, time.UTC, timeFormats)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1249
	// _ = "end of CoverTab[119188]"
}

// StringToDateInDefaultLocation casts an empty interface to a time.Time,
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1252
// interpreting inputs without a timezone to be in the given location,
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1252
// or the local timezone if nil.
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1255
func StringToDateInDefaultLocation(s string, location *time.Location) (time.Time, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1255
	_go_fuzz_dep_.CoverTab[119189]++
										return parseDateWith(s, location, timeFormats)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1256
	// _ = "end of CoverTab[119189]"
}

type timeFormatType int

const (
	timeFormatNoTimezone	timeFormatType	= iota
	timeFormatNamedTimezone
	timeFormatNumericTimezone
	timeFormatNumericAndNamedTimezone
	timeFormatTimeOnly
)

type timeFormat struct {
	format	string
	typ	timeFormatType
}

func (f timeFormat) hasTimezone() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1274
	_go_fuzz_dep_.CoverTab[119190]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1277
	return f.typ >= timeFormatNumericTimezone && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1277
		_go_fuzz_dep_.CoverTab[119191]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1277
		return f.typ <= timeFormatNumericAndNamedTimezone
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1277
		// _ = "end of CoverTab[119191]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1277
	}()
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1277
	// _ = "end of CoverTab[119190]"
}

var (
	timeFormats = []timeFormat{
		timeFormat{time.RFC3339, timeFormatNumericTimezone},
		timeFormat{"2006-01-02T15:04:05", timeFormatNoTimezone},
		timeFormat{time.RFC1123Z, timeFormatNumericTimezone},
		timeFormat{time.RFC1123, timeFormatNamedTimezone},
		timeFormat{time.RFC822Z, timeFormatNumericTimezone},
		timeFormat{time.RFC822, timeFormatNamedTimezone},
		timeFormat{time.RFC850, timeFormatNamedTimezone},
		timeFormat{"2006-01-02 15:04:05.999999999 -0700 MST", timeFormatNumericAndNamedTimezone},
		timeFormat{"2006-01-02T15:04:05-0700", timeFormatNumericTimezone},
		timeFormat{"2006-01-02 15:04:05Z0700", timeFormatNumericTimezone},
		timeFormat{"2006-01-02 15:04:05", timeFormatNoTimezone},
		timeFormat{time.ANSIC, timeFormatNoTimezone},
		timeFormat{time.UnixDate, timeFormatNamedTimezone},
		timeFormat{time.RubyDate, timeFormatNumericTimezone},
		timeFormat{"2006-01-02 15:04:05Z07:00", timeFormatNumericTimezone},
		timeFormat{"2006-01-02", timeFormatNoTimezone},
		timeFormat{"02 Jan 2006", timeFormatNoTimezone},
		timeFormat{"2006-01-02 15:04:05 -07:00", timeFormatNumericTimezone},
		timeFormat{"2006-01-02 15:04:05 -0700", timeFormatNumericTimezone},
		timeFormat{time.Kitchen, timeFormatTimeOnly},
		timeFormat{time.Stamp, timeFormatTimeOnly},
		timeFormat{time.StampMilli, timeFormatTimeOnly},
		timeFormat{time.StampMicro, timeFormatTimeOnly},
		timeFormat{time.StampNano, timeFormatTimeOnly},
	}
)

func parseDateWith(s string, location *time.Location, formats []timeFormat) (d time.Time, e error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1309
	_go_fuzz_dep_.CoverTab[119192]++

										for _, format := range formats {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1311
		_go_fuzz_dep_.CoverTab[119194]++
											if d, e = time.Parse(format.format, s); e == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1312
			_go_fuzz_dep_.CoverTab[119195]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1317
			if format.typ <= timeFormatNamedTimezone {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1317
				_go_fuzz_dep_.CoverTab[119197]++
													if location == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1318
					_go_fuzz_dep_.CoverTab[119199]++
														location = time.Local
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1319
					// _ = "end of CoverTab[119199]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1320
					_go_fuzz_dep_.CoverTab[119200]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1320
					// _ = "end of CoverTab[119200]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1320
				}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1320
				// _ = "end of CoverTab[119197]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1320
				_go_fuzz_dep_.CoverTab[119198]++
													year, month, day := d.Date()
													hour, min, sec := d.Clock()
													d = time.Date(year, month, day, hour, min, sec, d.Nanosecond(), location)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1323
				// _ = "end of CoverTab[119198]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1324
				_go_fuzz_dep_.CoverTab[119201]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1324
				// _ = "end of CoverTab[119201]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1324
			}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1324
			// _ = "end of CoverTab[119195]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1324
			_go_fuzz_dep_.CoverTab[119196]++

												return
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1326
			// _ = "end of CoverTab[119196]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1327
			_go_fuzz_dep_.CoverTab[119202]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1327
			// _ = "end of CoverTab[119202]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1327
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1327
		// _ = "end of CoverTab[119194]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1328
	// _ = "end of CoverTab[119192]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1328
	_go_fuzz_dep_.CoverTab[119193]++
										return d, fmt.Errorf("unable to parse date: %s", s)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1329
	// _ = "end of CoverTab[119193]"
}

// jsonStringToObject attempts to unmarshall a string as JSON into
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1332
// the object passed as pointer.
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1334
func jsonStringToObject(s string, v interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1334
	_go_fuzz_dep_.CoverTab[119203]++
										data := []byte(s)
										return json.Unmarshal(data, v)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1336
	// _ = "end of CoverTab[119203]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1337
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/caste.go:1337
var _ = _go_fuzz_dep_.CoverTab
