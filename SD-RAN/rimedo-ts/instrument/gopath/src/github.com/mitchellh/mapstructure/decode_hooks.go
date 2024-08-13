//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:1
package mapstructure

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:1
)

import (
	"encoding"
	"errors"
	"fmt"
	"net"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// typedDecodeHook takes a raw DecodeHookFunc (an interface{}) and turns
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:14
// it into the proper DecodeHookFunc type, such as DecodeHookFuncType.
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:16
func typedDecodeHook(h DecodeHookFunc) DecodeHookFunc {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:16
	_go_fuzz_dep_.CoverTab[116174]++
												// Create variables here so we can reference them with the reflect pkg
												var f1 DecodeHookFuncType
												var f2 DecodeHookFuncKind
												var f3 DecodeHookFuncValue

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:24
	potential := []interface{}{f1, f2, f3}

	v := reflect.ValueOf(h)
	vt := v.Type()
	for _, raw := range potential {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:28
		_go_fuzz_dep_.CoverTab[116176]++
													pt := reflect.ValueOf(raw).Type()
													if vt.ConvertibleTo(pt) {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:30
			_go_fuzz_dep_.CoverTab[116177]++
														return v.Convert(pt).Interface()
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:31
			// _ = "end of CoverTab[116177]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:32
			_go_fuzz_dep_.CoverTab[116178]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:32
			// _ = "end of CoverTab[116178]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:32
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:32
		// _ = "end of CoverTab[116176]"
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:33
	// _ = "end of CoverTab[116174]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:33
	_go_fuzz_dep_.CoverTab[116175]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:35
	// _ = "end of CoverTab[116175]"
}

// DecodeHookExec executes the given decode hook. This should be used
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:38
// since it'll naturally degrade to the older backwards compatible DecodeHookFunc
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:38
// that took reflect.Kind instead of reflect.Type.
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:41
func DecodeHookExec(
	raw DecodeHookFunc,
	from reflect.Value, to reflect.Value) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:43
	_go_fuzz_dep_.CoverTab[116179]++

												switch f := typedDecodeHook(raw).(type) {
	case DecodeHookFuncType:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:46
		_go_fuzz_dep_.CoverTab[116180]++
													return f(from.Type(), to.Type(), from.Interface())
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:47
		// _ = "end of CoverTab[116180]"
	case DecodeHookFuncKind:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:48
		_go_fuzz_dep_.CoverTab[116181]++
													return f(from.Kind(), to.Kind(), from.Interface())
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:49
		// _ = "end of CoverTab[116181]"
	case DecodeHookFuncValue:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:50
		_go_fuzz_dep_.CoverTab[116182]++
													return f(from, to)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:51
		// _ = "end of CoverTab[116182]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:52
		_go_fuzz_dep_.CoverTab[116183]++
													return nil, errors.New("invalid decode hook signature")
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:53
		// _ = "end of CoverTab[116183]"
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:54
	// _ = "end of CoverTab[116179]"
}

// ComposeDecodeHookFunc creates a single DecodeHookFunc that
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:57
// automatically composes multiple DecodeHookFuncs.
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:57
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:57
// The composed funcs are called in order, with the result of the
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:57
// previous transformation.
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:62
func ComposeDecodeHookFunc(fs ...DecodeHookFunc) DecodeHookFunc {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:62
	_go_fuzz_dep_.CoverTab[116184]++
												return func(f reflect.Value, t reflect.Value) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:63
		_go_fuzz_dep_.CoverTab[116185]++
													var err error
													data := f.Interface()

													newFrom := f
													for _, f1 := range fs {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:68
			_go_fuzz_dep_.CoverTab[116187]++
														data, err = DecodeHookExec(f1, newFrom, t)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:70
				_go_fuzz_dep_.CoverTab[116189]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:71
				// _ = "end of CoverTab[116189]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:72
				_go_fuzz_dep_.CoverTab[116190]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:72
				// _ = "end of CoverTab[116190]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:72
			}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:72
			// _ = "end of CoverTab[116187]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:72
			_go_fuzz_dep_.CoverTab[116188]++
														newFrom = reflect.ValueOf(data)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:73
			// _ = "end of CoverTab[116188]"
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:74
		// _ = "end of CoverTab[116185]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:74
		_go_fuzz_dep_.CoverTab[116186]++

													return data, nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:76
		// _ = "end of CoverTab[116186]"
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:77
	// _ = "end of CoverTab[116184]"
}

// StringToSliceHookFunc returns a DecodeHookFunc that converts
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:80
// string to []string by splitting on the given sep.
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:82
func StringToSliceHookFunc(sep string) DecodeHookFunc {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:82
	_go_fuzz_dep_.CoverTab[116191]++
												return func(
		f reflect.Kind,
		t reflect.Kind,
		data interface{}) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:86
		_go_fuzz_dep_.CoverTab[116192]++
													if f != reflect.String || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:87
			_go_fuzz_dep_.CoverTab[116195]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:87
			return t != reflect.Slice
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:87
			// _ = "end of CoverTab[116195]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:87
		}() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:87
			_go_fuzz_dep_.CoverTab[116196]++
														return data, nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:88
			// _ = "end of CoverTab[116196]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:89
			_go_fuzz_dep_.CoverTab[116197]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:89
			// _ = "end of CoverTab[116197]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:89
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:89
		// _ = "end of CoverTab[116192]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:89
		_go_fuzz_dep_.CoverTab[116193]++

													raw := data.(string)
													if raw == "" {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:92
			_go_fuzz_dep_.CoverTab[116198]++
														return []string{}, nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:93
			// _ = "end of CoverTab[116198]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:94
			_go_fuzz_dep_.CoverTab[116199]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:94
			// _ = "end of CoverTab[116199]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:94
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:94
		// _ = "end of CoverTab[116193]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:94
		_go_fuzz_dep_.CoverTab[116194]++

													return strings.Split(raw, sep), nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:96
		// _ = "end of CoverTab[116194]"
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:97
	// _ = "end of CoverTab[116191]"
}

// StringToTimeDurationHookFunc returns a DecodeHookFunc that converts
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:100
// strings to time.Duration.
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:102
func StringToTimeDurationHookFunc() DecodeHookFunc {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:102
	_go_fuzz_dep_.CoverTab[116200]++
												return func(
		f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:106
		_go_fuzz_dep_.CoverTab[116201]++
													if f.Kind() != reflect.String {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:107
			_go_fuzz_dep_.CoverTab[116204]++
														return data, nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:108
			// _ = "end of CoverTab[116204]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:109
			_go_fuzz_dep_.CoverTab[116205]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:109
			// _ = "end of CoverTab[116205]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:109
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:109
		// _ = "end of CoverTab[116201]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:109
		_go_fuzz_dep_.CoverTab[116202]++
													if t != reflect.TypeOf(time.Duration(5)) {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:110
			_go_fuzz_dep_.CoverTab[116206]++
														return data, nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:111
			// _ = "end of CoverTab[116206]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:112
			_go_fuzz_dep_.CoverTab[116207]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:112
			// _ = "end of CoverTab[116207]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:112
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:112
		// _ = "end of CoverTab[116202]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:112
		_go_fuzz_dep_.CoverTab[116203]++

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:115
		return time.ParseDuration(data.(string))
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:115
		// _ = "end of CoverTab[116203]"
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:116
	// _ = "end of CoverTab[116200]"
}

// StringToIPHookFunc returns a DecodeHookFunc that converts
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:119
// strings to net.IP
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:121
func StringToIPHookFunc() DecodeHookFunc {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:121
	_go_fuzz_dep_.CoverTab[116208]++
												return func(
		f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:125
		_go_fuzz_dep_.CoverTab[116209]++
													if f.Kind() != reflect.String {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:126
			_go_fuzz_dep_.CoverTab[116213]++
														return data, nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:127
			// _ = "end of CoverTab[116213]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:128
			_go_fuzz_dep_.CoverTab[116214]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:128
			// _ = "end of CoverTab[116214]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:128
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:128
		// _ = "end of CoverTab[116209]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:128
		_go_fuzz_dep_.CoverTab[116210]++
													if t != reflect.TypeOf(net.IP{}) {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:129
			_go_fuzz_dep_.CoverTab[116215]++
														return data, nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:130
			// _ = "end of CoverTab[116215]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:131
			_go_fuzz_dep_.CoverTab[116216]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:131
			// _ = "end of CoverTab[116216]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:131
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:131
		// _ = "end of CoverTab[116210]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:131
		_go_fuzz_dep_.CoverTab[116211]++

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:134
		ip := net.ParseIP(data.(string))
		if ip == nil {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:135
			_go_fuzz_dep_.CoverTab[116217]++
														return net.IP{}, fmt.Errorf("failed parsing ip %v", data)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:136
			// _ = "end of CoverTab[116217]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:137
			_go_fuzz_dep_.CoverTab[116218]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:137
			// _ = "end of CoverTab[116218]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:137
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:137
		// _ = "end of CoverTab[116211]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:137
		_go_fuzz_dep_.CoverTab[116212]++

													return ip, nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:139
		// _ = "end of CoverTab[116212]"
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:140
	// _ = "end of CoverTab[116208]"
}

// StringToIPNetHookFunc returns a DecodeHookFunc that converts
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:143
// strings to net.IPNet
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:145
func StringToIPNetHookFunc() DecodeHookFunc {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:145
	_go_fuzz_dep_.CoverTab[116219]++
												return func(
		f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:149
		_go_fuzz_dep_.CoverTab[116220]++
													if f.Kind() != reflect.String {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:150
			_go_fuzz_dep_.CoverTab[116223]++
														return data, nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:151
			// _ = "end of CoverTab[116223]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:152
			_go_fuzz_dep_.CoverTab[116224]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:152
			// _ = "end of CoverTab[116224]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:152
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:152
		// _ = "end of CoverTab[116220]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:152
		_go_fuzz_dep_.CoverTab[116221]++
													if t != reflect.TypeOf(net.IPNet{}) {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:153
			_go_fuzz_dep_.CoverTab[116225]++
														return data, nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:154
			// _ = "end of CoverTab[116225]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:155
			_go_fuzz_dep_.CoverTab[116226]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:155
			// _ = "end of CoverTab[116226]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:155
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:155
		// _ = "end of CoverTab[116221]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:155
		_go_fuzz_dep_.CoverTab[116222]++

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:158
		_, net, err := net.ParseCIDR(data.(string))
													return net, err
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:159
		// _ = "end of CoverTab[116222]"
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:160
	// _ = "end of CoverTab[116219]"
}

// StringToTimeHookFunc returns a DecodeHookFunc that converts
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:163
// strings to time.Time.
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:165
func StringToTimeHookFunc(layout string) DecodeHookFunc {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:165
	_go_fuzz_dep_.CoverTab[116227]++
												return func(
		f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:169
		_go_fuzz_dep_.CoverTab[116228]++
													if f.Kind() != reflect.String {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:170
			_go_fuzz_dep_.CoverTab[116231]++
														return data, nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:171
			// _ = "end of CoverTab[116231]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:172
			_go_fuzz_dep_.CoverTab[116232]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:172
			// _ = "end of CoverTab[116232]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:172
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:172
		// _ = "end of CoverTab[116228]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:172
		_go_fuzz_dep_.CoverTab[116229]++
													if t != reflect.TypeOf(time.Time{}) {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:173
			_go_fuzz_dep_.CoverTab[116233]++
														return data, nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:174
			// _ = "end of CoverTab[116233]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:175
			_go_fuzz_dep_.CoverTab[116234]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:175
			// _ = "end of CoverTab[116234]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:175
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:175
		// _ = "end of CoverTab[116229]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:175
		_go_fuzz_dep_.CoverTab[116230]++

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:178
		return time.Parse(layout, data.(string))
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:178
		// _ = "end of CoverTab[116230]"
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:179
	// _ = "end of CoverTab[116227]"
}

// WeaklyTypedHook is a DecodeHookFunc which adds support for weak typing to
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:182
// the decoder.
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:182
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:182
// Note that this is significantly different from the WeaklyTypedInput option
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:182
// of the DecoderConfig.
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:187
func WeaklyTypedHook(
	f reflect.Kind,
	t reflect.Kind,
	data interface{}) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:190
	_go_fuzz_dep_.CoverTab[116235]++
												dataVal := reflect.ValueOf(data)
												switch t {
	case reflect.String:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:193
		_go_fuzz_dep_.CoverTab[116237]++
													switch f {
		case reflect.Bool:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:195
			_go_fuzz_dep_.CoverTab[116239]++
														if dataVal.Bool() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:196
				_go_fuzz_dep_.CoverTab[116246]++
															return "1", nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:197
				// _ = "end of CoverTab[116246]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:198
				_go_fuzz_dep_.CoverTab[116247]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:198
				// _ = "end of CoverTab[116247]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:198
			}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:198
			// _ = "end of CoverTab[116239]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:198
			_go_fuzz_dep_.CoverTab[116240]++
														return "0", nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:199
			// _ = "end of CoverTab[116240]"
		case reflect.Float32:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:200
			_go_fuzz_dep_.CoverTab[116241]++
														return strconv.FormatFloat(dataVal.Float(), 'f', -1, 64), nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:201
			// _ = "end of CoverTab[116241]"
		case reflect.Int:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:202
			_go_fuzz_dep_.CoverTab[116242]++
														return strconv.FormatInt(dataVal.Int(), 10), nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:203
			// _ = "end of CoverTab[116242]"
		case reflect.Slice:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:204
			_go_fuzz_dep_.CoverTab[116243]++
														dataType := dataVal.Type()
														elemKind := dataType.Elem().Kind()
														if elemKind == reflect.Uint8 {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:207
				_go_fuzz_dep_.CoverTab[116248]++
															return string(dataVal.Interface().([]uint8)), nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:208
				// _ = "end of CoverTab[116248]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:209
				_go_fuzz_dep_.CoverTab[116249]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:209
				// _ = "end of CoverTab[116249]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:209
			}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:209
			// _ = "end of CoverTab[116243]"
		case reflect.Uint:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:210
			_go_fuzz_dep_.CoverTab[116244]++
														return strconv.FormatUint(dataVal.Uint(), 10), nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:211
			// _ = "end of CoverTab[116244]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:211
		default:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:211
			_go_fuzz_dep_.CoverTab[116245]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:211
			// _ = "end of CoverTab[116245]"
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:212
		// _ = "end of CoverTab[116237]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:212
	default:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:212
		_go_fuzz_dep_.CoverTab[116238]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:212
		// _ = "end of CoverTab[116238]"
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:213
	// _ = "end of CoverTab[116235]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:213
	_go_fuzz_dep_.CoverTab[116236]++

												return data, nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:215
	// _ = "end of CoverTab[116236]"
}

func RecursiveStructToMapHookFunc() DecodeHookFunc {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:218
	_go_fuzz_dep_.CoverTab[116250]++
												return func(f reflect.Value, t reflect.Value) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:219
		_go_fuzz_dep_.CoverTab[116251]++
													if f.Kind() != reflect.Struct {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:220
			_go_fuzz_dep_.CoverTab[116254]++
														return f.Interface(), nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:221
			// _ = "end of CoverTab[116254]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:222
			_go_fuzz_dep_.CoverTab[116255]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:222
			// _ = "end of CoverTab[116255]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:222
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:222
		// _ = "end of CoverTab[116251]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:222
		_go_fuzz_dep_.CoverTab[116252]++

													var i interface{} = struct{}{}
													if t.Type() != reflect.TypeOf(&i).Elem() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:225
			_go_fuzz_dep_.CoverTab[116256]++
														return f.Interface(), nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:226
			// _ = "end of CoverTab[116256]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:227
			_go_fuzz_dep_.CoverTab[116257]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:227
			// _ = "end of CoverTab[116257]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:227
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:227
		// _ = "end of CoverTab[116252]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:227
		_go_fuzz_dep_.CoverTab[116253]++

													m := make(map[string]interface{})
													t.Set(reflect.ValueOf(m))

													return f.Interface(), nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:232
		// _ = "end of CoverTab[116253]"
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:233
	// _ = "end of CoverTab[116250]"
}

// TextUnmarshallerHookFunc returns a DecodeHookFunc that applies
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:236
// strings to the UnmarshalText function, when the target type
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:236
// implements the encoding.TextUnmarshaler interface
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:239
func TextUnmarshallerHookFunc() DecodeHookFuncType {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:239
	_go_fuzz_dep_.CoverTab[116258]++
												return func(
		f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:243
		_go_fuzz_dep_.CoverTab[116259]++
													if f.Kind() != reflect.String {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:244
			_go_fuzz_dep_.CoverTab[116263]++
														return data, nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:245
			// _ = "end of CoverTab[116263]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:246
			_go_fuzz_dep_.CoverTab[116264]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:246
			// _ = "end of CoverTab[116264]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:246
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:246
		// _ = "end of CoverTab[116259]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:246
		_go_fuzz_dep_.CoverTab[116260]++
													result := reflect.New(t).Interface()
													unmarshaller, ok := result.(encoding.TextUnmarshaler)
													if !ok {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:249
			_go_fuzz_dep_.CoverTab[116265]++
														return data, nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:250
			// _ = "end of CoverTab[116265]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:251
			_go_fuzz_dep_.CoverTab[116266]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:251
			// _ = "end of CoverTab[116266]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:251
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:251
		// _ = "end of CoverTab[116260]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:251
		_go_fuzz_dep_.CoverTab[116261]++
													if err := unmarshaller.UnmarshalText([]byte(data.(string))); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:252
			_go_fuzz_dep_.CoverTab[116267]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:253
			// _ = "end of CoverTab[116267]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:254
			_go_fuzz_dep_.CoverTab[116268]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:254
			// _ = "end of CoverTab[116268]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:254
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:254
		// _ = "end of CoverTab[116261]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:254
		_go_fuzz_dep_.CoverTab[116262]++
													return result, nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:255
		// _ = "end of CoverTab[116262]"
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:256
	// _ = "end of CoverTab[116258]"
}

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:257
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/decode_hooks.go:257
var _ = _go_fuzz_dep_.CoverTab
