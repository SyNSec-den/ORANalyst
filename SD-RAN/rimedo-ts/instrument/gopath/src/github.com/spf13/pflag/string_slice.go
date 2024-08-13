//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:1
package pflag

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:1
)

import (
	"bytes"
	"encoding/csv"
	"strings"
)

// -- stringSlice Value
type stringSliceValue struct {
	value	*[]string
	changed	bool
}

func newStringSliceValue(val []string, p *[]string) *stringSliceValue {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:15
	_go_fuzz_dep_.CoverTab[120459]++
											ssv := new(stringSliceValue)
											ssv.value = p
											*ssv.value = val
											return ssv
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:19
	// _ = "end of CoverTab[120459]"
}

func readAsCSV(val string) ([]string, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:22
	_go_fuzz_dep_.CoverTab[120460]++
											if val == "" {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:23
		_go_fuzz_dep_.CoverTab[120462]++
												return []string{}, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:24
		// _ = "end of CoverTab[120462]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:25
		_go_fuzz_dep_.CoverTab[120463]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:25
		// _ = "end of CoverTab[120463]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:25
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:25
	// _ = "end of CoverTab[120460]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:25
	_go_fuzz_dep_.CoverTab[120461]++
											stringReader := strings.NewReader(val)
											csvReader := csv.NewReader(stringReader)
											return csvReader.Read()
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:28
	// _ = "end of CoverTab[120461]"
}

func writeAsCSV(vals []string) (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:31
	_go_fuzz_dep_.CoverTab[120464]++
											b := &bytes.Buffer{}
											w := csv.NewWriter(b)
											err := w.Write(vals)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:35
		_go_fuzz_dep_.CoverTab[120466]++
												return "", err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:36
		// _ = "end of CoverTab[120466]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:37
		_go_fuzz_dep_.CoverTab[120467]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:37
		// _ = "end of CoverTab[120467]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:37
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:37
	// _ = "end of CoverTab[120464]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:37
	_go_fuzz_dep_.CoverTab[120465]++
											w.Flush()
											return strings.TrimSuffix(b.String(), "\n"), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:39
	// _ = "end of CoverTab[120465]"
}

func (s *stringSliceValue) Set(val string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:42
	_go_fuzz_dep_.CoverTab[120468]++
											v, err := readAsCSV(val)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:44
		_go_fuzz_dep_.CoverTab[120471]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:45
		// _ = "end of CoverTab[120471]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:46
		_go_fuzz_dep_.CoverTab[120472]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:46
		// _ = "end of CoverTab[120472]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:46
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:46
	// _ = "end of CoverTab[120468]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:46
	_go_fuzz_dep_.CoverTab[120469]++
											if !s.changed {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:47
		_go_fuzz_dep_.CoverTab[120473]++
												*s.value = v
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:48
		// _ = "end of CoverTab[120473]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:49
		_go_fuzz_dep_.CoverTab[120474]++
												*s.value = append(*s.value, v...)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:50
		// _ = "end of CoverTab[120474]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:51
	// _ = "end of CoverTab[120469]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:51
	_go_fuzz_dep_.CoverTab[120470]++
											s.changed = true
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:53
	// _ = "end of CoverTab[120470]"
}

func (s *stringSliceValue) Type() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:56
	_go_fuzz_dep_.CoverTab[120475]++
											return "stringSlice"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:57
	// _ = "end of CoverTab[120475]"
}

func (s *stringSliceValue) String() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:60
	_go_fuzz_dep_.CoverTab[120476]++
											str, _ := writeAsCSV(*s.value)
											return "[" + str + "]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:62
	// _ = "end of CoverTab[120476]"
}

func (s *stringSliceValue) Append(val string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:65
	_go_fuzz_dep_.CoverTab[120477]++
											*s.value = append(*s.value, val)
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:67
	// _ = "end of CoverTab[120477]"
}

func (s *stringSliceValue) Replace(val []string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:70
	_go_fuzz_dep_.CoverTab[120478]++
											*s.value = val
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:72
	// _ = "end of CoverTab[120478]"
}

func (s *stringSliceValue) GetSlice() []string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:75
	_go_fuzz_dep_.CoverTab[120479]++
											return *s.value
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:76
	// _ = "end of CoverTab[120479]"
}

func stringSliceConv(sval string) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:79
	_go_fuzz_dep_.CoverTab[120480]++
											sval = sval[1 : len(sval)-1]

											if len(sval) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:82
		_go_fuzz_dep_.CoverTab[120482]++
												return []string{}, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:83
		// _ = "end of CoverTab[120482]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:84
		_go_fuzz_dep_.CoverTab[120483]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:84
		// _ = "end of CoverTab[120483]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:84
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:84
	// _ = "end of CoverTab[120480]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:84
	_go_fuzz_dep_.CoverTab[120481]++
											return readAsCSV(sval)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:85
	// _ = "end of CoverTab[120481]"
}

// GetStringSlice return the []string value of a flag with the given name
func (f *FlagSet) GetStringSlice(name string) ([]string, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:89
	_go_fuzz_dep_.CoverTab[120484]++
											val, err := f.getFlagType(name, "stringSlice", stringSliceConv)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:91
		_go_fuzz_dep_.CoverTab[120486]++
												return []string{}, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:92
		// _ = "end of CoverTab[120486]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:93
		_go_fuzz_dep_.CoverTab[120487]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:93
		// _ = "end of CoverTab[120487]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:93
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:93
	// _ = "end of CoverTab[120484]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:93
	_go_fuzz_dep_.CoverTab[120485]++
											return val.([]string), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:94
	// _ = "end of CoverTab[120485]"
}

// StringSliceVar defines a string flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:97
// The argument p points to a []string variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:97
// Compared to StringArray flags, StringSlice flags take comma-separated value as arguments and split them accordingly.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:97
// For example:
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:97
//
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:97
//	--ss="v1,v2" --ss="v3"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:97
//
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:97
// will result in
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:97
//
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:97
//	[]string{"v1", "v2", "v3"}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:104
func (f *FlagSet) StringSliceVar(p *[]string, name string, value []string, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:104
	_go_fuzz_dep_.CoverTab[120488]++
											f.VarP(newStringSliceValue(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:105
	// _ = "end of CoverTab[120488]"
}

// StringSliceVarP is like StringSliceVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) StringSliceVarP(p *[]string, name, shorthand string, value []string, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:109
	_go_fuzz_dep_.CoverTab[120489]++
											f.VarP(newStringSliceValue(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:110
	// _ = "end of CoverTab[120489]"
}

// StringSliceVar defines a string flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:113
// The argument p points to a []string variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:113
// Compared to StringArray flags, StringSlice flags take comma-separated value as arguments and split them accordingly.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:113
// For example:
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:113
//
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:113
//	--ss="v1,v2" --ss="v3"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:113
//
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:113
// will result in
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:113
//
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:113
//	[]string{"v1", "v2", "v3"}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:120
func StringSliceVar(p *[]string, name string, value []string, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:120
	_go_fuzz_dep_.CoverTab[120490]++
											CommandLine.VarP(newStringSliceValue(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:121
	// _ = "end of CoverTab[120490]"
}

// StringSliceVarP is like StringSliceVar, but accepts a shorthand letter that can be used after a single dash.
func StringSliceVarP(p *[]string, name, shorthand string, value []string, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:125
	_go_fuzz_dep_.CoverTab[120491]++
											CommandLine.VarP(newStringSliceValue(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:126
	// _ = "end of CoverTab[120491]"
}

// StringSlice defines a string flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:129
// The return value is the address of a []string variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:129
// Compared to StringArray flags, StringSlice flags take comma-separated value as arguments and split them accordingly.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:129
// For example:
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:129
//
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:129
//	--ss="v1,v2" --ss="v3"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:129
//
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:129
// will result in
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:129
//
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:129
//	[]string{"v1", "v2", "v3"}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:136
func (f *FlagSet) StringSlice(name string, value []string, usage string) *[]string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:136
	_go_fuzz_dep_.CoverTab[120492]++
											p := []string{}
											f.StringSliceVarP(&p, name, "", value, usage)
											return &p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:139
	// _ = "end of CoverTab[120492]"
}

// StringSliceP is like StringSlice, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) StringSliceP(name, shorthand string, value []string, usage string) *[]string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:143
	_go_fuzz_dep_.CoverTab[120493]++
											p := []string{}
											f.StringSliceVarP(&p, name, shorthand, value, usage)
											return &p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:146
	// _ = "end of CoverTab[120493]"
}

// StringSlice defines a string flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:149
// The return value is the address of a []string variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:149
// Compared to StringArray flags, StringSlice flags take comma-separated value as arguments and split them accordingly.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:149
// For example:
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:149
//
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:149
//	--ss="v1,v2" --ss="v3"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:149
//
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:149
// will result in
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:149
//
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:149
//	[]string{"v1", "v2", "v3"}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:156
func StringSlice(name string, value []string, usage string) *[]string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:156
	_go_fuzz_dep_.CoverTab[120494]++
											return CommandLine.StringSliceP(name, "", value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:157
	// _ = "end of CoverTab[120494]"
}

// StringSliceP is like StringSlice, but accepts a shorthand letter that can be used after a single dash.
func StringSliceP(name, shorthand string, value []string, usage string) *[]string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:161
	_go_fuzz_dep_.CoverTab[120495]++
											return CommandLine.StringSliceP(name, shorthand, value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:162
	// _ = "end of CoverTab[120495]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:163
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string_slice.go:163
var _ = _go_fuzz_dep_.CoverTab
