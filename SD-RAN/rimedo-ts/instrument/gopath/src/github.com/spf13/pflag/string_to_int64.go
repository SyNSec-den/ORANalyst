//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:1
package pflag

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:1
)

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// -- stringToInt64 Value
type stringToInt64Value struct {
	value	*map[string]int64
	changed	bool
}

func newStringToInt64Value(val map[string]int64, p *map[string]int64) *stringToInt64Value {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:16
	_go_fuzz_dep_.CoverTab[120539]++
											ssv := new(stringToInt64Value)
											ssv.value = p
											*ssv.value = val
											return ssv
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:20
	// _ = "end of CoverTab[120539]"
}

// Format: a=1,b=2
func (s *stringToInt64Value) Set(val string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:24
	_go_fuzz_dep_.CoverTab[120540]++
											ss := strings.Split(val, ",")
											out := make(map[string]int64, len(ss))
											for _, pair := range ss {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:27
		_go_fuzz_dep_.CoverTab[120543]++
												kv := strings.SplitN(pair, "=", 2)
												if len(kv) != 2 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:29
			_go_fuzz_dep_.CoverTab[120545]++
													return fmt.Errorf("%s must be formatted as key=value", pair)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:30
			// _ = "end of CoverTab[120545]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:31
			_go_fuzz_dep_.CoverTab[120546]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:31
			// _ = "end of CoverTab[120546]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:31
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:31
		// _ = "end of CoverTab[120543]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:31
		_go_fuzz_dep_.CoverTab[120544]++
												var err error
												out[kv[0]], err = strconv.ParseInt(kv[1], 10, 64)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:34
			_go_fuzz_dep_.CoverTab[120547]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:35
			// _ = "end of CoverTab[120547]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:36
			_go_fuzz_dep_.CoverTab[120548]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:36
			// _ = "end of CoverTab[120548]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:36
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:36
		// _ = "end of CoverTab[120544]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:37
	// _ = "end of CoverTab[120540]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:37
	_go_fuzz_dep_.CoverTab[120541]++
											if !s.changed {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:38
		_go_fuzz_dep_.CoverTab[120549]++
												*s.value = out
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:39
		// _ = "end of CoverTab[120549]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:40
		_go_fuzz_dep_.CoverTab[120550]++
												for k, v := range out {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:41
			_go_fuzz_dep_.CoverTab[120551]++
													(*s.value)[k] = v
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:42
			// _ = "end of CoverTab[120551]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:43
		// _ = "end of CoverTab[120550]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:44
	// _ = "end of CoverTab[120541]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:44
	_go_fuzz_dep_.CoverTab[120542]++
											s.changed = true
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:46
	// _ = "end of CoverTab[120542]"
}

func (s *stringToInt64Value) Type() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:49
	_go_fuzz_dep_.CoverTab[120552]++
											return "stringToInt64"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:50
	// _ = "end of CoverTab[120552]"
}

func (s *stringToInt64Value) String() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:53
	_go_fuzz_dep_.CoverTab[120553]++
											var buf bytes.Buffer
											i := 0
											for k, v := range *s.value {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:56
		_go_fuzz_dep_.CoverTab[120555]++
												if i > 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:57
			_go_fuzz_dep_.CoverTab[120557]++
													buf.WriteRune(',')
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:58
			// _ = "end of CoverTab[120557]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:59
			_go_fuzz_dep_.CoverTab[120558]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:59
			// _ = "end of CoverTab[120558]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:59
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:59
		// _ = "end of CoverTab[120555]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:59
		_go_fuzz_dep_.CoverTab[120556]++
												buf.WriteString(k)
												buf.WriteRune('=')
												buf.WriteString(strconv.FormatInt(v, 10))
												i++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:63
		// _ = "end of CoverTab[120556]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:64
	// _ = "end of CoverTab[120553]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:64
	_go_fuzz_dep_.CoverTab[120554]++
											return "[" + buf.String() + "]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:65
	// _ = "end of CoverTab[120554]"
}

func stringToInt64Conv(val string) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:68
	_go_fuzz_dep_.CoverTab[120559]++
											val = strings.Trim(val, "[]")

											if len(val) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:71
		_go_fuzz_dep_.CoverTab[120562]++
												return map[string]int64{}, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:72
		// _ = "end of CoverTab[120562]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:73
		_go_fuzz_dep_.CoverTab[120563]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:73
		// _ = "end of CoverTab[120563]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:73
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:73
	// _ = "end of CoverTab[120559]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:73
	_go_fuzz_dep_.CoverTab[120560]++
											ss := strings.Split(val, ",")
											out := make(map[string]int64, len(ss))
											for _, pair := range ss {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:76
		_go_fuzz_dep_.CoverTab[120564]++
												kv := strings.SplitN(pair, "=", 2)
												if len(kv) != 2 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:78
			_go_fuzz_dep_.CoverTab[120566]++
													return nil, fmt.Errorf("%s must be formatted as key=value", pair)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:79
			// _ = "end of CoverTab[120566]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:80
			_go_fuzz_dep_.CoverTab[120567]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:80
			// _ = "end of CoverTab[120567]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:80
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:80
		// _ = "end of CoverTab[120564]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:80
		_go_fuzz_dep_.CoverTab[120565]++
												var err error
												out[kv[0]], err = strconv.ParseInt(kv[1], 10, 64)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:83
			_go_fuzz_dep_.CoverTab[120568]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:84
			// _ = "end of CoverTab[120568]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:85
			_go_fuzz_dep_.CoverTab[120569]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:85
			// _ = "end of CoverTab[120569]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:85
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:85
		// _ = "end of CoverTab[120565]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:86
	// _ = "end of CoverTab[120560]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:86
	_go_fuzz_dep_.CoverTab[120561]++
											return out, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:87
	// _ = "end of CoverTab[120561]"
}

// GetStringToInt64 return the map[string]int64 value of a flag with the given name
func (f *FlagSet) GetStringToInt64(name string) (map[string]int64, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:91
	_go_fuzz_dep_.CoverTab[120570]++
											val, err := f.getFlagType(name, "stringToInt64", stringToInt64Conv)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:93
		_go_fuzz_dep_.CoverTab[120572]++
												return map[string]int64{}, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:94
		// _ = "end of CoverTab[120572]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:95
		_go_fuzz_dep_.CoverTab[120573]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:95
		// _ = "end of CoverTab[120573]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:95
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:95
	// _ = "end of CoverTab[120570]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:95
	_go_fuzz_dep_.CoverTab[120571]++
											return val.(map[string]int64), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:96
	// _ = "end of CoverTab[120571]"
}

// StringToInt64Var defines a string flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:99
// The argument p point64s to a map[string]int64 variable in which to store the values of the multiple flags.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:99
// The value of each argument will not try to be separated by comma
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:102
func (f *FlagSet) StringToInt64Var(p *map[string]int64, name string, value map[string]int64, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:102
	_go_fuzz_dep_.CoverTab[120574]++
											f.VarP(newStringToInt64Value(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:103
	// _ = "end of CoverTab[120574]"
}

// StringToInt64VarP is like StringToInt64Var, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) StringToInt64VarP(p *map[string]int64, name, shorthand string, value map[string]int64, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:107
	_go_fuzz_dep_.CoverTab[120575]++
											f.VarP(newStringToInt64Value(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:108
	// _ = "end of CoverTab[120575]"
}

// StringToInt64Var defines a string flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:111
// The argument p point64s to a map[string]int64 variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:111
// The value of each argument will not try to be separated by comma
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:114
func StringToInt64Var(p *map[string]int64, name string, value map[string]int64, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:114
	_go_fuzz_dep_.CoverTab[120576]++
											CommandLine.VarP(newStringToInt64Value(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:115
	// _ = "end of CoverTab[120576]"
}

// StringToInt64VarP is like StringToInt64Var, but accepts a shorthand letter that can be used after a single dash.
func StringToInt64VarP(p *map[string]int64, name, shorthand string, value map[string]int64, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:119
	_go_fuzz_dep_.CoverTab[120577]++
											CommandLine.VarP(newStringToInt64Value(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:120
	// _ = "end of CoverTab[120577]"
}

// StringToInt64 defines a string flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:123
// The return value is the address of a map[string]int64 variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:123
// The value of each argument will not try to be separated by comma
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:126
func (f *FlagSet) StringToInt64(name string, value map[string]int64, usage string) *map[string]int64 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:126
	_go_fuzz_dep_.CoverTab[120578]++
											p := map[string]int64{}
											f.StringToInt64VarP(&p, name, "", value, usage)
											return &p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:129
	// _ = "end of CoverTab[120578]"
}

// StringToInt64P is like StringToInt64, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) StringToInt64P(name, shorthand string, value map[string]int64, usage string) *map[string]int64 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:133
	_go_fuzz_dep_.CoverTab[120579]++
											p := map[string]int64{}
											f.StringToInt64VarP(&p, name, shorthand, value, usage)
											return &p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:136
	// _ = "end of CoverTab[120579]"
}

// StringToInt64 defines a string flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:139
// The return value is the address of a map[string]int64 variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:139
// The value of each argument will not try to be separated by comma
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:142
func StringToInt64(name string, value map[string]int64, usage string) *map[string]int64 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:142
	_go_fuzz_dep_.CoverTab[120580]++
											return CommandLine.StringToInt64P(name, "", value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:143
	// _ = "end of CoverTab[120580]"
}

// StringToInt64P is like StringToInt64, but accepts a shorthand letter that can be used after a single dash.
func StringToInt64P(name, shorthand string, value map[string]int64, usage string) *map[string]int64 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:147
	_go_fuzz_dep_.CoverTab[120581]++
											return CommandLine.StringToInt64P(name, shorthand, value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:148
	// _ = "end of CoverTab[120581]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:149
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int64.go:149
var _ = _go_fuzz_dep_.CoverTab
