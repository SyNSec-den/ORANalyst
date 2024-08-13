//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:1
package pflag

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:1
)

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"strings"
)

// -- stringToString Value
type stringToStringValue struct {
	value	*map[string]string
	changed	bool
}

func newStringToStringValue(val map[string]string, p *map[string]string) *stringToStringValue {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:16
	_go_fuzz_dep_.CoverTab[120582]++
											ssv := new(stringToStringValue)
											ssv.value = p
											*ssv.value = val
											return ssv
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:20
	// _ = "end of CoverTab[120582]"
}

// Format: a=1,b=2
func (s *stringToStringValue) Set(val string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:24
	_go_fuzz_dep_.CoverTab[120583]++
											var ss []string
											n := strings.Count(val, "=")
											switch n {
	case 0:
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:28
		_go_fuzz_dep_.CoverTab[120587]++
												return fmt.Errorf("%s must be formatted as key=value", val)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:29
		// _ = "end of CoverTab[120587]"
	case 1:
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:30
		_go_fuzz_dep_.CoverTab[120588]++
												ss = append(ss, strings.Trim(val, `"`))
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:31
		// _ = "end of CoverTab[120588]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:32
		_go_fuzz_dep_.CoverTab[120589]++
												r := csv.NewReader(strings.NewReader(val))
												var err error
												ss, err = r.Read()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:36
			_go_fuzz_dep_.CoverTab[120590]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:37
			// _ = "end of CoverTab[120590]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:38
			_go_fuzz_dep_.CoverTab[120591]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:38
			// _ = "end of CoverTab[120591]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:38
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:38
		// _ = "end of CoverTab[120589]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:39
	// _ = "end of CoverTab[120583]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:39
	_go_fuzz_dep_.CoverTab[120584]++

											out := make(map[string]string, len(ss))
											for _, pair := range ss {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:42
		_go_fuzz_dep_.CoverTab[120592]++
												kv := strings.SplitN(pair, "=", 2)
												if len(kv) != 2 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:44
			_go_fuzz_dep_.CoverTab[120594]++
													return fmt.Errorf("%s must be formatted as key=value", pair)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:45
			// _ = "end of CoverTab[120594]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:46
			_go_fuzz_dep_.CoverTab[120595]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:46
			// _ = "end of CoverTab[120595]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:46
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:46
		// _ = "end of CoverTab[120592]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:46
		_go_fuzz_dep_.CoverTab[120593]++
												out[kv[0]] = kv[1]
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:47
		// _ = "end of CoverTab[120593]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:48
	// _ = "end of CoverTab[120584]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:48
	_go_fuzz_dep_.CoverTab[120585]++
											if !s.changed {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:49
		_go_fuzz_dep_.CoverTab[120596]++
												*s.value = out
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:50
		// _ = "end of CoverTab[120596]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:51
		_go_fuzz_dep_.CoverTab[120597]++
												for k, v := range out {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:52
			_go_fuzz_dep_.CoverTab[120598]++
													(*s.value)[k] = v
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:53
			// _ = "end of CoverTab[120598]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:54
		// _ = "end of CoverTab[120597]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:55
	// _ = "end of CoverTab[120585]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:55
	_go_fuzz_dep_.CoverTab[120586]++
											s.changed = true
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:57
	// _ = "end of CoverTab[120586]"
}

func (s *stringToStringValue) Type() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:60
	_go_fuzz_dep_.CoverTab[120599]++
											return "stringToString"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:61
	// _ = "end of CoverTab[120599]"
}

func (s *stringToStringValue) String() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:64
	_go_fuzz_dep_.CoverTab[120600]++
											records := make([]string, 0, len(*s.value)>>1)
											for k, v := range *s.value {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:66
		_go_fuzz_dep_.CoverTab[120603]++
												records = append(records, k+"="+v)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:67
		// _ = "end of CoverTab[120603]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:68
	// _ = "end of CoverTab[120600]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:68
	_go_fuzz_dep_.CoverTab[120601]++

											var buf bytes.Buffer
											w := csv.NewWriter(&buf)
											if err := w.Write(records); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:72
		_go_fuzz_dep_.CoverTab[120604]++
												panic(err)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:73
		// _ = "end of CoverTab[120604]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:74
		_go_fuzz_dep_.CoverTab[120605]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:74
		// _ = "end of CoverTab[120605]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:74
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:74
	// _ = "end of CoverTab[120601]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:74
	_go_fuzz_dep_.CoverTab[120602]++
											w.Flush()
											return "[" + strings.TrimSpace(buf.String()) + "]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:76
	// _ = "end of CoverTab[120602]"
}

func stringToStringConv(val string) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:79
	_go_fuzz_dep_.CoverTab[120606]++
											val = strings.Trim(val, "[]")

											if len(val) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:82
		_go_fuzz_dep_.CoverTab[120610]++
												return map[string]string{}, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:83
		// _ = "end of CoverTab[120610]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:84
		_go_fuzz_dep_.CoverTab[120611]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:84
		// _ = "end of CoverTab[120611]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:84
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:84
	// _ = "end of CoverTab[120606]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:84
	_go_fuzz_dep_.CoverTab[120607]++
											r := csv.NewReader(strings.NewReader(val))
											ss, err := r.Read()
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:87
		_go_fuzz_dep_.CoverTab[120612]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:88
		// _ = "end of CoverTab[120612]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:89
		_go_fuzz_dep_.CoverTab[120613]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:89
		// _ = "end of CoverTab[120613]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:89
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:89
	// _ = "end of CoverTab[120607]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:89
	_go_fuzz_dep_.CoverTab[120608]++
											out := make(map[string]string, len(ss))
											for _, pair := range ss {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:91
		_go_fuzz_dep_.CoverTab[120614]++
												kv := strings.SplitN(pair, "=", 2)
												if len(kv) != 2 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:93
			_go_fuzz_dep_.CoverTab[120616]++
													return nil, fmt.Errorf("%s must be formatted as key=value", pair)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:94
			// _ = "end of CoverTab[120616]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:95
			_go_fuzz_dep_.CoverTab[120617]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:95
			// _ = "end of CoverTab[120617]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:95
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:95
		// _ = "end of CoverTab[120614]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:95
		_go_fuzz_dep_.CoverTab[120615]++
												out[kv[0]] = kv[1]
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:96
		// _ = "end of CoverTab[120615]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:97
	// _ = "end of CoverTab[120608]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:97
	_go_fuzz_dep_.CoverTab[120609]++
											return out, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:98
	// _ = "end of CoverTab[120609]"
}

// GetStringToString return the map[string]string value of a flag with the given name
func (f *FlagSet) GetStringToString(name string) (map[string]string, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:102
	_go_fuzz_dep_.CoverTab[120618]++
												val, err := f.getFlagType(name, "stringToString", stringToStringConv)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:104
		_go_fuzz_dep_.CoverTab[120620]++
													return map[string]string{}, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:105
		// _ = "end of CoverTab[120620]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:106
		_go_fuzz_dep_.CoverTab[120621]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:106
		// _ = "end of CoverTab[120621]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:106
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:106
	// _ = "end of CoverTab[120618]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:106
	_go_fuzz_dep_.CoverTab[120619]++
												return val.(map[string]string), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:107
	// _ = "end of CoverTab[120619]"
}

// StringToStringVar defines a string flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:110
// The argument p points to a map[string]string variable in which to store the values of the multiple flags.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:110
// The value of each argument will not try to be separated by comma
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:113
func (f *FlagSet) StringToStringVar(p *map[string]string, name string, value map[string]string, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:113
	_go_fuzz_dep_.CoverTab[120622]++
												f.VarP(newStringToStringValue(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:114
	// _ = "end of CoverTab[120622]"
}

// StringToStringVarP is like StringToStringVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) StringToStringVarP(p *map[string]string, name, shorthand string, value map[string]string, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:118
	_go_fuzz_dep_.CoverTab[120623]++
												f.VarP(newStringToStringValue(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:119
	// _ = "end of CoverTab[120623]"
}

// StringToStringVar defines a string flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:122
// The argument p points to a map[string]string variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:122
// The value of each argument will not try to be separated by comma
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:125
func StringToStringVar(p *map[string]string, name string, value map[string]string, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:125
	_go_fuzz_dep_.CoverTab[120624]++
												CommandLine.VarP(newStringToStringValue(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:126
	// _ = "end of CoverTab[120624]"
}

// StringToStringVarP is like StringToStringVar, but accepts a shorthand letter that can be used after a single dash.
func StringToStringVarP(p *map[string]string, name, shorthand string, value map[string]string, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:130
	_go_fuzz_dep_.CoverTab[120625]++
												CommandLine.VarP(newStringToStringValue(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:131
	// _ = "end of CoverTab[120625]"
}

// StringToString defines a string flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:134
// The return value is the address of a map[string]string variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:134
// The value of each argument will not try to be separated by comma
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:137
func (f *FlagSet) StringToString(name string, value map[string]string, usage string) *map[string]string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:137
	_go_fuzz_dep_.CoverTab[120626]++
												p := map[string]string{}
												f.StringToStringVarP(&p, name, "", value, usage)
												return &p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:140
	// _ = "end of CoverTab[120626]"
}

// StringToStringP is like StringToString, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) StringToStringP(name, shorthand string, value map[string]string, usage string) *map[string]string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:144
	_go_fuzz_dep_.CoverTab[120627]++
												p := map[string]string{}
												f.StringToStringVarP(&p, name, shorthand, value, usage)
												return &p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:147
	// _ = "end of CoverTab[120627]"
}

// StringToString defines a string flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:150
// The return value is the address of a map[string]string variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:150
// The value of each argument will not try to be separated by comma
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:153
func StringToString(name string, value map[string]string, usage string) *map[string]string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:153
	_go_fuzz_dep_.CoverTab[120628]++
												return CommandLine.StringToStringP(name, "", value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:154
	// _ = "end of CoverTab[120628]"
}

// StringToStringP is like StringToString, but accepts a shorthand letter that can be used after a single dash.
func StringToStringP(name, shorthand string, value map[string]string, usage string) *map[string]string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:158
	_go_fuzz_dep_.CoverTab[120629]++
												return CommandLine.StringToStringP(name, shorthand, value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:159
	// _ = "end of CoverTab[120629]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:160
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_string.go:160
var _ = _go_fuzz_dep_.CoverTab
