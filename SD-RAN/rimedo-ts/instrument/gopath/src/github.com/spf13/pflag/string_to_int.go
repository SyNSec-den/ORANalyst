//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:1
package pflag

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:1
)

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// -- stringToInt Value
type stringToIntValue struct {
	value	*map[string]int
	changed	bool
}

func newStringToIntValue(val map[string]int, p *map[string]int) *stringToIntValue {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:16
	_go_fuzz_dep_.CoverTab[120496]++
											ssv := new(stringToIntValue)
											ssv.value = p
											*ssv.value = val
											return ssv
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:20
	// _ = "end of CoverTab[120496]"
}

// Format: a=1,b=2
func (s *stringToIntValue) Set(val string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:24
	_go_fuzz_dep_.CoverTab[120497]++
											ss := strings.Split(val, ",")
											out := make(map[string]int, len(ss))
											for _, pair := range ss {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:27
		_go_fuzz_dep_.CoverTab[120500]++
												kv := strings.SplitN(pair, "=", 2)
												if len(kv) != 2 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:29
			_go_fuzz_dep_.CoverTab[120502]++
													return fmt.Errorf("%s must be formatted as key=value", pair)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:30
			// _ = "end of CoverTab[120502]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:31
			_go_fuzz_dep_.CoverTab[120503]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:31
			// _ = "end of CoverTab[120503]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:31
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:31
		// _ = "end of CoverTab[120500]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:31
		_go_fuzz_dep_.CoverTab[120501]++
												var err error
												out[kv[0]], err = strconv.Atoi(kv[1])
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:34
			_go_fuzz_dep_.CoverTab[120504]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:35
			// _ = "end of CoverTab[120504]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:36
			_go_fuzz_dep_.CoverTab[120505]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:36
			// _ = "end of CoverTab[120505]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:36
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:36
		// _ = "end of CoverTab[120501]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:37
	// _ = "end of CoverTab[120497]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:37
	_go_fuzz_dep_.CoverTab[120498]++
											if !s.changed {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:38
		_go_fuzz_dep_.CoverTab[120506]++
												*s.value = out
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:39
		// _ = "end of CoverTab[120506]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:40
		_go_fuzz_dep_.CoverTab[120507]++
												for k, v := range out {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:41
			_go_fuzz_dep_.CoverTab[120508]++
													(*s.value)[k] = v
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:42
			// _ = "end of CoverTab[120508]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:43
		// _ = "end of CoverTab[120507]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:44
	// _ = "end of CoverTab[120498]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:44
	_go_fuzz_dep_.CoverTab[120499]++
											s.changed = true
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:46
	// _ = "end of CoverTab[120499]"
}

func (s *stringToIntValue) Type() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:49
	_go_fuzz_dep_.CoverTab[120509]++
											return "stringToInt"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:50
	// _ = "end of CoverTab[120509]"
}

func (s *stringToIntValue) String() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:53
	_go_fuzz_dep_.CoverTab[120510]++
											var buf bytes.Buffer
											i := 0
											for k, v := range *s.value {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:56
		_go_fuzz_dep_.CoverTab[120512]++
												if i > 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:57
			_go_fuzz_dep_.CoverTab[120514]++
													buf.WriteRune(',')
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:58
			// _ = "end of CoverTab[120514]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:59
			_go_fuzz_dep_.CoverTab[120515]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:59
			// _ = "end of CoverTab[120515]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:59
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:59
		// _ = "end of CoverTab[120512]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:59
		_go_fuzz_dep_.CoverTab[120513]++
												buf.WriteString(k)
												buf.WriteRune('=')
												buf.WriteString(strconv.Itoa(v))
												i++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:63
		// _ = "end of CoverTab[120513]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:64
	// _ = "end of CoverTab[120510]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:64
	_go_fuzz_dep_.CoverTab[120511]++
											return "[" + buf.String() + "]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:65
	// _ = "end of CoverTab[120511]"
}

func stringToIntConv(val string) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:68
	_go_fuzz_dep_.CoverTab[120516]++
											val = strings.Trim(val, "[]")

											if len(val) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:71
		_go_fuzz_dep_.CoverTab[120519]++
												return map[string]int{}, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:72
		// _ = "end of CoverTab[120519]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:73
		_go_fuzz_dep_.CoverTab[120520]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:73
		// _ = "end of CoverTab[120520]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:73
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:73
	// _ = "end of CoverTab[120516]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:73
	_go_fuzz_dep_.CoverTab[120517]++
											ss := strings.Split(val, ",")
											out := make(map[string]int, len(ss))
											for _, pair := range ss {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:76
		_go_fuzz_dep_.CoverTab[120521]++
												kv := strings.SplitN(pair, "=", 2)
												if len(kv) != 2 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:78
			_go_fuzz_dep_.CoverTab[120523]++
													return nil, fmt.Errorf("%s must be formatted as key=value", pair)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:79
			// _ = "end of CoverTab[120523]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:80
			_go_fuzz_dep_.CoverTab[120524]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:80
			// _ = "end of CoverTab[120524]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:80
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:80
		// _ = "end of CoverTab[120521]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:80
		_go_fuzz_dep_.CoverTab[120522]++
												var err error
												out[kv[0]], err = strconv.Atoi(kv[1])
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:83
			_go_fuzz_dep_.CoverTab[120525]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:84
			// _ = "end of CoverTab[120525]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:85
			_go_fuzz_dep_.CoverTab[120526]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:85
			// _ = "end of CoverTab[120526]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:85
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:85
		// _ = "end of CoverTab[120522]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:86
	// _ = "end of CoverTab[120517]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:86
	_go_fuzz_dep_.CoverTab[120518]++
											return out, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:87
	// _ = "end of CoverTab[120518]"
}

// GetStringToInt return the map[string]int value of a flag with the given name
func (f *FlagSet) GetStringToInt(name string) (map[string]int, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:91
	_go_fuzz_dep_.CoverTab[120527]++
											val, err := f.getFlagType(name, "stringToInt", stringToIntConv)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:93
		_go_fuzz_dep_.CoverTab[120529]++
												return map[string]int{}, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:94
		// _ = "end of CoverTab[120529]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:95
		_go_fuzz_dep_.CoverTab[120530]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:95
		// _ = "end of CoverTab[120530]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:95
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:95
	// _ = "end of CoverTab[120527]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:95
	_go_fuzz_dep_.CoverTab[120528]++
											return val.(map[string]int), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:96
	// _ = "end of CoverTab[120528]"
}

// StringToIntVar defines a string flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:99
// The argument p points to a map[string]int variable in which to store the values of the multiple flags.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:99
// The value of each argument will not try to be separated by comma
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:102
func (f *FlagSet) StringToIntVar(p *map[string]int, name string, value map[string]int, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:102
	_go_fuzz_dep_.CoverTab[120531]++
											f.VarP(newStringToIntValue(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:103
	// _ = "end of CoverTab[120531]"
}

// StringToIntVarP is like StringToIntVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) StringToIntVarP(p *map[string]int, name, shorthand string, value map[string]int, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:107
	_go_fuzz_dep_.CoverTab[120532]++
											f.VarP(newStringToIntValue(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:108
	// _ = "end of CoverTab[120532]"
}

// StringToIntVar defines a string flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:111
// The argument p points to a map[string]int variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:111
// The value of each argument will not try to be separated by comma
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:114
func StringToIntVar(p *map[string]int, name string, value map[string]int, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:114
	_go_fuzz_dep_.CoverTab[120533]++
											CommandLine.VarP(newStringToIntValue(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:115
	// _ = "end of CoverTab[120533]"
}

// StringToIntVarP is like StringToIntVar, but accepts a shorthand letter that can be used after a single dash.
func StringToIntVarP(p *map[string]int, name, shorthand string, value map[string]int, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:119
	_go_fuzz_dep_.CoverTab[120534]++
											CommandLine.VarP(newStringToIntValue(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:120
	// _ = "end of CoverTab[120534]"
}

// StringToInt defines a string flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:123
// The return value is the address of a map[string]int variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:123
// The value of each argument will not try to be separated by comma
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:126
func (f *FlagSet) StringToInt(name string, value map[string]int, usage string) *map[string]int {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:126
	_go_fuzz_dep_.CoverTab[120535]++
											p := map[string]int{}
											f.StringToIntVarP(&p, name, "", value, usage)
											return &p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:129
	// _ = "end of CoverTab[120535]"
}

// StringToIntP is like StringToInt, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) StringToIntP(name, shorthand string, value map[string]int, usage string) *map[string]int {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:133
	_go_fuzz_dep_.CoverTab[120536]++
											p := map[string]int{}
											f.StringToIntVarP(&p, name, shorthand, value, usage)
											return &p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:136
	// _ = "end of CoverTab[120536]"
}

// StringToInt defines a string flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:139
// The return value is the address of a map[string]int variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:139
// The value of each argument will not try to be separated by comma
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:142
func StringToInt(name string, value map[string]int, usage string) *map[string]int {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:142
	_go_fuzz_dep_.CoverTab[120537]++
											return CommandLine.StringToIntP(name, "", value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:143
	// _ = "end of CoverTab[120537]"
}

// StringToIntP is like StringToInt, but accepts a shorthand letter that can be used after a single dash.
func StringToIntP(name, shorthand string, value map[string]int, usage string) *map[string]int {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:147
	_go_fuzz_dep_.CoverTab[120538]++
											return CommandLine.StringToIntP(name, shorthand, value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:148
	// _ = "end of CoverTab[120538]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:149
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_to_int.go:149
var _ = _go_fuzz_dep_.CoverTab
