//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:1
package pflag

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:1
)

// -- stringArray Value
type stringArrayValue struct {
	value	*[]string
	changed	bool
}

func newStringArrayValue(val []string, p *[]string) *stringArrayValue {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:9
	_go_fuzz_dep_.CoverTab[120427]++
											ssv := new(stringArrayValue)
											ssv.value = p
											*ssv.value = val
											return ssv
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:13
	// _ = "end of CoverTab[120427]"
}

func (s *stringArrayValue) Set(val string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:16
	_go_fuzz_dep_.CoverTab[120428]++
											if !s.changed {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:17
		_go_fuzz_dep_.CoverTab[120430]++
												*s.value = []string{val}
												s.changed = true
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:19
		// _ = "end of CoverTab[120430]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:20
		_go_fuzz_dep_.CoverTab[120431]++
												*s.value = append(*s.value, val)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:21
		// _ = "end of CoverTab[120431]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:22
	// _ = "end of CoverTab[120428]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:22
	_go_fuzz_dep_.CoverTab[120429]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:23
	// _ = "end of CoverTab[120429]"
}

func (s *stringArrayValue) Append(val string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:26
	_go_fuzz_dep_.CoverTab[120432]++
											*s.value = append(*s.value, val)
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:28
	// _ = "end of CoverTab[120432]"
}

func (s *stringArrayValue) Replace(val []string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:31
	_go_fuzz_dep_.CoverTab[120433]++
											out := make([]string, len(val))
											for i, d := range val {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:33
		_go_fuzz_dep_.CoverTab[120435]++
												var err error
												out[i] = d
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:36
			_go_fuzz_dep_.CoverTab[120436]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:37
			// _ = "end of CoverTab[120436]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:38
			_go_fuzz_dep_.CoverTab[120437]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:38
			// _ = "end of CoverTab[120437]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:38
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:38
		// _ = "end of CoverTab[120435]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:39
	// _ = "end of CoverTab[120433]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:39
	_go_fuzz_dep_.CoverTab[120434]++
											*s.value = out
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:41
	// _ = "end of CoverTab[120434]"
}

func (s *stringArrayValue) GetSlice() []string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:44
	_go_fuzz_dep_.CoverTab[120438]++
											out := make([]string, len(*s.value))
											for i, d := range *s.value {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:46
		_go_fuzz_dep_.CoverTab[120440]++
												out[i] = d
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:47
		// _ = "end of CoverTab[120440]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:48
	// _ = "end of CoverTab[120438]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:48
	_go_fuzz_dep_.CoverTab[120439]++
											return out
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:49
	// _ = "end of CoverTab[120439]"
}

func (s *stringArrayValue) Type() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:52
	_go_fuzz_dep_.CoverTab[120441]++
											return "stringArray"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:53
	// _ = "end of CoverTab[120441]"
}

func (s *stringArrayValue) String() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:56
	_go_fuzz_dep_.CoverTab[120442]++
											str, _ := writeAsCSV(*s.value)
											return "[" + str + "]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:58
	// _ = "end of CoverTab[120442]"
}

func stringArrayConv(sval string) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:61
	_go_fuzz_dep_.CoverTab[120443]++
											sval = sval[1 : len(sval)-1]

											if len(sval) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:64
		_go_fuzz_dep_.CoverTab[120445]++
												return []string{}, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:65
		// _ = "end of CoverTab[120445]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:66
		_go_fuzz_dep_.CoverTab[120446]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:66
		// _ = "end of CoverTab[120446]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:66
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:66
	// _ = "end of CoverTab[120443]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:66
	_go_fuzz_dep_.CoverTab[120444]++
											return readAsCSV(sval)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:67
	// _ = "end of CoverTab[120444]"
}

// GetStringArray return the []string value of a flag with the given name
func (f *FlagSet) GetStringArray(name string) ([]string, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:71
	_go_fuzz_dep_.CoverTab[120447]++
											val, err := f.getFlagType(name, "stringArray", stringArrayConv)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:73
		_go_fuzz_dep_.CoverTab[120449]++
												return []string{}, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:74
		// _ = "end of CoverTab[120449]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:75
		_go_fuzz_dep_.CoverTab[120450]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:75
		// _ = "end of CoverTab[120450]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:75
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:75
	// _ = "end of CoverTab[120447]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:75
	_go_fuzz_dep_.CoverTab[120448]++
											return val.([]string), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:76
	// _ = "end of CoverTab[120448]"
}

// StringArrayVar defines a string flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:79
// The argument p points to a []string variable in which to store the values of the multiple flags.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:79
// The value of each argument will not try to be separated by comma. Use a StringSlice for that.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:82
func (f *FlagSet) StringArrayVar(p *[]string, name string, value []string, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:82
	_go_fuzz_dep_.CoverTab[120451]++
											f.VarP(newStringArrayValue(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:83
	// _ = "end of CoverTab[120451]"
}

// StringArrayVarP is like StringArrayVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) StringArrayVarP(p *[]string, name, shorthand string, value []string, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:87
	_go_fuzz_dep_.CoverTab[120452]++
											f.VarP(newStringArrayValue(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:88
	// _ = "end of CoverTab[120452]"
}

// StringArrayVar defines a string flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:91
// The argument p points to a []string variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:91
// The value of each argument will not try to be separated by comma. Use a StringSlice for that.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:94
func StringArrayVar(p *[]string, name string, value []string, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:94
	_go_fuzz_dep_.CoverTab[120453]++
											CommandLine.VarP(newStringArrayValue(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:95
	// _ = "end of CoverTab[120453]"
}

// StringArrayVarP is like StringArrayVar, but accepts a shorthand letter that can be used after a single dash.
func StringArrayVarP(p *[]string, name, shorthand string, value []string, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:99
	_go_fuzz_dep_.CoverTab[120454]++
											CommandLine.VarP(newStringArrayValue(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:100
	// _ = "end of CoverTab[120454]"
}

// StringArray defines a string flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:103
// The return value is the address of a []string variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:103
// The value of each argument will not try to be separated by comma. Use a StringSlice for that.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:106
func (f *FlagSet) StringArray(name string, value []string, usage string) *[]string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:106
	_go_fuzz_dep_.CoverTab[120455]++
											p := []string{}
											f.StringArrayVarP(&p, name, "", value, usage)
											return &p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:109
	// _ = "end of CoverTab[120455]"
}

// StringArrayP is like StringArray, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) StringArrayP(name, shorthand string, value []string, usage string) *[]string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:113
	_go_fuzz_dep_.CoverTab[120456]++
											p := []string{}
											f.StringArrayVarP(&p, name, shorthand, value, usage)
											return &p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:116
	// _ = "end of CoverTab[120456]"
}

// StringArray defines a string flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:119
// The return value is the address of a []string variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:119
// The value of each argument will not try to be separated by comma. Use a StringSlice for that.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:122
func StringArray(name string, value []string, usage string) *[]string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:122
	_go_fuzz_dep_.CoverTab[120457]++
											return CommandLine.StringArrayP(name, "", value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:123
	// _ = "end of CoverTab[120457]"
}

// StringArrayP is like StringArray, but accepts a shorthand letter that can be used after a single dash.
func StringArrayP(name, shorthand string, value []string, usage string) *[]string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:127
	_go_fuzz_dep_.CoverTab[120458]++
											return CommandLine.StringArrayP(name, shorthand, value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:128
	// _ = "end of CoverTab[120458]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:129
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_array.go:129
var _ = _go_fuzz_dep_.CoverTab
