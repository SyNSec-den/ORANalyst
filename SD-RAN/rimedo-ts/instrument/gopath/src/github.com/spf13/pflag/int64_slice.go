//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:1
package pflag

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:1
)

import (
	"fmt"
	"strconv"
	"strings"
)

// -- int64Slice Value
type int64SliceValue struct {
	value	*[]int64
	changed	bool
}

func newInt64SliceValue(val []int64, p *[]int64) *int64SliceValue {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:15
	_go_fuzz_dep_.CoverTab[120162]++
											isv := new(int64SliceValue)
											isv.value = p
											*isv.value = val
											return isv
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:19
	// _ = "end of CoverTab[120162]"
}

func (s *int64SliceValue) Set(val string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:22
	_go_fuzz_dep_.CoverTab[120163]++
											ss := strings.Split(val, ",")
											out := make([]int64, len(ss))
											for i, d := range ss {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:25
		_go_fuzz_dep_.CoverTab[120166]++
												var err error
												out[i], err = strconv.ParseInt(d, 0, 64)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:28
			_go_fuzz_dep_.CoverTab[120167]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:29
			// _ = "end of CoverTab[120167]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:30
			_go_fuzz_dep_.CoverTab[120168]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:30
			// _ = "end of CoverTab[120168]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:30
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:30
		// _ = "end of CoverTab[120166]"

	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:32
	// _ = "end of CoverTab[120163]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:32
	_go_fuzz_dep_.CoverTab[120164]++
											if !s.changed {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:33
		_go_fuzz_dep_.CoverTab[120169]++
												*s.value = out
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:34
		// _ = "end of CoverTab[120169]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:35
		_go_fuzz_dep_.CoverTab[120170]++
												*s.value = append(*s.value, out...)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:36
		// _ = "end of CoverTab[120170]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:37
	// _ = "end of CoverTab[120164]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:37
	_go_fuzz_dep_.CoverTab[120165]++
											s.changed = true
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:39
	// _ = "end of CoverTab[120165]"
}

func (s *int64SliceValue) Type() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:42
	_go_fuzz_dep_.CoverTab[120171]++
											return "int64Slice"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:43
	// _ = "end of CoverTab[120171]"
}

func (s *int64SliceValue) String() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:46
	_go_fuzz_dep_.CoverTab[120172]++
											out := make([]string, len(*s.value))
											for i, d := range *s.value {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:48
		_go_fuzz_dep_.CoverTab[120174]++
												out[i] = fmt.Sprintf("%d", d)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:49
		// _ = "end of CoverTab[120174]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:50
	// _ = "end of CoverTab[120172]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:50
	_go_fuzz_dep_.CoverTab[120173]++
											return "[" + strings.Join(out, ",") + "]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:51
	// _ = "end of CoverTab[120173]"
}

func (s *int64SliceValue) fromString(val string) (int64, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:54
	_go_fuzz_dep_.CoverTab[120175]++
											return strconv.ParseInt(val, 0, 64)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:55
	// _ = "end of CoverTab[120175]"
}

func (s *int64SliceValue) toString(val int64) string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:58
	_go_fuzz_dep_.CoverTab[120176]++
											return fmt.Sprintf("%d", val)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:59
	// _ = "end of CoverTab[120176]"
}

func (s *int64SliceValue) Append(val string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:62
	_go_fuzz_dep_.CoverTab[120177]++
											i, err := s.fromString(val)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:64
		_go_fuzz_dep_.CoverTab[120179]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:65
		// _ = "end of CoverTab[120179]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:66
		_go_fuzz_dep_.CoverTab[120180]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:66
		// _ = "end of CoverTab[120180]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:66
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:66
	// _ = "end of CoverTab[120177]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:66
	_go_fuzz_dep_.CoverTab[120178]++
											*s.value = append(*s.value, i)
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:68
	// _ = "end of CoverTab[120178]"
}

func (s *int64SliceValue) Replace(val []string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:71
	_go_fuzz_dep_.CoverTab[120181]++
											out := make([]int64, len(val))
											for i, d := range val {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:73
		_go_fuzz_dep_.CoverTab[120183]++
												var err error
												out[i], err = s.fromString(d)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:76
			_go_fuzz_dep_.CoverTab[120184]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:77
			// _ = "end of CoverTab[120184]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:78
			_go_fuzz_dep_.CoverTab[120185]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:78
			// _ = "end of CoverTab[120185]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:78
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:78
		// _ = "end of CoverTab[120183]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:79
	// _ = "end of CoverTab[120181]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:79
	_go_fuzz_dep_.CoverTab[120182]++
											*s.value = out
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:81
	// _ = "end of CoverTab[120182]"
}

func (s *int64SliceValue) GetSlice() []string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:84
	_go_fuzz_dep_.CoverTab[120186]++
											out := make([]string, len(*s.value))
											for i, d := range *s.value {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:86
		_go_fuzz_dep_.CoverTab[120188]++
												out[i] = s.toString(d)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:87
		// _ = "end of CoverTab[120188]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:88
	// _ = "end of CoverTab[120186]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:88
	_go_fuzz_dep_.CoverTab[120187]++
											return out
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:89
	// _ = "end of CoverTab[120187]"
}

func int64SliceConv(val string) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:92
	_go_fuzz_dep_.CoverTab[120189]++
											val = strings.Trim(val, "[]")

											if len(val) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:95
		_go_fuzz_dep_.CoverTab[120192]++
												return []int64{}, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:96
		// _ = "end of CoverTab[120192]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:97
		_go_fuzz_dep_.CoverTab[120193]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:97
		// _ = "end of CoverTab[120193]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:97
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:97
	// _ = "end of CoverTab[120189]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:97
	_go_fuzz_dep_.CoverTab[120190]++
											ss := strings.Split(val, ",")
											out := make([]int64, len(ss))
											for i, d := range ss {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:100
		_go_fuzz_dep_.CoverTab[120194]++
												var err error
												out[i], err = strconv.ParseInt(d, 0, 64)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:103
			_go_fuzz_dep_.CoverTab[120195]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:104
			// _ = "end of CoverTab[120195]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:105
			_go_fuzz_dep_.CoverTab[120196]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:105
			// _ = "end of CoverTab[120196]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:105
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:105
		// _ = "end of CoverTab[120194]"

	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:107
	// _ = "end of CoverTab[120190]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:107
	_go_fuzz_dep_.CoverTab[120191]++
											return out, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:108
	// _ = "end of CoverTab[120191]"
}

// GetInt64Slice return the []int64 value of a flag with the given name
func (f *FlagSet) GetInt64Slice(name string) ([]int64, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:112
	_go_fuzz_dep_.CoverTab[120197]++
											val, err := f.getFlagType(name, "int64Slice", int64SliceConv)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:114
		_go_fuzz_dep_.CoverTab[120199]++
												return []int64{}, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:115
		// _ = "end of CoverTab[120199]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:116
		_go_fuzz_dep_.CoverTab[120200]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:116
		// _ = "end of CoverTab[120200]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:116
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:116
	// _ = "end of CoverTab[120197]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:116
	_go_fuzz_dep_.CoverTab[120198]++
											return val.([]int64), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:117
	// _ = "end of CoverTab[120198]"
}

// Int64SliceVar defines a int64Slice flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:120
// The argument p points to a []int64 variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:122
func (f *FlagSet) Int64SliceVar(p *[]int64, name string, value []int64, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:122
	_go_fuzz_dep_.CoverTab[120201]++
											f.VarP(newInt64SliceValue(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:123
	// _ = "end of CoverTab[120201]"
}

// Int64SliceVarP is like Int64SliceVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Int64SliceVarP(p *[]int64, name, shorthand string, value []int64, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:127
	_go_fuzz_dep_.CoverTab[120202]++
											f.VarP(newInt64SliceValue(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:128
	// _ = "end of CoverTab[120202]"
}

// Int64SliceVar defines a int64[] flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:131
// The argument p points to a int64[] variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:133
func Int64SliceVar(p *[]int64, name string, value []int64, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:133
	_go_fuzz_dep_.CoverTab[120203]++
											CommandLine.VarP(newInt64SliceValue(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:134
	// _ = "end of CoverTab[120203]"
}

// Int64SliceVarP is like Int64SliceVar, but accepts a shorthand letter that can be used after a single dash.
func Int64SliceVarP(p *[]int64, name, shorthand string, value []int64, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:138
	_go_fuzz_dep_.CoverTab[120204]++
											CommandLine.VarP(newInt64SliceValue(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:139
	// _ = "end of CoverTab[120204]"
}

// Int64Slice defines a []int64 flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:142
// The return value is the address of a []int64 variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:144
func (f *FlagSet) Int64Slice(name string, value []int64, usage string) *[]int64 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:144
	_go_fuzz_dep_.CoverTab[120205]++
											p := []int64{}
											f.Int64SliceVarP(&p, name, "", value, usage)
											return &p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:147
	// _ = "end of CoverTab[120205]"
}

// Int64SliceP is like Int64Slice, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Int64SliceP(name, shorthand string, value []int64, usage string) *[]int64 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:151
	_go_fuzz_dep_.CoverTab[120206]++
											p := []int64{}
											f.Int64SliceVarP(&p, name, shorthand, value, usage)
											return &p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:154
	// _ = "end of CoverTab[120206]"
}

// Int64Slice defines a []int64 flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:157
// The return value is the address of a []int64 variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:159
func Int64Slice(name string, value []int64, usage string) *[]int64 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:159
	_go_fuzz_dep_.CoverTab[120207]++
											return CommandLine.Int64SliceP(name, "", value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:160
	// _ = "end of CoverTab[120207]"
}

// Int64SliceP is like Int64Slice, but accepts a shorthand letter that can be used after a single dash.
func Int64SliceP(name, shorthand string, value []int64, usage string) *[]int64 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:164
	_go_fuzz_dep_.CoverTab[120208]++
											return CommandLine.Int64SliceP(name, shorthand, value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:165
	// _ = "end of CoverTab[120208]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:166
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64_slice.go:166
var _ = _go_fuzz_dep_.CoverTab
